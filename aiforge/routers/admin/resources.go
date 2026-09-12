package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

const (
	tplResourceQueue         base.TplName = "admin/resources/queue"
	tplResourceSpecification base.TplName = "admin/resources/specification"
	tplResourceScene         base.TplName = "admin/resources/scene"
)

func GetQueuePage(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminResources"] = true
	ctx.Data["PageIsAdminResourcesQueue"] = true
	ctx.HTML(200, tplResourceQueue)
}

func GetSpecificationPage(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminResources"] = true
	ctx.Data["PageIsAdminResourcesSpecification"] = true
	ctx.HTML(200, tplResourceSpecification)
}

func GetScenePage(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminResources"] = true
	ctx.Data["PageIsAdminResourcesScene"] = true
	ctx.HTML(200, tplResourceScene)
}

func GetResourceQueueList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	cluster := ctx.Query("cluster")
	aiCenterCode := ctx.Query("center")
	computeResource := ctx.Query("resource")
	accCardType := ctx.Query("card")
	hasInternet := ctx.QueryInt("hasInternet")
	queueType := ctx.Query("queueType")
	isAvailable := ctx.QueryInt("isAvailable")
	isQueueExclusive := ctx.QueryInt("isQueueExclusive")
	enableVisualization := ctx.QueryInt("enableVisualization")

	if pageSize > 1000 {
		log.Error("GetResourceQueueList pageSize too large.")
		ctx.JSON(http.StatusOK, response.ServerError("pageSize too large"))
		return
	}
	list, err := resource.GetResourceQueueList(models.SearchResourceQueueOptions{
		ListOptions:         models.ListOptions{Page: page, PageSize: pageSize},
		Cluster:             cluster,
		AiCenterCode:        aiCenterCode,
		ComputeResource:     computeResource,
		AccCardType:         accCardType,
		HasInternet:         models.SpecInternetQuery(hasInternet),
		QueueType:           queueType,
		IsAvailable:         isAvailable,
		IsQueueExclusive:    isQueueExclusive,
		EnableVisualization: enableVisualization,
	})
	if err != nil {
		log.Error("GetResourceQueueList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(list))
}

func GetResourceQueueCodes(ctx *context.Context) {
	var cluster []string
	clusterStr := ctx.Query("cluster")
	if clusterStr != "" {
		cluster = strings.Split(clusterStr, "||")
	}
	list, err := resource.GetResourceQueueCodes(models.GetQueueCodesOptions{Cluster: cluster})
	if err != nil {
		log.Error("GetResourceQueueCodes error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(list))
}

func GetResourceAiCenters(ctx *context.Context) {
	list, err := resource.GetResourceAiCenters()
	if err != nil {
		log.Error("GetResourceAiCenters error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	for i := 0; i < len(list); i++ {
		list[i].Tr(ctx.Language())
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(list))
}

func AddResourceQueue(ctx *context.Context, req models.ResourceQueueReq) {
	req.IsAutomaticSync = false
	req.CreatorId = ctx.User.ID
	err := resource.AddResourceQueue(req)
	if err != nil {
		log.Error("AddResourceQueue error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func UpdateResourceQueue(ctx *context.Context, req models.ResourceQueueReq) {
	queueId := ctx.ParamsInt64(":id")
	err := resource.UpdateResourceQueue(queueId, req)
	if err != nil {
		log.Error("UpdateResourceQueue error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func SyncGrampusQueue(ctx *context.Context) {
	err := resource.SyncGrampusQueue(ctx.User.ID)
	if err != nil {
		log.Error("AddResourceQueue error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	err = resource.SyncGrampusAICenter()
	if err != nil {
		log.Error("SyncGrampusAICenter error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func SyncGrampusImage(ctx *context.Context) {
	err := resource.SyncGrampusImage(ctx.User.ID)
	if err != nil {
		log.Error("sync image error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func GetResourceSpecificationList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	queue := ctx.QueryInt64("queue")
	status := ctx.QueryInt("status")
	cluster := ctx.Query("cluster")
	aiCenterCode := ctx.Query("center")
	available := ctx.QueryInt("available")
	cardsNum := ctx.QueryInt("cardsNum")
	computeResource := ctx.Query("resource")
	cardType := ctx.Query("cardType")
	hasInternet := ctx.QueryInt("hasInternet")
	enableVisualization := ctx.QueryInt("enableVisualization")

	if pageSize > 1000 {
		log.Error("GetResourceSpecificationList pageSize too large.")
		ctx.JSON(http.StatusOK, response.ServerError("pageSize too large"))
		return
	}
	list, err := resource.GetResourceSpecificationList(models.SearchResourceSpecificationOptions{
		ListOptions:         models.ListOptions{Page: page, PageSize: pageSize},
		QueueId:             queue,
		Status:              status,
		Cluster:             cluster,
		AiCenterCode:        aiCenterCode,
		AvailableCode:       available,
		OrderBy:             models.SearchSpecOrderById,
		AccCardsNum:         cardsNum,
		ComputeResource:     computeResource,
		AccCardType:         cardType,
		HasInternet:         models.SpecInternetQuery(hasInternet),
		EnableVisualization: enableVisualization,
	})
	if err != nil {
		log.Error("GetResourceSpecificationList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(list))
}

func GetAllResourceSpecificationList(ctx *context.Context) {
	status := ctx.QueryInt("status")
	cluster := ctx.Query("cluster")
	available := ctx.QueryInt("available")
	computeResource := ctx.Query("resource")
	list, err := resource.GetAllResourceSpecification(models.SearchResourceSpecificationOptions{
		Status:          status,
		Cluster:         cluster,
		AvailableCode:   available,
		AccCardsNum:     -1,
		ComputeResource: computeResource,
	})
	if err != nil {
		log.Error("GetResourceSpecificationList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	r := map[string]interface{}{"Specs": list}
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func GetResourceSpecificationScenes(ctx *context.Context) {
	specId := ctx.ParamsInt64(":id")
	list, err := resource.GetResourceSpecificationScenes(specId)
	if err != nil {
		log.Error("GetResourceSpecificationScenes error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	rolename := []string{}
	roleList, err := models.ListAiforgeRole("")
	if err != nil {
		log.Info("ListAiforgeRole error." + err.Error())
	}
	for _, role := range roleList {
		rightInfo := role.RightInfo
		operRights := make([]*models.RightInfo, 0)
		if rightInfo != "" {
			jsonerr := json.Unmarshal([]byte(rightInfo), &operRights)
			if jsonerr != nil {
				log.Info("un json error=" + jsonerr.Error())
			} else {
				for _, right := range operRights {
					if right.SpecId == fmt.Sprint(specId) {
						rolename = append(rolename, role.Name)
					}
				}
			}
		}
	}

	r := make(map[string]interface{})
	r["List"] = list
	r["RoleName"] = rolename
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func AddResourceSpecification(ctx *context.Context, req models.ResourceSpecificationReq) {
	req.IsAutomaticSync = false
	req.CreatorId = ctx.User.ID
	err := resource.AddResourceSpecification(ctx.User.ID, req)
	if err != nil {
		log.Error("AddResourceQueue error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func UpdateResourceSpecification(ctx *context.Context, req models.ResourceSpecificationReq) {
	id := ctx.ParamsInt64(":id")
	action := ctx.Query("action")
	comment := ctx.Query("comment")

	var err *response.BizError
	switch action {
	case "edit":
		if req.UnitPrice < 0 {
			ctx.JSON(http.StatusOK, response.ServerError("param error"))
			return
		}
		//only UnitPrice and  permitted to change
		err = resource.UpdateSpecUnitPrice(ctx.User.ID, id, req.UnitPrice, comment)
	case "on-shelf":
		err = resource.ResourceSpecOnShelf(ctx.User.ID, id, req.UnitPrice, comment)
	case "off-shelf":
		err = resource.ResourceSpecOffShelf(ctx.User.ID, id, comment)
	}

	if err != nil {
		log.Error("UpdateResourceSpecification error. %v", err)
		ctx.JSON(http.StatusOK, response.ResponseBizError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

// 获取资源规格操作日志
func GetResourceSpecLogs(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")
	specId := ctx.QueryInt("spec_id")       // 筛选特定规格
	operateType := ctx.Query("OperateType") // 筛选操作类型 (on-shelf, off-shelf, edit)
	startTime := ctx.Query("start_time")    // 格式: 时间戳
	endTime := ctx.Query("end_time")        // 格式: 时间戳
	computeResource := ctx.Query("compute_resource")
	accCardType := ctx.Query("acc_card_type")
	accCardNum := ctx.QueryInt("acc_card_num")

	aiCenterCode := ctx.Query("ai_center_code")
	resourceQueueID := ctx.QueryInt("resource_queue_id")

	opts := models.SearchSpecLogOptions{
		Page:            page,
		PageSize:        pageSize,
		SpecId:          specId,
		OperateType:     operateType,
		StartTime:       startTime,
		EndTime:         endTime,
		ComputeResource: computeResource,
		AccCardType:     accCardType,
		AccCardNum:      accCardNum,
		AiCenterCode:    aiCenterCode,
		ResourceQueueID: resourceQueueID,
	}

	total, logs, err := models.SearchResourceSpecLogs(opts)

	if err != nil {
		log.Error("SearchResourceSpecLogs error: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ServerError("查询失败"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(map[string]interface{}{
		"list":      logs,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}))
}

func SyncGrampusSpecs(ctx *context.Context) {
	err := resource.SyncGrampusSpecs(ctx.User.ID)
	if err != nil {
		log.Error("AddResourceQueue error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func GetResourceSceneList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	jobType := ctx.Query("jobType")
	aiCenterCode := ctx.Query("center")
	queueId := ctx.QueryInt64("queue")
	isExclusive := ctx.Query("isSpecExclusive")
	sceneType := ctx.Query("sceneType")
	computeResource := ctx.Query("resource")
	cardType := ctx.Query("cardType")
	cluster := ctx.Query("cluster")
	hasInternet := ctx.QueryInt("hasInternet")
	enableVisualization := ctx.QueryInt("enableVisualization")

	if pageSize > 1000 {
		log.Error("GetResourceSceneList pageSize too large.")
		ctx.JSON(http.StatusOK, response.ServerError("pageSize too large"))
		return
	}
	list, err := resource.GetResourceSceneList(models.SearchResourceSceneOptions{
		ListOptions:         models.ListOptions{Page: page, PageSize: pageSize},
		JobType:             jobType,
		IsSpecExclusive:     isExclusive,
		AiCenterCode:        aiCenterCode,
		QueueId:             queueId,
		ComputeResource:     computeResource,
		AccCardType:         cardType,
		Cluster:             cluster,
		HasInternet:         models.SpecInternetQuery(hasInternet),
		SceneType:           sceneType,
		EnableVisualization: enableVisualization,
	})
	if err != nil {
		log.Error("GetResourceSceneList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(list))
}

func AddResourceScene(ctx *context.Context, req models.ResourceSceneReq) {
	req.CreatorId = ctx.User.ID
	req.ExclusiveOrg = strings.ReplaceAll(req.ExclusiveOrg, " ", "")
	err := resource.AddResourceScene(req)
	if err != nil {
		log.Error("AddResourceScene error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func UpdateResourceScene(ctx *context.Context, req models.ResourceSceneReq) {
	id := ctx.ParamsInt64(":id")
	action := ctx.Query("action")

	req.ID = id
	var err error
	switch action {
	case "edit":
		req.ExclusiveOrg = strings.ReplaceAll(req.ExclusiveOrg, " ", "")
		err = resource.UpdateResourceScene(req)
	case "delete":
		err = resource.DeleteResourceScene(id)
	}

	if err != nil {
		log.Error("UpdateResourceScene error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func RefreshHistorySpec(ctx *context.Context) {
	scope := ctx.Query("scope")
	list := ctx.Query("list")

	var scopeAll = false
	if scope == "all" {
		scopeAll = true
	}
	var ids = make([]int64, 0)
	if list != "" {
		strs := strings.Split(list, "|")
		for _, s := range strs {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
				return
			}
			ids = append(ids, i)
		}

	}

	total, success, err := resource.RefreshHistorySpec(scopeAll, ids)
	if err != nil {
		log.Error("RefreshHistorySpec error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	r := make(map[string]interface{}, 0)
	r["success"] = success
	r["total"] = total
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func RefreshReposHistoryCnt(ctx *context.Context) {
	scope := ctx.Query("scope")
	list := ctx.Query("list")

	var scopeAll = false
	if scope == "all" {
		scopeAll = true
	}
	var ids = make([]int64, 0)
	if list != "" {
		strs := strings.Split(list, "|")
		for _, s := range strs {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
				return
			}
			ids = append(ids, i)
		}

	}

	total, success, err := resource.RefreshHistorySpec(scopeAll, ids)
	if err != nil {
		log.Error("RefreshHistorySpec error. %v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	r := make(map[string]interface{}, 0)
	r["success"] = success
	r["total"] = total
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}
