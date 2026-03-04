package supplier

import (
	"context"
	"fmt"
	"time"

	"github.com/k9xR7vA2/recharge-common/model/mongo/stats"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

// mongoUpdateOpts MongoDB upsert 选项（包级变量，collector 复用）
var mongoUpdateOpts = mongooptions.UpdateOptions{Upsert: boolPtr(true)}

func boolPtr(b bool) *bool { return &b }

// SupplierStatService 供货商订单统计服务
//
// 三个项目统一使用：
//
//	svc := supplier.NewSupplierStatService(rdb, mongoClient, logger)
//
// tenant_notify：
//
//	svc.Record(ctx, RecordInput{...})
//
// tenant_api 首页面板：
//
//	svc.GetTenantTotal(ctx, TodayQuery{...})
//	svc.GetAllSupplierStats(ctx, TodayQuery{...}, supplierIDs)
//	svc.GetAllProductOnlyStats(ctx, TodayQuery{...}, productCodes)
//
// tenant_server 管理后台：
//
//	svc.GetHistoryStats(ctx, HistoryQuery{...})
//
// cron 归集：
//
//	svc.FlushPrevHour(ctx, tenants, bizTypes)
type SupplierStatService struct {
	rdb    redis.UniversalClient
	mongo  *qmgo.Client
	logger Logger
}

func NewSupplierStatService(rdb redis.UniversalClient, mongoClient *qmgo.Client, logger Logger) *SupplierStatService {
	return &SupplierStatService{rdb: rdb, mongo: mongoClient, logger: logger}
}

// ---- 写入 ----

// RecordCreated 下单成功时调用（tenant_api domain 层，保存订单后）
// 只累加 total，不影响 success
func (s *SupplierStatService) RecordCreated(ctx context.Context, input RecordInput) error {
	return RecordOrderCreated(ctx, s.rdb, input)
}

// RecordResult 订单终态时调用（tenant_notify，查单成功或失败）
// 成功累加 success；失败不写（fail = total - success）
func (s *SupplierStatService) RecordResult(ctx context.Context, input RecordInput) error {
	return RecordOrderResult(ctx, s.rdb, input)
}

// ---- 今日读取（Redis）----

// GetTenantTotal 层级1：租户今日汇总
func (s *SupplierStatService) GetTenantTotal(ctx context.Context, query TodayQuery) (*TenantTotalResult, error) {
	return GetTenantTotal(ctx, s.rdb, query)
}

// GetAllSupplierStats 层级2：所有供货商今日统计
func (s *SupplierStatService) GetAllSupplierStats(ctx context.Context, query TodayQuery, supplierIDs []uint) ([]SupplierStatResult, error) {
	return GetAllSupplierStats(ctx, s.rdb, query, supplierIDs)
}

// GetAllProductStats 层级3：某供货商下所有产品今日统计
func (s *SupplierStatService) GetAllProductStats(ctx context.Context, query SupplierQuery, productCodes []string) ([]ProductStatResult, error) {
	return GetAllProductStats(ctx, s.rdb, query, productCodes)
}

// GetAmountStats 层级4：某产品所有面额今日统计（含供货商维度）
func (s *SupplierStatService) GetAmountStats(ctx context.Context, query ProductQuery, amounts []string) ([]AmountStatResult, error) {
	return GetAmountStats(ctx, s.rdb, query, amounts)
}

// GetProductOnlyStat 层级5：某产品今日统计（跨供货商）
func (s *SupplierStatService) GetProductOnlyStat(ctx context.Context, query ProductOnlyQuery) (*ProductStatResult, error) {
	return GetProductOnlyStat(ctx, s.rdb, query)
}

// GetAllProductOnlyStats 层级5：多产品今日统计（跨供货商，产品管理页面板）
func (s *SupplierStatService) GetAllProductOnlyStats(ctx context.Context, query TodayQuery, productCodes []string) ([]ProductStatResult, error) {
	return GetAllProductOnlyStats(ctx, s.rdb, query, productCodes)
}

// ---- 历史读取（MongoDB）----

// GetHistoryStats 历史统计，按时区日期查询
func (s *SupplierStatService) GetHistoryStats(ctx context.Context, query HistoryQuery) ([]stats.SupplierOrderHourlyStat, error) {
	return GetHistoryStats(ctx, s.mongo, query)
}

// GetHistoryDayAgg 历史统计按日聚合（常用快捷方法）
func (s *SupplierStatService) GetHistoryDayAgg(ctx context.Context, tenantID uint, bizType, date, timezone string, level int8) (map[string]StatNumbers, error) {
	records, err := GetHistoryStats(ctx, s.mongo, HistoryQuery{
		TenantID:     tenantID,
		BusinessType: bizType,
		Timezone:     timezone,
		Date:         date,
		StatLevel:    level,
	})
	if err != nil {
		return nil, err
	}
	return AggregateHistoryByDay(records), nil
}

// ---- 归集 ----

// Flush 归集指定 UTC 小时数据到 MongoDB
func (s *SupplierStatService) Flush(ctx context.Context, input FlushInput) error {
	return Flush(ctx, s.rdb, s.mongo, input, s.logger)
}

// FlushPrevHour 归集上一小时（cron 常用）
func (s *SupplierStatService) FlushPrevHour(ctx context.Context, tenants []FlushTenant, bizTypes []string) error {
	return s.Flush(ctx, FlushInput{
		TargetHourUTC: prevHourInt(time.Now().UTC()),
		Tenants:       tenants,
		BusinessTypes: bizTypes,
	})
}

func (s *SupplierStatService) tenantDB(tenantID uint) *qmgo.Database {
	return s.mongo.Database(fmt.Sprintf("tenant_%d", tenantID))
}
