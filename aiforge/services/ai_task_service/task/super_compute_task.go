package task

import (
	"strconv"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type SuperComputeTaskTemplate struct {
	GrampusNoteBookTaskTemplate
}

func init() {
	t := &SuperComputeTaskTemplate{
		GrampusNoteBookTaskTemplate: GrampusNoteBookTaskTemplate{
			DefaultAITaskTemplate: DefaultAITaskTemplate{
				ClusterType: entity.C2Net,
				JobType:     models.JobTypeSuperCompute,
				Config:      GetGrampusNoteBookConfig,
			},
		},
	}
	RegisterTask(models.JobTypeSuperCompute, entity.C2Net, t)
}

func (g SuperComputeTaskTemplate) GetImages(computeSource models.ComputeSource, accCardType string, queues ...models.ResourceQueue) ([]entity.ClusterImage, bool, *response.BizError) {
	c := g.GetMyCluster()
	if c == nil {
		log.Error("Get cluster failed")
		return nil, false, response.SYSTEM_ERROR
	}
	l, f, err := c.GetNotebookImages(entity.GetImageReq{
		ComputeSource: computeSource,
		JobType:       models.JobTypeDebug,
		AccCardType:   accCardType,
	})
	if err != nil {
		log.Error("GetImages err.computeSource=%s err =%v", computeSource.Name, err)
		return nil, false, response.NewBizError(err)
	}
	return l, f, nil
}

func (t SuperComputeTaskTemplate) Create(ctx *context.CreationContext) (*entity.CreateTaskRes, *response.BizError) {
	c := &CreateOperator{}
	err := c.Next(t.CheckParamFormat).
		Next(t.CheckMultiRequest).
		Next(t.CheckDisplayJobName).
		Next(t.LoadSpec).
		Next(t.CheckPointBalance).
		Next(t.CheckDatasets).
		Next(t.CheckBranchExists).
		Next(t.CheckModels).
		Next(t.InsertCloudbrainRecord4Async).
		AsyncNextWithErrFun(t.BuildContainerData, t.CallCreationAPI, t.AfterCallCreationAPI4Async, t.NotifyCreation, t.HandleErr4Async).
		Operate(ctx)
	if err != nil {
		log.Error("create GrampusNoteBookTask err.%v", err)
		return nil, err
	}
	return &entity.CreateTaskRes{ID: ctx.NewCloudbrain.ID}, nil
}

func (SuperComputeTaskTemplate) NotifyCreation(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request

	var actionType = models.ActionCreateSuperComputeTask

	task, err := models.GetCloudbrainByCloudbrainID(ctx.NewCloudbrain.ID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err.Error())
		return response.NewBizError(err)
	}
	stringId := strconv.FormatInt(task.ID, 10)
	notification.NotifyOtherTask(ctx.User, nil, stringId, req.DisplayJobName, actionType)
	return nil
}

func (g SuperComputeTaskTemplate) GetDisplayJobName(userName string) string {
	t := time.Now()
	millisecondStr := strconv.FormatInt((t.UnixNano()%1e6/1e3)%1000, 10)
	return "mmlspark-" + t.Format("20060102150405") + millisecondStr
}
