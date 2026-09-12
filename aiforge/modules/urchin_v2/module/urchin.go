package module

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	DefaultUClientReqTimeout    = 10
	DefaultUClientMaxConnection = 500

	UrchinClientHeaderUserId  = "X-User-Id"
	UrchinClientHeaderToken   = "X-Token"
	UrchinClientHeaderquestId = "X-Request-Id"

	UrchinClientCreateInitiateMultipartUploadSignedUrlInterface            = "/v1/object/auth/create_init_multi_part_upload_signed_url"
	UrchinClientCreateUploadPartSignedUrlInterface                         = "/v1/object/auth/create_upload_part_signed_url"
	UrchinClientCreateListPartsSignedUrlInterface                          = "/v1/object/auth/create_list_parts_signed_url"
	UrchinClientCreateCompleteMultipartUploadSignedUrlInterface            = "/v1/object/auth/create_complete_multi_part_upload_signed_url"
	UrchinClientCreateInitiateMultipartUploadSignedUrlInterfaceWithoutTask = "/v1/object/auth/create_init_multi_part_upload_signed_url_without_task"
	UrchinClientCreateUploadPartSignedUrlInterfaceWithoutTask              = "/v1/object/auth/create_upload_part_signed_url_without_task"
	UrchinClientCreateListPartsSignedUrlInterfaceWithoutTask               = "/v1/object/auth/create_list_parts_signed_url_without_task"
	UrchinClientCreateCompleteMultipartUploadSignedUrlInterfaceWithoutTask = "/v1/object/auth/create_complete_multi_part_upload_signed_url_without_task"
	UrchinClientCreateAbortMultipartUploadSignedUrlInterface               = "/v1/object/auth/create_abort_multi_part_upload_signed_url"
	UrchinClientCreatePutObjectSignedUrlInterface                          = "/v1/object/auth/create_put_object_signed_url"
	UrchinClientCreatePutObjectSignedUrlInterfaceWithoutTask               = "/v1/object/auth/create_put_object_signed_url_without_task"
	UrchinClientCreateGetObjectMetadataSignedUrlInterface                  = "/v1/object/auth/create_get_object_metadata_signed_url"
	UrchinClientCreateGetObjectMetadataSignedUrlInterfaceWithoutTask       = "/v1/object/auth/create_get_object_metadata_signed_url_without_task"
	UrchinClientCreateGetObjectSignedUrlInterface                          = "/v1/object/auth/create_get_object_signed_url"
	UrchinClientCreateGetObjectSignedUrlInterfaceWithoutTask               = "/v1/object/auth/create_get_object_signed_url_without_task"
	UrchinClientCreateListObjectsSignedUrlInterface                        = "/v1/object/auth/create_list_objects_signed_url"
	UrchinClientCreateListObjectsSignedUrlInterfaceWithoutTask             = "/v1/object/auth/create_list_objects_signed_url_without_task"
	UrchinClientGetIpfsTokenInterface                                      = "/v1/object/auth/get_ipfs_token"
	UrchinClientCreateJCSPreSignedObjectListInterface                      = "/v1/object/auth/create_jcs_pre_signed_list"
	UrchinClientCreateJCSPreSignedObjectUploadInterface                    = "/v1/object/auth/create_jcs_pre_signed_object_upload"
	UrchinClientCreateJCSPreSignedObjectNewMultipartUploadInterface        = "/v1/object/auth/create_jcs_pre_signed_new_multi_part_upload"
	UrchinClientCreateJCSPreSignedObjectUploadPartInterface                = "/v1/object/auth/create_jcs_pre_signed_upload_part"
	UrchinClientCreateJCSPreSignedObjectCompleteMultipartUploadInterface   = "/v1/object/auth/create_jcs_pre_signed_complete_multi_part_upload"
	UrchinClientCreateJCSPreSignedObjectDownloadInterface                  = "/v1/object/auth/create_jcs_pre_signed_download"

	UrchinClientCreateObjectInterface               = "/v1/object/create"
	UrchinClientRelateObjectInterface               = "/v1/object/relate"
	UrchinClientUpdateObjectLastModifyTimeInterface = "/v1/object/update_last_modify_time"
	UrchinClientUploadObjectInterface               = "/v1/object/upload"
	UrchinClientDownloadObjectInterface             = "/v1/object/download"
	UrchinClientLoadObjectInterface                 = "/v1/object/load"
	UrchinClientMigrateObjectInterface              = "/v1/object/migrate"
	UrchinClientCopyObjectInterface                 = "/v1/object/copy"
	UrchinClientGetObjectInterface                  = "/v1/object"
	UrchinClientDeleteObjectInterface               = "/v1/object"
	UrchinClientPutObjectDeploymentInterface        = "/v1/object/deployment"
	UrchinClientDeleteObjectDeploymentInterface     = "/v1/object/deployment"
	UrchinClientListObjectsInterface                = "/v1/object/list"
	UrchinClientGetObjectMetadataInterface          = "/v1/object/metadata"
	UrchinClientListPartsInterface                  = "/v1/object/part/list"

	UrchinClientUploadFileInterface        = "/v1/object/file/upload"
	UrchinClientDownloadFileInterface      = "/v1/object/file/download"
	UrchinClientDeleteFileInterface        = "/v1/object/file"
	UrchinClientReportSizeChangedInterface = "/v1/object/count"

	UrchinClientGetTaskInterface             = "/v1/task"
	UrchinClientFinishTaskInterface          = "/v1/task/finish"
	UrchinClientRetryTaskInterface           = "/v1/task/retry"
	UrchinClientReportTaskHeartbeatInterface = "/v1/task/report_heartbeat"

	StorageCategoryEIpfs      = 1
	StorageCategoryEObs       = 2
	StorageCategoryEMinio     = 3
	StorageCategoryEJcs       = 4
	StorageCategoryEEos       = 5
	StorageCategoryEStarLight = 6
	StorageCategoryEParaCloud = 7
	StorageCategoryEScow      = 8
	StorageCategoryESugon     = 9

	TaskTypeUpload       = 1
	TaskTypeDownload     = 2
	TaskTypeMigrate      = 3
	TaskTypeCopy         = 4
	TaskTypeUploadFile   = 5
	TaskTypeDownloadFile = 6
	TaskTypeLoad         = 7

	TaskFResultESuccess = 1
	TaskFResultEFailed  = 2

	UrchinSuccessCode         = 0
	ObjectStatusInvalidCode   = 10350
	ErrCodeObjectAlreadyExist = 10001
)

