package models

import (
	"encoding/json"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

const demoComputingPowerRoleName = "普惠算力演示"
const demoComputingPowerQueuePrefix = "demo-"

type demoQueueSeed struct {
	Code            string
	Name            string
	Cluster         string
	AiCenterCode    string
	AiCenterName    string
	ComputeResource string
	AccCardType     string
	CardsTotalNum   int
	Specs           []demoSpecSeed
}

type demoSpecSeed struct {
	SourceSpecId string
	AccCardsNum  int
	CpuCores     int
	MemGiB       float32
	GPUMemGiB    float32
	ShareMemGiB  float32
	UnitPrice    float64
}

type demoXPUSeed struct {
	CardType     string
	CardTypeShow string
	ResourceType string
	Company      string
	AccessTime   string
	All          demoXPUStat
	Week         demoXPUStat
	Month        demoXPUStat
}

type demoXPUStat struct {
	UsedDuration int64
	UsedCardHour int64
	UserCount    int64
	TaskCount    int64
}

func EnsureDemoComputingPowerData() error {
	if err := ensureDemoResourceCatalog(); err != nil {
		return err
	}
	if err := EnsureDemoXPUStatistic(); err != nil {
		return err
	}
	return nil
}

func ensureDemoResourceCatalog() error {
	existed, err := x.Where("queue_code LIKE ?", demoComputingPowerQueuePrefix+"%").Count(new(ResourceQueue))
	if err != nil {
		return err
	}

	var specIDs []int64
	if existed == 0 {
		specIDs, err = insertDemoResourceCatalog()
		if err != nil {
			return err
		}
	} else {
		specs := make([]ResourceSpecification, 0)
		err = x.Table("resource_specification").
			Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
			Where("resource_queue.queue_code LIKE ?", demoComputingPowerQueuePrefix+"%").
			Find(&specs)
		if err != nil {
			return err
		}
		specIDs = make([]int64, 0, len(specs))
		for _, spec := range specs {
			specIDs = append(specIDs, spec.ID)
		}
	}

	return ensureDemoComputingPowerRole(specIDs)
}

func insertDemoResourceCatalog() ([]int64, error) {
	seeds := []demoQueueSeed{
		{
			Code: "demo-sz-a100", Name: "鹏城云脑一号 A100 共享队列", Cluster: "C2Net",
			AiCenterCode: "cloudbrain_one", AiCenterName: "鹏城云脑一号",
			ComputeResource: "GPU", AccCardType: "A100", CardsTotalNum: 64,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-a100-1", AccCardsNum: 1, CpuCores: 16, MemGiB: 64, GPUMemGiB: 80, ShareMemGiB: 16, UnitPrice: 8.5},
				{SourceSpecId: "demo-a100-4", AccCardsNum: 4, CpuCores: 64, MemGiB: 256, GPUMemGiB: 80, ShareMemGiB: 32, UnitPrice: 32},
			},
		},
		{
			Code: "demo-sz-4090", Name: "鹏城云脑二号 4090 共享队列", Cluster: "C2Net",
			AiCenterCode: "cloudbrain_two", AiCenterName: "鹏城云脑二号",
			ComputeResource: "GPU", AccCardType: "4090", CardsTotalNum: 48,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-4090-1", AccCardsNum: 1, CpuCores: 12, MemGiB: 48, GPUMemGiB: 24, ShareMemGiB: 8, UnitPrice: 4.2},
				{SourceSpecId: "demo-4090-2", AccCardsNum: 2, CpuCores: 24, MemGiB: 96, GPUMemGiB: 24, ShareMemGiB: 16, UnitPrice: 8},
			},
		},
		{
			Code: "demo-hf-910", Name: "合肥智算 昇腾 910 队列", Cluster: "C2Net",
			AiCenterCode: "hefei", AiCenterName: "合肥类脑智能开放平台",
			ComputeResource: "NPU", AccCardType: "ASCEND910", CardsTotalNum: 32,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-910-1", AccCardsNum: 1, CpuCores: 24, MemGiB: 96, GPUMemGiB: 32, ShareMemGiB: 16, UnitPrice: 6.8},
				{SourceSpecId: "demo-910-8", AccCardsNum: 8, CpuCores: 192, MemGiB: 768, GPUMemGiB: 32, ShareMemGiB: 64, UnitPrice: 48},
			},
		},
		{
			Code: "demo-cd-910b", Name: "成都智算 昇腾 910B 队列", Cluster: "C2Net",
			AiCenterCode: "chengdu", AiCenterName: "成都人工智能计算中心",
			ComputeResource: "NPU", AccCardType: "ASCEND-D910B", CardsTotalNum: 40,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-910b-2", AccCardsNum: 2, CpuCores: 48, MemGiB: 192, GPUMemGiB: 64, ShareMemGiB: 24, UnitPrice: 18},
			},
		},
		{
			Code: "demo-xc-mlu290", Name: "中原智算 MLU290 队列", Cluster: "C2Net",
			AiCenterCode: "xuchang", AiCenterName: "中原人工智能计算中心",
			ComputeResource: "MLU", AccCardType: "MLU290", CardsTotalNum: 24,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-mlu290-1", AccCardsNum: 1, CpuCores: 16, MemGiB: 64, GPUMemGiB: 32, ShareMemGiB: 8, UnitPrice: 5.6},
			},
		},
		{
			Code: "demo-wh-t20", Name: "武汉智算 燧原 T20 队列", Cluster: "C2Net",
			AiCenterCode: "wuhan", AiCenterName: "武汉人工智能计算中心",
			ComputeResource: "GCU", AccCardType: "ENFLAME-T20", CardsTotalNum: 20,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-t20-1", AccCardsNum: 1, CpuCores: 20, MemGiB: 80, GPUMemGiB: 32, ShareMemGiB: 12, UnitPrice: 5.2},
			},
		},
		{
			Code: "demo-xa-dcu", Name: "西安智算 海光 DCU 队列", Cluster: "C2Net",
			AiCenterCode: "xian", AiCenterName: "西安未来人工智能计算中心",
			ComputeResource: "DCU", AccCardType: "DCU", CardsTotalNum: 16,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-dcu-1", AccCardsNum: 1, CpuCores: 32, MemGiB: 128, GPUMemGiB: 32, ShareMemGiB: 16, UnitPrice: 4.8},
			},
		},
		{
			Code: "demo-sz-biv100", Name: "鹏城云计算所 天数 BI-V100 队列", Cluster: "C2Net",
			AiCenterCode: "pclcci", AiCenterName: "鹏城云计算所",
			ComputeResource: "BIREN-GPU", AccCardType: "BI-V100", CardsTotalNum: 12,
			Specs: []demoSpecSeed{
				{SourceSpecId: "demo-biv100-1", AccCardsNum: 1, CpuCores: 16, MemGiB: 64, GPUMemGiB: 32, ShareMemGiB: 8, UnitPrice: 4.5},
			},
		},
	}

	now := time.Now().Unix()
	specIDs := make([]int64, 0, 16)
	for _, seed := range seeds {
		queue := ResourceQueue{
			QueueCode:       seed.Code,
			QueueName:       seed.Name,
			QueueType:       QueueTypePublic,
			Cluster:         seed.Cluster,
			AiCenterCode:    seed.AiCenterCode,
			AiCenterName:    seed.AiCenterName,
			ComputeResource: seed.ComputeResource,
			AccCardType:     seed.AccCardType,
			CardsTotalNum:   seed.CardsTotalNum,
			HasInternet:     int(HasInternet),
			IsAvailable:     true,
			CreatedBy:       1,
			UpdatedBy:       1,
		}
		if _, err := x.Insert(&queue); err != nil {
			return nil, err
		}
		for _, specSeed := range seed.Specs {
			spec := ResourceSpecification{
				QueueId:      queue.ID,
				SourceSpecId: specSeed.SourceSpecId,
				AccCardsNum:  specSeed.AccCardsNum,
				CpuCores:     specSeed.CpuCores,
				MemGiB:       specSeed.MemGiB,
				GPUMemGiB:    specSeed.GPUMemGiB,
				ShareMemGiB:  specSeed.ShareMemGiB,
				UnitPrice:    specSeed.UnitPrice,
				Status:       SpecOnShelf,
				IsAvailable:  true,
				CreatedBy:    1,
				UpdatedBy:    1,
			}
			if _, err := x.Insert(&spec); err != nil {
				return nil, err
			}
			specIDs = append(specIDs, spec.ID)
		}
	}

	scene := ResourceScene{
		SceneName:       "普惠算力演示",
		JobType:         "TRAIN",
		Cluster:         "C2Net",
		ComputeResource: "GPU",
		IsSpecExclusive: SpecPublic,
		SceneType:       SceneTypePublic,
		CreatedBy:       1,
		UpdatedBy:       1,
	}
	if _, err := x.Insert(&scene); err != nil {
		return nil, err
	}
	for _, specID := range specIDs {
		if _, err := x.Insert(&ResourceSceneSpec{
			SceneId:     scene.ID,
			SpecId:      specID,
			CreatedTime: timeutil.TimeStamp(now),
		}); err != nil {
			return nil, err
		}
	}

	log.Info("Seeded %d demo computing-power specs", len(specIDs))
	return specIDs, nil
}

