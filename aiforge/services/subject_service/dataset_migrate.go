package subject_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/generate"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

func CreateDataset4OldDataset(req *models.DatasetRegistry) (*models.DatasetRegistry, *response.BizError) {
	if !NamePattern.MatchString(req.Name) {
		return nil, response.DATASET_NAME_INVALID
	}
	if req.Alias == "" {
		req.Alias = req.Name
	}
	if !AlphaDashDotChinese.MatchString(req.Alias) {
		return nil, response.DATASET_ALIAS_INVALID
	}
	uploaderHelper := GetUploadHelper(models.DatasetSubject)
	storageType := uploaderHelper.GetStorageType()
	if storageType == "" {
		log.Info("dataset storage type setting is empty")
		return nil, response.SYSTEM_ERROR
	}
	stroageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if stroageHelper == nil {
		log.Info("dataset storage type %s not support", storageType)
		return nil, response.SYSTEM_ERROR
	}
	redisKey := redis_key.DatasetNameLock(req.OwnerID, strings.ToLower(req.Name))
	lock := redis_lock.NewDistributeLock(redisKey)
	isOk, err := lock.LockWithWait(5*time.Second, 5*time.Second)
	if err != nil {
		log.Error("CreateDataset LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
		return nil, response.NewBizError(err)
	}
	if !isOk {
		log.Error("CreateDataset LockWithWait failed, redisKey=%s, err=%v", redisKey, response.DATASET_NAME_EXIST.ToError())
		return nil, response.DATASET_NAME_EXIST
	}
	defer func() {
		err := lock.UnLock()
		if err != nil {
			log.Error("CreateDataset UnLock failed, redisKey=%s, err=%v", redisKey, err)
		}
	}()

	record, err := models.GetDatasetRegistryByOwnerAndName(req.OwnerID, req.Name)
	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("CreateDataset GetDatasetRegistryByOwnerAndName failed, ownerId=%d, name=%s, err=%v", req.OwnerID, req.Name, err)
			return nil, response.NewBizError(err)
		}
	}
	if record != nil {
		return nil, response.DATASET_NAME_EXIST
	}
	record, err = models.GetDatasetRegistryByOwnerAndAlias(req.OwnerID, req.Alias)
	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("CreateDataset GetDatasetRegistryByOwnerAndAlias failed, ownerId=%d, alias=%s, err=%v", req.OwnerID, req.Alias, err)
			return nil, response.NewBizError(err)
		}
	}
	if record != nil {
		return nil, response.DATASET_ALIAS_EXIST
	}
	record, err = models.GetDatasetRegistryByID(req.ID)
	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("CreateDataset GetDatasetRegistryByID failed, ownerId=%d, name=%s, err=%v", req.OwnerID, req.Name, err)
			return nil, response.NewBizError(err)
		}
	}
	if record != nil {
		return nil, response.DATASET_EXIST
	}

	if req.Path == "" {
		return nil, response.DATASET_PATH_EMPTY
	}

	pathShouldBe := strings.TrimPrefix(path.Join(DATASET_PREFIX, path.Join(req.ID[0:1], req.ID[1:2], req.ID+req.ID)), "/") + "/"

	if pathShouldBe != req.Path {
		if !strings.HasPrefix(req.Path, strings.TrimSuffix(pathShouldBe, "/")) {
			return nil, response.DATASET_PATH_INCORRECT
		}
	}
	doer, _ := models.GetUserByID(req.CreatorID)
	if doer == nil {
		doer = &models.User{
			ID: req.CreatorID,
		}
	}
	err = models.CreateDatasetRegistry4Old(req, doer)
	if err != nil {
		log.Error("CreateDatasetRegistry failed, name=%s, err=%v", req.Name, err)
		return nil, response.NewBizError(err)
	}

	return req, nil
}

