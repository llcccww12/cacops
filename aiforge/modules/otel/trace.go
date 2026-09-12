// Package otel 包提供了对接otel-sdk埋点上报traces/logs/metrics/profiles的工具类方法
package otel

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	logger "code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Trace struct {
	TraceID trace.TraceID
	SpanID  trace.SpanID
}

func (t *Trace) IsValid() bool {
	return t != nil && t.TraceID.IsValid() && t.SpanID.IsValid()
}

func (t *Trace) IsValidExt() bool {
	return t != nil && t.TraceID.IsValid() && (t.SpanID.IsValid() || t.SpanID == trace.SpanID{})
}

func (t Trace) String() string {
	return fmt.Sprintf("traceID: %s, spanID: %s", t.TraceID.String(), t.SpanID.String())
}

// CustomTraceIDToOTEL 将自定义32位十六进制字符串转为OTEL的TraceID
// 若字符串长度不足/过长，自动补零/截断到16字节（32位）
func CustomTraceIDToOTEL(customTraceID string) trace.TraceID {
	var traceID trace.TraceID
	// 解码十六进制字符串为字节数组
	bytes, err := hex.DecodeString(customTraceID)
	if err != nil {
		return traceID // 解码失败返回空ID
	}
	// 复制到TraceID（16字节），超出部分截断，不足部分补零
	copy(traceID[:], bytes)
	return traceID
}

// CustomSpanIDToOTEL 将自定义16位十六进制字符串转为OTEL的SpanID
// 若字符串长度不足/过长，自动补零/截断到8字节（16位）
func CustomSpanIDToOTEL(customSpanID string) trace.SpanID {
	var spanID trace.SpanID
	// 解码十六进制字符串为字节数组
	bytes, err := hex.DecodeString(customSpanID)
	if err != nil {
		return spanID // 解码失败返回空ID
	}
	// 复制到SpanID（8字节），超出部分截断，不足部分补零
	copy(spanID[:], bytes)
	return spanID
}

func GenerateRandomTrace() Trace {
	return Trace{
		TraceID: CustomTraceIDToOTEL(GetUUIDNoHyphen()),
		SpanID:  CustomSpanIDToOTEL(GetUUIDNoHyphen()),
	}
}

func IsValidSpan(span trace.Span) bool {
	return span != nil && span.IsRecording() && span.SpanContext().IsValid()
}

// ExtractTraceParent 从*http.Request中提取traceparent头，并解析Trace ID和Span ID
// 如果不存在traceparent header或格式校验失败，返回默认Trace对象，其中TraceID和SpanID均为空
func ExtractTraceParent(r *http.Request) (t Trace, err error) {
	defer func() {
		if err != nil {
			logger.Error("Failed to extract traceparent header from api http request: %v", err)
		}
	}()

	// 检测是否存在traceparent header
	traceParentHeader := r.Header.Get(TRACE_PARENT)
	if traceParentHeader == "" {
		logger.Debug("traceparent header is empty in api http request")
		return
	}

	// 按W3C规范分割traceparent字符串（格式：00-<TraceID>-<SpanID>-<TraceFlags>）
	parts := strings.Split(traceParentHeader, "-")
	// 验证分割后的片段数量是否为4（版本、TraceID、SpanID、TraceFlags）
	if len(parts) != 4 {
		err = fmt.Errorf("invalid %s format, wrong number (#%d) of segments: %s",
			TRACE_PARENT, len(parts), traceParentHeader)
		return
	}

	// 提取各字段并做合法性校验
	version := parts[0]
	traceID := parts[1]
	spanID := parts[2]
	// traceFlags := parts[3] // 若需要采样标志，可保留该字段；01表示采样，00表示不采样

	// 校验版本（当前W3C规范主流版本为00，暂不兼容其他版本）
	if version != "00" {
		err = fmt.Errorf("unsupported %s version: %s, only 00 is supported", TRACE_PARENT, version)
		return
	}

	// 校验TraceID是否符合32个字符
	if len(traceID) != 32 {
		err = fmt.Errorf("invalid traceID format in %s header: %s", TRACE_PARENT, traceID)
		return
	}

	// 校验SpanID是否符合16个字符
	if len(spanID) != 16 {
		err = fmt.Errorf("invalid spanID format in %s header: %s", TRACE_PARENT, spanID)
		spanID = trace.SpanID{}.String()
	}

	t = Trace{
		TraceID: CustomTraceIDToOTEL(traceID),
		SpanID:  CustomSpanIDToOTEL(spanID),
	}
	logger.Info("Found valid trace header from upstream service: traceID = %s, parentSpanID = %s",
		t.TraceID.String(), t.SpanID.String())
	return
}

