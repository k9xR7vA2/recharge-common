package supplier

import (
	"context"
	"fmt"
	"time"

	"github.com/k9xR7vA2/recharge-common/model/mongo/stats"
	"github.com/qiniu/qmgo"
	qmgoOpts "github.com/qiniu/qmgo/options"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

// FlushInput 归集任务输入
type FlushInput struct {
	// 要归集的 UTC 小时（int64，如 2026030408）；留零则自动取上一个完整小时
	TargetHourUTC int64

	// 活跃租户列表，调用方从 DB 查出后传入，避免扫全量 Redis
	Tenants []FlushTenant

	// 业务类型列表，如 ["mobile", "india_mobile"]
	BusinessTypes []string
}

// FlushTenant 单个租户的归集元数据
type FlushTenant struct {
	TenantID     uint
	SupplierIDs  []uint              // 层级2
	Products     map[uint][]string   // 层级3：supplierID → []productCode
	Amounts      map[string][]string // 层级4："supplierID:productCode" → []amount
	ProductCodes []string            // 层级5：所有产品编码（跨供货商）
}

// Flush 将指定 UTC 小时的 Redis 数据 upsert 到 MongoDB
// 幂等：相同小时多次调用结果一致（$set 覆盖）
func Flush(ctx context.Context, rdb redis.UniversalClient, mongoClient *qmgo.Client, input FlushInput, logger Logger) error {
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
		coll := mongoClient.Database(dbName).Collection(stats.CollSupplierOrderHourlyStats)

		for _, bizType := range input.BusinessTypes {
			docs, err := buildFlushDocs(ctx, rdb, tenant, bizType, hourKey, targetHour, dateUTC, now)
			if err != nil {
				logger.Warn(fmt.Sprintf("stats/collector: build docs failed tenant=%d biz=%s err=%v",
					tenant.TenantID, bizType, err))
				continue
			}
			for _, doc := range docs {
				if err = upsertStat(ctx, coll, doc); err != nil {
					logger.Warn(fmt.Sprintf("stats/collector: upsert failed tenant=%d level=%d err=%v",
						tenant.TenantID, doc.StatLevel, err))
					continue
				}
				totalFlushed++
			}
		}
	}

	logger.Info(fmt.Sprintf("stats/collector: flush complete hour_utc=%d flushed=%d", targetHour, totalFlushed))
	return nil
}

// buildFlushDocs 从 Redis 读取指定小时的数据，构建全部层级的待写文档
func buildFlushDocs(
	ctx context.Context,
	rdb redis.UniversalClient,
	tenant FlushTenant,
	bizType, hourKey string,
	hourUTC int64,
	dateUTC string,
	now int64,
) ([]stats.SupplierOrderHourlyStat, error) {
	var docs []stats.SupplierOrderHourlyStat

	// 层级1：租户汇总
	if doc, ok := readHashDoc(ctx, rdb, keyL1Tenant(tenant.TenantID, bizType, hourKey),
		newBaseDoc(tenant.TenantID, bizType, stats.StatLevelTenant, 0, "", "", hourUTC, dateUTC, now)); ok {
		docs = append(docs, doc)
	}

	// 层级2：供货商
	for _, sid := range tenant.SupplierIDs {
		if doc, ok := readHashDoc(ctx, rdb, keyL2Supplier(tenant.TenantID, bizType, sid, hourKey),
			newBaseDoc(tenant.TenantID, bizType, stats.StatLevelSupplier, sid, "", "", hourUTC, dateUTC, now)); ok {
			docs = append(docs, doc)
		}

		// 层级3：供货商+产品
		for _, pc := range tenant.Products[sid] {
			if doc, ok := readHashDoc(ctx, rdb, keyL3Product(tenant.TenantID, bizType, sid, pc, hourKey),
				newBaseDoc(tenant.TenantID, bizType, stats.StatLevelProduct, sid, pc, "", hourUTC, dateUTC, now)); ok {
				docs = append(docs, doc)
			}

			// 层级4：供货商+产品+面额
			amtKey := fmt.Sprintf("%d:%s", sid, pc)
			for _, amt := range tenant.Amounts[amtKey] {
				if doc, ok := readHashDoc(ctx, rdb, keyL4Amount(tenant.TenantID, bizType, sid, pc, amt, hourKey),
					newBaseDoc(tenant.TenantID, bizType, stats.StatLevelAmount, sid, pc, amt, hourUTC, dateUTC, now)); ok {
					docs = append(docs, doc)
				}
			}
		}
	}

	// 层级5：仅产品（跨供货商）
	for _, pc := range tenant.ProductCodes {
		if doc, ok := readHashDoc(ctx, rdb, keyL5ProductOnly(tenant.TenantID, bizType, pc, hourKey),
			newBaseDoc(tenant.TenantID, bizType, stats.StatLevelProductOnly, 0, pc, "", hourUTC, dateUTC, now)); ok {
			docs = append(docs, doc)
		}
	}

	return docs, nil
}

// readHashDoc 从 Redis 读取 Hash，若有数据则填充文档并返回
func readHashDoc(ctx context.Context, rdb redis.UniversalClient, key string, base stats.SupplierOrderHourlyStat) (stats.SupplierOrderHourlyStat, bool) {
	vals, err := rdb.HGetAll(ctx, key).Result()
	if err != nil || len(vals) == 0 {
		return stats.SupplierOrderHourlyStat{}, false
	}
	base.TotalOrders = parseInt(vals[fieldTotal])
	base.SuccessOrders = parseInt(vals[fieldSuccess])
	base.SuccessAmount = parseInt(vals[fieldSuccessAmount])
	base.TotalAmount = parseInt(vals[fieldTotalAmount])
	return base, true
}

// upsertStat 幂等写入一条统计文档到 MongoDB
func upsertStat(ctx context.Context, coll *qmgo.Collection, doc stats.SupplierOrderHourlyStat) error {
	filter := bson.M{
		"tenant_id":     doc.TenantID,
		"business_type": doc.BusinessType,
		"stat_level":    doc.StatLevel,
		"supplier_id":   doc.SupplierID,
		"product_code":  doc.ProductCode,
		"amount":        doc.Amount,
		"hour_utc":      doc.HourUTC,
	}
	update := bson.M{
		"$set": bson.M{
			"total_orders":   doc.TotalOrders,
			"success_orders": doc.SuccessOrders,
			"success_amount": doc.SuccessAmount,
			"total_amount":   doc.TotalAmount,
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

// newBaseDoc 构建文档骨架
func newBaseDoc(tenantID uint, bizType string, level stats.StatLevel, supplierID uint, productCode, amount string, hourUTC int64, dateUTC string, now int64) stats.SupplierOrderHourlyStat {
	return stats.SupplierOrderHourlyStat{
		TenantID:     tenantID,
		BusinessType: bizType,
		StatLevel:    level,
		SupplierID:   supplierID,
		ProductCode:  productCode,
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
