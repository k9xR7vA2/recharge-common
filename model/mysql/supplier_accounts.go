package mysql

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/gorm"
	"time"
)

type SupplierAccount struct {
	ID              uint                           `json:"id"               gorm:"primaryKey;autoIncrement"`
	SupplierID      uint                           `json:"supplier_id"      gorm:"not null;comment:供货商ID"`
	TenantID        uint                           `json:"tenant_id"        gorm:"not null;comment:租户ID"`
	BusinessType    constant.BusinessType          `json:"business_type"    gorm:"type:varchar(20);not null;comment:业务类型 electric/oil/dth"`
	ProductCode     string                         `json:"product_code"     gorm:"type:varchar(50);not null;comment:产品编码"`
	Account         string                         `json:"account"          gorm:"type:varchar(100);not null;comment:充值账号（电表号/车牌/DTH号）"`
	AccountName     string                         `json:"account_name"     gorm:"type:varchar(100);comment:账号归属人姓名"`
	Contact         string                         `json:"contact"          gorm:"type:varchar(50);comment:联系方式"`
	Region          string                         `json:"region"           gorm:"type:varchar(100);comment:所属地区"`
	Amount          uint                           `json:"amount" gorm:"type:int unsigned;not null;default:0;comment:总待充金额"`
	ChargedAmount   uint                           `json:"charged_amount" gorm:"type:int unsigned;not null;default:0;comment:已充成功金额"`
	LockAmount      uint                           `json:"lock_amount" gorm:"type:int unsigned;not null;default:0;comment:冻结中金额"`
	SplitCharge     int                            `json:"split_charge"      gorm:"type:tinyint;not null;default:1;comment:是否允许拆分充值 1允许 2不允许"`
	Source          int                            `json:"source"           gorm:"type:tinyint;not null;default:1;comment:来源 1供货商导入 2核销员录入"`
	VerifierID      *uint                          `json:"verifier_id"      gorm:"comment:核销员ID（source=2时有值）"`
	VerifierName    string                         `json:"verifier_name"    gorm:"type:varchar(50);comment:核销员姓名"`
	Status          constant.SupplierAccountStatus `json:"status"           gorm:"type:tinyint;not null;default:1;comment:1待审核 2审核通过 3充值中 4已完成 5已拒绝 6禁用"`
	RejectReason    string                         `json:"reject_reason"    gorm:"type:varchar(255);comment:拒绝原因"`
	Remark          string                         `json:"remark"           gorm:"type:varchar(255);comment:备注"`
	SupplierOrderSn string                         `json:"supplier_order_sn" gorm:"type:varchar(64);uniqueIndex;comment:供货商订单号"`
	NotifyUrl       string                         `json:"notify_url"       gorm:"type:varchar(255);comment:回调通知地址"`
	ApprovedAt      *time.Time                     `json:"approved_at"      gorm:"comment:审核时间"`
	ApprovedBy      string                         `json:"approved_by"      gorm:"type:varchar(50);comment:审核人"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	CompletedAt     *time.Time                     `json:"completed_at" gorm:"comment:充值完成时间"` // 新增
	DeletedAt       gorm.DeletedAt                 `json:"-"                gorm:"index"`
	Supplier        Supplier                       `gorm:"foreignKey:SupplierID;references:SupplierID" json:"supplier,omitempty"`
	Tenant          Tenant                         `gorm:"->;foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`
}

func (SupplierAccount) TableName() string { return "as_supplier_accounts" }