var TaskTypeOnlyServiceRetry = map[int32]bool{
	TaskTypeMigrate: true,
	TaskTypeCopy:    true}

type UrchinError interface {
	ToError() error
	IsSuccess() bool
}

type ErrObjectStatusInvalid struct {
}

func IsErrObjectStatusInvalid(err error) bool {
	_, ok := err.(ErrObjectStatusInvalid)
	return ok
}

func (err ErrObjectStatusInvalid) Error() string {
	return fmt.Sprintf("dataset.resource_in_use")
}

type BaseResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
}

func (r *BaseResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *BaseResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type CreateInitiateMultipartUploadSignedUrlReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateInitiateMultipartUploadSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateUploadPartSignedUrlReq struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"part_number"
	PartNumber int32 `protobuf:"varint,2,opt,name=part_number,proto3" json:"part_number"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,3,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
}

type CreateUploadPartSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"part_number"
	PartNumber int32 `protobuf:"varint,2,opt,name=part_number,proto3" json:"part_number"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
}

type CreateListPartsSignedUrlReq struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
	// @inject_tag: json:"max_parts"
	MaxParts *int32 `protobuf:"varint,4,opt,name=max_parts,proto3,oneof" json:"max_parts"`
	// @inject_tag: json:"part_number_marker"
	PartNumberMarker *int32 `protobuf:"varint,5,opt,name=part_number_marker,proto3,oneof" json:"part_number_marker"`
}

type CreateListPartsSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
	// @inject_tag: json:"max_parts"
	MaxParts *int32 `protobuf:"varint,4,opt,name=max_parts,proto3,oneof" json:"max_parts"`
	// @inject_tag: json:"part_number_marker"
	PartNumberMarker *int32 `protobuf:"varint,5,opt,name=part_number_marker,proto3,oneof" json:"part_number_marker"`
}

type CreateCompleteMultipartUploadSignedUrlReq struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
}

type CreateCompleteMultipartUploadSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
}

type CreateAbortMultipartUploadSignedUrlReq struct {
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,1,opt,name=upload_id,proto3" json:"upload_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
}

type CreatePutObjectSignedUrlReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreatePutObjectSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateGetObjectMetadataSignedUrlReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateGetObjectMetadataSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,1,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateGetObjectSignedUrlReq struct {
	// @inject_tag: json:"range_start"
	RangeStart *int64 `protobuf:"varint,1,opt,name=range_start,proto3,oneof" json:"range_start"`
	// @inject_tag: json:"range_end"
	RangeEnd *int64 `protobuf:"varint,2,opt,name=range_end,proto3,oneof" json:"range_end"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,3,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
}

type CreateGetObjectSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"range_start"
	RangeStart *int64 `protobuf:"varint,1,opt,name=range_start,proto3,oneof" json:"range_start"`
	// @inject_tag: json:"range_end"
	RangeEnd *int64 `protobuf:"varint,2,opt,name=range_end,proto3,oneof" json:"range_end"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,3,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
	// @inject_tag: json:"query_params"
	QueryParams map[string]string `protobuf:"bytes,5,rep,name=query_params,proto3" json:"query_params"`
}

type CreateListObjectsSignedUrlReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"marker"
	Marker *string `protobuf:"bytes,2,opt,name=marker,proto3,oneof" json:"marker"`
}

type CreateListObjectsSignedUrlReqWithoutTask struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,1,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"marker"
	Marker *string `protobuf:"bytes,2,opt,name=marker,proto3,oneof" json:"marker"`
	// @inject_tag: json:"prefix"
	Prefix *string `protobuf:"bytes,3,opt,name=source,proto3" json:"prefix"`
}

type CreateSignedUrlResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"signed_url"
	SignedUrl string `protobuf:"bytes,4,opt,name=signed_url,proto3" json:"signed_url"`
	// @inject_tag: json:"header"
	Header map[string]*HeaderValues `protobuf:"bytes,5,rep,name=header,proto3" json:"header" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (r *CreateSignedUrlResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *CreateSignedUrlResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type HeaderValues struct {
	// @inject_tag: json:"values"
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values"`
}

type GetIpfsTokenReq struct {
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,1,opt,name=node_name,proto3" json:"node_name" url:"node_name"`
}

type GetIpfsTokenResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"url"
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token"`
}

type CreateJCSPreSignedObjectListReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"continuation_token"
	ContinuationToken *string `protobuf:"bytes,2,opt,name=continuation_token,proto3,oneof" json:"continuation_token"`
}

type CreateJCSPreSignedObjectUploadReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateJCSPreSignedObjectNewMultipartUploadReq struct {
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
}

type CreateJCSPreSignedObjectUploadPartReq struct {
	// @inject_tag: json:"object_id"
	ObjectId int32 `protobuf:"varint,1,opt,name=object_id,proto3" json:"object_id"`
	// @inject_tag: json:"index"
	Index int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,3,opt,name=task_id,proto3" json:"task_id"`
}

type CreateJCSPreSignedObjectCompleteMultipartUploadReq struct {
	// @inject_tag: json:"object_id"
	ObjectId int32 `protobuf:"varint,1,opt,name=object_id,proto3" json:"object_id"`
	// @inject_tag: json:"indexes"
	Indexes []int32 `protobuf:"varint,2,rep,packed,name=indexes,proto3" json:"indexes"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,3,opt,name=task_id,proto3" json:"task_id"`
}

type CreateJCSPreSignedObjectDownloadReq struct {
	// @inject_tag: json:"object_id"
	ObjectId int32 `protobuf:"varint,1,opt,name=object_id,proto3" json:"object_id"`
	// @inject_tag: json:"offset"
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset"`
	// @inject_tag: json:"length"
	Length int64 `protobuf:"varint,3,opt,name=length,proto3" json:"length"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
}

type CreateObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"desc"
	Desc *string `protobuf:"bytes,3,opt,name=desc,proto3,oneof" json:"desc"`
	// @inject_tag: json:"node_name"
	NodeName *string `protobuf:"bytes,4,opt,name=node_name,proto3,oneof" json:"node_name"`
}

type CreateObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,4,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"deployment_info"
	DeploymentInfo *ObjDeploymentData `protobuf:"bytes,5,opt,name=deployment_info,proto3" json:"deployment_info"`
}

func (r *CreateObjectResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *CreateObjectResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type UploadObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"source_local_path"
	SourceLocalPath string `protobuf:"bytes,3,opt,name=source_local_path,proto3" json:"source_local_path"`
	// @inject_tag: json:"desc"
	Desc *string `protobuf:"bytes,4,opt,name=desc,proto3,oneof" json:"desc"`
	// @inject_tag: json:"node_name"
	NodeName *string `protobuf:"bytes,5,opt,name=node_name,proto3,oneof" json:"node_name"`
}

type UploadObjectTaskParams struct {
	// @inject_tag: json:"request"
	Request *UploadObjectReq `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	// @inject_tag: json:"uuid"
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid"`
	// @inject_tag: json:"node_id"
	NodeId int32 `protobuf:"varint,3,opt,name=node_id,proto3" json:"node_id"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,4,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,5,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"location"
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,7,opt,name=request_id,proto3" json:"request_id"`
}

type UploadObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,5,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,6,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,7,opt,name=obj_uuid,proto3" json:"obj_uuid"`
}

func (r *UploadObjectResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *UploadObjectResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type UploadFileReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
	// @inject_tag: json:"source_local_path"
	SourceLocalPath string `protobuf:"bytes,4,opt,name=source_local_path,proto3" json:"source_local_path"`
	// @inject_tag: json:"size"
	Size *int32 `protobuf:"varint,5,opt,name=size,proto3,oneof" json:"size"`
}

type UploadFileTaskParams struct {
	// @inject_tag: json:"request"
	Request *UploadFileReq `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	// @inject_tag: json:"node_id"
	NodeId int32 `protobuf:"varint,2,opt,name=node_id,proto3" json:"node_id"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,3,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,4,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"location"
	Location string `protobuf:"bytes,5,opt,name=location,proto3" json:"location"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,6,opt,name=request_id,proto3" json:"request_id"`
}

type UploadFileResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,5,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,6,opt,name=node_name,proto3" json:"node_name"`
}

func (r *UploadFileResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *UploadFileResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type DownloadFileReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source"`
	// @inject_tag: json:"target_local_path"
	TargetLocalPath string `protobuf:"bytes,4,opt,name=target_local_path,proto3" json:"target_local_path"`
	// @inject_tag: json:"node_name"
	NodeName *string `protobuf:"bytes,5,opt,name=node_name,proto3,oneof" json:"node_name"`
}

