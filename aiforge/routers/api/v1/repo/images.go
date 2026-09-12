package repo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	routeRepo "code.gitea.io/gitea/routers/repo"
)

type NPUImageINFO struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func GetRecommendImages(ctx *context.APIContext) {
	uid := getUID(ctx)
	var keyword string
	if ctx.Query("jobType") == string(models.JobTypeComfyuiExperience) {
		keyword = ctx.Query("q") + "," + models.DefaultComfyuiTag
	} else {
		keyword = ctx.Query("q")
	}
	opts := models.SearchImageOptions{
		UID:                 uid,
		Keyword:             keyword,
		Topics:              ctx.Query("topic"),
		IncludeOfficialOnly: true,
		Status:              models.IMAGE_STATUS_SUCCESS,
		CloudbrainType:      ctx.QueryInt("cloudbrainType"),
		TrainType:           ctx.Query("trainType"),
	}

	routeRepo.GetImages(ctx.Context, &opts)

}

func GetPublicImages(ctx *context.APIContext) {
	uid := getUID(ctx)
	opts := models.SearchImageOptions{
		IncludePublicOnly:   true,
		UID:                 uid,
		Keyword:             ctx.Query("q"),
		Topics:              ctx.Query("topic"),
		IncludeOfficialOnly: ctx.QueryBool("recommend"),
		Status:              models.IMAGE_STATUS_SUCCESS,
		CloudbrainType:      ctx.QueryInt("cloudbrainType"),
	}

	routeRepo.GetImages(ctx.Context, &opts)

}

func GetCustomImages(ctx *context.APIContext) {
	uid := getUID(ctx)
	var keyword string
	if ctx.Query("jobType") == string(models.JobTypeComfyuiExperience) {
		keyword = ctx.Query("q") + "," + models.DefaultComfyuiTag
	} else {
		keyword = ctx.Query("q")
	}
	opts := models.SearchImageOptions{
		UID:              uid,
		IncludeOwnerOnly: true,
		Keyword:          keyword,
		Topics:           ctx.Query("topic"),
		Status:           -1,
		CloudbrainType:   ctx.QueryInt("cloudbrainType"),
		TrainType:        ctx.Query("trainType"),
	}
	routeRepo.GetImages(ctx.Context, &opts)

}
func GetStarImages(ctx *context.APIContext) {

	uid := getUID(ctx)
	var keyword string
	if ctx.Query("jobType") == string(models.JobTypeComfyuiExperience) {
		keyword = ctx.Query("q") + "," + models.DefaultComfyuiTag
	} else {
		keyword = ctx.Query("q")
	}
	opts := models.SearchImageOptions{
		UID:             uid,
		IncludeStarByMe: true,
		Keyword:         keyword,
		Topics:          ctx.Query("topic"),
		Status:          models.IMAGE_STATUS_SUCCESS,
		CloudbrainType:  ctx.QueryInt("cloudbrainType"),
		TrainType:       "",
	}
	routeRepo.GetImages(ctx.Context, &opts)

}

func GetNpuImages(ctx *context.APIContext) {
	cloudbrainType := ctx.QueryInt("type")
	jobType := ctx.Query("jobType")
	if cloudbrainType == 0 { //modelarts
		getModelArtsImages(ctx)
	} else { //c2net
		getC2netNpuImages(ctx, jobType)
	}
}

func GetAvailableFilerInfo(ctx *context.APIContext) {
	columns := []string{"framework", "framework_version", "python_version", "cuda_version"}

	i := ctx.QueryInt("index")

	frameWork := ctx.Query("framework")
	version := ctx.Query("version")
	pythonVersion := ctx.Query("python")
	compute_resource := ctx.Query("compute_resource")
	onlyRecommend := ctx.QueryBool("recommend")
	includeMineOnly := ctx.QueryBool("mine")
	starByMe := ctx.QueryBool("star")
	jobType := ctx.Query("job_type")
	hasInternet := ctx.QueryInt("has_internet")
	visualizeRequired := ctx.QueryBool("visualize_required")
	if i < 0 || i >= len(columns) {
		i = 0
	}
	columnName := columns[i]

	if compute_resource == models.NPUResource && i == 3 {
		columnName = "cann_version"
	}

	opts := &models.SearchAvailableValueOptions{
		Column:           columnName,
		Framework:        frameWork,
		FrameworkVersion: version,
		PythonVersion:    pythonVersion,
		ComputeResource:  compute_resource,
		OnlyRecommend:    onlyRecommend,
		IncludeOwnerOnly: includeMineOnly,
		IncludeStarByMe:  starByMe,
		UID:              getUID(ctx),
	}

	specId := ctx.QueryInt64("spec")
	if specId > 0 {
		spec, _ := models.GetSpecificationById(specId)
		if spec != nil {
			queues := spec.GetAvailableQueues(models.GetAvailableCenterIdOpts{
				UserId:            opts.UID,
				JobType:           models.JobType(jobType),
				HasInternet:       models.SpecInternetQuery(hasInternet),
				VisualizeRequired: visualizeRequired,
			})
			if len(queues) > 0 {
				var aiCenterCodes []string
				for i := 0; i < len(queues); i++ {
					q := queues[i]
					aiCenterCodes = append(aiCenterCodes, q.AiCenterCode)
				}
				opts.AiCenterIds = aiCenterCodes
			}
		}
	}

	ctx.JSON(http.StatusOK, models.BaseMessageWithDataApi{
		Data: models.GetImageAvailableColumnValues(opts),
	})

}

func getModelArtsImages(ctx *context.APIContext) {

	var versionInfos modelarts.VersionInfo
	_ = json.Unmarshal([]byte(setting.EngineVersions), &versionInfos)
	var npuImageInfos []NPUImageINFO
	for _, info := range versionInfos.Version {
		npuImageInfos = append(npuImageInfos, NPUImageINFO{
			ID:    strconv.Itoa(info.ID),
			Value: info.Value,
		})
	}
	ctx.JSON(http.StatusOK, npuImageInfos)

}

func getC2netNpuImages(ctx *context.APIContext, jobType string) {
	jobTypeQuery := string(models.JobTypeTrain)
	if jobType == string(models.JobTypeDebug) {
		jobTypeQuery = string(models.JobTypeDebug)
	}
	images, err := grampus.GetImages(grampus.ProcessorTypeNPU, jobTypeQuery)
	var npuImageInfos []NPUImageINFO
	if err != nil {
		log.Error("GetImages failed:", err.Error())
		ctx.JSON(http.StatusOK, []NPUImageINFO{})
	} else {
		for _, info := range images.Infos {
			npuImageInfos = append(npuImageInfos, NPUImageINFO{
				ID:    info.ID,
				Value: info.Name,
			})
		}
		ctx.JSON(http.StatusOK, npuImageInfos)
	}
}

func getUID(ctx *context.APIContext) int64 {
	var uid int64 = -1
	if ctx.IsSigned {
		uid = ctx.User.ID
	}
	return uid
}
