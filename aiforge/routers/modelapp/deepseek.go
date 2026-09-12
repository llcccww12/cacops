package modelapp

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

var modelDeepSeekChatPageTpl base.TplName = "model/deepseek/index"

func DeepSeekChatPage(ctx *context.Context) {
	ctx.HTML(200, modelDeepSeekChatPageTpl)
}
