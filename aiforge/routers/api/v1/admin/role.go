package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/markup"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/role"
)

type RightUser struct {
	// the user's id
	ID int64 `json:"id"`
	// the user's username
	UserName string `json:"name"`
	// the user's full name
	FullName string `json:"fullName"`
	// Is the user an administrator
	AvatarURL string    `json:"avatarURL"`
	IsAdmin   bool      `json:"isAdmin"`
	Update    time.Time `json:"updateUnix,omitempty"`
	// swagger:strfmt date-time
	Created     time.Time `json:"createdUnix,omitempty"`
	RightUpdate time.Time `json:"rightUpdateUnix"`

	OperRole     []*models.AiforgeRole `json:"operRole"`
	ResourceRole []*models.AiforgeRole `json:"resourceRole"`
	StorageRole  []*models.AiforgeRole `json:"storageRole"`
}

func getRightJson(roleType int, ctx *context.APIContext) string {
	rights := make([]models.RightInfo, 0)
	if roleType == models.OperType {
		operNamesStrs := ctx.Query("operName")
		operNumsStrs := ctx.Query("operNum")
		computeResourceStrs := ctx.Query("computeResource")
		jobTypeStrs := ctx.Query("jobType")

		if operNamesStrs != "" {
			operNames := strings.Split(operNamesStrs, ",")
			operNums := make([]string, 0)
			computeResources := make([]string, 0)
			jobTypes := make([]string, 0)
			if operNumsStrs != "" {
				operNums = strings.Split(operNumsStrs, ",")
			}
			if computeResourceStrs != "" {
				computeResources = strings.Split(computeResourceStrs, ",")
			}
			if jobTypeStrs != "" {
				jobTypes = strings.Split(jobTypeStrs, ",")

			}
			for k, operName := range operNames {
				right := models.RightInfo{
					OperName: operName,
				}
				if k < len(operNums) && operNums[k] != "" {
					right.Num = operNums[k]

				}
				if k < len(computeResources) && computeResources[k] != "" {
					right.ComputeResource = computeResources[k]

				}
				if k < len(jobTypes) && jobTypes[k] != "" {
					right.TaskType = jobTypes[k]
				}
				rights = append(rights, right)
			}
		}
	} else if roleType == models.ResourceType {
		jobType := ctx.Query("jobType")
		specIdStrs := ctx.Query("specId")
		if specIdStrs != "" {
			jobTypes := strings.Split(jobType, ",")
			specIds := strings.Split(specIdStrs, ",")
			if len(jobTypes) != len(specIds) {
				log.Info("input error. jobTypes=" + jobType + " specid=" + specIdStrs)
			}
			for k, specId := range specIds {
				right := models.RightInfo{
					SpecId:   specId,
					TaskType: jobTypes[k],
				}
				rights = append(rights, right)
			}
		}
	} else if roleType == models.StorageType {
		right := models.RightInfo{
			Num:        ctx.Query("operNum"),
			CodeSize:   ctx.Query("codeSize"),
			OutputSize: ctx.Query("outputSize"),
		}
		rights = append(rights, right)
	}
	rightsJson, err := json.Marshal(rights)
	if err != nil {
		log.Info("getRightJson json error." + err.Error())
	}
	return string(rightsJson)
}

func AddAiforgeRole(ctx *context.APIContext) {
	resultMap := make(map[string]string, 0)
	resultMap["code"] = "0"
	name := ctx.Query("name")
	log.Info("name=" + name)
	temp := []rune(name)
	if len(temp) > 80 {
		resultMap["code"] = "-1"
		resultMap["msg"] = "角色名称长度不能超过80个字符。"
		ctx.JSON(200, resultMap)
		return
	}
	roleType := ctx.QueryInt("type")
	iscommon := ctx.QueryInt("isCommon")
	description := ctx.Query("description")
	temp = []rune(description)
	if len(temp) > 800 {
		resultMap["code"] = "-1"
		resultMap["msg"] = "角色描述长度不能超过800个字符。"
		ctx.JSON(200, resultMap)
		return
	}
	exist, err := models.QueryAiforgeRoleByName(name)
	if err == nil {
		if len(exist) > 0 {
			resultMap["code"] = "-1"
			resultMap["msg"] = "角色名称已经存在。"
			ctx.JSON(200, resultMap)
			return
		}
	}

	rightsJson := getRightJson(roleType, ctx)

	ar := models.AiforgeRole{
		Name:          name,
		Type:          roleType,
		IsCommon:      iscommon,
		Description:   description,
		RightInfo:     rightsJson,
		CreatedUserId: ctx.User.ID,
	}
	re, err := models.AddAiforgeRole(ar)
	if err == nil {
		log.Info("re=" + fmt.Sprint(re))
	} else {
		log.Info("error=" + err.Error())
		resultMap["code"] = "-1"
		resultMap["msg"] = "Add role error."
	}
	ctx.JSON(200, resultMap)

}

