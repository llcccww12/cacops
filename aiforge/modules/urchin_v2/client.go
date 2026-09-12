package urchin_v2

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"code.gitea.io/gitea/modules/log"
	urchin_client "code.gitea.io/gitea/modules/urchin_v2/client"
	urchin_module "code.gitea.io/gitea/modules/urchin_v2/module"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type UrchinV2Client struct {
	UrchinClient urchin_client.UrchinClient
	Address      string
	ObsClient    *obs.ObsClient
}

func NewUrchinClient(address string,
	reqTimeout int64,
	maxConnection int) *UrchinV2Client {

	var obsClientSocketTimeout,
		obsClientMaxConnection,
		obsClientMaxRetryCount = 3600, 100, 3
	obsClient, _ := obs.New(
		"",
		"",
		"magicalParam",
		obs.WithSignature(obs.SignatureObs),
		obs.WithSocketTimeout(obsClientSocketTimeout),
		obs.WithMaxConnections(obsClientMaxConnection),
		obs.WithMaxRetryCount(obsClientMaxRetryCount))

	var urchinClient urchin_client.UrchinClient

	urchinClient.Init(context.Background(), address, reqTimeout, maxConnection)
	return &UrchinV2Client{
		UrchinClient: urchinClient,
		Address:      address,
		ObsClient:    obsClient,
	}
}

func (client *UrchinV2Client) CreateCollection(name string) (string, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", name)

	createObjectReq := new(urchin_module.CreateObjectReq)
	createObjectReq.Name = name
	err, createObjectResp := client.UrchinClient.CreateObject(ctx, createObjectReq)
	if nil != err {
		return "", err
	}
	if !createObjectResp.IsSuccess() {
		return "", createObjectResp.ToError()
	}
	return createObjectResp.ObjUuid, nil
}

func (client *UrchinV2Client) CreatePutObjectSignedUrl(objUuid string, relativePath string) (string, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objUuid)
	req := &urchin_module.CreatePutObjectSignedUrlReqWithoutTask{
		ObjUuid: objUuid,
		Source:  relativePath,
	}
	err, resp := client.UrchinClient.CreatePutObjectSignedUrlWithoutTask(
		ctx,
		req)
	if nil != err {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", resp.ToError()
	}

	return resp.SignedUrl, nil

}

func (client *UrchinV2Client) UploadLocalDir(objUuid string, localPath, targetRelativePath string) error {
	targetRelativePath = strings.Trim(targetRelativePath, "/")
	files, err := readDir(localPath)
	if err != nil {
		log.Error("readDir(%s) failed: %s", localPath, err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			dirRelPath := path.Join(targetRelativePath, file.Name())
			err = client.MKDIR(objUuid, dirRelPath+"/")
			// _, err = storage.ObsCli.PutObject(input)
			if err != nil {
				log.Error("PutObject(%s) failed: %s", dirRelPath+"/", err.Error())
				return err
			}

			if err = client.UploadLocalDir(objUuid, localPath+file.Name()+"/", dirRelPath); err != nil {
				log.Error("UploadLocalDir(%s) failed: %s", file.Name(), err.Error())
				return err
			}
		} else {
			filePath := localPath + file.Name()
			err = client.UploadLocalFile(objUuid, filePath, path.Join(targetRelativePath, file.Name()))
			if err != nil {
				log.Error("PutFile(%s) failed: %s", filePath, err.Error())
				return err
			}
		}
	}
	return nil

}

func (client *UrchinV2Client) UploadLocalFile(objUuid string, localPath, targetRelativePath string) error {
	abs, err := filepath.Abs(filepath.Clean(localPath))
	if err != nil {
		return err
	}

	fi, err := os.Stat(abs)
	if err != nil {
		return err
	}
	if !fi.Mode().IsRegular() {
		log.Error("[%s]%s not a regular file", objUuid, localPath)
		return nil
	}

	f, err := os.Open(abs)
	if err != nil {
		return err
	}
	defer f.Close()

	err = client.UploadFile(objUuid, targetRelativePath, f)
	if err != nil {
		return err
	}
	return nil

}

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

