package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	PeriodNotCycle = "NOT_CYCLE"
	PeriodDaily    = "DAILY"
)

type TaskType string

const (
	TaskCreatePublicRepo       TaskType = "CreatePublicRepo"
	TaskCreateIssue            TaskType = "CreateIssue"
	TaskCreatePullRequest      TaskType = "CreatePullRequest"
	TaskCommentIssue           TaskType = "CommentIssue"
	TaskUploadAttachment       TaskType = "UploadAttachment"
	TaskCreateNewModelTask     TaskType = "CreateNewModelTask"
	TaskBindWechat             TaskType = "BindWechat"
	TaskCreateCloudbrainTask   TaskType = "CreateCloudbrainTask"
	TaskCreateSuperComputeTask TaskType = "CreateSuperComputeTask"
	TaskDatasetRecommended     TaskType = "DatasetRecommended"
	TaskCreateImage            TaskType = "CreateImage"
	TaskImageRecommend         TaskType = "ImageRecommend"
	TaskChangeUserAvatar       TaskType = "ChangeUserAvatar"
	TaskPushCommits            TaskType = "PushCommits"
	TaskInviteFriendRegister   TaskType = "TaskInviteFriendRegister"
	TaskCreateDataset          TaskType = "TaskCreateDataset"
	TaskCreateAimodel          TaskType = "TaskCreateAimodel"
	TaskAimodelRecommended     TaskType = "TaskAimodelRecommended"
)

func (t TaskType) ChineseName() string {
	switch t {
	case TaskCreatePublicRepo:
		return "创建公开项目"
	case TaskCreateIssue:
		return "每日提出任务"
	case TaskCreatePullRequest:
		return "每日提出PR"
	case TaskCommentIssue:
		return "发表评论"
	case TaskUploadAttachment:
		return "上传数据集文件"
	case TaskCreateNewModelTask:
		return "导入新模型"
	case TaskBindWechat:
		return "完成微信扫码验证"
	case TaskCreateCloudbrainTask:
		return "每日运行云脑任务"
	case TaskDatasetRecommended:
		return "数据集被平台推荐"
	case TaskCreateImage:
		return "提交新公开镜像"
	case TaskImageRecommend:
		return "镜像被平台推荐"
	case TaskChangeUserAvatar:
		return "首次更换头像"
	case TaskPushCommits:
		return "每日commit"
	case TaskInviteFriendRegister:
		return "邀请好友"
	case TaskAimodelRecommended:
		return "模型被平台推荐"
	case TaskCreateAimodel:
		return "创建模型"
	}
	return "--"
}

func GetTaskTypeFromAction(a ActionType) TaskType {
	switch a {
	case ActionCreateDebugGPUTask,
		ActionCreateDebugNPUTask,
		ActionCreateTrainTask,
		ActionCreateInferenceTask,
		ActionCreateBenchMarkTask,
		ActionCreateGPUTrainTask,
		ActionCreateGrampusGPUDebugTask,
		ActionCreateGrampusNPUDebugTask,
		ActionCreateGrampusNPUTrainTask,
		ActionCreateGrampusGCUDebugTask,
		ActionCreateGrampusGCUTrainTask,
		ActionCreateGrampusMLUDebugTask,
		ActionCreateGrampusDCUDebugTask,
		ActionCreateGrampusILUVATARDebugTask,
		ActionCreateGrampusMETAXDebugTask,
		ActionCreateSuperComputeTask,
		ActionCreateGrampusGPUOnlineInferTask,
		ActionCreateGrampusGPUTrainTask,
		ActionCreateGrampusGPUInferenceTask,
		ActionCreateGrampusILUVATARInferenceTask,
		ActionCreateGrampusILUVATARTrainTask,
		ActionCreateGrampusModelExperienceTask,
		ActionCreateGrampusFinetuneTask,
		ActionCreateGrampusDCUTrainTask,
		ActionCreateGrampusMETAXGPGPUTrainTask,
		ActionCreateGrampusSdFinetuneTask,
		ActionCreateGrampusComfyuiExperienceTask,
		ActionCreateGrampusEvalTask,
		ActionCreateGrampusBIRENGPUDebugTask,
		ActionCreateGrampusBIRENGPUTrainTask,
		ActionCreateGeneralGPUTask:
		return TaskCreateCloudbrainTask
	case ActionCreateRepo:
		return TaskCreatePublicRepo
	case ActionCreatePullRequest:
		return TaskCreatePullRequest
	case ActionCommentIssue:
		return TaskCommentIssue
	case ActionUploadAttachment:
		return TaskUploadAttachment
	case ActionCreateNewModelTask:
		return TaskCreateNewModelTask
	case ActionBindWechat:
		return TaskBindWechat
	case ActionDatasetRecommended:
		return TaskDatasetRecommended
	case ActionImageRecommend:
		return TaskImageRecommend
	case ActionCreateImage:
		return TaskCreateImage
	case ActionChangeUserAvatar:
		return TaskChangeUserAvatar
	case ActionCommitRepo,
		ActionDeleteBranch,
		ActionPushTag,
		ActionDeleteTag:
		return TaskPushCommits
	case ActionCreateIssue:
		return TaskCreateIssue
	case ActionInviteFriendRegister:
		return TaskInviteFriendRegister
	case ActionCreateDataset:
		return TaskCreateDataset
	case ActionCreateAimodel:
		return TaskCreateAimodel
	case ActionAimodelRecommended:
		return TaskAimodelRecommended
	}
	return ""
}

