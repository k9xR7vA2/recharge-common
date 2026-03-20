package mysql

import (
	"gorm.io/gorm"
	"time"
)

type SupplierVerifier struct {
	ID          uint           `json:"id"            gorm:"primaryKey;autoIncrement"`
	SupplierID  uint           `json:"supplier_id"   gorm:"not null"`
	TenantID    uint           `json:"tenant_id"     gorm:"not null"`
	Name        string         `json:"name"          gorm:"type:varchar(50);not null"`
	Username    string         `json:"username"      gorm:"type:varchar(50);not null;uniqueIndex"`
	Password    string         `json:"-"             gorm:"type:varchar(255);not null"`
	Status      int            `json:"status"        gorm:"tinyint;not null;default:1"`
	Remark      string         `json:"remark"        gorm:"type:varchar(255)"`
	LastLoginAt *time.Time     `json:"last_login_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"             gorm:"index"`
}

func (SupplierVerifier) TableName() string { return "as_supplier_verifiers" }
