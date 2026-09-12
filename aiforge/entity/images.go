package entity

import "code.gitea.io/gitea/models"

type ClusterImage struct {
	ImageId   string `json:"image_id"`
	ImageName string `json:"image_name"`
	ImageUrl  string `json:"image_url"`
}

type GetImageReq struct {
	ComputeSource models.ComputeSource
	JobType       models.JobType
	AccCardType   string
}
