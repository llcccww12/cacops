package models

import (
	"encoding/json"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type AITaskTemplateSearchOrder string

const (
	AITaskTemplateSearchOrderByDefault        AITaskTemplateSearchOrder = "recommend DESC,use_count DESC,num_collections DESC,updated_unix DESC,name ASC"
	AITaskTemplateSearchOrderByNewest         AITaskTemplateSearchOrder = "created_unix DESC,name ASC"
	AITaskTemplateSearchOrderByRecentUpdated  AITaskTemplateSearchOrder = "updated_unix DESC,name ASC"
	AITaskTemplateSearchOrderByNumCollections AITaskTemplateSearchOrder = "num_collections DESC,name ASC"
	AITaskTemplateSearchOrderByUseCount       AITaskTemplateSearchOrder = "use_count DESC,name ASC"
	AITaskTemplateSearchOrderByNameAsc        AITaskTemplateSearchOrder = `name COLLATE "zh-x-icu" ASC, id ASC`
	AITaskTemplateSearchOrderByNameDesc       AITaskTemplateSearchOrder = `name COLLATE "zh-x-icu" DESC, id ASC`
)

func (d AITaskTemplateSearchOrder) String() string {
	return string(d)
}

func GetAITaskTemplateSearchOrderBy(orderType string) AITaskTemplateSearchOrder {
	var orderBy AITaskTemplateSearchOrder
	switch orderType {
	case OrderByNewest:
		orderBy = AITaskTemplateSearchOrderByNewest
	case OrderByRecentUpdate:
		orderBy = AITaskTemplateSearchOrderByRecentUpdated
	case OrderByCollections:
		orderBy = AITaskTemplateSearchOrderByNumCollections
	case OrderByUseCount:
		orderBy = AITaskTemplateSearchOrderByUseCount
	case OrderByNameDesc:
		orderBy = AITaskTemplateSearchOrderByNameDesc
	case OrderByNameAsc:
		orderBy = AITaskTemplateSearchOrderByNameAsc
	default:
		orderBy = AITaskTemplateSearchOrderByDefault
	}
	return orderBy
}

type AITaskTemplate struct {
	//template base
	ID             string             `xorm:"pk uuid"`
	Name           string             `xorm:"INDEX"`
	LowerName      string             `xorm:"INDEX"`
	Description    string             `xorm:"TEXT"`
	Tags           []string           `xorm:"jsonb"`
	IsPrivate      bool               `xorm:"DEFAULT false"`
	UseCount       int64              `xorm:"INDEX DEFAULT 0"`
	NumCollections int                `xorm:"INDEX DEFAULT 0"`
	OwnerID        int64              `xorm:"INDEX"`
	Recommend      bool               `xorm:"DEFAULT false"`
	CreatedUnix    timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix    timeutil.TimeStamp `xorm:"INDEX updated"`

	//task base
	JobType           string `xorm:"NOT NULL DEFAULT 'DEBUG'"`
	Cluster           string `xorm:"NOT NULL DEFAULT 'C2Net'"`
	ComputeSource     string `xorm:"NOT NULL DEFAULT 'GPU'"`
	HasInternet       int    `xorm:"NOT NULL DEFAULT 0"`
	VisualizeRequired bool   `xorm:"DEFAULT false"`
	Parameters        string `xorm:"text"`

	//spec
	AccCardsNum int
	AccCardType string
	CpuCores    int
	MemGiB      float32
	GPUMemGiB   float32
	ShareMemGiB float32

	//image
	ImageID   string `xorm:"INDEX"`
	ImageName string
	ImageUrl  string

	//repo
	RepoId        int64 `xorm:"INDEX"`
	RepoName      string
	RepoOwnerName string
	BranchName    string
	BootFile      string
	IsRepoDeleted bool `xorm:"-"`

	Owner       *User                  `xorm:"-"`
	Repository  *Repository            `xorm:"-"`
	DatasetList []*TemplateDatasetInfo `xorm:"-"`
	ModelList   []*TemplateModelInfo   `xorm:"-"`

	IsCollected bool `xorm:"-"`
}

type TemplateDatasets struct {
	TemplateID   string `xorm:"uuid INDEX"`
	DatasetID    string `xorm:"uuid INDEX"`
	DatasetName  string `xorm:""`
	DatasetAlias string `xorm:""`
	OwnerName    string `xorm:""`
}

type TemplateModels struct {
	TemplateID string `xorm:"uuid INDEX"`
	ModelID    string `xorm:"INDEX"`
	ModelName  string `xorm:""`
	ModelAlias string `xorm:""`
	OwnerName  string `xorm:""`
}

type TemplateDatasetInfo struct {
	TemplateID   string
	DatasetID    string
	DatasetName  string
	DatasetAlias string
	OwnerName    string
	IsDeleted    bool
}

func (*TemplateDatasetInfo) TableName() string {
	return "template_datasets"
}

func (t *TemplateDatasetInfo) ToDTO() *TemplateDatasets {
	return &TemplateDatasets{
		TemplateID:   t.TemplateID,
		DatasetID:    t.DatasetID,
		DatasetName:  t.DatasetName,
		DatasetAlias: t.DatasetAlias,
		OwnerName:    t.OwnerName,
	}
}

type TemplateModelInfo struct {
	TemplateID string
	ModelID    string
	ModelName  string
	ModelAlias string
	OwnerName  string
	IsDeleted  bool
}

func (*TemplateModelInfo) TableName() string {
	return "template_models"
}

func (t *TemplateModelInfo) ToDTO() *TemplateModels {
	return &TemplateModels{
		TemplateID: t.TemplateID,
		ModelID:    t.ModelID,
		ModelName:  t.ModelName,
		ModelAlias: t.ModelAlias,
		OwnerName:  t.OwnerName,
	}
}

type SearchAITaskTemplateReq struct {
	ListOptions
	Keyword            string   `json:"q"`
	Tags               []string `json:"tags"`
	JobType            string
	ComputeSource      string
	Recommend          string `json:"recommend"` // all,only
	DatasetID          string
	ModelID            string
	RepoID             string
	OwnerName          string `json:"owner_name"`
	Scope              string `json:"scope"`      // all,collected
	Visibility         string `json:"visibility"` // public, private, all
	OrderBy            string `json:"order_by"`
	UseAdminPermission bool
	Operator           *User
}

type AITaskTemplateList []*AITaskTemplate

func (t AITaskTemplateList) loadAttributes(userId int64) error {
	t.loadDatasets()
	t.loadModels()
	t.loadOwners()
	t.loadRepos()
	t.loadCollection(userId)
	return nil
}

func (t AITaskTemplateList) loadCollection(userId int64) error {
	if userId <= 0 {
		return nil
	}
	templateIds := make([]string, 0)
	for _, template := range t {
		templateIds = append(templateIds, template.ID)
	}
	collectionMap, err := GetAITaskTemplateCollectionsMap(userId, templateIds)
	if err != nil {
		return err
	}
	for i := 0; i < len(t); i++ {
		t[i].IsCollected = collectionMap[t[i].ID]
	}
	return nil
}

func (t AITaskTemplateList) loadDatasets() error {
	templateIds := make([]string, 0)

	for _, template := range t {
		templateIds = append(templateIds, template.ID)
	}

	if len(templateIds) == 0 {
		return nil
	}

	datasets := make([]*TemplateDatasetInfo, 0)

	err := x.Select("template_datasets.template_id,COALESCE(dataset_registry.id, template_datasets.dataset_id) AS dataset_id,COALESCE(dataset_registry.name, template_datasets.dataset_name) AS dataset_name,COALESCE(dataset_registry.alias, template_datasets.dataset_alias) AS dataset_alias,COALESCE(public.user.name, template_datasets.owner_name) AS owner_name,(dataset_registry.id is null) as is_deleted").
		Table("template_datasets").
		Join("left", "dataset_registry", "dataset_registry.id = template_datasets.dataset_id").
		Join("left", "public.user", "dataset_registry.owner_id = public.user.id").
		Where(builder.In("template_datasets.template_id", templateIds)).
		Find(&datasets)

	if err != nil {
		return err
	}

	for _, template := range t {
		tmpMap := make(map[string]int, 0)
		templateDatasets := make([]*TemplateDatasetInfo, 0)
		for _, d := range datasets {
			if template.ID == d.TemplateID {
				if _, has := tmpMap[d.DatasetID]; !has {
					templateDatasets = append(templateDatasets, d)
					tmpMap[d.DatasetID] = 1
				}
			}

		}
		template.DatasetList = templateDatasets
	}

	return nil
}

func (t AITaskTemplateList) loadModels() error {
	templateIds := make([]string, 0)

	for _, template := range t {
		templateIds = append(templateIds, template.ID)
	}

	if len(templateIds) == 0 {
		return nil
	}

	templateModels := make([]*TemplateModelInfo, 0)

	err := x.Select("template_models.template_id,COALESCE(ai_model_manage.id, template_models.model_id) AS model_id,COALESCE(ai_model_manage.name, template_models.model_name) AS model_name,COALESCE(ai_model_manage.alias, template_models.model_alias) AS model_alias,COALESCE(public.user.name, template_models.owner_name) AS owner_name,(ai_model_manage.id is null) as is_deleted").
		Table("template_models").
		Join("left", "ai_model_manage", "ai_model_manage.id = template_models.model_id").
		Join("left", "public.user", "ai_model_manage.owner_id = public.user.id").
		Where(builder.In("template_models.template_id", templateIds)).
		Find(&templateModels)

	if err != nil {
		return err
	}

	for _, template := range t {
		tmpMap := make(map[string]int, 0)
		tmpModels := make([]*TemplateModelInfo, 0)
		for _, m := range templateModels {
			if template.ID == m.TemplateID {
				if _, has := tmpMap[m.ModelID]; !has {
					tmpModels = append(tmpModels, m)
					tmpMap[m.ModelID] = 1
				}
			}

		}
		template.ModelList = tmpModels
	}

	return nil
}

func (t AITaskTemplateList) loadOwners() error {
	ownerIds := make([]int64, 0)
	for _, template := range t {
		if template.OwnerID > 0 {
			ownerIds = append(ownerIds, template.OwnerID)
		}
	}
	if len(ownerIds) == 0 {
		return nil
	}
	owners, err := GetUsersByIDs(ownerIds)
	if err != nil {
		return err
	}
	for i := 0; i < len(t); i++ {
		if t[i].OwnerID > 0 {
			for _, owner := range owners {
				if owner.ID == t[i].OwnerID {
					t[i].Owner = owner
					break
				}
			}
		}
	}
	return nil
}

func (t AITaskTemplateList) loadRepos() error {
	repoIds := make([]int64, 0)

	for _, template := range t {
		repoIds = append(repoIds, template.RepoId)
	}

	if len(repoIds) == 0 {
		return nil
	}

	repoMaps, err := GetRepositoriesMapByIDs(repoIds)
	if err != nil {
		log.Error("AITaskTemplateList loadRepos error.%v", err)
		return err
	}

	for _, template := range t {
		if template.RepoId == 0 {
			continue
		}
		repo := repoMaps[template.RepoId]
		if repo == nil {
			template.IsRepoDeleted = true
		} else {
			template.IsRepoDeleted = false
			template.RepoName = repo.Name
			template.RepoOwnerName = repo.OwnerName
		}
	}

	return nil
}

func (t AITaskTemplate) CanRead(user *User) bool {
	if user == nil && t.IsPrivate {
		return false
	}

	if t.IsPrivate && user.ID != t.OwnerID && !user.IsAdmin {
		return false
	}

	return true
}
func (t AITaskTemplate) CanEdit(user *User) bool {
	if user == nil {
		return false
	}

	if user.ID != t.OwnerID && !user.IsAdmin {
		return false
	}

	return true
}

func (t AITaskTemplate) CanDelete(user *User) bool {
	if user == nil {
		return false
	}

	if user.ID != t.OwnerID && !user.IsAdmin {
		return false
	}

	return true
}

func CreateAITaskTemplate(template *AITaskTemplate) error {
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

	_, err = sess.Insert(template)
	if err != nil {
		return err
	}

	if len(template.DatasetList) > 0 {
		var tempDatasets = make([]*TemplateDatasets, 0, len(template.DatasetList))
		for i := 0; i < len(template.DatasetList); i++ {
			tempDatasets = append(tempDatasets, template.DatasetList[i].ToDTO())
		}
		_, err = sess.Insert(tempDatasets)
		if err != nil {
			return err
		}
	}

	if len(template.ModelList) > 0 {
		var tempModels = make([]*TemplateModels, 0, len(template.ModelList))
		for i := 0; i < len(template.ModelList); i++ {
			tempModels = append(tempModels, template.ModelList[i].ToDTO())
		}
		_, err = sess.Insert(tempModels)
		if err != nil {
			return err
		}
	}

	sess.Commit()
	return nil
}

// CountUserTemplates 统计个人的ai模板个数
func CountUserTemplates(userId int64) (totalCount int64, useTotal int64, err error) {
	var templates = make([]*AITaskTemplate, 0)

	query := builder.NewCond()

	// 统计个人的
	ownedCond := builder.NewCond().And(builder.Eq{"owner_id": userId})
	query = query.And(ownedCond)

	if err = x.Table(&AITaskTemplate{}).Where(query).Find(&templates); err != nil {
		return 0, 0, err
	}

	for _, template := range templates {
		useTotal += template.UseCount
	}

	return int64(len(templates)), useTotal, err
}

func UpdateAITaskTemplate(template *AITaskTemplate) error {
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

	var cols = "id,name,lower_name,description,tags,is_private,owner_id,job_type,cluster,compute_source,has_internet,visualize_required,parameters,acc_cards_num,acc_card_type,cpu_cores,mem_gi_b,gpu_mem_gi_b,share_mem_gi_b,image_id,image_name,image_url,repo_id,repo_name,repo_owner_name,branch_name,boot_file"
	_, err = sess.ID(template.ID).Cols(cols).Update(template)
	if err != nil {
		return err
	}

	_, err = sess.Where("template_id = ?", template.ID).Delete(&TemplateDatasets{})
	if err != nil {
		return err
	}
	_, err = sess.Where("template_id = ?", template.ID).Delete(&TemplateModels{})
	if err != nil {
		return err
	}

	if len(template.DatasetList) > 0 {
		var tempDatasets = make([]*TemplateDatasets, 0, len(template.DatasetList))
		for i := 0; i < len(template.DatasetList); i++ {
			tempDatasets = append(tempDatasets, template.DatasetList[i].ToDTO())
		}
		_, err = sess.Insert(tempDatasets)
		if err != nil {
			return err
		}
	}

	if len(template.ModelList) > 0 {
		var tempModels = make([]*TemplateModels, 0, len(template.ModelList))
		for i := 0; i < len(template.ModelList); i++ {
			tempModels = append(tempModels, template.ModelList[i].ToDTO())
		}
		_, err = sess.Insert(tempModels)
		if err != nil {
			return err
		}
	}

	sess.Commit()
	return nil
}

func GetAITaskTemplateByID(id string) (*AITaskTemplate, error) {
	template := &AITaskTemplate{}
	has, err := x.ID(id).Get(template)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return template, nil
}

func DelAITaskTemplateByID(id string) error {
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

	_, err = sess.ID(id).Delete(&AITaskTemplate{})
	if err != nil {
		return err
	}
	_, err = sess.Where("template_id = ?", id).Delete(&TemplateDatasets{})
	if err != nil {
		return err
	}
	_, err = sess.Where("template_id = ?", id).Delete(&TemplateModels{})
	if err != nil {
		return err
	}
	sess.Commit()
	return nil
}

func (t *AITaskTemplate) GetOwner() *User {
	if t.Owner != nil {
		return t.Owner
	}
	owner, err := GetUserByID(t.OwnerID)
	if err != nil {
		log.Error("GetUserByID failed, ownerId=%d", t.OwnerID)
		return nil
	}
	t.Owner = owner
	return owner
}

func (t *AITaskTemplate) LoadAtrribute(userId int64) error {
	t.GetOwner()
	t.loadDatasets()
	t.loadModels()
	t.loadRepo()
	t.loadCollection(userId)
	return nil
}

func (t *AITaskTemplate) loadDatasets() error {
	if len(t.DatasetList) > 0 {
		return nil
	}
	datasets := make([]*TemplateDatasetInfo, 0)

	err := x.Select("template_datasets.template_id,COALESCE(dataset_registry.id, template_datasets.dataset_id) AS dataset_id,COALESCE(dataset_registry.name, template_datasets.dataset_name) AS dataset_name,COALESCE(dataset_registry.alias, template_datasets.dataset_alias) AS dataset_alias,COALESCE(public.user.name, template_datasets.owner_name) AS owner_name,(dataset_registry.id is null) as is_deleted").
		Table("template_datasets").
		Join("left", "dataset_registry", "dataset_registry.id = template_datasets.dataset_id").
		Join("left", "public.user", "dataset_registry.owner_id = public.user.id").
		Where(builder.Eq{"template_datasets.template_id": t.ID}).
		Find(&datasets)
	if err != nil {
		return err
	}
	t.DatasetList = datasets
	return nil
}

func (t *AITaskTemplate) loadModels() error {
	if len(t.ModelList) > 0 {
		return nil
	}
	templateModels := make([]*TemplateModelInfo, 0)

	err := x.Select("template_models.template_id,COALESCE(ai_model_manage.id, template_models.model_id) AS model_id,COALESCE(ai_model_manage.name, template_models.model_name) AS model_name,COALESCE(ai_model_manage.alias, template_models.model_alias) AS model_alias,COALESCE(public.user.name, template_models.owner_name) AS owner_name,(ai_model_manage.id is null) as is_deleted").
		Table("template_models").
		Join("left", "ai_model_manage", "ai_model_manage.id = template_models.model_id").
		Join("left", "public.user", "ai_model_manage.owner_id = public.user.id").
		Where(builder.Eq{"template_models.template_id": t.ID}).
		Find(&templateModels)

	if err != nil {
		return err
	}
	t.ModelList = templateModels
	return nil
}

func (t *AITaskTemplate) loadRepo() {
	if t.RepoId == 0 {
		return
	}
	r, _ := GetRepositoryByID(t.RepoId)

	if r == nil {
		t.IsRepoDeleted = true
		return
	}

	t.RepoName = r.Name
	t.RepoOwnerName = r.OwnerName
	t.IsRepoDeleted = false
}

func (t *AITaskTemplate) loadCollection(userId int64) error {
	if userId <= 0 {
		return nil
	}

	isCollected := IsAITaskTemplateCollecting(userId, t.ID)
	t.IsCollected = isCollected
	return nil
}

func SearchAITaskTempalte(opts SearchAITaskTemplateReq) ([]*AITaskTemplate, int64, error) {
	var userId int64
	if opts.Operator != nil {
		userId = opts.Operator.ID
	}

	query := builder.NewCond()
	if opts.Keyword != "" {
		query = query.And(builder.Or(builder.Like{"ai_task_template.lower_name", "%" + strings.ToLower(opts.Keyword) + "%"}, builder.Like{"ai_task_template.description", "%" + strings.ToLower(opts.Keyword) + "%"}))
	}
	if opts.JobType != "" {
		query = query.And(builder.Eq{"ai_task_template.job_type": opts.JobType})
	}
	if opts.ComputeSource != "" {
		query = query.And(builder.Eq{"ai_task_template.compute_source": opts.ComputeSource})
	}
	if opts.OwnerName != "" {
		query = query.And(builder.In("ai_task_template.owner_id",
			builder.Select("id").
				From("public.user").
				Where(builder.Eq{"lower_name": strings.ToLower(opts.OwnerName)})))
	}
	if len(opts.Tags) > 0 {
		tagsJSON, _ := json.Marshal(opts.Tags)
		if len(tagsJSON) > 0 {
			query = query.And(builder.Expr("ai_task_template.tags @> ?", string(tagsJSON)))
		}
	}
	if opts.Visibility == VisibilityPublic {
		query = query.And(builder.Eq{"ai_task_template.is_private": false})
	} else if opts.Visibility == VisibilityPrivate {
		query = query.And(builder.Eq{"ai_task_template.is_private": true})
	}

	if opts.Scope == ScopeCollected {
		query = query.And(builder.In("ai_task_template.id",
			builder.Select("template_id").
				From("ai_task_template_collection").
				Where(builder.Eq{"user_id": userId})))
	} else if opts.Scope == ScopeCreated {
		query = query.And(builder.Eq{"ai_task_template.owner_id": userId})
	}

	if opts.DatasetID != "" {
		query = query.And(builder.In("ai_task_template.id",
			builder.Select("template_id").
				From("template_datasets").
				Where(builder.Eq{"dataset_id": opts.DatasetID})))
	}
	if opts.ModelID != "" {
		query = query.And(builder.In("ai_task_template.id",
			builder.Select("template_id").
				From("template_models").
				Where(builder.Eq{"model_id": opts.ModelID})))
	}
	if opts.RepoID != "" {
		query = query.And(builder.Eq{"repo_id": opts.RepoID})
	}
	if opts.Recommend == RecommendOnly {
		query = query.And(builder.Eq{"recommend": true})
	}

	totalCount, err := x.Where(query).Count(new(AITaskTemplate))
	if err != nil {
		return nil, 0, err
	}

	orderBy := GetAITaskTemplateSearchOrderBy(opts.OrderBy).String()
	templates := make(AITaskTemplateList, 0, opts.PageSize)
	err = x.Where(query).Limit(opts.PageSize, opts.PageSize*(opts.Page-1)).OrderBy(orderBy).Find(&templates)
	if err != nil {
		return nil, 0, err
	}
	if templates == nil {
		return nil, 0, nil
	}

	templates.loadAttributes(userId)
	return templates, totalCount, nil

}

func RecommendAITaskTemplate(templateId string, recommend bool) error {
	template := AITaskTemplate{Recommend: recommend}
	_, err := x.ID(templateId).Cols("recommend").Update(template)
	return err
}