func HandleAllOldDatasets() error {
	page := 1
	pageSize := 200
	count := 0
	for {
		log.Info("handleOldDataset new page start,page=%d pageSize=%d", page, pageSize)
		oldDatasets, err := models.GetUnhandledOldDataset(page, pageSize)
		if err != nil {
			log.Error("handleOldDataset GetUnhandledOldDataset err page=%d pageSize=%d.err=%v", page, pageSize, err)
			return err
		}
		if len(oldDatasets) == 0 {
			log.Info("handleOldDataset No more old datasets to handle.")
			return nil
		}
		count += len(oldDatasets)
		log.Info("handleOldDataset Handling old datasets, page=%d, datasaet lenth=%d", page, len(oldDatasets))
		for _, oldDataset := range oldDatasets {
			err := handleOldDataset(oldDataset)
			if err != nil {
				log.Error("handleOldDataset handleOldDataset error: %v, oldDataset: %+v", err, oldDataset)
				continue
			}
		}
		page++
		if len(oldDatasets) < pageSize {
			log.Info("handleOldDataset Handled all old datasets in this batch, exiting.")
			return nil
		}
		if page*pageSize > 100000 {
			log.Error("handleOldDataset Handled too many old datasets, stopping to avoid overload. Current page: %d", page)
			return nil
		}
		log.Info("handleOldDataset page end,count=%d page=%d pageSize=%d", count, page, pageSize)

	}
}

func handleOldDataset(oldDataset models.OldDataset) error {
	uuid := oldDataset.Attachment.UUID
	redisKey := redis_key.HandleOldDatasetLock(uuid)
	lock := redis_lock.NewDistributeLock(redisKey)
	isOk, err := lock.LockWithWait(20*time.Second, 20*time.Second)
	if err != nil {
		log.Error("handleOldDataset HandleOldDatasetLock LockWithWait failed, redisKey=%s, err=%v", redisKey, err)
		return err
	}
	if !isOk {
		log.Error("handleOldDataset HandleOldDatasetLock LockWithWait failed, redisKey=%s, err=%v", redisKey, response.DATASET_NAME_EXIST.ToError())
		return fmt.Errorf("get HandleOldDatasetLock failed")
	}
	defer func() {
		err := lock.UnLock()
		if err != nil {
			log.Error("handleOldDataset HandleOldDatasetLock UnLock failed, redisKey=%s, err=%v", redisKey, err)
		}
	}()
	k := redis_key.ProcessingOldDataset()
	redis_client.SAdd(k, uuid)
	redis_client.Expire(k, 24*time.Hour)
	defer redis_client.SRem(k, uuid)

	//1、插入新数据集表，要处理同名逻辑和解压路径一致的问题
	record, err := models.CreateOrFindOldDatasetProcessRecord(&models.OldDatasetProcessRecord{
		ID:     uuid,
		Status: 0,
	})
	if err != nil {
		return err
	}
	log.Info("handleOldDataset process start,id=%s status=%d", uuid, record.Status)
	if record.Status == 0 {
		log.Info("handleOldDataset process status 0 start,id=%s ", uuid)
		err = CreateNewDataset(oldDataset)
		if err != nil {
			log.Error("handleOldDataset process status 0 failed, err=%v", err)
			models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
			return err
		}
		models.UpdateUserDatasetNum(oldDataset.Dataset.UserID)
		models.UpdateOldDatasetProcessRecordStatus(uuid, 1)
		record.Status = 1
		log.Info("handleOldDataset process status 0 success,id=%s ", uuid)

	}
	dataset, err := models.GetDatasetRegistryByID(oldDataset.Attachment.UUID)
	if err != nil {
		log.Error("handleOldDataset GetDatasetRegistryByID failed, attachmentUUID=%s, err=%v", oldDataset.Attachment.UUID, err)
		models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
		return err
	}
	if record.Status == 1 {
		//2、创建readme和版本文件
		log.Info("handleOldDataset process status 1 start,id=%s ", uuid)
		err = CreateNewReadme(oldDataset, dataset)
		if err != nil {
			log.Error("handleOldDataset process status 1  CreateNewReadme failed, err=%v", err)
			models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
			return err
		}
		err = CreateNewVersionFile(oldDataset, dataset)
		if err != nil {
			log.Error("handleOldDataset process status 1  CreateNewVersionFile failed, err=%v", err)
			models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
			return err
		}
		models.UpdateOldDatasetProcessRecordStatus(uuid, 2)
		record.Status = 2
		log.Info("handleOldDataset process status 1 success,id=%s ", uuid)
	}
	if record.Status == 2 {
		//3、添加协作者权限
		log.Info("handleOldDataset process status 2 start,id=%s ", uuid)
		err = HandleColaborator(oldDataset, dataset)
		if err != nil {
			log.Error("handleOldDataset process status 2  HandleColaborator failed,id=%s err=%v", uuid, err)
			models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
			return err
		}
		models.UpdateOldDatasetProcessRecordStatus(uuid, 3)
		record.Status = 3
		log.Info("handleOldDataset process status 2 success,id=%s ", uuid)

	}
	//4、收藏关系
	if record.Status == 3 {
		log.Info("handleOldDataset process status 3 start,id=%s ", uuid)
		err = HandleCollections(oldDataset, dataset)
		if err != nil {
			log.Error("handleOldDataset process status 3 HandleCollections failed, err=%v", err)
			models.UpdateOldDatasetProcessRecordRemark(uuid, err.Error())
			return err
		}
		models.UpdateOldDatasetProcessRecordStatus(uuid, 4)
		record.Status = 4
		log.Info("handleOldDataset process status 3 success,id=%s ", uuid)

	}
	log.Info("handleOldDataset process finish,id=%s ", uuid)
	return nil
}

