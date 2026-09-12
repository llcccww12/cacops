// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2020 The Gitea Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"
	"strings"

	"code.gitea.io/gitea/modules/auth/sso"
	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/api/v1/utils"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/role"

	"github.com/unknwon/com"
)

// Search search users
func Search(ctx *context.APIContext) {
	// swagger:operation GET /users/search user userSearch
	// ---
	// summary: Search for users
	// produces:
	// - application/json
	// parameters:
	// - name: q
	//   in: query
	//   description: keyword
	//   type: string
	// - name: uid
	//   in: query
	//   description: ID of the user to search for
	//   type: integer
	//   format: int64
	// - name: page
	//   in: query
	//   description: page number of results to return (1-based)
	//   type: integer
	// - name: limit
	//   in: query
	//   description: page size of results, maximum page size is 50
	//   type: integer
	// responses:
	//   "200":
	//     description: "SearchResults of a successful search"
	//     schema:
	//       type: object
	//       properties:
	//         ok:
	//           type: boolean
	//         data:
	//           type: array
	//           items:
	//             "$ref": "#/definitions/User"

	opts := &models.SearchUserOptions{
		Keyword:     strings.Trim(ctx.Query("q"), " "),
		UID:         com.StrTo(ctx.Query("uid")).MustInt64(),
		Type:        models.UserTypeIndividual,
		ListOptions: utils.GetListOptions(ctx),
	}

	users, _, err := models.SearchUsers(opts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	results := make([]*api.User, len(users))
	for i := range users {
		results[i] = convert.ToUser(users[i], ctx.IsSigned, ctx.User != nil && ctx.User.IsAdmin)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ok":   true,
		"data": results,
	})
}

// GetInfo get user's information
func GetInfo(ctx *context.APIContext) {
	// swagger:operation GET /users/{username} user userGet
	// ---
	// summary: Get a user
	// produces:
	// - application/json
	// parameters:
	// - name: username
	//   in: path
	//   description: username of user to get
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/User"
	//   "404":
	//     "$ref": "#/responses/notFound"

	u, err := models.GetUserByName(ctx.Params(":username"))
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.NotFound()
		} else {
			ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, convert.ToUser(u, ctx.IsSigned, ctx.User != nil && (ctx.User.ID == u.ID || ctx.User.IsAdmin)))
}

// GetAuthenticatedUser get current user's information
func GetAuthenticatedUser(ctx *context.APIContext) {
	// swagger:operation GET /user user userGetCurrent
	// ---
	// summary: Get the authenticated user
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/User"
	var authed = false
	var scopes = make([]models.GrantScope, 0)
	value, exists := ctx.Data["AccessTokenAuthContext"]
	if exists {
		oauthCtx, ok := value.(*models.OAuthContext)
		if !ok {
			scopes = models.DefaultScopes
		} else {
			scopes = oauthCtx.Scopes
		}
	} else if ctx.IsSigned {
		authed = true
	}

	ctx.JSON(http.StatusOK, convert.ToAuthenticatedUser(ctx.User, scopes, authed))
}

// GetUserHeatmapData is the handler to get a users heatmap
func GetUserHeatmapData(ctx *context.APIContext) {
	// swagger:operation GET /users/{username}/heatmap user userGetHeatmapData
	// ---
	// summary: Get a user's heatmap
	// produces:
	// - application/json
	// parameters:
	// - name: username
	//   in: path
	//   description: username of user to get
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/UserHeatmapData"
	//   "404":
	//     "$ref": "#/responses/notFound"

	// Get the user to throw an error if it does not exist
	user, err := models.GetUserByName(ctx.Params(":username"))
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.Status(http.StatusNotFound)
		} else {
			ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
		}
		return
	}

	heatmap, err := models.GetUserHeatmapDataByUser(user)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetUserHeatmapDataByUser", err)
		return
	}
	ctx.JSON(http.StatusOK, heatmap)
}

func IsRewardPointAdmin(ctx *context.APIContext) {
	//isAdmin := role.UserHasRole(ctx.User.ID, models.RewardPointAdmin)
	isAdmin := role.UserHasOper(ctx.User.ID, role.ROLE_OPER_RewardPointAdmin)
	r := map[string]interface{}{}
	r["is_admin"] = isAdmin
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}
func IsSubscriber(ctx *context.APIContext) {
	r := map[string]interface{}{}
	r["isSubscriber"] = role.UserHasRole(ctx.User.ID, models.Subscriber)
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}

func HasAccess(ctx *context.APIContext) {
	appId := getAppClientId(ctx)
	if appId == "" {
		ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{"has_access": false}))
		return
	}
	operaName := "access_app_" + appId
	hasAccess := role.UserHasOper(ctx.User.ID, operaName)
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{"has_access": hasAccess}))
	return
}

func getAppClientId(ctx *context.APIContext) string {
	tokenSHA := ctx.Query("access_token")
	if len(tokenSHA) == 0 || !strings.Contains(tokenSHA, ".") {
		ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{"has_access": false}))
		return ""
	}
	applicationId := sso.GetApplicationIDFromAccessToken(tokenSHA)
	if applicationId == 0 {
		return ""
	}
	app, err := models.GetOAuth2ApplicationByID(applicationId)
	if err != nil {
		log.Error("GetOAuth2ApplicationByID err.%v", err)
		return ""
	}
	return app.ClientID
}
