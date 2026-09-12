package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/structs"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
	"xorm.io/xorm"
)

var HF_MODEL_ORG_NAME = setting.ExternalTransfer.HfModelOwner

const (
	MODEL_LOCAL_TYPE  = 1
	MODEL_ONLINE_TYPE = 0
	MODEL_HF_TYPE     = 2
	MODEL_ALL_TYPE    = -1
)

type AiModelManage struct {
	// base
	ID          string `xorm:"pk" json:"id"`
	Name        string `xorm:"INDEX NOT NULL" json:"name"`
	LowerName   string `xorm:"INDEX"`
	Alias       string `xorm:"INDEX"`
	LowerAlias  string `xorm:"INDEX"`
	UserId      int64  `xorm:"NOT NULL" json:"userId"` //creator id
	OwnerID     int64  `xorm:"INDEX"`
	ModelType   int    `xorm:"NULL" json:"modelType"`
	License     string `xorm:"NULL" json:"license"`
	Engine      int64  `xorm:"NOT NULL DEFAULT 0" json:"engine"`
	IsPrivate   bool   `xorm:"DEFAULT true" json:"isPrivate"`
	StorageType string `xorm:"varchar(20)"`
	Path        string `xorm:"varchar(400) NOT NULL" json:"path"`

	ExternalName    string             `xorm:"text NULL" json:"externalName"`
	Size            int64              `xorm:"NOT NULL" json:"size"`
	Label           string             `xorm:"varchar(1000)" json:"label"`
	DownloadCount   int                `xorm:"NOT NULL DEFAULT 0" json:"downloadCount"`
	Status          int                `xorm:"NOT NULL DEFAULT 0" json:"status"`
	StatusDesc      string             `xorm:"varchar(500)" json:"statusDesc"`
	Recommend       int                `xorm:"NOT NULL DEFAULT 0" json:"recommend"`
	TrainTaskInfo   string             `xorm:"text NULL" json:"trainTaskInfo"`
	DerivativeCount int                `xorm:"NOT NULL DEFAULT 0" json:"derivativeCount"`
	ReferenceCount  int                `xorm:"NOT NULL DEFAULT 0" json:"referenceCount"`
	CollectedCount  int                `xorm:"NOT NULL DEFAULT 0" json:"collectedCount"`
	CreatedUnix     timeutil.TimeStamp `xorm:"created" json:"createdUnix"`
	UpdatedUnix     timeutil.TimeStamp `xorm:"updated" json:"updatedUnix"`

	// extra
	Owner             *User                    `xorm:"-"`
	UserName          string                   `xorm:"-" json:"userName"`
	CreatorName       string                   `xorm:"-" json:"creator_name"`
	UserRelAvatarLink string                   `xorm:"-" json:"userRelAvatarLink"`
	IsCanOper         bool                     `xorm:"-" json:"isCanOper"`
	IsCanDelete       bool                     `xorm:"-" json:"isCanDelete"`
	IsCanDownload     bool                     `xorm:"-" json:"isCanDownload"`
	IsCollected       bool                     `xorm:"-" json:"isCollected"`
	RepoName          string                   `xorm:"-" json:"repoName"`
	RepoDisplayName   string                   `xorm:"-" json:"repoDisplayName"`
	RepoOwnerName     string                   `xorm:"-" json:"repoOwnerName"`
	DatasetInfo       []*DatasetDownload       `xorm:"-" json:"datasetInfo"`
	ModelFileList     []storage.FileInfo       `xorm:"-" json:"modelFileList"`
	OnlineInfo        []map[string]interface{} `xorm:"-" json:"onlineInfo"`
	UsedCloudbrain    []map[string]interface{} `xorm:"-" json:"usedCloudbrain"`

	// unused: to be decided
	Description     string `xorm:"varchar(2000)" json:"description"`
	ComputeResource string `json:"computeResource"`
	Version         string `xorm:"NOT NULL" json:"version"`
	VersionCount    int    `xorm:"NOT NULL DEFAULT 0" json:"versionCount"`
	New             int    `xorm:"NOT NULL" json:"new"`
	Type            int    `xorm:"NOT NULL" json:"type"`
	Accuracy        string `xorm:"varchar(1000)" json:"accuracy"`
	AttachmentId    string `xorm:"NULL" json:"attachmentId"`
	RepoId          int64  `xorm:"INDEX NULL" json:"repoId"`
	CodeBranch      string `xorm:"varchar(400) NULL" json:"codeBranch"`
	CodeCommitID    string `xorm:"NULL" json:"codeCommitID"`
	HasOnlineUrl    int    `xorm:"NOT NULL DEFAULT 0" json:"hasOnlineUrl"`
}

type AiModelFile struct {
	ID            int64              `xorm:"pk autoincr"`
	ModelID       string             `xorm:"UNIQUE(s)"`
	Name          string             `xorm:"varchar(400) UNIQUE(s)"`
	Path          string             `xorm:"varchar(400) NULL"`
	Description   string             `xorm:"varchar(400) NULL"`
	DownloadCount int64              `xorm:"DEFAULT 0"`
	Size          int64              `xorm:"DEFAULT 0"`
	CreatedUnix   timeutil.TimeStamp `xorm:"created"`
}

