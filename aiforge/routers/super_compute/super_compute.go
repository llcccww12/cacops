package super_compute

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplAPPList  base.TplName = "repo/supercompute/index"
	tplJobList  base.TplName = "repo/supercompute/list"
	tplCreateUI base.TplName = "repo/supercompute/create"
	tplDetail   base.TplName = "repo/supercompute/detail"
)

func GetAPPList(ctx *context.Context) {
	ctx.Data["PageIsSuperCompute"] = true
	ctx.HTML(200, tplAPPList)
}

func GetJobList(ctx *context.Context) {
	ctx.Data["PageIsSuperCompute"] = true
	ctx.HTML(200, tplJobList)
}

func CreateUI(ctx *context.Context) {
	ctx.Data["PageIsSuperCompute"] = true
	ctx.HTML(200, tplCreateUI)
}

func GetDetailUI(ctx *context.Context) {
	ctx.Data["PageIsSuperCompute"] = true
	ctx.HTML(200, tplDetail)
}
