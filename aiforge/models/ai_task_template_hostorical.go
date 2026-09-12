package models

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	gouuid "github.com/satori/go.uuid"
	"xorm.io/builder"
)

//select count(*) from repository  where updated_unix > 1718012867 and is_empty = false and size > 0 and is_mirror = false and ai_task_cnt > 0
func GetHistoricalTemplatePaging(page, pageSize int, updatedUnix timeutil.TimeStamp) ([]Repository, error) {
	var repos = make([]Repository, 0)
	err := x.Where("updated_unix > ? and is_empty = false and size > 0 and is_mirror = false and is_fork = false", updatedUnix).OrderBy("updated_unix asc").Limit(pageSize, (page-1)*pageSize).Find(&repos)
	return repos, err
}

func GetRepoReaders(repoId int64) ([]int64, error) {
	var userIds = make([]int64, 0)

	err := x.Table("access").Select("distinct(user_id)").Where("repo_id = ?", repoId).Find(&userIds)
	return userIds, err
}

func DeleteUserTemplateAndAddNew(repo Repository, userIds []int64, newTemplates []AITaskTemplate) error {
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	_, err = sess.Where(builder.In("owner_id", userIds).And(builder.Eq{"repo_id": repo.ID})).Delete(&AITaskTemplate{})
	if err != nil {
		return err
	}

	for _, id := range userIds {
		for i := 0; i < len(newTemplates); i++ {
			templateId := gouuid.NewV4().String()
			newTemplates[i].ID = templateId
			newTemplates[i].OwnerID = id
			for j := 0; j < len(newTemplates[i].DatasetList); j++ {
				newTemplates[i].DatasetList[j].TemplateID = templateId
			}
			for k := 0; k < len(newTemplates[i].ModelList); k++ {
				newTemplates[i].ModelList[k].TemplateID = templateId
			}
			_, err = sess.NoAutoTime().Insert(&newTemplates[i])
			if err != nil {
				return err
			}

			if len(newTemplates[i].DatasetList) > 0 {
				var tempDatasets = make([]*TemplateDatasets, 0, len(newTemplates[i].DatasetList))
				for l := 0; l < len(newTemplates[i].DatasetList); l++ {
					tempDatasets = append(tempDatasets, newTemplates[i].DatasetList[l].ToDTO())
				}
				_, err = sess.Insert(tempDatasets)
				if err != nil {
					return err
				}

			}

			if len(newTemplates[i].ModelList) > 0 {
				var tempModels = make([]*TemplateModels, 0, len(newTemplates[i].ModelList))
				for m := 0; m < len(newTemplates[i].ModelList); m++ {
					tempModels = append(tempModels, newTemplates[i].ModelList[m].ToDTO())
				}
				_, err = sess.Insert(tempModels)
				if err != nil {
					return err
				}
			}
		}

	}
	return sess.Commit()
}

// func ReadLatestFileInRepoWithValidCommiter(repo Repository, refName, treePath string) (string, []*User, *time.Time, *time.Time, error) {
// 	var err error
// 	repoPath := RepoPath(repo.OwnerName, repo.Name)
// 	gitRepo, err := git.OpenRepository(repoPath)
// 	if err != nil {
// 		log.Error("ReadLatestFileInRepoWithValidCommiter error when OpenRepository,error=%v", err)
// 		return "", nil, nil, nil, err
// 	}
// 	_, err = gitRepo.GetBranchCommitID(refName)
// 	if err != nil {
// 		log.Error("ReadLatestFileInRepoWithValidCommiter error when GetBranchCommitID,error=%v", err)
// 		return "", nil, nil, nil, err
// 	}

// 	commit, err := gitRepo.GetBranchCommit(refName)
// 	if err != nil {
// 		log.Error("ReadLatestFileInRepo error when GetBranchCommit,error=%v", err)
// 		return "", nil, nil, nil, err
// 	}

// 	blob, err := commit.GetBlobByPath(treePath)
// 	if err != nil {
// 		log.Error("ReadLatestFileInRepo error when GetBlobByPath,error=%v", err)
// 		return "", nil, nil, nil, err
// 	}

// 	reader, err := blob.DataAsync()
// 	if err != nil {
// 		return "", nil, nil, nil, err
// 	}
// 	defer func() {
// 		if err = reader.Close(); err != nil {
// 			log.Error("ReadLatestFileInRepo: Close: %v", err)
// 		}
// 	}()
// 	d, _ := ioutil.ReadAll(reader)

