package admin

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplThirdPrtyApp         base.TplName = "admin/third-party-app/list"
)

func GetThirdPartyAppPage(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminThirdPartyApp"] = true
	ctx.HTML(200, tplThirdPrtyApp)
}
