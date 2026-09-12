package otel

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
)

// SendLog 发送关联 Trace/Span 的日志（实现日志与追踪链路关联）
func SendLog(ctx context.Context, level log.Severity, msg string) {
	// 获取日志器（指定日志器名称）
	logger := global.Logger("otel-raw-logger")

	var otelRecord log.Record
	otelRecord.SetTimestamp(time.Now())
	otelRecord.SetBody(log.StringValue(msg))
	otelRecord.SetSeverity(level)
	otelRecord.SetSeverityText(level.String())

	// 注入追踪上下文（trace_id/span_id）
	InjectTraceContext(ctx, &otelRecord)

	logger.Emit(ctx, otelRecord)
}
