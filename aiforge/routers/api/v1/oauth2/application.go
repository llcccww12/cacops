package oauth2

import (
	"fmt"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/oauth2_service"
)

func GetApplicaitonList(ctx *context.APIContext) {
	scopeList := ctx.Query("scopes")
	var scopes = make([]models.GrantScope, 0)
	if scopeList != "" {
		for _, val := range strings.Split(scopeList, ",") {
			if val != "" {
				scopes = append(scopes, models.GrantScope(val))
			}
		}
	}
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")
	q := ctx.Query("q")
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	verifyFlag := ctx.QueryInt("verify_flag")
	req := models.SearchOauth2ApplicationReq{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Keyword:    q,
		VerifyFlag: verifyFlag,
		ScopeList:  scopes,
	}
	list, total, err := oauth2_service.GetOauth2ApplicationList(req)
	if err != nil {
		log.Error("GetApplicaitonList GetOauth2ApplicationList err.req=%+v err=%v")
		ctx.JSON(200, response.OuterTrBizError(err, ctx.Locale))
		return
	}

	m := make(map[string]interface{}, 0)
	m["list"] = list
	m["total"] = total
	m["page"] = page
	m["page_size"] = pageSize
	ctx.JSON(200, response.OuterSuccessWithData(m))

}

func EditApplicaiton(ctx *context.APIContext, req models.EditOauth2ApplicationReq) {
	if req.ID <= 0 {
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}
	var scopes = make([]models.GrantScope, 0)
	if req.ScopeList != "" {
		for _, val := range strings.Split(req.ScopeList, ",") {
			if val != "" {
				scopes = append(scopes, models.GrantScope(val))
			}
		}
	}
	if len(scopes) == 0 {
		scopes = models.DefaultScopes
	}
	//未认证时不能设置非默认scope
	if req.VerifyFlag != models.ApplicationVerified {
		for _, scope := range scopes {
			find := false
			for _, d := range models.DefaultScopes {
				if d == scope {
					find = true
				}
			}
			if !find {
				ctx.JSON(200, response.OuterBizError(response.NewBizError(fmt.Errorf("Can not add these scopes when application is not verified"))))
				return
			}
		}
	}

	err := models.UpdateApplicationVerifyStatusAndScopes(models.OAuth2Application{
		ID:            req.ID,
		AllowedScopes: scopes,
		VerifyFlag:    req.VerifyFlag,
	})
	if err != nil {
		log.Error("EditApplicaiton UpdateApplicationVerifyStatusAndScopes err.req=%+v err=%v")
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	ctx.JSON(200, response.OuterSuccess())

}