// GetTraceParentHTTPHeader 将Trace结构体转换为符合W3C规范的traceparent HTTP头字符串
// traceparent: "00-<traceId>-<spanId>-01"，例如 00-f818da1847874000a6e6fdf84f351840-fa983b6c7dc168a7-01
// 其中前缀 00 表示版本号，后缀 01 表示采样标志（01 表示采样，00 表示不采样），固定即可
func GetTraceParentHTTPHeader(t Trace) string {
	return fmt.Sprintf("00-%s-%s-01", t.TraceID.String(), t.SpanID.String())
}

func SetTraceParentHTTPHeader(r *http.Request, t Trace) {
	r.Header.Set(TRACE_PARENT, GetTraceParentHTTPHeader(t))
}

// MarkSpanError 通用的Span错误标记函数
// ctx: 携带Span的Context
// err: 需要记录的错误
// msg: 额外的错误描述（可选）
func MarkSpanError(ctx context.Context, err error) {
	if ctx == nil {
		return
	}

	span := trace.SpanFromContext(ctx)
	if !IsValidSpan(span) { // 检查Span是否有效且处于记录状态
		return
	}

	if err == nil {
		span.SetStatus(codes.Ok, "span success")
	} else {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	}
}

// GetTraceFromCtx 从上下文context提取 traceID/spanID，返回 Trace 结构体
func GetTraceFromCtx(ctx context.Context) (t Trace) {
	if ctx == nil {
		return
	}
	spanContext := trace.SpanContextFromContext(ctx)
	if !spanContext.IsValid() {
		span := trace.SpanFromContext(ctx)
		if IsValidSpan(span) {
			spanContext = span.SpanContext()
		}
	}
	if !spanContext.IsValid() { // 检查SpanContext是否有效
		return
	}
	return Trace{
		TraceID: spanContext.TraceID(),
		SpanID:  spanContext.SpanID(),
	}
}

// SetTrace 在传入的上下文中设置Trace信息，返回新的上下文
// ctx: 传入不包含Trace信息的原始上下文，如HTTP请求的上下文 r.Context()、context.Background()、context.WithValue()等
// t: 包含TraceID和SpanID的Trace结构体，即使TraceID或SpanID为空也支持设置
// isRemote: 标记Trace是否来自远程Parent服务（如HTTP请求头）
// 返回: 新的上下文，包含设置好的SpanContext
func SetTrace(ctx context.Context, t Trace, isRemote bool) context.Context {
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    t.TraceID,
		SpanID:     t.SpanID,
		TraceFlags: trace.FlagsSampled,
		Remote:     isRemote,
		// 若该SpanContext来自远程服务（如HTTP请求头），请设为true
	})
	// 这里生成的spanContext或span是valid的，可用于后续创建childspan，但是调用span.End()无法上报到otel collector
	// 只有通过GetTracer().Start()创建的span，才能上报
	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	logger.Info("Set trace to context: %s", t)
	return ctx
}

// CreateSpan 基于ctx中的trace信息，创建新的可上报的span
// (1) 如果ctx中包含trace信息，且traceID和spanID均有效，则创建同一traceID下以spanID为父Span的一个新的子Span；
//
//	注意这里子Span的SpanID是新生成的，与父Span不同
//
// (2) 如果ctx中包含trace信息，traceID有效但spanID无效（如全零为空），则创建同一traceID下的一个新的根Span；
//
//	注意这里根Span的SpanID也是新生成的
//
// (3) 如果ctx中不包含trace信息，则创建一个新的根Span；
//
//	这里根Span的TraceID和SpanID都是新生成的
//
// 总结一下，支持TraceID外设以关联上游服务父Span，但必须生成一个新的Span才能上报，SpanID不支持外设
func CreateSpan(ctx context.Context, err *error, spanName string) (
	newCtx context.Context, newSpan trace.Span, deferFunc func()) {
	newCtx = ctx
	newSpan = trace.SpanFromContext(ctx)
	deferFunc = func() {}

	if ctx == nil || !WithOtel() {
		return
	}

	// 只有通过GetTracer().Start()创建的span，才能成功上报追踪数据
	// 新创建的span已经放入newCtx中，后续可读取
	// spanName为创建的新span的名称，Grafana等可视化工具会显示该名称
	newCtx, newSpan = GetTracer().Start(ctx, spanName)
	oldCtx := GlsSet(GetGlsKey(), newCtx)

	// 这里deferFunc用于在span结束时，标记错误、设置结束时间、上报span最终化操作
	// 此函数请在span对应的业务逻辑执行完毕后调用，确保span最终化操作被记录
	deferFunc = func() {
		if err != nil {
			MarkSpanError(newCtx, *err)
		}
		newSpan.SetAttributes(
			attribute.String("finish.time", GetCurrentTimeStr()),
		)
		sendMetrics(newCtx)
		newSpan.End()
		GlsDelete(GetGlsKey(), oldCtx)
	}

	return
}

