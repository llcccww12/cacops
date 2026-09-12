package notice

import (
	"encoding/json"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/util"
	"github.com/patrickmn/go-cache"
)

var noticeCache = cache.New(2*time.Minute, 1*time.Minute)

const (
	NOTICE_CACHE_KEY = "notice"
)

type Notice struct {
	Title   string
	TitleEn string
	Link    string
	Visible int //0 invisible, 1 visible
	Date    string
}

type NoticeResponse struct {
	Notices  []*Notice
	CommitId string
}

var lock int32 = 0

func GetNewestNotice() (*NoticeResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recover error", err)
		}
	}()

	notice, err := getNewestNoticesFromCacheAndDisk()
	if err != nil {
		return nil, err
	}
	return notice, nil
}

func getNoticeTimeout() time.Duration {
	return time.Duration(setting.CacheTimeOutSecond) * time.Second
}

func getNewestNoticesFromDisk() (*NoticeResponse, error) {
	log.Debug("Get notice from disk")
	repo, err := models.GetRepositoryByOwnerAndAlias(setting.UserNameOfNoticeRepo, setting.RepoNameOfNoticeRepo)
	if err != nil {
		log.Error("get notice repo failed, error=%v", err)
		return nil, err
	}
	repoFile, err := models.ReadLatestFileInRepo(repo.OwnerName, repo.Name, setting.RefNameOfNoticeRepo, setting.TreePathOfNoticeRepo)
	if err != nil {
		log.Error("GetNewestNotice failed, error=%v", err)
		return nil, err
	}
	res := &NoticeResponse{}
	json.Unmarshal(repoFile.Content, res)
	if res == nil || len(res.Notices) == 0 {
		return nil, err
	}
	res.CommitId = repoFile.CommitId
	return res, nil
}

var updateNoticeLock util.NonBlockingLock
var defaultNotice *NoticeResponse

func getNewestNoticesFromCacheAndDisk() (*NoticeResponse, error) {
	v, success := noticeCache.Get(NOTICE_CACHE_KEY)
	if success {
		log.Debug("Get notice from cache,value = %v", v)
		if v == nil {
			return nil, nil
		}
		n := v.(*NoticeResponse)
		return n, nil
	}
	if updateNoticeLock.TryLock() {
		defer updateNoticeLock.Unlock()
		v, success = noticeCache.Get(NOTICE_CACHE_KEY)
		if success {
			log.Debug("Get notice from cache,value = %v", v)
			if v == nil {
				return nil, nil
			}
			n := v.(*NoticeResponse)
			return n, nil
		}
		log.Info("try to update notice")
		notice, err := getNewestNoticesFromDisk()
		if err != nil {
			log.Error("GetNewestNotice failed, error=%v", err)
			noticeCache.Set(NOTICE_CACHE_KEY, nil, 30*time.Second)
			return nil, err
		}
		noticeCache.Set(NOTICE_CACHE_KEY, notice, getNoticeTimeout())
		defaultNotice = notice
		return notice, nil
	}
	//未拿到锁的线程返回旧值
	return defaultNotice, nil
}
