package user

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"net/http"
)

// OverViewAction 查询用户的概览页面动态
func OverViewAction(ctx *context.APIContext) {
	actions, err := GetFeeds(ctx, models.GetFeedsOptions{
		RequestedUser:  ctx.User,
		Actor:          ctx.User,
		IncludePrivate: true,
	})

	if err != nil {
		log.Error("GetFeeds: %v", err)
		ctx.JSON(http.StatusOK, response.SuccessWithData(actions))
		return
	}

	// 这里的动态数据过多，我们需要裁剪一下
	resp := make([]*models.OverViewAction, 0, len(actions))
	for _, action := range actions {
		putAction := &models.OverViewAction{
			OpType:        action.OpType,
			CreatedUnix:   action.CreatedUnix,
			IsDeleted:     action.IsDeleted,
			RefName:       action.RefName,
			IsPrivate:     action.IsPrivate,
			IsTransformed: action.IsTransformed,
			Content:       action.Content,
		}

		// 用户
		if action.ActUser != nil {
			putAction.UserName = action.ActUser.Name
		}

		// 评论
		if action.Comment != nil {
			putAction.Comment = &models.OverViewComment{
				ID:    action.Comment.ID,
				Issue: action.Comment.Issue,
			}
		}

		// 项目
		if action.Repo != nil {
			putAction.Repo = &models.OverViewRepository{
				ID:        action.Repo.ID,
				Name:      action.Repo.Name,
				Alias:     action.Repo.Alias,
				OwnerName: action.Repo.OwnerName,
			}
		}

		// 数据集
		if action.Dataset != nil && action.Dataset.Owner != nil {
			putAction.Dataset = &models.OverViewDatasetRegistry{
				ID:        action.Dataset.ID,
				Name:      action.Dataset.Name,
				Alias:     action.Dataset.Alias,
				OwnerName: action.Dataset.Owner.Name,
			}
		}

		// 云脑
		if action.Cloudbrain != nil {
			putAction.CloudbrainId = action.Cloudbrain.ID
		}

		if action.Aimodel != nil && action.Aimodel.Owner != nil {
			putAction.Aimodel = &models.OverViewAiModel{
				ID:        action.Aimodel.ID,
				Name:      action.Aimodel.Name,
				Alias:     action.Aimodel.Alias,
				OwnerName: action.Aimodel.Owner.Name,
			}
		}

		resp = append(resp, putAction)
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
	return
}

// GetFeeds 获取动态的数据
func GetFeeds(ctx *context.APIContext, options models.GetFeedsOptions) (actions []*models.Action, err error) {
	log.Info("GetFeeds start reqId[%v] options[%+v]", ctx.ReqId, options)
	actions, err = models.GetFeedsV2(options)
	if err != nil {
		return
	}

	userMap := map[int64]*models.User{options.Actor.ID: options.Actor}
	if options.Actor != nil {
		userMap[options.Actor.ID] = options.Actor
	}
	for _, act := range actions {
		if act.ActUser != nil {
			userMap[act.ActUserID] = act.ActUser
		}
		act.FilterCloudbrainInfo()
	}

	for _, act := range actions {
		if act.Repo != nil {
			repoOwner, ok := userMap[act.Repo.OwnerID]
			if !ok {
				repoOwner, err = models.GetUserByID(act.Repo.OwnerID)
				if err != nil {
					if models.IsErrUserNotExist(err) {
						continue
					}
					return
				}
				userMap[repoOwner.ID] = repoOwner
			}
			act.Repo.Owner = repoOwner
		}
		if act.Dataset != nil {
			datasetOwner, ok := userMap[act.Dataset.OwnerID]
			if !ok {
				datasetOwner, err = models.GetUserByID(act.Dataset.OwnerID)
				if err != nil {
					if models.IsErrUserNotExist(err) {
						continue
					}
					return
				}
				userMap[datasetOwner.ID] = datasetOwner
			}
			act.Dataset.Owner = datasetOwner
		}
		if act.Aimodel != nil {
			aimodelOwner, ok := userMap[act.Aimodel.OwnerID]
			if !ok {
				aimodelOwner, err = models.GetUserByID(act.Aimodel.OwnerID)
				if err != nil {
					if models.IsErrUserNotExist(err) {
						continue
					}
					return
				}
				userMap[aimodelOwner.ID] = aimodelOwner
			}
		}
	}

	return
}
