// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	repo_module "code.gitea.io/gitea/modules/repository"
	"code.gitea.io/gitea/modules/setting"
	pull_service "code.gitea.io/gitea/services/pull"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"xorm.io/xorm"
)

var promoteHTTPClient = &http.Client{Timeout: 3 * time.Second}

const SHELL_FLAG_ON = 1

// CreateRepository creates a repository for the user/organization.
func CreateRepository(doer, owner *models.User, opts models.CreateRepoOptions) (*models.Repository, error) {
	repo, err := repo_module.CreateRepository(doer, owner, opts)
	if err != nil {
		if repo != nil {
			if errDelete := models.DeleteRepository(doer, owner.ID, repo.ID); errDelete != nil {
				log.Error("Rollback deleteRepository: %v", errDelete)
			}
		}
		return nil, err
	}

	notification.NotifyCreateRepository(doer, owner, repo)

	return repo, nil
}

// ForkRepository forks a repository
func ForkRepository(doer, u *models.User, oldRepo *models.Repository, name, desc, alias string) (*models.Repository, error) {
	repo, err := repo_module.ForkRepository(doer, u, oldRepo, name, desc, alias)
	if err != nil {
		if repo != nil {
			if errDelete := models.DeleteRepository(doer, u.ID, repo.ID); errDelete != nil {
				log.Error("Rollback deleteRepository: %v", errDelete)
			}
		}
		return nil, err
	}

	notification.NotifyForkRepository(doer, oldRepo, repo)

	return repo, nil
}

// DeleteRepository deletes a repository for a user or organization.
func DeleteRepository(doer *models.User, repo *models.Repository) error {
	if err := pull_service.CloseRepoBranchesPulls(doer, repo); err != nil {
		log.Error("CloseRepoBranchesPulls failed: %v", err)
	}

	if err := models.DeleteRepository(doer, repo.OwnerID, repo.ID); err != nil {
		return err
	}

	notification.NotifyDeleteRepository(doer, repo)

	return nil
}

// PushCreateRepo creates a repository when a new repository is pushed to an appropriate namespace
func PushCreateRepo(authUser, owner *models.User, repoName string) (*models.Repository, error) {
	if !authUser.IsAdmin {
		if owner.IsOrganization() {
			if ok, err := owner.CanCreateOrgRepo(authUser.ID); err != nil {
				return nil, err
			} else if !ok {
				return nil, fmt.Errorf("cannot push-create repository for org")
			}
		} else if authUser.ID != owner.ID {
			return nil, fmt.Errorf("cannot push-create repository for another user")
		}
	}

	repo, err := CreateRepository(authUser, owner, models.CreateRepoOptions{
		Name:      repoName,
		IsPrivate: true,
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func GetRecommendCourseKeyWords() ([]string, error) {

	url := setting.RecommentRepoAddr + "course_keywords"
	result, err := RecommendFromPromote(url)

	if err != nil {
		return []string{}, err
	}
	return result, err

}

func GetRecommendRepoFromPromote(repoMap []map[string]string) ([]map[string]interface{}, error) {
	resultRepo := make([]map[string]interface{}, 0)

	//resultRepo := make([]*models.Repository, 0)
	for _, record := range repoMap {
		repoName := record["project_url"]
		//log.Info("repoName=" + repoName + " tmpIndex1=" + fmt.Sprint(tmpIndex1) + " len(repoName)=" + fmt.Sprint(len(repoName)))
		tmpIndex := strings.Index(repoName, "/")
		if tmpIndex == -1 {
			log.Info("error repo name format.")
		} else {
			ownerName := strings.Trim(repoName[0:tmpIndex], " ")
			repoName := strings.Trim(repoName[tmpIndex+1:], " ")
			repo, err := models.GetRepositoryByOwnerAndAlias(ownerName, repoName)
			if err == nil {
				repoMap := make(map[string]interface{})
				repoMap["ID"] = fmt.Sprint(repo.ID)
				repoMap["Name"] = repo.Name
				repoMap["Alias"] = repo.Alias
				repoMap["Label"] = record["class"]
				repoMap["Label_en"] = record["class_en"]
				repoMap["OwnerName"] = repo.OwnerName
				repoMap["NumStars"] = repo.NumStars
				repoMap["NumForks"] = repo.NumForks
				repoMap["Description"] = repo.Description
				repoMap["NumWatchs"] = repo.NumWatches
				repoMap["Topics"] = repo.Topics
				repoMap["Avatar"] = repo.RelAvatarLink()
				resultRepo = append(resultRepo, repoMap)
			} else {
				log.Info("query repo error," + err.Error())
			}
		}
	}
	return resultRepo, nil
}

func RecommendContentFromPromote(url string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("not error.", err)
			return
		}
	}()
	resp, err := promoteHTTPClient.Get(url)
	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Info("Get organizations url error")
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return "", err
	}
	allLineStr := string(bytes)
	return allLineStr, nil
}

func RecommendFromPromote(url string) ([]string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("not error.", err)
			return
		}
	}()
	resp, err := promoteHTTPClient.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Info("Get organizations url error")
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Info("Get organizations url error=" + err.Error())
		return nil, err
	}

	allLineStr := string(bytes)
	lines := strings.Split(allLineStr, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		log.Info("i=" + fmt.Sprint(i) + " line=" + line)
		result[i] = strings.Trim(line, " ")
	}
	return result, nil
}

