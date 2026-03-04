package merchant

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/stats/base"
	"time"

	"github.com/k9xR7vA2/recharge-common/model/mongo/stats"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

// ================== 今日数据（Redis）==================

// GetTenantTotal 层级1：租户今日汇总（首页顶部卡片）
func GetTenantTotal(ctx context.Context, rdb redis.UniversalClient, query TodayQuery) (*TenantTotalResult, error) {
	agg, err := readHourRange(ctx, rdb, query.Timezone, func(hour string) string {
		return keyL1Tenant(query.TenantID, query.BusinessType, hour)
	})
	if err != nil {
		return nil, err
	}
	return &TenantTotalResult{StatNumbers: agg.toStatNumbers()}, nil
}

// GetAllMerchantStats 层级2：批量查询所有商户今日统计（首页商户列表）
func GetAllMerchantStats(ctx context.Context, rdb redis.UniversalClient, query TodayQuery, merchantIDs []uint) ([]MerchantStatResult, error) {
	hourRange, err := base.TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[uint]*hashAgg, len(merchantIDs))
	for _, mid := range merchantIDs {
		aggMap[mid] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		mid uint
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, mid := range merchantIDs {
			k := keyL2Merchant(query.TenantID, query.BusinessType, mid, hour)
			cmds = append(cmds, ref{mid, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/merchant reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.mid].merge(vals)
		}
	}

	results := make([]MerchantStatResult, 0, len(merchantIDs))
	for _, mid := range merchantIDs {
		results = append(results, MerchantStatResult{
			MerchantID:  mid,
			StatNumbers: aggMap[mid].toStatNumbers(),
		})
	}
	return results, nil
}

// GetAllChannelStats 层级3：批量查询某商户下所有通道今日统计
func GetAllChannelStats(ctx context.Context, rdb redis.UniversalClient, query MerchantQuery, channelCodes []string) ([]ChannelStatResult, error) {
	hourRange, err := base.TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[string]*hashAgg, len(channelCodes))
	for _, cc := range channelCodes {
		aggMap[cc] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		cc  string
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, cc := range channelCodes {
			k := keyL3Channel(query.TenantID, query.BusinessType, query.MerchantID, cc, hour)
			cmds = append(cmds, ref{cc, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/merchant reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.cc].merge(vals)
		}
	}

	results := make([]ChannelStatResult, 0, len(channelCodes))
	for _, cc := range channelCodes {
		a := aggMap[cc]
		results = append(results, ChannelStatResult{
			MerchantID:  query.MerchantID,
			ChannelCode: cc,
			ChannelType: a.channelType,
			StatNumbers: a.toStatNumbers(),
		})
	}
	return results, nil
}

// GetAmountStats 层级4：查询某商户某通道所有面额今日统计
func GetAmountStats(ctx context.Context, rdb redis.UniversalClient, query ChannelQuery, amounts []string) ([]AmountStatResult, error) {
	hourRange, err := base.TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[string]*hashAgg, len(amounts))
	for _, a := range amounts {
		aggMap[a] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		amount string
		cmd    *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, a := range amounts {
			k := keyL4Amount(query.TenantID, query.BusinessType, query.MerchantID, query.ChannelCode, a, hour)
			cmds = append(cmds, ref{a, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/merchant reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.amount].merge(vals)
		}
	}

	results := make([]AmountStatResult, 0, len(amounts))
	for _, a := range amounts {
		ag := aggMap[a]
		results = append(results, AmountStatResult{
			MerchantID:  query.MerchantID,
			ChannelCode: query.ChannelCode,
			ChannelType: ag.channelType,
			Amount:      a,
			StatNumbers: ag.toStatNumbers(),
		})
	}
	return results, nil
}

// GetAllChannelOnlyStats 层级5：批量查询所有通道今日统计（跨商户，通道成功率面板）
// channelCodes 由调用方从 DB 查出后传入
// 返回结果已按成功率倒序排列，前端直接渲染
func GetAllChannelOnlyStats(ctx context.Context, rdb redis.UniversalClient, query TodayQuery, channelCodes []string, channelTypeMap map[string]int) ([]ChannelStatResult, error) {
	hourRange, err := base.TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[string]*hashAgg, len(channelCodes))
	for _, cc := range channelCodes {
		aggMap[cc] = &hashAgg{}
		// channel_type 从调用方传入的 map 补充（通道表是 MySQL，不存在 Redis 里）
		if ct, ok := channelTypeMap[cc]; ok {
			aggMap[cc].channelType = ct
		}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		cc  string
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, cc := range channelCodes {
			k := keyL5ChannelOnly(query.TenantID, query.BusinessType, cc, hour)
			cmds = append(cmds, ref{cc, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/merchant reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.cc].merge(vals)
		}
	}

	results := make([]ChannelStatResult, 0, len(channelCodes))
	for _, cc := range channelCodes {
		a := aggMap[cc]
		results = append(results, ChannelStatResult{
			MerchantID:  0,
			ChannelCode: cc,
			ChannelType: a.channelType,
			StatNumbers: a.toStatNumbers(),
		})
	}
	// 按成功率倒序排（前端 ChannelList 组件展示用）
	sortBySuccessRate(results)
	return results, nil
}

