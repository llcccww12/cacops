package pay_computility

type ResourceType struct {
	Label      string   `json:"label,omitempty"`
	Category   string   `json:"category,omitempty"`
	CardTypes  []string `json:"cardTypes,omitempty"`  // 联动卡类型
	CardCounts []int64  `json:"cardCounts,omitempty"` // 联动卡数
	PriceUnit  string   `json:"priceUnit,omitempty"`  // 价格单位
	Providers  []string `json:"providers,omitempty"`  // 联动的供应商
}

type Product struct {
	ID           int64   `json:"id,omitempty"`           // 产品序号
	Name         string  `json:"name"`                   // 产品名称，如：RTX4090 24G / 独享
	Description  string  `json:"description,omitempty"`  // 产品描述
	GPU          string  `json:"gpu,omitempty"`          // GPU配置
	CPU          string  `json:"cpu,omitempty"`          // CPU配置
	SourceType   string  `json:"source_type,omitempty"`  // 资源类型 独享
	Memory       string  `json:"memory,omitempty"`       // 内存容量，如：512GB
	Storage      string  `json:"storage,omitempty"`      // 存储配置
	Price        float64 `json:"price,omitempty"`        // 价格，如：6000.00
	PriceUnit    string  `json:"priceUnit,omitempty"`    // 价格单位，如：元/台/月
	PurchaseLink string  `json:"purchaseLink,omitempty"` // 立即购买跳转链接
	Category     string  `json:"category,omitempty"`     // 产品类别：bare_metal/cloud_server/container_cloud
	CardType     string  `json:"cardType,omitempty"`     // 联动卡类型
	CardCount    int64   `json:"cardCount,omitempty"`    // 联动卡数
	Provider     string  `json:"provider,omitempty"`     // 提供方
}

type FilterData struct {
	ResourceTypes []ResourceType `json:"resourceTypes"`
	Products      []Product      `json:"products"`
}

func GetFilterData() FilterData {
	return FilterData{
		ResourceTypes: []ResourceType{
			{
				Label:      "云服务器",
				Category:   "cloud_server",
				CardTypes:  []string{"910B", "A800", "H100", "H20", "H800", "L40", "L40S", "RTX3090", "RTX4090", "Tesla A100", "Tesla V100"},
				CardCounts: []int64{1, 2, 4, 8}, // 云服务器隐藏卡数
				PriceUnit:  "元/卡/时",
			},
			{
				Label:      "裸金属",
				Category:   "bare_metal",
				CardTypes:  []string{"910B", "A800", "H100", "H20", "H800", "L40", "L40S", "RTX3090", "RTX4090"},
				CardCounts: nil,
				PriceUnit:  "元/台/月",
			},
			{
				Label:      "容器云",
				Category:   "container_cloud",
				CardTypes:  []string{"联通核时包"}, // 容器云隐藏卡类型
				CardCounts: nil,               // 容器云隐藏卡数
				PriceUnit:  "元/核时",
			},
		},
		Products: []Product{},
	}
}

func GetResourceTypeEmpty() []ResourceType {
	return []ResourceType{
		{
			Label:      "云服务器",
			Category:   "cloud_server",
			CardTypes:  []string{},
			CardCounts: []int64{},
			PriceUnit:  "元/卡/时",
		},
		{
			Label:      "裸金属",
			Category:   "bare_metal",
			CardTypes:  []string{},
			CardCounts: nil,
			PriceUnit:  "元/台/月",
		},
		{
			Label:      "容器云",
			Category:   "container_cloud",
			CardTypes:  []string{},
			CardCounts: nil,
			PriceUnit:  "元/核时",
		},
	}
}
