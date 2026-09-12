package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type DatasetRegistry struct {
	ID             string             `xorm:"pk uuid"`
	Name           string             `xorm:"INDEX NOT NULL"`
	LowerName      string             `xorm:"INDEX NOT NULL"`
	Alias          string             `xorm:"INDEX"`
	LowerAlias     string             `xorm:"INDEX"`
	Tags           []string           `xorm:"jsonb"`
	License        string             `xorm:"varchar(200)"`
	Tasks          []string           `xorm:"jsonb"`
	StorageType    string             `xorm:"varchar(20)"`
	IsPrivate      bool               `xorm:"DEFAULT false"`
	UseCount       int64              `xorm:"INDEX DEFAULT 0"`
	DownloadCount  int64              `xorm:"INDEX DEFAULT 0"`
	NumCollections int                `xorm:"INDEX NOT NULL DEFAULT 0"`
	Recommend      bool               `xorm:"DEFAULT false"`
	CreatorID      int64              `xorm:"INDEX"`
	OwnerID        int64              `xorm:"INDEX"`
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`
	Size           int64              `xorm:"INDEX DEFAULT 0"`
	Path           string             `xorm:"varchar(400)"`
	Owner          *User              `xorm:"-"`
	CreatorName    string             `xorm:"-" json:"creator_name"`
	IsCollected    bool               `xorm:"-"`
}

type DatasetRegistryList []*DatasetRegistry

func (d DatasetRegistryList) loadAttributes(userId int64) error {
	d.loadOwners()
	d.loadCollection(userId)
	return nil
}

func (d DatasetRegistryList) loadCollection(userId int64) error {
	if userId <= 0 {
		return nil
	}
	datasetIds := make([]string, 0)
	for _, dataset := range d {
		datasetIds = append(datasetIds, dataset.ID)
	}
	collectionMap, err := GetDatasetCollectionsMap(userId, datasetIds)
	if err != nil {
		return err
	}
	for i := 0; i < len(d); i++ {
		d[i].IsCollected = collectionMap[d[i].ID]
	}
	return nil
}

func (d DatasetRegistryList) loadOwners() error {
	ownerIds := make([]int64, 0)
	for _, dataset := range d {
		if dataset.OwnerID > 0 {
			ownerIds = append(ownerIds, dataset.OwnerID)
		}
	}
	if len(ownerIds) == 0 {
		return nil
	}
	owners, err := GetUsersByIDs(ownerIds)
	if err != nil {
		return err
	}
	for _, dataset := range d {
		if dataset.OwnerID > 0 {
			for _, owner := range owners {
				if owner.ID == dataset.OwnerID {
					dataset.Owner = owner
					break
				}
			}
		}
	}
	return nil
}

type DatasetSearchOrder string

const (
	DatasetSearchOrderByDefault        DatasetSearchOrder = "recommend DESC,num_collections DESC,updated_unix DESC,name ASC"
	DatasetSearchOrderByNewest         DatasetSearchOrder = "created_unix DESC,name ASC"
	DatasetSearchOrderByRecentUpdated  DatasetSearchOrder = "updated_unix DESC,name ASC"
	DatasetSearchOrderByDownloadCounts DatasetSearchOrder = "download_count DESC,name ASC"
	DatasetSearchOrderByNumCollections DatasetSearchOrder = "num_collections DESC,name ASC"
	DatasetSearchOrderByUseCount       DatasetSearchOrder = "use_count DESC,name ASC"
	DatasetSearchOrderBySizeAsc        DatasetSearchOrder = "size ASC,name ASC"
	DatasetSearchOrderBySizeDesc       DatasetSearchOrder = "size DESC,name ASC"
	DatasetSearchOrderByAliasAsc       DatasetSearchOrder = `alias COLLATE "zh-x-icu" ASC, id ASC`
	DatasetSearchOrderByAliasDesc      DatasetSearchOrder = `alias COLLATE "zh-x-icu" DESC, id ASC`
)

const (
	// Visibility constants
	VisibilityPublic  = "public"
	VisibilityPrivate = "private"
	VisibilityAll     = "all"

	// Scope constants
	ScopeAll          = "all"
	ScopeCollected    = "collected"
	ScopeOwned        = "owned"
	ScopeCollaborated = "collaborated"
	ScopeAccessible   = "accessible"
	ScopeInvolved     = "involved"
	ScopeCreated      = "created"

	// Recommend constants
	RecommendAll  = "all"
	RecommendOnly = "only"

	// OrderBy constants
	OrderByNewest          = "newest"
	OrderByRecentUpdate    = "recentupdate"
	OrderByDownloadCount   = "downloadcount"
	OrderByCollections     = "collections"
	OrderByUseCount        = "usecount"
	OrderBySizeDesc        = "size"
	OrderBySizeAsc         = "size_asc"
	OrderByDerivativeCount = "derivativecount"
	OrderByAliasDesc       = "alias"
	OrderByAliasAsc        = "alias_asc"
	OrderByNameDesc        = "name"
	OrderByNameAsc         = "name_asc"

	// Owner type constants
	OwnerTypeIndividual   = "individual"
	OwnerTypeOrganization = "organization"
)

func (d DatasetSearchOrder) String() string {
	return string(d)
}

func GetDatasetSearchOrderBy(orderType string) DatasetSearchOrder {
	var orderBy DatasetSearchOrder
	switch orderType {
	case OrderByNewest:
		orderBy = DatasetSearchOrderByNewest
	case OrderByRecentUpdate:
		orderBy = DatasetSearchOrderByRecentUpdated
	case OrderByDownloadCount:
		orderBy = DatasetSearchOrderByDownloadCounts
	case OrderByCollections:
		orderBy = DatasetSearchOrderByNumCollections
	case OrderByUseCount:
		orderBy = DatasetSearchOrderByUseCount
	case OrderBySizeDesc:
		orderBy = DatasetSearchOrderBySizeDesc
	case OrderBySizeAsc:
		orderBy = DatasetSearchOrderBySizeAsc
	case OrderByAliasDesc:
		orderBy = DatasetSearchOrderByAliasDesc
	case OrderByAliasAsc:
		orderBy = DatasetSearchOrderByAliasAsc
	default:
		orderBy = DatasetSearchOrderByDefault
	}
	return orderBy
}

type DatasetRegistryShow4Action struct {
	ID        string
	Name      string
	OwnerName string
	Alias     string
}

func (dataset *DatasetRegistry) ToActionShow() *DatasetRegistryShow4Action {
	dataset.GetOwner()
	var ownerName string
	if dataset.Owner != nil {
		ownerName = dataset.Owner.Name
	}
	return &DatasetRegistryShow4Action{
		ID:        dataset.ID,
		Name:      dataset.Name,
		OwnerName: ownerName,
		Alias:     dataset.Alias,
	}
}

func (d DatasetRegistry) DisplayName() string {
	if d.Alias != "" {
		return d.Alias
	}
	return d.Name
}

func CreateDatasetRegistry(dataset *DatasetRegistry, doer *User) error {
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	_, err = sess.Insert(dataset)
	if err != nil {
		return err
	}
	if _, err = sess.Incr("num_datasets").ID(dataset.OwnerID).Update(new(User)); err != nil {
		return fmt.Errorf("increment user num_datasets: %v", err)
	}
	if err = dataset.getOwner(sess); err != nil {
		return fmt.Errorf("dataset getOwner: %v", err)
	}
	subjectCtx := dataset.ConvertSubjectAccessContext()
	owner := dataset.Owner
	if owner.IsOrganization() {
		if err = owner.GetTeams(&SearchTeamOptions{}); err != nil {
			return fmt.Errorf("GetTeams: %v", err)
		}
		for _, t := range owner.Teams {
			if t.IncludesAllDatasets {
				if err = t.addDataset(sess, dataset); err != nil {
					return fmt.Errorf("addDataset: %v", err)
				}
			}
		}

		if isAdmin, err := isUserDatasetAdmin(sess, dataset, doer); err != nil {
			return fmt.Errorf("isUserDatasetAdmin: %v", err)
		} else if !isAdmin {
			if err = subjectCtx.addCollaborator(sess, doer); err != nil {
				return fmt.Errorf("AddCollaborator: %v", err)
			}
			if err = subjectCtx.changeCollaborationAccessMode(sess, doer.ID, AccessModeAdmin); err != nil {
				return fmt.Errorf("ChangeCollaborationAccessMode: %v", err)
			}
		}
	} else if err = subjectCtx.recalculateAccesses(sess); err != nil {
		return fmt.Errorf("recalculateAccesses: %v", err)
	}
	sess.Commit()
	return nil
}

func UpdateDatasetPathByID(path string, datasetId string) error {
	_, err := x.ID(datasetId).Cols("path").Update(&DatasetRegistry{
		Path: path,
	})
	if err != nil {
		return err
	}

	return nil
}

func CreateDatasetRegistry4Old(dataset *DatasetRegistry, doer *User) error {
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	_, err = sess.NoAutoTime().Insert(dataset)
	if err != nil {
		return err
	}
	if _, err = sess.Incr("num_datasets").ID(dataset.OwnerID).Update(new(User)); err != nil {
		return fmt.Errorf("increment user num_datasets: %v", err)
	}
	sess.Commit()
	return nil
}

func GetDatasetRegistryByID(id string) (*DatasetRegistry, error) {
	dataset := &DatasetRegistry{}
	has, err := x.ID(id).Get(dataset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return dataset, nil
}

func GetDatasetRegistryListByIDs(ids []string) (DatasetRegistryList, error) {
	datasetList := make(DatasetRegistryList, 0)
	err := x.In("id", ids).Find(&datasetList)
	if err != nil {
		return nil, err
	}
	datasetList.loadAttributes(0)
	return datasetList, nil
}

func QueryDatasetRegistryMapsByIds(ids []string) (map[string]*DatasetRegistry, error) {
	re := make([]*DatasetRegistry, 0)
	err := x.Table(new(DatasetRegistry)).In("id", ids).Find(&re)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]*DatasetRegistry, 0)
	for _, m := range re {
		resultMap[m.ID] = m
	}
	return resultMap, nil
}

func UpdateDatasetRegistry(dataset DatasetRegistry) error {
	_, err := x.ID(dataset.ID).Cols("name", "lower_name", "tags", "license", "tasks", "is_private", "alias", "lower_alias").Unscoped().Update(dataset)
	if err != nil {
		return err
	}
	return nil
}

func (dataset *DatasetRegistry) ConvertSubjectAccessContext() *SubjectAccessContext {
	return &SubjectAccessContext{
		SubjectType: DatasetSubject,
		SubjectID:   dataset.ID,
		OwnerID:     dataset.OwnerID,
		IsPrivate:   dataset.IsPrivate,
		Dataset:     dataset,
	}
}

func DeleteDatasetRegistry(ctx DBContext, doer *User, uid int64, datasetID string) error {
	// In case is a organization.
	org, err := GetUserByID(uid)
	if err != nil {
		return err
	}
	if org.IsOrganization() {
		if err = org.GetTeams(&SearchTeamOptions{}); err != nil {
			return err
		}
	}

	sess := ctx.e

	dataset := &DatasetRegistry{ID: datasetID, OwnerID: uid}
	has, err := sess.Get(dataset)
	if err != nil {
		return err
	} else if !has {
		return ErrRecordNotExist{}
	}

	if cnt, err := sess.ID(datasetID).Delete(&DatasetRegistry{}); err != nil {
		return err
	} else if cnt != 1 {
		return ErrRecordNotExist{}
	}

	if org.IsOrganization() {
		for _, t := range org.Teams {
			if !t.hasSubject(sess, datasetID, DatasetSubject) {
				continue
			} else if err = t.removeDataset(sess, dataset, false); err != nil {
				return err
			}
		}
	}

	if err = deleteBeans(sess,
		&SubjectAccess{SubjectID: datasetID, SubjectType: int(DatasetSubject)},
		&Action{DatasetID: &datasetID},
		&DatasetCollection{DatasetID: datasetID},
		&SubjectCollaboration{SubjectID: datasetID, SubjectType: int(DatasetSubject)},
		&OfficialDatasetRegistry{DatasetId: datasetID},
	); err != nil {
		return fmt.Errorf("deleteBeans: %v", err)
	}

	return nil
}

func (dataset *DatasetRegistry) GetOwner() (err error) {
	return dataset.getOwner(x)
}

func (dataset *DatasetRegistry) getOwner(e Engine) (err error) {
	if dataset.Owner != nil {
		return nil
	}

	dataset.Owner, err = getUserByID(e, dataset.OwnerID)
	return err
}

func GetDatasetRegistryByOwnerAndName(ownerID int64, name string) (*DatasetRegistry, error) {
	dataset := &DatasetRegistry{}
	has, err := x.Where("owner_id = ? AND lower_name = ?", ownerID, strings.ToLower(name)).Get(dataset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return dataset, nil

}

func GetDatasetRegistryByOwnerAndAlias(ownerID int64, alias string) (*DatasetRegistry, error) {
	dataset := &DatasetRegistry{}
	has, err := x.Where("owner_id = ? AND lower_alias = ?", ownerID, strings.ToLower(alias)).Get(dataset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return dataset, nil

}

func GetDatasetRegistryByPath(path string) (*DatasetRegistry, error) {
	dataset := &DatasetRegistry{}
	has, err := x.Where("path = ? ", path).Get(dataset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return dataset, nil

}

func GetDatasetRegistryByOwnerNameAndDatasetName(ownerName, datasetName string) (*DatasetRegistry, error) {
	dataset := &DatasetRegistry{}
	has, err := x.Join("INNER", "public.user", "public.user.id = dataset_registry.owner_id").Where("dataset_registry.lower_name = ? and public.user.lower_name = ?", strings.ToLower(datasetName), strings.ToLower(ownerName)).Get(dataset)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return dataset, nil

}

func GetUsedDatasetRegistrySizeByUser(userId int64) (int64, error) {

	total, err := x.Where("owner_id = ? ", userId).Sum(new(DatasetRegistry), "size")

	return int64(total), err

}

func UpdateDatasetRegistrySize(id string, size int64) error {
	_, err := x.ID(id).Cols("size", "updated_unix").Update(&DatasetRegistry{
		Size:        size,
		UpdatedUnix: timeutil.TimeStampNow(),
	})
	if err != nil {
		return err
	}
	return nil
}

func IncreaseDatasetRegistryDownloadCount(id string) error {
	sess := x.NewSession()
	defer sess.Close()
	if _, err := sess.Exec("UPDATE `dataset_registry` SET download_count = download_count + 1 WHERE id = ?", id); err != nil {
		return err
	}

	return nil
}

type SearchDatasetReq struct {
	ListOptions
	Keyword            string   `json:"q"`
	Tasks              []string `json:"tasks"`
	Tags               []string `json:"tags"`
	Liscense           string   `json:"license"`
	CreatorName        string   `json:"creator_name"`
	OwnerType          string   `json:"owner_type"` // individual, organization
	Visibility         string   `json:"visibility"` // public, private, all
	Scope              string   `json:"scope"`      // all,collected,owned,collaborated,accessible,involved
	Recommend          string   `json:"recommend"`  // all,only
	OrderBy            string   `json:"order_by"`   // newest, recentupdate, downloadcount, collections, usecount
	OwnerName          string   `json:"owner_name"`
	From               string   `json:"from"`
	DatasetName        string   `json:"dataset_name"`
	OwnerId            int64    `json:"owner_id"`
	MinAccessMode      string   `json:"min_access_mode"` // read, write, admin, owner
	UseAdminPermission bool
	User               *User `json:"-"`
}

func SearchDatasetRegistry(opts SearchDatasetReq) (DatasetRegistryList, int64, error) {
	query := builder.NewCond()
	if opts.Keyword != "" {
		query = query.And(builder.Or(builder.Like{"lower_alias", "%" + strings.ToLower(opts.Keyword) + "%"}, builder.Like{"lower_name", "%" + strings.ToLower(opts.Keyword) + "%"}))
	}
	if len(opts.Tasks) > 0 {
		tasksJSON, _ := json.Marshal(opts.Tasks)
		if len(tasksJSON) > 0 {
			query = query.And(builder.Expr("tasks @> ?", string(tasksJSON)))
		}

	}
	if opts.DatasetName != "" {
		query = query.And(builder.Eq{"lower_name": strings.ToLower(opts.DatasetName)})
	}
	if opts.OwnerName != "" {
		var cond builder.Cond
		if opts.From == "profile" {
			cond = builder.Eq{"lower_name": strings.ToLower(opts.OwnerName)}
		} else {
			cond = builder.Like{"lower_name", "%" + strings.ToLower(opts.OwnerName) + "%"}
		}
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(cond)))
	}
	if opts.CreatorName != "" {
		query = query.And(builder.In("creator_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Like{"lower_name", "%" + strings.ToLower(opts.CreatorName) + "%"})))
	}
	if len(opts.Tags) > 0 {
		tagsJSON, _ := json.Marshal(opts.Tags)
		if len(tagsJSON) > 0 {
			query = query.And(builder.Expr("tags @> ?", string(tagsJSON)))
		}
	}
	if opts.Liscense != "" {
		query = query.And(builder.Eq{"license": opts.Liscense})
	}
	if opts.Visibility == VisibilityPublic {
		query = query.And(builder.Eq{"is_private": false})
	} else if opts.Visibility == VisibilityPrivate {
		query = query.And(builder.Eq{"is_private": true})
	}
	if opts.OwnerId > 0 {
		query = query.And(builder.Eq{"owner_id": opts.OwnerId})
	}
	if opts.OwnerType == OwnerTypeIndividual {
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Eq{"type": UserTypeIndividual})))
	} else if opts.OwnerType == OwnerTypeOrganization {
		query = query.And(builder.In("owner_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Eq{"type": UserTypeOrganization})))
	}
	if opts.Scope == ScopeOwned {
		ownedCond := builder.NewCond()
		ownedCond = ownedCond.Or(builder.Eq{"owner_id": opts.User.ID})
		ownedCond = ownedCond.Or(builder.In("id",
			builder.Select("`subject_access`.subject_id").
				From("subject_access").
				Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
					And(builder.Eq{"`subject_access`.subject_type": DatasetSubject}).
					And(builder.Eq{"`subject_access`.mode": int(AccessModeOwner)}))))
		query = query.And(ownedCond)
	} else if opts.Scope == ScopeCollected {
		query = query.And(builder.In("id",
			builder.Select("`dataset_collection`.dataset_id").
				From("dataset_collection").
				Where(builder.Eq{"`dataset_collection`.user_id": opts.User.ID})))
	} else if opts.Scope == ScopeCollaborated {

		if opts.MinAccessMode != "" {
			mode := ParseAccessMode(opts.MinAccessMode)
			query = query.And(
				builder.Neq{"owner_id": opts.User.ID},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": DatasetSubject}).
							And(builder.Gte{"`subject_access`.mode": int(mode)}).
							And(builder.Neq{"`subject_access`.mode": int(AccessModeOwner)}))),
					builder.In("id", builder.Select("`team_subject`.subject_id").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": opts.User.ID}.
							And(builder.Gte{"`team`.dataset_authorize": int(mode)}).
							And(builder.Neq{"`team`.dataset_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				),
			)

		} else {
			query = query.And(
				builder.Neq{"owner_id": opts.User.ID},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": DatasetSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin))))),
					builder.In("id", builder.Select("`team_subject`.subject_id").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": opts.User.ID}.
							And(builder.Neq{"`team`.dataset_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				),
			)
		}

	} else if opts.Scope == ScopeAccessible {
		if !opts.UseAdminPermission {

			accessibleCond := builder.NewCond()
			if opts.User != nil {
				// accessible = owned + collaborated + public
				accessibleCond = accessibleCond.Or(builder.Eq{"owner_id": opts.User.ID})
				accessibleCond = accessibleCond.Or(builder.In("id",
					builder.Select("`subject_access`.subject_id").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": opts.User.ID}.
							And(builder.Eq{"`subject_access`.subject_type": DatasetSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin), int(AccessModeOwner))))))
				accessibleCond = accessibleCond.Or(builder.Eq{"is_private": false})
			} else {
				// accessible = public
				accessibleCond = builder.Eq{"is_private": false}
			}

			query = query.And(accessibleCond)
		}
	}

	if opts.Recommend == RecommendOnly {
		query = query.And(builder.Eq{"recommend": true})
	}

	totalCount, err := x.Where(query).Count(new(DatasetRegistry))
	if err != nil {
		return nil, 0, err
	}

	orderBy := GetDatasetSearchOrderBy(opts.OrderBy).String()
	datasets := make(DatasetRegistryList, 0, opts.PageSize)
	err = x.Where(query).Limit(opts.PageSize, opts.PageSize*(opts.Page-1)).OrderBy(orderBy).Find(&datasets)
	if err != nil {
		return nil, 0, err
	}
	if datasets == nil {
		return nil, 0, nil
	}
	var userId int64
	if opts.User != nil {
		userId = opts.User.ID
	}
	datasets.loadAttributes(userId)
	return datasets, totalCount, nil
}

// CountPersonalDataset 统计个人的数据集个数
func CountPersonalDataset(userId int64) (allCount, privateCount, publicCount int64, err error) {
	query := builder.NewCond()

	ownedCond := builder.NewCond()
	ownedCond = ownedCond.Or(builder.Eq{"owner_id": userId})
	query = query.And(ownedCond)

	// 统计私有的
	privateCond := query
	privateCond = privateCond.And(builder.Eq{"is_private": true})
	privateCount, err = x.Where(privateCond).Count(new(DatasetRegistry))
	if err != nil {
		return
	}

	// 统计公开的
	publicCond := query
	publicCond = publicCond.And(builder.Eq{"is_private": false})
	publicCount, err = x.Where(publicCond).Count(new(DatasetRegistry))
	if err != nil {
		return
	}

	// 所有的
	allCount = privateCount + publicCount

	return
}

func GetOwnedPublicDatasetRegistrysByUserID(userID int64) (DatasetRegistryList, error) {
	datasets := make(DatasetRegistryList, 0)
	err := x.Where("owner_id = ? and is_private = false", userID).OrderBy("updated_unix desc").Find(&datasets)
	if err != nil {
		return nil, err
	}
	datasets.loadAttributes(userID)
	return datasets, nil
}

func RecommendDatasetRegistry(datasetId string, recommend bool) error {
	dataset := DatasetRegistry{Recommend: recommend}
	_, err := x.ID(datasetId).Cols("recommend").Update(dataset)
	return err
}

func GetAccessibleDatasetIDs(user *User) ([]string, error) {
	datasetIDs := make([]string, 0, 10)
	if err := x.
		Table("dataset_registry").
		Cols("id").
		Where(accessibleDatasetCondition(user)).
		Find(&datasetIDs); err != nil {
		return nil, fmt.Errorf("GetAccessibleDatasetIDs: %v", err)
	}
	return datasetIDs, nil
}

func AccessibleDatasetIDsQuery(user *User) *builder.Builder {
	return builder.Select("id").From("dataset_registry").Where(accessibleDatasetCondition(user))
}

func accessibleDatasetCondition(user *User) builder.Cond {
	var cond = builder.NewCond()

	if user == nil || !user.IsRestricted || user.ID <= 0 {
		orgVisibilityLimit := []structs.VisibleType{structs.VisibleTypePrivate}
		if user == nil || user.ID <= 0 {
			orgVisibilityLimit = append(orgVisibilityLimit, structs.VisibleTypeLimited)
		}
		// 1. Be able to see all non-private datasets that either:
		cond = cond.Or(builder.And(
			builder.Eq{"`dataset_registry`.is_private": false},
			// 2. Aren't in an private organisation or limited organisation if we're not logged in
			builder.NotIn("`dataset_registry`.owner_id", builder.Select("id").From("`user`").Where(
				builder.And(
					builder.Eq{"type": UserTypeOrganization},
					builder.In("visibility", orgVisibilityLimit)),
			))))
	}

	if user != nil {
		cond = cond.Or(
			// 2. Be able to see all datasets that we have access to
			builder.In("`dataset_registry`.id", builder.Select("subject_id").
				From("`subject_access`").
				Where(builder.And(
					builder.Eq{"user_id": user.ID},
					builder.Eq{"subject_type": DatasetSubject},
					builder.Gt{"mode": int(AccessModeNone)}))),
			// 3. Datasets that we directly own
			builder.Eq{"`dataset_registry`.owner_id": user.ID},
			// 4. Be able to see all datasets that we are in a team
			builder.In("`dataset_registry`.id", builder.Select("`team_subject`.subject_id").
				From("team_subject").
				Where(builder.And(
					builder.Eq{"`team_user`.uid": user.ID},
					builder.Eq{"`team_subject`.subject_type": DatasetSubject},
				)).
				Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id")),
			// 5. Be able to see all public datasets in private organizations that we are an org_user of
			builder.And(builder.Eq{"`dataset_registry`.is_private": false},
				builder.In("`dataset_registry`.owner_id",
					builder.Select("`org_user`.org_id").
						From("org_user").
						Where(builder.Eq{"`org_user`.uid": user.ID}))))
	}

	return cond
}

type OldDataset struct {
	Attachment *Attachment
	Dataset    *Dataset
	Repo       *Repository
}

func GetUnhandledOldDataset(page, pageSize int) ([]OldDataset, error) {
	var attachments []*Attachment
	var err error
	globalStart := (page - 1) * pageSize
	globalEnd := globalStart + pageSize
	// 1. 先获取热点数据集信息
	hotDatasetIds := GetHotDatasetIds()
	hotTotal := len(hotDatasetIds)

	// 2. 判断需要从热点数据集取多少
	if globalStart < hotTotal {
		hotEnd := globalEnd
		if hotEnd > hotTotal {
			hotEnd = hotTotal
		}
		hotAttachments, err := GetHotDataset(globalStart, hotEnd)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, hotAttachments...)
	}

	// 3. 判断是否需要从数据库补充
	if globalEnd > hotTotal {
		dbStart := 0
		if globalStart > hotTotal {
			dbStart = globalStart - hotTotal
		}
		dbPageSize := globalEnd - hotTotal - dbStart
		dbAttachments := make([]*Attachment, 0)

		err = x.
			Table("attachment").
			Where("dataset_id > ?", 0).
			And("decompress_state = ?", 1).
			And("NOT EXISTS (SELECT 1 FROM old_dataset_process_record WHERE old_dataset_process_record.id = attachment.uuid AND old_dataset_process_record.status = ?)", 4).
			OrderBy("id desc").Limit(dbPageSize, dbStart).Find(&dbAttachments)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, dbAttachments...)
	}
	var datasetIds []int64

	if len(attachments) == 0 {
		return nil, nil
	}

	attachmentIdMap := make(map[string]int, 0)
	distinctAttachments := make([]*Attachment, 0)
	for i := 0; i < len(attachments); i++ {
		attach := attachments[i]
		if _, ok := attachmentIdMap[attach.UUID]; ok {
			continue
		}
		distinctAttachments = append(distinctAttachments, attach)
		attachmentIdMap[attach.UUID] = 0
	}

	for _, attachment := range distinctAttachments {
		datasetIds = append(datasetIds, attachment.DatasetID)
	}
	var datasets []*Dataset
	err = x.In("id", datasetIds).Find(&datasets)
	if err != nil {
		return nil, err
	}
	if len(datasets) == 0 {
		return nil, nil
	}
	datasetMap := make(map[int64]*Dataset, len(datasets))
	for _, dataset := range datasets {
		datasetMap[dataset.ID] = dataset
	}
	var repoIds []int64
	for _, dataset := range datasets {
		if dataset.RepoID > 0 {
			repoIds = append(repoIds, dataset.RepoID)
		}
	}
	var repos []*Repository
	if len(repoIds) > 0 {
		err = x.In("id", repoIds).Find(&repos)
		if err != nil {
			return nil, err
		}
	}
	repoMap := make(map[int64]*Repository, len(repos))
	for _, repo := range repos {
		repoMap[repo.ID] = repo
	}
	oldDatasets := make([]OldDataset, 0, len(distinctAttachments))
	for _, attachment := range distinctAttachments {
		dataset, ok := datasetMap[attachment.DatasetID]
		if !ok {
			continue
		}
		repo, ok := repoMap[dataset.RepoID]
		if !ok {
			continue
		}
		oldDatasets = append(oldDatasets, OldDataset{
			Attachment: attachment,
			Dataset:    dataset,
			Repo:       repo,
		})
	}
	return oldDatasets, nil
}

func GetOldDataset(uuid string) (*OldDataset, error) {
	attachment, err := GetAttachmentByUUID(uuid)
	if err != nil {
		return nil, err
	}
	if attachment.DatasetID == 0 || attachment.DecompressState != DecompressStateDone {
		return nil, nil
	}
	dataset, err := GetDatasetByID(attachment.DatasetID)
	if err != nil {
		return nil, err
	}
	repo, err := GetRepositoryByID(dataset.RepoID)
	if err != nil {
		return nil, err
	}
	return &OldDataset{
		Attachment: attachment,
		Dataset:    dataset,
		Repo:       repo,
	}, nil
}

var hotDataset = make([]string, 0)
var hotDatasetLock = sync.Mutex{}
var initFlag = false

func GetHotDatasetIds() []string {
	if initFlag {
		return hotDataset
	}
	hotDatasetLock.Lock()
	defer hotDatasetLock.Unlock()
	if !initFlag {
		client := &http.Client{Timeout: 10 * time.Second}
		url := setting.HOT_DATASET_LIST_URL
		if url == "" {
			hotDataset = []string{}
			initFlag = true
			return []string{}
		}
		resp, err := client.Get(url)
		if err != nil {
			log.Error("http get GetHotDatasetIds err.%v", err)
			return []string{}
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			log.Error("http get GetHotDatasetIds status failed.%d:%s", resp.StatusCode, body)
			return []string{}
		}

		var result []string
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Error("GetHotDatasetIds JSON Unmarshal err : %v", err)
			return []string{}
		}
		hotDataset = result
		initFlag = true
		return result
	}
	return hotDataset
}

func GetHotDataset(start, end int) ([]*Attachment, error) {
	resultAttachments := make([]*Attachment, 0)
	hotDatasetIds := GetHotDatasetIds()
	if len(hotDatasetIds) == 0 {
		return resultAttachments, nil
	}

	if start >= len(hotDatasetIds) {
		return resultAttachments, nil
	}

	if end > len(hotDatasetIds) {
		end = len(hotDatasetIds)
	}

	currentDatasetIds := hotDatasetIds[start:end]
	if len(currentDatasetIds) == 0 {
		return resultAttachments, nil
	}
	err := x.
		Table("attachment").
		Where("dataset_id > ?", 0).
		And("decompress_state = ?", 1).
		And(builder.In("uuid", currentDatasetIds)).
		Find(&resultAttachments)
	if err != nil {
		return nil, err
	}
	return resultAttachments, nil
}

func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

type OldDatasetProcessRecord struct {
	ID          string             `xorm:"pk uuid"`
	Status      int                //0:unhandled,1: record_created, 2:readme_created , 3: collaborator_handled,4:collection_handled
	Remark      string             `xorm:"varchar(2000)"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`
}

func CreateOrFindOldDatasetProcessRecord(record *OldDatasetProcessRecord) (*OldDatasetProcessRecord, error) {
	p := &OldDatasetProcessRecord{}
	has, _ := x.ID(record.ID).Get(p)
	if has {
		return p, nil
	}
	_, err := x.Insert(record)
	if err != nil {
		return nil, err
	}
	x.ID(record.ID).Get(p)
	return p, nil
}

func UpdateOldDatasetProcessRecordStatus(id string, status int) error {
	_, err := x.ID(id).Cols("status").Update(&OldDatasetProcessRecord{
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateOldDatasetProcessRecordRemark(id string, remark string) error {
	_, err := x.ID(id).Cols("remark").Update(&OldDatasetProcessRecord{
		Remark: remark,
	})
	if err != nil {
		return err
	}
	return nil
}

func HandleOrg4OldDataset(orgID int64) error {
	_, err := x.Exec("update public.user set dataset_admin_change_team_access =repo_admin_change_team_access where id =? ", orgID)
	return err
}

func HandleTeam4OldDataset(t *Team, datasetAuthChanged bool, includeAllDatasetChanged bool) (err error) {

	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.ID(t.ID).Cols("can_create_org_dataset", "dataset_authorize", "includes_all_datasets").Update(t); err != nil {
		return fmt.Errorf("update: %v", err)
	}

	// Update access for team members if needed.
	if datasetAuthChanged {
		if err = t.getDatasets(sess); err != nil {
			return fmt.Errorf("getDatasets err: %v", err)
		}

		for _, dataset := range t.Datasets {
			subjectCtx := dataset.ConvertSubjectAccessContext()
			if err = subjectCtx.recalculateTeamAccesses(sess, 0, DatasetSubject); err != nil {
				return fmt.Errorf("recalculateTeamAccesses: %v", err)
			}
		}
	}

	if includeAllDatasetChanged && t.IncludesAllDatasets {
		err = t.addAllDatasets(sess)
		if err != nil {
			return fmt.Errorf("addAllDatasets: %v", err)
		}
	}

	return sess.Commit()
}

type StorageDeleteFailedDataset struct {
	ID             string             `xorm:"pk uuid"`
	Name           string             `xorm:"INDEX NOT NULL"`
	LowerName      string             `xorm:"INDEX NOT NULL"`
	Tags           []string           `xorm:"jsonb"`
	License        string             `xorm:"varchar(200)"`
	Tasks          []string           `xorm:"jsonb"`
	StorageType    string             `xorm:"varchar(20)"`
	IsPrivate      bool               `xorm:"DEFAULT false"`
	UseCount       int64              `xorm:"INDEX DEFAULT 0"`
	DownloadCount  int64              `xorm:"INDEX DEFAULT 0"`
	NumCollections int                `xorm:"INDEX NOT NULL DEFAULT 0"`
	Recommend      bool               `xorm:"DEFAULT false"`
	CreatorID      int64              `xorm:"INDEX"`
	OwnerID        int64              `xorm:"INDEX"`
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`
	Size           int64              `xorm:"INDEX DEFAULT 0"`
	Path           string             `xorm:"varchar(400)"`
	Owner          *User              `xorm:"-"`
	IsCollected    bool               `xorm:"-"`
}

