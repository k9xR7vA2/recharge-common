package mysql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// SystemConfig 系统配置表结构体
type SystemConfig struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	ConfigKey   string    `json:"configKey" gorm:"uniqueIndex:uk_config_key;size:100;not null;comment:配置键名"`
	ConfigValue string    `json:"configValue" gorm:"type:text;comment:配置值"`
	ValueType   string    `json:"valueType" gorm:"size:20;not null;default:string;comment:值类型:string/number/boolean/json"`
	Category    string    `json:"category" gorm:"size:50;not null;default:system;comment:配置类别"`
	Name        string    `json:"name" gorm:"size:100;not null;comment:配置名称（显示名称）"`
	Description string    `json:"description" gorm:"size:500;comment:配置描述"`
	IsSensitive bool      `json:"isSensitive" gorm:"not null;default:false;comment:是否敏感配置"`
	IsInternal  bool      `json:"isInternal" gorm:"not null;default:false;comment:是否内部配置（不可删除）"`
	SortOrder   int       `json:"sortOrder" gorm:"default:0;comment:排序顺序"`
	CreatedAt   time.Time `json:"createdAt" `
	UpdatedAt   time.Time `json:"updatedAt" `
}

// TableName 指定表名
func (SystemConfig) TableName() string {
	return "system_config"
}

// JSONConfigValue 处理JSON类型的配置值的辅助结构体
type JSONConfigValue map[string]interface{}

// Scan 实现 sql.Scanner 接口，处理从数据库读取的JSON配置值
func (j *JSONConfigValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &j)
}

// Value 实现 driver.Valuer 接口，处理将JSON配置值保存到数据库
func (j JSONConfigValue) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}
