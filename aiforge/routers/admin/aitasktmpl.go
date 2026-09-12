package admin

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplAiTaskTmpl     base.TplName = "admin/aitasktmpl/list"
)

func AiTaskTmpl(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminAiTaskTmpl"] = true
	ctx.HTML(200, tplAiTaskTmpl)
}
