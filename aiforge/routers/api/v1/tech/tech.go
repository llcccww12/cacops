package tech

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/role"
	techService "code.gitea.io/gitea/services/tech"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetFilterInfo(ctx *context.APIContext) {
	filterType := ctx.QueryInt("type")
	limit := ctx.QueryInt("limit")
	if filterType == 0 {
		projectTypes := models.GetProjectTypes()
		allInstitutions := models.GetAllInstitutions()
		applyYears, executeYears := models.GetApplyExecuteYears()
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"type_name":        projectTypes,
			"institution_name": allInstitutions,
			"execute_year":     executeYears,
			"apply_year":       applyYears,
		})
		return
	}
	if filterType == 1 {
		projectNames := models.GetProjectNames()
		topics := models.GetTechRepoTopics(limit)
		allInstitutions := models.GetAllInstitutions()
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"topic":            topics,
			"institution_name": allInstitutions,
			"project_name":     projectNames,
		})
		return

	}
	ctx.Error(http.StatusBadRequest, "parameter error", "")

}

func GetAdminRepoInfo(ctx *context.APIContext) {
	go techService.UpdateTechStatus()
	opts := &models.SearchUserRepoOpt{}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	opts.Page = page
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = 15
	}
	opts.PageSize = pageSize
	infos, total, err := techService.GetAdminRepoInfo(opts)
	if err != nil {
		log.Error("search has error", err)
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": 0, "data": []*models.TechRepoInfoAdmin{}})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": total, "data": infos})
	}

}

func Action(ctx *context.APIContext, ids ActionIDs) {
	if len(ids.ID) == 0 {
		ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
	}
	var err error
	switch ctx.Params(":action") {

	case "show":
		err = models.ShowTechRepo(ids.ID, true)
	case "hide":
		err = models.ShowTechRepo(ids.ID, false)
	}
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("tech.show_or_hide_fail", ctx.Params(":action"))))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}

}

type ActionIDs struct {
	ID []int64 `json:"id"`
}

func GetMyRepoInfo(ctx *context.APIContext) {
	go techService.UpdateTechStatus()
	opts := &models.SearchUserRepoOpt{}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	opts.Page = page
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = 15
	}
	opts.PageSize = pageSize
	opts.User = ctx.User
	infos, total, err := techService.GetMyRepoInfo(opts)
	if err != nil {
		log.Error("search has error", err)
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": 0, "data": []*models.TechRepoInfoUser{}})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": total, "data": infos})
	}

}

func SearchRepoInfo(ctx *context.APIContext) {
	go techService.UpdateTechStatus()
	opts := &models.SearchRepoOpt{}
	opts.Q = ctx.Query("name")
	opts.ProjectName = ctx.Query("tech_name")
	opts.Topic = ctx.Query("topic")
	opts.Institution = ctx.Query("institution_name")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	opts.Page = page
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = 30
	}
	opts.PageSize = pageSize
	orderBy := ctx.Query("sort")
	opts.OrderBy = orderBy
	infos, total, err := techService.SearchRepoInfoWithInstitution(opts)
	if err != nil {
		log.Error("search has error", err)
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": 0, "data": []*models.RepoWithInstitution{}})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": total, "data": infos})
	}

}

func SearchTechProjectInfo(ctx *context.APIContext) {
	go techService.UpdateTechStatus()
	opts := &models.SearchTechOpt{}
	opts.Q = ctx.Query("name")
	opts.ApplyYear = ctx.QueryInt("apply_year")
	opts.ExecuteYear = ctx.QueryInt("execute_year")
	opts.Institution = ctx.Query("institution_name")
	opts.ProjectType = ctx.Query("type_name")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	opts.Page = page
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = 15
	}
	opts.PageSize = pageSize
	orderBy := ctx.Query("sort")
	if orderBy == "" || orderBy == "count" {
		opts.OrderBy = "order by repo_count desc, max desc "
	} else {
		opts.OrderBy = "order by max desc, repo_count desc "
	}

	infos, total, err := models.SearchTechRepoInfo(opts)
	if err != nil {

		log.Error("search has error", err)
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": 0, "data": []*models.TechRepoInfo{}})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface {
		}{"total": total, "data": infos})
	}

}

