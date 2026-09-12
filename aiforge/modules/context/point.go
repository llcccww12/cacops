package context

import (
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/reward/point/account"
	"gitea.com/macaron/macaron"
)

// PointAccount returns a macaron to get request user's point account
func PointAccount() macaron.Handler {
	return func(ctx *Context) {
		a, err := account.GetAccount(ctx.User.ID)
		if err != nil {
			ctx.ServerError("GetPointAccount", err)
			return
		}
		ctx.Data["PointAccount"] = a
		ctx.Data["CloudBrainPaySwitch"] = setting.CloudBrainPaySwitch
		ctx.Next()
	}
}
