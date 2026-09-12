package authentication

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth/wechat"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/wechat_service"
	"encoding/json"
	"strconv"
)

const tplBindPage base.TplName = "repo/wx_autorize"

// GetQRCode4Bind get QR code for wechat binding
func GetQRCode4Bind(ctx *context.Context) {
	userId := ctx.User.ID

	r, err := wechat_service.CreateQRCode4Bind(models.CreateQRCodeReq{
		UserId: userId,
		Type:   models.QR_CODE_TYPE_NORMAL,
	})
	if err != nil {
		log.Error("GetQRCode4Bind failed,error=%v", err)
		ctx.JSON(200, map[string]interface{}{
			"code": "9999",
			"msg":  "Get QR code failed",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "00",
		"msg":  "success",
		"data": r,
	})
}

// GetBindStatus   the web page will poll the service to get bind status
func GetBindStatus(ctx *context.Context) {
	sceneStr := ctx.Query("sceneStr")
	val, _ := redis_client.Get(redis_key.WechatBindingUserIdKey(sceneStr))
	if val == "" {
		ctx.JSON(200, map[string]interface{}{
			"code": "00",
			"msg":  "QR code expired",
			"data": map[string]interface{}{
				"status": models.BIND_STATUS_EXPIRED,
			},
		})
		return
	}
	qrCache := new(models.QRCode4BindCache)
	json.Unmarshal([]byte(val), qrCache)
	ctx.JSON(200, map[string]interface{}{
		"code": "00",
		"msg":  "success",
		"data": map[string]interface{}{
			"status": qrCache.Status,
			"extend": qrCache.Data,
		},
	})
}

// UnbindWechat
func UnbindWechat(ctx *context.Context) {
	if ctx.User.WechatOpenId != "" {
		wechat.UnbindWechat(ctx.User.ID, ctx.User.WechatOpenId)
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "00",
		"msg":  "success",
	})
}

// GetBindPage
func GetBindPage(ctx *context.Context) {
	userId := ctx.User.ID
	r, _ := wechat_service.CreateQRCode4Bind(models.CreateQRCodeReq{
		UserId: userId,
		Type:   models.QR_CODE_TYPE_NORMAL,
	})
	if r != nil {
		ctx.Data["qrcode"] = r
	}
	redirectUrl := ctx.Query("redirect_to")
	if redirectUrl != "" {
		ctx.SetCookie("redirect_to", setting.AppSubURL+redirectUrl, 0, setting.AppSubURL)
	}
	ctx.HTML(200, tplBindPage)
}

//
//func createQRCode4Bind(userId int64) (*models.QRCodeResponse, error) {
//	log.Info("start to create qr-code for binding.userId=%d", userId)
//	sceneStr := gouuid.NewV4().String()
//	r := wechat.GetWechatQRCode4Bind(sceneStr)
//	if r == nil {
//		return nil, errors.New("createQRCode4Bind failed")
//	}
//
//	jsonStr, _ := json.Marshal(&wechat.QRCode4BindCache{
//		UserId: userId,
//		Status: wechat.BIND_STATUS_UNBIND,
//	})
//	isOk, err := redis_client.Setex(redis_key.WechatBindingUserIdKey(sceneStr), string(jsonStr), time.Duration(setting.WechatQRCodeExpireSeconds)*time.Second)
//	if err != nil {
//		log.Error("createQRCode4Bind failed.e=%+v", err)
//		return nil, err
//	}
//	if !isOk {
//		log.Error("createQRCode4Bind failed.redis reply is not ok")
//		return nil, errors.New("reply is not ok when set WechatBindingUserIdKey")
//	}
//
//	result := &models.QRCodeResponse{
//		Url:           r.Url,
//		Ticket:        r.Ticket,
//		SceneStr:      sceneStr,
//		ExpireSeconds: setting.WechatQRCodeExpireSeconds,
//	}
//	return result, nil
//}

// GetMaterial
func GetMaterial(ctx *context.Context) {
	mType := ctx.Query("type")
	offsetStr := ctx.Query("offset")
	countStr := ctx.Query("count")
	var offset, count int
	if offsetStr == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(offsetStr)
	}
	if countStr == "" {
		count = 20
	} else {
		count, _ = strconv.Atoi(countStr)
	}
	r := wechat.GetWechatMaterial(mType, offset, count)
	ctx.JSON(200, response.SuccessWithData(r))
}
