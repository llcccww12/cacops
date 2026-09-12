package repo

import (
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplStorageManage base.TplName = "repo/storage/index"
)

func GetStorageManage(ctx *context.Context) {
	ctx.HTML(200, tplStorageManage)
}
