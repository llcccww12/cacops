package container_builder

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

func DownloadCode(ctx *context.CreationContext, codeLocalPath string, uncompressed bool) error {
	dir, err := ioutil.ReadDir(codeLocalPath)
	log.Info("codeLocalPath=" + codeLocalPath)
	//ReqCommitID为空时需要下载最新的代码，把旧的删掉
	if len(dir) != 0 && ctx.Request.ReqCommitID == "" {
		if err == nil {
			os.RemoveAll(codeLocalPath)
		}
	}
	var commitId string
	//目录为空时需要下载代码
	if len(dir) == 0 {
		if uncompressed {
			commitId, err = storage_helper.DownloadCode(ctx.GitRepo, ctx.Repository, codeLocalPath, ctx.Request.BranchName)
		} else {
			commitId, err = storage_helper.DownloadZipCode(ctx.GitRepo, codeLocalPath, ctx.Request.BranchName)
		}
		if err != nil {
			log.Error("downloadZipCode failed, server timed out: %s (%v)", ctx.Repository.FullName(), err)
			return errors.New("cloudbrain.load_code_failed")
		}
	}
	ctx.CommitID = commitId
	return nil
}

var obsUploader = &storage_helper.OBSHelper{}
var minioUploader = &storage_helper.MinioHelper{}

const CLONE_FILE_PREFIX = "file:///"

func DownloadBranch(repo *models.Repository, codePath, branchName string) error {
	//add "file:///" prefix to make the depth valid
	if err := git.Clone(CLONE_FILE_PREFIX+repo.RepoPath(), codePath, git.CloneRepoOptions{Branch: branchName, Depth: 1}); err != nil {
		log.Error("Failed to clone repository: %s (%v)", repo.FullName(), err)
		return err
	}

	configFile, err := os.OpenFile(codePath+"/.git/config", os.O_RDWR, 0666)
	if err != nil {
		log.Error("open file(%s) failed:%v", codePath+"/,git/config", err)
		return err
	}

	defer configFile.Close()

	pos := int64(0)
	reader := bufio.NewReader(configFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Error("not find the remote-url")
				return nil
			} else {
				log.Error("read error: %v", err)
				return err
			}
		}

		if strings.Contains(line, "url") && strings.Contains(line, ".git") {
			originUrl := "\turl = " + repo.CloneLink().HTTPS + "\n"
			if len(line) > len(originUrl) {
				originUrl += strings.Repeat(" ", len(line)-len(originUrl))
			}
			bytes := []byte(originUrl)
			_, err := configFile.WriteAt(bytes, pos)
			if err != nil {
				log.Error("WriteAt failed:%v", err)
				return err
			}
			break
		}

		pos += int64(len(line))
	}

	return nil
}
