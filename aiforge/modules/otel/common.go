package otel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	logger "code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

const (
	TRACE_ID     string = "trace_id"
	SPAN_ID      string = "span_id"
	TRACE_PARENT string = "traceparent"
)

func WithOtel() bool {
	return setting.OtelService.Enabled
}

func OtelWrapHandler() bool {
	return setting.OtelService.WrapHandler
}

func GetOtelEndpoint() string {
	return setting.OtelService.Endpoint
}

func WithOtelInsecure() bool {
	return setting.OtelService.Insecure
}

func GetOtelMetricsInterval() time.Duration {
	return time.Duration(setting.OtelService.Metrics.Interval) * time.Second

}
func GetProjName() string {
	return setting.OtelService.ProjectName
}
func GetServiceDisplayName() string {
	return setting.OtelService.ServiceDisplayName
}
func WithOtelProfiling() bool {
	return setting.OtelService.Profiling.Enabled
}
func GetPyroscopeAddr() string {
	return setting.OtelService.Profiling.Url
}
func WithOtelLogrus() bool {
	return setting.OtelService.WithLog
}
func WithOtelSlog() bool {
	return setting.OtelService.WithSlog
}

func GetGlsKey() string {
	return setting.OtelService.GlsKey
}

func GetOtelMetricsNamespace() string {
	return setting.OtelService.Metrics.Namespace
}
func WithOtelMetrics() bool {
	return setting.OtelService.Metrics.Enabled
}

func OtelMetricsPush() bool {
	return setting.OtelService.Metrics.WithPush
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowTSSuffix() string {
	now := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d%03d", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
}
func GetRootSpanName(suffix string) string {
	return GetProjName() + "-service:http-api:" + suffix + ":" + GetNowTSSuffix()
}

func GetContext() context.Context {
	return context.Background()
}

func GenerateHeader(span trace.Span) string {
	if !IsValidSpan(span) {
		return ""
	}
	return fmt.Sprintf("00-%s-%s-01", span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String())

}

func GenerateCacheValue(span trace.Span) string {
	return span.SpanContext().TraceID().String()

}

func GetScheme(r *http.Request) string {
	if r.URL.Scheme != "" {
		return r.URL.Scheme
	}
	// 优先通过 X-Forwarded-Proto 头（代理场景，如Nginx转发）
	if scheme := r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	// 判断是否为HTTPS连接（直接连接场景）
	if r.TLS != nil {
		return "https"
	}
	// 默认返回http
	return "http"
}
func ToStr(v any) string {
	str, ok := v.(string)
	if ok {
		return str
	}
	if IsCompType(v) {
		return ToJSONString(v)
	} else {
		return ToRawStr(v)
	}
}

func IsCompType(v interface{}) bool {
	return IsStruct(v) || IsMap(v) || IsArray(v) || IsSlice(v)
}

func ToRawStr(v any) string {
	return fmt.Sprintf("%v", v)
}

func ToJSONString(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		logger.Error("Failed to marshal to JSON: err = %v, data = %v", err, v)
		return ToRawStr(v)
	}
	return string(b)
}

func OfType(v interface{}, kind reflect.Kind) bool {
	// 获取v的反射类型
	t := reflect.TypeOf(v)
	if t == nil {
		return false
	}
	// 判断类型的Kind
	return t.Kind() == kind || t.Kind() == reflect.Ptr && t.Elem() != nil && t.Elem().Kind() == kind
}

func IsStruct(v interface{}) bool {
	return OfType(v, reflect.Struct)
}

func IsMap(v interface{}) bool {
	return OfType(v, reflect.Map)
}

func IsArray(v interface{}) bool {
	return OfType(v, reflect.Array)
}

func IsSlice(v interface{}) bool {
	return OfType(v, reflect.Slice)
}

func GetMilli() int64 {
	return time.Now().UnixMilli()
}

func GetInterval(ts int64) int64 {
	return GetMilli() - ts
}

func GetIntervalStr(ts int64) string {
	return fmt.Sprintf("timeTaken = %d ms", GetInterval(ts))
}

func GetUUID() string {
	return uuid.New().String()
}

func GetUUIDNoHyphen() string {
	return strings.ReplaceAll(GetUUID(), "-", "")
}
