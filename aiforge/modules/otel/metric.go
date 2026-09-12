package otel

import (
	"context"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	logger "github.com/sirupsen/logrus"
)

// 定义服务自身相关的Prometheus指标样例
// 使用 promauto 可以自动将指标注册到默认的Registry中，更方便
var (
	// apiRequestTotalCounterVec 记录API请求总数
	apiRequestTotalCounterVec = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "The total number of processed HTTP requests",
	}, []string{"path", "method", "code"}) // 按路径、方法、状态码划分

	// apiRequestDurationHistVec 记录API请求耗时
	// 时长单位毫秒
	apiRequestDurationHistVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_request_duration_seconds",
		Help:    "The duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets, // 使用默认的桶：[.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10]
	}, []string{"path", "method"}) // 按路径、方法划分

	// apiRequestSizeHistVec 记录请求体大小
	apiRequestSizeHistVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_request_size_bytes",
		Help:    "The size of HTTP request body in bytes",
		Buckets: prometheus.ExponentialBuckets(100, 2, 10), // 100B, 200B, 400B... 直到 51200B
	}, []string{"path", "method"})

	// apiResponseSizeHistVec 记录响应体大小
	apiResponseSizeHistVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_response_size_bytes",
		Help:    "The size of HTTP response body in bytes",
		Buckets: prometheus.ExponentialBuckets(100, 2, 10),
	}, []string{"path", "method", "code"})
)

var (
	apiRequestSize      metric.Float64Histogram
	apiResponseSize     metric.Float64Histogram
	apiResponseDuration metric.Float64Histogram
	apiRequestTotal     metric.Float64Counter
)

func getMeter() metric.Meter {
	metricProvider := GetMetricsProvider()
	if metricProvider == nil {
		return nil
	}
	return metricProvider.Meter(GetOtelMetricsNamespace())
}

func getAPIReqTotal(ctx context.Context) metric.Float64Counter {
	if apiRequestTotal != nil {
		return apiRequestTotal
	}
	meter := getMeter()
	if meter == nil {
		return nil
	}
	counter, err := meter.Float64Counter(
		GetProjName()+".api.request.total.count",
		metric.WithDescription("API 请求总数记录"),
	)
	if err != nil {
		logger.WithContext(ctx).Errorf("Failed to create metric counter for api request total: %v", err)
		return nil
	} else {
		logger.WithContext(ctx).Infof("Successfully create metric counter for api request total")
	}
	apiRequestTotal = counter
	return apiRequestTotal
}

func getAPIRespDur(ctx context.Context) metric.Float64Histogram {
	if apiResponseDuration != nil {
		return apiResponseDuration
	}
	meter := getMeter()
	if meter == nil {
		return nil
	}
	// 时长单位毫秒
	hist, err := meter.Float64Histogram(
		GetProjName()+".api.response.duration.hist",
		metric.WithDescription("API响应耗时分布直方图 (毫秒)"),
		metric.WithUnit("milliseconds"),
		metric.WithExplicitBucketBoundaries(0, 100, 1000, 10*1000, 60*1000, 10*60*1000, 30*60*1000, 60*60*1000),
	)
	if err != nil {
		logger.WithContext(ctx).Errorf("Failed to create metric histogram for api response duration: %v", err)
		return nil
	} else {
		logger.WithContext(ctx).Infof("Successfully create metric histogram for api response duration")
	}
	apiResponseDuration = hist
	return apiResponseDuration
}

func getAPIReqSize(ctx context.Context) metric.Float64Histogram {
	if apiRequestSize != nil {
		return apiRequestSize
	}
	meter := getMeter()
	if meter == nil {
		return nil
	}
	hist, err := meter.Float64Histogram(
		GetProjName()+".api.request.size.hist",
		metric.WithDescription("API 请求体大小分布直方图 (字节)"),
		metric.WithUnit("bytes"),
		metric.WithExplicitBucketBoundaries(0, 1024, 1024*1024, 1024*1024*1024),
	)
	if err != nil {
		logger.WithContext(ctx).Errorf("Failed to create metric histogram for api request size: %v", err)
		return nil
	} else {
		logger.WithContext(ctx).Infof("Successfully create metric histogram for api request size")
	}
	apiRequestSize = hist
	return apiRequestSize
}

func getAPIRespSize(ctx context.Context) metric.Float64Histogram {
	if apiResponseSize != nil {
		return apiResponseSize
	}
	meter := getMeter()
	if meter == nil {
		return nil
	}
	// 时长单位毫秒
	hist, err := meter.Float64Histogram(
		GetProjName()+".api.response.size.hist",
		metric.WithDescription("API 响应体大小分布直方图 (字节)"),
		metric.WithUnit("bytes"),
		metric.WithExplicitBucketBoundaries(0, 1024, 1024*1024, 1024*1024*1024),
	)
	if err != nil {
		logger.WithContext(ctx).Errorf("Failed to create metric histogram for api response size: %v", err)
		return nil
	} else {
		logger.WithContext(ctx).Infof("Successfully create metric histogram for api response size")
	}
	apiResponseSize = hist
	return apiResponseSize
}

