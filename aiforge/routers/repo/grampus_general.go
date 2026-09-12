package repo

import (
	"net/http"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplGrampusGeneralIndex base.TplName = "repo/grampus/general/list"
	tplGrampusGeneralShow  base.TplName = "repo/grampus/general/show"
	tplGrampusGeneralNew   base.TplName = "repo/grampus/general/new"
)

func GrampusGeneralNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusGeneralNew)
}

func GrampusGeneralShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusGeneralShow)
}

func GrampusGeneralIndex(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusGeneralIndex)
}
