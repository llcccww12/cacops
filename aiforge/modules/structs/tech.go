package structs

type NotOpenITechRepo struct {
	Url         string   `json:"url"  binding:"Required"`
	TechNo      string   `json:"no"`
	Institution string   `json:"institution"`
	UID         int64    `json:"uid"` //启智项目uid
	RepoName    string   `json:"repo_name" binding:"Required;AlphaDashDot;MaxSize(100)"`
	Alias       string   `json:"alias" binding:"Required;AlphaDashDotChinese;MaxSize(100)"`
	Topics      []string `json:"topics"` //关键词
	Description string   `json:"description" binding:"MaxSize(255)"`
}

type OpenITechRepo struct {
	Url         string `json:"url"  binding:"Required"`
	TechNo      string `json:"no"`
	Institution string `json:"institution"`
}
