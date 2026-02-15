package utils

import (
	"SaasApi/pkg/constant/constant"
	"github.com/gin-gonic/gin"
)

func ToSupplierContext(c *gin.Context) (*mysql.SupplierContext, bool) {
	// 从原始 context 中获取 supplierContext
	if sc, exists := c.Get(constant.SupplierContextKey); exists {
		if supplierCtx, ok := sc.(*mysql2.SupplierContext); ok {
			return supplierCtx, true
		}
	}
	return nil, false
}

func ToMerchantContext(c *gin.Context) (*mysql2.MerchantContext, bool) {
	// 从原始 context 中获取 MerchantContext
	if mc, exists := c.Get(constant.MerchantContextKey); exists {
		if merchantCtx, ok := mc.(*mysql2.MerchantContext); ok {
			return merchantCtx, true
		}
	}
	return nil, false
}

func GetTraceID(c *gin.Context) (string, bool) {
	if traceID, exists := c.Get(constant.TraceIdContent); exists {
		if traceIDStr, ok := traceID.(string); ok {
			return traceIDStr, true
		}
	}
	return "", false
}

//func ToPaymentOrderContext(c *gin.Context) (*mysql.PaymentOrderContext, bool) {
//	// 从原始 context 中获取 MerchantContext
//	if mc, exists := c.Get(merchantConstant.PaymentOrderContext); exists {
//		if paymentOrderCtx, ok := mc.(*mysql.PaymentOrderContext); ok {
//			return paymentOrderCtx, true
//		}
//	}
//	return nil, false
//}
