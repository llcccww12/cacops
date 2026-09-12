package modelbase

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

var modelbasePortalTpl base.TplName = "model/base/portal"

func ModelBasePortalUI(ctx *context.Context) {
	ctx.HTML(200, modelbasePortalTpl)
}
