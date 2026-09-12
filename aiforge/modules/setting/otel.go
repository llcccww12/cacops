package setting

type Otel struct {
	Enabled            bool
	ProjectName        string
	ServiceDisplayName string
	Endpoint           string
	CacheTimeout       int
	Insecure           bool
	WrapHandler        bool
	WithLog            bool
	WithSlog           bool
	GlsKey             string
	Profiling          *Profiling
	Metrics            *OtelMetrics
}
type Profiling struct {
	Enabled bool
	Url     string
}
type OtelMetrics struct {
	Enabled   bool
	Namespace string
	Interval  int
	WithPush  bool
}

var OtelService *Otel

func InitOTEL() {
	OtelService = iniOtelConfig()
}

func iniOtelConfig() *Otel {

	sec := Cfg.Section("otel")
	if !sec.Key("enabled").MustBool() {
		return &Otel{
			Enabled: false,
		}
	} else {
		return &Otel{
			Enabled:            true,
			ProjectName:        sec.Key("project_name").MustString("openi"),
			ServiceDisplayName: sec.Key("service_display_name").MustString("openi ai service"),
			Endpoint:           sec.Key("endpoint").MustString(""),
			Insecure:           sec.Key("insecure").MustBool(true),
			CacheTimeout:       sec.Key("cach_timeout").MustInt(240),
			WrapHandler:        sec.Key("wrap_handler").MustBool(false),
			WithLog:            sec.Key("with_log").MustBool(false),
			WithSlog:           sec.Key("with_slog").MustBool(false),
			GlsKey:             sec.Key("gls_key").MustString("openi-trace-cxt"),
			Profiling: &Profiling{
				Enabled: sec.Key("profiling_enabled").MustBool(false),
				Url:     sec.Key("profiling_url").MustString("")},
			Metrics: &OtelMetrics{
				Enabled:   sec.Key("metrics_enabled").MustBool(false),
				Namespace: sec.Key("metrics_namespace").MustString("openi"),
				Interval:  sec.Key("metrics_interval").MustInt(10),
				WithPush:  sec.Key("metrics_with_push").MustBool(false)},
		}

	}

}
