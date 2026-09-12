package models

import (
	"code.gitea.io/gitea/modules/timeutil"
)

type CloudbrainSpec struct {
	CloudbrainID    int64 `xorm:"pk"`
	SpecId          int64 `xorm:"index"`
	SourceSpecId    string
	AccCardsNum     int
	AccCardType     string `xorm:"index"`
	CpuCores        int
	MemGiB          float32
	GPUMemGiB       float32
	ShareMemGiB     float32
	ComputeResource string
	UnitPrice       float64
	QueueId         int64
	QueueCode       string
	QueueName       string
	QueueType       string
	Cluster         string
	HasInternet     int
	AiCenterCode    string `xorm:"index"`
	AiCenterName    string
	IsExclusive     bool
	ExclusiveOrg    string
	EnableVisualization bool
	CreatedTime     timeutil.TimeStamp `xorm:"created"`
	UpdatedTime     timeutil.TimeStamp `xorm:"updated"`
}

func (s CloudbrainSpec) ExclusiveType() string {
	if s.IsExclusive {
		return SpecExclusive
	}
	return SpecPublic
}

func (s CloudbrainSpec) ConvertToSpecification() *Specification {
	return &Specification{
		ID:                s.SpecId,
		SourceSpecId:      s.SourceSpecId,
		AccCardsNum:       s.AccCardsNum,
		AccCardType:       s.AccCardType,
		CpuCores:          s.CpuCores,
		MemGiB:            s.MemGiB,
		GPUMemGiB:         s.GPUMemGiB,
		ShareMemGiB:       s.ShareMemGiB,
		ComputeResource:   s.ComputeResource,
		UnitPrice:         s.UnitPrice,
		QueueId:           s.QueueId,
		QueueCode:         s.QueueCode,
		QueueName:         s.QueueName,
		QueueType:         s.QueueType,
		Cluster:           s.Cluster,
		AiCenterCode:      s.AiCenterCode,
		AiCenterName:      s.AiCenterName,
		SpecExclusiveType: s.ExclusiveType(),
		ExclusiveOrg:      s.ExclusiveOrg,
		HasInternet:       s.HasInternet,
	}
}

func NewCloudBrainSpec(cloudbrainId int64, s Specification) CloudbrainSpec {
	return CloudbrainSpec{
		CloudbrainID:        cloudbrainId,
		SpecId:              s.ID,
		SourceSpecId:        s.SourceSpecId,
		AccCardsNum:         s.AccCardsNum,
		AccCardType:         s.AccCardType,
		CpuCores:            s.CpuCores,
		MemGiB:              s.MemGiB,
		GPUMemGiB:           s.GPUMemGiB,
		ShareMemGiB:         s.ShareMemGiB,
		ComputeResource:     s.ComputeResource,
		UnitPrice:           s.UnitPrice,
		QueueId:             s.QueueId,
		QueueCode:           s.QueueCode,
		QueueName:           s.QueueName,
		QueueType:           s.QueueType,
		Cluster:             s.Cluster,
		AiCenterCode:        s.AiCenterCode,
		AiCenterName:        s.AiCenterName,
		IsExclusive:         s.IsSpecExclusive(),
		ExclusiveOrg:        s.ExclusiveOrg,
		HasInternet:         s.HasInternet,
		EnableVisualization: s.EnableVisualization,
	}
}

var StatusChangeChan = make(chan *Cloudbrain, 50)

func InsertCloudbrainSpec(c CloudbrainSpec) (int64, error) {
	return x.Insert(&c)
}

func GetCloudbrainSpecByID(cloudbrainId int64) (*CloudbrainSpec, error) {
	r := &CloudbrainSpec{}
	if has, err := x.Where("cloudbrain_id = ?", cloudbrainId).Get(r); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func GetSpecAccCardsNumByID(cloudbrainId int64) (*CloudbrainSpec, error) {
	r := &CloudbrainSpec{}
	if has, err := x.Select("acc_cards_num").Where("cloudbrain_id = ?", cloudbrainId).Get(r); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func FindCloudbrainTask(page, pageSize int) ([]*Cloudbrain, error) {
	r := make([]*Cloudbrain, 0)
	err := x.Unscoped().
		Limit(pageSize, (page-1)*pageSize).
		OrderBy("cloudbrain.id").
		Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CountNoSpecHistoricTask() (int64, error) {
	n, err := x.Unscoped().
		Where(" 1=1 and not exists (select 1 from cloudbrain_spec  where cloudbrain.id = cloudbrain_spec.cloudbrain_id)").
		Count(&Cloudbrain{})
	if err != nil {
		return 0, err
	}
	return n, nil
}

// GetResourceSpecMapByCloudbrainIDs
func GetResourceSpecMapByCloudbrainIDs(ids []int64) (map[int64]*Specification, error) {
	specs := make([]*CloudbrainSpec, 0)
	if err := x.In("cloudbrain_id", ids).Find(&specs); err != nil {
		return nil, err
	}
	r := make(map[int64]*Specification, len(ids))
	for _, s := range specs {
		r[s.CloudbrainID] = s.ConvertToSpecification()
	}
	return r, nil
}

func GetCloudbrainTaskUnitPrice(task Cloudbrain) (float64, error) {
	s, err := GetCloudbrainSpecByID(task.ID)
	if err != nil {
		return 0, err
	}
	var n = 1
	if task.WorkServerNumber > 1 {
		n = task.WorkServerNumber
	}
	return s.UnitPrice * float64(n), nil
}

func UpdateCloudbrainSpec(cloudbrainId int64, s *Specification) (int64, error) {
	new := CloudbrainSpec{
		CloudbrainID:    cloudbrainId,
		SpecId:          s.ID,
		SourceSpecId:    s.SourceSpecId,
		AccCardsNum:     s.AccCardsNum,
		AccCardType:     s.AccCardType,
		CpuCores:        s.CpuCores,
		MemGiB:          s.MemGiB,
		GPUMemGiB:       s.GPUMemGiB,
		ShareMemGiB:     s.ShareMemGiB,
		ComputeResource: s.ComputeResource,
		UnitPrice:       s.UnitPrice,
		QueueId:         s.QueueId,
		QueueCode:       s.QueueCode,
		QueueName:       s.QueueName,
		QueueType:       s.QueueType,
		Cluster:         s.Cluster,
		AiCenterCode:    s.AiCenterCode,
		AiCenterName:    s.AiCenterName,
		IsExclusive:     s.IsSpecExclusive(),
		ExclusiveOrg:    s.ExclusiveOrg,
		HasInternet:     s.HasInternet,
	}
	return x.Where("cloudbrain_id = ?", cloudbrainId).Update(&new)
}
