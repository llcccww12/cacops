package wechat

import "code.gitea.io/gitea/modules/log"

func GetWechatQRCode4Bind(sceneStr string) *QRCodeResponse {
	result, retryFlag := callQRCodeCreate(sceneStr)
	if retryFlag {
		log.Info("retry wechat qr-code calling,sceneStr=%s", sceneStr)
		refreshAccessToken()
		result, _ = callQRCodeCreate(sceneStr)
	}
	return result
}
