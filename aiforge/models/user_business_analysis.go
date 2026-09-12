package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	PAGE_SIZE         = 2000
	BATCH_INSERT_SIZE = 50
)

type UserBusinessAnalysisQueryOptions struct {
	ListOptions
	UserName  string
	SortType  string
	StartTime int64
	EndTime   int64
	IsAll     bool
}

type UserBusinessAnalysisList []*UserBusinessAnalysis

func (ulist UserBusinessAnalysisList) Swap(i, j int) { ulist[i], ulist[j] = ulist[j], ulist[i] }
func (ulist UserBusinessAnalysisList) Len() int      { return len(ulist) }
func (ulist UserBusinessAnalysisList) Less(i, j int) bool {
	return ulist[i].ID > ulist[j].ID
}

func QueryMetricsPage(start int64, end int64) ([]*UserMetrics, int64) {

	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	cond := "count_date >" + fmt.Sprint(start) + " and count_date<" + fmt.Sprint(end)

	userMetricsList := make([]*UserMetrics, 0)
	//.Limit(pageSize, page*pageSize)
	if err := statictisSess.Table(new(UserMetrics)).Where(cond).OrderBy("count_date desc").
		Find(&userMetricsList); err != nil {
		return nil, 0
	}
	postUserMetricsList := postDeal(userMetricsList)
	return postUserMetricsList, int64(len(postUserMetricsList))
}

func QueryMetrics(start int64, end int64) ([]*UserMetrics, int) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	userMetricsList := make([]*UserMetrics, 0)
	if err := statictisSess.Table(new(UserMetrics)).Where("count_date >" + fmt.Sprint(start) + " and count_date<" + fmt.Sprint(end)).OrderBy("count_date desc").
		Find(&userMetricsList); err != nil {
		return nil, 0
	}
	postUserMetricsList := postDeal(userMetricsList)
	return postUserMetricsList, int(len(postUserMetricsList))
}

func duplicateRemoval(userMetricsList []*UserMetrics) []*UserMetrics {
	userMetricsResult := make([]*UserMetrics, 0)
	for i := 0; i < len(userMetricsList); i++ {
		if i > 0 {
			if userMetricsList[i].DataDate == userMetricsList[i-1].DataDate {
				continue
			}
		}
		userMetricsResult = append(userMetricsResult, userMetricsList[i])
	}
	return userMetricsResult
}

func postDeal(userMetricsList []*UserMetrics) []*UserMetrics {
	duplicateRemovalUserMetricsList := duplicateRemoval(userMetricsList)
	for _, userMetrics := range duplicateRemovalUserMetricsList {
		userMetrics.DisplayDate = userMetrics.DataDate
		userMetrics.TotalRegistUser = userMetrics.ActivateRegistUser + userMetrics.NotActivateRegistUser
		userMetrics.TotalNotActivateRegistUser = userMetrics.TotalUser - userMetrics.TotalActivateRegistUser
	}
	return duplicateRemovalUserMetricsList
}

func QueryMetricsForAll(start int64, end int64) []*UserMetrics {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	userMetricsList := make([]*UserMetrics, 0)
	if err := statictisSess.Table(new(UserMetrics)).Where("count_date >" + fmt.Sprint(start) + " and count_date<" + fmt.Sprint(end)).OrderBy("count_date desc").
		Find(&userMetricsList); err != nil {
		return nil
	}
	duplicateRemovalUserMetricsList := duplicateRemoval(userMetricsList)
	return makeResultForMonth(duplicateRemovalUserMetricsList, len(duplicateRemovalUserMetricsList))
}

func QueryMetricsForYear(start int64, end int64) []*UserMetrics {

	allUserInfo, count := QueryMetrics(start, end)

	return makeResultForMonth(allUserInfo, count)
}

func makeResultForMonth(allUserInfo []*UserMetrics, count int) []*UserMetrics {
	monthMap := make(map[string]*UserMetrics)
	if count > 0 {
		for _, userMetrics := range allUserInfo {
			dateTime := time.Unix(userMetrics.CountDate, 0)
			mInt := int(dateTime.Month())
			mString := fmt.Sprint(mInt)
			if mInt < 10 {
				mString = "0" + mString
			}
			month := fmt.Sprint(dateTime.Year()) + "-" + mString
			if _, ok := monthMap[month]; !ok {
				monthUserMetrics := &UserMetrics{
					DisplayDate:                month,
					ActivateRegistUser:         userMetrics.ActivateRegistUser,
					RegistActivityUser:         userMetrics.RegistActivityUser,
					NotActivateRegistUser:      userMetrics.NotActivateRegistUser,
					TotalUser:                  userMetrics.TotalUser,
					TotalNotActivateRegistUser: userMetrics.TotalUser - userMetrics.TotalActivateRegistUser,
					TotalActivateRegistUser:    userMetrics.TotalActivateRegistUser,
					TotalHasActivityUser:       userMetrics.TotalHasActivityUser,
					HasActivityUser:            userMetrics.HasActivityUser,
					DaysForMonth:               1,
					TotalRegistUser:            userMetrics.ActivateRegistUser + userMetrics.NotActivateRegistUser,
				}
				monthMap[month] = monthUserMetrics
			} else {
				value := monthMap[month]
				value.ActivateRegistUser += userMetrics.ActivateRegistUser
				value.NotActivateRegistUser += userMetrics.NotActivateRegistUser
				value.HasActivityUser += userMetrics.HasActivityUser
				value.RegistActivityUser += userMetrics.RegistActivityUser
				value.TotalRegistUser += userMetrics.ActivateRegistUser + userMetrics.NotActivateRegistUser
				value.ActivateIndex = float64(value.ActivateRegistUser) / float64(value.TotalRegistUser)
				value.DaysForMonth += 1
			}
		}
	}
	result := make([]*UserMetrics, 0)
	for _, value := range monthMap {
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].DisplayDate, result[j].DisplayDate) > 0
	})
	return result
}

func QueryRankList(key string, tableName string, limit int) ([]*UserBusinessAnalysisAll, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	userBusinessAnalysisAllList := make([]*UserBusinessAnalysisAll, 0)
	if err := statictisSess.Table(tableName).OrderBy(key+" desc,id desc").Limit(limit, 0).
		Find(&userBusinessAnalysisAllList); err != nil {
		return nil, 0
	}
	return userBusinessAnalysisAllList, int64(len(userBusinessAnalysisAllList))
}

func QueryUserInvitationDataByTableName(start int, pageSize int, tableName string, queryObj interface{}, userName string, invitationNum int) ([]*UserBusinessAnalysisAll, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	var cond = builder.NewCond()
	if len(userName) > 0 {
		cond = cond.And(
			builder.Like{"lower(name)", strings.ToLower(userName)},
		)
	}
	cond = cond.And(
		builder.Gte{"invitation_user_num": invitationNum},
	)

	allCount, err := statictisSess.Where(cond).Count(queryObj)
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}
	log.Info("query return total:" + fmt.Sprint(allCount))
	userBusinessAnalysisAllList := make([]*UserBusinessAnalysisAll, 0)
	if err := statictisSess.Table(tableName).Where(cond).OrderBy("invitation_user_num desc,id asc").Limit(pageSize, start).
		Find(&userBusinessAnalysisAllList); err != nil {
		return nil, 0
	}
	return userBusinessAnalysisAllList, allCount
}

func QueryUserStaticDataByTableName(start int, pageSize int, tableName string, queryObj interface{}, userName string) ([]*UserBusinessAnalysisAll, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	var cond = builder.NewCond()
	if len(userName) > 0 {
		cond = cond.And(
			builder.Like{"lower(name)", strings.ToLower(userName)},
		)
	}

	allCount, err := statictisSess.Where(cond).Count(queryObj)
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}
	log.Info("query return total:" + fmt.Sprint(allCount))
	userBusinessAnalysisAllList := make([]*UserBusinessAnalysisAll, 0)
	orderby := "user_index desc,id desc"
	if len(userName) > 0 {
		orderby = "CASE WHEN name='" + userName + "' THEN 0 ELSE 1 END," + orderby
	}
	if err := statictisSess.Table(tableName).Where(cond).OrderBy(orderby).Limit(pageSize, start).
		Find(&userBusinessAnalysisAllList); err != nil {
		return nil, 0
	}
	return userBusinessAnalysisAllList, allCount
}

func QueryUserStaticDataAll(opts *UserBusinessAnalysisQueryOptions) ([]*UserBusinessAnalysisAll, int64) {
	log.Info("query startTime =" + fmt.Sprint(opts.StartTime) + " endTime=" + fmt.Sprint(opts.EndTime) + " isAll=" + fmt.Sprint(opts.IsAll))

	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	allCount, err := statictisSess.Count(new(UserBusinessAnalysisAll))
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}
	log.Info("query return total:" + fmt.Sprint(allCount))

	pageSize := PAGE_SIZE
	totalPage := int(allCount) / pageSize
	userBusinessAnalysisReturnList := make([]*UserBusinessAnalysisAll, 0)
	for i := 0; i <= int(totalPage); i++ {
		userBusinessAnalysisAllList := make([]*UserBusinessAnalysisAll, 0)
		if err := statictisSess.Table("user_business_analysis_all").OrderBy("id desc").Limit(pageSize, i*pageSize).
			Find(&userBusinessAnalysisAllList); err != nil {
			return nil, 0
		}
		log.Info("query " + fmt.Sprint(i+1) + " result size=" + fmt.Sprint(len(userBusinessAnalysisAllList)))
		for _, userRecord := range userBusinessAnalysisAllList {
			userBusinessAnalysisReturnList = append(userBusinessAnalysisReturnList, userRecord)
		}
	}

	log.Info("return size=" + fmt.Sprint(len(userBusinessAnalysisReturnList)))
	return userBusinessAnalysisReturnList, allCount
}

func QueryDataForUserDefineFromDb(opts *UserBusinessAnalysisQueryOptions, key string) ([]*UserBusinessAnalysis, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	var cond = builder.NewCond()
	cond = cond.And(
		builder.Eq{"data_date": key},
	)
	if len(opts.UserName) > 0 {
		cond = cond.And(
			builder.Like{"name", opts.UserName},
		)
	}
	allCount, err := statictisSess.Where(cond).Count(new(UserBusinessAnalysis))
	if err == nil {
		if allCount > 0 {
			userBusinessAnalysisList := make([]*UserBusinessAnalysis, 0)
			if err := statictisSess.Table("user_business_analysis").Where(cond).OrderBy("id desc").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
				Find(&userBusinessAnalysisList); err != nil {
				return nil, 0
			}
			return userBusinessAnalysisList, allCount
		}
	}
	return nil, 0
}

func WriteDataToDb(dataList []*UserBusinessAnalysis, key string) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	log.Info("write to db, size=" + fmt.Sprint(len(dataList)))
	userBusinessAnalysisList := make([]*UserBusinessAnalysis, 0)
	for _, data := range dataList {
		data.DataDate = key
		userBusinessAnalysisList = append(userBusinessAnalysisList, data)
		if len(userBusinessAnalysisList) > BATCH_INSERT_SIZE {
			statictisSess.Insert(userBusinessAnalysisList)
			userBusinessAnalysisList = make([]*UserBusinessAnalysis, 0)
		}
	}
	if len(userBusinessAnalysisList) > 0 {
		statictisSess.Insert(userBusinessAnalysisList)
	}
}

