package repo

import (
	"net/http"

	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplGrampusOnlineInferIndex base.TplName = "repo/grampus/onlineinfer/list"
	tplGrampusOnlineInferShow  base.TplName = "repo/grampus/onlineinfer/show"
	tplGrampusOnlineInferNew   base.TplName = "repo/grampus/onlineinfer/new"
)

func GrampusOnlineInferNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusOnlineInferNew)
}

func GrampusOnlineInferShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusOnlineInferShow)
}

func GrampusOnlineInferIndex(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusOnlineInferIndex)

}
