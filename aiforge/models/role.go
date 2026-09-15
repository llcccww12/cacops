package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type RoleType string

const (
	TechProgramAdmin RoleType = "TechProgramAdmin"
	RewardPointAdmin RoleType = "RewardPointAdmin"
	MonitorAdmin     RoleType = "MonitorAdmin"
	Subscriber       RoleType = "Subscriber"
)

const (
	OperType     = 0
	ResourceType = 1
	StorageType  = 2
)

const (
	CommonType = 1
	MinusType  = -1
)

type Role struct {
	Type        RoleType
	Name        string
	Description string
}

type AiforgeRole struct {
	ID                int64 `xorm:"pk autoincr"`
	Name              string
	Type              int    //0，操作角色，1，资源角色 2,存储角色
	IsCommon          int    //0，通用角色，1，指定用户角色
	Description       string `xorm:"varchar(2000)"`
	RightInfo         string `xorm:"text NULL"`
	CreatedUserId     int64
	UserName          string             `xorm:"-" json:"userName"`
	UserRelAvatarLink string             `xorm:"-" json:"userRelAvatarLink"`
	UpdatedUnix       timeutil.TimeStamp `xorm:"updated"`
	CreatedUnix       timeutil.TimeStamp `xorm:"INDEX created"`
}

type AiforgeUserRole struct {
	ID     int64 `xorm:"pk autoincr"`
	UserId int64 `xorm:"INDEX"`
	RoleId int64
	Type   int //1,默认值，表示此用户拥有此角色，-1,该用户不拥有此普通角色，仅用于通用角色排除。
	//由于生产环境用户数量已经达到10W+，如果给每个用户都配置角色，则数据库记录达100W，影响查询，因此通用
	//角色默认所有用户都具有，如果某些用户不具有通用角色权限，则type=-1时，将排除roleid指定的角色权限。
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

type AiforgeOperation struct {
	ID          int64 `xorm:"pk autoincr"`
	Name        string
	Description string
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

type RightInfo struct {
	OperName        string `json:"operName"`
	Num             string `json:"num"`
	TaskType        string `json:"taskType"`
	SpecId          string `json:"specId"`
	ComputeResource string `json:"computeResource"`
	CodeSize        string `json:"codeSize"`   // code目录大小限制，单位GB，仅启智混合集群生效
	OutputSize      string `json:"outputSize"` // output目录大小限制，单位GB，仅启智混合集群生效
}

type UserRole struct {
	ID          int64              `xorm:"pk autoincr"`
	UserId      int64              `xorm:"INDEX UNIQUE(uq_user_role)"`
	RoleType    RoleType           `xorm:"UNIQUE(uq_user_role)"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

type FindAvailSpecsOptions struct {
	JobType           JobType
	ComputeResource   string
	Cluster           string
	UserId            int64
	SpecId            int64
	SourceSpecId      string
	VisualizeRequired bool
	HasInternet       SpecInternetQuery
}

type RightUserOptions struct {
	Page           int
	PageSize       int
	OperRoleId     int64
	ResourceRoleId int64
	StorageRoleId  int64
	OrderBy        string
	UserName       string
	UserId         int64
}

func AddAiforgeRole(ar AiforgeRole) (int64, error) {
	return x.Insert(&ar)
}

func DelAiforgeRole(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(&AiforgeRole{})
}

func QueryAiforgeRole(id int64) (*AiforgeRole, error) {
	r := &AiforgeRole{}
	has, err := x.Where("id = ?", id).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func QueryAiforgeRoleByName(name string) ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)
	err := x.Select("*").Where("name = ?", name).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateAiforgeRole(role *AiforgeRole) error {
	_, err := x.ID(role.ID).Cols("name", "description", "right_info").Update(role)
	if err != nil {
		return err
	}
	return nil
}

func ListAiforgeRole(name string) ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)
	var err error
	if name != "" {
		err = x.Select("*").Where("name ILIKE '%" + name + "%'").OrderBy("updated_unix desc").Find(&r)
	} else {
		err = x.Select("*").OrderBy("updated_unix desc").Find(&r)
	}
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ListNonCommonAiforgeRoleByNameAndType(name string, roleType int) ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)
	var err error
	if name != "" {
		err = x.Select("*").Where(builder.Like{"name", name}.And(builder.Eq{"type": roleType}).And(builder.Neq{"is_common": 0})).OrderBy("updated_unix desc").Find(&r)
	} else {
		err = x.Select("*").Where(builder.Eq{"type": roleType}.And(builder.Neq{"is_common": 0})).OrderBy("updated_unix desc").Find(&r)
	}
	if err != nil {
		return nil, err
	}
	return r, nil
}

// ListCommonAiforgeRole 查询默认角色列表
func ListCommonAiforgeRole() ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)

	cond := builder.Eq{"is_common": 0}

	err := x.Where(cond).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func QueryAiforgeRoleByIdsAndCommonRole(ids []int64) ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)
	var cond builder.Cond
	if len(ids) > 0 {
		log.Info("ids len > 0")
		cond = builder.In("id", ids)
		cond = cond.Or(builder.Eq{"is_common": 0})
	} else {
		cond = builder.Eq{"is_common": 0}
	}

	err := x.Where(cond).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func QueryAiforgeRoleByIds(ids []int64) ([]*AiforgeRole, error) {
	r := make([]*AiforgeRole, 0)
	var cond builder.Cond
	log.Info("ids len > 0")
	cond = builder.In("id", ids)

	err := x.Where(cond).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func AddAiforgeUserRole(ro AiforgeUserRole) (int64, error) {
	return x.Insert(&ro)
}

func DelAiforgeUserRole(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(&AiforgeUserRole{})
}

func DelAiforgeUserRoleByUserId(userId int64) (int64, error) {
	return x.Where("user_id = ?", userId).Delete(&AiforgeUserRole{})
}

func DelAiforgeUserRoleByRoleId(roleId int64) (int64, error) {
	return x.Where("role_id = ?", roleId).Delete(&AiforgeUserRole{})
}

func QueryAvailResourceNotDistinct(opts FindAvailSpecsOptions) ([]*Specification, error) {
	return queryAvailResource(opts)
}

// 这个是历史的老方法，随机选了一条分中心的资源规格，容易出问题
func QueryDistinctResourceWithRandomQueue(opts FindAvailSpecsOptions) ([]*Specification, error) {
	specs, err := queryAvailResource(opts)
	specsfilter := make([]*Specification, 0)
	if err == nil {
		if opts.VisualizeRequired {
			for _, t := range specs {
				if t.EnableVisualization {
					specsfilter = append(specsfilter, t)
				}
			}
		} else {
			specsfilter = specs
		}
		//distinct by sourceSpecId
		specs = DistinctSpecs(specsfilter)
	}
	return specs, err
}

func QueryAggregatedResource(opts FindAvailSpecsOptions) ([]*AggregateSpecification, error) {
	specs, err := queryAvailResource(opts)
	specsfilter := make([]*Specification, 0)
	if err != nil {
		return nil, err
	}

	if opts.VisualizeRequired {
		for _, t := range specs {
			if t.EnableVisualization {
				specsfilter = append(specsfilter, t)
			}
		}
	} else {
		specsfilter = specs
	}
	var result = make([]*AggregateSpecification, 0, len(specs))

	sourceSpecIdMap := make(map[string]QueueCapabilityAggregate, 0)
	for i := 0; i < len(specs); i++ {
		spec := specs[i]
		if spec.SourceSpecId == "" {
			result = append(result, &AggregateSpecification{
				ID:                  spec.ID,
				SourceSpecId:        spec.SourceSpecId,
				AccCardsNum:         spec.AccCardsNum,
				AccCardType:         spec.AccCardType,
				CpuCores:            spec.CpuCores,
				MemGiB:              spec.MemGiB,
				GPUMemGiB:           spec.GPUMemGiB,
				ShareMemGiB:         spec.ShareMemGiB,
				ComputeResource:     spec.ComputeResource,
				UnitPrice:           spec.UnitPrice,
				QueueId:             spec.QueueId,
				QueueCode:           spec.QueueCode,
				QueueName:           spec.QueueName,
				QueueType:           spec.QueueType,
				HasInternet:         spec.HasInternet,
				EnableVisualization: spec.EnableVisualization,
				Cluster:             spec.Cluster,
				AiCenterCode:        spec.AiCenterCode,
				AiCenterName:        spec.AiCenterName,
				SpecExclusiveType:   spec.SpecExclusiveType, // xorm:"is_spec_exclusive"
				ExclusiveOrg:        spec.ExclusiveOrg,
				Remark:              spec.Remark,
				Capability:          spec.ParseQueueCapability(),
			})
			continue
		}
		if _, has := sourceSpecIdMap[spec.SourceSpecId]; has {
			capability := sourceSpecIdMap[spec.SourceSpecId]
			if spec.HasInternet == int(HasInternet) {
				capability.NetworkCapableQueuesExist = true
			} else {
				capability.NetworkIncapableQueuesExist = true

			}
			if spec.EnableVisualization {
				capability.VisualizeCapableQueuesExist = true
			} else {
				capability.VisualizeIncapableQueuesExist = true

			}
			sourceSpecIdMap[spec.SourceSpecId] = capability
			continue
		}
		capability := spec.ParseQueueCapability()
		result = append(result, &AggregateSpecification{
			ID:                  spec.ID,
			SourceSpecId:        spec.SourceSpecId,
			AccCardsNum:         spec.AccCardsNum,
			AccCardType:         spec.AccCardType,
			CpuCores:            spec.CpuCores,
			MemGiB:              spec.MemGiB,
			GPUMemGiB:           spec.GPUMemGiB,
			ShareMemGiB:         spec.ShareMemGiB,
			ComputeResource:     spec.ComputeResource,
			UnitPrice:           spec.UnitPrice,
			QueueId:             spec.QueueId,
			QueueCode:           spec.QueueCode,
			QueueName:           spec.QueueName,
			QueueType:           spec.QueueType,
			HasInternet:         spec.HasInternet,
			EnableVisualization: spec.EnableVisualization,
			Cluster:             spec.Cluster,
			AiCenterCode:        spec.AiCenterCode,
			AiCenterName:        spec.AiCenterName,
			SpecExclusiveType:   spec.SpecExclusiveType,
			ExclusiveOrg:        spec.ExclusiveOrg,
			Remark:              spec.Remark,
			Capability:          capability,
		})
		sourceSpecIdMap[spec.SourceSpecId] = capability
	}
	for i := 0; i < len(result); i++ {
		if result[i].SourceSpecId == "" {
			continue
		}
		if _, has := sourceSpecIdMap[result[i].SourceSpecId]; has {
			result[i].Capability = sourceSpecIdMap[result[i].SourceSpecId]
		}
	}

	return result, err
}

func queryAvailResource(opts FindAvailSpecsOptions) ([]*Specification, error) {
	re, err := QueryAiforgeUserRoleByUserId(opts.UserId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	minusRoleIdMap := make(map[int64]int64, 0)
	for _, tmp := range re {
		ids = append(ids, tmp.RoleId)
		if tmp.Type == MinusType {
			minusRoleIdMap[tmp.RoleId] = 1
		}
	}
	roles, err := QueryAiforgeRoleByIdsAndCommonRole(ids)
	specIdMap := make(map[int64]int64, 0)
	if err == nil {
		for _, tmp := range roles {
			if minusRoleIdMap[tmp.ID] == 1 {
				continue
			}
			setAvailSpecIdToMap(tmp, string(opts.JobType), specIdMap)
		}
	}
	specIds := make([]int64, 0)
	for k, _ := range specIdMap {
		if opts.SpecId > 0 {
			if k == opts.SpecId {
				specIds = append(specIds, k)
			}
		} else {
			specIds = append(specIds, k)
		}
	}
	if len(specIds) > 0 {
		specs, err := FindSpecsForNewRight(FindSpecsOptions{
			ComputeResource: opts.ComputeResource,
			Cluster:         opts.Cluster,
			Ids:             specIds,
			SpecStatus:      SpecOnShelf,
			SourceSpecId:    opts.SourceSpecId,
			HasInternet:     opts.HasInternet,
		})
		if err != nil {
			return make([]*Specification, 0), nil
		}
		return specs, nil
	} else {
		log.Info("Not found spec id,so return empty.")
		return make([]*Specification, 0), nil
	}
}

// setAvailSpecIdToMap 将RightInfo转换成map返回，同时key是对应的specId
func setAvailSpecIdToMap(tmp *AiforgeRole, jobType string, specIdMap map[int64]int64) {
	rightInfo := tmp.RightInfo
	tmpsRights := make([]*RightInfo, 0)
	if rightInfo != "" {
		jsonerr := json.Unmarshal([]byte(rightInfo), &tmpsRights)
		if jsonerr != nil {
			log.Info("un json error=" + jsonerr.Error())
			return
		} else {
			if tmp.Type == ResourceType {
				for _, tmpRight := range tmpsRights {
					if tmpRight.TaskType == jobType {
						num, err := strconv.ParseInt(tmpRight.SpecId, 10, 64)
						if err == nil {
							specIdMap[num] = 0
						}
					}
				}
			}
		}
	}
}

// GetSpecIdsByList 将RightInfo组装成多个返回specIds
func GetSpecIdsByList(roles []*AiforgeRole) []int64 {
	var (
		specIdMap  = make(map[int64]int64)
		specIds    = make([]int64, 0, 0)
		tmpsRights []*RightInfo
	)

	for _, role := range roles {
		if role.RightInfo != "" {
			jsonerr := json.Unmarshal([]byte(role.RightInfo), &tmpsRights)
			if jsonerr != nil {
				log.Info("TransSpecIdToMap json error=" + jsonerr.Error())
				continue
			}

			for _, tmpRight := range tmpsRights {
				specId, _ := strconv.ParseInt(tmpRight.SpecId, 10, 64)
				if specId > 0 {
					specIdMap[specId] = 0
				}
			}
		}
	}

	for specId := range specIdMap {
		specIds = append(specIds, specId)
	}

	return specIds
}

func QueryAiforgeUserRoleByUserId(userid int64) ([]*AiforgeUserRole, error) {
	r := make([]*AiforgeUserRole, 0)
	err := x.Where("user_id = ?", userid).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func AddRightOperation(ro AiforgeOperation) (int64, error) {
	return x.Insert(&ro)
}

func DelRightOperation(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(&AiforgeOperation{})
}

func ListRightOperation() ([]*AiforgeOperation, error) {
	r := make([]*AiforgeOperation, 0)
	err := x.Select("*").Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func EnsureDefaultAiforgeOperations() error {
	defaults := []AiforgeOperation{
		{Name: "online_infer_path", Description: "允许创建在线推理服务"},
		{Name: "debug_time", Description: "允许使用调试任务时长配额"},
		{Name: "multi_node", Description: "运行训练任务时可以使用多节点计算资源"},
		{Name: "multi_task", Description: "可以同时运行多个计算任务"},
		{Name: "ignore_flow_control", Description: "忽略上传流控限制"},
		{Name: "RewardPointAdmin", Description: "算力积分管理员（可给用户发放/扣减积分）"},
		{Name: "MonitorAdmin", Description: "任务监控管理员"},
		{Name: "KANBANAdmin", Description: "看板管理员"},
		{Name: "TechProgramAdmin", Description: "科技计划项目管理员"},
	}
	existing, err := ListRightOperation()
	if err != nil {
		return err
	}
	have := make(map[string]bool, len(existing))
	for _, item := range existing {
		have[item.Name] = true
	}
	for _, item := range defaults {
		if have[item.Name] {
			continue
		}
		if _, err := AddRightOperation(item); err != nil {
			return err
		}
	}
	return nil
}

func QueryRightOperation(id int64) (*AiforgeOperation, error) {
	r := &AiforgeOperation{}
	has, err := x.Where("id = ?", id).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func NewUserRole(r UserRole) (int64, error) {
	return x.Insert(&r)
}
func DeleteUserRole(roleType RoleType, userId int64) (int64, error) {
	return x.Where("role_type = ? and user_id = ?", roleType, userId).Delete(&UserRole{})
}

func GetUserRoleByUserAndRole(userId int64, roleType RoleType) (*UserRole, error) {
	r := &UserRole{}
	has, err := x.Where("role_type = ? and user_id = ?", roleType, userId).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func GetRightUserList(opts RightUserOptions) (users []*User, _ int64, _ error) {
	//sql := "select user.* from public.user "
	r := make([]*User, 0)

	userTableName := "\"user\""
	where := userTableName + ".type = 0"

	if opts.UserName != "" {
		where += " and "
		where += userTableName + ".lower_name like '%" + strings.ToLower(opts.UserName) + "%'"
	}
	if opts.UserId > 0 {
		where += " and "
		where += userTableName + ".id = " + fmt.Sprint(opts.UserId)
	}
	if opts.OperRoleId > 0 {
		where += " and "
		where += userTableName + ".id in (select b.user_id from public.aiforge_user_role b where b.role_id=" + fmt.Sprint(opts.OperRoleId) + " )"
	}
	if opts.ResourceRoleId > 0 {
		role, err := QueryAiforgeRole(opts.ResourceRoleId)
		if err == nil {
			where += " and "
			if role.IsCommon == 0 {
				where += userTableName + ".id not in (select b.user_id from public.aiforge_user_role b where b.role_id=" + fmt.Sprint(opts.ResourceRoleId) + " )"
			} else {
				where += userTableName + ".id in (select b.user_id from public.aiforge_user_role b where b.role_id=" + fmt.Sprint(opts.ResourceRoleId) + " )"
			}
		}
	}
	if opts.StorageRoleId > 0 {
		role, err := QueryAiforgeRole(opts.StorageRoleId)
		if err == nil {
			where += " and "
			if role.IsCommon == 0 {
				where += userTableName + ".id not in (select b.user_id from public.aiforge_user_role b where b.role_id=" + fmt.Sprint(opts.StorageRoleId) + " )"
			} else {
				where += userTableName + ".id in (select b.user_id from public.aiforge_user_role b where b.role_id=" + fmt.Sprint(opts.StorageRoleId) + " )"
			}
		}
	}
	sess := x.NewSession()
	defer sess.Close()
	count, err := sess.Where(where).Count(new(User))
	if err == nil {
		if opts.Page >= 0 && opts.PageSize > 0 {
			var start int
			if opts.Page == 0 {
				start = 0
			} else {
				start = (opts.Page - 1) * opts.PageSize
			}
			sess.Limit(opts.PageSize, start)
		}
		orderBy := userTableName + ".right_unix desc," + userTableName + ".updated_unix desc"
		if opts.OrderBy != "" {
			orderBy = userTableName + "." + opts.OrderBy
		}
		if opts.UserName != "" {
			orderBy = "CASE WHEN " + userTableName + ".name='" + opts.UserName + "' THEN 0 ELSE 1 END," + orderBy
		}
		err = sess.Select("*").Where(where).OrderBy(orderBy).Find(&r)
	}
	if err == nil {
		return r, count, nil
	} else {
		log.Info("query error.")
	}

	return nil, 0, nil
}

// GetRightUserList
func GetRightOrgList(opts RightUserOptions) (users []*User, total int64, err error) {
	r := make([]*User, 0)

	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"public.user.type": UserTypeOrganization})
	//组织只能配存储类角色
	if opts.StorageRoleId > 0 {
		role, err := QueryAiforgeRole(opts.StorageRoleId)
		if err != nil {
			return nil, 0, err
		}
		if role.IsCommon == 0 {
			//默认角色对组织不生效，此时返回空列表
			return r, 0, nil
		} else {
			cond = cond.And(builder.In("public.user.id", builder.Select("user_id").From("aiforge_user_role").Where(builder.Eq{"role_id": opts.StorageRoleId})))
		}

	}
	if opts.UserName != "" {
		cond = cond.And(builder.Like{"public.user.lower_name", strings.ToLower(opts.UserName)})
	}
	if opts.UserId > 0 {
		cond = cond.And(builder.Eq{"public.user.id": opts.UserId})
	}
	count, err := x.Where(cond).Count(new(User))
	if err != nil {
		return nil, 0, err
	}
	if opts.PageSize <= 0 || opts.PageSize > 100 {
		opts.PageSize = 15
	}
	if opts.Page <= 0 {
		opts.Page = 1
	}
	x.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize)
	orderBy := "public.user.right_unix desc, public.user.updated_unix desc"
	if opts.OrderBy != "" {
		orderBy = opts.OrderBy
	}
	err = x.Where(cond).OrderBy(orderBy).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&r)
	if err != nil {
		return nil, 0, err
	}
	return r, count, nil
}

func GetUserRoleList(userId int64) ([]UserRole, error) {
	r := make([]UserRole, 0)
	err := x.Where("user_id = ?", userId).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetRoleByCode(code string) (*Role, error) {
	r := &Role{}
	has, err := x.Where("code = ?", code).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

type ErrRoleNotExists struct {
}

func IsErrRoleNotExists(err error) bool {
	_, ok := err.(ErrRoleNotExists)
	return ok
}

func (err ErrRoleNotExists) Error() string {
	return fmt.Sprintf("role is not exists")
}

type OperateRoleReq struct {
	UserName string   `json:"user_name" binding:"Required"`
	RoleType RoleType `json:"role_type" binding:"Required"`
}
