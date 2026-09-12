package pay_computility

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	HAOSUANPartnerName     = "computility.haoSuan"
	PARATERPartnerName     = "computility.paratera"
	CHINAMOBILEPartnerName = "computility.ecloud"
)

var (
	ProvidersList = []string{HAOSUANPartnerName, PARATERPartnerName, CHINAMOBILEPartnerName}
)

type ComputilityManager struct {
	ctx        *context.APIContext
	filterData FilterData
	errorList  []error
	wg         sync.WaitGroup
}

func InitComputilityManager(ctx *context.APIContext) *ComputilityManager {
	return &ComputilityManager{
		ctx: ctx,
		filterData: FilterData{
			ResourceTypes: make([]ResourceType, 0),
			Products:      make([]Product, 0),
		},
		errorList: []error{},
		wg:        sync.WaitGroup{},
	}
}

// ComputilityFilter 算力过滤，后续叠加算力
func (c *ComputilityManager) ComputilityFilter() error {
	c.GetJson()

	if len(c.errorList) > 0 {
		// 错误可以做特殊处理或者日志打印

		for _, err := range c.errorList {
			log.Error("Computility error: %+v", err)
		}
		// 随意返回一条
		return c.errorList[0]
	}

	return nil
}

// GetJson 获取配置的json
func (c *ComputilityManager) GetJson() *ComputilityManager {

	defer func() {
		if r := recover(); r != nil {
			log.Error("ComputilityManager panic: %+v", r)
		}
	}()

	var (
		haosuanJsonData     = FilterData{}
		perateraJsonData    = FilterData{}
		chinaMobileJsonData = FilterData{}
		manager             = InitComputilityConfigManager()
	)

	// 增加对应的数据
	// step 1 HAOSUAN 配置
	haosuanConfigList, _ := manager.hsConfig.GetConfig(GitComputilityConfigPath, false)

	// 不用断言会出现panic
	if str, ok := haosuanConfigList.(string); ok && haosuanConfigList != nil {
		if err := json.Unmarshal([]byte(str), &haosuanJsonData); err != nil {
			// 错误放到列表里面去
			c.errorList = append(c.errorList, err)
		}
	}

	//打上产品提供方标签
	c.filterData = c.FillProductType(HAOSUANPartnerName, haosuanJsonData)

	// 如果昊算出现了配置问题，那么也要不影响其他第三方场景的使用
	if len(c.filterData.ResourceTypes) == 0 {
		c.filterData.ResourceTypes = GetResourceTypeEmpty()
	}

	// step 1.1 并行科技配置
	perateraConfigList, _ := manager.peraConfig.GetConfig(GitParateraComputilityConfigPath, false)

	// 不用断言会出现panic
	if str, ok := perateraConfigList.(string); ok && perateraConfigList != nil {
		if err := json.Unmarshal([]byte(str), &perateraJsonData); err != nil {
			// 错误放到列表里面去
			c.errorList = append(c.errorList, err)
		}
	}

	// 打上产品提供方标签
	perateraJsonData = c.FillProductType(PARATERPartnerName, perateraJsonData)

	// step 1.2 增加中国移动
	mobildConfigList, _ := manager.mobileConfig.GetConfig(GitChinaMobileComputilityConfigPath, false)

	// 不用断言会出现panic
	if str, ok := mobildConfigList.(string); ok && mobildConfigList != nil {
		if err := json.Unmarshal([]byte(str), &chinaMobileJsonData); err != nil {
			// 错误放到列表里面去
			c.errorList = append(c.errorList, err)
		}
	}

	//  打上产品提供方标签
	chinaMobileJsonData = c.FillProductType(CHINAMOBILEPartnerName, chinaMobileJsonData)

	// step 2 组装数据
	c.ComposeFilterData(perateraJsonData, chinaMobileJsonData)

	// step 3 补充筛选列表提供方
	c.FillProviders()

	// step 4 排序返回
	c.Sort()

	// step 5 中英文转换
	c.Translate()
	return c
}

