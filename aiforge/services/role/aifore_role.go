package role

import (
	"encoding/json"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

const (
	ROLE_OPER_ONLINE_INFER_PATH   = "online_infer_path"
	ROLE_OPER_DEBUG_TIME          = "debug_time"
	ROLE_OPER_TechProgramAdmin    = "TechProgramAdmin"
	ROLE_OPER_RewardPointAdmin    = "RewardPointAdmin"
	ROLE_OPER_MonitorAdmin        = "MonitorAdmin"
	ROLE_OPER_KANBANAdmin         = "KANBANAdmin"
	ROLE_OPER_IGNORE_FLOW_CONTROL = "ignore_flow_control"
	ROLE_OPER_MULTI_NODE          = "multi_node"
	ROLE_OPER_MULTI_TASK          = "multi_task"
)

func QueryUserRole(userId int64) ([]*models.AiforgeRole, []*models.AiforgeRole, []*models.AiforgeRole) {
	operRights := make([]*models.AiforgeRole, 0)
	resourceRights := make([]*models.AiforgeRole, 0)
	storageRights := make([]*models.AiforgeRole, 0)
	re, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		return nil, nil, nil
	}
	ids := make([]int64, 0)
	minusRoleIdMap := make(map[int64]int64, 0)
	for _, tmp := range re {
		ids = append(ids, tmp.RoleId)
		if tmp.Type == models.MinusType {
			minusRoleIdMap[tmp.RoleId] = 1
		}
	}
	allRoles, err := models.QueryAiforgeRoleByIdsAndCommonRole(ids)
	if err == nil {
		for _, tmp := range allRoles {
			if minusRoleIdMap[tmp.ID] == 1 {
				continue
			}
			role, err := models.QueryAiforgeRole(tmp.ID)
			if err == nil {
				if role.Type == models.OperType {
					operRights = append(operRights, role)
				}
				if role.Type == models.ResourceType {
					resourceRights = append(resourceRights, role)
				}
				if role.Type == models.StorageType {
					storageRights = append(storageRights, role)
				}
			}
		}
	}
	return operRights, resourceRights, storageRights
}

// 组织只查询存储类非默认角色
func QueryOrgRole(userId int64) ([]*models.AiforgeRole, []*models.AiforgeRole, []*models.AiforgeRole) {
	operRights := make([]*models.AiforgeRole, 0)
	resourceRights := make([]*models.AiforgeRole, 0)
	storageRights := make([]*models.AiforgeRole, 0)
	re, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		return nil, nil, nil
	}
	ids := make([]int64, 0)
	for _, tmp := range re {
		if tmp.Type == models.MinusType {
			continue
		}
		ids = append(ids, tmp.RoleId)
	}
	allRoles, err := models.QueryAiforgeRoleByIds(ids)
	if err != nil {
		log.Error("QueryAiforgeRoleByIds err %v", err)
		return operRights, resourceRights, storageRights
	}
	for _, tmp := range allRoles {
		//默认角色跳过
		if tmp.IsCommon == 0 {
			continue
		}
		role, err := models.QueryAiforgeRole(tmp.ID)
		if err != nil {
			log.Error("QueryAiforgeRoleByIds err %v", err)
			return operRights, resourceRights, storageRights
		}
		if role.Type == models.StorageType {
			storageRights = append(storageRights, role)
		}
	}
	return operRights, resourceRights, storageRights
}

func QueryUserStorageRole(userId int64) []*models.AiforgeRole {

	storageRights := make([]*models.AiforgeRole, 0)
	re, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		return nil
	}
	ids := make([]int64, 0)
	minusRoleIdMap := make(map[int64]int64, 0)
	for _, tmp := range re {
		ids = append(ids, tmp.RoleId)
		if tmp.Type == models.MinusType {
			minusRoleIdMap[tmp.RoleId] = 1
		}
	}
	allRoles, err := models.QueryAiforgeRoleByIdsAndCommonRole(ids)
	if err == nil {
		for _, tmp := range allRoles {
			if minusRoleIdMap[tmp.ID] == 1 {
				continue
			}

			if tmp.Type == models.StorageType {
				storageRights = append(storageRights, tmp)
			}

		}
	}
	return storageRights
}

func QueryOrgStorageRole(userId int64) []*models.AiforgeRole {
	//组织只有非默认角色才生效
	storageRights := make([]*models.AiforgeRole, 0)
	re, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		return nil
	}
	ids := make([]int64, 0)
	for _, tmp := range re {
		if tmp.Type == models.MinusType {
			continue
		}
		ids = append(ids, tmp.RoleId)

	}
	allRoles, err := models.QueryAiforgeRoleByIds(ids)
	if err == nil {
		for _, tmp := range allRoles {
			if tmp.Type == models.StorageType {
				storageRights = append(storageRights, tmp)
			}

		}
	}
	return storageRights
}

func UserHasOper(userId int64, operName string) bool {
	roles, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err == nil {
		for _, v := range roles {
			role, err := models.QueryAiforgeRole(v.RoleId)
			if err == nil {
				if role.Type == models.OperType {
					rightInfo := role.RightInfo
					operRights := make([]*models.RightInfo, 0)
					if rightInfo != "" {
						jsonerr := json.Unmarshal([]byte(rightInfo), &operRights)
						if jsonerr != nil {
							log.Info("un json error=" + jsonerr.Error())
						} else {
							for _, right := range operRights {
								if right.OperName == operName {
									return true
								}
							}
						}
					}
				}
			} else {
				log.Info("find role error=" + err.Error())
			}
		}
	} else {
		log.Info("find user role error=" + err.Error())
	}
	return false
}

