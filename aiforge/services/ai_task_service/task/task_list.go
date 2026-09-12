package task

import (
	"io/ioutil"
	"net/http"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
)

type ModelExperienceList struct {
}

type ModelExperience struct {
	Name  string             `json:"name"`
	Descr string             `json:"descr"`
	Ops   ModelExperienceOps `json:"ops"`
}

type ModelExperienceOps struct {
	Repo_owner_name string `json:"repo_owner_name"`
	Repo_name       string `json:"repo_name"`
	Spec_id         string `json:"spec_id"`
	Model_id        string `json:"model_id"`
	Image_url       string `json:"image_url"`
	Boot_file       string `json:"boot_file"`
	Compute_source  string `json:"compute_source"`
}

func ReadFromPromote(url string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("not error.", err)
			return
		}
	}()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}
	return bytes, nil
}

func GetRepoAITaskList(req entity.GetTaskListReq) (*entity.AITaskListRes, *response.BizError) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	computeResource := ""
	if req.ComputeSource != nil {
		computeResource = req.ComputeSource.GetCloudbrainFormat()
	}
	var jobStatuses []string
	jobStatusNot := false
	if req.JobStatus == "other" {
		jobStatusNot = true
		jobStatuses = append(jobStatuses, req.ExcludeStatus...)
	} else if req.JobStatus != "all" && req.JobStatus != "" {
		jobStatuses = append(jobStatuses, req.JobStatus)
	}
	excludeJobTypes := []string{string(models.JobTypeSdFinetune), string(models.JobTypeFINETUNE), string(models.JobTypeEval), string(models.JobTypeModelExperience), string(models.JobTypeComfyuiExperience)} //需要固定排除掉微调任务和在线体验任务
	if len(req.ExcludeJobTypes) > 0 {
		excludeJobTypes = append(excludeJobTypes, req.ExcludeJobTypes...)
	}
	tasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
		ListOptions:     req.ListOptions,
		RepoID:          req.RepoID,
		ComputeResource: computeResource,
		Type:            models.TypeCloudBrainAll,
		IsLatestVersion: modelarts.IsLatestVersion,
		JobTypes:        req.JobTypes,
		AiCenter:        req.AICenter,
		JobStatus:       jobStatuses,
		JobStatusNot:    jobStatusNot,
		Cluster:         req.Cluster,
		ExcludeJobTypes: excludeJobTypes,
		IsAimRequired:   req.IsAimRequired,
	})
	if err != nil {
		log.Error("GetAITaskList query cloudbrains err. req=%+v,err=%v", req, err)
		return nil, response.NewBizError(err)
	}

	repo, _ := models.GetRepositoryByID(req.RepoID)
	models.LoadSpecs4CloudbrainInfo(tasks)
	r := make([]*entity.AITaskInfo4List, len(tasks))

	for i := 0; i < len(tasks); i++ {
		tasks[i].Repo = repo
		r[i] = &entity.AITaskInfo4List{
			Task:              entity.ConvertCloudbrainToAITaskBriefInfo(&tasks[i].Cloudbrain).ClearNonPublicFields(),
			Creator:           *entity.ConvertUserToBrief(&tasks[i].User),
			CanModify:         tasks[i].CanUserModify(req.Operator),
			CanDelete:         tasks[i].CanUserDelete(req.Operator, req.IsRepoOwner),
			CanCreateTemplate: tasks[i].CanCreateTemplate(req.Operator),
			CanExperience:     tasks[i].CanExperience(req.Operator),
		}
	}

	return &entity.AITaskListRes{
		Tasks:    r,
		Total:    count,
		PageSize: req.PageSize,
		Page:     page,
	}, nil
}

