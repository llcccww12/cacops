// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"

	"github.com/unknwon/com"
	"xorm.io/builder"
)

// ActionType represents the type of an action.
type ActionType int

// Possible action types.
const (
	ActionCreateRepo         ActionType = iota + 1 // 1
	ActionRenameRepo                               // 2
	ActionStarRepo                                 // 3
	ActionWatchRepo                                // 4
	ActionCommitRepo                               // 5
	ActionCreateIssue                              // 6
	ActionCreatePullRequest                        // 7
	ActionTransferRepo                             // 8
	ActionPushTag                                  // 9
	ActionCommentIssue                             // 10
	ActionMergePullRequest                         // 11
	ActionCloseIssue                               // 12
	ActionReopenIssue                              // 13
	ActionClosePullRequest                         // 14
	ActionReopenPullRequest                        // 15
	ActionDeleteTag                                // 16
	ActionDeleteBranch                             // 17
	ActionMirrorSyncPush                           // 18
	ActionMirrorSyncCreate                         // 19
	ActionMirrorSyncDelete                         // 20
	ActionApprovePullRequest                       // 21
	ActionRejectPullRequest                        // 22
	ActionCommentPull                              // 23

	ActionUploadAttachment                   //24
	ActionCreateDebugGPUTask                 //25
	ActionCreateDebugNPUTask                 //26
	ActionCreateTrainTask                    //27
	ActionCreateInferenceTask                // 28
	ActionCreateBenchMarkTask                //29
	ActionCreateNewModelTask                 //30
	ActionCreateGPUTrainTask                 //31
	ActionCreateGrampusNPUTrainTask          //32
	ActionCreateGrampusGPUTrainTask          //33
	ActionBindWechat                         //34
	ActionDatasetRecommended                 //35
	ActionCreateImage                        //36
	ActionImageRecommend                     //37
	ActionChangeUserAvatar                   //38
	ActionCreateGrampusNPUDebugTask          //39
	ActionCreateGrampusGPUDebugTask          //40
	ActionCreateGrampusGCUDebugTask          //41
	ActionCreateGrampusGCUTrainTask          //42
	ActionCreateGrampusMLUDebugTask          //43
	ActionCreateGrampusMLUTrainTask          //44
	ActionCreateGrampusGPUOnlineInferTask    //45
	ActionCreateGrampusDCUDebugTask          //46
	ActionCreateSuperComputeTask             //47
	ActionCreateGrampusILUVATARDebugTask     //48
	ActionCreateGrampusMETAXDebugTask        //49
	ActionCreateGrampusGPUInferenceTask      //50
	ActionCreateGrampusILUVATARInferenceTask //51
	ActionInviteFriendRegister               //52
	ActionCreateGrampusILUVATARTrainTask     //53
	ActionCreateGeneralGPUTask               //54
	ActionCreateGrampusModelExperienceTask   //55
	ActionCreateGrampusFinetuneTask          //56
	ActionCreateGrampusDCUTrainTask          //57
	ActionCreateIFLYTEKFinetuneTask          //58
	ActionCreateIFLYTEKModelExperienceTask   //59
	ActionCreateGrampusSdFinetuneTask        //60
	ActionCreateGrampusComfyuiExperienceTask //61
	ActionCreateDataset                      //62
	ActionDeleteDataset                      //63
	ActionRenameDataset                      //64
	ActionCreateGrampusMETAXGPGPUTrainTask   //65
	ActionCreateGrampusEvalTask              //66
	ActionCreateAimodel                      //67
	ActionDeleteAimodel                      //68
	ActionRenameAimodel                      //69
	ActionAimodelRecommended                 //70
	ActionCreateGrampusBIRENGPUDebugTask     //71
	ActionCreateGrampusBIRENGPUTrainTask     //72
)

