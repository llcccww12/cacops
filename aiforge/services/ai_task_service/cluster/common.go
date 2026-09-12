package cluster

import (
	"archive/zip"
	"bufio"
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"io"
	"strings"
)

func GetLocalLog(r io.Reader, startLine, endLine int64) (content string, realEndLine int64, total int64) {
	if startLine > endLine {
		return "", 0, 0
	}
	re := ""
	fileEndLine := endLine
	reader := bufio.NewReader(r)
	var countLine = int64(1)
	//跳过开始行之前的内容
	for countLine < startLine {
		_, err := reader.ReadString('\n')
		if err != nil {
			log.Error("GetLocalLog ReadString err. %v", err)
			return "", 0, 0
		}
		countLine++
	}
	//读取指定的开始行到结束行的内容
	for countLine >= startLine && countLine <= endLine {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				re = re + line
				countLine++
			}
			log.Error("GetLocalLog ReadString err. %v", err)
			break
		}
		re = re + line
		countLine++
	}
	fileEndLine = countLine - 1

	return re, fileEndLine, fileEndLine - startLine + 1
}

func getAllLineFromFile(helper storage_helper.StorageHelper, filePath string) int64 {
	var count int64
	r, err := helper.OpenFile(filePath)
	defer r.Close()
	if err != nil {
		log.Info("error:" + err.Error())
		return 0
	}

	reader := bufio.NewReader(r)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				count++
			}
			log.Error("GetLocalLog ReadString err. %v", err)
			break
		}
		count++
	}

	return count
}

func getLogFilesInStorage(helper storage_helper.StorageHelper, objectKeyPrefix string, logSuffix string) []storage.FileInfo {
	//获取日志输出目录下文件列表

	fileList, err := helper.GetOneLevelObjectsUnderDir(objectKeyPrefix)
	if err != nil {
		log.Error("GetTrainLog read dir err.objectKeyPrefix=%s,err=%v", objectKeyPrefix, err)
		return nil
	}
	if len(fileList) == 0 {
		return nil
	}

	logFiles := make([]storage.FileInfo, 0)
	for _, f := range fileList {
		if f.IsDir {
			continue
		}
		if strings.HasSuffix(f.FileName, logSuffix) {
			logFiles = append(logFiles, f)
		}
	}
	return logFiles
}

func DownloadAllOutput(opts entity.DownloadOutputOpts) error {
	helper := storage_helper.SelectStorageHelperFromStorageType(opts.StorageType)
	var err error
	fileList, err := helper.GetAllObjectsUnderDir(opts.Path)
	if err != nil {
		log.Error("GetAllObjectsUnderDir err.objectKeyPrefix=%s,err=%v", opts.Path, err)
		return err
	}
	if len(fileList) == 0 {
		return nil
	}

	for i := 0; i < len(fileList); i++ {
		file := fileList[i]
		if file.IsDir {
			continue
		}
		err = openAndWrite2ZIP(helper, file, opts.ZIPWriter)
		if err != nil {
			log.Error("openAndWrite2ZIP err.%v", err)
			return err
		}
	}

	return nil
}

func openAndWrite2ZIP(helper storage_helper.StorageHelper, file storage.FileInfo, zipWriter *zip.Writer) error {
	var reader io.ReadCloser
	reader, err := helper.OpenFile(file.RelativePath)
	if err != nil {
		log.Error("openAndWrite2ZIP OpenFile err.filePath=%+v,err =%v", file.RelativePath, err)
		return err
	}
	defer reader.Close()

	fDest, err := zipWriter.Create(file.FileName)
	if err != nil {
		log.Error("zipWriter.Create error.%v", err)
		return err
	}
	p := make([]byte, 1024)
	var readErr error
	var readCount int
	// 读取对象内容
	for {
		readCount, readErr = reader.Read(p)
		if readCount > 0 {
			fDest.Write(p[:readCount])
		}
		if readErr != nil {
			break
		}
	}
	return nil
}