func AddServiceMetrics(ctx context.Context, r *http.Request, code, respSize int, ts int64) {
	if !WithOtelMetrics() {
		return
	}

	uri := r.URL.Path
	method := r.Method
	codeStr := ToStr(code)
	duration := float64(GetInterval(ts))
	reqSize := float64(r.ContentLength)
	var mode string

	if OtelMetricsPush() {
		mode = "PUSH"
		t := GetTraceFromCtx(ctx)
		if !t.IsValid() {
			return
		}

		attrs := metric.WithAttributes(
			attribute.String("trace_id", t.TraceID.String()),
			attribute.String("span_id", t.SpanID.String()),
			attribute.String("http_uri", uri),
			attribute.String("http_method", method),
			attribute.String("http_status_code", codeStr),
		)

		counter := getAPIReqTotal(ctx)
		if counter != nil {
			counter.Add(ctx, 1, attrs)
		}
		hist := getAPIRespDur(ctx)
		if hist != nil {
			hist.Record(ctx, duration, attrs)
		}
		hist = getAPIReqSize(ctx)
		if hist != nil {
			hist.Record(ctx, reqSize, attrs)
		}
		hist = getAPIRespSize(ctx)
		if hist != nil {
			hist.Record(ctx, float64(respSize), attrs)
		}
	} else {
		mode = "PULL"
		apiRequestTotalCounterVec.WithLabelValues(uri, method, codeStr).Inc()
		apiRequestDurationHistVec.WithLabelValues(uri, method).Observe(duration)
		apiRequestSizeHistVec.WithLabelValues(uri, method).Observe(reqSize)
		apiResponseSizeHistVec.WithLabelValues(uri, method, codeStr).Observe(float64(respSize))
	}
	logger.WithContext(ctx).Infof("Add service metrics in %s mode: uri = %s, method = %s, code = %d, duration = %f, reqSize = %f, respSize = %d",
		mode, uri, method, code, duration, reqSize, respSize)
}

func sendMetrics(ctx context.Context) {
	if ctx == nil {
		return
	}

	if !WithOtelMetrics() {
		return
	}

	t := GetTraceFromCtx(ctx)
	if !t.IsValid() {
		return
	}

	meter := getMeter()
	if meter == nil {
		return
	}

	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		logger.Warnf("Failed to get gopsutil process info: %v", err)
		return
	}

	// 参数是"指标命名空间"（自定义，用于分类）
	attrs := metric.WithAttributes(
		attribute.String("trace_id", t.TraceID.String()),
		attribute.String("span_id", t.SpanID.String()),
	)

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		logger.Warnf("Failed to get gopsutil CPU percent: %v", err)
		return
	}
	cpuPercentGauge, err := meter.Float64Gauge(
		GetProjName()+".cpu.percent",
		metric.WithDescription("CPU 使用率 (百分比, 0-100)"),
		metric.WithUnit("%"),
	)
	if err != nil {
		logger.Warnf("创建 CPU 指标失败: %v", err)
		return
	}
	cpuInfos, err := cpu.Info()
	totalCores := 0
	for _, info := range cpuInfos {
		totalCores += int(info.Cores)
	}
	if err != nil {
		logger.Warnf("Failed to get CPU info: %v", err)
		return
	}
	cpuPercent = cpuPercent * 100 / (float64)(totalCores)
	cpuPercentGauge.Record(ctx, cpuPercent, attrs)

	memPercent, err := p.MemoryPercent()
	if err != nil {
		logger.Warnf("Failed to get memory percent: %v", err)
		return
	}
	memPercentGauge, err := meter.Float64Gauge(
		GetProjName()+".mem.percent",
		metric.WithDescription("内存使用率 (百分比, 0-100)"),
		metric.WithUnit("%"),
	)
	if err != nil {
		logger.Warnf("创建内存指标失败: %v", err)
		return
	}
	memPercentGauge.Record(ctx, (float64)(memPercent*100), attrs)

	memInfo, err := p.MemoryInfo()
	if err != nil {
		logger.Warnf("Failed to get memory info: %v", err)
		return
	}
	memRssGauge, err := meter.Float64Gauge(
		GetProjName()+".mem.rss",
		metric.WithDescription("内存RSS使用量 (MB)"),
		metric.WithUnit("MiB"),
	)
	if err != nil {
		logger.Warnf("创建内存指标失败: %v", err)
		return
	}
	memRssGauge.Record(ctx, (float64)(memInfo.RSS)/1024/1024, attrs)

	cpuHist, err := meter.Float64Histogram(
		GetProjName()+".cpu.percent.hist",
		metric.WithDescription("CPU 使用率分布直方图 (百分比, 0-100)"),
		metric.WithUnit("%"),
		metric.WithExplicitBucketBoundaries(0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100),
	)
	if err != nil {
		logger.Warnf("创建 CPU直方图指标失败: %v", err)
		return
	}
	cpuHist.Record(ctx, cpuPercent*100)

}