func ensureDemoComputingPowerRole(specIDs []int64) error {
	if len(specIDs) == 0 {
		return nil
	}

	rights := make([]*RightInfo, 0, len(specIDs)*2)
	for _, specID := range specIDs {
		id := strconv.FormatInt(specID, 10)
		for _, jobType := range []string{"TRAIN", "DEBUG", "GENERAL"} {
			rights = append(rights, &RightInfo{
				TaskType: jobType,
				SpecId:   id,
			})
		}
	}
	payload, err := json.Marshal(rights)
	if err != nil {
		return err
	}

	existing, err := QueryAiforgeRoleByName(demoComputingPowerRoleName)
	if err != nil {
		return err
	}
	if len(existing) > 0 {
		role := existing[0]
		role.RightInfo = string(payload)
		role.Type = ResourceType
		role.IsCommon = 0
		role.Description = "演示环境：向所有用户开放普惠算力规格"
		_, err = x.ID(role.ID).Cols("right_info", "type", "is_common", "description").Update(role)
		return err
	}

	_, err = AddAiforgeRole(AiforgeRole{
		Name:          demoComputingPowerRoleName,
		Type:          ResourceType,
		IsCommon:      0,
		Description:   "演示环境：向所有用户开放普惠算力规格",
		RightInfo:     string(payload),
		CreatedUserId: 1,
	})
	return err
}

