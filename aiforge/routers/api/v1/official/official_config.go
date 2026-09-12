package official

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"sync"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/notice"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/dynconfig"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
	"code.gitea.io/gitea/services/repository"

	"github.com/unknwon/com"
)

var (
	gitNoticeLocalConfigHelper *dynconfig.DyncConfigHelper
	gitLocalOnce               sync.Once
)

func getGitNoticeLocalConfigHelper() *dynconfig.DyncConfigHelper {
	gitLocalOnce.Do(func() {
		gitNoticeLocalConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.DyncConfigRepoOwner, setting.DyncConfigRepoName, setting.DyncConfigBranch),
		)
	})
	return gitNoticeLocalConfigHelper
}

// OfficialConfig 获取一些官方的配置数据
func OfficialConfig(ctx *context.APIContext) {
	var (
		wg   sync.WaitGroup
		resp = &models.OfficialConfig{}
	)

	// Step 1:  获取对应的公告
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetNotice(resp); err != nil {
			log.Error("GetNotice err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 2:  获取对应的官方活动图片
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetDashboardPicture(resp); err != nil {
			log.Error("GetNotice err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	wg.Wait()

	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
}

// GetNotice 获取公告列表
// 目前对于此列表展示的是visible为1的数据
func GetNotice(resp *models.OfficialConfig) error {
	//获取公告
	gitHelper := getGitNoticeLocalConfigHelper()
	noticeList, _ := gitHelper.GetConfig("notice/notice.json", false)
	noticeListStr, ok := noticeList.(string)
	if ok {
		var noticeResponse notice.NoticeResponse
		err := json.Unmarshal([]byte(noticeListStr), &noticeResponse)
		if err != nil {
			return err
		}
		noticeListJson := noticeResponse.Notices
		for i := 0; i < len(noticeListJson); i++ {
			visible := noticeListJson[i].Visible
			if visible == 0 {
				noticeListJson = append(noticeListJson[:i], noticeListJson[i+1:]...)
				i--
			}
		}

		// 转换标准格式
		noticeTrans := make([]*models.NoticeInfo, 0, len(noticeListJson))
		for _, noticeIndex := range noticeListJson {
			noticeTrans = append(noticeTrans, &models.NoticeInfo{
				Title:   noticeIndex.Title,
				TitleEn: noticeIndex.TitleEn,
				Link:    noticeIndex.Link,
				Visible: noticeIndex.Visible,
				Date:    noticeIndex.Date,
			})
		}

		resp.NoticeInfo.Notices = noticeTrans
	}
	if len(resp.NoticeInfo.Notices) == 0 {
		resp.NoticeInfo.Notices = loadFallbackNotices()
	}
	return nil
}

func loadFallbackNotices() []*models.NoticeInfo {
	customPath := path.Join(setting.CustomPath, "conf", "notice.json")
	if com.IsFile(customPath) {
		content, err := ioutil.ReadFile(customPath)
		if err == nil {
			var noticeResponse notice.NoticeResponse
			if err := json.Unmarshal(content, &noticeResponse); err == nil {
				notices := make([]*models.NoticeInfo, 0, len(noticeResponse.Notices))
				for _, noticeIndex := range noticeResponse.Notices {
					if noticeIndex.Visible == 0 {
						continue
					}
					notices = append(notices, &models.NoticeInfo{
						Title:   noticeIndex.Title,
						TitleEn: noticeIndex.TitleEn,
						Link:    noticeIndex.Link,
						Visible: noticeIndex.Visible,
						Date:    noticeIndex.Date,
					})
				}
				if len(notices) > 0 {
					return notices
				}
			}
		}
	}
	return []*models.NoticeInfo{
		{
			Title:   "欢迎使用 CacOps 智算服务平台",
			TitleEn: "Welcome to CacOps",
			Link:    "/dashboard",
			Visible: 1,
			Date:    time.Now().Format("2006-01-02"),
		},
	}
}

// GetDashboardPicture 获取面板的图片
func GetDashboardPicture(resp *models.OfficialConfig) error {
	pictureInfo, err := getImageInfo("dashboard-picture")
	if err == nil && len(pictureInfo) > 0 {
		log.Info("set image info=" + pictureInfo[0]["url"])
		resp.ActivityImage = models.ActivityImageOverview{
			ImageUrl:  pictureInfo[0]["url"],
			ImageLink: pictureInfo[0]["image_link"],
		}

		if len(pictureInfo) > 1 {
			resp.ActivityImage.InviteImageUrl = pictureInfo[1]["url"]
			resp.ActivityImage.InviteImageLink = pictureInfo[1]["image_link"]
		}
	}
	return nil
}

// getImageInfo 沿用方法
func getImageInfo(filename string) ([]map[string]string, error) {
	url := setting.RecommentRepoAddr + filename
	result, err := repository.RecommendFromPromote(url)

	if err != nil {
		return nil, err
	}
	imageInfo := make([]map[string]string, 0)
	for i := 0; i < (len(result) - 1); i++ {
		line := result[i]
		imageMap := make(map[string]string)
		if line[0:4] == "url=" {
			url := line[4:]
			imageMap["url"] = url
			if result[i+1][0:11] == "image_link=" {
				image_link := result[i+1][11:]
				imageMap["image_link"] = image_link
			}
		}
		imageInfo = append(imageInfo, imageMap)
		i = i + 1
	}
	return imageInfo, nil
}
