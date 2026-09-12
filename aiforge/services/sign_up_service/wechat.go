package sign_up_service

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/wechat_service"
	"encoding/json"
)

func GetSignUpWechatQRCode(data map[string]interface{}) (*entity.SignUp4Wechat, error) {
	signUpId := util.UUID()
	res, err := wechat_service.CreateQRCode4Bind(models.CreateQRCodeReq{
		Type: models.QR_CODE_TYPE_SING_UP,
		UUID: signUpId,
		Data: data,
	})
	if err != nil {
		log.Error("GetSignUpWechatQRCode err when CreateQRCode4Bind,err=%v", err)
		return nil, err
	}
	return &entity.SignUp4Wechat{SignUpId: signUpId, QRCode: res}, nil
}

func GetWechatOpenIdBySignUpId(signUpId string) string {
	key := redis_key.SignUpWechatKey(signUpId)
	val, _ := redis_client.Get(key)
	if val == "" {
		return ""
	}
	cache := new(models.SignUpWechatCache)
	json.Unmarshal([]byte(val), cache)
	if cache == nil {
		return ""
	}
	return cache.WechatOpenId
}

func IsWechatOpenIdAvailable(wechatOpenId string) bool {
	user := models.GetUserByWechatOpenId(wechatOpenId)
	return user == nil
}
