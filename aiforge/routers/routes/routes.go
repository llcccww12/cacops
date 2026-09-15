// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"path"
	"text/template"
	"time"

	"code.gitea.io/gitea/routers/ai_task"
	"code.gitea.io/gitea/routers/pipeline"

	"code.gitea.io/gitea/routers/card_request"

	"code.gitea.io/gitea/routers/super_compute"

	"code.gitea.io/gitea/routers/tech"

	"code.gitea.io/gitea/routers/badge"
	"code.gitea.io/gitea/routers/reward/point"
	"code.gitea.io/gitea/routers/task"
	badge_service "code.gitea.io/gitea/services/badge"
	"code.gitea.io/gitea/services/reward"
	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/routers/agent"
	"code.gitea.io/gitea/routers/modelapp"
	"code.gitea.io/gitea/routers/modelbase"

	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/slideimage"

	"code.gitea.io/gitea/routers/image"

	"code.gitea.io/gitea/routers/authentication"

	"code.gitea.io/gitea/modules/cloudbrain"

	"code.gitea.io/gitea/routers/operation"
	"code.gitea.io/gitea/routers/private"

	"code.gitea.io/gitea/routers/secure"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/lfs"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/metrics"
	"code.gitea.io/gitea/modules/options"
	"code.gitea.io/gitea/modules/public"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/templates"
	"code.gitea.io/gitea/modules/validation"
	"code.gitea.io/gitea/routers"
	"code.gitea.io/gitea/routers/admin"
	apiv1 "code.gitea.io/gitea/routers/api/v1"
	"code.gitea.io/gitea/routers/dev"
	"code.gitea.io/gitea/routers/events"
	"code.gitea.io/gitea/routers/guide"
	"code.gitea.io/gitea/routers/org"
	"code.gitea.io/gitea/routers/repo"
	"code.gitea.io/gitea/routers/user"
	userSetting "code.gitea.io/gitea/routers/user/setting"
	"code.gitea.io/gitea/services/mailer"

	// to registers all internal adapters
	_ "code.gitea.io/gitea/modules/session"

	"gitea.com/macaron/binding"
	"gitea.com/macaron/cache"
	"gitea.com/macaron/captcha"
	"gitea.com/macaron/cors"
	"gitea.com/macaron/csrf"
	"gitea.com/macaron/gzip"
	"gitea.com/macaron/i18n"
	"gitea.com/macaron/macaron"
	"gitea.com/macaron/session"
	"gitea.com/macaron/toolbox"
	"github.com/prometheus/client_golang/prometheus"
	gouuid "github.com/satori/go.uuid"
	"github.com/tstranex/u2f"
)

type routerLoggerOptions struct {
	Ctx            *macaron.Context
	Identity       *string
	Start          *time.Time
	ResponseWriter *macaron.ResponseWriter
}

func setupAccessLogger(m *macaron.Macaron) {
	logger := log.GetLogger("access")

	logTemplate, _ := template.New("log").Parse(setting.AccessLogTemplate)
	m.Use(func(ctx *macaron.Context) {
		start := time.Now()
		ctx.Next()
		identity := "-"
		if val, ok := ctx.Data["SignedUserName"]; ok {
			if stringVal, ok := val.(string); ok && stringVal != "" {
				identity = stringVal
			}
		}
		rw := ctx.Resp.(macaron.ResponseWriter)

		buf := bytes.NewBuffer([]byte{})
		err := logTemplate.Execute(buf, routerLoggerOptions{
			Ctx:            ctx,
			Identity:       &identity,
			Start:          &start,
			ResponseWriter: &rw,
		})
		if err != nil {
			log.Error("Could not set up macaron access logger: %v", err.Error())
		}

		err = logger.SendLog(log.INFO, "", "", 0, buf.String(), ctx.Data["msgID"].(string), "")
		if err != nil {
			log.Error("Could not set up macaron access logger: %v", err.Error())
		}
	})
}

// RouterHandler is a macaron handler that will log the routing to the default gitea log
func RouterHandler(level log.Level) func(ctx *macaron.Context) {
	return func(ctx *macaron.Context) {
		start := time.Now()

		_ = log.GetLogger("router").Log(0, level, "Started %s %s for %s", log.ColoredMethod(ctx.Req.Method), ctx.Req.URL.RequestURI(), ctx.RemoteAddr())

		rw := ctx.Resp.(macaron.ResponseWriter)
		ctx.Next()

		status := rw.Status()
		_ = log.GetLogger("router").Log(0, level, "Completed %s %s %v %s in %v", log.ColoredMethod(ctx.Req.Method), ctx.Req.URL.RequestURI(), log.ColoredStatus(status), log.ColoredStatus(status, http.StatusText(rw.Status())), log.ColoredTime(time.Since(start)))
	}
}

// SetLogMsgID set msgID in Context
func SetLogMsgID() macaron.Handler {
	return func(ctx *macaron.Context) {
		start := time.Now()

		uuid := gouuid.NewV4().String()
		ctx.Data["MsgID"] = uuid

		log.Info("%s Started %s %s for %s", ctx.Data["SignedUserName"], log.ColoredMethod(ctx.Req.Method), ctx.Req.URL.RequestURI(), ctx.RemoteAddr(), ctx.Data["MsgID"])

		rw := ctx.Resp.(macaron.ResponseWriter)
		ctx.Next()

		status := rw.Status()
		log.Info("Completed %s %s %v %s in %v", log.ColoredMethod(ctx.Req.Method), ctx.Req.URL.RequestURI(), log.ColoredStatus(status), log.ColoredStatus(status, http.StatusText(rw.Status())), log.ColoredTime(time.Since(start)), ctx.Data["MsgID"])
	}
}

// NewMacaron initializes Macaron instance.
func NewMacaron() *macaron.Macaron {
	gob.Register(&u2f.Challenge{})
	var m *macaron.Macaron
	if setting.RedirectMacaronLog {
		loggerAsWriter := log.NewLoggerAsWriter("INFO", log.GetLogger("macaron"))
		m = macaron.NewWithLogger(loggerAsWriter)
		if !setting.DisableRouterLog && setting.RouterLogLevel != log.NONE {
			if log.GetLogger("router").GetLevel() <= setting.RouterLogLevel {
				m.Use(RouterHandler(setting.RouterLogLevel))
			}
		}
	} else {
		m = macaron.New()
		if !setting.DisableRouterLog {
			m.Use(macaron.Logger())
		}
	}
	//m.Use(SetLogMsgID())
	// Access Logger is similar to Router Log but more configurable and by default is more like the NCSA Common Log format
	if setting.EnableAccessLog {
		setupAccessLogger(m)
	}
	m.Use(macaron.Recovery())
	if setting.EnableGzip {
		m.Use(gzip.Middleware())
	}
	if setting.Protocol == setting.FCGI || setting.Protocol == setting.FCGIUnix {
		m.SetURLPrefix(setting.AppSubURL)
	}
	m.Use(public.Custom(
		&public.Options{
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: setting.StaticCacheTime,
		},
	))
	m.Use(public.Static(
		&public.Options{
			Directory:    path.Join(setting.StaticRootPath, "public"),
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: setting.StaticCacheTime,
		},
	))
	m.Use(public.CustomJson(
		&public.Options{
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: 0,
		},
	))
	m.Use(public.StaticHandler(
		setting.AvatarUploadPath,
		&public.Options{
			Prefix:       "avatars",
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: setting.StaticCacheTime,
		},
	))
	m.Use(public.StaticHandler(
		setting.IconUploadPath,
		&public.Options{
			Prefix:       "icons",
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: setting.StaticCacheTime,
		},
	))
	m.Use(public.StaticHandler(
		setting.RepositoryAvatarUploadPath,
		&public.Options{
			Prefix:       "repo-avatars",
			SkipLogging:  setting.DisableRouterLog,
			ExpiresAfter: setting.StaticCacheTime,
		},
	))

	m.Use(templates.HTMLRenderer())
	mailer.InitMailRender(templates.Mailer())

	localeNames, err := options.Dir("locale")

	if err != nil {
		log.Fatal("Failed to list locale files: %v", err)
	}

	localFiles := make(map[string][]byte)

	for _, name := range localeNames {
		localFiles[name], err = options.Locale(name)

		if err != nil {
			log.Fatal("Failed to load %s locale file. %v", name, err)
		}
	}

	m.Use(i18n.I18n(i18n.Options{
		SubURL:       setting.AppSubURL,
		Files:        localFiles,
		Langs:        setting.Langs,
		Names:        setting.Names,
		DefaultLang:  "en-US",
		Redirect:     false,
		CookieDomain: setting.SessionConfig.Domain,
	}))
	m.Use(cache.Cacher(cache.Options{
		Adapter:       setting.CacheService.Adapter,
		AdapterConfig: setting.CacheService.Conn,
		Interval:      setting.CacheService.Interval,
	}))
	m.Use(captcha.Captchaer(captcha.Options{
		SubURL: setting.AppSubURL,
	}))
	m.Use(slideimage.SlideImager(slideimage.Options{
		SubURL:       setting.AppSubURL,
		SampleImages: setting.SlideImagesCount,
	}))
	m.Use(session.Sessioner(session.Options{
		Provider:       setting.SessionConfig.Provider,
		ProviderConfig: setting.SessionConfig.ProviderConfig,
		CookieName:     setting.SessionConfig.CookieName,
		CookiePath:     setting.SessionConfig.CookiePath,
		Gclifetime:     setting.SessionConfig.Gclifetime,
		Maxlifetime:    setting.SessionConfig.Maxlifetime,
		Secure:         setting.SessionConfig.Secure,
		Domain:         setting.SessionConfig.Domain,
	}))
	m.Use(auth.AutoLoginner())
	m.Use(csrf.Csrfer(csrf.Options{
		Secret:         setting.SecretKey,
		Cookie:         setting.CSRFCookieName,
		SetCookie:      true,
		Secure:         setting.SessionConfig.Secure,
		CookieHttpOnly: setting.CSRFCookieHTTPOnly,
		Header:         "X-Csrf-Token",
		CookieDomain:   setting.SessionConfig.Domain,
		CookiePath:     setting.AppSubURL,
	}))
	m.Use(toolbox.Toolboxer(m, toolbox.Options{
		HealthCheckFuncs: []*toolbox.HealthCheckFuncDesc{
			{
				Desc: "Database connection",
				Func: models.Ping,
			},
		},
		DisableDebug: !setting.EnablePprof,
	}))
	m.Use(context.Contexter())
	m.Use(SetLogMsgID())
	// OK we are now set-up enough to allow us to create a nicer recovery than
	// the default macaron recovery
	m.Use(context.Recovery())
	m.SetAutoHead(true)
	return m
}

