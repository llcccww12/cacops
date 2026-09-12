package wechat

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

func BindWechat(userId int64, wechatOpenId string) error {
	if !IsWechatAccountUsed(userId, wechatOpenId) {
		log.Error("bind wechat failed, because user use wrong wechat account to bind,userId=%d  wechatOpenId=%s", userId, wechatOpenId)
		return models.NewWechatBindError(setting.BindReplyWechatAccountUsed)
	}

	return models.BindWechatOpenId(userId, wechatOpenId)
}

func UnbindWechat(userId int64, oldWechatOpenId string) error {
	return models.UnbindWechatOpenId(userId, oldWechatOpenId)
}

//IsUserAvailableForWechatBind  if user has bound wechat and the bound openId is not the given wechatOpenId,return false
//otherwise,return true
func IsUserAvailableForWechatBind(userId int64, wechatOpenId string) bool {
	currentOpenId := models.GetUserWechatOpenId(userId)
	return currentOpenId == "" || currentOpenId == wechatOpenId
}

//IsWechatAccountUsed  if wechat account used by another account,return false
//if wechat account not used or used by the given user,return true
func IsWechatAccountUsed(userId int64, wechatOpenId string) bool {
	user := models.GetUserByWechatOpenId(wechatOpenId)
	if user != nil && user.WechatOpenId != "" && user.ID != userId {
		return false
	}
	return true
}