func HandleCollections(oldDataset models.OldDataset, dataset *models.DatasetRegistry) error {
	return models.HandleOldDatasetCollections(oldDataset.Dataset.ID, dataset.ID, dataset.OwnerID)
}

func HandleColaborator(oldDataset models.OldDataset, dataset *models.DatasetRegistry) error {
	dataset.GetOwner()
	datasetOwner := dataset.Owner
	isOrgDataset := datasetOwner.IsOrganization()

	repo := oldDataset.Repo

	//删掉可能已存在的数据集协作者列表
	subjectCtx := dataset.ConvertSubjectAccessContext()
	datasetCollaborators, err := subjectCtx.GetCollaborators(models.ListOptions{})
	if err != nil {
		log.Error("subjectCtx GetCollaborators err.id=%s err=%v", oldDataset.Attachment.UUID, err)
		return err
	}
	for _, collaborator := range datasetCollaborators {
		subjectCtx.DeleteCollaboration(collaborator.ID)
	}

	collaborators, err := repo.GetCollaborators(models.ListOptions{})
	if err != nil {
		log.Error("repo GetCollaborators err.id=%s err=%v", oldDataset.Attachment.UUID, err)
		return err
	}
	//依次添加项目协作者到数据集
	for _, collaborator := range collaborators {
		err := AddCollaborator4Old(collaborator.User, dataset, collaborator.Collaboration.Mode)
		if err != nil {
			log.Error("handleOldDataset  AddCollaborator4Old error.user=%+v dataset=%+v", collaborator.User, dataset)
			continue
		}
	}
	if isOrgDataset {
		//修改组织、团队设置
		err := models.HandleOrg4OldDataset(dataset.OwnerID)
		if err != nil {
			log.Error("handleOldDataset HandleOrg4OldDataset error.OwnerID=%d", dataset.OwnerID)
			return err
		}
		teams, err := repo.GetRepoTeams()
		if err != nil {
			log.Error("handleOldDataset GetRepoTeams error.OwnerID=%d", dataset.OwnerID)
			return err
		}
		for _, team := range teams {
			//初始化时使用项目的相关配置
			datasetAuthChanged := team.DatasetAuthorize != team.Authorize
			includeAllDatasetChanged := team.IncludesAllDatasets != team.IncludesAllRepositories
			team.DatasetAuthorize = team.Authorize
			team.IncludesAllDatasets = team.IncludesAllRepositories
			team.CanCreateOrgDataset = team.CanCreateOrgRepo
			err := models.HandleTeam4OldDataset(team, datasetAuthChanged, includeAllDatasetChanged)
			if err != nil {
				log.Error("handleOldDataset HandleTeam4OldDataset error.OwnerID=%d", dataset.OwnerID)
				continue
			}

		}
		//删除数据集已有团队
		datasetTeams, err := subjectCtx.GetTeams()
		if err != nil {
			log.Error("handleOldDataset get datasetTeams error.datasetID=%d", dataset.ID)
			return err
		}
		for _, team := range datasetTeams {
			err := team.RemoveDataset4Old(dataset.ID)
			if err != nil {
				log.Error("handleOldDataset RemoveDataset error.team=%+v dataset=%+v", team, dataset)
				return err
			}
		}
		//依次添加项目的团队到数据集
		for _, team := range teams {
			err := team.AddDataset(dataset)
			if err != nil {
				log.Error("handleOldDataset AddDataset error.team=%+v dataset=%+v", team, dataset)
				continue
			}
			//更新团队数据集数量
			models.UpdateTeamDatasetnum(team.ID)
		}
	}
	return nil
}