// RegisterRoutes routes routes to Macaron
func RegisterRoutes(m *macaron.Macaron) {
	reqSignIn := context.Toggle(&context.ToggleOptions{SignInRequired: true})
	ignSignIn := context.Toggle(&context.ToggleOptions{SignInRequired: setting.Service.RequireSignInView})
	ignSignInAndCsrf := context.Toggle(&context.ToggleOptions{DisableCSRF: true})
	reqSignOut := context.Toggle(&context.ToggleOptions{SignOutRequired: true})
	reqBasicAuth := context.Toggle(&context.ToggleOptions{BasicAuthRequired: true, DisableCSRF: true})
	reqWechatBind := context.Toggle(&context.ToggleOptions{WechatAuthRequired: true})
	//reqWeChatStandard := context.Toggle(&context.ToggleOptions{WechatAuthRequiredStandard: true})

	bindIgnErr := binding.BindIgnErr
	validation.AddBindingRules()

	openIDSignInEnabled := func(ctx *context.Context) {
		if !setting.Service.EnableOpenIDSignIn {
			ctx.Error(403)
			return
		}
	}

	openIDSignUpEnabled := func(ctx *context.Context) {
		if !setting.Service.EnableOpenIDSignUp {
			ctx.Error(403)
			return
		}
	}

	reqMilestonesDashboardPageEnabled := func(ctx *context.Context) {
		if !setting.Service.ShowMilestonesDashboardPage {
			ctx.Error(403)
			return
		}
	}

	m.Use(user.GetNotificationCount)
	m.Use(func(ctx *context.Context) {
		ctx.Data["UnitWikiGlobalDisabled"] = models.UnitTypeWiki.UnitGlobalDisabled()
		ctx.Data["UnitIssuesGlobalDisabled"] = models.UnitTypeIssues.UnitGlobalDisabled()
		ctx.Data["UnitPullsGlobalDisabled"] = models.UnitTypePullRequests.UnitGlobalDisabled()

	})

	// FIXME: not all routes need go through same middlewares.
	// Especially some AJAX requests, we can reduce middleware number to improve performance.
	// Routers.
	// for health check
	m.Head("/", func() string {
		return ""
	})
	m.Get("/", routers.Home)
	m.Get("/dashboard", routers.OverView)
	m.Get("/repostemp", routers.Dashboard1)
	m.Get("/^:type(issues|pulls)$", reqSignIn, user.Issues1)
	m.Group("/repositories", func() {
		m.Get("", routers.Dashboard)
		m.Get("/^:type(issues|pulls)$", reqSignIn, user.Issues)
	})
	go routers.SocketManager.Run()
	go task.RunTask()
	go reward.AcceptStatusChangeAction()

	if otel.WithOtel() {
		otel.InitOTEL()
	}

	m.Get("/action/notification", routers.ActionNotification)
	m.Get("/action/latest_actions", routers.LatestActions)
	m.Get("/recommend/home", routers.RecommendHomeInfo)
	m.Get("/dashboard/invitation", routers.GetMapInfo)
	m.Get("/resource_desc", routers.HomeResoruceDesc)
	//m.Get("/recommend/org", routers.RecommendOrgFromPromote)
	//m.Get("/recommend/repo", routers.RecommendRepoFromPromote)
	m.Get("/recommend/userrank/:index", routers.GetUserRankFromPromote)
	//m.Get("/recommend/imageinfo", routers.GetImageInfoFromPromote)
	m.Post("/all/search/", routers.Search)
	m.Get("/all/search/", routers.EmptySearch)
	m.Get("/all/dosearch/", routers.SearchApi)
	m.Post("/user/login/kanban", user.SignInPostAPI, reqSignOut)
	m.Get("/home/term", routers.HomeTerm)
	m.Get("/home/annual_privacy", routers.HomeAnnual)
	m.Get("/home/model_privacy", routers.HomeWenxin)
	m.Get("/home/notice", routers.HomeNoticeTmpl)
	m.Get("/home/privacy", routers.HomePrivacy)
	m.Get("/home/auth_term", routers.HomeAuthTerm)

	m.Post("/user/saveOtherInfo", user.SaveUserOtherInfo, reqSignIn)

	m.Group("/modelsquare", func() {
		// m.Get("/main", repo.ModelSquareTmpl)
		// m.Get("/main_my", reqSignIn, repo.ModelSquareTmpl)
		m.Get("/main_query_data", repo.ModelSquareData)
		m.Put("/modify_model_collect", repo.ModelCollect)
		m.Get("/main_query_label", repo.QueryModelLabel)
		m.Post("/batch_delete_model", repo.BatchDeleteModel)

	}, ignSignIn)

	m.Group("/agent", func() {
		m.Get("", agent.AgentPortalUI)
		m.Get("/*", agent.AgentPortalUI)
	}, ignSignIn)

	m.Group("/modelbase", func() {
		m.Get("", modelbase.ModelBasePortalUI)
		m.Get("/*", modelbase.ModelBasePortalUI)
	}, ignSignIn)

	m.Group("/extension", func() {
		// m.Get("", modelapp.ModelMainPage)
		m.Get("/mind", ignSignIn, modelapp.MindPage)
		m.Get("/tuomin/upload", modelapp.ProcessImageUI)
		m.Post("/tuomin/upload", reqSignIn, modelapp.ProcessImage)
		m.Get("/wenxin", modelapp.WenXinPage)
		m.Group("/modelexperience", func() {
			m.Get("/sd", reqSignIn, reqWechatBind, modelapp.ModelExperienceSd)
			m.Get("/chat", reqSignIn, reqWechatBind, modelapp.LLMChatPage)
			m.Get("/tts", reqSignIn, reqWechatBind, modelapp.TtsPage)
			m.Get("/create", reqSignIn, reqWechatBind, modelapp.ModelExperienceCreate)
			m.Get("/queryall", reqSignIn, modelapp.QueryAllModelExperience)
		})
		m.Get("/deepseek/chat", reqSignIn, reqWechatBind, modelapp.DeepSeekChatPage)
		m.Get("/wenxin/paint_new", reqSignIn, modelapp.WenXinPaintNew)
		m.Get("/sd/paint_new", reqSignIn, modelapp.SdPaintNew)
		m.Get("/wenxin/query_paint_result", reqSignIn, modelapp.QueryWenXinPaintResult)
		m.Get("/wenxin/query_paint_image", reqSignIn, modelapp.QueryWenXinPaintById)
		m.Group("/modelbase", func() {
			m.Get("", modelapp.ModelBaseUI)
			// m.Group("/pangufinetune", func() {
			// 	m.Get("", modelapp.PanguFinetuneUI)
			m.Get("/create", reqSignIn, reqWechatBind, modelapp.PanguFinetuneCreateUI)
			// 	m.Get("/inference", reqSignIn, modelapp.PanguInferenceUI)
			// })
		}, ignSignIn)
	})

	m.Group("/computingpower", func() {
		m.Get("/demand", routers.ComputingPowerDemand)
		m.Get("/domestic", routers.ComputingPowerDomestic)
	}, ignSignIn)

	operationReq := context.Toggle(&context.ToggleOptions{SignInRequired: true, OperationRequired: true})
	m.Group("/explore", func() {
		m.Get("", func(ctx *context.Context) {
			ctx.Redirect(setting.AppSubURL + "/explore/repos")
		})

		m.Group("/repos", func() {
			//m.Get("", routers.ExploreRepos)
			m.Get("", routers.GetRepoSearchPage)
			m.Group("/square", func() {
				m.Get("", routers.GetRepoSquarePage)
				m.Get("/tab", routers.RepoSquare)
				m.Get("/active-user", routers.ActiveUser)
				m.Get("/active-org", routers.ActiveOrg)
			})

			m.Get("/search", routers.RepoFind)
		})
		m.Group("/datasets", func() {
			m.Get("", routers.ExploreDatasetsUI)
			// m.Get("/:id", routers.ExploreDatasetsDetailUI)
			// m.Get("/:ownername/:name", routers.ExploreDatasetsDetailUI)
		})
		m.Group("/models", func() {
			m.Get("", routers.ExploreModelsUI)
			// m.Get("/:ownername/:name", routers.ExploreModelsDetailUI)
			// m.Get("/model_migrating", repo.ModelMigratingTmpl)
		})
		m.Get("/datasets_my", reqSignIn, routers.ExploreDatasetsUI)
		m.Get("/models_my", reqSignIn, routers.ExploreModelsUI)
		m.Get("/public_datasets", routers.ExploreDatasets)
		m.Get("/my_datasets", reqSignIn, routers.ExploreMyDatasets)
		m.Get("/my_favorite_datasets", reqSignIn, routers.ExploreFavoriteDatasets)
		m.Get("/users", routers.ExploreUsers)
		m.Get("/organizations", routers.ExploreOrganizations)
		m.Get("/org_my", routers.ExploreOrgMy)
		m.Get("/code", routers.ExploreCode)
		m.Get("/images", routers.ExploreImages)
		m.Get("/images_my", reqSignIn, routers.ExploreImages)
		m.Get("/center_map", reqSignIn, routers.CenterMapUI)
		m.Get("/c2net_map", routers.C2NetUI)
		m.Group("/card_request", func() {
			m.Get("/creation/required", card_request.GetCreationInfo)
			m.Get("/list", card_request.GetCardRequestList)
			m.Get("/resource/list", card_request.GetResourceList)
		}, ignSignIn)

		m.Group("/card_request", func() {
			m.Post("/create", binding.Bind(structs.CardReq{}), card_request.CreateCardRequest)
			m.Get("/my_list", card_request.GetMyCardRequestList)
			m.Get("/admin_list", operationReq, card_request.GetAdminCardRequestList)
			m.Get("/specification/list", operationReq, admin.GetAllResourceSpecificationList)
			m.Get("/resources/queue/centers", operationReq, admin.GetResourceAiCenters)
			m.Get("/resources/queue/codes", operationReq, admin.GetResourceQueueCodes)
			m.Put("/update/:id", binding.Bind(structs.CardReq{}), card_request.UpdateCardRequest)
			m.Put("/admin/update/:id", operationReq, bindIgnErr(structs.CardReq{}), card_request.UpdateCardRequestAndSpec)

		}, reqSignIn)

	}, ignSignIn)
	m.Combo("/install", routers.InstallInit).Get(routers.Install).
		Post(bindIgnErr(auth.InstallForm{}), routers.InstallPost)

	m.Get("/milestones", reqSignIn, reqMilestonesDashboardPageEnabled, user.Milestones)
	m.Group("/cloudbrains", func() {
		m.Get("", reqSignIn, user.Cloudbrains)
		m.Get("/create", reqSignIn, user.CloudbrainCreate)
		m.Get("/detail/:jobid", reqSignIn, user.CloudbrainDetail)
	})

	m.Group("/ai_task_tmpl", func() {
		m.Get("/list", ai_task.GetAiTaskTmplListTpl)
		m.Get("/list_my", reqSignIn, ai_task.GetAiTaskTmplListMyTpl)
		m.Get("/create", reqSignIn, ai_task.GetAiTaskTmplCreateTpl)
		m.Get("/detail/:id", ai_task.GetAiTaskTmplDetailTpl)
		m.Get("/edit/:id", reqSignIn, ai_task.GetAiTaskTmplEditTpl)
	})
	m.Get("/datasets/detail/:ownername/:name", routers.ExploreDatasetsDetailUI)
	m.Get("/models/detail/:ownername/:name", routers.ExploreModelsDetailUI)
	m.Get("/models/detail/model_migrating", repo.ModelMigratingTmpl)
	m.Group("/ros-hmci", func() {
		m.Get("", routers.ExploreRosHmci)
		m.Get("/*", routers.ExploreRosHmci)
	})
	m.Group("/storages", func() {
		m.Get("", reqSignIn, repo.GetStorageManage)
	})

	// ***** START: User *****
	m.Group("/user", func() {
		m.Get("/login", user.SignIn)
		m.Get("/login/cloud_brain", user.SignInCloudBrain)
		m.Post("/login/cloud_brain", bindIgnErr(auth.SignInForm{}), user.SignInCloudBrainPost)
		m.Post("/login", bindIgnErr(auth.SignInForm{}), user.SignInPost)
		m.Get("/login/wechat", user.SignInWeChat)
		m.Get("/login/phone", user.SignInPhone)
		m.Post("/login/phone", bindIgnErr(auth.PhoneNumberCodeForm{}), user.SignInPhonePost)
		m.Get("/login/wechat/qrcode", user.GetLoginWechatQRCode)
		m.Post("/login/wechat", bindIgnErr(auth.WechatForm{}), user.LoginByWechat)
		m.Group("", func() {
			m.Combo("/login/openid").
				Get(user.SignInOpenID).
				Post(bindIgnErr(auth.SignInOpenIDForm{}), user.SignInOpenIDPost)
		}, openIDSignInEnabled)

		m.Group("/openid", func() {
			m.Combo("/connect").
				Get(user.ConnectOpenID).
				Post(bindIgnErr(auth.ConnectOpenIDForm{}), user.ConnectOpenIDPost)
			m.Group("/register", func() {
				m.Combo("").
					Get(user.RegisterOpenID, openIDSignUpEnabled).
					Post(bindIgnErr(auth.SignUpOpenIDForm{}), user.RegisterOpenIDPost)
			}, openIDSignUpEnabled)
		}, openIDSignInEnabled)
		m.Get("/sign_up", user.SignUp)
		m.Get("/sign_up_scancode", user.SignUpScanCode)
		m.Group("/sign_up/wechat", func() {
			m.Get("", user.GetSignUpWechatQRCode)
		})
		m.Post("/sign_up", bindIgnErr(auth.RegisterForm{}), user.SignUpPost)
		m.Group("/oauth2", func() {
			m.Get("/:provider", user.SignInOAuth)
			m.Get("/:provider/callback", user.SignInOAuthCallback)
		})
		m.Get("/link_account", user.LinkAccount)
		m.Post("/link_account_signin", bindIgnErr(auth.SignInForm{}), user.LinkAccountPostSignIn)
		m.Post("/link_account_signup", bindIgnErr(auth.RegisterForm{}), user.LinkAccountPostRegister)

		m.Get("/link_account_new", user.LinkAccountNew)
		m.Post("/link_account_signup_new", bindIgnErr(auth.RegisterOauthForm{}), user.LinkAccountPostRegisterNew)
		m.Group("/two_factor", func() {
			m.Get("", user.TwoFactor)
			m.Post("", bindIgnErr(auth.TwoFactorAuthForm{}), user.TwoFactorPost)
			m.Get("/scratch", user.TwoFactorScratch)
			m.Post("/scratch", bindIgnErr(auth.TwoFactorScratchAuthForm{}), user.TwoFactorScratchPost)
		})
		m.Group("/u2f", func() {
			m.Get("", user.U2F)
			m.Get("/challenge", user.U2FChallenge)
			m.Post("/sign", bindIgnErr(u2f.SignResponse{}), user.U2FSign)

		})
		m.Get("/guest_verify", user.GuestVerify)
	}, reqSignOut)

	m.Any("/user/events", reqSignIn, events.Events)
	m.Get("/slideImage", user.CreateSlideImageInfo)
	m.Post("/verifySlideImage", bindIgnErr(auth.SlideImageForm{}), user.VerifySlideImage)
	m.Post("/verifyGuestSlideImage", bindIgnErr(auth.SlideImageForm{}), user.VerifyGuestSlideImage)
	m.Post("/sendVerifyCode", bindIgnErr(auth.PhoneNumberForm{}), user.SendVerifyCode)
	m.Post("/bindPhone", reqSignIn, bindIgnErr(auth.PhoneNumberCodeForm{}), user.BindPhone)

	m.Group("/login/oauth", func() {
		m.Get("/authorize", bindIgnErr(auth.AuthorizationForm{}), user.AuthorizeOAuth)
		m.Post("/grant", bindIgnErr(auth.GrantApplicationForm{}), user.GrantApplicationOAuth)
		// TODO manage redirection
		m.Post("/authorize", bindIgnErr(auth.AuthorizationForm{}), user.AuthorizeOAuth)
	}, ignSignInAndCsrf, reqSignIn)
	m.Post("/login/oauth/access_token", bindIgnErr(auth.AccessTokenForm{}), ignSignInAndCsrf, user.AccessTokenOAuth)

	m.Group("/authentication/wechat", func() {
		m.Get("/qrCode4Bind", authentication.GetQRCode4Bind)
		//m.Post("/unbind", authentication.UnbindWechat)
		m.Get("/bind", authentication.GetBindPage)
	}, reqSignIn)
	m.Get("/authentication/wechat/bindStatus", authentication.GetBindStatus)
	m.Group("/user/settings", func() {
		m.Get("", userSetting.Profile)
		m.Post("", bindIgnErr(auth.UpdateProfileForm{}), userSetting.ProfilePost)
		m.Get("/change_password", user.MustChangePassword)
		m.Post("/change_password", bindIgnErr(auth.MustChangePasswordForm{}), user.MustChangePasswordPost)
		m.Post("/avatar", binding.MultipartForm(auth.AvatarForm{}), userSetting.AvatarPost)
		m.Post("/avatar/diffusion", bindIgnErr(auth.AvatarDiffusionForm{}), userSetting.AvatarDiffusionGeneratePost)
		m.Post("/avatar/change_diffusion", bindIgnErr(auth.AvatarDiffusionUpdateForm{}), userSetting.AvatarDiffusionUpdatePost)
		m.Post("/avatar/delete", userSetting.DeleteAvatar)
		m.Group("/account", func() {
			m.Combo("").Get(userSetting.Account).Post(bindIgnErr(auth.ChangePasswordForm{}), userSetting.AccountPost)
			m.Post("/email", bindIgnErr(auth.AddEmailForm{}), userSetting.EmailPost)
			m.Post("/email/delete", userSetting.DeleteEmail)
			m.Post("/delete", userSetting.DeleteAccount)
			m.Post("/theme", bindIgnErr(auth.UpdateThemeForm{}), userSetting.UpdateUIThemePost)
		})
		m.Group("/security", func() {
			m.Get("", userSetting.Security)
			m.Group("/two_factor", func() {
				m.Post("/regenerate_scratch", userSetting.RegenerateScratchTwoFactor)
				m.Post("/disable", userSetting.DisableTwoFactor)
				m.Get("/enroll", userSetting.EnrollTwoFactor)
				m.Post("/enroll", bindIgnErr(auth.TwoFactorAuthForm{}), userSetting.EnrollTwoFactorPost)
			})
			m.Group("/u2f", func() {
				m.Post("/request_register", bindIgnErr(auth.U2FRegistrationForm{}), userSetting.U2FRegister)
				m.Post("/register", bindIgnErr(u2f.RegisterResponse{}), userSetting.U2FRegisterPost)
				m.Post("/delete", bindIgnErr(auth.U2FDeleteForm{}), userSetting.U2FDelete)
			})
			m.Group("/openid", func() {
				m.Post("", bindIgnErr(auth.AddOpenIDForm{}), userSetting.OpenIDPost)
				m.Post("/delete", userSetting.DeleteOpenID)
				m.Post("/toggle_visibility", userSetting.ToggleOpenIDVisibility)
			}, openIDSignInEnabled)
			m.Post("/account_link", userSetting.DeleteAccountLink)
		})
		m.Group("/applications/oauth2", func() {
			m.Get("/:id", userSetting.OAuth2ApplicationShow)
			m.Post("/:id", bindIgnErr(auth.EditOAuth2ApplicationForm{}), userSetting.OAuthApplicationsEdit)
			m.Post("/:id/regenerate_secret", userSetting.OAuthApplicationsRegenerateSecret)
			m.Post("", bindIgnErr(auth.EditOAuth2ApplicationForm{}), userSetting.OAuthApplicationsPost)
			m.Post("/delete", userSetting.DeleteOAuth2Application)
			m.Post("/revoke", userSetting.RevokeOAuth2Grant)
		})

		m.Combo("/applications").Get(userSetting.Applications).
			Post(bindIgnErr(auth.NewAccessTokenForm{}), userSetting.ApplicationsPost)
		m.Post("/applications/delete", userSetting.DeleteApplication)
		m.Combo("/keys").Get(userSetting.Keys).
			Post(bindIgnErr(auth.AddKeyForm{}), userSetting.KeysPost)
		m.Post("/keys/delete", userSetting.DeleteKey)
		m.Get("/organization", userSetting.Organization)
		m.Get("/repos", userSetting.Repos)
	}, reqSignIn, func(ctx *context.Context) {
		ctx.Data["PageIsUserSettings"] = true
		ctx.Data["AllThemes"] = setting.UI.Themes
	})

	m.Group("/user", func() {
		// r.Get("/feeds", binding.Bind(auth.FeedsForm{}), user.Feeds)
		m.Any("/activate", user.Activate, reqSignIn)
		m.Any("/activate_email", user.ActivateEmail)
		m.Post("/update_email", bindIgnErr(auth.UpdateEmailForm{}), user.UpdateEmailPost)
		m.Get("/avatar/:username/:size", user.Avatar)
		m.Get("/email2user", user.Email2User)
		m.Get("/recover_account", user.ResetPasswd)
		m.Post("/recover_account", user.ResetPasswdPost)
		m.Post("/recover_account_by_phone", bindIgnErr(auth.ResetPassWordByPhoneForm{}), user.ResetPasswdByPhonePost)
		m.Get("/forgot_password", user.ForgotPasswd)
		m.Post("/forgot_password", user.ForgotPasswdPost)
		m.Post("/logout", user.SignOut)
		m.Get("/invitation_code", reqSignIn, user.GetInvitaionCode)
		m.Get("/invitation_tpl", reqSignIn, user.InviationTpl)
	})
	// ***** END: User *****
	m.Group("/:username/:reponame", func() {
		m.Group("/ai_task", func() {
			m.Get("/can_create", ai_task.CanCreateAITask)
		})
	}, reqSignIn, context.RepoAssignment(), context.RepoRefByType(context.RepoRefBranch))

	m.Get("/avatar/:hash", user.AvatarByEmailHash)

	m.Get("/show/icon/:hash", badge.GetIcon)

	adminReq := context.Toggle(&context.ToggleOptions{SignInRequired: true, AdminRequired: true})

	// ***** START: Admin *****
	m.Group("/admin", func() {
		m.Get("", adminReq, admin.Dashboard)
		m.Post("", adminReq, bindIgnErr(auth.AdminDashboardForm{}), admin.DashboardPost)
		m.Get("/config", admin.Config)
		m.Post("/config/test_mail", admin.SendTestMail)
		m.Group("/monitor", func() {
			m.Get("", admin.Monitor)
			m.Post("/cancel/:pid", admin.MonitorCancel)
			m.Group("/queue/:qid", func() {
				m.Get("", admin.Queue)
				m.Post("/set", admin.SetQueueSettings)
				m.Post("/add", admin.AddWorkers)
				m.Post("/cancel/:pid", admin.WorkerCancel)
				m.Post("/flush", admin.Flush)
			})
		})

		m.Group("/users", func() {
			m.Get("", admin.Users)
			m.Combo("/new").Get(admin.NewUser).Post(bindIgnErr(auth.AdminCreateUserForm{}), admin.NewUserPost)
			m.Combo("/:userid").Get(admin.EditUser).Post(bindIgnErr(auth.AdminEditUserForm{}), admin.EditUserPost)
			m.Post("/:userid/delete", admin.DeleteUser)
		})

		m.Group("/emails", func() {
			m.Get("", admin.Emails)
			m.Post("/activate", admin.ActivateEmail)
		})

		m.Group("/orgs", func() {
			m.Get("", admin.Organizations)
		})

		m.Group("/repos", func() {
			m.Get("", admin.Repos)
			m.Post("/delete", admin.DeleteRepo)
		})

		m.Group("/datasets", func() {
			m.Get("", admin.Datasets)
			m.Put("/:id/action/:action", admin.DatasetAction)
			// m.Post("/delete", admin.DeleteDataset)
		})
		m.Group("/model", func() {
			m.Get("", admin.AdminModelManage)
			m.Put("/action", admin.ModifyModelRecommend)
		})
		m.Group("/cloudbrains", func() {
			m.Get("", admin.CloudBrains)
			m.Get("/download", admin.DownloadCloudBrains)
		})
		m.Group("/images", func() {
			m.Get("", admin.Images)
			m.Get("/data", repo.GetAllImages)
			m.Get("/commit_image", admin.CloudBrainCommitImageShow)
			m.Post("/commit_image", bindIgnErr(auth.CommitAdminImageCloudBrainForm{}), repo.CloudBrainAdminCommitImage)
		})
		m.Put("/image/:id/action/:action", bindIgnErr(auth.ReviewImageForm{}), image.Action)
		m.Group("/ai_task_tmpl", func() {
			m.Get("", admin.AiTaskTmpl)
		})
		m.Group("/^:configType(hooks|system-hooks)$", func() {
			m.Get("", admin.DefaultOrSystemWebhooks)
			m.Post("/delete", admin.DeleteDefaultOrSystemWebhook)
			m.Get("/:type/new", repo.WebhooksNew)
			m.Post("/gitea/new", bindIgnErr(auth.NewWebhookForm{}), repo.GiteaHooksNewPost)
			m.Post("/gogs/new", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksNewPost)
			m.Post("/slack/new", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksNewPost)
			m.Post("/discord/new", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksNewPost)
			m.Post("/dingtalk/new", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksNewPost)
			m.Post("/telegram/new", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksNewPost)
			m.Post("/matrix/new", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksNewPost)
			m.Post("/msteams/new", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksNewPost)
			m.Post("/feishu/new", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksNewPost)
			m.Get("/:id", repo.WebHooksEdit)
			m.Post("/gitea/:id", bindIgnErr(auth.NewWebhookForm{}), repo.WebHooksEditPost)
			m.Post("/gogs/:id", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksEditPost)
			m.Post("/slack/:id", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksEditPost)
			m.Post("/discord/:id", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksEditPost)
			m.Post("/dingtalk/:id", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksEditPost)
			m.Post("/telegram/:id", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksEditPost)
			m.Post("/matrix/:id", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksEditPost)
			m.Post("/msteams/:id", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksEditPost)
			m.Post("/feishu/:id", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksEditPost)
		})

		m.Group("/auths", func() {
			m.Get("", admin.Authentications)
			m.Combo("/new").Get(admin.NewAuthSource).Post(bindIgnErr(auth.AuthenticationForm{}), admin.NewAuthSourcePost)
			m.Combo("/:authid").Get(admin.EditAuthSource).
				Post(bindIgnErr(auth.AuthenticationForm{}), admin.EditAuthSourcePost)
			m.Post("/:authid/delete", admin.DeleteAuthSource)
		})

		m.Group("/notices", func() {
			m.Get("", admin.Notices)
			m.Post("/delete", admin.DeleteNotices)
			m.Post("/empty", admin.EmptyNotices)
		})

		m.Group("/resources", func() {
			m.Group("/queue", func() {
				m.Get("", admin.GetQueuePage)
				m.Get("/list", admin.GetResourceQueueList)
				m.Post("/grampus/sync", admin.SyncGrampusQueue)
				m.Get("/codes", admin.GetResourceQueueCodes)
				m.Get("/centers", admin.GetResourceAiCenters)
				m.Post("/add", binding.Bind(models.ResourceQueueReq{}), admin.AddResourceQueue)
				m.Post("/update/:id", binding.BindIgnErr(models.ResourceQueueReq{}), admin.UpdateResourceQueue)
			})
			m.Group("/specification", func() {
				m.Get("", admin.GetSpecificationPage)
				m.Get("/list", admin.GetResourceSpecificationList)
				m.Get("/list/all", admin.GetAllResourceSpecificationList)
				m.Get("/scenes/:id", admin.GetResourceSpecificationScenes)
				m.Post("/grampus/sync", admin.SyncGrampusSpecs)
				m.Post("/add", binding.Bind(models.ResourceSpecificationReq{}), admin.AddResourceSpecification)
				m.Post("/update/:id", binding.BindIgnErr(models.ResourceSpecificationReq{}), admin.UpdateResourceSpecification)
				m.Get("/log", binding.BindIgnErr(models.ResourceSpecificationReq{}), admin.GetResourceSpecLogs)
			})
			m.Group("/scene", func() {
				m.Get("", admin.GetScenePage)
				m.Get("/list", admin.GetResourceSceneList)
				m.Post("/add", binding.Bind(models.ResourceSceneReq{}), admin.AddResourceScene)
				m.Post("/update/:id", binding.BindIgnErr(models.ResourceSceneReq{}), admin.UpdateResourceScene)
			})
			m.Group("/image", func() {
				m.Post("/sync", admin.SyncGrampusImage)
			})
		})
		m.Group("/third-party-app", func() {
			m.Get("", admin.GetThirdPartyAppPage)
		})
		m.Group("/log", func() {
			m.Group("/operation", func() {
				m.Get("", admin.GetLogPage)
			})
		})
		m.Group("/ai_model", func() {
			m.Post("/update_version", repo.UpdateAllModelMeta)
			m.Get("/query_meta", repo.QueryModelMetaById)
		})
		m.Group("/access", func() {
			m.Get("", admin.Access)
			m.Get("/:id", admin.EditAccess)
			m.Get("/batch", admin.BatchAccess)
		})
		m.Group("/org_access", func() {
			m.Get("", admin.OrgAccess)
			m.Get("/:id", admin.EditOrgAccess)
			// m.Get("/batch", admin.BatchAccess)
		})
		m.Get("/roles", admin.Roles)
		m.Group("/points", func() {
			m.Get("", admin.Points)
			m.Get("/search", point.SearchPointAccount)
			m.Post("/operate", binding.Bind(models.AdminRewardOperateReq{}), point.OperatePointAccountBalance)
		})

	}, adminReq)
	// ***** END: Admin *****

	// ***** START: Operation *****
	m.Group("/operation", func() {
		m.Get("/config/recommend_org", operation.Organizations)
		m.Post("/config/recommend_org", bindIgnErr(operation.OrgInfos{}), operation.UpdateRecommendOrganizations)

		m.Group("/reward/point", func() {
			m.Combo("/limiter/single-daily").Get(point.GetSingleDailyPointLimitConfig).Post(bindIgnErr(models.LimitConfigVO{}), point.SetSingleDailyPointLimitConfig)
			m.Post("/limiter/delete", point.DeletePointLimitConfig)
			m.Get("/account/search", point.SearchPointAccount)
			m.Post("/account/operate", binding.Bind(models.AdminRewardOperateReq{}), point.OperatePointAccountBalance)
			m.Post("/account/batch_operate", point.BatchOperatePointAccountBalance)
			m.Get("/list", point.GetAdminRewardList)
		}, apiv1.HasOperRoleOrAdmin(role.ROLE_OPER_RewardPointAdmin))

		m.Group("/task/config", func() {
			m.Get("/list", task.GetTaskConfigList)
			m.Post("/add/batch", bindIgnErr(models.BatchLimitConfigVO{}), task.BatchAddTaskConfig)
			m.Post("/^:action(new|edit|del)$", bindIgnErr(models.TaskConfigWithLimit{}), task.OperateTaskConfig)
		})

		m.Group("/badge", func() {
			m.Group("/category", func() {
				m.Get("/list", badge.GetBadgeCategoryList)
				m.Post("/^:action(new|edit|del)$", bindIgnErr(models.BadgeCategory4Show{}), badge.OperateBadgeCategory)
			})
			m.Group("/customize", func() {
				m.Get("/list", badge.GetCustomizeBadgeList)
			})
			m.Group("/users", func() {
				m.Get("", badge.GetBadgeUsers)
				m.Post("/add", bindIgnErr(models.AddBadgeUsersReq{}), badge.AddOperateBadgeUsers)
				m.Post("/del", bindIgnErr(models.DelBadgeUserReq{}), badge.DelBadgeUsers)
			})
			m.Post("/^:action(new|edit|del)$", bindIgnErr(models.BadgeOperateReq{}), badge.OperateBadge)
		})
		m.Post("/icon/upload", bindIgnErr(badge_service.IconUploadForm{}), badge.UploadIcon)
	}, operationReq)
	// ***** END: Operation *****

	m.Group("", func() {
		m.Get("/:username", user.Profile)
	}, ignSignIn)

	m.Group("", func() {
		m.Get("/attachments/:uuid", repo.GetAttachment)
	})

	m.Group("/attachments", func() {
		m.Post("", repo.UploadAttachment)
		m.Post("/delete", repo.DeleteAttachment)
		m.Post("/batch_delete", repo.BatchDeleteAttachment)
		m.Get("/get_pre_url", repo.GetPresignedPutObjectURL)
		m.Post("/add", repo.AddAttachment)

		m.Post("/edit", bindIgnErr(auth.EditAttachmentForm{}), repo.EditAttachment)
		m.Post("/private", repo.UpdatePublicAttachment)
		m.Get("/get_chunks", repo.GetSuccessChunks)
		m.Get("/new_multipart", repo.NewMultipart)
		m.Put("/obs_proxy_multipart", repo.PutOBSProxyUpload)
		m.Get("/obs_proxy_download", repo.GetOBSProxyDownload)
		m.Get("/get_multipart_url", repo.GetMultipartUploadUrl)
	}, reqSignIn)

	m.Group("/attachments", func() {
		m.Post("/decompress_done_notify", repo.UpdateAttachmentDecompressState)
		m.Post("/complete_multipart", repo.CompleteMultipart)
	})

	m.Group("/attachments/model", func() {
		m.Get("/get_chunks", repo.GetModelChunks)
		m.Get("/new_multipart", repo.NewModelMultipart)
		m.Get("/get_multipart_url", repo.GetModelMultipartUploadUrl)
		m.Post("/complete_multipart", repo.CompleteModelMultipart)
	})

	m.Group("/attachments", func() {
		m.Get("/public/query", repo.QueryAllPublicDataset)
		m.Get("/private/:username", repo.QueryPrivateDataset)
	}, reqBasicAuth)

	m.Group("/:username", func() {
		m.Post("/action/:action", user.Action)
	}, reqSignIn)

	m.Group("/blockchain", func() {
		m.Post("/init_notify", repo.HandleBlockChainInitNotify)
		m.Post("/commit_notify", repo.HandleBlockChainCommitNotify)
	})

	if macaron.Env == macaron.DEV {
		m.Get("/template/*", dev.TemplatePreview)
	}

	reqRepoAdmin := context.RequireRepoAdmin()
	reqRepoCodeWriter := context.RequireRepoWriter(models.UnitTypeCode)
	reqRepoCodeReader := context.RequireRepoReader(models.UnitTypeCode)
	reqRepoReleaseWriter := context.RequireRepoWriter(models.UnitTypeReleases)
	reqRepoReleaseReader := context.RequireRepoReader(models.UnitTypeReleases)
	reqRepoWikiWriter := context.RequireRepoWriter(models.UnitTypeWiki)
	reqRepoIssueWriter := context.RequireRepoWriter(models.UnitTypeIssues)
	reqRepoIssueReader := context.RequireRepoReader(models.UnitTypeIssues)
	reqRepoPullsReader := context.RequireRepoReader(models.UnitTypePullRequests)
	reqRepoIssuesOrPullsWriter := context.RequireRepoWriterOr(models.UnitTypeIssues, models.UnitTypePullRequests)
	reqRepoIssuesOrPullsReader := context.RequireRepoReaderOr(models.UnitTypeIssues, models.UnitTypePullRequests)
	reqRepoDatasetReader := context.RequireRepoReader(models.UnitTypeDatasets)
	reqRepoDatasetWriter := context.RequireRepoWriter(models.UnitTypeDatasets)
	reqRepoDatasetReaderJson := context.RequireRepoReaderJson(models.UnitTypeDatasets)
	reqRepoDatasetWriterJson := context.RequireRepoWriterJson(models.UnitTypeDatasets)

	reqRepoCloudBrainReader := context.RequireRepoReader(models.UnitTypeCloudBrain)
	reqRepoCloudBrainWriter := context.RequireRepoWriter(models.UnitTypeCloudBrain)
	reqRepoHPCReader := context.RequireRepoReader(models.UnitTypeHPC)
	reqRepoHPCWriter := context.RequireRepoWriter(models.UnitTypeHPC)

	//reqRepoModelManageReader := context.RequireRepoReader(models.UnitTypeModelManage)
	//reqRepoModelManageWriter := context.RequireRepoWriter(models.UnitTypeModelManage)
	//reqRepoBlockChainReader := context.RequireRepoReader(models.UnitTypeBlockChain)
	//reqRepoBlockChainWriter := context.RequireRepoWriter(models.UnitTypeBlockChain)

	m.Group("/ide", func() {
		m.Get("", func(ctx *context.Context) {
			ctx.Redirect(setting.AppSubURL + "/project")
		})
		m.Get("/project", routers.IdeProject)

	}, reqSignIn, func(ctx *context.Context) {
		ctx.Data["PageIsUserSettings"] = true
		ctx.Data["AllThemes"] = setting.UI.Themes
	})

	// ***** START: Organization *****
	m.Group("/org", func() {
		m.Group("/:org", func() {
			m.Get("/members", org.Members)
			m.Group("/org_tag", func() {
				m.Get("/repo_list", org.GetTagRepos)
				m.Post("/repo_submit", bindIgnErr(auth.SubmitReposOfTagForm{}), org.SubmitTags)

				m.Get("/model_list", org.GetTagModel)
				m.Post("/model_submit", bindIgnErr(auth.SubmitModelOfTagForm{}), org.SubmitTagsToModel)

				m.Get("/dataset_list", org.GetTagDataset)
				m.Post("/dataset_submit", bindIgnErr(auth.SubmitDatasetOfTagForm{}), org.SubmitTagsToDataset)

			})

			m.Get("/org_list_repo", org.GetOrgRepoList)
			m.Get("/org_list_dataset", org.GetOrgDatasetList)
			m.Get("/org_list_model", org.GetOrgModelList)

			m.Get("/org_card_repo", org.GetOrgRepoCardList)
			m.Get("/org_card_dataset", org.GetOrgDatasetCardList)
			m.Get("/org_card_model", org.GetOrgModelCardList)

		}, context.OrgAssignment())
	})
	m.Group("/org", func() {
		m.Group("", func() {
			m.Get("/create", org.Create)
			m.Post("/create", bindIgnErr(auth.CreateOrgForm{}), org.CreatePost)
		})

		m.Group("/:org", func() {
			// m.Get("/repositories", user.Dashboard)
			m.Group("/repositories", func() {
				m.Get("", user.Dashboard)
				m.Get("/^:type(issues|pulls)$", user.Issues)
			})
			m.Get("/models", org.Models)
			m.Get("/datasets", org.Datasets)
			m.Get("/storages", org.Storages)
			// m.Get("/^:type(issues|pulls)$", user.Issues)
			m.Get("/milestones", reqMilestonesDashboardPageEnabled, user.Milestones)
			//m.Get("/members", org.Members)
			m.Post("/members/action/:action", org.MembersAction)

			m.Get("/teams", org.Teams)
		}, context.OrgAssignment(true))

		m.Group("/:org", func() {
			m.Get("/teams/:team", org.TeamMembers)
			m.Get("/teams/:team/repositories", org.TeamRepositories)
			m.Get("/teams/:team/datasets", org.TeamDatasaets)
			m.Get("/teams/:team/aimodels", org.TeamAimodels)
			m.Post("/teams/:team/action/:action", org.TeamsAction)
			m.Post("/teams/:team/action/repo/:action", org.TeamsRepoAction)
			m.Post("/teams/:team/action/dataset/:action", org.TeamsDatasetAction)
			m.Post("/teams/:team/action/aimodel/:action", org.TeamsAimodelAction)
		}, context.OrgAssignment(true, false, true))

		m.Group("/:org", func() {
			m.Get("/teams/new", org.NewTeam)
			m.Post("/teams/new", bindIgnErr(auth.CreateTeamForm{}), org.NewTeamPost)
			m.Get("/teams/:team/edit", org.EditTeam)
			m.Post("/teams/:team/edit", bindIgnErr(auth.CreateTeamForm{}), org.EditTeamPost)
			m.Post("/teams/:team/delete", org.DeleteTeam)

			m.Group("/settings", func() {
				m.Combo("").Get(org.Settings).
					Post(bindIgnErr(auth.UpdateOrgSettingForm{}), org.SettingsPost)
				m.Post("/avatar", binding.MultipartForm(auth.AvatarForm{}), org.SettingsAvatar)
				m.Post("/avatar/delete", org.SettingsDeleteAvatar)

				m.Group("/hooks", func() {
					m.Get("", org.Webhooks)
					m.Post("/delete", org.DeleteWebhook)
					m.Get("/:type/new", repo.WebhooksNew)
					m.Post("/gitea/new", bindIgnErr(auth.NewWebhookForm{}), repo.GiteaHooksNewPost)
					m.Post("/gogs/new", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksNewPost)
					m.Post("/slack/new", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksNewPost)
					m.Post("/discord/new", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksNewPost)
					m.Post("/dingtalk/new", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksNewPost)
					m.Post("/telegram/new", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksNewPost)
					m.Post("/matrix/new", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksNewPost)
					m.Post("/msteams/new", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksNewPost)
					m.Post("/feishu/new", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksNewPost)
					m.Get("/:id", repo.WebHooksEdit)
					m.Post("/gitea/:id", bindIgnErr(auth.NewWebhookForm{}), repo.WebHooksEditPost)
					m.Post("/gogs/:id", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksEditPost)
					m.Post("/slack/:id", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksEditPost)
					m.Post("/discord/:id", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksEditPost)
					m.Post("/dingtalk/:id", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksEditPost)
					m.Post("/telegram/:id", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksEditPost)
					m.Post("/matrix/:id", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksEditPost)
					m.Post("/msteams/:id", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksEditPost)
					m.Post("/feishu/:id", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksEditPost)
				})

				m.Group("/labels", func() {
					m.Get("", org.RetrieveLabels, org.Labels)
					m.Post("/new", bindIgnErr(auth.CreateLabelForm{}), org.NewLabel)
					m.Post("/edit", bindIgnErr(auth.CreateLabelForm{}), org.UpdateLabel)
					m.Post("/delete", org.DeleteLabel)
					m.Post("/initialize", bindIgnErr(auth.InitializeLabelsForm{}), org.InitializeLabels)
				})

				m.Route("/delete", "GET,POST", org.SettingsDelete)
			})
		}, context.OrgAssignment(true, true))
	}, reqSignIn)
	// ***** END: Organization *****

	m.Group("/course", func() {
		m.Get("/create", repo.CreateCourse)
		m.Post("/create", bindIgnErr(auth.CreateCourseForm{}), repo.CreateCoursePost)
		m.Get("/addOrg", repo.AddCourseOrg)

	}, reqSignIn)

	// ***** START: Repository *****
	m.Group("/repo", func() {
		m.Get("/create", repo.Create)
		m.Post("/create", bindIgnErr(auth.CreateRepoForm{}), repo.CreatePost)
		m.Get("/migrate", repo.Migrate)
		m.Post("/migrate", bindIgnErr(auth.MigrateRepoForm{}), repo.MigratePost)
		m.Get("/model_migrate", repo.ModelMigrateTmpl)
		m.Group("/fork", func() {
			m.Combo("/:repoid").Get(repo.Fork).
				Post(bindIgnErr(auth.CreateRepoForm{}), repo.ForkPost)
		}, context.RepoIDAssignment(), context.UnitTypes(), reqRepoCodeReader)
		m.Get("/check_name", repo.CheckName)
	}, reqSignIn)

	m.Group("/tech", func() {
		m.Get("/new", tech.Create)
		m.Get("/tech_view", tech.TechView)
		m.Get("/repo_view", tech.RepoView)
		m.Get("/admin_view", apiv1.HasOperRole(role.ROLE_OPER_TechProgramAdmin), tech.AdminView)
		m.Get("/my_view", tech.MyView)

	}, reqSignIn)

	// ***** Release Attachment Download without Signin
	m.Get("/:username/:reponame/releases/download/:vTag/:fileName", ignSignIn, context.LowLimiter(), context.RepoAssignment(), repo.MustBeNotEmpty, repo.RedirectDownload)

	m.Group("/:username/:reponame", func() {
		m.Group("/settings", func() {
			m.Combo("").Get(repo.Settings).
				Post(bindIgnErr(auth.RepoSettingForm{}), repo.SettingsPost)
			m.Post("/avatar", binding.MultipartForm(auth.AvatarForm{}), repo.SettingsAvatar)
			m.Post("/avatar/delete", repo.SettingsDeleteAvatar)

			m.Group("/collaboration", func() {
				m.Combo("").Get(repo.Collaboration).Post(repo.CollaborationPost)
				m.Post("/access_mode", repo.ChangeCollaborationAccessMode)
				m.Post("/delete", repo.DeleteCollaboration)
				m.Group("/team", func() {
					m.Post("", repo.AddTeamPost)
					m.Post("/delete", repo.DeleteTeam)
				})
			})
			m.Group("/branches", func() {
				m.Combo("").Get(repo.ProtectedBranch).Post(repo.ProtectedBranchPost)
				m.Combo("/*").Get(repo.SettingsProtectedBranch).
					Post(bindIgnErr(auth.ProtectBranchForm{}), context.RepoMustNotBeArchived(), repo.SettingsProtectedBranchPost)
			}, repo.MustBeNotEmpty)

			m.Group("/hooks", func() {
				m.Get("", repo.Webhooks)
				m.Post("/delete", repo.DeleteWebhook)
				m.Get("/:type/new", repo.WebhooksNew)
				m.Post("/gitea/new", bindIgnErr(auth.NewWebhookForm{}), repo.GiteaHooksNewPost)
				m.Post("/gogs/new", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksNewPost)
				m.Post("/slack/new", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksNewPost)
				m.Post("/discord/new", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksNewPost)
				m.Post("/dingtalk/new", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksNewPost)
				m.Post("/telegram/new", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksNewPost)
				m.Post("/matrix/new", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksNewPost)
				m.Post("/msteams/new", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksNewPost)
				m.Post("/feishu/new", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksNewPost)
				m.Get("/:id", repo.WebHooksEdit)
				m.Post("/:id/test", repo.TestWebhook)
				m.Post("/gitea/:id", bindIgnErr(auth.NewWebhookForm{}), repo.WebHooksEditPost)
				m.Post("/gogs/:id", bindIgnErr(auth.NewGogshookForm{}), repo.GogsHooksEditPost)
				m.Post("/slack/:id", bindIgnErr(auth.NewSlackHookForm{}), repo.SlackHooksEditPost)
				m.Post("/discord/:id", bindIgnErr(auth.NewDiscordHookForm{}), repo.DiscordHooksEditPost)
				m.Post("/dingtalk/:id", bindIgnErr(auth.NewDingtalkHookForm{}), repo.DingtalkHooksEditPost)
				m.Post("/telegram/:id", bindIgnErr(auth.NewTelegramHookForm{}), repo.TelegramHooksEditPost)
				m.Post("/matrix/:id", bindIgnErr(auth.NewMatrixHookForm{}), repo.MatrixHooksEditPost)
				m.Post("/msteams/:id", bindIgnErr(auth.NewMSTeamsHookForm{}), repo.MSTeamsHooksEditPost)
				m.Post("/feishu/:id", bindIgnErr(auth.NewFeishuHookForm{}), repo.FeishuHooksEditPost)

				m.Group("/git", func() {
					m.Get("", repo.GitHooks)
					m.Combo("/:name").Get(repo.GitHooksEdit).
						Post(repo.GitHooksEditPost)
				}, context.GitHookService())
			})

			m.Group("/keys", func() {
				m.Combo("").Get(repo.DeployKeys).
					Post(bindIgnErr(auth.AddKeyForm{}), repo.DeployKeysPost)
				m.Post("/delete", repo.DeleteDeployKey)
			})

			m.Group("/lfs", func() {
				m.Get("", repo.LFSFiles)
				m.Get("/show/:oid", repo.LFSFileGet)
				m.Post("/delete/:oid", repo.LFSDelete)
				m.Get("/pointers", repo.LFSPointerFiles)
				m.Post("/pointers/associate", repo.LFSAutoAssociate)
				m.Get("/find", repo.LFSFileFind)
				m.Group("/locks", func() {
					m.Get("/", repo.LFSLocks)
					m.Post("/", repo.LFSLockFile)
					m.Post("/:lid/unlock", repo.LFSUnlock)
				})
			})

		}, func(ctx *context.Context) {
			ctx.Data["PageIsSettings"] = true
			ctx.Data["LFSStartServer"] = setting.LFS.StartServer
		})
	}, reqSignIn, context.RepoAssignment(), context.UnitTypes(), reqRepoAdmin, context.RepoRef())

	m.Post("/:username/:reponame/action/:action", reqSignIn, context.RepoAssignment(), context.UnitTypes(), repo.Action)

	// Grouping for those endpoints not requiring authentication
	m.Group("/:username/:reponame", func() {
		m.Get("/contributors", repo.Contributors)
		m.Get("/contributors/list", repo.ContributorsAPI)
	}, context.GuestHandler(false), context.LowLimiter(), context.RepoAssignment(), context.UnitTypes())

	m.Group("/:username/:reponame", func() {
		m.Group("/milestone", func() {
			m.Get("/:id", repo.MilestoneIssuesAndPulls)
		}, reqRepoIssuesOrPullsReader, context.RepoRef())
		m.Combo("/compare/*", repo.MustBeNotEmpty, reqRepoCodeReader, repo.SetEditorconfigIfExists).
			Get(reqSignIn, repo.SetDiffViewStyle, repo.CompareDiff).
			Post(reqSignIn, context.RepoMustNotBeArchived(), reqRepoPullsReader, repo.MustAllowPulls, bindIgnErr(auth.CreateIssueForm{}), repo.CompareAndPullRequestPost)

	}, context.GuestHandler(false), context.LowLimiter(), context.RepoAssignment(), context.UnitTypes())

	// Grouping for those endpoints that do require authentication
	m.Group("/:username/:reponame", func() {
		m.Group("/issues", func() {
			m.Combo("/new").Get(context.RepoRef(), repo.NewIssue).
				Post(bindIgnErr(auth.CreateIssueForm{}), repo.NewIssuePost)
		}, context.RepoMustNotBeArchived(), reqRepoIssueReader)
		// FIXME: should use different URLs but mostly same logic for comments of issue and pull reuqest.
		// So they can apply their own enable/disable logic on routers.
		m.Group("/issues", func() {
			m.Group("/:index", func() {
				m.Post("/title", repo.UpdateIssueTitle)
				m.Post("/content", repo.UpdateIssueContent)
				m.Post("/watch", repo.IssueWatch)
				m.Post("/setstaytop", repo.SetIssueStayTop)
				m.Group("/dependency", func() {
					m.Post("/add", repo.AddDependency)
					m.Post("/delete", repo.RemoveDependency)
				})
				m.Combo("/comments").Post(repo.MustAllowUserComment, bindIgnErr(auth.CreateCommentForm{}), repo.NewComment)
				m.Group("/times", func() {
					m.Post("/add", bindIgnErr(auth.AddTimeManuallyForm{}), repo.AddTimeManually)
					m.Group("/stopwatch", func() {
						m.Post("/toggle", repo.IssueStopwatch)
						m.Post("/cancel", repo.CancelStopwatch)
					})
				})
				m.Post("/reactions/:action", bindIgnErr(auth.ReactionForm{}), repo.ChangeIssueReaction)
				m.Post("/lock", reqRepoIssueWriter, bindIgnErr(auth.IssueLockForm{}), repo.LockIssue)
				m.Post("/unlock", reqRepoIssueWriter, repo.UnlockIssue)
				m.Get("/attachments", repo.GetIssueAttachments)
			}, context.RepoMustNotBeArchived())

			m.Post("/labels", reqRepoIssuesOrPullsWriter, repo.UpdateIssueLabel)
			m.Post("/milestone", reqRepoIssuesOrPullsWriter, repo.UpdateIssueMilestone)
			m.Post("/assignee", reqRepoIssuesOrPullsWriter, repo.UpdateIssueAssignee)
			m.Post("/ref", reqRepoIssuesOrPullsWriter, repo.UpdateIssueRef)
			m.Post("/request_review", reqRepoIssuesOrPullsReader, repo.UpdatePullReviewRequest)
			m.Post("/status", reqRepoIssuesOrPullsWriter, repo.UpdateIssueStatus)
			m.Post("/resolve_conversation", reqRepoIssuesOrPullsReader, repo.UpdateResolveConversation)
		}, context.RepoMustNotBeArchived())
		m.Group("/comments/:id", func() {
			m.Post("", repo.UpdateCommentContent)
			m.Post("/delete", repo.DeleteComment)
			m.Post("/reactions/:action", bindIgnErr(auth.ReactionForm{}), repo.ChangeCommentReaction)
			m.Get("/attachments", repo.GetCommentAttachments)
		}, context.RepoMustNotBeArchived())
		m.Group("/labels", func() {
			m.Post("/new", bindIgnErr(auth.CreateLabelForm{}), repo.NewLabel)
			m.Post("/edit", bindIgnErr(auth.CreateLabelForm{}), repo.UpdateLabel)
			m.Post("/delete", repo.DeleteLabel)
			m.Post("/initialize", bindIgnErr(auth.InitializeLabelsForm{}), repo.InitializeLabels)
		}, context.RepoMustNotBeArchived(), reqRepoIssuesOrPullsWriter, context.RepoRef())
		m.Group("/milestones", func() {
			m.Combo("/new").Get(repo.NewMilestone).
				Post(bindIgnErr(auth.CreateMilestoneForm{}), repo.NewMilestonePost)
			m.Get("/:id/edit", repo.EditMilestone)
			m.Post("/:id/edit", bindIgnErr(auth.CreateMilestoneForm{}), repo.EditMilestonePost)
			m.Post("/:id/:action", repo.ChangeMilestonStatus)
			m.Post("/delete", repo.DeleteMilestone)
		}, context.RepoMustNotBeArchived(), reqRepoIssuesOrPullsWriter, context.RepoRef())
		m.Group("/pull", func() {
			m.Post("/:index/target_branch", repo.UpdatePullRequestTarget)
		}, context.RepoMustNotBeArchived())

		m.Group("", func() {
			m.Group("", func() {
				m.Combo("/_edit/*").Get(repo.EditFile).
					Post(bindIgnErr(auth.EditRepoFileForm{}), repo.EditFilePost)
				m.Combo("/_new/*").Get(repo.NewFile).
					Post(bindIgnErr(auth.EditRepoFileForm{}), repo.NewFilePost)
				m.Post("/_preview/*", bindIgnErr(auth.EditPreviewDiffForm{}), repo.DiffPreviewPost)
				m.Combo("/_delete/*").Get(repo.DeleteFile).
					Post(bindIgnErr(auth.DeleteRepoFileForm{}), repo.DeleteFilePost)
				m.Combo("/_upload/*", repo.MustBeAbleToUpload).
					Get(repo.UploadFile).
					Post(bindIgnErr(auth.UploadRepoFileForm{}), repo.UploadFilePost)
				m.Post("/_rename/*", bindIgnErr(auth.RenameRepoFileForm{}), repo.RenameFilePost)
			}, context.RepoRefByType(context.RepoRefBranch), repo.MustBeEditable)
			m.Group("", func() {
				m.Post("/upload-file", repo.UploadFileToServer)
				m.Post("/upload-remove", bindIgnErr(auth.RemoveUploadFileForm{}), repo.RemoveUploadFileFromServer)
			}, context.RepoRef(), repo.MustBeEditable, repo.MustBeAbleToUpload)
		}, context.RepoMustNotBeArchived(), reqRepoCodeWriter, repo.MustBeNotEmpty)

		m.Group("/branches", func() {
			m.Group("/_new/", func() {
				m.Post("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.CreateBranch)
				m.Post("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.CreateBranch)
				m.Post("/commit/*", context.RepoRefByType(context.RepoRefCommit), repo.CreateBranch)
			}, bindIgnErr(auth.NewBranchForm{}))
			m.Post("/delete", repo.DeleteBranchPost)
			m.Post("/restore", repo.RestoreBranchPost)
		}, context.RepoMustNotBeArchived(), reqRepoCodeWriter, repo.MustBeNotEmpty)

	}, reqSignIn, context.RepoAssignment(), context.UnitTypes())

	// Releases
	m.Group("/:username/:reponame", func() {
		m.Group("/releases", func() {
			m.Get("/", repo.Releases)
			m.Get("/tag/:tag", repo.SingleRelease)
			m.Get("/latest", repo.LatestRelease)
		}, repo.MustBeNotEmpty, context.RepoRef())
		m.Group("/releases", func() {
			m.Get("/new", repo.NewRelease)
			m.Post("/new", bindIgnErr(auth.NewReleaseForm{}), repo.NewReleasePost)
			m.Post("/delete", repo.DeleteRelease)
		}, reqSignIn, repo.MustBeNotEmpty, context.RepoMustNotBeArchived(), reqRepoReleaseWriter, context.RepoRef())
		m.Group("/releases", func() {
			m.Get("/edit/*", repo.EditRelease)
			m.Post("/edit/*", bindIgnErr(auth.EditReleaseForm{}), repo.EditReleasePost)
		}, reqSignIn, repo.MustBeNotEmpty, context.RepoMustNotBeArchived(), reqRepoReleaseWriter, func(ctx *context.Context) {
			var err error
			ctx.Repo.Commit, err = ctx.Repo.GitRepo.GetBranchCommit(ctx.Repo.Repository.DefaultBranch)
			if err != nil {
				ctx.ServerError("GetBranchCommit", err)
				return
			}
			ctx.Repo.CommitsCount, err = ctx.Repo.GetCommitsCount()
			if err != nil {
				ctx.ServerError("GetCommitsCount", err)
				return
			}
			ctx.Data["CommitsCount"] = ctx.Repo.CommitsCount
		})
	}, ignSignIn, context.LowLimiter(), context.RepoAssignment(), context.UnitTypes(), reqRepoReleaseReader)

	m.Group("/:username/:reponame", func() {
		m.Post("/topics", repo.TopicsPost)
	}, context.RepoAssignment(), context.RepoMustNotBeArchived(), reqRepoAdmin)

	m.Group("/image/:id", func() {
		m.Get("", repo.GetImage)
		m.Get("/apply", cloudbrain.AdminOrImageCreaterRight, repo.CloudBrainImageApplyRecommend)
		m.Get("/:from", cloudbrain.AdminOrImageCreaterRight, repo.CloudBrainImageEdit)
		m.Post("", cloudbrain.AdminOrImageCreaterRight, bindIgnErr(auth.EditImageCloudBrainForm{}), repo.CloudBrainImageEditPost)
		m.Post("/apply", cloudbrain.AdminOrImageCreaterRight, bindIgnErr(auth.EditImageCloudBrainForm{}), repo.CloudBrainImageRecommendApplyPost)
		m.Delete("", cloudbrain.AdminOrImageCreaterRight, repo.CloudBrainImageDelete)
		m.Put("/action/:action", reqSignIn, image.UserAction)
	})
	m.Group("/:username/:reponame", func() {
		m.Group("", func() {
			m.Get("/^:type(issues|pulls)$", repo.Issues)
		}, context.RepoRef())
	}, context.GuestHandler(false), reqSignIn, context.LowestLimiter(), context.RepoAssignment(), context.UnitTypes())
	m.Group("/:username/:reponame", func() {
		m.Group("", func() {
			m.Get("/^:type(issues|pulls)$/:index", context.GuestHandler(false), context.RepoRef(), reqSignIn, repo.ViewIssue)
			m.Get("/labels/", context.RepoRef(), reqRepoIssuesOrPullsReader, repo.RetrieveLabels, repo.Labels)
			m.Get("/milestones", context.RepoRef(), reqRepoIssuesOrPullsReader, repo.Milestones)
		})

		m.Group("/datasets", func() {
			m.Get("", reqRepoDatasetReader, repo.DatasetIndex)
			m.Get("/reference_datasets", reqRepoDatasetReader, repo.ReferenceDataset)
			m.Get("/reference_datasets_data", reqRepoDatasetReaderJson, repo.ReferenceDatasetData)
			m.Delete("/reference_datasets/:id", reqRepoDatasetWriterJson, repo.ReferenceDatasetDelete)
			m.Put("/:id/:action", reqSignIn, reqRepoDatasetReader, repo.DatasetAction)
			m.Get("/create", reqRepoDatasetWriter, repo.CreateDataset)
			m.Post("/create", reqRepoDatasetWriter, bindIgnErr(auth.CreateDatasetForm{}), repo.CreateDatasetPost)
			m.Get("/edit/:id", reqRepoDatasetWriter, repo.EditDataset)
			m.Post("/reference_datasets", reqRepoDatasetWriterJson, bindIgnErr(auth.ReferenceDatasetForm{}), repo.ReferenceDatasetPost)
			m.Post("/edit", reqRepoDatasetWriter, bindIgnErr(auth.EditDatasetForm{}), repo.EditDatasetPost)

			m.Get("/current_repo_m", repo.CurrentRepoDatasetMultiple)
			m.Get("/my_datasets_m", repo.MyDatasetsMultiple)
			m.Get("/public_datasets_m", repo.PublicDatasetMultiple)

			m.Get("/reference_datasets_available", repo.ReferenceDatasetAvailable)
			m.Get("/my_favorite_m", repo.MyFavoriteDatasetMultiple)

			m.Group("/status", func() {
				m.Get("/:uuid", repo.GetDatasetStatus)
			})

			m.Group("/attachments", func() {
				// m.Get("/upload", repo.UploadAttachmentUI)
				m.Get("/edit/:id", repo.EditAttachmentUI)
			}, reqSignIn)

			m.Group("/dirs", func() {
				m.Get("/:uuid", reqRepoDatasetReader, repo.DirIndex)
			})
			m.Group("/label", func() {
				m.Get("/:uuid", reqRepoDatasetReader, repo.LabelIndex)
			})
			m.Group("/model", func() {
				m.Get("/getcurrentdataset", reqRepoDatasetReader, repo.GetCurrentDataSet)
				m.Get("/getmodelfile", reqRepoDatasetReader, repo.GetDataSetSelectItemByJobId)
				m.Get("/getprogress", reqRepoDatasetReader, repo.GetExportDataSetByMsgId)
				m.Post("/export_exist_dataset", reqRepoDatasetWriterJson, repo.ExportModelToExistDataSet)
			})

		}, context.RepoRef())
	}, ignSignIn, context.LowLimiter(), context.RepoAssignment(), context.UnitTypes())
	m.Group("/notebook", func() {
		m.Group("/:id", func() {
			m.Get("/commit_image", cloudbrain.AdminOrJobCreaterRight, repo.GrampusCommitImageShow)
			m.Post("/commit_image", cloudbrain.AdminOrJobCreaterRight, bindIgnErr(auth.CommitImageGrampusForm{}), repo.GrampusCommitImage)
		})
	})

	m.Group("/:username/:reponame", func() {

		m.Group("/cloudbrain", func() {
			m.Group("/:id", func() {
				m.Get("", reqRepoCloudBrainReader, repo.CloudBrainShow)
				m.Get("/debug", cloudbrain.AdminOrJobCreaterRight, repo.CloudBrainDebug)
				m.Get("/commit_image", cloudbrain.AdminOrJobCreaterRight, repo.CloudBrainCommitImageShow)
				m.Post("/commit_image/check", cloudbrain.AdminOrJobCreaterRight, bindIgnErr(auth.CommitImageCloudBrainForm{}), repo.CloudBrainCommitImageCheck)
				m.Post("/commit_image", cloudbrain.AdminOrJobCreaterRight, bindIgnErr(auth.CommitImageCloudBrainForm{}), repo.CloudBrainCommitImage)
				m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.CloudBrainStop)
				m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.CloudBrainDel)
				m.Get("/rate", reqRepoCloudBrainReader, repo.GetRate)
				m.Get("/models", reqRepoCloudBrainReader, repo.CloudBrainShowModels)
				m.Get("/download_model", cloudbrain.AdminOrJobCreaterRight, repo.CloudBrainDownloadModel)
				m.Get("/download_multi_model", cloudbrain.AdminOrJobCreaterRight, repo.CloudBrainDownloadMultiModel)
			})
			m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.CloudBrainNew)

			m.Group("/benchmark", func() {
				m.Get("", reqRepoCloudBrainReader, repo.CloudBrainBenchmarkIndex)
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.CloudBrainBenchMarkShow)
					m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.CloudBrainStop)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.BenchmarkDel)
					m.Get("/rate", reqRepoCloudBrainReader, repo.GetRate)
				})
				m.Get("/get_child_types", repo.GetChildTypes)
			})

			m.Group("/train-job", func() {
				m.Group("/:jobid", func() {
					m.Get("", reqRepoCloudBrainReader, repo.CloudBrainTrainJobShow)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.CloudBrainTrainJobDel)
					//m.Get("/models", reqRepoCloudBrainReader, repo.CloudBrainShowModels)
					m.Get("/download_model", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.CloudBrainDownloadModel)
					m.Get("/download_multi_model", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.CloudBrainDownloadMultiModel)
					//m.Get("/get_log", cloudbrain.AdminOrJobCreaterRightForTrain, repo.GetLogFromModelDir)
					//m.Post("/create_version", reqWechatBind, cloudbrain.AdminOrJobCreaterRightForTrain, bindIgnErr(auth.CreateModelArtsTrainJobForm{}), repo.TrainJobCreateVersion)
					m.Get("/create_version", reqWechatBind, cloudbrain.AdminOrJobCreaterRightForTrain, context.PointAccount(), repo.CloudBrainTrainJobVersionNew)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.CloudBrainTrainJobNew)
			})
			m.Group("/inference-job", func() {
				m.Group("/:jobid", func() {
					m.Get("", reqRepoCloudBrainReader, repo.InferenceCloudBrainJobShow)
					m.Get("/result_download", cloudbrain.AdminOrJobCreaterRightForTrain, repo.CloudBrainDownloadInferenceResult)
					m.Get("/download_multi_model", cloudbrain.AdminOrJobCreaterRightForTrain, repo.CloudBrainDownloadMultiModel)
					m.Get("/downloadall", cloudbrain.AdminOrJobCreaterRightForTrain, repo.DownloadGPUInferenceResultFile)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.InferenceCloudBrainJobNew)
			})
		}, context.RepoRef())
		m.Group("/supercompute", func() {
			m.Get("", reqRepoHPCReader, super_compute.GetAPPList)
			m.Get("/job", reqRepoHPCReader, super_compute.GetJobList)
			m.Get("/job/create", reqRepoHPCWriter, super_compute.CreateUI)
			m.Get("/job/:id", reqRepoHPCReader, super_compute.GetDetailUI)
		})
		m.Group("/grampus", func() {
			m.Group("/notebook", func() {
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.GrampusNotebookShow)
					m.Get("/debug", reqWechatBind, cloudbrain.AdminOrJobCreaterRight, repo.GrampusNotebookDebug)
					m.Get("/commit_image", cloudbrain.AdminOrJobCreaterRight, repo.GrampusCommitImageShow)
					m.Post("/commit_image", cloudbrain.AdminOrJobCreaterRight, bindIgnErr(auth.CommitImageGrampusForm{}), repo.GrampusCommitImage)
					m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.GrampusStopJob)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.GrampusNotebookDel)
				})

				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusNotebookNew)
			})

			m.Group("/onlineinfer", func() {
				m.Get("", reqRepoCloudBrainReader, repo.GrampusOnlineInferIndex)
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.GrampusOnlineInferShow)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusOnlineInferNew)
			})

			m.Group("/general", func() {
				m.Get("", reqRepoCloudBrainReader, repo.GrampusGeneralIndex)
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.GrampusGeneralShow)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusGeneralNew)
			})

			m.Group("/train-job", func() {
				m.Group("/:jobid", func() {
					m.Get("", reqRepoCloudBrainReader, repo.GrampusTrainJobShow)
					m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.GrampusStopJob)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.GrampusTrainJobDel)
					m.Get("/model_download", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.ModelDownload)
					m.Get("/download_multi_model", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.MultiModelDownload)
					m.Post("/create_version", reqWechatBind, cloudbrain.AdminOrJobCreaterRightForTrain, bindIgnErr(auth.CreateGrampusTrainJobForm{}), context.PointAccount(), repo.GrampusTrainJobVersionCreate)
				})
				m.Group("/gpu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobGPUNew)
					m.Post("/create", reqWechatBind, reqRepoCloudBrainWriter, bindIgnErr(auth.CreateGrampusTrainJobForm{}), context.PointAccount(), repo.GrampusTrainJobGpuCreate)
				})
				m.Group("/npu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobNPUNew)
					m.Post("/create", reqWechatBind, reqRepoCloudBrainWriter, bindIgnErr(auth.CreateGrampusTrainJobForm{}), context.PointAccount(), repo.GrampusTrainJobNpuCreate)
				})
				m.Group("/gcu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobGCUNew)
				})
				m.Group("/dcu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobDCUNew)
					// m.Post("/create", reqWechatBind, reqRepoCloudBrainWriter, bindIgnErr(auth.CreateGrampusTrainJobForm{}), context.PointAccount(), repo.GrampusTrainJobGcuCreate)
				})
				m.Group("/iluvatar-gpgpu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobIluvatarGPGPUNew)
				})
				m.Group("/metax-gpgpu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobIluvatarGPGPUNew)
				})
				m.Group("/biren-gpu", func() {
					m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusTrainJobBIRENGPUNew)
				})
			})

			m.Group("/inference-job", func() {
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.GrampusInferenceNew)
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.GrampusInferenceShow)
				})
			})
		}, context.RepoRef())

		//m.Group("/modelmanage", func() {
		//	m.Get("/create_local_model", repo.CreateLocalModel)
		//	m.Get("/create_online_model", repo.CreateOnlineModel)
		//	m.Post("/create_local_model", repo.SaveLocalModel)
		//	m.Delete("/delete_model_file", repo.DeleteModelFile)
		//
		//	m.Get("/model_readme_tmpl", repo.ModelReadMeTmpl)
		//	m.Get("/model_readme_data", repo.QueryModelReadMe)
		//	m.Post("/model_readme_data", repo.ModifyModelReadMe)
		//	m.Get("/model_filelist_tmpl", repo.ModelFileListTmpl)
		//	m.Get("/model_fileupload_tmpl", repo.CreateLocalModelForUpload)
		//	m.Get("/model_setting", repo.ModelFileSettingTmpl)
		//	m.Get("/model_evolution_map", repo.ModelEvolutionMapTmpl)
		//	m.Get("/model_evolution_map_data", repo.ModelEvolutionMapData)
		//	m.Get("/model_migrating", repo.ModelMigratingTmpl)
		//
		//	m.Post("/create_model", repo.SaveModel)
		//	m.Post("/create_model_convert", reqWeChatStandard, reqRepoModelManageWriter, repo.SaveModelConvert)
		//	m.Post("/create_new_model", repo.SaveNewNameModel)
		//	m.Delete("/delete_model", repo.DeleteModel)
		//	m.Post("/delete_model_convert/:id", repo.DeleteModelConvert)
		//	m.Post("/convert_stop/:id", repo.StopModelConvert)
		//	m.Put("/modify_model", reqRepoModelManageWriter, repo.ModifyModelInfo)
		//	m.Put("/modify_model_status", reqRepoModelManageWriter, repo.ModifyModelPrivate)
		//	m.Get("/show_model_collect_num", reqRepoModelManageReader, repo.QueryModelCollectNum)
		//	m.Get("/show_model", reqRepoModelManageReader, repo.ShowModelTemplate)
		//	m.Get("/convert_model", reqRepoModelManageReader, repo.ConvertModelTemplate)
		//	m.Get("/show_model_info", repo.ShowModelInfo)
		//	m.Get("/show_model_convert_info", repo.ShowModelConvertInfo)
		//	m.Get("/show_model_info_api", repo.ShowSingleModel)
		//	m.Get("/show_model_api", repo.ShowModelPageInfo)
		//	m.Get("/show_model_child_api", repo.ShowOneVersionOtherModel)
		//	m.Get("/query_train_job", reqRepoCloudBrainReader, repo.QueryTrainJobList)
		//	m.Get("/query_train_model", reqRepoCloudBrainReader, repo.QueryTrainModelList)
		//	m.Get("/query_train_job_version", reqRepoCloudBrainReader, repo.QueryTrainJobVersionList)
		//	m.Get("/query_model_for_predict", reqRepoModelManageReader, repo.QueryModelListForPredict)
		//	m.Get("/query_modelfile_for_predict", reqRepoModelManageReader, repo.QueryModelFileForPredict)
		//	m.Get("/query_onelevel_modelfile", reqRepoModelManageReader, repo.QueryOneLevelModelFile)
		//	m.Get("/download_model_convert/:id", reqRepoModelManageReader, repo.ModelConvertDownloadModel)
		//	m.Group("/:ID", func() {
		//		m.Get("", repo.ShowSingleModel)
		//		m.Get("/downloadsingle", repo.DownloadSingleModelFile)
		//	})
		//	m.Get("/downloadall", repo.DownloadMultiModelFile)
		//}, context.RepoRef())

		m.Group("/modelsafety", func() {
			m.Group("/:id", func() {
				m.Get("/show", reqRepoCloudBrainReader, repo.GetAiSafetyTaskTmpl)
				m.Get("", reqRepoCloudBrainReader, repo.GetAiSafetyTask)
				m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.StopAiSafetyTask)
				m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.DelAiSafetyTask)
			})
			m.Get("/create_gpu", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.AiSafetyCreateForGetGPU)
			m.Get("/create_npu", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.AiSafetyCreateForGetNPU)
		}, context.RepoRef())

		m.Group("/debugjob", func() {
			m.Get("", reqRepoCloudBrainReader, repo.DebugJobIndex)
		}, context.RepoRef())

		m.Group("/modelarts", func() {
			m.Group("/notebook", func() {
				m.Group("/:id", func() {
					m.Get("", reqRepoCloudBrainReader, repo.NotebookShow)
					m.Get("/debug", cloudbrain.AdminOrJobCreaterRight, repo.NotebookDebug2)
					m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.NotebookStop)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.NotebookDel)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.NotebookNew)
			})

			m.Group("/train-job", func() {
				m.Get("", reqRepoCloudBrainReader, repo.TrainJobIndex)
				m.Group("/:jobid", func() {
					m.Get("", reqRepoCloudBrainReader, repo.TrainJobShow)
					m.Post("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.TrainJobStop)
					m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.TrainJobDel)
					m.Get("/model_download", cloudbrain.AdminOrJobCreaterRightForTrain, repo.ModelDownload)
					m.Get("/download_multi_model", cloudbrain.AdminOrJobCreaterRightForTrain, repo.MultiModelDownload)
					m.Get("/download_log_file", cloudbrain.AdminOrJobCreaterRightForTrain, repo.TrainJobDownloadLogFile)
					m.Get("/create_version", reqWechatBind, cloudbrain.AdminOrJobCreaterRightForTrain, context.PointAccount(), repo.TrainJobNewVersion)
					m.Post("/create_version", reqWechatBind, cloudbrain.AdminOrJobCreaterRightForTrain, bindIgnErr(auth.CreateModelArtsTrainJobForm{}), context.PointAccount(), repo.TrainJobCreateVersion)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.TrainJobNew)

				m.Get("/para-config-list", reqRepoCloudBrainReader, repo.TrainJobGetConfigList)
			})

			m.Group("/inference-job", func() {
				m.Get("", reqRepoCloudBrainReader, repo.InferenceJobIndex)
				m.Group("/:jobid", func() {
					m.Get("", reqRepoCloudBrainReader, repo.InferenceJobShow)
					m.Get("/result_download", cloudbrain.AdminOrJobCreaterRightForTrain, repo.ResultDownload)
					m.Get("/downloadall", cloudbrain.AdminOrJobCreaterRightForTrain, repo.DownloadMultiResultFile)
				})
				m.Get("/create", reqWechatBind, reqRepoCloudBrainWriter, context.PointAccount(), repo.InferenceJobNew)
			})
		}, context.RepoRef())

		m.Group("/blockchain", func() {
			m.Get("", repo.BlockChainIndex)
		}, context.RepoRef())
	}, ignSignIn, context.LowLimiter(), context.RepoAssignment(), context.UnitTypes())

	m.Group("/:username/:reponame", func() {

		m.Group("/activity", func() {
			m.Get("", repo.Activity)
			m.Get("/:period", repo.Activity)
		}, context.RepoRef(), repo.MustBeNotEmpty, context.RequireRepoReaderOr(models.UnitTypePullRequests, models.UnitTypeIssues, models.UnitTypeReleases))

		m.Group("/raw", func() {
			m.Get("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.SingleDownload)
			m.Get("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.SingleDownload)
			m.Get("/commit/*", context.RepoRefByType(context.RepoRefCommit), repo.SingleDownload)
			m.Get("/blob/:sha", context.RepoRefByType(context.RepoRefBlob), repo.DownloadByID)
			// "/*" route is deprecated, and kept for backward compatibility
			m.Get("/*", context.RepoRefByType(context.RepoRefLegacy), repo.SingleDownload)
		}, repo.MustBeNotEmpty, reqRepoCodeReader)

		m.Group("/activity_author_data", func() {
			m.Get("", repo.ActivityAuthors)
			m.Get("/:period", repo.ActivityAuthors)
		}, context.RepoRef(), repo.MustBeNotEmpty, context.RequireRepoReaderOr(models.UnitTypeCode))

		m.Get("/archive/*", repo.MustBeNotEmpty, reqRepoCodeReader, repo.Download)

		m.Get("/status", reqRepoCodeReader, repo.Status)

		m.Group("/blob_excerpt", func() {
			m.Get("/:sha", repo.SetEditorconfigIfExists, repo.SetDiffViewStyle, repo.ExcerptBlob)
		}, repo.MustBeNotEmpty, context.RepoRef(), reqRepoCodeReader)

		m.Group("", func() {
			m.Get("/graph", repo.Graph)
		}, repo.MustBeNotEmpty, context.RepoRef(), reqRepoCodeReader)

		m.Group("", func() {
			m.Get("/forks", repo.Forks)
		}, context.RepoRef(), reqRepoCodeReader)

	}, ignSignIn, context.HighLimiter(), context.RepoAssignment(), context.UnitTypes())
	m.Group("/:username/:reponame", func() {

		m.Group("/wiki", func() {
			m.Get("/?:page", repo.Wiki)
			m.Get("/_pages", repo.WikiPages)
			m.Get("/:page/_revision", repo.WikiRevision)
			m.Get("/commit/:sha([a-f0-9]{7,40})$", repo.SetEditorconfigIfExists, repo.SetDiffViewStyle, repo.Diff)
			m.Get("/commit/:sha([a-f0-9]{7,40})\\.:ext(patch|diff)", repo.RawDiff)

			m.Group("", func() {
				m.Combo("/_new").Get(repo.NewWiki).
					Post(bindIgnErr(auth.NewWikiForm{}), repo.NewWikiPost)
				m.Combo("/:page/_edit").Get(repo.EditWiki).
					Post(bindIgnErr(auth.NewWikiForm{}), repo.EditWikiPost)
				m.Post("/:page/delete", repo.DeleteWikiPagePost)
			}, context.RepoMustNotBeArchived(), reqSignIn, reqRepoWikiWriter)
		}, repo.MustEnableWiki, context.RepoRef(), func(ctx *context.Context) {
			ctx.Data["PageIsWiki"] = true
		})

		m.Group("/wiki", func() {
			m.Get("/raw/*", repo.WikiRaw)
		}, repo.MustEnableWiki)

		m.Group("/branches", func() {
			m.Get("", repo.Branches)
		}, context.GuestHandler(false), repo.MustBeNotEmpty, context.RepoRef(), reqRepoCodeReader)

		m.Group("/pulls/:index", func() {
			m.Get(".diff", repo.DownloadPullDiff)
			m.Get(".patch", repo.DownloadPullPatch)
			m.Get("/commits", context.RepoRef(), repo.ViewPullCommits)
			m.Post("/merge", context.RepoMustNotBeArchived(), bindIgnErr(auth.MergePullRequestForm{}), repo.MergePullRequest)
			m.Post("/update", repo.UpdatePullRequest)
			m.Post("/cleanup", context.RepoMustNotBeArchived(), context.RepoRef(), repo.CleanUpPullRequest)
			m.Group("/files", func() {
				m.Get("", context.RepoRef(), repo.SetEditorconfigIfExists, repo.SetDiffViewStyle, repo.SetWhitespaceBehavior, repo.ViewPullFiles)
				m.Group("/reviews", func() {
					m.Post("/comments", bindIgnErr(auth.CodeCommentForm{}), repo.CreateCodeComment)
					m.Post("/submit", bindIgnErr(auth.SubmitReviewForm{}), repo.SubmitReview)
				}, context.RepoMustNotBeArchived())
			})
		}, repo.MustAllowPulls)

		m.Group("/commits", func() {
			m.Get("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.RefCommits)
			m.Get("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.RefCommits)
			m.Get("/commit/*", context.RepoRefByType(context.RepoRefCommit), repo.RefCommits)
			// "/*" route is deprecated, and kept for backward compatibility
			m.Get("/*", context.RepoRefByType(context.RepoRefLegacy), repo.RefCommits)
		}, repo.MustBeNotEmpty, reqRepoCodeReader)
		m.Group("/media", func() {
			m.Get("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.SingleDownloadOrLFS)
			m.Get("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.SingleDownloadOrLFS)
			m.Get("/commit/*", context.RepoRefByType(context.RepoRefCommit), repo.SingleDownloadOrLFS)
			m.Get("/blob/:sha", context.RepoRefByType(context.RepoRefBlob), repo.DownloadByIDOrLFS)
			// "/*" route is deprecated, and kept for backward compatibility
			m.Get("/*", context.RepoRefByType(context.RepoRefLegacy), repo.SingleDownloadOrLFS)
		}, repo.MustBeNotEmpty, reqRepoCodeReader)

		m.Group("/blame", func() {
			m.Get("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.RefBlame)
			m.Get("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.RefBlame)
			m.Get("/commit/*", context.RepoRefByType(context.RepoRefCommit), repo.RefBlame)
		}, repo.MustBeNotEmpty, reqRepoCodeReader)

		m.Group("", func() {
			m.Get("/commit/:sha([a-f0-9]{7,40})$", reqSignIn, repo.SetEditorconfigIfExists, repo.SetDiffViewStyle, repo.Diff)
		}, repo.MustBeNotEmpty, context.RepoRef(), reqRepoCodeReader)

		m.Group("/src", func() {
			m.Get("/tag/*", context.RepoRefByType(context.RepoRefTag), repo.Home)
			m.Get("/commit/*", reqSignIn, context.RepoRefByType(context.RepoRefCommit), repo.Home)
			// "/*" route is deprecated, and kept for backward compatibility
			m.Get("/*", context.RepoRefByType(context.RepoRefLegacy), repo.Home)
		}, repo.SetEditorconfigIfExists)
		m.Get("/commit/:sha([a-f0-9]{7,40})\\.:ext(patch|diff)",
			reqSignIn, repo.MustBeNotEmpty, reqRepoCodeReader, repo.RawDiff)
	}, ignSignIn, context.GuestHandler(false), context.RepoAssignment(), context.UnitTypes())
	m.Group("/:username/:reponame", func() {
		m.Group("/src", func() {
			m.Get("/branch/*", context.RepoRefByType(context.RepoRefBranch), repo.Home)
		}, repo.SetEditorconfigIfExists)

	}, ignSignIn, context.GuestHandler(true), context.RepoAssignment(), context.UnitTypes())
	m.Group("/:username/:reponame", func() {
		m.Get("/stars", repo.Stars)
		m.Get("/watchers", repo.Watchers)
		m.Get("/search", reqRepoCodeReader, repo.Search)
	}, ignSignIn, context.LowLimiter(), context.RepoAssignment(), context.RepoRef(), context.UnitTypes())

	m.Group("/:username", func() {
		m.Group("/:reponame", func() {
			m.Get("", repo.SetEditorconfigIfExists, repo.Home)
			m.Get("\\.git$", repo.SetEditorconfigIfExists, repo.Home)
		}, ignSignIn, context.HighLimiter(), context.RepoAssignment(), context.RepoRef(), context.UnitTypes())

		m.Group("/:reponame", func() {
			m.Group("\\.git/info/lfs", func() {
				m.Post("/objects/batch", lfs.BatchHandler)
				m.Get("/objects/:oid/:filename", lfs.ObjectOidHandler)
				m.Any("/objects/:oid", lfs.ObjectOidHandler)
				m.Post("/objects", lfs.PostHandler)
				m.Post("/verify", lfs.VerifyHandler)
				m.Group("/locks", func() {
					m.Get("/", lfs.GetListLockHandler)
					m.Post("/", lfs.PostLockHandler)
					m.Post("/verify", lfs.VerifyLockHandler)
					m.Post("/:lid/unlock", lfs.UnLockHandler)
				})
				m.Any("/*", func(ctx *context.Context) {
					ctx.NotFound("", nil)
				})
			}, ignSignInAndCsrf)
			m.Any("/*", ignSignInAndCsrf, repo.HTTP)
			m.Head("/tasks/trigger", repo.TriggerTask)
		})
	})
	// ***** END: Repository *****
	m.Group("/notifications", func() {
		m.Get("", user.Notifications)
		m.Post("/status", user.NotificationStatusPost)
		m.Post("/purge", user.NotificationPurgePost)
	}, reqSignIn)

	m.Group("/reward/point", func() {
		m.Get("", point.GetPointPage)
		m.Get("/rule", point.GetRulePage)
		m.Get("/rule/config", point.GetRuleConfig)
		m.Get("/account", point.GetPointAccount)
		m.Get("/record/list", point.GetPointRecordList)
	}, reqSignIn)

	m.Group("/resources", func() {
		m.Group("/queue", func() {
			m.Get("/centers", admin.GetResourceAiCenters)
		})
	})

	if setting.API.EnableSwagger {
		m.Get("/swagger.v1.json", templates.JSONRenderer(), routers.SwaggerV1Json)
	}

	var handlers []macaron.Handler
	if setting.CORSConfig.Enabled {
		handlers = append(handlers, cors.CORS(cors.Options{
			Scheme:           setting.CORSConfig.Scheme,
			AllowDomain:      setting.CORSConfig.AllowDomain,
			AllowSubdomain:   setting.CORSConfig.AllowSubdomain,
			Methods:          setting.CORSConfig.Methods,
			MaxAgeSeconds:    int(setting.CORSConfig.MaxAge.Seconds()),
			AllowCredentials: setting.CORSConfig.AllowCredentials,
		}))
	}
	handlers = append(handlers, ignSignIn)
	m.Group("/api", func() {
		apiv1.RegisterRoutes(m)
	}, handlers...)

	//secure api,
	m.Group("/secure", func() {
		m.Post("/user", binding.BindIgnErr(structs.CreateUserOption{}), secure.CreateUser)
	}, reqBasicAuth)

	m.Group("/api/internal", func() {
		// package name internal is ideal but Golang is not allowed, so we use private as package name.
		private.RegisterRoutes(m)
	})

	// robots.txt
	m.Get("/robots.txt", func(ctx *context.Context) {
		if setting.HasRobotsTxt {
			ctx.ServeFileContent(path.Join(setting.CustomPath, "robots.txt"))
		} else {
			ctx.NotFound("", nil)
		}
	})

	m.Get("/apple-touch-icon.png", func(ctx *context.Context) {
		ctx.Redirect(path.Join(setting.StaticURLPrefix, "img/apple-touch-icon.png"), 301)
	})

	// prometheus metrics endpoint
	if setting.Metrics.Enabled {
		c := metrics.NewCollector()
		prometheus.MustRegister(c)

		m.Get("/metrics", routers.Metrics)
	}

	// Beginner's Guide
	m.Group("/guide", func() {
		m.Get("/create_dataset", reqSignIn, guide.GetCreateDataset)
		m.Get("/create_model", reqSignIn, guide.GetCreateModel)
	})

	m.Group("/pipeline", func() {
		m.Get("", reqSignIn, repo.PipelineMenu)
	})

	m.Group("/:username/:reponame", func() {
		m.Get("/pipeline", reqSignIn, repo.PipelineIndex)
		m.Get("/pipeline/template", reqSignIn, repo.TemplateIndex)
		m.Get("/pipeline/comparison", reqSignIn, repo.ComparisonIndex)
	}, ignSignIn, context.LowLimiter(), context.RepoAssignment(), context.RepoRef(), context.UnitTypes())
	m.Group(pipeline.UrlGroup, func() {
		m.Any("/*", reqSignIn, pipeline.Redirect)
	})

	// Not found handler.
	m.NotFound(routers.NotFound)
}
