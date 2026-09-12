package setting

type Oauth2Config struct {
	Provider     string
	ClientId     string
	ClientSecret string
	AuthUrl      string
	IconUrl      string
	ProfileUrl   string
	TokenUrl     string
}

var EduCoderConfig = Oauth2Config{}

var GitlinkConfig = Oauth2Config{}
var OsredmConfig = Oauth2Config{}

func getOauth2Config() {
	getEduCoderConfig()
	getGitlinkConfig()
	getOsredmConfig()
}

func getEduCoderConfig() {
	sec := Cfg.Section("auth.educoder")
	EduCoderConfig.Provider = sec.Key("provider").MustString("educoder")
	EduCoderConfig.AuthUrl = sec.Key("auth_url").MustString("")
	EduCoderConfig.TokenUrl = sec.Key("token_url").MustString("")
	EduCoderConfig.ProfileUrl = sec.Key("profile_url").MustString("")
}

func getGitlinkConfig() {
	sec := Cfg.Section("auth.gitlink")
	GitlinkConfig.Provider = sec.Key("provider").MustString("gitlink")
	GitlinkConfig.AuthUrl = sec.Key("auth_url").MustString("")
	GitlinkConfig.TokenUrl = sec.Key("token_url").MustString("")
	GitlinkConfig.ProfileUrl = sec.Key("profile_url").MustString("")
}

func getOsredmConfig() {
	sec := Cfg.Section("auth.osredm")
	OsredmConfig.Provider = sec.Key("provider").MustString("osredm")
	OsredmConfig.AuthUrl = sec.Key("auth_url").MustString("")
	OsredmConfig.TokenUrl = sec.Key("token_url").MustString("")
	OsredmConfig.ProfileUrl = sec.Key("profile_url").MustString("")
}
