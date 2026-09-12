package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/markup/markdown"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/modelmanage"
	"code.gitea.io/gitea/services/repository"
)

const (
	// tplModelSquareIndex  = "model/square/index"
	tplModelSquareReadMe = "repo/modelmanage/readme"
	tplModelFileList     = "repo/modelmanage/filelist"
	tplModelSetting      = "repo/modelmanage/setting"
	tplModelEvolutionMap = "repo/modelmanage/evolution_map"
	tplModelMigrate      = "repo/modelmanage/migrate"
	tplModelMigrating    = "repo/modelmanage/migrating"
	README_FILE_NAME     = "README.md"
)

type ModelMap struct {
	Type            int //0:repo; 1:model
	IsParent        bool
	IsCurrent       bool
	RepoName        string
	RepoOwnerName   string
	RepoDisplayName string
	RepoId          int64
	IsCanOper       bool
	IsPrivate       bool
	Model           *models.AiModelManage
	Models4Parent   []*models.AiModelManage
	Next            []*ModelMap
}

// func ModelSquareTmpl(ctx *context.Context) {
// 	ctx.HTML(200, tplModelSquareIndex)
// }

func ModelSquareData(ctx *context.Context) {
	log.Info("ShowModel Square Info start.")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	isRecommend := ctx.QueryBool("recommend")
	isOnlineUrl := ctx.QueryBool("hasOnlineUrl")
	queryType := ctx.QueryInt("queryType")
	labelFilter := ctx.Query("label")
	frameFilterStr := ctx.Query("frame")
	orderBy := ctx.Query("orderBy")
	Namelike := ctx.Query("q")
	TypeStr := ctx.Query("type")
	needModelFile := ctx.QueryBool("needModelFile")
	frameFilterInt := -1
	if frameFilterStr != "" {
		frameFilterInt, _ = strconv.Atoi(frameFilterStr)
	}
	notNeedEmpty := ctx.QueryBool("notNeedEmpty")
	computeResourceFilter := ctx.Query("compute_resource")
	statusFinish := ctx.QueryBool("statusFinish")
	var IsQueryPrivate bool
	var user_id int64
	var IsQueryCollect bool
	var collected_user_id int64
	var repo_id int64
	var typeInt int
	var myExternal bool
	hasOnlineUrl := 0
	if isOnlineUrl {
		hasOnlineUrl = 1
	}
	typeInt = -1
	if TypeStr != "" {
		//typeInt, _ = strconv.Atoi(TypeStr)
	}
	myExternal = false
	modelStatus := -1
	if statusFinish {
		modelStatus = 0
	}
	if queryType == 1 {
		IsQueryPrivate = false
		user_id = 0
	} else if queryType == 2 {
		IsQueryPrivate = true
		if ctx.User == nil {
			log.Info("the user not login.")
			ctx.JSON(http.StatusOK, nil)
			return
		}
		user_id = ctx.User.ID
	} else if queryType == 3 {
		IsQueryCollect = true
		IsQueryPrivate = true
		user_id = 0
		if ctx.User == nil {
			log.Info("the user not login.")
			ctx.JSON(http.StatusOK, nil)
			return
		}
		collected_user_id = ctx.User.ID
	} else if queryType == 4 {
		if ctx.User == nil {
			log.Info("the user not login.")
			ctx.JSON(http.StatusOK, nil)
			return
		}
		IsQueryPrivate = true
		repoName := ctx.Query("repoName")
		repoOwnerName := ctx.Query("repoOwnerName")
		repo, err := models.GetRepositoryByOwnerAndName(repoOwnerName, repoName)
		if err == nil {
			repo_id = repo.ID
		} else {
			log.Info("the repo is not exist.repoName=" + repoName + " repoOwnerName=" + repoOwnerName)
			ctx.JSON(http.StatusOK, nil)
			return
		}
	} else if queryType == 5 {
		if ctx.User == nil {
			log.Info("the user not login.")
			ctx.JSON(http.StatusOK, nil)
			return
		}
		user_id = ctx.User.ID
		myExternal = true
	} else {
		log.Info("not support")
		ctx.JSON(http.StatusOK, nil)
		return
	}
	SortType := "ai_model_manage.recommend desc,ai_model_manage.collected_count desc,ai_model_manage.download_count desc,ai_model_manage.reference_count desc"
	if orderBy != "" {
		if orderBy == "derivativeCount" {
			orderBy = "derivative_count"
		}
		if orderBy == "size_asc" {
			SortType = "ai_model_manage.size"
		} else {
			SortType = "ai_model_manage." + orderBy + " DESC"
		}
	}
	modelResult, count, err := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Type:                  typeInt,
		New:                   -1,
		Status:                modelStatus,
		IsQueryPrivate:        IsQueryPrivate,
		IsRecommend:           isRecommend,
		UserID:                user_id,
		IsCollected:           IsQueryCollect,
		CollectedUserId:       collected_user_id,
		LabelFilter:           labelFilter,
		FrameFilter:           frameFilterInt,
		ComputeResourceFilter: computeResourceFilter,
		Namelike:              Namelike,
		SortType:              SortType,
		RepoID:                repo_id,
		NotNeedEmpty:          notNeedEmpty,
		HasOnlineUrl:          hasOnlineUrl,
		MyExternal:            myExternal,
	})
	if err != nil {
		ctx.ServerError("Cloudbrain", err)
		return
	}
	userIds := make([]int64, len(modelResult))
	modelIds := make([]string, len(modelResult))
	repoIds := make([]int64, len(modelResult))
	for i, model := range modelResult {
		userIds[i] = model.UserId
		modelIds[i] = model.ID
		repoIds[i] = model.RepoId
	}
	repoInfo, err := queryRepoInfoByIds(repoIds)
	userNameMap := models.QueryUserName(userIds)
	var modelCollect map[string]*models.AiModelCollect
	if ctx.User != nil && queryType != 4 {
		modelCollect = models.QueryModelCollectedStatus(modelIds, ctx.User.ID)
	}

	for _, model := range modelResult {
		//removeIpInfo(model)
		model.TrainTaskInfo = ""
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
		if repoInfo != nil {
			repo := repoInfo[model.RepoId]
			if repo != nil {
				model.RepoName = repo.Name
				model.RepoOwnerName = repo.OwnerName
				model.RepoDisplayName = repo.DisplayName()
			}
			model.IsCanDownload = isModelCanDownload(ctx.User, repo, model)
		}
		if ctx.User != nil && modelCollect != nil {
			value := modelCollect[model.ID]
			if value != nil {
				model.IsCollected = true
			}
		} else {
			model.IsCollected = false
		}
		if needModelFile && len(model.Path) > 0 {
			//查询模型文件列表
			model.ModelFileList = modelmanage.QueryModelFileByModel(model)
		}
	}
	mapInterface := make(map[string]interface{})
	mapInterface["data"] = modelResult
	mapInterface["count"] = count
	ctx.JSON(http.StatusOK, mapInterface)
}

