package orderpool

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/orderpool/keys"
	"github.com/small-cat1/recharge-common/orderpool/options"
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

func NewMobileOrderPool(redisClient redis.UniversalClient) *MobileOrderPool {
	return &MobileOrderPool{
		redisClient: redisClient,
	}
}

// AddOrderToPool 订单入池
func (m *MobileOrderPool) AddOrderToPool(ctx context.Context, opts options.IMobileHandlerOptions) error {
	script := `
		local errExists = ARGV[7]
		local errAddFailed = ARGV[8]

		local exists = redis.call('EXISTS', KEYS[2])
		if exists == 1 then
			return {0, errExists}
		end

		local msgId = redis.call('XADD', KEYS[1], '*', 'order_sn', ARGV[1], 'expire_at', ARGV[2], 'retry_count', 0, 'create_time', ARGV[3])
		if not msgId then
			return {0, errAddFailed}
		end

		redis.call('HSET', KEYS[2], 
		  'msg_id', msgId, 
		  'pool_key', KEYS[1], 
		  'order_info', ARGV[4], 
		  'create_time', ARGV[3],
		  'priority', ARGV[6])
		redis.call('EXPIRE', KEYS[2], ARGV[5])

		return {1, msgId}
	`

	validTime := opts.GetValidTimeArg() + constant.ExtraRetentionSeconds
	scriptKeys := []string{
		opts.GetPoolKey(),
		opts.GetOrderKey(),
	}
	now := time.Now().UnixMilli()
	args := []interface{}{
		opts.GetSystemOrderSnArg(), // ARGV[1]
		opts.GetExpiredAtArg(),     // ARGV[2]
		now,                        // ARGV[3]
		opts.GetOrderDataArg(),     // ARGV[4]
		validTime,                  // ARGV[5]
		opts.GetPriority(),         // ARGV[6]
		ErrCodeOrderExists,         // ARGV[7]
		ErrCodeOrderAddFailed,      // ARGV[8]
	}

	result, err := m.redisClient.Eval(ctx, script, scriptKeys, args).Result()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrPoolOperationFailed, err)
	}

	resultArray, ok := result.([]interface{})
	if !ok || len(resultArray) < 2 {
		return ErrUnknownResultFormat
	}

	status, _ := resultArray[0].(int64)
	message, _ := resultArray[1].(string)

	if status == 0 {
		return resolveError(message)
	}

	// 执行非核心操作
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

