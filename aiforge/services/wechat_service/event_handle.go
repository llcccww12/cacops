package wechat_service

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth/wechat"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

//<xml>
//  <ToUserName><![CDATA[toUser]]></ToUserName>
//  <FromUserName><![CDATA[FromUser]]></FromUserName>
//  <CreateTime>123456789</CreateTime>
//  <MsgType><![CDATA[event]]></MsgType>
//  <Event><![CDATA[SCAN]]></Event>
//  <EventKey><![CDATA[SCENE_VALUE]]></EventKey>
//  <Ticket><![CDATA[TICKET]]></Ticket>
//</xml>
type WechatMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
	Ticket       string
	Content      string
	MsgID        string
	MsgDataId    string
	Idx          string
	Status       string
}

type MsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

type TextMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
type ImageMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Image        ImageContent
}
type VoiceMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Voice        VoiceContent
}
type VideoMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Video        VideoContent
}
type MusicMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Music        MusicContent
}
type NewsMsgReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	ArticleCount int
	Articles     ArticleItem
}

type ArticleItem struct {
	Item []ArticlesContent
}

type ImageContent struct {
	MediaId string
}
type VoiceContent struct {
	MediaId string
}
type VideoContent struct {
	MediaId     string
	Title       string
	Description string
}
type MusicContent struct {
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
}
type ArticlesContent struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

const (
	WECHAT_EVENT_SUBSCRIBE = "subscribe"
	WECHAT_EVENT_SCAN      = "SCAN"
)

const (
	WECHAT_MSG_TYPE_TEXT  = "text"
	WECHAT_MSG_TYPE_EVENT = "event"
)

func HandleScanEvent(we WechatMsg) string {
	replyMsg := setting.BindReplySuccess
	needNotify := false

	eventKey := we.EventKey
	if eventKey == "" {
		return ""
	}
	sceneStr := strings.TrimPrefix(eventKey, "qrscene_")
	key := redis_key.WechatBindingUserIdKey(sceneStr)
	val, _ := redis_client.Get(key)
	if val == "" {
		return ""
	}
	qrCache := new(models.QRCode4BindCache)
	json.Unmarshal([]byte(val), qrCache)
	if qrCache.Status == models.BIND_STATUS_UNBIND {
		switch qrCache.Type {
		case models.QR_CODE_TYPE_SING_UP:
			u := models.GetUserByWechatOpenId(we.FromUserName)
			if u != nil {
				replyMsg = setting.BindReplySignUpErr4WechatUsed
				qrCache.Data["wechatUsed"] = true
			} else {
				qrCache.Data["wechatUsed"] = false
				sharedUser := ""
				if _, ok := qrCache.Data["sharedUser"]; ok {
					sharedUser = qrCache.Data["sharedUser"].(string)
				}
				c := models.SignUpWechatCache{WechatOpenId: we.FromUserName}
				json, _ := json.Marshal(c)
				redis_client.Setex(redis_key.SignUpWechatKey(qrCache.UUID), string(json), 10*time.Minute)
				replyMsg = fmt.Sprintf(setting.BindReplySignUp, strings.TrimSuffix(setting.AppURL, "/")+"/user/sign_up"+"?signUpId="+qrCache.UUID+"&sharedUser="+sharedUser)
				needNotify = true
			}

		case models.QR_CODE_TYPE_LOGIN:
			user := models.GetUserByWechatOpenId(we.FromUserName)
			c := models.LoginWechatCache{UserId: 0}
			if user == nil {
				json, _ := json.Marshal(c)
				redis_client.Setex(redis_key.LoginWechatKey(qrCache.UUID), string(json), 10*time.Minute)
				replyMsg = setting.BindReplyNoUser
			} else {
				c.UserId = user.ID
				json, _ := json.Marshal(c)
				redis_client.Setex(redis_key.LoginWechatKey(qrCache.UUID), string(json), 10*time.Minute)
				replyMsg = setting.BindReplyLoginSuccess
			}
		default:
			needNotify = true
			err := wechat.BindWechat(qrCache.UserId, we.FromUserName)
			if err != nil {
				redis_client.Del(key)
				if err, ok := err.(models.WechatBindError); ok {
					return err.Reply
				}
				return setting.BindReplyFailedDefault
			}
		}
		qrCache.Status = models.BIND_STATUS_BOUND
		jsonStr, _ := json.Marshal(qrCache)
		redis_client.Setex(redis_key.WechatBindingUserIdKey(sceneStr), string(jsonStr), 60*time.Second)
	}
	if needNotify && qrCache.UserId > 0 {
		u, err := models.GetUserByID(qrCache.UserId)
		if err == nil {
			notification.NotifyWechatBind(u, we.FromUserName)
		}
	}

	return replyMsg
}

func HandleSubscribeEvent(we WechatMsg) *WechatReplyContent {
	r, err := LoadReplyFromCacheAndDisk(SubscribeReply)
	if err != nil || len(r) == 0 {
		return nil
	}
	return r[0]
}
