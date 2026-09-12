package context

import (
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/setting"
	"gitea.com/macaron/macaron"
)

func GuestHandler(filterRepo bool) macaron.Handler {

	return func(ctx *Context) {

		if !ctx.IsSigned {
			var repoInWhiteList bool = false
			if filterRepo {
				userName := ctx.Params(":username")
				repoName := ctx.Params(":reponame")

				if userName != "" && repoName != "" {
					for _, v := range setting.RateLimitConfig.WhiteList {
						if v == userName+"/"+repoName {
							repoInWhiteList = true
						}
					}
				}
			}

			if !repoInWhiteList {

				ip := ctx.RemoteAddr()
				if setting.CustomRealUrlHeader != "" {
					addr := ctx.Req.Header.Get(setting.CustomRealUrlHeader)
					if len(addr) > 0 {
						ip = addr
					}
				}

				if ip != "" {
					guest := auth.GetGuest()
					if guest.Enabled {

						if guest.IsInBlackList(ip) {
							guest.ResetMaxFailedCount(ip)
							ctx.Error(403)
							return
						} else if !guest.IsInWhiteList(ip) {
							if !guest.IsInTempList(ip) {
								guest.AddTempList(ip)
								guest.AddMaxFailedCount(ip)
							}
							if guest.GetMaxFailedCount(ip) > guest.MaxFailedCount {
								guest.AddBlackList(ip)
								guest.ResetMaxFailedCount(ip)
								ctx.Error(403)
								return
							}

							ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
							ctx.Redirect(setting.AppSubURL + "/user/guest_verify")
							return

						} else {
							guest.ResetMaxFailedCount(ip)
						}
					}
				}

			}

		}
		ctx.Next()
	}
}
