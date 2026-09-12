package repo

import (
	"code.gitea.io/gitea/modules/httplib"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

const (
	tplDirIndex base.TplName = "repo/datasets/dirs/index"
)

type RespGetDirs struct {
	ResultCode string `json:"resultCode"`
	FileInfos  string `json:"fileInfos"`
}

func DeleteAllUnzipFile(attachment *models.Attachment, parentDir string) {

	if !isCanDecompress(attachment.Name) {
		log.Error("The file is not zip file, can not query the dir")
		return
	} else if attachment.DecompressState != models.DecompressStateDone {
		log.Error("The file has not been decompressed completely now")
		return
	}
	if len(attachment.UnzipPath) > 0 {
		log.Info("start to delete Decompress path=" + attachment.UnzipPath)
		if attachment.Type == models.StorageTypeMinio {
			storage.Attachments.DeleteDir(attachment.UnzipPath)
		}
		if attachment.Type == models.StorageTypeObs {
			err := storage.ObsRemoveObject(setting.Bucket, attachment.UnzipPath)
			if err != nil {
				log.Info("delete file error.")
			}
		}
	}
}

func DirIndex(ctx *context.Context) {
	uuid := ctx.Params("uuid")
	parentDir := ctx.Query("parentDir")
	dirArray := strings.Split(parentDir, "/")

	attachment, err := models.GetAttachmentByUUID(uuid)
	if err != nil {
		ctx.ServerError("GetDatasetAttachments", err)
		return
	}

	//if !strings.HasSuffix(attachment.Name, ".zip") {
	if !isCanDecompress(attachment.Name) {
		log.Error("The file is not zip file, can not query the dir")
		ctx.ServerError("The file is not zip file, can not query the dir", errors.New("The file is not zip file, can not query the dir"))
		return
	} else if attachment.DecompressState != models.DecompressStateDone {
		log.Error("The file has not been decompressed completely now")
		ctx.ServerError("The file has not been decompressed completely now", errors.New("The file has not been decompressed completely now"))
		return
	}

	dirArray = append([]string{attachment.Name}, dirArray...)
	if parentDir == "" {
		dirArray = []string{attachment.Name}
	}

	/*
		dirs, err := GetDatasetDirs(uuid, parentDir)
		if err != nil {
			log.Error("getDatasetDirs failed:", err.Error())
			ctx.ServerError("getDatasetDirs failed:", err)
			return
		}
	*/

	//var fileInfos []FileInfo

	/*
		err = json.Unmarshal([]byte(dirs), &fileInfos)
		if err != nil {
			log.Error("json.Unmarshal failed:", err.Error())
			ctx.ServerError("json.Unmarshal failed:", err)
			return
		}
	*/

	ctx.Data["Path"] = dirArray
	ctx.Data["Dirs"] = true
	ctx.Data["Uuid"] = uuid
	ctx.Data["Type"] = attachment.Type
	ctx.Data["PageIsDataset"] = true

	ctx.HTML(200, tplDirIndex)
}

func GetDatasetDirs(uuid string, parentDir string) (string, error) {
	var req string
	dataActualPath := setting.Attachment.Minio.RealPath +
		setting.Attachment.Minio.Bucket + "/" +
		setting.Attachment.Minio.BasePath +
		models.AttachmentRelativePath(uuid) +
		uuid + "/"
	if parentDir == "" {
		req = "baseDir=" + dataActualPath
	} else {
		req = "baseDir=" + dataActualPath + "&parentDir=" + parentDir
	}

	return getDirs(req)
}

func getDirs(req string) (string, error) {
	var dirs string

	url := setting.DecompressAddress + "/dirs?" + req
	reqHttp, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error("http.NewRequest failed:", err.Error())
		return dirs, err
	}

	reqHttp.SetBasicAuth(setting.AuthUser, setting.AuthPassword)

	client := httplib.NewClientTimeOut(&httplib.Config{})
	res, err := client.Do(reqHttp)
	if err != nil {
		log.Error("send http to decompress failed:", err.Error())
		return dirs, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error("the response from decompress is failed")
		return dirs, errors.New("the response from decompress is failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("read resp body failed:", err.Error())
		return dirs, err
	}

	var resp RespGetDirs
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Error("unmarshal resp failed:", err.Error())
		return dirs, err
	}

	if resp.ResultCode != "0" {
		log.Error("GetDirs failed:", resp.ResultCode)
		return dirs, errors.New("GetDirs failed")
	}

	dirs = resp.FileInfos
	return dirs, nil
}