func queryRepoInfoByIds(intSlice []int64) (map[int64]*models.Repository, error) {
	keys := make(map[int64]string)
	uniqueElements := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = ""
			uniqueElements = append(uniqueElements, entry)
		}
	}
	re, err := models.GetRepositoriesMapByIDs(uniqueElements)
	return re, err
}

func ModelReadMeTmpl(ctx *context.Context) {
	ctx.HTML(200, tplModelSquareReadMe)
}

func ModelFileListTmpl(ctx *context.Context) {
	ctx.Data["max_model_size"] = setting.MaxModelSize * MODEL_MAX_SIZE
	ctx.HTML(200, tplModelFileList)
}

func ModelFileSettingTmpl(ctx *context.Context) {
	ctx.HTML(200, tplModelSetting)
}

func ModelEvolutionMapTmpl(ctx *context.Context) {
	ctx.HTML(200, tplModelEvolutionMap)
}

func ModelMigrateTmpl(ctx *context.Context) {
	ctx.Data["HfModelOwner"] = setting.ExternalTransfer.HfModelOwner
	ctx.HTML(200, tplModelMigrate)
}

func ModelMigratingTmpl(ctx *context.Context) {
	ctx.HTML(200, tplModelMigrating)
}

func setModelUser(model *models.AiModelManage) {
	user, err := models.GetUserByID(model.UserId)
	if err == nil {
		model.UserName = user.Name
		model.UserRelAvatarLink = user.RelAvatarLink()
	}
}

