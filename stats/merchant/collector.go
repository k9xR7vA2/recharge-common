package merchant

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/stats/base"
	"strconv"
	"time"

	"github.com/k9xR7vA2/recharge-common/model/mongo/stats"
	"github.com/qiniu/qmgo"
	qmgoOpts "github.com/qiniu/qmgo/options"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUpdateOpts = mongooptions.UpdateOptions{Upsert: boolPtr(true)}

func boolPtr(b bool) *bool { return &b }

// FlushInput 归集任务输入
type FlushInput struct {
	TargetHourUTC int64 // 留零则自动取上一个完整小时
	Tenants       []FlushTenant
	BusinessTypes []string
}

// FlushTenant 单个租户的归集元数据
type FlushTenant struct {
	TenantID    uint
	MerchantIDs []uint              // 层级2
	Channels    map[uint][]string   // 层级3：merchantID → []channelCode
	Amounts     map[string][]string // 层级4："merchantID:channelCode" → []amount
	// 层级5：所有通道（跨商户）及其类型
	AllChannelCodes []string       // 层级5 key 枚举
	ChannelTypeMap  map[string]int // channelCode → channelType（从 MySQL 查出后传入）
}

// Flush 将指定 UTC 小时的 Redis 数据 upsert 到 MongoDB（幂等）
func Flush(ctx context.Context, rdb redis.UniversalClient, mongoClient *qmgo.Client, input FlushInput, logger base.Logger) error {
	targetHour := input.TargetHourUTC
	if targetHour == 0 {
		targetHour = prevHourInt(time.Now().UTC())
	}
	hourKey := fmt.Sprintf("%d", targetHour)
	hourTime := parseHourInt(targetHour)
	dateUTC := hourTime.Format("2006-01-02")
	now := time.Now().Unix()

	var totalFlushed int
	for _, tenant := range input.Tenants {
		dbName := fmt.Sprintf("tenant_%d", tenant.TenantID)
		coll := mongoClient.Database(dbName).Collection(stats.CollMerchantOrderHourlyStats)

		for _, bizType := range input.BusinessTypes {
			docs := buildFlushDocs(ctx, rdb, tenant, bizType, hourKey, targetHour, dateUTC, now)
			for _, doc := range docs {
				if err := upsertStat(ctx, coll, doc); err != nil {
					logger.Warn(fmt.Sprintf("stats/merchant collector: upsert failed tenant=%d level=%d err=%v",
						tenant.TenantID, doc.StatLevel, err))
					continue
				}
				totalFlushed++
			}
		}
	}
	logger.Info(fmt.Sprintf("stats/merchant collector: flush complete hour_utc=%d flushed=%d", targetHour, totalFlushed))
	return nil
}

func buildFlushDocs(ctx context.Context, rdb redis.UniversalClient, tenant FlushTenant, bizType, hourKey string, hourUTC int64, dateUTC string, now int64) []stats.MerchantOrderHourlyStat {
	var docs []stats.MerchantOrderHourlyStat

	// 层级1：租户汇总
	if doc, ok := readHashDoc(ctx, rdb, keyL1Tenant(tenant.TenantID, bizType, hourKey),
		newBaseDoc(tenant.TenantID, bizType, stats.MerStatLevelTenant, 0, "", 0, "", hourUTC, dateUTC, now)); ok {
		docs = append(docs, doc)
	}

	// 层级2：商户
	for _, mid := range tenant.MerchantIDs {
		if doc, ok := readHashDoc(ctx, rdb, keyL2Merchant(tenant.TenantID, bizType, mid, hourKey),
			newBaseDoc(tenant.TenantID, bizType, stats.MerStatLevelMerchant, mid, "", 0, "", hourUTC, dateUTC, now)); ok {
			docs = append(docs, doc)
		}

		// 层级3：商户+通道
		for _, cc := range tenant.Channels[mid] {
			ct := tenant.ChannelTypeMap[cc]
			if doc, ok := readHashDoc(ctx, rdb, keyL3Channel(tenant.TenantID, bizType, mid, cc, hourKey),
				newBaseDoc(tenant.TenantID, bizType, stats.MerStatLevelChannel, mid, cc, ct, "", hourUTC, dateUTC, now)); ok {
				docs = append(docs, doc)
			}

			// 层级4：商户+通道+面额
			amtKey := fmt.Sprintf("%d:%s", mid, cc)
			for _, amt := range tenant.Amounts[amtKey] {
				if doc, ok := readHashDoc(ctx, rdb, keyL4Amount(tenant.TenantID, bizType, mid, cc, amt, hourKey),
					newBaseDoc(tenant.TenantID, bizType, stats.MerStatLevelAmount, mid, cc, ct, amt, hourUTC, dateUTC, now)); ok {
					docs = append(docs, doc)
				}
			}
		}
	}

	// 层级5：仅通道（跨商户）
	for _, cc := range tenant.AllChannelCodes {
		ct := tenant.ChannelTypeMap[cc]
		if doc, ok := readHashDoc(ctx, rdb, keyL5ChannelOnly(tenant.TenantID, bizType, cc, hourKey),
			newBaseDoc(tenant.TenantID, bizType, stats.MerStatLevelChannelOnly, 0, cc, ct, "", hourUTC, dateUTC, now)); ok {
			docs = append(docs, doc)
		}
	}

	return docs
}

