// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"io"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/urchin_v2"
	"github.com/minio/minio-go"
)

const (
	MinioStorageType = "minio"
	LocalStorageType = "local"
)

// ObjectStorage represents an object storage to handle a bucket and files
type ObjectStorage interface {
	Save(path string, r io.Reader) (int64, error)
	Open(path string) (io.ReadCloser, error)
	DownloadAFile(bucket string, objectName string) (io.ReadCloser, error)
	Delete(path string) error
	DeleteDir(dir string) error
	PresignedGetURL(path string, fileName string) (string, error)
	PresignedPutURL(path string) (string, error)
	HasObject(path string) (bool, error)
	UploadObject(fileName, filePath string) error
	UploadContent(bucketName string, path string, r io.Reader) (int64, error)
	UploadContentWithSize(bucketName string, path string, size int64, r io.Reader) (int64, error)
}

// Copy copys a file from source ObjectStorage to dest ObjectStorage
func Copy(dstStorage ObjectStorage, dstPath string, srcStorage ObjectStorage, srcPath string) (int64, error) {
	f, err := srcStorage.Open(srcPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return dstStorage.Save(dstPath, f)
}

var (
	// Attachments represents attachments storage
	Attachments  ObjectStorage
	ObsCli       *obs.ObsClient
	MinioCore    *minio.Core
	UrchinClient *urchin_v2.UrchinV2Client
)

// Init init the stoarge
func Init() error {
	var err error
	switch setting.Attachment.StoreType {
	case LocalStorageType:
		Attachments, err = NewLocalStorage(setting.Attachment.Path)
		log.Info("local storage inited.")
	case MinioStorageType:
		m := setting.Attachment.Minio
		Attachments, err = NewMinioStorage(
			m.Endpoint,
			m.AccessKeyID,
			m.SecretAccessKey,
			m.Bucket,
			m.Location,
			m.BasePath,
			m.UseSSL,
		)
		if err != nil {
			log.Info("minio storage inited failed." + err.Error())
		} else {
			log.Info("minio storage inited succeed.")
		}
		MinioCore, err = minio.NewCore(m.Endpoint, m.AccessKeyID, m.SecretAccessKey, m.UseSSL)
		if err != nil {
			log.Error("init ScheduleMinioCore err.%v", err)
		}
	default:
		return fmt.Errorf("Unsupported attachment store type: %s", setting.Attachment.StoreType)
	}

	ObsCli, err = obs.New(setting.AccessKeyID, setting.SecretAccessKey, setting.Endpoint)
	if err != nil {
		log.Error("obs.New failed:", err)
		//return err
	}
	log.Info("obs cli inited.")
	UrchinClient = urchin_v2.NewUrchinClient(setting.UrchinAddress, setting.UrchinReqTimeoutSecond, setting.UrchinMaxConnection)
	log.Info("urchin client inited.")
	return nil
}

func SelectFileByPrefixAndSuffix(list []FileInfo, prefix, suffix string) []FileInfo {
	r := make([]FileInfo, 0)
	for _, l := range list {
		if l.IsDir {
			continue
		}
		if !strings.HasPrefix(l.FileName, prefix) {
			continue
		}
		if !strings.HasSuffix(l.FileName, suffix) {
			continue
		}
		r = append(r, l)
	}
	return r
}