var AI_TASK_ACTION_LIST = []ActionType{
	ActionCreateDebugGPUTask,
	ActionCreateDebugNPUTask,
	ActionCreateTrainTask,
	ActionCreateInferenceTask,
	ActionCreateBenchMarkTask,
	ActionCreateGPUTrainTask,
	ActionCreateGrampusGPUDebugTask,
	ActionCreateGrampusGPUOnlineInferTask,
	ActionCreateGrampusNPUDebugTask,
	ActionCreateGrampusNPUTrainTask,
	ActionCreateGrampusGPUTrainTask,
	ActionCreateGrampusGCUTrainTask,
	ActionCreateGrampusGCUDebugTask,
	ActionCreateGrampusDCUDebugTask,
	ActionCreateGrampusMLUDebugTask,
	ActionCreateGrampusILUVATARDebugTask,
	ActionCreateGrampusMETAXDebugTask,
	ActionCreateSuperComputeTask,
	ActionCreateGrampusILUVATARInferenceTask,
	ActionCreateGrampusGPUInferenceTask,
	ActionCreateGrampusILUVATARTrainTask,
	ActionCreateGrampusModelExperienceTask,
	ActionCreateGrampusSdFinetuneTask,
	ActionCreateGrampusComfyuiExperienceTask,
	ActionCreateGrampusFinetuneTask,
	ActionCreateGrampusDCUTrainTask,
	ActionCreateGrampusMETAXGPGPUTrainTask,
	ActionCreateGeneralGPUTask,
	ActionCreateIFLYTEKFinetuneTask,
	ActionCreateGrampusEvalTask,
	ActionCreateIFLYTEKModelExperienceTask,
}

// Action represents user operation type and other information to
// repository. It implemented interface base.Actioner so that can be
// used in template render.
type Action struct {
	ID            int64 `xorm:"pk autoincr"`
	UserID        int64 `xorm:"INDEX"` // Receiver user id.
	OpType        ActionType
	ActUserID     int64       `xorm:"INDEX"` // Action user id.
	ActUser       *User       `xorm:"-"`
	RepoID        int64       `xorm:"INDEX"`
	Repo          *Repository `xorm:"-"`
	CommentID     int64       `xorm:"INDEX"`
	Comment       *Comment    `xorm:"-"`
	IsDeleted     bool        `xorm:"INDEX NOT NULL DEFAULT false"`
	RefName       string
	IsPrivate     bool               `xorm:"INDEX NOT NULL DEFAULT false"`
	IsTransformed bool               `xorm:"INDEX NOT NULL DEFAULT false"`
	Content       string             `xorm:"TEXT"`
	CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created"`
	DatasetID     *string            `xorm:"uuid INDEX  DEFAULT null"`
	Dataset       *DatasetRegistry   `xorm:"-"`
	Cloudbrain    *Cloudbrain        `xorm:"-"`
	AimodelID     *string            `xorm:"uuid INDEX  DEFAULT null"` // for future use
	Aimodel       *AiModelManage     `xorm:"-"`
}

type ActionShow struct {
	OpType                   ActionType
	TaskType                 TaskType
	RepoLink                 string
	ShortRepoFullDisplayName string
	Content                  string
	RefName                  string
	IssueInfos               []string
	CommentLink              string
	Cloudbrain               *CloudbrainShow4Action
	Dataset                  *DatasetRegistryShow4Action
	Aimodel                  *AimodelShow4Action
	Data                     map[string]interface{}
}

func (a *ActionShow) AddData(key string, val interface{}) {
	if a.Data == nil {
		a.Data = map[string]interface{}{key: val}
	} else {
		a.Data[key] = val
	}
}

// GetOpType gets the ActionType of this action.
func (a *Action) GetOpType() ActionType {
	return a.OpType
}

func (a *Action) loadActUser() {
	if a.ActUser != nil {
		return
	}
	var err error
	a.ActUser, err = GetUserByID(a.ActUserID)
	if err == nil {
		return
	} else if IsErrUserNotExist(err) {
		a.ActUser = NewGhostUser()
	} else {
		log.Error("GetUserByID(%d): %v", a.ActUserID, err)
	}
}

func (a *Action) FilterCloudbrainInfo() {
	if a.Cloudbrain == nil {
		return
	}

	if a.Cloudbrain.DeletedAt.IsZero() {
		newCloudbrain := &Cloudbrain{}
		newCloudbrain.ID = a.Cloudbrain.ID
		newCloudbrain.JobType = a.Cloudbrain.JobType
		newCloudbrain.ComputeResource = a.Cloudbrain.ComputeResource
		newCloudbrain.DisplayJobName = a.Cloudbrain.DisplayJobName
		a.Cloudbrain = newCloudbrain
	} else {
		a.Cloudbrain = nil
	}
}

func (a *Action) FilterDatasetInfo() {
	if a.Dataset == nil {
		return
	}

	ownerName := ""
	if a.Dataset.Owner != nil {
		ownerName = a.Dataset.Owner.Name
	}
	a.Dataset = &DatasetRegistry{
		ID:    a.Dataset.ID,
		Name:  a.Dataset.Name,
		Alias: a.Dataset.Alias,
		Owner: &User{
			Name: ownerName,
		},
	}

}