func setModelRepo(model *models.AiModelManage) {
	repo, err := models.GetRepositoryByID(model.RepoId)
	if err == nil {
		model.RepoName = repo.Name
		model.RepoOwnerName = repo.OwnerName
		model.RepoDisplayName = repo.DisplayName()
	}
}

func ModelEvolutionMapData(ctx *context.Context) {
	id := ctx.Query("id")
	model, err := models.QueryModelById(id)
	re := map[string]interface{}{
		"code": "-1",
	}
	var userId int64
	userId = -1
	if ctx.User != nil {
		userId = ctx.User.ID
	}
	if err == nil {
		modelMap := models.QueryModelEvalueAnalizeMap()
		removeIpInfo(model)
		repo, err := models.GetRepositoryByID(model.RepoId)
		if err == nil {
			model.RepoName = repo.Name
			model.RepoOwnerName = repo.OwnerName
			model.RepoDisplayName = repo.DisplayName()
			setModelUser(model)
			setModelDataSet(model)
			IsCanOper := true
			if repo.IsPrivate {
				if userId == -1 {
					IsCanOper = false
				} else {
					IsCanOper, _ = models.HasAccess(userId, repo)
				}
			}
			currentNode := &ModelMap{
				Type:            1,
				IsCurrent:       true,
				RepoName:        repo.Name,
				RepoOwnerName:   repo.OwnerName,
				RepoDisplayName: repo.DisplayName(),
				RepoId:          repo.ID,
				Model:           model,
				IsCanOper:       IsCanOper,
				IsPrivate:       repo.IsPrivate,
			}
			ParentNode := findParent(model, userId)
			if ParentNode != nil {
				nexts := make([]*ModelMap, 0)
				nexts = append(nexts, currentNode)
				ParentNode.Next = nexts
			} else {
				ParentNode = currentNode
			}
			findChild2(currentNode, modelMap, ctx.User)
			re["code"] = "0"
			re["node"] = ParentNode
			ctx.JSON(200, ParentNode)
		}
	} else {
		re["msg"] = "No such model."
		ctx.JSON(200, re)
	}
}

func setModelDataSet(model *models.AiModelManage) {
	if model.TrainTaskInfo != "" {
		var task models.Cloudbrain
		err1 := json.Unmarshal([]byte(model.TrainTaskInfo), &task)
		if err1 == nil {
			model.DatasetInfo = GetCloudBrainDataSetInfo(task.Uuid, task.DatasetName, false)
			tmpTask, err1 := models.GetCloudbrainByIDWithDeleted(fmt.Sprint(task.ID))
			if err1 == nil {
				cloudbrainTaskJson, _ := json.Marshal(tmpTask)
				model.TrainTaskInfo = string(cloudbrainTaskJson)
			}
		}
	}
}

