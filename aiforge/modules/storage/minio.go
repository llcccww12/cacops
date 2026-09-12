// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package storage

import (
	"io"
	"net/url"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	"github.com/minio/minio-go"
)

var (
	_ ObjectStorage = &MinioStorage{}
)

const (
	PresignedGetUrlExpireTime = time.Hour * 24 * 7
	PresignedPutUrlExpireTime = time.Hour * 24 * 7
)

// MinioStorage returns a minio bucket storage
type MinioStorage struct {
	client   *minio.Client
	bucket   string
	basePath string
}

// NewMinioStorage returns a minio storage
func NewMinioStorage(endpoint, accessKeyID, secretAccessKey, bucket, location, basePath string, useSSL bool) (*MinioStorage, error) {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return nil, err
	}

	if err := minioClient.MakeBucket(bucket, location); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucket)
		if !exists || errBucketExists != nil {
			return nil, err
		}
	}

	return &MinioStorage{
		client:   minioClient,
		bucket:   bucket,
		basePath: basePath,
	}, nil
}

func (m *MinioStorage) buildMinioPath(p string) string {
	return strings.TrimPrefix(path.Join(m.basePath, p), "/")
}

func (m *MinioStorage) DownloadAFile(bucket string, objectName string) (io.ReadCloser, error) {

	var opts = minio.GetObjectOptions{}
	object, err := m.client.GetObject(m.bucket, objectName, opts)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// Open open a file
func (m *MinioStorage) Open(path string) (io.ReadCloser, error) {
	var opts = minio.GetObjectOptions{}
	object, err := m.client.GetObject(m.bucket, m.buildMinioPath(path), opts)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// Save save a file to minio
func (m *MinioStorage) Save(path string, r io.Reader) (int64, error) {
	return m.client.PutObject(m.bucket, m.buildMinioPath(path), r, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
}

// Delete delete a file
func (m *MinioStorage) Delete(path string) error {
	return m.client.RemoveObject(m.bucket, m.buildMinioPath(path))
}

// Delete delete a file
func (m *MinioStorage) DeleteDir(dir string) error {
	objectsCh := make(chan string)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		for object := range m.client.ListObjects(m.bucket, dir, true, nil) {
			if object.Err != nil {
				log.Error("ListObjects failed:%v", object.Err)
			}
			objectsCh <- object.Key
		}
	}()

	m.client.RemoveObjects(m.bucket, objectsCh)

	return nil
}

//Get Presigned URL for get object
func (m *MinioStorage) PresignedGetURL(path string, fileName string) (string, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

	var preURL *url.URL
	preURL, err := m.client.PresignedGetObject(m.bucket, path, PresignedGetUrlExpireTime, reqParams)
	if err != nil {
		return "", err
	}

	return preURL.String(), nil
}

//Get Presigned URL for put object
func (m *MinioStorage) PresignedPutURL(path string) (string, error) {
	var preURL *url.URL
	preURL, err := m.client.PresignedPutObject(m.bucket, m.buildMinioPath(path), PresignedPutUrlExpireTime)
	if err != nil {
		return "", err
	}

	return preURL.String(), nil
}

//check if has the object
func (m *MinioStorage) HasObject(path string) (bool, error) {
	hasObject := false
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)
	//objectCh := m.client.ListObjects(m.bucket, m.buildMinioPath(path), false, doneCh)
	objectCh := m.client.ListObjects(m.bucket, path, false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return hasObject, object.Err
		}

		hasObject = true
	}

	return hasObject, nil
}

//upload object
func (m *MinioStorage) UploadObject(fileName, filePath string) error {
	_, err := m.client.FPutObject(m.bucket, fileName, filePath, minio.PutObjectOptions{})
	return err
}

func (m *MinioStorage) UploadContent(bucketName string, path string, r io.Reader) (int64, error) {
	return m.client.PutObject(bucketName, path, r, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
}

func (m *MinioStorage) UploadContentWithSize(bucketName string, path string, size int64, r io.Reader) (int64, error) {
	return m.client.PutObject(bucketName, path, r, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
}

func GetMinioPath(jobName, suffixPath string) string {
	return setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + setting.CBCodePathPrefix + jobName + suffixPath
}