func InsertStorageDeleteFailedDataset(dataset *DatasetRegistry) error {
	r := &StorageDeleteFailedDataset{
		ID:             dataset.ID,
		Name:           dataset.Name,
		LowerName:      dataset.LowerName,
		Tags:           dataset.Tags,
		License:        dataset.License,
		Tasks:          dataset.Tasks,
		StorageType:    dataset.StorageType,
		IsPrivate:      dataset.IsPrivate,
		UseCount:       dataset.UseCount,
		DownloadCount:  dataset.DownloadCount,
		NumCollections: dataset.NumCollections,
		Recommend:      dataset.Recommend,
		CreatorID:      dataset.CreatorID,
		OwnerID:        dataset.OwnerID,
		CreatedUnix:    dataset.CreatedUnix,
		UpdatedUnix:    dataset.UpdatedUnix,
		Size:           dataset.Size,
		Path:           dataset.Path,
	}
	_, err := x.Insert(r)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserDatasetNum(userId int64) error {
	_, err := x.Exec("UPDATE public.user SET num_datasets = (SELECT COUNT(*) FROM dataset_registry WHERE owner_id = ?) WHERE id = ? ", userId, userId)
	if err != nil {
		return err
	}
	return nil
}

type DatasetTags struct {
	Tag     []string
	Task    []string
	License []string
}

func GetOwnedPublicDatasetTags(userID int64) (*DatasetTags, error) {
	datasets := make([]DatasetRegistry, 0)

	err := x.Cols("tags", "license", "tasks").Where("owner_id =? and is_private = false", userID).Find(&datasets)
	if err != nil {
		return nil, err
	}
	tagsMap := make(map[string]int, 0)
	licenseMap := make(map[string]int, 0)
	tasksMap := make(map[string]int, 0)

	tags := make([]string, 0)
	licenses := make([]string, 0)
	tasks := make([]string, 0)

	for _, dataset := range datasets {
		for _, tag := range dataset.Tags {
			if tag == "" {
				continue
			}
			if _, ok := tagsMap[tag]; !ok {
				tagsMap[tag] = 0
				tags = append(tags, tag)
			}

		}
		if dataset.License != "" {
			if _, ok := licenseMap[dataset.License]; !ok {
				licenseMap[dataset.License] = 0
				licenses = append(licenses, dataset.License)
			}
		}
		for _, task := range dataset.Tasks {
			if task == "" {
				continue
			}
			if _, ok := tasksMap[task]; !ok {
				tasksMap[task] = 0
				tasks = append(tasks, task)
			}
		}

	}

	return &DatasetTags{
		Tag:     tags,
		Task:    tasks,
		License: licenses,
	}, nil
}

func GetHandledOldDatasetRegistryListByIDs(ids []string) (DatasetRegistryList, error) {
	datasetList := make(DatasetRegistryList, 0)
	err := x.In("dataset_registry.id", ids).Join("inner", "old_dataset_process_record", "old_dataset_process_record.id = dataset_registry.id").Where("old_dataset_process_record.status = 4").Find(&datasetList)
	if err != nil {
		return nil, err
	}
	datasetList.loadAttributes(0)
	return datasetList, nil
}

func GetHandledOldDatasetByIDs(ids []string) ([]OldDatasetProcessRecord, error) {
	records := make([]OldDatasetProcessRecord, 0)
	err := x.In("id", ids).Where("status = 4").Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func UpdateTeamDatasetnum(teamID int64) error {
	_, err := x.Exec("update team set num_datasets = (select count(*) from team_subject inner join dataset_registry on dataset_registry.id = team_subject.subject_id  where team_id = ? and subject_type = 1) where id = ?", teamID, teamID)
	return err
}
