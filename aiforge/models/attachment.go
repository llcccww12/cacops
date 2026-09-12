// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"

	gouuid "github.com/satori/go.uuid"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	DecompressStateInit int32 = iota
	DecompressStateDone
	DecompressStateIng
	DecompressStateFailed
)

// Attachment represent a attachment of issue/comment/release.
type Attachment struct {
	ID              int64  `xorm:"pk autoincr"`
	UUID            string `xorm:"uuid UNIQUE"`
	IssueID         int64  `xorm:"INDEX"`
	DatasetID       int64  `xorm:"INDEX DEFAULT 0"`
	ReleaseID       int64  `xorm:"INDEX"`
	UploaderID      int64  `xorm:"INDEX DEFAULT 0"` // Notice: will be zero before this column added
	CommentID       int64
	Name            string
	Description     string             `xorm:"TEXT"`
	DownloadCount   int64              `xorm:"DEFAULT 0"`
	UseNumber       int64              `xorm:"DEFAULT 0"`
	Size            int64              `xorm:"DEFAULT 0"`
	IsPrivate       bool               `xorm:"DEFAULT false"`
	DecompressState int32              `xorm:"DEFAULT 0"`
	Type            int                `xorm:"DEFAULT 0"`
	CreatedUnix     timeutil.TimeStamp `xorm:"created"`
	Path            string             `xorm:"varchar(400)"`
	UnzipPath       string             `xorm:"varchar(400)"`

	FileChunk            *FileChunk `xorm:"-"`
	CanDel               bool       `xorm:"-"`
	Uploader             *User      `xorm:"-"`
	Md5                  string     `xorm:"-"`
	Dataset              *Dataset   `xorm:"-"`
	DatasetMigrateStatus int        `xorm:"-"` // 0 不可升级 1 未升级 2 已升级 9 升级中
	NewDatasetName       string     `xorm:"-"`
	NewOwnerName         string     `xorm:"-"`
}

type AttachmentUsername struct {
	Attachment `xorm:"extends"`
	Name       string
}

type SearchDatasetAttachmentOptions struct {
	UploaderID int64
	ListOptions
	OrderBy string
}

type GetDatasetSummaryOptions struct {
	UploaderID int64
	ListOptions
	OrderBy string
}

func (a *Attachment) AfterUpdate() {
	if a.DatasetID > 0 {
		datasetIsPublicCount, err := x.Where("dataset_id = ? AND is_private = ?", a.DatasetID, false).Count(new(Attachment))
		if err != nil {
			return
		}
		if datasetIsPublicCount > 0 {
			x.Table(new(Dataset)).ID(a.DatasetID).Update(map[string]interface{}{"status": DatasetStatusPublic})
		} else {
			x.Table(new(Dataset)).ID(a.DatasetID).Update(map[string]interface{}{"status": DatasetStatusPrivate})
		}
	}
}

// IncreaseDownloadCount is update download count + 1
func (a *Attachment) IncreaseDownloadCount() error {
	// Update download count.
	if _, err := x.Exec("UPDATE `attachment` SET download_count=download_count+1 WHERE id=?", a.ID); err != nil {
		return fmt.Errorf("increase attachment count: %v", err)
	}

	return nil
}

func increaseAttachmentUseNumber(uuid string) error {
	uuidArray := strings.Split(uuid, ";")

	if len(uuidArray) == 0 {
		return nil
	}

	_, err := x.Table("attachment").
		Where(builder.In("uuid", uuidArray)).
		Incr("use_number", 1).
		Update(&Attachment{})

	if err != nil {
		return fmt.Errorf("increase attachment use count: %v", err)
	}

	return nil
}

func (a *Attachment) UpdateDatasetUpdateUnix() error {
	// Update download count.
	if _, err := x.Exec("UPDATE `dataset` SET updated_unix="+fmt.Sprint(time.Now().Unix())+" WHERE id=?", a.DatasetID); err != nil {
		return fmt.Errorf("UpdateDatasetUpdateUnix: %v", err)
	}
	return nil
}