func (a *Action) FilterAimodelInfo() {
	if a.Aimodel == nil {
		return
	}

	ownerName := ""
	if a.Aimodel.Owner != nil {
		ownerName = a.Aimodel.Owner.Name
	}
	a.Aimodel = &AiModelManage{
		ID:    a.Aimodel.ID,
		Name:  a.Aimodel.Name,
		Alias: a.Aimodel.Alias,
		Owner: &User{
			Name: ownerName,
		},
	}

}

func (a *Action) loadRepo() {
	if a.Repo != nil {
		return
	}
	var err error
	if a.RepoID == 0 {
		return
	}
	a.Repo, err = GetRepositoryByID(a.RepoID)
	if err != nil {
		log.Error("GetRepositoryByID(%d): %v", a.RepoID, err)
	}
}
func (a *Action) loadDataset() {
	if a.Dataset != nil {
		return
	}
	if a.DatasetID == nil || *a.DatasetID == "" {
		return
	}
	var err error
	dataset, err := GetDatasetRegistryByID(*a.DatasetID)
	if dataset != nil {
		dataset.GetOwner()
		a.Dataset = dataset
	}
	if err != nil {
		log.Error("GetDatasetRegistryByID(%s): %v", *a.DatasetID, err)
	}
}
func (a *Action) loadAimodel() {
	if a.Aimodel != nil {
		return
	}
	if a.AimodelID == nil || *a.AimodelID == "" {
		return
	}
	var err error
	aimodel, err := GetAimodelByID(*a.AimodelID)
	if aimodel != nil {
		aimodel.GetOwner()
		a.Aimodel = aimodel
	}
	if err != nil {
		log.Error("GetAimodelByID(%s): %v", *a.AimodelID, err)
	}
}

func (a *Action) loadComment() {
	if a.CommentID == 0 {
		return
	}
	if a.Comment != nil {
		return
	}
	a.Comment, _ = GetCommentByID(a.CommentID)
}

func (a *Action) loadCloudbrain() {
	if !a.IsCloudbrainAction() {
		return
	}
	cloudbrain := &Cloudbrain{}
	cloudbrainId, _ := strconv.ParseInt(a.Content, 10, 64)
	jobId := a.Content

	//由于各个类型的云脑任务在发布action的时候，content字段保存的ID含义不同，部分取的是ID,部分取的是jobId
	//所以在查询action对应的cloudbrain对象时，以这两个字段做为条件查询
	if has, err := x.
		Where(builder.Or(builder.Eq{"id": cloudbrainId}).Or(builder.Eq{"job_id": jobId})).Unscoped().
		Get(cloudbrain); err != nil || !has {
		return
	}

	a.Cloudbrain = cloudbrain

}

// GetActFullName gets the action's user full name.
func (a *Action) GetActFullName() string {
	a.loadActUser()
	return a.ActUser.FullName
}

// GetActUserName gets the action's user name.
func (a *Action) GetActUserName() string {
	a.loadActUser()
	return a.ActUser.Name
}

// ShortActUserName gets the action's user name trimmed to max 20
// chars.
func (a *Action) ShortActUserName() string {
	return base.EllipsisString(a.GetActUserName(), 20)
}

// GetDisplayName gets the action's display name based on DEFAULT_SHOW_FULL_NAME, or falls back to the username if it is blank.
func (a *Action) GetDisplayName() string {
	if setting.UI.DefaultShowFullName {
		trimmedFullName := strings.TrimSpace(a.GetActFullName())
		if len(trimmedFullName) > 0 {
			return trimmedFullName
		}
	}
	return a.ShortActUserName()
}

// GetDisplayNameTitle gets the action's display name used for the title (tooltip) based on DEFAULT_SHOW_FULL_NAME
func (a *Action) GetDisplayNameTitle() string {
	if setting.UI.DefaultShowFullName {
		return a.ShortActUserName()
	}
	return a.GetActFullName()
}

// GetActAvatar the action's user's avatar link
func (a *Action) GetActAvatar() string {
	a.loadActUser()
	return a.ActUser.RelAvatarLink()
}

// GetRepoUserName returns the name of the action repository owner.
func (a *Action) GetRepoUserName() string {
	a.loadRepo()
	if a.Repo == nil {
		return ""
	}
	return a.Repo.OwnerName
}