// FetchOrder 取出订单
func (m *MobileOrderPool) FetchOrder(ctx context.Context, opts options.IFetchMobileOrderOptions, workerID string) (string, string, error) {
	fetchScript := `
    local poolCount = (#KEYS) / 2
    local orderKeyPrefix = ARGV[3]
    local groupName = "processors"

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
                    local orderKey = orderKeyPrefix .. order_sn
                    redis.call('DEL', orderKey)
                end
                expiredCounts[priority] = expiredCounts[priority] + 1
            else
                local order_sn = values["order_sn"]
                local orderKey = orderKeyPrefix .. order_sn
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
	KeyGenerator := keys.NewRedisKeysGenerate(opts.GetTenantId(), keys.RoleSupplier, opts.GetBusinessType())
	highPriorityPoolKey := KeyGenerator.GenerateMobilePoolKey(constant.HighPriority.String(), opts.GetPoolArgs())
	normalPriorityPoolKey := KeyGenerator.GenerateMobilePoolKey(constant.NormalPriority.String(), opts.GetPoolArgs())
	orderKeyPrefix := KeyGenerator.OrderKeyPrefix()

	scriptKeys := []string{
		highPriorityPoolKey,
		normalPriorityPoolKey,
		constant.HighPriority.String(),
		constant.NormalPriority.String(),
	}

	args := []interface{}{
		workerID,
		time.Now().UnixMilli(),
		orderKeyPrefix,
	}

	result, err := m.redisClient.Eval(ctx, fetchScript, scriptKeys, args).Result()
	if err != nil {
		if err == redis.Nil {
			return "", "", ErrNoOrderAvailable
		}
		return "", "", fmt.Errorf("%w: %v", ErrPoolOperationFailed, err)
	}

	if result == nil {
		return "", "", ErrNoOrderAvailable
	}

	resultArr, ok := result.([]interface{})
	if !ok || len(resultArr) < 4 {
		return "", "", ErrUnknownResultFormat
	}

	highExpired, _ := resultArr[2].(int64)
	normalExpired, _ := resultArr[3].(int64)
	poolArgs := opts.GetPoolArgs()

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

	orderInfoStr, _ := resultArr[0].(string)
	priority, _ := resultArr[1].(string)
	if orderInfoStr == "" {
		return "", "", ErrNoOrderAvailable
	}

	var tempData map[string]interface{}
	if err := json.Unmarshal([]byte(orderInfoStr), &tempData); err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrOrderDataParseFailed, err)
	}
	systemOrderSn, ok := tempData["system_order_sn"].(string)
	if !ok || systemOrderSn == "" {
		return "", "", ErrOrderSnMissing
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

// CancelOrRemoveOrder 撤销/删除订单
func (m *MobileOrderPool) CancelOrRemoveOrder(ctx context.Context, opts options.IMobileHandlerOptions, event EventType) error {
	cancelScript := `
		local orderKey = KEYS[1]
		local errNotFound = ARGV[1]
		local errProcessing = ARGV[2]

		local orderInfo = redis.call('HGETALL', orderKey)
		if #orderInfo == 0 then
			return {"error", errNotFound}
		end

		local orderData = {}
		for i = 1, #orderInfo, 2 do
			orderData[orderInfo[i]] = orderInfo[i+1]
		end

		local status = orderData["status"] or ""
		if status == "processing" then
			return {"error", errProcessing}
		end

		local poolKey = orderData["pool_key"]
		local messageID = orderData["msg_id"]
		local priority = orderData["priority"] or "normal"

		local deleted = 0
		if poolKey and messageID then
			pcall(redis.call, 'XACK', poolKey, 'processors', messageID)
			deleted = redis.call('XDEL', poolKey, messageID)
		end

		redis.call('DEL', orderKey)

		return {"success", poolKey, messageID, priority, tostring(deleted)}
	`

	scriptKeys := []string{
		opts.GetOrderKey(),
	}

	args := []interface{}{
		ErrCodeOrderNotFound,   // ARGV[1]
		ErrCodeOrderProcessing, // ARGV[2]
	}

	result, err := m.redisClient.Eval(ctx, cancelScript, scriptKeys, args).Result()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrPoolOperationFailed, err)
	}

	resultArr, ok := result.([]interface{})
	if !ok || len(resultArr) == 0 {
		return ErrUnknownResultFormat
	}

	status, _ := resultArr[0].(string)
	if status == "error" && len(resultArr) > 1 {
		errCode, _ := resultArr[1].(string)
		return resolveError(errCode)
	}

	if status == "success" && len(resultArr) >= 5 {
		poolKey, _ := resultArr[1].(string)
		msgID, _ := resultArr[2].(string)
		priority, _ := resultArr[3].(string)
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

	return ErrUnknownResultFormat
}

// GetOrderInfo 获取订单信息
func (m *MobileOrderPool) GetOrderInfo(ctx context.Context, orderKey string) (map[string]interface{}, error) {
	result, err := m.redisClient.HGetAll(ctx, orderKey).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrOrderNotFound
	}
	orderInfo := make(map[string]interface{})
	for k, v := range result {
		orderInfo[k] = v
	}
	return orderInfo, nil
}

// executeNonCoreOperations 执行非核心操作（统计更新和事件记录）
func (m *MobileOrderPool) executeNonCoreOperations(ctx context.Context,
	params NonCoreOpParams,
	eventType EventType,
	eventDetails interface{},
	incr int64) error {

	nonCoreLua := `
        local statsKey = KEYS[1]
        local incrValue = tonumber(ARGV[11])

        redis.call('HINCRBY', statsKey, 'amount:'..ARGV[1], incrValue)
        redis.call('HINCRBY', statsKey, 'carrier:'..ARGV[2], incrValue)
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3], incrValue)
        redis.call('HINCRBY', statsKey, 'area:'..ARGV[4], incrValue)
        redis.call('HINCRBY', statsKey, 'province:'..ARGV[5], incrValue)
        redis.call('HINCRBY', statsKey, 'pool_orders', incrValue)
        redis.call('HINCRBY', statsKey, 'priority:'..ARGV[6], incrValue)
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':carrier:'..ARGV[2], incrValue)
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':amount:'..ARGV[1], incrValue)
        redis.call('HINCRBY', statsKey, 'carrier:'..ARGV[2]..':amount:'..ARGV[1], incrValue)
        redis.call('HINCRBY', statsKey, 'charge_speed:'..ARGV[3]..':carrier:'..ARGV[2]..':amount:'..ARGV[1], incrValue)

        local eventType = ARGV[8]
        if eventType == 'processing' then
            redis.call('HINCRBY', statsKey, 'processing_orders', 1)
        elseif eventType == 'completed' then
            redis.call('HINCRBY', statsKey, 'processing_orders', -1)
        end

        local eventsKey = KEYS[2]
        redis.call('XADD', eventsKey, '*', 
            'order_sn', ARGV[7], 
            'event', ARGV[8], 
            'timestamp', ARGV[9], 
            'details', ARGV[10]
        )
        redis.call('EXPIRE', eventsKey, 21600)
        
        return 'OK'
    `

	var detailsStr string
	switch v := eventDetails.(type) {
	case string:
		detailsStr = v
	default:
		data, _ := json.Marshal(v)
		detailsStr = string(data)
	}

	now := time.Now().UnixMilli()
	scriptKeys := []string{
		params.StatsKey,
		params.EventsKey,
	}
	args := []interface{}{
		params.Amount,
		params.Carrier,
		params.ChargeSpeed,
		params.Region,
		params.Province,
		params.Priority,
		params.SystemOrderSn,
		eventType.String(),
		now,
		detailsStr,
		incr,
	}

	_, err := m.redisClient.Eval(ctx, nonCoreLua, scriptKeys, args).Result()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStatsUpdateFailed, err)
	}
	return nil
}