func DelAiforgeRole(ctx *context.APIContext) {
	roleId := ctx.QueryInt64("id")
	resultMap := make(map[string]string, 0)
	resultMap["code"] = "0"
	_, err := models.DelAiforgeRole(roleId)
	if err != nil {
		log.Info("error=" + err.Error())
		resultMap["code"] = "-1"
		resultMap["msg"] = "Del role error."
	} else {
		_, err1 := models.DelAiforgeUserRoleByRoleId(roleId)
		if err1 != nil {
			log.Info("error1=" + err1.Error())
		}
	}
	ctx.JSON(200, resultMap)
}

func UpdateAiforgeRole(ctx *context.APIContext) {
	resultMap := make(map[string]string, 0)
	resultMap["code"] = "0"
	roleId := ctx.QueryInt64("id")
	log.Info("id=" + fmt.Sprint(roleId))

	re, dberr := models.QueryAiforgeRole(roleId)
	if dberr != nil {
		log.Info("error=" + dberr.Error())
		resultMap["code"] = "-1"
		resultMap["msg"] = "Not found this role."
		return
	}
	roleType := re.Type
	description := ctx.Query("description")
	temp := []rune(description)
	if len(temp) > 800 {
		resultMap["code"] = "-1"
		resultMap["msg"] = "角色描述长度不能超过800个字符。"
		ctx.JSON(200, resultMap)
		return
	}
	name := ctx.Query("name")
	if name == "" {
		name = re.Name
	} else {
		if name != re.Name {
			exist, err := models.QueryAiforgeRoleByName(name)
			if err == nil {
				if len(exist) > 0 {
					resultMap["code"] = "-1"
					resultMap["msg"] = "角色名称已经存在。"
					ctx.JSON(200, resultMap)
					return
				}
			}
			temp := []rune(name)
			if len(temp) > 80 {
				resultMap["code"] = "-1"
				resultMap["msg"] = "角色名称长度不能超过80个字符。"
				ctx.JSON(200, resultMap)
				return
			}
		}
	}
	rightsJson := getRightJson(roleType, ctx)

	ar := models.AiforgeRole{
		Name:        name,
		ID:          roleId,
		Description: description,
		RightInfo:   rightsJson,
	}
	err := models.UpdateAiforgeRole(&ar)
	if err == nil {
		log.Info("succeed update aiforge role.")
	} else {
		log.Info("error=" + err.Error())
		resultMap["code"] = "-1"
		resultMap["msg"] = "update role error."
	}
	ctx.JSON(200, resultMap)
}

func QueryAiforgeRole(ctx *context.APIContext) {

}

