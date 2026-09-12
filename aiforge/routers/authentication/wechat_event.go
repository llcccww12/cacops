package authentication

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	wechat "code.gitea.io/gitea/services/wechat_service"
	"encoding/xml"
	"io/ioutil"
	"time"
)

// AcceptWechatEvent
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_event_pushes.html
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html
func AcceptWechatEvent(ctx *context.Context) {
	b, _ := ioutil.ReadAll(ctx.Req.Request.Body)
	we := wechat.WechatMsg{}
	xml.Unmarshal(b, &we)
	switch we.MsgType {
	case wechat.WECHAT_MSG_TYPE_EVENT:
		HandleEventMsg(ctx, we)
	case wechat.WECHAT_MSG_TYPE_TEXT:
		HandleTextMsg(ctx, we)
	}
	log.Info("accept wechat event= %+v", we)

}

// ValidEventSource
func ValidEventSource(ctx *context.Context) {
	echostr := ctx.Query("echostr")
	ctx.Write([]byte(echostr))
	return
}

func HandleEventMsg(ctx *context.Context, msg wechat.WechatMsg) {
	switch msg.Event {
	case wechat.WECHAT_EVENT_SCAN:
		HandleEventScan(ctx, msg)
	case wechat.WECHAT_EVENT_SUBSCRIBE:
		if msg.EventKey != "" {
			HandleEventScan(ctx, msg)
		} else {
			HandleEventSubscribe(ctx, msg)
		}

	}
}

func HandleEventScan(ctx *context.Context, msg wechat.WechatMsg) {
	replyStr := wechat.HandleScanEvent(msg)
	if replyStr == "" {
		log.Info("reply str is empty")
		return
	}
	reply := &wechat.MsgReply{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      wechat.WECHAT_MSG_TYPE_TEXT,
		Content:      replyStr,
	}
	ctx.XML(200, reply)
}

func HandleEventSubscribe(ctx *context.Context, msg wechat.WechatMsg) {
	r := wechat.HandleSubscribeEvent(msg)
	if r == nil {
		return
	}
	reply := buildReplyContent(msg, r)
	ctx.XML(200, reply)
}

func HandleTextMsg(ctx *context.Context, msg wechat.WechatMsg) {
	r := wechat.GetAutomaticReply(msg.Content)
	if r == nil {
		log.Info("TextMsg reply is empty")
		return
	}
	reply := buildReplyContent(msg, r)
	ctx.XML(200, reply)
}

func buildReplyContent(msg wechat.WechatMsg, r *wechat.WechatReplyContent) interface{} {
	reply := &wechat.MsgReply{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      r.ReplyType,
	}
	switch r.ReplyType {
	case wechat.ReplyTypeText:
		return &wechat.TextMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			Content:      r.Reply.Content,
		}

	case wechat.ReplyTypeImage:
		return &wechat.ImageMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			Image: wechat.ImageContent{
				MediaId: r.Reply.MediaId,
			},
		}
	case wechat.ReplyTypeVoice:
		return &wechat.VoiceMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			Voice: wechat.VoiceContent{
				MediaId: r.Reply.MediaId,
			},
		}
	case wechat.ReplyTypeVideo:
		return &wechat.VideoMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			Video: wechat.VideoContent{
				MediaId:     r.Reply.MediaId,
				Title:       r.Reply.Title,
				Description: r.Reply.Description,
			},
		}
	case wechat.ReplyTypeMusic:
		return &wechat.MusicMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			Music: wechat.MusicContent{
				Title:        r.Reply.Title,
				Description:  r.Reply.Description,
				MusicUrl:     r.Reply.MusicUrl,
				HQMusicUrl:   r.Reply.HQMusicUrl,
				ThumbMediaId: r.Reply.ThumbMediaId,
			},
		}
	case wechat.ReplyTypeNews:
		return &wechat.NewsMsgReply{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      r.ReplyType,
			ArticleCount: len(r.Reply.Articles),
			Articles: wechat.ArticleItem{
				Item: r.Reply.Articles},
		}

	}
	return reply
}