func findParent(model *models.AiModelManage, userId int64) *ModelMap {
	if model.TrainTaskInfo != "" {
		var task models.Cloudbrain
		err := json.Unmarshal([]byte(model.TrainTaskInfo), &task)
		if err != nil {
			log.Info("error=" + err.Error())
		} else {
			log.Info("find parent model name." + task.ModelName)
			if task.ModelName != "" {
				result := &ModelMap{
					Type:      1,
					IsParent:  true,
					IsCanOper: true,
					IsPrivate: model.IsPrivate,
				}
				modelsArray := make([]*models.AiModelManage, 0)
				if task.ModelId != "" {
					modelIdArray := task.GetModelIdArray()
					for i := 0; i < len(modelIdArray); i++ {
						modelId := modelIdArray[i]
						parentModel, err := models.QueryModelById(modelId)
						setModelRepo(parentModel)
						setModelUser(parentModel)
						setModelDataSet(parentModel)
						if err != nil {
							return nil
						}
						modelsArray = append(modelsArray, parentModel)
					}
					result.Models4Parent = modelsArray
					return result
				} else {
					modelNameArray := task.GetModelNameArray()
					for i := 0; i < len(modelNameArray); i++ {
						modelName := modelNameArray[i]
						modelList := models.QueryModelByName(modelName, task.RepoID)
						if modelList != nil && len(modelList) > 0 {
							for _, parentModel := range modelList {
								setModelUser(parentModel)
								setModelRepo(parentModel)
								setModelDataSet(parentModel)
								if parentModel.Version == task.ModelVersion {
									modelsArray = append(modelsArray, parentModel)
								}
							}
						}
					}
					result.Models4Parent = modelsArray
					return result
				}
			}
		}
	}
	return nil
}

func findChild2(currentModelNode *ModelMap, modelMap map[string][]*models.AiModelManage, user *models.User) {
	log.Info("find child2 start.")
	if currentModelNode.Model != nil {
		currentModel := currentModelNode.Model
		childModelsList := modelMap[currentModel.ID]
		if childModelsList != nil {
			for _, childModel := range childModelsList {
				var repoNode *ModelMap
				find := false
				if currentModelNode.Next != nil {
					for _, repo := range currentModelNode.Next {
						if childModel.RepoId == repo.RepoId {
							repoNode = repo
							find = true
							break
						}
					}
				}
				var repo *models.Repository
				var err error
				if !find {
					repo, err = models.GetRepositoryByID(childModel.RepoId)
					if err == nil {
						IsCanOper := true
						if repo.IsPrivate {
							if user == nil {
								IsCanOper = false
							} else {
								IsCanOper, _ = models.HasAccess(user.ID, repo)
							}
						}
						repoNode = &ModelMap{
							Type:            0,
							RepoName:        repo.Name,
							RepoOwnerName:   repo.OwnerName,
							RepoDisplayName: repo.DisplayName(),
							RepoId:          repo.ID,
							IsCanOper:       IsCanOper,
							IsPrivate:       repo.IsPrivate,
						}
						if currentModelNode.Next != nil {
							currentModelNode.Next = append(currentModelNode.Next, repoNode)
						} else {
							repoNodes := make([]*ModelMap, 0)
							repoNodes = append(repoNodes, repoNode)
							currentModelNode.Next = repoNodes
						}
					}
				}
				if repoNode != nil && repoNode.RepoName != "" {
					childModelNodes := repoNode.Next
					if childModelNodes == nil {
						childModelNodes = make([]*ModelMap, 0)
					}
					childModel.RepoName = repoNode.RepoName
					childModel.RepoOwnerName = repoNode.RepoOwnerName
					childModel.RepoDisplayName = repoNode.RepoDisplayName
					setModelUser(childModel)
					setModelDataSet(childModel)
					IsCanOper := true
					if childModel.IsPrivate {
						if user == nil {
							IsCanOper = false
						} else {
							if repo != nil {
								IsCanOper = isModelCanDownload(user, repo, childModel)
							} else {
								repo, err = models.GetRepositoryByID(repoNode.RepoId)
								if err == nil {
									IsCanOper = isModelCanDownload(user, repo, childModel)
								} else {
									IsCanOper = repoNode.IsCanOper
								}
							}
						}
					}
					modelNode := &ModelMap{
						Type:      1,
						Model:     childModel,
						IsCanOper: IsCanOper,
						IsPrivate: childModel.IsPrivate,
					}
					childModelNodes = append(childModelNodes, modelNode)
					repoNode.Next = childModelNodes
					findChild2(modelNode, modelMap, user)
				}
			}
		}

	} else {
		log.Info("the current model is nil.")
	}

}

