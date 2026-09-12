package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"errors"
	"strings"
	"xorm.io/builder"
)

const (
	Exclusive = iota + 1
	NotExclusive
)

const (
	SceneTypePublic    = "public"
	SceneTypeExclusive = "exclusive"

	SpecPublic    = "public"
	SpecExclusive = "exclusive"
)

type ResourceScene struct {
	ID              int64 `xorm:"pk autoincr"`
	SceneName       string
	JobType         string
	Cluster         string
	ComputeResource string
	IsSpecExclusive string
	SceneType       string //共享或者独占场景
	ExclusiveOrg    string
	CreatedTime     timeutil.TimeStamp `xorm:"created"`
	CreatedBy       int64
	UpdatedTime     timeutil.TimeStamp `xorm:"updated"`
	UpdatedBy       int64
	DeleteTime      timeutil.TimeStamp `xorm:"deleted"`
	DeletedBy       int64
}

type ResourceSceneSpec struct {
	ID          int64              `xorm:"pk autoincr"`
	SceneId     int64              `xorm:"unique(idx_scene_spec)"`
	SpecId      int64              `xorm:"unique(idx_scene_spec)"`
	CreatedTime timeutil.TimeStamp `xorm:"created"`
}

type ResourceSceneReq struct {
	ID                int64
	SceneName         string
	JobType           string
	SceneType         string
	Cluster           string
	Resource          string
	ExclusiveQueueIds []int64
	IsSpecExclusive   string
	ExclusiveOrg      string
	CreatorId         int64
	SpecIds           []int64
}

type SearchResourceSceneOptions struct {
	ListOptions
	JobType             string
	IsSpecExclusive     string
	AiCenterCode        string
	QueueId             int64
	ComputeResource     string
	AccCardType         string
	Cluster             string
	HasInternet         SpecInternetQuery
	SceneType           string
	EnableVisualization int
}

type ResourceSceneListRes struct {
	TotalSize int64
	List      []ResourceSceneRes
}

func NewResourceSceneListRes(totalSize int64, list []ResourceSceneRes) *ResourceSceneListRes {

	return &ResourceSceneListRes{
		TotalSize: totalSize,
		List:      list,
	}
}

type ResourceSceneRes struct {
	ID              int64
	SceneName       string
	JobType         JobType
	IsSpecExclusive string
	Cluster         string
	ComputeResource string
	SceneType       string //共享或者独占场景
	ExclusiveOrg    string
	Specs           []ResourceSpecInfo
}

func (ResourceSceneRes) TableName() string {
	return "resource_scene"
}

type ResourceSceneBriefRes struct {
	ID        int64
	SceneName string
}

func (ResourceSceneBriefRes) TableName() string {
	return "resource_scene"
}

type ResourceSpecInfo struct {
	ID           int64
	SourceSpecId string
	AccCardsNum  int
	CpuCores     int
	MemGiB       float32
	GPUMemGiB    float32
	ShareMemGiB  float32
	UnitPrice    float64
	Status       int
	UpdatedTime  timeutil.TimeStamp
	SceneId      int64
	//queue
	Cluster             string
	AiCenterCode        string
	AiCenterName        string
	QueueCode           string
	QueueType           string
	QueueName           string
	QueueId             int64
	ComputeResource     string
	AccCardType         string
	HasInternet         int
	IsQueueExclusive    bool
	EnableVisualization bool
}

func (ResourceSpecInfo) TableName() string {
	return "resource_specification"
}

func InsertResourceScene(r ResourceSceneReq) error {
	sess := x.NewSession()
	defer sess.Close()

	err := sess.Begin()
	if err != nil {
		log.Error("InsertResourceScene start transaction err. %v", err)
		return err
	}
	//check
	specs := make([]ResourceSpecification, 0)
	cond := builder.In("id", r.SpecIds)
	if err := sess.Where(cond).Find(&specs); err != nil {
		return err
	}
	if len(specs) < len(r.SpecIds) {
		return errors.New("specIds not correct")
	}

	rs := ResourceScene{
		SceneName:       r.SceneName,
		JobType:         r.JobType,
		IsSpecExclusive: r.IsSpecExclusive,
		SceneType:       r.SceneType,
		Cluster:         r.Cluster,
		ComputeResource: r.Resource,
		ExclusiveOrg:    r.ExclusiveOrg,
		CreatedBy:       r.CreatorId,
		UpdatedBy:       r.CreatorId,
	}
	_, err = sess.InsertOne(&rs)
	if err != nil {
		sess.Rollback()
		return err
	}

	if len(r.SpecIds) > 0 {
		rss := make([]ResourceSceneSpec, len(r.SpecIds))
		for i, v := range r.SpecIds {
			rss[i] = ResourceSceneSpec{
				SceneId: rs.ID,
				SpecId:  v,
			}
		}

		_, err = sess.Insert(&rss)
		if err != nil {
			sess.Rollback()
			return err
		}
	}

	if r.SceneType == SceneTypeExclusive && r.ExclusiveOrg != "" && len(r.SpecIds) > 0 {
		pools := make([]ResourceExclusivePool, 0)
		queueIds := make([]int64, 0)
		err = sess.Table("resource_specification").Distinct("queue_id").In("id", r.SpecIds).Find(&queueIds)
		if err != nil {
			sess.Rollback()
			return err
		}
		for _, org := range strings.Split(r.ExclusiveOrg, ";") {
			if org == "" {
				continue
			}
			for _, id := range queueIds {
				pools = append(pools, ResourceExclusivePool{
					SceneId:         rs.ID,
					OrgName:         org,
					JobType:         r.JobType,
					Cluster:         r.Cluster,
					QueueId:         id,
					ComputeResource: r.Resource,
					CreatedBy:       r.CreatorId,
					UpdatedBy:       r.CreatorId,
				})
			}
		}
		_, err = sess.Insert(pools)
		if err != nil {
			sess.Rollback()
			return err
		}
	}

	return sess.Commit()
}

