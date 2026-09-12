package setting

import "code.gitea.io/gitea/modules/log"

type Guest struct {
	Enabled         bool
	WhiteExpiration int // 白名单有效期
	BlackExpiration int // 黑名单有效期
	TempExpiration  int // 临时名单有效期
	MaxFailedCount  int
}

var (
	// Guest  verify info
	GuestInfo *Guest
)

func newGuestInfo() {
	sec := Cfg.Section("guest")
	// Check phone setting.

	GuestInfo = &Guest{
		Enabled:         sec.Key("ENABLED").MustBool(true),
		WhiteExpiration: sec.Key("WhiteExpiration").MustInt(86400),
		BlackExpiration: sec.Key("BlackExpiration").MustInt(86400),
		TempExpiration:  sec.Key("TempExpiration").MustInt(90),
		MaxFailedCount:  sec.Key("MaxFailedCount").MustInt(5),
	}

	log.Info("guest Service Enabled")
}