func AddCollaborator4Old(u *models.User, dataset *models.DatasetRegistry, mode models.AccessMode) error {
	// Organization is not allowed to be added as a collaborator.
	if u.IsOrganization() {
		return response.ORG_NOT_ALLOWED_TO_BE_COLLABORATOR.ToError()

	}
	subjectCtx := dataset.ConvertSubjectAccessContext()
	if subjectCtx.OwnerID == u.ID {
		return response.OWNER_NOT_ALLOWED_TO_BE_COLLABORATOR.ToError()
	}

	if got, err := subjectCtx.IsCollaborator(u.ID); err == nil && got {
		return nil

	}

	if err := subjectCtx.AddCollaborator(u); err != nil {
		return err
	}
	if err := subjectCtx.ChangeCollaborationAccessMode(u.ID, mode); err != nil {
		return err
	}
	return nil
}

func makeValidDatasetPath(input string) string {
	validName := input
	// 使用正则表达式将非字母数字下划线的字符替换为下划线
	re := regexp.MustCompile(`[^a-zA-Z0-9\_.-]`)
	validName = re.ReplaceAllString(validName, "_")
	re = regexp.MustCompile(`_+`)
	validName = re.ReplaceAllString(validName, "_")
	re = regexp.MustCompile(`-+`)
	validName = re.ReplaceAllString(validName, "-")
	re = regexp.MustCompile(`\.+`)
	validName = re.ReplaceAllString(validName, ".")
	validName = strings.Trim(validName, "_")

	if validName == "" || validName == "_" || validName == "." || validName == "-" {
		randStr, _ := generate.GetRandomString(3)
		if randStr == "" {
			randStr = "dfg"
		}
		randStr = strings.ToLower(randStr)
		validName = "dataset_" + randStr
	}
	if len(validName) > 100 {
		validName = validName[:100]
	}
	runes := []rune(validName)
	if !isAlNum(runes[0]) {
		runes[0] = 'a'
	}
	if !isAlNum(runes[len(runes)-1]) {
		runes[len(runes)-1] = '0'
	}
	return string(runes)
}

func isAlNum(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
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
		s = "数据集" + randStr
	}
	runes := []rune(s)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	s = string(runes)
	return s
}