func (client *UrchinV2Client) UploadFile(objUuid string, targetRelativePath string, r io.Reader) error {
	url, err := client.CreatePutObjectSignedUrl(objUuid, targetRelativePath)
	if err != nil {
		return err
	}

	_, err = client.ObsClient.PutObjectWithSignedUrl(
		url,
		http.Header{},
		r)
	if nil != err {
		return err
	}
	return nil
}

func (client *UrchinV2Client) MKDIR(objUuid, path string) error {
	url, err := client.CreatePutObjectSignedUrl(objUuid, path)
	if err != nil {
		return err
	}

	_, err = client.ObsClient.PutObjectWithSignedUrl(
		url,
		http.Header{},
		nil)
	if nil != err {
		return err
	}
	return nil
}

func (client *UrchinV2Client) ReportSizeChanged(objUuid string) error {
	ctx := context.Background()
	err, _ := client.UrchinClient.ReportSizeChanged(ctx, &urchin_module.ReportSizeChangedReq{
		ObjUuid: objUuid,
	})
	return err
}

func (client *UrchinV2Client) RelateObject(name string, desc string, nodeName string, location string, userId string) (error, string) {
	ctx := context.Background()
	err, res := client.UrchinClient.RelateObject(ctx, &urchin_module.RelateObjectReq{
		Name:     name,
		Desc:     GetStringPointer(desc),
		NodeName: nodeName,
		Location: location,
		UserId:   userId,
	})
	if nil != err {
		return err, ""
	}
	if !res.IsSuccess() {
		if res.Code == urchin_module.ErrCodeObjectAlreadyExist && res.ObjUuid != "" {
			return nil, res.ObjUuid
		}
		return res.ToError(), ""
	}
	return nil, res.ObjUuid
}

func (client *UrchinV2Client) DeleteObjectDeployment(objUuid string) error {
	ctx := context.Background()
	var force = new(bool)
	*force = true
	err, resp := client.UrchinClient.DeleteObjectDeployment(ctx, &urchin_module.DeleteObjectDeploymentReq{
		ObjUuid: objUuid,
		Force:   force,
	})
	if !resp.IsSuccess() {
		return resp.ToError()
	}
	return err
}

func (client *UrchinV2Client) UpdateObjectLastModifyTime(objUuid string) error {
	ctx := context.Background()
	err, resp := client.UrchinClient.UpdateObjectLastModifyTime(ctx, &urchin_module.UpdateObjectLastModifyTimeReq{
		ObjUuid: objUuid,
	})
	if !resp.IsSuccess() {
		return resp.ToError()
	}
	return err
}

func (client *UrchinV2Client) ListPrefixObjectsWithMarkerAndDelimeter(objuuid, prefix, marker string, maxKey int32) (*urchin_module.ListObjectsResp, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)

	var delimeter = "/"
	req := &urchin_module.ListObjectsReq{
		ObjUuid:   objuuid,
		Prefix:    GetStringPointer(prefix),
		Marker:    GetStringPointer(marker),
		MaxKeys:   &maxKey,
		Delimiter: &delimeter,
	}
	err, resp := client.UrchinClient.ListObjects(ctx, req)
	if nil != err {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, resp.ToError()
	}
	return resp, err
}

func GetStringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (client *UrchinV2Client) GetObject(objuuid, path string) (*obs.GetObjectOutput, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)

	err, objectResp := client.UrchinClient.CreateGetObjectSignedUrlWithoutTask(ctx, &urchin_module.CreateGetObjectSignedUrlReqWithoutTask{
		ObjUuid: objuuid,
		Source:  path,
	})
	if err != nil {
		return nil, err
	}
	if !objectResp.IsSuccess() {
		return nil, objectResp.ToError()
	}
	output, err := client.ObsClient.GetObjectWithSignedUrl(
		objectResp.SignedUrl,
		http.Header{})
	if nil != err {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return output, nil
}

