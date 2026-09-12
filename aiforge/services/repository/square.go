package repository

import (
	"encoding/json"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/patrickmn/go-cache"
)

var repoSquareCache = cache.New(2*time.Minute, 1*time.Minute)

const (
	RREFERED_CACHE    = "PreferredRepos"
	REPO_BANNER_CACHE = "RepoBanner"
	TOPICS_CACHE      = "RepoTopics"
	RECOMMEND_CACHE   = "RecommendRepos"
)

func GetBanners() []map[string]string {
	v, success := repoSquareCache.Get(REPO_BANNER_CACHE)
	if success {
		log.Debug("GetBanners from cache,value = %v", v)
		if v == nil {
			return nil
		}
		r := v.([]map[string]string)
		return r
	}
	repoMap := getMapContent("repos/square_banner")
	repoSquareCache.Set(REPO_BANNER_CACHE, repoMap, 1*time.Minute)
	return repoMap
}

func GetTopics() []string {
	v, success := repoSquareCache.Get(TOPICS_CACHE)
	if success {
		log.Debug("GetTopics from cache,value = %v", v)
		if v == nil {
			return nil
		}
		r := v.([]string)
		return r
	}
	topics := getArrayContent("repos/recommend_topics")
	repoSquareCache.Set(TOPICS_CACHE, topics, 1*time.Minute)
	return topics
}

func getMapContent(fileName string) []map[string]string {
	url := setting.RecommentRepoAddr + fileName
	result, err := RecommendContentFromPromote(url)
	remap := make([]map[string]string, 0)
	if err == nil {
		json.Unmarshal([]byte(result), &remap)
	}
	return remap
}

func getArrayContent(fileName string) []string {
	url := setting.RecommentRepoAddr + fileName
	result, err := RecommendContentFromPromote(url)
	r := make([]string, 0)
	if err == nil {
		json.Unmarshal([]byte(result), &r)
	}
	return r
}

func GetRecommendRepos() []map[string]interface{} {
	v, success := repoSquareCache.Get(RECOMMEND_CACHE)
	if success {
		log.Debug("GetRecommendRepos from cache,value = %v", v)
		if v == nil {
			return nil
		}
		r := v.([]map[string]interface{})
		return r
	}
	repoMap := getMapContent("home/projects")
	r, _ := GetRecommendRepoFromPromote(repoMap)
	repoSquareCache.Set(RECOMMEND_CACHE, r, 1*time.Minute)
	return r
}