func UpdateResourceScene(r ResourceSceneReq) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()
	err = sess.Begin()
	if err != nil {
		log.Error("UpdateResourceScene start transaction err. %v", err)
		return err
	}
	// find old scene
	old := ResourceScene{}
	if has, _ := sess.ID(r.ID).Get(&old); !has {
		return errors.New("ResourceScene not exist")
	}
	//check specification
	specs := make([]ResourceSpecification, 0)
	cond := builder.In("id", r.SpecIds)
	if err := sess.Where(cond).Find(&specs); err != nil {
		return err
	}
	if len(specs) < len(r.SpecIds) {
		return errors.New("specIds not correct")
	}

	//update scene
	rs := ResourceScene{
		SceneName:       r.SceneName,
		IsSpecExclusive: r.IsSpecExclusive,
		ExclusiveOrg:    r.ExclusiveOrg,
		SceneType:       r.SceneType,
	}
	if _, err = sess.ID(r.ID).Cols("is_spec_exclusive", "scene_name", "exclusive_org", "scene_type").Update(&rs); err != nil {
		return err
	}

	//delete scene spec relation
	if _, err = sess.Where("scene_id = ? ", r.ID).Delete(&ResourceSceneSpec{}); err != nil {
		return err
	}

	if len(r.SpecIds) > 0 {
		//build new scene spec relation
		rss := make([]ResourceSceneSpec, len(r.SpecIds))
		for i, v := range r.SpecIds {
			rss[i] = ResourceSceneSpec{
				SceneId: r.ID,
				SpecId:  v,
			}
		}
		if _, err = sess.Insert(&rss); err != nil {
			return err
		}

	}

	if _, err = sess.Where("scene_id = ? ", r.ID).Delete(&ResourceExclusivePool{}); err != nil {
		return err
	}

	if r.SceneType == SceneTypeExclusive && r.ExclusiveOrg != "" && len(r.SpecIds) > 0 {
		pools := make([]ResourceExclusivePool, 0)
		queueIds := make([]int64, 0)
		err = sess.Table("resource_specification").Distinct("queue_id").In("id", r.SpecIds).Find(&queueIds)
		if err != nil {
			return err
		}
		for _, org := range strings.Split(r.ExclusiveOrg, ";") {
			if org == "" {
				continue
			}
			for _, id := range queueIds {
				pools = append(pools, ResourceExclusivePool{
					SceneId:         r.ID,
					OrgName:         org,
					JobType:         r.JobType,
					Cluster:         r.Cluster,
					QueueId:         id,
					ComputeResource: r.Resource,
					CreatedBy:       r.CreatorId,
					UpdatedBy:       r.CreatorId,
				})
			}
		}
		_, err = sess.Insert(pools)
		if err != nil {
			return err
		}
	}

	return sess.Commit()
}

func DeleteResourceScene(sceneId int64) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()
	err = sess.Begin()
	if err != nil {
		log.Error("DeleteResourceScene start transaction err. %v", err)
		return err
	}

	if _, err = sess.ID(sceneId).Delete(&ResourceScene{}); err != nil {
		return err
	}
	if _, err = sess.Where("scene_id = ? ", sceneId).Delete(&ResourceSceneSpec{}); err != nil {
		return err
	}
	if _, err = sess.Where("scene_id = ? ", sceneId).Delete(&ResourceExclusivePool{}); err != nil {
		return err
	}
	return sess.Commit()
}

