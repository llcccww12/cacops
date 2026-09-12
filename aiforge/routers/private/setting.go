package private

import (
	"code.gitea.io/gitea/modules/setting"
	"gitea.com/macaron/macaron"
)

func RefreshSetting(ctx *macaron.Context) {

	setting.Cfg.Reload()
	setting.NewScreenMapConfig()
	setting.GetGrampusConfig()
	setting.GetDiffusionAvatarConfig()
	setting.GetModelartsConfig()
	setting.GetModelartsCDConfig()
	ctx.PlainText(200, []byte("success"))

}
