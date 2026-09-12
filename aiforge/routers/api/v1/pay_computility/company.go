package pay_computility

// Company 公司基本信息
type Company struct {
	Provider    string  `json:"provider,omitempty"`    // 提供方
	Name        string  `json:"name"`                  // 公司名称
	Logo        string  `json:"logo"`                  // Logo图片路径或URL
	Sort        int     `json:"sort"`                  // 排序值
	Height      float64 `json:"height"`                // 排序值
	Description string  `json:"description,omitempty"` // 描述（可选）
	Website     string  `json:"website,omitempty"`     // 官网（可选）
}

type ResponseCompany struct {
	ComputingPlatform []Company `json:"computing_platform"`
	AiCenter          []Company `json:"ai_center"`
	ChipVendor        []Company `json:"chip_vendor"`
}

func GetCompanies() ResponseCompany {
	return ResponseCompany{
		ComputingPlatform: []Company{
			{
				Provider: "haosuan",
				Name:     "联通昊算",
				Logo:     "",
			},
			{
				Provider: "paratera",
				Name:     "并行科技",
				Logo:     "",
			},
			{
				Provider: "supercomputing-internet",
				Name:     "超算互联网平台",
				Logo:     "",
			},
			{
				Provider: "ctyun",
				Name:     "天翼云",
				Logo:     "",
			},
		},
		AiCenter: []Company{
			{
				Provider: "pku-center",
				Name:     "北大分中心",
				Logo:     "",
			},
			{
				Provider: "guangzhou-supercomputing",
				Name:     "广州超算",
				Logo:     "",
			},
			{
				Provider: "hebei-ai-center",
				Name:     "河北人工智能计算中心",
				Logo:     "",
			},
		},
		ChipVendor: []Company{
			{
				Provider: "huawei",
				Name:     "华为",
				Logo:     "",
			},
			{
				Provider: "tiansu",
				Name:     "天数智芯",
				Logo:     "",
			},
			{
				Provider: "suiyuan",
				Name:     "燧原科技",
				Logo:     "",
			},
			{
				Provider: "haiguang",
				Name:     "海光",
				Logo:     "",
			},
			{
				Provider: "muxi",
				Name:     "沐曦",
				Logo:     "",
			},
			{
				Provider: "cambricon",
				Name:     "寒武纪",
				Logo:     "",
			},
		},
	}
}
