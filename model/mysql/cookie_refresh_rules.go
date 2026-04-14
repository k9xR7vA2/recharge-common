package mysql

import "time"

type CookieRefreshRule struct {
	ID           uint      `json:"id"            gorm:"primaryKey;autoIncrement"`
	ChannelCode  string    `json:"channel_code"  gorm:"type:varchar(50);not null;uniqueIndex:uk_channel_code"`
	IntervalHour int       `json:"interval_hour" gorm:"type:int;not null;default:28"`
	Status       int       `json:"status"        gorm:"type:tinyint;not null;default:1"`
	Remark       string    `json:"remark"        gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (CookieRefreshRule) TableName() string {
	return "as_cookie_refresh_rules"
}
