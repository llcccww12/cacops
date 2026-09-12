package entity

import "code.gitea.io/gitea/models"

var DynConfigType string

const (
	// DynConfigTypePromote is the type for promotion configuration
	DynConfigTypeRaw = "raw"
	// DynConfigTypeDataset is the type for dataset configuration
	DynConfigTypeDataset = "dataset"
	// DynConfigTypeModel is the type for model configuration
	DynConfigTypeModel = "model"
	// DynConfigTypeRepo is the type for repository configuration
	DynConfigTypeRepo = "repo"
	// DynConfigTypeNotice is the type for notice configuration
	DynConfigTypeNotice         = "notice"
	DynConfigTypeAITaskTemplate = "ai_task_template"
)

const (
	ValueTypeObject = "object"
	ValueTypeList   = "list"
)

type DynConfig struct {
	Name       string      `json:"name"`
	ConfigType string      `json:"config_type"`
	ValueType  string      `json:"value_type"`
	Value      interface{} `json:"value"`
}

type PromoteConfig struct {
	Titile     string
	TitleEn    string
	Desc       string
	DescEn     string
	Link       string
	Background string
	//头像图片链接
	Avatar string
}

type DatasetConfig4Show struct {
	ID        string
	Name      string
	Alias     string
	Tags      []string
	Tasks     []string
	License   string
	OwnerName string
	Recommend bool
}

type ModelConfig4Show struct {
	ID           string
	Name         string
	Alias        string
	ExternalName string
	Label        string
	OwnerName    string
	Avatar       string
	ModelType    int
	Recommend    bool
	Engine       int64
}

type RepoConfig4Show struct {
	Name        string
	Description string
	RepoLink    string
	Avatar      string
	NumWatches  int
	NumStars    int
	NumForks    int
}

type AITaskTemplateConfig4Show struct {
	ID            string
	Name          string
	Tags          []string
	ComputeSource string
	JobType       string
	Recommend     bool
	RepoName      string
	RepoOwnerName string
	ImageID       string
	ImageName     string
	ImageUrl      string
	Owner         *models.User4Front
	DatasetLists  []*models.TemplateDatasetInfo
	ModelLists    []*models.TemplateModelInfo
}

type AgentConfigShow struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AgentType  string `json:"agent_type"`
	TemplateID int64  `json:"template_id"`
	Official   bool   `json:"official"` // 是否为官方智能体
}