func CreateNewDataset(oldDataset models.OldDataset) error {
	attachment := oldDataset.Attachment
	repo := oldDataset.Repo
	dataset := oldDataset.Dataset
	repo.GetOwner()
	owner := repo.Owner
	if owner == nil {
		log.Error("handleOldDataset GetOwner failed, repoId=%d", repo.ID)
		return fmt.Errorf("handleOldDataset owner not found for repoId=%d", repo.ID)
	}
	// oldDatasetName := dataset.Title
	fileName := strings.TrimSuffix(strings.TrimSuffix(attachment.Name, ".zip"), ".tar.gz")
	// fileName = FilterInvalidChars(fileName)
	alias := makeValidAlias(fileName)
	newDatasetName := makeValidDatasetPath(fileName)
	// if oldDatasetName != "" && strings.ToLower(oldDatasetName) != strings.ToLower(fileName) {
	// 	newDatasetName = oldDatasetName + "_" + fileName
	// }

	tags := make([]string, 0)
	if dataset.Category != "" {
		tags = append(tags, dataset.Category)
	}
	tasks := make([]string, 0)
	if dataset.Task != "" {
		tasks = append(tasks, dataset.Task)
	}
	req := &models.DatasetRegistry{
		ID:             attachment.UUID,
		Name:           newDatasetName,
		LowerName:      strings.ToLower(newDatasetName),
		Alias:          alias,
		LowerAlias:     strings.ToLower(alias),
		Tags:           tags,
		Tasks:          tasks,
		License:        dataset.License,
		StorageType:    string(entity.GetStorageTypeFromCloudbrainType(attachment.Type)),
		IsPrivate:      attachment.IsPrivate,
		UseCount:       attachment.UseNumber,
		DownloadCount:  attachment.DownloadCount,
		NumCollections: 0, //实际处理收藏关系时使用
		Recommend:      dataset.Recommend,
		CreatorID:      attachment.UploaderID,
		OwnerID:        owner.ID,
		CreatedUnix:    attachment.CreatedUnix,
		UpdatedUnix:    attachment.CreatedUnix,
		Size:           attachment.Size,
		Path:           attachment.UnzipPath,
	}
	for i := 0; i < 5; i++ {
		//尝试创建数据集
		//如果数据集名称已存在，则重新生成一个新的名称
		//最多尝试5次
		log.Info("handleOldDataset CreateDataset try %d, name=%s, ownerId=%d uuid=%s", i+1, newDatasetName, owner.ID, attachment.UUID)
		res, bizErr := CreateDataset4OldDataset(req)
		if bizErr == nil {
			log.Info("handleOldDataset CreateDataset success, datasetId=%s", res)
			break
		}
		if i == 4 {
			log.Error("handleOldDataset CreateDataset failed after 5 attempts, dataset name=%s, ownerId=%d uuid=%s", newDatasetName, owner.ID, attachment.UUID)
			return fmt.Errorf("CreateDataset failed after 5 attempts, dataset name=%s, ownerId=%d uuid=%s", newDatasetName, owner.ID, attachment.UUID)
		}
		if bizErr.Code == response.DATASET_NAME_EXIST.Code {
			log.Error("handleOldDataset CreateDataset failed, dataset name already exists, name=%s, ownerId=%d, uuid=%s, err=%v", newDatasetName, owner.ID, bizErr.ToError(), attachment.UUID)
			randStr, err := generate.GetRandomString(3)
			if err != nil {
				log.Error("handleOldDataset GetRandomString failed, uuid = %s,err=%v", attachment.UUID, err)
				return fmt.Errorf("GetRandomString failed, err=%v", err)
			}
			randStr = strings.ToLower(randStr)

			if len(newDatasetName) <= 96 {
				newDatasetName = strings.TrimSuffix(newDatasetName, "_") + "_" + randStr
			} else {
				runes := []rune(newDatasetName)

				var out []rune
				out = make([]rune, 100)
				copy(out, runes[:96])
				copy(out[96:], []rune("_"+randStr))
				newDatasetName = string(out)
			}
			req.Name = newDatasetName
			req.LowerName = strings.ToLower(req.Name)
			newAlias := makeValidAlias(newDatasetName)
			req.Alias = newAlias
			req.LowerAlias = strings.ToLower(newAlias)
			log.Info("handleOldDataset Retrying CreateDataset with new name=%s, ownerId=%d uuid = %s", newDatasetName, owner.ID, attachment.UUID)
			continue

		}
		if bizErr.Code == response.DATASET_ALIAS_EXIST.Code {
			log.Error("handleOldDataset CreateDataset failed, dataset alias already exists, alias=%s, ownerId=%d, uuid=%s, err=%v", req.Alias, owner.ID, bizErr.ToError(), attachment.UUID)
			randStr, err := generate.GetRandomString(3)
			if err != nil {
				log.Error("handleOldDataset GetRandomString failed, uuid = %s,err=%v", attachment.UUID, err)
				return fmt.Errorf("GetRandomString failed, err=%v", err)
			}
			randStr = strings.ToLower(randStr)
			var newAlias = req.Alias
			if len(req.Alias) <= 96 {
				newAlias = strings.TrimSuffix(newAlias, "_") + "_" + randStr
			} else {
				runes := []rune(newAlias)

				var out []rune
				out = make([]rune, 100)
				copy(out, runes[:96])
				copy(out[96:], []rune("_"+randStr))
				newAlias = string(out)
			}
			req.Alias = newAlias
			req.LowerAlias = strings.ToLower(newAlias)
			log.Info("handleOldDataset Retrying CreateDataset with new alias=%s, ownerId=%d uuid = %s", newAlias, owner.ID, attachment.UUID)
			continue
		}
		if bizErr.Code == response.DATASET_PATH_INCORRECT.Code {
			//此时原数据集文件解压路径不对，需要copy到新地址
			dataId := req.ID
			sourcePath := attachment.UnzipPath
			newPath := strings.TrimPrefix(path.Join(DATASET_PREFIX, path.Join(dataId[0:1], dataId[1:2], dataId+dataId)), "/") + "/"
			oldStorageType := storage_helper.GetStorageTypeFromIntType(attachment.Type)
			// if oldStorageType != entity.OBS {
			// 	log.Error("handleOldDataset source storage is minio,dataId =%s", dataId)
			// 	return fmt.Errorf("source storage is minio,dataId = %s", dataId)
			// }
			err := storage_helper.Copy(oldStorageType, oldStorageType, sourcePath, newPath)
			if err != nil {
				//识别出来是文件已存在？
				return err
			}
			req.Path = newPath
			continue
		}
		if bizErr.Code == response.DATASET_PATH_EMPTY.Code {
			log.Error("handleOldDataset dataset path is empty,dataId =%s", req.ID)
			return fmt.Errorf("sdataset path is empty,dataId =%s", req.ID)
		}
		if bizErr.Code == response.DATASET_EXIST.Code {
			log.Info("handleOldDataset dataset record exists.dataId =%s", req.ID)
			return nil
		}
		log.Error("handleOldDataset CreateDataset error, dataset name=%s, ownerId=%d uuid=%s err=%v", newDatasetName, owner.ID, attachment.UUID, bizErr)

		return bizErr.ToError()

	}
	return nil
}

