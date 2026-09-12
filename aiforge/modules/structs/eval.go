package structs

type EvalRequest struct {
	ModelName string `json:"model"`
	Datasets  string `json:"datasets"`
	TaskId    int64  `json:"task_id"`
	Limit     int    `json:"limit"`
}

type TaskDetailRequest struct {
	TaskId   int64  `json:"task_id"`
	Filter   string `json:"filter"`
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
	Result   int    `json:"result"`
}

type TaskRequest struct {
	TaskId int64 `json:"task_id"`
}

type EvalResultResponse struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Result  []EvalResult `json:"result"`
}

type EvalResult struct {
	Name    string  `json:"name"`
	Class   string  `json:"class"`
	ClassEN string  `json:"classen"`
	Value   float32 `json:"value"`
}

type EvalDetailResultResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Count   int                `json:"count"`
	Data    []EvalDetailResult `json:"data"`
}
type EvalDetailResult struct {
	Gold     string `json:"gold"`
	Pred     string `json:"pred"`
	Predict  string `json:"predict"`
	Type     string `json:"type"`
	RawInput string `json:"raw_input"`
}
