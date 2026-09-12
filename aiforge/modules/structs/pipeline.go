package structs

type Pipeline struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type NodeInfo struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PipelineNotification struct {
	Type          int      `json:"type"`
	Username      string   `json:"username"`
	Reponame      string   `json:"reponame"`
	Pipeline      Pipeline `json:"pipeline"`
	PipelineRunId string   `json:"pipeline_run_id"`
	Node          NodeInfo `json:"node"`
	OccurTime     int64    `json:"occur_time"`
}