// PointTaskConfig Only add and delete are allowed, edit is not allowed
// so if you want to edit config for some task code,please delete first and add new one
type TaskConfig struct {
	ID          int64  `xorm:"pk autoincr"`
	TaskCode    string `xorm:"NOT NULL"`
	Title       string
	AwardType   string  `xorm:"NOT NULL"`
	AwardAmount float64 `xorm:"NOT NULL"`
	CreatorId   int64   `xorm:"NOT NULL"`
	CreatorName string
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	DeletedAt   timeutil.TimeStamp `xorm:"deleted"`
	DeleterId   int64
	DeleterName string
}

type TaskConfigWithLimit struct {
	ID          int64
	TaskCode    string
	Title       string
	AwardType   string
	AwardAmount float64
	Creator     string
	IsDeleted   bool
	CreatedUnix timeutil.TimeStamp
	DeleteAt    timeutil.TimeStamp
	Limiters    []*LimitConfigVO
}

type TaskConfigWithLimitResponse struct {
	Records  []*TaskConfigWithSingleLimit
	Total    int64
	PageSize int
	Page     int
}

type TaskConfigWithSingleLimit struct {
	ID          int64
	TaskCode    string
	AwardType   string
	AwardAmount float64
	Creator     string
	IsDeleted   bool
	CreatedUnix timeutil.TimeStamp
	DeleteAt    timeutil.TimeStamp
	RefreshRate string
	LimitNum    float64
}

type TaskAndLimiterConfig struct {
	TaskConfig  TaskConfig  `xorm:"extends"`
	LimitConfig LimitConfig `xorm:"extends"`
}

type PointRule struct {
	UserDailyLimit float64
	TaskRules      []TaskRule
}

type TaskRule struct {
	TaskCode    string
	AwardType   string
	AwardAmount float64
	RefreshRate string
	LimitNum    float64
}

func (TaskAndLimiterConfig) TableName() string {
	return "task_config"
}

type BatchLimitConfigVO struct {
	ConfigList []TaskConfigWithLimit
}

func getTaskConfig(t *TaskConfig) (*TaskConfig, error) {
	has, err := x.Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return t, nil
}

func GetTaskConfigByTaskCode(taskCode string) (*TaskConfig, error) {
	t := &TaskConfig{
		TaskCode: taskCode,
	}
	return getTaskConfig(t)
}

func GetTaskConfigByID(id int64) (*TaskConfig, error) {
	t := &TaskConfig{
		ID: id,
	}
	return getTaskConfig(t)
}

