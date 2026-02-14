package mysql

import "time"

type BlacklistToken struct {
	ID        uint   `gorm:"primarykey"`
	Token     string `gorm:"type:varchar(512);uniqueIndex"`
	Reason    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (BlacklistToken) TableName() string {
	return "tenant_blacklists"
}
