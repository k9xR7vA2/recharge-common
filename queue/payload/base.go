package payload

const ExpiredBuffer = 10

type BasePayload struct {
	TraceCtx map[string]string `json:"trace_ctx,omitempty"`
}

// SetTraceCtx 实现 queue.TraceInjectable 接口
func (b *BasePayload) SetTraceCtx(carrier map[string]string) {
	b.TraceCtx = carrier
}