func CreateChildSpanFromTraceParent(ctx context.Context, traceParent Trace,
	err *error, spanName string, isRemote bool) (
	newCtx context.Context, childSpan trace.Span, deferFunc func()) {
	ctx = SetTrace(ctx, traceParent, isRemote)
	return CreateSpan(ctx, err, spanName)
}

func CreateRootSpan(ctx context.Context, traceID string,
	err *error, spanName string, isRemote bool) (
	newCtx context.Context, childSpan trace.Span, deferFunc func()) {
	trace := Trace{
		TraceID: CustomTraceIDToOTEL(traceID),
		SpanID:  trace.SpanID{},
	}
	return CreateChildSpanFromTraceParent(ctx, trace, err, spanName, isRemote)
}

func UseExistingTrace(ctx context.Context, traceID trace.TraceID, err *error) (context.Context, func()) {
	if !traceID.IsValid() {
		return ctx, nil
	}
	span := trace.SpanFromContext(ctx)
	spanName := GetSpanName(span)
	newCtx, newSpan, deferFunc := CreateRootSpan(ctx, traceID.String(), err, spanName, true)
	attrs := GetSpanAttrs(span)
	if len(attrs) != 0 {
		newSpan.SetAttributes(attrs...)
	}
	return newCtx, deferFunc
}

func InitContextTraceHTTP(ctx context.Context, r *http.Request, err *error, spanName string) (
	newCtx context.Context, span trace.Span, deferFunc func()) {
	t, _ := ExtractTraceParent(r)
	return CreateChildSpanFromTraceParent(ctx, t, err, spanName, true)
}

func FinalizeSpan(ctx context.Context, err error) {
	if !WithOtel() {
		return
	}

	MarkSpanError(ctx, err)
	span := trace.SpanFromContext(ctx)
	if !IsValidSpan(span) {
		return
	}
	span.SetAttributes(
		attribute.String("finish.time", GetCurrentTimeStr()),
	)
	span.End()
}

func GetParentSpanCtx(span trace.Span) trace.SpanContext {
	newspan, ok := span.(sdktrace.ReadOnlySpan)
	if ok {
		return newspan.Parent()
	}
	newspan, ok = span.(sdktrace.ReadWriteSpan)
	if ok {
		return newspan.Parent()
	}
	return trace.SpanContext{}
}

func GetParentSpanID(span trace.Span) trace.SpanID {
	newspan, ok := span.(sdktrace.ReadOnlySpan)
	if ok {
		return newspan.Parent().SpanID()
	}
	newspan, ok = span.(sdktrace.ReadWriteSpan)
	if ok {
		return newspan.Parent().SpanID()
	}
	return trace.SpanID{}
}

func GetSpanName(span trace.Span) string {
	newspan, ok := span.(sdktrace.ReadOnlySpan)
	if ok {
		return newspan.Name()
	}
	newspan, ok = span.(sdktrace.ReadWriteSpan)
	if ok {
		return newspan.Name()
	}
	return ""
}

func GetSpanAttrs(span trace.Span) []attribute.KeyValue {
	if !IsValidSpan(span) {
		return nil
	}
	newspan, ok := span.(sdktrace.ReadOnlySpan)
	if ok {
		return newspan.Attributes()
	}
	newspan, ok = span.(sdktrace.ReadWriteSpan)
	if ok {
		return newspan.Attributes()
	}
	return nil
}

func IsReadOnlySpan(span trace.Span) bool {
	_, ok := span.(sdktrace.ReadOnlySpan)
	return ok
}

func IsReadWriteSpan(span trace.Span) bool {
	_, ok := span.(sdktrace.ReadWriteSpan)
	return ok
}

// GetNopSpan 返回一个非记录 Span，不为NIL，traceID/SpanID全0，用于在不记录追踪数据的场景下使用
func GetNopSpan() trace.Span {
	return trace.SpanFromContext(context.Background())
}