func findChild(currentNode *ModelMap, modelMap map[string][]*models.AiModelManage) {
	log.Info("find child start.")
	if currentNode.Model != nil {
		currentModel := currentNode.Model
		re, err := models.GetCloudBrainByModelId(currentModel.ID)
		if err != nil || len(re) == 0 {
			re, err = models.GetCloudBrainByRepoIdAndModelName(currentModel.RepoId, currentModel.Name)
		}
		log.Info("start to load child model.")
		if err == nil && len(re) > 0 {
			repoNodes := getRepoNodes(re)
			currentNode.Next = repoNodes
			for _, node := range repoNodes {
				childModels := models.QueryModelByRepoId(node.RepoId)
				childNodes := make([]*ModelMap, 0)
				if childModels != nil && len(childModels) > 0 {
					for _, childModel := range childModels {
						if childModel.TrainTaskInfo != "" {
							childModel.RepoName = node.RepoName
							childModel.RepoOwnerName = node.RepoOwnerName
							childModel.RepoDisplayName = node.RepoDisplayName
							log.Info("childModel.RepoName=" + childModel.RepoName)
							log.Info("childModel.RepoOwnerName=" + childModel.RepoOwnerName)
							var task models.Cloudbrain
							err := json.Unmarshal([]byte(childModel.TrainTaskInfo), &task)
							if err != nil {
								log.Info("error=" + err.Error())
							} else {
								log.Info("task.ModelId=%v,currentModel.ID=%v", task.ModelId, currentModel.ID)
								if task.ModelId != "" && task.HasUseModel(currentModel.ID) {
									setModelUser(childModel)
									setModelDataSet(childModel)
									modelMap := &ModelMap{
										Type:  1,
										Model: childModel,
									}
									childNodes = append(childNodes, modelMap)
								} else {
									log.Info("task.ModelName=%v,currentModel.Name=%v", task.ModelName, currentModel.Name)
									log.Info("task.ModelVersion=%v,currentModel.Version=%v", task.ModelVersion, currentModel.Version)
									if task.ModelName == currentModel.Name && task.ModelVersion == currentModel.Version {
										setModelUser(childModel)
										setModelDataSet(childModel)
										modelMap := &ModelMap{
											Type:  1,
											Model: childModel,
										}
										childNodes = append(childNodes, modelMap)
									}
								}
							}
						}
					}

				}
				node.Next = childNodes
				for _, child := range childNodes {
					findChild(child, modelMap)
				}
			}
		}
	} else {
		log.Info("the current model is nil.")
	}

}

func getRepoNodes(re []*models.Cloudbrain) []*ModelMap {
	result := make([]*ModelMap, 0)
	repoMap := make(map[int64]string, 0)
	for _, task := range re {
		repo, err := models.GetRepositoryByID(task.RepoID)
		if err == nil {
			if _, ok := repoMap[repo.ID]; !ok {
				modelMap := &ModelMap{
					Type:            0,
					RepoName:        repo.Name,
					RepoOwnerName:   repo.OwnerName,
					RepoDisplayName: repo.DisplayName(),
					RepoId:          repo.ID,
				}
				result = append(result, modelMap)
			}
			repoMap[repo.ID] = "true"
		}
	}
	return result
}

