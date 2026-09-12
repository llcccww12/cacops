package models

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	DatasetStatusPrivate int32 = iota
	DatasetStatusPublic
	DatasetStatusDeleted
)
const (
	StorageTypeMinio int = iota
	StorageTypeObs
)

type Dataset struct {
	ID            int64  `xorm:"pk autoincr"`
	Title         string `xorm:"INDEX NOT NULL""`
	Status        int32  `xorm:"INDEX""` // normal_private: 0, pulbic: 1, is_delete: 2
	Category      string
	Description   string `xorm:"TEXT"`
	DownloadTimes int64
	UseCount      int64 `xorm:"DEFAULT 0"`
	NumStars      int   `xorm:"INDEX NOT NULL DEFAULT 0"`
	Recommend     bool  `xorm:"INDEX NOT NULL DEFAULT false"`
	License       string
	Task          string
	ReleaseID     int64              `xorm:"INDEX"`
	UserID        int64              `xorm:"INDEX"`
	RepoID        int64              `xorm:"INDEX"`
	Repo          *Repository        `xorm:"-"`
	CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"INDEX updated"`

	User        *User         `xorm:"-"`
	Attachments []*Attachment `xorm:"-"`
}

type DatasetWithStar struct {
	Dataset
	IsStaring bool
}

func (d *Dataset) IsPrivate() bool {
	switch d.Status {
	case DatasetStatusPrivate:
		return true
	case DatasetStatusPublic:
		return false
	case DatasetStatusDeleted:
		return false
	default:
		return false
	}
}

type DatasetList []*Dataset

func (datasets DatasetList) loadAttributes(e Engine) error {
	if len(datasets) == 0 {
		return nil
	}

	set := make(map[int64]struct{})
	userIdSet := make(map[int64]struct{})
	datasetIDs := make([]int64, len(datasets))
	for i := range datasets {
		userIdSet[datasets[i].UserID] = struct{}{}
		set[datasets[i].RepoID] = struct{}{}
		datasetIDs[i] = datasets[i].ID
	}

	// Load owners.
	users := make(map[int64]*User, len(userIdSet))
	repos := make(map[int64]*Repository, len(set))
	if err := e.
		Where("id > 0").
		In("id", keysInt64(userIdSet)).
		Cols("id", "lower_name", "name", "full_name", "email").
		Find(&users); err != nil {
		return fmt.Errorf("find users: %v", err)
	}
	if err := e.
		Where("id > 0").
		In("id", keysInt64(set)).
		Cols("id", "owner_id", "owner_name", "lower_name", "name", "description", "alias", "lower_alias", "is_private").
		Find(&repos); err != nil {
		return fmt.Errorf("find repos: %v", err)
	}
	for i := range datasets {
		datasets[i].User = users[datasets[i].UserID]
		datasets[i].Repo = repos[datasets[i].RepoID]
	}

	return nil
}

func (datasets DatasetList) loadAttachmentAttributes(opts *SearchDatasetOptions) error {
	if len(datasets) == 0 {
		return nil
	}
	datasetIDs := make([]int64, len(datasets))
	for i := range datasets {
		datasetIDs[i] = datasets[i].ID
	}
	attachments, err := AttachmentsByDatasetOption(datasetIDs, opts)
	if err != nil {
		return fmt.Errorf("GetAttachmentsByDatasetIds failed error: %v", err)
	}

	permissionMap := make(map[int64]bool, len(datasets))

	for _, attachment := range attachments {

		for i := range datasets {
			if attachment.DatasetID == datasets[i].ID {

				if !attachment.IsPrivate {
					datasets[i].Attachments = append(datasets[i].Attachments, attachment)
				} else {
					permission, ok := permissionMap[datasets[i].ID]
					if !ok {

						permission = false
						datasets[i].Repo.GetOwner()
						if !permission {
							if datasets[i].Repo.OwnerID == opts.User.ID {
								permission = true
							} else {
								isCollaborator, _ := datasets[i].Repo.IsCollaborator(opts.User.ID)
								isInRepoTeam, _ := datasets[i].Repo.IsInRepoTeam(opts.User.ID)

								if isCollaborator || isInRepoTeam {
									permission = true
								}
							}

						}

						permissionMap[datasets[i].ID] = permission
					}

					if permission {
						datasets[i].Attachments = append(datasets[i].Attachments, attachment)
					}
				}

			}

		}

	}

	for i := range datasets {
		if datasets[i].Attachments == nil {
			datasets[i].Attachments = []*Attachment{}
		}
		datasets[i].Repo.Owner = nil
	}
	return nil

}

type SearchDatasetOptions struct {
	Keyword          string
	OwnerID          int64
	User             *User
	RepoID           int64
	IncludePublic    bool
	RecommendOnly    bool
	Category         string
	Task             string
	License          string
	DatasetIDs       []int64
	ExcludeDatasetId int64
	ListOptions
	SearchOrderBy
	IsOwner              bool
	StarByMe             bool
	CloudBrainType       int //0 cloudbrain   1  modelarts   -1 all
	PublicOnly           bool
	JustNeedZipFile      bool
	NeedAttachment       bool
	UploadAttachmentByMe bool
	ExploreMine          bool
	QueryReference       bool
}

func CreateDataset(dataset *Dataset) (err error) {

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	datasetByRepoId := &Dataset{RepoID: dataset.RepoID}
	has, err := sess.Get(datasetByRepoId)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("The dataset already exists.")
	}

	if _, err = sess.Insert(dataset); err != nil {
		return err
	}
	return sess.Commit()

}

func RecommendDataset(dataSetId int64, recommend bool) error {

	dataset := Dataset{Recommend: recommend}
	_, err := x.ID(dataSetId).Cols("recommend").Update(dataset)
	return err
}

func SearchMyDataset(opts *SearchDatasetOptions) (DatasetList, int64, error) {
	cond := SearchMyDatasetCondition(opts)
	return SearchMyDatasetByCondition(opts, cond)
}

func SearchMyDatasetCondition(opts *SearchDatasetOptions) builder.Cond {
	var cond = builder.NewCond()
	cond = cond.And(builder.Neq{"dataset.status": DatasetStatusDeleted})

	cond = generateBasicFilterCond(opts, cond)

	if len(opts.DatasetIDs) > 0 && opts.UploadAttachmentByMe {
		cond = cond.And(builder.In("dataset.id", opts.DatasetIDs))
	}

	return cond
}
func SearchMyDatasetByCondition(opts *SearchDatasetOptions, cond builder.Cond) (DatasetList, int64, error) {
	if opts.Page <= 0 {
		opts.Page = 1
	}

	var err error
	sess := x.NewSession()
	defer sess.Close()

	datasets := make(DatasetList, 0, opts.PageSize)

	if len(opts.DatasetIDs) == 0 {
		return datasets, 0, nil
	}
	selectColumnsSql := "distinct dataset.id,dataset.title, dataset.status, dataset.category, dataset.description, dataset.download_times, dataset.license, dataset.task, dataset.release_id, dataset.user_id, dataset.repo_id, dataset.created_unix,dataset.updated_unix,dataset.num_stars,dataset.recommend,dataset.use_count"

	count, err := sess.Distinct("dataset.id").Join("INNER", "repository", "repository.id = dataset.repo_id").
		Where(cond).Count(new(Dataset))

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	builderQuery := builder.Dialect(setting.Database.Type).Select("id", "title", "status", "category", "description", "download_times", "license", "task", "release_id", "user_id", "repo_id", "created_unix", "updated_unix", "num_stars", "recommend", "use_count").From(builder.Dialect(setting.Database.Type).Select(selectColumnsSql).From("dataset").Join("INNER", "repository", "repository.id = dataset.repo_id").
		Where(cond), "d").OrderBy(opts.SearchOrderBy.String())

	if opts.PageSize > 0 {
		builderQuery.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize)
	}

	if err = sess.SQL(builderQuery).Find(&datasets); err != nil {
		return nil, 0, fmt.Errorf("Dataset: %v", err)
	}

	if err = datasets.loadAttributes(sess); err != nil {
		return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
	}

	if opts.NeedAttachment {
		if err = datasets.loadAttachmentAttributes(opts); err != nil {
			return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
		}
	}

	return datasets, count, nil
}

func SearchDataset(opts *SearchDatasetOptions) (DatasetList, int64, error) {
	cond := SearchDatasetCondition(opts)
	return SearchDatasetByCondition(opts, cond)
}

func SearchDatasetCondition(opts *SearchDatasetOptions) builder.Cond {
	var cond = builder.NewCond()
	cond = cond.And(builder.Neq{"dataset.status": DatasetStatusDeleted})

	cond = generateFilterCond(opts, cond)

	if opts.RepoID > 0 {
		cond = cond.And(builder.Eq{"dataset.repo_id": opts.RepoID})
	}

	if opts.ExcludeDatasetId > 0 {
		cond = cond.And(builder.Neq{"dataset.id": opts.ExcludeDatasetId})
	}

	if opts.PublicOnly {
		cond = cond.And(builder.Eq{"dataset.status": DatasetStatusPublic})
		cond = cond.And(builder.Eq{"attachment.is_private": false})
	} else if opts.IncludePublic {
		cond = cond.And(builder.Eq{"dataset.status": DatasetStatusPublic})
		cond = cond.And(builder.Eq{"attachment.is_private": false})
		if opts.OwnerID > 0 {
			subCon := builder.NewCond()
			subCon = subCon.And(builder.Eq{"repository.owner_id": opts.OwnerID})
			subCon = generateFilterCond(opts, subCon)
			cond = cond.Or(subCon)

		}
	} else if opts.OwnerID > 0 && !opts.StarByMe && !opts.UploadAttachmentByMe {
		cond = cond.And(builder.Eq{"repository.owner_id": opts.OwnerID})
		if !opts.IsOwner {
			cond = cond.And(builder.Eq{"dataset.status": DatasetStatusPublic})
			cond = cond.And(builder.Eq{"attachment.is_private": false})
		}
	}
	if len(opts.DatasetIDs) > 0 {
		if opts.StarByMe || (opts.RepoID == 0 && opts.QueryReference) {
			cond = cond.And(builder.In("dataset.id", opts.DatasetIDs))
		} else {
			subCon := builder.NewCond()
			subCon = subCon.And(builder.In("dataset.id", opts.DatasetIDs))
			subCon = generateFilterCond(opts, subCon)
			cond = cond.Or(subCon)
		}
	} else {
		if opts.StarByMe {
			cond = cond.And(builder.Eq{"dataset.id": -1})
		}
	}

	return cond
}

func generateBasicFilterCond(opts *SearchDatasetOptions, cond builder.Cond) builder.Cond {
	if len(opts.Keyword) > 0 {
		cond = cond.And(builder.Or(builder.Like{"LOWER(dataset.title)", strings.ToLower(opts.Keyword)}, builder.Like{"LOWER(dataset.description)", strings.ToLower(opts.Keyword)}))
	}

	if len(opts.Category) > 0 {
		cond = cond.And(builder.Eq{"dataset.category": opts.Category})
	}

	if len(opts.Task) > 0 {
		cond = cond.And(builder.Eq{"dataset.task": opts.Task})
	}
	if len(opts.License) > 0 {
		cond = cond.And(builder.Eq{"dataset.license": opts.License})
	}

	if opts.RecommendOnly {
		cond = cond.And(builder.Eq{"dataset.recommend": opts.RecommendOnly})
	}

	return cond
}

func generateFilterCond(opts *SearchDatasetOptions, cond builder.Cond) builder.Cond {
	if len(opts.Keyword) > 0 {
		cond = cond.And(builder.Or(builder.Like{"LOWER(dataset.title)", strings.ToLower(opts.Keyword)}, builder.Like{"LOWER(dataset.description)", strings.ToLower(opts.Keyword)}))
	}

	if len(opts.Category) > 0 {
		cond = cond.And(builder.Eq{"dataset.category": opts.Category})
	}

	if len(opts.Task) > 0 {
		cond = cond.And(builder.Eq{"dataset.task": opts.Task})
	}
	if len(opts.License) > 0 {
		cond = cond.And(builder.Eq{"dataset.license": opts.License})
	}

	if opts.RecommendOnly {
		cond = cond.And(builder.Eq{"dataset.recommend": opts.RecommendOnly})
	}

	if opts.JustNeedZipFile {
		cond = cond.And(builder.Gt{"attachment.decompress_state": 0})
	}

	if opts.CloudBrainType >= 0 {
		cond = cond.And(builder.Eq{"attachment.type": opts.CloudBrainType})
	}
	if opts.UploadAttachmentByMe {
		cond = cond.And(builder.Eq{"attachment.uploader_id": opts.User.ID})
	}

	return cond
}

func SearchDatasetByCondition(opts *SearchDatasetOptions, cond builder.Cond) (DatasetList, int64, error) {
	if opts.Page <= 0 {
		opts.Page = 1
	}

	var err error
	sess := x.NewSession()
	defer sess.Close()

	datasets := make(DatasetList, 0, opts.PageSize)
	selectColumnsSql := "distinct dataset.id,dataset.title, dataset.status, dataset.category, dataset.description, dataset.download_times, dataset.license, dataset.task, dataset.release_id, dataset.user_id, dataset.repo_id, dataset.created_unix,dataset.updated_unix,dataset.num_stars,dataset.recommend,dataset.use_count"

	count, err := sess.Distinct("dataset.id").Join("INNER", "repository", "repository.id = dataset.repo_id").
		Join("INNER", "attachment", "attachment.dataset_id=dataset.id").
		Where(cond).Count(new(Dataset))

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	builderQuery := builder.Dialect(setting.Database.Type).Select("id", "title", "status", "category", "description", "download_times", "license", "task", "release_id", "user_id", "repo_id", "created_unix", "updated_unix", "num_stars", "recommend", "use_count").From(builder.Dialect(setting.Database.Type).Select(selectColumnsSql).From("dataset").Join("INNER", "repository", "repository.id = dataset.repo_id").
		Join("INNER", "attachment", "attachment.dataset_id=dataset.id").
		Where(cond), "d").OrderBy(opts.SearchOrderBy.String())

	if opts.PageSize > 0 {
		builderQuery.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize)
	}

	if err = sess.SQL(builderQuery).Find(&datasets); err != nil {
		return nil, 0, fmt.Errorf("Dataset: %v", err)
	}

	if err = datasets.loadAttributes(sess); err != nil {
		return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
	}

	if opts.NeedAttachment {
		if err = datasets.loadAttachmentAttributes(opts); err != nil {
			return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
		}
	}

	return datasets, count, nil
}

type datasetMetaSearch struct {
	ID  []int64
	Rel []*Dataset
}

func (s datasetMetaSearch) Len() int {
	return len(s.ID)
}
func (s datasetMetaSearch) Swap(i, j int) {
	s.ID[i], s.ID[j] = s.ID[j], s.ID[i]
	s.Rel[i], s.Rel[j] = s.Rel[j], s.Rel[i]
}
func (s datasetMetaSearch) Less(i, j int) bool {
	return s.ID[i] < s.ID[j]
}

func GetDatasetAttachments(typeCloudBrain int, isSigned bool, user *User, rels ...*Dataset) (err error) {
	return getDatasetAttachments(x, typeCloudBrain, isSigned, user, rels...)
}

func getDatasetAttachments(e Engine, typeCloudBrain int, isSigned bool, user *User, rels ...*Dataset) (err error) {
	if len(rels) == 0 {
		return
	}

	// To keep this efficient as possible sort all datasets by id,
	//    select attachments by dataset id,
	//    then merge join them

	// Sort
	var sortedRels = datasetMetaSearch{ID: make([]int64, len(rels)), Rel: make([]*Dataset, len(rels))}
	var attachments []*Attachment
	for index, element := range rels {
		element.Attachments = []*Attachment{}
		sortedRels.ID[index] = element.ID
		sortedRels.Rel[index] = element
	}
	sort.Sort(sortedRels)

	// Select attachments
	if typeCloudBrain == -1 {
		err = e.
			Asc("dataset_id").
			In("dataset_id", sortedRels.ID).
			Find(&attachments, Attachment{})
		if err != nil {
			return err
		}
	} else {
		err = e.
			Asc("dataset_id").
			In("dataset_id", sortedRels.ID).
			And("type = ?", typeCloudBrain).
			Find(&attachments, Attachment{})
		if err != nil {
			return err
		}
	}

	// merge join
	var currentIndex = 0
	for _, attachment := range attachments {
		for sortedRels.ID[currentIndex] < attachment.DatasetID {
			currentIndex++
		}
		fileChunks := make([]*FileChunk, 0, 10)
		err = e.
			Where("uuid = ?", attachment.UUID).
			Find(&fileChunks)
		if err != nil {
			return err
		}
		if len(fileChunks) > 0 {
			attachment.Md5 = fileChunks[0].Md5
		} else {
			log.Error("has attachment record, but has no file_chunk record")
			attachment.Md5 = "no_record"
		}

		attachment.CanDel = CanDelAttachment(isSigned, user, attachment)
		sortedRels.Rel[currentIndex].Attachments = append(sortedRels.Rel[currentIndex].Attachments, attachment)
	}

	return
}

// AddDatasetAttachments adds a Dataset attachments
func AddDatasetAttachments(DatasetID int64, attachmentUUIDs []string) (err error) {
	// Check attachments
	attachments, err := GetAttachmentsByUUIDs(attachmentUUIDs)
	if err != nil {
		return fmt.Errorf("GetAttachmentsByUUIDs [uuids: %v]: %v", attachmentUUIDs, err)
	}

	for i := range attachments {
		attachments[i].DatasetID = DatasetID
		// No assign value could be 0, so ignore AllCols().
		if _, err = x.ID(attachments[i].ID).Update(attachments[i]); err != nil {
			return fmt.Errorf("update attachment [%d]: %v", attachments[i].ID, err)
		}
	}

	return
}

func UpdateDataset(ctx DBContext, rel *Dataset) error {
	_, err := ctx.e.ID(rel.ID).AllCols().Update(rel)
	return err
}

func increaseDatasetUseCount(uuid string) {

	increaseAttachmentUseNumber(uuid)
	attachments, _ := GetAttachmentsByUUIDs(strings.Split(uuid, ";"))

	countMap := make(map[int64]int)

	for _, attachment := range attachments {
		value, ok := countMap[attachment.DatasetID]
		if ok {
			countMap[attachment.DatasetID] = value + 1
		} else {
			countMap[attachment.DatasetID] = 1
		}

	}

	for key, value := range countMap {
		x.Exec("UPDATE `dataset` SET use_count=use_count+? WHERE id=?", value, key)
	}

}
func increaseDatasetRegistryUseCount(datasetUUIDStr string) {
	if datasetUUIDStr == "" {
		return
	}
	uuids := strings.Split(datasetUUIDStr, ";")
	x.In("id", uuids).Cols("use_count").Incr("use_count").NoAutoTime().Update(&DatasetRegistry{})

}

// GetDatasetByID returns Dataset with given ID.
func GetDatasetByID(id int64) (*Dataset, error) {
	rel := new(Dataset)
	has, err := x.
		ID(id).
		Get(rel)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrDatasetNotExist{id}
	}

	return rel, nil
}

func GetDatasetByRepo(repo *Repository) (*Dataset, error) {
	dataset := &Dataset{RepoID: repo.ID}
	has, err := x.Get(dataset)
	if err != nil {
		return nil, err
	}
	if has {
		return dataset, nil
	} else {
		return nil, ErrNotExist{repo.ID}
	}
}

func IsDatasetExistInRepo(repoID int64) bool {
	dataset := &Dataset{RepoID: repoID}
	has, _ := x.Get(dataset)
	return has
}

func GetDatasetStarByUser(user *User) ([]*DatasetStar, error) {
	datasetStars := make([]*DatasetStar, 0)
	err := x.Cols("id", "uid", "dataset_id", "created_unix").Where("uid=?", user.ID).Find(&datasetStars)
	return datasetStars, err
}

func DeleteDataset(datasetID int64, uid int64) error {
	var err error
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	dataset := &Dataset{ID: datasetID, UserID: uid}
	has, err := sess.Get(dataset)
	if err != nil {
		return err
	} else if !has {
		return errors.New("not found")
	}

	if cnt, err := sess.ID(datasetID).Delete(new(Dataset)); err != nil {
		return err
	} else if cnt != 1 {
		return errors.New("not found")
	}
	if err = sess.Commit(); err != nil {
		sess.Close()
		return fmt.Errorf("Commit: %v", err)
	}
	return nil
}

func GetOwnerDatasetByID(id int64, user *User) (*Dataset, error) {
	dataset, err := GetDatasetByID(id)
	if err != nil {
		return nil, err
	}

	if !dataset.IsPrivate() {
		return dataset, nil
	}
	if dataset.IsPrivate() && user != nil && user.ID == dataset.UserID {
		return dataset, nil
	}
	return nil, errors.New("dataset not fount")
}

func IncreaseDownloadCount(datasetID int64) error {
	// Update download count.
	if _, err := x.Exec("UPDATE `dataset` SET download_times=download_times+1 WHERE id=?", datasetID); err != nil {
		return fmt.Errorf("increase dataset count: %v", err)
	}

	return nil
}

func GetCollaboratorDatasetIdsByUserID(userID int64) []int64 {
	var datasets []int64
	_ = x.Table("dataset").Join("INNER", "collaboration", "dataset.repo_id = collaboration.repo_id and collaboration.mode>0 and collaboration.user_id=?", userID).
		Cols("dataset.id").Find(&datasets)
	return datasets
}

func GetOwnedDatasetIdsByUserID(userID int64) []int64 {
	var datasets []int64
	_ = x.Table("dataset").Join("INNER", "repository", "(dataset.repo_id = repository.id  and repository.owner_id=?)", userID).
		Cols("dataset.id").Find(&datasets)
	var datasets1 []int64
	_ = x.Table("dataset").Where("dataset.user_id=?", userID).
		Cols("dataset.id").Find(&datasets1)
	for _, datasetId := range datasets1 {
		if !containsInt64(datasets, datasetId) {
			datasets = append(datasets, datasetId)
		}
	}
	return datasets
}
func containsInt64(s []int64, t int64) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}

	return false
}

func GetTeamDatasetIdsByUserID(userID int64) []int64 {
	var datasets []int64
	_ = x.Table("dataset").Join("INNER", "team_repo", "dataset.repo_id = team_repo.repo_id").
		Join("INNER", "team_user", "team_repo.team_id=team_user.team_id and team_user.uid=?", userID).
		Cols("dataset.id").Find(&datasets)
	return datasets
}

func UpdateDatasetCreateUser(ID int64, user *User) error {
	_, err := x.Where("id = ?", ID).Cols("user_id").Update(&Dataset{
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func QueryDatasetGroupByTask(num int) ([]map[string]interface{}, error) {
	rows, err := x.QueryInterface(fmt.Sprintf("SELECT count(*) as total,task FROM public.dataset where task <>'' group by task order by total desc limit %d", num))
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetOwnedPublicDatasetsByUserID(userID int64) DatasetList {
	datasets := make(DatasetList, 0)
	err := x.Table("dataset").Join("INNER", "repository", "(dataset.repo_id = repository.id and repository.is_private=false and repository.owner_id=?)", userID).OrderBy("dataset.updated_unix desc").Find(&datasets)
	if err != nil {
		log.Info("err:" + err.Error())
	}
	datasets.loadAttributes(x)
	return datasets
}

func GetOwnedPublicDatasetsByUserIDPage(orgId int64, keyword, categoryFilter, taskFilter, licenseFilter string, orderBy SearchOrderBy, page int, pageSize int, uid int64, isAdmin bool) (DatasetList, int64) {
	datasets := make(DatasetList, 0, pageSize)
	start := (page - 1) * pageSize

	var cond = builder.NewCond()
	cond = cond.And(builder.Eq{"1": 1})
	if len(keyword) > 0 {
		cond = cond.And(builder.Or(builder.Like{"LOWER(dataset.title)", strings.ToLower(keyword)}, builder.Like{"LOWER(dataset.description)", strings.ToLower(keyword)}))
	}
	if len(categoryFilter) > 0 {
		cond = cond.And(builder.Eq{"dataset.category": categoryFilter})
	}
	if len(taskFilter) > 0 {
		cond = cond.And(builder.Eq{"dataset.task": taskFilter})
	}
	if len(licenseFilter) > 0 {
		cond = cond.And(builder.Eq{"dataset.license": licenseFilter})
	}
	orderByStr := "dataset." + string(orderBy)
	if len(keyword) > 0 {
		orderByStr = "CASE WHEN dataset.title='" + keyword + "' THEN 0 ELSE 1 END," + orderByStr
	}
	var count int64
	if isAdmin {
		innnerCondition := "dataset.repo_id = repository.id and repository.owner_id=?"
		count1, err := x.Table("dataset").Join("INNER", "repository", innnerCondition, orgId).Where(cond).Count()
		if err != nil {
			log.Info("err:" + err.Error())
		}
		err = x.Table("dataset").Join("INNER", "repository", innnerCondition, orgId).Where(cond).OrderBy(orderByStr).Limit(pageSize, start).Find(&datasets)
		if err != nil {
			log.Info("err:" + err.Error())
		}
		count = count1
	} else {
		innnerCondition := "dataset.repo_id = repository.id and ((repository.is_private=false and repository.owner_id=?) or (repository.is_private=true and dataset.user_id=? and repository.owner_id=?))"
		count1, err := x.Table("dataset").Join("INNER", "repository", innnerCondition, orgId, uid, orgId).Where(cond).Count()
		if err != nil {
			log.Info("err:" + err.Error())
		}
		err = x.Table("dataset").Join("INNER", "repository", innnerCondition, orgId, uid, orgId).Where(cond).OrderBy(orderByStr).Limit(pageSize, start).Find(&datasets)
		if err != nil {
			log.Info("err:" + err.Error())
		}
		count = count1
	}

	datasets.loadAttributes(x)
	return datasets, count
}

func GetAllOwnedPublicDatasetsByOrgID(userID int64) DatasetList {
	datasets := make(DatasetList, 0)

	err := x.Table("dataset").Join("INNER", "repository", "(dataset.repo_id = repository.id and repository.is_private=false and repository.owner_id=?)", userID).Find(&datasets)
	if err != nil {
		log.Info("err:" + err.Error())
	}

	return datasets
}

func GetDefaultStorageType() int {
	if setting.StorageDefaultType == "obs" {
		return StorageTypeObs
	}
	return StorageTypeMinio
}