func FilterInvalidChars(input string) string {
	var validChars []rune
	for _, r := range input {
		if (r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == '-' || r == '_' || r == '.' {
			validChars = append(validChars, r)
		}
	}
	filtered := string(validChars)
	filtered = strings.TrimFunc(filtered, func(r rune) bool {
		return r == '-' || r == '_' || r == '.'
	})
	if len(filtered) > 80 {
		filtered = filtered[:80] // 截断至100字符
	}
	return filtered
}

func CreateNewReadme(oldDataset models.OldDataset, dataset *models.DatasetRegistry) error {
	var desc, datasetDesc, attachmentDesc string
	if oldDataset.Dataset != nil {
		datasetDesc = oldDataset.Dataset.Description
	}
	if oldDataset.Attachment != nil {
		attachmentDesc = oldDataset.Attachment.Description
	}
	if datasetDesc == "" && attachmentDesc == "" {
		log.Info("handleOldDataset Old dataset attachment or description is empty, skipping readme creation.")
		return nil
	}
	if datasetDesc != "" {
		desc = datasetDesc
	}
	if attachmentDesc != "" && strings.ToLower(desc) != strings.ToLower(attachmentDesc) {
		if desc != "" {
			desc += "\n\n"
		}
		desc += attachmentDesc
	}

	PutDatasetReadme(entity.DatasetReadmeReq{
		Content: desc,
		Dataset: dataset,
	})
	return nil
}

func CreateNewVersionFile(oldDataset models.OldDataset, dataset *models.DatasetRegistry) error {
	info := map[string]interface{}{versionKey: oldDataset.Attachment.CreatedUnix}
	helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if helper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", dataset.StorageType)
		return errors.New("Data type error")
	}

	metaInfo, _ := json.Marshal(info)
	objectKey := path.Join(helper.TrimBucketPrefix(dataset.Path), versionFileName)
	err := helper.UploadFile(objectKey, strings.NewReader(string(metaInfo)))
	return err
}

