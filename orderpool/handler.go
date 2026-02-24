package orderpool

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/orderpool/keys"
	"github.com/small-cat1/recharge-common/orderpool/options"
	"strconv"
	"time"
)

type NonCoreOpParams struct {
	StatsKey      string
	EventsKey     string
	Amount        string
	Carrier       string
	ChargeSpeed   string
	Region        string
	Province      string
	Priority      string
	SystemOrderSn string
}

type MobileOrderPool struct {
	redisClient redis.UniversalClient
}

// NewMobileOrderPool 话费订单池
func NewMobileOrderPool(redisClient redis.UniversalClient) *MobileOrderPool {
	return &MobileOrderPool{
		redisClient: redisClient,
	}
}

// AddOrderToPool 订单入池，这里会根据池子的key进入不同的订单池优先级高和优先级低的池子,订单再次入池也是这个逻辑
func (m *MobileOrderPool) AddOrderToPool(ctx context.Context, opts options.IMobileHandlerOptions) error {
	// === 调试开始 ===
	fmt.Printf("[DEBUG AddOrderToPool] PoolKey: %s\n", opts.GetPoolKey())
	fmt.Printf("[DEBUG AddOrderToPool] OrderKey: %s\n", opts.GetOrderKey())
	fmt.Printf("[DEBUG AddOrderToPool] SystemOrderSn: %s\n", opts.GetSystemOrderSnArg())
	fmt.Printf("[DEBUG AddOrderToPool] ValidTime: %d\n", opts.GetValidTimeArg())
	fmt.Printf("[DEBUG AddOrderToPool] ExpiredAt: %d\n", opts.GetExpiredAtArg())
	fmt.Printf("[DEBUG AddOrderToPool] Priority: %s\n", opts.GetPriority())
	// === 调试结束 ===
	// 1. 准备Lua脚本和参数（确保核心操作原子性）
	script := `
		-- 修改Lua脚本，先检查，再写入
		local exists = redis.call('EXISTS', KEYS[2])
		if exists == 1 then
			return {0, "ORDER_EXISTS"}
		end

	  	-- 添加订单到流
		local msgId = redis.call('XADD', KEYS[1], '*', 'order_sn', ARGV[1], 'expire_at', ARGV[2], 'retry_count', 0, 'create_time', ARGV[3])
		if not msgId then
			return {0, "ADD_FAILED"}
		end

		-- 设置订单信息
		redis.call('HSET', KEYS[2], 
		  'msg_id', msgId, 
		  'pool_key', KEYS[1], 
		  'order_info', ARGV[4], 
		  'create_time', ARGV[3],
		  'priority', ARGV[6])
		redis.call('EXPIRE', KEYS[2], ARGV[5])

		return {1, msgId}
    `
	// 2. 执行核心原子操作
	validTime := opts.GetValidTimeArg() + constant.ExtraRetentionSeconds
	scriptKeys := []string{
		opts.GetPoolKey(),  // KEYS[1] - 订单池键
		opts.GetOrderKey(), // KEYS[2] - 订单信息键
	}
	now := time.Now().UnixMilli()
	args := []interface{}{
		opts.GetSystemOrderSnArg(), // ARGV[1] - 系统订单号
		opts.GetExpiredAtArg(),     // ARGV[2] - 过期时间
		now,                        // ARGV[3] - 当前时间
		opts.GetOrderDataArg(),     // ARGV[4] - 订单数据
		validTime,                  // ARGV[5] - 有效时间
		opts.GetPriority(),         // ARGV[6] - 优先级
	}
	// 3. 执行Lua脚本并处理返回值
	result, err := m.redisClient.Eval(ctx, script, scriptKeys, args).Result()
	if err != nil {

		return fmt.Errorf("订单入池核心操作失败: %w", err)
	}
	// 4. 解析结果
	resultArray, ok := result.([]interface{})
	if !ok || len(resultArray) < 2 {
		return fmt.Errorf("订单入池返回结果格式错误")
	}
	// 5. 检查操作状态
	status, _ := resultArray[0].(int64)
	message, _ := resultArray[1].(string)

	if status == 0 {
		if message == "ORDER_EXISTS" {
			return fmt.Errorf("订单信息已存在")
		}
		return fmt.Errorf("订单入池失败: %s", message)
	}
	// 6. 执行非核心操作,记录统计和事件
	eventDetails := map[string]interface{}{
		"pool_args": opts.GetPoolArgs(),
		"expire_at": opts.GetExpiredAtArg(),
		"priority":  opts.GetPriority(),
		"msg_id":    message,
	}
	poolArgs := opts.GetPoolArgs()
	params := NonCoreOpParams{
		StatsKey:      opts.GetStatsKey(),
		EventsKey:     opts.GetEventsKey(),
		Amount:        poolArgs.Amount,
		Carrier:       poolArgs.Carrier,
		ChargeSpeed:   poolArgs.ChargeSpeed,
		Region:        poolArgs.Region,
		Province:      poolArgs.Province,
		Priority:      opts.GetPriority(),
		SystemOrderSn: opts.GetSystemOrderSnArg(),
	}
	return m.executeNonCoreOperations(ctx, params, EventTypeCreated, eventDetails, 1)
}