func (client *UrchinV2Client) GetObjectMeta(objuuid, path string) (*obs.GetObjectMetadataOutput, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	err, objectResp := client.UrchinClient.CreateGetObjectMetadataSignedUrlWithoutTask(ctx, &urchin_module.CreateGetObjectMetadataSignedUrlReqWithoutTask{
		ObjUuid: objuuid,
		Source:  path,
	})
	if err != nil {
		return nil, err
	}
	if !objectResp.IsSuccess() {
		return nil, objectResp.ToError()
	}
	output, err := client.ObsClient.GetObjectMetadataWithSignedUrl(
		objectResp.SignedUrl,
		http.Header{})

	if output == nil {
		return nil, fmt.Errorf("GetObjectMeta(%s:%s) failed", objuuid, path)
	}
	return output, nil
}

const (
	TASKSUCCESS int32 = 1
	TASKFAILED  int32 = 2
)

func (client *UrchinV2Client) FinishSuccessTask(taskId int32) error {
	ctx := context.Background()
	req := &urchin_module.FinishTaskReq{
		TaskId: taskId,
		Result: TASKSUCCESS,
	}
	err, resp := client.UrchinClient.FinishTask(ctx, req)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return resp.ToError()
	}
	return nil
}

func (client *UrchinV2Client) FinishFailedTask(taskId int32) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", fmt.Sprint(taskId))

	req := &urchin_module.FinishTaskReq{
		TaskId: taskId,
		Result: TASKFAILED,
	}
	err, resp := client.UrchinClient.FinishTask(ctx, req)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return resp.ToError()
	}
	return nil
}

func (client *UrchinV2Client) GetObjectSignedUrl(objuuid, path string, queryParams map[string]string) (string, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)

	err, objectResp := client.UrchinClient.CreateGetObjectSignedUrlWithoutTask(ctx, &urchin_module.CreateGetObjectSignedUrlReqWithoutTask{
		ObjUuid:     objuuid,
		Source:      path,
		QueryParams: queryParams,
	})
	if err != nil {
		return "", err
	}
	if !objectResp.IsSuccess() {
		return "", objectResp.ToError()
	}
	return objectResp.SignedUrl, nil
}

func (client *UrchinV2Client) ListPrefixObjectsWithMarker(objuuid string, marker string, prefix string, maxKey int32) (*urchin_module.ListObjectsResp, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)

	req := &urchin_module.ListObjectsReq{
		ObjUuid: objuuid,
		Prefix:  GetStringPointer(prefix),
		Marker:  GetStringPointer(marker),
		MaxKeys: &maxKey,
	}
	err, resp := client.UrchinClient.ListObjects(ctx, req)
	if nil != err {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, resp.ToError()
	}
	return resp, err
}

func (client *UrchinV2Client) CopyObject(sourceObjUuid, sourcePath, targetObjUuid, targetPath string) error {
	////先下载再上传，海胆后续版本会支持直接copy

	res, err := client.GetObject(sourceObjUuid, sourcePath)
	if err != nil {
		return err
	}
	return client.UploadFile(targetObjUuid, targetPath, res.Body)
}

func (client *UrchinV2Client) DeleteFile(objuuid, sourcePath string) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	req := &urchin_module.DeleteFileReq{
		ObjUuid: objuuid,
		Source:  sourcePath,
	}
	err, resp := client.UrchinClient.DeleteFile(ctx, req)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return resp.ToError()
	}

	return nil
}

func (client *UrchinV2Client) DeleteObject(objuuid string) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	req := &urchin_module.DeleteObjectReq{
		ObjUuid: objuuid,
	}
	err, resp := client.UrchinClient.DeleteObject(ctx, req)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return resp.ToError()
	}

	return nil
}

func dirOf(name string) string {
	c := filepath.Clean(name)
	if !strings.ContainsRune(c, filepath.Separator) {
		return "/"
	}
	return filepath.Dir(c)
}

func (client *UrchinV2Client) InitiateMultipartUpload(objuuid, path string) (string, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	err, objectResp := client.UrchinClient.CreateInitiateMultipartUploadSignedUrlWithoutTask(ctx, &urchin_module.CreateInitiateMultipartUploadSignedUrlReqWithoutTask{
		ObjUuid: objuuid,
		Source:  path,
	})
	if err != nil {
		return "", err
	}
	if !objectResp.IsSuccess() {
		return "", objectResp.ToError()
	}
	output, err := client.ObsClient.InitiateMultipartUploadWithSignedUrl(
		objectResp.SignedUrl,
		http.Header{})

	if err != nil {
		return "", err
	}
	// client.UrchinClient.FinishTask()
	return output.UploadId, nil
}