type DownloadFileTaskParams struct {
	// @inject_tag: json:"request"
	Request *DownloadFileReq `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	// @inject_tag: json:"node_id"
	NodeId int32 `protobuf:"varint,2,opt,name=node_id,proto3" json:"node_id"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,3,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,4,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"bucket_name"
	BucketName string `protobuf:"bytes,5,opt,name=bucket_name,proto3" json:"bucket_name"`
	// @inject_tag: json:"location"
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,7,opt,name=request_id,proto3" json:"request_id"`
}

type DownloadFileResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,5,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"bucket_name"
	BucketName string `protobuf:"bytes,6,opt,name=bucket_name,proto3" json:"bucket_name"`
}

type GetObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id" url:"user_id"`
	// @inject_tag: json:"page_index"
	PageIndex int32 `protobuf:"varint,2,opt,name=page_index,proto3" json:"page_index" url:"page_index"`
	// @inject_tag: json:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,proto3" json:"page_size" url:"page_size"`
	// @inject_tag: json:"sort_by"
	SortBy *string `protobuf:"bytes,4,opt,name=sort_by,proto3,oneof" json:"sort_by" url:"sort_by,omitempty"`
	// @inject_tag: json:"order_by"
	OrderBy *string `protobuf:"bytes,5,opt,name=order_by,proto3,oneof" json:"order_by" url:"order_by,omitempty"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid *string `protobuf:"bytes,6,opt,name=obj_uuid,proto3,oneof" json:"obj_uuid" url:"obj_uuid,omitempty"`
}

type GetObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"data"
	Data *GetObjectRespData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

type GetObjectRespData struct {
	// @inject_tag: json:"total"
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// @inject_tag: json:"list"
	List []*ObjectInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list"`
}

type ObjectInfo struct {
	// @inject_tag: json:"meta_info"
	MetaInfo *DataObjData `protobuf:"bytes,1,opt,name=meta_info,proto3" json:"meta_info"`
	// @inject_tag: json:"deployment_info"
	DeploymentInfo []*ObjDeploymentData `protobuf:"bytes,2,rep,name=deployment_info,proto3" json:"deployment_info"`
}

type DataObjData struct {
	// @inject_tag: json:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"uuid"
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"desc"
	Desc string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc"`
	// @inject_tag: json:"status"
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,6,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"version"
	Version int32 `protobuf:"varint,7,opt,name=version,proto3" json:"version"`
	// @inject_tag: json:"create_time"
	CreateTime string `protobuf:"bytes,8,opt,name=create_time,proto3" json:"create_time"`
	// @inject_tag: json:"update_time"
	UpdateTime string `protobuf:"bytes,9,opt,name=update_time,proto3" json:"update_time"`
	// @inject_tag: json:"delete_time"
	DeleteTime string `protobuf:"bytes,10,opt,name=delete_time,proto3" json:"delete_time"`
}

type ObjDeploymentData struct {
	// @inject_tag: json:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"node_id"
	NodeId int32 `protobuf:"varint,3,opt,name=node_id,proto3" json:"node_id"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,4,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"location"
	Location string `protobuf:"bytes,5,opt,name=location,proto3" json:"location"`
	// @inject_tag: json:"space"
	Space string `protobuf:"bytes,6,opt,name=space,proto3" json:"space"`
	// @inject_tag: json:"type"
	Type int32 `protobuf:"varint,7,opt,name=type,proto3" json:"type"`
	// @inject_tag: json:"status"
	Status int32 `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	// @inject_tag: json:"ref_count"
	RefCount int32 `protobuf:"varint,9,opt,name=ref_count,proto3" json:"ref_count"`
	// @inject_tag: json:"version"
	Version int32 `protobuf:"varint,10,opt,name=version,proto3" json:"version"`
	// @inject_tag: json:"create_time"
	CreateTime string `protobuf:"bytes,11,opt,name=create_time,proto3" json:"create_time"`
	// @inject_tag: json:"update_time"
	UpdateTime string `protobuf:"bytes,12,opt,name=update_time,proto3" json:"update_time"`
	// @inject_tag: json:"delete_time"
	DeleteTime string `protobuf:"bytes,13,opt,name=delete_time,proto3" json:"delete_time"`
}

type DownloadObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"target_local_path"
	TargetLocalPath string `protobuf:"bytes,3,opt,name=target_local_path,proto3" json:"target_local_path"`
	// @inject_tag: json:"node_name"
	NodeName *string `protobuf:"bytes,4,opt,name=node_name,proto3,oneof" json:"node_name"`
}

type DownloadObjectTaskParams struct {
	// @inject_tag: json:"request"
	Request *DownloadObjectReq `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	// @inject_tag: json:"node_id"
	NodeId int32 `protobuf:"varint,2,opt,name=node_id,proto3" json:"node_id"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,3,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,4,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"bucket_name"
	BucketName string `protobuf:"bytes,5,opt,name=bucket_name,proto3" json:"bucket_name"`
	// @inject_tag: json:"location"
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,7,opt,name=request_id,proto3" json:"request_id"`
}

type DownloadObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"node_type"
	NodeType int32 `protobuf:"varint,5,opt,name=node_type,proto3" json:"node_type"`
	// @inject_tag: json:"bucket_name"
	BucketName string `protobuf:"bytes,6,opt,name=bucket_name,proto3" json:"bucket_name"`
}

type LoadObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source_node_name"
	SourceNodeName *string `protobuf:"bytes,3,opt,name=source_node_name,proto3,oneof" json:"source_node_name"`
	// @inject_tag: json:"target_node_name"
	TargetNodeName string `protobuf:"bytes,4,opt,name=target_node_name,proto3" json:"target_node_name"`
	// @inject_tag: json:"cache_local_path"
	CacheLocalPath string `protobuf:"bytes,5,opt,name=cache_local_path,proto3" json:"cache_local_path"`
}

type LoadObjectTaskParams struct {
	// @inject_tag: json:"request"
	Request *LoadObjectReq `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	// @inject_tag: json:"source_node_id"
	SourceNodeId int32 `protobuf:"varint,2,opt,name=source_node_id,proto3" json:"source_node_id"`
	// @inject_tag: json:"source_node_name"
	SourceNodeName string `protobuf:"bytes,3,opt,name=source_node_name,proto3" json:"source_node_name"`
	// @inject_tag: json:"source_node_type"
	SourceNodeType int32 `protobuf:"varint,4,opt,name=source_node_type,proto3" json:"source_node_type"`
	// @inject_tag: json:"source_bucket_name"
	SourceBucketName string `protobuf:"bytes,5,opt,name=source_bucket_name,proto3" json:"source_bucket_name"`
	// @inject_tag: json:"source_location"
	SourceLocation string `protobuf:"bytes,6,opt,name=source_location,proto3" json:"source_location"`
	// @inject_tag: json:"target_node_id"
	TargetNodeId int32 `protobuf:"varint,7,opt,name=target_node_id,proto3" json:"target_node_id"`
	// @inject_tag: json:"target_node_name"
	TargetNodeName string `protobuf:"bytes,8,opt,name=target_node_name,proto3" json:"target_node_name"`
	// @inject_tag: json:"target_node_type"
	TargetNodeType int32 `protobuf:"varint,9,opt,name=target_node_type,proto3" json:"target_node_type"`
	// @inject_tag: json:"target_bucket_name"
	TargetBucketName string `protobuf:"bytes,10,opt,name=target_bucket_name,proto3" json:"target_bucket_name"`
	// @inject_tag: json:"target_location"
	TargetLocation string `protobuf:"bytes,11,opt,name=target_location,proto3" json:"target_location"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,12,opt,name=request_id,proto3" json:"request_id"`
}

type LoadObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"source_node_id"
	SourceNodeId int32 `protobuf:"varint,5,opt,name=source_node_id,proto3" json:"source_node_id"`
	// @inject_tag: json:"source_node_name"
	SourceNodeName string `protobuf:"bytes,6,opt,name=source_node_name,proto3" json:"source_node_name"`
	// @inject_tag: json:"source_node_type"
	SourceNodeType int32 `protobuf:"varint,7,opt,name=source_node_type,proto3" json:"source_node_type"`
	// @inject_tag: json:"source_bucket_name"
	SourceBucketName string `protobuf:"bytes,8,opt,name=source_bucket_name,proto3" json:"source_bucket_name"`
	// @inject_tag: json:"source_location"
	SourceLocation string `protobuf:"bytes,9,opt,name=source_location,proto3" json:"source_location"`
	// @inject_tag: json:"target_node_id"
	TargetNodeId int32 `protobuf:"varint,10,opt,name=target_node_id,proto3" json:"target_node_id"`
	// @inject_tag: json:"target_node_name"
	TargetNodeName string `protobuf:"bytes,11,opt,name=target_node_name,proto3" json:"target_node_name"`
	// @inject_tag: json:"target_node_type"
	TargetNodeType int32 `protobuf:"varint,12,opt,name=target_node_type,proto3" json:"target_node_type"`
	// @inject_tag: json:"target_location"
	TargetLocation string `protobuf:"bytes,13,opt,name=target_location,proto3" json:"target_location"`
	// @inject_tag: json:"target_deployment_info"
	TargetDeploymentInfo *ObjDeploymentData `protobuf:"bytes,14,opt,name=target_deployment_info,proto3" json:"target_deployment_info"`
}

type MigrateObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"source_node_name"
	SourceNodeName *string `protobuf:"bytes,3,opt,name=source_node_name,proto3,oneof" json:"source_node_name"`
	// @inject_tag: json:"target_node_name"
	TargetNodeName string `protobuf:"bytes,4,opt,name=target_node_name,proto3" json:"target_node_name"`
}

type MigrateObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"target_deployment_info"
	TargetDeploymentInfo *ObjDeploymentData `protobuf:"bytes,5,opt,name=target_deployment_info,proto3" json:"target_deployment_info"`
}

type CopyObjectReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"source_obj_uuid"
	SourceObjUuid string `protobuf:"bytes,2,opt,name=source_obj_uuid,proto3" json:"source_obj_uuid"`
	// @inject_tag: json:"object_keys"
	ObjectKeys []string `protobuf:"bytes,3,rep,name=object_keys,proto3" json:"object_keys"`
	// @inject_tag: json:"target_obj_uuid"
	TargetObjUuid *string `protobuf:"bytes,4,opt,name=target_obj_uuid,proto3,oneof" json:"target_obj_uuid"`
	// @inject_tag: json:"name"
	Name *string `protobuf:"bytes,5,opt,name=name,proto3,oneof" json:"name"`
	// @inject_tag: json:"desc"
	Desc *string `protobuf:"bytes,6,opt,name=desc,proto3,oneof" json:"desc"`
}

type CopyObjectResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,4,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"dest_obj_uuid"
	DestObjUuid string `protobuf:"bytes,5,opt,name=dest_obj_uuid,proto3" json:"dest_obj_uuid"`
}

type PutObjectDeploymentReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"node_name"
	NodeName string `protobuf:"bytes,3,opt,name=node_name,proto3" json:"node_name"`
	// @inject_tag: json:"location"
	Location *string `protobuf:"bytes,4,opt,name=location,proto3,oneof" json:"location"`
}

type DeleteObjectReq struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
}

type DeleteFileReq struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,1,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
	// @inject_tag: json:"source"
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source" url:"source"`
}

type DeleteObjectDeploymentReq struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
	// @inject_tag: json:"force"
	Force *bool `protobuf:"varint,4,opt,name=force,proto3,oneof" json:"force" url:"force"`
}

type ListObjectsReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id" url:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
	// @inject_tag: json:"prefix"
	Prefix *string `protobuf:"bytes,3,opt,name=prefix,proto3,oneof" json:"prefix" url:"prefix,omitempty"`
	// @inject_tag: json:"marker"
	Marker *string `protobuf:"bytes,4,opt,name=marker,proto3,oneof" json:"marker" url:"marker,omitempty"`
	// @inject_tag: json:"max_keys"
	MaxKeys *int32 `protobuf:"varint,5,opt,name=max_keys,proto3,oneof" json:"max_keys" url:"max_keys,omitempty"`
	// @inject_tag: json:"delimiter"
	Delimiter *string `protobuf:"bytes,6,opt,name=delimiter,proto3,oneof" json:"delimiter" url:"delimiter,omitempty"`
}

type ListObjectsResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"data"
	Data ListObjectsRespData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

func (r *ListObjectsResp) ToError() error {
	if r.Code == ObjectStatusInvalidCode {
		return ErrObjectStatusInvalid{}
	}
	return fmt.Errorf("urchin operate error.[%d]%s", r.Code, r.Message)
}

func (r *ListObjectsResp) IsSuccess() bool {
	return r.Code == UrchinSuccessCode
}

type ListObjectsRespData struct {
	// @inject_tag: json:"next_marker"
	NextMarker string `protobuf:"bytes,1,opt,name=next_marker,proto3" json:"next_marker"`
	// @inject_tag: json:"list"
	List []*ObjectContent `protobuf:"bytes,2,rep,name=list,proto3" json:"list"`
	// @inject_tag: json:"common_prefixes"
	CommonPrefixes []string `protobuf:"bytes,3,rep,name=common_prefixes,proto3" json:"common_prefixes"`
	// @inject_tag: json:"is_truncated"
	IsTruncated bool `protobuf:"varint,4,opt,name=is_truncated,proto3" json:"is_truncated"`
}

type ObjectContent struct {
	// @inject_tag: json:"key"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// @inject_tag: json:"etag"
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag"`
	// @inject_tag: json:"size"
	// Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	Size string `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// @inject_tag: json:"last_modified"
	LastModified time.Time `protobuf:"bytes,4,opt,name=last_modified,proto3" json:"last_modified"`
}

type GetObjectMetadataReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"object_key"
	ObjectKey string `protobuf:"bytes,3,opt,name=object_key,proto3" json:"object_key"`
}

type GetObjectMetadataResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"data"
	Data *GetObjectMetadataData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

type GetObjectMetadataData struct {
	// @inject_tag: json:"content_length"
	ContentLength int64 `protobuf:"varint,1,opt,name=content_length,proto3" json:"content_length"`
	// @inject_tag: json:"etag"
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag"`
	// @inject_tag: json:"last_modified"
	LastModified string `protobuf:"bytes,3,opt,name=last_modified,proto3" json:"last_modified"`
}

type ListPartsReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id" url:"user_id"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,2,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
	// @inject_tag: json:"upload_id"
	UploadId string `protobuf:"bytes,3,opt,name=upload_id,proto3" json:"upload_id" url:"upload_id"`
	// @inject_tag: json:"max_parts"
	MaxParts *int32 `protobuf:"varint,4,opt,name=max_parts,proto3,oneof" json:"max_parts" url:"max_parts,omitempty"`
	// @inject_tag: json:"part_number_marker"
	PartNumberMarker *int32 `protobuf:"varint,5,opt,name=part_number_marker,proto3,oneof" json:"part_number_marker" url:"part_number_marker,omitempty"`
}

type ListPartsResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"data"
	Data *ListPartsRespData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

