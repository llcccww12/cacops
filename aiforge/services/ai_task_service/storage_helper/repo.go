package storage_helper

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"github.com/unknwon/com"
)

func DownloadZipCode(gitRepo *git.Repository, codePath, branchName string) (commitId string, err error) {
	archiveType := git.ZIP
	archivePath := codePath

	if !com.IsDir(archivePath) {
		if err := os.MkdirAll(archivePath, os.ModePerm); err != nil {
			log.Info("MkdirAll failed,archivePath=" + archivePath)
			log.Error("MkdirAll failed:" + err.Error())
			return "", err
		}
	}

	// Get corresponding commit.
	var commit *git.Commit

	if err != nil {
		log.Error("OpenRepository failed:" + err.Error())
		return "", err
	}

	if gitRepo.IsBranchExist(branchName) {
		commit, err = gitRepo.GetBranchCommit(branchName)

		if err != nil {
			log.Error("GetBranchCommit failed:" + err.Error())
			return "", err
		}
	} else {
		log.Error("the branch is not exist: " + branchName)
		return "", fmt.Errorf("The branch does not exist.")
	}

	archivePath = path.Join(archivePath, grampus.CodeArchiveName)
	if !com.IsFile(archivePath) {
		if err := commit.CreateArchive(archivePath, git.CreateArchiveOpts{
			Format: archiveType,
			Prefix: setting.Repository.PrefixArchiveFiles,
		}); err != nil {
			log.Error("CreateArchive failed:" + err.Error())
			return "", err
		}
	}

	return commit.ID.String(), nil
}

func UploadDirToMinio(codePath, objectKeyPrefix, parentDir string) error {
	files, err := readDir(codePath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", codePath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err = UploadDirToMinio(codePath+file.Name()+"/", objectKeyPrefix, parentDir+file.Name()+"/"); err != nil {
				log.Error("uploadCodeToMinio(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			destObject := objectKeyPrefix + "/" + parentDir + file.Name()
			sourceFile := codePath + file.Name()
			err = storage.Attachments.UploadObject(destObject, sourceFile)
			if err != nil {
				log.Error("UploadObject(%s) failed: %s", file.Name(), err.Error())
				if strings.Contains(err.Error(), "no such file or directory") {
					continue
				}
				return err
			}
		}
	}

	return nil
}

func UploadDirToObs(codePath, objectKeyPrefix, parentDir string) error {
	files, err := readDir(codePath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", codePath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			input := &obs.PutObjectInput{}
			input.Bucket = setting.Bucket
			input.Key = parentDir + file.Name() + "/"
			_, err = storage.ObsCli.PutObject(input)
			if err != nil {
				log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
				return err
			}

			if err = UploadDirToObs(codePath+file.Name()+"/", objectKeyPrefix, parentDir+file.Name()+"/"); err != nil {
				log.Error("uploadCodeToObs(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			input := &obs.PutFileInput{}
			input.Bucket = setting.Bucket
			input.Key = objectKeyPrefix + "/" + parentDir + file.Name()
			input.SourceFile = codePath + file.Name()
			_, err = storage.ObsCli.PutFile(input)
			if err != nil {
				log.Error("PutFile(%s) failed: %s", input.SourceFile, err.Error())
				return err
			}
		}
	}

	return nil
}

// readDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(0)
	f.Close()
	if err != nil {
		//todo: can not upload empty folder
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	//sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

const CLONE_FILE_PREFIX = "file:///"

func DownloadCode(gitRepo *git.Repository, repo *models.Repository, codePath, branchName string) (string, error) {
	//add "file:///" prefix to make the depth valid
	codePath = path.Join(codePath, repo.Name)
	if err := git.Clone(CLONE_FILE_PREFIX+repo.RepoPath(), codePath, git.CloneRepoOptions{Branch: branchName, Depth: 1}); err != nil {
		log.Error("Failed to clone repository: %s (%v)", repo.FullName(), err)
		return "", err
	}

	configFile, err := os.OpenFile(codePath+"/.git/config", os.O_RDWR, 0666)
	if err != nil {
		log.Error("open file(%s) failed:%v", codePath+"/,git/config", err)
		return "", err
	}

	defer configFile.Close()

	pos := int64(0)
	reader := bufio.NewReader(configFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Error("not find the remote-url")
				return "", nil
			} else {
				log.Error("read error: %v", err)
				return "", err
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
				return "", err
			}
			break
		}

		pos += int64(len(line))
	}
	commitID, _ := gitRepo.GetBranchCommitID(branchName)

	return commitID, nil
}

func createZipArchive(sourceDir, archivePath string) error {
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	zipWriter := zip.NewWriter(archiveFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		// 排除压缩包本身
		if strings.TrimSuffix(relPath, "/") == filepath.Base(archivePath) {
			return nil
		}

		if info.IsDir() {
			// 创建目录项
			zipEntry, err := zipWriter.Create(filepath.ToSlash(relPath) + "/") // 目录项以斜杠结尾
			if err != nil {
				return err
			}
			_ = zipEntry // 忽略目录项

			return nil
		}

		zipEntry, err := zipWriter.Create(filepath.ToSlash(relPath)) // 使用正斜杠分隔符
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipEntry, file)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
