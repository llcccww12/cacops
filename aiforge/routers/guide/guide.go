package guide

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplCreateDataset base.TplName = "repo/guide/create_dataset"
	tplCreateModel   base.TplName = "repo/guide/create_model"
)

func GetCreateDataset(ctx *context.Context) {
	ctx.HTML(200, tplCreateDataset)
}

func GetCreateModel(ctx *context.Context) {
	ctx.HTML(200, tplCreateModel)
}
