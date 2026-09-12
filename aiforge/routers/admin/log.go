package admin

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplLogOperation        base.TplName = "admin/log/operation"
)

func GetLogPage(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminLog"] = true
	ctx.Data["PageIsAdminLogOperation"] = true
	ctx.HTML(200, tplLogOperation)
}