func ListAiforgeRole(ctx *context.APIContext) {

	name := ctx.Query("name")

	roleList, err := models.ListAiforgeRole(name)
	userIds := make([]int64, len(roleList))
	for i, role := range roleList {
		userIds[i] = role.CreatedUserId
	}

	list, err := resource.GetAllResourceSpecification(models.SearchResourceSpecificationOptions{
		AvailableCode: 1,
		AccCardsNum:   -1,
	})
	specMap := make(map[string]*models.ResourceSpecInfo, 0)
	if err == nil {
		for _, v := range list {
			specMap[fmt.Sprint(v.ID)] = v
		}
	}

	userNameMap := models.QueryUserName(userIds)

	for _, role := range roleList {
		filterOperRights := make([]*models.RightInfo, 0)
		rightInfo := role.RightInfo
		operRights := make([]*models.RightInfo, 0)
		if rightInfo != "" {
			jsonerr := json.Unmarshal([]byte(rightInfo), &operRights)
			if jsonerr != nil {
				log.Info("un json error=" + jsonerr.Error())
			} else {
				for _, right := range operRights {
					if right.SpecId == "" {
						filterOperRights = append(filterOperRights, right)
					} else {
						if specMap[right.SpecId] != nil {
							filterOperRights = append(filterOperRights, right)
						}
					}
				}
			}
		}

		rightsJson, err := json.Marshal(filterOperRights)
		if err != nil {
			log.Info("getRightJson json error." + err.Error())
		}
		role.RightInfo = string(rightsJson)

		value := userNameMap[role.CreatedUserId]
		if value != nil {
			role.UserName = value.Name
			role.UserRelAvatarLink = value.RelAvatarLink()
		}
	}

	if err == nil {
		ctx.JSON(200, roleList)
	} else {
		ctx.JSON(200, make([]*models.AiforgeRole, 0))
	}
}

func ListAiforgeOrgRole(ctx *context.APIContext) {

	name := ctx.Query("name")
	//只允许存储类角色
	roleList, err := models.ListNonCommonAiforgeRoleByNameAndType(name, 2)
	userIds := make([]int64, len(roleList))
	for i, role := range roleList {
		userIds[i] = role.CreatedUserId
	}

	userNameMap := models.QueryUserName(userIds)

	for _, role := range roleList {
		value := userNameMap[role.CreatedUserId]
		if value != nil {
			role.UserName = value.Name
			role.UserRelAvatarLink = value.RelAvatarLink()
		}
	}

	if err == nil {
		ctx.JSON(200, roleList)
	} else {
		ctx.JSON(200, make([]*models.AiforgeRole, 0))
	}
}

func ListResourceQuene(ctx *context.APIContext) {
	re, err := models.GetAccCardInfoByAccCardTypeAndQueueTypeList()
	if err == nil {
		ctx.JSON(200, re)
	} else {
		ctx.JSON(200, make([]models.ResourceQueue, 0))
	}
}

func ListRightUser(ctx *context.APIContext) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	if pageSize == 0 {
		pageSize = setting.UI.Admin.UserPagingNum
	}
	userName := ctx.Query("userName")
	userId := ctx.QueryInt64("userId")
	operRoleId := ctx.QueryInt64("operRoleId")
	resourceRoleId := ctx.QueryInt64("resourceRoleId")
	storageRoleId := ctx.QueryInt64("storageRoleId")
	orderBy := ctx.Query("orderBy")

	opts := models.RightUserOptions{
		Page:           page,
		PageSize:       pageSize,
		OperRoleId:     operRoleId,
		ResourceRoleId: resourceRoleId,
		StorageRoleId:  storageRoleId,
		UserName:       userName,
		UserId:         userId,
		OrderBy:        orderBy,
	}
	users, count, err := models.GetRightUserList(opts)
	// users, _, err := models.SearchUsers(&models.SearchUserOptions{
	// 	Type:    models.UserTypeIndividual,
	// 	OrderBy: models.SearchOrderByRecentRightUpdated,
	// 	ListOptions: models.ListOptions{
	// 		Page:     page,
	// 		PageSize: pageSize,
	// 	},
	// })
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "ListRightUser", err)
		return
	}

	results := make([]*RightUser, len(users))
	for i, user := range users {
		results[i] = &RightUser{
			ID:          user.ID,
			UserName:    user.Name,
			AvatarURL:   user.AvatarLink(),
			FullName:    markup.Sanitize(user.FullName),
			Update:      user.UpdatedUnix.AsTime(),
			Created:     user.CreatedUnix.AsTime(),
			RightUpdate: user.RightUnix.AsTime(),
		}
		results[i].OperRole, results[i].ResourceRole, results[i].StorageRole = role.QueryUserRole(user.ID)
	}
	resultMap := make(map[string]interface{}, 0)
	resultMap["count"] = count
	resultMap["data"] = results
	ctx.JSON(http.StatusOK, resultMap)
}

