package modelapp

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

var modelMindTpl base.TplName = "model/mind/index"

func MindPage(ctx *context.Context) {
	ctx.HTML(200, modelMindTpl)
}