func QueryUserStaticDataForUserDefine(opts *UserBusinessAnalysisQueryOptions, wikiCountMap map[string]int) ([]*UserBusinessAnalysis, int64) {
	log.Info("start to count other user info data")
	sess := x.NewSession()
	defer sess.Close()

	currentTimeNow := time.Now()
	log.Info("current time:" + currentTimeNow.Format("2006-01-02 15:04:05"))

	start_unix := opts.StartTime

	end_unix := opts.EndTime
	CountDate := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 1, 0, 0, currentTimeNow.Location())
	DataDate := currentTimeNow.Format("2006-01-02 15:04")

	CodeMergeCountMap := queryPullRequest(start_unix, end_unix)
	CommitCountMap := queryCommitAction(start_unix, end_unix, 5)
	IssueCountMap := queryCreateIssue(start_unix, end_unix)

	CommentCountMap := queryComment(start_unix, end_unix)
	FocusRepoCountMap := queryWatch(start_unix, end_unix)
	StarRepoCountMap := queryStar(start_unix, end_unix)
	WatchedCountMap, WatchOtherMap := queryFollow(start_unix, end_unix)
	CommitCodeSizeMap := make(map[string]*git.UserKPIStats, 0)
	// StartTime := time.Unix(start_unix, 0)
	// EndTime := time.Unix(end_unix, 0)
	//CommitCodeSizeMap, err := GetAllUserKPIStats(StartTime, EndTime)
	// if err != nil {
	// 	log.Info("query commit code errr.")
	// } else {
	// 	log.Info("query commit code size, len=" + fmt.Sprint(len(CommitCodeSizeMap)))
	// 	//CommitCodeSizeMapJson, _ := json.Marshal(CommitCodeSizeMap)
	// 	//log.Info("CommitCodeSizeMapJson=" + string(CommitCodeSizeMapJson))
	// }
	//CommitDatasetSizeMap, CommitDatasetNumMap, _ := queryDatasetSize(start_unix, end_unix)
	SolveIssueCountMap := querySolveIssue(start_unix, end_unix)
	CreateRepoCountMap, _, _, _ := queryUserCreateRepo(start_unix, end_unix)
	LoginCountMap, _ := queryLoginCount(start_unix, end_unix)
	LoginActionCountMap := queryLoginActionCount(start_unix, end_unix)
	OpenIIndexMap := queryUserRepoOpenIIndex(start_unix, end_unix)
	CloudBrainTaskMap, CloudBrainTaskItemMap, _, _ := queryCloudBrainTask(start_unix, end_unix)
	//AiModelManageMap := queryUserModel(start_unix, end_unix)
	//AiModelConvertMap := queryUserModelConvert(start_unix, end_unix)

	// 新版模型
	AimodelMap, AimodelCollectedMap, AimodelRecommendedMap := queryUserAimodel(start_unix, end_unix)
	AimodelCollectMap := queryAimodelCollection(start_unix, end_unix)

	// 新版数据集
	DatasetRegistryMap, DatasetCollectedMap, DatasetRecommendedMap := queryUserDatasetRegistry(start_unix, end_unix)
	DatasetCollectMap := queryDatasetCollection(start_unix, end_unix)

	//CollectDataset, CollectedDataset := queryDatasetStars(start_unix, end_unix)
	//RecommendDataset, _ := queryRecommedDataSet(start_unix, end_unix)
	CollectImage, CollectedImage := queryImageStars(start_unix, end_unix)
	RecommendImage := queryRecommedImage(start_unix, end_unix)

	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	cond := "type != 1 and (is_active=true or wechat_bind_unix > 0 or login_type=6)"
	count, _ := sess.Where(cond).Count(new(User))

	ParaWeight := getParaWeight()
	ResultList := make([]*UserBusinessAnalysis, 0)
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("`user`.*").Table("user").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userList := make([]*User, 0)
		sess.Find(&userList)

		for _, userRecord := range userList {
			var dateRecord UserBusinessAnalysis
			dateRecord.ID = userRecord.ID
			//log.Info("i=" + fmt.Sprint(i) + " userName=" + userRecord.Name)
			dateRecord.CountDate = CountDate.Unix()
			dateRecord.DataDate = DataDate
			dateRecord.Email = userRecord.Email
			dateRecord.Occupation = userRecord.Occupation
			dateRecord.Phone = userRecord.PhoneNumber
			dateRecord.RegistDate = userRecord.CreatedUnix
			dateRecord.Name = userRecord.Name
			dateRecord.UserLocation = userRecord.Location
			dateRecord.GiteaAgeMonth = subMonth(currentTimeNow, userRecord.CreatedUnix.AsTime())

			dateRecord.CodeMergeCount = getMapValue(dateRecord.ID, CodeMergeCountMap)
			dateRecord.CommitCount = getMapValue(dateRecord.ID, CommitCountMap)
			dateRecord.IssueCount = getMapValue(dateRecord.ID, IssueCountMap)
			dateRecord.CommentCount = getMapValue(dateRecord.ID, CommentCountMap)
			dateRecord.FocusRepoCount = getMapValue(dateRecord.ID, FocusRepoCountMap)
			dateRecord.StarRepoCount = getMapValue(dateRecord.ID, StarRepoCountMap)
			dateRecord.WatchedCount = getMapValue(dateRecord.ID, WatchedCountMap)
			dateRecord.FocusOtherUser = getMapValue(dateRecord.ID, WatchOtherMap)
			if _, ok := CommitCodeSizeMap[dateRecord.Email]; !ok {
				dateRecord.CommitCodeSize = 0
			} else {
				dateRecord.CommitCodeSize = int(CommitCodeSizeMap[dateRecord.Email].CommitLines)
			}
			//dateRecord.CommitDatasetSize = getMapValue(dateRecord.ID, CommitDatasetSizeMap)
			//dateRecord.CommitDatasetNum = getMapValue(dateRecord.ID, CommitDatasetNumMap)
			dateRecord.SolveIssueCount = getMapValue(dateRecord.ID, SolveIssueCountMap)

			dateRecord.EncyclopediasCount = getMapKeyStringValue(dateRecord.Name, wikiCountMap)

			dateRecord.CreateRepoCount = getMapValue(dateRecord.ID, CreateRepoCountMap)

			dateRecord.LoginCount = getMapValue(dateRecord.ID, LoginCountMap)
			dateRecord.LoginActionCount = getMapValue(dateRecord.ID, LoginActionCountMap)
			if _, ok := OpenIIndexMap[dateRecord.ID]; !ok {
				dateRecord.OpenIIndex = 0
			} else {
				dateRecord.OpenIIndex = OpenIIndexMap[dateRecord.ID]
			}

			dateRecord.CloudBrainTaskNum = getMapValue(dateRecord.ID, CloudBrainTaskMap)
			dateRecord.GpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuDebugJob", CloudBrainTaskItemMap)
			dateRecord.NpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuDebugJob", CloudBrainTaskItemMap)
			dateRecord.GpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuTrainJob", CloudBrainTaskItemMap)
			dateRecord.NpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuTrainJob", CloudBrainTaskItemMap)
			dateRecord.NpuInferenceJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuInferenceJob", CloudBrainTaskItemMap)
			dateRecord.GpuBenchMarkJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuBenchMarkJob", CloudBrainTaskItemMap)
			dateRecord.CloudBrainRunTime = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_CloudBrainRunTime", CloudBrainTaskItemMap)
			//dateRecord.CommitModelCount = getMapValue(dateRecord.ID, AiModelManageMap)
			//dateRecord.ModelConvertCount = getMapValue(dateRecord.ID, AiModelConvertMap)

			//dateRecord.CollectDataset = getMapValue(dateRecord.ID, CollectDataset)
			//dateRecord.CollectedDataset = getMapValue(dateRecord.ID, CollectedDataset)
			//dateRecord.RecommendDataset = getMapValue(dateRecord.ID, RecommendDataset)
			dateRecord.CollectImage = getMapValue(dateRecord.ID, CollectImage)
			dateRecord.CollectedImage = getMapValue(dateRecord.ID, CollectedImage)
			dateRecord.RecommendImage = getMapValue(dateRecord.ID, RecommendImage)

			// 新版模型
			dateRecord.OwnAimodel = getMapValue(dateRecord.ID, AimodelMap)
			dateRecord.CollectAimodel = getMapValue(dateRecord.ID, AimodelCollectMap)
			dateRecord.CollectedAimodel = getMapValue(dateRecord.ID, AimodelCollectedMap)
			dateRecord.RecommendedAimodel = getMapValue(dateRecord.ID, AimodelRecommendedMap)

			// 新版数据集
			dateRecord.CommitDatasetNumNew = getMapValue(dateRecord.ID, DatasetRegistryMap)
			dateRecord.CollectDatasetNew = getMapValue(dateRecord.ID, DatasetCollectMap)
			dateRecord.CollectedDatasetNew = getMapValue(dateRecord.ID, DatasetCollectedMap)
			dateRecord.RecommendDatasetNew = getMapValue(dateRecord.ID, DatasetRecommendedMap)

			dateRecord.UserIndexPrimitive = getUserIndex(dateRecord, ParaWeight)
			ResultList = append(ResultList, &dateRecord)
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	log.Info("query user define,count=" + fmt.Sprint(len(ResultList)))
	return ResultList, int64(len(ResultList))
}

func QueryUserStaticDataPage(opts *UserBusinessAnalysisQueryOptions) ([]*UserBusinessAnalysis, int64) {

	log.Info("query startTime =" + fmt.Sprint(opts.StartTime) + " endTime=" + fmt.Sprint(opts.EndTime) + " isAll=" + fmt.Sprint(opts.IsAll))
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	//currentTimeNow := time.Now()
	//pageStartTime := getLastCountDate()
	//pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location()).Unix()

	var cond = builder.NewCond()
	if len(opts.UserName) > 0 {
		cond = cond.And(
			builder.Like{"name", opts.UserName},
		)
	}
	cond = cond.And(
		builder.Gte{"count_date": opts.StartTime},
	)
	cond = cond.And(
		builder.Lte{"count_date": opts.EndTime},
	)

	count, err := statictisSess.Where(cond).Count(new(UserBusinessAnalysis))
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}

	if opts.Page >= 0 && opts.PageSize > 0 {
		var start int
		if opts.Page == 0 {
			start = 0
		} else {
			start = (opts.Page - 1) * opts.PageSize
		}
		statictisSess.Limit(opts.PageSize, start)
	}

	userBusinessAnalysisList := make([]*UserBusinessAnalysis, 0)
	if err := statictisSess.Table("user_business_analysis").Where(cond).OrderBy("count_date,id desc").
		Find(&userBusinessAnalysisList); err != nil {
		return nil, 0
	}

	resultMap := make(map[int64]*UserBusinessAnalysis)

	if len(userBusinessAnalysisList) > 0 {
		var newAndCond = builder.NewCond()
		var newOrCond = builder.NewCond()
		for _, userRecord := range userBusinessAnalysisList {
			newOrCond = newOrCond.Or(
				builder.Eq{"id": userRecord.ID},
			)
		}
		newAndCond = newAndCond.And(
			newOrCond,
		)
		if !opts.IsAll {
			newAndCond = newAndCond.And(
				builder.Gte{"count_date": opts.StartTime},
			)
			newAndCond = newAndCond.And(
				builder.Lte{"count_date": opts.EndTime},
			)
		}

		allCount, err := statictisSess.Where(newAndCond).Count(new(UserBusinessAnalysis))
		if err != nil {
			log.Info("query error." + err.Error())
			return nil, 0
		}

		pageSize := 1000
		totalPage := int(allCount) / pageSize

		for i := 0; i <= int(totalPage); i++ {
			userBusinessAnalysisList = make([]*UserBusinessAnalysis, 0)
			if err := statictisSess.Table("user_business_analysis").Where(newAndCond).OrderBy("count_date desc").Limit(pageSize, i*pageSize).
				Find(&userBusinessAnalysisList); err != nil {
				return nil, 0
			}
			log.Info("query result size=" + fmt.Sprint(len(userBusinessAnalysisList)))
			for _, userRecord := range userBusinessAnalysisList {
				if _, ok := resultMap[userRecord.ID]; !ok {
					resultMap[userRecord.ID] = userRecord
				} else {
					resultMap[userRecord.ID].CodeMergeCount += userRecord.CodeMergeCount
					resultMap[userRecord.ID].CommitCount += userRecord.CommitCount
					resultMap[userRecord.ID].IssueCount += userRecord.IssueCount
					resultMap[userRecord.ID].CommentCount += userRecord.CommentCount
					resultMap[userRecord.ID].FocusRepoCount += userRecord.FocusRepoCount
					resultMap[userRecord.ID].StarRepoCount += userRecord.StarRepoCount
					resultMap[userRecord.ID].WatchedCount += userRecord.WatchedCount
					resultMap[userRecord.ID].CommitCodeSize += userRecord.CommitCodeSize
					resultMap[userRecord.ID].CommitDatasetSize += userRecord.CommitDatasetSize
					resultMap[userRecord.ID].CommitDatasetNum += userRecord.CommitDatasetNum
					resultMap[userRecord.ID].CommitModelCount += userRecord.CommitModelCount
					resultMap[userRecord.ID].ModelConvertCount += userRecord.ModelConvertCount
					resultMap[userRecord.ID].SolveIssueCount += userRecord.SolveIssueCount
					resultMap[userRecord.ID].EncyclopediasCount += userRecord.EncyclopediasCount
					resultMap[userRecord.ID].CreateRepoCount += userRecord.CreateRepoCount
					resultMap[userRecord.ID].LoginCount += userRecord.LoginCount
				}
			}
		}
	}

	userBusinessAnalysisReturnList := UserBusinessAnalysisList{}
	for _, v := range resultMap {
		userBusinessAnalysisReturnList = append(userBusinessAnalysisReturnList, v)
	}
	sort.Sort(userBusinessAnalysisReturnList)
	log.Info("return size=" + fmt.Sprint(len(userBusinessAnalysisReturnList)))
	return userBusinessAnalysisReturnList, count
}

