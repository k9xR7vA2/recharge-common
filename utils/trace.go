package utils

import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

// ExtractTraceID 放在utils包中
func ExtractTraceID(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.IsValid() {
		return spanCtx.TraceID().String()
	}
	return "no-trace-id"
}
