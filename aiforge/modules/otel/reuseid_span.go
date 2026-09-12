package otel

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// ReuseIDSpan 直接嵌入 sdktrace.ReadWriteSpan
// 嵌入后自动继承 sdktrace.ReadWriteSpan 的所有方法（包括 embedded.Span 私有接口实现）
type ReuseIDSpan struct {
	sdktrace.ReadWriteSpan               // 保留原生 Span 的所有关联信息
	customTraceID          trace.TraceID // 自定义 TraceID
	customSpanID           trace.SpanID  // 自定义 SpanID
}

// SpanContext 唯一重写：返回自定义的 TraceID/SpanID，其他属性从原生 Span 继承
// 其他所有方法（End/AddEvent/RecordError 等）都从 *sdktrace.Span 继承，无需手动实现
// 尝试了不成功，上报还是替换前的spanId，不能成功替换上报spanId为自定义的spanId（customSpanID）
func (r *ReuseIDSpan) SpanContext() trace.SpanContext {
	// 获取原生 Span 上下文（包含完整的上报关联信息）
	originalCtx := r.ReadWriteSpan.SpanContext()
	// 构建新的 SpanContext：替换 TraceID/SpanID，保留其他原生属性
	newCtx := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    r.customTraceID,          // 替换为自定义 TraceID
		SpanID:     r.customSpanID,           // 替换为自定义 SpanID
		TraceFlags: originalCtx.TraceFlags(), // 保留原生采样标记
		TraceState: originalCtx.TraceState(), // 保留原生 TraceState
		Remote:     originalCtx.IsRemote(),   // 保留原生远程标记
	})
	return newCtx
}

func SetReuseIDSpan(ctx context.Context, spanName string) context.Context {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		ctx, span := GetTracer().Start(ctx, spanName)
		nativeRWSpan, ok := span.(sdktrace.ReadWriteSpan)
		if ok {
			reuseSpan := &ReuseIDSpan{
				ReadWriteSpan: nativeRWSpan,          // 赋值：将官方原生 Span 实例传给嵌入的接口字段
				customTraceID: spanContext.TraceID(), // 赋值自定义 TraceID
				customSpanID:  spanContext.SpanID(),  // 赋值自定义 SpanID
			}
			return trace.ContextWithSpan(ctx, reuseSpan)
		}
	}
	return ctx
}
