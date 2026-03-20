package mysql

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type SupplierAccount struct {
	ID            uint            `json:"id"               gorm:"primaryKey;autoIncrement"`
	SupplierID    uint            `json:"supplier_id"      gorm:"not null"`
	TenantID      uint            `json:"tenant_id"        gorm:"not null"`
	BusinessType  string          `json:"business_type"    gorm:"type:varchar(20);not null"`
	Account       string          `json:"account"          gorm:"type:varchar(100);not null"`
	AccountName   string          `json:"account_name"     gorm:"type:varchar(100)"`
	Contact       string          `json:"contact"          gorm:"type:varchar(50)"`
	Region        string          `json:"region"           gorm:"type:varchar(100)"`
	Amount        decimal.Decimal `json:"amount"           gorm:"type:decimal(12,2);not null;default:0"`
	Source        int             `json:"source"           gorm:"tinyint;not null;default:1"`
	VerifierID    *uint           `json:"verifier_id"`
	VerifierName  string          `json:"verifier_name"    gorm:"type:varchar(50)"`
	Status        int             `json:"status"           gorm:"tinyint;not null;default:1"`
	RejectReason  string          `json:"reject_reason"    gorm:"type:varchar(255)"`
	Remark        string          `json:"remark"           gorm:"type:varchar(255)"`
	SystemOrderSn string          `json:"system_order_sn"  gorm:"type:varchar(64)"`
	ApprovedAt    *time.Time      `json:"approved_at"`
	ApprovedBy    string          `json:"approved_by"      gorm:"type:varchar(50)"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `json:"-"                gorm:"index"`
}

func (SupplierAccount) TableName() string { return "as_supplier_accounts" }