// InjectSpan 将源 context 中的 SpanContext 注入到目标 context 中，返回新的 context
// 若源 context 中无 SpanContext，直接返回目标 context
func InjectSpan(srcCtx, destCtx context.Context) context.Context {
	// 从源 context 中提取 Span
	srcSpan := trace.SpanFromContext(srcCtx)
	if srcSpan.SpanContext().IsValid() {
		// 通过 OTEL 的 trace.ContextWithSpanContext 将 SpanContext 注入到目标 context
		destCtx = trace.ContextWithSpan(destCtx, srcSpan)
	}
	return destCtx
}

// InjectTraceContext 从上下文提取 trace_id/span_id，注入到 Record
func InjectTraceContext(ctx context.Context, record *log.Record) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.IsValid() {
		record.AddAttributes(
			log.KeyValue{
				Key:   TRACE_ID,
				Value: log.StringValue(spanCtx.TraceID().String()),
			},
			log.KeyValue{
				Key:   SPAN_ID,
				Value: log.StringValue(spanCtx.SpanID().String()),
			},
		)
	}
}

func AddRootSpan(r *http.Request, suffix string) (context.Context, trace.Span, *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)
	if WithOtel() && !OtelWrapHandler() && !IsValidSpan(span) {
		tr, _ := ExtractTraceParent(r)
		if tr.IsValidExt() {
			ctx = SetTrace(ctx, tr, true)
		}
		spanName := GetRootSpanName(suffix)
		ctx, span = GetTracer().Start(ctx, spanName)
		uri := r.URL.Path
		span.SetAttributes(
			attribute.Key("http.request.method").String(r.Method),
			attribute.Key("http.request.url_path").String(uri),
			attribute.Key("http.request.url_scheme").String(GetScheme(r)),
			attribute.Key("http.request.url_host").String(r.URL.Host),
			// attribute.Key("http.request.url_username").String(r.URL.User.Username()),
			attribute.Key("http.request.remote_addr").String(r.RemoteAddr),
			attribute.Key("http.request.user_agent").String(r.UserAgent()),
			attribute.Key("http.request.content_length").Int64(r.ContentLength),
			attribute.Key("http.request.protocol").String(r.Proto),
		)
		r = r.WithContext(ctx)
		logger.Info("Add new service root span for api %s: traceID = %s, spanID = %s, spanName = %s",
			uri, span.SpanContext().TraceID(), span.SpanContext().SpanID(), spanName)
	}
	return ctx, span, r
}

func SpantoStr(span trace.Span) string {
	return fmt.Sprintf("traceID = %s, spanID = %s, spanName = %s",
		span.SpanContext().TraceID(), span.SpanContext().SpanID(), GetSpanName(span))
}

func GetTraceInfo(cloudbrain *models.Cloudbrain, spanName string, parentCtx *context.Context) *entity.TraceInfo {
	return &entity.TraceInfo{
		TaskId:    cloudbrain.ID,
		TaskName:  cloudbrain.DisplayJobName,
		JobType:   cloudbrain.JobType,
		SpanName:  spanName,
		ParentCtx: parentCtx,
	}

}

func StartTraceParent(traceInfo *entity.TraceInfo, RequestMethod ...string) (context.Context, trace.Span) {
	if !WithOtel() {
		return GetContext(), GetNopSpan()
	}

	newCtx, span := GetTracer().Start(GetContext(), traceInfo.SpanName, trace.WithSpanKind(trace.SpanKindServer))

	if len(RequestMethod) > 0 {
		span.SetAttributes(
			attribute.String("http.request.method", RequestMethod[0]),
		)
	} else {
		span.SetAttributes(
			attribute.String("http.request.method", "GET"),
		)
	}

	span.SetAttributes(
		attribute.String("task_name", traceInfo.TaskName),
		attribute.String("job_type", traceInfo.JobType),
	)

	return newCtx, span

}

func StartTraceSimple(spanName string, RequestMethod ...string) (context.Context, context.Context, trace.Span) {
	if !WithOtel() {
		return GetContext(), GetContext(), GetNopSpan()
	}

	newCtx, span := GetTracer().Start(GetContext(), spanName, trace.WithSpanKind(trace.SpanKindServer))

	childCtx, childSpan := GetTracer().Start(newCtx, spanName+" client", trace.WithSpanKind(trace.SpanKindClient))
	if len(RequestMethod) > 0 {
		span.SetAttributes(
			attribute.String("http.request.method", RequestMethod[0]),
		)
	} else {
		span.SetAttributes(
			attribute.String("http.request.method", "GET"),
		)
	}
	return newCtx, childCtx, childSpan

}