// APIFormat converts models.Attachment to api.Attachment
func (a *Attachment) APIFormat() *api.Attachment {
	return &api.Attachment{
		ID:            a.ID,
		Name:          a.Name,
		Created:       a.CreatedUnix.AsTime(),
		DownloadCount: a.DownloadCount,
		Size:          a.Size,
		UUID:          a.UUID,
		DownloadURL:   a.DownloadURL(),
		S3DownloadURL: a.S3DownloadURL(),
	}
}

// DownloadURL returns the download url of the attached file
func (a *Attachment) DownloadURL() string {
	return fmt.Sprintf("%sattachments/%s?type=%d", setting.AppURL, a.UUID, a.Type)
}

// S3DownloadURL returns the s3 download url of the attached file
func (a *Attachment) S3DownloadURL() string {
	url := ""
	if a.Type == TypeCloudBrainOne {
		url, _ = storage.Attachments.PresignedGetURL(a.Path, a.Name)
	} else if a.Type == TypeCloudBrainTwo {
		url, _ = storage.ObsGetPreSignedUrl(a.Path, a.Name)
	}
	return url
}

// AttachmentRelativePath returns the relative path
func AttachmentRelativePath(uuid string) string {
	return path.Join(uuid[0:1], uuid[1:2], uuid)
}

// RelativePath returns the relative path of the attachment
func (a *Attachment) RelativePath() string {
	return AttachmentRelativePath(a.UUID)
}

func (a *Attachment) RealPath() string {
	if a.Path != "" {
		return a.Path
	}
	if a.Type == TypeCloudBrainOne {
		return setting.Attachment.Minio.BasePath + a.RelativePath()
	} else {
		return setting.BasePath + a.RelativePath()
	}
}

// LinkedRepository returns the linked repo if any
func (a *Attachment) LinkedRepository() (*Repository, UnitType, error) {
	if a.IssueID != 0 {
		iss, err := GetIssueByID(a.IssueID)
		if err != nil {
			return nil, UnitTypeIssues, err
		}
		repo, err := GetRepositoryByID(iss.RepoID)
		unitType := UnitTypeIssues
		if iss.IsPull {
			unitType = UnitTypePullRequests
		}
		return repo, unitType, err
	} else if a.ReleaseID != 0 {
		rel, err := GetReleaseByID(a.ReleaseID)
		if err != nil {
			return nil, UnitTypeReleases, err
		}
		repo, err := GetRepositoryByID(rel.RepoID)
		return repo, UnitTypeReleases, err
	}
	return nil, -1, nil
}

// NewAttachment creates a new attachment object.
func NewAttachment(attach *Attachment, buf []byte, file io.Reader) (_ *Attachment, err error) {
	attach.UUID = gouuid.NewV4().String()

	size, err := storage.Attachments.Save(attach.RelativePath(), io.MultiReader(bytes.NewReader(buf), file))
	if err != nil {
		return nil, fmt.Errorf("Create: %v", err)
	}
	attach.Size = size
	attach.Path = setting.Attachment.Minio.BasePath + attach.RelativePath()

	if _, err := x.Insert(attach); err != nil {
		return nil, err
	}

	return attach, nil
}

// GetAttachmentByID returns attachment by given id
func GetAttachmentByID(id int64) (*Attachment, error) {
	return getAttachmentByID(x, id)
}

func getAttachmentByID(e Engine, id int64) (*Attachment, error) {
	attach := new(Attachment)

	if has, err := e.Where("id = ?", id).Get(attach); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrAttachmentNotExist{ID: id, UUID: ""}
	}
	return attach, nil
}

func getAttachmentByUUID(e Engine, uuid string) (*Attachment, error) {
	attach := new(Attachment)
	has, err := e.Where("uuid = ?", uuid).Get(attach)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrAttachmentNotExist{0, uuid}
	}
	return attach, nil
}

// GetAttachmentsByUUIDs returns attachment by given UUID list.
func GetAttachmentsByUUIDs(uuids []string) ([]*Attachment, error) {
	return getAttachmentsByUUIDs(x, uuids)
}

func getAttachmentsByUUIDs(e Engine, uuids []string) ([]*Attachment, error) {
	if len(uuids) == 0 {
		return []*Attachment{}, nil
	}

	// Silently drop invalid uuids.
	attachments := make([]*Attachment, 0, len(uuids))
	return attachments, e.In("uuid", uuids).Find(&attachments)
}

