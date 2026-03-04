package supplier

import (
	"context"
	"fmt"
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

// GetAllSupplierStats 层级2：批量查询所有供货商今日统计（首页供货商列表）
// supplierIDs 由调用方从 DB 查出后传入
func GetAllSupplierStats(ctx context.Context, rdb redis.UniversalClient, query TodayQuery, supplierIDs []uint) ([]SupplierStatResult, error) {
	hourRange, err := TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[uint]*hashAgg, len(supplierIDs))
	for _, sid := range supplierIDs {
		aggMap[sid] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		sid uint
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, sid := range supplierIDs {
			k := keyL2Supplier(query.TenantID, query.BusinessType, sid, hour)
			cmds = append(cmds, ref{sid, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.sid].merge(vals)
		}
	}

	results := make([]SupplierStatResult, 0, len(supplierIDs))
	for _, sid := range supplierIDs {
		results = append(results, SupplierStatResult{
			SupplierID:  sid,
			StatNumbers: aggMap[sid].toStatNumbers(),
		})
	}
	return results, nil
}

// GetAllProductStats 层级3：批量查询某供货商下所有产品今日统计
// productCodes 由调用方从 DB 查出后传入
func GetAllProductStats(ctx context.Context, rdb redis.UniversalClient, query SupplierQuery, productCodes []string) ([]ProductStatResult, error) {
	hourRange, err := TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[string]*hashAgg, len(productCodes))
	for _, pc := range productCodes {
		aggMap[pc] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		pc  string
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, pc := range productCodes {
			k := keyL3Product(query.TenantID, query.BusinessType, query.SupplierID, pc, hour)
			cmds = append(cmds, ref{pc, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.pc].merge(vals)
		}
	}

	results := make([]ProductStatResult, 0, len(productCodes))
	for _, pc := range productCodes {
		results = append(results, ProductStatResult{
			SupplierID:  query.SupplierID,
			ProductCode: pc,
			StatNumbers: aggMap[pc].toStatNumbers(),
		})
	}
	return results, nil
}

// GetAmountStats 层级4：查询某产品所有面额今日统计（产品详情页，按供货商维度）
// amounts 由调用方从产品配置中取出后传入
func GetAmountStats(ctx context.Context, rdb redis.UniversalClient, query ProductQuery, amounts []string) ([]AmountStatResult, error) {
	hourRange, err := TodayHourRange(query.Timezone)
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
			k := keyL4Amount(query.TenantID, query.BusinessType, query.SupplierID, query.ProductCode, a, hour)
			cmds = append(cmds, ref{a, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.amount].merge(vals)
		}
	}

	results := make([]AmountStatResult, 0, len(amounts))
	for _, a := range amounts {
		results = append(results, AmountStatResult{
			SupplierID:  query.SupplierID,
			ProductCode: query.ProductCode,
			Amount:      a,
			StatNumbers: aggMap[a].toStatNumbers(),
		})
	}
	return results, nil
}

// GetProductOnlyStat 层级5：查询某产品今日统计（跨供货商，不限定 supplierID）
func GetProductOnlyStat(ctx context.Context, rdb redis.UniversalClient, query ProductOnlyQuery) (*ProductStatResult, error) {
	agg, err := readHourRange(ctx, rdb, query.Timezone, func(hour string) string {
		return keyL5ProductOnly(query.TenantID, query.BusinessType, query.ProductCode, hour)
	})
	if err != nil {
		return nil, err
	}
	return &ProductStatResult{
		SupplierID:  0,
		ProductCode: query.ProductCode,
		StatNumbers: agg.toStatNumbers(),
	}, nil
}

