package traceConstant

// ExporterType 追踪导出器类型
type ExporterType string

const (
	ExporterJaeger    ExporterType = "jaeger"
	ExporterAliARMS   ExporterType = "aliarms"
	ExporterTencentAP ExporterType = "tencentap"
)
