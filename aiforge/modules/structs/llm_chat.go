package structs

import "mime/multipart"

type LLMChatHistory struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMChatMessage struct {
	ModelName string           `json:"model_name" binding:"Required"`
	Query     string           `json:"query" binding:"Required"`
	History   []LLMChatHistory `json:"history"`
	Stream    bool             `json:"stream" `
}

type KBChatMessage struct {
	ModelName         string           `json:"model_name" binding:"Required"`
	Query             string           `json:"query" binding:"Required"`
	KnowledgeBaseName string           `json:"knowledge_base_name" binding:"Required"`
	History           []LLMChatHistory `json:"history"`
	Stream            bool             `json:"stream"`
	TopK              int              `json:"top_k"`
	ScoreThreshold    float64          `json:"score_threshold"`
}

type CreateKnowledgeBaseParams struct {
	KnowledgeBaseName string `json:"knowledge_base_name"`
	VectorStoreType   string `json:"vector_store_type"`
	EmbedModel        string `json:"embed_model"`
}

type SearchDocParams struct {
	Query             string  `json:"query"`
	KnowledgeBaseName string  `json:"knowledge_base_name"`
	TopK              int     `json:"top_k"`
	ScoreThreshold    float64 `json:"score_threshold"`
}

type DeleteDocParams struct {
	KnowledgeBaseName string   `json:"knowledge_base_name" binding:"Required"`
	FileNames         []string `json:"file_names" binding:"Required"`
	DeleteContent     bool     `json:"delete_content"`
	NotRefreshVsCache bool     `json:"not_refresh_vs_cache"`
}

type UpdateDocParams struct {
	KnowledgeBaseName string `json:"knowledge_base_name"`
	FileNames         string `json:"file_names"`
	NotRefreshVsCache bool   `json:"not_refresh_vs_cache"`
}

type RecreateVectorStoreParams struct {
	KnowledgeBaseName string `json:"knowledge_base_name"`
	AllowEmptyKb      bool   `json:"allow_empty_kb"`
	VsType            string `json:"vs_type"`
	EmbedModel        string `json:"embed_model"`
}

type LLMChatCountsResults struct {
	MaxTries int64 `json:"max_tries"`
	Counts   int64 `json:"counts"`
	CanChat  bool  `json:"can_chat"`
	//FirstVisit bool  `json:"first_visit"`
}

type KBChatAnswer struct {
	Answer string `json:"answer"`
}

type KBChatDocs struct {
	Docs []string `json:"docs"`
}
type LegalTextParams struct {
	Text string `json:"text"`
}

type LLMChatUploadForm struct {
	KnowledgeBaseName string                  `form:"knowledge_base_name"`
	Files             []*multipart.FileHeader `form:"files"`
	Override          bool                    `form:"override"`
}
