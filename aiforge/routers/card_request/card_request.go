package card_request

import (
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/log"

	api "code.gitea.io/gitea/modules/structs"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	cardrequestservice "code.gitea.io/gitea/services/card_request"
)

func GetCreationInfo(ctx *context.Context) {

	data, err := cardrequestservice.GetCreationInfo()

	if err != nil {
		log.Error("can not get creation info", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.BaseMessageWithDataApi{Data: data})

}

func GetCardRequestList(ctx *context.Context) {

	page := ctx.QueryInt("page")
	if page < 1 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize < 1 {
		pageSize = setting.UI.DatasetPagingNum
	}

	opts := &models.CardRequestOptions{
		OrderBy:  models.OrderByIDDesc,
		NeedSpec: false,
	}
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	getRequestShowList(ctx, opts, false)
}

func GetResourceList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	r := ctx.QueryStrings("resource")
	accCardType := ctx.Query("accCardType")
	accCardNum := ctx.QueryInt("accCardNum")
	excludeAccCardNumStr := ctx.Query("excludeAccCardNums")
	centerCode := ctx.Query("centerCode")
	minPrice := ctx.QueryFloat64("minPrice")
	maxPrice := ctx.QueryFloat64("maxPrice")
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}
	excludeAccCardNums := make([]int, 0)
	if excludeAccCardNumStr != "" {
		numStrArray := strings.Split(excludeAccCardNumStr, "|")
		for _, s := range numStrArray {
			if s == "" {
				continue
			}
			n, err := strconv.Atoi(s)
			if err == nil {
				excludeAccCardNums = append(excludeAccCardNums, n)
			}
		}
	}
	opts := models.GetResourceListOpts{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Resource:           r,
		AccCardType:        accCardType,
		AccCardNum:         accCardNum,
		ExcludeAccCardNums: excludeAccCardNums,
		AICenterCode:       centerCode,
		MinPrice:           minPrice,
		MaxPrice:           maxPrice,
	}
	res, total, err := resource.GetResourceListPagingV2(opts)
	if err != nil {
		log.Error("GetResourceList err.opts=%+v,%v", opts, err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	if res != nil {
		for i := 0; i < len(res); i++ {
			res[i].Tr(ctx.Language())
		}
	}
	resultMap := make(map[string]interface{})
	resultMap["list"] = res
	resultMap["total"] = total
	resultMap["page"] = page
	resultMap["pageSize"] = pageSize
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(resultMap))
}

func GetMyCardRequestList(ctx *context.Context) {

	page := ctx.QueryInt("page")
	if page < 1 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize < 1 {
		pageSize = setting.UI.DatasetPagingNum
	}

	opts := &models.CardRequestOptions{
		UserID:   ctx.User.ID,
		OrderBy:  models.OrderByIDDesc,
		NeedSpec: true,
	}
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	getRequestShowList(ctx, opts, true)
}

func GetAdminCardRequestList(ctx *context.Context) {

	page := ctx.QueryInt("page")
	if page < 1 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize < 1 {
		pageSize = setting.UI.DatasetPagingNum
	}
	useBeginTime, _ := time.Parse(cardrequestservice.DATE_LAYOUT, ctx.Query("useBeginTime"))
	useBeginTime.Unix()

	opts := &models.CardRequestOptions{
		OrderBy:           models.OrderByStatus,
		NeedSpec:          true,
		Keyword:           strings.Trim(ctx.Query("q"), " "),
		AiCenterCode:      ctx.Query("center"),
		Cluster:           ctx.Query("cluster"),
		ComputeResource:   ctx.Query("resource"),
		AccCardType:       ctx.Query("cardType"),
		QueueId:           ctx.QueryInt64("queue"),
		IsResearchProject: ctx.QueryInt("isResearchProject"),
		UseBeginTime:      getTimeUnix(ctx.Query("useBeginTime")),
		UseEndTime:        getTimeUnix(ctx.Query("useEndTime")),
		BeginTimeUnix:     getTimeUnix(ctx.Query("beginTime")),
		EndTimeUnix:       getTimeUnix(ctx.Query("endTime")),
	}
	opts.ListOptions = models.ListOptions{
		Page:     page,
		PageSize: pageSize,
	}
	getRequestShowList(ctx, opts, true)
}