func ListRightOrg(ctx *context.APIContext) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	if pageSize == 0 {
		pageSize = setting.UI.Admin.UserPagingNum
	}
	userName := ctx.Query("userName")
	userId := ctx.QueryInt64("userId")
	storageRoleId := ctx.QueryInt64("storageRoleId")
	orderBy := ctx.Query("orderBy")

	opts := models.RightUserOptions{
		Page:          page,
		PageSize:      pageSize,
		StorageRoleId: storageRoleId,
		UserName:      userName,
		UserId:        userId,
		OrderBy:       orderBy,
	}
	users, count, err := models.GetRightOrgList(opts)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "ListRightUser", err)
		return
	}

	results := make([]*RightUser, len(users))
	for i, user := range users {
		results[i] = &RightUser{
			ID:          user.ID,
			UserName:    user.Name,
			AvatarURL:   user.AvatarLink(),
			FullName:    markup.Sanitize(user.FullName),
			Update:      user.UpdatedUnix.AsTime(),
			Created:     user.CreatedUnix.AsTime(),
			RightUpdate: user.RightUnix.AsTime(),
		}
		results[i].OperRole, results[i].ResourceRole, results[i].StorageRole = role.QueryOrgRole(user.ID)
	}
	resultMap := make(map[string]interface{}, 0)
	resultMap["count"] = count
	resultMap["data"] = results
	ctx.JSON(http.StatusOK, resultMap)
}

func BatchDelRoleToUser(ctx *context.APIContext) {
	resultMap := make(map[string]interface{}, 0)
	resultMap["code"] = "0"

	failedUserMap := make(map[int64]string, 0)

	roleId := ctx.Query("roleIds")
	userIdStr := ctx.Query("userIds")
	commonRoleMap := make(map[int64]*models.AiforgeRole, 0)
	if userIdStr != "" {
		userIds := strings.Split(userIdStr, ",")
		roleIds := strings.Split(roleId, ",")
		commonRole, err := models.ListCommonAiforgeRole()
		if err == nil {
			for _, t := range commonRole {
				commonRoleMap[t.ID] = t
			}
		}

		for _, userId := range userIds {
			log.Info("batch del role userId=" + userId)
			userIdInt, err := strconv.ParseInt(userId, 10, 64)

			if err == nil {
				user, err3 := models.GetUserByID(userIdInt)
				if err3 != nil || user == nil {
					failedUserMap[userIdInt] = "用户不存在。"
					continue
				} else {
					if user.IsRestricted || user.ProhibitLogin {
						failedUserMap[userIdInt] = "用户被锁定或者被禁止登录。"
						continue
					}
				}

				dbRoles, err2 := models.QueryAiforgeUserRoleByUserId(userIdInt)
				if err2 != nil {
					log.Info("query user role relation error." + err2.Error())
					continue
				}
				dbRoleMap := make(map[int64]*models.AiforgeUserRole, 0)
				for _, v := range dbRoles {
					dbRoleMap[v.RoleId] = v
				}
				for _, ts := range roleIds {
					tRoleId, _ := strconv.ParseInt(ts, 10, 64)
					tCommonRole := commonRoleMap[tRoleId]
					if tCommonRole == nil {
						//付费角色
						tDbRole := dbRoleMap[tRoleId]
						if tDbRole != nil {
							//数据库有则删除
							models.DelAiforgeUserRole(tDbRole.ID)
						}
					} else {
						//取消通用角色
						tDbRole := dbRoleMap[tRoleId]
						if tDbRole == nil {
							ur := models.AiforgeUserRole{
								UserId: userIdInt,
								RoleId: tRoleId,
								Type:   models.MinusType,
							}
							models.AddAiforgeUserRole(ur)
						}
					}
				}
				tmpUser := &models.User{
					ID:        userIdInt,
					RightUnix: timeutil.TimeStamp(time.Now().Unix()),
				}
				models.UpdateUserCols(tmpUser, "right_unix")
			}
		}
	}
	if len(failedUserMap) > 0 {
		resultMap["code"] = "-1"
		resultMap["failedUsers"] = failedUserMap
	}
	ctx.JSON(200, resultMap)
}

