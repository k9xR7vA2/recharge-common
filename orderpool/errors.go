package orderpool

import (
	"errors"
	"fmt"
)

// ====== 错误码常量（传入 Lua 使用）======
const (
	ErrCodeOrderExists     = "ORDER_EXISTS"
	ErrCodeOrderAddFailed  = "ADD_FAILED"
	ErrCodeOrderProcessing = "ORDER_PROCESSING"
	ErrCodeOrderNotFound   = "ORDER_NOT_FOUND"
	ErrCodeNoOrder         = "NO_ORDER_AVAILABLE"
)

// ====== 哨兵错误 ======
var (
	// Lua 层返回的业务错误
	ErrOrderExists      = errors.New(ErrCodeOrderExists)
	ErrOrderAddFailed   = errors.New(ErrCodeOrderAddFailed)
	ErrOrderProcessing  = errors.New(ErrCodeOrderProcessing)
	ErrOrderNotFound    = errors.New(ErrCodeOrderNotFound)
	ErrNoOrderAvailable = errors.New(ErrCodeNoOrder)

	// Go 层操作错误
	ErrUnknownResultFormat  = errors.New("unknown result format")
	ErrPoolOperationFailed  = errors.New("pool operation failed")
	ErrOrderDataParseFailed = errors.New("order data parse failed")
	ErrOrderSnMissing       = errors.New("order data missing system_order_sn")
	ErrStatsUpdateFailed    = errors.New("stats update failed")
)

// resolveError 根据错误码返回对应的哨兵错误
func resolveError(errCode string) error {
	switch errCode {
	case ErrCodeOrderExists:
		return ErrOrderExists
	case ErrCodeOrderAddFailed:
		return ErrOrderAddFailed
	case ErrCodeOrderProcessing:
		return ErrOrderProcessing
	case ErrCodeOrderNotFound:
		return ErrOrderNotFound
	case ErrCodeNoOrder:
		return ErrNoOrderAvailable
	default:
		return fmt.Errorf("unknown error code: %s", errCode)
	}
}
