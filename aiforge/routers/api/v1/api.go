// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package v1 OpenI API.
//
// This documentation describes the OpenI API.
//
//	Schemes: http, https
//	BasePath: /api/v1
//	Version: 1.1.1
//	License: MIT http://opensource.org/licenses/MIT
//
//	Consumes:
//	- application/json
//	- text/plain
//
//	Produces:
//	- application/json
//	- text/html
//
//	Security:
//	- BasicAuth :
//	- Token :
//	- AccessToken :
//	- AuthorizationHeaderToken :
//	- SudoParam :
//	- SudoHeader :
//
//	SecurityDefinitions:
//	BasicAuth:
//	     type: basic
//	Token:
//	     type: apiKey
//	     name: token
//	     in: query
//	AccessToken:
//	     type: apiKey
//	     name: access_token
//	     in: query
//	AuthorizationHeaderToken:
//	     type: apiKey
//	     name: Authorization
//	     in: header
//	     description: API tokens must be prepended with "token" followed by a space.
//	SudoParam:
//	     type: apiKey
//	     name: sudo
//	     in: query
//	     description: Sudo API request as the user provided as the key. Admin privileges are required.
//	SudoHeader:
//	     type: apiKey
//	     name: Sudo
//	     in: header
//	     description: Sudo API request as the user provided as the key. Admin privileges are required.
//
// swagger:meta
package v1

import (
	"net/http"
	"strings"

	"code.gitea.io/gitea/routers/api/v1/official"

	"code.gitea.io/gitea/routers/api/v1/oauth2"
	"code.gitea.io/gitea/routers/api/v1/pay_computility"
	"code.gitea.io/gitea/routers/private"

	"code.gitea.io/gitea/routers/api/v1/ai_task_template"

	"code.gitea.io/gitea/routers/api/v1/sdk"

	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/routers/api/v1/access"
	"code.gitea.io/gitea/routers/api/v1/aimodel"
	"code.gitea.io/gitea/routers/api/v1/dataset"
	"code.gitea.io/gitea/routers/api/v1/dync_config"
	"code.gitea.io/gitea/routers/api/v1/modelscope"
	"code.gitea.io/gitea/routers/api/v1/monitor"
	"code.gitea.io/gitea/routers/api/v1/upload"

	"code.gitea.io/gitea/routers/api/v1/sd"

	"code.gitea.io/gitea/routers/api/v1/hf_model"

	"code.gitea.io/gitea/routers/resources"

	"code.gitea.io/gitea/routers/reward/point"

	"code.gitea.io/gitea/services/memory"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/entity"

	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"

	"code.gitea.io/gitea/routers/ai_task"

	"code.gitea.io/gitea/services/ai_task_service/task"

	"code.gitea.io/gitea/routers/api/v1/finetune"

	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/routers/api/v1/tech"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/api/v1/admin"
	"code.gitea.io/gitea/routers/api/v1/llm_chat"
	"code.gitea.io/gitea/routers/api/v1/misc"
	"code.gitea.io/gitea/routers/api/v1/notify"
	"code.gitea.io/gitea/routers/api/v1/org"
	"code.gitea.io/gitea/routers/api/v1/repo"
	"code.gitea.io/gitea/routers/api/v1/storage"
	_ "code.gitea.io/gitea/routers/api/v1/swagger" // for swagger generation
	"code.gitea.io/gitea/routers/api/v1/user"
	"code.gitea.io/gitea/routers/authentication"
	repo_ext "code.gitea.io/gitea/routers/repo"

	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

func sudo() macaron.Handler {
	return func(ctx *context.APIContext) {
		sudo := ctx.Query("sudo")
		if len(sudo) == 0 {
			sudo = ctx.Req.Header.Get("Sudo")
		}

		if len(sudo) > 0 {
			if ctx.IsSigned && ctx.User.IsAdmin {
				user, err := models.GetUserByName(sudo)
				if err != nil {
					if models.IsErrUserNotExist(err) {
						ctx.NotFound()
					} else {
						ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
					}
					return
				}
				log.Trace("Sudo from (%s) to: %s", ctx.User.Name, user.Name)
				ctx.User = user
			} else {
				ctx.JSON(http.StatusForbidden, map[string]string{
					"message": "Only administrators allowed to sudo.",
				})
				return
			}
		}
	}
}

func reqAITaskInRepo() macaron.Handler {
	return func(ctx *context.APIContext) {
		if ctx.Repo == nil {
			ctx.Context.Error(http.StatusUnauthorized)
			return
		}
		id := ctx.QueryInt64("id")
		if id <= 0 {
			ctx.Context.Error(http.StatusUnauthorized)
			return
		}
		t, err := models.GetCloudbrainByCloudbrainID(id)
		if err != nil {
			ctx.Context.Error(http.StatusUnauthorized)
			return
		}
		if t.RepoID != ctx.Repo.Repository.ID {
			ctx.Context.Error(http.StatusUnauthorized)
			return
		}
	}
}

func repoAssignmentRepoCanNull() macaron.Handler {
	return func(ctx *context.APIContext) {
		userName := ctx.Params(":username")
		repoName := ctx.Params(":reponame")

		var (
			owner *models.User
			err   error
		)
		if userName != "-" && repoName != "-" {

			// Check if the user is the same as the repository owner.
			if ctx.IsSigned && ctx.User.LowerName == strings.ToLower(userName) {
				owner = ctx.User
			} else {
				owner, err = models.GetUserByName(userName)
				if err != nil {
					if models.IsErrUserNotExist(err) {
						ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
					} else {
						ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
					}
					return
				}
			}
			ctx.Repo.Owner = owner

			// Get repository.
			repo, err := models.GetRepositoryByName(owner.ID, repoName)
			if err != nil {
				if models.IsErrRepoNotExist(err) {
					redirectRepoID, err := models.LookupRepoRedirect(owner.ID, repoName)
					if err == nil {
						context.RedirectToRepo(ctx.Context, redirectRepoID)
					} else if models.IsErrRepoRedirectNotExist(err) {
						ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
					} else {
						ctx.Error(http.StatusInternalServerError, "LookupRepoRedirect", err)
					}
				} else {
					ctx.Error(http.StatusInternalServerError, "GetRepositoryByName", err)
				}
				return
			}

			repo.Owner = owner
			ctx.Repo.Repository = repo

			ctx.Repo.Permission, err = models.GetUserRepoPermission(repo, ctx.User)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, "GetUserRepoPermission", err)
				return
			}

			if !ctx.Repo.HasAccess() {
				ctx.NotFound()
				return
			}
		}
	}

}

func repoAssignment() macaron.Handler {
	return func(ctx *context.APIContext) {
		userName := ctx.Params(":username")
		repoName := ctx.Params(":reponame")

		var (
			owner *models.User
			err   error
		)

		// Check if the user is the same as the repository owner.
		if ctx.IsSigned && ctx.User.LowerName == strings.ToLower(userName) {
			owner = ctx.User
		} else {
			owner, err = models.GetUserByName(userName)
			if err != nil {
				if models.IsErrUserNotExist(err) {
					ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
				} else {
					ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
				}
				return
			}
		}
		ctx.Repo.Owner = owner

		// Get repository.
		repo, err := models.GetRepositoryByName(owner.ID, repoName)
		if err != nil {
			if models.IsErrRepoNotExist(err) {
				redirectRepoID, err := models.LookupRepoRedirect(owner.ID, repoName)
				if err == nil {
					context.RedirectToRepo(ctx.Context, redirectRepoID)
				} else if models.IsErrRepoRedirectNotExist(err) {
					ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
				} else {
					ctx.Error(http.StatusInternalServerError, "LookupRepoRedirect", err)
				}
			} else {
				ctx.Error(http.StatusInternalServerError, "GetRepositoryByName", err)
			}
			return
		}

		repo.Owner = owner
		ctx.Repo.Repository = repo

		ctx.Repo.Permission, err = models.GetUserRepoPermission(repo, ctx.User)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "GetUserRepoPermission", err)
			return
		}

		if !ctx.Repo.HasAccess() {
			ctx.NotFound()
			return
		}
	}
}

func accessAssignment() macaron.Handler {
	return func(ctx *context.APIContext) {
		var (
			owner *models.User
			err   error
		)
		subjectID := ctx.Query("subject_id")
		subjectType := ctx.QueryInt("subject_type")

		subjectContext, err := models.GetSubjectContext(subjectID, models.SubjectType(subjectType))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Subject not found")
			return
		}
		if ctx.IsSigned && ctx.User.ID == subjectContext.OwnerID {
			owner = ctx.User
			subjectContext.Owner = owner
		} else {
			err = subjectContext.GetOwner()
			if err != nil {
				if models.IsErrUserNotExist(err) {
					ctx.JSON(http.StatusInternalServerError, "owner not exists")
				} else {
					ctx.JSON(http.StatusInternalServerError, "")
				}
				return
			}
			owner = subjectContext.Owner
		}

		ctx.Context.AccessContext.SubjectAccessContext = subjectContext

		ctx.AccessContext.Permission, err = models.GetUserSubjectPermission(subjectContext, ctx.User)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "GetUserSubjectPermission", err)
			return
		}

		if !ctx.AccessContext.HasAccess() {
			ctx.JSON(http.StatusNotFound, "")
			return
		}
	}
}