func BatchAddRoleToUser(ctx *context.APIContext) {
	resultMap := make(map[string]interface{}, 0)
	resultMap["code"] = "0"

	failedUserMap := make(map[int64]string, 0)

	roleId := ctx.Query("roleIds")
	userIdStr := ctx.Query("userIds")
	commonRoleMap := make(map[int64]*models.AiforgeRole, 0)
	if userIdStr != "" {
		userIds := strings.Split(userIdStr, ",")
		roleIds := strings.Split(roleId, ",")
		commonRole, err := models.ListCommonAiforgeRole()
		if err == nil {
			for _, t := range commonRole {
				commonRoleMap[t.ID] = t
			}
		}

		for _, userId := range userIds {
			log.Info("batch add role userId=" + userId)
			userIdInt, err := strconv.ParseInt(userId, 10, 64)
			if err == nil {
				user, err3 := models.GetUserByID(userIdInt)
				if err3 != nil || user == nil {
					failedUserMap[userIdInt] = "用户不存在。"
					continue
				} else {
					if user.IsRestricted || user.ProhibitLogin {
						failedUserMap[userIdInt] = "用户被锁定或者被禁止登录。"
						continue
					}
				}

				dbRoles, err2 := models.QueryAiforgeUserRoleByUserId(userIdInt)
				if err2 != nil {
					log.Info("query user role relation error." + err2.Error())
					continue
				}
				dbRoleMap := make(map[int64]*models.AiforgeUserRole, 0)
				for _, v := range dbRoles {
					dbRoleMap[v.RoleId] = v
				}
				for _, ts := range roleIds {
					tRoleId, _ := strconv.ParseInt(ts, 10, 64)
					tCommonRole := commonRoleMap[tRoleId]
					if tCommonRole == nil {
						//付费角色
						tDbRole := dbRoleMap[tRoleId]
						if tDbRole == nil {
							//数据库中没有，插入，如果有则跳过
							ur := models.AiforgeUserRole{
								UserId: userIdInt,
								RoleId: tRoleId,
								Type:   models.CommonType,
							}
							models.AddAiforgeUserRole(ur)
						}
					} else {
						//通用角色
						tDbRole := dbRoleMap[tRoleId]
						if tDbRole != nil {
							if tDbRole.Type == models.MinusType {
								//如果通用角色被取消，现在要加回来。加回来直接删除掉记录即可。
								models.DelAiforgeUserRole(tDbRole.ID)
							}
						}
					}
				}
				tmpUser := &models.User{
					ID:        userIdInt,
					RightUnix: timeutil.TimeStamp(time.Now().Unix()),
				}
				models.UpdateUserCols(tmpUser, "right_unix")
			}
		}
	}

	if len(failedUserMap) > 0 {
		resultMap["code"] = "-1"
		resultMap["failedUsers"] = failedUserMap
	}
	ctx.JSON(200, resultMap)
}

func SetAiforgeRoleToUser(ctx *context.APIContext) {
	resultMap := make(map[string]string, 0)
	resultMap["code"] = "0"

	roleId := ctx.Query("roleIds")
	userIdStr := ctx.Query("userIds")

	if userIdStr != "" {
		userIds := strings.Split(userIdStr, ",")
		roleIds := strings.Split(roleId, ",")
		commonRole, err := models.ListCommonAiforgeRole()
		minusTypeList := make([]int64, 0)
		addTypeList := make([]int64, 0)
		if err == nil {
			//检查通用角色是否被取消了，如果被取消了，则要新增记录
			for _, t := range commonRole {
				find := false
				for _, ts := range roleIds {
					if fmt.Sprint(t.ID) == ts {
						find = true
						break
					}
				}
				if !find {
					minusTypeList = append(minusTypeList, t.ID)
				}
			}
			//非通用角色，需要新增记录
			for _, ts := range roleIds {
				find := false
				tRoleId, _ := strconv.ParseInt(ts, 10, 64)
				for _, t := range commonRole {
					if t.ID == tRoleId {
						find = true
						break
					}
				}
				if !find {
					addTypeList = append(addTypeList, tRoleId)
				}
			}
		}

		for _, userId := range userIds {
			log.Info("insert userId=" + userId)
			userIdInt, err := strconv.ParseInt(userId, 10, 64)

			if err == nil {
				_, err1 := models.DelAiforgeUserRoleByUserId(userIdInt)
				if err1 != nil {
					log.Info("delete user role relation error." + err1.Error())
				}
				for _, t := range minusTypeList {
					ur := models.AiforgeUserRole{
						UserId: userIdInt,
						RoleId: t,
						Type:   models.MinusType,
					}
					models.AddAiforgeUserRole(ur)
				}

				for _, t := range addTypeList {
					ur := models.AiforgeUserRole{
						UserId: userIdInt,
						RoleId: t,
						Type:   models.CommonType,
					}
					models.AddAiforgeUserRole(ur)
				}

				tmpUser := &models.User{
					ID:        userIdInt,
					RightUnix: timeutil.TimeStamp(time.Now().Unix()),
				}
				models.UpdateUserCols(tmpUser, "right_unix")
			}
		}
	}

	ctx.JSON(200, resultMap)
}