// ShortRepoUserName returns the name of the action repository owner
// trimmed to max 20 chars.
func (a *Action) ShortRepoUserName() string {
	return base.EllipsisString(a.GetRepoUserName(), 20)
}

// GetRepoName returns the name of the action repository.
func (a *Action) GetRepoName() string {
	a.loadRepo()
	if a.Repo == nil {
		return ""
	}
	return a.Repo.Name
}

// GetRepoName returns the name of the action repository.
func (a *Action) GetRepoDisplayName() string {
	a.loadRepo()
	if a.Repo == nil {
		return ""
	}
	return a.Repo.DisplayName()
}

// ShortRepoName returns the name of the action repository
// trimmed to max 33 chars.
func (a *Action) ShortRepoName() string {
	return base.EllipsisString(a.GetRepoName(), 33)
}

// ShortRepoName returns the name of the action repository
// trimmed to max 33 chars.
func (a *Action) ShortRepoDisplayName() string {
	return base.EllipsisString(a.GetRepoDisplayName(), 33)
}

// GetRepoPath returns the virtual path to the action repository.
func (a *Action) GetRepoPath() string {
	return path.Join(a.GetRepoUserName(), a.GetRepoName())
}

// ShortRepoPath returns the virtual path to the action repository
// trimmed to max 20 + 1 + 33 chars.
func (a *Action) ShortRepoPath() string {
	return path.Join(a.ShortRepoUserName(), a.ShortRepoName())
}

// ShortRepoPath returns the virtual path to the action repository
// trimmed to max 20 + 1 + 33 chars.
func (a *Action) ShortRepoFullDisplayName() string {
	return path.Join(a.ShortRepoUserName(), a.ShortRepoDisplayName())
}

// GetRepoLink returns relative link to action repository.
func (a *Action) GetRepoLink() string {
	if len(setting.AppSubURL) > 0 {
		return path.Join(setting.AppSubURL, a.GetRepoPath())
	}
	return "/" + a.GetRepoPath()
}

func (a *Action) ToShow() *ActionShow {
	actionShow := &ActionShow{}
	actionShow.OpType = a.OpType
	actionShow.TaskType = GetTaskTypeFromAction(a.OpType)
	actionShow.Content = a.Content
	actionShow.RefName = a.RefName

	if strings.Contains(a.Content, "|") && a.IsIssueAction() {
		actionShow.IssueInfos = a.GetIssueInfos()
	}
	if strings.Contains(a.Content, "|") && a.IsInviteAction() {
		ids := strings.Split(a.Content, "|")
		if len(ids) >= 2 {
			var invitedId int64
			var invitedName string
			if len(ids) >= 4 {
				invitedName = ids[3]
			}
			invitedId, _ = strconv.ParseInt(ids[1], 10, 64)
			if invitedId > 0 {
				invitedUser, _ := GetUserByID(invitedId)
				if invitedUser != nil {
					actionShow.AddData("InvitedUserName", invitedUser.Name)
					actionShow.AddData("InvitedUserNotExists", false)
				} else {
					actionShow.AddData("InvitedUserName", invitedName)
					actionShow.AddData("InvitedUserNotExists", true)
				}
			}
		}
		actionShow.IssueInfos = a.GetIssueInfos()
	}
	if a.RepoID > 0 {
		a.loadRepo()
		if a.Repo != nil {
			actionShow.RepoLink = a.GetRepoLink()
			actionShow.ShortRepoFullDisplayName = a.ShortRepoFullDisplayName()
		}
	}
	if a.CommentID > 0 {
		a.loadComment()
		if a.Comment != nil {
			actionShow.CommentLink = a.GetCommentLink()
		}
	}
	if a.DatasetID != nil && *a.DatasetID != "" {
		a.loadDataset()
		if a.Dataset != nil {
			actionShow.Dataset = a.Dataset.ToActionShow()
		}
	}
	if a.AimodelID != nil && *a.AimodelID != "" {
		a.loadAimodel()
		if a.Aimodel != nil {
			actionShow.Aimodel = a.Aimodel.ToActionShow()
		}
	}

	if a.IsCloudbrainAction() {
		a.loadCloudbrain()
		if a.Cloudbrain != nil {
			c := &CloudbrainShow4Action{
				ID:              a.Cloudbrain.ID,
				JobID:           a.Cloudbrain.JobID,
				Type:            a.Cloudbrain.Type,
				JobType:         a.Cloudbrain.JobType,
				DisplayJobName:  a.Cloudbrain.DisplayJobName,
				ComputeResource: a.Cloudbrain.ComputeResource,
			}
			actionShow.Cloudbrain = c
		}
	}

	return actionShow
}