func GetPreferredRepos() ([]*models.Repository4Card, error) {
	v, success := repoSquareCache.Get(RREFERED_CACHE)
	if success {
		log.Debug("GetPreferredRepos from cache,value = %v", v)
		if v == nil {
			return nil, nil
		}
		r := v.([]*models.Repository4Card)
		return r, nil
	}

	repos, err := models.GetSelectedRepos(models.FindSelectedReposOpts{
		ListOptions: models.ListOptions{
			PageSize: 10,
			Page:     1,
		},
		OnlyPublic: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Repository4Card, len(repos))
	for i, r := range repos {
		result[i] = r.ToCardFormat()
	}

	repoSquareCache.Set(RREFERED_CACHE, result, 1*time.Minute)
	return result, nil
}

func GetIncubationRepos() ([]*models.Repository4Card, error) {
	org, err := models.GetOrgByName(setting.IncubationSourceOrgName)
	if models.IsErrOrgNotExist(err) {
		return make([]*models.Repository4Card, 0), nil
	}
	if err != nil {
		return nil, err
	}
	repos, err := models.GetSelectedRepos(models.FindSelectedReposOpts{
		ListOptions: models.ListOptions{
			PageSize: 10,
			Page:     1,
		},
		OrgId:      org.ID,
		OnlyPublic: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Repository4Card, len(repos))
	for i, r := range repos {
		result[i] = r.ToCardFormat()
	}
	return result, nil
}

func GetHotPaperRepos() ([]*models.Repository4Card, error) {
	rlist, _, err := models.SearchRepository(&models.SearchRepoOptions{
		ListOptions: models.ListOptions{
			Page:     1,
			PageSize: 10,
		},
		OrderBy:   models.SearchOrderByLastMonthVisitsReverse + "," + models.SearchOrderByRecentUpdated,
		TopicOnly: true,
		TopicName: setting.PaperRepoTopicName,
		AllPublic: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Repository4Card, len(rlist))
	for i, r := range rlist {
		result[i] = r.ToCardFormat()
	}
	return result, nil
}

type FindReposOptions struct {
	models.ListOptions
	Actor            *models.User
	Sort             string
	Keyword          string
	Topics           []string
	IsTopicIntersect bool
	Private          bool
	OwnerID          int64
	RepoIds          []int64
}

func FindRepos(opts FindReposOptions) (*models.FindReposResponse, error) {

	var (
		repos   []*models.Repository
		count   int64
		err     error
		orderBy models.SearchOrderBy
	)

	switch opts.Sort {
	//1.近期热门：按最近1个月浏览量倒序排序，最近1个月浏览量>最近更新>项目名称升序
	case "mostpopular":
		orderBy = models.SearchOrderByLastMonthVisitsReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//2.近期活跃：按提交增长量（最近4个月commit数）倒序排序，提交增长量>最近更新>项目名称升序。
	case "mostactive":
		orderBy = models.SearchOrderByLastFourMonthCommitsReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//3.最近更新：按最近更新>项目名称升序排序。
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//4.最近创建：按项目创建时间排序，最近的排前面。最近创建>项目名称升序。
	case "newest":
		orderBy = models.SearchOrderByNewest + "," + models.SearchOrderByAlphabetically
		//5.点赞最多：按点赞数倒序排序。点赞数>最近更新>项目名称升序。
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//6.派生最多：按派生数倒序排序。派生数>最近更新>项目名称升序。
	case "mostforks":
		orderBy = models.SearchOrderByForksReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//7.数据集最多：按项目包含的数据集文件数量倒序排序，数据集文件数>最近更新>项目名称升序。
	case "mostdatasets":
		orderBy = models.SearchOrderByDatasetCntReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//8.AI任务最多：按项目包含的AI任务数量倒序排序，AI任务数>最近更新>项目名称升序。
	case "mostaitasks":
		orderBy = models.SearchOrderByAiTaskCntReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
		//9.模型最多：按项目包含的模型数量倒序排序，模型大小为0则不统计。模型数>最近更新>项目名称升序。
	case "mostmodels":
		orderBy = models.SearchOrderByModelCntReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically

	default:
		orderBy = models.SearchOrderByLastMonthVisitsReverse + "," + models.SearchOrderByRecentUpdated + "," + models.SearchOrderByAlphabetically
	}

	repos, count, err = models.SearchRepository(&models.SearchRepoOptions{
		ListOptions:        opts.ListOptions,
		Actor:              opts.Actor,
		OrderBy:            orderBy,
		Private:            opts.Private,
		Keyword:            opts.Keyword,
		OwnerID:            opts.OwnerID,
		AllPublic:          true,
		AllLimited:         true,
		MultiTopicNames:    opts.Topics,
		IsTopicIntersect:   opts.IsTopicIntersect,
		IncludeDescription: setting.UI.SearchRepoDescription,
		RepoIds:            opts.RepoIds,
	})
	if err != nil {
		log.Error("FindRepos error when SearchRepository.%v", err)
		return nil, err
	}
	result := make([]*models.Repository4Card, len(repos))
	for i, r := range repos {
		t := r.ToCardFormat()
		result[i] = t
	}

	return &models.FindReposResponse{
		Repos:    result,
		Total:    count,
		Page:     opts.Page,
		PageSize: opts.PageSize,
	}, nil
}

type ActiveUser struct {
	User       *models.User4Front
	Followed   bool
	ShowButton bool
}

func GetActiveUser4Square(currentUserId int64) ([]*ActiveUser, error) {
	result := make([]*ActiveUser, 0)
	userIds, err := models.QueryLast30DaysHighestIndexUsers(5)
	if err != nil {
		log.Error("ActiveUser err. %v", err)
		return result, err
	}
	if len(userIds) == 0 {
		return result, nil
	}

	users, err := models.GetUsersByIDs(userIds)
	if err != nil {
		return result, nil
	}
	usersMap := make(map[int64]*models.User)
	for _, v := range users {
		usersMap[v.ID] = v
	}

	for i := 0; i < len(userIds); i++ {
		userId := userIds[i]
		user := usersMap[userId]
		if user == nil {
			continue
		}
		isFollowed := false
		if currentUserId != 0 {
			isFollowed = models.IsFollowing(currentUserId, userId)
		}
		a := &ActiveUser{
			Followed:   isFollowed,
			User:       user.ToFrontFormat(),
			ShowButton: currentUserId != userId,
		}
		result = append(result, a)
	}
	return result, nil
}

func GetActiveOrgs() ([]*models.User4Front, error) {
	orgScores, err := models.FindTopNOpenIOrgs(5)
	if err != nil {
		return nil, err
	}
	orgs := make([]*models.User4Front, len(orgScores))
	for i, v := range orgScores {
		orgs[i] = v.ToFrontFormat()
	}
	return orgs, nil
}

func RefreshRepoStatData() {
	repos, err := models.GetAllRepositories()
	if err != nil {
		log.Error("RefreshRepoStatData GetAllRepositories failed: %v", err.Error())
		return
	}
	for _, repo := range repos {
		models.SyncStatDataToRepo(repo)
	}
}
