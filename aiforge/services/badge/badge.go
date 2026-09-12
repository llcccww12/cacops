package badge

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/admin/operate_log"
	"errors"
)

func GetBadgeList(opts models.GetBadgeOpts) (int64, []*models.Badge4AdminShow, error) {
	total, list, err := models.GetBadgeList(opts)
	if err != nil {
		return 0, nil, err
	}
	if len(list) == 0 {
		return 0, nil, nil
	}
	r := make([]*models.Badge4AdminShow, len(list))
	for i := 0; i < len(list); i++ {
		r[i] = list[i].ToShow()
	}

	return total, r, nil
}

func AddBadge(m models.BadgeOperateReq, doer *models.User) *response.BizError {
	_, err := models.GetBadgeCategoryById(m.CategoryId)

	if err != nil {
		if models.IsErrRecordNotExist(err) {
			return response.NewBizError(errors.New("badge category is not available"))
		}
		return response.NewBizError(err)
	}
	_, err = models.AddBadge(m.ToDTO())
	if err != nil {
		return response.NewBizError(err)
	}
	operate_log.Log4Add(operate_log.BadgeOperate, m, doer.ID, "新增了勋章")
	return nil
}

func EditBadge(m models.BadgeOperateReq, doer *models.User) *response.BizError {
	if m.ID == 0 {
		log.Error(" EditBadge param error")
		return response.NewBizError(errors.New("param error"))
	}
	old, err := models.GetBadgeById(m.ID)
	if err != nil {
		return response.NewBizError(err)
	}
	_, err = models.UpdateBadgeById(m.ID, m.ToDTO())
	if err != nil {
		return response.NewBizError(err)
	}
	operate_log.Log4Edit(operate_log.BadgeOperate, old, m.ToDTO(), doer.ID, "修改了勋章")
	return nil
}

func DelBadge(id int64, doer *models.User) *response.BizError {
	if id == 0 {
		log.Error(" DelBadge param error")
		return response.NewBizError(errors.New("param error"))
	}
	old, err := models.GetBadgeById(id)
	if err != nil {
		return response.NewBizError(err)
	}
	n, _, err := models.GetBadgeUsers(id, models.ListOptions{PageSize: 1, Page: 1})
	if err != nil {
		return response.NewBizError(err)
	}
	if n > 0 {
		return response.BADGES_STILL_HAS_USERS
	}
	_, err = models.DelBadge(id)
	operate_log.Log4Del(operate_log.BadgeOperate, old, doer.ID, "删除了勋章")
	return nil
}