func EnsureDemoXPUStatistic() error {
	if xStatistic == nil {
		return nil
	}
	count, err := xStatistic.Count(new(XPUInfoBase))
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	now := time.Now().Unix()
	seeds := []demoXPUSeed{
		{CardType: "ASCEND910", CardTypeShow: "昇腾 910", ResourceType: "NPU", Company: "华为", AccessTime: "2024-03-12",
			All: demoXPUStat{186420, 128600, 1860, 9420}, Week: demoXPUStat{12680, 8640, 312, 980}, Month: demoXPUStat{48600, 32800, 860, 3120}},
		{CardType: "ASCEND-D910B", CardTypeShow: "昇腾 910B", ResourceType: "NPU", Company: "华为", AccessTime: "2024-08-20",
			All: demoXPUStat{96800, 64200, 980, 4680}, Week: demoXPUStat{8420, 5680, 210, 640}, Month: demoXPUStat{28600, 19400, 520, 1860}},
		{CardType: "ENFLAME-T20", CardTypeShow: "燧原 T20", ResourceType: "GCU", Company: "燧原科技", AccessTime: "2024-05-18",
			All: demoXPUStat{74200, 48600, 760, 3520}, Week: demoXPUStat{6180, 4120, 168, 480}, Month: demoXPUStat{21400, 14200, 390, 1280}},
		{CardType: "MLU290", CardTypeShow: "寒武纪 MLU290", ResourceType: "MLU", Company: "寒武纪", AccessTime: "2024-04-09",
			All: demoXPUStat{56800, 36400, 620, 2680}, Week: demoXPUStat{4920, 3180, 142, 360}, Month: demoXPUStat{16800, 10800, 310, 920}},
		{CardType: "BI-V100", CardTypeShow: "天数 BI-V100", ResourceType: "BIREN-GPU", Company: "天数智芯", AccessTime: "2024-06-03",
			All: demoXPUStat{38600, 24600, 410, 1860}, Week: demoXPUStat{3280, 2140, 96, 240}, Month: demoXPUStat{11200, 7200, 210, 640}},
		{CardType: "DCU", CardTypeShow: "海光 DCU", ResourceType: "DCU", Company: "海光信息", AccessTime: "2024-02-21",
			All: demoXPUStat{29400, 18600, 360, 1420}, Week: demoXPUStat{2460, 1580, 78, 180}, Month: demoXPUStat{8600, 5400, 160, 480}},
		{CardType: "C500", CardTypeShow: "沐曦 C500", ResourceType: "METAX-GPGPU", Company: "沐曦", AccessTime: "2024-09-15",
			All: demoXPUStat{18600, 12400, 240, 860}, Week: demoXPUStat{1680, 1120, 54, 120}, Month: demoXPUStat{5400, 3600, 110, 280}},
		{CardType: "TIANSHE-R300", CardTypeShow: "昆仑芯 R300", ResourceType: "NPU", Company: "百度昆仑芯", AccessTime: "2024-01-16",
			All: demoXPUStat{12800, 8200, 180, 620}, Week: demoXPUStat{980, 640, 36, 80}, Month: demoXPUStat{3600, 2400, 86, 190}},
	}

	for _, seed := range seeds {
		base := &XPUInfoBase{
			CardType:     seed.CardType,
			CardTypeShow: seed.CardTypeShow,
			ResourceType: seed.ResourceType,
			Company:      seed.Company,
			AccessTime:   seed.AccessTime,
		}
		if _, err := xStatistic.Insert(base); err != nil {
			return err
		}
		stats := []XPUInfoStatistic{
			{InfoID: base.ID, Type: TypeAllDays, UsedDuration: seed.All.UsedDuration, UsedCardHour: seed.All.UsedCardHour, UserCount: seed.All.UserCount, TaskCount: seed.All.TaskCount, UpdatedUnix: now},
			{InfoID: base.ID, Type: TypeSevenDays, UsedDuration: seed.Week.UsedDuration, UsedCardHour: seed.Week.UsedCardHour, UserCount: seed.Week.UserCount, TaskCount: seed.Week.TaskCount, UpdatedUnix: now},
			{InfoID: base.ID, Type: TypeThirtyDays, UsedDuration: seed.Month.UsedDuration, UsedCardHour: seed.Month.UsedCardHour, UserCount: seed.Month.UserCount, TaskCount: seed.Month.TaskCount, UpdatedUnix: now},
		}
		if _, err := xStatistic.Insert(&stats); err != nil {
			return err
		}
	}

	log.Info("Seeded %d demo domestic XPU statistic cards", len(seeds))
	return nil
}
