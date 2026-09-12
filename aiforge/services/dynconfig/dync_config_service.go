package dynconfig

import (
	"sync"
	"time"

	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
)

var (
	gitLocalConfigHelper *DyncConfigHelper
	gitLocalOnce         sync.Once
)

func getGitLocalConfigHelper() *DyncConfigHelper {
	gitLocalOnce.Do(func() {
		gitLocalConfigHelper = NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.DyncConfigRepoOwner, setting.DyncConfigRepoName, setting.DyncConfigBranch),
		)
	})
	return gitLocalConfigHelper
}

func GetGitLocalConfig(key string, noCache bool) (interface{}, error) {
	gitHelper := getGitLocalConfigHelper()
	return gitHelper.GetConfig(key, noCache)
}