func GetMyAITaskList(req entity.GetMyTaskListReq) (*entity.AITaskListRes, *response.BizError) {
	var jobTypes []string
	jobTypeNot := false
	if req.JobType == string(models.JobTypeBenchmark) {
		jobTypes = models.AllBenchMarkJobType()
	} else if req.JobType != "all" && req.JobType != "" {
		jobTypes = append(jobTypes, req.JobType)
	}

	var jobStatuses []string
	jobStatusNot := false
	if req.JobStatus == "other" {
		jobStatusNot = true
		jobStatuses = append(jobStatuses, req.ExcludeStatus...)
	} else if req.JobStatus == "nostop" {
		jobStatuses = append(jobStatuses, cloudbrainTask.AlleNotFinalStatuses...)
	} else if req.JobStatus != "all" && req.JobStatus != "" {
		jobStatuses = append(jobStatuses, req.JobStatus)
	}

	computeSourceName := ""
	if req.ComputeSource != nil {
		computeSourceName = req.ComputeSource.GetCloudbrainFormat()
	}
	tasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Keyword:         req.Keyword,
		UserID:          req.User.ID,
		JobTypeNot:      jobTypeNot,
		JobStatusNot:    jobStatusNot,
		JobStatus:       jobStatuses,
		JobTypes:        jobTypes,
		NeedRepoInfo:    true,
		IsLatestVersion: modelarts.IsLatestVersion,
		ComputeResource: computeSourceName,
		Type:            models.TypeCloudBrainAll,
		AiCenter:        req.AICenter,
		Cluster:         req.Cluster,
		AppName:         req.AppName,
		OrderBy:         req.OrderBy,
	})
	if err != nil {
		log.Error("GetMyAITaskList query cloudbrains err. req=%+v,err=%v", req, err)
		return nil, response.NewBizError(err)
	}
	models.LoadSpecs4CloudbrainInfo(tasks)
	r := make([]*entity.AITaskInfo4List, len(tasks))

	for i := 0; i < len(tasks); i++ {
		var repoName, ownerName, repoAlias string
		if tasks[i].Repo != nil {
			repoName = tasks[i].Repo.Name
			ownerName = tasks[i].Repo.OwnerName
			repoAlias = tasks[i].Repo.Alias
		}
		r[i] = &entity.AITaskInfo4List{
			Task:                 entity.ConvertCloudbrainToAITaskBriefInfo(&tasks[i].Cloudbrain).ClearNonPublicFields(),
			CanModify:            tasks[i].CanUserModify(req.User),
			CanDelete:            tasks[i].CanUserDelete(req.User, req.IsRepoOwner),
			CanCreateTemplate:    tasks[i].CanCreateTemplate(req.User),
			CanExperience:        tasks[i].CanExperience(req.User),
			CanFintuneExperience: tasks[i].CanFintuneExperience(req.User),
			RepoID:               tasks[i].RepoID,
			RepoName:             repoName,
			RepoAlias:            repoAlias,
			OwnerName:            ownerName,
		}
	}

	return &entity.AITaskListRes{
		Tasks:    r,
		Total:    count,
		PageSize: req.PageSize,
		Page:     req.Page,
	}, nil
}

func GetAITaskList4Admin(req entity.GetMyTaskListReq) (*entity.AITaskListRes, *response.BizError) {
	var jobTypes []string
	jobTypeNot := false
	if req.JobType == string(models.JobTypeBenchmark) {
		jobTypes = models.AllBenchMarkJobType()
	} else if req.JobType != "all" && req.JobType != "" {
		jobTypes = append(jobTypes, req.JobType)
	}

	var jobStatuses []string
	jobStatusNot := false
	if req.JobStatus == "other" {
		jobStatusNot = true
		jobStatuses = append(jobStatuses, req.ExcludeStatus...)
	} else if req.JobStatus != "all" && req.JobStatus != "" {
		jobStatuses = append(jobStatuses, req.JobStatus)
	}

	computeSourceName := ""
	if req.ComputeSource != nil {
		computeSourceName = req.ComputeSource.GetCloudbrainFormat()
	}

	tasks, count, err := models.Cloudbrains(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Keyword:         req.Keyword,
		JobTypeNot:      jobTypeNot,
		JobStatusNot:    jobStatusNot,
		JobStatus:       jobStatuses,
		JobTypes:        jobTypes,
		NeedRepoInfo:    true,
		IsLatestVersion: modelarts.IsLatestVersion,
		ComputeResource: computeSourceName,
		Type:            models.TypeCloudBrainAll,
		AiCenter:        req.AICenter,
		Cluster:         req.Cluster,
		BeginTimeUnix:   req.BeginTimeUnix,
		EndTimeUnix:     req.EndTimeUnix,
	})
	if err != nil {
		log.Error("GetMyAITaskList query cloudbrains err. req=%+v,err=%v", req, err)
		return nil, response.NewBizError(err)
	}
	models.LoadSpecs4CloudbrainInfo(tasks)
	r := make([]*entity.AITaskInfo4List, len(tasks))

	for i := 0; i < len(tasks); i++ {
		var repoName, ownerName, repoAlias string

		if tasks[i].Repo != nil {
			repoName = tasks[i].Repo.Name
			ownerName = tasks[i].Repo.OwnerName
			repoAlias = tasks[i].Repo.Alias

		}
		creator := entity.UserBriefInfo{}
		if tasks[i].User.ID > 0 {
			creator = *entity.ConvertUserToBrief(&tasks[i].User)
		}
		r[i] = &entity.AITaskInfo4List{
			Task:          entity.ConvertCloudbrainToAITaskBriefInfo(&tasks[i].Cloudbrain),
			Creator:       creator,
			CanModify:     true,
			CanDelete:     true,
			CanExperience: true,
			RepoName:      repoName,
			RepoAlias:     repoAlias,
			OwnerName:     ownerName,
			RepoID:        tasks[i].RepoID,
		}
	}

	return &entity.AITaskListRes{
		Tasks:    r,
		Total:    count,
		PageSize: req.PageSize,
		Page:     req.Page,
	}, nil
}
