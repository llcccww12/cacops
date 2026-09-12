package dync_config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/notice"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/dynconfig"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/fetcher"

	"github.com/unknwon/com"
)

var (
	gitLocalConfigHelper         *dynconfig.DyncConfigHelper
	datasetStatisticConfigHelper *dynconfig.DyncConfigHelper
	gitLocalOnce                 sync.Once
	datasetStatisticOnce         sync.Once
)

func getGitLocalConfigHelper() *dynconfig.DyncConfigHelper {
	gitLocalOnce.Do(func() {
		gitLocalConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.DyncConfigRepoOwner, setting.DyncConfigRepoName, setting.DyncConfigBranch),
		)
	})
	return gitLocalConfigHelper
}

func getDatasetStatisticConfigHelper() *dynconfig.DyncConfigHelper {
	datasetStatisticOnce.Do(func() {
		datasetStatisticConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewDatasetStatisticFetcher(),
		)
	})
	return datasetStatisticConfigHelper
}

func loadCustomDynConfig(key string) interface{} {
	if key == "" {
		return nil
	}
	customPath := path.Join(setting.CustomPath, "conf", key)
	if !com.IsFile(customPath) {
		return nil
	}
	content, err := ioutil.ReadFile(customPath)
	if err != nil {
		log.Error("loadCustomDynConfig failed, path=%s, error=%v", customPath, err)
		return nil
	}
	return string(content)
}

func defaultDynConfigValue(key string) interface{} {
	switch {
	case strings.HasSuffix(key, "dataset_square.json"), strings.HasSuffix(key, "dataset_square_en.json"):
		return `{"tagInformation":{"category":{},"task":{},"license":{}}}`
	case strings.HasSuffix(key, "model_square.json"), strings.HasSuffix(key, "model_square_en.json"):
		return `{"tagInformation":{"frame":{},"label":{}}}`
	default:
		return nil
	}
}

func GetDynicConfig(ctx *context.APIContext) {
	key := ctx.Query("key")
	noCache := ctx.QueryBool("no_cache")
	val, err := dynconfig.GetGitLocalConfig(key, noCache)
	if err != nil {
		log.Error("GetDynicConfig failed, error=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	if val == nil {
		val = loadCustomDynConfig(key)
	}
	if val == nil {
		val = defaultDynConfigValue(key)
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(val))
	return
}

func GetHomeConfig(ctx *context.APIContext) {
	gitHelper := getGitLocalConfigHelper()
	datasetStatisticHelper := getDatasetStatisticConfigHelper()
	//获取公告，只返回visible=1的公告
	noticeList, _ := gitHelper.GetConfig("notice/notice.json", false)
	noticeListStr, ok := noticeList.(string)
	if ok {
		var noticeResponse notice.NoticeResponse
		err := json.Unmarshal([]byte(noticeListStr), &noticeResponse)
		if err == nil {
			noticeListJson := noticeResponse.Notices
			for i := 0; i < len(noticeListJson); i++ {
				visible := noticeListJson[i].Visible
				if visible == 0 {
					noticeListJson = append(noticeListJson[:i], noticeListJson[i+1:]...)
					i--
				}
			}
			noticeList = noticeListJson
		}
	}
	//获取推荐模型体验
	modelExperienceList, _ := gitHelper.GetConfig("home_v2/model_experience.json", false)
	//获取推荐项目
	repoList, _ := gitHelper.GetConfig("home_v2/repo.json", false)
	//获取推荐数据集
	datasetList, _ := datasetStatisticHelper.GetConfig("dataset_statistic", false)
	//获取推荐模型
	modelList, _ := gitHelper.GetConfig("home_v2/model.json", false)
	//获取推荐模板
	aiTaskTemplateList, _ := gitHelper.GetConfig("home_v2/ai_task_template.json", false)
	//获取推荐活动
	activityList, _ := gitHelper.GetConfig("home_v2/activity.json", false)

	r := make(map[string]interface{})
	r["Notice"] = noticeList
	r["ModelExperience"] = modelExperienceList
	r["Repo"] = repoList
	r["Dataset"] = datasetList
	r["Model"] = modelList
	r["AITaskTemplate"] = aiTaskTemplateList
	r["Activity"] = activityList
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
	return
}
