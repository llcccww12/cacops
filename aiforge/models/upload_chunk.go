package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	DataFileNotUploaded int = iota
	DataFileUploaded
)

type UploadChunk struct {
	ID             int64  `xorm:"pk autoincr"`
	UUID           string `xorm:"INDEX"`
	Md5            string `xorm:"INDEX"`
	SubjectID      string `xorm:"INDEX"`
	SubjectType    int
	ObjectName     string `xorm:"DEFAULT ''"`
	IsUploaded     int    `xorm:"DEFAULT 0"` // not uploaded: 0, uploaded: 1
	UploadID       string `xorm:"UNIQUE"`    //minio upload id
	TotalChunks    int
	Size           int64
	UserID         int64              `xorm:"INDEX"`
	StorageType    string             `xorm:"INDEX"`
	CompletedParts []string           `xorm:"DEFAULT ''"` // chunkNumber+etag eg: ,1-asqwewqe21312312.2-123hjkas
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`
}

func GetUploadChunkByFileNameAndUser(filename string, userID int64, storageType string, subjectId string) (*UploadChunk, error) {
	fileChunk := new(UploadChunk)

	if has, err := x.Where("LOWER(object_name) = LOWER(?) and user_id = ? and storage_type = ? and subject_id = ?", filename, userID, storageType, subjectId).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{filename, ""}
	}
	return fileChunk, nil
}

func DeleteUploadChunk(fileChunk *UploadChunk) error {
	_, err := x.ID(fileChunk.ID).Delete(fileChunk)
	return err
}

func GetUploadChunkByMD5AndUser(md5 string, userID int64, storageType string, dataId string) (*UploadChunk, error) {
	fileChunk := new(UploadChunk)

	if has, err := x.Where("md5 = ? and user_id = ? and storage_type = ? storage_type = ?", md5, userID, storageType, dataId).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{md5, ""}
	}
	return fileChunk, nil
}

func UpdateChunkUploadedStatusByUUID(isUploaded int, uuid string) error {
	f := &UploadChunk{
		IsUploaded: isUploaded,
	}
	_, err := x.Cols("is_uploaded").Where("uuid = ?", uuid).Update(f)
	return err
}

func InsertUploadChunk(fileChunk *UploadChunk) (_ *UploadChunk, err error) {
	if _, err := x.Insert(fileChunk); err != nil {
		return nil, err
	}
	return fileChunk, nil
}

func GetUploadChunkByUUID(uuid string) (*UploadChunk, error) {
	fileChunk := new(UploadChunk)

	if has, err := x.Where("uuid = ?", uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{"", uuid}
	}
	return fileChunk, nil
}

func GetUploadChunksByUserId(userId int64, lastTime int64, isUploadFinished bool) ([]*UploadChunk, error) {
	return getUploadChunksByUserId(x, userId, lastTime, isUploadFinished)
}

func getUploadChunksByUserId(e Engine, userId int64, lastTime int64, isUploadFinished bool) ([]*UploadChunk, error) {
	fileChunks := make([]*UploadChunk, 0)
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"user_id": userId})
	if lastTime > 0 {
		cond = cond.And(builder.Gte{"created_unix": lastTime})
	}
	if !isUploadFinished {
		cond = cond.And(builder.Eq{"is_uploaded": 0})
	}
	if err := e.Where(cond).Find(&fileChunks); err != nil {
		return nil, err
	}
	return fileChunks, nil
}
