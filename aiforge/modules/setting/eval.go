package setting

type Eval struct {
	Enabled   bool
	ServerUrl string
}

var (
	// MailService the global mailer
	EvalService *Eval
)

func newEvalService() {
	sec := Cfg.Section("eval")
	// Check mailer setting.
	if !sec.Key("Enabled").MustBool() {
		EvalService = &Eval{}
		return
	}

	EvalService = &Eval{
		Enabled:   true,
		ServerUrl: sec.Key("ServerUrl").MustString(""),
	}
}
