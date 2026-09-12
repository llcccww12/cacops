package otel

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

// otelHandler 包装slog.Handler，添加追踪上下文字段
type otelHandler struct {
	handler slog.Handler
}

func NewOtelSlogHandler(w *os.File, opts slog.HandlerOptions) *otelHandler {
	return &otelHandler{
		handler: slog.NewJSONHandler(w, &opts), // 输出JSON格式日志（Loki更易解析）
	}
}

// Handle 处理日志记录，添加trace_id/span_id
func (h *otelHandler) Handle(ctx context.Context, r slog.Record) error {
	// 从上下文中提取追踪信息
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// 添加trace_id和span_id字段
		r.AddAttrs(
			slog.String(TRACE_ID, span.SpanContext().TraceID().String()),
			slog.String(SPAN_ID, span.SpanContext().SpanID().String()),
		)
	}
	return h.handler.Handle(ctx, r)
}

// 以下为slog.Handler的默认实现
func (h *otelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *otelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &otelHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *otelHandler) WithGroup(name string) slog.Handler {
	return &otelHandler{handler: h.handler.WithGroup(name)}
}