func readHashDoc(ctx context.Context, rdb redis.UniversalClient, key string, baseStats stats.MerchantOrderHourlyStat) (stats.MerchantOrderHourlyStat, bool) {
	vals, err := rdb.HGetAll(ctx, key).Result()
	if err != nil || len(vals) == 0 {
		return stats.MerchantOrderHourlyStat{}, false
	}
	baseStats.TotalOrders = base.ParseInt(vals[fieldTotal])
	baseStats.SuccessOrders = base.ParseInt(vals[fieldSuccess])
	baseStats.SuccessAmount = base.ParseInt(vals[fieldSuccessAmount])
	baseStats.TotalAmount = base.ParseInt(vals[fieldTotalAmount])
	// channel_type 优先从 Redis 取（写入时已 HSET），base 里的作为兜底
	if ct, ok := vals[fieldChannelType]; ok && ct != "" {
		if v, err := strconv.Atoi(ct); err == nil && v > 0 {
			baseStats.ChannelType = v
		}
	}
	return baseStats, true
}

func upsertStat(ctx context.Context, coll *qmgo.Collection, doc stats.MerchantOrderHourlyStat) error {
	filter := bson.M{
		"tenant_id":     doc.TenantID,
		"business_type": doc.BusinessType,
		"stat_level":    doc.StatLevel,
		"merchant_id":   doc.MerchantID,
		"channel_code":  doc.ChannelCode,
		"amount":        doc.Amount,
		"hour_utc":      doc.HourUTC,
	}
	update := bson.M{
		"$set": bson.M{
			"total_orders":   doc.TotalOrders,
			"success_orders": doc.SuccessOrders,
			"success_amount": doc.SuccessAmount,
			"total_amount":   doc.TotalAmount,
			"channel_type":   doc.ChannelType,
			"date_utc":       doc.DateUTC,
			"flush_at":       doc.FlushAt,
			"updated_at":     doc.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"created_at": doc.CreatedAt,
		},
	}
	return coll.UpdateOne(ctx, filter, update, qmgoOpts.UpdateOptions{UpdateOptions: &mongoUpdateOpts})
}

func newBaseDoc(tenantID uint, bizType string, level stats.MerStatLevel, merchantID uint, channelCode string, channelType int, amount string, hourUTC int64, dateUTC string, now int64) stats.MerchantOrderHourlyStat {
	return stats.MerchantOrderHourlyStat{
		TenantID:     tenantID,
		BusinessType: bizType,
		StatLevel:    level,
		MerchantID:   merchantID,
		ChannelCode:  channelCode,
		ChannelType:  channelType,
		Amount:       amount,
		HourUTC:      hourUTC,
		DateUTC:      dateUTC,
		FlushAt:      now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func prevHourInt(t time.Time) int64 {
	var result int64
	fmt.Sscanf(t.Add(-time.Hour).UTC().Format("2006010215"), "%d", &result)
	return result
}

func parseHourInt(h int64) time.Time {
	t, _ := time.Parse("2006010215", fmt.Sprintf("%d", h))
	return t
}
