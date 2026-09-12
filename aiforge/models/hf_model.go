package models

import (
	"fmt"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	HfTransferStatusWaiting = -1
	HfTransferStatusSuccess = 1
	HfTransferStatusOngoing = 0
	HfTransferStatusFailure = 2

	HfOperationTypeUpdate = 1
	HfOperationTypeRetry  = 2
)

type HfModelFile struct {
	ID       int64  `xorm:"pk autoincr" json:"id"`
	Filename string `xorm:"varchar(400)" json:"filename"`
	Status   int    `xorm:"NOT NULL DEFAULT 0" json:"status"`
	Size     int64  `xorm:"DEFAULT 0 NOT NULL" json:"size"`

	HfRepoId    string `xorm:"varchar(400)" json:"hf_repo_id"`
	OpeniRepoId string `xorm:"varchar(400)" json:"openi_repo_id"`
	ModelName   string `xorm:"varchar(400)" json:"model_name"`
	ModelId     string `xorm:"varchar(400)" json:"model_id"`
	RepoId      int64  `xorm:"DEFAULT 0 " json:"repo_id"`
	OwnerId     int64  `xorm:"DEFAULT 0" json:"owner_id"`
	UserId      int64  `xorm:"DEFAULT 0" json:"user_id"`

	Path      string `xorm:"varchar(400) NULL" json:"path"`
	HfBlobId  string `xorm:"varchar(400)" json:"hf_blob_id"`
	HfSha     string `xorm:"varchar(400)" json:"hf_sha"`
	HfLicense string `xorm:"varchar(400)" json:"hf_license"`
	HfTags    string `xorm:"varchar(400)" json:"hf_tags"`

	CreatedUnix   timeutil.TimeStamp `xorm:"created" json:"created_unix"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"updated" json:"updated_unix"`
	LastHeartbeat timeutil.TimeStamp `json:"last_heartbeat"`
	HeartbeatSec  int                `xorm:"NOT NULL DEFAULT 0" json:"heartbeat_sec"`
	HfToken       string             `xorm:"varchar(400)" json:"hf_token"`
}

type HfModelOperation struct {
	ID          int64              `xorm:"pk autoincr" json:"id"`
	ModelId     string             `xorm:"varchar(400)" json:"model_id"`
	UserId      int64              `xorm:"DEFAULT 0" json:"user_id"`
	OpType      int64              `xorm:"DEFAULT 0" json:"op_type"`
	CreatedUnix timeutil.TimeStamp `xorm:"created" json:"created_unix"`
}

func AddOperation(operation *HfModelOperation) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Insert(operation)
	if err != nil {
		return err
	}
	return nil
}

func GetOperations(modelId string) ([]*HfModelOperation, error) {
	sess := x.NewSession()
	defer sess.Close()
	var operations []*HfModelOperation
	err := sess.Where("model_id = ?", modelId).Find(&operations)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetUniqueOperatorIds(modelId string) ([]int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	var userIds []int64
	err := sess.Table("hf_model_operation").
		Where("model_id = ?", modelId).
		Distinct("user_id").
		Cols("user_id").
		Find(&userIds)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

func GetOperatorModelIds(userId int64) ([]string, error) {
	sess := x.NewSession()
	defer sess.Close()
	var modelIds []string
	err := sess.Table("hf_model_operation").
		Where("user_id = ?", userId).
		Distinct("model_id").
		Cols("model_id").
		Find(&modelIds)
	if err != nil {
		return nil, err
	}
	return modelIds, nil
}

func GetCreatorModelIds(userId int64) ([]string, error) {
	sess := x.NewSession()
	defer sess.Close()
	var modelIds []string
	err := sess.Table("hf_model_file").
		Where("user_id = ?", userId).
		GroupBy("model_id").
		Cols("model_id").
		Find(&modelIds)
	if err != nil {
		return nil, err
	}
	return modelIds, nil
}

func UpdateHfFileStatusAndHeartbeat(fileRecordId int64, statusNew int, heartbeat time.Time, heartbeatSec int) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("id = ?", fileRecordId).
		Cols("status", "last_heartbeat", "heartbeat_sec").
		Update(&HfModelFile{Status: statusNew, LastHeartbeat: timeutil.TimeStamp(heartbeat.Unix()), HeartbeatSec: heartbeatSec})
	if err != nil {
		return err
	}
	return nil
}

func UpdateHfFileHeartbeat(fileRecordId int64, heartbeat time.Time) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("id = ?", fileRecordId).
		Cols("last_heartbeat").
		Update(&HfModelFile{LastHeartbeat: timeutil.TimeStamp(heartbeat.Unix())})
	if err != nil {
		return err
	}
	return nil
}

