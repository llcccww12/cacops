package models

import (
	"encoding/json"
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type UserBusinessAnalysisForActivity struct {
	ID        int64 `xorm:"pk"`
	CountDate int64 `xorm:"pk"`
	//action :ActionMergePullRequest                         // 11
	CodeMergeCount int `xorm:"NOT NULL DEFAULT 0"`
	//action :ActionCommitRepo
	CommitCount int `xorm:"NOT NULL DEFAULT 0"`
	//issue                              // 10
	IssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//comment table current date
	CommentCount int `xorm:"NOT NULL DEFAULT 0"`

	//follow table
	WatchedCount int `xorm:"NOT NULL DEFAULT 0"`

	CommitCodeSize int `xorm:"NOT NULL DEFAULT 0"`
	//issue, issueassigees
	SolveIssueCount int `xorm:"NOT NULL DEFAULT 0"`
	//use
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`

	//user
	Email      string `xorm:"NOT NULL"`
	Occupation int64  `xorm:"NOT NULL DEFAULT 0"` // 0空，1其他，2学生，3科研人员，4企业人员

	Phone string `xorm:"NULL"`
	//user
	Name     string `xorm:"NOT NULL"`
	DataDate string `xorm:"NULL"`

	//TODO
	CloudBrainTaskNum int `xorm:"NOT NULL DEFAULT 0"`
	CommitDatasetNum  int `xorm:"NOT NULL DEFAULT 0"`
	//0
	CommitModelCount int `xorm:"NOT NULL DEFAULT 0"`

	//login  TODO
	LoginCount int `xorm:"NOT NULL DEFAULT 0"`

	//Create repo nums TODO
	CreateRepoCount int `xorm:"NOT NULL DEFAULT 0"`

	//model convert TODO
	ModelConvertCount int `xorm:"NOT NULL DEFAULT 0"`

	//Collect Dataset nums
	CollectDataset int `xorm:"NOT NULL DEFAULT 0"`
}

func QueryDataForActivity(startTime time.Time, endTime time.Time) []*UserBusinessAnalysisForActivity {
	sess := x.NewSession()
	defer sess.Close()

	result := make([]*UserBusinessAnalysisForActivity, 0)
	publicRepo, publicUserRepo := queryPublicRepo()
	start_unix := startTime.Unix()
	end_unix := endTime.Unix()

	CodeMergeCountMap := queryPullRequestPublic(start_unix, end_unix, publicRepo)
	CommitCodeSizeMap, err := GetAllUserPublicRepoKPIStats(startTime, endTime)
	if err != nil {
		log.Info("error,info=" + err.Error())
	}
	CommitCountMap := queryCommitActionPublic(start_unix, end_unix, 5, publicRepo)
	IssueCountMap, publicRepoIssueIdMap := queryCreateIssuePublic(start_unix, end_unix, publicRepo)
	SolveIssueCountMap := querySolveIssuePublic(start_unix, end_unix, publicRepoIssueIdMap)
	WatchedCountMap, _ := queryFollow(start_unix, end_unix)
	CommentCountMap := queryCommentPublic(start_unix, end_unix, publicRepoIssueIdMap)
	PublicDataSet := queryAllPublicDataSet(publicRepo)
	DatasetFileNums := queryPublicDatasetFileNums(start_unix, end_unix, PublicDataSet)
	AiModelManageMap := queryUserModelPublic(start_unix, end_unix, publicRepo)
	LoginCountMap, _ := queryLoginCount(start_unix, end_unix)
	CloudBrainTaskMap := queryCloudBrainTaskPublic(start_unix, end_unix, publicRepo)
	CollectDataset, _ := queryDatasetStars(start_unix, end_unix)
	AiModelConvertMap := queryUserModelConvertPublic(start_unix, end_unix, publicRepo)

	cond := "type != 1 and is_active=true"
	count, err := sess.Where(cond).Count(new(User))

	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("`user`.*").Table("user").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userList := make([]*User, 0)
		sess.Find(&userList)

		for i, userRecord := range userList {
			var dateRecord UserBusinessAnalysisForActivity
			dateRecord.ID = userRecord.ID
			log.Info("i=" + fmt.Sprint(i) + " userName=" + userRecord.Name)
			dateRecord.Email = userRecord.Email
			dateRecord.Occupation = userRecord.Occupation
			dateRecord.Phone = userRecord.PhoneNumber
			dateRecord.RegistDate = userRecord.CreatedUnix
			dateRecord.Name = userRecord.Name

			dateRecord.CodeMergeCount = getMapValue(dateRecord.ID, CodeMergeCountMap)
			dateRecord.CommitCount = getMapValue(dateRecord.ID, CommitCountMap)
			dateRecord.WatchedCount = getMapValue(dateRecord.ID, WatchedCountMap)
			dateRecord.CommitDatasetNum = getMapValue(dateRecord.ID, DatasetFileNums)
			dateRecord.IssueCount = getMapValue(dateRecord.ID, IssueCountMap)
			dateRecord.CommentCount = getMapValue(dateRecord.ID, CommentCountMap)
			dateRecord.LoginCount = getMapValue(dateRecord.ID, LoginCountMap)

			if _, ok := CommitCodeSizeMap[dateRecord.Email]; !ok {
				dateRecord.CommitCodeSize = 0
			} else {
				dateRecord.CommitCodeSize = int(CommitCodeSizeMap[dateRecord.Email].CommitLines)
			}
			dateRecord.SolveIssueCount = getMapValue(dateRecord.ID, SolveIssueCountMap)

			dateRecord.CommitModelCount = getMapValue(dateRecord.ID, AiModelManageMap)
			dateRecord.CloudBrainTaskNum = getMapValue(dateRecord.ID, CloudBrainTaskMap)

			dateRecord.ModelConvertCount = getMapValue(dateRecord.ID, AiModelConvertMap)
			dateRecord.CreateRepoCount = getMapValue(dateRecord.ID, publicUserRepo)
			dateRecord.CollectDataset = getMapValue(dateRecord.ID, CollectDataset)

			result = append(result, &dateRecord)
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return result
}

func queryUserModelConvertPublic(start_unix int64, end_unix int64, publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(AiModelConvert))
	if err != nil {
		log.Info("query AiModelConvert error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,repo_id").Table("ai_model_convert").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		aiModelList := make([]*AiModelConvert, 0)
		sess.Find(&aiModelList)
		log.Info("query AiModelConvert size=" + fmt.Sprint(len(aiModelList)))
		for _, aiModelRecord := range aiModelList {
			if isPublicRepo(aiModelRecord.RepoId, publicAllRepo) {
				if _, ok := resultMap[aiModelRecord.UserId]; !ok {
					resultMap[aiModelRecord.UserId] = 1
				} else {
					resultMap[aiModelRecord.UserId] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryCloudBrainTaskPublic(start_unix int64, end_unix int64, publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Unscoped().Count(new(Cloudbrain))
	if err != nil {
		log.Info("query cloudbrain error. return.")
		return resultMap
	}
	log.Info("cloudbrain count=" + fmt.Sprint(count))
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,repo_id,job_type,user_id,duration,train_job_duration,type,compute_resource").Table("cloudbrain").Unscoped().Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		cloudTaskList := make([]*Cloudbrain, 0)
		sess.Find(&cloudTaskList)
		log.Info("query cloudbrain size=" + fmt.Sprint(len(cloudTaskList)))
		for _, cloudTaskRecord := range cloudTaskList {
			if isPublicRepo(cloudTaskRecord.RepoID, publicAllRepo) {
				if _, ok := resultMap[cloudTaskRecord.UserID]; !ok {
					resultMap[cloudTaskRecord.UserID] = 1
				} else {
					resultMap[cloudTaskRecord.UserID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func querySolveIssuePublic(start_unix int64, end_unix int64, publicRepoIssueIdMap map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := "issue.is_closed=true and issue.closed_unix>=" + fmt.Sprint(start_unix) + " and issue.closed_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Table("issue_assignees").Join("inner", "issue", "issue.id=issue_assignees.issue_id").Where(cond).Count(new(IssueAssignees))
	if err != nil {
		log.Info("query issue error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		issueAssigneesList := make([]*IssueAssignees, 0)
		sess.Select("issue_assignees.*").Table("issue_assignees").
			Join("inner", "issue", "issue.id=issue_assignees.issue_id").
			Where(cond).OrderBy("issue_assignees.id asc").Limit(PAGE_SIZE, int(indexTotal))

		sess.Find(&issueAssigneesList)

		log.Info("query IssueAssignees size=" + fmt.Sprint(len(issueAssigneesList)))
		for _, issueAssigneesRecord := range issueAssigneesList {
			if isPublicRepo(issueAssigneesRecord.IssueID, publicRepoIssueIdMap) {
				if _, ok := resultMap[issueAssigneesRecord.AssigneeID]; !ok {
					resultMap[issueAssigneesRecord.AssigneeID] = 1
				} else {
					resultMap[issueAssigneesRecord.AssigneeID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap
}

func queryPublicRepo() (map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	resultUserRepoMap := make(map[int64]int)
	count, err := sess.Table("repository").Count(new(Repository))
	if err != nil {
		log.Info("query Repository error. return.")
		return resultMap, resultUserRepoMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		repositoryList := make([]*Repository, 0)
		sess.Select("*").Table("repository").OrderBy("id desc").Limit(PAGE_SIZE, int(indexTotal))
		sess.Find(&repositoryList)
		log.Info("query repo size=" + fmt.Sprint(len(repositoryList)))
		for _, repositoryRecord := range repositoryList {
			if repositoryRecord.IsPrivate {
				continue
			}
			if _, ok := resultMap[repositoryRecord.ID]; !ok {
				resultMap[repositoryRecord.ID] = 1
			}
			if _, ok := resultUserRepoMap[repositoryRecord.OwnerID]; !ok {
				resultUserRepoMap[repositoryRecord.OwnerID] = 1
			} else {
				resultUserRepoMap[repositoryRecord.OwnerID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap, resultUserRepoMap
}

func isPublicRepo(repoId int64, publicAllRepo map[int64]int) bool {
	if _, ok := publicAllRepo[repoId]; !ok {
		return false
	}
	return true
}

func queryPullRequestPublic(start_unix int64, end_unix int64, publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := "issue.created_unix>=" + fmt.Sprint(start_unix) + " and issue.created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Table("issue").Join("inner", "pull_request", "issue.id=pull_request.issue_id").Where(cond).Count(new(Issue))
	if err != nil {
		log.Info("query issue error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		issueList := make([]*Issue, 0)
		sess.Select("issue.*").Table("issue").Join("inner", "pull_request", "issue.id=pull_request.issue_id").Where(cond).OrderBy("issue.id asc").Limit(PAGE_SIZE, int(indexTotal))
		sess.Find(&issueList)
		log.Info("query issue(PR) size=" + fmt.Sprint(len(issueList)))
		for _, issueRecord := range issueList {
			if isPublicRepo(issueRecord.RepoID, publicAllRepo) {
				if _, ok := resultMap[issueRecord.PosterID]; !ok {
					resultMap[issueRecord.PosterID] = 1
				} else {
					resultMap[issueRecord.PosterID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryCommitActionPublic(start_unix int64, end_unix int64, actionType int64, publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)

	cond := "user_id=act_user_id and op_type=" + fmt.Sprint(actionType) + " and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Action))
	if err != nil {
		log.Info("query action error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,op_type,act_user_id,repo_id").Table("action").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		actionList := make([]*Action, 0)
		sess.Find(&actionList)

		log.Info("query action size=" + fmt.Sprint(len(actionList)))
		for _, actionRecord := range actionList {
			if isPublicRepo(actionRecord.RepoID, publicAllRepo) {
				if _, ok := resultMap[actionRecord.UserID]; !ok {
					resultMap[actionRecord.UserID] = 1
				} else {
					resultMap[actionRecord.UserID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap
}

func queryCreateIssuePublic(start_unix int64, end_unix int64, publicAllRepo map[int64]int) (map[int64]int, map[int64]int) {

	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	publicRepoIssueIdMap := make(map[int64]int)
	cond := "is_pull=false and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Issue))
	if err != nil {
		log.Info("query Issue error. return.")
		return resultMap, publicRepoIssueIdMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,poster_id,repo_id").Table("issue").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		issueList := make([]*Issue, 0)
		sess.Find(&issueList)
		log.Info("query issue size=" + fmt.Sprint(len(issueList)))
		for _, issueRecord := range issueList {
			if isPublicRepo(issueRecord.RepoID, publicAllRepo) {
				if _, ok := resultMap[issueRecord.PosterID]; !ok {
					resultMap[issueRecord.PosterID] = 1
				} else {
					resultMap[issueRecord.PosterID] += 1
				}
				publicRepoIssueIdMap[issueRecord.ID] = 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap, publicRepoIssueIdMap

}

func queryCommentPublic(start_unix int64, end_unix int64, publicRepoIssueIdMap map[int64]int) map[int64]int {

	sess := x.NewSession()
	defer sess.Close()
	cond := "created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	resultMap := make(map[int64]int)
	count, err := sess.Where(cond).Count(new(Comment))
	if err != nil {
		log.Info("query Comment error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,type,poster_id,issue_id").Table("comment").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		commentList := make([]*Comment, 0)
		sess.Find(&commentList)
		log.Info("query Comment size=" + fmt.Sprint(len(commentList)))
		for _, commentRecord := range commentList {
			if isPublicRepo(commentRecord.IssueID, publicRepoIssueIdMap) {
				if _, ok := resultMap[commentRecord.PosterID]; !ok {
					resultMap[commentRecord.PosterID] = 1
				} else {
					resultMap[commentRecord.PosterID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryAllPublicDataSet(publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	publicDataSetIdMap := make(map[int64]int)
	count, err := sess.Count(new(Dataset))
	if err != nil {
		log.Info("query dataset error. return.")
		return publicDataSetIdMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,repo_id").Table(new(Dataset)).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		datasetList := make([]*Dataset, 0)
		sess.Find(&datasetList)
		log.Info("query datasetList size=" + fmt.Sprint(len(datasetList)))
		for _, datasetRecord := range datasetList {
			if isPublicRepo(datasetRecord.RepoID, publicAllRepo) {
				publicDataSetIdMap[datasetRecord.ID] = 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return publicDataSetIdMap
}

func queryPublicDatasetFileNums(start_unix int64, end_unix int64, publicDataSetIdMap map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultNumMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Attachment))
	if err != nil {
		log.Info("query attachment error. return.")
		return resultNumMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uploader_id,size,dataset_id").Table("attachment").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		attachmentList := make([]*Attachment, 0)
		sess.Find(&attachmentList)

		log.Info("query Attachment size=" + fmt.Sprint(len(attachmentList)))
		for _, attachRecord := range attachmentList {
			if isPublicRepo(attachRecord.DatasetID, publicDataSetIdMap) {
				if _, ok := resultNumMap[attachRecord.UploaderID]; !ok {
					resultNumMap[attachRecord.UploaderID] = 1
				} else {
					resultNumMap[attachRecord.UploaderID] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultNumMap
}

func queryUserModelPublic(start_unix int64, end_unix int64, publicAllRepo map[int64]int) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(AiModelManage))
	if err != nil {
		log.Info("query AiModelManage error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,repo_id").Table("ai_model_manage").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		aiModelList := make([]*AiModelManage, 0)
		sess.Find(&aiModelList)
		log.Info("query AiModelManage size=" + fmt.Sprint(len(aiModelList)))
		for _, aiModelRecord := range aiModelList {
			if isPublicRepo(aiModelRecord.RepoId, publicAllRepo) {
				if _, ok := resultMap[aiModelRecord.UserId]; !ok {
					resultMap[aiModelRecord.UserId] = 1
				} else {
					resultMap[aiModelRecord.UserId] += 1
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func QueryUserLoginInfo(userIds []int64) []*UserLoginLog {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	var cond = builder.NewCond()
	cond = cond.And(builder.In("u_id", userIds))
	statictisSess.Select("*").Table(new(UserLoginLog)).Where(cond)
	loginList := make([]*UserLoginLog, 0)

	statictisSess.Find(&loginList)

	return loginList
}

var WeekBonusData = make(map[int64][]int)

func QueryUserAnnualReport(userId int64) *UserSummaryCurrentYear {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	log.Info("userId=" + fmt.Sprint(userId))
	if len(WeekBonusData) == 0 {
		//WeekBonusData = getBonusWeekDataMap()
	}

	reList := make([]*UserSummaryCurrentYear, 0)
	err := statictisSess.Select("*").Table(new(UserSummaryCurrentYear)).Where("id=" + fmt.Sprint(userId)).Find(&reList)
	if err == nil {
		if len(reList) > 0 {
			record, ok := WeekBonusData[userId]
			if ok {
				bonusInfo := make(map[string]int)
				bonusInfo["order"] = record[0]
				bonusInfo["money"] = record[1]
				bonusInfo["week"] = record[2]
				bonusInfo["num"] = record[3]
				bonusInfoJson, _ := json.Marshal(bonusInfo)
				reList[0].WeekBonusData = string(bonusInfoJson)
			}
			return reList[0]
		}
	} else {
		log.Info("error:=" + err.Error())
	}
	dbuser, err := GetUserByID(userId)
	if err == nil {
		return &UserSummaryCurrentYear{
			ID:         dbuser.ID,
			Name:       dbuser.Name,
			RegistDate: dbuser.CreatedUnix,
		}
	}
	return nil
}

func GetLastModifyTime() string {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	reList := make([]*UserBusinessAnalysisLastMonth, 0)
	err := statictisSess.Select("*").Table(new(UserBusinessAnalysisLastMonth)).Limit(1, 0).Find(&reList)
	if err == nil {
		return reList[0].DataDate
	}
	return ""
}
