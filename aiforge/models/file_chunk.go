package models

import (
	"fmt"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	FileNotUploaded int = iota
	FileUploaded
)

type FileChunk struct {
	ID             int64  `xorm:"pk autoincr"`
	UUID           string `xorm:"uuid UNIQUE"`
	Md5            string `xorm:"INDEX"`
	IsUploaded     int    `xorm:"DEFAULT 0"` // not uploaded: 0, uploaded: 1
	UploadID       string
	TotalChunks    int
	Size           int64
	UserID         int64              `xorm:"INDEX"`
	Type           int                `xorm:"INDEX DEFAULT 0"`
	CompletedParts []string           `xorm:"DEFAULT ''"` // chunkNumber+etag eg: ,1-asqwewqe21312312.2-123hjkas
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`
}

type ModelFileChunk struct {
	ID             int64  `xorm:"pk autoincr"`
	UUID           string `xorm:"INDEX"`
	Md5            string `xorm:"INDEX"`
	ModelUUID      string `xorm:"INDEX"`
	ObjectName     string `xorm:"DEFAULT ''"`
	IsUploaded     int    `xorm:"DEFAULT 0"` // not uploaded: 0, uploaded: 1
	UploadID       string `xorm:"UNIQUE"`    //minio upload id
	TotalChunks    int
	Size           int64
	UserID         int64              `xorm:"INDEX"`
	Type           int                `xorm:"INDEX DEFAULT 0"`
	CompletedParts []string           `xorm:"DEFAULT ''"` // chunkNumber+etag eg: ,1-asqwewqe21312312.2-123hjkas
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`
}

// GetFileChunkByMD5 returns fileChunk by given id
func GetFileChunkByMD5(md5 string) (*FileChunk, error) {
	return getFileChunkByMD5(x, md5)
}

func getFileChunkByMD5(e Engine, md5 string) (*FileChunk, error) {
	fileChunk := new(FileChunk)

	if has, err := e.Where("md5 = ?", md5).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{md5, ""}
	}
	return fileChunk, nil
}

func GetFileChunkByMD5AndType(md5 string, storageType int, size int64) ([]*FileChunk, error) {
	return getFileChunkByMD5AndType(x, md5, storageType, size)
}

func getFileChunkByMD5AndType(e Engine, md5 string, storageType int, size int64) ([]*FileChunk, error) {
	fileChunks := make([]*FileChunk, 0)
	if err := e.Where("md5 = ? and  type = ? and size=?", md5, storageType, size).OrderBy("id desc").Find(&fileChunks); err != nil {
		return nil, err
	}
	return fileChunks, nil
}

// GetFileChunkByMD5 returns fileChunk by given id
func GetFileChunkByMD5AndUser(md5 string, userID int64, storageType int) (*FileChunk, error) {
	return getFileChunkByMD5AndUser(x, md5, userID, storageType)
}

func GetModelFileChunkByMD5AndUser(md5 string, userID int64, storageType int, uuid string) (*ModelFileChunk, error) {
	return getModelFileChunkByMD5AndUser(x, md5, userID, storageType, uuid)
}

func GetModelFileChunkByFileNameAndUser(md5 string, userID int64, storageType int, uuid string) (*ModelFileChunk, error) {
	return getModelFileChunkByFileNameAndUser(x, md5, userID, storageType, uuid)
}

func getModelFileChunkByFileNameAndUser(e Engine, filename string, userID int64, storageType int, uuid string) (*ModelFileChunk, error) {
	fileChunk := new(ModelFileChunk)

	if has, err := e.Where("object_name = ? and user_id = ? and type = ? and model_uuid= ?", filename, userID, storageType, uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{filename, ""}
	}
	return fileChunk, nil
}

func getModelFileChunkByMD5AndUser(e Engine, md5 string, userID int64, storageType int, uuid string) (*ModelFileChunk, error) {
	fileChunk := new(ModelFileChunk)

	if has, err := e.Where("md5 = ? and user_id = ? and type = ? and model_uuid= ?", md5, userID, storageType, uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{md5, ""}
	}
	return fileChunk, nil
}

func getFileChunkByMD5AndUser(e Engine, md5 string, userID int64, storageType int) (*FileChunk, error) {
	fileChunk := new(FileChunk)

	if has, err := e.Where("md5 = ? and user_id = ? and type = ?", md5, userID, storageType).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{md5, ""}
	}
	return fileChunk, nil
}