func CreateCardRequest(ctx *context.Context, cardReq api.CardReq) {
	/**data, err := cardrequestservice.GetCreationInfo()
	if err != nil {
		log.Error("can not get creation info", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	if v, ok := data[cardReq.ComputeResource]; !ok {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	} else {
		if !slices.Contains(v, cardReq.CardType) {
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
			return
		}
	}*/

	if cardReq.ResourceType != models.RESOURCE_TYPE_SHARE && cardReq.ResourceType != models.RESOURCE_TYPE_EXCLUSIVE {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	}

	err := cardrequestservice.CreateCardRequest(cardReq, ctx.User.ID)
	if err != nil {
		log.Error("can not create card request", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("card_request.create_fail")))
	} else {
		ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
	}

}

func UpdateCardRequestAndSpec(ctx *context.Context, cardReq api.CardReq) {

	id := ctx.ParamsInt64(":id")
	action := ctx.Query("action")

	cardReq.ID = id
	var err error
	switch action {
	case "agree":
		err = cardrequestservice.AgreeRequest(cardReq)
	case "disagree":
		err = cardrequestservice.DisagreeRequest(cardReq)
	case "modify":
		err = cardrequestservice.UpdateCardRequestAdmin(cardReq)
	}

	if err != nil {
		log.Error("Update  error. %v", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("card_request.update_fail")))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)

}

func UpdateCardRequest(ctx *context.Context, cardReq api.CardReq) {
	id := ctx.ParamsInt64(":id")

	cardRequestInDB, err := models.GetCardRequestById(id)
	if err != nil {
		log.Error("can not get card request record", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("card_request.update_fail_no_record")))
		return
	}
	if ctx.User.ID != cardRequestInDB.UID || cardRequestInDB.Status == models.CARD_REQEST_DISAGREE {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("common_error.insufficient_permission")))
		return
	}

	/**data, err := cardrequestservice.GetCreationInfo()
	if err != nil {
		log.Error("can not get creation info", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	if v, ok := data[cardReq.ComputeResource]; !ok {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	} else {
		if !slices.Contains(v, cardReq.CardType) {
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
			return
		}
	}*/

	if cardReq.ResourceType != models.RESOURCE_TYPE_SHARE && cardReq.ResourceType != models.RESOURCE_TYPE_EXCLUSIVE {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.parameter_is_wrong")))
		return
	}
	cardReq.ID = id
	err = cardrequestservice.UpdateCardRequest(cardReq, cardRequestInDB)
	if err != nil {
		log.Error("Update card request failed", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("card_request.update_fail")))
		return
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
}

func getRequestShowList(ctx *context.Context, opts *models.CardRequestOptions, containsAllParams bool) {
	total, res, err := models.SearchCardRequest(opts)
	if err != nil {
		log.Error("search card request err", err)
		ctx.JSON(http.StatusOK, models.BaseMessageWithDataApi{Data: models.CardRequestShowList{
			CardRequestList: []*models.CardRequestSpecShow{},
		}})
		return
	}

	var show = make([]*models.CardRequestSpecShow, 0)
	for _, v := range res {

		customShow := &models.CardRequestSpecShow{
			ID:              v.ID,
			ComputeResource: v.ComputeResource,
			CardType:        v.CardType,
			AccCardsNum:     v.AccCardsNum,
			BeginDate:       v.BeginDate,
			EndDate:         v.EndDate,
			CreatedUnix:     v.CreatedUnix,
			Status:          v.Status,
		}
		if containsAllParams {
			customShow.UID = v.UID
			customShow.UserName = v.UserName
			customShow.Review = v.Review
			customShow.PhoneNumber = v.PhoneNumber
			customShow.EmailAddress = v.EmailAddress
			customShow.Wechat = v.Wechat
			customShow.IsResearchProject = v.IsResearchProject
			customShow.InstitutionName = v.InstitutionName
			customShow.ProjectName = v.ProjectName
			customShow.ProjectCode = v.ProjectCode
			customShow.Contact = v.Contact
			customShow.Specs = v.Specs
			customShow.Org = v.Org
			customShow.Description = v.Description
			customShow.ResourceType = v.ResourceType
			customShow.DiskCapacity = v.DiskCapacity

		}

		show = append(show, customShow)
	}
	ctx.JSON(http.StatusOK, models.BaseMessageWithDataApi{Data: models.CardRequestShowList{
		Total:           total,
		CardRequestList: show,
	}})
}

func getTimeUnix(value string) int64 {
	timeParse, err := time.Parse(cardrequestservice.DATE_LAYOUT, value)
	if err != nil {
		return 0
	}
	return timeParse.Unix()
}
