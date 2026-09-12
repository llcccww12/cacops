package user

import (
	"errors"
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
)

const (
	tplInvitation base.TplName = "user/settings/invite"
)

func GetInvitaionCode(ctx *context.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"disabled":               true,
		"invitation_code":        "",
		"invitation_users":       []interface{}{},
		"invitation_users_count": 0,
	})
}

func InviationTpl(ctx *context.Context) {
	ctx.Redirect(setting.AppSubURL + "/dashboard", http.StatusFound)
}

func RegisteUserByInvitaionCode(invitationcode string, newUser *models.User) error {
	user := parseInvitaionCode(invitationcode)
	if user == nil {
		return errors.New("The invitated user not existed.")
	}

	if newUser.PhoneNumber != "" {
		re := models.QueryInvitaionByPhone(newUser.PhoneNumber)
		if re != nil {
			if len(re) > 0 {
				log.Info("The phone has been invitated. so ingore it.")
				return errors.New("The phone has been invitated.")
			}
		}
	} else {
		log.Info("the phone number is null. user name=" + user.Name)
	}

	invitation := &models.Invitation{
		SrcUserID: user.ID,
		UserID:    newUser.ID,
		Phone:     newUser.PhoneNumber,
		Email:     newUser.Email,
	}

	err := models.InsertInvitaion(invitation)
	if err != nil {
		log.Info("insert error," + err.Error())
	} else {
		notification.NotifyInviteFriendRegister(user, newUser)
	}
	return err
}

func getInvitaionCode(ctx *context.Context) string {
	return ctx.User.Name
}

func parseInvitaionCode(invitationcode string) *models.User {
	user, err := models.GetUserByName(invitationcode)
	if err == nil {
		return user
	}
	return nil
}