// 全局状态管理器
type ProductionManager struct {
	mu                 sync.Mutex
	isRunning          bool
	ctx                context.Context
	cancel             context.CancelFunc
	taskQueue          chan models.OldDataset
	wg                 sync.WaitGroup
	stats              Stats
	producerCount      int
	consumerCount      int
	closeTaskQueueOnce sync.Once
}

var (
	pm     *ProductionManager
	pmOnce sync.Once
)

// 获取生产管理器实例（单例）
func GetProductionManager() *ProductionManager {
	pmOnce.Do(func() {
		pm = NewProductionManager(2, 5)
	})
	return pm
}

type Stats struct {
	Produced     int64 `json:"produced"`
	Consumed     int64 `json:"consumed"`
	ProducerErrs int64 `json:"producer_errors"`
	ConsumerErrs int64 `json:"consumer_errors"`
}

func NewProductionManager(producerCount, consumerCount int) *ProductionManager {
	return &ProductionManager{
		taskQueue: make(chan models.OldDataset, 1000),
	}
}

// 启动生产消费流程
func (pm *ProductionManager) Start(producerCount, consumerCount int) error {

	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.isRunning {
		return errors.New("production is already running")
	}
	// 重新初始化 taskQueue（如果已关闭）
	if pm.taskQueue == nil || isChanClosed(pm.taskQueue) {
		pm.taskQueue = make(chan models.OldDataset, 1000)
	}
	pm.producerCount = producerCount
	pm.consumerCount = consumerCount
	// 重置 context 和 waitgroup
	pm.ctx, pm.cancel = context.WithCancel(context.Background())
	pm.wg = sync.WaitGroup{}
	pm.isRunning = true
	pm.stats = Stats{} // 重置统计

	// 启动生产者
	pm.wg.Add(pm.producerCount)
	for i := 0; i < pm.producerCount; i++ {
		go pm.runProducer(i)
	}

	// 启动消费者
	pm.wg.Add(pm.consumerCount)
	for i := 0; i < pm.consumerCount; i++ {
		go pm.runConsumer(i)
	}

	return nil
}

func isChanClosed(ch chan models.OldDataset) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func (pm *ProductionManager) Stop() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if !pm.isRunning {
		return errors.New("production is not running")
	}

	pm.cancel()
	pm.isRunning = false

	// 异步等待所有 goroutine 完成后再关闭 channel
	go func() {
		defer func() {
			if err := recover(); err != nil {
				combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
				log.Error("PANIC:%v", combinedErr)
			}
		}()

		pm.wg.Wait()

		// 使用 sync.Once 确保 channel 只关闭一次
		pm.closeTaskQueueOnce.Do(func() {
			pm.mu.Lock()
			defer pm.mu.Unlock()
			if pm.taskQueue != nil {
				close(pm.taskQueue)
			}
		})
	}()

	return nil
}

func (pm *ProductionManager) GetStatus() (bool, Stats) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	return pm.isRunning, pm.stats
}