type ListPartsRespData struct {
	// @inject_tag: json:"is_truncated"
	IsTruncated bool `protobuf:"varint,1,opt,name=is_truncated,proto3" json:"is_truncated"`
	// @inject_tag: json:"next_part_number_marker"
	NextPartNumberMarker int32 `protobuf:"varint,2,opt,name=next_part_number_marker,proto3" json:"next_part_number_marker"`
	// @inject_tag: json:"list"
	List []*PartContent `protobuf:"bytes,3,rep,name=list,proto3" json:"list"`
}

type PartContent struct {
	// @inject_tag: json:"part_number"
	PartNumber int32 `protobuf:"varint,1,opt,name=part_number,proto3" json:"part_number"`
	// @inject_tag: json:"etag"
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag"`
	// @inject_tag: json:"size"
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// @inject_tag: json:"last_modified"
	LastModified string `protobuf:"bytes,4,opt,name=last_modified,proto3" json:"last_modified"`
}

type GetTaskReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id" url:"user_id"`
	// @inject_tag: json:"page_index"
	PageIndex int32 `protobuf:"varint,2,opt,name=page_index,proto3" json:"page_index" url:"page_index"`
	// @inject_tag: json:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,proto3" json:"page_size" url:"page_size"`
	// @inject_tag: json:"sort_by"
	SortBy *string `protobuf:"bytes,4,opt,name=sort_by,proto3,oneof" json:"sort_by" url:"sort_by,omitempty"`
	// @inject_tag: json:"order_by"
	OrderBy *string `protobuf:"bytes,5,opt,name=order_by,proto3,oneof" json:"order_by" url:"order_by,omitempty"`
	// @inject_tag: json:"task_id"
	TaskId *int32 `protobuf:"varint,6,opt,name=task_id,proto3,oneof" json:"task_id" url:"task_id,omitempty"`
}

type GetTaskResp struct {
	// @inject_tag: json:"code"
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// @inject_tag: json:"request_id"
	RequestId string `protobuf:"bytes,3,opt,name=request_id,proto3" json:"request_id"`
	// @inject_tag: json:"data"
	Data *GetTaskRespData `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

type GetTaskRespData struct {
	// @inject_tag: json:"total"
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// @inject_tag: json:"list"
	List []*TaskDetail `protobuf:"bytes,2,rep,name=list,proto3" json:"list"`
}

type TaskDetail struct {
	// @inject_tag: json:"task"
	Task *TaskData `protobuf:"bytes,1,opt,name=task,proto3" json:"task"`
	// @inject_tag: json:"execs"
	Execs []*TaskExecData `protobuf:"bytes,2,rep,name=execs,proto3" json:"execs"`
}

type TaskData struct {
	// @inject_tag: json:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"type"
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"params"
	Params string `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
	// @inject_tag: json:"status"
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status"`
	// @inject_tag: json:"task_exec_id"
	TaskExecId int32 `protobuf:"varint,6,opt,name=task_exec_id,proto3" json:"task_exec_id"`
	// @inject_tag: json:"result"
	Result int32 `protobuf:"varint,7,opt,name=result,proto3" json:"result"`
	// @inject_tag: json:"return"
	Return string `protobuf:"bytes,8,opt,name=return,proto3" json:"return"`
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,9,opt,name=obj_uuid,proto3" json:"obj_uuid"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,10,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"start_time"
	StartTime string `protobuf:"bytes,11,opt,name=start_time,proto3" json:"start_time"`
	// @inject_tag: json:"finish_time"
	FinishTime string `protobuf:"bytes,12,opt,name=finish_time,proto3" json:"finish_time"`
	// @inject_tag: json:"create_time"
	CreateTime string `protobuf:"bytes,13,opt,name=create_time,proto3" json:"create_time"`
	// @inject_tag: json:"update_time"
	UpdateTime string `protobuf:"bytes,14,opt,name=update_time,proto3" json:"update_time"`
}

type TaskExecData struct {
	// @inject_tag: json:"id"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"context"
	Context string `protobuf:"bytes,3,opt,name=context,proto3" json:"context"`
	// @inject_tag: json:"status"
	Status int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status"`
	// @inject_tag: json:"result"
	Result int32 `protobuf:"varint,5,opt,name=result,proto3" json:"result"`
	// @inject_tag: json:"return"
	Return string `protobuf:"bytes,6,opt,name=return,proto3" json:"return"`
	// @inject_tag: json:"start_time"
	StartTime string `protobuf:"bytes,7,opt,name=start_time,proto3" json:"start_time"`
	// @inject_tag: json:"finish_time"
	FinishTime string `protobuf:"bytes,8,opt,name=finish_time,proto3" json:"finish_time"`
	// @inject_tag: json:"create_time"
	CreateTime string `protobuf:"bytes,9,opt,name=create_time,proto3" json:"create_time"`
	// @inject_tag: json:"update_time"
	UpdateTime string `protobuf:"bytes,10,opt,name=update_time,proto3" json:"update_time"`
}

type FinishTaskReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
	// @inject_tag: json:"result"
	Result int32 `protobuf:"varint,3,opt,name=result,proto3" json:"result"`
	// @inject_tag: json:"return"
	Return *string `protobuf:"bytes,4,opt,name=return,proto3,oneof" json:"return"`
}

type RetryTaskReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
}

type ReportTaskHeartbeatReq struct {
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"task_id"
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,proto3" json:"task_id"`
}

type XIpfsUpload struct {
	XMLName xml.Name `xml:"IpfsUpload"`
	CId     string   `xml:"CId"`
	Result  int      `xml:"Result"`
}

type XIpfsDownload struct {
	XMLName xml.Name `xml:"IpfsDownload"`
	Result  int      `xml:"Result"`
}

type StorageNodeConfig struct {
	// @inject_tag: json:"endpoint"
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint"`
	// @inject_tag: json:"access_key"
	AccessKey string `protobuf:"bytes,2,opt,name=access_key,proto3" json:"access_key"`
	// @inject_tag: json:"secret_key"
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,proto3" json:"secret_key"`
	// @inject_tag: json:"user"
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user"`
	// @inject_tag: json:"pass"
	Pass string `protobuf:"bytes,5,opt,name=pass,proto3" json:"pass"`
	// @inject_tag: json:"bucket_name"
	BucketName string `protobuf:"bytes,6,opt,name=bucket_name,proto3" json:"bucket_name"`
	// @inject_tag: json:"auth_service"
	AuthService string `protobuf:"bytes,7,opt,name=auth_service,proto3" json:"auth_service"`
	// @inject_tag: json:"auth_region"
	AuthRegion string `protobuf:"bytes,8,opt,name=auth_region,proto3" json:"auth_region"`
	// @inject_tag: json:"user_id"
	UserId int32 `protobuf:"varint,9,opt,name=user_id,proto3" json:"user_id"`
	// @inject_tag: json:"bucket_id"
	BucketId int32 `protobuf:"varint,10,opt,name=bucket_id,proto3" json:"bucket_id"`
	// @inject_tag: json:"base_path"
	BasePath string `protobuf:"bytes,11,opt,name=base_path,proto3" json:"base_path"`
	// @inject_tag: json:"url"
	Url string `protobuf:"bytes,12,opt,name=url,proto3" json:"url"`
	// @inject_tag: json:"cluster_id"
	ClusterId string `protobuf:"bytes,13,opt,name=cluster_id,proto3" json:"cluster_id"`
	// @inject_tag: json:"org_id"
	OrgId string `protobuf:"bytes,14,opt,name=org_id,proto3" json:"org_id"`
	// @inject_tag: json:"pass_magic"
	PassMagic string `protobuf:"bytes,15,opt,name=pass_magic,proto3" json:"pass_magic"`
	// @inject_tag: json:"req_timeout"
	ReqTimeout int32 `protobuf:"varint,16,opt,name=req_timeout,proto3" json:"req_timeout"`
	// @inject_tag: json:"max_connection"
	MaxConnection int32 `protobuf:"varint,17,opt,name=max_connection,proto3" json:"max_connection"`
	// @inject_tag: json:"concurrency"
	Concurrency *StorageNodeConcurrencyConfig `protobuf:"bytes,18,opt,name=concurrency,proto3" json:"concurrency"`
	// @inject_tag: json:"rate"
	Rate *StorageNodeRateConfig `protobuf:"bytes,19,opt,name=rate,proto3" json:"rate"`
	// @inject_tag: json:"part"
	Part *StorageNodePartConfig `protobuf:"bytes,20,opt,name=part,proto3" json:"part"`
}

type StorageNodeConcurrencyConfig struct {
	UploadFileTaskNum    int32 `protobuf:"varint,1,opt,name=upload_file_task_num,proto3" json:"upload_file_task_num,omitempty"`
	UploadMultiTaskNum   int32 `protobuf:"varint,2,opt,name=upload_multi_task_num,proto3" json:"upload_multi_task_num,omitempty"`
	DownloadFileTaskNum  int32 `protobuf:"varint,3,opt,name=download_file_task_num,proto3" json:"download_file_task_num,omitempty"`
	DownloadMultiTaskNum int32 `protobuf:"varint,4,opt,name=download_multi_task_num,proto3" json:"download_multi_task_num,omitempty"`
}

type StorageNodeRateConfig struct {
	Limit float32 `protobuf:"fixed32,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Burst int32   `protobuf:"varint,2,opt,name=burst,proto3" json:"burst,omitempty"`
}

type StorageNodePartConfig struct {
	PartLimit int64 `protobuf:"varint,1,opt,name=part_limit,proto3" json:"part_limit,omitempty"`
	PartSize  int64 `protobuf:"varint,2,opt,name=part_size,proto3" json:"part_size,omitempty"`
}

type ReportSizeChangedReq struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,1,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
}

type RelateObjectReq struct {
	UserId   string  `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Desc     *string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc"`
	NodeName string  `protobuf:"bytes,4,opt,name=node_name,proto3" json:"node_name"`
	Location string  `protobuf:"bytes,5,opt,name=location,proto3" json:"location"`
}

type UpdateObjectLastModifyTimeReq struct {
	// @inject_tag: json:"obj_uuid"
	ObjUuid string `protobuf:"bytes,1,opt,name=obj_uuid,proto3" json:"obj_uuid" url:"obj_uuid"`
}
