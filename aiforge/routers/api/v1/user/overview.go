package user

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/reward/point/account"
	"code.gitea.io/gitea/services/storage_limit"
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const TTL = 3 // 缓存过期时间3秒

// Overview 获取当前用户的概览页面（采用go group并发获取数据的方式）
func Overview(ctx *context.APIContext) {
	var (
		wg        sync.WaitGroup
		resp      = &models.Overview{}
		reqAction = ctx.Query("action")
	)

	log.Info("Overview start reqId[%d] reqAction[%v]", ctx.ReqId, reqAction)

	// 防止panic
	defer func() {
		if err := recover(); err != nil {
			stack := make([]byte, 64*1024)       // 64KB
			length := runtime.Stack(stack, true) // true: 包含所有goroutine

			log.Error("overview panic:[%v] reqId[%v] \n 堆栈信息:\n %s",
				err, ctx.ReqId, string(stack[:length]))
		}
	}()

	if ctx.User == nil {
		log.Error("ctx.User is nil")
		ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
		return
	}

	// 如果命中redis直接返回
	resp = GetCacheOverview(ctx, ctx.User.ID)
	if resp != nil && reqAction != "no-cache" {
		ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
		return
	} else {
		// 重新初始化
		resp = &models.Overview{}
	}

	// Step 1: 获取用户绑定信息
	GetUserIsBinding(ctx, resp)

	// Step 2: 获取计算任务总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetCountCloudBrains(ctx, resp); err != nil {
			log.Error("GetCountCloudBrains err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 3: 获取存储配额总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetUserStorageSummary(ctx, resp); err != nil {
			log.Error("GetUserStorageSummary err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 4:  获取对应的积分配置
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetCountUserPoint(ctx, resp); err != nil {
			log.Error("GetUserStorageSummary err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 5:  获取项目总数量
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetCountRepository(ctx, resp); err != nil {
			log.Error("GetCountRepository err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 6:  获取数据集总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetCountDataset(ctx, resp); err != nil {
			log.Error("GetCountDataset err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	// Step 7:  获取模型总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := GetCountModel(ctx, resp); err != nil {
			log.Error("GetCountModel err:%v reqId[%d]", err, ctx.ReqId)
			return
		}
	}()

	wg.Wait()

	go func() {
		// 异步设置缓存
		SetCacheOverview(ctx, ctx.User.ID, resp)
	}()

	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
	return
}

// GetUserIsBinding 查看用户是否已经绑定
func GetUserIsBinding(ctx *context.APIContext, resp *models.Overview) {
	resp.BindingInfo = models.BindingOverview{
		IsBindWechat: ctx.User.WechatOpenId != "",
	}
}

// GetCountCloudBrains 统计云脑的数据
func GetCountCloudBrains(ctx *context.APIContext, resp *models.Overview) error {
	// 获取我得所有云脑任务
	allTaskCount, err := models.CountCloudBrainsByUserId(ctx.User.ID)
	if err != nil {
		return err
	}

	// 获取我的（运行中）的任务数量
	runningTaskCount, err := models.CountCloudBrainsByJobStatus(ctx.User.ID, models.JobRunning)
	if err != nil {
		return err
	}

	allCardDurationCount, err := models.CountAccCardTotalDuration(ctx.User.ID)
	if err != nil {
		return err
	}

	// 计算任务模板
	templatesCount, useTotal, err := models.CountUserTemplates(ctx.User.ID)
	if err != nil {
		return err
	}

	resp.AiTaskInfo = models.AiTaskOverview{
		TotalAITasks:         allTaskCount,
		RunningAITasks:       runningTaskCount,
		AllCardDurationCount: allCardDurationCount / 3600,
		TemplateCount:        templatesCount,
		TemplateUseCount:     useTotal,
	}

	return nil
}

// GetCountRepository 获取项目总数
// 参考接口
func GetCountRepository(ctx *context.APIContext, resp *models.Overview) error {
	// 获取项目的总数
	sum, sourceSum, forkSum, mirrorSum, collaborateSum, err := models.CountRepositoryOverView(ctx.User.ID)
	if err != nil {
		return err
	}

	resp.RepositoryInfo = models.RepositoryOverview{
		RepositoryCount:    sum,
		SelfBuiltCount:     sourceSum,
		ForkedCount:        forkSum,
		MirroredCount:      mirrorSum,
		CollaborationCount: collaborateSum,
	}
	return nil
}

// GetCountDataset 获取数据集总数
func GetCountDataset(ctx *context.APIContext, resp *models.Overview) error {

	// 数据集的数量
	allCount, privateCount, publicCount, err := models.CountPersonalDataset(ctx.User.ID)
	if err != nil {
		return err
	}

	resp.DatasetInfo = models.DatasetOverview{
		DatasetCount: allCount,
		PublicCount:  publicCount,
		PrivateCount: privateCount,
	}
	return nil
}

// GetCountModel 获取模型总数
func GetCountModel(ctx *context.APIContext, resp *models.Overview) error {
	// 模型的数量
	allCount, privateCount, publicCount, err := models.CountPersonalAiModels(ctx.User.ID)
	if err != nil {
		return err
	}

	resp.ModelInfo = models.ModelOverview{
		ModelCount:   allCount,
		PublicCount:  publicCount,
		PrivateCount: privateCount,
	}
	return nil
}

// GetUserStorageSummary 获取用户的存储资源配置
func GetUserStorageSummary(ctx *context.APIContext, resp *models.Overview) error {
	// 获取用户的存储资源
	summary, err := storage_limit.GetUserStorageSummary(ctx.User)
	if err != nil {
		return err
	}

	resp.StorageInfo = models.StorageOverview{
		StorageLimit:       summary.StorageLimit,
		DatasetUsedStorage: summary.DatasetUsedStorage,
		ModelUsedStorage:   summary.ModelUsedStorage,
		UsedStorage:        summary.UsedStorage,
		RemainingStorage:   summary.RemainingStorage,
	}
	return nil
}

// GetCountUserPoint 获取用户的积分
func GetCountUserPoint(ctx *context.APIContext, resp *models.Overview) error {
	pointCount, err := account.GetAccount(ctx.User.ID)
	if err != nil {
		return err
	}

	resp.PointInfo = models.PointOverview{
		TotalEarned:   pointCount.TotalEarned,
		Balance:       pointCount.Balance,
		TotalConsumed: pointCount.TotalConsumed,
	}
	return nil
}

// GetCacheOverview 快速获取数据（redis）
func GetCacheOverview(ctx *context.APIContext, userId int64) *models.Overview {
	userOverviewKey := redis_key.GetOverviewKeyByUserId(userId)
	val, err := redis_client.Get(userOverviewKey)
	if val != "" {
		resp := &models.Overview{}
		_ = json.Unmarshal([]byte(val), resp)
		return resp
	}

	if err != nil {
		log.Error("GetCacheOverview key [%] err:%v reqId[%d]", userId, err, ctx.ReqId)
	}

	return nil
}

// SetCacheOverview 设置缓存数据（redis）
func SetCacheOverview(ctx *context.APIContext, userId int64, resp *models.Overview) {
	userOverviewKey := redis_key.GetOverviewKeyByUserId(userId)
	byteResp, _ := json.Marshal(resp)
	_, err := redis_client.Setex(userOverviewKey, string(byteResp), TTL*time.Second)
	if err != nil {
		log.Error("GetCacheOverview key [%] err:%v reqId[%d]", userId, err, ctx.ReqId)
	}

	return
}
