package setting

import "strings"

var RateLimitConfig *RateLimit

type RateLimit struct {
	Enabled   bool
	LowRate   int
	LowBurst  int
	LowestRate   int
	LowestBurst  int
	HighRate  int
	HighBurst int
	WhiteList []string
}

func initRateLimitConfig() {
	sec := Cfg.Section("rate_limit")
	RateLimitConfig = &RateLimit{
		Enabled:   sec.Key("Enabled").MustBool(true),
		LowestRate:   sec.Key("LowestRate").MustInt(10),
		LowestBurst:  sec.Key("LowestBurst").MustInt(15),
		LowRate:   sec.Key("LowRate").MustInt(10),
		LowBurst:  sec.Key("LowBurst").MustInt(20),
		HighRate:  sec.Key("HighRate").MustInt(20),
		HighBurst: sec.Key("HighBurst").MustInt(30),
		WhiteList: strings.Split(sec.Key("WhiteList").MustString(""), ","),
	}

}
