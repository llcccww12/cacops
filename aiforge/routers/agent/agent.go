package agent

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

var agentPortalTpl base.TplName = "agent/portal"

func AgentPortalUI(ctx *context.Context) {
	ctx.HTML(200, agentPortalTpl)
}