// GetAttachmentByID returns attachment by given id
func GetFileChunksByUserId(userId int64, lastTime int64, isUploadFinished bool) ([]*FileChunk, error) {
	return getFileChunksByUserId(x, userId, lastTime, isUploadFinished)
}

func getFileChunksByUserId(e Engine, userId int64, lastTime int64, isUploadFinished bool) ([]*FileChunk, error) {
	fileChunks := make([]*FileChunk, 0)
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

// GetAttachmentByID returns attachment by given id
func GetFileChunkByUUID(uuid string) (*FileChunk, error) {
	return getFileChunkByUUID(x, uuid)
}

func getFileChunkByUUID(e Engine, uuid string) (*FileChunk, error) {
	fileChunk := new(FileChunk)

	if has, err := e.Where("uuid = ?", uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{"", uuid}
	}
	return fileChunk, nil
}

func GetModelFileChunkByUUID(uuid string) (*ModelFileChunk, error) {
	return getModelFileChunkByUUID(x, uuid)
}

func getModelFileChunkByUUID(e Engine, uuid string) (*ModelFileChunk, error) {
	fileChunk := new(ModelFileChunk)

	if has, err := e.Where("uuid = ?", uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{"", uuid}
	}
	return fileChunk, nil
}

func GetModelFileChunksByUserId(userId int64, lastTime int64, isUploadFinished bool) ([]*ModelFileChunk, error) {
	return getModelFileChunksByUserId(x, userId, lastTime, isUploadFinished)
}

func getModelFileChunksByUserId(e Engine, userId int64, lastTime int64, isUploadFinished bool) ([]*ModelFileChunk, error) {
	fileChunks := make([]*ModelFileChunk, 0)
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

// InsertFileChunk insert a record into file_chunk.
func InsertFileChunk(fileChunk *FileChunk) (_ *FileChunk, err error) {
	if _, err := x.Insert(fileChunk); err != nil {
		return nil, err
	}

	return fileChunk, nil
}

// InsertFileChunk insert a record into file_chunk.
func InsertModelFileChunk(fileChunk *ModelFileChunk) (_ *ModelFileChunk, err error) {
	if _, err := x.Insert(fileChunk); err != nil {
		return nil, err
	}
	return fileChunk, nil
}

func DeleteFileChunkById(uuid string) (*FileChunk, error) {
	return deleteFileChunkById(x, uuid)
}

func deleteFileChunkById(e Engine, uuid string) (*FileChunk, error) {
	fileChunk := new(FileChunk)
	if has, err := e.Where("uuid = ?", uuid).Get(fileChunk); err != nil {
		return nil, err
	} else if !has {
		return nil, ErrFileChunkNotExist{"", uuid}
	}

	err := deleteFileChunk(e, fileChunk)
	log.Info("delete the filechunk,id=" + fmt.Sprint(fileChunk.ID))
	if err != nil {
		return nil, err
	} else {
		return fileChunk, nil
	}
}

func UpdateModelFileChunk(fileChunk *ModelFileChunk) error {
	return updateModelFileChunk(x, fileChunk)
}

func updateModelFileChunk(e Engine, fileChunk *ModelFileChunk) error {
	var sess *xorm.Session
	sess = e.Where("uuid = ?", fileChunk.UUID)
	_, err := sess.Cols("is_uploaded").Update(fileChunk)
	return err
}

// UpdateFileChunk updates the given file_chunk in database
func UpdateFileChunk(fileChunk *FileChunk) error {
	return updateFileChunk(x, fileChunk)
}

func updateFileChunk(e Engine, fileChunk *FileChunk) error {
	var sess *xorm.Session
	sess = e.Where("uuid = ?", fileChunk.UUID)
	_, err := sess.Cols("is_uploaded").Update(fileChunk)
	return err
}

// DeleteFileChunk delete the given file_chunk in database
func DeleteFileChunk(fileChunk *FileChunk) error {
	return deleteFileChunk(x, fileChunk)
}

func deleteFileChunk(e Engine, fileChunk *FileChunk) error {
	_, err := e.ID(fileChunk.ID).Delete(fileChunk)
	return err
}

func DeleteModelFileChunk(fileChunk *ModelFileChunk) error {
	return deleteModelFileChunk(x, fileChunk)
}

func deleteModelFileChunk(e Engine, fileChunk *ModelFileChunk) error {
	_, err := e.ID(fileChunk.ID).Delete(fileChunk)
	return err
}