// Translate 中英文转换
func (c *ComputilityManager) Translate() {
	for i := range c.filterData.ResourceTypes {
		// 筛选类型
		c.filterData.ResourceTypes[i].Label = c.ctx.Locale.Tr(fmt.Sprintf("computility.%s", c.filterData.ResourceTypes[i].Category))
	}
	return
}

// FillProviders 补充提供方
func (c *ComputilityManager) FillProviders() {
	if len(c.filterData.ResourceTypes) == 0 {
		return
	}
	localizedProviders := make([]string, len(ProvidersList))
	for i := range ProvidersList {
		localizedProviders[i] = c.ctx.Locale.Tr(ProvidersList[i])
	}

	for i := range c.filterData.ResourceTypes {
		c.filterData.ResourceTypes[i].Providers = localizedProviders
	}

	return
}

// FillProductType  补充产品类型
func (c *ComputilityManager) FillProductType(partner string, filterData FilterData) FilterData {
	if len(filterData.Products) == 0 {
		return filterData
	}
	for i := range filterData.Products {
		filterData.Products[i].Provider = c.ctx.Locale.Tr(partner)
	}

	return filterData
}

// Sort 根据CardType按字母A-Z排序（不区分大小写）
func (c *ComputilityManager) Sort() *ComputilityManager {
	if c == nil || len(c.filterData.Products) == 0 {
		return c
	}

	// 在排序卡类型,否则id降序排序
	sort.SliceStable(c.filterData.Products, func(i, j int) bool {
		cardTypeI := strings.ToUpper(c.filterData.Products[i].CardType)
		cardTypeJ := strings.ToUpper(c.filterData.Products[j].CardType)

		if cardTypeI != cardTypeJ {
			return cardTypeI < cardTypeJ
		}

		// CardType相同时，按ID降序排序
		return c.filterData.Products[i].ID > c.filterData.Products[j].ID
	})

	return c
}

// ComposeFilterData 组合返回资源
// 主要处理【c.filterData】里面的数据
func (c *ComputilityManager) ComposeFilterData(newFilterDataList ...FilterData) *ComputilityManager {

	if c.filterData.ResourceTypes == nil {
		c.filterData = GetFilterData()
	}

	var (
		categoryMap   = make(map[string]ResourceType)
		cardTypesMap  = make(map[string]bool)
		cardCountMap  = make(map[int64]bool)
		resourceTypes = make([]ResourceType, 0, len(c.filterData.ResourceTypes))
	)

	for _, newFilterData := range newFilterDataList {
		if newFilterData.Products != nil {
			c.filterData.Products = append(c.filterData.Products, newFilterData.Products...)
		}

		// 处理之前的数据,主要是筛选类型
		for _, index := range c.filterData.ResourceTypes {
			if index.Category == "" {
				continue
			}

			categoryMap[index.Category] = index
			for _, cardType := range index.CardTypes {
				cardTypesMap[cardType] = true
			}
			for _, cardCount := range index.CardCounts {
				cardCountMap[cardCount] = true
			}
		}

		// 双列表合并
		for _, index := range newFilterData.Products {
			if categoryMap[index.Category].Category == "" {
				continue
			}

			isNotExistCardType := !cardTypesMap[index.CardType]
			isNotExistCardCount := !cardCountMap[index.CardCount]
			category := categoryMap[index.Category]

			// 如果不存在对应数据，就添加
			if isNotExistCardType {
				category.CardTypes = append(category.CardTypes, index.CardType)
			}

			if isNotExistCardCount && index.CardCount > 0 {
				category.CardCounts = append(category.CardCounts, index.CardCount)
			}

			categoryMap[index.Category] = category
		}
	}

	// 同时补充筛选条件
	for _, index := range categoryMap {
		resourceTypes = append(resourceTypes, index)
	}

	if len(resourceTypes) > 0 {
		c.filterData.ResourceTypes = resourceTypes
	}

	return c
}

// GetApi  获取第三方api
func (c *ComputilityManager) GetApi() *ComputilityManager {
	// TODO 若需要对接bx算力，需要在这里合并返回
	return c
}