func GetTaskConfigList() ([]*TaskConfig, error) {
	r := make([]*TaskConfig, 0)
	err := x.Find(&r)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

type GetTaskConfigOpts struct {
	ListOptions
	Status   int //1 normal 2 deleted
	TaskType string
}

func GetTaskConfigPageWithDeleted(opt GetTaskConfigOpts) ([]*TaskAndLimiterConfig, int64, error) {
	if opt.Page <= 0 {
		opt.Page = 1
	}
	cond := builder.NewCond()
	if opt.TaskType != "" {
		cond = cond.And(builder.Eq{"task_code": opt.TaskType})
	}

	var count int64
	var err error
	if opt.Status == 1 {
		subCond := builder.NewCond()
		subCond = subCond.Or(builder.IsNull{"task_config.deleted_at"})
		subCond = subCond.Or(builder.Eq{"task_config.deleted_at": 0})
		cond = cond.And(subCond)
	} else if opt.Status == 2 {
		cond = cond.And(builder.Gt{"task_config.deleted_at": 0})
	}
	count, err = x.Unscoped().Where(cond).Count(&TaskConfig{})
	if err != nil {
		return nil, 0, err
	}
	r := make([]*TaskAndLimiterConfig, 0)
	err = x.Join("LEFT", "limit_config", "task_config.id = limit_config.related_id").
		Unscoped().Where(cond).Limit(opt.PageSize, (opt.Page-1)*opt.PageSize).
		OrderBy("task_config.deleted_at desc,task_config.id desc").Find(&r)

	if len(r) == 0 {
		return nil, 0, ErrRecordNotExist{}
	}
	return r, count, nil
}

func EditTaskConfig(config TaskConfigWithLimit, doer *User) error {
	sess := x.NewSession()
	defer sess.Close()

	//delete old task config
	p := &TaskConfig{
		ID: config.ID,
	}
	_, err := sess.Delete(p)
	if err != nil {
		sess.Rollback()
		return err
	}
	//update deleter
	p.DeleterId = doer.ID
	p.DeleterName = doer.Name
	sess.Where("id = ?", config.ID).Unscoped().Update(p)

	//add new config
	t := &TaskConfig{
		TaskCode:    config.TaskCode,
		Title:       config.Title,
		AwardType:   config.AwardType,
		AwardAmount: config.AwardAmount,
		CreatorId:   doer.ID,
		CreatorName: doer.Name,
	}
	_, err = sess.InsertOne(t)
	if err != nil {
		sess.Rollback()
		return err
	}

	//delete old limiter config
	lp := &LimitConfig{
		RelatedId: config.ID,
	}
	_, err = sess.Delete(lp)
	if err != nil {
		sess.Rollback()
		return err
	}
	lp.DeleterName = doer.Name
	lp.DeleterId = doer.ID
	//update deleter
	sess.Where("related_id = ?", config.ID).Unscoped().Update(lp)

	//add new limiter config
	if config.Limiters != nil && len(config.Limiters) > 0 {
		for _, v := range config.Limiters {
			//add new config
			l := &LimitConfig{
				Title:       v.Title,
				RefreshRate: v.RefreshRate,
				Scope:       v.Scope,
				LimitNum:    v.LimitNum,
				LimitCode:   config.TaskCode,
				LimitType:   LimitTypeTask.Name(),
				CreatorId:   doer.ID,
				CreatorName: doer.Name,
				RelatedId:   t.ID,
			}
			_, err = sess.Insert(l)
			if err != nil {
				sess.Rollback()
				return err
			}
		}
	}
	sess.Commit()
	return nil
}

func NewTaskConfig(config TaskConfigWithLimit, doer *User) error {
	sess := x.NewSession()
	defer sess.Close()

	//add new config
	t := &TaskConfig{
		TaskCode:    config.TaskCode,
		Title:       config.Title,
		AwardType:   config.AwardType,
		AwardAmount: config.AwardAmount,
		CreatorId:   doer.ID,
		CreatorName: doer.Name,
	}
	_, err := sess.InsertOne(t)
	if err != nil {
		sess.Rollback()
		return err
	}

	//add new limiter config
	if config.Limiters != nil && len(config.Limiters) > 0 {
		for _, v := range config.Limiters {
			//add new config
			l := &LimitConfig{
				RelatedId:   t.ID,
				Title:       v.Title,
				RefreshRate: v.RefreshRate,
				Scope:       v.Scope,
				LimitNum:    v.LimitNum,
				LimitCode:   config.TaskCode,
				LimitType:   LimitTypeTask.Name(),
				CreatorId:   doer.ID,
				CreatorName: doer.Name,
			}
			_, err = sess.Insert(l)
			if err != nil {
				sess.Rollback()
				return err
			}
		}
	}
	sess.Commit()
	return nil
}

func DelTaskConfig(id int64, doer *User) error {
	sess := x.NewSession()
	defer sess.Close()

	//delete old task config
	p := &TaskConfig{
		ID: id,
	}
	_, err := sess.Delete(p)
	if err != nil {
		sess.Rollback()
		return err
	}
	//update deleter
	p.DeleterId = doer.ID
	p.DeleterName = doer.Name
	sess.Where("id = ?", id).Unscoped().Update(p)
	//delete old limiter config
	lp := &LimitConfig{
		RelatedId: id,
	}
	_, err = sess.Delete(lp)
	if err != nil {
		sess.Rollback()
		return err
	}
	lp.DeleterName = doer.Name
	lp.DeleterId = doer.ID
	//update deleter
	sess.Where("related_id = ?", id).Unscoped().Update(lp)
	sess.Commit()
	return nil
}
