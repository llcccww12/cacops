package baiduai

type LegalTextResponse struct {
	Conclusion     string `json:"conclusion"`
	LogId          string `json:"log_id"`
	IsHitMd5       bool   `json:"isHitMd5"`
	ConclusionType int    `json:"conclusionType"`
	Data           []Data `json:"data"`
}

type Data struct {
	Msg            string `json:"msg"`
	Conclusion     string `json:"conclusion"`
	SubType        int    `json:"subType"`
	ConclusionType int    `json:"conclusionType"`
	Type           int    `json:"type"`
	Hits           []Hit  `json:"hits"`
}

type Hit struct {
	Probability       int               `json:"probability"`
	DatasetName       string            `json:"datasetName"`
	Words             []string          `json:"words"`
	ModelHitPositions [][]float64       `json:"modelHitPositions"`
	WordHitPositions  []WordHitPosition `json:"wordHitPositions"`
}

type WordHitPosition struct {
	Positions [][]int `json:"positions"`
	Label     string  `json:"label"`
	Keyword   string  `json:"keyword"`
}
