package pay_computility

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/routers/response"
	"net/http"
)

// GetComputilityConfig 获取付费算力的配置
func GetComputilityConfig(ctx *context.APIContext) {
	// step 1: 初始化算力管理器
	cm := InitComputilityManager(ctx)

	// step 2: 算力过滤/获取
	_ = cm.ComputilityFilter()

	// step 3: 如果返回为空，塞入默认值返回
	if len(cm.filterData.Products) == 0 {
		cm.filterData = GetFilterData()
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(cm.filterData))
	return
}