func (pm *ProductionManager) runProducer(producerID int) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	defer pm.wg.Done()

	page := producerID + 1 // 简单分配策略，每个生产者从不同页开始
	pageSize := 200

	for {
		select {
		case <-pm.ctx.Done():
			fmt.Printf("Producer %d shutting down...\n", producerID)
			return
		default:
			records, err := models.GetUnhandledOldDataset(page, pageSize)
			// records, err := pm.db.QueryByPage(pm.ctx, page, pageSize)
			if err != nil {
				atomic.AddInt64(&pm.stats.ProducerErrs, 1)
				continue
			}

			if len(records) == 0 {
				fmt.Printf("Producer %d completed (no more records)\n", producerID)
				return
			}

			for _, record := range records {
				if pm.taskQueue == nil {
					return
				}
				select {
				case pm.taskQueue <- record:
					atomic.AddInt64(&pm.stats.Produced, 1)
					fmt.Printf("Producer %d produced record %d\n", producerID, record.Attachment.ID)
				case <-pm.ctx.Done():
					return
				}
			}
			page += pm.producerCount // 跳到下一个分配给该生产者的页
		}
	}
}

func (pm *ProductionManager) runConsumer(consumerID int) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()

	defer pm.wg.Done()

	for {
		select {
		case <-pm.ctx.Done():
			fmt.Printf("Consumer %d shutting down...\n", consumerID)
			return
		case record, ok := <-pm.taskQueue:
			if !ok {
				fmt.Printf("Consumer %d completed (channel closed)\n", consumerID)
				return
			}

			if err := handleOldDataset(record); err != nil {
				atomic.AddInt64(&pm.stats.ConsumerErrs, 1)
				continue
			}
			atomic.AddInt64(&pm.stats.Consumed, 1)
			fmt.Printf("Consumer %d processed record %d\n", consumerID, record.Attachment.ID)
		}
	}
}

func StartDatasetMigrate(producerCount, consumerCount int) error {
	pm := GetProductionManager()
	if err := pm.Start(producerCount, consumerCount); err != nil {
		return err
	}
	return nil
}

func StopDatasetMigrate() error {
	pm := GetProductionManager()
	if err := pm.Stop(); err != nil {
		return err
	}
	return nil
}

func GetDatasetMigrateStatus() map[string]interface{} {
	pm := GetProductionManager()
	m := make(map[string]interface{}, 0)
	m["isRunning"] = pm.isRunning
	m["producerCount"] = pm.producerCount
	m["consumerCount"] = pm.consumerCount
	m["stats"] = pm.stats

	return m
}

func HandleOneOldDataset(uuid string) error {
	dataset, err := models.GetOldDataset(uuid)
	if err != nil {
		log.Error("GetOldDataset err.uuid=%s err=%v", uuid, err)
		return err
	}
	if dataset == nil {
		return errors.New("dataset not exsits or can not migrate")
	}
	k := redis_key.ProcessingOldDataset()
	redis_client.SAdd(k, uuid)
	redis_client.Expire(k, 24*time.Hour)
	go handleOldDataset(*dataset)
	return nil
}

func IsOldDatasetNameChanged(uuid string, oldCloudbrainId int64) bool {
	task, err := models.GetCloudbrainByCloudbrainID(oldCloudbrainId)
	if err != nil {
		return true
	}
	dataset, _ := models.GetDatasetRegistryByID(uuid)
	if dataset == nil {
		return true
	}

	uuidStr := task.Uuid
	nameStr := task.DatasetName
	if uuidStr == "" || nameStr == "" {
		return true
	}
	ids := strings.Split(uuidStr, ";")
	names := strings.Split(nameStr, ";")
	if len(names) != len(ids) {
		return true
	}

	index := -1
	for i := 0; i < len(ids); i++ {
		if ids[i] == uuid {
			index = i
		}
	}
	oldName := ""
	if index >= 0 {
		oldName = names[index]
	}
	if oldName == "" {
		return true
	}
	oldName = strings.TrimSuffix(strings.TrimSuffix(oldName, ".zip"), ".tar.gz")
	return dataset.Name != oldName
}
