package models

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

type UserSummaryCurrentYear struct {
	ID    int64  `xorm:"pk"`
	Email string `xorm:"NOT NULL"`
	//user
	Name  string `xorm:"NOT NULL"`
	Phone string `xorm:"NULL"`
	//user
	RegistDate timeutil.TimeStamp `xorm:"NOT NULL"`

	DateCount     int    `xorm:"NOT NULL DEFAULT 0"`
	MostActiveDay string `xorm:"varchar(1000)"` //08.05
	RepoInfo      string `xorm:"varchar(1000)"` //创建了XX 个项目，公开项目XX 个，私有项目XX 个累计被下载XXX 次，其中《XXXXXXX 》项目，获得了最高XXX 次下载
	//2023：增加春、夏、秋、冬四季创建的项目信息。

	DataSetInfo string `xorm:"varchar(500)"` //创建了XX 个数据集，上传了XX 个数据集文件，累计被下载XX 次，被收藏XX 次
	CodeInfo    string `xorm:"varchar(500)"` //代码提交次数，提交总代码行数，最晚的提交时间
	//2023：所有人的代码行要进行排序，并设置每个人的排位

	CloudBrainInfo string `xorm:"varchar(6000)"` //，创建了XX 个云脑任务，调试任务XX 个，训练任务XX 个，推理任务XX 个，累计运行了XXXX 卡时，累计节省xxxxx 元
	//这些免费的算力资源分别有，XX% 来自鹏城云脑1，XX% 来自鹏城云脑2，XX% 来自智算网络
	//2023：有XX 个使用启智集群资源,有XX 个使用智算网络集群,使用过的计算资源有GPU NPU GCU
	//2023：你的所有任务累计运行了XXX卡时,其中 GPU资源运行XX卡时 NPU资源运行XX卡时  GCU资源运行XX卡时

	PlayARoll           string `xorm:"varchar(500)"` //你参加了XX 次“我为开源打榜狂”活动，累计上榜XX 次，总共获得了社区XXX 元的激励
	OpenSource          string `xorm:"varchar(500)"` //开源实习活动
	WeekBonusData       string `xorm:"-"`
	Label               string `xorm:"varchar(500)"`
	IssueInfo           string `xorm:"varchar(500)"`       //2023：ISSUE相关信息，包括创建issue，关闭issue,评论issue
	ModelInfo           string `xorm:"varchar(500)"`       //2023：模型相关信息，包括创建的模型个数，要给出下载次数最多、引用次数最多的模型
	LoginCount          int    `xorm:"NOT NULL DEFAULT 0"` //2023：当年的登录次数
	ActionCount         int    `xorm:"NOT NULL DEFAULT 0"` //2025：当年的开发行为次数（action）
	LoginInfo           string `xorm:"varchar(500)"`       //2024：登录最频繁的月份，及登录次数 //2025：location 登录次数超过了百分之多少的用户
	ActionInfo          string `xorm:"varchar(500)"`       //2023：最勤奋的一个月，取个人活动页面最多的一个月，还要给出该月活动最多的一天。
	AccumulatePointInfo string `xorm:"varchar(500)"`       //2023：获取的积分总数及消耗的总数，还剩余多少积分。
	ForumInfo           string `xorm:"varchar(500)"`       //2023: 你在OpenI论坛发表了XX篇贴  总共获得了XX次浏   XX次点赞  其中《XXXXX》 获得了最高XX次浏览
	CourseInfo          string `xorm:"varchar(500)"`       //2023:你在OpenI学习了XX门课程  总共完成了XX章节实训   40%课程为机器学习，列出4个最多的分类
	ActionDays          string `xorm:"varchar(500)"`       //2023:你在OpenI中有产生活动的天数。
	DatasetStarCount    int    `xorm:"NOT NULL DEFAULT 0"`
	ModelStarCount      int    `xorm:"NOT NULL DEFAULT 0"`
}

