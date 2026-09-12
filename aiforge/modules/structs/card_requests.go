package structs

type CardReq struct {
	ID                int64   `json:"id"`
	ComputeResource   string  `json:"compute_resource" binding:"Required"`
	CardType          string  `json:"card_type" binding:"Required"`
	AccCardsNum       string  `json:"acc_cards_num" binding:"Required"`
	DiskCapacity      int64   `json:"disk_capacity"`
	ResourceType      int     `json:"resource_type" binding:"Required"`
	BeginDate         string  `json:"begin_date" binding:"Required"`
	EndDate           string  `json:"end_date" binding:"Required"`
	Contact           string  `json:"contact" binding:"Required"`
	PhoneNumber       string  `json:"phone_number" binding:"Required"`
	EmailAddress      string  `json:"email_address" binding:"Required;Email;MaxSize(254)"`
	Wechat            string  `json:"wechat" binding:"Required;MaxSize(254)"`
	IsResearchProject bool    `json:"is_research_project"`
	InstitutionName   string  `json:"institution_name" binding:"MaxSize(254)"`
	ProjectName       string  `json:"project_name" binding:"MaxSize(254)"`
	ProjectCode       string  `json:"project_code" binding:"MaxSize(254)"`
	Org               string  `json:"org" binding:"MaxSize(500)"`
	Description       string  `json:"description" binding:"MaxSize(3000)"`
	Review            string  `json:"review"`
	SpecIds           []int64 `json:"spec_ids"`
}

type RequestSpecInfo struct {
	ID           int64
	SourceSpecId string
	AccCardsNum  int
	CpuCores     int
	MemGiB       float32
	GPUMemGiB    float32
	ShareMemGiB  float32
	UnitPrice    int
	Status       int
	UpdatedTime  int64
	RequestId    int64
	//queue
	Cluster         string
	AiCenterCode    string
	AiCenterName    string
	QueueCode       string
	QueueId         int64
	ComputeResource string
	AccCardType     string
	HasInternet     int
}