func ModifyModelReadMe(ctx *context.Context) {
	id := ctx.Query("id")
	model, err := models.QueryModelById(id)
	re := map[string]string{
		"code": "-1",
	}
	if err == nil {
		content := ctx.Query("content")
		hasOnlineUrl := hasOnlineUrlInReadme(content)
		log.Info(" model.License =")
		if model.License != "" {
			content = addLicenseMDToReadme(content, model.License)
		}
		path := Model_prefix + models.AttachmentRelativePath(id) + "/"
		helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
		err = helper.PutString(path+README_FILE_NAME, content)
		if err != nil {
			re["msg"] = "Failed to created readme file."
			log.Info("Failed to created readme file. as:" + err.Error())
		} else {
			re["code"] = "0"
		}

		if hasOnlineUrl {
			models.ModifyModelHashOnlineUrl(id, 1)
		} else {
			models.ModifyModelHashOnlineUrl(id, 0)
		}
		ctx.JSON(200, re)
	} else {
		re["msg"] = "No such model."
		ctx.JSON(200, re)
	}
}

func updateReadMeMd(model *models.AiModelManage, content string) {
	path := Model_prefix + models.AttachmentRelativePath(model.ID) + "/"

	helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
	err := helper.PutString(path+README_FILE_NAME, content)

	if err != nil {
		log.Info("Failed to update readme file. as:" + err.Error())
	}

}

func removeLicenseFromMD(content string, licenseId string) string {
	lines := strings.Split(content, "\n")
	licenseMap := getLicenseMap(licenseId)
	if licenseMap == nil {
		return content
	}
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if line == "## 许可证" || strings.HasPrefix(line, "["+licenseMap["name"]+"]") {
			log.Info("line is equal. + line=" + line)
			continue
		}
		newLines = append(newLines, line)
	}
	return strings.Join(newLines, "\n")
}

func getLicenseMap(licenseId string) map[string]string {
	url := setting.RecommentRepoAddr + "model/license.json"
	result, err := repository.RecommendContentFromPromote(url)
	log.Info("license result=" + result)
	remap := make([]map[string]string, 0)
	if err == nil {
		err = json.Unmarshal([]byte(result), &remap)
		if err != nil {
			log.Info("error=" + err.Error())
			return nil
		}
		for _, licenseMap := range remap {
			if licenseMap["id"] == licenseId {
				return licenseMap
			}
		}
	} else {
		log.Info("error=" + err.Error())
	}
	return nil
}

func addLicenseMDToReadme(content string, licenseId string) string {
	if hasLicense(content) {
		return content
	}
	log.Info("start to add license to md file.")
	licenseMap := getLicenseMap(licenseId)
	if licenseMap == nil {
		return content
	}
	content += "\n## 许可证\n"
	content += "[" + licenseMap["name"] + "](" + licenseMap["linkUrl"] + ")\n"
	return content
}

func hasLicense(content string) bool {
	re := regexp.MustCompile(`## 许可证\s+(.+)`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return true
	}
	return false
}

func hasOnlineUrlInReadme(content string) bool {
	re := regexp.MustCompile(`在线体验地址\s+(.+)`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		url := match[1]
		log.Info("find online url=" + url)
		re := regexp.MustCompile(`^\s*|\r?\n`)
		filteredStr := re.ReplaceAllString(url, "")
		if strings.HasPrefix(filteredStr, "http") {
			return true
		}
	}
	return false
}

func getMDContent(model *models.AiModelManage) string {
	files := queryOneLevelModelFile(model, "")
	var content []byte
	for _, file := range files {
		if strings.ToLower(file.FileName) == strings.ToLower(README_FILE_NAME) {
			path := Model_prefix + models.AttachmentRelativePath(model.ID) + "/"
			storageHelper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
			body, err := storageHelper.OpenFile(path + file.FileName)
			if err != nil {
				log.Info("download file failed: %s\n", err.Error())
				break
			} else {
				defer body.Close()
				content, err = ioutil.ReadAll(body)
				return string(content)
			}
		}
	}
	return ""
}

