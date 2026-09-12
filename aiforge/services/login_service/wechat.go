package login_service

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/wechat_service"
)

func GetLoginWechatQRCode() (*entity.Login4Wechat, error) {
	loginId := util.UUID()
	res, err := wechat_service.CreateQRCode4Bind(models.CreateQRCodeReq{
		Type: models.QR_CODE_TYPE_LOGIN,
		UUID: loginId,
	})
	if err != nil {
		log.Error("GetSignUpWechatQRCode err when CreateQRCode4Bind,err=%v", err)
		return nil, err
	}
	return &entity.Login4Wechat{LoginId: loginId, QRCode: res}, nil
}