// GetChannelOnlyStat 层级5：查询单个通道今日统计
func GetChannelOnlyStat(ctx context.Context, rdb redis.UniversalClient, query ChannelOnlyQuery, channelType int) (*ChannelStatResult, error) {
	agg, err := readHourRange(ctx, rdb, query.Timezone, func(hour string) string {
		return keyL5ChannelOnly(query.TenantID, query.BusinessType, query.ChannelCode, hour)
	})
	if err != nil {
		return nil, err
	}
	if agg.channelType == 0 {
		agg.channelType = channelType
	}
	return &ChannelStatResult{
		MerchantID:  0,
		ChannelCode: query.ChannelCode,
		ChannelType: agg.channelType,
		StatNumbers: agg.toStatNumbers(),
	}, nil
}

// ================== 历史数据（MongoDB）==================

// GetHistoryStats 从 MongoDB 查历史统计（按时区日期）
func GetHistoryStats(ctx context.Context, mongoClient *qmgo.Client, query HistoryQuery) ([]stats.MerchantOrderHourlyStat, error) {
	hourRange, err := base.DateHourRange(query.Date, query.Timezone)
	if err != nil {
		return nil, err
	}

	dbName := fmt.Sprintf("tenant_%d", query.TenantID)
	coll := mongoClient.Database(dbName).Collection(stats.CollMerchantOrderHourlyStats)

	filter := bson.M{
		"tenant_id":     query.TenantID,
		"business_type": query.BusinessType,
		"stat_level":    query.StatLevel,
		"hour_utc":      bson.M{"$gte": hourRange.MinHour, "$lte": hourRange.MaxHour},
	}
	if query.StatLevel >= 2 && query.MerchantID > 0 {
		filter["merchant_id"] = query.MerchantID
	}
	if query.StatLevel >= 3 && query.ChannelCode != "" {
		filter["channel_code"] = query.ChannelCode
	}
	if query.StatLevel >= 4 && query.Amount != "" {
		filter["amount"] = query.Amount
	}
	if query.StatLevel == 5 {
		if query.ChannelCode != "" {
			filter["channel_code"] = query.ChannelCode
		}
		// 支持按通道类型过滤（0=不过滤）
		if query.ChannelType > 0 {
			filter["channel_type"] = query.ChannelType
		}
	}

	var result []stats.MerchantOrderHourlyStat
	if err = coll.Find(ctx, filter).All(&result); err != nil {
		return nil, fmt.Errorf("stats/merchant reader: mongo query: %w", err)
	}
	return result, nil
}

// AggregateHistoryByDay 将小时记录聚合为当日汇总
// 返回 map key 格式："{merchantID}:{channelCode}:{amount}"
func AggregateHistoryByDay(records []stats.MerchantOrderHourlyStat) map[string]StatNumbers {
	type key struct {
		merchantID  uint
		channelCode string
		amount      string
	}
	aggMap := make(map[key]*hashAgg)
	for _, r := range records {
		k := key{r.MerchantID, r.ChannelCode, r.Amount}
		if aggMap[k] == nil {
			aggMap[k] = &hashAgg{channelType: r.ChannelType}
		}
		aggMap[k].total += r.TotalOrders
		aggMap[k].success += r.SuccessOrders
		aggMap[k].successAmount += r.SuccessAmount
		aggMap[k].totalAmount += r.TotalAmount
	}

	result := make(map[string]StatNumbers, len(aggMap))
	for k, agg := range aggMap {
		result[fmt.Sprintf("%d:%s:%s", k.merchantID, k.channelCode, k.amount)] = agg.toStatNumbers()
	}
	return result
}

// ================== 内部工具 ==================

func readHourRange(ctx context.Context, rdb redis.UniversalClient, timezone string, keyFn func(hour string) string) (hashAgg, error) {
	hourRange, err := base.TodayHourRange(timezone)
	if err != nil {
		return hashAgg{}, err
	}
	var agg hashAgg
	for _, hour := range hourRange.Hours {
		vals, e := rdb.HGetAll(ctx, keyFn(hour)).Result()
		if e != nil || len(vals) == 0 {
			continue
		}
		agg.merge(vals)
	}
	return agg, nil
}

// sortBySuccessRate 按成功率倒序排列（插入排序，数量不大）
func sortBySuccessRate(results []ChannelStatResult) {
	for i := 1; i < len(results); i++ {
		cur := results[i]
		j := i - 1
		for j >= 0 && results[j].SuccessRate < cur.SuccessRate {
			results[j+1] = results[j]
			j--
		}
		results[j+1] = cur
	}
}

// ShouldReadFromRedis 判断是否读 Redis（今日）还是 MongoDB（历史）
func shouldReadFromRedis(dateStr, timezone string) (bool, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return false, err
	}
	target, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return false, err
	}
	now := time.Now().In(loc)
	return target.Year() == now.Year() && target.YearDay() == now.YearDay(), nil
}
