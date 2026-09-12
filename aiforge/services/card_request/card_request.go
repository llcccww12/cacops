package card_request

import (
	"time"

	"code.gitea.io/gitea/models"
	api "code.gitea.io/gitea/modules/structs"
)

const DATE_LAYOUT = "2006-01-02"

func GetCreationInfo() (map[string][]string, error) {

	xpuInfoBase, err := models.GetXPUInfos()

	if err != nil {
		return nil, err
	}

	var xpuInfoMap = make(map[string][]string)

	for _, xpuInfo := range xpuInfoBase {
		if _, ok := xpuInfoMap[xpuInfo.ResourceType]; ok {
			xpuInfoMap[xpuInfo.ResourceType] = append(xpuInfoMap[xpuInfo.ResourceType], xpuInfo.CardType)
		} else {
			xpuInfoMap[xpuInfo.ResourceType] = []string{xpuInfo.CardType}
		}
	}
	return xpuInfoMap, nil

}
func AgreeRequest(cardReq api.CardReq) error {
	return models.AgreeCardRequest(models.CardRequestReview{
		ID:      cardReq.ID,
		SpecIds: cardReq.SpecIds,
	})
}

func DisagreeRequest(cardReq api.CardReq) error {
	return models.DisagreeCardRequest(models.CardRequestReview{
		ID:     cardReq.ID,
		Review: cardReq.Review,
	})
}

func UpdateCardRequestAdmin(cardReq api.CardReq) error {
	request := models.CardRequest{
		ID:                cardReq.ID,
		ComputeResource:   cardReq.ComputeResource,
		CardType:          cardReq.CardType,
		AccCardsNum:       cardReq.AccCardsNum,
		EmailAddress:      cardReq.EmailAddress,
		DiskCapacity:      cardReq.DiskCapacity,
		Contact:           cardReq.Contact,
		PhoneNumber:       cardReq.PhoneNumber,
		Wechat:            cardReq.Wechat,
		IsResearchProject: cardReq.IsResearchProject,
		InstitutionName:   cardReq.InstitutionName,
		ProjectName:       cardReq.ProjectName,
		ProjectCode:       cardReq.ProjectCode,
		BeginDate:         cardReq.BeginDate,
		EndDate:           cardReq.EndDate,
		Description:       cardReq.Description,
		Org:               cardReq.Org,
		ResourceType:      cardReq.ResourceType,
	}
	beginTime, err := time.Parse(DATE_LAYOUT, cardReq.BeginDate)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(DATE_LAYOUT, cardReq.EndDate)
	if err != nil {
		return err

	}
	request.BeginUnix = beginTime.Unix()
	request.EndUnix = endTime.Unix()
	return models.UpdateCardRequest(&request)
}

func UpdateCardRequest(cardReq api.CardReq, request *models.CardRequest) error {

	request.DiskCapacity = cardReq.DiskCapacity
	request.Org = cardReq.Org
	request.Description = cardReq.Description

	if request.Status == models.CARD_REQUEST_COMMIT {

		request.ComputeResource = cardReq.ComputeResource
		request.CardType = cardReq.CardType
		request.AccCardsNum = cardReq.AccCardsNum
		request.ResourceType = cardReq.ResourceType
		request.BeginDate = cardReq.BeginDate
		request.EndDate = cardReq.EndDate

		request.Contact = cardReq.Contact
		request.EmailAddress = cardReq.EmailAddress
		request.PhoneNumber = cardReq.PhoneNumber
		request.Wechat = cardReq.Wechat
		request.IsResearchProject = cardReq.IsResearchProject
		request.InstitutionName = cardReq.InstitutionName
		request.ProjectName = cardReq.ProjectName
		request.ProjectCode = cardReq.ProjectCode

		beginTime, err := time.Parse(DATE_LAYOUT, cardReq.BeginDate)
		if err != nil {
			return err
		}
		endTime, err := time.Parse(DATE_LAYOUT, cardReq.EndDate)
		if err != nil {
			return err

		}
		request.BeginUnix = beginTime.Unix()
		request.EndUnix = endTime.Unix()

	}

	return models.UpdateCardRequest(request)

}

func CreateCardRequest(cardReq api.CardReq, uid int64) error {

	bean := &models.CardRequest{
		UID:               uid,
		ComputeResource:   cardReq.ComputeResource,
		CardType:          cardReq.CardType,
		AccCardsNum:       cardReq.AccCardsNum,
		EmailAddress:      cardReq.EmailAddress,
		DiskCapacity:      cardReq.DiskCapacity,
		Contact:           cardReq.Contact,
		PhoneNumber:       cardReq.PhoneNumber,
		Wechat:            cardReq.Wechat,
		IsResearchProject: cardReq.IsResearchProject,
		InstitutionName:   cardReq.InstitutionName,
		ProjectName:       cardReq.ProjectName,
		ProjectCode:       cardReq.ProjectCode,
		BeginDate:         cardReq.BeginDate,
		EndDate:           cardReq.EndDate,
		Description:       cardReq.Description,
		Org:               cardReq.Org,
		ResourceType:      cardReq.ResourceType,
		Status:            models.CARD_REQUEST_COMMIT,
	}

	beginTime, err := time.Parse(DATE_LAYOUT, cardReq.BeginDate)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(DATE_LAYOUT, cardReq.EndDate)
	if err != nil {
		return err

	}
	bean.BeginUnix = beginTime.Unix()
	bean.EndUnix = endTime.Unix()

	return models.CreateCardRequest(bean)

}