func GetUsedDatasetSizeByUser(userId int64) (int64, error) {

	total, err := x.Where("uploader_id = ? and dataset_id!=0", userId).Sum(new(Attachment), "size")

	return int64(total), err

}

// GetAttachmentByUUID returns attachment by given UUID.
func GetAttachmentByUUID(uuid string) (*Attachment, error) {
	return getAttachmentByUUID(x, uuid)
}

// GetAttachmentByReleaseIDFileName returns attachment by given releaseId and fileName.
func GetAttachmentByReleaseIDFileName(releaseID int64, fileName string) (*Attachment, error) {
	return getAttachmentByReleaseIDFileName(x, releaseID, fileName)
}

func getAttachmentsByIssueID(e Engine, issueID int64) ([]*Attachment, error) {
	attachments := make([]*Attachment, 0, 10)
	return attachments, e.Where("issue_id = ? AND comment_id = 0", issueID).Find(&attachments)
}

// GetAttachmentsByIssueID returns all attachments of an issue.
func GetAttachmentsByIssueID(issueID int64) ([]*Attachment, error) {
	return getAttachmentsByIssueID(x, issueID)
}

// GetAttachmentsByCommentID returns all attachments if comment by given ID.
func GetAttachmentsByCommentID(commentID int64) ([]*Attachment, error) {
	return getAttachmentsByCommentID(x, commentID)
}

func GetAttachmentByDatasetIdFileName(fileName string, datasetId int64) (*Attachment, error) {
	attach := &Attachment{DatasetID: datasetId, Name: fileName}
	has, err := x.Get(attach)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, err
	}
	return attach, nil
}

func getAttachmentsByCommentID(e Engine, commentID int64) ([]*Attachment, error) {
	attachments := make([]*Attachment, 0, 10)
	return attachments, e.Where("comment_id=?", commentID).Find(&attachments)
}

func GetAttachmentsByPath(path string) ([]*Attachment, error) {
	return getAttachmentsByPath(x, path)
}

func getAttachmentsByPath(e Engine, path string) ([]*Attachment, error) {
	attachments := make([]*Attachment, 0, 10)
	return attachments, e.Where("path=?", path).Find(&attachments)
}

// getAttachmentByReleaseIDFileName return a file based on the the following infos:
func getAttachmentByReleaseIDFileName(e Engine, releaseID int64, fileName string) (*Attachment, error) {
	attach := &Attachment{ReleaseID: releaseID, Name: fileName}
	has, err := e.Get(attach)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, err
	}
	return attach, nil
}

// DeleteAttachment deletes the given attachment and optionally the associated file.
func DeleteAttachment(a *Attachment, remove bool) error {
	_, err := DeleteAttachments([]*Attachment{a}, remove)
	return err
}

// DeleteAttachment deletes the given attachment and optionally the associated file.
func DeleteSingleAttachment(a *Attachment) error {

	if a.Type == TypeCloudBrainOne {
		if err := storage.Attachments.Delete(a.RelativePath()); err != nil {
			log.Info("Minio Message:%s\n", err.Error())
		}
	}
	if a.Type == TypeCloudBrainTwo {
		input := &obs.DeleteObjectInput{}
		input.Bucket = setting.Bucket
		input.Key = a.Path
		log.Info("delete obs file:" + input.Key)
		output, err := storage.ObsCli.DeleteObject(input)
		if err == nil {
			log.Info("RequestId:%s\n", output.RequestId)
		} else if obsError, ok := err.(obs.ObsError); ok {
			log.Info("Code:%s\n", obsError.Code)
			log.Info("Message:%s\n", obsError.Message)
		}
	}

	return nil
}

