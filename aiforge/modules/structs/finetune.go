package structs

type FinetuneJobShow struct {
	ID             int64  `json:"id"`
	JobID          string `json:"job_id"`
	JobType        string `json:"job_type"`
	Type           int    `json:"type"`
	DisplayJobName string `json:"display_job_name"`
	Status         string `json:"status"`
	CreatedUnix    int64  `json:"created_unix"`
	JobCategory    int    `json:"job_category"`
	DeployStatus   string `json:"deploy_status"`
	Cleared        bool   `json:"cleared"`
}
