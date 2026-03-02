package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

// AsInterface [...]
type AsInterface struct {
	InterfaceID   uint                          `json:"interface_id,omitempty" gorm:"autoIncrement:true;primaryKey;column:interface_id;type:int unsigned;not null" `
	BusinessType  constant.BusinessType         `json:"business_type,omitempty" gorm:"not null;comment:业务类型"`
	InterfaceName string                        `json:"interface_name,omitempty" gorm:"column:interface_name;type:varchar(255);not null;comment:'接口名称'" `        // 接口名称
	InterfaceCode string                        `json:"interface_code,omitempty" gorm:"column:interface_code;unique;type:varchar(255);not null;comment:'接口编码'" ` // 接口编码
	RateType      constant.RateType             `json:"rate_type,omitempty" gorm:"column:rate_type;default:NULL;comment:'费率类型，1百分比，2千分比,3单笔固定'"`
	ClientLogin   uint                          `json:"client_login" gorm:"column:client_login;type:tinyint;not null;default:2;comment:客户端登录: 1-开启2-关闭"`
	Rate          int64                         `json:"rate,omitempty" gorm:"column:rate;NOT NULL;comment:'费率'"`
	Status        constant.TenantBusinessStatus `json:"status,omitempty" gorm:"column:status;type:tinyint;not null;comment:'状态1开启，2关闭'" `                           // 状态1开启，2关闭
	NeedCK        constant.NeedCkStatus         `json:"need_ck" gorm:"column:need_ck;type:tinyint;not null;comment:'下单CK1需要，2不需要'" `                                // 下单CK1需要，2不需要
	ConfigParams  datatypes.JSON                `json:"config_params,omitempty" gorm:"column:config_params;type:json;default:null;comment:'配置参数'" `                 // 配置参数
	PaySeconds    int                           `json:"pay_seconds,omitempty" gorm:"column:pay_seconds;type:int;not null;default:90;comment:'支付秒数'" `               // 支付秒数
	QuerySeconds  int                           `json:"query_seconds,omitempty" gorm:"column:query_seconds;type:int unsigned;not null;default:140;comment:'查单秒数'" ` // 查单秒数
	CreatedAt     *time.Time                    `json:"created_at,omitempty"`                                                                                       // 创建时间
	UpdatedAt     *time.Time                    `json:"updated_at,omitempty"`                                                                                       // 更新时间
	DeletedAt     *time.Time                    `json:"-" gorm:"index" `                                                                                            // 删除时间
}

// TableName get sql table name.获取数据库表名
func (AsInterface) TableName() string {
	return "as_interfaces"
}

type InterfaceConfigParams struct {
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	PayUrl        string `json:"pay_url"`        // 下单完整URL
	QueryUrl      string `json:"query_url"`      // 查单完整URL
	QueryInterval int    `json:"query_interval"` // 查单频率（秒）
}

// 在 Interface 模型上加个方法
func (i *AsInterface) GetConfigParams() (*InterfaceConfigParams, error) {
	var params InterfaceConfigParams
	if err := json.Unmarshal(i.ConfigParams, &params); err != nil {
		return nil, fmt.Errorf("解析接口配置参数失败: %w", err)
	}
	return &params, nil
}