// FetchOrder 取出订单 - 使用Lua脚本实现核心逻辑
func (m *MobileOrderPool) FetchOrder(ctx context.Context, opts options.IFetchMobileOrderOptions, workerID string) (string, string, error) {
	fetchScript := `
    local poolCount = (#KEYS - 2) / 2
    local tenantID = KEYS[#KEYS - 1]
    local role = KEYS[#KEYS]
    local groupName = "processors"

    -- 收集各优先级过期订单数量
    local expiredCounts = {}
    local validOrder = nil

    for i = 1, poolCount do
        local poolKey = KEYS[i]
        local priority = KEYS[i + poolCount]
        expiredCounts[priority] = 0

        pcall(redis.call, 'XGROUP', 'CREATE', poolKey, groupName, '0', 'MKSTREAM')

        local maxAttempts = 100
        for attempt = 1, maxAttempts do
            if validOrder then
                break
            end

            local messages = redis.call('XREADGROUP', 'GROUP', groupName, ARGV[1], 'COUNT', 1, 'STREAMS', poolKey, '>')

            if not messages or #messages == 0 or #messages[1][2] == 0 then
                break
            end

            local msg = messages[1][2][1]
            local msgId = msg[1]
            local values = {}

            for j = 1, #msg[2], 2 do
                values[msg[2][j]] = msg[2][j+1]
            end

            local expireAt = tonumber(values["expire_at"] or 0)
            local now = tonumber(ARGV[2])

            if now > expireAt then
                redis.call('XACK', poolKey, groupName, msgId)
                redis.call('XDEL', poolKey, msgId)
                local order_sn = values["order_sn"]
                if order_sn then
                    local orderKey = "tenant:" .. tenantID .. ":" .. role .. ":" .. ARGV[3] .. ":order:" .. order_sn
                    redis.call('DEL', orderKey)
                end
                expiredCounts[priority] = expiredCounts[priority] + 1
            else
                local order_sn = values["order_sn"]
                local orderKey = "tenant:" .. tenantID .. ":" .. role .. ":" .. ARGV[3] .. ":order:" .. order_sn
                local orderInfo = redis.call('HGETALL', orderKey)

                if #orderInfo == 0 then
                    redis.call('XACK', poolKey, groupName, msgId)
                    redis.call('XDEL', poolKey, msgId)
                else
                    redis.call('XACK', poolKey, groupName, msgId)
                    redis.call('XDEL', poolKey, msgId)

                    local orderData = {}
                    for j = 1, #orderInfo, 2 do
                        orderData[orderInfo[j]] = orderInfo[j+1]
                    end

                    redis.call('HSET', orderKey, 'status', 'processing', 'fetch_time', ARGV[2], 'worker_id', ARGV[1])

                    validOrder = {orderData["order_info"], priority}
                end
            end
        end
    end

    -- 统一返回：order_info, priority, high过期数, normal过期数
    local highPriority = KEYS[1 + poolCount]
    local normalPriority = KEYS[2 + poolCount]
    local highExpired = expiredCounts[highPriority] or 0
    local normalExpired = expiredCounts[normalPriority] or 0

    if validOrder then
        return {validOrder[1], validOrder[2], highExpired, normalExpired}
    else
        return {"", "", highExpired, normalExpired}
    end
`
	tenantIdStr := strconv.Itoa(int(opts.GetTenantId()))
	KeyGenerator := keys.NewRedisKeysGenerate(opts.GetTenantId(), keys.RoleSupplier, opts.GetBusinessType())
	highPriorityPoolKey := KeyGenerator.GenerateMobilePoolKey(constant.HighPriority.String(), opts.GetPoolArgs())
	normalPriorityPoolKey := KeyGenerator.GenerateMobilePoolKey(constant.NormalPriority.String(), opts.GetPoolArgs())

	scriptKeys := []string{
		highPriorityPoolKey,              // KEYS[1]
		normalPriorityPoolKey,            // KEYS[2]
		constant.HighPriority.String(),   // KEYS[3]
		constant.NormalPriority.String(), // KEYS[4]
		tenantIdStr,                      // KEYS[5]
		opts.GetRoleType(),               // KEYS[6]
	}

	args := []interface{}{
		workerID,               // ARGV[1]
		time.Now().UnixMilli(), // ARGV[2]
		opts.GetBusinessType(), // ARGV[3]
	}

	// === 调试开始 ===
	fmt.Printf("[DEBUG FetchOrder] scriptKeys: %v\n", scriptKeys)
	fmt.Printf("[DEBUG FetchOrder] args: %v\n", args)

	// 查 Stream 长度
	highLen, _ := m.redisClient.XLen(ctx, highPriorityPoolKey).Result()
	normalLen, _ := m.redisClient.XLen(ctx, normalPriorityPoolKey).Result()
	fmt.Printf("[DEBUG FetchOrder] highPool XLEN: %d, normalPool XLEN: %d\n", highLen, normalLen)

	// 查订单 Hash 是否存在（用最新那条消息的 order_sn）
	normalMsgs, _ := m.redisClient.XRange(ctx, normalPriorityPoolKey, "-", "+").Result()
	for _, msg := range normalMsgs {
		orderSn := msg.Values["order_sn"]
		orderKey := fmt.Sprintf("tenant:%s:%s:%s:order:%s", tenantIdStr, opts.GetRoleType(), opts.GetBusinessType(), orderSn)
		exists, _ := m.redisClient.Exists(ctx, orderKey).Result()
		orderInfo, _ := m.redisClient.HGetAll(ctx, orderKey).Result()
		fmt.Printf("[DEBUG FetchOrder] orderKey: %s, exists: %d, info: %v\n", orderKey, exists, orderInfo)
	}
	result, err := m.redisClient.Eval(ctx, fetchScript, scriptKeys, args).Result()
	fmt.Printf("[DEBUG FetchOrder] Eval result: %+v, err: %v\n", result, err)
	if err != nil {
		if err == redis.Nil {
			return "", "", fmt.Errorf("no orders available")
		}
		return "", "", fmt.Errorf("获取订单失败: %w", err)
	}

	if result == nil {
		return "", "", fmt.Errorf("no orders available")
	}

	resultArr, ok := result.([]interface{})
	if !ok || len(resultArr) < 4 {
		return "", "", fmt.Errorf("获取订单失败：未知响应格式")
	}

	highExpired, _ := resultArr[2].(int64)
	normalExpired, _ := resultArr[3].(int64)
	poolArgs := opts.GetPoolArgs()

	// 先处理过期订单的统计递减
	if highExpired > 0 {
		params := NonCoreOpParams{
			StatsKey:      opts.GetStatsKey(),
			EventsKey:     opts.GetEventsKey("expired_batch"),
			Amount:        poolArgs.Amount,
			Carrier:       poolArgs.Carrier,
			ChargeSpeed:   poolArgs.ChargeSpeed,
			Region:        poolArgs.Region,
			Province:      poolArgs.Province,
			Priority:      constant.HighPriority.String(),
			SystemOrderSn: "expired_batch",
		}
		eventDetails := map[string]interface{}{
			"type":          "expired_cleanup",
			"expired_count": highExpired,
		}
		_ = m.executeNonCoreOperations(ctx, params, EventTypeExpired, eventDetails, -highExpired)
	}

	if normalExpired > 0 {
		params := NonCoreOpParams{
			StatsKey:      opts.GetStatsKey(),
			EventsKey:     opts.GetEventsKey("expired_batch"),
			Amount:        poolArgs.Amount,
			Carrier:       poolArgs.Carrier,
			ChargeSpeed:   poolArgs.ChargeSpeed,
			Region:        poolArgs.Region,
			Province:      poolArgs.Province,
			Priority:      constant.NormalPriority.String(),
			SystemOrderSn: "expired_batch",
		}
		eventDetails := map[string]interface{}{
			"type":          "expired_cleanup",
			"expired_count": normalExpired,
		}
		_ = m.executeNonCoreOperations(ctx, params, EventTypeExpired, eventDetails, -normalExpired)
	}

	// 再处理正常取出的订单
	orderInfoStr, _ := resultArr[0].(string)
	priority, _ := resultArr[1].(string)
	if orderInfoStr == "" {
		return "", "", fmt.Errorf("no orders available")
	}

	var tempData map[string]interface{}
	if err := json.Unmarshal([]byte(orderInfoStr), &tempData); err != nil {
		return "", "", fmt.Errorf("订单数据解析失败: %w", err)
	}
	systemOrderSn, ok := tempData["system_order_sn"].(string)
	if !ok || systemOrderSn == "" {
		return "", "", fmt.Errorf("订单数据中缺少 system_order_sn")
	}

	eventDetails := map[string]interface{}{
		"worker_id": workerID,
		"priority":  priority,
	}
	eventKey := opts.GetEventsKey(systemOrderSn)
	params := NonCoreOpParams{
		StatsKey:      opts.GetStatsKey(),
		EventsKey:     eventKey,
		Amount:        poolArgs.Amount,
		Carrier:       poolArgs.Carrier,
		ChargeSpeed:   poolArgs.ChargeSpeed,
		Region:        poolArgs.Region,
		Province:      poolArgs.Province,
		Priority:      priority,
		SystemOrderSn: systemOrderSn,
	}
	err = m.executeNonCoreOperations(ctx, params, EventTypeProcessing, eventDetails, -1)
	if err != nil {
		return "", "", err
	}

	return orderInfoStr, priority, nil
}

