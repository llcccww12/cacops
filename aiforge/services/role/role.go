package role

import "code.gitea.io/gitea/models"

var roleMap = map[models.RoleType]*models.Role{
	models.TechProgramAdmin: {
		Type:        models.TechProgramAdmin,
		Name:        "科技项目管理员",
		Description: "拥有科技项目管理相关功能的管理员权限",
	},
	models.RewardPointAdmin: {
		Type:        models.RewardPointAdmin,
		Name:        "奖励积分管理员",
		Description: "拥有奖励积分管理相关功能的管理员权限",
	},
	models.MonitorAdmin: {
		Type:        models.MonitorAdmin,
		Name:        "监测管理员",
		Description: "拥有监测的管理员权限",
	},
	models.Subscriber: {
		Type:        models.Subscriber,
		Name:        "Subscriber",
		Description: "Subscriber",
	},
}

func GetRole(roleType models.RoleType) *models.Role {
	return roleMap[roleType]
}
