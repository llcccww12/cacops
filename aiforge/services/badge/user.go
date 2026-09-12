package badge

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func GetBadgeUsers(badgeId int64, opts models.ListOptions) (int64, []*models.BadgeUser4SHow, error) {
	total, list, err := models.GetBadgeUsers(badgeId, opts)
	if err != nil {
		return 0, nil, err
	}
	if len(list) == 0 {
		return 0, nil, nil
	}
	r := make([]*models.BadgeUser4SHow, len(list))
	for i := 0; i < len(list); i++ {
		r[i] = list[i].ToShow()
	}

	return total, r, nil
}

func AddBadgeUsers(badgeId int64, userIds []int64) (int, []int64, error) {
	errorList := make([]int64, 0)
	if len(userIds) == 0 {
		return 0, errorList, nil
	}
	successCount := 0
	for _, v := range userIds {
		_, err := models.GetUserByID(v)
		if err != nil {
			log.Error("AddBadgeUser get user error in loop, userId=%d. e=%v", v, err)
			errorList = append(errorList, v)
			continue
		}
		m := models.BadgeUser{
			UserId:  v,
			BadgeId: badgeId,
		}
		_, err = models.AddBadgeUser(m)
		if err != nil {
			if models.IsPGErrUniqueViolation(err) {
				//means record exists,user already has the badge
				continue
			}
			log.Error("AddBadgeUser err in loop, m=%+v. e=%v", m, err)
			errorList = append(errorList, v)
			continue
		}
		successCount++
	}
	return successCount, errorList, nil
}

func DelBadgeUser(id int64) error {
	_, err := models.DelBadgeUser(id)
	return err
}

//GetUserBadges Only Returns badges the user has earned
func GetUserBadges(userId int64, opts models.ListOptions) ([]*models.Badge4UserShow, error) {
	badges, err := models.GetUserBadgesPaging(userId, models.GetUserBadgesOpts{ListOptions: opts})
	if err != nil {
		return nil, err
	}
	r := make([]*models.Badge4UserShow, len(badges))
	for i, v := range badges {
		r[i] = v.ToUserShow()
	}
	return r, nil
}

func CountUserBadges(userId int64) (int64, error) {
	return models.CountUserBadges(userId)
}

func GetUserAllBadges(userId int64) ([]models.UserAllBadgeInCategory, error) {
	categoryList, err := models.GetBadgeCategoryList()
	if err != nil {
		return nil, err
	}
	r := make([]models.UserAllBadgeInCategory, 0)
	for _, v := range categoryList {
		badges, err := models.GetBadgeByCategoryId(v.ID)
		if badges == nil || len(badges) == 0 {
			continue
		}
		userBadgeMap, err := getUserBadgesMap(userId, v.ID)
		if err != nil {
			return nil, err
		}
		t := models.UserAllBadgeInCategory{
			CategoryName: v.Name,
			CategoryId:   v.ID,
			LightedNum:   len(userBadgeMap),
		}
		bArray := make([]*models.BadgeShowWithStatus, len(badges))
		for j, v := range badges {
			b := &models.BadgeShowWithStatus{Badge: v.ToUserShow()}
			if _, has := userBadgeMap[v.ID]; has {
				b.IsLighted = true
			}
			bArray[j] = b
		}
		t.Badges = bArray
		r = append(r, t)
	}
	return r, nil
}

func getUserBadgesMap(userId, categoryId int64) (map[int64]*models.Badge, error) {
	userBadges, err := models.GetUserBadges(userId, categoryId)
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*models.Badge, 0)
	for _, v := range userBadges {
		m[v.ID] = v
	}
	return m, nil
}
