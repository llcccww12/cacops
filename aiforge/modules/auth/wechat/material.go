package wechat

import "code.gitea.io/gitea/modules/log"

func GetWechatMaterial(mType string, offset, count int) interface{} {
	result, retryFlag := getMaterial(mType, offset, count)
	if retryFlag {
		log.Info("retryGetWechatMaterial calling")
		refreshAccessToken()
		result, _ = getMaterial(mType, offset, count)
	}
	return result
}
