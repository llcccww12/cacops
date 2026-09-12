package pay_computility

import (
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/dynconfig"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
	"context"
	"sync"
	"time"
)

const (
	GitComputilityConfigPath            = "haosuan_config.json"
	GitParateraComputilityConfigPath    = "paratera.json"
	GitChinaMobileComputilityConfigPath = "ecloud.json"
)

type ComputilityConfigManager struct {
	ctx          context.Context
	hsConfig     *dynconfig.DyncConfigHelper
	peraConfig   *dynconfig.DyncConfigHelper
	mobileConfig *dynconfig.DyncConfigHelper
}

func InitComputilityConfigManager() *ComputilityConfigManager {
	return &ComputilityConfigManager{
		ctx:          context.Background(),
		hsConfig:     getHaosuanConfigHelper(),
		peraConfig:   getPerateraConfigHelper(),
		mobileConfig: getChinaMobileConfigHelper(),
	}
}

func getHaosuanConfigHelper() *dynconfig.DyncConfigHelper {
	gitComputilityConfigHelper := &dynconfig.DyncConfigHelper{}
	gitLocalOnce := &sync.Once{}

	gitLocalOnce.Do(func() {
		gitComputilityConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.PayConputilityConfigRepoOwner, setting.PayConputilityConfigRepoName, setting.PayConputilitycConfigBranch),
		)
	})
	return gitComputilityConfigHelper
}

func getPerateraConfigHelper() *dynconfig.DyncConfigHelper {
	gitComputilityConfigHelper := &dynconfig.DyncConfigHelper{}
	gitLocalOnce := sync.Once{}
	gitLocalOnce.Do(func() {
		gitComputilityConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.PeraPayConputilityConfigRepoOwner, setting.PeraPayConputilityConfigRepoName, setting.PeraPayConputilitycConfigBranch),
		)
	})
	return gitComputilityConfigHelper
}

func getChinaMobileConfigHelper() *dynconfig.DyncConfigHelper {
	gitComputilityConfigHelper := &dynconfig.DyncConfigHelper{}
	gitLocalOnce := sync.Once{}
	gitLocalOnce.Do(func() {
		gitComputilityConfigHelper = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.ChinaMobilePayConputilityConfigRepoOwner, setting.ChinaMobilePayConputilityConfigRepoName, setting.ChinaMobilePayConputilitycConfigBranch),
		)
	})
	return gitComputilityConfigHelper
}
