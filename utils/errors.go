package utils

import (
	"errors"
	"gorm.io/gorm"
)

func HandleDBResult(result *gorm.DB, notFoundErr error) error {
	if result.Error == nil {
		return nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return notFoundErr // 传入自定义的"未找到"错误
	}

	return result.Error // 返回原始数据库错误
}