func QueryFileById(fileId int64) (*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()

	var record HfModelFile
	_, err := sess.Where("id = ?", fileId).Get(&record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func GetExpiredHfFiles() ([]*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []*HfModelFile
	err := sess.Where("status = ? AND last_heartbeat < ?",
		HfTransferStatusOngoing, time.Now().Add(-5*time.Minute)).
		Find(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetUserIdByModelId(modelId string) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	var file HfModelFile
	has, err := sess.Where("model_id = ?", modelId).Get(&file)
	if err != nil {
		return 0, err
	}
	if !has {
		return 0, nil
	}
	return file.UserId, nil
}

func SaveRecord(record *HfModelFile) error {
	sess := x.NewSession()
	defer sess.Close()

	_, err := sess.Insert(record)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(hfRepoId string, modelOwner *User, filename string, BlobIdNew string, SizeNew int64, statusNew int, hfToken string) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("hf_repo_id = ? AND filename = ? AND owner_id = ?", hfRepoId, filename, modelOwner.ID).
		Cols("hf_blob_id", "size", "status", "hf_token").
		Update(&HfModelFile{HfBlobId: BlobIdNew, Size: SizeNew, Status: statusNew, HfToken: hfToken})
	if err != nil {
		log.Error("[hf_model_service] UpdateHfModelFile: %v", err)
		return err
	}
	return nil
}

func DeleteRecord(hfRepoId string, modelOwner *User, filename string) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("hf_repo_id = ? AND filename = ? AND owner_id = ?", hfRepoId, filename, modelOwner.ID).Delete(&HfModelFile{})
	if err != nil {
		log.Error("[hf_model_service] DeleteHfModelFile: %v", err)
		return err
	}
	return nil
}

func UpdateHfFile(record *HfModelFile) error {
	sess := x.NewSession()
	defer sess.Close()

	_, err := sess.ID(record.ID).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func QueryHfFileById(fileId int64) (*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var record HfModelFile
	_, err := sess.Where("id = ?", fileId).Get(&record)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func QueryModelFiles(hfRepoId string, modelOwner *User) ([]*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []*HfModelFile
	err := sess.Where("hf_repo_id = ? AND owner_id = ?", hfRepoId, modelOwner.ID).OrderBy("filename asc").Find(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetWaitlist(modelOwner *User, numFiles int, hfRepoId string) ([]*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []*HfModelFile

	cond := builder.Eq{
		"status":   HfTransferStatusWaiting,
		"owner_id": modelOwner.ID,
	}
	if hfRepoId != "" {
		cond["hf_repo_id"] = hfRepoId
	}

	query := sess.Where(cond)
	query = query.OrderBy("created_unix ASC")
	if numFiles > 0 {
		query = query.Limit(numFiles)
	}
	err := query.Find(&files)

	if err != nil {
		return nil, fmt.Errorf("failed to query waitlist: %w", err)
	}

	return files, nil
}

func UpdateHfFileStatus(fileRecordId int64, statusNew int) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Where("id = ?", fileRecordId).Cols("status", "last_heartbeat").Update(&HfModelFile{Status: statusNew, LastHeartbeat: timeutil.TimeStampNow()})
	if err != nil {
		return err
	}
	return nil
}

func BatchUpdateHfFileStatus(fileIds []int64, statusNew int) error {
	if len(fileIds) == 0 {
		return nil
	}
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.In("id", fileIds).
		Cols("status", "last_heartbeat").
		Update(&HfModelFile{Status: statusNew, LastHeartbeat: timeutil.TimeStampNow()})
	if err != nil {
		return err
	}
	return nil
}

func BatchUpdateHfFileHeartbeat(fileIds []int64) error {
	if len(fileIds) == 0 {
		return nil
	}
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.In("id", fileIds).
		Cols("last_heartbeat").
		Update(&HfModelFile{LastHeartbeat: timeutil.TimeStampNow()})
	if err != nil {
		return err
	}
	return nil
}

func GetFailedFiles(modelId string) ([]*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []*HfModelFile
	err := sess.Where("model_id = ? AND status = ?", modelId, HfTransferStatusFailure).Find(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func IsTransferComplete(modelId string) (bool, error) {
	sess := x.NewSession()
	defer sess.Close()

	var files []HfModelFile
	err := sess.Where("model_id = ?", modelId).Find(&files)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.Status == HfTransferStatusOngoing || file.Status == HfTransferStatusWaiting {
			return false, nil
		}
	}

	return true, nil
}

func IsTransferSuccess(modelId string) (bool, error) {
	sess := x.NewSession()
	defer sess.Close()

	var files []HfModelFile
	err := sess.Where("model_id = ?", modelId).Find(&files)
	if err != nil {
		return false, err
	}

	var totalCount, totalStatus int
	for _, file := range files {
		totalCount++
		totalStatus += file.Status
	}

	return totalCount == totalStatus, nil
}

func DeleteHfFilesByModelId(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Delete(&HfModelFile{
		ModelId: id,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteHfFilesByModelIdWithContext(ctx DBContext, id string) error {
	sess := ctx.e
	_, err := sess.Delete(&HfModelFile{
		ModelId: id,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetHfModelSizeByModelId(modelId string) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []HfModelFile
	err := sess.Where("model_id = ?", modelId).Find(&files)
	if err != nil {
		return 0, err
	}
	var totalSize int64
	for _, file := range files {
		totalSize += file.Size
	}
	return totalSize, nil
}

func GetRunningFiles() ([]*HfModelFile, error) {
	sess := x.NewSession()
	defer sess.Close()
	var files []*HfModelFile
	err := sess.Where("status = ?", HfTransferStatusOngoing).Find(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetRunningModelByUser(userId int64) ([]string, error) {
	sess := x.NewSession()
	defer sess.Close()
	var hfRepoIds []string
	err := sess.Table("hf_model_file").
		Where("user_id = ? AND (status = ? OR status = ?)", userId, HfTransferStatusOngoing, HfTransferStatusWaiting).
		GroupBy("hf_repo_id").
		Cols("hf_repo_id").
		Find(&hfRepoIds)
	if err != nil {
		return nil, err
	}
	return hfRepoIds, nil
}

type ModelQueryOptions struct {
	Page             int
	PageSize         int
	Keyword          string
	KeywordType      int
	Status           int
	CreatedUnixStart int
	CreatedUnixEnd   int
}

type ModelQueryData struct {
	TotalCount        int64                    `json:"total_count"`
	ID                string                   `xorm:"id" json:"id"`
	Name              string                   `xorm:"name" json:"name"`
	HFRepoID          string                   `xorm:"hf_repo_id" json:"hf_repo_id"`
	OpenIRepoID       string                   `xorm:"openi_repo_id" json:"openi_repo_id"`
	Source            string                   `xorm:"source" json:"source"`
	UserId            int64                    `xorm:"user_id" json:"user_id"`
	Username          string                   `xorm:"username" json:"username"`
	TotalFiles        int64                    `xorm:"total_files" json:"total_files"`
	FinishedFiles     int64                    `xorm:"finished_files" json:"finished_files"`
	TotalSize         int64                    `xorm:"total_size" json:"total_size"`
	FinishedSize      int64                    `xorm:"finished_size" json:"finished_size"`
	CreatedUnix       int64                    `xorm:"created_unix" json:"created_unix"`
	AiModelStatus     int64                    `xorm:"ai_model_status" json:"ai_model_status"`
	OperatorsList     []ModelOperatorQueryData `xorm:"operators_list" json:"operators_list"`
	OperatorsNameList []string                 `xorm:"operators_name_list" json:"operators_name_list"`
}

type ModelOperatorQueryData struct {
	OperationUnix int64  `xorm:"operation_unix" json:"operation_unix"`
	OperationType int64  `xorm:"operation_type" json:"operation_type"`
	Username      string `xorm:"username" json:"username"`
	UserId        int64  `xorm:"user_id" json:"user_id"`
}

func QueryHfModelStats(opts *ModelQueryOptions) ([]*ModelQueryData, int64, error) {
	sess := x.NewSession()
	defer sess.Close()

	baseSQL := `SELECT model_stats.*, COUNT(*) OVER() AS total_count
	FROM (
		SELECT ai.id,
			   ai.name,
			   hf.hf_repo_id,
			   hf.openi_repo_id,
			   'huggingface'                                                                  AS source,
			   hf.user_id                                                                     AS user_id,
			   COALESCE(creator.name, 'Ghost')                                                AS username,
			   hf.total_files,
			   hf.finished_files,
			   hf.total_size,
			   hf.finished_size,
			   ai.created_unix,
			   CASE
				   WHEN ai.status = 1 AND hf.hf_sum_status = -1 * hf.total_files THEN -1
				   ELSE ai.status END                                                        AS ai_model_status,
			   COALESCE(hf_ops.operators, '[]'::jsonb)                                       AS operators_list,
			   COALESCE(hf_ops.operators_names, '[]'::jsonb)                                 AS operators_name_list
		FROM ai_model_manage AS ai
		LEFT JOIN (
			SELECT model_id,
				   hf_repo_id,
				   openi_repo_id,
				   user_id,
				   COUNT(*)                                       AS total_files,
				   SUM(size)                                      AS total_size,
				   SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END)    AS finished_files,
				   SUM(CASE WHEN status = 1 THEN size ELSE 0 END) AS finished_size,
				   SUM(status)                                    AS hf_sum_status
			FROM hf_model_file
			GROUP BY model_id, hf_repo_id, openi_repo_id, user_id) AS hf ON ai.id = hf.model_id
		LEFT JOIN "user" AS creator ON hf.user_id = creator.id
		LEFT JOIN (
			SELECT model_id,
				   jsonb_agg(jsonb_build_object(
						   'operation_type', hf_op.op_type,
						   'username', COALESCE(u.name, 'Ghost'),
						   'user_id', u.id,
						   'operation_unix', hf_op.created_unix)) AS operators,
				   jsonb_agg(COALESCE(u.name, 'Ghost'))           AS operators_names
				FROM hf_model_operation AS hf_op
				LEFT JOIN "user" AS u ON hf_op.user_id = u.id
				GROUP BY model_id) AS hf_ops ON ai.id = hf_ops.model_id
		WHERE ai.model_type = 2) AS model_stats
	`

	// Building conditions
	var conditions []string

	//-1=waiting, 1=running, 0=finished, 2=error
	if opts.Status >= -1 && opts.Status <= 2 {
		conditions = append(conditions, fmt.Sprintf(" ai_model_status = %d ", opts.Status))
	}

	if opts.Keyword != "" {
		var queryCond string
		// search_model=1, search_hf_repo=2, search_user=3,
		switch opts.KeywordType {
		case 1:
			queryCond = fmt.Sprintf(" name ILIKE '%%%s%%' ", opts.Keyword)
		case 2:
			queryCond = fmt.Sprintf(" hf_repo_id ILIKE '%%%s%%' ", opts.Keyword)
		case 3:
			queryCond = fmt.Sprintf(" username ILIKE '%%%s%%' OR  operators_name_list @> '[\"%s\"]'::jsonb ", opts.Keyword, opts.Keyword)
		}
		conditions = append(conditions, queryCond)
	}

	if opts.CreatedUnixStart > 0 && opts.CreatedUnixEnd > 0 {
		conditions = append(conditions, fmt.Sprintf(" created_unix BETWEEN %d AND %d ", opts.CreatedUnixStart, opts.CreatedUnixEnd))
	}

	if len(conditions) > 0 {
		baseSQL += " WHERE " + strings.Join(conditions, " AND ")
	}

	baseSQL += " ORDER BY ARRAY_POSITION(ARRAY[1,-1,0,2], ai_model_status), created_unix DESC "

	if opts.Page > 0 && opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		baseSQL += fmt.Sprintf(" LIMIT %d OFFSET %d ", opts.PageSize, offset)
	}

	var modelList []*ModelQueryData
	err := sess.SQL(baseSQL).Find(&modelList)
	if err != nil {
		log.Error("QueryModel Error: ", err)
		return nil, 0, err
	}

	var total int64
	if len(modelList) > 0 {
		total = modelList[0].TotalCount
	}

	return modelList, total, nil
}
