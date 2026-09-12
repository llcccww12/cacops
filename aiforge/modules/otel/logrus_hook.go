package otel

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
)

// -------------------------- logrus Hook 定义 --------------------------
// otelHook logrus 的 OTEL 日志转发 Hook，适配 otel/log v0.14.0
type otelHook struct {
	otelLogger log.Logger     // v0.14.0 的 log.Logger 接口（核心接口层）
	levels     []logrus.Level // 需要转发的日志级别
}

// NewOtelLogrusHook 创建 logrus 的 OTEL Hook 实例
// levels：指定转发的日志级别，传 nil 则转发所有级别
func NewOtelLogrusHook(levels []logrus.Level) *otelHook {
	// 从全局 LoggerProvider 获取 v0.14.0 的 OTEL Logger 实例
	otelLogger := global.GetLoggerProvider().Logger("logrus-otel-bridge")
	if levels == nil {
		levels = logrus.AllLevels
	}
	return &otelHook{
		otelLogger: otelLogger,
		levels:     levels,
	}
}

// Levels 实现 logrus.Hook 接口，返回需要监听的日志级别
func (h *otelHook) Levels() []logrus.Level {
	return h.levels
}

// Fire 实现 logrus.Hook 接口，核心转发逻辑（适配新 Record 结构体）
func (h *otelHook) Fire(entry *logrus.Entry) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithError(fmt.Errorf("panic: %v", r)).Error("OTel Hook 转发日志失败")
		}
	}()

	ctx := entry.Context
	if ctx == nil {
		// 从自定义GLS中提取存储的r.Context()（带trace_id/span_id）
		glsVal, ok := GlsGet(GetGlsKey())
		if ok {
			// 类型断言为context.Context
			if c, ok := glsVal.(context.Context); ok {
				ctx = c
			}
		}
	}
	if ctx == nil {
		ctx = context.Background()
	}

	// 1. 创建空的 Record 实例（新结构体需先初始化，再通过方法设值）
	var otelRecord log.Record

	// 2. 设置日志基础字段（通过 Record 的公开方法赋值，私有字段无法直接修改）
	otelRecord.SetTimestamp(entry.Time)                           // 设置日志时间
	otelRecord.SetBody(log.StringValue(entry.Message))            // 设置日志体
	otelRecord.SetSeverity(convertLogrusLevelToOTEL(entry.Level)) // 设置日志级别
	otelRecord.SetSeverityText(otelRecord.Severity().String())    // 设置级别文本描述

	// 3. 注入追踪上下文（trace_id/span_id）
	InjectTraceContext(ctx, &otelRecord)

	// 4. 注入 logrus 自定义字段
	h.injectCustomFields(entry.Data, &otelRecord)

	// 5. 为 Panic 级别添加自定义标记
	if entry.Level == logrus.PanicLevel {
		otelRecord.AddAttributes(log.KeyValue{
			Key:   "logrus_level",
			Value: log.StringValue("panic"),
		})
	}

	// 6. 发送日志到 OTel Collector
	h.otelLogger.Emit(ctx, otelRecord)

	return nil
}

// -------------------------- 工具方法 --------------------------
// convertLogrusLevelToOTEL 将 logrus 级别转换为 OTel Severity 枚举
func convertLogrusLevelToOTEL(level logrus.Level) log.Severity {
	switch level {
	case logrus.DebugLevel:
		return log.SeverityDebug
	case logrus.InfoLevel:
		return log.SeverityInfo
	case logrus.WarnLevel:
		return log.SeverityWarn
	case logrus.ErrorLevel:
		return log.SeverityError
	case logrus.FatalLevel, logrus.PanicLevel:
		return log.SeverityFatal
	default:
		return log.SeverityInfo
	}
}

// injectCustomFields 将 logrus 自定义字段注入到 Record
func (h *otelHook) injectCustomFields(fields logrus.Fields, record *log.Record) {
	for k, v := range fields {
		record.AddAttributes(convertValueToOTELAttr(k, v))
	}
}

// convertValueToOTELAttr 将任意类型转换为 OTel KeyValue
func convertValueToOTELAttr(key string, value interface{}) log.KeyValue {
	var val log.Value
	switch v := value.(type) {
	case string:
		val = log.StringValue(v)
	case int:
		val = log.Int64Value(int64(v))
	case int64:
		val = log.Int64Value(v)
	case int32:
		val = log.Int64Value(int64(v))
	case uint:
		val = log.Int64Value(int64(v))
	case uint64:
		val = log.StringValue(strconv.FormatUint(v, 10)) // 大数值转字符串避免溢出
	case uint32:
		val = log.Int64Value(int64(v))
	case float64:
		val = log.Float64Value(v)
	case float32:
		val = log.Float64Value(float64(v))
	case bool:
		val = log.BoolValue(v)
	case time.Time:
		val = log.StringValue(v.Format(time.RFC3339))
	case error:
		val = log.StringValue(v.Error())
	default:
		val = log.StringValue(fmt.Sprintf("%v", v))
	}
	return log.KeyValue{Key: key, Value: val}
}
