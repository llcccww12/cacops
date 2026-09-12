// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
	"encoding/base64"
	"net/http"
	"strings"

	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	"gitea.com/macaron/csrf"
	"gitea.com/macaron/macaron"

	marc_auth "github.com/go-macaron/auth"
)

// ToggleOptions contains required or check options
type ToggleOptions struct {
	SignInRequired             bool
	SignOutRequired            bool
	AdminRequired              bool
	DisableCSRF                bool
	BasicAuthRequired          bool
	OperationRequired          bool
	WechatAuthRequired         bool
	WechatAuthRequiredForAPI   bool
	WechatAuthRequiredStandard bool
}

// Toggle returns toggle options as middleware
func Toggle(options *ToggleOptions) macaron.Handler {
	return func(ctx *Context) {
		// Cannot view any page before installation.
		if !setting.InstallLock {
			ctx.Redirect(setting.AppSubURL + "/install")
			return
		}

		// Check prohibit login users.
		if ctx.IsSigned {
			if ctx.User.ProhibitLogin {
				log.Info("Failed authentication attempt for %s from %s", ctx.User.Name, ctx.RemoteAddr())
				ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
				ctx.HTML(200, "user/auth/prohibit_login")
				return
			} else if setting.PhoneService.Enabled && ctx.User.PhoneNumber == "" && ctx.Req.URL.Path != "/bindPhone" {
				ctx.Data["Title"] = ctx.Tr("phone.bind_phone")
				ctx.HTML(200, "user/auth/bind_phone")
				return
			}

			if ctx.User.MustChangePassword {
				if ctx.Req.URL.Path != "/user/settings/change_password" {
					ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
					ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/change_password"
					ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
					ctx.Redirect(setting.AppSubURL + "/user/settings/change_password")
					return
				}
			} else if ctx.Req.URL.Path == "/user/settings/change_password" {
				// make sure that the form cannot be accessed by users who don't need this
				ctx.Redirect(setting.AppSubURL + "/")
				return
			}

			if ctx.QueryBool("course") {
				ctx.Redirect(setting.AppSubURL + "/" + setting.Course.OrgName)
				return
			}
		}

		// Redirect to dashboard if user tries to visit any non-login page.
		if options.SignOutRequired && ctx.IsSigned && ctx.Req.URL.RequestURI() != "/" {
			redirectTo := ctx.Query("redirect_to")
			if len(redirectTo) > 0 {
				ctx.Redirect(redirectTo)
			} else {
				ctx.Redirect(setting.AppSubURL + "/")
			}
			return
		}

		if !options.SignOutRequired && !options.DisableCSRF && ctx.Req.Method == "POST" && !auth.IsAPIPath(ctx.Req.URL.Path) {
			csrf.Validate(ctx.Context, ctx.csrf)
			if ctx.Written() {
				return
			}
		}

		if options.SignInRequired {
			if !ctx.IsSigned {
				// Restrict API calls with error message.
				if auth.IsAPIPath(ctx.Req.URL.Path) {
					ctx.JSON(403, map[string]string{
						"message": "Only signed in user is allowed to call APIs.",
					})
					return
				}

				tempUrl := ctx.Req.URL.RequestURI()

				if strings.Contains(tempUrl, "action/star?") || strings.Contains(tempUrl, "action/watch?") {
					redirectForStarAndWatch(ctx, tempUrl)

				} else {
					ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
				}
				ctx.Redirect(setting.AppSubURL + "/user/login")
				return
			}
			if ctx.IsSigned && auth.IsAPIPath(ctx.Req.URL.Path) && ctx.IsBasicAuth {
				twofa, err := models.GetTwoFactorByUID(ctx.User.ID)
				if err != nil {
					if models.IsErrTwoFactorNotEnrolled(err) {
						return // No 2FA enrollment for this user
					}
					ctx.Error(500)
					return
				}
				otpHeader := ctx.Req.Header.Get("X-Gitea-OTP")
				ok, err := twofa.ValidateTOTP(otpHeader)
				if err != nil {
					ctx.Error(500)
					return
				}
				if !ok {
					ctx.JSON(403, map[string]string{
						"message": "Only signed in user is allowed to call APIs.",
					})
					return
				}
			}
		}

		if setting.WechatAuthSwitch && options.WechatAuthRequired {
			if !ctx.IsSigned {
				ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
				ctx.Redirect(setting.AppSubURL + "/user/login")
				return
			}
			if ctx.User.WechatOpenId == "" {
				ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
				ctx.Redirect(setting.AppSubURL + "/authentication/wechat/bind")
			}
		}

		if setting.WechatAuthSwitch && options.WechatAuthRequiredForAPI {
			if !ctx.IsSigned {
				ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
				ctx.Redirect(setting.AppSubURL + "/user/login")
				return
			}
			if ctx.User.WechatOpenId == "" {
				redirectUrl := ctx.Query("redirect_to")
				if redirectUrl == "" {
					redirectUrl = ctx.Req.URL.RequestURI()
				}
				ctx.SetCookie("redirect_to", setting.AppSubURL+redirectUrl, 0, setting.AppSubURL)
				ctx.JSON(200, map[string]string{
					"WechatRedirectUrl": setting.AppSubURL + "/authentication/wechat/bind",
				})
			}
		}
		if setting.WechatAuthSwitch && options.WechatAuthRequiredStandard {
			if !ctx.IsSigned {
				ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
				ctx.Redirect(setting.AppSubURL + "/user/login")
				return
			}
			if ctx.User.WechatOpenId == "" {
				redirectUrl := ctx.Query("redirect_to")
				if redirectUrl == "" {
					redirectUrl = ctx.Req.URL.RequestURI()
				}
				ctx.SetCookie("redirect_to", setting.AppSubURL+redirectUrl, 0, setting.AppSubURL)
				ctx.JSON(http.StatusOK, response.OuterTrBizError(response.WECHAT_NOT_BIND, ctx))
				return
			}
		}

		// Redirect to log in page if auto-signin info is provided and has not signed in.
		if !options.SignOutRequired && !ctx.IsSigned && !auth.IsAPIPath(ctx.Req.URL.Path) &&
			len(ctx.GetCookie(setting.CookieUserName)) > 0 {
			ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
			ctx.Redirect(setting.AppSubURL + "/user/login")
			return
		}

		if options.AdminRequired {
			if !ctx.User.IsAdmin {
				ctx.Error(403)
				return
			}
			ctx.Data["PageIsAdmin"] = true
		}

		if options.BasicAuthRequired {
			if !basicAuth(ctx) {
				basicUnauthorized(ctx.Resp)
				return
			}
		}

		if options.OperationRequired {
			if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_KANBANAdmin) {
				ctx.Error(403)
				return
			}
			ctx.Data["PageIsOperation"] = true
		}
	}
}

func redirectForStarAndWatch(ctx *Context, tempUrl string) {
	splits := strings.Split(tempUrl, "?")
	if len(splits) > 1 {
		redirectArguments := strings.Split(splits[1], "=")

		if len(redirectArguments) > 0 && redirectArguments[0] == "redirect_to" {
			ctx.SetCookie("redirect_to", setting.AppSubURL+strings.Replace(redirectArguments[1], "%2f", "/", -1), 0, setting.AppSubURL)
		}
	}
}

func basicAuth(ctx *Context) bool {
	var siteAuth = base64.StdEncoding.EncodeToString([]byte(setting.CBAuthUser + ":" + setting.CBAuthPassword))
	auth := ctx.Req.Header.Get("Authorization")

	if !marc_auth.SecureCompare(auth, "Basic "+siteAuth) {
		return false
	}

	return true

}

func basicUnauthorized(res http.ResponseWriter) {
	res.Header().Set("WWW-Authenticate", "Basic realm=\""+marc_auth.BasicRealm+"\"")
	http.Error(res, "Not Authorized", http.StatusUnauthorized)
}
