package orderpool

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/orderpool/keys"
	"github.com/small-cat1/recharge-common/orderpool/options"
	"time"
)

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

	return nil
}

// FetchOrder 取出订单
func (m *MobileOrderPool) FetchOrder(ctx context.Context, opts options.IFetchMobileOrderOptions, workerID string) (string, string, error) {
	fetchScript := `
		local poolCount = (#KEYS) / 2
		local orderKeyPrefix = ARGV[3]
		local groupName = "processors"
		local validOrder = nil

		for i = 1, poolCount do
			local poolKey = KEYS[i]
			local priority = KEYS[i + poolCount]

			pcall(redis.call, 'XGROUP', 'CREATE', poolKey, groupName, '0', 'MKSTREAM')

			local maxAttempts = 100
			for attempt = 1, maxAttempts do
				if validOrder then break end

				local messages = redis.call('XREADGROUP', 'GROUP', groupName, ARGV[1], 'COUNT', 1, 'STREAMS', poolKey, '>')
				if not messages or #messages == 0 or #messages[1][2] == 0 then break end

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
						redis.call('DEL', orderKeyPrefix .. order_sn)
					end
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

		if validOrder then
			return {validOrder[1], validOrder[2]}
		else
			return {"", ""}
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
	if !ok || len(resultArr) < 2 {
		return "", "", ErrUnknownResultFormat
	}

	orderInfoStr, _ := resultArr[0].(string)
	priority, _ := resultArr[1].(string)
	if orderInfoStr == "" {
		return "", "", ErrNoOrderAvailable
	}

	return orderInfoStr, priority, nil
}

// CancelOrRemoveOrder 撤销/删除订单
func (m *MobileOrderPool) CancelOrRemoveOrder(ctx context.Context, opts options.IMobileHandlerOptions) error {
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

		if poolKey and messageID then
			pcall(redis.call, 'XACK', poolKey, 'processors', messageID)
			redis.call('XDEL', poolKey, messageID)
		end

		redis.call('DEL', orderKey)

		return {"success"}
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

	return nil
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