func IsNeedToDeleteAttachPath(attachment *Attachment) bool {
	objectKey := "attachment/" + attachment.UUID[0:1] + "/" + attachment.UUID[1:2] + "/" + attachment.UUID + "/" + attachment.Name
	if attachment.Type == TypeCloudBrainOne {
		objectKey = "attachments/" + attachment.UUID[0:1] + "/" + attachment.UUID[1:2] + "/" + attachment.UUID
	}
	if objectKey == attachment.Path {
		attachments, err := GetAttachmentsByPath(objectKey)
		if err == nil {
			if len(attachments) > 1 { //有其它记录在使用此文件，不能删除。
				log.Info("Not need to delete the attach, as it has been used.1")
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	} else { //复用的其它记录文件路径，因此不能实际删除文件
		log.Info("Not need to delete the attach, as it has been used.2")
		//这里要再判断一下，可能只剩下它在使用了
		attachments, err := GetAttachmentsByPath(attachment.Path)
		if err == nil {
			if len(attachments) <= 1 { //
				return true
			}
		}
		return false
	}
}

// DeleteAttachments deletes the given attachments and optionally the associated files.
func DeleteAttachments(attachments []*Attachment, remove bool) (int, error) {
	if len(attachments) == 0 {
		return 0, nil
	}

	if remove {
		for _, a := range attachments {
			if IsNeedToDeleteAttachPath(a) {
				DeleteSingleAttachment(a)
				if len(a.UnzipPath) > 0 {
					log.Info("start to delete Decompress path=" + a.UnzipPath)
					if a.Type == TypeCloudBrainOne {
						storage.Attachments.DeleteDir(a.UnzipPath)
					}
					if a.Type == TypeCloudBrainTwo {
						err := storage.ObsRemoveObject(setting.Bucket, a.UnzipPath)
						if err != nil {
							log.Info("delete unzip file error.")
						}
					}
				}
			}
			DeleteAttachmentByID(a.ID)
			DeleteFileChunkById(a.UUID)
		}
	}

	return len(attachments), nil
}

func DeleteAttachmentByID(id int64) error {
	_, err := x.Where("id = ?", id).Delete(&Attachment{})
	return err
}

// DeleteAttachmentsByIssue deletes all attachments associated with the given issue.
func DeleteAttachmentsByIssue(issueID int64, remove bool) (int, error) {
	attachments, err := GetAttachmentsByIssueID(issueID)

	if err != nil {
		return 0, err
	}

	return DeleteAttachments(attachments, remove)
}

// DeleteAttachmentsByComment deletes all attachments associated with the given comment.
func DeleteAttachmentsByComment(commentID int64, remove bool) (int, error) {
	attachments, err := GetAttachmentsByCommentID(commentID)

	if err != nil {
		return 0, err
	}

	return DeleteAttachments(attachments, remove)
}

func UpdateRepositoryAndAttachment(repo *Repository, atta *Attachment) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if err = updateRepository(sess, repo, true); err != nil {
		return fmt.Errorf("updateRepository: %v", err)
	}
	if err = updateAttachment(sess, atta); err != nil {
		return fmt.Errorf("updateAttachment: %v", err)
	}

	return sess.Commit()
}

// UpdateAttachment updates the given attachment in database
func UpdateAttachment(atta *Attachment) error {
	return updateAttachment(x, atta)
}
func UpdateAttachmentDescription(atta *Attachment) error {
	return updateAttachmentDescription(x, atta)
}

func updateAttachmentDescription(e Engine, atta *Attachment) error {
	var sess *xorm.Session

	sess = e.ID(atta.ID)

	_, err := sess.Cols("description").Update(atta)
	return err
}

func updateAttachment(e Engine, atta *Attachment) error {
	var sess *xorm.Session
	if atta.ID != 0 && atta.UUID == "" {
		sess = e.ID(atta.ID)
	} else {
		// Use uuid only if id is not set and uuid is set
		sess = e.Where("uuid = ?", atta.UUID)
	}
	_, err := sess.Cols("name", "issue_id", "release_id", "comment_id", "download_count", "is_private", "decompress_state").Update(atta)
	return err
}

// DeleteAttachmentsByRelease deletes all attachments associated with the given release.
func DeleteAttachmentsByRelease(releaseID int64) error {
	_, err := x.Where("release_id = ?", releaseID).Delete(&Attachment{})
	return err
}

// IterateAttachment iterates attachments
func IterateAttachment(f func(attach *Attachment) error) error {
	var start int
	const batchSize = 100
	for {
		var attachments = make([]*Attachment, 0, batchSize)
		if err := x.Limit(batchSize, start).Find(&attachments); err != nil {
			return err
		}
		if len(attachments) == 0 {
			return nil
		}
		start += len(attachments)

		for _, attach := range attachments {
			if err := f(attach); err != nil {
				return err
			}
		}
	}
}

// LinkedDataSet returns the linked data_set if any
func (a *Attachment) LinkedDataSet() (*Dataset, error) {
	if a.DatasetID != 0 {
		return GetDatasetByID(a.DatasetID)
	}
	return nil, nil
}

// InsertAttachment insert a record into attachment.
func InsertAttachment(attach *Attachment) (_ *Attachment, err error) {

	if _, err := x.Insert(attach); err != nil {
		return nil, err
	}

	return attach, nil
}

// GetUnDecompressAttachments query the attachments unDecompressed
func GetUnDecompressAttachments() ([]*Attachment, error) {
	return getUnDecompressAttachments(x)
}

func getUnDecompressAttachments(e Engine) ([]*Attachment, error) {
	attachments := make([]*Attachment, 0, 10)
	return attachments, e.Where("decompress_state = ? and dataset_id != 0 and (name like '%.zip' or name like '%.tar.gz' or name like '%.tgz')", DecompressStateInit).Find(&attachments)
}

func GetAllPublicAttachments() ([]*AttachmentUsername, error) {
	return getAllPublicAttachments(x)
}

func getAllPublicAttachments(e Engine) ([]*AttachmentUsername, error) {
	attachments := make([]*AttachmentUsername, 0, 10)
	if err := e.Table("attachment").Join("LEFT", "`user`", "attachment.uploader_id "+
		"= `user`.id").Where("decompress_state= ? and is_private= ? and attachment.type = ?", DecompressStateDone, false, TypeCloudBrainOne).Find(&attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

func GetPrivateAttachments(username string) ([]*AttachmentUsername, error) {
	user, err := getUserByName(x, username)
	if err != nil {
		log.Error("getUserByName(%s) failed:%v", username, err)
		return nil, err
	}
	return getPrivateAttachments(x, user.ID)
}

func getPrivateAttachments(e Engine, userID int64) ([]*AttachmentUsername, error) {
	attachments := make([]*AttachmentUsername, 0, 10)
	if err := e.Table("attachment").Join("LEFT", "`user`", "attachment.uploader_id "+
		"= `user`.id").Where("decompress_state= ? and uploader_id= ? and attachment.type = ?", DecompressStateDone, userID, TypeCloudBrainOne).Find(&attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

func getModelArtsUserAttachments(e Engine, userID int64) ([]*AttachmentUsername, error) {
	attachments := make([]*AttachmentUsername, 0, 10)
	if err := e.Table("attachment").Join("LEFT", "`user`", "attachment.uploader_id "+
		"= `user`.id").Where("attachment.type = ? and (uploader_id= ? or is_private = ?)", TypeCloudBrainTwo, userID, false).Find(&attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

func GetModelArtsUserAttachments(userID int64) ([]*AttachmentUsername, error) {
	return getModelArtsUserAttachments(x, userID)
}

func getModelArtsTrainAttachments(e Engine, userID int64) ([]*AttachmentUsername, error) {
	attachments := make([]*AttachmentUsername, 0, 10)
	if err := e.Table("attachment").Join("LEFT", "`user`", "attachment.uploader_id "+
		"= `user`.id").Where("attachment.type = ? and (uploader_id= ? or is_private = ?) and attachment.decompress_state = ?", TypeCloudBrainTwo, userID, false, DecompressStateDone).Find(&attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

func GetModelArtsTrainAttachments(userID int64) ([]*AttachmentUsername, error) {
	return getModelArtsTrainAttachments(x, userID)
}

func CanDelAttachment(isSigned bool, user *User, attach *Attachment) bool {
	if !isSigned {
		return false
	}
	dataset, err := GetDatasetByID(attach.DatasetID)
	if err != nil {
		log.Error("GetDatasetByID failed:%v", err.Error())
		return false
	}
	repo, _ := GetRepositoryByID(dataset.RepoID)
	if err != nil {
		log.Error("GetRepositoryByID failed:%v", err.Error())
		return false
	}
	permission, _ := GetUserRepoPermission(repo, user)
	if err != nil {
		log.Error("GetUserRepoPermission failed:%v", err.Error())
		return false
	}

	if user.ID == attach.UploaderID || user.IsAdmin || permission.AccessMode >= AccessModeAdmin {
		return true
	}
	return false
}

func GetAttachmentSizeByDatasetID(datasetID int64) (int64, error) {
	total, err := x.Where("dataset_id = ?", datasetID).SumInt(&Attachment{}, "size")
	if err != nil {
		return 0, err
	}

	return total, nil
}

func AttachmentsByDatasetOption(datasets []int64, opts *SearchDatasetOptions) ([]*Attachment, error) {
	sess := x.NewSession()
	defer sess.Close()
	var cond = builder.NewCond()
	cond = cond.And(builder.In("attachment.dataset_id", datasets))
	if opts.JustNeedZipFile {
		cond = cond.And(builder.Gt{"attachment.decompress_state": 0})
	}
	if opts.PublicOnly {
		cond = cond.And(builder.Eq{"attachment.is_private": false})
	}
	if opts.CloudBrainType >= 0 {
		cond = cond.And(builder.Eq{"attachment.type": opts.CloudBrainType})
	}
	if opts.UploadAttachmentByMe {
		cond = cond.And(
			builder.Eq{"attachment.uploader_id": opts.User.ID},
		)
	}

	attachments := make([]*Attachment, 0)
	if err := sess.Table(&Attachment{}).Where(cond).Desc("id").
		Find(&attachments); err != nil {
		return nil, fmt.Errorf("Find: %v", err)
	}
	return attachments, nil

}

func GetAllAttachmentSize() (int64, error) {
	return x.SumInt(&Attachment{}, "size")
}

func GetAllDatasetContributorByDatasetId(datasetId int64) ([]*User, error) {
	r := make([]*User, 0)
	if err := x.Select("distinct(public.user.*)").Table("attachment").Join("LEFT", "user", "public.user.ID = attachment.uploader_id").Where("attachment.dataset_id = ?", datasetId).Find(&r); err != nil {
		return nil, err
	}
	return r, nil
}

type DatasetAttachmentList []*Attachment

func (attachments DatasetAttachmentList) loadAttachmentAttributes(e Engine) error {

	for _, attachment := range attachments {
		dataset := new(Dataset)

		has, err := e.ID(attachment.DatasetID).Get(dataset)

		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("Get Dataset by id %d error", attachment.DatasetID)
		}
		attachment.Dataset = dataset
		if attachment.Dataset != nil {
			repo := new(Repository)
			has, err = e.ID(attachment.Dataset.RepoID).Get(repo)
			if err != nil {
				return err
			}
			if !has {
				return fmt.Errorf("Get Repository by id %d error", attachment.Dataset.RepoID)
			}
			attachment.Dataset.Repo = repo

		}

	}

	return nil
}
func SearchDatasetAttachments(opts *SearchDatasetAttachmentOptions) (DatasetAttachmentList, int64, error) {
	cond := searchDatasetAttachmentCondition(opts)
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 {
		opts.PageSize = 10
	}

	var err error
	sess := x.NewSession()
	defer sess.Close()

	datasetAttachments := make(DatasetAttachmentList, 0, opts.PageSize)

	count, err := sess.Where(cond).Count(new(Attachment))
	if err != nil {
		return nil, 0, err
	}

	orderBy := "dataset_id,id desc"
	if opts.OrderBy == "size" {
		orderBy = "size desc"
	} else if opts.OrderBy == "size_asc" {
		orderBy = "size"
	}

	if err := sess.Where(cond).OrderBy(orderBy).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&datasetAttachments); err != nil {
		return nil, 0, err
	}

	datasetAttachments.loadAttachmentAttributes(sess)
	return datasetAttachments, count, err

}

func searchDatasetAttachmentCondition(opts *SearchDatasetAttachmentOptions) builder.Cond {
	var cond = builder.NewCond()
	//dataset id !=0 才是数据集文件
	cond = cond.And(builder.Neq{"dataset_id": 0})

	if opts.UploaderID > 0 {
		cond = cond.And(builder.Eq{"uploader_id": opts.UploaderID})
	}

	return cond
}

func GetTemporaryAttachments() ([]*Attachment, error) {
	attachments := make([]*Attachment, 0)
	if err := x.Where("issue_id=0 and dataset_id=0 and release_id=0 and comment_id=0 and created_unix < ?", time.Now().Add(-time.Hour*24).Unix()).Find(&attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}
