package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// OperationRecordLogs  MongoDB版本的操作记录结构
type OperationRecordLogs struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`            // MongoDB的唯一ID
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`       // 创建时间
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`       // 更新时间
	Ip           string             `bson:"ip" json:"ip"`                       // 请求IP
	Method       string             `bson:"method" json:"method"`               // 请求方法
	Path         string             `bson:"path" json:"path"`                   // 请求路径
	Status       int                `bson:"status" json:"status"`               // 请求状态
	Latency      int64              `bson:"latency" json:"latency"`             // 延迟(纳秒)
	Agent        string             `bson:"agent" json:"agent"`                 // 代理
	ErrorMessage string             `bson:"error_message" json:"error_message"` // 错误信息
	Body         string             `bson:"body" json:"body"`                   // 请求Body
	Resp         string             `bson:"resp" json:"resp"`                   // 响应Body
	UserID       int                `bson:"user_id" json:"user_id"`             // 用户ID
	TenantID     uint               `bson:"tenant_id" json:"tenant_id"`         // 租户ID(新增字段)
	UserName     string             `bson:"user_name" json:"user_name"`         // 用户名(替代关联查询)
}

func (receiver OperationRecordLogs) CollName() string {
	return "operation_record_logs"
}

// CreateIndexes CreateIndex 创建MongoDB索引的辅助方法
func (OperationRecordLogs) CreateIndexes() []interface{} {
	return []interface{}{
		map[string]interface{}{
			"key": map[string]int{
				"created_at": -1, // 按创建时间降序索引
			},
			"name": "created_at_index",
		},
		map[string]interface{}{
			"key": map[string]int{
				"tenant_id": 1,
				"user_id":   1,
			},
			"name": "tenant_user_index",
		},
		map[string]interface{}{
			"key": map[string]int{
				"status": 1,
			},
			"name": "status_index",
		},
	}
}

// 添加TTL索引方法(用于自动删除旧记录)
func (OperationRecordLogs) CreateTTLIndex(expireAfterDays int) interface{} {
	return map[string]interface{}{
		"key": map[string]int{
			"created_at": 1,
		},
		"name":               "ttl_index",
		"expireAfterSeconds": 60 * 60 * 24 * expireAfterDays,
	}
}