type AiModelCollect struct {
	ID          int64              `xorm:"pk autoincr"`
	ModelID     string             `xorm:"UNIQUE(s)"`
	UserId      int64              `xorm:"UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

type AiModelConvert struct {
	ID                 string             `xorm:"pk" json:"id"`
	Name               string             `xorm:"INDEX NOT NULL" json:"name"`
	Status             string             `xorm:"NULL" json:"status"`
	StatusResult       string             `xorm:"NULL" json:"statusResult"`
	SrcEngine          int                `xorm:"NOT NULL DEFAULT 0" json:"srcEngine"`
	RepoId             int64              `xorm:"INDEX NULL" json:"repoId"`
	ModelId            string             `xorm:"NOT NULL" json:"modelId"`
	ModelName          string             `xorm:"NULL" json:"modelName"`
	ModelVersion       string             `xorm:"NOT NULL" json:"modelVersion"`
	ModelPath          string             `xorm:"NULL" json:"modelPath"`
	DestFormat         int                `xorm:"NOT NULL DEFAULT 0" json:"destFormat"`
	NetOutputFormat    int                `xorm:"NULL" json:"netOutputFormat"`
	UserId             int64              `xorm:"NOT NULL" json:"userId"`
	CloudBrainTaskId   string             `xorm:"NULL" json:"cloudBrainTaskId"`
	ModelArtsVersionId string             `xorm:"NULL" json:"modelArtsVersionId"`
	ContainerID        string             `json:"containerID"`
	ContainerIp        string             `json:"containerIp"`
	RunTime            int64              `xorm:"NULL" json:"runTime"`
	TrainJobDuration   string             `json:"trainJobDuration"`
	InputShape         string             `xorm:"varchar(2000)" json:"inputShape"`
	InputDataFormat    string             `xorm:"NOT NULL" json:"inputDataFormat"`
	Description        string             `xorm:"varchar(2000)" json:"description"`
	Path               string             `xorm:"varchar(400) NOT NULL" json:"path"`
	CreatedUnix        timeutil.TimeStamp `xorm:"created" json:"createdUnix"`
	UpdatedUnix        timeutil.TimeStamp `xorm:"updated" json:"updatedUnix"`
	StartTime          timeutil.TimeStamp `json:"startTime"`
	EndTime            timeutil.TimeStamp `json:"endTime"`
	UserName           string             `xorm:"-" json:"userName"`
	UserRelAvatarLink  string             `xorm:"-" json:"userRelAvatarLink"`
	IsCanOper          bool               `xorm:"-" json:"isCanOper"`
	IsCanDelete        bool               `xorm:"-" json:"isCanDelete"`
}

type AiModelQueryOptions struct {
	ListOptions
	RepoID   int64 // include all repos if empty
	UserID   int64
	ModelID  string
	SortType string
	New      int
	// JobStatus     CloudbrainStatus
	Type                  int
	Status                int
	IsOnlyThisRepo        bool
	IsQueryPrivate        bool
	IsRecommend           bool
	IsCollected           bool
	CollectedUserId       int64
	Namelike              string
	LabelFilter           string
	FrameFilter           int
	ComputeResourceFilter string
	NotNeedEmpty          bool
	HasOnlineUrl          int
	MyExternal            bool
}

func (a *AiModelConvert) IsGpuTrainTask() bool {
	if a.SrcEngine == 0 || a.SrcEngine == 1 || a.SrcEngine == 4 || a.SrcEngine == 6 {
		return true
	}
	return false
}

func ModelComputeAndSetDuration(task *AiModelConvert, result JobResultPayload) {
	if task.StartTime == 0 {
		task.StartTime = timeutil.TimeStamp(result.JobStatus.CreatedTime / 1000)
	}
	if task.EndTime == 0 {
		if result.JobStatus.CompletedTime > 0 {
			task.EndTime = timeutil.TimeStamp(result.JobStatus.CompletedTime / 1000)
		}
	}
	var d int64
	if task.StartTime == 0 {
		d = 0
	} else if task.EndTime == 0 {
		d = time.Now().Unix() - task.StartTime.AsTime().Unix()
	} else {
		d = task.EndTime.AsTime().Unix() - task.StartTime.AsTime().Unix()
	}

	if d < 0 {
		d = 0
	}
	task.RunTime = d
	task.TrainJobDuration = ConvertDurationToStr(d)
}

func ModelConvertSetDuration(task *AiModelConvert) {
	var d int64
	if task.StartTime == 0 {
		d = 0
	} else if task.EndTime == 0 {
		d = time.Now().Unix() - task.StartTime.AsTime().Unix()
	} else {
		d = task.EndTime.AsTime().Unix() - task.StartTime.AsTime().Unix()
	}

	if d < 0 {
		d = 0
	}
	task.RunTime = d
	task.TrainJobDuration = ConvertDurationToStr(d)
}

func UpdateModelConvertModelArts(id string, CloudBrainTaskId string, VersionId string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("cloud_brain_task_id,model_arts_version_id").Update(&AiModelConvert{
		CloudBrainTaskId:   CloudBrainTaskId,
		ModelArtsVersionId: VersionId,
	})
	if err != nil {
		return err
	}
	log.Info("success to update cloud_brain_task_id from  db.re=" + fmt.Sprint((re)))
	return nil
}

func UpdateModelConvertFailed(id string, status string, statusResult string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("status", "status_result").Update(&AiModelConvert{
		Status:       status,
		StatusResult: statusResult,
	})
	if err != nil {
		return err
	}
	log.Info("success to update cloud_brain_task_id from  db.re=" + fmt.Sprint((re)))
	return nil
}

func UpdateModelConvertCBTI(id string, CloudBrainTaskId string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("cloud_brain_task_id").Update(&AiModelConvert{
		CloudBrainTaskId: CloudBrainTaskId,
	})
	if err != nil {
		return err
	}
	log.Info("success to update cloud_brain_task_id from  db.re=" + fmt.Sprint((re)))
	return nil
}

func UpdateResultMigrateFlag(id string, resultMigrateFlag string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("status_result").Update(&AiModelConvert{
		StatusResult: resultMigrateFlag,
	})
	if err != nil {
		return err
	}
	log.Info("success to update resultMigrateFlag from  db.re=" + fmt.Sprint((re)))
	return nil
}

func UpdateModelConvert(job *AiModelConvert) error {
	return updateModelConvert(x, job)
}

func updateModelConvert(e Engine, job *AiModelConvert) error {
	var sess *xorm.Session
	sess = e.Where("id = ?", job.ID)
	_, err := sess.Cols("status", "train_job_duration", "run_time", "start_time", "end_time", "updated_unix").Update(job)
	return err
}

func SaveModelConvert(modelConvert *AiModelConvert) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Insert(modelConvert)
	if err != nil {
		log.Info("insert modelConvert error." + err.Error())
		return err
	}
	log.Info("success to save modelConvert db.re=" + fmt.Sprint((re)))
	return nil
}

func SaveModelToDb(model *AiModelManage) error {
	sess := x.NewSession()
	defer sess.Close()

	re, err := sess.Insert(model)
	if err != nil {
		log.Info("insert error." + err.Error())
		return err
	}
	log.Info("success to save db.re=" + fmt.Sprint((re)))
	return nil
}

func QueryModelConvertByName(name string, repoId int64) ([]*AiModelConvert, error) {
	sess := x.NewSession()
	defer sess.Close()
	sess.Select("*").Table(new(AiModelConvert)).
		Where("name='" + name + "' and repo_id=" + fmt.Sprint(repoId)).OrderBy("created_unix desc")
	aiModelManageConvertList := make([]*AiModelConvert, 0)
	err := sess.Find(&aiModelManageConvertList)
	if err == nil {
		return aiModelManageConvertList, nil
	}
	return nil, err
}

func QueryModelConvertById(id string) (*AiModelConvert, error) {
	sess := x.NewSession()
	defer sess.Close()
	re := new(AiModelConvert)
	isExist, err := sess.Table(new(AiModelConvert)).ID(id).Get(re)
	if err == nil && isExist {
		return re, nil
	}
	return nil, err
}

func QueryHfModelByStatus(status int) ([]*AiModelManage, error) {
	sess := x.NewSession()
	defer sess.Close()
	var models []*AiModelManage
	err := sess.Where("model_type=2 AND status=" + fmt.Sprint(status)).Find(&models)
	if err != nil {
		return nil, err
	}
	return models, nil

}

func QueryModelById(id string) (*AiModelManage, error) {
	sess := x.NewSession()
	defer sess.Close()
	re := new(AiModelManage)
	isExist, err := sess.Table(new(AiModelManage)).ID(id).Get(re)
	if err != nil {
		return nil, err
	} else if !isExist {
		return nil, ErrPretrainModelNotExist{}
	}
	return re, nil
}

func DeleteModelConvertById(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Delete(&AiModelConvert{
		ID: id,
	})
	if err != nil {
		return err
	}
	log.Info("success to delete AiModelManageConvert from  db.re=" + fmt.Sprint((re)))
	return nil
}

func DeleteModelById(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Delete(&AiModelManage{
		ID: id,
	})
	if err != nil {
		return err
	}
	log.Info("success to delete from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelDescription(id string, description string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("description").Update(&AiModelManage{
		Description: description,
	})
	if err != nil {
		return err
	}
	log.Info("success to update description from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelHashOnlineUrl(id string, hasOnlineUrl int) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("has_online_url").Update(&AiModelManage{
		HasOnlineUrl: hasOnlineUrl,
	})
	if err != nil {
		return err
	}
	log.Info("success to update hasOnlineUrl from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelPrivate(id string, isPrivate bool) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("is_private").Update(&AiModelManage{
		IsPrivate: isPrivate,
	})
	if err != nil {
		return err
	}
	log.Info("success to update isPrivate from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelRecommend(id string, recommend int) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("recommend").Update(&AiModelManage{
		Recommend: recommend,
	})
	if err != nil {
		return err
	}
	log.Info("success to update recommend from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelCollectedNum(id string, collectedNum int) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("collected_count").Update(&AiModelManage{
		CollectedCount: collectedNum,
	})
	if err != nil {
		return err
	}
	log.Info("success to update collectedNum from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyLocalModel(id string, name, label, description string, engine int, isPrivate bool, license string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("name", "label", "description", "engine", "is_private", "license").Update(&AiModelManage{
		Description: description,
		Name:        name,
		Label:       label,
		Engine:      int64(engine),
		IsPrivate:   isPrivate,
		License:     license,
	})
	if err != nil {
		return err
	}
	log.Info("success to update description from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelSize(id string, size int64) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("size").Update(&AiModelManage{
		Size: size,
	})
	if err != nil {
		return err
	}
	log.Info("success to update size from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelStatus(id string, modelSize int64, status int, modelPath string, statusDesc string) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("size", "status", "path", "status_desc").Update(&AiModelManage{
		Size:       modelSize,
		Status:     status,
		Path:       modelPath,
		StatusDesc: statusDesc,
	})
	if err != nil {
		return err
	}
	log.Info("success to update ModelStatus from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelNewProperty(id string, new int, versioncount int) error {
	var sess *xorm.Session
	sess = x.ID(id)
	defer sess.Close()
	re, err := sess.Cols("new", "version_count").Update(&AiModelManage{
		New:          new,
		VersionCount: versioncount,
	})
	if err != nil {
		return err
	}
	log.Info("success to update new property from  db.re=" + fmt.Sprint((re)))
	return nil
}

func ModifyModelDerivativeCount(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("UPDATE `ai_model_manage` SET derivative_count = derivative_count + 1 WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}

func ModifyModelDownloadCount(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("UPDATE `ai_model_manage` SET download_count = download_count + 1 WHERE id = ?", id); err != nil {
		return err
	}

	return nil
}

func QueryModelByName(name string, repoId int64) []*AiModelManage {
	sess := x.NewSession()
	defer sess.Close()
	sess.Select("*").Table("ai_model_manage").
		Where("name='" + name + "' and repo_id=" + fmt.Sprint(repoId)).OrderBy("created_unix desc")
	aiModelManageList := make([]*AiModelManage, 0)
	sess.Find(&aiModelManageList)
	return aiModelManageList
}

func QueryModelByRepoId(repoId int64) []*AiModelManage {
	sess := x.NewSession()
	defer sess.Close()
	sess.Select("*").Table("ai_model_manage").
		Where("repo_id=?", repoId)
	aiModelManageList := make([]*AiModelManage, 0)
	sess.Find(&aiModelManageList)
	return aiModelManageList
}

func DeleteModelByRepoId(repoId int64) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Delete(&AiModelManage{
		RepoId: repoId,
	})
	if err != nil {
		return err
	}
	log.Info("success to delete DeleteModelByRepoId from  db.re=" + fmt.Sprint((re)))
	return nil
}

func QueryModelByPath(path string) (*AiModelManage, error) {
	modelManage := new(AiModelManage)
	has, err := x.Where("path=?", path).Get(modelManage)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrNotExist{}
	}
	return modelManage, nil
}

func QueryModel(opts *AiModelQueryOptions) ([]*AiModelManage, int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	if opts.Namelike != "" {
		opts.Namelike = strings.ReplaceAll(opts.Namelike, "'", "")

	}
	var where string
	where += " ai_model_manage.user_id > 0 "
	if opts.RepoID > 0 {
		where += " and  ai_model_manage.repo_id= " + fmt.Sprint(opts.RepoID)
	}
	var hfTable string
	if opts.MyExternal {
		hfTable = `(SELECT DISTINCT model_id, user_id AS hf_user_id FROM hf_model_file
				UNION
				SELECT DISTINCT model_id, user_id AS hf_user_id FROM hf_model_operation) AS hf_table`
	}

	if opts.UserID > 0 {
		if opts.MyExternal {
			sess.Join("INNER", hfTable, "ai_model_manage.id = hf_table.model_id")
			where += " and hf_table.hf_user_id=" + fmt.Sprint(opts.UserID)
		} else {
			where += " and ai_model_manage.user_id=" + fmt.Sprint(opts.UserID)
		}
	}

	if opts.New >= 0 {
		where += " and ai_model_manage.new=" + fmt.Sprint(opts.New)
	}

	if len(opts.ModelID) > 0 {
		where += " and ai_model_manage.id='" + fmt.Sprint(opts.ModelID) + "'"
	}

	if (opts.Type) >= 0 {
		where += " and ai_model_manage.type=" + fmt.Sprint(opts.Type)
	}

	if (opts.Status) >= 0 {
		where += " and ai_model_manage.status=" + fmt.Sprint(opts.Status)
	}
	if !opts.IsQueryPrivate {
		where += " and ai_model_manage.is_private=false"
	}
	if opts.IsRecommend {
		where += " and ai_model_manage.recommend=1"
	}
	if opts.FrameFilter >= 0 {
		if opts.FrameFilter == 2 {
			where += " and ai_model_manage.engine in (2,121,122)"
		} else {
			where += " and ai_model_manage.engine=" + fmt.Sprint(opts.FrameFilter)
		}
	}
	if opts.LabelFilter != "" {
		where += " and ai_model_manage.label ILIKE '%" + opts.LabelFilter + "%'"
	}
	if opts.ComputeResourceFilter != "" {
		where += " and ai_model_manage.compute_resource ILIKE '%" + opts.ComputeResourceFilter + "%'"
	}
	if opts.Namelike != "" {
		where += " and ( ai_model_manage.name ILIKE '%" + opts.Namelike + "%'"
		where += " or ai_model_manage.description ILIKE '%" + opts.Namelike + "%'"
		where += " or ai_model_manage.label ILIKE '%" + opts.Namelike + "%')"
	}
	if opts.NotNeedEmpty {
		where += " and ai_model_manage.size > 0 "
	}
	if opts.HasOnlineUrl > 0 {
		where += " and ai_model_manage.has_online_url =1 "
	}
	var count int64
	var err error
	if opts.IsCollected {
		where += " and ai_model_collect.user_id=" + fmt.Sprint(opts.CollectedUserId)

		count, err = sess.Join("INNER", "ai_model_collect", "ai_model_manage.id = ai_model_collect.model_id").Where(where).Count(new(AiModelManage))
		if err != nil {
			log.Info("error=" + err.Error())
			return nil, 0, fmt.Errorf("Count: %v", err)
		}
	} else {
		count, err = sess.Where(where).Count(new(AiModelManage))
		if err != nil {
			log.Info("error=" + err.Error())
			return nil, 0, fmt.Errorf("Count: %v", err)
		}
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}
	if opts.IsCollected {
		sess.Join("INNER", "ai_model_collect", "ai_model_manage.id = ai_model_collect.model_id")
	}
	if opts.MyExternal {
		sess.Join("INNER", hfTable, "ai_model_manage.id = hf_table.model_id")
	}
	orderby := "ai_model_manage.created_unix desc"
	if opts.SortType != "" {
		orderby = opts.SortType
	}
	sess.OrderBy(orderby)
	aiModelManages := make([]*AiModelManage, 0, setting.UI.IssuePagingNum)
	if err := sess.Table("ai_model_manage").Where(where).
		Find(&aiModelManages); err != nil {
		log.Info("error=" + err.Error())
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	return aiModelManages, count, nil
}

func QueryModelConvertCountByRepoID(repoId int64) int64 {
	convert := new(AiModelConvert)
	total, _ := x.Where("repo_id =?", repoId).Count(convert)
	return total
}

func QueryModelConvertByRepoID(repoId int64) ([]*AiModelConvert, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"ai_model_convert.repo_id": repoId},
	)
	sess.OrderBy("ai_model_convert.created_unix DESC")
	aiModelManageConvert := make([]*AiModelConvert, 0)
	if err := sess.Table(new(AiModelConvert)).Where(cond).
		Find(&aiModelManageConvert); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}
	return aiModelManageConvert, nil
}

func QueryModelConvertByUserID(userID int64) ([]*AiModelConvert, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"ai_model_convert.user_id": userID},
	)
	sess.OrderBy("ai_model_convert.created_unix DESC")
	aiModelManageConvert := make([]*AiModelConvert, 0)
	if err := sess.Table(new(AiModelConvert)).Where(cond).
		Find(&aiModelManageConvert); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}
	return aiModelManageConvert, nil
}

func QueryModelConvert(opts *AiModelQueryOptions) ([]*AiModelConvert, int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	if opts.RepoID > 0 {
		cond = cond.And(
			builder.Eq{"ai_model_convert.repo_id": opts.RepoID},
		)
	}
	if opts.UserID > 0 {
		cond = cond.And(
			builder.Eq{"ai_model_convert.user_id": opts.UserID},
		)
	}
	count, err := sess.Where(cond).Count(new(AiModelConvert))
	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}
	sess.OrderBy("ai_model_convert.created_unix DESC")
	aiModelManageConvert := make([]*AiModelConvert, 0, setting.UI.IssuePagingNum)
	if err := sess.Table(new(AiModelConvert)).Where(cond).
		Find(&aiModelManageConvert); err != nil {
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	return aiModelManageConvert, count, nil
}

func SaveModelCollect(modelCollect *AiModelCollect) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Insert(modelCollect)
	if err != nil {
		log.Info("insert AiModelCollect error." + err.Error())
		return err
	}
	log.Info("success to save AiModelCollect db.re=" + fmt.Sprint((re)))
	return nil
}

func DeleteModelCollect(modelCollect *AiModelCollect) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Delete(modelCollect)
	if err != nil {
		log.Info("delete AiModelCollect error." + err.Error())
		return err
	}
	log.Info("success to delete AiModelCollect db.re=" + fmt.Sprint((re)))
	return nil
}

func QueryModelCollectNum(modelId string) int {
	sess := x.NewSession()
	defer sess.Close()
	modelCollects := make([]*AiModelCollect, 0)
	err := sess.Table(new(AiModelCollect)).Where("model_id=?", modelId).Find(&modelCollects)
	if err == nil {
		return len(modelCollects)
	}
	return 0
}
func QueryModelCollectByUserId(modelId string, userId int64) []*AiModelCollect {
	sess := x.NewSession()
	defer sess.Close()
	modelCollects := make([]*AiModelCollect, 0)
	err := sess.Table(new(AiModelCollect)).Where("model_id=? and user_id=?", modelId, userId).Find(&modelCollects)
	if err == nil {
		return modelCollects
	}
	return nil
}

func QueryModelCollectedStatus(modelIds []string, userId int64) map[string]*AiModelCollect {
	sess := x.NewSession()
	defer sess.Close()
	modelCollects := make([]*AiModelCollect, 0)
	var cond = builder.NewCond()
	cond = cond.And(
		builder.In("model_id", modelIds),
	)
	cond = cond.And(
		builder.Eq{"user_id": userId},
	)
	result := make(map[string]*AiModelCollect, 0)
	err := sess.Table(new(AiModelCollect)).Where(cond).Find(&modelCollects)
	if err == nil {
		for _, v := range modelCollects {
			result[v.ModelID] = v
		}
	}
	return result
}

func SaveModelFile(modelFile *AiModelFile) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Insert(modelFile)
	if err != nil {
		log.Info("insert modelFile error." + err.Error())
		return err
	}
	log.Info("success to save modelFile db.re=" + fmt.Sprint((re)))
	return nil
}

func DeleteModelFile(modelFile *AiModelFile) error {
	sess := x.NewSession()
	defer sess.Close()
	re, err := sess.Delete(modelFile)
	if err != nil {
		log.Info("delete modelFile error." + err.Error())
		return err
	}
	log.Info("success to delete modelFile db.re=" + fmt.Sprint((re)))
	return nil
}

func QueryModelFileByModelId(modelId string) []*AiModelFile {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"model_id": modelId},
	)
	result := make([]*AiModelFile, 0)
	err := sess.Table(new(AiModelFile)).Where(cond).Find(&result)
	if err != nil {
		log.Info("query AiModelFile failed, err=" + err.Error())
	}
	return result
}

func QueryModelForSearch(opts *AiModelQueryOptions, isAdmin bool) ([]*AiModelManage, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	query := builder.NewCond()
	if opts.Namelike != "" {
		query = query.And(builder.Or(builder.Like{"lower_alias", "%" + strings.ToLower(opts.Namelike) + "%"}, builder.Like{"lower_name", "%" + strings.ToLower(opts.Namelike) + "%"}))
	}
	if !isAdmin {
		query = query.And(
			builder.Or(
				builder.Eq{"owner_id": opts.UserID},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id::text").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.UserID}.
							And(builder.Eq{"`subject_access`.subject_type": AimodelSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeOwner), int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin))))),
					builder.In("id", builder.Select("`team_subject`.subject_id::text").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": opts.UserID}.
							And(builder.Neq{"`team`.aimodel_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				),
			),
		)
	}
	query = query.And(builder.Eq{"is_private": true})

	var count int64
	var err error
	count, err = sess.Where(query).Count(new(AiModelManage))
	if err != nil {
		log.Info("error=" + err.Error())
		return nil, 0, fmt.Errorf("Count: %v", err)
	}
	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		sess.Limit(opts.PageSize, start)
	}

	orderby := "ai_model_manage.created_unix desc"
	if opts.SortType != "" {
		orderby = opts.SortType
	}
	sess.OrderBy(orderby)
	aiModelManages := make([]*AiModelManage, 0, setting.UI.IssuePagingNum)
	if err := sess.Table("ai_model_manage").Where(query).
		Find(&aiModelManages); err != nil {
		log.Info("error=" + err.Error())
		return nil, 0, fmt.Errorf("Find: %v", err)
	}

	return aiModelManages, count, nil
}

func QueryModelRepoByModelID(modelId string) (*Repository, error) {
	r := &Repository{}
	has, err := x.Where(builder.NewCond().
		And(builder.Eq{"id": builder.Select("repo_id").
			From("ai_model_manage").
			Where(builder.Eq{"id": modelId})})).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, &ErrRecordNotExist{}
	}
	return r, nil
}

func QueryModelMapsByIds(ids []string) (map[string]*AiModelManage, error) {
	sess := x.NewSession()
	defer sess.Close()
	re := make([]*AiModelManage, 0)
	err := sess.Table(new(AiModelManage)).In("id", ids).Find(&re)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]*AiModelManage, 0)
	for _, m := range re {
		resultMap[m.ID] = m
	}
	return resultMap, nil
}

func QueryModelByIds(ids []string) (AimodelList, error) {
	sess := x.NewSession()
	defer sess.Close()
	aimodels := make(AimodelList, 0)
	err := sess.Table(new(AiModelManage)).In("id", ids).Find(&aimodels)
	if err != nil {
		return nil, err
	}
	aimodels.loadAttributes(0)
	return aimodels, nil
}

// created_unix
func QueryModelIdsByPaging(pageSize, pageNum int, sort string) ([]string, error) {
	sess := x.NewSession()
	defer sess.Close()
	re := make([]string, 0)
	start := (pageNum - 1) * pageSize
	err := sess.Table("ai_model_manage").Cols("id").OrderBy(sort).Limit(pageSize, start).Find(&re)
	return re, err
}

func QueryModelEvalueAnalizeMap() map[string][]*AiModelManage {
	sess := x.NewSession()
	defer sess.Close()
	result := make(map[string][]*AiModelManage, 0)
	cond := "train_task_info <> ''"
	var indexTotal int64
	count, err := sess.Where(cond).Count(new(AiModelManage))
	if err == nil {
		for {
			sess.Select("*").Table(new(AiModelManage)).Where(cond).OrderBy("created_unix desc").Limit(PAGE_SIZE, int(indexTotal))
			aiModelManageList := make([]*AiModelManage, 0)
			err1 := sess.Find(&aiModelManageList)
			if err1 == nil {
				for _, model := range aiModelManageList {
					if model.TrainTaskInfo != "" {
						var task Cloudbrain
						err := json.Unmarshal([]byte(model.TrainTaskInfo), &task)
						if err != nil {
							log.Info("error=" + err.Error())
						} else {
							if task.ModelId != "" {
								tmps := strings.Split(task.ModelId, ";")
								for _, v := range tmps {
									if tmpList, ok := result[v]; !ok {
										list := make([]*AiModelManage, 0)
										list = append(list, model)
										result[v] = list
									} else {
										tmpList = append(tmpList, model)
										result[v] = tmpList
									}
								}
							}
						}
					}
				}
			} else {
				log.Info("query aimodelmanage error:" + err1.Error())
				break
			}

			indexTotal += PAGE_SIZE
			if indexTotal >= count {
				break
			}
		}
	}
	return result
}

func QueryPublicModelByRepoIdS(repoIds []int64) []*AiModelManage {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.In("repo_id", repoIds),
	)
	cond = cond.And(
		builder.Eq{"is_private": false},
	)
	sess.Select("*").Table("ai_model_manage").OrderBy("ai_model_manage.updated_unix desc").
		Where(cond)
	aiModelManageList := make([]*AiModelManage, 0)
	sess.Find(&aiModelManageList)
	return aiModelManageList
}

func QueryPublicModelAndOwnByRepoIdS(orgId int64, uid int64, keyword, labelFilter string, orderBy SearchOrderBy, page int, pageSize int, isAdmin bool) ([]*AiModelManage, int64) {
	sess := x.NewSession()
	defer sess.Close()
	where := "1=1"

	if labelFilter != "" {
		where += " and label ILIKE '%" + labelFilter + "%'"
	}
	orderByStr := "ai_model_manage." + string(orderBy)
	if keyword != "" {
		where += " and (ai_model_manage.name ILIKE '%" + keyword + "%'"
		where += " or ai_model_manage.description ILIKE '%" + keyword + "%'"
		where += " or ai_model_manage.label ILIKE '%" + keyword + "%')"
		orderByStr = "CASE WHEN ai_model_manage.name='" + keyword + "' THEN 0 ELSE 1 END," + orderByStr
	}
	start := (page - 1) * pageSize
	aiModelManageList := make([]*AiModelManage, 0)
	var count int64
	if isAdmin {
		innnerCondition := "ai_model_manage.repo_id = repository.id and repository.owner_id=? "
		count1, err := x.Table("ai_model_manage").Join("INNER", "repository", innnerCondition, orgId).Where(where).Count()
		if err != nil {
			log.Info("err:" + err.Error())
		}
		err = x.Table("ai_model_manage").Join("INNER", "repository", innnerCondition, orgId).Where(where).OrderBy(orderByStr).Limit(pageSize, start).Find(&aiModelManageList)
		if err != nil {
			log.Info("err:" + err.Error())
		}
		count = count1
	} else {
		innnerCondition := "ai_model_manage.repo_id = repository.id and ((repository.is_private=false and repository.owner_id=? and ai_model_manage.is_private=false) or (ai_model_manage.user_id=? and repository.owner_id=?))"
		count1, err := x.Table("ai_model_manage").Join("INNER", "repository", innnerCondition, orgId, uid, orgId).Where(where).Count()
		if err != nil {
			log.Info("err:" + err.Error())
		}
		err = x.Table("ai_model_manage").Join("INNER", "repository", innnerCondition, orgId, uid, orgId).Where(where).OrderBy(orderByStr).Limit(pageSize, start).Find(&aiModelManageList)
		if err != nil {
			log.Info("err:" + err.Error())
		}
		count = count1
	}

	return aiModelManageList, count
}

func QueryPublicModelLabelAndOwnByRepoIdS(repoIds []int64, uid int64) []*AiModelManage {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(
		builder.In("repo_id", repoIds),
	)

	var or2Cond = builder.NewCond()
	or2Cond = or2Cond.And(
		builder.Eq{"is_private": true},
	)
	or2Cond = or2Cond.And(
		builder.Eq{"user_id": uid},
	)
	var orCond = builder.NewCond()
	orCond = orCond.Or(builder.Eq{"is_private": false}, or2Cond)

	cond = cond.And(orCond)

	aiModelManageList := make([]*AiModelManage, 0)

	sess.Select("id,label").Table("ai_model_manage").Where(cond).Find(&aiModelManageList)
	return aiModelManageList
}

func GetUsedModelSizeByUser(userId int64) (int64, error) {

	total, err := x.Where("user_id = ? and model_type!= ?", userId, MODEL_HF_TYPE).Sum(new(AiModelManage), "size")

	return int64(total), err

}

// new
func GetUsedModelSizeByOwner(userId int64) (int64, error) {
	total, err := x.Where("owner_id = ? and model_type!= ?", userId, MODEL_HF_TYPE).Sum(new(AiModelManage), "size")
	return int64(total), err
}

func GetModelByOwnerNameAndAimodelName(ownerName, aimodelName string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.Join("INNER", "public.user", "public.user.id = ai_model_manage.owner_id").Where("ai_model_manage.lower_name = ? and public.user.lower_name = ?", strings.ToLower(aimodelName), strings.ToLower(ownerName)).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil
}

func GetAimodelByID(id string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.ID(id).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil
}

func DeleteAimodel(ctx DBContext, doer *User, uid int64, aimodelId string) error {
	// In case is a organization.
	org, err := GetUserByID(uid)
	if err != nil {
		return err
	}
	if org.IsOrganization() {
		if err = org.GetTeams(&SearchTeamOptions{}); err != nil {
			return err
		}
	}

	sess := ctx.e
	aimodel := &AiModelManage{ID: aimodelId, OwnerID: uid}
	has, err := sess.Get(aimodel)
	if err != nil {
		return err
	} else if !has {
		return ErrRecordNotExist{}
	}

	if cnt, err := sess.ID(aimodelId).Delete(&AiModelManage{}); err != nil {
		return err
	} else if cnt != 1 {
		return ErrRecordNotExist{}
	}

	if org.IsOrganization() {
		for _, t := range org.Teams {
			if !t.hasSubject(sess, aimodelId, AimodelSubject) {
				continue
			} else if err = t.removeAimodel(sess, aimodel, false); err != nil {
				return err
			}
		}
	}

	if err = deleteBeans(sess,
		&SubjectAccess{SubjectID: aimodelId, SubjectType: int(AimodelSubject)},
		&Action{AimodelID: &aimodelId},
		&AiModelCollect{ModelID: aimodelId},
		&SubjectCollaboration{SubjectID: aimodelId, SubjectType: int(AimodelSubject)},
		&OfficialModel{ModelID: aimodelId},
	); err != nil {
		return fmt.Errorf("deleteBeans: %v", err)
	}

	return nil
}

func (aimodel *AiModelManage) ConvertSubjectAccessContext() *SubjectAccessContext {
	return &SubjectAccessContext{
		SubjectType: AimodelSubject,
		SubjectID:   aimodel.ID,
		OwnerID:     aimodel.OwnerID,
		IsPrivate:   aimodel.IsPrivate,
		Aimodel:     aimodel,
	}
}

func (aimodel *AiModelManage) GetOwner() (err error) {
	return aimodel.getOwner(x)
}

func (aimodel *AiModelManage) getOwner(e Engine) (err error) {
	if aimodel.Owner != nil {
		return nil
	}

	aimodel.Owner, err = getUserByID(e, aimodel.OwnerID)
	return err
}

func (aimodel *AiModelManage) DisplayName() string {
	if aimodel.Alias != "" {
		return aimodel.Alias
	}
	return aimodel.Name
}

func GetAiModelByByOwnerAndName(ownerId int64, name string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.Where("owner_id = ? and lower_name = ?", ownerId, strings.ToLower(name)).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil
}

func GetAiModelByOwnerNameAndName(ownerName string, name string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.Join("INNER", "public.user", "public.user.id = dataset_registry.owner_id").Where("ai_model_manage.lower_name = ? and public.user.lower_name = ?", strings.ToLower(name), strings.ToLower(ownerName)).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil

}

func GetAiModelByOwnerAndAlias(ownerId int64, alias string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.Where("owner_id = ? and lower_alias = ?", ownerId, strings.ToLower(alias)).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil
}

func isUserAimodelAdmin(e Engine, aimodel *AiModelManage, user *User) (bool, error) {
	if user == nil || aimodel == nil {
		return false, nil
	}
	if user.IsAdmin {
		return true, nil
	}

	mode, err := subjectAccessLevel(e, user, aimodel.ConvertSubjectAccessContext())
	if err != nil {
		return false, err
	}
	if mode >= AccessModeAdmin {
		return true, nil
	}

	teams, err := getUserSubjectTeams(e, aimodel.OwnerID, user.ID, aimodel.ID, AimodelSubject)
	if err != nil {
		return false, err
	}

	for _, team := range teams {
		if team.Authorize >= AccessModeAdmin {
			return true, nil
		}
	}
	return false, nil
}

func CreateAimodel(aimodel *AiModelManage, doer *User) error {
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	_, err = sess.Insert(aimodel)
	if err != nil {
		return err
	}
	if _, err = sess.Incr("num_aimodels").ID(aimodel.OwnerID).Update(new(User)); err != nil {
		return fmt.Errorf("increment user num_aimodels: %v", err)
	}
	if err = aimodel.getOwner(sess); err != nil {
		return fmt.Errorf("aimodel getOwner: %v", err)
	}
	subjectCtx := aimodel.ConvertSubjectAccessContext()
	owner := aimodel.Owner
	if owner.IsOrganization() {
		if err = owner.GetTeams(&SearchTeamOptions{}); err != nil {
			return fmt.Errorf("GetTeams: %v", err)
		}
		for _, t := range owner.Teams {
			if t.IncludesAllAimodels {
				if err = t.addAimodel(sess, aimodel); err != nil {
					return fmt.Errorf("addAimodel: %v", err)
				}
			}
		}

		if isAdmin, err := isUserAimodelAdmin(sess, aimodel, doer); err != nil {
			return fmt.Errorf("isUserDatasetAdmin: %v", err)
		} else if !isAdmin {
			if err = subjectCtx.addCollaborator(sess, doer); err != nil {
				return fmt.Errorf("AddCollaborator: %v", err)
			}
			if err = subjectCtx.changeCollaborationAccessMode(sess, doer.ID, AccessModeAdmin); err != nil {
				return fmt.Errorf("ChangeCollaborationAccessMode: %v", err)
			}
		}
	} else if err = subjectCtx.recalculateAccesses(sess); err != nil {
		return fmt.Errorf("recalculateAccesses: %v", err)
	}
	sess.Commit()
	return nil
}

type AimodelSearchOrder string

const (
	AimodelSearchOrderByDefault       AimodelSearchOrder = "recommend DESC,collected_count DESC,updated_unix DESC,name ASC"
	AimodelSearchOrderByNewest        AimodelSearchOrder = "created_unix DESC,name ASC"
	AimodelSearchOrderByRecentUpdated AimodelSearchOrder = "updated_unix DESC,name ASC"
	AimodelSearchOrderByDownload      AimodelSearchOrder = "download_count DESC,name ASC"
	AimodelSearchOrderByCollected     AimodelSearchOrder = "collected_count DESC,name ASC"
	AimodelSearchOrderByReference     AimodelSearchOrder = "reference_count DESC,name ASC"
	AimodelSearchOrferByDerivative    AimodelSearchOrder = "derivative_count DESC,name ASC"
	AimodelSearchOrderBySizeAsc       AimodelSearchOrder = "size ASC,name ASC"
	AimodelSearchOrderBySizeDesc      AimodelSearchOrder = "size DESC,name ASC"
	AimodelSearchOrderByAliasAsc      AimodelSearchOrder = `alias COLLATE "zh-x-icu" ASC, id ASC`
	AimodelSearchOrderByAliasDesc     AimodelSearchOrder = `alias COLLATE "zh-x-icu" DESC, id ASC`
)

func (d AimodelSearchOrder) String() string {
	return string(d)
}

func GetAimodelSearchOrderBy(orderType string) AimodelSearchOrder {
	var orderBy AimodelSearchOrder
	switch orderType {
	case OrderByNewest:
		orderBy = AimodelSearchOrderByNewest
	case OrderByRecentUpdate:
		orderBy = AimodelSearchOrderByRecentUpdated
	case OrderByDownloadCount:
		orderBy = AimodelSearchOrderByDownload
	case OrderByCollections:
		orderBy = AimodelSearchOrderByCollected
	case OrderByUseCount:
		orderBy = AimodelSearchOrderByReference
	case OrderByDerivativeCount:
		orderBy = AimodelSearchOrferByDerivative
	case OrderBySizeAsc:
		orderBy = AimodelSearchOrderBySizeAsc
	case OrderBySizeDesc:
		orderBy = AimodelSearchOrderBySizeDesc
	case OrderByAliasDesc:
		orderBy = AimodelSearchOrderByAliasDesc
	case OrderByAliasAsc:
		orderBy = AimodelSearchOrderByAliasAsc
	default:
		orderBy = AimodelSearchOrderByDefault
	}
	return orderBy
}

type SearchAimodelReq struct {
	ListOptions
	Keyword            string `json:"q"`
	LabelFilter        string `json:"label_filter"`
	EngineFilter       int    `json:"engine_filter"`
	AimodelName        string `json:"aimodel_name"`
	CreatorName        string `json:"creator_name"`
	OwnerName          string `json:"owner_name"`
	From               string `json:"from"` // 请求来源标识，如 "profile"；用于区分身份边界场景走精确匹配 vs 管理后台模糊匹配
	OwnerId            int64  `json:"owner_id"`
	OwnerType          string `json:"owner_type"`      // individual, organization
	AimodelType        int    `json:"aimodel_type"`    // -1 all, 0 online, 1 local, 2 external(hf)
	MinAccessMode      string `json:"min_access_mode"` // read, write, admin, owner
	Visibility         string `json:"visibility"`      // public, private, all
	Scope              string `json:"scope"`           // all,collected,owned,collaborated,accessible,involved
	Recommend          string `json:"recommend"`       // all,only
	OrderBy            string `json:"order_by"`        // newest, recentupdate, downloadcount, collections, usecount
	UseAdminPermission bool
	User               *User `json:"-"`
}

type AimodelList []*AiModelManage

func (aml AimodelList) loadAttributes(userId int64) error {
	aml.loadOwners()
	aml.loadCollection(userId)
	return nil
}

func (aml AimodelList) loadCollection(userId int64) error {
	if userId <= 0 {
		return nil
	}
	modelIds := make([]string, 0)
	for _, aimodel := range aml {
		modelIds = append(modelIds, aimodel.ID)
	}
	collectionMap := QueryModelCollectedStatus(modelIds, userId)
	for i := 0; i < len(aml); i++ {
		_, ok := collectionMap[aml[i].ID]
		aml[i].IsCollected = ok
	}
	return nil
}

func (aml AimodelList) loadOwners() error {
	ownerIds := make([]int64, 0)
	for _, dataset := range aml {
		if dataset.OwnerID > 0 {
			ownerIds = append(ownerIds, dataset.OwnerID)
		}
	}
	if len(ownerIds) == 0 {
		return nil
	}
	owners, err := GetUsersByIDs(ownerIds)
	if err != nil {
		return err
	}
	for _, dataset := range aml {
		if dataset.OwnerID > 0 {
			for _, owner := range owners {
				if owner.ID == dataset.OwnerID {
					dataset.Owner = owner
					break
				}
			}
		}
	}
	return nil
}

func SearchAimodel(opts SearchAimodelReq) (AimodelList, int64, error) {
	query := builder.NewCond()
	if opts.Keyword != "" {
		query = query.And(builder.Or(builder.Like{"lower_alias", "%" + strings.ToLower(opts.Keyword) + "%"}, builder.Like{"lower_name", "%" + strings.ToLower(opts.Keyword) + "%"}))
	}

	if opts.LabelFilter != "" {
		query = query.And(builder.Like{"label", "%" + opts.LabelFilter + "%"})
	}

	if opts.EngineFilter >= 0 {
		if opts.EngineFilter == 2 {
			query = query.And(builder.In("engine", 2, 121, 122))
		} else {
			query = query.And(builder.Eq{"engine": opts.EngineFilter})
		}
	}

	if opts.AimodelName != "" {
		query = query.And(builder.Eq{"lower_name": strings.ToLower(opts.AimodelName)})
	}

	if opts.OwnerName != "" {
		// 桥：根据 From 标识决定 owner_name 是精确匹配还是 LIKE 模糊匹配。
		// 个人主页（From="profile"）是身份边界场景，必须按精确 lower_name 匹配，
		// 否则 LIKE '%wutianhh1%' 会把 "wutianhh1" 的模型挂到 "wutianhh11"
		// 的个人主页下。其他场景（如管理后台按部分名搜作者）保持原 LIKE 行为。
		var ownerNameCond builder.Cond
		if opts.From == "profile" {
			ownerNameCond = builder.Eq{"lower_name": strings.ToLower(opts.OwnerName)}
		} else {
			ownerNameCond = builder.Like{"lower_name", "%" + strings.ToLower(opts.OwnerName) + "%"}
		}
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(ownerNameCond)))
	}
	if opts.CreatorName != "" {
		query = query.And(builder.In("user_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Like{"lower_name", "%" + strings.ToLower(opts.CreatorName) + "%"})))
	}

	if opts.AimodelType != MODEL_ALL_TYPE {
		query = query.And(builder.Eq{"model_type": opts.AimodelType})
	}

	if opts.Visibility == VisibilityPublic {
		query = query.And(builder.Eq{"is_private": false})
	} else if opts.Visibility == VisibilityPrivate {
		query = query.And(builder.Eq{"is_private": true})
	}

	if opts.OwnerId > 0 {
		query = query.And(builder.Eq{"owner_id": opts.OwnerId})
	}

	if opts.OwnerType == OwnerTypeIndividual {
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Eq{"type": UserTypeIndividual})))
	} else if opts.OwnerType == OwnerTypeOrganization {
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Eq{"type": UserTypeOrganization})))
	}

	if opts.Scope == ScopeOwned {
		ownedCond := builder.NewCond()
		ownedCond = ownedCond.Or(builder.Eq{"owner_id": opts.User.ID})
		ownedCond = ownedCond.Or(builder.In("id",
			builder.Select("`subject_access`.subject_id::text").
				From("subject_access").
				Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
					And(builder.Eq{"`subject_access`.subject_type": AimodelSubject}).
					And(builder.Eq{"`subject_access`.mode": int(AccessModeOwner)}))))
		query = query.And(ownedCond)
	} else if opts.Scope == ScopeCollected {
		query = query.And(builder.In("id",
			builder.Select("`ai_model_collect`.model_id").
				From("ai_model_collect").
				Where(builder.Eq{"`ai_model_collect`.user_id": opts.User.ID})))
	} else if opts.Scope == ScopeCollaborated {
		if opts.MinAccessMode != "" {
			mode := ParseAccessMode(opts.MinAccessMode)
			query = query.And(
				builder.Neq{"owner_id": opts.User.ID},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id::text").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": AimodelSubject}).
							And(builder.Gte{"`subject_access`.mode": int(mode)}).
							And(builder.Neq{"`subject_access`.mode": int(AccessModeOwner)}))),
					builder.In("id", builder.Select("`team_subject`.subject_id::text").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": opts.User.ID}.
							And(builder.Gte{"`team`.aimodel_authorize": int(mode)}).
							And(builder.Neq{"`team`.aimodel_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				),
			)
		} else {
			query = query.And(
				builder.Neq{"owner_id": opts.User.ID},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id::text").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": AimodelSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin))))),
					builder.In("id", builder.Select("`team_subject`.subject_id::text").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": opts.User.ID}.
							And(builder.Neq{"`team`.aimodel_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				),
			)
		}

	} else if opts.Scope == ScopeAccessible {
		if !opts.UseAdminPermission {
			accessibleCond := builder.NewCond()
			if opts.User != nil {
				// accessible = owned + collaborated + public
				accessibleCond = accessibleCond.Or(builder.Eq{"owner_id": opts.User.ID})
				accessibleCond = accessibleCond.Or(builder.In("id",
					builder.Select("`subject_access`.subject_id::text").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": AimodelSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin), int(AccessModeOwner))))))
				accessibleCond = accessibleCond.Or(builder.Eq{"is_private": false})
			} else {
				// accessible = public
				accessibleCond = builder.Eq{"is_private": false}
			}
			query = query.And(accessibleCond)
		}
	} else if opts.Scope == ScopeInvolved {
		if opts.AimodelType == MODEL_HF_TYPE {
			query = query.And(builder.In("id",
				builder.Expr(
					"(SELECT model_id FROM hf_model_file WHERE user_id = ? GROUP BY model_id) "+
						"UNION "+
						"(SELECT model_id FROM hf_model_operation WHERE user_id = ? GROUP BY model_id)",
					opts.User.ID, opts.User.ID,
				),
			))
		}
	}

	if opts.Recommend == RecommendOnly {
		query = query.And(builder.Eq{"recommend": 1})
	}

	totalCount, err := x.Where(query).Count(new(AiModelManage))
	if err != nil {
		return nil, 0, err
	}

	orderBy := GetAimodelSearchOrderBy(opts.OrderBy).String()
	aimodels := make(AimodelList, 0, opts.PageSize)
	err = x.Where(query).Limit(opts.PageSize, opts.PageSize*(opts.Page-1)).OrderBy(orderBy).Find(&aimodels)
	if err != nil {
		return nil, 0, err
	}
	if aimodels == nil {
		return nil, 0, nil
	}
	var userId int64
	if opts.User != nil {
		userId = opts.User.ID
	}
	aimodels.loadAttributes(userId)
	return aimodels, totalCount, nil
}

// CountPersonalAiModels 统计个人的ai模型（概览页面用的）
func CountPersonalAiModels(userId int64) (allCount, privateCount, publicCount int64, err error) {
	query := builder.NewCond()

	// 统计个人的
	ownedCond := builder.NewCond().And(builder.Eq{"owner_id": userId})
	query = query.And(ownedCond)

	// 统计私有的
	privateCond := query
	privateCond = privateCond.And(builder.Eq{"is_private": true})
	privateCount, err = x.Where(privateCond).Count(new(AiModelManage))
	if err != nil {
		return
	}

	// 统计公开的
	publicCond := query
	publicCond = publicCond.And(builder.Eq{"is_private": false})
	publicCount, err = x.Where(publicCond).Count(new(AiModelManage))
	if err != nil {
		return
	}

	allCount = privateCount + publicCount
	return
}

func UpdateAimodelBySize(id string, size int64) error {
	_, err := x.ID(id).Cols("size", "updated_unix").Update(&AiModelManage{
		Size:        size,
		UpdatedUnix: timeutil.TimeStampNow(),
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateAimodelByStatus(id string, status int, status_desc string) error {
	_, err := x.ID(id).Cols("status", "status_desc").Update(&AiModelManage{
		Status:     status,
		StatusDesc: status_desc,
	})
	return err
}

func UpdateAimodel(aimodel AiModelManage) error {
	_, err := x.ID(aimodel.ID).Cols("name", "lower_name", "label", "license", "engine", "is_private", "alias", "lower_alias").Unscoped().Update(aimodel)
	if err != nil {
		return err
	}
	return nil
}

type StorageDeleteFailedAimodel struct {
	ID              string             `xorm:"pk" json:"id"`
	Name            string             `xorm:"INDEX NOT NULL" json:"name"`
	LowerName       string             `xorm:"INDEX NOT NULL"`
	Alias           string             `xorm:"INDEX"`
	LowerAlias      string             `xorm:"INDEX"`
	UserId          int64              `xorm:"NOT NULL" json:"userId"` //creator id
	OwnerID         int64              `xorm:"INDEX"`
	ModelType       int                `xorm:"NULL" json:"modelType"`
	License         string             `xorm:"NULL" json:"license"`
	Engine          int64              `xorm:"NOT NULL DEFAULT 0" json:"engine"`
	IsPrivate       bool               `xorm:"DEFAULT true" json:"isPrivate"`
	StorageType     string             `xorm:"varchar(20)"`
	Path            string             `xorm:"varchar(400) NOT NULL" json:"path"`
	ExternalName    string             `xorm:"text NULL" json:"externalName"`
	Size            int64              `xorm:"NOT NULL" json:"size"`
	Label           string             `xorm:"varchar(1000)" json:"label"`
	DownloadCount   int                `xorm:"NOT NULL DEFAULT 0" json:"downloadCount"`
	Status          int                `xorm:"NOT NULL DEFAULT 0" json:"status"`
	StatusDesc      string             `xorm:"varchar(500)" json:"statusDesc"`
	Recommend       int                `xorm:"NOT NULL DEFAULT 0" json:"recommend"`
	TrainTaskInfo   string             `xorm:"text NULL" json:"trainTaskInfo"`
	DerivativeCount int                `xorm:"NOT NULL DEFAULT 0" json:"derivativeCount"`
	ReferenceCount  int                `xorm:"NOT NULL DEFAULT 0" json:"referenceCount"`
	CollectedCount  int                `xorm:"NOT NULL DEFAULT 0" json:"collectedCount"`
	CreatedUnix     timeutil.TimeStamp `xorm:"created" json:"createdUnix"`
	UpdatedUnix     timeutil.TimeStamp `xorm:"updated" json:"updatedUnix"`
	Owner           *User              `xorm:"-"`
	IsCollected     bool               `xorm:"-"`
}

func InsertStorageDeleteFailedAimodel(aimodel *AiModelManage) error {
	r := &StorageDeleteFailedAimodel{
		ID:              aimodel.ID,
		Name:            aimodel.Name,
		LowerName:       aimodel.LowerName,
		Alias:           aimodel.Alias,
		LowerAlias:      aimodel.LowerAlias,
		UserId:          aimodel.UserId,
		OwnerID:         aimodel.OwnerID,
		ModelType:       aimodel.ModelType,
		License:         aimodel.License,
		Engine:          aimodel.Engine,
		IsPrivate:       aimodel.IsPrivate,
		StorageType:     aimodel.StorageType,
		Path:            aimodel.Path,
		ExternalName:    aimodel.ExternalName,
		Size:            aimodel.Size,
		Label:           aimodel.Label,
		DownloadCount:   aimodel.DownloadCount,
		Status:          aimodel.Status,
		StatusDesc:      aimodel.StatusDesc,
		Recommend:       aimodel.Recommend,
		TrainTaskInfo:   aimodel.TrainTaskInfo,
		DerivativeCount: aimodel.DerivativeCount,
		ReferenceCount:  aimodel.ReferenceCount,
		CollectedCount:  aimodel.CollectedCount,
		CreatedUnix:     aimodel.CreatedUnix,
		UpdatedUnix:     aimodel.UpdatedUnix,
		Owner:           aimodel.Owner,
		IsCollected:     aimodel.IsCollected,
	}
	_, err := x.Insert(r)
	if err != nil {
		return err
	}
	return nil
}

func GetOwnedPublicAimodelsByUserID(userID int64) (AimodelList, error) {
	aimodels := make(AimodelList, 0)
	err := x.Where("owner_id = ? and is_private = false", userID).OrderBy("updated_unix desc").Find(&aimodels)
	if err != nil {
		return nil, err
	}
	aimodels.loadAttributes(userID)
	return aimodels, nil
}

type AimodelShow4Action struct {
	ID        string
	Name      string
	OwnerName string
	Alias     string
}

func (aimodel *AiModelManage) ToActionShow() *AimodelShow4Action {
	aimodel.GetOwner()
	var ownerName string
	if aimodel.Owner != nil {
		ownerName = aimodel.Owner.Name
	}
	return &AimodelShow4Action{
		ID:        aimodel.ID,
		Name:      aimodel.Name,
		OwnerName: ownerName,
		Alias:     aimodel.Alias,
	}
}

type AimodelLabels struct {
	Label  []string
	Engine []int64
}

func GetOwnedPublicAimodelLabels(userId int64) (*AimodelLabels, error) {
	aimodels := make([]AiModelManage, 0)

	err := x.Cols("label", "engine").Where("owner_id =? and is_private = false", userId).Find(&aimodels)
	if err != nil {
		return nil, err
	}
	labelMap := make(map[string]int, 0)
	engineMap := make(map[int64]int, 0)

	labels := make([]string, 0)
	engines := make([]int64, 0)

	for _, aimodel := range aimodels {
		labelSplit := strings.Split(aimodel.Label, " ")
		for _, label := range labelSplit {
			if label == "" {
				continue
			}
			if _, ok := labelMap[label]; !ok {
				labelMap[label] = 0
				labels = append(labels, label)
			}

		}
		engine := aimodel.Engine
		if _, ok := engineMap[engine]; !ok {
			engineMap[engine] = 0
			engines = append(engines, engine)
		}
	}

	return &AimodelLabels{
		Label:  labels,
		Engine: engines,
	}, nil
}

func QueryEvolutionTaskAndModelMap() (map[string][]string, map[string][]int64) {
	exportedModelMap := make(map[string][]string) // map[task.JobID][childMode1.ID,childModel2.ID,..]
	usedByTaskMap := make(map[string][]int64)     // map[parentModel.ID][task1.JobID,task2.JobID,..]

	// 辅助去重集合
	exportedSet := make(map[string]map[string]struct{}) // task.JobID -> set(childModel.ID)
	usedBySet := make(map[string]map[int64]struct{})    // parentModel.ID -> set(task.JobID)

	sess := x.NewSession()
	defer sess.Close()

	cond := "train_task_info <> ''"
	var indexTotal int64
	count, err := sess.Where(cond).Count(new(AiModelManage))
	if err != nil {
		return exportedModelMap, usedByTaskMap
	}

	for {
		var aiModelManageList []*AiModelManage
		err1 := sess.
			Select("*").
			Table(new(AiModelManage)).
			Where(cond).
			OrderBy("created_unix desc").
			Limit(PAGE_SIZE, int(indexTotal)).
			Find(&aiModelManageList)

		if err1 != nil {
			log.Info("query aimodelmanage error:" + err1.Error())
			break
		}

		for _, childModel := range aiModelManageList {
			if childModel.TrainTaskInfo == "" {
				continue
			}

			var task Cloudbrain
			if err := json.Unmarshal([]byte(childModel.TrainTaskInfo), &task); err != nil {
				log.Info("error=" + err.Error())
				continue
			}

			// === exportedModelMap: task.JobID -> childModel.ID ===
			if _, ok := exportedSet[task.JobID]; !ok {
				exportedSet[task.JobID] = make(map[string]struct{})
			}
			if _, exists := exportedSet[task.JobID][childModel.ID]; !exists {
				exportedSet[task.JobID][childModel.ID] = struct{}{}
				exportedModelMap[task.JobID] = append(exportedModelMap[task.JobID], childModel.ID)
			}

			// === usedByTaskMap: parentModelId -> task.JobID ===
			if task.ModelId != "" {
				for _, parentModelId := range strings.Split(task.ModelId, ";") {
					if _, ok := usedBySet[parentModelId]; !ok {
						usedBySet[parentModelId] = make(map[int64]struct{})
					}
					if _, exists := usedBySet[parentModelId][task.ID]; !exists {
						usedBySet[parentModelId][task.ID] = struct{}{}
						usedByTaskMap[parentModelId] = append(usedByTaskMap[parentModelId], task.ID)
					}
				}
			}
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return exportedModelMap, usedByTaskMap
}

func GetAccessibleAimodelIDs(user *User) ([]string, error) {
	aimodelIDs := make([]string, 0, 10)
	if err := x.
		Table("ai_model_manage").
		Cols("id").
		Where(accessibleAimodelCondition(user)).
		Find(&aimodelIDs); err != nil {
		return nil, fmt.Errorf("GetAccessibleAimodelIDs: %v", err)
	}
	return aimodelIDs, nil
}

func AccessibleAimodelIDsQuery(user *User) *builder.Builder {
	return builder.Select("id::uuid").From("ai_model_manage").Where(accessibleAimodelCondition(user))
}

func accessibleAimodelCondition(user *User) builder.Cond {
	var cond = builder.NewCond()

	if user == nil || !user.IsRestricted || user.ID <= 0 {
		orgVisibilityLimit := []structs.VisibleType{structs.VisibleTypePrivate}
		if user == nil || user.ID <= 0 {
			orgVisibilityLimit = append(orgVisibilityLimit, structs.VisibleTypeLimited)
		}
		// 1. Be able to see all non-private datasets that either:
		cond = cond.Or(builder.And(
			builder.Eq{"`ai_model_manage`.is_private": false},
			// 2. Aren't in an private organisation or limited organisation if we're not logged in
			builder.NotIn("`ai_model_manage`.owner_id", builder.Select("id").From("`user`").Where(
				builder.And(
					builder.Eq{"type": UserTypeOrganization},
					builder.In("visibility", orgVisibilityLimit)),
			))))
	}

	if user != nil {
		cond = cond.Or(
			// 2. Be able to see all datasets that we have access to
			builder.In("`ai_model_manage`.id::uuid", builder.Select("subject_id").
				From("`subject_access`").
				Where(builder.And(
					builder.Eq{"user_id": user.ID},
					builder.Eq{"subject_type": AimodelSubject},
					builder.Gt{"mode": int(AccessModeNone)}))),
			// 3. Datasets that we directly own
			builder.Eq{"`ai_model_manage`.owner_id": user.ID},
			// 4. Be able to see all datasets that we are in a team
			builder.In("`ai_model_manage`.id::uuid", builder.Select("`team_subject`.subject_id").
				From("team_subject").
				Where(builder.And(
					builder.Eq{"`team_user`.uid": user.ID},
					builder.Eq{"`team_subject`.subject_type": DatasetSubject},
				)).
				Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id")),
			// 5. Be able to see all public datasets in private organizations that we are an org_user of
			builder.And(builder.Eq{"`ai_model_manage`.is_private": false},
				builder.In("`ai_model_manage`.owner_id",
					builder.Select("`org_user`.org_id").
						From("org_user").
						Where(builder.Eq{"`org_user`.uid": user.ID}))))
	}

	return cond
}

func QueryChildAiModelVague(aimodelID string) ([]*AiModelManage, error) {
	sess := x.NewSession()
	defer sess.Close()

	aiModelManageList := make([]*AiModelManage, 0)

	err := sess.Table(new(AiModelManage)).
		Where("train_task_info ILIKE ?", fmt.Sprintf("%%%s%%", aimodelID)).
		Find(&aiModelManageList)
	if err != nil {
		return nil, err
	}

	return aiModelManageList, nil
}

func QueryAimodelByExternalName(name string) (*AiModelManage, error) {
	aimodel := &AiModelManage{}
	has, err := x.Where("external_name = ?", name).Get(aimodel)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return aimodel, nil
}
