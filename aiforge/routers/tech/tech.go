package tech

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplCreate base.TplName = "tech/create"
	tplTech   base.TplName = "tech/tech_view"
	tplRepo   base.TplName = "tech/repo_view"
	tplAdmin  base.TplName = "tech/admin_view"
	tplMine   base.TplName = "tech/my_view"
)

func Create(ctx *context.Context) {
	ctx.HTML(200, tplCreate)
}

func TechView(ctx *context.Context) {
	ctx.HTML(200, tplTech)
}
func RepoView(ctx *context.Context) {
	ctx.HTML(200, tplRepo)
}
func AdminView(ctx *context.Context) {
	ctx.HTML(200, tplAdmin)
}

func MyView(ctx *context.Context) {
	ctx.HTML(200, tplMine)
}
