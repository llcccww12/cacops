package setting

import "strings"

var ScreenMap = struct {
	ShowData      bool
	MinValue      int
	MaxValue      int
	ExcludeCenter []string
}{}

var IPInfo = struct {
	Host  string
	Token string
}{}

func NewScreenMapConfig() {
	sec := Cfg.Section("Screen")
	ScreenMap.ShowData = sec.Key("ShowData").MustBool(false)
	ScreenMap.MinValue = sec.Key("MinValue").MustInt(130)
	ScreenMap.MaxValue = sec.Key("MaxValue").MustInt(190)
	ScreenMap.ExcludeCenter = []string{}
	excludeCenterStr := sec.Key("ExcludeCenter").MustString("")
	if excludeCenterStr != "" {
		ScreenMap.ExcludeCenter = strings.Split(excludeCenterStr, ",")
	}

	sec = Cfg.Section("IPInfo")

	IPInfo.Host = sec.Key("Host").MustString("https://ipinfo.io")
	IPInfo.Token = sec.Key("Token").MustString("df2b002afe582a")
}