func RefreshUserYearTable(pageStartTime time.Time, pageEndTime time.Time) {
	sess := x.NewSession()
	defer sess.Close()
	log.Info("RefreshUserYearTable start....")
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	log.Info("UserYear StartTime:" + pageStartTime.Format("2006-01-02 15:04:05"))
	log.Info("UserYear EndTime time:" + pageEndTime.Format("2006-01-02 15:04:05"))

	start_unix := pageStartTime.Unix()
	end_unix := pageEndTime.Unix()

	CodeMergeCountMap := queryPullRequest(start_unix, end_unix)
	CommitCountMap := queryCommitAction(start_unix, end_unix, 5)
	actionCountMap, mostActiveMap, mostActiveMonthMap, mostActivePeriodMap, mostActiveStartEndMap, mostActiveStartEndDayMap, mostActiveSeazonMap, actionTaskTypeMap := queryMostActiveCommitAction(start_unix, end_unix)
	IssueCountMap := queryCreateIssue(start_unix, end_unix)
	UserYearModel := queryUserYearModel(start_unix, end_unix)
	CommentCountMap := queryComment(start_unix, end_unix)
	LoginMap, UserMonthLoginMap, loginValueArray := queryLoginCountForYear(start_unix, end_unix)

	commitCodeSizeMap, err := GetAllUserKPIStats(pageStartTime, pageEndTime)

	commitCodeArrays := make([]int, 0)
	if err == nil {
		for _, v := range commitCodeSizeMap {
			commitCodeArrays = append(commitCodeArrays, int(v.CommitLines))
		}
	}

	sort.Ints(commitCodeArrays)
	//log.Info("query commit code size, len=" + fmt.Sprint(len(existCommitCodeSize)))

	//CommitDatasetSizeMap, CommitDatasetNumMap, dataSetDownloadMap := queryDatasetSize(start_unix, end_unix)
	SolveIssueCountMap := querySolveIssue(start_unix, end_unix)
	CreateRepoCountMap, DetailInfoMap, MostDownloadMap, fourSeasonMap := queryUserCreateRepo(start_unix, end_unix)

	CloudBrainTaskMap, CloudBrainTaskItemMap, resourceItemMap, jobTypeItemMap, cardTypeItemMap, maxDurationJobTypeMap, maxDurationMap, aiCenterMap, modelMap := queryCloudBrainTaskForYear(start_unix, end_unix)

	CollectedDataset := queryDatasetStarsForYear(start_unix, end_unix)
	CollectedModel := queryAIModelStarsForYear(start_unix, end_unix)
	_, CreatedDataset, userIdMostDownloadDatasetMap, userIdMostDownloadDatasetNameMap, userIdMostQuoteDatasetMap, userIdMostQuoteDatasetNameMap, privateDatasetMap := queryRecommedDataSet(start_unix, end_unix)

	PointMap := queryPointInfo(start_unix, end_unix)

	//MostActiveDayMap := queryMostActiveMonth(start_unix, end_unix)

	//bonusMap := getBonusMap()
	//forumMap := getForumMap()
	//courseMap := getCourseMap()
	//openSourceMap := getOpenSourceMap()
	//actionDays := queryActionDays(start_unix, end_unix)
	log.Info("truncate all data from table:user_summary_current_year ")
	statictisSess.Exec("TRUNCATE TABLE user_summary_current_year")

	cond := "type != 1"
	count, err := sess.Where(cond).Count(new(User))
	if err != nil {
		log.Info("query user error. return.")
		return
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("`user`.*").Table("user").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userList := make([]*User, 0)
		sess.Find(&userList)
		userCount := len(userList)
		for _, userRecord := range userList {
			var dateRecordAll UserBusinessAnalysisAll
			dateRecordAll.ID = userRecord.ID
			dateRecordAll.Email = userRecord.Email
			dateRecordAll.Phone = userRecord.PhoneNumber
			dateRecordAll.RegistDate = userRecord.CreatedUnix
			dateRecordAll.Name = userRecord.Name

			dateRecordAll.CodeMergeCount = getMapValue(dateRecordAll.ID, CodeMergeCountMap)
			dateRecordAll.CommitCount = getMapValue(dateRecordAll.ID, CommitCountMap)
			dateRecordAll.IssueCount = getMapValue(dateRecordAll.ID, IssueCountMap)
			dateRecordAll.CommentCount = getMapValue(dateRecordAll.ID, CommentCountMap)

			dateRecordAll.CommitCodeSize = 0
			if _, ok := commitCodeSizeMap[dateRecordAll.Email]; ok {
				dateRecordAll.CommitCodeSize = int(commitCodeSizeMap[dateRecordAll.Email].CommitLines)
			}

			dateRecordAll.SolveIssueCount = getMapValue(dateRecordAll.ID, SolveIssueCountMap)
			dateRecordAll.CreateRepoCount = getMapValue(dateRecordAll.ID, CreateRepoCountMap)
			dateRecordAll.LoginCount = getMapValue(dateRecordAll.ID, LoginMap)
			dateRecordAll.CloudBrainTaskNum = getMapValue(dateRecordAll.ID, CloudBrainTaskMap)
			//dateRecordAll.GpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuDebugJob", CloudBrainTaskItemMap)
			//dateRecordAll.NpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuDebugJob", CloudBrainTaskItemMap)
			//dateRecordAll.GpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuTrainJob", CloudBrainTaskItemMap)
			//dateRecordAll.NpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuTrainJob", CloudBrainTaskItemMap)
			//dateRecordAll.NpuInferenceJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuInferenceJob", CloudBrainTaskItemMap)
			//dateRecordAll.GpuBenchMarkJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuBenchMarkJob", CloudBrainTaskItemMap)
			dateRecordAll.CloudBrainRunTime = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_CloudBrainRunTime", CloudBrainTaskItemMap)

			//年度数据
			subTime := time.Now().UTC().Sub(dateRecordAll.RegistDate.AsTime().UTC())
			mostActiveDay := ""
			if userInfo, ok := mostActiveMap[dateRecordAll.ID]; ok {
				userMonthInfo, _ := mostActiveMonthMap[dateRecordAll.ID]
				userPeriodInfo, _ := mostActivePeriodMap[dateRecordAll.ID]
				userStartEndInfo, _ := mostActiveStartEndMap[dateRecordAll.ID]
				userStartEndDayInfo, _ := mostActiveStartEndDayMap[dateRecordAll.ID]
				userSeazonInfo, _ := mostActiveSeazonMap[dateRecordAll.ID]
				userTaskTypeInfo, _ := actionTaskTypeMap[dateRecordAll.ID]

				mostActiveDay = getMostActiveJson(userInfo, userMonthInfo, userPeriodInfo, userStartEndInfo, userStartEndDayInfo, userSeazonInfo, userTaskTypeInfo)
			}
			mostLoginMonth := ""

			if userMonthLoginInfo, ok := UserMonthLoginMap[dateRecordAll.ID]; ok {

				mostLoginMonth = getMostLoginMonthJson(userMonthLoginInfo, dateRecordAll.LoginCount, loginValueArray, userCount)
			}
			scoreMap := make(map[string]float64)
			repoInfo := getRepoDetailInfo(DetailInfoMap, dateRecordAll.ID, MostDownloadMap, fourSeasonMap)
			dataSetInfo, datasetscore := getDataSetInfo(dateRecordAll.ID, CreatedDataset, userIdMostDownloadDatasetMap, userIdMostDownloadDatasetNameMap, userIdMostQuoteDatasetMap, userIdMostQuoteDatasetNameMap, privateDatasetMap)
			scoreMap["datasetscore"] = datasetscore
			codeInfo, codescore := getCodeInfo(&dateRecordAll, commitCodeArrays)
			scoreMap["codescore"] = codescore
			cloudBrainInfo := getCloudBrainInfo(&dateRecordAll, jobTypeItemMap[dateRecordAll.ID], scoreMap, resourceItemMap[dateRecordAll.ID], cardTypeItemMap[dateRecordAll.ID], maxDurationJobTypeMap[dateRecordAll.ID], maxDurationMap[dateRecordAll.ID], aiCenterMap[dateRecordAll.ID], modelMap[dateRecordAll.ID])
			//playARoll := getPlayARoll(bonusMap, dateRecordAll.Name)
			//course := getCourseInfo(courseMap, dateRecordAll.ID)
			//openSource := getPlayARoll(openSourceMap, dateRecordAll.Name)
			//forumInfo := getForumInfo(forumMap, dateRecordAll.Name)
			exteral := 0
			if int(subTime.Hours())%24 > 0 {
				exteral = 1
			}
			re := &UserSummaryCurrentYear{
				ID:                  dateRecordAll.ID,
				Name:                dateRecordAll.Name,
				Email:               dateRecordAll.Email,
				Phone:               dateRecordAll.Phone,
				RegistDate:          dateRecordAll.RegistDate,
				DateCount:           int(subTime.Hours())/24 + exteral,
				MostActiveDay:       mostActiveDay,
				RepoInfo:            repoInfo,
				DataSetInfo:         dataSetInfo,
				CodeInfo:            codeInfo,
				CloudBrainInfo:      cloudBrainInfo,
				PlayARoll:           "", //playARoll,
				IssueInfo:           getIssueInfo(&dateRecordAll),
				ModelInfo:           getModelInfo(&dateRecordAll, UserYearModel),
				LoginCount:          dateRecordAll.LoginCount,
				ActionCount:         actionCountMap[dateRecordAll.ID],
				LoginInfo:           mostLoginMonth,
				AccumulatePointInfo: getPointInfo(&dateRecordAll, PointMap),
				ActionInfo:          "", //getActionInfo(&dateRecordAll, MostActiveDayMap),
				CourseInfo:          "", //course,
				OpenSource:          "", //openSource,
				ForumInfo:           "", // no need forumInfo in 2024
				ActionDays:          "", //getActionDaysInfo(actionDays[dateRecordAll.ID]), no need 2025
				DatasetStarCount:    CollectedDataset[dateRecordAll.ID],
				ModelStarCount:      CollectedModel[dateRecordAll.ID],
			}
			_, err := statictisSess.Insert(re)
			if err != nil {
				log.Error("insert user year data error. ", err)
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	log.Info("update user year data finished. ")
}

func getAllLoginCount(loginValueArray []int) int {
	sum := 0
	for _, loginValue := range loginValueArray {
		sum += loginValue
	}
	return sum
}

func queryDatasetStarsForYear(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	datasetCollect := make(map[int64]int)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(DatasetCollection))
	if err != nil {
		log.Info("query follow error. return.")
		return datasetCollect
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,dataset_id").Table(new(DatasetCollection)).Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		datasetStarList := make([]*DatasetCollection, 0)
		sess.Find(&datasetStarList)
		log.Info("query datasetStarList size=" + fmt.Sprint(len(datasetStarList)))
		for _, datasetStarRecord := range datasetStarList {
			if _, ok := datasetCollect[datasetStarRecord.UserID]; !ok {
				datasetCollect[datasetStarRecord.UserID] = 1
			} else {
				datasetCollect[datasetStarRecord.UserID] += 1
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return datasetCollect
}

func queryAIModelStarsForYear(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	modelCollect := make(map[int64]int)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(AiModelCollect))
	if err != nil {
		log.Info("query follow error. return.")
		return modelCollect
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,model_id").Table(new(AiModelCollect)).Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		modelStarList := make([]*AiModelCollect, 0)
		sess.Find(&modelStarList)
		log.Info("query modelStarList size=" + fmt.Sprint(len(modelStarList)))
		for _, modelStarRecord := range modelStarList {
			if _, ok := modelCollect[modelStarRecord.UserId]; !ok {
				modelCollect[modelStarRecord.UserId] = 1
			} else {
				modelCollect[modelStarRecord.UserId] += 1
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return modelCollect
}

func queryCloudBrainTaskForYear(start_unix int64, end_unix int64) (map[int64]int, map[string]int, map[int64]map[string]int, map[int64]map[string]int, map[int64]map[string]int, map[int64]string, map[int64]int, map[int64]map[string]int, map[int64]map[string]int) {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	resultItemMap := make(map[string]int)
	resourceItemMap := make(map[int64]map[string]int)
	jobTypeItemMap := make(map[int64]map[string]int)
	cardTypeItemMap := make(map[int64]map[string]int)
	maxDurationJobTypeMap := make(map[int64]string)
	maxDurationMap := make(map[int64]int)
	aiCenterMap := make(map[int64]map[string]int)
	modelMap := make(map[int64]map[string]int)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Unscoped().Count(new(Cloudbrain))
	if err != nil {
		log.Info("query cloudbrain error. return.")
		return resultMap, resultItemMap, resourceItemMap, jobTypeItemMap, cardTypeItemMap, maxDurationJobTypeMap, maxDurationMap, aiCenterMap, modelMap
	}
	log.Info("cloudbrain count=" + fmt.Sprint(count))
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,job_type,user_id,duration,train_job_duration,type,compute_resource,ai_center,model_name,model_id").Table("cloudbrain").Unscoped().Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		cloudTaskList := make([]*Cloudbrain, 0)
		sess.Find(&cloudTaskList)
		log.Info("query cloudbrain size=" + fmt.Sprint(len(cloudTaskList)))
		for _, cloudTaskRecord := range cloudTaskList {
			if _, ok := resultMap[cloudTaskRecord.UserID]; !ok {
				resultMap[cloudTaskRecord.UserID] = 1
			} else {
				resultMap[cloudTaskRecord.UserID] += 1
			}
			if _, ok := resourceItemMap[cloudTaskRecord.UserID]; !ok {
				resourceItemMap[cloudTaskRecord.UserID] = make(map[string]int)
			}

			if _, ok := jobTypeItemMap[cloudTaskRecord.UserID]; !ok {
				jobTypeItemMap[cloudTaskRecord.UserID] = make(map[string]int)
			}

			if _, ok := cardTypeItemMap[cloudTaskRecord.UserID]; !ok {
				cardTypeItemMap[cloudTaskRecord.UserID] = make(map[string]int)
			}
			if _, ok := aiCenterMap[cloudTaskRecord.UserID]; !ok {
				aiCenterMap[cloudTaskRecord.UserID] = make(map[string]int)
			}
			if _, ok := modelMap[cloudTaskRecord.UserID]; !ok {
				modelMap[cloudTaskRecord.UserID] = make(map[string]int)
			}

			spec, _ := GetCloudbrainSpecByID(cloudTaskRecord.ID)
			if spec != nil {

				if cloudTaskRecord.Duration < 100000000 && cloudTaskRecord.Duration > 0 {
					cardSeconds := getCardSeconds(spec.AccCardsNum, cloudTaskRecord.Duration, cloudTaskRecord.WorkServerNumber)

					setMapKey("CloudBrainRunTime", cloudTaskRecord.UserID, cardSeconds, resultItemMap)

					resourceItemMap[cloudTaskRecord.UserID][cloudTaskRecord.ComputeResource] = resourceItemMap[cloudTaskRecord.UserID][cloudTaskRecord.ComputeResource] + int(cloudTaskRecord.Duration)
					cardTypeItemMap[cloudTaskRecord.UserID][spec.AccCardType] = cardTypeItemMap[cloudTaskRecord.UserID][spec.AccCardType] + int(cloudTaskRecord.Duration)
					if int(cloudTaskRecord.Duration) > maxDurationMap[cloudTaskRecord.UserID] {
						maxDurationMap[cloudTaskRecord.UserID] = int(cloudTaskRecord.Duration)
						maxDurationJobTypeMap[cloudTaskRecord.UserID] = cloudTaskRecord.JobType
					}
				}
			}

			if _, ok := jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType]; !ok {
				jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType] = 1
			} else {
				jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType] += 1
			}

			centerName := getAiCenterName(cloudTaskRecord.AiCenter)
			if centerName != "" {
				if _, ok := aiCenterMap[cloudTaskRecord.UserID][centerName]; !ok {
					aiCenterMap[cloudTaskRecord.UserID][centerName] = 1
				} else {
					aiCenterMap[cloudTaskRecord.UserID][centerName] += 1
				}
			}

			if cloudTaskRecord.JobType == string(JobTypeModelExperience) || cloudTaskRecord.JobType == string(JobTypeComfyuiExperience) {
				if cloudTaskRecord.ModelId != "" {
					modelIds := strings.Split(cloudTaskRecord.ModelId, ";")
					modelNames := strings.Split(cloudTaskRecord.ModelName, ";")
					for i, modelId := range modelIds {
						var modelName string

						model, _ := QueryModelById(modelId)
						if model != nil {
							modelName = model.DisplayName()
						} else {
							if i < len(modelNames) {
								modelName = modelNames[i]
							}
						}
						if modelName != "" {

							if _, ok := modelMap[cloudTaskRecord.UserID][modelName]; !ok {
								modelMap[cloudTaskRecord.UserID][modelName] = 1
							} else {
								modelMap[cloudTaskRecord.UserID][modelName] = modelMap[cloudTaskRecord.UserID][modelName] + 1
							}
						}
					}

				}

			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap, resultItemMap, resourceItemMap, jobTypeItemMap, cardTypeItemMap, maxDurationJobTypeMap, maxDurationMap, aiCenterMap, modelMap
}

func getAiCenterName(aiCenter string) string {
	if aiCenter != "" {
		values := strings.Split(aiCenter, "+")
		if len(values) > 1 {
			return values[1]
		}
	}
	return ""
}
func getCardSeconds(accCardNum int, duration int64, nodes int) int {
	tempNodes := 1
	if nodes > 0 {
		tempNodes = nodes
	}
	if accCardNum > 0 {
		return accCardNum * tempNodes * int(duration)
	}
	return int(duration) * tempNodes
}

func getMostLoginMonthJson(userMonthLoginInfo map[string]int, userLoginCount int, loginValueArray []int, userCount int) string {
	mostMap := make(map[string]string)
	max := 0
	max_month := ""
	for key, value := range userMonthLoginInfo {
		if value > max {
			max = value
			max_month = key
		}
	}
	mostMap["most_login_month"] = max_month
	mostMap["most_login_month_num"] = fmt.Sprint(max)

	location := binarySearchDesc(loginValueArray, userLoginCount)
	if location >= 0 {
		mostMap["location"] = fmt.Sprintf("%.2f", float64(userCount-location)/float64(userCount))
	} else {
		mostMap["location"] = "0"
	}

	mostMapJson, _ := json.Marshal(mostMap)
	return string(mostMapJson)

}

func binarySearchDesc(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			right = mid - 1 // 目标值较大，在左侧
		} else {
			left = mid + 1 // 目标值较小，在右侧
		}
	}

	return -1
}

func getActionDaysInfo(days int) string {
	return fmt.Sprint(days)
}

func getForumInfo(forumMap map[string]map[string]interface{}, name string) string {
	record, ok := forumMap[name]
	if ok {
		forumJson, _ := json.Marshal(record)
		return string(forumJson)
	}
	return ""
}

func getActionInfo(dateRecordAll *UserBusinessAnalysisAll, MostActiveDayMap map[int64]map[string]int) string {
	record, ok := MostActiveDayMap[dateRecordAll.ID]
	if ok {
		actionJson, _ := json.Marshal(record)
		return string(actionJson)
	}
	return ""
}

func getPointInfo(dateRecordAll *UserBusinessAnalysisAll, PointMap map[int64]map[string]float64) string {
	convertMap := make(map[string]string)
	record, ok := PointMap[dateRecordAll.ID]

	if ok {
		convertMap["decrease_count"] = strconv.FormatFloat(record["decrease_count"], 'f', 0, 64)
		convertMap["increase_count"] = strconv.FormatFloat(record["increase_count"], 'f', 0, 64)
		convertMap["decrease"] = strconv.FormatFloat(record["decrease"], 'f', 1, 64)
		convertMap["increase"] = strconv.FormatFloat(record["increase"], 'f', 1, 64)

		pointJson, _ := json.Marshal(convertMap)
		return string(pointJson)
	}
	return ""
}

func getModelInfo(dateRecordAll *UserBusinessAnalysisAll, userModelMap map[int64]map[string]interface{}) string {
	record, ok := userModelMap[dateRecordAll.ID]

	if ok {
		userModelJson, _ := json.Marshal(record)
		return string(userModelJson)
	}
	return ""
}

func getIssueInfo(dateRecordAll *UserBusinessAnalysisAll) string {
	issueInfo := make(map[string]int)

	issueInfo["create_count"] = dateRecordAll.IssueCount
	issueInfo["solve_count"] = dateRecordAll.SolveIssueCount
	issueInfo["comment_count"] = dateRecordAll.CommentCount

	if dateRecordAll.IssueCount == 0 && dateRecordAll.SolveIssueCount == 0 && dateRecordAll.CommentCount == 0 {
		return ""
	}

	issueInfoJson, _ := json.Marshal(issueInfo)
	return string(issueInfoJson)
}

func getCourseMap() map[int64]map[string]int {
	courseMap := make(map[int64]map[string]int)
	url := setting.RecommentRepoAddr + "bonus/training.txt"
	content, err := GetContentFromPromote(url)
	if err == nil {
		filenames := strings.Split(content, "\n")
		for i := 0; i < len(filenames); i++ {
			filenames[i] = strings.TrimSuffix(filenames[i], "\r")
			url = setting.RecommentRepoAddr + "bonus/" + filenames[i]
			log.Info("bonus url=" + url)
			csvContent, err1 := GetContentFromPromote(url)
			if err1 == nil {
				//read csv
				lines := strings.Split(csvContent, "\n")
				for j := 1; j < len(lines); j++ {
					newLine := strings.TrimSuffix(lines[j], "\r")
					aLine := strings.Split(newLine, ",")
					if len(aLine) < 7 {
						continue
					}
					userIDStr := aLine[0]
					userID, err := strconv.ParseInt(userIDStr, 10, 64)
					if err != nil {
						continue
					}
					record, ok := courseMap[userID]
					if !ok {
						record = make(map[string]int)
						courseMap[userID] = record
					}
					record["course"] = getIntValue(aLine[1])
					record["chapter"] = getIntValue(aLine[2])
					record["deep_learning"] = getIntValue(aLine[3])
					record["marchine_learning"] = getIntValue(aLine[4])
					record["big_data"] = getIntValue(aLine[5])
					record["llm"] = getIntValue(aLine[6])

				}
			}
		}
	}
	return courseMap

}

func getOpenSourceMap() map[string]map[string]int {
	opensourceMap := make(map[string]map[string]int)
	url := setting.RecommentRepoAddr + "bonus/open_source.txt"
	content, err := GetContentFromPromote(url)
	if err == nil {
		filenames := strings.Split(content, "\n")
		for i := 0; i < len(filenames); i++ {
			filenames[i] = strings.TrimSuffix(filenames[i], "\r")
			url = setting.RecommentRepoAddr + "bonus/" + filenames[i]
			log.Info("bonus url=" + url)
			csvContent, err1 := GetContentFromPromote(url)
			if err1 == nil {
				//read csv
				lines := strings.Split(csvContent, "\n")
				for j := 1; j < len(lines); j++ {
					newLine := strings.TrimSuffix(lines[j], "\r")
					aLine := strings.Split(newLine, ",")
					if len(aLine) < 3 {
						continue
					}
					userName := aLine[0]
					record, ok := opensourceMap[userName]
					if !ok {
						record = make(map[string]int)
						opensourceMap[userName] = record
					}
					record["task"] = getIntValue(aLine[1])
					record["money"] = getIntValue(aLine[2])

				}
			}
		}
	}
	return opensourceMap
}

func getBonusMap() map[string]map[string]int {
	bonusMap := make(map[string]map[string]int)
	url := setting.RecommentRepoAddr + "bonus/2024.txt"
	content, err := GetContentFromPromote(url)
	if err == nil {
		filenames := strings.Split(content, "\n")
		for i := 0; i < len(filenames); i++ {
			filenames[i] = strings.TrimSuffix(filenames[i], "\r")
			url = setting.RecommentRepoAddr + "bonus/" + filenames[i]
			log.Info("bonus url=" + url)
			csvContent, err1 := GetContentFromPromote(url)
			if err1 == nil {
				//read csv
				lines := strings.Split(csvContent, "\n")
				for j := 1; j < len(lines); j++ {
					newLine := strings.TrimSuffix(lines[j], "\r")
					aLine := strings.Split(newLine, ",")
					if len(aLine) < 5 {
						continue
					}
					userName := aLine[0]
					record, ok := bonusMap[userName]
					if !ok {
						record = make(map[string]int)
						bonusMap[userName] = record
					}
					record["times"] = getIntValue(aLine[4])
					record["total_model"] = getIntValue(aLine[1])
					record["total_reward"] = getIntValue(aLine[2])
					record["total_point"] = getIntValue(aLine[3])

				}
			}
		}
	}
	return bonusMap
}

func getForumMap() map[string]map[string]interface{} {
	sep := "(-2023-)"
	forumMap := make(map[string]map[string]interface{})
	url := setting.RecommentRepoAddr + "forum/data.txt"
	content, err := GetContentFromPromote(url)
	if err == nil {
		lines := strings.Split(content, "\n")
		for i := 0; i < len(lines); i++ {
			newLine := strings.TrimSuffix(lines[i], "\r")
			//read a line
			aLine := strings.Split(newLine, sep)
			if len(aLine) < 6 {
				continue
			}
			userName := aLine[0]
			record, ok := forumMap[userName]
			if !ok {
				record = make(map[string]interface{})
				forumMap[userName] = record
			}
			record["max_subject"] = aLine[1]
			record["total_view_num"] = getIntValue(aLine[2])
			record["total_num"] = getIntValue(aLine[3])
			record["total_star_num"] = getIntValue(aLine[4])
			record["max_view_count"] = getIntValue(aLine[5])
		}
	}
	return forumMap
}

func getPlayARoll(bonusMap map[string]map[string]int, userName string) string {
	record, ok := bonusMap[userName]
	if ok {

		bonusInfoJson, _ := json.Marshal(record)
		return string(bonusInfoJson)
	} else {
		return ""
	}
}

func getCourseInfo(couseMap map[int64]map[string]int, userID int64) string {
	record, ok := couseMap[userID]
	if ok {

		courseInfoJson, _ := json.Marshal(record)
		return string(courseInfoJson)
	} else {
		return ""
	}
}

func queryUserYearModel(start_unix int64, end_unix int64) map[int64]map[string]interface{} {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]map[string]interface{})

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(AiModelManage))
	if err != nil {
		log.Info("query AiModelManage error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,name,user_id,download_count,reference_count,model_type, is_private,alias").Table("ai_model_manage").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		aiModelList := make([]*AiModelManage, 0)
		sess.Find(&aiModelList)
		log.Info("query user year AiModelManage size=" + fmt.Sprint(len(aiModelList)))
		for _, aiModelRecord := range aiModelList {
			if aiModelRecord.ModelType == 2 {
				realUserID := getRealUserId(aiModelRecord.ID)
				if realUserID == 0 {
					continue
				}
				if _, ok := resultMap[realUserID]; !ok {
					modelmap := make(map[string]interface{})
					modelmap["count"] = 1
					modelmap["max_download_count"] = aiModelRecord.DownloadCount
					modelmap["max_reference_count"] = aiModelRecord.ReferenceCount
					modelmap["download_name"] = aiModelRecord.DisplayName()
					modelmap["reference_name"] = aiModelRecord.DisplayName()
					modelmap["migration_count"] = 1
					if aiModelRecord.IsPrivate {
						modelmap["private_count"] = 1
					} else {
						modelmap["private_count"] = 0
					}

					resultMap[realUserID] = modelmap
				} else {
					resultMap[realUserID]["count"] = resultMap[realUserID]["count"].(int) + 1
					if resultMap[realUserID]["max_download_count"].(int) < aiModelRecord.DownloadCount {
						resultMap[realUserID]["max_download_count"] = aiModelRecord.DownloadCount

						resultMap[realUserID]["download_name"] = aiModelRecord.DisplayName()
					}
					if resultMap[realUserID]["max_reference_count"].(int) < aiModelRecord.ReferenceCount {
						resultMap[realUserID]["max_reference_count"] = aiModelRecord.ReferenceCount
						resultMap[realUserID]["reference_name"] = aiModelRecord.DisplayName()
					}
					resultMap[realUserID]["migration_count"] = resultMap[realUserID]["migration_count"].(int) + 1

					if aiModelRecord.IsPrivate {
						resultMap[realUserID]["private_count"] = resultMap[realUserID]["private_count"].(int) + 1
					}

				}

			} else {
				if _, ok := resultMap[aiModelRecord.UserId]; !ok {
					modelmap := make(map[string]interface{})
					modelmap["count"] = 1
					modelmap["max_download_count"] = aiModelRecord.DownloadCount
					modelmap["max_reference_count"] = aiModelRecord.ReferenceCount
					modelmap["download_name"] = aiModelRecord.DisplayName()
					modelmap["reference_name"] = aiModelRecord.DisplayName()
					modelmap["migration_count"] = 0
					if aiModelRecord.IsPrivate {
						modelmap["private_count"] = 1
					} else {
						modelmap["private_count"] = 0
					}

					resultMap[aiModelRecord.UserId] = modelmap
				} else {
					resultMap[aiModelRecord.UserId]["count"] = resultMap[aiModelRecord.UserId]["count"].(int) + 1
					if resultMap[aiModelRecord.UserId]["max_download_count"].(int) < aiModelRecord.DownloadCount {
						resultMap[aiModelRecord.UserId]["max_download_count"] = aiModelRecord.DownloadCount

						resultMap[aiModelRecord.UserId]["download_name"] = aiModelRecord.DisplayName()
					}
					if resultMap[aiModelRecord.UserId]["max_reference_count"].(int) < aiModelRecord.ReferenceCount {
						resultMap[aiModelRecord.UserId]["max_reference_count"] = aiModelRecord.ReferenceCount
						resultMap[aiModelRecord.UserId]["reference_name"] = aiModelRecord.DisplayName()
					}
					if aiModelRecord.IsPrivate {
						resultMap[aiModelRecord.UserId]["private_count"] = resultMap[aiModelRecord.UserId]["private_count"].(int) + 1
					}

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

func getRealUserId(modelId string) int64 {
	var ints []int64
	err := x.Table("hf_model_file").Cols("user_id").Where("model_id=?", modelId).Find(&ints)
	if err != nil || len(ints) == 0 {
		return 0
	}
	return ints[0]

}

func queryPointInfo(start_unix int64, end_unix int64) map[int64]map[string]float64 {

	sess := x.NewSession()
	defer sess.Close()
	scoreInfoMap := make(map[int64]map[string]float64)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(RewardOperateRecord))
	if err != nil {
		log.Info("query PointAccountLog error. return.")
		return scoreInfoMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("user_id,amount,operate_type").Table("reward_operate_record").Where(cond).Limit(PAGE_SIZE, int(indexTotal))
		rewardOperateRecordList := make([]*RewardOperateRecord, 0)
		sess.Find(&rewardOperateRecordList)
		log.Info("query user year AiModelManage size=" + fmt.Sprint(len(rewardOperateRecordList)))
		for _, rewardOperateRecord := range rewardOperateRecordList {
			if _, ok := scoreInfoMap[rewardOperateRecord.UserId]; !ok {
				pointMap := make(map[string]float64)
				scoreInfoMap[rewardOperateRecord.UserId] = pointMap
				scoreInfoMap[rewardOperateRecord.UserId]["increase"] = 0
				scoreInfoMap[rewardOperateRecord.UserId]["increase_count"] = 0
				scoreInfoMap[rewardOperateRecord.UserId]["decrease"] = 0
				scoreInfoMap[rewardOperateRecord.UserId]["decrease_count"] = 0

			}
			if rewardOperateRecord.OperateType == "INCREASE" && rewardOperateRecord.Amount > 0 {
				scoreInfoMap[rewardOperateRecord.UserId]["increase"] += rewardOperateRecord.Amount
				scoreInfoMap[rewardOperateRecord.UserId]["increase_count"] += 1
			}
			if rewardOperateRecord.OperateType == "DECREASE" && rewardOperateRecord.Amount > 0 {
				scoreInfoMap[rewardOperateRecord.UserId]["decrease"] += rewardOperateRecord.Amount
				scoreInfoMap[rewardOperateRecord.UserId]["decrease_count"] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return scoreInfoMap
}

func queryMostActiveMonth(start_unix int64, end_unix int64) map[int64]map[string]int {

	sess := x.NewSession()
	defer sess.Close()

	actionInfoMap := make(map[int64]map[string]int)

	cond := "(op_type<=" + fmt.Sprint(17) + " or op_type>=21) and user_id=act_user_id and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(Action))
	if err != nil {
		log.Info("query Action error. return.")
		return actionInfoMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		actionList, err := sess.QueryInterface("select id,user_id,op_type,act_user_id,created_unix from public.action where " + cond + " order by id asc limit " + fmt.Sprint(PAGE_SIZE) + " offset " + fmt.Sprint(indexTotal))
		if err != nil {
			log.Info("error:" + err.Error())
			continue
		}
		log.Info("query action size=" + fmt.Sprint(len(actionList)))
		for _, actionRecord := range actionList {
			userId := convertInterfaceToInt64(actionRecord["user_id"])
			created_unix := convertInterfaceToInt64(actionRecord["created_unix"])
			created_time := time.Unix(created_unix, 0)
			str_time := created_time.Format("20060102")
			str_time = str_time[4:]
			if _, ok := actionInfoMap[userId]; !ok {
				actionMap := make(map[string]int)
				actionInfoMap[userId] = actionMap
			}
			actionInfoMap[userId][str_time] += 1
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	returnMap := make(map[int64]map[string]int)
	for k, v := range actionInfoMap {
		returnMap[k] = getActionMostActiveDayMap(v)
	}
	return returnMap
}

func getActionMostActiveDayMap(record map[string]int) map[string]int {
	reMap := make(map[string]int)
	monthMap := make(map[string]int)
	for k, v := range record {
		month := k[0:2]
		monthMap[month] += v
	}
	maxMonthCount := 0
	maxMonthCountKey := ""
	for k, v := range monthMap {
		if v > maxMonthCount {
			maxMonthCount = v
			maxMonthCountKey = k
		}
	}
	reMap[maxMonthCountKey] = maxMonthCount

	maxDayCount := 0
	maxDayCountKey := ""
	for k, v := range record {
		month := k[0:2]
		if month == maxMonthCountKey {
			if v > maxDayCount {
				maxDayCount = v
				maxDayCountKey = k
			}
		}
	}
	reMap[maxDayCountKey] = maxDayCount

	return reMap
}

func getRepoDetailInfo(repoDetailInfoMap map[string]int, userId int64, mostDownload map[int64]string, fourSeason map[string]map[string]interface{}) string {
	repoDetailInfo := make(map[string]interface{})
	if total, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_total"]; ok {
		repoDetailInfo["repo_total"] = fmt.Sprint(total)
	}
	if private, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_is_private"]; ok {
		repoDetailInfo["repo_is_private"] = fmt.Sprint(private)
	}
	if public, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_is_public"]; ok {
		repoDetailInfo["repo_is_public"] = fmt.Sprint(public)
	}
	if fork, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_is_fork"]; ok {
		repoDetailInfo["repo_is_fork"] = fmt.Sprint(fork)
	} else {
		repoDetailInfo["repo_is_fork"] = "0"
	}
	if download, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_total_download"]; ok {
		repoDetailInfo["repo_total_download"] = fmt.Sprint(download)
	}
	if mostdownload, ok := repoDetailInfoMap[fmt.Sprint(userId)+"_most_download"]; ok {
		repoDetailInfo["repo_most_download_count"] = fmt.Sprint(mostdownload)
	}
	if mostdownloadName, ok := mostDownload[userId]; ok {
		repoDetailInfo["repo_most_download_name"] = mostdownloadName
	}
	if springRepo, ok := fourSeason[fmt.Sprint(userId)+"_spring"]; ok {
		repoDetailInfo["repo_spring"] = springRepo
	}
	if summerRepo, ok := fourSeason[fmt.Sprint(userId)+"_summer"]; ok {
		repoDetailInfo["repo_summer"] = summerRepo
	}
	if autumnRepo, ok := fourSeason[fmt.Sprint(userId)+"_autumn"]; ok {
		repoDetailInfo["repo_autumn"] = autumnRepo
	}
	if winterRepo, ok := fourSeason[fmt.Sprint(userId)+"_winter"]; ok {
		repoDetailInfo["repo_winter"] = winterRepo
	}
	if len(repoDetailInfo) > 0 {
		repoDetailInfoJson, _ := json.Marshal(repoDetailInfo)
		return string(repoDetailInfoJson)
	} else {
		return ""
	}
}

func getCodeInfo(dateRecordAll *UserBusinessAnalysisAll, commitCodeArrays []int) (string, float64) {
	if dateRecordAll.CommitCount > 0 {
		codeInfo := make(map[string]string)
		codeInfo["commit_count"] = fmt.Sprint(dateRecordAll.CommitCount)
		codeInfo["commit_line"] = fmt.Sprint(dateRecordAll.CommitCodeSize)
		score := 0.0
		score = float64(dateRecordAll.CommitCodeSize) / float64(dateRecordAll.CommitCount) / float64(20000)
		if score < (float64(dateRecordAll.CommitCount) / float64(100)) {
			score = float64(dateRecordAll.CommitCount) / float64(100)
		}
		log.Info("len(commitCodeArrays)=" + fmt.Sprint(len(commitCodeArrays)))
		location := binarySearch(commitCodeArrays, dateRecordAll.CommitCodeSize)

		codeInfo["location"] = fmt.Sprintf("%.2f", float64(location+1)/float64(len(commitCodeArrays)))

		codeInfo["score"] = fmt.Sprintf("%.2f", score)

		codeInfoJson, _ := json.Marshal(codeInfo)
		return string(codeInfoJson), score
	} else {
		return "", 0
	}
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return mid // 如果没有找到，返回-1
}

func getMostActiveJson(userInfo map[string]int, userMonthInfo map[string]int, userPeriodInfo map[string]int, userStartEndInfo map[string]int, userStartEndDayInfo map[string]timeutil.TimeStamp, userSeazonInfo map[string]int, actionTaskTypeMap map[int]int) string {
	mostActiveMap := make(map[string]string)
	if day, ok := userInfo["hour_day"]; ok {
		hour := userInfo["hour_hour"]
		month := userInfo["hour_month"]
		year := userInfo["hour_year"]
		delete(userInfo, "hour_day")
		delete(userInfo, "hour_hour")
		delete(userInfo, "hour_month")
		delete(userInfo, "hour_year")
		mostActiveMap["before_dawn"] = fmt.Sprint(year) + "/" + fmt.Sprint(month) + "/" + fmt.Sprint(day) + " " + fmt.Sprint(hour)
	}
	max := 0
	max_day := ""
	for key, value := range userInfo {
		if value > max {
			max = value
			max_day = key
		}
	}
	mostActiveMap["most_active_day"] = max_day
	mostActiveMap["most_active_num"] = fmt.Sprint(max)

	max = 0
	max_month := ""
	for key, value := range userMonthInfo {
		if value > max {
			max = value
			max_month = key
		}
	}
	mostActiveMap["most_active_month"] = max_month
	mostActiveMap["most_active_month_num"] = fmt.Sprint(max)

	max = 0
	max_seazon := ""
	for key, value := range userSeazonInfo {
		if value > max {
			max = value
			max_seazon = key
		}
	}
	mostActiveMap["s_1"] = fmt.Sprint(userSeazonInfo["1"])
	mostActiveMap["s_2"] = fmt.Sprint(userSeazonInfo["2"])
	mostActiveMap["s_3"] = fmt.Sprint(userSeazonInfo["3"])
	mostActiveMap["s_4"] = fmt.Sprint(userSeazonInfo["4"])

	mostActiveMap["most_active_seazon"] = max_seazon
	mostActiveMap["most_active_seazon_num"] = fmt.Sprint(max)
	mostActiveMap["task"] = sortedKeysToString(actionTaskTypeMap)

	max = 0
	max_period := ""
	for key, value := range userPeriodInfo {
		if value > max {
			max = value
			max_period = key
		}
	}
	mostActiveMap["most_active_period"] = max_period
	//mostActiveMap["most_active_period_num"] = fmt.Sprint(max)

	mostActiveMap["most_start_active"] = getHourAndMinuteStr(userStartEndInfo["start_unix"])
	mostActiveMap["most_end_active"] = getHourAndMinuteStr(userStartEndInfo["end_unix"])

	mostActiveMap["most_start_day_active"] = getDate(userStartEndDayInfo["start_unix"])
	mostActiveMap["most_end_day_active"] = getDate(userStartEndDayInfo["end_unix"])

	mostActiveMapJson, _ := json.Marshal(mostActiveMap)
	return string(mostActiveMapJson)
}

func sortedKeysToString(m map[int]int) string {
	// 获取有序的 keys
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 将 int 数组转为 string 数组
	strKeys := make([]string, len(keys))
	for i, key := range keys {
		strKeys[i] = strconv.Itoa(key)
	}

	// 用逗号连接
	return strings.Join(strKeys, ",")
}

func queryActionDays(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	timeMap := make(map[int64]string)
	cond := "user_id=act_user_id  and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(Action))
	if err != nil {
		log.Info("query action error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		actionList, err := sess.QueryInterface("select id,user_id,op_type,act_user_id,created_unix from public.action where " + cond + " order by id asc limit " + fmt.Sprint(PAGE_SIZE) + " offset " + fmt.Sprint(indexTotal))
		if err != nil {
			log.Info("error:" + err.Error())
			continue
		}
		log.Info("query action size=" + fmt.Sprint(len(actionList)))
		for _, actionRecord := range actionList {
			userId := convertInterfaceToInt64(actionRecord["user_id"])
			created_unix := convertInterfaceToInt64(actionRecord["created_unix"])
			created_time := time.Unix(created_unix, 0)
			time_str := created_time.Format("2006-01-02")
			if timeMap[userId] != time_str {
				if _, ok := resultMap[userId]; !ok {
					resultMap[userId] = 1
				} else {
					resultMap[userId] += 1
				}
				timeMap[userId] = time_str
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}