func QueryModelReadMe(ctx *context.Context) {
	id := ctx.Query("id")
	model, err := models.QueryModelById(id)

	re := map[string]string{
		"code": "-1",
	}
	if err == nil {

		if model.IsPrivate && !isCanReadPrivateModel(ctx, model.UserId) {
			re["msg"] = "No right to read."
			ctx.JSON(401, re)
			return
		}

		mdContent := getMDContent(model)
		metas := map[string]string{"include_toc": "true"}
		if mdContent != "" {
			re["isExistMDFile"] = "true"
			re["fileName"] = README_FILE_NAME
			re["content"] = mdContent
			re["htmlcontent"] = string(markdown.RenderRaw([]byte(mdContent), "", false, metas))
		} else {
			re["isExistMDFile"] = "false"
			re["fileName"] = README_FILE_NAME
			url := setting.RecommentRepoAddr + "model/" + README_FILE_NAME
			result, err := repository.RecommendContentFromPromote(url)
			if err == nil {
				re["content"] = result
				re["htmlcontent"] = string(markdown.RenderRaw([]byte(result), "", false, metas))
			}
		}
		re["code"] = "0"
		ctx.JSON(200, re)
	} else {
		re["msg"] = "No such model."
		ctx.JSON(200, re)
	}

}

func QueryModelLabel(ctx *context.Context) {
	url := setting.RecommentRepoAddr + "model/label.json"
	result, err := repository.RecommendContentFromPromote(url)
	log.Info("label result=" + result)
	remap := make([]map[string]string, 0)
	if err == nil {
		err = json.Unmarshal([]byte(result), &remap)
		if err != nil {
			log.Info("error=" + err.Error())
		}
	} else {
		log.Info("error=" + err.Error())
	}
	if err == nil {
		ctx.JSON(200, remap)
	} else {
		ctx.JSON(200, "")
	}
}

func setModelRelateInfo(model *models.AiModelManage) {
	online := make([]map[string]interface{}, 0)
	othertaskList := make([]map[string]interface{}, 0)
	re, err := models.GetCloudBrainByModelId(model.ID)
	if err == nil && len(re) > 0 {
		for _, task := range re {
			if task.JobType == string(models.JobTypeOnlineInference) {
				if task.Status == string(models.JobRunning) {
					repo, err1 := models.GetRepositoryByID(task.RepoID)
					if err1 == nil && !repo.IsPrivate {
						onlineTask := make(map[string]interface{})
						onlineTask["repoName"] = repo.Name
						onlineTask["ownerName"] = repo.OwnerName
						onlineTask["repoDisplayName"] = repo.DisplayName()
						onlineTask["repoId"] = repo.ID
						onlineTask["onlineInferUrl"] = getOnlineInferUrl(task.JobID)
						online = append(online, onlineTask)
						continue
					}
				}
			}

			if len(othertaskList) >= 5 {
				continue
			}
			otherTask := make(map[string]interface{})
			repo, err1 := models.GetRepositoryByID(task.RepoID)
			if err1 == nil && !repo.IsPrivate {
				otherTask["repoName"] = repo.Name
				otherTask["ownerName"] = repo.OwnerName
				otherTask["repoDisplayName"] = repo.DisplayName()
				otherTask["jobId"] = task.JobID
				otherTask["id"] = task.ID
				otherTask["jobType"] = task.JobType
				otherTask["type"] = task.Type
				otherTask["computeResource"] = task.ComputeResource
				otherTask["jobName"] = task.JobName
				otherTask["displayJobName"] = task.DisplayJobName
				othertaskList = append(othertaskList, otherTask)
			}

		}
	}
	setModelDataSet(model)
	model.OnlineInfo = online
	model.UsedCloudbrain = othertaskList
}

func getOnlineInferUrl(jobId string) string {
	result, err := grampus.GetNotebookJob(jobId)
	if err != nil {
		return ""
	}
	return cloudbrainTask.GrampusNotebookUrl(result)
}
