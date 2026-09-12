package setting

var AimConfig *Aim

type Aim struct {
	Enabled    bool
	EnvHost    string
	DetailUrl  string
	CompareUrl string
}

func iniAimConfig() {
	sec := Cfg.Section("aim")
	AimConfig = &Aim{
		Enabled:    sec.Key("Enabled").MustBool(true),
		EnvHost:    sec.Key("Env_Host").MustString(""),
		DetailUrl:  sec.Key("Detail_Url").MustString(""),
		CompareUrl: sec.Key("Compare_Url").MustString(""),
	}
}