func refreshUserStaticTable(wikiCountMap map[string]int, tableName string, pageStartTime time.Time, pageEndTime time.Time, userMetrics map[string]int) {
	sess := x.NewSession()
	defer sess.Close()

	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	//var CommitCodeSizeMap map[string]*git.UserKPIStats
	CommitCodeSizeMap := make(map[string]*git.UserKPIStats, 0)
	var err error
	//var existCommitCodeSize map[int64]int
	existCommitCodeSize := make(map[int64]int, 0)
	if tableName == "user_business_analysis_all" || tableName == "user_business_analysis_current_year" || tableName == "user_business_analysis_current_month" {
		oneDayStartTime := pageEndTime.AddDate(0, 0, -1)
		oneDayStartTime = time.Date(oneDayStartTime.Year(), oneDayStartTime.Month(), oneDayStartTime.Day(), 0, 0, 0, 1, oneDayStartTime.Location())
		if oneDayStartTime.Format("2006-01-02") == pageStartTime.Format("2006-01-02") {
			existCommitCodeSize = make(map[int64]int, 0)
		} else {
			existCommitCodeSize = queryCommitCodeSizeFromDb("public." + tableName)
		}
		oneDayEndTime := time.Date(oneDayStartTime.Year(), oneDayStartTime.Month(), oneDayStartTime.Day(), 23, 59, 59, 1, oneDayStartTime.Location())
		log.Info("GetAllUserKPIStats oneDayStartTime=" + oneDayStartTime.Format("2006-01-02 15:04:05"))
		log.Info("GetAllUserKPIStats pageEndTime=" + oneDayEndTime.Format("2006-01-02 15:04:05"))
		log.Info("existCommitCodeSize len=" + fmt.Sprint(len(existCommitCodeSize)))
		//CommitCodeSizeMap, err = GetAllUserKPIStats(oneDayStartTime, oneDayEndTime)
		if err != nil {
			log.Info("query commit code errr.")
		} else {
			log.Info("query commit code size, len=" + fmt.Sprint(len(CommitCodeSizeMap)))
		}
	} else {
		//existCommitCodeSize = make(map[int64]int, 0)
		//CommitCodeSizeMap = make(map[string]*git.UserKPIStats, 0)
		//CommitCodeSizeMap, err = GetAllUserKPIStats(pageStartTime, pageEndTime)
		// if err != nil {
		// 	log.Info("query commit code errr.")
		// } else {
		// 	log.Info("query commit code size, len=" + fmt.Sprint(len(CommitCodeSizeMap)))
		// }
	}

	log.Info("truncate all data from table: " + tableName)
	statictisSess.Exec("TRUNCATE TABLE " + tableName)

	log.Info("pageStartTime:" + pageStartTime.Format("2006-01-02 15:04:05"))
	log.Info("pageEndTime time:" + pageEndTime.Format("2006-01-02 15:04:05"))

	start_unix := pageStartTime.Unix()
	end_unix := pageEndTime.Unix()

	currentTimeNow := time.Now()
	startTime := currentTimeNow.AddDate(0, 0, -1)

	CodeMergeCountMap := queryPullRequest(start_unix, end_unix)
	CommitCountMap := queryCommitAction(start_unix, end_unix, 5)
	IssueCountMap := queryCreateIssue(start_unix, end_unix)

	CommentCountMap := queryComment(start_unix, end_unix)
	FocusRepoCountMap := queryWatch(start_unix, end_unix)
	StarRepoCountMap := queryStar(start_unix, end_unix)
	WatchedCountMap, WatchOtherMap := queryFollow(start_unix, end_unix)

	SolveIssueCountMap := querySolveIssue(start_unix, end_unix)
	CreateRepoCountMap, _, _, _ := queryUserCreateRepo(start_unix, end_unix)
	LoginCountMap, _ := queryLoginCount(start_unix, end_unix)

	LoginActionCountMap := queryLoginActionCount(start_unix, end_unix)

	OpenIIndexMap := queryUserRepoOpenIIndex(startTime.Unix(), end_unix)
	CloudBrainTaskMap, CloudBrainTaskItemMap, _, _ := queryCloudBrainTask(start_unix, end_unix)

	// 新版模型
	AimodelMap, AimodelCollectedMap, AimodelRecommendedMap := queryUserAimodel(start_unix, end_unix)
	AimodelCollectMap := queryAimodelCollection(start_unix, end_unix)

	// 新版数据集
	DatasetRegistryMap, DatasetCollectedMap, DatasetRecommendedMap := queryUserDatasetRegistry(start_unix, end_unix)
	DatasetCollectMap := queryDatasetCollection(start_unix, end_unix)

	CollectImage, CollectedImage := queryImageStars(start_unix, end_unix)
	RecommendImage := queryRecommedImage(start_unix, end_unix)

	InvitationMap := queryUserInvitationCount(start_unix, end_unix)

	DataDate := currentTimeNow.Format("2006-01-02") + " 00:01"

	cond := "type != 1 and (is_active=true or wechat_bind_unix > 0 or login_type=6)"
	count, err := sess.Where(cond).Count(new(User))
	if err != nil {
		log.Info("query user error. return.")
		return
	}
	ParaWeight := getParaWeight()
	var indexTotal int64
	indexTotal = 0
	insertCount := 0
	userIndexMap := make(map[int64]float64, 0)
	maxUserIndex := 0.0
	minUserIndex := 100000000.0
	for {
		sess.Select("`user`.*").Table("user").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userList := make([]*User, 0)
		sess.Find(&userList)
		dateRecordBatch := make([]UserBusinessAnalysisAll, 0)
		for _, userRecord := range userList {
			var dateRecordAll UserBusinessAnalysisAll
			dateRecordAll.ID = userRecord.ID
			dateRecordAll.Email = userRecord.Email
			dateRecordAll.Occupation = userRecord.Occupation
			dateRecordAll.Phone = userRecord.PhoneNumber
			dateRecordAll.RegistDate = userRecord.CreatedUnix
			dateRecordAll.Name = userRecord.Name
			dateRecordAll.GiteaAgeMonth = subMonth(currentTimeNow, userRecord.CreatedUnix.AsTime())
			dateRecordAll.DataDate = DataDate
			dateRecordAll.UserLocation = userRecord.Location

			dateRecordAll.CodeMergeCount = getMapValue(dateRecordAll.ID, CodeMergeCountMap)
			dateRecordAll.CommitCount = getMapValue(dateRecordAll.ID, CommitCountMap)
			dateRecordAll.IssueCount = getMapValue(dateRecordAll.ID, IssueCountMap)
			dateRecordAll.CommentCount = getMapValue(dateRecordAll.ID, CommentCountMap)
			dateRecordAll.FocusRepoCount = getMapValue(dateRecordAll.ID, FocusRepoCountMap)
			dateRecordAll.FocusOtherUser = getMapValue(dateRecordAll.ID, WatchOtherMap)
			dateRecordAll.StarRepoCount = getMapValue(dateRecordAll.ID, StarRepoCountMap)
			dateRecordAll.WatchedCount = getMapValue(dateRecordAll.ID, WatchedCountMap)
			if _, ok := CommitCodeSizeMap[dateRecordAll.Email]; !ok {
				dateRecordAll.CommitCodeSize = getMapValue(dateRecordAll.ID, existCommitCodeSize)
			} else {
				dateRecordAll.CommitCodeSize = int(CommitCodeSizeMap[dateRecordAll.Email].CommitLines) + getMapValue(dateRecordAll.ID, existCommitCodeSize)
			}
			dateRecordAll.SolveIssueCount = getMapValue(dateRecordAll.ID, SolveIssueCountMap)
			dateRecordAll.EncyclopediasCount = getMapKeyStringValue(dateRecordAll.Name, wikiCountMap)
			dateRecordAll.CreateRepoCount = getMapValue(dateRecordAll.ID, CreateRepoCountMap)
			dateRecordAll.LoginCount = getMapValue(dateRecordAll.ID, LoginCountMap)
			dateRecordAll.LoginActionCount = getMapValue(dateRecordAll.ID, LoginActionCountMap)

			if _, ok := OpenIIndexMap[dateRecordAll.ID]; !ok {
				dateRecordAll.OpenIIndex = 0
			} else {
				dateRecordAll.OpenIIndex = OpenIIndexMap[dateRecordAll.ID]
			}

			dateRecordAll.CloudBrainTaskNum = getMapValue(dateRecordAll.ID, CloudBrainTaskMap)
			dateRecordAll.GpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuDebugJob", CloudBrainTaskItemMap)
			dateRecordAll.NpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuDebugJob", CloudBrainTaskItemMap)
			dateRecordAll.GpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuTrainJob", CloudBrainTaskItemMap)
			dateRecordAll.NpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuTrainJob", CloudBrainTaskItemMap)
			dateRecordAll.NpuInferenceJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_NpuInferenceJob", CloudBrainTaskItemMap)
			dateRecordAll.GpuBenchMarkJob = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_GpuBenchMarkJob", CloudBrainTaskItemMap)
			dateRecordAll.CloudBrainRunTime = getMapKeyStringValue(fmt.Sprint(dateRecordAll.ID)+"_CloudBrainRunTime", CloudBrainTaskItemMap)
			dateRecordAll.CollectImage = getMapValue(dateRecordAll.ID, CollectImage)
			dateRecordAll.CollectedImage = getMapValue(dateRecordAll.ID, CollectedImage)
			dateRecordAll.RecommendImage = getMapValue(dateRecordAll.ID, RecommendImage)
			dateRecordAll.InvitationUserNum = getMapValue(dateRecordAll.ID, InvitationMap)
			dateRecordAll.UserIndexPrimitive = getUserIndexFromAnalysisAll(dateRecordAll, ParaWeight)

			// 新版模型
			dateRecordAll.OwnAimodel = getMapValue(dateRecordAll.ID, AimodelMap)
			dateRecordAll.CollectAimodel = getMapValue(dateRecordAll.ID, AimodelCollectMap)
			dateRecordAll.CollectedAimodel = getMapValue(dateRecordAll.ID, AimodelCollectedMap)
			dateRecordAll.RecommendedAimodel = getMapValue(dateRecordAll.ID, AimodelRecommendedMap)

			// 新版数据集
			dateRecordAll.CommitDatasetNumNew = getMapValue(dateRecordAll.ID, DatasetRegistryMap)
			dateRecordAll.CollectDatasetNew = getMapValue(dateRecordAll.ID, DatasetCollectMap)
			dateRecordAll.CollectedDatasetNew = getMapValue(dateRecordAll.ID, DatasetCollectedMap)
			dateRecordAll.RecommendDatasetNew = getMapValue(dateRecordAll.ID, DatasetRecommendedMap)

			userIndexMap[dateRecordAll.ID] = dateRecordAll.UserIndexPrimitive
			if maxUserIndex < dateRecordAll.UserIndexPrimitive {
				maxUserIndex = dateRecordAll.UserIndexPrimitive
			}
			if minUserIndex > dateRecordAll.UserIndexPrimitive {
				minUserIndex = dateRecordAll.UserIndexPrimitive
			}

			dateRecordBatch = append(dateRecordBatch, dateRecordAll)
			if len(dateRecordBatch) >= BATCH_INSERT_SIZE {
				err := insertTable(dateRecordBatch, tableName, statictisSess)
				insertCount += BATCH_INSERT_SIZE
				if err != nil {
					log.Info("insert all data failed." + err.Error())
				}
				dateRecordBatch = make([]UserBusinessAnalysisAll, 0)
			}
			if tableName == "user_business_analysis_all" {
				tValue := getUserActivateAll(dateRecordAll)
				if tValue > 0 {
					log.Info("dateRecordAll name=" + dateRecordAll.Name + "  value=" + fmt.Sprint(tValue))
					userMetrics["TotalHasActivityUser"] = getMapKeyStringValue("TotalHasActivityUser", userMetrics) + 1
				}
			}
		}
		if len(dateRecordBatch) > 0 {
			err := insertTable(dateRecordBatch, tableName, statictisSess)
			insertCount += len(dateRecordBatch)
			if err != nil {
				log.Info("insert all data failed." + err.Error())
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	if tableName == "user_business_analysis_all" {
		log.Info("TotalHasActivityUser=" + fmt.Sprint(userMetrics["TotalHasActivityUser"]))
	}
	//normalization
	log.Info("start refresh userIndex start=" + tableName)
	updatesql := make([]string, 0)
	for k, v := range userIndexMap {
		tmpResult := (v - minUserIndex) / (maxUserIndex - minUserIndex)
		if tmpResult > 0.99 {
			tmpResult = 0.99
		}
		sql := updateUserIndex(tableName, statictisSess, k, tmpResult)
		updatesql = append(updatesql, sql)
		if len(updatesql) >= BATCH_INSERT_SIZE {
			updateTable(updatesql, statictisSess)
			updatesql = make([]string, 0)
		}
	}
	if len(updatesql) >= 0 {
		updateTable(updatesql, statictisSess)
	}
	log.Info("end refresh userIndex start=" + tableName)
	log.Info("refresh data finished.tableName=" + tableName + " total record:" + fmt.Sprint(insertCount))
}

func getBonusWeekDataMap() map[int64][]int {
	bonusMap := make(map[int64][]int)
	url := setting.RecommentRepoAddr + "bonus/weekdata/record.txt"
	content, err := GetContentFromPromote(url)
	if err == nil {
		filenames := strings.Split(content, "\n")
		for i := 0; i < len(filenames); i++ {
			if strings.HasSuffix(filenames[i], "\r") {
				filenames[i] = filenames[i][0 : len(filenames[i])-len("\r")]
			}
			url = setting.RecommentRepoAddr + "bonus/weekdata/" + filenames[i]
			csvContent, err1 := GetContentFromPromote(url)
			if err1 == nil {
				//read csv
				lines := strings.Split(csvContent, "\n")
				for j := 1; j < len(lines); j++ {
					if strings.HasSuffix(lines[j], "\r") {
						lines[j] = lines[j][0 : len(lines[j])-len("\r")]
					}
					log.Info("aLine=" + lines[j])
					aLine := strings.Split(lines[j], ",")
					if len(aLine) < 4 {
						continue
					}
					userId := getInt64Value(aLine[0])
					order := getIntValue(aLine[2])
					money := getIntValue(aLine[3])
					week, num := getWeekAndNum(filenames[i])
					//log.Info("userId=" + fmt.Sprint(userId) + " order=" + fmt.Sprint(order) + " money=" + fmt.Sprint(money) + " week=" + fmt.Sprint(week) + " num=" + fmt.Sprint(num))
					//email := lines[2]
					record, ok := bonusMap[userId]
					if !ok {
						record = make([]int, 4)
						record[0] = order
						record[1] = money
						record[2] = week
						record[3] = num
						bonusMap[userId] = record
					} else {
						if record[0] > order {
							record[0] = order
							record[1] = money
							record[2] = week
							record[3] = num
						} else {
							if record[0] == order && record[1] < money {
								record[1] = money
								record[2] = week
								record[3] = num
							}
						}
					}
				}
			}
		}
	}
	return bonusMap
}

func getWeekAndNum(name string) (int, int) {
	name = name[0 : len(name)-4]
	tmp := strings.Split(name, "_")
	if len(tmp) == 2 {
		week := getIntValue(tmp[0])
		num := getIntValue(tmp[1])
		return week, num
	}
	return 0, 0
}

func getIntValue(val string) int {
	i, err := strconv.Atoi(val)
	if err == nil {
		return i
	}
	return 0
}

func getInt64Value(val string) int64 {
	i, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return i
	}
	return 0
}

func sortedStringKeysToString(m map[string]int) string {
	// 获取有序的 keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 用逗号连接
	return strings.Join(keys, ",")
}

// 2023：有XX 个使用启智集群资源,有XX 个使用智算网络集群,使用过的计算资源有GPU NPU GCU
// 2023：你的所有任务累计运行了XXX卡时,其中 GPU资源运行XX卡时 NPU资源运行XX卡时  GCU资源运行XX卡时
func getCloudBrainInfo(dateRecordAll *UserBusinessAnalysisAll, jobTypeItemMap map[string]int, scoreMap map[string]float64, resourceItemMap map[string]int, cardTypeItemMap map[string]int, mostJobType string, mostJobTypeDuration int, centerMap map[string]int, modelMap map[string]int) string {
	trainscore := 0.0
	debugscore := 0.0
	runtime := 0.0
	if dateRecordAll.CloudBrainTaskNum > 0 {
		cloudBrainInfo := make(map[string]interface{})
		cloudBrainInfo["create_task_num"] = fmt.Sprint(dateRecordAll.CloudBrainTaskNum)
		cloudBrainInfo["debug_task_num"] = jobTypeItemMap["DEBUG"]

		cloudBrainInfo["train_task_num"] = jobTypeItemMap["TRAIN"]

		cloudBrainInfo["online_inference_task_num"] = jobTypeItemMap["ONLINEINFERENCE"]
		cloudBrainInfo["experience_task_num"] = jobTypeItemMap["MODELEXPERIENCE"]
		cloudBrainInfo["finetue_task_num"] = jobTypeItemMap["FINETUNE"]
		cloudBrainInfo["hpc_task_num"] = jobTypeItemMap["HPC"]
		cloudBrainInfo["card_runtime"] = fmt.Sprintf("%.2f", float64(dateRecordAll.CloudBrainRunTime)/3600)
		cloudBrainInfo["card_runtime_money"] = fmt.Sprint(dateRecordAll.CloudBrainRunTime * 5)
		cloudBrainInfo["card_type"] = sortedStringKeysToString(cardTypeItemMap)
		max_card_type := ""
		max_card_type_duration := 0
		for key, value := range cardTypeItemMap {
			if value > max_card_type_duration {
				max_card_type = key
				max_card_type_duration = value
			}
		}
		cloudBrainInfo["most_card_type"] = max_card_type
		cloudBrainInfo["most_card_type_duration"] = fmt.Sprint(max_card_type_duration)

		cloudBrainInfo["most_job_type"] = mostJobType
		cloudBrainInfo["most_job_type_duration"] = fmt.Sprint(mostJobTypeDuration)
		cloudBrainInfo["center"] = sortedStringKeysToString(centerMap)
		cloudBrainInfo["experience_model_count"] = len(modelMap)
		cloudBrainInfo["experience_model"] = modelMap

		if resourceItemMap != nil {
			cloudBrainInfo["computer_resource"] = resourceItemMap
		}
		cloudBrainInfoJson, _ := json.Marshal(cloudBrainInfo)
		scoreMap["trainscore"] = trainscore
		scoreMap["debugscore"] = debugscore

		return string(cloudBrainInfoJson)
	} else {
		scoreMap["trainscore"] = trainscore
		scoreMap["debugscore"] = debugscore
		scoreMap["runtime"] = runtime
		return ""
	}
}

func getDataSetInfo(userId int64, CreatedDataset map[int64]int, dataSetMostDownloadMap map[int64]int, dataSetMostDownloadNameMap map[int64]string, dataSetMostUsedMap map[int64]int, dataSetMostUsedNameMap map[int64]string, privateDataset map[int64]int) (string, float64) {
	datasetInfo := make(map[string]string)
	score := 0.0
	if create_count, ok := CreatedDataset[userId]; ok {
		datasetInfo["create_count"] = fmt.Sprint(create_count)
		score = float64(create_count) / 10
		datasetInfo["private_count"] = fmt.Sprint(privateDataset[userId])
	}

	if download_count, ok := dataSetMostDownloadMap[userId]; ok {
		datasetInfo["download_count"] = fmt.Sprint(download_count)
		datasetInfo["download_name"] = dataSetMostDownloadNameMap[userId]
	}

	if use_count, ok := dataSetMostUsedMap[userId]; ok {
		datasetInfo["use_count"] = fmt.Sprint(use_count)
		datasetInfo["use_name"] = dataSetMostUsedNameMap[userId]
	}

	if len(datasetInfo) > 0 {
		datasetInfoJson, _ := json.Marshal(datasetInfo)
		return string(datasetInfoJson), score
	} else {
		return "", score
	}
}

func updateUserIndex(tableName string, statictisSess *xorm.Session, userId int64, userIndex float64) string {
	updateSql := "UPDATE public." + tableName + " set user_index=" + fmt.Sprint(userIndex*100) + " where id=" + fmt.Sprint(userId)
	return updateSql
	//statictisSess.Exec(updateSql)
}

func updateTable(updatesql []string, statictisSess *xorm.Session) error {
	udpateBatchSql := ""
	for i, sql := range updatesql {
		udpateBatchSql += sql
		if i < (len(updatesql) - 1) {
			udpateBatchSql += ";"
		}
	}
	_, err := statictisSess.Exec(udpateBatchSql)
	return err
}

func insertTable(dateRecords []UserBusinessAnalysisAll, tableName string, statictisSess *xorm.Session) error {
	fields := []string{
		"id",
		"count_date",
		"code_merge_count",
		"commit_count",
		"issue_count",
		"comment_count",
		"focus_repo_count",
		"star_repo_count",
		"watched_count",
		"gitea_age_month",
		"commit_code_size",
		"solve_issue_count",
		"encyclopedias_count",
		"regist_date",
		"create_repo_count",
		"login_count",
		"open_i_index",
		"email",
		"name",
		"data_date",
		"cloud_brain_task_num",
		"gpu_debug_job",
		"npu_debug_job",
		"gpu_train_job",
		"npu_train_job",
		"npu_inference_job",
		"gpu_bench_mark_job",
		"cloud_brain_run_time",
		"user_index",
		"user_location",
		"focus_other_user",
		"collect_image",
		"collected_image",
		"recommend_image",
		"user_index_primitive",
		"phone",
		"invitation_user_num",
		"login_action_count",
		"occupation",
		"commit_dataset_num_new",
		"collect_dataset_new",
		"collected_dataset_new",
		"recommend_dataset_new",
		"own_aimodel",
		"collect_aimodel",
		"collected_aimodel",
		"recommended_aimodel",
	}

	insertSql := "INSERT INTO public." + tableName +
		" (" + strings.Join(fields, ", ") + ") VALUES "

	valueRows := make([]string, 0, len(dateRecords))
	for _, r := range dateRecords {
		values := []string{
			fmt.Sprint(r.ID),
			fmt.Sprint(r.CountDate),
			fmt.Sprint(r.CodeMergeCount),
			fmt.Sprint(r.CommitCount),
			fmt.Sprint(r.IssueCount),
			fmt.Sprint(r.CommentCount),
			fmt.Sprint(r.FocusRepoCount),
			fmt.Sprint(r.StarRepoCount),
			fmt.Sprint(r.WatchedCount),
			fmt.Sprint(r.GiteaAgeMonth),
			fmt.Sprint(r.CommitCodeSize),
			fmt.Sprint(r.SolveIssueCount),
			fmt.Sprint(r.EncyclopediasCount),
			fmt.Sprint(r.RegistDate),
			fmt.Sprint(r.CreateRepoCount),
			fmt.Sprint(r.LoginCount),
			fmt.Sprint(r.OpenIIndex),
			quote(r.Email),
			quote(r.Name),
			quote(r.DataDate),
			fmt.Sprint(r.CloudBrainTaskNum),
			fmt.Sprint(r.GpuDebugJob),
			fmt.Sprint(r.NpuDebugJob),
			fmt.Sprint(r.GpuTrainJob),
			fmt.Sprint(r.NpuTrainJob),
			fmt.Sprint(r.NpuInferenceJob),
			fmt.Sprint(r.GpuBenchMarkJob),
			fmt.Sprint(r.CloudBrainRunTime),
			fmt.Sprint(r.UserIndex),
			quote(r.UserLocation),
			fmt.Sprint(r.FocusOtherUser),
			fmt.Sprint(r.CollectImage),
			fmt.Sprint(r.CollectedImage),
			fmt.Sprint(r.RecommendImage),
			fmt.Sprint(r.UserIndexPrimitive),
			quote(r.Phone),
			fmt.Sprint(r.InvitationUserNum),
			fmt.Sprint(r.LoginActionCount),
			fmt.Sprint(r.Occupation),
			fmt.Sprint(r.CommitDatasetNumNew),
			fmt.Sprint(r.CollectDatasetNew),
			fmt.Sprint(r.CollectedDatasetNew),
			fmt.Sprint(r.CollectedDatasetNew),
			fmt.Sprint(r.OwnAimodel),
			fmt.Sprint(r.CollectAimodel),
			fmt.Sprint(r.CollectedAimodel),
			fmt.Sprint(r.RecommendedAimodel),
		}
		rowSql := "(" + strings.Join(values, ", ") + ")"
		valueRows = append(valueRows, rowSql)
	}
	insertSql += strings.Join(valueRows, ", ")

	_, err := statictisSess.Exec(insertSql)
	return err
}

// 对 string 做统一 quote 处理
func quote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'" // 防止 SQL 注入
}

func RefreshUserStaticAllTabel(wikiCountMap map[string]int, userMetrics map[string]int) {
	currentTimeNow := time.Now()
	pageStartTime := time.Date(2021, 11, 5, 0, 0, 0, 0, currentTimeNow.Location())
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_all", pageStartTime, pageEndTime, userMetrics)
	log.Info("refresh all data finished.")

	pageStartTime = time.Date(currentTimeNow.Year(), 1, 1, 0, 0, 0, 0, currentTimeNow.Location())
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_current_year", pageStartTime, pageEndTime, userMetrics)

	thisMonth := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 0, 0, 0, 0, currentTimeNow.Location())
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_current_month", thisMonth, pageEndTime, userMetrics)

	offset := int(time.Monday - currentTimeNow.Weekday())
	if offset > 0 {
		offset = -6
	}
	pageStartTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_current_week", pageStartTime, pageEndTime, userMetrics)

	pageEndTime = pageStartTime
	pageEndTime = time.Date(pageEndTime.Year(), pageEndTime.Month(), pageEndTime.Day(), 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	pageStartTime = pageStartTime.AddDate(0, 0, -7)

	refreshUserStaticTable(wikiCountMap, "user_business_analysis_last_week", pageStartTime, pageEndTime, userMetrics)

	pageStartTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -30)
	pageEndTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_last30_day", pageStartTime, pageEndTime, userMetrics)

	pageStartTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	pageEndTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_yesterday", pageStartTime, pageEndTime, userMetrics)

	pageStartTime = thisMonth.AddDate(0, -1, 0)
	pageEndTime = time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	refreshUserStaticTable(wikiCountMap, "user_business_analysis_last_month", pageStartTime, pageEndTime, userMetrics)
}

