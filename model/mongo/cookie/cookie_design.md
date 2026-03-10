# Cookie 系统设计文档

---

## 一、各表协作关系

### 表职责总览

```
┌─────────────────────────────────────────────────────────────────┐
│                        MySQL（配置 & 统计）                        │
│   PlatformRule          CookieBanRecord       CookieDailyStats  │
│   平台规则配置             封控历史事件              运营日报快照        │
└────────────┬────────────────────┬──────────────────┬────────────┘
             │ 读取规则              │ 写入封控事件         │ 定时聚合
┌────────────▼────────────────────▼──────────────────▼────────────┐
│                       MongoDB（主数据）                            │
│   AccountCookie         AccountStats          CookieUsageLog    │
│   Cookie主体+风控状态      日/周期计数统计           操作流水日志        │
└────────────┬────────────────────┬──────────────────────────────┘
             │ 初始化/重建           │ lastUsedAt 同步
┌────────────▼──────────────────────────────────────────────────┐
│                       Redis（调度运行时）                          │
│   pool:{tenant}:{platform}        pool:cooldown / pool:suspend │
│   可用池（ZSET）                    冷却池 / 封控池（ZSET）          │
│   pool:inuse（SET）                                             │
└───────────────────────────────────────────────────────────────┘
```

---

### 完整业务流程中的表协作

#### 1. 取 Cookie（调度）

```
请求进来
  │
  ├─ 读 PlatformRule          → 获取 minInterval / dailyMax / monthlyMax
  │
  ├─ ZPOPMIN pool:{t}:{p}     → Redis 原子取出 cookieId（防并发）
  │   └─ 池子为空？→ 返回无可用Cookie
  │
  ├─ SADD pool:inuse           → 标记使用中
  │
  └─ MongoDB AccountCookie     → 用 cookieId 取完整数据（cookie/ua/proxy）
```

#### 2. 使用结束回写

```
执行结果
  │
  ├─ 成功
  │   ├─ AccountCookie        → 更新 lastUsedAt
  │   ├─ AccountStats         → DailyUsed++ / PeriodUsed++
  │   ├─ CookieUsageLog       → 写一条 action=release result=1
  │   ├─ Redis SREM inuse     → 移出使用中
  │   └─ Redis ZADD pool      → score = now + minIntervalSec（更新可用时间）
  │
  └─ 失败
      ├─ AccountCookie        → FailCount++，更新 HealthScore
      ├─ CookieUsageLog       → 写一条 action=fail result=2
      ├─ Redis SREM inuse     → 移出使用中
      │
      └─ FailCount >= MaxFailCount?
          ├─ 是（进冷却）
          │   ├─ AccountCookie  → Status=Cooldown，CooldownUntil=now+CooldownSec，FailCount=0
          │   ├─ CookieBanRecord → 写一条封控历史事件
          │   ├─ Redis ZREM pool → 移出可用池
          │   └─ Redis ZADD pool:cooldown → score=CooldownUntil
          │
          └─ 否（留在池子）
              └─ Redis ZADD pool → 更新 score 延迟重试
```

#### 3. 冷却到期（定时任务，每分钟）

```
扫描 Redis pool:cooldown
  │
  ZRANGEBYSCORE 0 now
  │
  ├─ 有到期的 cookieId
  │   ├─ AccountCookie   → Status=Normal / Probing，FailCount=0
  │   ├─ Redis ZREM pool:cooldown
  │   │
  │   └─ ProbeEnabled = true?
  │       ├─ 是 → ZADD pool:probing，等待探测任务处理
  │       └─ 否 → 直接 ZADD pool 恢复可用
  │
  └─ 无到期的 → 跳过
```

#### 4. 探测（定时任务，每 ProbeIntervalSec）

```
扫描 pool:probing
  │
  ├─ 发起探测请求
  │
  ├─ 探测成功
  │   ├─ AccountCookie    → Status=Normal，FailCount=0，NextProbeAt=nil
  │   ├─ CookieBanRecord  → ResolvedAt=now
  │   ├─ CookieUsageLog   → action=probe result=1
  │   ├─ Redis ZREM pool:probing
  │   └─ Redis ZADD pool  → 恢复可用
  │
  └─ 探测失败
      ├─ AccountCookie    → Status=Suspend，SuspendUntil=now+SuspendSec，NextProbeAt=now+ProbeIntervalSec
      ├─ CookieUsageLog   → action=probe result=2
      ├─ Redis ZREM pool:probing
      └─ Redis ZADD pool:suspend → score=SuspendUntil
```

#### 5. 每日运营快照（定时任务，每天凌晨）

```
遍历所有租户+平台
  │
  ├─ AccountStats  → 聚合当日 DailyUsed
  ├─ AccountCookie → 统计 Status=Normal 的数量（Available）
  ├─ CookieBanRecord → 统计当日新增封控数（Banned）
  └─ 写入 CookieDailyStats
```

