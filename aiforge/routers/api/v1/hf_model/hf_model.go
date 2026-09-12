package hf_model

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/services/subject_service"
	"fmt"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/hf_model_service"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	routerRepo "code.gitea.io/gitea/routers/repo"
)

const (
	RepoIdPatterns     = "^[.a-zA-Z0-9_-]+/[.a-zA-Z0-9_-]+$"
	HfModelLabel       = "huggingface"
	FileUpdateSame     = 0
	FileUpdateNew      = 1
	FileUpdateDeleted  = 2
	FileUpdateModified = 3
)

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func initResponseJson() *APIResponse {
	return &APIResponse{
		Code: -1,
		Msg:  "",
		Data: nil,
	}
}

func handleInternalError(ctx *context.APIContext, err error, msg string) {
	log.Error("[hf_model]: " + err.Error())
	ctx.JSON(http.StatusOK, &APIResponse{
		Code: -1,
		Msg:  ctx.Tr("hf_model.internal_error", msg),
	})
}

func CreateAimodelNew(doer *models.User, userReq NewTransferRequest, hfRepoId string, modelName string) (*models.AiModelManage, error) {
	modelOwner, err := GetModelOwner()
	if err != nil {
		log.Error("[hf_model]: cannot get model owner: " + err.Error())
		return nil, err
	}

	req := entity.CreateAimodelReq{
		Name:         modelName,
		Alias:        modelName,
		AimodelType:  models.MODEL_HF_TYPE,
		Engine:       int64(userReq.Engine),
		License:      "",
		Label:        HfModelLabel + " " + userReq.Label,
		IsPrivate:    userReq.IsPrivate,
		CreatorId:    doer.ID,
		OwnerId:      modelOwner.ID,
		ExternalName: hfRepoId,
	}
	res, errRes := subject_service.CreateAimodel(doer, modelOwner, req)
	if errRes != nil {
		return nil, errRes.ToError()
	}
	return res, nil
}

func CheckModelSize(files *[]hf_model_service.Sibling, maxSizeGb int) bool {
	modelSizeTotal := 0
	maxSizeBytes := maxSizeGb * 1024 * 1024 * 1024
	for _, fileData := range *files {
		modelSizeTotal += int(fileData.Size)
	}
	return modelSizeTotal <= maxSizeBytes
}

func SplitHfRepoID(hfRepoId string) (string, string, error) {
	_, err := regexp.MatchString(RepoIdPatterns, hfRepoId)
	if err != nil {
		return "", "", err
	}

	splitResult := strings.Split(hfRepoId, "/")
	repoName, modelName := splitResult[0], splitResult[1]

	return repoName, modelName, nil
}

func GetModelOwner() (*models.User, error) {
	username := setting.ExternalTransfer.HfModelOwner
	if username == "" {
		return nil, fmt.Errorf("hf_model_owner not set")
	}

	modelOwner, err := models.GetUserByName(username)
	if err != nil {
		log.Error("[hf_model]: cannot get model owner: " + err.Error())
		return nil, err
	}

	return modelOwner, nil
}

func CheckMaxModelTransfer(ctx *context.APIContext, resp *APIResponse) *APIResponse {
	creatorModelIds, _ := models.GetCreatorModelIds(ctx.User.ID)
	operatorModelIds, _ := models.GetOperatorModelIds(ctx.User.ID)
	modelIds := append(creatorModelIds, operatorModelIds...)

	var runningCnt int
	for _, modelId := range modelIds {
		aiModel, _ := models.QueryModelById(modelId)
		if aiModel == nil {
			continue
		}
		if aiModel.Status == models.HfTransferStatusOngoing || aiModel.Status == models.HfTransferStatusWaiting {
			runningCnt++
		}
	}

	if runningCnt >= setting.ExternalTransfer.MaxPerUser {
		resp.Msg = ctx.Tr("hf_model.transfer_running", setting.ExternalTransfer.MaxPerUser)
		resp.Code = -1
	}
	return resp
}

