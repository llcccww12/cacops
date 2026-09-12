package entity

type ResourceUsage struct {
	Interval    int           `json:"interval"`
	MetricsInfo []MetricsInfo `json:"metrics_info"`
}

type MetricsInfo struct {
	Name  string    `json:"name"`
	Value []float32 `json:"value"`
}