---

## 二、Redis Cookie Pool 设计

### Key 结构

```
# 可用池（有序集合）
pool:available:{tenantId}:{platform}
  score = 下次可用时间戳（lastUsedAt + minIntervalSec）
  value = cookieId

# 冷却池（有序集合）
pool:cooldown:{tenantId}:{platform}
  score = cooldownUntil 时间戳
  value = cookieId

# 封控池（有序集合）
pool:suspend:{tenantId}:{platform}
  score = suspendUntil 时间戳
  value = cookieId

# 探测池（有序集合）
pool:probing:{tenantId}:{platform}
  score = nextProbeAt 时间戳
  value = cookieId

# 使用中（集合，防并发重复取）
pool:inuse:{tenantId}:{platform}
  value = cookieId

# 池子元数据（Hash，存基础信息避免回查MongoDB）
pool:meta:{cookieId}
  field: last_used_at
  field: fail_count
  field: health_score
  TTL: 24小时（定期从MongoDB刷新）
```

---

### 核心操作

#### 取 Cookie
```
ZRANGEBYSCORE pool:available:{t}:{p} 0 {now} LIMIT 0 1
  → 取 score <= now 的第一个（即当前可用的）

if 取到:
    ZREM  pool:available:{t}:{p} cookieId   # 移出可用池
    SADD  pool:inuse:{t}:{p}    cookieId   # 加入使用中
    return cookieId

if 未取到:
    return nil  # 无可用Cookie
```

#### 释放 Cookie（成功）
```
SREM pool:inuse:{t}:{p} cookieId
ZADD pool:available:{t}:{p} {now + minIntervalSec} cookieId
```

#### 释放 Cookie（失败，未达阈值）
```
SREM pool:inuse:{t}:{p} cookieId
ZADD pool:available:{t}:{p} {now + minIntervalSec * 2} cookieId  # 适当延长间隔
```

#### 进入冷却
```
SREM pool:inuse:{t}:{p}    cookieId
ZREM pool:available:{t}:{p} cookieId
ZADD pool:cooldown:{t}:{p} {cooldownUntil} cookieId
```

#### 进入封控
```
ZREM pool:cooldown:{t}:{p} cookieId
ZADD pool:suspend:{t}:{p}  {suspendUntil} cookieId
```

---

### 池子初始化 / 重建

服务启动时或 Redis 重启后，从 MongoDB 重建：

```go
func RebuildPool(tenantId, platform string) {
    // 1. 查 MongoDB 所有有效 Cookie
    cookies := mongo.Find({
        status:     {$in: [Normal, Cooldown, Suspend]},
        platform:   platform,
        is_deleted: false,
    })

    for _, c := range cookies {
        switch c.Status {
        case Normal:
            score := max(c.LastUsedAt + minIntervalSec, now)
            ZADD pool:available score c.ID

        case Cooldown:
            ZADD pool:cooldown c.CooldownUntil c.ID

        case Suspend:
            if ProbeEnabled && c.NextProbeAt != nil:
                ZADD pool:probing c.NextProbeAt c.ID
            else:
                ZADD pool:suspend c.SuspendUntil c.ID
        }

        // 写 meta
        HSET pool:meta:{c.ID}
            last_used_at c.LastUsedAt
            fail_count   c.FailCount
            health_score c.HealthScore
        EXPIRE pool:meta:{c.ID} 86400
    }
}
```

---

### 各池子数量监控

```
# 可用数
ZCOUNT pool:available:{t}:{p} 0 +inf

# 当前使用中
SCARD pool:inuse:{t}:{p}

# 冷却中
ZCOUNT pool:cooldown:{t}:{p} 0 +inf

# 封控中
ZCOUNT pool:suspend:{t}:{p} 0 +inf

# 探测中
ZCOUNT pool:probing:{t}:{p} 0 +inf
```

这几个指标每分钟采集一次，配合 `CookieDailyStats` 做运营大盘。

---

### 状态流转总览

```
                    导入
                     │
                     ▼
              ┌─── Normal ───┐
              │   可用池       │
              │  ZSET score  │  失败次数 >= MaxFailCount
              └──────────────┘ ──────────────────────────►  Cooldown
                     ▲                                       冷却池
                     │ 探测成功                               ZSET score=cooldownUntil
                     │                                            │
                  Probing  ◄── 冷却到期 + ProbeEnabled=true ──────┘
                  探测池                                           │
                     │                                    ProbeEnabled=false
                     │ 探测失败                                     │
                     ▼                                            ▼
                  Suspend ◄───────────────────────────────────  Normal
                  封控池                                         直接恢复
              ZSET score=suspendUntil
                     │
                     │ 封控到期
                     ▼
                  Probing（再次探测）
```