// CancelOrRemoveOrder  撤销/删除订单
func (m *MobileOrderPool) CancelOrRemoveOrder(ctx context.Context, opts options.IMobileHandlerOptions, event EventType) error {
	cancelScript := `
		local orderKey = KEYS[1]
		-- 获取订单信息
		local orderInfo = redis.call('HGETALL', orderKey)
		if #orderInfo == 0 then
			return {"error", "订单信息不存在"}
		end

		-- 将订单信息转换为table
		local orderData = {}
		for i = 1, #orderInfo, 2 do
			orderData[orderInfo[i]] = orderInfo[i+1]
		end

		-- 如果订单已经被消费，不允许取消
		local status = orderData["status"] or ""
		if status == "processing" then
			return {"error", "订单已被领取，无法取消"}
		end

		local poolKey = orderData["pool_key"]
		local messageID = orderData["msg_id"]
		local priority = orderData["priority"] or "normal"

		-- 从池中删除订单，并检查实际删除数量
		local deleted = 0
		if poolKey and messageID then
			-- 先 XACK 再 XDEL，清理 PEL
			pcall(redis.call, 'XACK', poolKey, 'processors', messageID)
			deleted = redis.call('XDEL', poolKey, messageID)
		end

		-- 删除订单 Hash 信息
		redis.call('DEL', orderKey)

		return {"success", poolKey, messageID, priority, tostring(deleted)}
	`

	scriptKeys := []string{
		opts.GetOrderKey(), // KEYS[1] - 订单信息键
	}

	result, err := m.redisClient.Eval(ctx, cancelScript, scriptKeys, []interface{}{}).Result()
	if err != nil {
		return fmt.Errorf("删除订单失败: %w", err)
	}

	resultArr, ok := result.([]interface{})
	if !ok || len(resultArr) == 0 {
		return fmt.Errorf("lua脚本删除订单失败：未知错误")
	}

	status, _ := resultArr[0].(string)
	if status == "error" && len(resultArr) > 1 {
		errMsg, _ := resultArr[1].(string)
		return fmt.Errorf("%s", errMsg)
	}

	if status == "success" && len(resultArr) >= 5 {
		poolKey, _ := resultArr[1].(string)
		msgID, _ := resultArr[2].(string)
		priority, _ := resultArr[3].(string)
		// deleted, _ := resultArr[4].(string) // 可用于日志记录
		eventDetails := map[string]interface{}{
			"by_request": true,
			"pool_key":   poolKey,
			"msg_id":     msgID,
		}
		poolArgs := opts.GetPoolArgs()
		params := NonCoreOpParams{
			StatsKey:      opts.GetStatsKey(),
			EventsKey:     opts.GetEventsKey(),
			Amount:        poolArgs.Amount,
			Carrier:       poolArgs.Carrier,
			ChargeSpeed:   poolArgs.ChargeSpeed,
			Region:        poolArgs.Region,
			Province:      poolArgs.Province,
			Priority:      priority,
			SystemOrderSn: opts.GetSystemOrderSnArg(),
		}
		return m.executeNonCoreOperations(ctx, params, event, eventDetails, -1)
	}

	return fmt.Errorf("删除订单失败：未知响应格式")
}

