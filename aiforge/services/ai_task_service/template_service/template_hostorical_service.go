package template_service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/generate"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/timeutil"
)

func HandleHistoricalAITaskTemplate(forceFull bool) (int, error) {
	lockKey := "ai_task_template:historical:lock"
	lock := redis_lock.NewDistributeLock(lockKey)

	if ok, err := lock.Lock(24 * time.Hour); !ok {
		log.Error("HandleHistoricalAITaskTemplate try get lock err.%v", err)
		return 0, err
	}
	defer lock.UnLock()

	path := ".openi-aitemplate"

	var currentUnix int64
	var lastUnix timeutil.TimeStamp
	lastUnix = 1718012867

	lastTimeKey := "ai_task_template:historical:last_unix"
	if !forceFull {
		lastTimeStr, _ := redis_client.Get(lastTimeKey)
		if lastTimeStr != "" {
			tempUnix, _ := strconv.ParseInt(lastTimeStr, 10, 64)
			if tempUnix > int64(lastUnix) {
				lastUnix = timeutil.TimeStamp(tempUnix)
			}
		}
	}

	page := 0
	pageSize := 100
	count := 0
	for {
		page++
		log.Info("HandleHistoricalAITaskTemplate go into page %d pageSize =%d current count=%d", page, pageSize, count)
		//1、根据更新时间遍历查询所有仓库，每一个时间戳之前的都存下来
		//select count(*) from repository  where updated_unix > 1718012867 and is_empty = false and size > 0 and is_mirror = false and ai_task_cnt > 0
		repos, err := models.GetHistoricalTemplatePaging(page, pageSize, lastUnix)
		if err != nil {
			log.Error("HandleHistoricalAITaskTemplate GetHistoricalTemplatePaging err. count=%d err=%v", count, err)
			return count, err
		}
		if len(repos) == 0 {
			log.Info("HandleHistoricalAITaskTemplate len(repos) == 0,page %d pageSize =%d", page, pageSize)
			return count, nil
		}

		for i := 0; i < len(repos); i++ {
			repo := repos[i]
			if !forceFull && int64(repo.UpdatedUnix) > currentUnix {
				currentUnix = int64(repo.UpdatedUnix) - 1
				redis_client.Setex(lastTimeKey, fmt.Sprint(currentUnix), 30*24*time.Hour)
			}
			//2、读取仓库文件master分支中是否存在模板文件
			content, users, createdAt, updatedAt, err := models.ReadLatestFileInRepoWithValidCommiter(repo, path, 100)
			if err != nil {
				log.Error("HandleHistoricalAITaskTemplate ReadLatestFileInRepo failed,repo.OwnerName=%s, repo.Name=%s,  path=%s, error=%v", repo.OwnerName, repo.Name, path, err)
				continue
			}
			if users == nil || len(users) == 0 {
				log.Info("HandleHistoricalAITaskTemplate find template file but no valid commiters. repo.OwnerName=%s, repo.Name=%s", repo.OwnerName, repo.Name)
				continue
			}
			log.Info("HandleHistoricalAITaskTemplate find template file repo.OwnerName=%s, repo.Name=%s", repo.OwnerName, repo.Name)

			//3、如果有的话查出有权限的用户名单
			userIds := make([]int64, 0, len(users))
			for _, user := range users {
				userIds = append(userIds, user.ID)
			}

			//4、删所有有权限用户下该项目的模板并重新插入模板文件中的模板
			newTemplates, err := parseTemplates(content, repo, createdAt, updatedAt)
			if err != nil {
				log.Error("HandleHistoricalAITaskTemplate parseTemplates failed,repo.OwnerName=%s, repo.Name=%s,path=%s, error=%v", repo.OwnerName, repo.Name, path, err)
				continue
			}
			err = models.DeleteUserTemplateAndAddNew(repo, userIds, newTemplates)
			if err != nil {
				log.Error("HandleHistoricalAITaskTemplate parseTemplates failed,repo.OwnerName=%s, repo.Name=%s,  path=%s, error=%v", repo.OwnerName, repo.Name, path, err)
				continue
			}

			count++
		}
		if len(repos) < pageSize {
			log.Info("HandleHistoricalAITaskTemplate parseTemplates finished,count = %d", count)
			break

		}
	}

	return count, nil
}

