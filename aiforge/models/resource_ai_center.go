package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type ResourceAICenter struct {
	AICenterCode string `xorm:"pk"`
	AICenterName string
	City         string
	Province     string
	AccessTime   timeutil.TimeStamp
	ComputeScale int
	CreatedTime  timeutil.TimeStamp `xorm:"created"`
	UpdatedTime  timeutil.TimeStamp `xorm:"updated"`
}

type ActiveAICenterReq struct {
	ComputeSource string
	AccCardType   string
}

type AICenterWithCardsUsage struct {
	AICenterCode       string
	TotalCardsUsageNum int
}

type AICenterDetailInfo struct {
	AICenterCode       string
	AICenterName       string
	City               string
	Province           string
	AccessTime         timeutil.TimeStamp
	ComputeScale       int
	TotalCardsUsageNum int
	ComputeSourceList  []string
	AccCardTypeList    []string
}

func (a *AICenterDetailInfo) Tr(language string) {
	a.AICenterName = GetAiCenterShow(a.AICenterCode, a.AICenterName, language)
}

type AICenterDetailInfoList []AICenterDetailInfo

func (a AICenterDetailInfoList) Len() int {
	return len(a)
}

func (a AICenterDetailInfoList) Less(i, j int) bool {
	if a[i].TotalCardsUsageNum > a[j].TotalCardsUsageNum {
		return true
	}
	if a[i].TotalCardsUsageNum < a[j].TotalCardsUsageNum {
		return false
	}
	return a[i].AccessTime > a[j].AccessTime
}

func (a AICenterDetailInfoList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetAICenterInfoMap(aiCenterList []string) (map[string]ResourceAICenter, error) {
	res := make([]ResourceAICenter, 0)
	err := x.Where(builder.In("ai_center_code", aiCenterList)).Find(&res)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]ResourceAICenter, 0)
	for _, v := range res {
		resultMap[v.AICenterCode] = v
	}
	return resultMap, nil
}

func GetResourceAICenter(r *ResourceAICenter) (*ResourceAICenter, error) {
	has, err := x.Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func SyncGrampusAICenters(updateList []ResourceAICenter, insertList []ResourceAICenter) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	//update exists centers
	if len(updateList) > 0 {
		for _, v := range updateList {
			if _, err = sess.Where("ai_center_code = ?", v.AICenterCode).Update(&v); err != nil {
				return err
			}
		}

	}

	//insert new centers
	if len(insertList) > 0 {
		if _, err = sess.Insert(insertList); err != nil {
			return err
		}
	}

	return sess.Commit()
}