func CheckPushSizeLimit4Web(repo *models.Repository, fileIds []string) error {
	if err := CheckRepoNumOnceLimit(len(fileIds)); err != nil {
		return err
	}
	totalSize, err := CountUploadFileSizeByIds(fileIds)
	if err != nil {
		return UploadFileInvalidErr{}
	}
	if err := CheckRepoTotalSizeLimit(repo, totalSize); err != nil {
		return err
	}
	return nil
}

func CheckPushSizeLimit4Http(repo *models.Repository, uploadFileSize int64) error {
	if err := CheckRepoOnceTotalSizeLimit(uploadFileSize); err != nil {
		return err
	}
	if err := CheckRepoTotalSizeLimit(repo, uploadFileSize); err != nil {
		return err
	}
	return nil
}

func CheckRepoTotalSizeLimit(repo *models.Repository, uploadFileSize int64) error {
	if repo.Size+uploadFileSize > setting.Repository.RepoMaxSize*1024*1024 {
		return RepoTooLargeErr{}
	}
	return nil
}

func CheckRepoOnceTotalSizeLimit(uploadFileSize int64) error {
	if uploadFileSize > setting.Repository.Upload.TotalMaxSize*1024*1024 {
		return UploadFileTooLargeErr{}
	}
	return nil
}

func CheckRepoNumOnceLimit(uploadFileNum int) error {
	if uploadFileNum > setting.Repository.Upload.MaxFiles {
		return UploadFileTooMuchErr{}
	}
	return nil
}

func CountUploadFileSizeByIds(fileIds []string) (int64, error) {
	if len(fileIds) == 0 {
		return 0, nil
	}
	uploads, err := models.GetUploadsByUUIDs(fileIds)
	if err != nil {
		return 0, fmt.Errorf("CountUploadFileSizeByIds error [uuids: %v]: %v", fileIds, err)
	}
	var totalSize int64
	for _, upload := range uploads {
		size, err := GetUploadFileSize(upload)
		if err != nil {
			return 0, err
		}
		totalSize += size
	}
	return totalSize, nil
}

func GetUploadFileSize(upload *models.Upload) (int64, error) {
	info, err := os.Lstat(upload.LocalPath())

	if err != nil {
		return 0, err
	}
	return info.Size(), nil

}

type RepoTooLargeErr struct {
}

func (RepoTooLargeErr) Error() string {
	return fmt.Sprintf("Repository can not exceed %d MB. Please remove some unnecessary files and try again", setting.Repository.RepoMaxSize)
}

func IsRepoTooLargeErr(err error) bool {
	_, ok := err.(RepoTooLargeErr)
	return ok
}

type UploadFileTooLargeErr struct {
}

func (UploadFileTooLargeErr) Error() string {
	return fmt.Sprintf("Upload files can not exceed %d MB at a time", setting.Repository.Upload.TotalMaxSize)
}

func IsUploadFileTooLargeErr(err error) bool {
	_, ok := err.(UploadFileTooLargeErr)
	return ok
}

type RepoFileTooLargeErr struct {
}

func (RepoFileTooLargeErr) Error() string {
	return "repository file is too large"
}

func IsRepoFileTooLargeErr(err error) bool {
	_, ok := err.(RepoFileTooLargeErr)
	return ok
}

type UploadFileTooMuchErr struct {
}

func (UploadFileTooMuchErr) Error() string {
	return "upload files are too lmuch"
}

func IsUploadFileTooMuchErr(err error) bool {
	_, ok := err.(UploadFileTooMuchErr)
	return ok
}

type UploadFileInvalidErr struct {
}

func (UploadFileInvalidErr) Error() string {
	return "upload files are invalid"
}

func IsUploadFileInvalidErr(err error) bool {
	_, ok := err.(UploadFileInvalidErr)
	return ok
}

func IncreaseRepoDatasetNum(datasetID int64, engines ...*xorm.Engine) error {
	dataset, err := models.GetDatasetByID(datasetID)
	if err != nil {
		return err
	}
	return models.OperateRepoDatasetNum(dataset.RepoID, 1, engines...)
}

func IncreaseRepoModelNum(repoId int64, engines ...*xorm.Engine) error {
	return models.OperateRepoModelNum(repoId, 1, engines...)
}

func ResetRepoModelNum(repoId int64) error {
	return models.ResetRepoModelNum(repoId)
}

func DecreaseRepoDatasetNum(datasetID int64, engines ...*xorm.Engine) error {
	dataset, err := models.GetDatasetByID(datasetID)
	if err != nil {
		return err
	}
	return models.OperateRepoDatasetNum(dataset.RepoID, -1, engines...)
}

func DecreaseRepoModelNum(repoId int64, engines ...*xorm.Engine) error {
	return models.OperateRepoModelNum(repoId, -1, engines...)
}

func GetDefaultBranchName(repo *models.Repository) string {
	gitRepo, err := git.OpenRepository(repo.RepoPath())
	if err != nil {
		return ""
	}
	defer gitRepo.Close()
	if len(repo.DefaultBranch) > 0 && gitRepo.IsBranchExist(repo.DefaultBranch) {
		return repo.DefaultBranch
	}
	brs, _, err := gitRepo.GetBranches(0, 0)
	if len(brs) > 0 {
		return brs[0]
	}
	return ""
}