func UserStorageNum(userId int64) int64 {
	user, err := models.GetUserByID(userId)
	if err != nil {
		return 0
	}
	var roles []*models.AiforgeRole
	if user.IsOrganization() {
		roles = QueryOrgStorageRole(userId)
	} else {
		roles = QueryUserStorageRole(userId)
	}

	if len(roles) == 0 {
		return 0
	}

	userStorage := 0

	for _, role := range roles {
		rightInfo := role.RightInfo
		storageRights := make([]*models.RightInfo, 0)
		if rightInfo != "" {
			jsonerr := json.Unmarshal([]byte(rightInfo), &storageRights)
			if jsonerr != nil {
				log.Info("un json error=" + jsonerr.Error())
			} else {
				for _, right := range storageRights {
					if right.Num != "" {
						num, err := strconv.Atoi(right.Num)
						if err != nil {
							continue
						} else {
							if num == -1 {
								return int64(num)
							}
							if num > userStorage {
								userStorage = num
							}

						}
					} else {
						continue
					}
				}
			}
		}
	}

	return int64(userStorage)

}

func AllUserOperNum(userId int64, operName string, operResourceType ...string) []int {
	nums := make([]int, 0)
	roles, _, _ := QueryUserRole(userId)

	for _, role := range roles {
		if role.Type == models.OperType {
			rightInfo := role.RightInfo
			operRights := make([]*models.RightInfo, 0)
			if rightInfo != "" {
				jsonerr := json.Unmarshal([]byte(rightInfo), &operRights)
				if jsonerr != nil {
					log.Info("un json error=" + jsonerr.Error())
				} else {
					for _, right := range operRights {
						if right.OperName == operName {
							if right.Num != "" {
								num, err := strconv.Atoi(right.Num)
								if err == nil {

									if len(operResourceType) > 0 {

										if operName == ROLE_OPER_MULTI_NODE && right.ComputeResource == operResourceType[0] {
											nums = append(nums, num)
										} else if operName == ROLE_OPER_MULTI_TASK && right.TaskType == operResourceType[0] {
											nums = append(nums, num)
										}

									} else {
										nums = append(nums, num)
									}

								}
							}
						}
					}
				}
			}
		}

	}

	return nums

}

func UserOperNum(userId int64, operName string, jobType string) int {
	roles, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err == nil {
		for _, v := range roles {
			role, err := models.QueryAiforgeRole(v.RoleId)
			if err == nil {
				if role.Type == models.OperType {
					rightInfo := role.RightInfo
					operRights := make([]*models.RightInfo, 0)
					if rightInfo != "" {
						jsonerr := json.Unmarshal([]byte(rightInfo), &operRights)
						if jsonerr != nil {
							log.Info("un json error=" + jsonerr.Error())
						} else {
							for _, right := range operRights {
								if right.OperName == operName {
									if right.Num != "" {
										num, err := strconv.Atoi(right.Num)
										if err != nil {
											return num
										} else {
											return getDefaultNum(jobType)
										}
									} else {
										return getDefaultNum(jobType)
									}
								}
							}
						}
					}
				}
			} else {
				log.Info("find role error=" + err.Error())
			}
		}
	} else {
		log.Info("find user role error=" + err.Error())
	}
	return 0
}

func getDefaultNum(jobType string) int {
	roleLimitMap := setting.ROLE_MULTI_LIMIT_MAP[string(jobType)]
	if roleLimitMap == nil || len(roleLimitMap) == 0 {
		return 1
	}
	return roleLimitMap["Subscriber"]
}

// GetUserContainerStorageLimits 查询用户容器内储存配额角色，返回code目录和output目录的大小限制（GB），0表示不限制
func GetUserContainerStorageLimits(userId int64) (codeSize int, outputSize int) {
	userRoles, err := models.QueryAiforgeUserRoleByUserId(userId)
	if err != nil {
		return 0, 0
	}
	ids := make([]int64, 0)
	minusRoleIdMap := make(map[int64]int64)
	for _, tmp := range userRoles {
		ids = append(ids, tmp.RoleId)
		if tmp.Type == models.MinusType {
			minusRoleIdMap[tmp.RoleId] = 1
		}
	}
	allRoles, err := models.QueryAiforgeRoleByIdsAndCommonRole(ids)
	if err != nil {
		return 0, 0
	}
	for _, tmp := range allRoles {
		if minusRoleIdMap[tmp.ID] == 1 {
			continue
		}
		if tmp.Type != models.StorageType {
			continue
		}
		rights := make([]*models.RightInfo, 0)
		if tmp.RightInfo == "" {
			continue
		}
		if err := json.Unmarshal([]byte(tmp.RightInfo), &rights); err != nil {
			log.Info("GetUserContainerStorageLimits unmarshal error=" + err.Error())
			continue
		}
		for _, right := range rights {
			if right.CodeSize != "" {
				if n, err := strconv.Atoi(right.CodeSize); err == nil && n > codeSize {
					codeSize = n
				}
			}
			if right.OutputSize != "" {
				if n, err := strconv.Atoi(right.OutputSize); err == nil && n > outputSize {
					outputSize = n
				}
			}
		}
	}
	return codeSize, outputSize
}