func SearchResourceScene(opts SearchResourceSceneOptions) (int64, []ResourceSceneRes, error) {
	var cond = builder.NewCond()
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.JobType != "" {
		cond = cond.And(builder.Eq{"resource_scene.job_type": opts.JobType})
	}
	if opts.IsSpecExclusive != "" {
		cond = cond.And(builder.Eq{"resource_scene.is_spec_exclusive": opts.IsSpecExclusive})
	}
	if opts.AiCenterCode != "" {
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AiCenterCode})
	}
	if opts.QueueId > 0 {
		cond = cond.And(builder.Eq{"resource_queue.id": opts.QueueId})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"resource_queue.compute_resource": opts.ComputeResource})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"resource_queue.acc_card_type": opts.AccCardType})
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"resource_queue.cluster": opts.Cluster})
	}
	if opts.HasInternet == QueryHasInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": HasInternet})
	} else if opts.HasInternet == QueryNoInternetSpecs {
		cond = cond.And(builder.Eq{"resource_queue.has_internet": NoInternet})
	}
	if opts.SceneType != "" {
		cond = cond.And(builder.Eq{"resource_scene.scene_type": opts.SceneType})
	}
	if opts.EnableVisualization > 0 {
		if opts.EnableVisualization == 1 {
			cond = cond.And(builder.Neq{"resource_queue.enable_visualization": true})
		} else if opts.EnableVisualization == 2 {
			cond = cond.And(builder.Eq{"resource_queue.enable_visualization": true})
		}
	}
	cond = cond.And(builder.NewCond().Or(builder.Eq{"resource_scene.delete_time": 0}).Or(builder.IsNull{"resource_scene.delete_time"}))
	cols := []string{"resource_scene.id", "resource_scene.scene_name", "resource_scene.job_type", "resource_scene.is_spec_exclusive",
		"resource_scene.exclusive_org", "resource_scene.scene_type", "resource_scene.cluster", "resource_scene.compute_resource"}
	count, err := x.Where(cond).
		Distinct("resource_scene.id").
		Join("INNER", "resource_scene_spec", "resource_scene_spec.scene_id = resource_scene.id").
		Join("INNER", "resource_specification", "resource_specification.id = resource_scene_spec.spec_id").
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Count(&ResourceSceneRes{})
	if err != nil {
		return 0, nil, err
	}

	r := make([]ResourceSceneRes, 0)
	if err = x.Where(cond).Distinct(cols...).
		Join("INNER", "resource_scene_spec", "resource_scene_spec.scene_id = resource_scene.id").
		Join("INNER", "resource_specification", "resource_specification.id = resource_scene_spec.spec_id").
		Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
		Desc("resource_scene.id").
		Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
		Find(&r); err != nil {
		return 0, nil, err
	}

	if len(r) == 0 {
		return 0, r, err
	}

	//find related specs
	sceneIds := make([]int64, 0, len(r))
	for _, v := range r {
		sceneIds = append(sceneIds, v.ID)
	}

	specs := make([]ResourceSpecInfo, 0)

	if err := x.Cols("resource_specification.id", "resource_specification.source_spec_id",
		"resource_specification.acc_cards_num", "resource_specification.cpu_cores",
		"resource_specification.mem_gi_b", "resource_specification.gpu_mem_gi_b",
		"resource_specification.share_mem_gi_b", "resource_specification.unit_price",
		"resource_specification.status", "resource_specification.updated_time",
		"resource_scene_spec.scene_id", "resource_queue.cluster",
		"resource_queue.ai_center_code", "resource_queue.acc_card_type",
		"resource_queue.id as queue_id", "resource_queue.compute_resource",
		"resource_queue.queue_code", "resource_queue.ai_center_name",
		"resource_queue.has_internet", "resource_queue.queue_name",
		"resource_queue.queue_type", "resource_queue.enable_visualization",
	).In("resource_scene_spec.scene_id", sceneIds).
		Join("INNER", "resource_scene_spec", "resource_scene_spec.spec_id = resource_specification.id").
		Join("INNER", "resource_queue", "resource_queue.ID = resource_specification.queue_id").
		OrderBy("resource_specification.acc_cards_num").
		Find(&specs); err != nil {
		return 0, nil, err
	}

	specsMap := make(map[int64][]ResourceSpecInfo, 0)
	for _, v := range specs {
		if _, ok := specsMap[v.SceneId]; !ok {
			specsMap[v.SceneId] = []ResourceSpecInfo{v}
		} else {
			specsMap[v.SceneId] = append(specsMap[v.SceneId], v)
		}
	}

	for i, v := range r {
		s := specsMap[v.ID]
		if s == nil {
			s = make([]ResourceSpecInfo, 0)
		}
		r[i].Specs = s
	}

	return count, r, nil
}
