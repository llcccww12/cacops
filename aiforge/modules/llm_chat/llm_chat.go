package llm_chat

type LLMChatResponse struct {
	Answer string `json:"answer"`
}

type KBChatResponse struct {
	Answer string   `json:"answer"`
	Docs   []string `json:"docs"`
}

type SearchDocResponse struct {
	Results []SearchDocResult `json:"results"`
}

type SearchDocResult struct {
	PageContent string `json:"page_content"`
	Metadata    struct {
	} `json:"metadata"`
	Score float64 `json:"score"`
}

type RecreateVectorStoreResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Total    int    `json:"total"`
	Finished int    `json:"finished"`
	Doc      string `json:"doc"`
}

type LLMBasicMsgWithData struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type LLMDeleteDocMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FailedFiles map[string]string `json:"failed_files"`
	} `json:"data"`
}

type LLMBasicMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LLMErrorMsg struct {
	Detail []struct {
		Loc  []interface{} `json:"loc"`
		Msg  string        `json:"msg"`
		Type string        `json:"type"`
	} `json:"detail"`
}
