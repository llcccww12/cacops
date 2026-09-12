package entity

import "code.gitea.io/gitea/models"

type SignUp4Wechat struct {
	SignUpId string
	QRCode   *models.QRCodeResponse
}
