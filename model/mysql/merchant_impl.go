package mysql

import (
	"github.com/gin-gonic/gin"
)

type MerchantContext struct {
	*gin.Context
	MerchantInfo *Merchant
	TenantInfo   *Tenant
}

// GetMerchant 获取供应商信息的辅助方法
func (mc *MerchantContext) GetMerchant() *Merchant {
	return mc.MerchantInfo
}

func (mc *MerchantContext) GetTenantInfo() *Tenant {
	return mc.TenantInfo
}
