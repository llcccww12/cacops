package upload

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/subject_service"
)

func GetUploadChunks(ctx *context.APIContext) {
	fileMD5 := ctx.Query("md5")
	fileName := ctx.Query("file_name")
	subjectId := ctx.Query("subject_id")
	subjectType := ctx.QueryInt("subject_type")

	if fileMD5 == "" || subjectId == "" || subjectType < 0 {
		log.Error("GetUploadChunks error, fileMD5=%s, subjectId=%s, subjectType=%d", fileMD5, subjectId, subjectType)
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}

	req := entity.UploadChunkRequest{
		FileMD5:     fileMD5,
		FileName:    fileName,
		SubjectID:   subjectId,
		SubjectType: subjectType,
		UserId:      ctx.User.ID,
	}
	res, err := subject_service.GetFileChunks(req)
	if err != nil {
		log.Error("GetFileChunks error, req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
	}

	ctx.JSON(200, response.OuterSuccessWithData(res))
}

func NewUploadMultipart(ctx *context.APIContext) {
	req := entity.NewMultipartRequest{
		SubjectId:            ctx.Query("subject_id"),
		FileName:             ctx.Query("file_name"),
		FileType:             ctx.Query("file_type"),
		SubjectType:          ctx.QueryInt("subject_type"),
		TotalChunkCounts:     ctx.QueryInt("total_chunk_counts"),
		Size:                 ctx.QueryInt64("size"),
		MD5:                  ctx.Query("md5"),
		User:                 ctx.User,
		SubjectAccessContext: ctx.AccessContext.SubjectAccessContext,
	}
	if req.SubjectId == "" || req.FileName == "" || req.SubjectType < 0 || req.TotalChunkCounts <= 0 || req.Size <= 0 || req.MD5 == "" {
		log.Error("NewUploadMultipart error, req=%+v", req)
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}
	res, err := subject_service.NewUploadMultipart(req)
	if err != nil {
		log.Error("NewUploadMultipart error,req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccessWithData(res))
}

func GetMultipartUploadUrl(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	partNumber := ctx.QueryInt("chunk_number")
	size := ctx.QueryInt64("size")

	req := entity.GetMultipartUrlRequest{
		UUID:       uuid,
		PartNumber: partNumber,
		Size:       size,
	}

	url, err := subject_service.GetMultipartUploadUrl(req)
	if err != nil {
		log.Error("GetMultipartUploadUrl error, req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccessWithData(map[string]string{
		"url": url,
	}))

}

func GetUploadUrl(ctx *context.Context) {
	subjectCtx := ctx.AccessContext.SubjectAccessContext
	req := entity.GetUploadUrlRequest{
		SubjectContext: subjectCtx,
		FileName:       ctx.Query("file_name"),
		Size:           ctx.QueryInt64("size"),
		FileType:       ctx.Query("file_type"),
		User:           ctx.User,
	}
	url, err := subject_service.GetUploadUrl(req)
	if err != nil {
		log.Error("GetUploadUrl error, req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccessWithData(map[string]string{
		"url": url,
	}))

}

func CompleteUploadMultipart(ctx *context.Context) {
	req := entity.CompleteMultipartRequest{
		UUID: ctx.Query("uuid"),
	}
	err := subject_service.CompleteModelMultipart(req)
	if err != nil {
		log.Error("CompleteModelMultipart error, req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccess())

}

func CompleteUpload(ctx *context.Context, req entity.CompleteUploadRequest) {
	req.SubjectContext = ctx.AccessContext.SubjectAccessContext
	if len(req.FileNameList) == 0 {
		ctx.JSON(200, response.OuterSuccess())
		return
	}
	err := subject_service.CompleteUpload(req)
	if err != nil {
		log.Error("CompleteUpload error, req=%+v err=%v", req, err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccess())

}