// GetRepositoryFromMatch returns a *Repository from a username and repo strings
func GetRepositoryFromMatch(ownerName string, repoName string) (*Repository, error) {
	var err error
	refRepo, err := GetRepositoryByOwnerAndName(ownerName, repoName)
	if err != nil {
		if IsErrRepoNotExist(err) {
			log.Warn("Repository referenced in commit but does not exist: %v", err)
			return nil, err
		}
		log.Error("GetRepositoryByOwnerAndName: %v", err)
		return nil, err
	}
	return refRepo, nil
}

// GetCommentLink returns link to action comment.
func (a *Action) GetCommentLink() string {
	return a.getCommentLink(x)
}

func (a *Action) getCommentLink(e Engine) string {
	if a == nil {
		return "#"
	}
	if a.Comment == nil && a.CommentID != 0 {
		a.Comment, _ = getCommentByID(e, a.CommentID)
	}
	if a.Comment != nil {
		return a.Comment.HTMLURL()
	}
	if len(a.GetIssueInfos()) == 0 {
		return "#"
	}
	//Return link to issue
	issueIDString := a.GetIssueInfos()[0]
	issueID, err := strconv.ParseInt(issueIDString, 10, 64)
	if err != nil {
		return "#"
	}

	issue, err := getIssueByID(e, issueID)
	if err != nil {
		return "#"
	}

	if err = issue.loadRepo(e); err != nil {
		return "#"
	}

	return issue.HTMLURL()
}

// GetBranch returns the action's repository branch.
func (a *Action) GetBranch() string {
	return a.RefName
}

// GetContent returns the action's content.
func (a *Action) GetContent() string {
	return a.Content
}

// GetCreate returns the action creation time.
func (a *Action) GetCreate() time.Time {
	return a.CreatedUnix.AsTime()
}

// GetIssueInfos returns a list of issues associated with
// the action.
func (a *Action) GetIssueInfos() []string {
	return strings.SplitN(a.Content, "|", 2)
}

// GetIssueTitle returns the title of first issue associated
// with the action.
func (a *Action) GetIssueTitle() string {
	index := com.StrTo(a.GetIssueInfos()[0]).MustInt64()
	issue, err := GetIssueByIndex(a.RepoID, index)
	if err != nil {
		log.Error("GetIssueByIndex: %v", err)
		return "500 when get issue"
	}
	return issue.Title
}

// GetIssueContent returns the content of first issue associated with
// this action.
func (a *Action) GetIssueContent() string {
	index := com.StrTo(a.GetIssueInfos()[0]).MustInt64()
	issue, err := GetIssueByIndex(a.RepoID, index)
	if err != nil {
		log.Error("GetIssueByIndex: %v", err)
		return "500 when get issue"
	}
	return issue.Content
}

func (a *Action) IsCloudbrainAction() bool {
	switch a.OpType {
	case ActionCreateDebugGPUTask,
		ActionCreateDebugNPUTask,
		ActionCreateTrainTask,
		ActionCreateInferenceTask,
		ActionCreateBenchMarkTask,
		ActionCreateGPUTrainTask,
		ActionCreateGrampusGPUDebugTask,
		ActionCreateGrampusGPUOnlineInferTask,
		ActionCreateGrampusNPUDebugTask,
		ActionCreateGrampusNPUTrainTask,
		ActionCreateGrampusGPUTrainTask,
		ActionCreateGrampusGCUTrainTask,
		ActionCreateGrampusGCUDebugTask,
		ActionCreateGrampusDCUDebugTask,
		ActionCreateGrampusMLUDebugTask,
		ActionCreateGrampusILUVATARDebugTask,
		ActionCreateGrampusMETAXDebugTask,
		ActionCreateSuperComputeTask,
		ActionCreateGrampusILUVATARInferenceTask,
		ActionCreateGrampusGPUInferenceTask,
		ActionCreateGrampusILUVATARTrainTask,
		ActionCreateGrampusModelExperienceTask,
		ActionCreateGrampusSdFinetuneTask,
		ActionCreateGrampusComfyuiExperienceTask,
		ActionCreateGrampusFinetuneTask,
		ActionCreateGrampusDCUTrainTask,
		ActionCreateGrampusMETAXGPGPUTrainTask,
		ActionCreateGeneralGPUTask,
		ActionCreateIFLYTEKFinetuneTask,
		ActionCreateGrampusEvalTask,
		ActionCreateGrampusBIRENGPUDebugTask,
		ActionCreateGrampusBIRENGPUTrainTask,
		ActionCreateIFLYTEKModelExperienceTask:
		return true
	}
	return false
}