// GetOrderInfo 获取订单信息
func (m *MobileOrderPool) GetOrderInfo(ctx context.Context, orderKey string) (map[string]interface{}, error) {
	result, err := m.redisClient.HGetAll(ctx, orderKey).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("订单信息不存在")
	}
	orderInfo := make(map[string]interface{})
	for k, v := range result {
		orderInfo[k] = v
	}
	return orderInfo, nil
}

// 执行非核心操作（统计更新和事件记录）
func (m *MobileOrderPool) executeNonCoreOperations(ctx context.Context,
	params NonCoreOpParams,
	eventType EventType,
	eventDetails interface{},
	incr int64) error {

	// 准备Lua脚本
	nonCoreLua := `
        -- 更新统计数据
        local statsKey = KEYS[1]
        local incrValue = tonumber(ARGV[11])  -- 新增参数，指定增减值

        redis.call('HINCRBY', statsKey, 'amount:'..ARGV[1], incrValue)
        redis.call('HINCRBY', statsKey, 'carrier:'..ARGV[2], incrValue)
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3], incrValue)
        redis.call('HINCRBY', statsKey, 'area:'..ARGV[4], incrValue)
        redis.call('HINCRBY', statsKey, 'province:'..ARGV[5], incrValue)
        redis.call('HINCRBY', statsKey, 'pool_orders', incrValue)
        redis.call('HINCRBY', statsKey, 'priority:'..ARGV[6], incrValue)
         -- 交叉维度统计 - 充值速度 x 运营商
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':carrier:'..ARGV[2], incrValue)
        -- 交叉维度统计 - 充值速度 x 金额
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':amount:'..ARGV[1], incrValue)
        -- 交叉维度统计 - 运营商 x 金额
        redis.call('HINCRBY', statsKey, 'carrier:'..ARGV[2]..':amount:'..ARGV[1], incrValue)
        -- 三维交叉统计 - 充值速度 x 运营商 x 金额（用于前端交叉表）
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':carrier:'..ARGV[2]..':amount:'..ARGV[1], incrValue)

        -- 新增：处理中订单数
		local eventType = ARGV[8]
		if eventType == 'processing' then
			-- 取出订单：pool -1 已经做了，processing +1
			redis.call('HINCRBY', statsKey, 'processing_orders', 1)
		elseif eventType == 'completed' then
			-- 订单结束：processing -1
			redis.call('HINCRBY', statsKey, 'processing_orders', -1)
		end

        -- 记录事件
        local eventsKey = KEYS[2]
        redis.call('XADD', eventsKey, '*', 
            'order_sn', ARGV[7], 
            'event', ARGV[8], 
            'timestamp', ARGV[9], 
            'details', ARGV[10]
        )
        
        -- 设置事件流过期时间
        redis.call('EXPIRE', eventsKey, 21600) -- 6小时
        
        return 'OK'
    `
	// 序列化事件详情
	var detailsStr string
	switch v := eventDetails.(type) {
	case string:
		detailsStr = v
	case map[string]interface{}:
		data, _ := json.Marshal(v)
		detailsStr = string(data)
	default:
		data, _ := json.Marshal(v)
		detailsStr = string(data)
	}
	now := time.Now().UnixMilli()
	// 准备键和参数
	scriptKeys := []string{
		params.StatsKey,
		params.EventsKey,
	}
	args := []interface{}{
		params.Amount,        // ARGV[1] - 金额
		params.Carrier,       // ARGV[2] - 运营商
		params.ChargeSpeed,   // ARGV[3] - 充值类型
		params.Region,        // ARGV[4] - 区域
		params.Province,      // ARGV[5] - 省份
		params.Priority,      // ARGV[6] - 优先级
		params.SystemOrderSn, // ARGV[7] - 系统订单号
		eventType.String(),   // ARGV[8] - 事件类型
		now,                  // ARGV[9] - 当前时间戳
		detailsStr,           // ARGV[10] - 事件详情
		incr,                 // ARGV[11] - 统计值增减
	}

	// 执行Lua脚本
	_, err := m.redisClient.Eval(ctx, nonCoreLua, scriptKeys, args).Result()
	if err != nil {
		return fmt.Errorf("执行更新统计数据和记录事件lua脚本失败,error:%w", err)
	}

	return nil
}