func SetAiforgeRoleToOrg(ctx *context.APIContext) {
	resultMap := make(map[string]string, 0)
	resultMap["code"] = "0"

	roleId := ctx.Query("roleIds")
	userIdStr := ctx.Query("userIds")

	if userIdStr == "" {
		ctx.JSON(200, resultMap)
		return
	}

	userIds := strings.Split(userIdStr, ",")
	roleIds := strings.Split(roleId, ",")
	commonRole, err := models.ListCommonAiforgeRole()
	if err != nil {
		log.Error("ListCommonAiforgeRole err.%v", err)
		ctx.JSON(200, resultMap)
		return
	}
	availableRoleId := make([]int64, 0, len(roleIds))
	//只添加非默认角色
	for _, ts := range roleIds {
		find := false
		tRoleId, _ := strconv.ParseInt(ts, 10, 64)
		for _, t := range commonRole {
			if t.ID == tRoleId {
				find = true
				break
			}
		}
		if find {
			continue
		}
		availableRoleId = append(availableRoleId, tRoleId)
	}

	for _, userId := range userIds {
		log.Info("insert userId=" + userId)
		userIdInt, err := strconv.ParseInt(userId, 10, 64)

		if err != nil {
			log.Error("userid format error.id=%s", userId)
			ctx.JSON(200, resultMap)
			return

		}
		_, err1 := models.DelAiforgeUserRoleByUserId(userIdInt)
		if err1 != nil {
			log.Info("delete user role relation error." + err1.Error())
			ctx.JSON(200, resultMap)
			return
		}

		for _, t := range availableRoleId {
			ur := models.AiforgeUserRole{
				UserId: userIdInt,
				RoleId: t,
				Type:   models.CommonType,
			}
			models.AddAiforgeUserRole(ur)
		}

		tmpUser := &models.User{
			ID:        userIdInt,
			RightUnix: timeutil.TimeStamp(time.Now().Unix()),
		}
		models.UpdateUserCols(tmpUser, "right_unix")
	}

	ctx.JSON(200, resultMap)
}

func AddUserRole(ctx *context.APIContext, form models.OperateRoleReq) {
	user, err := models.GetUserByName(form.UserName)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError("User not exists"))
		return
	}
	err = role.AddUserRole(user.ID, form.RoleType)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}
func DeleteUserRole(ctx *context.APIContext, form models.OperateRoleReq) {
	user, err := models.GetUserByName(form.UserName)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError("User not exists"))
		return
	}
	err = role.DeleteUserRole(user.ID, form.RoleType)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func AddOperation(ctx *context.APIContext) {
	name := ctx.Query("name")
	description := ctx.Query("description")
	re := models.AiforgeOperation{
		Name:        name,
		Description: description,
	}
	_, err := models.AddRightOperation(re)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, response.OuterSuccess())
	}
}

func DelOperation(ctx *context.APIContext) {
	id := ctx.QueryInt64("id")
	_, err := models.DelRightOperation(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, response.OuterSuccess())
	}
}

func ListOperation(ctx *context.APIContext) {
	re, err := models.ListRightOperation()
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, re)
	}
}
