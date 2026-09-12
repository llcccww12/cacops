package wechat_service

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

var WechatReplyCache = cache.New(2*time.Minute, 1*time.Minute)

const (
	WECHAT_REPLY_CACHE_KEY = "wechat_response"
)

const (
	ReplyTypeText  = "text"
	ReplyTypeImage = "image"
	ReplyTypeVoice = "voice"
	ReplyTypeVideo = "video"
	ReplyTypeMusic = "music"
	ReplyTypeNews  = "news"
)

type ReplyConfigType string

const (
	SubscribeReply ReplyConfigType = "subscribe"
	AutoMsgReply   ReplyConfigType = "autoMsg"
)

func (r ReplyConfigType) Name() string {
	switch r {
	case SubscribeReply:
		return "subscribe"
	case AutoMsgReply:
		return "autoMsg"
	}
	return ""
}

func (r ReplyConfigType) TreePath() string {
	switch r {
	case SubscribeReply:
		return setting.TreePathOfSubscribe
	case AutoMsgReply:
		return setting.TreePathOfAutoMsgReply
	}
	return ""
}

type WechatReplyContent struct {
	Reply       *ReplyContent
	ReplyType   string
	KeyWords    []string
	IsFullMatch int
}

type ReplyContent struct {
	Content      string
	MediaId      string
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
	Articles     []ArticlesContent
}

func GetAutomaticReply(msg string) *WechatReplyContent {
	r, err := LoadReplyFromCacheAndDisk(AutoMsgReply)
	if err != nil {
		return nil
	}
	if r == nil || len(r) == 0 {
		return nil
	}
	for i := 0; i < len(r); i++ {
		if r[i].IsFullMatch == 0 {
			for _, v := range r[i].KeyWords {
				if strings.Contains(msg, v) {
					return r[i]
				}
			}
		} else if r[i].IsFullMatch > 0 {
			for _, v := range r[i].KeyWords {
				if msg == v {
					return r[i]
				}
			}
		}
	}
	return nil

}

func loadReplyFromDisk(replyConfig ReplyConfigType) ([]*WechatReplyContent, error) {
	log.Info("LoadReply from  disk")
	repo, err := models.GetRepositoryByOwnerAndAlias(setting.UserNameOfWechatReply, setting.RepoNameOfWechatReply)
	if err != nil {
		log.Error("get AutomaticReply repo failed, error=%v", err)
		return nil, err
	}
	repoFile, err := models.ReadLatestFileInRepo(setting.UserNameOfWechatReply, repo.Name, setting.RefNameOfWechatReply, replyConfig.TreePath())
	if err != nil {
		log.Error("get AutomaticReply failed, error=%v", err)
		return nil, err
	}
	res := make([]*WechatReplyContent, 0)
	json.Unmarshal(repoFile.Content, &res)
	if res == nil || len(res) == 0 {
		return nil, err
	}
	return res, nil
}

func LoadReplyFromCacheAndDisk(replyConfig ReplyConfigType) ([]*WechatReplyContent, error) {
	v, success := WechatReplyCache.Get(replyConfig.Name())
	if success {
		log.Info("LoadReply from cache,value = %v", v)
		if v == nil {
			return nil, nil
		}
		n := v.([]*WechatReplyContent)
		return n, nil
	}

	content, err := loadReplyFromDisk(replyConfig)
	if err != nil {
		log.Error("LoadReply failed, error=%v", err)
		WechatReplyCache.Set(replyConfig.Name(), nil, 30*time.Second)
		return nil, err
	}
	WechatReplyCache.Set(replyConfig.Name(), content, 60*time.Second)
	return content, nil
}
