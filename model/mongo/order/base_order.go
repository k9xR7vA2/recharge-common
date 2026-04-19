package order

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/utils"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderBase struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	TraceId    string             `bson:"trace_id" json:"trace_id"`       //链路日志ID
	TenantID   uint               `bson:"tenant_id" json:"tenant_id"`     // 租户ID
	TenantName string             `bson:"tenant_name" json:"tenant_name"` // 租户名称

	AppID string `bson:"app_id" json:"app_id"` // 账户AppID

	BusinessType         constant.BusinessType `bson:"business_type"  json:"business_type"`                  // 业务类型
	SystemOrderSn        string                `bson:"system_order_sn" json:"system_order_sn"`               // 系统订单号
	OfficialSerialNumber string                `bson:"official_serial_number" json:"official_serial_number"` // 官方流水号
	TransactionID        string                `bson:"transaction_id" json:"transaction_id"`                 // 交易ID
	RechargeAccount      string                `bson:"recharge_account" json:"recharge_account"`             // 充值账号
	Amount               string                `bson:"amount" json:"amount"`                                 // 充值金额
	HandingFees          string                `bson:"handing_fees" json:"handing_fees"`                     // 手续费
	IsTest               bool                  `bson:"is_test" json:"is_test"`                               // 是否测试单

	RechargeURL string `bson:"recharge_url" json:"recharge_url"` // 充值地址
	QueryUrl    string `bson:"query_url" json:"query_url"`       // 查单地址
	PullAt      int64  `bson:"pull_at" json:"pull_at"`           // 充值时间
	SuccessAt   int64  `bson:"success_at" json:"success_at"`     // 到账时间

	NotifyURL     string                      `bson:"notify_url" json:"notify_url"`         // 回调地址
	NotifyAt      int64                       `bson:"notify_at" json:"notify_at"`           // 回调时间
	NotifyNumbers uint                        `bson:"notify_numbers" json:"notify_numbers"` // 回调次数
	NotifyStatus  constant.GlobalNotifyStatus `bson:"notify_status" json:"notify_status"`   // 回调状态1准备通知，2通知中，3通知成功，4通知异常，5通知超时
	Remark        string                      `bson:"remark" json:"remark"`                 // 备注
	CallbackLogs  []CallbackLogEntry          `bson:"callback_logs,omitempty"`

	CreatedAt int64 `bson:"created_at" json:"created_at"` // 创建时间
	UpdatedAt int64 `bson:"updated_at" json:"updated_at"` // 更新时间
}

type CallbackLogEntry struct {
	URL          string `bson:"url"`
	Status       int    `bson:"status"`
	StatusCode   int    `bson:"status_code"`
	RetryCount   int    `bson:"retry_count"`
	RequestBody  string `bson:"request_body"`
	ResponseBody string `bson:"response_body"`
	ErrorMsg     string `bson:"error_msg,omitempty"`
	CreatedAt    int64  `bson:"created_at"`
}

// 基础方法
func (o OrderBase) GetID() primitive.ObjectID              { return o.ID }
func (o OrderBase) GetTenantID() uint                      { return o.TenantID }
func (o OrderBase) GetTenantName() string                  { return o.TenantName }
func (o OrderBase) GetAppId() string                       { return o.AppID }
func (o OrderBase) GetSystemOrderSn() string               { return o.SystemOrderSn }
func (o OrderBase) GetAmount() string                      { return o.Amount }
func (o OrderBase) GetBusinessType() constant.BusinessType { return o.BusinessType }

// 追踪相关
func (o OrderBase) GetTraceID() string { return o.TraceId }

// 订单状态相关
func (o OrderBase) GetOfficialSerialNumber() string { return o.OfficialSerialNumber }
func (o OrderBase) GetRechargeAccount() string      { return o.RechargeAccount }
func (o OrderBase) GetHandingFees() decimal.Decimal { return utils.FormatAmount(o.HandingFees) }
func (o OrderBase) GetRechargeURL() string          { return o.RechargeURL }
func (o OrderBase) GetQueryUrl() string             { return o.QueryUrl }

// 回调相关
func (o OrderBase) GetNotifyURL() string                         { return o.NotifyURL }     //回调地址
func (o OrderBase) GetNotifyAt() int64                           { return o.NotifyAt }      //回调时间
func (o OrderBase) GetNotifyNumbers() uint                       { return o.NotifyNumbers } //回调次数
func (o OrderBase) GetNotifyStatus() constant.GlobalNotifyStatus { return o.NotifyStatus }  //回调状态

// 其他信息
func (o OrderBase) GetRemark() string   { return o.Remark } //订单备注
func (o OrderBase) GetCreatedAt() int64 { return o.CreatedAt }
func (o OrderBase) GetPullAt() int64 {
	return o.PullAt
}
func (o OrderBase) GetSuccessAt() int64 {
	return o.SuccessAt
}
func (o OrderBase) GetUpdatedAt() int64 { return o.UpdatedAt } // 更新时间