// 	newcommit, err := gitRepo.GetCommitByPath(treePath)
// 	if err != nil {
// 		log.Error("ReadLatestFileInRepoWithValidCommiter error when GetCommitByPath,error=%v", err)
// 		return "", nil, nil, nil, err
// 	}
// 	commits, err := newcommit.CommitsByRange(1, 100)
// 	if err != nil {
// 		return "", nil, nil, nil, err
// 	}

// 	var createdAt, updatedAt *time.Time
// 	if commits.Len() > 0 {
// 		earliest := commits.Back().Value.(*git.Commit)
// 		oldest := commits.Front().Value.(*git.Commit)
// 		createdAt = &earliest.Author.When
// 		updatedAt = &oldest.Author.When
// 	}
// 	emailSet := make(map[string]*User)
// 	for e := commits.Front(); e != nil; e = e.Next() {
// 		c := e.Value.(*git.Commit)

// 		if c.Author == nil {
// 			continue
// 		}

// 		email := c.Author.Email

// 		if _, exists := emailSet[email]; exists {
// 			continue
// 		}

// 		if user, _ := GetUserByEmail(email); user != nil {
// 			emailSet[email] = user
// 		}
// 	}

// 	users := make([]*User, 0, len(emailSet))
// 	for _, user := range emailSet {
// 		users = append(users, user)
// 	}

// 	if len(users) == 0 {
// 		//该文件没有有效的提交者平台账户，直接返回
// 		return "", nil, nil, nil, nil
// 	}

// 	return string(d), users, createdAt, updatedAt, nil
// }

func ReadLatestFileInRepoWithValidCommiter(repo Repository, treePath string, maxCommits int) (string, []*User, *time.Time, *time.Time, error) {
	// 1. 打开仓库
	repoPath := RepoPath(repo.OwnerName, repo.Name)
	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Error("Open repository failed: %v", err)
		return "", nil, nil, nil, err
	}

	// 2. 获取文件提交历史
	commits, err := GetFileCommits(gitRepo, treePath, maxCommits)
	if err != nil {
		log.Error("Get file commits failed: %v", err)
		return "", nil, nil, nil, err
	}

	// 3. 提取时间和用户信息
	if commits.Len() == 0 {
		return "", nil, nil, nil, fmt.Errorf("no commits found for file %s", treePath)
	}

	// 最早提交 = 文件创建时间
	earliestCommit := commits.Back().Value.(*git.Commit)
	createdAt := &earliestCommit.Author.When

	// 最晚提交 = 文件最后更新时间
	latestCommit := commits.Front().Value.(*git.Commit)
	updatedAt := &latestCommit.Author.When

	// 提取去重用户
	emailSet := make(map[string]*User)
	for e := commits.Front(); e != nil; e = e.Next() {
		c := e.Value.(*git.Commit)
		if c.Author == nil {
			continue
		}

		email := c.Author.Email
		if _, exists := emailSet[email]; exists {
			continue
		}

		if user, _ := GetUserByEmail(email); user != nil {
			emailSet[email] = user
		}
	}

	users := make([]*User, 0, len(emailSet))
	for _, user := range emailSet {
		users = append(users, user)
	}

	if len(users) == 0 {
		return "", nil, createdAt, updatedAt, nil
	}

	// 使用 latestCommit 获取最新版本
	blob, err := latestCommit.GetBlobByPath(treePath)
	if err != nil {
		log.Error("Get blob failed: %v", err)
		return "", nil, nil, nil, err
	}

	reader, err := blob.DataAsync()
	if err != nil {
		return "", nil, nil, nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", nil, nil, nil, err
	}

	return string(data), users, createdAt, updatedAt, nil
}

// GetFileCommits 获取文件的真实提交历史
func GetFileCommits(gitRepo *git.Repository, treePath string, limit int) (*list.List, error) {
	cmd := git.NewCommand(
		"log",
		"--follow",           // 追踪重命名
		"--pretty=format:%H", // 只输出 SHA
		"-n", strconv.Itoa(limit),
		"--", treePath, // 指定文件路径
	)

	output, err := cmd.RunInDir(gitRepo.Path)
	if err != nil {
		return nil, err
	}

	commits := list.New()
	for _, line := range strings.Split(output, "\n") {
		sha := strings.TrimSpace(line)
		if sha == "" {
			continue
		}

		commit, err := gitRepo.GetCommit(sha)
		if err != nil {
			log.Warn("Failed to get commit %s: %v", sha, err)
			continue
		}
		commits.PushBack(commit)
	}

	return commits, nil
}