func (a *Action) IsIssueAction() bool {
	switch a.OpType {
	case ActionCreateIssue,
		ActionCloseIssue,
		ActionClosePullRequest,
		ActionReopenIssue,
		ActionReopenPullRequest,
		ActionCommentPull,
		ActionCommentIssue,
		ActionCreatePullRequest,
		ActionApprovePullRequest,
		ActionRejectPullRequest,
		ActionMergePullRequest:
		return true
	}
	return false
}

func (a *Action) IsInviteAction() bool {
	switch a.OpType {
	case ActionInviteFriendRegister:
		return true
	}
	return false
}

// GetFeedsOptions options for retrieving feeds
type GetFeedsOptions struct {
	RequestedUser   *User // the user we want activity for
	Actor           *User // the user viewing the activity
	IncludePrivate  bool  // include private actions
	OnlyPerformedBy bool  // only actions performed by requested user
	IncludeDeleted  bool  // include deleted actions
}

// GetFeeds returns actions according to the provided options
func GetFeeds(opts GetFeedsOptions) ([]*Action, error) {
	cond := builder.NewCond()
	var repoIDs []int64
	var datasetIDs []string
	var actorID int64
	var aimodelIDs []string

	if opts.Actor != nil {
		actorID = opts.Actor.ID
	}

	if opts.RequestedUser.IsOrganization() {
		env, err := opts.RequestedUser.AccessibleReposEnv(actorID)
		if err != nil {
			return nil, fmt.Errorf("AccessibleReposEnv: %v", err)
		}
		if repoIDs, err = env.RepoIDs(1, opts.RequestedUser.NumRepos); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if datasetIDs, err = opts.RequestedUser.GetOrgDatasetIds4User(actorID); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if aimodelIDs, err = opts.RequestedUser.GetOrgAimodelIds4User(actorID); err != nil {
			return nil, fmt.Errorf("GetOrgAimodelIds4User: %v", err)
		}

		cond = cond.And(builder.Or(builder.In("repo_id", repoIDs), builder.In("dataset_id", datasetIDs), builder.In("aimodel_id", aimodelIDs)))
	} else {

		cond = cond.And(builder.Or(builder.In("op_type", AI_TASK_ACTION_LIST), builder.In("repo_id", AccessibleRepoIDsQuery(opts.Actor)), builder.In("dataset_id", AccessibleDatasetIDsQuery(opts.Actor)), builder.In("aimodel_id", AccessibleAimodelIDsQuery(opts.Actor))))
	}

	cond = cond.And(builder.Eq{"user_id": opts.RequestedUser.ID})

	if opts.OnlyPerformedBy {
		cond = cond.And(builder.Eq{"act_user_id": opts.RequestedUser.ID})
	}
	if !opts.IncludePrivate {
		cond = cond.And(builder.Eq{"is_private": false})
	}

	if !opts.IncludeDeleted {
		cond = cond.And(builder.Eq{"is_deleted": false})
	}

	actions := make([]*Action, 0, 20)

	if err := x.Limit(20).Desc("id").Where(cond).Find(&actions); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}

	if err := ActionList(actions).LoadAllAttributes(); err != nil {
		return nil, fmt.Errorf("LoadAttributes: %v", err)
	}
	return actions, nil
}