func CounDataByDateAndReCount(wikiCountMap map[string]int, startTime time.Time, endTime time.Time, isReCount bool) error {

	log.Info("start to count other user info data")
	sess := x.NewSession()
	defer sess.Close()

	currentTimeNow := time.Now()
	log.Info("current time:" + currentTimeNow.Format("2006-01-02 15:04:05"))

	start_unix := startTime.Unix()
	log.Info("DB query time:" + startTime.Format("2006-01-02 15:04:05"))

	end_unix := endTime.Unix()
	CountDate := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 1, 0, 0, currentTimeNow.Location())
	if isReCount {
		CountDate = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 1, 0, 0, currentTimeNow.Location())
	}

	DataDate := CountDate.Format("2006-01-02")
	CodeMergeCountMap := queryPullRequest(start_unix, end_unix)
	CommitCountMap := queryCommitAction(start_unix, end_unix, 5)
	IssueCountMap := queryCreateIssue(start_unix, end_unix)

	CommentCountMap := queryComment(start_unix, end_unix)
	FocusRepoCountMap := queryWatch(start_unix, end_unix)
	StarRepoCountMap := queryStar(start_unix, end_unix)
	WatchedCountMap, WatchOtherMap := queryFollow(start_unix, end_unix)
	CommitCodeSizeMap := make(map[string]*git.UserKPIStats, 0)
	//CommitCodeSizeMap, err := GetAllUserKPIStats(startTime, endTime)
	// if err != nil {
	// 	log.Info("query commit code errr.")
	// } else {
	// 	//log.Info("query commit code size, len=" + fmt.Sprint(len(CommitCodeSizeMap)))
	// 	//CommitCodeSizeMapJson, _ := json.Marshal(CommitCodeSizeMap)
	// 	//log.Info("CommitCodeSizeMapJson=" + string(CommitCodeSizeMapJson))
	// }
	SolveIssueCountMap := querySolveIssue(start_unix, end_unix)
	CreateRepoCountMap, _, _, _ := queryUserCreateRepo(start_unix, end_unix)
	LoginCountMap, _ := queryLoginCount(start_unix, end_unix)
	OpenIIndexMap := queryUserRepoOpenIIndex(start_unix, end_unix)
	CloudBrainTaskMap, CloudBrainTaskItemMap, _, _ := queryCloudBrainTask(start_unix, end_unix)

	// 新版模型
	AimodelMap, AimodelCollectedMap, AimodelRecommendedMap := queryUserAimodel(start_unix, end_unix)
	AimodelCollectMap := queryAimodelCollection(start_unix, end_unix)

	// 新版数据集
	DatasetRegistryMap, DatasetCollectedMap, DatasetRecommendedMap := queryUserDatasetRegistry(start_unix, end_unix)
	DatasetCollectMap := queryDatasetCollection(start_unix, end_unix)

	CollectImage, CollectedImage := queryImageStars(start_unix, end_unix)
	RecommendImage := queryRecommedImage(start_unix, end_unix)

	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	log.Info("truncate all data from table:user_business_analysis ")
	statictisSess.Exec("TRUNCATE TABLE user_business_analysis")

	cond := "type != 1"
	count, err := sess.Where(cond).Count(new(User))
	if err != nil {
		log.Info("query user error. return.")
		return err
	}
	userNewAddActivity := make(map[int64]map[int64]int64)
	userAcitvateJsonMap := make(map[int64]map[int64]int64)
	userCurrentDayRegistMap := make(map[int64]map[int64]int64)
	ParaWeight := getParaWeight()
	userMetrics := make(map[string]int)
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("`user`.*").Table("user").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userList := make([]*User, 0)
		sess.Find(&userList)

		for i, userRecord := range userList {
			var dateRecord UserBusinessAnalysis
			dateRecord.ID = userRecord.ID
			log.Info("i=" + fmt.Sprint(i) + " userName=" + userRecord.Name)
			dateRecord.CountDate = CountDate.Unix()

			dateRecord.Email = userRecord.Email
			dateRecord.Occupation = userRecord.Occupation
			dateRecord.Phone = userRecord.PhoneNumber
			dateRecord.RegistDate = userRecord.CreatedUnix
			dateRecord.Name = userRecord.Name
			dateRecord.GiteaAgeMonth = subMonth(currentTimeNow, userRecord.CreatedUnix.AsTime())
			dateRecord.DataDate = DataDate
			dateRecord.UserLocation = userRecord.Location
			dateRecord.CodeMergeCount = getMapValue(dateRecord.ID, CodeMergeCountMap)
			dateRecord.CommitCount = getMapValue(dateRecord.ID, CommitCountMap)
			dateRecord.IssueCount = getMapValue(dateRecord.ID, IssueCountMap)
			dateRecord.CommentCount = getMapValue(dateRecord.ID, CommentCountMap)
			dateRecord.FocusRepoCount = getMapValue(dateRecord.ID, FocusRepoCountMap)
			dateRecord.StarRepoCount = getMapValue(dateRecord.ID, StarRepoCountMap)
			dateRecord.WatchedCount = getMapValue(dateRecord.ID, WatchedCountMap)
			dateRecord.FocusOtherUser = getMapValue(dateRecord.ID, WatchOtherMap)
			if _, ok := CommitCodeSizeMap[dateRecord.Email]; !ok {
				dateRecord.CommitCodeSize = 0
			} else {
				dateRecord.CommitCodeSize = int(CommitCodeSizeMap[dateRecord.Email].CommitLines)
			}
			dateRecord.SolveIssueCount = getMapValue(dateRecord.ID, SolveIssueCountMap)

			dateRecord.EncyclopediasCount = getMapKeyStringValue(dateRecord.Name, wikiCountMap)

			dateRecord.CreateRepoCount = getMapValue(dateRecord.ID, CreateRepoCountMap)

			dateRecord.LoginCount = getMapValue(dateRecord.ID, LoginCountMap)

			if _, ok := OpenIIndexMap[dateRecord.ID]; !ok {
				dateRecord.OpenIIndex = 0
			} else {
				dateRecord.OpenIIndex = OpenIIndexMap[dateRecord.ID]
			}

			dateRecord.CloudBrainTaskNum = getMapValue(dateRecord.ID, CloudBrainTaskMap)
			dateRecord.GpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuDebugJob", CloudBrainTaskItemMap)
			dateRecord.NpuDebugJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuDebugJob", CloudBrainTaskItemMap)
			dateRecord.GpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuTrainJob", CloudBrainTaskItemMap)
			dateRecord.NpuTrainJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuTrainJob", CloudBrainTaskItemMap)
			dateRecord.NpuInferenceJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_NpuInferenceJob", CloudBrainTaskItemMap)
			dateRecord.GpuBenchMarkJob = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_GpuBenchMarkJob", CloudBrainTaskItemMap)
			dateRecord.CloudBrainRunTime = getMapKeyStringValue(fmt.Sprint(dateRecord.ID)+"_CloudBrainRunTime", CloudBrainTaskItemMap)
			dateRecord.CollectImage = getMapValue(dateRecord.ID, CollectImage)
			dateRecord.CollectedImage = getMapValue(dateRecord.ID, CollectedImage)
			dateRecord.RecommendImage = getMapValue(dateRecord.ID, RecommendImage)

			// 新版模型
			dateRecord.OwnAimodel = getMapValue(dateRecord.ID, AimodelMap)
			dateRecord.CollectAimodel = getMapValue(dateRecord.ID, AimodelCollectMap)
			dateRecord.CollectedAimodel = getMapValue(dateRecord.ID, AimodelCollectedMap)
			log.Info("dateRecord.ID=" + fmt.Sprint(dateRecord.ID) + " AimodelCollectedMap=" + fmt.Sprint(getMapValue(dateRecord.ID, AimodelCollectedMap)))
			dateRecord.RecommendedAimodel = getMapValue(dateRecord.ID, AimodelRecommendedMap)

			// 新版数据集
			dateRecord.CommitDatasetNumNew = getMapValue(dateRecord.ID, DatasetRegistryMap)
			dateRecord.CollectDatasetNew = getMapValue(dateRecord.ID, DatasetCollectMap)
			dateRecord.CollectedDatasetNew = getMapValue(dateRecord.ID, DatasetCollectedMap)
			dateRecord.RecommendDatasetNew = getMapValue(dateRecord.ID, DatasetRecommendedMap)

			dateRecord.UserIndexPrimitive = getUserIndex(dateRecord, ParaWeight)
			setUserMetrics(userMetrics, userRecord, start_unix, end_unix, dateRecord)
			if getUserActivate(dateRecord) > 0 {
				log.Info("has activity." + userRecord.Name)
				addUserToMap(userNewAddActivity, userRecord.CreatedUnix, dateRecord.ID, currentTimeNow)
			}
			if userRecord.IsActive {
				addUserToMap(userAcitvateJsonMap, userRecord.CreatedUnix, dateRecord.ID, currentTimeNow)
			}
			addUserToMap(userCurrentDayRegistMap, userRecord.CreatedUnix, dateRecord.ID, currentTimeNow)
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	RefreshUserStaticAllTabel(wikiCountMap, userMetrics)
	log.Info("start to update UserMetrics")
	//insert userMetrics table
	var useMetrics UserMetrics
	useMetrics.CountDate = CountDate.Unix()
	statictisSess.Delete(&useMetrics)

	useMetrics.DataDate = DataDate
	useMetrics.ActivateRegistUser = getMapKeyStringValue("ActivateRegistUser", userMetrics)
	useMetrics.HasActivityUser = getMapKeyStringValue("HasActivityUser", userMetrics)
	useMetrics.RegistActivityUser = 0
	useMetrics.NotActivateRegistUser = getMapKeyStringValue("NotActivateRegistUser", userMetrics)
	useMetrics.TotalActivateRegistUser = getMapKeyStringValue("TotalActivateRegistUser", userMetrics)
	useMetrics.TotalHasActivityUser = getMapKeyStringValue("TotalHasActivityUser", userMetrics)
	useMetrics.CurrentDayRegistUser = getMapKeyStringValue("CurrentDayRegistUser", userMetrics)
	count, err = sess.Where("type=0 and created_unix<=" + fmt.Sprint(end_unix)).Count(new(User))
	if err != nil {
		log.Info("query user error. return.")
	}
	useMetrics.TotalUser = int(count)
	if useMetrics.ActivateRegistUser+useMetrics.NotActivateRegistUser == 0 {
		useMetrics.ActivateIndex = 0
	} else {
		useMetrics.ActivateIndex = float64(useMetrics.ActivateRegistUser) / float64(useMetrics.ActivateRegistUser+useMetrics.NotActivateRegistUser)
	}
	statictisSess.Insert(&useMetrics)
	//update new user activity
	updateNewUserAcitivity(userNewAddActivity, userAcitvateJsonMap, userCurrentDayRegistMap, statictisSess)
	return nil
}

