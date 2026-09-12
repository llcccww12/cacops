package badge

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/admin/operate_log"
	"errors"
)

func GetBadgeCategoryList(opts models.ListOptions) (int64, []*models.BadgeCategory4Show, error) {
	total, list, err := models.GetBadgeCategoryListPaging(opts)
	if err != nil {
		return 0, nil, err
	}
	if len(list) == 0 {
		return 0, nil, nil
	}
	r := make([]*models.BadgeCategory4Show, len(list))
	for i := 0; i < len(list); i++ {
		r[i] = list[i].ToShow()
	}

	return total, r, nil
}

func AddBadgeCategory(m models.BadgeCategory4Show, doer *models.User) *response.BizError {
	_, err := models.AddBadgeCategory(m.ToDTO())
	if err != nil {
		return response.NewBizError(err)
	}
	operate_log.Log4Add(operate_log.BadgeCategoryOperate, m, doer.ID, "新增了勋章分类")
	return nil
}

func EditBadgeCategory(m models.BadgeCategory4Show, doer *models.User) *response.BizError {
	if m.ID == 0 {
		log.Error(" EditBadgeCategory param error")
		return response.NewBizError(errors.New("param error"))
	}
	old, err := models.GetBadgeCategoryById(m.ID)
	if err != nil {
		return response.NewBizError(err)
	}
	_, err = models.UpdateBadgeCategoryById(m.ID, m.ToDTO())
	if err != nil {
		return response.NewBizError(err)
	}
	operate_log.Log4Edit(operate_log.BadgeCategoryOperate, old, m.ToDTO(), doer.ID, "修改了勋章分类")
	return nil
}

func DelBadgeCategory(id int64, doer *models.User) *response.BizError {
	if id == 0 {
		log.Error(" DelBadgeCategory param error")
		return response.NewBizError(errors.New("param error"))
	}
	old, err := models.GetBadgeCategoryById(id)
	if err != nil {
		return response.NewBizError(err)
	}
	badges, err := models.GetBadgeByCategoryId(id)
	if err != nil {
		return response.NewBizError(err)
	}
	if len(badges) > 0 {
		return response.CATEGORY_STILL_HAS_BADGES
	}
	_, err = models.DelBadgeCategory(id)
	operate_log.Log4Del(operate_log.BadgeCategoryOperate, old, doer.ID, "删除了勋章分类")
	return nil
}