func StartTraceChild(traceInfo *entity.TraceInfo, RequestMethod ...string) (context.Context, trace.Span) {
	if !WithOtel() || traceInfo == nil || traceInfo.TaskId == 0 || traceInfo.ParentCtx == nil {
		return GetContext(), GetNopSpan()
	}
	patentSpan := trace.SpanFromContext(*traceInfo.ParentCtx)

	childCtx, childSpan := GetTracer().Start(*traceInfo.ParentCtx, traceInfo.SpanName+" client", trace.WithSpanKind(trace.SpanKindClient))

	if traceInfo.JobType != "" {
		patentSpan.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
			attribute.String("task_name", traceInfo.TaskName),
			attribute.String("job_type", traceInfo.JobType),
		)
		childSpan.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
			attribute.String("task_name", traceInfo.TaskName),
			attribute.String("job_type", traceInfo.JobType),
		)
	} else {
		patentSpan.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
		)
		childSpan.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
		)

	}
	if len(RequestMethod) > 0 {
		patentSpan.SetAttributes(
			attribute.String("http.request.method", RequestMethod[0]),
		)
		childSpan.SetAttributes(
			attribute.String("http.request.method", RequestMethod[0]),
		)
	} else {
		patentSpan.SetAttributes(
			attribute.String("http.request.method", "GET"),
		)
		childSpan.SetAttributes(
			attribute.String("http.request.method", "GET"),
		)
	}
	traceValue, _ := redis_client.Get(GetTraceKey(traceInfo.TaskId))
	if traceValue == "" {
		setTraceCache(traceInfo.TaskId, GenerateCacheValue(patentSpan))
	}
	return childCtx, childSpan

}

func StartTrace(traceInfo *entity.TraceInfo, overWriteCache bool, RequestMethod ...string) (context.Context, context.Context, trace.Span) {
	if !WithOtel() || traceInfo == nil || traceInfo.TaskId == 0 {
		return GetContext(), GetContext(), GetNopSpan()
	}

	var link *trace.Link

	traceValue, _ := redis_client.Get(GetTraceKey(traceInfo.TaskId))
	if traceValue != "" && !overWriteCache {
		traceID, err := trace.TraceIDFromHex(traceValue)
		if err != nil {
			logger.Warn("Invalid traceID: %s", err)

		}
		link = &trace.Link{
			SpanContext: trace.NewSpanContext(trace.SpanContextConfig{
				TraceID: traceID,
			}),
			Attributes: []attribute.KeyValue{
				attribute.String("link.type", "task-reference"),
				attribute.String("task.id", strconv.FormatInt(traceInfo.TaskId, 10)),
			},
		}

	}
	var span trace.Span
	var newCtx context.Context
	if link != nil {
		newCtx, span = GetTracer().Start(GetContext(), traceInfo.SpanName, trace.WithLinks(*link), trace.WithSpanKind(trace.SpanKindServer))
	} else {
		newCtx, span = GetTracer().Start(GetContext(), traceInfo.SpanName, trace.WithSpanKind(trace.SpanKindServer))
	}

	childCtx, childSpan := GetTracer().Start(newCtx, traceInfo.SpanName+" client", trace.WithSpanKind(trace.SpanKindClient))

	if traceInfo.JobType != "" {
		span.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
			attribute.String("task_name", traceInfo.TaskName),
			attribute.String("job_type", traceInfo.JobType),
		)
	} else {
		span.SetAttributes(
			attribute.Int64("task_id", traceInfo.TaskId),
		)
	}
	if len(RequestMethod) > 0 {
		span.SetAttributes(
			attribute.String("http.request.method", RequestMethod[0]),
		)
	} else {
		span.SetAttributes(
			attribute.String("http.request.method", "GET"),
		)
	}
	if traceValue == "" || (traceValue != "" && overWriteCache) {
		setTraceCache(traceInfo.TaskId, GenerateCacheValue(span))
	}
	return newCtx, childCtx, childSpan
}

func UpdateTraceCache(oldId int64, newId int64) {
	if !WithOtel() {
		return
	}
	if oldId != 0 && newId != 0 {
		traceValue, _ := redis_client.Get(GetTraceKey(oldId))
		if traceValue != "" {

			setTraceCache(newId, traceValue)
		}

		redis_client.Del(GetTraceKey(oldId))
	}
}
func setTraceCache(id int64, value string) {
	redis_client.Setnx(GetTraceKey(id), value, time.Duration(setting.ClearStrategy.ResultSaveDays*24)*time.Hour)
}

func RemoveTraceCache(id int64, value string) {
	if !WithOtel() {
		return
	}
	traceId, _ := redis_client.Get(GetTraceKey(id))
	if traceId == value {
		redis_client.Del(GetTraceKey(id))
	}
}

func GetTraceKey(id int64) string {

	return "otel" + strconv.FormatInt(id, 10)

}