func fetchHfModelInfo(ctx *context.APIContext, resp *APIResponse, hfRepoId string, hfToken string) (*hf_model_service.HfModelInfo, int) {
	hfModelInfo, statusCode, err := hf_model_service.GetHfModelInfoResty(hfRepoId, hfToken)
	if err != nil {
		resp.Code = -1
		switch statusCode {
		case 404:
			resp.Msg = ctx.Tr("hf_model.model_not_exist")
		case 502, 503:
			resp.Msg = ctx.Tr("hf_model.server_error")
		case 401, 403:
			resp.Msg = ctx.Tr("hf_model.model_no_access", hfRepoId)
		default:
			resp.Msg = err.Error()
		}
		resp.Data = nil
		ctx.JSON(http.StatusOK, resp)
		return nil, statusCode
	}
	return hfModelInfo, statusCode
}

func validateHfToken(ctx *context.APIContext, resp *APIResponse, hfRepoId string, hfToken string) string {
	if hfToken == "" {
		resp.Msg = ctx.Tr("hf_model.model_no_access", hfRepoId)
		resp.Data = nil
		ctx.JSON(http.StatusOK, resp)
		return ""
	}

	code, _ := hf_model_service.ValidTokenAccess(hfRepoId, hfToken)
	log.Info("validateHfToken status code: %v\n", code)
	if code != 200 {
		resp.Msg = ctx.Tr("hf_model.token_invalid", hfToken, hfRepoId)
		resp.Data = nil
		ctx.JSON(http.StatusOK, resp)
		return ""
	}

	return hfToken
}

