package entity

import "code.gitea.io/gitea/models"

type Login4Wechat struct {
	LoginId string
	QRCode  *models.QRCodeResponse
}

type LoginBindStatus struct {
	Status int
}
