package repo

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	pipelineIndex    base.TplName = "repo/pipeline/list"
	pipelineMenu     base.TplName = "pipeline/list"
	pipelineTplIndex base.TplName = "repo/pipeline/template"
	comparisonTplIndex base.TplName = "repo/pipeline/comparison"
)

func PipelineIndex(ctx *context.Context) {
	ctx.Data["PageIsPipeline"] = true
	ctx.HTML(200, pipelineIndex)
}

func PipelineMenu(ctx *context.Context) {
	ctx.HTML(200, pipelineMenu)
}

func TemplateIndex(ctx *context.Context) {
	ctx.Data["PageIsPipelineTemplate"] = true
	ctx.HTML(200, pipelineTplIndex)
}

func ComparisonIndex(ctx *context.Context) {
	ctx.Data["PageIsComparisonTemplate"] = true
	ctx.HTML(200, comparisonTplIndex)
}
