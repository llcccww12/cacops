package models

// OfficialConfig 概览结构体
type OfficialConfig struct {
	NoticeInfo    NoticeOverview        `json:"notice_info"`         // 公告
	ActivityImage ActivityImageOverview `json:"activity_image_info"` // 活动图片
}

// NoticeInfo 公告和notice\notice.go相对应
type NoticeInfo struct {
	Title   string `json:"title"`             // Title 公告标题（中文）
	TitleEn string `json:"title_en"`          // TitleEn 公告标题（英文）
	Link    string `json:"link"`              // Link 公告详情链接
	Visible int    `json:"visible,omitempty"` // 0 invisible, 1 visible
	Date    string `json:"date"`              // Date 公告发布日期
}

type NoticeOverview struct {
	Notices  []*NoticeInfo `json:"notices"` // 公告列表
	CommitId string        `json:"commit_id,omitempty"`
}

// ActivityImageOverview 活动图片对象
type ActivityImageOverview struct {
	ImageUrl        string `json:"image_url"`                   // 图片地址
	ImageLink       string `json:"image_link"`                  // 图片链接
	InviteImageUrl  string `json:"invite_image_url,omitempty"`  // 邀请图片地址
	InviteImageLink string `json:"invite_image_link,omitempty"` // 邀请图片链接
}