func datasetAssignment() macaron.Handler {
	return func(ctx *context.APIContext) {
		var (
			owner          *models.User
			err            error
			subjectContext *models.SubjectAccessContext
		)
		datasetID := ctx.Query("dataset_id")
		datasetNameStr := ctx.Query("dataset_name")
		if datasetID == "" && datasetNameStr == "" {
			ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
			return
		}
		if datasetID != "" {
			subjectContext, err = models.GetSubjectContext(datasetID, models.DatasetSubject)
		}
		if subjectContext == nil {
			if datasetNameStr == "" {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			parts := strings.Split(datasetNameStr, "/")
			if len(parts) != 2 {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}

			ownerName := strings.TrimSpace(parts[0])
			datasetName := strings.TrimSpace(parts[1])
			if ownerName == "" || datasetName == "" {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			dataset, err := models.GetDatasetRegistryByOwnerNameAndDatasetName(ownerName, datasetName)
			if err != nil {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			subjectContext = dataset.ConvertSubjectAccessContext()
		}

		if ctx.IsSigned && ctx.User.ID == subjectContext.OwnerID {
			owner = ctx.User
			subjectContext.Owner = owner
		} else {
			err = subjectContext.GetOwner()
			if err != nil {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			owner = subjectContext.Owner
		}

		ctx.Context.AccessContext.SubjectAccessContext = subjectContext

		ctx.AccessContext.Permission, err = models.GetUserSubjectPermission(subjectContext, ctx.User)
		if err != nil {
			ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
			return
		}
	}
}

func aimodelAssignment() macaron.Handler {
	return func(ctx *context.APIContext) {
		var (
			owner          *models.User
			err            error
			subjectContext *models.SubjectAccessContext
		)
		aimodelID := ctx.Query("aimodel_id")
		aimodelNameStr := ctx.Query("aimodel_name")
		if aimodelID == "" && aimodelNameStr == "" {
			ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
			return
		}
		if aimodelID != "" {
			subjectContext, err = models.GetSubjectContext(aimodelID, models.AimodelSubject)
		}
		if subjectContext == nil {
			if aimodelNameStr == "" {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			parts := strings.Split(aimodelNameStr, "/")
			if len(parts) != 2 {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}

			ownerName := strings.TrimSpace(parts[0])
			aimodelName := strings.TrimSpace(parts[1])
			if ownerName == "" || aimodelName == "" {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			aimodel, err := models.GetModelByOwnerNameAndAimodelName(ownerName, aimodelName)
			if err != nil {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			subjectContext = aimodel.ConvertSubjectAccessContext()
		}

		if ctx.IsSigned && ctx.User.ID == subjectContext.OwnerID {
			owner = ctx.User
			subjectContext.Owner = owner
		} else {
			err = subjectContext.GetOwner()
			if err != nil {
				ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
				return
			}
			owner = subjectContext.Owner
		}

		ctx.Context.AccessContext.SubjectAccessContext = subjectContext
		ctx.AccessContext.Permission, err = models.GetUserSubjectPermission(subjectContext, ctx.User)
		if err != nil {
			ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
			return
		}
	}
}

func repoAssignmentWithNoAccess() macaron.Handler {
	return func(ctx *context.APIContext) {
		userName := ctx.Params(":username")
		repoName := ctx.Params(":reponame")

		var (
			owner *models.User
			err   error
		)

		// Check if the user is the same as the repository owner.
		if ctx.IsSigned && ctx.User.LowerName == strings.ToLower(userName) {
			owner = ctx.User
		} else {
			owner, err = models.GetUserByName(userName)
			if err != nil {
				if models.IsErrUserNotExist(err) {
					ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
				} else {
					ctx.Error(http.StatusInternalServerError, "GetUserByName", err)
				}
				return
			}
		}
		ctx.Repo.Owner = owner

		// Get repository.
		repo, err := models.GetRepositoryByName(owner.ID, repoName)
		if err != nil {
			if models.IsErrRepoNotExist(err) {
				redirectRepoID, err := models.LookupRepoRedirect(owner.ID, repoName)
				if err == nil {
					context.RedirectToRepo(ctx.Context, redirectRepoID)
				} else if models.IsErrRepoRedirectNotExist(err) {
					ctx.NotFound(ctx.Tr("repo.RepoNotFound", userName+"/"+repoName))
				} else {
					ctx.Error(http.StatusInternalServerError, "LookupRepoRedirect", err)
				}
			} else {
				ctx.Error(http.StatusInternalServerError, "GetRepositoryByName", err)
			}
			return
		}

		repo.Owner = owner
		ctx.Repo.Repository = repo

		ctx.Repo.Permission, err = models.GetUserRepoPermission(repo, ctx.User)

	}
}

// Contexter middleware already checks token for user sign in process.
func reqToken() macaron.Handler {
	return func(ctx *context.APIContext) {
		if true == ctx.Data["IsApiToken"] {
			return
		}
		if ctx.Context.IsBasicAuth {
			ctx.CheckForOTP()
			return
		}
		if ctx.IsSigned {
			ctx.RequireCSRF()
			return
		}
		ctx.Context.Error(http.StatusUnauthorized)
	}
}

func reqBasicAuth() macaron.Handler {
	return func(ctx *context.APIContext) {
		if !ctx.Context.IsBasicAuth {
			ctx.Context.Error(http.StatusUnauthorized)
			return
		}
		ctx.CheckForOTP()
	}
}

// reqSiteAdmin user should be the site admin
func reqSiteAdmin() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsUserSiteAdmin() {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

// reqTechAdmin user should be the tech program admin
// func HasRole(roleType models.RoleType) macaron.Handler {
// 	return func(ctx *context.Context) {
// 		if !ctx.IsSigned || !role.UserHasRole(ctx.User.ID, roleType) {
// 			ctx.Error(http.StatusForbidden)
// 			return
// 		}
// 	}
// }

// reqTechAdmin user should be the tech program admin
func HasOperRole(operName string) macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsSigned || !role.UserHasOper(ctx.User.ID, operName) {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

func HasOperRoleOrAdmin(operName string) macaron.Handler {
	return func(ctx *context.Context) {
		if ctx.IsSigned && ctx.User != nil && (ctx.User.IsAdmin || role.UserHasOper(ctx.User.ID, operName)) {
			return
		}
		ctx.Error(http.StatusForbidden)
	}
}

// reqOwner user should be the owner of the repo or site admin.
func reqOwner() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsUserRepoOwner() && !ctx.IsUserSiteAdmin() {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

// reqAdmin user should be an owner or a collaborator with admin write of a repository, or site admin
func reqAdmin() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsUserRepoAdmin() && !ctx.IsUserSiteAdmin() {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

// reqRepoWriter user should have a permission to write to a repo, or be a site admin
func reqRepoWriter(unitTypes ...models.UnitType) macaron.Handler {
	return func(ctx *context.Context) {
		log.Info("reqRepoWriter enter")
		if !ctx.IsUserRepoWriter(unitTypes) && !ctx.IsUserRepoAdmin() && !ctx.IsUserSiteAdmin() {
			log.Info("!ctx.IsUserRepoWriter(unitTypes) && !ctx.IsUserRepoAdmin() && !ctx.IsUserSiteAdmin() forbidden")
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

// reqAdminOrAITaskCreator
func reqAdminOrAITaskCreator() macaron.Handler {

	return func(ctx *context.Context) {

		var id = ctx.Params(":id")
		if id == "" {
			id = ctx.Query("id")
		}
		job, err := cloudbrain.GetCloudBrainByIdOrJobId(id, "id")
		if err != nil {
			log.Error("GetCloudbrainByID failed:%v", err.Error())
			ctx.Error(http.StatusForbidden)
		}
		ctx.Cloudbrain = job
		if !isAdminOrJobCreater(ctx, job, err) {
			if err != nil {
				log.Error("!isAdminOrJobCreater error:%v", err)
			}

			ctx.Error(http.StatusForbidden)
		}

	}
}

// reqAdminOrOwnerAITaskCreator
func reqAdminOrOwnerAITaskCreator() macaron.Handler {

	return func(ctx *context.Context) {

		var id = ctx.Params(":id")
		if id == "" {
			id = ctx.Query("id")
		}
		job, err := cloudbrain.GetCloudBrainByIdOrJobId(id, "id")
		if err != nil {
			log.Error("GetCloudbrainByID failed:%v", err.Error())
			ctx.Error(http.StatusForbidden)
		}
		ctx.Cloudbrain = job
		if !isAdminOrOwnerOrJobCreator(ctx, job, err) {
			if err != nil {
				log.Error("!isAdminOrJobCreater error:%v", err)
			}
			ctx.Error(http.StatusForbidden)
		}

	}
}

func reqAccessAdmin() macaron.Handler {

	return func(ctx *context.Context) {
		if !ctx.IsSigned {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
			return
		}
		if isAdminOrSubjectOwner(ctx) {
			return
		}
		if !ctx.AccessContext.IsAdmin() {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
			return
		}
	}
}

func reqSubjectOwner() macaron.Handler {

	return func(ctx *context.Context) {
		if !ctx.IsSigned {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
			return
		}
		if !isAdminOrSubjectOwner(ctx) {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
			return
		}
	}
}

func isAdminOrSubjectOwner(ctx *context.Context) bool {
	if !ctx.IsSigned {
		return false
	}
	if ctx.User.IsAdmin || ctx.User.ID == ctx.AccessContext.SubjectAccessContext.OwnerID {
		return true
	}
	acessCtx := ctx.AccessContext.SubjectAccessContext

	if err := acessCtx.GetOwner(); err != nil {
		return false
	}

	if acessCtx.Owner.IsOrganization() {
		isOwner, err := acessCtx.Owner.IsOwnedBy(ctx.User.ID)
		if err != nil {
			return false
		} else if isOwner {
			return true
		}
	}
	return false
}

func reqAccessRead() macaron.Handler {
	return func(ctx *context.Context) {
		//公开资源游客和所有用户都能访问
		if !ctx.AccessContext.IsPrivate {
			return
		}
		if isAdminOrSubjectOwner(ctx) {
			return
		}
		if !ctx.AccessContext.IsReader() {
			ctx.JSON(200, response.OuterTrBizError(response.NOT_EXISTS_OR_NO_PERMISSION, ctx.Locale))
			return
		}
	}
}

func reqAccessWrite() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsSigned {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
			return
		}
		if isAdminOrSubjectOwner(ctx) {
			return
		}
		if !ctx.AccessContext.IsWriter() {
			ctx.JSON(http.StatusForbidden, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
		}
	}
}

func isAdminOrOwnerOrJobCreator(ctx *context.Context, job *models.Cloudbrain, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {

		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserRepoOwner() || ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func isAdminOrJobCreater(ctx *context.Context, job *models.Cloudbrain, err error) bool {
	if !ctx.IsSigned {
		return false
	}
	if err != nil {
		return ctx.IsUserSiteAdmin()
	} else {
		return ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func reqWeChat() macaron.Handler {
	return func(ctx *context.Context) {
		if setting.WechatAuthSwitch && ctx.User.WechatOpenId == "" {
			ctx.JSON(http.StatusForbidden, models.BaseErrorMessageApi("settings.no_wechat_bind"))
			return
		}
	}
}

func reqWeChatStandard() macaron.Handler {
	return func(ctx *context.Context) {
		if setting.WechatAuthSwitch && ctx.User.WechatOpenId == "" {
			ctx.JSON(http.StatusOK, response.OuterTrBizError(response.WECHAT_NOT_BIND, ctx))
			return
		}
	}
}

// reqAnyRepoReader user should have any permission to read repository or permissions of site admin
func reqAnyRepoReader() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsUserRepoReaderAny() && !ctx.IsUserSiteAdmin() {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

// reqOrgOwnership user should be an organization owner, or a site admin
func reqOrgOwnership() macaron.Handler {
	return func(ctx *context.APIContext) {
		if ctx.Context.IsUserSiteAdmin() {
			return
		}

		var orgID int64
		if ctx.Org.Organization != nil {
			orgID = ctx.Org.Organization.ID
		} else if ctx.Org.Team != nil {
			orgID = ctx.Org.Team.OrgID
		} else {
			ctx.Error(http.StatusInternalServerError, "", "reqOrgOwnership: unprepared context")
			return
		}

		isOwner, err := models.IsOrganizationOwner(orgID, ctx.User.ID)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "IsOrganizationOwner", err)
			return
		} else if !isOwner {
			if ctx.Org.Organization != nil {
				ctx.Error(http.StatusForbidden, "", "Must be an organization owner")
			} else {
				ctx.NotFound()
			}
			return
		}
	}
}

// reqTeamMembership user should be an team member, or a site admin
func reqTeamMembership() macaron.Handler {
	return func(ctx *context.APIContext) {
		if ctx.Context.IsUserSiteAdmin() {
			return
		}
		if ctx.Org.Team == nil {
			ctx.Error(http.StatusInternalServerError, "", "reqTeamMembership: unprepared context")
			return
		}

		var orgID = ctx.Org.Team.OrgID
		isOwner, err := models.IsOrganizationOwner(orgID, ctx.User.ID)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "IsOrganizationOwner", err)
			return
		} else if isOwner {
			return
		}

		if isTeamMember, err := models.IsTeamMember(orgID, ctx.Org.Team.ID, ctx.User.ID); err != nil {
			ctx.Error(http.StatusInternalServerError, "IsTeamMember", err)
			return
		} else if !isTeamMember {
			isOrgMember, err := models.IsOrganizationMember(orgID, ctx.User.ID)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, "IsOrganizationMember", err)
			} else if isOrgMember {
				ctx.Error(http.StatusForbidden, "", "Must be a team member")
			} else {
				ctx.NotFound()
			}
			return
		}
	}
}

// reqOrgMembership user should be an organization member, or a site admin
func reqOrgMembership() macaron.Handler {
	return func(ctx *context.APIContext) {
		if ctx.Context.IsUserSiteAdmin() {
			return
		}

		var orgID int64
		if ctx.Org.Organization != nil {
			orgID = ctx.Org.Organization.ID
		} else if ctx.Org.Team != nil {
			orgID = ctx.Org.Team.OrgID
		} else {
			ctx.Error(http.StatusInternalServerError, "", "reqOrgMembership: unprepared context")
			return
		}

		if isMember, err := models.IsOrganizationMember(orgID, ctx.User.ID); err != nil {
			ctx.Error(http.StatusInternalServerError, "IsOrganizationMember", err)
			return
		} else if !isMember {
			if ctx.Org.Organization != nil {
				ctx.Error(http.StatusForbidden, "", "Must be an organization member")
			} else {
				ctx.NotFound()
			}
			return
		}
	}
}

func reqGitHook() macaron.Handler {
	return func(ctx *context.APIContext) {
		if !ctx.User.CanEditGitHook() {
			ctx.Error(http.StatusForbidden, "", "must be allowed to edit Git hooks")
			return
		}
	}
}

// reqRepoReader user should have specific read permission or be a repo admin or a site admin
func reqRepoReader(unitType models.UnitType) macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsUserRepoReaderSpecific(unitType) && !ctx.IsUserRepoAdmin() && !ctx.IsUserSiteAdmin() {
			ctx.Error(http.StatusForbidden)
			return
		}
	}
}

func reqRepoReaderCanNull(unitType models.UnitType) macaron.Handler {
	return func(ctx *context.Context) {
		if ctx.Repo.Repository != nil {

			if !ctx.IsUserRepoReaderSpecific(unitType) && !ctx.IsUserRepoAdmin() && !ctx.IsUserSiteAdmin() {
				ctx.Error(http.StatusForbidden)
				return
			}
		}
	}
}

func orgAssignment(args ...bool) macaron.Handler {
	var (
		assignOrg  bool
		assignTeam bool
	)
	if len(args) > 0 {
		assignOrg = args[0]
	}
	if len(args) > 1 {
		assignTeam = args[1]
	}
	return func(ctx *context.APIContext) {
		ctx.Org = new(context.APIOrganization)

		var err error
		if assignOrg {
			ctx.Org.Organization, err = models.GetOrgByName(ctx.Params(":org"))
			if err != nil {
				if models.IsErrOrgNotExist(err) {
					ctx.NotFound()
				} else {
					ctx.Error(http.StatusInternalServerError, "GetOrgByName", err)
				}
				return
			}
		}

		if assignTeam {
			ctx.Org.Team, err = models.GetTeamByID(ctx.ParamsInt64(":teamid"))
			if err != nil {
				if models.IsErrUserNotExist(err) {
					ctx.NotFound()
				} else {
					ctx.Error(http.StatusInternalServerError, "GetTeamById", err)
				}
				return
			}
		}
	}
}

func mustEnableIssues(ctx *context.APIContext) {
	if !ctx.Repo.CanRead(models.UnitTypeIssues) {
		if log.IsTrace() {
			if ctx.IsSigned {
				log.Trace("Permission Denied: User %-v cannot read %-v in Repo %-v\n"+
					"User in Repo has Permissions: %-+v",
					ctx.User,
					models.UnitTypeIssues,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			} else {
				log.Trace("Permission Denied: Anonymous user cannot read %-v in Repo %-v\n"+
					"Anonymous user in Repo has Permissions: %-+v",
					models.UnitTypeIssues,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			}
		}
		ctx.NotFound()
		return
	}
}

func mustAllowPulls(ctx *context.APIContext) {
	if !(ctx.Repo.Repository.CanEnablePulls() && ctx.Repo.CanRead(models.UnitTypePullRequests)) {
		if ctx.Repo.Repository.CanEnablePulls() && log.IsTrace() {
			if ctx.IsSigned {
				log.Trace("Permission Denied: User %-v cannot read %-v in Repo %-v\n"+
					"User in Repo has Permissions: %-+v",
					ctx.User,
					models.UnitTypePullRequests,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			} else {
				log.Trace("Permission Denied: Anonymous user cannot read %-v in Repo %-v\n"+
					"Anonymous user in Repo has Permissions: %-+v",
					models.UnitTypePullRequests,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			}
		}
		ctx.NotFound()
		return
	}
}

func mustEnableIssuesOrPulls(ctx *context.APIContext) {
	if !ctx.Repo.CanRead(models.UnitTypeIssues) &&
		!(ctx.Repo.Repository.CanEnablePulls() && ctx.Repo.CanRead(models.UnitTypePullRequests)) {
		if ctx.Repo.Repository.CanEnablePulls() && log.IsTrace() {
			if ctx.IsSigned {
				log.Trace("Permission Denied: User %-v cannot read %-v and %-v in Repo %-v\n"+
					"User in Repo has Permissions: %-+v",
					ctx.User,
					models.UnitTypeIssues,
					models.UnitTypePullRequests,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			} else {
				log.Trace("Permission Denied: Anonymous user cannot read %-v and %-v in Repo %-v\n"+
					"Anonymous user in Repo has Permissions: %-+v",
					models.UnitTypeIssues,
					models.UnitTypePullRequests,
					ctx.Repo.Repository,
					ctx.Repo.Permission)
			}
		}
		ctx.NotFound()
		return
	}
}

func mustEnableUserHeatmap(ctx *context.APIContext) {
	if !setting.Service.EnableUserHeatmap {
		ctx.NotFound()
		return
	}
}

func mustNotBeArchived(ctx *context.APIContext) {
	if ctx.Repo.Repository.IsArchived {
		ctx.NotFound()
		return
	}
}

// RegisterRoutes registers all v1 APIs routes to web application.
// FIXME: custom form error response
func RegisterRoutes(m *macaron.Macaron) {
	bind := binding.Bind

	if setting.API.EnableSwagger {
		m.Get("/swagger", misc.Swagger) // Render V1 by default
	}

	m.Group("/v1", func() {

		// for openi-sdk
		m.Group("/sdk", func() {
			m.Get("/notification", sdk.SendNotification)
		})

		//m.Group("/:username/:reponame", func() {
		//	m.Group("/sdk", func() {
		//		m.Get("/get_dataset", repo.ListDatasetFilesForSDK)
		//		m.Get("/download_dataset_file/:uuid", repo.GetAttachment)
		//		m.Get("/get_model", repo.ListModelFilesForSDK)
		//		m.Get("/download_model_file/:ID", repo.DownloadModelSingle)
		//		m.Get("/generate_model_download_code", ai_task.GenerateModelDownloadCode)
		//
		//		m.Get("/generate_model_upload_code", ai_task.GenerateModelUploadCode)
		//		m.Get("/generate_model_upload_file_code", ai_task.GenerateModelFileUploadCode)
		//		m.Get("/generate_dataset_download_code", ai_task.GenerateDatasetDownloadCode)
		//		m.Get("/generate_dataset_upload_code", ai_task.GenerateDatasetUploadCode)
		//	}, repoAssignment())
		//})

		m.Group("/:username/:reponame", func() {
			m.Group("/ai_task", func() {
				m.Post("agent/create", reqWeChatStandard(), reqRepoReaderCanNull(models.UnitTypeCode), bind(entity.CreateReq{}), ai_task.CreateAITask)
				m.Post("/create", reqWeChatStandard(), reqRepoReaderCanNull(models.UnitTypeCode), bind(entity.CreateReq{}), ai_task.CreateAITask)
				m.Get("/creation/required", reqWeChatStandard(), ai_task.GetCreationRequiredInfo)
				m.Get("/creation/image_by_spec", reqWeChatStandard(), ai_task.GetImageInfoBySelectedSpec)
				m.Post("/restart", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.RestartAITask)
				m.Get("/debug_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetNotebookUrl)
				m.Get("/visualize_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetVisualizeUrl)
				m.Get("/self_endpoint_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetSelfEndPointUrl)
				m.Get("/repo_branch", reqWeChatStandard(), reqRepoReader(models.UnitTypeCode), ai_task.GetCreationRequiredRepoInfo)
				m.Post("/output/reschedule", reqAdminOrAITaskCreator(), ai_task.RetryModelSchedule)
				m.Group("/template", func() {
					m.Group("/get_edit_address", func() {
						m.Get("", ai_task.GetTemplateEditAddress)
						m.Get("/branch/*", ai_task.GetTemplateEditAddress)
					}, reqRepoWriter(models.UnitTypeCloudBrain))
					m.Get("/get_edit_address", reqRepoWriter(models.UnitTypeCloudBrain), ai_task.GetTemplateEditAddress)
				})
			}, reqToken(), context.RepoRef())
			m.Group("/ai_task", func() {
				m.Get("", reqAdminOrAITaskCreator(), ai_task.GetAITaskInfo)
				m.Get("/brief", reqAdminOrAITaskCreator(), ai_task.GetAITaskBriefInfo)
				//m.Get("/list", reqRepoReader(models.UnitTypeCloudBrain), ai_task.GetAITaskList)
				m.Get("/operation_profile", reqAdminOrAITaskCreator(), ai_task.GetAITaskOperationProfile)
				m.Get("/log", reqAdminOrAITaskCreator(), ai_task.GetAITaskLog)
				m.Get("/log/download", reqAdminOrAITaskCreator(), ai_task.DownloadAITaskLog)
				m.Get("/node_info", reqAdminOrAITaskCreator(), ai_task.GetNodeInfo)
				m.Get("/output", reqAdminOrAITaskCreator(), ai_task.GetAITaskOutput)
				m.Get("/output/download", reqAdminOrAITaskCreator(), ai_task.DownloadOutputFile)
				m.Get("/output/download/all", reqAdminOrAITaskCreator(), ai_task.DownloadAllOutputFile)
				m.Get("/output/all", reqAdminOrAITaskCreator(), ai_task.GetAllAITaskOutput)
				m.Get("/resource_usage", reqAdminOrAITaskCreator(), ai_task.GetAITaskResourceUsage)
				m.Get("/loss", reqAdminOrAITaskCreator(), ai_task.GetAITaskLoss)
				m.Get("/eval_result", reqAdminOrAITaskCreator(), ai_task.GetEvalResult)
				m.Post("/eval_detail_result", reqAdminOrAITaskCreator(), bind(api.TaskDetailRequest{}), ai_task.GetEvalDetailResult)
				m.Post("/export2dataset", reqAdminOrAITaskCreator(), datasetAssignment(), reqAccessWrite(), bind(entity.ExportTaskResultReq{}), dataset.ExportTaskResult)

			})
		}, repoAssignmentRepoCanNull())
		m.Group("/:username/:reponame", func() {
			m.Group("/ai_task", func() {
				m.Post("/stop", reqAdminOrOwnerAITaskCreator(), ai_task.StopAITask)
				m.Post("/del", reqAdminOrOwnerAITaskCreator(), ai_task.DelAITask)
			}, reqToken())
		})

		// Agent 广场列表接口（无需鉴权，路径 /api/v1/ai_task/list）
		m.Group("/agent", func() {
			m.Get("/list", ai_task.GetAgentList)
		})

		m.Group("/ai_task", func() {
			m.Get("/creation/required", reqWeChatStandard(), ai_task.GetCreationRequiredInfo)
			m.Get("/creation/image_by_spec", reqWeChatStandard(), ai_task.GetImageInfoBySelectedSpec)
			m.Get("/generate_sdk_code", ai_task.GenerateSDKCode)
			m.Get("/my_list", ai_task.GetMyAITaskList)

			// Agent 智能体路由
			m.Group("/agent", func() {
				m.Get("/user_list", ai_task.GetUserAgentList) // 获取当前用户的智能体列表（带 job_status）

				// 状态管理路由
				m.Post("/:id/run", ai_task.RunAgent)
				m.Post("/:id/stop", ai_task.StopAgent)

				// 删除用户智能体关联（通过 id）
				m.Delete("/user/:id", ai_task.DeleteUserAgentByCloudBrainID)
			})

			// Agent 智能体路由
			m.Group("/agentAdmin/", func() {
				m.Post("/create", bind(task.CreateAgentReq{}), ai_task.CreateAgent)
				m.Put("/:id", bind(task.UpdateAgentReq{}), ai_task.UpdateAgent)
				m.Delete("/:id", ai_task.DeleteAgent)
				m.Put("/:id/status", bind(task.UpdateAgentStatusReq{}), ai_task.UpdateAgentStatus)
				m.Get("/records", ai_task.GetUserAgentRecordList) // 获取用户智能体使用历史记录

			}, reqSiteAdmin())

			m.Post("/batch_del", bind(entity.DeleteIDs{}), ai_task.BatchDelAITask)
			m.Post("/restart", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.RestartAITask)
			m.Get("/debug_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetNotebookUrl)
			m.Get("/visualize_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetVisualizeUrl)
			m.Get("/self_endpoint_url", reqWeChatStandard(), reqAdminOrAITaskCreator(), ai_task.GetSelfEndPointUrl)
			m.Post("/output/reschedule", reqAdminOrAITaskCreator(), ai_task.RetryModelSchedule)
			m.Get("", reqAdminOrAITaskCreator(), ai_task.GetAITaskInfo)
			m.Get("/brief", reqAdminOrAITaskCreator(), ai_task.GetAITaskBriefInfo)
			m.Get("/operation_profile", reqAdminOrAITaskCreator(), ai_task.GetAITaskOperationProfile)
			m.Get("/log", reqAdminOrAITaskCreator(), ai_task.GetAITaskLog)
			m.Get("/log/download", reqAdminOrAITaskCreator(), ai_task.DownloadAITaskLog)
			m.Get("/node_info", reqAdminOrAITaskCreator(), ai_task.GetNodeInfo)
			m.Get("/output", reqAdminOrAITaskCreator(), ai_task.GetAITaskOutput)
			m.Get("/output/download", reqAdminOrAITaskCreator(), ai_task.DownloadOutputFile)
			m.Get("/output/download/all", reqAdminOrAITaskCreator(), ai_task.DownloadAllOutputFile)
			m.Get("/output/all", reqAdminOrAITaskCreator(), ai_task.GetAllAITaskOutput)
			m.Get("/resource_usage", reqAdminOrAITaskCreator(), ai_task.GetAITaskResourceUsage)
			m.Get("/loss", reqAdminOrAITaskCreator(), ai_task.GetAITaskLoss)
			m.Get("/eval_result", reqAdminOrAITaskCreator(), ai_task.GetEvalResult)
			m.Post("/eval_detail_result", reqAdminOrAITaskCreator(), bind(api.TaskDetailRequest{}), ai_task.GetEvalDetailResult)
			m.Post("/export2dataset", reqAdminOrAITaskCreator(), datasetAssignment(), reqAccessWrite(), bind(entity.ExportTaskResultReq{}), dataset.ExportTaskResult)
			m.Post("/export2aimodel", reqAdminOrAITaskCreator(), reqToken(), bind(entity.CreateAimodelReq{}), aimodel.ExportTaskResult2NewAimodel)
			m.Post("/stop", reqAdminOrAITaskCreator(), ai_task.StopAITask)
			m.Post("/del", reqAdminOrAITaskCreator(), ai_task.DelAITask)

		}, reqToken())

		m.Group("/platform", func() {
			m.Get("/overview", user.Overview)
			m.Get("/official_config", official.OfficialConfig)
			m.Get("/action", user.OverViewAction)
			m.Get("/org", user.UserOrganization)
		}, reqToken())

		//aimodel
		m.Group("/aimodel", func() {
			m.Group("/create", func() {
				m.Post("", reqToken(), bind(entity.CreateAimodelReq{}), aimodel.CreateAimodel)
				m.Get("/available_users", aimodel.GetAimodelCreationAvailableUsers)
			}, reqToken())
			m.Group("/list", func() {
				m.Get("/public", aimodel.SearchAllPublicAimodels)
				m.Get("/owned", reqToken(), aimodel.SearchMyOwnedAimodels)
				m.Get("/collaborated", reqToken(), aimodel.SearchMyCollaboratedAimodels)
				m.Get("/collected", reqToken(), aimodel.SearchMyCollectedAimodels)
				m.Get("/accessible", reqToken(), aimodel.SearchMyAccessibleAimodels)
				m.Get("/involved", reqToken(), aimodel.SearchInvolvedAimodels)
			})
			m.Post("/batch_delete", reqToken(), aimodel.BatchDeleteAimodel)
			m.Get("/get_export_process", reqToken(), aimodel.GetExportAimodelByMsgId)
			m.Get("/tags", reqToken(), aimodel.GetUserAimodelLabels)
		})
		m.Group("/aimodel", func() {
			m.Combo("").Get(reqAccessRead(), aimodel.GetAimodel).Delete(reqSubjectOwner(), aimodel.DelAimodel)
			m.Post("/readme", reqAccessWrite(), bind(entity.AimodelReadmeReq{}), aimodel.PutAimodelReadme)
			m.Get("/readme", reqAccessRead(), aimodel.GetAimodelReadme)
			m.Get("/files", reqAccessRead(), aimodel.GetAimodelFileList)
			m.Get("/file", reqAccessRead(), aimodel.DownloadAimodelFile)
			m.Get("/file/meta", reqAccessRead(), aimodel.GetAimodelFileMeta)
			m.Delete("/file", reqAccessWrite(), aimodel.DeleteAimodelFile)
			m.Get("/download", reqAccessRead(), aimodel.DownloadAimodel)
			//m.Get("/download/meta", reqAccessRead(), aimodel.GetDownloadAimodelMeta)
			m.Post("/edit", reqAccessAdmin(), bind(entity.CreateAimodelReq{}), aimodel.EditAimodel)
			m.Combo("/collect").Put(reqAccessRead(), aimodel.CollectAimodel).Delete(aimodel.UnCollectAimodel)
			m.Get("/sdk_download_code", reqAccessRead(), ai_task.GenerateModelDownloadCode)
			m.Get("/evolution", reqAccessRead(), aimodel.AimodelEvolutionGraph)
			m.Get("/related", reqAccessRead(), aimodel.GetAimodelRelatedData)
		}, aimodelAssignment())

		//dataset
		m.Group("/dataset", func() {
			m.Group("/create", func() {
				m.Post("", reqToken(), bind(entity.CreateDatasetReq{}), dataset.CreateDataset)
				m.Get("/available_users", reqToken(), dataset.GetDatasetCreationAvailableUsers)
			}, reqToken())
			m.Group("/list", func() {
				m.Get("/public", dataset.SearchAllPublicDataset)
				m.Get("/owned", reqToken(), dataset.SearchMyOwnedDatasets)
				m.Get("/collaborated", reqToken(), dataset.SearchMyCollaboratedDatasets)
				m.Get("/collected", reqToken(), dataset.SearchMyCollectedDatasets)
				m.Get("/accessible", reqToken(), dataset.SearchMyAccessibleDatasets)
			})
			m.Post("/batch_delete", reqToken(), dataset.BatchDeleteDatasets)
			m.Get("/get_export_process", reqToken(), dataset.GetExportDatasetByMsgId)
			m.Get("/check_old", reqToken(), dataset.CheckOld)
		})

		m.Group("/dataset", func() {
			m.Combo("").Get(reqAccessRead(), dataset.GetDataset).Delete(reqSubjectOwner(), dataset.DelDataset)
			m.Post("/readme", reqAccessWrite(), bind(entity.DatasetReadmeReq{}), dataset.PutDatasetReadme)
			m.Get("/readme", reqAccessRead(), dataset.GetDatasetReadme)
			m.Get("/files", reqAccessRead(), dataset.GetDatasetFileList)
			m.Get("/file", reqAccessRead(), dataset.DownloadDatasetFile)
			m.Get("/file/meta", reqAccessRead(), dataset.GetDatasetFileMeta)
			m.Delete("/file", reqAccessWrite(), dataset.DeleteDatasetFile)
			m.Get("/download", reqAccessRead(), dataset.DownloadDataset)
			m.Get("/download/meta", reqAccessRead(), dataset.GetDownloadDatasetMeta)
			m.Get("/preview", reqAccessRead(), dataset.PreviewDatasetFile)
			m.Post("/edit", reqAccessAdmin(), bind(entity.CreateDatasetReq{}), dataset.EditDataset)
			m.Combo("/collect").Put(reqAccessRead(), dataset.CollectDataset).Delete(dataset.UnCollectDataset)
			m.Post("/export_task_result", reqAccessWrite(), bind(entity.ExportTaskResultReq{}), dataset.ExportTaskResult)
			m.Get("/sdk_download_code", reqAccessRead(), ai_task.GenerateDatasetDownloadCode)
		}, datasetAssignment())

		m.Group("/access/collaboration", func() {
			m.Combo("").Get(access.GetCollaborators).Post(access.AddCollaborator).Delete(access.DeleteCollaboration)
			m.Post("/access_mode", access.ChangeCollaborationAccessMode)
			m.Post("/delete", access.DeleteCollaboration)
			m.Group("/dataset", func() {
				m.Group("/team", func() {
					m.Post("", access.AddDatasetTeam)
					m.Post("/delete", access.DeleteDatasetTeam)
				})
			})
			m.Group("/aimodel", func() {
				m.Group("/team", func() {
					m.Post("", access.AddAimodelTeam)
					m.Post("/delete", access.DeleteAimodelTeam)
				})
			})
		}, reqToken(), accessAssignment(), reqAccessAdmin())

		m.Group("/ai_task_template", func() {
			m.Post("/create", reqToken(), bind(entity.CreateAITaskTemplateReq{}), ai_task_template.CreateAITaskTemplate)
			m.Post("/edit", reqToken(), bind(entity.CreateAITaskTemplateReq{}), ai_task_template.EditAITaskTemplate)
			m.Combo("").Get(ai_task_template.GetAITaskTemplate).Delete(reqToken(), ai_task_template.DelAITaskTemplate)
			m.Combo("/collect").Put(ai_task_template.CollectAITaskTemplate).Delete(ai_task_template.UnCollectAITaskTemplate)
			m.Group("/list", func() {
				m.Get("/public", ai_task_template.SearchAllPublicTemplates)
				m.Get("/all", ai_task_template.SearchAllTemplates)
				m.Get("/collected", ai_task_template.SearchCollectedTemplates)
				m.Get("/created", ai_task_template.SearchCreatedTemplates)
				m.Get("/accessible", reqToken(), ai_task_template.SearchAccessibleTemplates)
			})
		})

		// Miscellaneous
		if setting.API.EnableSwagger {
			m.Get("/swagger", misc.Swagger)
		}
		m.Get("/version", misc.Version)
		m.Get("/signing-key.gpg", misc.SigningKey)
		m.Post("/markdown", bind(api.MarkdownOption{}), misc.Markdown)
		m.Post("/markdown/raw", misc.MarkdownRaw)

		m.Group("/monitor", func() {
			m.Get("/summary", HasOperRole(role.ROLE_OPER_KANBANAdmin), monitor.GetTaskMonitorSummary)
			m.Get("/alerts", HasOperRole(role.ROLE_OPER_KANBANAdmin), monitor.GetAlerts)
			m.Get("/aim_right", repo.HasAIMRight)
			m.Post("/aim_url", HasOperRole(role.ROLE_OPER_MonitorAdmin), bind(repo.AimJobName{}), repo.GetAimUrl)
			m.Get("/mlops_right", repo.HasMLopsRight)
			m.Get("/pprof", HasOperRole(role.ROLE_OPER_MonitorAdmin), memory.PprofMemory)
			m.Group("/:username/:reponame", func() {
				m.Group("/ai_task", func() {
					m.Get("", HasOperRole(role.ROLE_OPER_MonitorAdmin), ai_task.GetAITaskInfo)
					m.Get("/log/download", HasOperRole(role.ROLE_OPER_MonitorAdmin), ai_task.DownloadAITaskLog)
					m.Get("/resource_usage", HasOperRole(role.ROLE_OPER_MonitorAdmin), ai_task.GetAITaskResourceUsage)
				})

			})
			m.Get("/ai_tasks", HasOperRole(role.ROLE_OPER_MonitorAdmin), ai_task.GetAllMonitorAITask)
		})

		m.Group("/images", func() {
			m.Get("/recommend", repo.GetRecommendImages)
			m.Get("/public", repo.GetPublicImages)
			m.Get("/custom", reqToken(), repo.GetCustomImages)
			m.Get("/star", reqToken(), repo.GetStarImages)
			m.Get("/npu", reqToken(), repo.GetNpuImages)
			m.Get("/availableFilter", reqToken(),
				repo.GetAvailableFilerInfo)

		})
		m.Group("/storage", func() {
			m.Get("/summary", storage.GetUserStorageSummary)
			m.Get("/dataset_attachment", storage.SearchDataSetAttachments)
			m.Get("/dataset", dataset.SearchDatasetsForPersonalStorage)
			m.Get("/aimodel", aimodel.SearchAimodelsForPersonalStorage)
		}, reqToken())

		m.Group("/tech", func() {
			m.Get("", tech.FindTech)
			m.Post("/basic", HasOperRole(role.ROLE_OPER_TechProgramAdmin), tech.ImportBasicInfo)
			m.Get("/filter", tech.GetFilterInfo)
			m.Get("/search", tech.SearchTechProjectInfo)
			m.Get("/repo_search", tech.SearchRepoInfo)
			m.Get("/admin", HasOperRole(role.ROLE_OPER_TechProgramAdmin), tech.GetAdminRepoInfo)
			m.Get("/my", tech.GetMyRepoInfo)
			m.Post("/admin/:action", HasOperRole(role.ROLE_OPER_TechProgramAdmin), bind(tech.ActionIDs{}), tech.Action)
			m.Post("/openi", bind(api.OpenITechRepo{}), tech.CommitOpenIRepo)
			m.Post("/no_openi", bind(api.NotOpenITechRepo{}), tech.CommitNotOpenIRepo)
			m.Get("/is_admin", tech.IsAdmin)
		}, reqToken())
		m.Group("/finetune", func() {
			m.Get("/checkRepo", finetune.CheckRepo)
			m.Get("/spec", finetune.GetSpec)
		}, reqToken())

		m.Group("/llm", func() {
			m.Post("/chat/completion", bind(llm_chat.ChatConfig{}), llm_chat.ChatCompletionAPI)
			// 云脑二modelarts部署deepseek服务
			m.Post("/chat/deepseek", bind(modelarts.ChatConfigDeepSeek{}), llm_chat.ChatDeepSeekAPINew)
			m.Post("/feedback/deepseek", bind(modelarts.ChatFeedBack{}), llm_chat.FeedBackDeepSeekAPI)
			//北京超算裸机部署服务
			//	m.Get("/stats", reqAdmin(), llm_chat.GetChatStats)
			//	m.Group("/chat", func() {
			//		m.Get("/counts", llm_chat.GetFreeTries)
			//		m.Post("/visit", llm_chat.NewVisit)
			//		m.Post("/agree", llm_chat.SaveAgreement)
			//		m.Post("/legaltext", bind(api.LegalTextParams{}), llm_chat.LegalText)
			//		m.Post("/chat", bind(api.LLMChatMessage{}), llm_chat.LLMChat)
			//		m.Post("/knowledge_base_chat", bind(api.KBChatMessage{}), llm_chat.KBChat)
			//	})
			//	m.Group("/knowledge_base", func() {
			//		m.Get("/list", llm_chat.ListKnowledgeBase)
			//		m.Post("/create", bind(api.CreateKnowledgeBaseParams{}), llm_chat.CreateKnowledgeBase)
			//		m.Post("/delete", llm_chat.DeleteKnowledgeBase)
			//		m.Get("/list_files", llm_chat.ListFiles)
			//		m.Post("/search_docs", bind(api.SearchDocParams{}), llm_chat.SearchDoc)
			//		m.Post("/delete_doc", bind(api.DeleteDocParams{}), llm_chat.DeleteDoc)
			//		m.Post("/update_doc", llm_chat.UpdateDoc)
			//		m.Post("/recreate_vector_store", llm_chat.RecreateVectorStore)
			//		m.Get("/upload_doc_url", llm_chat.UploadDocUrl)
			//		m.Post("/upload_doc", binding.MultipartForm(api.LLMChatUploadForm{}), llm_chat.UploadDoc)
			//		m.Get("/download_doc_url", llm_chat.DownloadDoc)
			//	})
		}, reqToken(), reqWeChatStandard())

		m.Group("/sd", func() {
			m.Group("/finetune", func() {
				m.Get("/stage", sd.GetTaskStage)
				m.Post("/train_lora", bind(sd.UserConfig{}), sd.SendTrainLora)
				m.Post("/text2img", bind(sd.UserInferPrompt{}), sd.SendText2Img)
			}, reqToken(), reqWeChatStandard())
		})

		m.Group("/comfyui", func() {
			m.Group("/experience", func() {
				m.Get("/online_url", sd.GetComfyuiExperienceUrl)
			}, reqToken(), reqWeChatStandard())
		})

		m.Group("/tts", func() {
			m.Get("/get_speaker_voice", llm_chat.GetSpeakerVoice)
			m.Post("/synthesize", bind(llm_chat.TTSConfig{}), llm_chat.SynthesizeVoice)
		}, reqToken(), reqWeChatStandard())

		m.Group("/hf_model", func() {
			m.Get("/get_transfer_user", hf_model.GetTransferUserAPI)
		})
		m.Group("/hf_model", func() {
			m.Post("/new", bind(hf_model.NewTransferRequest{}), hf_model.NewTransferAPI)
			m.Get("/status", hf_model.GetTransferStatusAPI)
			m.Post("/retry", bind(hf_model.RetryTransferRequest{}), hf_model.RetryTransferAPI)
			m.Get("/fetch_update", hf_model.FetchHfUpdateAPI)
			m.Post("/update", bind(hf_model.ModelUpdate{}), hf_model.UpdateHfModelAPI)
			m.Group("/transfer", func() {
				m.Post("/start_waitlist", hf_model.StartTransferFilesAPI)
				m.Post("/heartbeat", hf_model.UpdateFileStatusAPI)
			})
		}, reqToken())

		m.Group("/reward_point", func() {
			m.Get("/is_admin", user.IsRewardPointAdmin)
			m.Group("/list", func() {
				m.Get("/export", HasOperRole(role.ROLE_OPER_RewardPointAdmin), point.ExportAdminRewardList)
			})
		}, reqToken())

		m.Group("/attachments", func() {
			m.Get("/:uuid", repo.GetAttachment)
			m.Get("/get_chunks", repo.GetSuccessChunks)
			m.Get("/new_multipart", repo.NewMultipart)
			m.Get("/get_multipart_url", repo.GetMultipartUploadUrl)
			m.Post("/complete_multipart", repo.CompleteMultipart)
		}, reqToken())
		m.Group("/attachments", func() {
			m.Get("/get_dir", repo.GetAttachmentDir)
			m.Get("/get_image_content", repo.GetAttachmentImageContent)
			m.Get("/get_txt_content", repo.GetAttachmentTxtContent)
		})
		m.Group("/attachments/model", func() {
			m.Get("/get_chunks", repo.GetModelChunks)
			m.Get("/new_multipart", repo.NewModelMultipart)
			m.Get("/get_multipart_url", repo.GetModelMultipartUploadUrl)
			m.Post("/complete_multipart", repo.CompleteModelMultipart)
		})
		m.Group("/upload", func() {
			m.Get("/get_chunks", accessAssignment(), reqAccessWrite(), upload.GetUploadChunks)
			m.Get("/new_multipart", accessAssignment(), reqAccessWrite(), upload.NewUploadMultipart)
			m.Get("/get_multipart_url", upload.GetMultipartUploadUrl)
			m.Post("/complete_multipart", upload.CompleteUploadMultipart)
			m.Group("/direct", func() {
				m.Get("/get_upload_url", accessAssignment(), reqAccessWrite(), upload.GetUploadUrl)
				m.Post("/complete_upload", accessAssignment(), bind(entity.CompleteUploadRequest{}), upload.CompleteUpload)
			})
		}, reqToken())

		m.Group("/pipeline", func() {
			m.Post("/notification", bind(api.PipelineNotification{}), notify.PipelineNotify)

		}, reqToken())

		m.Group("/dync_config", func() {
			m.Get("", dync_config.GetDynicConfig)
		})

		m.Group("/modelscope", func() {
			m.Get("/datasets", modelscope.ListDatasets)
			m.Get("/models", modelscope.ListModels)
		})

		m.Group("/home", func() {
			m.Get("", dync_config.GetHomeConfig)
		})

		m.Group("/pay_computility", func() {
			m.Get("", pay_computility.GetComputilityConfig)
		})

		m.Group("/computility_partner", func() {
			m.Get("", pay_computility.ComputilityPartner)
		})

		m.Get("/compute-nodes", reqToken(), user.GetComputeNodes)

		// Notifications
		m.Group("/notifications", func() {
			m.Combo("").
				Get(notify.ListNotifications).
				Put(notify.ReadNotifications)
			m.Get("/new", notify.NewAvailable)
			m.Combo("/threads/:id").
				Get(notify.GetThread).
				Patch(notify.ReadThread)
		}, reqToken())

		operationReq := context.Toggle(&context.ToggleOptions{SignInRequired: true, OperationRequired: true})
		// hf model migration board
		m.Group("/hf_model/board", func() {
			m.Get("/model_stats", hf_model.GetModelStats)
			m.Get("/model_stats_download", hf_model.DownloadModelStats)
		}, operationReq)

		//Project board
		m.Group("/projectboard", func() {

			m.Get("/restoreFork", repo.RestoreForkNumber)
			m.Get("/downloadAll", repo.ServeAllProjectsPeriodStatisticsFile)
			m.Get("/downloadAllOpenI", repo.ServeAllProjectsOpenIStatisticsFile)
			m.Get("/summary", repo.GetLatestProjectsSummaryData)
			m.Get("/summary/period", repo.GetProjectsSummaryData)
			m.Get("/summary/download", repo.GetProjectsSummaryDataFile)
			m.Group("/project", func() {
				m.Get("", repo.GetAllProjectsPeriodStatistics)
				m.Get("/numVisit", repo.ProjectNumVisit)

				m.Group("/:id", func() {
					m.Get("", repo.GetProjectLatestStatistics)
					m.Get("/period", repo.GetProjectPeriodStatistics)

				})
			})
		}, operationReq)

		m.Get("/query_user_count_time_info", operationReq, repo_ext.QueryUserCountTimeInfo)
		m.Get("/query_metrics_current_month", operationReq, repo_ext.QueryUserMetricsCurrentMonth)
		m.Get("/query_metrics_current_week", operationReq, repo_ext.QueryUserMetricsCurrentWeek)
		m.Get("/query_metrics_current_year", operationReq, repo_ext.QueryUserMetricsCurrentYear)
		m.Get("/query_metrics_last30_day", operationReq, repo_ext.QueryUserMetricsLast30Day)
		m.Get("/query_metrics_last_month", operationReq, repo_ext.QueryUserMetricsLastMonth)
		m.Get("/query_metrics_yesterday", operationReq, repo_ext.QueryUserMetricsYesterday)
		m.Get("/query_metrics_all", operationReq, repo_ext.QueryUserMetricsAll)
		m.Get("/query_user_metrics_page", operationReq, repo_ext.QueryUserMetricDataPage)

		m.Post("/user_statistic_manually", operationReq, private.UserStatisticManually)
		m.Get("/download_user_define_file", operationReq, repo_ext.DownloadUserDefineFile)
		m.Get("/query_user_rank_list", operationReq, repo_ext.QueryRankingList)
		m.Get("/query_user_static_page", operationReq, repo_ext.QueryUserStaticDataPage)
		m.Get("/query_user_current_month", operationReq, repo_ext.QueryUserStaticCurrentMonth)
		m.Get("/query_user_current_week", operationReq, repo_ext.QueryUserStaticCurrentWeek)
		m.Get("/query_user_last_week", operationReq, repo_ext.QueryUserStaticLastWeek)
		m.Get("/query_user_current_year", operationReq, repo_ext.QueryUserStaticCurrentYear)
		m.Get("/query_user_last30_day", operationReq, repo_ext.QueryUserStaticLast30Day)
		m.Get("/query_user_last_month", operationReq, repo_ext.QueryUserStaticLastMonth)
		m.Get("/query_user_yesterday", operationReq, repo_ext.QueryUserStaticYesterday)
		m.Get("/query_user_all", operationReq, repo_ext.QueryUserStaticAll)
		m.Get("/query_user_activity", operationReq, repo_ext.QueryUserActivity)
		m.Get("/query_user_login", operationReq, repo_ext.QueryUserLoginInfo)

		m.Get("/query_invitation_current_month", operationReq, repo_ext.QueryInvitationCurrentMonth)
		m.Get("/query_invitation_current_week", operationReq, repo_ext.QueryInvitationCurrentWeek)
		m.Get("/query_invitation_last_week", operationReq, repo_ext.QueryInvitationLastWeek)
		m.Get("/query_invitation_current_year", operationReq, repo_ext.QueryInvitationCurrentYear)
		m.Get("/query_invitation_last30_day", operationReq, repo_ext.QueryInvitationLast30Day)
		m.Get("/query_invitation_last_month", operationReq, repo_ext.QueryInvitationLastMonth)
		m.Get("/query_invitation_yesterday", operationReq, repo_ext.QueryInvitationYesterday)
		m.Get("/query_invitation_all", operationReq, repo_ext.QueryInvitationAll)
		m.Get("/query_invitation_userdefine", operationReq, repo_ext.QueryUserDefineInvitationPage)
		m.Get("/query_user_annual_report", repo_ext.QueryUserAnnualReport)

		m.Get("/download_invitation_detail", operationReq, repo_ext.DownloadInvitationDetail)

		//cloudbrain board
		m.Get("/cloudbrainboard/cloudbrain/resource_queues", repo.GetResourceQueues)
		m.Get("/cloudbrainboard/ai_center_overview", repo.GetCloubrainOverviewGroupByAiCenter)
		m.Get("/cloudbrainboard/location", cloudbrainService.GetCloudbrainLocationInfo)
		m.Get("/cloudbrainboard/card_data", repo.GetCartStatisticData)

		m.Group("/cloudbrainboard", func() {
			m.Get("/downloadAll", repo.DownloadCloudBrainBoard)
			m.Group("/cloudbrain", func() {
				m.Get("/overview", repo.GetAllCloudbrainsOverview)
				m.Get("/overview_duration", repo.GetOverviewDuration)
				m.Get("/distribution", repo.GetAllCloudbrainsPeriodDistribution)
				m.Get("/trend", repo.GetAllCloudbrainsTrend)
				m.Get("/trend_detail_data", repo.GetAllCloudbrainsTrendDetail)
				m.Get("/status_analysis", repo.GetCloudbrainsStatusAnalysis)
				m.Get("/detail_data", repo.GetCloudbrainsDetailData)
				m.Get("/hours_data", repo.GetCloudbrainsCreateHoursData)
				m.Get("/waitting_top_data", repo.GetWaittingTop)
				m.Get("/running_top_data", repo.GetRunningTop)

				m.Get("/overview_resource", repo.GetCloudbrainResourceOverview)
				m.Get("/resource_usage_statistic", repo.GetDurationRateStatistic)
				m.Get("/resource_usage_rate_detail", repo.GetCloudbrainResourceUsageDetail)
				m.Get("/apitest_for_statistic", repo.CloudbrainDurationStatisticForTest)

				m.Get("/compute_center_task_status", repo.GetComputeCenterTaskStatus)
				m.Get("/compute_center_task_detail", repo.GetComputeCenterTaskDetail)
				m.Get("/queue_duration_usage_statistic", repo.GetQueueDurationUsageStat)
				m.Get("/queue_duration_usage_statistic_detail", repo.GetQueueDurationUsageDetail)

			})
		}, operationReq)

		// Users
		m.Group("/users", func() {
			m.Get("/search", reqToken(), user.Search)

			m.Group("/:username", func() {
				m.Get("", reqToken(), user.GetInfo)
				m.Get("/heatmap", mustEnableUserHeatmap, user.GetUserHeatmapData)

				m.Get("/repos", user.ListUserRepos)

				m.Group("/tokens", func() {
					m.Combo("").Get(user.ListAccessTokens).
						Post(bind(api.CreateAccessTokenOption{}), user.CreateAccessToken)
					m.Combo("/:id").Delete(user.DeleteAccessToken)
				}, reqBasicAuth())
			})
		})

		m.Group("/users", func() {
			m.Group("/:username", func() {
				m.Get("/keys", user.ListPublicKeys)
				m.Get("/gpg_keys", user.ListGPGKeys)

				m.Get("/followers", user.ListFollowers)
				m.Group("/following", func() {
					m.Get("", user.ListFollowing)
					m.Get("/:target", user.CheckFollowing)
				})

				m.Get("/starred", user.GetStarredRepos)

				m.Get("/subscriptions", user.GetWatchedRepos)
			})
		}, reqToken())

		m.Group("/user", func() {
			m.Get("", user.GetAuthenticatedUser)
			m.Get("/subscriber", user.IsSubscriber)
			m.Get("/point_account", user.GetPointAccount)
			m.Combo("/emails").Get(user.ListEmails).
				Post(bind(api.CreateEmailOption{}), user.AddEmail).
				Delete(bind(api.DeleteEmailOption{}), user.DeleteEmail)

			m.Get("/followers", user.ListMyFollowers)
			m.Group("/following", func() {
				m.Get("", user.ListMyFollowing)
				m.Combo("/:username").Get(user.CheckMyFollowing).Put(user.Follow).Delete(user.Unfollow)
			})
			m.Get("/get_latest_cloudbrain_repo", user.GetCloudbrainRepo)

			m.Group("/application", func() {
				m.Get("/has-access", user.HasAccess)
			})

			m.Group("/keys", func() {
				m.Combo("").Get(user.ListMyPublicKeys).
					Post(bind(api.CreateKeyOption{}), user.CreatePublicKey)
				m.Combo("/:id").Get(user.GetPublicKey).
					Delete(user.DeletePublicKey)
			})
			m.Group("/applications", func() {
				m.Combo("/oauth2").
					Get(user.ListOauth2Applications).
					Post(bind(api.CreateOAuth2ApplicationOptions{}), user.CreateOauth2Application)
				m.Combo("/oauth2/:id").
					Delete(user.DeleteOauth2Application).
					Patch(bind(api.CreateOAuth2ApplicationOptions{}), user.UpdateOauth2Application).
					Get(user.GetOauth2Application)
			}, reqToken())

			m.Group("/gpg_keys", func() {
				m.Combo("").Get(user.ListMyGPGKeys).
					Post(bind(api.CreateGPGKeyOption{}), user.CreateGPGKey)
				m.Combo("/:id").Get(user.GetGPGKey).
					Delete(user.DeleteGPGKey)
			})

			m.Combo("/repos").Get(user.ListMyRepos).
				Post(bind(api.CreateRepoOption{}), repo.Create)

			m.Group("/starred", func() {
				m.Get("", user.GetMyStarredRepos)
				m.Group("/:username/:reponame", func() {
					m.Get("", user.IsStarring)
					m.Put("", user.Star)
					m.Delete("", user.Unstar)
				}, repoAssignment())
			})
			m.Get("/times", repo.ListMyTrackedTimes)

			m.Get("/stopwatches", repo.GetStopwatches)

			m.Get("/subscriptions", user.GetMyWatchedRepos)

			m.Get("/teams", org.ListUserTeams)
		}, reqToken())

		// Repositories
		m.Post("/org/:org/repos", reqToken(), bind(api.CreateRepoOption{}), repo.CreateOrgRepoDeprecated)

		m.Combo("/repositories/:id", reqToken()).Get(repo.GetByID)

		m.Group("/datasets/:username/:reponame", func() {
			m.Get("", repo.CurrentRepoDatasetInfoWithoutAttachment)
			m.Get("/current_repo", repo.CurrentRepoDatasetMultiple)
			m.Get("/my_datasets", repo.MyDatasetsMultiple)
			m.Get("/public_datasets", repo.PublicDatasetMultiple)
			m.Get("/my_favorite", repo.MyFavoriteDatasetMultiple)
			m.Post("/create", reqRepoWriter(models.UnitTypeDatasets), bind(auth.CreateDatasetForm{}), repo.CreateDatasetAPI)
			m.Group("/model", func() {
				m.Get("/getmodelfile", repo.GetDataSetSelectItemByJobId)
				m.Get("/getprogress", repo.GetExportDataSetByMsgId)
				m.Post("/export_exist_dataset", repo.ExportModelToExistDataSet)
			})
		}, reqToken(), repoAssignment())

		m.Group("/file_notebook", func() {
			m.Get("", repo.GetFileNoteBookInfo)
			m.Post("/create", reqToken(), reqWeChatStandard(), bind(api.CreateFileNotebookJobOption{}), repo.CreateFileNoteBook)
			m.Post("/status", reqToken(), bind(api.CreateFileNotebookJobOption{}), repo.FileNoteBookStatus)
		})

		m.Group("/model_notebook", func() {
			m.Post("/create", reqToken(), reqWeChatStandard(), bind(api.CreateModelNoteBook{}), repo.CreateModelNoteBook)
		})

		m.Group("/repos", func() {
			m.Get("/search", reqToken(), repo.Search)
			// 工作台改造项目的查询
			m.Get("/platform/search", repo.PlatFormRepoSearch)

			m.Get("/search_for_ai_task", reqToken(), repo.ReposCanCreateCloudbrainJob)

			m.Get("/issues/search", repo.SearchIssues)

			m.Post("/migrate", reqToken(), bind(auth.MigrateRepoForm{}), repo.Migrate)
			m.Post("/migrate/submit", reqToken(), bind(auth.MigrateRepoForm{}), repo.MigrateSubmit)

			m.Group("/specification", func() {
				m.Get("", repo.GetResourceSpec)
			}, reqToken())

			m.Group("/:username/:reponame", func() {
				// 项目模块
				m.Group("/project", func() {
					// 置顶\取消置顶
					m.Group("/pinned", func() {
						m.Combo("").
							Post(repo.PinnedProject).
							Delete(repo.CancelPinnedProject)
					}, reqToken())

					// 删除项目
					m.Delete("", reqToken(), reqOwner(), repo.DeleteProject)
				})
				m.Get("/right", reqToken(), repo.GetRight)
				m.Get("/tagger", reqToken(), repo.ListTagger)
				m.Get("/cloudBrainJobId", repo.GetCloudBrainJobId)
				m.Combo("").Get(reqAnyRepoReader(), repo.Get).
					Delete(reqToken(), reqOwner(), repo.Delete).
					Patch(reqToken(), reqAdmin(), bind(api.EditRepoOption{}), context.RepoRef(), repo.Edit)
				m.Post("/transfer", reqOwner(), bind(api.TransferRepoOption{}), repo.Transfer)
				m.Combo("/notifications").
					Get(reqToken(), notify.ListRepoNotifications).
					Put(reqToken(), notify.ReadRepoNotifications)
				m.Group("/hooks", func() {
					m.Combo("").Get(repo.ListHooks).
						Post(bind(api.CreateHookOption{}), repo.CreateHook)
					m.Group("/:id", func() {
						m.Combo("").Get(repo.GetHook).
							Patch(bind(api.EditHookOption{}), repo.EditHook).
							Delete(repo.DeleteHook)
						m.Post("/tests", context.RepoRef(), repo.TestHook)
					})
					m.Group("/git", func() {
						m.Combo("").Get(repo.ListGitHooks)
						m.Group("/:id", func() {
							m.Combo("").Get(repo.GetGitHook).
								Patch(bind(api.EditGitHookOption{}), repo.EditGitHook).
								Delete(repo.DeleteGitHook)
						})
					}, reqGitHook(), context.ReferencesGitRepo(true))
				}, reqToken(), reqAdmin())
				m.Group("/collaborators", func() {
					m.Get("", reqAnyRepoReader(), repo.ListCollaborators)
					m.Combo("/:collaborator").Get(reqAnyRepoReader(), repo.IsCollaborator).
						Put(reqAdmin(), bind(api.AddCollaboratorOption{}), repo.AddCollaborator).
						Delete(reqAdmin(), repo.DeleteCollaborator)
				}, reqToken())
				m.Get("/raw/*", context.RepoRefByType(context.RepoRefAny), reqRepoReader(models.UnitTypeCode), repo.GetRawFile)
				m.Get("/archive/*", reqRepoReader(models.UnitTypeCode), repo.GetArchive)
				m.Combo("/forks").Get(repo.ListForks).
					Post(reqToken(), reqRepoReader(models.UnitTypeCode), bind(api.CreateForkOption{}), repo.CreateFork)
				m.Group("/branches", func() {
					m.Get("", repo.ListBranches)
					m.Get("/*", context.RepoRefByType(context.RepoRefBranch), repo.GetBranch)
					m.Delete("/*", reqRepoWriter(models.UnitTypeCode), context.RepoRefByType(context.RepoRefBranch), repo.DeleteBranch)
				}, reqRepoReader(models.UnitTypeCode))
				m.Group("/branch_protections", func() {
					m.Get("", repo.ListBranchProtections)
					m.Post("", bind(api.CreateBranchProtectionOption{}), repo.CreateBranchProtection)
					m.Group("/:name", func() {
						m.Get("", repo.GetBranchProtection)
						m.Patch("", bind(api.EditBranchProtectionOption{}), repo.EditBranchProtection)
						m.Delete("", repo.DeleteBranchProtection)
					})
				}, reqToken(), reqAdmin())
				m.Group("/tags", func() {
					m.Get("", repo.ListTags)
				}, reqRepoReader(models.UnitTypeCode), context.ReferencesGitRepo(true))
				m.Group("/keys", func() {
					m.Combo("").Get(repo.ListDeployKeys).
						Post(bind(api.CreateKeyOption{}), repo.CreateDeployKey)
					m.Combo("/:id").Get(repo.GetDeployKey).
						Delete(repo.DeleteDeploykey)
				}, reqToken(), reqAdmin())
				m.Group("/times", func() {
					m.Combo("").Get(repo.ListTrackedTimesByRepository)
					m.Combo("/:timetrackingusername").Get(repo.ListTrackedTimesByUser)
				}, mustEnableIssues, reqToken())
				m.Group("/issues", func() {
					m.Combo("").Get(repo.ListIssues).
						Post(reqToken(), mustNotBeArchived, bind(api.CreateIssueOption{}), repo.CreateIssue)
					m.Group("/comments", func() {
						m.Get("", repo.ListRepoIssueComments)
						m.Group("/:id", func() {
							m.Combo("").
								Get(repo.GetIssueComment).
								Post(mustNotBeArchived, reqToken(), bind(api.EditIssueCommentOption{}), repo.EditIssueComment).
								Delete(reqToken(), repo.DeleteIssueComment)
							m.Combo("/reactions").
								Get(repo.GetIssueCommentReactions).
								Post(bind(api.EditReactionOption{}), reqToken(), repo.PostIssueCommentReaction).
								Delete(bind(api.EditReactionOption{}), reqToken(), repo.DeleteIssueCommentReaction)
						})
					})
					m.Group("/:index", func() {
						m.Combo("").Get(repo.GetIssue).
							Post(reqToken(), bind(api.EditIssueOption{}), repo.EditIssue)
						m.Group("/comments", func() {
							m.Combo("").Get(repo.ListIssueComments).
								Post(reqToken(), mustNotBeArchived, bind(api.CreateIssueCommentOption{}), repo.CreateIssueComment)
							m.Combo("/:id", reqToken()).Patch(bind(api.EditIssueCommentOption{}), repo.EditIssueCommentDeprecated).
								Delete(repo.DeleteIssueCommentDeprecated)
						})
						m.Group("/labels", func() {
							m.Combo("").Get(repo.ListIssueLabels).
								Post(reqToken(), bind(api.IssueLabelsOption{}), repo.AddIssueLabels).
								Put(reqToken(), bind(api.IssueLabelsOption{}), repo.ReplaceIssueLabels).
								Delete(reqToken(), repo.ClearIssueLabels)
							m.Delete("/:id", reqToken(), repo.DeleteIssueLabel)
						})
						m.Group("/times", func() {
							m.Combo("").
								Get(repo.ListTrackedTimes).
								Post(bind(api.AddTimeOption{}), repo.AddTime).
								Delete(repo.ResetIssueTime)
							m.Delete("/:id", repo.DeleteTime)
						}, reqToken())
						m.Combo("/deadline").Post(reqToken(), bind(api.EditDeadlineOption{}), repo.UpdateIssueDeadline)
						m.Group("/stopwatch", func() {
							m.Post("/start", reqToken(), repo.StartIssueStopwatch)
							m.Post("/stop", reqToken(), repo.StopIssueStopwatch)
							m.Delete("/delete", reqToken(), repo.DeleteIssueStopwatch)
						})
						m.Group("/subscriptions", func() {
							m.Get("", repo.GetIssueSubscribers)
							m.Get("/check", reqToken(), repo.CheckIssueSubscription)
							m.Put("/:user", reqToken(), repo.AddIssueSubscription)
							m.Delete("/:user", reqToken(), repo.DelIssueSubscription)
						})
						m.Combo("/reactions").
							Get(repo.GetIssueReactions).
							Post(bind(api.EditReactionOption{}), reqToken(), repo.PostIssueReaction).
							Delete(bind(api.EditReactionOption{}), reqToken(), repo.DeleteIssueReaction)
					})
				}, mustEnableIssuesOrPulls)
				m.Group("/labels", func() {
					m.Combo("").Get(repo.ListLabels).
						Post(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), bind(api.CreateLabelOption{}), repo.CreateLabel)
					m.Combo("/:id").Get(repo.GetLabel).
						Post(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), bind(api.EditLabelOption{}), repo.EditLabel).
						Delete(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), repo.DeleteLabel)
				})
				m.Post("/markdown", bind(api.MarkdownOption{}), misc.Markdown)
				m.Post("/markdown/raw", misc.MarkdownRaw)
				m.Group("/milestones", func() {
					m.Combo("").Get(repo.ListMilestones).
						Post(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), bind(api.CreateMilestoneOption{}), repo.CreateMilestone)
					m.Combo("/:id").Get(repo.GetMilestone).
						Post(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), bind(api.EditMilestoneOption{}), repo.EditMilestone).
						Delete(reqToken(), reqRepoWriter(models.UnitTypeIssues, models.UnitTypePullRequests), repo.DeleteMilestone)
				})
				m.Get("/stargazers", repo.ListStargazers)
				m.Get("/subscribers", repo.ListSubscribers)
				m.Group("/subscription", func() {
					m.Get("", user.IsWatching)
					m.Put("", reqToken(), user.Watch)
					m.Delete("", reqToken(), user.Unwatch)
				})
				m.Group("/releases", func() {
					m.Combo("").Get(repo.ListReleases).
						Post(reqToken(), reqRepoWriter(models.UnitTypeReleases), context.ReferencesGitRepo(false), bind(api.CreateReleaseOption{}), repo.CreateRelease)
					m.Group("/:id", func() {
						m.Combo("").Get(repo.GetRelease).
							Patch(reqToken(), reqRepoWriter(models.UnitTypeReleases), context.ReferencesGitRepo(false), bind(api.EditReleaseOption{}), repo.EditRelease).
							Delete(reqToken(), reqRepoWriter(models.UnitTypeReleases), repo.DeleteRelease)
						m.Group("/assets", func() {
							m.Combo("").Get(repo.ListReleaseAttachments).
								Post(reqToken(), reqRepoWriter(models.UnitTypeReleases), repo.CreateReleaseAttachment)
							m.Combo("/:asset").Get(repo.GetReleaseAttachment).
								Patch(reqToken(), reqRepoWriter(models.UnitTypeReleases), bind(api.EditAttachmentOptions{}), repo.EditReleaseAttachment).
								Delete(reqToken(), reqRepoWriter(models.UnitTypeReleases), repo.DeleteReleaseAttachment)
						})
					})
				}, reqRepoReader(models.UnitTypeReleases))
				m.Post("/mirror-sync", reqToken(), reqRepoWriter(models.UnitTypeCode), repo.MirrorSync)
				m.Get("/editorconfig/:filename", context.RepoRef(), reqRepoReader(models.UnitTypeCode), repo.GetEditorconfig)
				m.Group("/pulls", func() {
					m.Combo("").Get(bind(api.ListPullRequestsOptions{}), repo.ListPullRequests).
						Post(reqToken(), mustNotBeArchived, bind(api.CreatePullRequestOption{}), repo.CreatePullRequest)
					m.Group("/:index", func() {
						m.Combo("").Get(repo.GetPullRequest).
							Patch(reqToken(), reqRepoWriter(models.UnitTypePullRequests), bind(api.EditPullRequestOption{}), repo.EditPullRequest)
						m.Combo("/merge").Get(repo.IsPullRequestMerged).
							Post(reqToken(), mustNotBeArchived, bind(auth.MergePullRequestForm{}), repo.MergePullRequest)
						m.Group("/reviews", func() {
							m.Combo("").
								Get(repo.ListPullReviews).
								Post(reqToken(), bind(api.CreatePullReviewOptions{}), repo.CreatePullReview)
							m.Group("/:id", func() {
								m.Combo("").
									Get(repo.GetPullReview).
									Delete(reqToken(), repo.DeletePullReview).
									Post(reqToken(), bind(api.SubmitPullReviewOptions{}), repo.SubmitPullReview)
								m.Combo("/comments").
									Get(repo.GetPullReviewComments)
							})
						})

					})
				}, mustAllowPulls, reqRepoReader(models.UnitTypeCode), context.ReferencesGitRepo(false))
				m.Group("/statuses", func() {
					m.Combo("/:sha").Get(repo.GetCommitStatuses).
						Post(reqToken(), bind(api.CreateStatusOption{}), repo.NewCommitStatus)
				}, reqRepoReader(models.UnitTypeCode))
				m.Group("/commits", func() {
					m.Get("", repo.GetAllCommits)
					m.Group("/:ref", func() {
						m.Get("/status", repo.GetCombinedCommitStatusByRef)
						m.Get("/statuses", repo.GetCommitStatusesByRef)
					})
				}, reqRepoReader(models.UnitTypeCode))
				m.Group("/git", func() {
					m.Group("/commits", func() {
						m.Get("/:sha", repo.GetSingleCommit)
					})
					m.Get("/refs", repo.GetGitAllRefs)
					m.Get("/refs/*", repo.GetGitRefs)
					m.Get("/trees/:sha", context.RepoRef(), repo.GetTree)
					m.Get("/blobs/:sha", context.RepoRef(), repo.GetBlob)
					m.Get("/tags/:sha", context.RepoRef(), repo.GetTag)
				}, reqRepoReader(models.UnitTypeCode))
				m.Group("/git", func() {
					m.Get("/codes/:currentBranch", repo.GetCodes)
				}, reqRepoWriter(models.UnitTypeCode))
				m.Group("/contents", func() {
					m.Get("", repo.GetContentsList)
					m.Get("/*", repo.GetContents)
					m.Post("/commit", bind(api.CommitMultiFileOptions{}), reqRepoWriter(models.UnitTypeCode), repo.CommitFiles)
					m.Group("/*", func() {
						m.Post("", bind(api.CreateFileOptions{}), repo.CreateFile)
						m.Put("", bind(api.UpdateFileOptions{}), repo.UpdateFile)
						m.Delete("", bind(api.DeleteFileOptions{}), repo.DeleteFile)
					}, reqRepoWriter(models.UnitTypeCode), reqToken())
				}, reqRepoReader(models.UnitTypeCode))
				m.Get("/signing-key.gpg", misc.SigningKey)
				m.Group("/topics", func() {
					m.Combo("").Get(repo.ListTopics).
						Put(reqToken(), reqAdmin(), bind(api.RepoTopicOptions{}), repo.UpdateTopics)
					m.Group("/:topic", func() {
						m.Combo("").Put(reqToken(), repo.AddTopic).
							Delete(reqToken(), repo.DeleteTopic)
					}, reqAdmin())
				}, reqAnyRepoReader())
				m.Group("/finetune", func() {

					m.Get("", reqToken(), finetune.GetFinetuneJobs)
					m.Post("/create", reqRepoWriter(models.UnitTypeCloudBrain), reqWeChat(), context.ReferencesGitRepo(false), bind(api.CreateFineTuneJobOption{}), finetune.CreateFineTune)
					m.Group("/deploy", func() {
						m.Group("/:jobid", func() {
							m.Get("", finetune.GetPanguDeployStatus)
						})
						m.Post("/create", reqWeChat(), context.ReferencesGitRepo(false), bind(api.CreateFineTuneDeployOption{}), finetune.FineTuneDeployCreate)
						m.Post("/inference", bind(api.PanguInferenceOption{}), finetune.ServiceInference)
						m.Post("/update", bind(api.UpdateFineTuneDeployOption{}), finetune.ServiceUpdate)
					})
				}, reqToken())
				m.Group("/cloudbrain", func() {
					m.Get("/:id", repo.GetCloudbrainTask)
					m.Get("/:id/log", repo.CloudbrainGetLog)
					m.Get("/:id/download_log_file", repo.CloudbrainDownloadLogFile)
					m.Group("/notebook", func() {

						m.Group("/:id", func() {
							m.Get("/debug", reqWeChat(), cloudbrain.AdminOrJobCreaterRight, repo.GrampusNoteBookDebug)
							m.Get("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.GrampusStopJob)
							m.Delete("/del", cloudbrain.AdminOrOwnerOrJobCreaterRight, repo.GrampusNotebookDel)
						})
					})
					m.Group("/train-job", func() {

						m.Group("/:jobid", func() {
							m.Get("", repo.GetModelArtsTrainJobVersion)
							m.Get("/detail", reqToken(), reqRepoReader(models.UnitTypeCloudBrain), repo.CloudBrainShow)
							m.Get("/model_list", repo.CloudBrainModelList)
							m.Get("/download_multi_model", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.MultiModelDownload)
							m.Post("/stop_version", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo_ext.CloudBrainStop)
							m.Put("/stop", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.GeneralCloudBrainJobStop)
							m.Group("/model", func() {
								m.Get("/schedule_status", repo.GetModelScheduleStatus)
								m.Post("/reschedule", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.RetryModelSchedule)
							})
						})
					})
					m.Group("/inference-job", func() {

						m.Group("/:jobid", func() {
							m.Get("", repo.GetCloudBrainInferenceJob)
							m.Get("/detail", reqToken(), reqRepoReader(models.UnitTypeCloudBrain), repo.CloudBrainShow)

							m.Post("/del", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo.DelCloudBrainJob)
							m.Get("/result_list", repo.InferencJobResultList)
						})
					})
				}, reqRepoReader(models.UnitTypeCloudBrain))
				//m.Group("/modelmanage", func() {
				//m.Post("/create_new_model", repo.CreateNewModel)
				//m.Post("/create_local_model", repo.SaveLocalModel)
				//m.Delete("/delete_model_file", repo.DeleteModelFile)
				//m.Get("/show_model_api", repo.ShowModelManageApi)
				//m.Delete("/delete_model", repo.DeleteModel)
				//m.Get("/downloadall", repo.DownloadModel)
				//m.Get("/downloadsingle/:ID", repo.DownloadModelSingle)
				//m.Get("/query_model_byId", repo.QueryModelById)
				//m.Get("/query_model_byName", repo.QueryModelByName)
				//m.Get("/query_model_for_predict", repo.QueryModelListForPredict)
				//m.Get("/query_modelfile_for_predict", repo.QueryModelFileForPredict)
				//m.Get("/query_train_job", repo.QueryTrainJobList)
				//m.Get("/query_train_job_version", repo.QueryTrainJobVersionList)
				//m.Get("/query_train_model", repo.QueryTrainModelList)
				//m.Post("/create_model_convert", repo.CreateModelConvert)
				//m.Post("/convert_stop", repo.StopModelConvert)
				//m.Get("/show_model_convert_page", repo.ShowModelConvertPage)
				//m.Get("/query_model_convert_byId", repo.QueryModelConvertById)
				//m.Get("/query_model_convert_byName", repo.QueryModelConvertByName)
				//m.Get("/query_model_convert_resultfile", repo.QueryModeConvertResultFile)
				//m.Get("/download_model_convert_resultfile", repo.DownloadModeConvertResultFile)
				//m.Get("/query_onelevel_modelfile", repo.QueryOneLevelModelFile)
				//m.Get("/:id", repo.GetCloudbrainModelConvertTask)
				//m.Get("/:id/log", repo.GrampusForModelConvertGetLog)
				//m.Get("/:id/modelartlog", repo.TrainJobForModelConvertGetLog)
				//m.Get("/:id/model_list", repo.CloudBrainModelConvertList)
				//
				//}, reqRepoReader(models.UnitTypeModelManage))
				m.Group("/modelarts", func() {
					m.Group("/notebook", func() {
						//m.Get("/:jobid", repo.GetModelArtsNotebook)
						m.Get("/:id", repo.GetModelArtsNotebook2)
					})
					m.Group("/train-job", func() {
						m.Group("/:jobid", func() {
							m.Get("", repo.GetModelArtsTrainJobVersion)
							m.Get("/log", repo.TrainJobGetLog)
							m.Post("/del_version", repo.DelTrainJobVersion)
							m.Post("/stop_version", repo.StopTrainJobVersion)
							m.Get("/model_list", repo.ModelList)
							m.Get("/metric_statistics", repo.TrainJobGetMetricStatistic)
							m.Get("/download_multi_model", cloudbrain.AdminOrJobCreaterRightForTrain, repo.MultiModelDownload)
						})
					})
					m.Group("/inference-job", func() {
						m.Group("/:jobid", func() {
							m.Get("", repo.GetModelArtsInferenceJob)
							m.Get("/log", repo.TrainJobGetLog)
							m.Post("/del_version", repo.DelTrainJobVersion)
							m.Post("/stop_version", repo.StopTrainJobVersion)
							m.Get("/result_list", repo.ResultList)
							m.Get("/downloadall", repo.DownloadMultiResultFile)
						})
					})
				}, reqRepoReader(models.UnitTypeCloudBrain))
				m.Group("/grampus", func() {
					m.Group("/notebook", func() {
						m.Get("/:id", repo_ext.GetGrampusNotebook)
						m.Get("/:id/job_event", repo_ext.GrampusDebugJobEvents)
					})
					m.Group("/train-job", func() {
						m.Group("/:jobid", func() {
							m.Get("", repo.GetModelArtsTrainJobVersion)
							m.Post("/stop_version", cloudbrain.AdminOrOwnerOrJobCreaterRightForTrain, repo_ext.GrampusStopJob)
							m.Get("/log", repo_ext.GrampusGetLog)
							m.Get("/metrics", repo_ext.GrampusMetrics)
							m.Get("/metrics/:nodeId", repo_ext.GrampusMetrics)
							m.Get("/log/:nodeId", repo_ext.GrampusGetLog)
							m.Get("/download_multi_model", cloudbrain.AdminOrJobCreaterRightForTrain, repo.MultiModelDownload)
							m.Get("/download_log", cloudbrain.AdminOrJobCreaterRightForTrain, repo_ext.GrampusDownloadLog)
							m.Get("/download_log/:nodeId", cloudbrain.AdminOrJobCreaterRightForTrain, repo_ext.GrampusDownloadLog)
							m.Get("/job_event", repo_ext.GrampusTrainJobEvents)
						})
					})
				}, reqRepoReader(models.UnitTypeCloudBrain))
			}, repoAssignment())
		})

		// Organizations
		m.Get("/user/orgs", reqToken(), org.ListMyOrgs)
		m.Get("/user/owners", reqToken(), org.GetMyOwners)
		m.Get("/users/:username/orgs", org.ListUserOrgs)
		m.Post("/orgs", reqToken(), bind(api.CreateOrgOption{}), org.Create)
		m.Get("/orgs", org.GetAll)
		m.Group("/orgs/:org", func() {
			m.Combo("").Get(org.Get).
				Patch(reqToken(), reqOrgOwnership(), bind(api.EditOrgOption{}), org.Edit).
				Delete(reqToken(), reqOrgOwnership(), org.Delete)
			m.Combo("/repos").Get(user.ListOrgRepos).
				Post(reqToken(), bind(api.CreateRepoOption{}), repo.CreateOrgRepo)
			m.Group("/members", func() {
				m.Get("", org.ListMembers)
				m.Combo("/:username").Get(org.IsMember).
					Delete(reqToken(), reqOrgOwnership(), org.DeleteMember)
			})
			m.Group("/public_members", func() {
				m.Get("", org.ListPublicMembers)
				m.Combo("/:username").Get(org.IsPublicMember).
					Put(reqToken(), reqOrgMembership(), org.PublicizeMember).
					Delete(reqToken(), reqOrgMembership(), org.ConcealMember)
			})
			m.Group("/teams", func() {
				m.Combo("", reqToken()).Get(org.ListTeams).
					Post(reqOrgOwnership(), bind(api.CreateTeamOption{}), org.CreateTeam)
				m.Get("/search", org.SearchTeam)
			}, reqOrgMembership())
			m.Group("/labels", func() {
				m.Get("", org.ListLabels)
				m.Post("", reqToken(), reqOrgOwnership(), bind(api.CreateLabelOption{}), org.CreateLabel)
				m.Combo("/:id").Get(org.GetLabel).
					Patch(reqToken(), reqOrgOwnership(), bind(api.EditLabelOption{}), org.EditLabel).
					Delete(reqToken(), reqOrgOwnership(), org.DeleteLabel)
			})
			m.Group("/hooks", func() {
				m.Combo("").Get(org.ListHooks).
					Post(bind(api.CreateHookOption{}), org.CreateHook)
				m.Combo("/:id").Get(org.GetHook).
					Patch(bind(api.EditHookOption{}), org.EditHook).
					Delete(org.DeleteHook)
			}, reqToken(), reqOrgOwnership())
			m.Group("/org_tag/dataset", func() {
				m.Get("/available", org.GetAvailableOrgOfficialDatasetRegistry)
				m.Post("/submit", bind(auth.SubmitDatasetOfTagForm{}), org.SubmitTagsToDatasetRegistry)
				m.Get("/list", org.GetOrgOfficialDatasetCardList)
			})
			m.Group("/org_tag/aimodel", func() {
				m.Get("/available", org.GetAvailableOrgOfficialAimodel)
				m.Post("/submit", bind(auth.SubmitModelOfTagForm{}), org.SubmitTagsToAimodel)
				m.Get("/list", org.GetOrgOfficialAimodelCardList)
			})
			m.Get("/dataset/list", dataset.SearchOrgDatasets)
			m.Get("/aimodel/list", aimodel.SearchOrgAimodels)
			m.Get("/dataset/tags", dataset.GetOwnedPublicDatasetTags)
			m.Get("/aimodel/tags", aimodel.GetOrgAimodelLabels)
			m.Group("/storage", func() {
				m.Get("/summary", storage.GetOrgStorageSummary)
				m.Get("/dataset", dataset.SearchStorage4Datasets)
				m.Get("/aimodel", aimodel.SearchStorage4Aimodels)
			})

		}, orgAssignment(true))
		m.Group("/teams/:teamid", func() {
			m.Combo("").Get(org.GetTeam).
				Patch(reqOrgOwnership(), bind(api.EditTeamOption{}), org.EditTeam).
				Delete(reqOrgOwnership(), org.DeleteTeam)
			m.Group("/members", func() {
				m.Get("", org.GetTeamMembers)
				m.Combo("/:username").
					Get(org.GetTeamMember).
					Put(reqOrgOwnership(), org.AddTeamMember).
					Delete(reqOrgOwnership(), org.RemoveTeamMember)
			})
			m.Group("/repos", func() {
				m.Get("", org.GetTeamRepos)
				m.Combo("/:org/:reponame").
					Put(org.AddTeamRepository).
					Delete(org.RemoveTeamRepository)
			})
		}, orgAssignment(false, true), reqToken(), reqTeamMembership())

		m.Any("/*", func(ctx *context.APIContext) {
			ctx.NotFound()
		})
		m.Post("/dataset/historic/one", reqToken(), dataset.HandleOneOldDataset)
		m.Group("/admin", func() {
			m.Get("/orgs", admin.GetAllOrgs)
			m.Group("/users", func() {
				m.Get("", admin.GetAllUsers)
				m.Post("", bind(api.CreateUserOption{}), admin.CreateUser)
				m.Post("/change_default_portrait", bind(api.UserNames{}), admin.ChangeDefaultUserPortrait)
				m.Group("/:username", func() {
					m.Combo("").Patch(bind(api.EditUserOption{}), admin.EditUser).
						Delete(admin.DeleteUser)
					m.Group("/keys", func() {
						m.Post("", bind(api.CreateKeyOption{}), admin.CreatePublicKey)
						m.Delete("/:id", admin.DeleteUserPublicKey)
					})
					m.Get("/orgs", org.ListUserOrgs)
					m.Post("/orgs", bind(api.CreateOrgOption{}), admin.CreateOrg)
					m.Post("/repos", bind(api.CreateRepoOption{}), admin.CreateRepo)
				})
			})
			m.Group("/role", func() {
				m.Combo("").Post(bind(models.OperateRoleReq{}), admin.AddUserRole).
					Delete(bind(models.OperateRoleReq{}), admin.DeleteUserRole)
				m.Post("/add", admin.AddAiforgeRole)
				m.Delete("/del", admin.DelAiforgeRole)
				m.Post("/update", admin.UpdateAiforgeRole)
				m.Get("/query", admin.QueryAiforgeRole)
				m.Get("/list", admin.ListAiforgeRole)
				m.Get("/org/list", admin.ListAiforgeOrgRole)

				m.Post("/addOperation", admin.AddOperation)
				m.Delete("/delOperation", admin.DelOperation)
				m.Get("/listOperation", admin.ListOperation)

				m.Get("/list_quene", admin.ListResourceQuene)
				m.Post("/settouser", admin.SetAiforgeRoleToUser)
				m.Post("/set_to_org", admin.SetAiforgeRoleToOrg)
				m.Post("/batchAddRoleToUser", admin.BatchAddRoleToUser)
				m.Post("/batchDelRoleToUser", admin.BatchDelRoleToUser)
				m.Get("/list_right_user", admin.ListRightUser)
				m.Get("/list_right_org", admin.ListRightOrg)
			})
			m.Group("/ai_task", func() {
				m.Get("/list", ai_task.GetAITaskList4Admin)
				m.Get("/download/list", repo.DownloadAitask4Admin)
			})
			m.Group("/aimodel", func() {
				m.Get("/list", aimodel.SearchAllAimodelsForAdmin)
				m.Combo("/recommend").Put(aimodelAssignment(), aimodel.RecommendAimodel).Delete(aimodelAssignment(), aimodel.UnRecommendAimodel)
			})
			m.Group("/dataset", func() {
				m.Get("/list", dataset.SearchAllDatasetsForAdmin)
				m.Combo("/recommend").Put(datasetAssignment(), dataset.RecommendDataset).Delete(datasetAssignment(), dataset.UnRecommendDataset)
				m.Group("/historic", func() {
					m.Post("/start", dataset.StartDatasetMigrate)
					m.Post("/stop", dataset.StopDatasetMigrate)
					m.Post("/status", dataset.GetDatasetMigrateStatus)
					m.Post("/one", dataset.HandleOneOldDataset)
					m.Post("/all", dataset.HandleAllOldDatasets)
				})
			})
			m.Post("/reward/action/historic/add", point.AddHistoricSourceContent)
			m.Post("/reward/action/historic/all", point.HandleHistoricSourceContent)
			m.Group("/oauth2", func() {
				m.Group("/application", func() {
					m.Get("/list", oauth2.GetApplicaitonList)
					m.Post("/edit", bind(models.EditOauth2ApplicationReq{}), oauth2.EditApplicaiton)
				})
			})
			m.Group("/ai_task_template", func() {
				m.Combo("/recommend").Put(ai_task_template.RecommendAITaskTemplate).Delete(ai_task_template.UnRecommendAITaskTemplate)
				m.Get("/list", ai_task_template.SearchAllTemplates4Admin)
				m.Post("/historical/all", ai_task_template.HandleHistoricalAITaskTemplate)
			})
			m.Post("/resource/historic/to_urchin/update", dataset.UpdateByBatch)
			m.Post("/resource/historic/to_urchin/create_batch", dataset.CreateBatch)
			m.Get("/resource/historic/to_urchin/query_batch_list", dataset.QueryBatchList)
			m.Get("/resource/historic/to_urchin/query_batch_detail", dataset.QueryBatchDetail)
			m.Post("/resource/historic/to_urchin/rollback", dataset.RollbackBatch)
			m.Post("/resource/historic/to_urchin/update/ids", dataset.UpdateToUrchinByResourceIds)
			m.Get("/resource/historic/to_urchin/total", dataset.CountResourceUpdatedData)

		}, reqToken(), reqSiteAdmin())

		m.Group("/topics", func() {
			m.Get("/search", repo.TopicSearch)
		})
		m.Group("/image/topics", func() {
			m.Get("/search", repo.ImageTopicSearch)
		})
		m.Group("/from_wechat", func() {
			m.Get("/event", authentication.ValidEventSource)
			m.Post("/event", authentication.AcceptWechatEvent)
			m.Get("/prd/event", authentication.ValidEventSource)
			m.Post("/prd/event", authentication.AcceptWechatEvent)
		})
		m.Group("/authentication/wechat", func() {
			m.Get("/qrCode4Bind", authentication.GetQRCode4Bind)
			m.Get("/bindStatus", authentication.GetBindStatus)
			m.Post("/unbind", authentication.UnbindWechat)
		}, reqToken())
		m.Get("/wechat/material", authentication.GetMaterial)
		m.Get("/cloudbrain/get_newest_job", repo.GetNewestJobs)
		m.Get("/cloudbrain/get_center_info", repo.GetAICenterInfo)
		m.Get("/all_model_data", repo.QueryAllModelFile)
		m.Group("/resources", func() {
			m.Get("/acc_card/list", resources.GetAccCardList)
			m.Group("/ai_center", func() {
				m.Get("/card_info/active", resources.GetActiveCardInfoList)
				m.Get("/active", resources.GetActiveAICenterList)
				m.Get("/available", resources.GetAvailableAICenterList)
			})
		})
	}, securityHeaders(), context.APIContexter(), sudo())
}

func securityHeaders() macaron.Handler {
	return func(ctx *macaron.Context) {
		ctx.Resp.Before(func(w macaron.ResponseWriter) {
			// CORB: https://www.chromium.org/Home/chromium-security/corb-for-developers
			// http://stackoverflow.com/a/3146618/244009
			w.Header().Set("x-content-type-options", "nosniff")
		})
	}
}
