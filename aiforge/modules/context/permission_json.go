package context

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"gitea.com/macaron/macaron"
)

func RequireRepoReaderJson(unitType models.UnitType) macaron.Handler {
	return func(ctx *Context) {
		if !ctx.Repo.CanRead(unitType) {
			if log.IsTrace() {
				if ctx.IsSigned {
					log.Trace("Permission Denied: User %-v cannot read %-v in Repo %-v\n"+
						"User in Repo has Permissions: %-+v",
						ctx.User,
						unitType,
						ctx.Repo.Repository,
						ctx.Repo.Permission)
				} else {
					log.Trace("Permission Denied: Anonymous user cannot read %-v in Repo %-v\n"+
						"Anonymous user in Repo has Permissions: %-+v",
						unitType,
						ctx.Repo.Repository,
						ctx.Repo.Permission)
				}
			}
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("error.no_right")))
			return
		}
	}
}

func RequireRepoWriterJson(unitType models.UnitType) macaron.Handler {
	return func(ctx *Context) {
		if !ctx.Repo.CanWrite(unitType) {
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("error.no_right")))
			return
		}
	}
}