func updateNewUserAcitivity(currentUserActivity map[int64]map[int64]int64, userAcitvateJsonMap map[int64]map[int64]int64, userCurrentDayRegistMap map[int64]map[int64]int64, statictisSess *xorm.Session) {
	for key, value := range userCurrentDayRegistMap {
		useMetrics := &UserMetrics{CountDate: key}
		userAcitvateValue := userAcitvateJsonMap[key]
		HuodongValue := currentUserActivity[key]
		has, err := statictisSess.Get(useMetrics)
		if err == nil && has {
			ActivityUserArray, HuodongTotal := setUniqueUserId(useMetrics.HasActivityUserJson, HuodongValue)
			useMetrics.HasActivityUser = HuodongTotal
			useMetrics.HasActivityUserJson = ActivityUserArray

			useMetrics.CurrentDayRegistUser = len(value)

			RegistUserArray, lenRegistUser := setUniqueUserId(useMetrics.ActivityUserJson, userAcitvateValue)
			useMetrics.ActivityUserJson = RegistUserArray
			useMetrics.ActivateRegistUser = lenRegistUser

			updateSql := "update public.user_metrics set has_activity_user_json='" + useMetrics.HasActivityUserJson +
				"',regist_activity_user=" + fmt.Sprint(useMetrics.HasActivityUser) +
				",activity_user_json='" + useMetrics.ActivityUserJson + "'" +
				",activate_regist_user=" + fmt.Sprint(useMetrics.ActivateRegistUser) +
				",not_activate_regist_user=" + fmt.Sprint(useMetrics.CurrentDayRegistUser-useMetrics.ActivateRegistUser) +
				",current_day_regist_user=" + fmt.Sprint(useMetrics.CurrentDayRegistUser) +
				",activate_index=" + fmt.Sprint(float64(useMetrics.ActivateRegistUser)/float64(useMetrics.CurrentDayRegistUser)) +
				",data_date='" + time.Unix(key, 0).Format("2006-01-02") + "'" +
				" where count_date=" + fmt.Sprint(key)

			statictisSess.Exec(updateSql)
		}
	}
}

func setUniqueUserId(jsonString string, value map[int64]int64) (string, int) {
	if value == nil {
		value = make(map[int64]int64, 0)
	}
	userIdArrays := strings.Split(jsonString, ",")
	for _, userIdStr := range userIdArrays {
		userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
		if err == nil {
			value[userIdInt] = userIdInt
		}
	}
	userIdArray := ""
	for _, tmpValue := range value {
		userIdArray += fmt.Sprint(tmpValue) + ","
	}
	if len(userIdArray) > 0 {
		return userIdArray[0 : len(userIdArray)-1], len(value)
	}
	return userIdArray, len(value)
}

func addUserToMap(currentUserActivity map[int64]map[int64]int64, registDate timeutil.TimeStamp, userId int64, currentTimeNow time.Time) {
	registTime := registDate.AsTimeInLocation(currentTimeNow.Location())
	CountDateTime := time.Date(registTime.Year(), registTime.Month(), registTime.Day(), 0, 1, 0, 0, currentTimeNow.Location())
	CountDate := CountDateTime.Unix()
	if _, ok := currentUserActivity[CountDate]; !ok {
		userIdMap := make(map[int64]int64, 0)
		userIdMap[userId] = userId
		currentUserActivity[CountDate] = userIdMap
	} else {
		currentUserActivity[CountDate][userId] = userId
	}
}

func setUserMetrics(userMetrics map[string]int, user *User, start_time int64, end_time int64, dateRecord UserBusinessAnalysis) {
	//ActivateRegistUser      int   `xorm:"NOT NULL DEFAULT 0"`
	//NotActivateRegistUser   int   `xorm:"NOT NULL DEFAULT 0"`
	//HasActivityUser         int   `xorm:"NOT NULL DEFAULT 0"`
	//TotalActivateRegistUser int   `xorm:"NOT NULL DEFAULT 0"`
	//TotalHasActivityUser
	regist_time := int64(user.CreatedUnix)
	if regist_time >= start_time && regist_time <= end_time {
		if user.IsActive || user.WechatBindUnix > 0 || user.LoginType == 6 {
			userMetrics["ActivateRegistUser"] = getMapKeyStringValue("ActivateRegistUser", userMetrics) + 1
		} else {
			userMetrics["NotActivateRegistUser"] = getMapKeyStringValue("NotActivateRegistUser", userMetrics) + 1
		}
		userMetrics["CurrentDayRegistUser"] = getMapKeyStringValue("CurrentDayRegistUser", userMetrics) + 1
	}
	if user.IsActive || user.WechatBindUnix > 0 || user.LoginType == 6 {
		userMetrics["TotalActivateRegistUser"] = getMapKeyStringValue("TotalActivateRegistUser", userMetrics) + 1
	}

	if getUserActivate(dateRecord) > 0 {
		userMetrics["HasActivityUser"] = getMapKeyStringValue("HasActivityUser", userMetrics) + 1
	}

}

func getParaWeight() map[string]float64 {
	result := make(map[string]float64)
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	statictisSess.Select("*").Table(new(UserAnalysisPara))
	paraList := make([]*UserAnalysisPara, 0)
	statictisSess.Find(&paraList)
	for _, paraRecord := range paraList {
		result[paraRecord.Key] = paraRecord.Value
	}
	return result
}

func getUserIndexFromAnalysisAll(dateRecord UserBusinessAnalysisAll, ParaWeight map[string]float64) float64 {
	var result float64
	// 	PR数	0.20
	// commit数	0.20
	// 提出任务数	0.20
	// 评论数	0.20
	// 关注项目数	0.10
	// 点赞项目数	0.10
	// 登录次数	0.10
	result = float64(dateRecord.CodeMergeCount) * getParaWeightValue("CodeMergeCount", ParaWeight, 0.2)
	result += float64(dateRecord.CommitCount) * getParaWeightValue("CommitCount", ParaWeight, 0.2)
	result += float64(dateRecord.IssueCount) * getParaWeightValue("IssueCount", ParaWeight, 0.2)
	result += float64(dateRecord.CommentCount) * getParaWeightValue("CommentCount", ParaWeight, 0.2)
	result += float64(dateRecord.FocusRepoCount) * getParaWeightValue("FocusRepoCount", ParaWeight, 0.1)
	result += float64(dateRecord.StarRepoCount) * getParaWeightValue("StarRepoCount", ParaWeight, 0.1)
	result += float64(dateRecord.LoginCount) * getParaWeightValue("LoginCount", ParaWeight, 0.1)
	result += float64(dateRecord.WatchedCount) * getParaWeightValue("WatchedCount", ParaWeight, 0.3)
	codeLine := float64(dateRecord.CommitCodeSize)
	limitCodeLine := getParaWeightValue("LimitCommitCodeSize", ParaWeight, 1000)
	if codeLine >= limitCodeLine {
		codeLine = limitCodeLine
	}
	result += codeLine * getParaWeightValue("CommitCodeSize", ParaWeight, 0.01)
	result += float64(dateRecord.SolveIssueCount) * getParaWeightValue("SolveIssueCount", ParaWeight, 0.2)
	result += float64(dateRecord.EncyclopediasCount) * getParaWeightValue("EncyclopediasCount", ParaWeight, 0.1)
	result += float64(dateRecord.CreateRepoCount) * getParaWeightValue("CreateRepoCount", ParaWeight, 0.05)
	result += float64(dateRecord.CloudBrainTaskNum) * getParaWeightValue("CloudBrainTaskNum", ParaWeight, 0.3)
	result += float64(dateRecord.OwnAimodel) * getParaWeightValue("OwnAimodel", ParaWeight, 0.2)
	result += float64(dateRecord.CollectAimodel) * getParaWeightValue("CollectAimodel", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedAimodel) * getParaWeightValue("CollectedAimodel", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendedAimodel) * getParaWeightValue("RecommendedAimodel", ParaWeight, 0.2)
	result += dateRecord.OpenIIndex * getParaWeightValue("OpenIIndex", ParaWeight, 0.1)

	result += float64(dateRecord.CommitDatasetNumNew) * getParaWeightValue("CommitDatasetNumNew", ParaWeight, 0.2)
	result += float64(dateRecord.CollectDatasetNew) * getParaWeightValue("CollectDatasetNew", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedDatasetNew) * getParaWeightValue("CollectedDatasetNew", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendDatasetNew) * getParaWeightValue("RecommendDatasetNew", ParaWeight, 0.2)
	result += float64(dateRecord.CollectImage) * getParaWeightValue("CollectImage", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedImage) * getParaWeightValue("CollectedImage", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendImage) * getParaWeightValue("RecommendImage", ParaWeight, 0.2)

	return result
}

func getUserActivateAll(dateRecord UserBusinessAnalysisAll) int {
	var result int
	result += dateRecord.CodeMergeCount
	result += dateRecord.CommitCount
	result += dateRecord.IssueCount
	result += dateRecord.CommentCount
	result += dateRecord.FocusRepoCount
	result += dateRecord.StarRepoCount
	result += dateRecord.SolveIssueCount
	result += dateRecord.EncyclopediasCount
	result += dateRecord.CreateRepoCount
	result += dateRecord.CloudBrainTaskNum
	result += dateRecord.OwnAimodel
	//result += dateRecord.ModelConvertCount
	result += dateRecord.CommitDatasetNumNew
	result += dateRecord.FocusOtherUser
	result += dateRecord.CollectDataset
	result += dateRecord.CollectImage
	result += dateRecord.CommitCodeSize
	return result
}

func getUserActivate(dateRecord UserBusinessAnalysis) int {
	var result int
	result += dateRecord.CodeMergeCount
	result += dateRecord.CommitCount
	result += dateRecord.IssueCount
	result += dateRecord.CommentCount
	result += dateRecord.FocusRepoCount
	result += dateRecord.StarRepoCount
	result += dateRecord.SolveIssueCount
	result += dateRecord.EncyclopediasCount
	result += dateRecord.CreateRepoCount
	result += dateRecord.CloudBrainTaskNum
	result += dateRecord.OwnAimodel
	//result += dateRecord.ModelConvertCount
	result += dateRecord.CommitDatasetNumNew
	result += dateRecord.FocusOtherUser
	result += dateRecord.CollectDataset
	result += dateRecord.CollectImage
	result += dateRecord.CommitCodeSize
	return result
}

func getUserIndex(dateRecord UserBusinessAnalysis, ParaWeight map[string]float64) float64 {
	var result float64
	// 	PR数	0.20
	// commit数	0.20
	// 提出任务数	0.20
	// 评论数	0.20
	// 关注项目数	0.10
	// 点赞项目数	0.10
	// 登录次数	0.10
	result = float64(dateRecord.CodeMergeCount) * getParaWeightValue("CodeMergeCount", ParaWeight, 0.2)
	result += float64(dateRecord.CommitCount) * getParaWeightValue("CommitCount", ParaWeight, 0.2)
	result += float64(dateRecord.IssueCount) * getParaWeightValue("IssueCount", ParaWeight, 0.2)
	result += float64(dateRecord.CommentCount) * getParaWeightValue("CommentCount", ParaWeight, 0.2)
	result += float64(dateRecord.FocusRepoCount) * getParaWeightValue("FocusRepoCount", ParaWeight, 0.1)
	result += float64(dateRecord.StarRepoCount) * getParaWeightValue("StarRepoCount", ParaWeight, 0.1)
	result += float64(dateRecord.LoginCount) * getParaWeightValue("LoginCount", ParaWeight, 0.1)
	result += float64(dateRecord.WatchedCount) * getParaWeightValue("WatchedCount", ParaWeight, 0.3)
	codeLine := float64(dateRecord.CommitCodeSize)
	limitCodeLine := getParaWeightValue("LimitCommitCodeSize", ParaWeight, 1000)
	if codeLine >= limitCodeLine {
		codeLine = limitCodeLine
	}
	result += codeLine * getParaWeightValue("CommitCodeSize", ParaWeight, 0.01)
	result += float64(dateRecord.SolveIssueCount) * getParaWeightValue("SolveIssueCount", ParaWeight, 0.2)
	result += float64(dateRecord.EncyclopediasCount) * getParaWeightValue("EncyclopediasCount", ParaWeight, 0.1)
	result += float64(dateRecord.CreateRepoCount) * getParaWeightValue("CreateRepoCount", ParaWeight, 0.05)
	result += float64(dateRecord.CloudBrainTaskNum) * getParaWeightValue("CloudBrainTaskNum", ParaWeight, 0.3)
	result += float64(dateRecord.OwnAimodel) * getParaWeightValue("OwnAimodel", ParaWeight, 0.2)
	result += float64(dateRecord.CollectAimodel) * getParaWeightValue("CollectAimodel", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedAimodel) * getParaWeightValue("CollectedAimodel", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendedAimodel) * getParaWeightValue("RecommendedAimodel", ParaWeight, 0.2)
	result += dateRecord.OpenIIndex * getParaWeightValue("OpenIIndex", ParaWeight, 0.1)

	result += float64(dateRecord.CommitDatasetNumNew) * getParaWeightValue("CommitDatasetNumNew", ParaWeight, 0.2)
	result += float64(dateRecord.CollectDatasetNew) * getParaWeightValue("CollectDatasetNew", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedDatasetNew) * getParaWeightValue("CollectedDatasetNew", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendDatasetNew) * getParaWeightValue("RecommendDatasetNew", ParaWeight, 0.2)
	result += float64(dateRecord.CollectImage) * getParaWeightValue("CollectImage", ParaWeight, 0.1)
	result += float64(dateRecord.CollectedImage) * getParaWeightValue("CollectedImage", ParaWeight, 0.1)
	result += float64(dateRecord.RecommendImage) * getParaWeightValue("RecommendImage", ParaWeight, 0.2)

	return result
}

