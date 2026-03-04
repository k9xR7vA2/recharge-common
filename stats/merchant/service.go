package merchant

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/stats/base"
	"time"

	"github.com/k9xR7vA2/recharge-common/model/mongo/stats"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
)

// MerchantStatService 商户订单统计服务
//
// tenant_api（下单成功后）：
//
//	svc.RecordCreated(ctx, RecordInput{...})
//
// tenant_notify（查单终态成功后）：
//
//	svc.RecordResult(ctx, RecordInput{...})
//
// tenant_api 首页面板：
//
//	svc.GetTenantTotal(ctx, TodayQuery{...})
//	svc.GetAllMerchantStats(ctx, TodayQuery{...}, merchantIDs)
//	svc.GetAllChannelOnlyStats(ctx, TodayQuery{...}, channelCodes, channelTypeMap)
//
// cron 归集：
//
//	svc.FlushPrevHour(ctx, tenants, bizTypes)
type MerchantStatService struct {
	rdb    redis.UniversalClient
	mongo  *qmgo.Client
	logger base.Logger
}

func NewMerchantStatService(rdb redis.UniversalClient, mongoClient *qmgo.Client, logger base.Logger) *MerchantStatService {
	return &MerchantStatService{rdb: rdb, mongo: mongoClient, logger: logger}
}

// ---- 写入 ----

// RecordCreated 下单成功时调用（tenant_api）
func (s *MerchantStatService) RecordCreated(ctx context.Context, input RecordInput) error {
	return RecordOrderCreated(ctx, s.rdb, input)
}

// RecordResult 订单终态时调用（tenant_notify，仅成功时有效）
func (s *MerchantStatService) RecordResult(ctx context.Context, input RecordInput) error {
	return RecordOrderResult(ctx, s.rdb, input)
}

// ---- 今日读取（Redis）----

// GetTenantTotal 层级1：租户今日汇总
func (s *MerchantStatService) GetTenantTotal(ctx context.Context, query TodayQuery) (*TenantTotalResult, error) {
	return GetTenantTotal(ctx, s.rdb, query)
}

// GetAllMerchantStats 层级2：所有商户今日统计（首页商户列表）
func (s *MerchantStatService) GetAllMerchantStats(ctx context.Context, query TodayQuery, merchantIDs []uint) ([]MerchantStatResult, error) {
	return GetAllMerchantStats(ctx, s.rdb, query, merchantIDs)
}

// GetAllChannelStats 层级3：某商户下所有通道今日统计
func (s *MerchantStatService) GetAllChannelStats(ctx context.Context, query MerchantQuery, channelCodes []string) ([]ChannelStatResult, error) {
	return GetAllChannelStats(ctx, s.rdb, query, channelCodes)
}

// GetAmountStats 层级4：某商户某通道所有面额今日统计
func (s *MerchantStatService) GetAmountStats(ctx context.Context, query ChannelQuery, amounts []string) ([]AmountStatResult, error) {
	return GetAmountStats(ctx, s.rdb, query, amounts)
}

// GetAllChannelOnlyStats 层级5：所有通道今日统计（跨商户，已按成功率倒序）
// channelTypeMap: channelCode → channelType，由调用方从 MySQL 查出后传入
func (s *MerchantStatService) GetAllChannelOnlyStats(ctx context.Context, query TodayQuery, channelCodes []string, channelTypeMap map[string]int) ([]ChannelStatResult, error) {
	return GetAllChannelOnlyStats(ctx, s.rdb, query, channelCodes, channelTypeMap)
}

// GetChannelOnlyStat 层级5：单个通道今日统计
func (s *MerchantStatService) GetChannelOnlyStat(ctx context.Context, query ChannelOnlyQuery, channelType int) (*ChannelStatResult, error) {
	return GetChannelOnlyStat(ctx, s.rdb, query, channelType)
}

// ---- 历史读取（MongoDB）----

// GetHistoryStats 历史统计，按时区日期查询
func (s *MerchantStatService) GetHistoryStats(ctx context.Context, query HistoryQuery) ([]stats.MerchantOrderHourlyStat, error) {
	return GetHistoryStats(ctx, s.mongo, query)
}

// GetHistoryDayAgg 历史统计按日聚合
func (s *MerchantStatService) GetHistoryDayAgg(ctx context.Context, tenantID uint, bizType, date, timezone string, level MerStatLevel) (map[string]StatNumbers, error) {
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
func (s *MerchantStatService) Flush(ctx context.Context, input FlushInput) error {
	return Flush(ctx, s.rdb, s.mongo, input, s.logger)
}

// FlushPrevHour 归集上一小时（cron 常用）
func (s *MerchantStatService) FlushPrevHour(ctx context.Context, tenants []FlushTenant, bizTypes []string) error {
	return s.Flush(ctx, FlushInput{
		TargetHourUTC: prevHourInt(time.Now().UTC()),
		Tenants:       tenants,
		BusinessTypes: bizTypes,
	})
}

func (s *MerchantStatService) tenantDB(tenantID uint) *qmgo.Database {
	return s.mongo.Database(fmt.Sprintf("tenant_%d", tenantID))
}
