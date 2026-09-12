package pay_computility

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/dynconfig"
	"code.gitea.io/gitea/services/dynconfig/config_cache"
	"code.gitea.io/gitea/services/dynconfig/fetcher"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

var partConfig *dynconfig.DyncConfigHelper

const GitComputilityPartnerConfigPath = "computility/computility_partners.json"

func getPartnerConfigHelper() *dynconfig.DyncConfigHelper {

	partConfig = &dynconfig.DyncConfigHelper{}
	gitLocalOnce := sync.Once{}
	gitLocalOnce.Do(func() {
		partConfig = dynconfig.NewDyncConfigHelper(
			config_cache.NewLocalCache(2*time.Minute, 1*time.Minute),
			fetcher.NewGitLocalFetcher(setting.ConputilityPartnerConfigRepoOwner, setting.ConputilityPartnerConfigRepoName, setting.ConputilityPartnerConfigBranch),
		)
	})

	return partConfig
}

// ComputilityPartner 获取算力合作伙伴
func ComputilityPartner(ctx *context.APIContext) {
	var (
		resp   = ResponseCompany{}
		helper = getPartnerConfigHelper()
	)

	partConfigList, _ := helper.GetConfig(GitComputilityPartnerConfigPath, false)

	if str, ok := partConfigList.(string); ok && partConfigList != nil {
		if err := json.Unmarshal([]byte(str), &resp); err != nil {
			ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
			return
		}
	}

	if len(resp.AiCenter) == 0 {
		resp = GetCompanies()
	}

	// 排序
	SortCompany(resp.AiCenter)
	SortCompany(resp.ChipVendor)
	SortCompany(resp.ComputingPlatform)

	// 获取算力合作伙伴配置
	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
	return
}

// SortCompany 排序【首先按 Sort 从大到小排序;Sort 相同时，按字母顺序 A-Z 排序】
func SortCompany(company []Company) {
	sort.SliceStable(company, func(i, j int) bool {
		// 首先按 Sort 从大到小排序
		if company[i].Sort != company[j].Sort {
			return company[i].Sort > company[j].Sort
		}

		// Sort 相同时，按字母顺序 A-Z 排序
		cardTypeI := strings.ToUpper(company[i].Provider)
		cardTypeJ := strings.ToUpper(company[j].Provider)

		return cardTypeI < cardTypeJ
	})
}