func getParaWeightValue(key string, valueMap map[string]float64, defaultValue float64) float64 {
	if _, ok := valueMap[key]; !ok {
		return defaultValue
	} else {
		return valueMap[key]
	}
}

func getMapKeyStringValue(key string, valueMap map[string]int) int {
	if _, ok := valueMap[key]; !ok {
		return 0
	} else {
		return valueMap[key]
	}
}

func getMapValue(userId int64, valueMap map[int64]int) int {
	if _, ok := valueMap[userId]; !ok {
		return 0
	} else {
		return valueMap[userId]
	}
}

func getInt(str string) int {
	re, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int(re)
}

func querySolveIssue(start_unix int64, end_unix int64) map[int64]int {
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
			if _, ok := resultMap[issueAssigneesRecord.AssigneeID]; !ok {
				resultMap[issueAssigneesRecord.AssigneeID] = 1
			} else {
				resultMap[issueAssigneesRecord.AssigneeID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap
}

func queryPullRequest(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := "pull_request.merged_unix>=" + fmt.Sprint(start_unix) + " and pull_request.merged_unix<=" + fmt.Sprint(end_unix)
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
			if _, ok := resultMap[issueRecord.PosterID]; !ok {
				resultMap[issueRecord.PosterID] = 1
			} else {
				resultMap[issueRecord.PosterID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryMostActiveCommitAction(start_unix int64, end_unix int64) (map[int64]int, map[int64]map[string]int, map[int64]map[string]int, map[int64]map[string]int, map[int64]map[string]int, map[int64]map[string]timeutil.TimeStamp, map[int64]map[string]int, map[int64]map[int]int) {
	sess := x.NewSession()
	defer sess.Close()
	actionCountMap := make(map[int64]int)
	actionTaskTypeMap := make(map[int64]map[int]int)
	mostActiveMap := make(map[int64]map[string]int)
	mostActiveMonthMap := make(map[int64]map[string]int)
	mostActiveSeazonMap := make(map[int64]map[string]int)
	mostActivePeriodMap := make(map[int64]map[string]int)
	mostActiveStartEndMap := make(map[int64]map[string]int)
	mostActiveStartEndDayMap := make(map[int64]map[string]timeutil.TimeStamp)

	cond := "user_id=act_user_id and op_type!=18 and op_type!=19 and op_type!=20 and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Action))
	if err != nil {
		log.Info("query action error. return.")
		return actionCountMap, mostActiveMap, mostActiveMonthMap, mostActivePeriodMap, mostActiveStartEndMap, mostActiveStartEndDayMap, mostActiveSeazonMap, actionTaskTypeMap
	}

	var indexTotal int64
	indexTotal = 0
	for {
		actionList, err := sess.QueryInterface("select id,user_id,op_type,act_user_id,created_unix from public.action where " + cond + " order by id asc limit " + fmt.Sprint(PAGE_SIZE) + " offset " + fmt.Sprint(indexTotal))
		if err != nil {
			log.Info("error:" + err.Error())
			continue
		}
		log.Info("query mostactive action size=" + fmt.Sprint(len(actionList)))
		for _, actionRecord := range actionList {
			userId := convertInterfaceToInt64(actionRecord["user_id"])
			if _, ok := actionCountMap[userId]; !ok {
				actionCountMap[userId] = 1
			} else {
				actionCountMap[userId] = actionCountMap[userId] + 1
			}
			created_unix := timeutil.TimeStamp(convertInterfaceToInt64(actionRecord["created_unix"]))
			key := getDate(created_unix)
			if _, ok := mostActiveMap[userId]; !ok {
				tmpMap := make(map[string]int)
				tmpMap[key] = 1
				mostActiveMap[userId] = tmpMap
			} else {
				mostActiveMap[userId][key] = getMapKeyStringValue(key, mostActiveMap[userId]) + 1
			}

			if _, ok := mostActiveStartEndDayMap[userId]; !ok {
				tmpMap := make(map[string]timeutil.TimeStamp)
				tmpMap["start_unix"] = created_unix
				tmpMap["end_unix"] = created_unix
				mostActiveStartEndDayMap[userId] = tmpMap
			} else {
				if mostActiveStartEndDayMap[userId]["start_unix"] > created_unix {
					mostActiveStartEndDayMap[userId]["start_unix"] = created_unix
				}
				if mostActiveStartEndDayMap[userId]["end_unix"] < created_unix {
					mostActiveStartEndDayMap[userId]["end_unix"] = created_unix
				}
			}

			if _, ok := mostActiveStartEndMap[userId]; !ok {
				tmpMap := make(map[string]int)
				tmpMap["start_unix"] = getHourAndMinute(created_unix)
				tmpMap["end_unix"] = getHourAndMinute(created_unix)
				mostActiveStartEndMap[userId] = tmpMap
			} else {
				if mostActiveStartEndMap[userId]["start_unix"] > getHourAndMinute(created_unix) {
					mostActiveStartEndMap[userId]["start_unix"] = getHourAndMinute(created_unix)
				}
				if mostActiveStartEndMap[userId]["end_unix"] < getHourAndMinute(created_unix) {
					mostActiveStartEndMap[userId]["end_unix"] = getHourAndMinute(created_unix)
				}
			}

			month_key := getMonth(created_unix)
			if _, ok := mostActiveMonthMap[userId]; !ok {
				tmpMap := make(map[string]int)
				tmpMap[month_key] = 1
				mostActiveMonthMap[userId] = tmpMap
			} else {
				mostActiveMonthMap[userId][month_key] = getMapKeyStringValue(month_key, mostActiveMonthMap[userId]) + 1
			}

			season_key := getSeason(created_unix)
			if _, ok := mostActiveSeazonMap[userId]; !ok {
				tmpMap := make(map[string]int)
				tmpMap[season_key] = 1
				mostActiveSeazonMap[userId] = tmpMap
			} else {
				mostActiveSeazonMap[userId][season_key] = getMapKeyStringValue(season_key, mostActiveSeazonMap[userId]) + 1
			}

			period_key := getPeriod(created_unix)
			if _, ok := mostActivePeriodMap[userId]; !ok {
				tmpMap := make(map[string]int)
				tmpMap[period_key] = 1
				mostActivePeriodMap[userId] = tmpMap
			} else {
				mostActivePeriodMap[userId][period_key] = getMapKeyStringValue(period_key, mostActivePeriodMap[userId]) + 1
			}

			task_type_key := getActionTaskType(int(convertInterfaceToInt64(actionRecord["op_type"])))
			if _, ok := actionTaskTypeMap[userId]; !ok {
				tmpMap := make(map[int]int)
				tmpMap[task_type_key] = 1
				actionTaskTypeMap[userId] = tmpMap
			} else {
				actionTaskTypeMap[userId][task_type_key] = actionTaskTypeMap[userId][task_type_key] + 1
			}

			utcTime := created_unix.AsTime()
			hour := utcTime.Hour()
			if hour >= 0 && hour <= 5 {
				key = "hour_hour"
				if getMapKeyStringValue(key, mostActiveMap[userId]) < hour {
					mostActiveMap[userId][key] = hour
					mostActiveMap[userId]["hour_day"] = utcTime.Day()
					mostActiveMap[userId]["hour_month"] = int(utcTime.Month())
					mostActiveMap[userId]["hour_year"] = utcTime.Year()
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return actionCountMap, mostActiveMap, mostActiveMonthMap, mostActivePeriodMap, mostActiveStartEndMap, mostActiveStartEndDayMap, mostActiveSeazonMap, actionTaskTypeMap
}

func queryCommitCodeSizeFromDb(tableName string) map[int64]int {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	resultMap := make(map[int64]int)
	count, err := statictisSess.Table(tableName).Count()
	if err != nil {
		log.Info("query " + tableName + " error. return." + err.Error())
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		commit_code_sizeList, err := statictisSess.QueryInterface("select id,commit_code_size from " + tableName + " order by id asc limit " + fmt.Sprint(PAGE_SIZE) + " offset " + fmt.Sprint(indexTotal))
		if err != nil {
			log.Info("error:" + err.Error())
			continue
		}
		log.Info("query " + tableName + " size=" + fmt.Sprint(len(commit_code_sizeList)))
		for _, record := range commit_code_sizeList {
			userId := convertInterfaceToInt64(record["id"])
			commit_code_size := convertInterfaceToInt64(record["commit_code_size"])
			if _, ok := resultMap[userId]; !ok {
				resultMap[userId] = int(commit_code_size)
			} else {
				resultMap[userId] += int(commit_code_size)
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryCommitAction(start_unix int64, end_unix int64, actionType int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := "op_type=" + fmt.Sprint(actionType) + " and user_id=act_user_id  and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
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

			if _, ok := resultMap[userId]; !ok {
				resultMap[userId] = 1
			} else {
				resultMap[userId] += 1
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func convertInterfaceToInt64(obj interface{}) int64 {
	switch obj.(type) {
	case int8:
		return int64(obj.(int8))
	case int16:
		return int64(obj.(int16))
	case int32:
		return int64(obj.(int32))
	case int64:
		return obj.(int64)
	}
	return 0
}

func getDate(createTime timeutil.TimeStamp) string {
	return createTime.Format("2006-01-02")
}

func getMonth(createTime timeutil.TimeStamp) string {
	date_str := createTime.Format("2006-01-02")
	if len(date_str) > 0 {
		return strings.Split(date_str, "-")[1]
	}
	return "0"
}
func getSeason(createTime timeutil.TimeStamp) string {
	month := getMonth(createTime)
	if month == "03" || month == "04" || month == "05" {
		return "1"
	}
	if month == "06" || month == "07" || month == "08" {
		return "2"
	}
	if month == "09" || month == "10" || month == "11" {
		return "3"
	}
	if month == "12" || month == "01" || month == "02" {
		return "4"
	}
	return "0"
}
func getPeriod(createTime timeutil.TimeStamp) string {
	hour := createTime.AsTime().Hour()
	if hour > 0 && hour <= 6 {
		return "0"
	} else if hour > 6 && hour <= 12 {
		return "1"
	} else if hour > 12 && hour <= 18 {
		return "2"
	} else {
		return "3"
	}

}

func getActionTaskType(opType int) int {

	if ActionType(opType) == ActionCreateRepo {
		return 0 //repo
	}
	if ActionType(opType) == ActionCreateIssue {
		return 1 //issue
	}
	if ActionType(opType) == ActionCreatePullRequest {
		return 2 //pr
	}
	if ActionType(opType) == ActionCreateImage {
		return 3 //image
	}
	if ActionType(opType) == ActionCreateDataset {
		return 4 //inference
	}
	if ActionType(opType) == ActionCreateAimodel {
		return 5 //model
	}
	if ActionType(opType) == ActionCommitRepo {
		return 6 //code
	}
	action := &Action{OpType: ActionType(opType)}
	if action.IsCloudbrainAction() {
		return 7 //cloudbrain
	}

	return 8

}
func getHourAndMinute(createTime timeutil.TimeStamp) int {
	hour := createTime.AsTime().Hour()
	minute := createTime.AsTime().Minute()
	return hour*60 + minute
}
func getHourAndMinuteStr(hourMinute int) string {
	hour := hourMinute / 60
	minute := hourMinute % 60
	hourstr := fmt.Sprint(hour)
	if hour < 10 {
		hourstr = "0" + fmt.Sprint(hour)
	}
	minutestr := fmt.Sprint(minute)
	if minute < 10 {
		minutestr = "0" + fmt.Sprint(minute)
	}
	return hourstr + ":" + minutestr
}

func queryCreateIssue(start_unix int64, end_unix int64) map[int64]int {

	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	cond := "is_pull=false and created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Issue))
	if err != nil {
		log.Info("query Issue error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,poster_id").Table("issue").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		issueList := make([]*Issue, 0)
		sess.Find(&issueList)
		log.Info("query issue size=" + fmt.Sprint(len(issueList)))
		for _, issueRecord := range issueList {
			if _, ok := resultMap[issueRecord.PosterID]; !ok {
				resultMap[issueRecord.PosterID] = 1
			} else {
				resultMap[issueRecord.PosterID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap

}

func queryComment(start_unix int64, end_unix int64) map[int64]int {

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
		userIssueIdMap := make(map[string]int)
		for _, commentRecord := range commentList {
			userIssueIdKey := strconv.FormatInt(commentRecord.IssueID, 10) + "_" + strconv.FormatInt(commentRecord.PosterID, 10)
			if _, ok := userIssueIdMap[userIssueIdKey]; !ok {
				if _, ok := resultMap[commentRecord.PosterID]; !ok {
					resultMap[commentRecord.PosterID] = 1
				} else {
					resultMap[commentRecord.PosterID] += 1
				}
				userIssueIdMap[userIssueIdKey] = 1

			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryWatch(start_unix int64, end_unix int64) map[int64]int {

	sess := x.NewSession()
	defer sess.Close()

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	resultMap := make(map[int64]int)
	count, err := sess.Where(cond).Count(new(Watch))
	if err != nil {
		log.Info("query issue error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		watchList := make([]*Watch, 0)
		sess.Select("id,user_id,repo_id").Table("watch").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		sess.Find(&watchList)

		log.Info("query Watch size=" + fmt.Sprint(len(watchList)))
		for _, watchRecord := range watchList {
			if _, ok := resultMap[watchRecord.UserID]; !ok {
				resultMap[watchRecord.UserID] = 1
			} else {
				resultMap[watchRecord.UserID] += 1
			}
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap

}

func queryStar(start_unix int64, end_unix int64) map[int64]int {

	sess := x.NewSession()
	defer sess.Close()

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	resultMap := make(map[int64]int)

	count, err := sess.Where(cond).Count(new(Star))
	if err != nil {
		log.Info("query star error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uid,repo_id").Table("star").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		starList := make([]*Star, 0)
		sess.Find(&starList)

		log.Info("query Star size=" + fmt.Sprint(len(starList)))
		for _, starRecord := range starList {
			if _, ok := resultMap[starRecord.UID]; !ok {
				resultMap[starRecord.UID] = 1
			} else {
				resultMap[starRecord.UID] += 1
			}
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryFollow(start_unix int64, end_unix int64) (map[int64]int, map[int64]int) {

	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	resultFocusedByOtherMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Follow))
	if err != nil {
		log.Info("query follow error. return.")
		return resultMap, resultFocusedByOtherMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id,follow_id").Table("follow").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		followList := make([]*Follow, 0)
		sess.Find(&followList)

		log.Info("query Follow size=" + fmt.Sprint(len(followList)))
		for _, followRecord := range followList {
			if _, ok := resultMap[followRecord.FollowID]; !ok {
				resultMap[followRecord.FollowID] = 1
			} else {
				resultMap[followRecord.FollowID] += 1
			}
			if _, ok := resultFocusedByOtherMap[followRecord.UserID]; !ok {
				resultFocusedByOtherMap[followRecord.UserID] = 1
			} else {
				resultFocusedByOtherMap[followRecord.UserID] += 1
			}
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap, resultFocusedByOtherMap
}

/**
 * 查询用户推荐数据集数量、创建数据集数量、下载量最多数据集下载次数，下载最多的数据集名称，引用最多的数据集引用次数，引用次数最多的数据集名称
 */
func queryRecommedDataSet(start_unix int64, end_unix int64) (map[int64]int, map[int64]int, map[int64]int, map[int64]string, map[int64]int, map[int64]string, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	userIdRecommentDatasetMap := make(map[int64]int)
	userIdCreateDatasetMap := make(map[int64]int)

	userIdMostDownloadDatasetMap := make(map[int64]int)
	userIdMostDownloadDatasetNameMap := make(map[int64]string)
	userIdMostQuoteDatasetMap := make(map[int64]int)
	userIdMostQuoteDatasetNameMap := make(map[int64]string)
	privateDatasetMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(DatasetRegistry))
	if err != nil {
		log.Info("query recommend dataset error. return.")
		return userIdRecommentDatasetMap, userIdCreateDatasetMap, userIdMostDownloadDatasetMap, userIdMostDownloadDatasetNameMap, userIdMostQuoteDatasetMap, userIdMostQuoteDatasetNameMap, privateDatasetMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,creator_id,recommend,use_count,download_count,name,is_private,alias").Where(cond).Table(new(DatasetRegistry)).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		datasetList := make([]*DatasetRegistry, 0)
		sess.Find(&datasetList)
		log.Info("query datasetList size=" + fmt.Sprint(len(datasetList)))
		for _, datasetRecord := range datasetList {
			if datasetRecord.Recommend {
				if _, ok := userIdRecommentDatasetMap[datasetRecord.CreatorID]; !ok {
					userIdRecommentDatasetMap[datasetRecord.CreatorID] = 1
				} else {
					userIdRecommentDatasetMap[datasetRecord.CreatorID] += 1
				}
			}
			userIdCreateDatasetMap[datasetRecord.CreatorID] = getMapValue(datasetRecord.CreatorID, userIdCreateDatasetMap) + 1
			if int(datasetRecord.DownloadCount) > getMapValue(datasetRecord.CreatorID, userIdMostDownloadDatasetMap) {
				userIdMostDownloadDatasetMap[datasetRecord.CreatorID] = int(datasetRecord.DownloadCount)
				userIdMostDownloadDatasetNameMap[datasetRecord.CreatorID] = datasetRecord.DisplayName()
			}
			if int(datasetRecord.UseCount) > getMapValue(datasetRecord.CreatorID, userIdMostQuoteDatasetMap) {
				userIdMostQuoteDatasetMap[datasetRecord.CreatorID] = int(datasetRecord.UseCount)
				userIdMostQuoteDatasetNameMap[datasetRecord.CreatorID] = datasetRecord.DisplayName()
			}

			if datasetRecord.IsPrivate {
				privateDatasetMap[datasetRecord.CreatorID] = getMapValue(datasetRecord.CreatorID, privateDatasetMap) + 1
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return userIdRecommentDatasetMap, userIdCreateDatasetMap, userIdMostDownloadDatasetMap, userIdMostDownloadDatasetNameMap, userIdMostQuoteDatasetMap, userIdMostQuoteDatasetNameMap, privateDatasetMap
}

func queryAllDataSet() (map[int64]int64, map[int64]int64) {
	sess := x.NewSession()
	defer sess.Close()
	datasetUserIdMap := make(map[int64]int64)
	userIdDdatasetMap := make(map[int64]int64)
	count, err := sess.Count(new(DatasetRegistry))
	if err != nil {
		log.Info("query dataset error. return.")
		return datasetUserIdMap, userIdDdatasetMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,user_id").Table(new(Dataset)).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		datasetList := make([]*Dataset, 0)
		sess.Find(&datasetList)
		log.Info("query datasetList size=" + fmt.Sprint(len(datasetList)))
		for _, datasetRecord := range datasetList {
			datasetUserIdMap[datasetRecord.ID] = datasetRecord.UserID
			if _, ok := userIdDdatasetMap[datasetRecord.UserID]; !ok {
				userIdDdatasetMap[datasetRecord.UserID] = 1
			} else {
				userIdDdatasetMap[datasetRecord.UserID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return datasetUserIdMap, userIdDdatasetMap
}

func queryRecommedImage(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()
	userIdImageMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix) + " and type=5"
	count, err := sess.Where(cond).Count(new(Image))
	if err != nil {
		log.Info("query recommend image error. return.")
		return userIdImageMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uid,type").Where(cond).Table(new(Image)).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		imageList := make([]*Image, 0)
		sess.Find(&imageList)
		log.Info("query imageList size=" + fmt.Sprint(len(imageList)))
		for _, imageRecord := range imageList {
			if _, ok := userIdImageMap[imageRecord.UID]; !ok {
				userIdImageMap[imageRecord.UID] = 1
			} else {
				userIdImageMap[imageRecord.UID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return userIdImageMap
}

func queryDatasetStars(start_unix int64, end_unix int64) (map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	datasetCollect := make(map[int64]int)
	datasetCollected := make(map[int64]int)
	datasetUserIdMap, _ := queryAllDataSet()
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(DatasetStar))
	if err != nil {
		log.Info("query follow error. return.")
		return datasetCollect, datasetCollected
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uid,dataset_id").Table(new(DatasetStar)).Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		datasetStarList := make([]*DatasetStar, 0)
		sess.Find(&datasetStarList)
		log.Info("query datasetStarList size=" + fmt.Sprint(len(datasetStarList)))
		for _, datasetStarRecord := range datasetStarList {
			if _, ok := datasetCollect[datasetStarRecord.UID]; !ok {
				datasetCollect[datasetStarRecord.UID] = 1
			} else {
				datasetCollect[datasetStarRecord.UID] += 1
			}
			if _, ok := datasetCollected[datasetUserIdMap[datasetStarRecord.DatasetID]]; !ok {
				datasetCollected[datasetUserIdMap[datasetStarRecord.DatasetID]] = 1
			} else {
				datasetCollected[datasetUserIdMap[datasetStarRecord.DatasetID]] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return datasetCollect, datasetCollected
}

func queryImageStars(start_unix int64, end_unix int64) (map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	imageCollect := make(map[int64]int)
	imageCollected := make(map[int64]int)
	imageUserIdMap, _ := queryAllDataSet()
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(ImageStar))
	if err != nil {
		log.Info("query follow error. return.")
		return imageCollect, imageCollected
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uid,image_id").Table(new(ImageStar)).Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		imageStarList := make([]*ImageStar, 0)
		sess.Find(&imageStarList)
		log.Info("query imageStarList size=" + fmt.Sprint(len(imageStarList)))
		for _, imageStarRecord := range imageStarList {
			if _, ok := imageCollect[imageStarRecord.UID]; !ok {
				imageCollect[imageStarRecord.UID] = 1
			} else {
				imageCollect[imageStarRecord.UID] += 1
			}
			if _, ok := imageCollected[imageUserIdMap[imageStarRecord.ImageID]]; !ok {
				imageCollected[imageUserIdMap[imageStarRecord.ImageID]] = 1
			} else {
				imageCollected[imageUserIdMap[imageStarRecord.ImageID]] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return imageCollect, imageCollected
}

func queryDatasetSize(start_unix int64, end_unix int64) (map[int64]int, map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	resultSizeMap := make(map[int64]int)
	resultNumMap := make(map[int64]int)
	resultDownloadMap := make(map[int64]int)
	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)

	count, err := sess.Where(cond).Count(new(Attachment))
	if err != nil {
		log.Info("query attachment error. return.")
		return resultSizeMap, resultNumMap, resultDownloadMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,uploader_id,size,download_count").Table("attachment").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		attachmentList := make([]*Attachment, 0)
		sess.Find(&attachmentList)

		log.Info("query Attachment size=" + fmt.Sprint(len(attachmentList)))
		for _, attachRecord := range attachmentList {
			if _, ok := resultSizeMap[attachRecord.UploaderID]; !ok {
				resultSizeMap[attachRecord.UploaderID] = int(attachRecord.Size / (1024 * 1024)) //MB
				resultNumMap[attachRecord.UploaderID] = 1
				resultDownloadMap[attachRecord.UploaderID] = int(attachRecord.DownloadCount)
			} else {
				resultSizeMap[attachRecord.UploaderID] += int(attachRecord.Size / (1024 * 1024)) //MB
				resultNumMap[attachRecord.UploaderID] += 1
				resultDownloadMap[attachRecord.UploaderID] += int(attachRecord.DownloadCount)
			}
		}

		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultSizeMap, resultNumMap, resultDownloadMap
}

func queryUserCreateRepo(start_unix int64, end_unix int64) (map[int64]int, map[string]int, map[int64]string, map[string]map[string]interface{}) {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)

	detailInfoMap := make(map[string]int)
	mostDownloadMap := make(map[int64]string)

	fourSeasonMap := make(map[string]map[string]interface{})

	cond := "created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Count(new(Repository))
	if err != nil {
		log.Info("query Repository error. return.")
		return resultMap, detailInfoMap, mostDownloadMap, fourSeasonMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,owner_id,name,is_private,is_fork,clone_cnt,alias,created_unix").Table("repository").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		repoList := make([]*Repository, 0)
		sess.Find(&repoList)
		log.Info("query Repository size=" + fmt.Sprint(len(repoList)))
		for _, repoRecord := range repoList {
			resultMap[repoRecord.OwnerID] = getMapValue(repoRecord.OwnerID, resultMap) + 1
			key := fmt.Sprint(repoRecord.OwnerID) + "_total"
			if !repoRecord.IsFork {
				detailInfoMap[key] = getMapKeyStringValue(key, detailInfoMap) + 1

			}

			if repoRecord.IsPrivate {

				key := fmt.Sprint(repoRecord.OwnerID) + "_is_private"
				detailInfoMap[key] = getMapKeyStringValue(key, detailInfoMap) + 1

			} else {

				key := fmt.Sprint(repoRecord.OwnerID) + "_is_public"
				detailInfoMap[key] = getMapKeyStringValue(key, detailInfoMap) + 1

			}
			if repoRecord.IsFork {
				key := fmt.Sprint(repoRecord.OwnerID) + "_is_fork"
				detailInfoMap[key] = getMapKeyStringValue(key, detailInfoMap) + 1
			}
			key = fmt.Sprint(repoRecord.OwnerID) + "_total_download"
			detailInfoMap[key] = getMapKeyStringValue(key, detailInfoMap) + int(repoRecord.CloneCnt)

			key = fmt.Sprint(repoRecord.OwnerID) + "_most_download"
			if int(repoRecord.CloneCnt) > getMapKeyStringValue(key, detailInfoMap) {
				detailInfoMap[key] = int(repoRecord.CloneCnt)
				mostDownloadMap[repoRecord.OwnerID] = repoRecord.DisplayName()
			}
			setFourSeasonData(repoRecord, fourSeasonMap)
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}

	return resultMap, detailInfoMap, mostDownloadMap, fourSeasonMap
}

func setFourSeasonData(repoRecord *Repository, fourSeason map[string]map[string]interface{}) {
	key := ""
	switch repoRecord.CreatedUnix.AsTime().Month() {
	case time.January, time.February, time.March:
		key = fmt.Sprint(repoRecord.OwnerID) + "_spring"
	case time.April, time.May, time.June:
		key = fmt.Sprint(repoRecord.OwnerID) + "_summer"
	case time.July, time.August, time.September:
		key = fmt.Sprint(repoRecord.OwnerID) + "_autumn"
	case time.October, time.November, time.December:
		key = fmt.Sprint(repoRecord.OwnerID) + "_winter"
	default:
		log.Info("no found")
	}
	repoInfo := make(map[string]interface{})
	repoInfo["time"] = repoRecord.CreatedUnix
	repoInfo["displayName"] = repoRecord.DisplayName()
	if _, ok := fourSeason[key]; !ok {
		fourSeason[key] = repoInfo
	}
}

func queryUserRepoOpenIIndex(start_unix int64, end_unix int64) map[int64]float64 {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	statictisSess.Select("id,repo_id,radar_total").Table("repo_statistic").Where("created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)).OrderBy("id desc")
	repoStatisticList := make([]*RepoStatistic, 0)
	statictisSess.Find(&repoStatisticList)
	repoOpenIIndexMap := make(map[int64]float64)
	log.Info("query repo_statistic size=" + fmt.Sprint(len(repoStatisticList)))
	for _, repoRecord := range repoStatisticList {
		if _, ok := repoOpenIIndexMap[repoRecord.RepoID]; !ok {
			repoOpenIIndexMap[repoRecord.RepoID] = repoRecord.RadarTotal
		}
	}

	sess := x.NewSession()
	defer sess.Close()
	sess.Select("id,owner_id,name").Table("repository").Where("is_fork=false")
	repoList := make([]*Repository, 0)
	sess.Find(&repoList)

	userMap := make(map[int64]float64)

	log.Info("query Repository size=" + fmt.Sprint(len(repoList)))
	for _, repoRecord := range repoList {
		if _, ok := userMap[repoRecord.OwnerID]; !ok {
			if _, ok := repoOpenIIndexMap[repoRecord.ID]; ok {
				userMap[repoRecord.OwnerID] = repoOpenIIndexMap[repoRecord.ID]
			}
		}
	}

	//query collaboration
	sess.Select("repo_id,user_id,mode").Table("collaboration")
	collaborationList := make([]*Collaboration, 0)
	sess.Find(&collaborationList)

	log.Info("query collaborationList size=" + fmt.Sprint(len(collaborationList)))

	for _, collaborationRecord := range collaborationList {
		if _, ok := userMap[collaborationRecord.UserID]; !ok {
			if _, ok := repoOpenIIndexMap[collaborationRecord.RepoID]; ok {
				userMap[collaborationRecord.UserID] = repoOpenIIndexMap[collaborationRecord.RepoID]
			}
		} else {
			if _, ok := repoOpenIIndexMap[collaborationRecord.RepoID]; ok {
				userMap[collaborationRecord.UserID] += repoOpenIIndexMap[collaborationRecord.RepoID]
			}
		}
	}

	log.Info("user openi index size=" + fmt.Sprint(len(userMap)))

	return userMap
}

func queryLoginCountForYear(start_unix int64, end_unix int64) (map[int64]int, map[int64]map[string]int, []int) {
	resultMap, userLoginLogMap := queryLoginCount(start_unix, end_unix)
	return resultMap, userLoginLogMap, getValuesSortedDesc(resultMap)

}

func queryLoginCount(start_unix int64, end_unix int64) (map[int64]int, map[int64]map[string]int) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	resultMap := make(map[int64]int)
	userLoginLogMap := make(map[int64]map[string]int)
	cond := "created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := statictisSess.Where(cond).Count(new(UserLoginLog))
	if err != nil {
		log.Info("query UserLoginLog error. return.")
		return resultMap, userLoginLogMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		statictisSess.Select("id,u_id,created_unix").Table("user_login_log").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userLoginLogList := make([]*UserLoginLog, 0)
		statictisSess.Find(&userLoginLogList)
		log.Info("query user login size=" + fmt.Sprint(len(userLoginLogList)))
		for _, loginRecord := range userLoginLogList {
			if _, ok := resultMap[loginRecord.UId]; !ok {
				resultMap[loginRecord.UId] = 1
			} else {
				resultMap[loginRecord.UId] += 1
			}
			month := getMonth(loginRecord.CreatedUnix)
			if _, ok := userLoginLogMap[loginRecord.UId]; !ok {
				userLoginLogMap[loginRecord.UId] = make(map[string]int)
				userLoginLogMap[loginRecord.UId][month] = 1
			} else {
				if _, ok := userLoginLogMap[loginRecord.UId][month]; !ok {
					userLoginLogMap[loginRecord.UId][month] = 1
				} else {
					userLoginLogMap[loginRecord.UId][month] += 1
				}
			}

		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	log.Info("user login size=" + fmt.Sprint(len(resultMap)))
	return resultMap, userLoginLogMap
}

func getValuesSortedDesc(m map[int64]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	// 从大到小排序
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values
}

func queryLoginActionCount(start_unix int64, end_unix int64) map[int64]int {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	resultMap := make(map[int64]int)
	cond := "created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := statictisSess.Where(cond).Count(new(UserLoginActionLog))
	if err != nil {
		log.Info("query UserLoginActionLog error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	cachemap := make(map[string]int)
	for {
		statictisSess.Select("id,u_id,created_unix").Table("user_login_action_log").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		userLoginActionLogList := make([]*UserLoginActionLog, 0)
		statictisSess.Find(&userLoginActionLogList)
		log.Info("query user login action size=" + fmt.Sprint(len(userLoginActionLogList)))

		for _, loginRecord := range userLoginActionLogList {
			strkey := loginRecord.CreatedUnix.AsTimePtrInLocation(time.Now().Location()).Format("2006-01-02") + strconv.FormatInt(loginRecord.UId, 10)
			if _, ok := cachemap[strkey]; !ok {
				if _, ok := resultMap[loginRecord.UId]; !ok {
					resultMap[loginRecord.UId] = 1
				} else {
					resultMap[loginRecord.UId] += 1
				}
				cachemap[strkey] = 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	log.Info("user login action size=" + fmt.Sprint(len(resultMap)))
	return resultMap
}

func queryDatasetCollection(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()

	datasetCollect := make(map[int64]int)

	rows := []struct {
		UserId int64 `xorm:"user_id"`
		Cnt    int   `xorm:"cnt"`
	}{}

	err := sess.Table("dataset_collection").
		Select("user_id, COUNT(DISTINCT id) AS cnt").
		Where("created_unix BETWEEN ? AND ?", start_unix, end_unix).
		GroupBy("user_id").
		Find(&rows)
	if err != nil {
		log.Warn("queryDatasetCollection failed: ", err)
		return datasetCollect
	}

	for _, row := range rows {
		datasetCollect[row.UserId] = row.Cnt
	}

	return datasetCollect
}

func queryUserDatasetRegistry(start_unix int64, end_unix int64) (map[int64]int, map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()
	datasetCnt := make(map[int64]int)
	recommendedCnt := make(map[int64]int)
	collectedCnt := make(map[int64]int)

	// 用户创建数据集,被推荐,被收藏的次数统计
	rows := []struct {
		UserId         int64 `xorm:"user_id"`
		Cnt            int   `xorm:"cnt"`
		RecommendedCnt int   `xorm:"recommended_cnt"`
		CollectedCnt   int   `xorm:"collected_cnt"`
	}{}
	err := sess.Table("dataset_registry").
		Select("owner_id AS user_id, COUNT(*) AS cnt, SUM(CASE WHEN recommend is true THEN 1 ELSE 0 END) AS recommended_cnt, SUM(num_collections) AS collected_cnt").
		Where("created_unix >= ? AND created_unix <= ?", start_unix, end_unix).
		GroupBy("owner_id").
		Find(&rows)
	if err != nil {
		log.Info("query queryUserDatasetRegistry error. return.")
		return datasetCnt, recommendedCnt, collectedCnt
	}

	for _, row := range rows {
		datasetCnt[row.UserId] = row.Cnt
		recommendedCnt[row.UserId] = row.RecommendedCnt
		collectedCnt[row.UserId] = row.CollectedCnt
	}
	return datasetCnt, collectedCnt, recommendedCnt
}

func queryAimodelCollection(start_unix int64, end_unix int64) map[int64]int {
	sess := x.NewSession()
	defer sess.Close()

	aimodelCollect := make(map[int64]int)

	rows := []struct {
		UserId int64 `xorm:"user_id"`
		Cnt    int   `xorm:"cnt"`
	}{}

	err := sess.Table("ai_model_collect").
		Select("user_id, COUNT(DISTINCT id) AS cnt").
		Where("created_unix BETWEEN ? AND ?", start_unix, end_unix).
		GroupBy("user_id").
		Find(&rows)
	if err != nil {
		log.Warn("queryAimodelCollection failed: ", err)
		return aimodelCollect
	}

	for _, row := range rows {
		aimodelCollect[row.UserId] = row.Cnt
	}

	return aimodelCollect
}

func queryUserAimodel(start_unix int64, end_unix int64) (map[int64]int, map[int64]int, map[int64]int) {
	sess := x.NewSession()
	defer sess.Close()

	aimodelCnt := make(map[int64]int)
	recommendedCnt := make(map[int64]int)
	collectedCnt := make(map[int64]int)

	rows := []struct {
		UserId         int64 `xorm:"user_id"`
		Cnt            int   `xorm:"cnt"`
		RecommendedCnt int   `xorm:"recommended_cnt"`
		CollectedCnt   int   `xorm:"collected_cnt"`
	}{}

	err := sess.Table("ai_model_manage").
		Select("owner_id AS user_id, COUNT(*) AS cnt, SUM(recommend) AS recommended_cnt, SUM(collected_count) AS collected_cnt").
		Where("created_unix >= ? AND created_unix <= ?", start_unix, end_unix).
		GroupBy("owner_id").
		Find(&rows)
	if err != nil {
		log.Info("query queryUserAimodel error. return.")
		return aimodelCnt, recommendedCnt, collectedCnt
	}

	for _, row := range rows {
		aimodelCnt[row.UserId] = row.Cnt
		recommendedCnt[row.UserId] = row.RecommendedCnt
		collectedCnt[row.UserId] = row.CollectedCnt
	}
	// add print to see if the query results in row is correctly added into three maps
	log.Info("queryUserAimodel result: ", rows)
	log.Info("aimodelCnt: ", aimodelCnt)
	log.Info("recommendedCnt: ", recommendedCnt)
	log.Info("collectedCnt: ", collectedCnt)

	return aimodelCnt, collectedCnt, recommendedCnt
}

func queryUserModel(start_unix int64, end_unix int64) map[int64]int {
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
		sess.Select("id,owner_id").Table("ai_model_manage").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		aiModelList := make([]*AiModelManage, 0)
		sess.Find(&aiModelList)
		log.Info("query AiModelManage size=" + fmt.Sprint(len(aiModelList)))
		for _, aiModelRecord := range aiModelList {
			if _, ok := resultMap[aiModelRecord.UserId]; !ok {
				resultMap[aiModelRecord.UserId] = 1
			} else {
				resultMap[aiModelRecord.UserId] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func queryUserModelConvert(start_unix int64, end_unix int64) map[int64]int {
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
		sess.Select("id,user_id").Table("ai_model_convert").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		aiModelList := make([]*AiModelConvert, 0)
		sess.Find(&aiModelList)
		log.Info("query AiModelConvert size=" + fmt.Sprint(len(aiModelList)))
		for _, aiModelRecord := range aiModelList {
			if _, ok := resultMap[aiModelRecord.UserId]; !ok {
				resultMap[aiModelRecord.UserId] = 1
			} else {
				resultMap[aiModelRecord.UserId] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap
}

func isBenchMark(JobType string) bool {
	if JobType == "BENCHMARK" || JobType == "MODELSAFETY" || JobType == "SNN4IMAGENET" || JobType == "BRAINSCORE" || JobType == "SNN4ECOSET" || JobType == "SIM2BRAIN_SNN" {
		return true
	}
	return false
}

func queryCloudBrainTask(start_unix int64, end_unix int64) (map[int64]int, map[string]int, map[int64]map[string]int, map[int64]map[string]int) {
	sess := x.NewSession()
	defer sess.Close()
	resultMap := make(map[int64]int)
	resultItemMap := make(map[string]int)
	resourceItemMap := make(map[int64]map[string]int)
	jobTypeItemMap := make(map[int64]map[string]int)

	cond := " created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := sess.Where(cond).Unscoped().Count(new(Cloudbrain))
	if err != nil {
		log.Info("query cloudbrain error. return.")
		return resultMap, resultItemMap, resourceItemMap, jobTypeItemMap
	}
	log.Info("cloudbrain count=" + fmt.Sprint(count))
	var indexTotal int64
	indexTotal = 0
	for {
		sess.Select("id,job_type,user_id,duration,train_job_duration,type,compute_resource").Table("cloudbrain").Unscoped().Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
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

			if cloudTaskRecord.Duration < 100000000 && cloudTaskRecord.Duration > 0 {
				setMapKey("CloudBrainRunTime", cloudTaskRecord.UserID, int(cloudTaskRecord.Duration), resultItemMap)
				resourceItemMap[cloudTaskRecord.UserID][cloudTaskRecord.ComputeResource] = resourceItemMap[cloudTaskRecord.UserID][cloudTaskRecord.ComputeResource] + int(cloudTaskRecord.Duration)
			}

			if _, ok := jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType]; !ok {
				jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType] = 1
			} else {
				jobTypeItemMap[cloudTaskRecord.UserID][cloudTaskRecord.JobType] += 1
			}

			if cloudTaskRecord.Type == 1 { //npu
				setMapKey("CloudBrainTwo", cloudTaskRecord.UserID, 1, resultItemMap)
				if cloudTaskRecord.JobType == "TRAIN" {
					setMapKey("NpuTrainJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else if cloudTaskRecord.JobType == "INFERENCE" {
					setMapKey("NpuInferenceJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else if isBenchMark(cloudTaskRecord.JobType) {
					setMapKey("GpuBenchMarkJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else {
					setMapKey("NpuDebugJob", cloudTaskRecord.UserID, 1, resultItemMap)
				}
			} else if cloudTaskRecord.Type == 0 { //type=0  gpu
				setMapKey("CloudBrainOne", cloudTaskRecord.UserID, 1, resultItemMap)
				if cloudTaskRecord.JobType == "TRAIN" {
					setMapKey("GpuTrainJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else if cloudTaskRecord.JobType == "INFERENCE" {
					setMapKey("GpuInferenceJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else if isBenchMark(cloudTaskRecord.JobType) {
					setMapKey("GpuBenchMarkJob", cloudTaskRecord.UserID, 1, resultItemMap)
				} else {
					setMapKey("GpuDebugJob", cloudTaskRecord.UserID, 1, resultItemMap)
				}
			} else if cloudTaskRecord.Type == 2 || cloudTaskRecord.Type == 3 {
				setMapKey("C2Net", cloudTaskRecord.UserID, 1, resultItemMap)
				if cloudTaskRecord.ComputeResource == NPUResource {
					if cloudTaskRecord.JobType == "TRAIN" {
						setMapKey("NpuTrainJob", cloudTaskRecord.UserID, 1, resultItemMap)
					} else {
						setMapKey("NpuDebugJob", cloudTaskRecord.UserID, 1, resultItemMap)
					}
				} else {
					if cloudTaskRecord.JobType == "TRAIN" {
						setMapKey("GpuTrainJob", cloudTaskRecord.UserID, 1, resultItemMap)
					} else if cloudTaskRecord.JobType == "ONLINEINFERENCE" {
						setMapKey("GpuInferenceJob", cloudTaskRecord.UserID, 1, resultItemMap)
					} else {
						setMapKey("GpuDebugJob", cloudTaskRecord.UserID, 1, resultItemMap)
					}
				}
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	return resultMap, resultItemMap, resourceItemMap, jobTypeItemMap
}

func queryUserInvitationCount(start_unix int64, end_unix int64) map[int64]int {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	resultMap := make(map[int64]int)
	cond := "created_unix>=" + fmt.Sprint(start_unix) + " and created_unix<=" + fmt.Sprint(end_unix)
	count, err := statictisSess.Where(cond).Count(new(Invitation))
	if err != nil {
		log.Info("query queryUserInvitationCount error. return.")
		return resultMap
	}
	var indexTotal int64
	indexTotal = 0
	for {
		statictisSess.Select("id,src_user_id,user_id").Table("invitation").Where(cond).OrderBy("id asc").Limit(PAGE_SIZE, int(indexTotal))
		invitationList := make([]*Invitation, 0)
		statictisSess.Find(&invitationList)
		log.Info("query invitationList size=" + fmt.Sprint(len(invitationList)))
		for _, invitationRecord := range invitationList {
			if _, ok := resultMap[invitationRecord.SrcUserID]; !ok {
				resultMap[invitationRecord.SrcUserID] = 1
			} else {
				resultMap[invitationRecord.SrcUserID] += 1
			}
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	log.Info("invitationList size=" + fmt.Sprint(len(resultMap)))
	return resultMap
}

func setMapKey(key string, userId int64, value int, resultItemMap map[string]int) {
	newKey := fmt.Sprint(userId) + "_" + key
	if _, ok := resultItemMap[newKey]; !ok {
		resultItemMap[newKey] = value
	} else {
		resultItemMap[newKey] += value
	}
}

func subMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	if month == 0 {
		month = 1
	}
	return month
}

func GetContentFromPromote(url string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("not error.", err)
			return
		}
	}()
	resp, err := http.Get(url)
	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return "", err
	}
	defer resp.Body.Close()
	if resp == nil {
		log.Info("respone is null")
		return "", errors.New("resp is null")
	}
	if resp.StatusCode != 200 {
		log.Info("respone code=" + fmt.Sprint(resp.StatusCode))
		return "", errors.New("resp is null")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return "", err
	}
	allLineStr := string(bytes)
	return allLineStr, nil
}

func QueryLast30DaysHighestIndexUsers(size int) ([]int64, error) {
	userIds := make([]int64, 0)
	err := xStatistic.Table("user_business_analysis_last30_day").Cols("id").OrderBy("user_index desc").Limit(size).Find(&userIds)
	return userIds, err
}