func parseTemplates(content string, repo models.Repository, createdAt, updatedAt *time.Time) ([]models.AITaskTemplate, error) {
	var oldTemplates = make([]OldTaskTemplate, 0)
	var newTemplates = make([]models.AITaskTemplate, 0)
	err := json.Unmarshal([]byte(content), &oldTemplates)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(oldTemplates); i++ {
		t := oldTemplates[i]
		var parameterStr string
		b, err := json.Marshal(t.Parameters)
		if err == nil && len(b) > 0 {
			parameterStr = string(b)
		}
		var datasetList = make([]*models.TemplateDatasetInfo, 0)
		for j := 0; j < len(t.DatasetList); j++ {
			d := t.DatasetList[j]
			datasetList = append(datasetList, &models.TemplateDatasetInfo{
				DatasetID:    d["ID"],
				DatasetName:  d["DatasetName"],
				DatasetAlias: d["DatasetAlias"],
				OwnerName:    d["OwnerName"],
			})
		}

		var modelList = make([]*models.TemplateModelInfo, 0)
		for k := 0; k < len(t.PretrainModelList); k++ {
			d := t.PretrainModelList[k]
			modelList = append(modelList, &models.TemplateModelInfo{
				ModelID:    d["ID"],
				ModelName:  d["ModelName"],
				ModelAlias: d["ModelAlias"],
				OwnerName:  d["OwnerName"],
			})
		}
		t.Name = makeValidAlias(t.Name)
		newTemplate := models.AITaskTemplate{
			Name:              t.Name,
			LowerName:         strings.ToLower(t.Name),
			Description:       t.Description,
			IsPrivate:         repo.IsPrivate,
			JobType:           t.JobType,
			Cluster:           t.Cluster,
			ComputeSource:     t.ComputeSource,
			HasInternet:       t.HasInternet,
			VisualizeRequired: t.VisualizeRequired,

			//spec
			AccCardsNum: t.AccCardsNum,
			AccCardType: t.AccCardType,
			CpuCores:    t.CpuCores,
			MemGiB:      t.MemGiB,
			GPUMemGiB:   t.GPUMemGiB,
			ShareMemGiB: t.ShareMemGiB,

			//repo
			RepoId:        repo.ID,
			RepoName:      repo.Name,
			RepoOwnerName: repo.OwnerName,
			BranchName:    t.BranchName,
			BootFile:      t.BootFile,

			ImageID:   t.Image["ImageID"],
			ImageName: t.Image["ImageName"],
			ImageUrl:  t.Image["ImageUrl"],

			Parameters:  parameterStr,
			ModelList:   modelList,
			DatasetList: datasetList,
		}
		if createdAt != nil {
			newTemplate.CreatedUnix = timeutil.TimeStamp(createdAt.Unix())
		} else {
			newTemplate.CreatedUnix = timeutil.TimeStampNow()
		}
		if updatedAt != nil && createdAt != nil {
			newTemplate.UpdatedUnix = timeutil.TimeStamp(updatedAt.Unix())
		} else {
			newTemplate.UpdatedUnix = timeutil.TimeStampNow()

		}
		newTemplates = append(newTemplates, newTemplate)

	}

	return newTemplates, nil
}

func makeValidAlias(s string) string {
	re := regexp.MustCompile(`[^\p{Han}A-Za-z0-9_.-]+`)
	s = re.ReplaceAllString(s, "_")
	re = regexp.MustCompile(`_+`)
	s = re.ReplaceAllString(s, "_")
	re = regexp.MustCompile(`-+`)
	s = re.ReplaceAllString(s, "-")
	re = regexp.MustCompile(`\.+`)
	s = re.ReplaceAllString(s, ".")
	s = strings.Trim(s, "_")

	if s == "" || s == "_" || s == "." || s == "-" {
		randStr, _ := generate.GetRandomString(3)
		if randStr == "" {
			randStr = "dfg"
		}
		randStr = strings.ToLower(randStr)
		s = "NewTempl" + randStr
	}
	runes := []rune(s)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	s = string(runes)
	return s
}

type OldTaskTemplate struct {
	Name        string `xorm:"INDEX"`
	Description string `xorm:"TEXT"`
	Recommend   bool   `xorm:"DEFAULT false"`

	JobType           string                   `xorm:"NOT NULL DEFAULT 'DEBUG'"`
	Cluster           string                   `xorm:"NOT NULL DEFAULT 'C2Net'"`
	ComputeSource     string                   `xorm:"NOT NULL DEFAULT 'GPU'"`
	HasInternet       int                      `xorm:"NOT NULL DEFAULT 0"`
	VisualizeRequired bool                     `xorm:"DEFAULT false"`
	Parameters        []map[string]interface{} `xorm:"text"`

	//spec
	AccCardsNum int
	AccCardType string
	CpuCores    int
	MemGiB      float32
	GPUMemGiB   float32
	ShareMemGiB float32

	//image
	Image             map[string]string
	PretrainModelList []map[string]string
	DatasetList       []map[string]string
	// ImageID   string `xorm:"INDEX"`
	// ImageName string
	// ImageUrl  string

	//repo
	BranchName string
	BootFile   string
}
