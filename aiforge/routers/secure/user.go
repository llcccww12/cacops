// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"net/http"
	"net/mail"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/password"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/services/mailer"
)

func parseLoginSource(ctx *context.Context, u *models.User, sourceID int64, loginName string) {
	if sourceID == 0 {
		return
	}

	source, err := models.GetLoginSourceByID(sourceID)
	if err != nil {
		if models.IsErrLoginSourceNotExist(err) {
			ctx.Error(http.StatusUnprocessableEntity, "", err.Error())
		} else {
			ctx.Error(http.StatusInternalServerError, "GetLoginSourceByID", err.Error())
		}
		return
	}

	u.LoginType = source.Type
	u.LoginSource = source.ID
	u.LoginName = loginName
}

// CreateUser create a user
func CreateUser(ctx *context.Context, form api.CreateUserOption) {
	// swagger:operation POST /admin/users admin adminCreateUser
	// ---
	// summary: Create a user
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/CreateUserOption"
	// responses:
	//   "201":
	//     "$ref": "#/responses/User"
	//   "400":
	//     "$ref": "#/responses/error"
	//   "403":
	//     "$ref": "#/responses/forbidden"
	//   "422":
	//     "$ref": "#/responses/validationError"

	_, err1 := mail.ParseAddress(form.Email)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error_msg": "Email format is wrong.",
		})
		return
	}

	u := &models.User{
		Name:               form.Username,
		FullName:           form.FullName,
		Email:              form.Email,
		Passwd:             form.Password,
		MustChangePassword: false,
		IsActive:           !setting.Service.RegisterEmailConfirm,
		LoginType:          models.LoginPlain,
	}
	if form.MustChangePassword != nil {
		u.MustChangePassword = *form.MustChangePassword
	}

	if strings.Contains(form.Email, " ") {
		log.Error("CreateUser failed: email(%s) contains blank space", form.Email, ctx.Data["MsgID"])
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error_msg": "Email contains blank space",
		})
		return
	}

	parseLoginSource(ctx, u, form.SourceID, form.LoginName)
	if ctx.Written() {
		return
	}
	if !password.IsComplexEnough(form.Password) || len(form.Password) < setting.MinPasswordLength {
		log.Error("CreateUser failed: PasswordComplexity", ctx.Data["MsgID"])
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error_msg": "PasswordComplexity",
		})
		return
	}
	if err := models.CreateUser(u); err != nil {
		if models.IsErrUserAlreadyExist(err) ||
			models.IsErrEmailAlreadyUsed(err) ||
			models.IsErrNameReserved(err) ||
			models.IsErrNameCharsNotAllowed(err) ||
			models.IsErrNamePatternNotAllowed(err) {
			log.Error("CreateUser failed:%v", err.Error(), ctx.Data["MsgID"])
			ctx.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error_msg": err.Error(),
			})
		} else {
			log.Error("CreateUser failed:%v", err.Error(), ctx.Data["MsgID"])
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error_msg": err.Error(),
			})
		}
		return
	}

	err := models.AddEmailAddress(&models.EmailAddress{
		UID: u.ID,
		Email: form.Email,
		IsActivated: !setting.Service.RegisterEmailConfirm,
	})

	if err != nil {
		log.Error("AddEmailAddress failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error_msg": err.Error(),
		})
		return
	}

	// Send confirmation email
	if setting.Service.RegisterEmailConfirm{
		mailer.SendActivateAccountMail(ctx.Locale, u)
		if err := ctx.Cache.Put("MailResendLimit_"+u.LowerName, u.LowerName, 180); err != nil {
			log.Error("Set cache(MailResendLimit) fail: %v", err)
		}
	}

	log.Trace("Account created (%s): %s", ctx.User.Name, u.Name, ctx.Data["MsgID"])

	// Send email notification.
	if form.SendNotify {
		mailer.SendRegisterNotifyMail(ctx.Locale, u)
	}
	ctx.JSON(http.StatusCreated, convert.ToUser(u, ctx.IsSigned, ctx.User.IsAdmin))
}