func (client *UrchinV2Client) CreateMultipartUploadSignedUrl(objuuid, uploadId, source string, partNumber int) (string, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	err, output := client.UrchinClient.CreateUploadPartSignedUrlWithoutTask(ctx, &urchin_module.CreateUploadPartSignedUrlReqWithoutTask{
		ObjUuid:    objuuid,
		Source:     source,
		UploadId:   uploadId,
		PartNumber: int32(partNumber),
	})
	if err != nil {
		return "", err
	}
	if !output.IsSuccess() {
		return "", output.ToError()
	}
	return output.SignedUrl, nil
}

func (client *UrchinV2Client) CompleteMultiPartUpload(objuuid, uploadId, source string, totalChunks int) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)
	err, urlOutput := client.UrchinClient.CreateCompleteMultipartUploadSignedUrlWithoutTask(ctx, &urchin_module.CreateCompleteMultipartUploadSignedUrlReqWithoutTask{
		ObjUuid:  objuuid,
		Source:   source,
		UploadId: uploadId,
	})
	if err != nil {
		return err
	}
	if !urlOutput.IsSuccess() {
		return urlOutput.ToError()
	}

	allParts, err := client.ListAllParts(objuuid, uploadId, source)
	if err != nil {
		return err
	}

	if len(allParts.Parts) != totalChunks {
		log.Error("listAllParts number(%d) is not equal the set total chunk number(%d)", len(allParts.Parts), totalChunks)
		return errors.New("the parts is not complete")
	}

	type completePart struct {
		PartNumber int    `xml:"PartNumber"`
		ETag       string `xml:"ETag"`
	}

	var parts []completePart
	for _, p := range allParts.Parts {
		parts = append(parts, completePart{
			PartNumber: p.PartNumber,
			ETag:       p.ETag,
		})
	}

	xmlBody, _ := xml.MarshalIndent(struct {
		XMLName xml.Name       `xml:"CompleteMultipartUpload"`
		Parts   []completePart `xml:"Part"`
	}{
		Parts: parts,
	}, "", "  ")
	data := bytes.NewReader(xmlBody)
	output, _ := client.ObsClient.CompleteMultipartUploadWithSignedUrl(
		urlOutput.SignedUrl,
		http.Header{}, data)

	if output == nil {
		return fmt.Errorf("CompleteMultipartUploadWithSignedUrl output empty")
	}

	return nil
}

func (client *UrchinV2Client) ListAllParts(objuuid, uploadID, path string) (output *obs.ListPartsOutput, err error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-Id", objuuid)

	output = &obs.ListPartsOutput{}
	var partNumberMarker int32 = 0
	var maxListParts int32 = 1000

	for {

		err, objectResp := client.UrchinClient.CreateListPartsSignedUrlWithoutTask(ctx, &urchin_module.CreateListPartsSignedUrlReqWithoutTask{
			ObjUuid:          objuuid,
			Source:           path,
			UploadId:         uploadID,
			PartNumberMarker: &partNumberMarker,
			MaxParts:         &maxListParts,
		})
		if err != nil {
			return nil, err
		}
		if !objectResp.IsSuccess() {
			return nil, objectResp.ToError()
		}
		temp, err := client.ObsClient.ListPartsWithSignedUrl(
			objectResp.SignedUrl,
			http.Header{})

		if err != nil {
			return nil, err
		}
		partNumberMarker = int32(temp.NextPartNumberMarker)
		for _, partInfo := range temp.Parts {
			output.Parts = append(output.Parts, obs.Part{
				PartNumber: partInfo.PartNumber,
				ETag:       partInfo.ETag,
			})
		}

		if !temp.IsTruncated {
			break
		} else {
			continue
		}
	}

	return output, nil
}