type NewTransferRequest struct {
	HfToken     string `json:"hf_token"`
	HfRepoID    string `json:"hf_repo_id"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Engine      int    `json:"engine,omitempty"`
	IsPrivate   bool   `json:"isPrivate omitempty"`
}

func NewTransferAPI(ctx *context.APIContext, data NewTransferRequest) {
	resp := initResponseJson()
	hfTokenUser := data.HfToken

	// check user's running or waiting model transfer
	resp = CheckMaxModelTransfer(ctx, resp)
	if resp.Msg != "" {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// get model_owner
	modelOwner, err := GetModelOwner()
	if err != nil {
		handleInternalError(ctx, err, "failed to get model owner")
	}

	// split repo/name
	hfRepoId := data.HfRepoID
	repoName, modelName, err := SplitHfRepoID(hfRepoId)
	if err != nil {
		handleInternalError(ctx, err, "failed to split hf_repo_id")
	}

	existedAimodel, err := models.QueryAimodelByExternalName(hfRepoId)
	if existedAimodel != nil {
		resp.Msg = ctx.Tr("hf_model.transfer_exist")
		resp.Code = 2
		ctx.JSON(http.StatusOK, resp)
		return

	}

	// validate and fetch model info
	hfToken := ""
	hfModelInfo, _ := fetchHfModelInfo(ctx, resp, hfRepoId, hfToken)
	if hfModelInfo == nil {
		return
	}

	// Gated model
	if hfModelInfo.Gated != "false" {
		hfToken = validateHfToken(ctx, resp, hfRepoId, hfTokenUser)
		if hfToken == "" {
			return
		}
	}

	// check model size
	files := hfModelInfo.Siblings
	maxSizeGb := setting.ExternalTransfer.MaxModelSizeGb
	ok := CheckModelSize(&files, maxSizeGb)
	if !ok {
		log.Error("[hf_model]: model size exceeds the limit")
		resp.Msg = ctx.Tr("hf_model.model_size_exceeds", maxSizeGb)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	aiModel, err := CreateAimodelNew(ctx.User, data, hfRepoId, modelName)
	if err != nil {
		handleInternalError(ctx, err, "failed to create model")
	}

	// insert data table
	//err = CreateFileRecords(hfModelInfo, aiModel.Path, aiModel.ID, modelRepo.ID, ctx.User.ID, modelOwner.ID)
	for _, fileData := range hfModelInfo.Siblings {
		filename := fileData.Rfilename
		storagePath := path.Join(aiModel.Path, filename)
		openiRepoId := path.Join(modelOwner.Name, repoName)
		record := &models.HfModelFile{
			Filename:    filename,
			Status:      models.HfTransferStatusWaiting,
			HfRepoId:    hfModelInfo.ModelId,
			OpeniRepoId: openiRepoId,
			ModelName:   modelName,
			Path:        storagePath,
			Size:        fileData.Size,
			HfBlobId:    fileData.BlobId,
			ModelId:     aiModel.ID,
			//RepoId:      modelRepo.ID,
			OwnerId: modelOwner.ID,
			UserId:  ctx.User.ID,
			HfToken: hfToken,
		}
		err = models.SaveRecord(record)
		if err != nil {
			handleInternalError(ctx, err, "failed to save file records")
		}
	}

	resp.Code = 1
	resp.Msg = "success"
	resp.Data = aiModel
	ctx.JSON(http.StatusOK, resp)

}

func GetTransferStatusAPI(ctx *context.APIContext) {
	resp := initResponseJson()

	hfRepoId := ctx.Query("hf_repo_id")
	modelOwner, err := GetModelOwner()
	if err != nil {
		handleInternalError(ctx, err, "failed to get model owner")
	}

	filesRecords, err := models.QueryModelFiles(hfRepoId, modelOwner)
	if err != nil {
		handleInternalError(ctx, err, "failed to fetch model files")
	}

	resp.Code = 1
	resp.Msg = "success"
	resp.Data = filesRecords
	ctx.JSON(http.StatusOK, resp)
}

type FileUpdate struct {
	Filename     string `json:"filename"`
	BlobId       string `json:"blobId"`
	Size         int64  `json:"size"`
	BlobIdNew    string `json:"blobId_new"`
	SizeNew      int64  `json:"size_new"`
	UpdateStatus int64  `json:"update_status"` // same, new, deleted, modified
}

type ModelUpdate struct {
	HfToken   string       `json:"hf_token"`
	IsUpdate  bool         `json:"is_update"`
	HfRepoId  string       `json:"hf_repo_id"`
	RepoId    int64        `json:"repo_id"`
	ModelId   string       `json:"model_id"`
	ModelPath string       `json:"model_path"`
	Files     []FileUpdate `json:"files"`
	Gated     bool         `json:"gated"`
}

func FetchHfUpdateAPI(ctx *context.APIContext) {
	resp := initResponseJson()
	hfRepoId := ctx.Query("hf_repo_id")
	modelOwner, err := GetModelOwner()
	if err != nil {
		handleInternalError(ctx, err, "failed to get model owner")
	}

	// validate and fetch model info
	hfToken := ""
	hfModelInfo, statusCode := fetchHfModelInfo(ctx, resp, hfRepoId, hfToken)
	if hfModelInfo == nil {
		return
	}

	// Gated model
	isGated := false
	if hfModelInfo.Gated != "false" {
		isGated = true
	}

	newFiles := hfModelInfo.Siblings
	oldFiles, err := models.QueryModelFiles(hfRepoId, modelOwner)
	if err != nil {
		handleInternalError(ctx, err, "failed to fetch model files")
	}

	aiModelId := oldFiles[0].ModelId
	aiModelPath := strings.TrimSuffix(oldFiles[0].Path, oldFiles[0].Filename)
	repoId := oldFiles[0].RepoId

	isUpdated := false
	fileUpdates := make([]FileUpdate, 0)
	oldFileMap := make(map[string]*models.HfModelFile)
	newFileMap := make(map[string]hf_model_service.Sibling)

	for _, oldFile := range oldFiles {
		oldFileMap[oldFile.Filename] = oldFile
	}

	for _, newFile := range newFiles {
		newFileMap[newFile.Rfilename] = newFile
	}

	allFilenames := make(map[string]bool)
	for filename := range oldFileMap {
		allFilenames[filename] = true
	}
	for filename := range newFileMap {
		allFilenames[filename] = true
	}

	for filename := range allFilenames {
		fileUpdate := FileUpdate{
			Filename: filename,
		}

		if oldFile, exists := oldFileMap[filename]; exists {
			fileUpdate.BlobId = oldFile.HfBlobId
			fileUpdate.Size = oldFile.Size
		}

		if newFile, exists := newFileMap[filename]; exists {
			fileUpdate.BlobIdNew = newFile.BlobId
			fileUpdate.SizeNew = newFile.Size
		}

		if fileUpdate.BlobId == "" {
			fileUpdate.UpdateStatus = FileUpdateNew
			isUpdated = true
		} else if fileUpdate.BlobIdNew == "" {
			fileUpdate.UpdateStatus = FileUpdateDeleted
			isUpdated = true
		} else if fileUpdate.BlobId != fileUpdate.BlobIdNew || fileUpdate.Size != fileUpdate.SizeNew {
			fileUpdate.UpdateStatus = FileUpdateModified
			isUpdated = true
		} else {
			fileUpdate.UpdateStatus = FileUpdateSame
		}

		fileUpdates = append(fileUpdates, fileUpdate)
	}

	modelUpdate := &ModelUpdate{
		IsUpdate:  isUpdated,
		ModelId:   aiModelId,
		RepoId:    repoId,
		ModelPath: aiModelPath,
		HfRepoId:  hfRepoId,
		Files:     fileUpdates,
		Gated:     isGated,
	}
	resp.Code = statusCode
	resp.Msg = "success"
	resp.Data = modelUpdate

	ctx.JSON(http.StatusOK, resp)
}

func UpdateHfModelAPI(ctx *context.APIContext, data ModelUpdate) {
	files := data.Files
	hfRepoId := data.HfRepoId
	repoId := data.RepoId
	aiModelId := data.ModelId
	aiModelPath := data.ModelPath
	hfToken := data.HfToken

	resp := initResponseJson()

	resp = CheckMaxModelTransfer(ctx, resp)
	if resp.Msg != "" {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// Gated model
	if hfToken != "" {
		hfToken = validateHfToken(ctx, resp, hfRepoId, hfToken)
		if hfToken == "" {
			return
		}
	}

	modelOwner, err := GetModelOwner()
	if err != nil {
		handleInternalError(ctx, err, "failed to get model owner")
	}

	modelUserId, err := models.GetUserIdByModelId(aiModelId)
	if err != nil {
		handleInternalError(ctx, err, "failed to fetch model user id")
	}

	// split repo/name
	repoName, modelName, err := SplitHfRepoID(hfRepoId)
	if err != nil {
		handleInternalError(ctx, err, "failed to split hf_repo_id")
	}
	openiRepoId := path.Join(modelOwner.Name, repoName)

	var aiModelSize int64
	maxAimodelSize := int64(setting.ExternalTransfer.MaxModelSizeGb * 1024 * 1024 * 1024)
	for _, fileUpdate := range files {
		aiModelSize += fileUpdate.SizeNew
	}
	if aiModelSize >= maxAimodelSize {
		log.Error("[hf_model]: model size exceeds the limit")
		resp.Msg = ctx.Tr("hf_model.model_size_exceeds", setting.ExternalTransfer.MaxModelSizeGb)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// 处理文件更新
	for _, fileUpdate := range files {
		filename := fileUpdate.Filename
		storagePath := path.Join(aiModelPath, filename)

		switch fileUpdate.UpdateStatus {
		case FileUpdateNew:
			// 添加新记录
			record := &models.HfModelFile{
				Filename:    filename,
				Status:      models.HfTransferStatusWaiting,
				HfRepoId:    hfRepoId,
				OpeniRepoId: openiRepoId,
				ModelName:   modelName,
				Path:        storagePath,
				Size:        fileUpdate.SizeNew,
				HfBlobId:    fileUpdate.BlobIdNew,
				ModelId:     aiModelId,
				RepoId:      repoId,
				OwnerId:     modelOwner.ID,
				UserId:      modelUserId,
				HfToken:     hfToken,
			}
			models.SaveRecord(record)
		case FileUpdateDeleted:
			// 删除记录和实际文件
			models.DeleteRecord(hfRepoId, modelOwner, filename)
			err = routerRepo.DeleteModelFileInternal(aiModelId, filename)
		case FileUpdateModified:
			// 更新记录和删除旧文件
			models.UpdateRecord(hfRepoId, modelOwner, filename, fileUpdate.BlobIdNew, fileUpdate.SizeNew, models.HfTransferStatusWaiting, hfToken)
			err = routerRepo.DeleteModelFileInternal(aiModelId, filename)
		case FileUpdateSame:
			// 不做任何操作
			continue
		}

		if err != nil {
			handleInternalError(ctx, err, "failed to update file records")
		}
	}

	models.UpdateAimodelByStatus(aiModelId, models.HfTransferStatusWaiting, "Updating")
	if err != nil {
		handleInternalError(ctx, err, "failed updating model records status")
	}

	err = models.AddOperation(&models.HfModelOperation{
		ModelId: aiModelId,
		UserId:  ctx.User.ID,
		OpType:  models.HfOperationTypeUpdate,
	})
	if err != nil {
		handleInternalError(ctx, err, "failed to add operation record")
	}

	resp.Code = 1
	resp.Msg = "success"
	ctx.JSON(http.StatusOK, resp)
}

func getUserById(userId int64) (*models.User, error) {
	user, err := models.GetUserByID(userId)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			user = models.NewGhostUser()
		} else {
			return nil, err
		}
	}
	return user, nil
}

func GetTransferUserAPI(ctx *context.APIContext) {
	modelId := ctx.Query("model_id")

	creatorId, err := models.GetUserIdByModelId(modelId)
	if err != nil {
		handleInternalError(ctx, err, "failed fetching model user")
		return
	}
	operatorIds, err := models.GetUniqueOperatorIds(modelId)
	if err != nil {
		handleInternalError(ctx, err, "failed fetching model user")
		return
	}

	uniqueUserIds := make(map[int64]struct{})
	uniqueUserIds[creatorId] = struct{}{}
	for _, operatorId := range operatorIds {
		uniqueUserIds[operatorId] = struct{}{}
	}

	var uniqueUsernames []string
	for userId := range uniqueUserIds {
		userObj, err := getUserById(userId)
		if err != nil {
			handleInternalError(ctx, err, "failed fetching model user")
			return
		}
		uniqueUsernames = append(uniqueUsernames, userObj.Name)
	}
	sort.Strings(uniqueUsernames)

	ctx.JSON(http.StatusOK, &APIResponse{
		Code: 1,
		Msg:  "success",
		Data: uniqueUsernames,
	})
}

type RetryTransferRequest struct {
	HfToken  string `json:"hf_token"`
	HfRepoID string `json:"hf_repo_id"`
}

func RetryTransferAPI(ctx *context.APIContext, data RetryTransferRequest) {
	resp := initResponseJson()
	hfRepoId := data.HfRepoID

	resp = CheckMaxModelTransfer(ctx, resp)
	if resp.Msg != "" {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	modelOwner, err := GetModelOwner()
	if err != nil {
		handleInternalError(ctx, err, "failed to get model owner")
	}

	filesRecords, err := models.QueryModelFiles(hfRepoId, modelOwner)
	if err != nil {
		handleInternalError(ctx, err, "failed to fetch model files")
	}

	aiModelId := filesRecords[0].ModelId
	newAimodelStatus := models.HfTransferStatusWaiting

	success, _ := models.IsTransferSuccess(aiModelId)
	if success {
		newAimodelStatus = models.HfTransferStatusSuccess
	} else {
		for _, file := range filesRecords {
			if file.Status == models.HfTransferStatusFailure {
				file.Status = models.HfTransferStatusWaiting
				err := models.UpdateHfFile(file)
				if err != nil {
					log.Error("[hf_model]:Error updating transfer record:" + err.Error())
					continue
				}
			}
		}
	}

	models.UpdateAimodelByStatus(aiModelId, newAimodelStatus, "")
	if err != nil {
		log.Error("UpdateAimodelByStatus faild." + err.Error())
	}
	if err != nil {
		handleInternalError(ctx, err, "failed updating model records status")
	}

	err = models.AddOperation(&models.HfModelOperation{
		ModelId: aiModelId,
		UserId:  ctx.User.ID,
		OpType:  models.HfOperationTypeRetry,
	})
	if err != nil {
		handleInternalError(ctx, err, "failed to add operation record")
	}

	resp.Code = 1
	resp.Msg = "success"
	resp.Data = filesRecords
	ctx.JSON(http.StatusOK, resp)
}