// GetAllProductOnlyStats 层级5：批量查询多个产品今日统计（产品管理页面板，跨供货商）
func GetAllProductOnlyStats(ctx context.Context, rdb redis.UniversalClient, query TodayQuery, productCodes []string) ([]ProductStatResult, error) {
	hourRange, err := TodayHourRange(query.Timezone)
	if err != nil {
		return nil, err
	}

	aggMap := make(map[string]*hashAgg, len(productCodes))
	for _, pc := range productCodes {
		aggMap[pc] = &hashAgg{}
	}

	pipe := rdb.Pipeline()
	type ref struct {
		pc  string
		cmd *redis.MapStringStringCmd
	}
	var cmds []ref
	for _, hour := range hourRange.Hours {
		for _, pc := range productCodes {
			k := keyL5ProductOnly(query.TenantID, query.BusinessType, pc, hour)
			cmds = append(cmds, ref{pc, pipe.HGetAll(ctx, k)})
		}
	}
	if _, err = pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("stats/reader: pipeline exec: %w", err)
	}
	for _, r := range cmds {
		if vals, e := r.cmd.Result(); e == nil && len(vals) > 0 {
			aggMap[r.pc].merge(vals)
		}
	}

	results := make([]ProductStatResult, 0, len(productCodes))
	for _, pc := range productCodes {
		results = append(results, ProductStatResult{
			SupplierID:  0,
			ProductCode: pc,
			StatNumbers: aggMap[pc].toStatNumbers(),
		})
	}
	return results, nil
}

// ================== 历史数据（MongoDB）==================

// GetHistoryStats 从 MongoDB 查历史统计（管理后台，按时区日期查询）
func GetHistoryStats(ctx context.Context, mongoClient *qmgo.Client, query HistoryQuery) ([]stats.SupplierOrderHourlyStat, error) {
	hourRange, err := DateHourRange(query.Date, query.Timezone)
	if err != nil {
		return nil, err
	}

	dbName := fmt.Sprintf("tenant_%d", query.TenantID)
	coll := mongoClient.Database(dbName).Collection(stats.CollSupplierOrderHourlyStats)

	filter := bson.M{
		"tenant_id":     query.TenantID,
		"business_type": query.BusinessType,
		"stat_level":    query.StatLevel,
		"hour_utc":      bson.M{"$gte": hourRange.MinHour, "$lte": hourRange.MaxHour},
	}
	if query.StatLevel >= 2 && query.SupplierID > 0 {
		filter["supplier_id"] = query.SupplierID
	}
	if query.StatLevel >= 3 && query.ProductCode != "" {
		filter["product_code"] = query.ProductCode
	}
	if query.StatLevel >= 4 && query.Amount != "" {
		filter["amount"] = query.Amount
	}
	// 层级5：仅产品，不带 supplier_id 过滤
	if query.StatLevel == 5 && query.ProductCode != "" {
		filter["product_code"] = query.ProductCode
	}

	var result []stats.SupplierOrderHourlyStat
	if err = coll.Find(ctx, filter).All(&result); err != nil {
		return nil, fmt.Errorf("stats/reader: mongo query: %w", err)
	}
	return result, nil
}

// AggregateHistoryByDay 将小时记录聚合为当日汇总
// 返回 map key 格式："{supplierID}:{productCode}:{amount}"
func AggregateHistoryByDay(records []stats.SupplierOrderHourlyStat) map[string]StatNumbers {
	type key struct {
		supplierID  uint
		productCode string
		amount      string
	}
	aggMap := make(map[key]*hashAgg)
	for _, r := range records {
		k := key{r.SupplierID, r.ProductCode, r.Amount}
		if aggMap[k] == nil {
			aggMap[k] = &hashAgg{}
		}
		aggMap[k].total += r.TotalOrders
		aggMap[k].success += r.SuccessOrders
		aggMap[k].successAmount += r.SuccessAmount
		aggMap[k].totalAmount += r.TotalAmount
	}

	result := make(map[string]StatNumbers, len(aggMap))
	for k, agg := range aggMap {
		result[fmt.Sprintf("%d:%s:%s", k.supplierID, k.productCode, k.amount)] = agg.toStatNumbers()
	}
	return result
}

// ShouldReadFromRedis 判断是否应读 Redis（今日）还是 MongoDB（历史）
func ShouldReadFromRedis(dateStr, timezone string) (bool, error) {
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

// ================== 内部工具 ==================

// readHourRange 封装单 key 场景下按小时范围遍历读取的公共逻辑
func readHourRange(ctx context.Context, rdb redis.UniversalClient, timezone string, keyFn func(hour string) string) (hashAgg, error) {
	hourRange, err := TodayHourRange(timezone)
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
