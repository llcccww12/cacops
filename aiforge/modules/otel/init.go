package otel

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	logger "code.gitea.io/gitea/modules/log"
	otelpyroscope "github.com/grafana/otel-profiling-go"
	pyroscope "github.com/grafana/pyroscope-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.38.0"
	t "go.opentelemetry.io/otel/trace"
)

// 全局变量：追踪器和日志器
var (
	tracerProvider *trace.TracerProvider
	loggerProvider *log.LoggerProvider
	meterProvider  *metric.MeterProvider
	tracer         = otel.Tracer("openi") // 自定义追踪器名称

	profiler *pyroscope.Profiler
)

func GetTracer() t.Tracer {
	return tracer
}

func GetTraceProvider() *trace.TracerProvider {
	return tracerProvider
}

func GetLogProvider() *log.LoggerProvider {
	return loggerProvider
}

func GetMetricsProvider() *metric.MeterProvider {
	return meterProvider
}

func getOtelTraceOptions() []otlptracehttp.Option {
	opts := make([]otlptracehttp.Option, 0)
	opts = append(opts, otlptracehttp.WithEndpoint(GetOtelEndpoint()))
	if WithOtelInsecure() {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	return opts
}

func getOtelLogOptions() []otlploghttp.Option {
	opts := make([]otlploghttp.Option, 0)
	opts = append(opts, otlploghttp.WithEndpoint(GetOtelEndpoint()))
	if WithOtelInsecure() {
		opts = append(opts, otlploghttp.WithInsecure())
	}
	return opts
}

func getOtelMetricOptions() []otlpmetrichttp.Option {
	opts := make([]otlpmetrichttp.Option, 0)
	opts = append(opts, otlpmetrichttp.WithEndpoint(GetOtelEndpoint()))
	opts = append(opts, otlpmetrichttp.WithTimeout(GetOtelMetricsInterval()))
	if WithOtelInsecure() {
		opts = append(opts, otlpmetrichttp.WithInsecure())
	}
	return opts
}

// InitOTEL 初始化OpenTelemetry：包含追踪（Tempo）和日志（Loki）
func InitOTEL() {
	var err error = nil
	defer func() {
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to init OTEL SDK: %v", err))
		} else {
			logger.Info("Successfully init OTEL SDK")
		}
	}()
	ctx := context.Background()
	// 1. 配置资源（服务元数据）
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(GetProjName()),                          // 服务名（必填）
			attribute.Key("service.display.name").String(GetServiceDisplayName()), // 服务中文展示名
		),
		resource.WithTelemetrySDK(),
	)
	if err != nil {
		return
	}

	// 2. 初始化追踪导出器（对接Tempo）
	traceExporter, err := otlptracehttp.New(ctx, getOtelTraceOptions()...)
	if err != nil {
		return
	}

	// 3. 配置TracerProvider（追踪器）
	tracerProvider = trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithBatchTimeout(time.Second*5), // 批量导出超时
		),
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()), // 采样率（开发环境全采样）
	)

	// 配置 Pyroscope OTEL Profiling 性能剖析（可选）
	if WithOtelProfiling() {
		// 用 otel-profiling-go 包装原生 TracerProvider，实现 Span 与 Profiling 数据的绑定
		wrappedTP := otelpyroscope.NewTracerProvider(tracerProvider)
		// 将包装后的 TracerProvider 设置为 OTEL 全局 TracerProvider
		otel.SetTracerProvider(wrappedTP)

		profiler, err = pyroscope.Start(pyroscope.Config{
			ApplicationName: GetProjName(),      // 必须与 OTEL 的 service.name 一致
			ServerAddress:   GetPyroscopeAddr(), // Pyroscope 服务地址
			ProfileTypes: []pyroscope.ProfileType{ // 要采集的剖析类型
				pyroscope.ProfileCPU,           // CPU 剖析
				pyroscope.ProfileInuseObjects,  // 内存（使用中对象）剖析
				pyroscope.ProfileAllocObjects,  // 内存（堆）剖析
				pyroscope.ProfileInuseSpace,    // 内存（使用中空间）剖析
				pyroscope.ProfileAllocSpace,    // 内存（分配空间）剖析
				pyroscope.ProfileGoroutines,    // Goroutine 剖析
				pyroscope.ProfileMutexCount,    // 锁竞争剖析
				pyroscope.ProfileMutexDuration, // 锁竞争剖析
				pyroscope.ProfileBlockCount,    // 阻塞数剖析
				pyroscope.ProfileBlockDuration, // 阻塞时间剖析
			},
			Logger: logrus.StandardLogger(), // 使用 logrus 作为日志记录器
		})
		if err != nil {
			return
		}
	} else {
		otel.SetTracerProvider(tracerProvider)
	}

	// 设置上下文传播器（跨服务追踪）
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// 4. 初始化日志导出器（对接Loki）
	logExporter, err := otlploghttp.New(ctx, getOtelLogOptions()...)
	if err != nil {
		return
	}

	// 5. 配置LoggerProvider（日志器）
	loggerProvider = log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
		log.WithResource(res),
	)
	global.SetLoggerProvider(loggerProvider)

	// 6. 创建 Prometheus 导出器（核心：将 OTel 指标转为 Prometheus 格式）
	stdExporter, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(), // 关键：格式化输出，避免乱码
	)
	if err == nil {
		logger.Info("Initialized OTEL STDOUT Metric Exporter for debugging purpose: %v", stdExporter)
	} else {
		logger.Error("Failed to initialize OTEL STDOUT Metric Exporter: %v", err)
	}
	meterExporter, err := otlpmetrichttp.New(ctx, getOtelMetricOptions()...)
	if err != nil {
		return
	}

	// view := metric.NewView(
	// 	// 第一个参数：Instrument（匹配条件）- 匹配指定名称和类型的直方图
	// 	metric.Instrument{
	// 	  Name:  "dolphin.cpu.percent.hist", // 精准匹配业务直方图名称
	// 	  Kind:  metric.InstrumentKindHistogram,
	// 	},
	// 	metric.Stream{
	// 	  Aggregation: metric.AggregationExplicitBucketHistogram{
	// 		Boundaries: []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
	// 		NoMinMax: true,
	// 	  },
	// 	},
	//   )

	// 7. 配置MeterProvider（指标器）
	meterProvider = metric.NewMeterProvider(
		metric.WithResource(res), // 关联资源（元数据）
		metric.WithReader(metric.NewPeriodicReader(
			meterExporter, // 绑定 OTLP HTTP 导出器（往 Collector 上报）
			metric.WithInterval(GetOtelMetricsInterval()), // 上报间隔（默认 10s，可按需调整为 5s/60s）
		)),
		metric.WithExemplarFilter(exemplar.TraceBasedFilter), // 启用示例过滤器（可选，按需配置）
		//metric.WithView(view),
	)

	// 8. 设置OTEL SDK全局错误处理器
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		logger.Error("OTEL SDK error: %v", err)
	}))

	// 9. 配置Logrus日志库的OTEL Hook（关联OTEL上下文）
	if WithOtelLogrus() {
		//创建并注册 OTel Logrus Hook 实例，指定转发的日志级别
		logrus.AddHook(NewOtelLogrusHook(nil))
	} else if WithOtelSlog() {
		// 初始化slog日志器（关联OTEL上下文）
		slog.SetDefault(slog.New(NewOtelSlogHandler(os.Stdout, slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))
	}
}

// ShutdownOTEL 关闭OTEL资源，确保数据刷出
func ShutdownOTEL() {
	withErr := false
	ctx := context.Background()
	// 关闭追踪器
	if tracerProvider != nil {
		err := tracerProvider.Shutdown(ctx)
		if err != nil {
			logger.Error("Failed to shutdown OTEL TracerProvider: %v", err)
			withErr = true
		}
	}

	// 关闭日志器
	if loggerProvider != nil {
		err := loggerProvider.Shutdown(ctx)
		if err != nil {
			logger.Error("Failed to shutdown OTEL LoggerProvider: %v", err)
			withErr = true
		}
	}

	if meterProvider != nil {
		err := meterProvider.Shutdown(ctx)
		if err != nil {
			logger.Error("Failed to shutdown OTEL MeterProvider: %v", err)
			withErr = true
		}
	}

	if profiler != nil {
		err := profiler.Stop()
		if err != nil {
			logger.Error("Failed to shutdown OTEL Profiler: %v", err)
			withErr = true
		}
	}

	if withErr {
		logger.Error("Failed to shutdown OTEL SDK")
	} else {
		logger.Info("Successfully shutdown OTEL SDK")
	}
}