func GetFeedsV2(opts GetFeedsOptions) ([]*Action, error) {
	cond := builder.NewCond()

	var actorID int64
	if opts.Actor != nil {
		actorID = opts.Actor.ID
	}

	var repoIDs []int64
	var datasetIDs []string
	var aimodelIDs []string
	if opts.RequestedUser.IsOrganization() {
		env, err := opts.RequestedUser.AccessibleReposEnv(actorID)
		if err != nil {
			return nil, fmt.Errorf("AccessibleReposEnv: %v", err)
		}
		if repoIDs, err = env.RepoIDs(1, opts.RequestedUser.NumRepos); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if datasetIDs, err = opts.RequestedUser.GetOrgDatasetIds4User(actorID); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if aimodelIDs, err = opts.RequestedUser.GetOrgAimodelIds4User(actorID); err != nil {
			return nil, fmt.Errorf("GetOrgAimodelIds4User: %v", err)
		}
	} else {
		var err error
		if repoIDs, err = GetAccessibleRepoIDs(opts.Actor); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if datasetIDs, err = GetAccessibleDatasetIDs(opts.Actor); err != nil {
			return nil, fmt.Errorf("GetUserRepositories: %v", err)
		}
		if aimodelIDs, err = GetAccessibleAimodelIDs(opts.Actor); err != nil {
			return nil, fmt.Errorf("GetOrgAimodelIds4User: %v", err)
		}

	}

	repoSet := make(map[int64]struct{}, len(repoIDs))
	for _, id := range repoIDs {
		repoSet[id] = struct{}{}
	}
	datasetSet := make(map[string]struct{}, len(datasetIDs))
	for _, id := range datasetIDs {
		datasetSet[id] = struct{}{}
	}
	aimodelSet := make(map[string]struct{}, len(aimodelIDs))
	for _, id := range aimodelIDs {
		aimodelSet[id] = struct{}{}
	}
	actTypeSet := make(map[ActionType]struct{}, len(AI_TASK_ACTION_LIST))
	for _, at := range AI_TASK_ACTION_LIST {
		actTypeSet[at] = struct{}{}
	}

	cond = cond.And(builder.Eq{"user_id": opts.RequestedUser.ID})
	if opts.OnlyPerformedBy {
		cond = cond.And(builder.Eq{"act_user_id": opts.RequestedUser.ID})
	}
	if !opts.IncludePrivate {
		cond = cond.And(builder.Eq{"is_private": false})
	}
	if !opts.IncludeDeleted {
		cond = cond.And(builder.Eq{"is_deleted": false})
	}

	actions := make([]*Action, 0, 200)
	if err := x.
		Limit(200).
		Desc("id").
		Where(cond).
		Find(&actions); err != nil {
		return nil, fmt.Errorf("Find actions: %w", err)
	}

	result := make([]*Action, 0, 20)
	for _, act := range actions {
		if len(result) >= 20 {
			break
		}

		match := false
		if act.RepoID > 0 {
			_, match = repoSet[act.RepoID]
		}
		if !match && act.DatasetID != nil {
			_, match = datasetSet[*act.DatasetID]
		}
		if !match && act.AimodelID != nil {
			_, match = aimodelSet[*act.AimodelID]
		}
		if !match {
			_, match = actTypeSet[act.OpType]
		}
		if match {
			result = append(result, act)
		}
	}

	if err := ActionList(result).LoadAllAttributes(); err != nil {
		return nil, fmt.Errorf("LoadAttributes: %v", err)
	}
	return result, nil
}

func GetLast20PublicFeeds(opTypes []int) ([]*Action, error) {
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"is_private": false})
	cond = cond.And(builder.Eq{"is_deleted": false})
	cond = cond.And(builder.Expr("user_id=act_user_id"))
	cond = cond.And(builder.In("op_type", opTypes))

	actions := make([]*Action, 0, 20)

	if err := x.Limit(30).Desc("id").Where(cond).Find(&actions); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}

	if err := ActionList(actions).LoadAllAttributes(); err != nil {
		return nil, fmt.Errorf("LoadAttributes: %v", err)
	}

	return actions, nil
}

func GetUnTransformedActions() ([]*Action, error) {
	actions := make([]*Action, 0, 10)
	err := x.Where("op_type = ?", ActionCommitRepo).
		And("content != ''").
		And("is_transformed = ?", false).
		And("to_timestamp(created_unix) >= ?", setting.CommitValidDate).
		Find(&actions)
	return actions, err
}

func GetActionByIds(ids []int64) ([]*Action, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	actions := make([]*Action, 0)
	err := x.In("id", ids).Find(&actions)
	if err != nil {
		return nil, err
	}
	if err := ActionList(actions).LoadAllAttributes(); err != nil {
		return nil, fmt.Errorf("ActionList loadAttributes: %v", err)
	}
	return actions, nil
}

func GetActionById(id int64) (*Action, error) {
	if id == 0 {
		return nil, ErrRecordNotExist{}
	}
	action := &Action{}
	has, err := x.ID(id).Get(action)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return action, nil
}
