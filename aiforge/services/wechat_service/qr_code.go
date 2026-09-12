package wechat_service

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth/wechat"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"errors"
	gouuid "github.com/satori/go.uuid"
	"time"
)

func CreateQRCode4Bind(req models.CreateQRCodeReq) (*models.QRCodeResponse, error) {
	log.Info("start to create qr-code for binding.req=%+v", req)
	sceneStr := gouuid.NewV4().String()
	r := wechat.GetWechatQRCode4Bind(sceneStr)
	if r == nil {
		return nil, errors.New("createQRCode4Bind failed")
	}

	jsonStr, _ := json.Marshal(&models.QRCode4BindCache{
		UserId: req.UserId,
		Status: models.BIND_STATUS_UNBIND,
		Type:   req.Type,
		UUID:   req.UUID,
		Data:   req.Data,
	})
	isOk, err := redis_client.Setex(redis_key.WechatBindingUserIdKey(sceneStr), string(jsonStr), time.Duration(setting.WechatQRCodeExpireSeconds)*time.Second)
	if err != nil {
		log.Error("createQRCode4Bind failed.e=%+v", err)
		return nil, err
	}
	if !isOk {
		log.Error("createQRCode4Bind failed.redis reply is not ok")
		return nil, errors.New("reply is not ok when set WechatBindingUserIdKey")
	}

	result := &models.QRCodeResponse{
		Url:           r.Url,
		Ticket:        r.Ticket,
		SceneStr:      sceneStr,
		ExpireSeconds: setting.WechatQRCodeExpireSeconds,
	}
	return result, nil
}
