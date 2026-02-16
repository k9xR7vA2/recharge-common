package mysql

import (
	"github.com/gin-gonic/gin"
)

type SupplierContext struct {
	*gin.Context
	SupplierInfo *Supplier
	TenantInfo   *Tenant
}

// GetSupplier 获取供应商信息的辅助方法
func (sc *SupplierContext) GetSupplier() *Supplier {
	return sc.SupplierInfo
}

func (sc *SupplierContext) GetTenantInfo() *Tenant {
	return sc.TenantInfo
}