func ImportBasicInfo(ctx *context.APIContext) {
	file, _, err := ctx.GetFile("file")
	if err != nil {
		log.Error("get file err", err)
		ctx.JSON(http.StatusBadRequest, models.BaseErrorMessageApi(ctx.Tr("tech.get_file_fail")))
		return
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Error("open file err", err)
		ctx.JSON(http.StatusBadRequest, models.BaseErrorMessageApi(ctx.Tr("tech.content_type_unsupported")))
		return
	}
	rows, err := f.GetRows(f.GetSheetName(1))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseErrorMessageApi(ctx.Tr("tech.content_type_unsupported")))
		return
	}
	for i, row := range rows {
		if i != 0 {
			baseInfo := &models.TechConvergeBaseInfo{}
			baseInfo.ProjectNumber = strings.TrimSpace(row[2])
			tbInDB, err := models.GetTechConvergeBaseInfoByProjectNumber(baseInfo.ProjectNumber)
			if err != nil && !models.IsErrTechConvergeBaseInfoNotExist(err) {
				log.Error("query err ", i, err)
				ctx.JSON(http.StatusBadRequest, models.BaseErrorMessageApi(ctx.Tr("tech.sql_err")))
				return
			}
			if tbInDB != nil {
				baseInfo = tbInDB
			}

			baseInfo.ApplyYear, err = strconv.Atoi(strings.TrimSpace(row[1]))
			if err != nil {
				log.Warn("base info apply year parse err ", i)
			}
			baseInfo.ProjectName = strings.TrimSpace(row[3])
			baseInfo.Type = strings.TrimSpace(row[16])
			baseInfo.Owner = strings.TrimSpace(row[8])
			baseInfo.Email = strings.TrimSpace(row[10])
			baseInfo.Phone = strings.TrimSpace(row[9])
			baseInfo.Recommend = strings.TrimSpace(row[7])
			baseInfo.Institution = strings.TrimSpace(row[4])
			baseInfo.Category = strings.TrimSpace(row[6])
			baseInfo.Province = strings.TrimSpace(row[5])
			baseInfo.Contact = strings.TrimSpace(row[11])
			baseInfo.ContactPhone = strings.TrimSpace(row[12])
			baseInfo.ContactEmail = strings.TrimSpace(row[13])
			baseInfo.ExecuteMonth, _ = strconv.Atoi(strings.TrimSpace(row[14]))
			baseInfo.ExecutePeriod = strings.TrimSpace(row[15])
			baseInfo.ExecuteStartYear, baseInfo.ExecuteEndYear = getBeginEndYear(baseInfo.ExecutePeriod)
			baseInfo.StartUp = strings.TrimSpace(row[17])
			baseInfo.NumberTopic, _ = strconv.Atoi(strings.TrimSpace(row[30]))
			baseInfo.Topic1 = strings.TrimSpace(row[31])
			baseInfo.Topic2 = strings.TrimSpace(row[32])
			baseInfo.Topic3 = strings.TrimSpace(row[33])
			baseInfo.Topic4 = strings.TrimSpace(row[34])
			baseInfo.Topic5 = strings.TrimSpace(row[35])

			allIns := replaceNewLineWithComma(strings.TrimSpace(row[36]))
			if allIns != "" {
				allInsArray := strings.Split(allIns, ",")
				var allInsValid []string
				for _, ins := range allInsArray {
					if ins != "" {
						allInsValid = append(allInsValid, ins)
					}
				}
				baseInfo.AllInstitution = strings.Join(allInsValid, ",")
				if baseInfo.AllInstitution == "" {
					baseInfo.AllInstitution = baseInfo.Institution
				}
			} else {
				baseInfo.AllInstitution = baseInfo.Institution
			}

			err = baseInfo.InsertOrUpdate()
			if err != nil {
				log.Error("update err", i, err)
				ctx.JSON(http.StatusBadRequest, models.BaseErrorMessageApi(ctx.Tr("tech.sql_err")))
				return
			}

		}
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
}

func replaceNewLineWithComma(s string) string {
	const replacement = ","

	var replacer = strings.NewReplacer(
		"\r\n", replacement,
		"\r", replacement,
		"\n", replacement,
		"\v", replacement,
		"\f", replacement,
		"\u0085", replacement,
		"\u2028", replacement,
		"\u2029", replacement,
	)
	return replacer.Replace(s)
}

/**
  input:2019.12-2023.12   output: 2019,2023
*/
func getBeginEndYear(executePeriod string) (int, int) {
	period := strings.Split(executePeriod, "-")
	if len(period) == 2 {
		start := strings.Split(period[0], ".")
		end := strings.Split(period[1], ".")

		if len(start) == 2 && len(end) == 2 {
			startYear, err := strconv.Atoi(start[0])
			endYear, err1 := strconv.Atoi(end[0])
			if err == nil && err1 == nil {
				return startYear, endYear
			}

		}

	}
	return 0, 0
}

func FindTech(ctx *context.APIContext) {
	techNo := ctx.Query("no")
	name := ctx.Query("name")
	institution := ctx.Query("institution")
	r, err := techService.FindTech(models.FindTechOpt{TechNo: techNo, ProjectName: name, Institution: institution})
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}

func IsAdmin(ctx *context.APIContext) {
	//isAdmin := role.UserHasRole(ctx.User.ID, models.TechProgramAdmin)
	isAdmin := role.UserHasOper(ctx.User.ID, role.ROLE_OPER_TechProgramAdmin)
	r := map[string]interface{}{}
	r["is_admin"] = isAdmin
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}
