package routers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"github.com/olivere/elastic/v7"
)

type SearchRes struct {
	Total        int64
	Result       []map[string]interface{}
	PrivateTotal int64
}

var client *elastic.Client

func InitESClient() {
	ESSearchUrl := setting.ESSearchURL
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(ESSearchUrl))
	if err != nil {
		log.Info("es init error.")
		//panic(err)
	}
}

func EmptySearch(ctx *context.Context) {
	log.Info("search template.")
	ctx.Data["Keyword"] = ""
	ctx.HTML(200, "explore/search_new")
}

func Search(ctx *context.Context) {
	log.Info("search template.")
	keyword := strings.Trim(ctx.Query("q"), " ")
	ctx.Data["Keyword"] = keyword
	ctx.Data["SortType"] = "newest"
	ctx.HTML(200, "explore/search_new")
}

func SearchApi(ctx *context.Context) {
	TableName := ctx.Query("TableName")
	Key := ctx.Query("Key")
	Page := ctx.QueryInt("Page")
	PageSize := ctx.QueryInt("PageSize")
	OnlyReturnNum := ctx.QueryBool("OnlyReturnNum")
	OnlySearchLabel := ctx.QueryBool("OnlySearchLabel")

	if Page <= 0 {
		Page = 1
	}
	if PageSize <= 0 || PageSize > 200 {
		PageSize = setting.UI.IssuePagingNum
	}
	if Key != "" && !OnlyReturnNum {
		go models.SaveSearchKeywordToDb(Key)
	}
	if TableName == "repository" {
		if OnlySearchLabel {
			searchRepoByLabel(ctx, Key, Page, PageSize)
		} else {
			searchRepo(ctx, "repository-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, OnlyReturnNum)
		}
		return
	} else if TableName == "issue" {
		searchIssueOrPr(ctx, "issue-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, OnlyReturnNum, "f")
		return
	} else if TableName == "user" {
		searchUserOrOrg(ctx, "user-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, true, OnlyReturnNum)
		return
	} else if TableName == "org" {
		searchUserOrOrg(ctx, "user-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, false, OnlyReturnNum)
		return
	} else if TableName == "dataset" {
		searchDataSet(ctx, "dataset-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, OnlyReturnNum)
		return
	} else if TableName == "pr" {
		searchIssueOrPr(ctx, "issue-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, OnlyReturnNum, "t")
		//searchPR(ctx, "issue-es-index", Key, Page, PageSize, OnlyReturnNum)
		return
	} else if TableName == "model" {
		searchModel(ctx, "model-es-index"+setting.INDEXPOSTFIX, Key, Page, PageSize, OnlyReturnNum)
		return
	}
}

func searchRepoByLabel(ctx *context.Context, Key string, Page int, PageSize int) {
	/*
			   项目， ES名称： repository-es-index
			   搜索:
			   name character varying(255) ,  项目名称
			   description text,  项目描述
			   topics json,  标签
			   排序:
			   updated_unix
			   num_watches,
		       num_stars,
		       num_forks,
	*/
	SortBy := ctx.Query("SortBy")
	PrivateTotal := ctx.QueryInt("PrivateTotal")
	WebTotal := ctx.QueryInt("WebTotal")
	ascending := ctx.QueryBool("Ascending")
	language := ctx.Query("language")
	if language == "" {
		language = "zh-CN"
	}
	from := (Page - 1) * PageSize
	resultObj := &SearchRes{}
	log.Info("WebTotal=" + fmt.Sprint(WebTotal))
	log.Info("PrivateTotal=" + fmt.Sprint(PrivateTotal))
	resultObj.Result = make([]map[string]interface{}, 0)
	if from == 0 {
		WebTotal = 0
	}
	if ctx.User != nil && (from < PrivateTotal || from == 0) {
		orderBy := models.SearchOrderByRecentUpdated
		switch SortBy {
		case "updated_unix.keyword":
			orderBy = models.SearchOrderByRecentUpdated
		case "num_stars":
			orderBy = models.SearchOrderByStarsReverse
		case "num_forks":
			orderBy = models.SearchOrderByForksReverse
		case "num_watches":
			orderBy = models.SearchOrderByWatches
		}
		log.Info("actor is null?:" + fmt.Sprint(ctx.User == nil))
		repos, count, err := models.SearchRepository(&models.SearchRepoOptions{
			ListOptions: models.ListOptions{
				Page:     Page,
				PageSize: PageSize,
			},
			Actor:              ctx.User,
			OrderBy:            orderBy,
			Private:            true,
			OnlyPrivate:        true,
			TopicOnly:          true,
			TopicName:          Key,
			IncludeDescription: setting.UI.SearchRepoDescription,
		})
		if err != nil {
			ctx.JSON(200, "")
			return
		}
		resultObj.PrivateTotal = count
		if repos.Len() > 0 {
			log.Info("Query private repo number is:" + fmt.Sprint(repos.Len()))
			makePrivateRepo(repos, resultObj, Key, language)
		} else {
			log.Info("not found private repo,keyword=" + Key)
		}
		if repos.Len() >= PageSize {
			if WebTotal > 0 {
				resultObj.Total = int64(WebTotal)
				ctx.JSON(200, resultObj)
				return
			}
		}
	} else {
		if ctx.User == nil {
			resultObj.PrivateTotal = 0
		} else {
			resultObj.PrivateTotal = int64(PrivateTotal)
		}
	}

	from = from - PrivateTotal
	if from < 0 {
		from = 0
	}
	Size := PageSize - len(resultObj.Result)

	log.Info("query searchRepoByLabel start")
	if Key != "" {
		boolQ := elastic.NewBoolQuery()
		topicsQuery := elastic.NewMatchQuery("topics", Key)
		boolQ.Should(topicsQuery)

		res, err := client.Search("repository-es-index").Query(boolQ).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Highlight(queryHighlight("topics")).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeRepoResult(res, "", false, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
			ctx.JSON(200, "")
		}
	} else {
		ctx.JSON(200, "")
	}
}

func getSort(SortBy string, ascending bool, secondSortBy string, secondAscending bool) []elastic.Sorter {
	sort := make([]elastic.Sorter, 0)
	if SortBy == "default" || SortBy == "" {
		sort = append(sort, elastic.NewScoreSort())
		if secondSortBy != "" {
			log.Info("SortBy=" + SortBy + " secondSortBy=" + secondSortBy)
			sort = append(sort, elastic.NewFieldSort(secondSortBy).Order(secondAscending))
		}
	} else {
		log.Info("SortBy=" + SortBy)
		sort = append(sort, elastic.NewFieldSort(SortBy).Order(ascending))
	}
	log.Info("sort size=" + fmt.Sprint(len(sort)))
	return sort
}

func searchRepo(ctx *context.Context, TableName string, Key string, Page int, PageSize int, OnlyReturnNum bool) {
	/*
			   项目， ES名称： repository-es-index
			   搜索:
			   name character varying(255) ,  项目名称
			   description text,  项目描述
			   topics json,  标签
			   排序:
			   updated_unix
			   num_watches,
		       num_stars,
		       num_forks,
	*/

	SortBy := ctx.Query("SortBy")
	PrivateTotal := ctx.QueryInt("PrivateTotal")
	WebTotal := ctx.QueryInt("WebTotal")
	ascending := ctx.QueryBool("Ascending")
	from := (Page - 1) * PageSize
	resultObj := &SearchRes{}
	log.Info("WebTotal=" + fmt.Sprint(WebTotal))
	log.Info("PrivateTotal=" + fmt.Sprint(PrivateTotal))
	resultObj.Result = make([]map[string]interface{}, 0)
	if from == 0 {
		WebTotal = 0
	}
	language := ctx.Query("language")
	if language == "" {
		language = "zh-CN"
	}
	if ctx.User != nil && (from < PrivateTotal || from == 0) {
		orderBy := models.SearchOrderByRecentUpdated
		switch SortBy {
		case "updated_unix.keyword":
			orderBy = models.SearchOrderByRecentUpdated
		case "num_stars":
			orderBy = models.SearchOrderByStarsReverse
		case "num_forks":
			orderBy = models.SearchOrderByForksReverse
		case "num_watches":
			orderBy = models.SearchOrderByWatches
		}
		log.Info("actor is null?:" + fmt.Sprint(ctx.User == nil))
		repos, count, err := models.SearchRepository(&models.SearchRepoOptions{
			ListOptions: models.ListOptions{
				Page:     Page,
				PageSize: PageSize,
			},
			Actor:              ctx.User,
			OrderBy:            orderBy,
			Private:            true,
			OnlyPrivate:        true,
			Keyword:            Key,
			IncludeDescription: setting.UI.SearchRepoDescription,
			OnlySearchPrivate:  true,
		})
		if err != nil {
			ctx.JSON(200, "")
			return
		}
		resultObj.PrivateTotal = count
		if repos.Len() > 0 {
			log.Info("Query private repo number is:" + fmt.Sprint(repos.Len()))
			makePrivateRepo(repos, resultObj, Key, language)
		} else {
			log.Info("not found private repo,keyword=" + Key)
		}
		if repos.Len() >= PageSize {
			if WebTotal > 0 {
				resultObj.Total = int64(WebTotal)
				ctx.JSON(200, resultObj)
				return
			}
		}
	} else {
		if ctx.User == nil {
			resultObj.PrivateTotal = 0
		} else {
			resultObj.PrivateTotal = int64(PrivateTotal)
		}
	}

	from = from - PrivateTotal
	if from < 0 {
		from = 0
	}
	Size := PageSize - len(resultObj.Result)

	log.Info("query searchRepo start")
	if Key != "" {
		boolQ := elastic.NewBoolQuery()
		sKey := Tokenize(Key)
		nameQuery := elastic.NewMatchQuery("alias", sKey).Boost(100).QueryName("f_first")
		topicsQuery := elastic.NewMatchQuery("topics", sKey).Boost(1).QueryName("f_second")
		lower_aliasQuery := elastic.NewMatchQuery("lower_alias", sKey).Boost(1).QueryName("f_three")
		lower_aliasQuery1 := elastic.NewMatchQuery("lower_alias", strings.ToLower(Key)).Boost(1024).QueryName("f_four")
		boolQ.Should(nameQuery, topicsQuery, lower_aliasQuery, lower_aliasQuery1)

		res, err := client.Search(TableName).Query(boolQ).SortBy(getSort(SortBy, ascending, "num_stars", false)...).From(from).Size(Size).Highlight(queryHighlight("alias", "topics")).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeRepoResult(res, Key, OnlyReturnNum, language)
			setForkRepoOrder(esresult, SortBy)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			isNeedSort := false
			if len(resultObj.Result) > 0 {
				isNeedSort = true
			}
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			if isNeedSort {
				sortRepo(resultObj.Result, SortBy, ascending)
			}
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
			ctx.JSON(200, "")
		}
	} else {
		log.Info("query all content.")
		//搜索的属性要指定{"timestamp":{"unmapped_type":"date"}}
		res, err := client.Search(TableName).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeRepoResult(res, "", OnlyReturnNum, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
			ctx.JSON(200, "")
		}
	}
}

func setForkRepoOrder(esresult *SearchRes, SortBy string) {
	if SortBy == "default" || SortBy == "" {
		forkidMap := make(map[string]int, 0)
		for index, re := range esresult.Result {
			if re["fork_id"] != nil {
				fork_id := re["fork_id"].(string)
				if _, ok := forkidMap[fork_id]; !ok {
					forkidMap[fork_id] = index
				}
			}
		}
		for key, value := range forkidMap {
			for index, re := range esresult.Result {
				if re["id"].(string) == key {
					if value < index { //swap
						tmp := esresult.Result[index]
						esresult.Result[index] = esresult.Result[value]
						esresult.Result[value] = tmp
						break
					}
				}
			}
		}
	}
}

func sortRepo(Result []map[string]interface{}, SortBy string, ascending bool) {
	orderBy := ""
	switch SortBy {
	case "updated_unix.keyword":
		orderBy = "updated_unix"
	case "num_stars":
		orderBy = "num_stars"
	case "num_forks":
		orderBy = "num_forks"
	case "num_watches":
		orderBy = "num_watches"
	}
	sort.Slice(Result, func(i, j int) bool {
		return getInt(Result[i][orderBy], orderBy) > getInt(Result[j][orderBy], orderBy)
	})
}

func getInt(tmp interface{}, orderBy string) int64 {
	timeInt, err := strconv.ParseInt(fmt.Sprint(tmp), 10, 64)
	if err == nil {
		return timeInt
	} else {
		log.Info("convert " + orderBy + " error type=" + fmt.Sprint(tmp))
	}
	return -1
}

func makePrivateRepo(repos models.RepositoryList, res *SearchRes, keyword string, language string) {

	for _, repo := range repos {
		record := make(map[string]interface{})
		record["id"] = repo.ID
		record["name"] = makeHighLight(keyword, repo.Name)
		record["real_name"] = repo.Name
		record["owner_name"] = repo.OwnerName
		record["description"] = truncLongText(makeHighLight(keyword, repo.Description), true)

		hightTopics := make([]string, 0)
		if len(repo.Topics) > 0 {
			for _, t := range repo.Topics {
				hightTopics = append(hightTopics, makeHighLight(keyword, t))
			}
		}
		record["hightTopics"] = hightTopics

		record["num_watches"] = repo.NumWatches
		record["num_stars"] = repo.NumStars
		record["num_forks"] = repo.NumForks
		record["alias"] = truncLongText(makeHighLight(keyword, repo.Alias), true)
		record["lower_alias"] = repo.LowerAlias
		record["topics"] = repo.Topics
		record["avatar"] = repo.RelAvatarLink()
		if len(repo.RelAvatarLink()) == 0 {
			record["avatar"] = setting.RepositoryAvatarFallbackImage
		}
		record["updated_unix"] = repo.UpdatedUnix
		record["updated_html"] = timeutil.TimeSinceUnix(repo.UpdatedUnix, language)
		lang, err := repo.GetTopLanguageStats(1)
		if err == nil && len(lang) > 0 {
			record["lang"] = lang[0].Language
		} else {
			record["lang"] = ""
		}
		record["is_private"] = true
		res.Result = append(res.Result, record)
	}
}

func makeHighLight(keyword string, dest string) string {

	dest = replaceIngoreUpperOrLower(dest, strings.ToLower(dest), strings.ToLower(keyword))

	return dest
}

func replaceIngoreUpperOrLower(dest string, destLower string, keywordLower string) string {
	re := ""
	last := 0
	lenDestLower := len(destLower)
	lenkeywordLower := len(keywordLower)
	for i := 0; i < lenDestLower; i++ {
		if destLower[i] == keywordLower[0] {
			isFind := true
			for j := 1; j < lenkeywordLower; j++ {
				if (i+j) < lenDestLower && keywordLower[j] != destLower[i+j] {
					isFind = false
					break
				}
			}
			if isFind && (i+lenkeywordLower) <= lenDestLower {
				re += dest[last:i] + "\u003cfont color='red'\u003e" + dest[i:(i+lenkeywordLower)] + "\u003c/font\u003e"
				i = i + lenkeywordLower
				last = i
			}
		}
	}
	if last < lenDestLower {
		re += dest[last:lenDestLower]
	}
	return re
}

func makeRepoResult(sRes *elastic.SearchResult, Key string, OnlyReturnNum bool, language string) *SearchRes {
	total := sRes.Hits.TotalHits.Value
	result := make([]map[string]interface{}, 0)
	if !OnlyReturnNum {
		for i, hit := range sRes.Hits.Hits {
			log.Info("this is repo query " + fmt.Sprint(i) + " result.")
			recordSource := make(map[string]interface{})
			source, err := hit.Source.MarshalJSON()

			if err == nil {
				err = json.Unmarshal(source, &recordSource)
				if err == nil {
					record := make(map[string]interface{})
					record["id"] = hit.Id
					record["alias"] = getLabelValue("alias", recordSource, hit.Highlight)
					record["real_name"] = recordSource["name"]
					record["owner_name"] = recordSource["owner_name"]
					if recordSource["description"] != nil {
						desc := getLabelValue("description", recordSource, hit.Highlight)
						record["description"] = dealLongText(desc, Key, hit.MatchedQueries)
					} else {
						record["description"] = ""
					}

					record["hightTopics"] = jsonStrToArray(getLabelValue("topics", recordSource, hit.Highlight))
					record["num_watches"] = recordSource["num_watches"]
					record["num_stars"] = recordSource["num_stars"]
					record["num_forks"] = recordSource["num_forks"]
					record["lower_alias"] = recordSource["lower_alias"]
					record["fork_id"] = recordSource["fork_id"]
					if recordSource["topics"] != nil {
						topicsStr := recordSource["topics"].(string)
						log.Info("topicsStr=" + topicsStr)
						if topicsStr != "null" {
							record["topics"] = jsonStrToArray(topicsStr)
						}
					}
					if recordSource["avatar"] != nil {
						avatarstr := recordSource["avatar"].(string)
						if len(avatarstr) == 0 {
							// record["avatar"] = setting.RepositoryAvatarFallbackImage
						} else {
							record["avatar"] = setting.AppSubURL + "/repo-avatars/" + avatarstr
						}
					}
					record["updated_unix"] = recordSource["updated_unix"]
					setUpdateHtml(record, recordSource["updated_unix"].(string), language)

					record["lang"] = recordSource["lang"]
					record["is_private"] = false
					result = append(result, record)
				} else {
					log.Info("deal repo source error," + err.Error())
				}
			} else {
				log.Info("deal repo source error," + err.Error())
			}
		}
	}
	returnObj := &SearchRes{
		Total:  total,
		Result: result,
	}

	return returnObj
}

func setUpdateHtml(record map[string]interface{}, updated_unix string, language string) {
	timeInt, err := strconv.ParseInt(updated_unix, 10, 64)
	if err == nil {
		record["updated_html"] = timeutil.TimeSinceUnix(timeutil.TimeStamp(timeInt), language)
	}
}

func jsonStrToArray(str string) []string {
	b := []byte(str)
	strs := make([]string, 0)
	err := json.Unmarshal(b, &strs)
	if err != nil {
		log.Info("convert str arrar error, str=" + str)
	}
	return strs
}

func dealLongText(text string, Key string, MatchedQueries []string) string {
	var isNeedToDealText bool
	isNeedToDealText = false
	if len(MatchedQueries) > 0 && Key != "" {
		if MatchedQueries[0] == "f_second" || MatchedQueries[0] == "f_third" {
			isNeedToDealText = true
		}
	}
	return truncLongText(text, isNeedToDealText)
}

func truncLongText(text string, isNeedToDealText bool) string {
	startStr := "color="
	textRune := []rune(text)
	stringlen := len(textRune)
	if isNeedToDealText && stringlen > 200 {
		index := findFont(textRune, []rune(startStr))
		if index > 0 {
			start := index - 50
			if start < 0 {
				start = 0
			}
			end := index + 150
			if end >= stringlen {
				end = stringlen
			}
			return trimFontHtml(textRune[start:end]) + "..."
		} else {
			return trimFontHtml(textRune[0:200]) + "..."
		}
	} else {
		if stringlen > 200 {
			return trimFontHtml(textRune[0:200]) + "..."
		} else {
			return text
		}
	}
}

func trimFontHtml(text []rune) string {
	startRune := rune('<')
	endRune := rune('>')
	count := 0
	i := 0
	for ; i < len(text); i++ {
		if text[i] == startRune { //start <
			re := false
			j := i + 1
			for ; j < len(text); j++ {
				if text[j] == endRune {
					re = true
					break
				}
			}
			if re { //found >
				i = j
				count++
			} else {
				if count%2 == 1 {
					return string(text[0:i]) + "</font>"
				} else {
					return string(text[0:i])
				}
			}
		}
	}
	if count%2 == 1 {
		return string(text[0:i]) + "</font>"
	} else {
		return string(text[0:i])
	}
}

func trimHrefHtml(result string) string {
	result = strings.Replace(result, "</a>", "", -1)
	result = strings.Replace(result, "\n", "", -1)
	var index int
	for {
		index = findSubstr(result, 0, "<a")
		if index != -1 {
			sIndex := findSubstr(result, index+2, ">")
			if sIndex != -1 {
				result = result[0:index] + result[sIndex+1:]
			} else {
				result = result[0:index] + result[index+2:]
			}
		} else {
			break
		}
	}
	return result
}

func findFont(text []rune, childText []rune) int {
	for i := 0; i < len(text); i++ {
		if text[i] == childText[0] {
			re := true
			for j, k := range childText {
				if k != text[i+j] {
					re = false
					break
				}
			}
			if re {
				return i
			}
		}
	}
	return -1
}

func findSubstr(text string, startindex int, childText string) int {
	for i := startindex; i < len(text); i++ {
		if text[i] == childText[0] {
			re := true
			for k := range childText {
				if childText[k] != text[i+k] {
					re = false
					break
				}
			}
			if re {
				return i
			}
		}
	}
	return -1
}

func searchUserOrOrg(ctx *context.Context, TableName string, Key string, Page int, PageSize int, IsQueryUser bool, OnlyReturnNum bool) {
	/*
	   用户或者组织 ES名称： user-es-index
	   搜索:
	   name  ,  名称
	   full_name   全名
	   description  描述或者简介
	   排序:
	   created_unix
	   名称字母序
	*/
	SortBy := ctx.Query("SortBy")
	ascending := ctx.QueryBool("Ascending")
	boolQ := elastic.NewBoolQuery()

	typeValue := 1
	if IsQueryUser {
		typeValue = 0
	}
	UserOrOrgQuery := elastic.NewTermQuery("type", typeValue)
	if Key != "" {
		boolKeyQ := elastic.NewBoolQuery()
		log.Info("user or org Key=" + Key)
		nameQuery := elastic.NewMatchQuery("name", Key).Boost(2).QueryName("f_first")
		full_nameQuery := elastic.NewMatchQuery("full_name", Key).Boost(1.5).QueryName("f_second")
		descriptionQuery := elastic.NewMatchQuery("description", Key).Boost(1).QueryName("f_third")
		boolKeyQ.Should(nameQuery, full_nameQuery, descriptionQuery)
		boolQ.Must(UserOrOrgQuery, boolKeyQ)
	} else {
		boolQ.Must(UserOrOrgQuery)
	}

	res, err := client.Search(TableName).Query(boolQ).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From((Page - 1) * PageSize).Size(PageSize).Highlight(queryHighlight("name", "full_name", "description")).Do(ctx.Req.Context())
	if err == nil {
		//searchJson, _ := json.Marshal(res)
		//log.Info("searchJson=" + string(searchJson))
		result := makeUserOrOrgResult(res, Key, ctx, OnlyReturnNum)
		ctx.JSON(200, result)
	} else {
		log.Info("query es error," + err.Error())
		ctx.JSON(200, "")
	}
}

func getLabelValue(key string, recordSource map[string]interface{}, searchHighliht elastic.SearchHitHighlight) string {
	if value, ok := searchHighliht[key]; !ok {
		if recordSource[key] != nil {
			return recordSource[key].(string)
		} else {
			return ""
		}
	} else {
		return value[0]
	}
}

func makeUserOrOrgResult(sRes *elastic.SearchResult, Key string, ctx *context.Context, OnlyReturnNum bool) *SearchRes {
	total := sRes.Hits.TotalHits.Value
	result := make([]map[string]interface{}, 0)
	if !OnlyReturnNum {
		for i, hit := range sRes.Hits.Hits {
			log.Info("this is user query " + fmt.Sprint(i) + " result.")
			recordSource := make(map[string]interface{})
			source, err := hit.Source.MarshalJSON()

			if err == nil {
				err = json.Unmarshal(source, &recordSource)
				if err == nil {
					record := make(map[string]interface{})
					record["id"] = hit.Id
					record["name"] = getLabelValue("name", recordSource, hit.Highlight)
					record["real_name"] = recordSource["name"]
					record["full_name"] = getLabelValue("full_name", recordSource, hit.Highlight)
					if recordSource["description"] != nil {
						desc := getLabelValue("description", recordSource, hit.Highlight)
						record["description"] = dealLongText(desc, Key, hit.MatchedQueries)
					} else {
						record["description"] = ""
					}
					if ctx.User != nil {
						record["email"] = recordSource["email"]
					} else {
						record["email"] = ""
					}

					record["location"] = recordSource["location"]
					record["website"] = recordSource["website"]
					record["num_repos"] = recordSource["num_repos"]
					record["num_teams"] = recordSource["num_teams"]
					record["num_members"] = recordSource["num_members"]

					record["avatar"] = strings.TrimRight(setting.AppSubURL, "/") + "/user/avatar/" + recordSource["name"].(string) + "/" + strconv.Itoa(-1)
					record["updated_unix"] = recordSource["updated_unix"]
					record["created_unix"] = recordSource["created_unix"]
					record["add_time"] = getAddTime(recordSource["created_unix"].(string))
					result = append(result, record)
				} else {
					log.Info("deal user source error," + err.Error())
				}
			} else {
				log.Info("deal user source error," + err.Error())
			}
		}
	}
	returnObj := &SearchRes{
		Total:  total,
		Result: result,
	}
	return returnObj
}

func getAddTime(time string) string {
	timeInt, err := strconv.ParseInt(time, 10, 64)
	if err == nil {
		t := timeutil.TimeStamp(timeInt)
		return t.FormatShort()
	}
	return ""
}

func searchDataSet(ctx *context.Context, TableName string, Key string, Page int, PageSize int, OnlyReturnNum bool) {
	/*
	   数据集，ES名称：dataset-es-index
	   搜索:
	   title  ,  名称
	   description   描述
	   category  标签
	   file_name  数据集文件名称
	   排序:
	   download_times

	*/
	log.Info("query searchdataset start")
	SortBy := ctx.Query("SortBy")
	ascending := ctx.QueryBool("Ascending")
	PrivateTotal := ctx.QueryInt("PrivateTotal")
	WebTotal := ctx.QueryInt("WebTotal")
	language := ctx.Query("language")
	if language == "" {
		language = "zh-CN"
	}
	from := (Page - 1) * PageSize
	if from == 0 {
		WebTotal = 0
	}
	resultObj := &SearchRes{}
	log.Info("WebTotal=" + fmt.Sprint(WebTotal))
	log.Info("PrivateTotal=" + fmt.Sprint(PrivateTotal))
	resultObj.Result = make([]map[string]interface{}, 0)

	if ctx.User != nil && (from < PrivateTotal || from == 0) {

		log.Info("actor is null?:" + fmt.Sprint(ctx.User == nil))

		datasets, count, err := models.SearchDatasetBySQL(Page, PageSize, Key, ctx.User.ID, ctx.User.IsAdmin)
		if err != nil {
			ctx.JSON(200, "")
			return
		}
		resultObj.PrivateTotal = count
		datasetSize := len(datasets)
		if datasetSize > 0 {
			log.Info("Query private dataset number is:" + fmt.Sprint(datasetSize) + " count=" + fmt.Sprint(count))
			makePrivateDataSet(datasets, resultObj, Key, language)
		} else {
			log.Info("not found private dataset, keyword=" + Key)
		}
		if datasetSize >= PageSize {
			if WebTotal > 0 { //next page， not first query.
				resultObj.Total = int64(WebTotal)
				ctx.JSON(200, resultObj)
				return
			}
		}
	} else {
		resultObj.PrivateTotal = int64(PrivateTotal)
	}

	from = from - PrivateTotal
	if from < 0 {
		from = 0
	}
	Size := PageSize - len(resultObj.Result)

	boolQ := elastic.NewBoolQuery()
	if Key != "" {
		nameQuery := elastic.NewMatchQuery("title", Key).Boost(2).QueryName("f_first")
		categoryQuery := elastic.NewMatchQuery("category", Key).Boost(1).QueryName("f_second")
		boolQ.Should(nameQuery, categoryQuery)
		res, err := client.Search(TableName).Query(boolQ).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Highlight(queryHighlight("title", "category")).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeDatasetResult(res, Key, OnlyReturnNum, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			log.Info("query dataset es count=" + fmt.Sprint(esresult.Total) + " total=" + fmt.Sprint(resultObj.Total))
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
		}
	} else {
		log.Info("query all datasets.")
		//搜索的属性要指定{"timestamp":{"unmapped_type":"date"}}
		res, err := client.Search(TableName).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeDatasetResult(res, "", OnlyReturnNum, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			log.Info("query dataset es count=" + fmt.Sprint(esresult.Total) + " total=" + fmt.Sprint(resultObj.Total))
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
			ctx.JSON(200, "")
		}
	}

}

func makePrivateDataSet(datasets []*models.DatasetRegistry, res *SearchRes, Key string, language string) {
	for _, dataset := range datasets {
		record := make(map[string]interface{})

		record["id"] = dataset.ID
		userId := dataset.OwnerID

		user, errUser := models.GetUserByID(userId)
		if errUser == nil {
			record["owerName"] = user.GetDisplayName()
			record["avatar"] = user.RelAvatarLink()
		}
		record["name"] = dataset.Name
		record["title"] = makeHighLight(Key, dataset.Alias)
		record["uuid"] = dataset.ID
		record["description"] = ""
		record["category"] = dataset.Tags
		record["task"] = dataset.Tasks
		record["download_times"] = dataset.DownloadCount
		record["created_unix"] = dataset.CreatedUnix
		record["updated_unix"] = dataset.UpdatedUnix
		record["updated_html"] = timeutil.TimeSinceUnix(dataset.UpdatedUnix, language)

		res.Result = append(res.Result, record)
	}
}

func makeDatasetResult(sRes *elastic.SearchResult, Key string, OnlyReturnNum bool, language string) *SearchRes {
	total := sRes.Hits.TotalHits.Value
	result := make([]map[string]interface{}, 0)
	if !OnlyReturnNum {
		for i, hit := range sRes.Hits.Hits {
			log.Info("this is dataset query " + fmt.Sprint(i) + " result.")
			recordSource := make(map[string]interface{})
			source, err := hit.Source.MarshalJSON()

			if err == nil {
				err = json.Unmarshal(source, &recordSource)
				if err == nil {
					record := make(map[string]interface{})
					record["id"] = hit.Id
					userIdStr := recordSource["user_id"].(string)
					userId, cerr := strconv.ParseInt(userIdStr, 10, 64)
					if cerr == nil {
						user, errUser := models.GetUserByID(userId)
						if errUser == nil {
							record["owerName"] = user.GetDisplayName()
							record["avatar"] = user.RelAvatarLink()
						}
					}
					record["name"] = recordSource["name"]
					record["title"] = getLabelValue("title", recordSource, hit.Highlight)
					record["category"] = getLabelValue("category", recordSource, hit.Highlight)
					record["description"] = ""
					record["uuid"] = recordSource["uuid"]
					record["file_name"] = ""
					record["task"] = recordSource["task"]
					record["download_times"] = recordSource["download_times"]
					record["created_unix"] = recordSource["created_unix"]
					setUpdateHtml(record, recordSource["updated_unix"].(string), language)
					result = append(result, record)
				} else {
					log.Info("deal dataset source error," + err.Error())
				}
			} else {
				log.Info("deal dataset source error," + err.Error())
			}
		}
	}
	returnObj := &SearchRes{
		Total:  total,
		Result: result,
	}

	return returnObj
}

func getDatasetFileName(fileName string) string {
	slices := strings.Split(fileName, "-#,#-")
	fileName = strings.Join(slices, ",  ")
	return fileName
}

func searchIssueOrPr(ctx *context.Context, TableName string, Key string, Page int, PageSize int, OnlyReturnNum bool, issueOrPr string) {

	/*
	   任务，合并请求  ES名称：issue-es-index
	   搜索:
	   name character varying(255) ,  标题
	   content text,  内容
	   comment text,  评论
	   排序:
	   updated_unix
	*/
	SortBy := ctx.Query("SortBy")
	ascending := ctx.QueryBool("Ascending")
	PrivateTotal := ctx.QueryInt("PrivateTotal")
	WebTotal := ctx.QueryInt("WebTotal")
	language := ctx.Query("language")
	if language == "" {
		language = "zh-CN"
	}
	from := (Page - 1) * PageSize
	if from == 0 {
		WebTotal = 0
	}
	resultObj := &SearchRes{}
	log.Info("WebTotal=" + fmt.Sprint(WebTotal))
	log.Info("PrivateTotal=" + fmt.Sprint(PrivateTotal))
	resultObj.Result = make([]map[string]interface{}, 0)
	isPull := false
	if issueOrPr == "t" {
		isPull = true
	}

	if ctx.User != nil && (from < PrivateTotal || from == 0) {

		log.Info("actor is null?:" + fmt.Sprint(ctx.User == nil))
		issues, count, err := models.SearchPrivateIssueOrPr(Page, PageSize, Key, isPull, ctx.User.ID)
		if err != nil {
			ctx.JSON(200, "")
			return
		}
		resultObj.PrivateTotal = count
		issuesSize := len(issues)
		if issuesSize > 0 {
			log.Info("Query private repo issue number is:" + fmt.Sprint(issuesSize) + " count=" + fmt.Sprint(count))
			makePrivateIssueOrPr(issues, resultObj, Key, language)
		} else {
			log.Info("not found private repo issue,keyword=" + Key)
		}
		if issuesSize >= PageSize {
			if WebTotal > 0 { //next page， not first query.
				resultObj.Total = int64(WebTotal)
				ctx.JSON(200, resultObj)
				return
			}
		}
	} else {
		resultObj.PrivateTotal = int64(PrivateTotal)
	}

	from = from - PrivateTotal
	if from < 0 {
		from = 0
	}
	Size := PageSize - len(resultObj.Result)

	boolQ := elastic.NewBoolQuery()
	isIssueQuery := elastic.NewTermQuery("is_pull", issueOrPr)

	if Key != "" {
		boolKeyQ := elastic.NewBoolQuery()
		log.Info("issue Key=" + Key)
		nameQuery := elastic.NewMatchQuery("name", Key).Boost(2).QueryName("f_first")
		contentQuery := elastic.NewMatchQuery("content", Key).Boost(1.5).QueryName("f_second")
		commentQuery := elastic.NewMatchQuery("comment", Key).Boost(1).QueryName("f_third")
		boolKeyQ.Should(nameQuery, contentQuery, commentQuery)
		boolQ.Must(isIssueQuery, boolKeyQ)
	} else {
		boolQ.Must(isIssueQuery)
	}

	res, err := client.Search(TableName).Query(boolQ).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Highlight(queryHighlight("name", "content", "comment")).Do(ctx.Req.Context())
	if err == nil {
		//searchJson, _ := json.Marshal(res)
		//log.Info("searchJson=" + string(searchJson))
		esresult := makeIssueResult(res, Key, OnlyReturnNum, language)

		resultObj.Total = resultObj.PrivateTotal + esresult.Total
		log.Info("query issue es count=" + fmt.Sprint(esresult.Total) + " total=" + fmt.Sprint(resultObj.Total))
		resultObj.Result = append(resultObj.Result, esresult.Result...)
		ctx.JSON(200, resultObj)
	} else {
		log.Info("query es error," + err.Error())
	}
}

func queryHighlight(names ...string) *elastic.Highlight {
	re := elastic.NewHighlight()
	for i := 0; i < len(names); i++ {
		field := &elastic.HighlighterField{
			Name: names[i],
		}
		re.Fields(field)
	}
	re.PreTags("<font color='red'>")
	re.PostTags("</font>")
	return re
}

func setRepoInfo(recordSource map[string]interface{}, record map[string]interface{}) {
	repoIdstr := recordSource["repo_id"].(string)
	repoId, cerr := strconv.ParseInt(repoIdstr, 10, 64)
	if cerr == nil {
		repo, errRepo := models.GetRepositoryByID(repoId)
		if errRepo == nil {
			log.Info("repo_url=" + repo.FullName())
			record["repoUrl"] = repo.FullName()
			record["avatar"] = repo.RelAvatarLink()
		} else {
			log.Info("repo err=" + errRepo.Error())
		}
	} else {
		log.Info("parse int err=" + cerr.Error())
	}
}

func makePrivateIssueOrPr(issues []*models.Issue, res *SearchRes, Key string, language string) {
	for _, issue := range issues {
		record := make(map[string]interface{})
		record["id"] = issue.ID
		record["repo_id"] = issue.RepoID

		repo, errRepo := models.GetRepositoryByID(issue.RepoID)
		if errRepo == nil {
			log.Info("repo_url=" + repo.FullName())
			record["repoUrl"] = repo.FullName()
			record["avatar"] = repo.RelAvatarLink()
		} else {
			log.Info("repo err=" + errRepo.Error())
		}
		record["name"] = makeHighLight(Key, issue.Title)
		record["content"] = truncLongText(makeHighLight(Key, issue.Content), true)

		if issue.IsPull {
			pr, err1 := issue.GetPullRequest()
			if err1 == nil && pr != nil {
				record["pr_id"] = pr.ID
			}
		}
		record["index"] = issue.Index
		record["num_comments"] = issue.NumComments
		record["is_closed"] = issue.IsClosed
		record["updated_unix"] = issue.UpdatedUnix
		record["updated_html"] = timeutil.TimeSinceUnix(issue.UpdatedUnix, language)
		res.Result = append(res.Result, record)
	}
}

func makeIssueResult(sRes *elastic.SearchResult, Key string, OnlyReturnNum bool, language string) *SearchRes {
	total := sRes.Hits.TotalHits.Value
	result := make([]map[string]interface{}, 0)
	if !OnlyReturnNum {
		for i, hit := range sRes.Hits.Hits {
			log.Info("this is issue query " + fmt.Sprint(i) + " result.")
			recordSource := make(map[string]interface{})
			source, err := hit.Source.MarshalJSON()

			if err == nil {
				err = json.Unmarshal(source, &recordSource)
				if err == nil {
					record := make(map[string]interface{})
					record["id"] = hit.Id
					record["repo_id"] = recordSource["repo_id"]
					log.Info("recordSource[\"repo_id\"]=" + fmt.Sprint(recordSource["repo_id"]))
					setRepoInfo(recordSource, record)
					record["name"] = getLabelValue("name", recordSource, hit.Highlight)
					if recordSource["content"] != nil {
						desc := getLabelValue("content", recordSource, hit.Highlight)
						record["content"] = dealLongText(desc, Key, hit.MatchedQueries)
						if _, ok := hit.Highlight["content"]; !ok {
							if _, ok_comment := hit.Highlight["comment"]; ok_comment {
								desc := getLabelValue("comment", recordSource, hit.Highlight)
								record["content"] = trimHrefHtml(dealLongText(desc, Key, hit.MatchedQueries))
							}
						}
					} else {
						if recordSource["comment"] != nil {
							desc := getLabelValue("comment", recordSource, hit.Highlight)
							record["content"] = dealLongText(desc, Key, hit.MatchedQueries)
						}
					}
					if recordSource["pr_id"] != nil {
						record["pr_id"] = recordSource["pr_id"]
					}
					log.Info("index=" + recordSource["index"].(string))
					record["index"] = recordSource["index"]
					record["num_comments"] = recordSource["num_comments"]
					record["is_closed"] = recordSource["is_closed"]
					record["updated_unix"] = recordSource["updated_unix"]
					setUpdateHtml(record, recordSource["updated_unix"].(string), language)
					result = append(result, record)
				} else {
					log.Info("deal issue source error," + err.Error())
				}
			} else {
				log.Info("deal issue source error," + err.Error())
			}
		}
	}
	returnObj := &SearchRes{
		Total:  total,
		Result: result,
	}

	return returnObj
}

func searchModel(ctx *context.Context, TableName string, Key string, Page int, PageSize int, OnlyReturnNum bool) {
	/*
	   模型，model-es-index
	   搜索:
	   name  ,  名称
	   description   描述
	   label  标签
	   file_name  数据集文件名称
	   排序:
	   download_count
	   reference_count
	   created_unix
	*/
	log.Info("query searchModel start")
	SortBy := ctx.Query("SortBy")
	ascending := ctx.QueryBool("Ascending")
	PrivateTotal := ctx.QueryInt("PrivateTotal")
	WebTotal := ctx.QueryInt("WebTotal")
	language := ctx.Query("language")
	if language == "" {
		language = "zh-CN"
	}
	from := (Page - 1) * PageSize
	if from == 0 {
		WebTotal = 0
	}
	resultObj := &SearchRes{}
	log.Info("WebTotal=" + fmt.Sprint(WebTotal))
	log.Info("PrivateTotal=" + fmt.Sprint(PrivateTotal))
	resultObj.Result = make([]map[string]interface{}, 0)

	if ctx.User != nil && (from < PrivateTotal || from == 0) {

		log.Info("actor is null?:" + fmt.Sprint(ctx.User == nil))
		sortBy := "ai_model_manage.reference_count desc,ai_model_manage.download_count desc,ai_model_manage.created_unix desc"
		if SortBy != "" && SortBy != "default" {
			tSortBy := SortBy
			if strings.HasSuffix(SortBy, ".keyword") {
				tSortBy = SortBy[0:(len(SortBy) - len(".keyword"))]
			}
			sortBy = "ai_model_manage." + tSortBy
			if ascending {
				sortBy += " asc"
			} else {
				sortBy += " desc"
			}
		}
		//Page, PageSize, Key, ctx.User.ID
		privateModels, count, err := models.QueryModelForSearch(&models.AiModelQueryOptions{
			ListOptions: models.ListOptions{
				Page:     Page,
				PageSize: PageSize,
			},
			UserID:   ctx.User.ID,
			Namelike: Key,
			SortType: sortBy,
		}, ctx.User.IsAdmin)
		if err != nil {
			ctx.JSON(200, "")
			return
		}
		resultObj.PrivateTotal = count
		modelSize := len(privateModels)
		if modelSize > 0 {
			log.Info("Query private model number is:" + fmt.Sprint(modelSize) + " count=" + fmt.Sprint(count))
			makePrivateModel(privateModels, resultObj, Key, language)
		} else {
			log.Info("not found private model, keyword=" + Key)
		}
		if modelSize >= PageSize {
			if WebTotal > 0 { //next page， not first query.
				resultObj.Total = int64(WebTotal)
				ctx.JSON(200, resultObj)
				return
			}
		}
	} else {
		resultObj.PrivateTotal = int64(PrivateTotal)
	}

	from = from - PrivateTotal
	if from < 0 {
		from = 0
	}
	Size := PageSize - len(resultObj.Result)

	boolQ := elastic.NewBoolQuery()
	if Key != "" {
		fileNameQuery := elastic.NewMatchQuery("file_name", Key).Boost(3).QueryName("f_first")
		nameQuery := elastic.NewMatchQuery("name", Key).Boost(2).QueryName("f_second")
		descQuery := elastic.NewMatchQuery("alias", Key).Boost(1.5).QueryName("f_three")
		labelQuery := elastic.NewMatchQuery("label", Key).Boost(1).QueryName("f_fourth")
		boolQ.Should(fileNameQuery, nameQuery, descQuery, labelQuery)
		res, err := client.Search(TableName).Query(boolQ).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Highlight(queryHighlight("file_name", "name", "alias", "label")).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeModelResult(res, Key, OnlyReturnNum, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			log.Info("query model es count=" + fmt.Sprint(esresult.Total) + " total=" + fmt.Sprint(resultObj.Total))
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
		}
	} else {
		log.Info("query all models.")
		//搜索的属性要指定{"timestamp":{"unmapped_type":"date"}}
		res, err := client.Search(TableName).SortBy(getSort(SortBy, ascending, "updated_unix.keyword", false)...).From(from).Size(Size).Do(ctx.Req.Context())
		if err == nil {
			//searchJson, _ := json.Marshal(res)
			//log.Info("searchJson=" + string(searchJson))
			esresult := makeModelResult(res, "", OnlyReturnNum, language)
			resultObj.Total = resultObj.PrivateTotal + esresult.Total
			log.Info("query model es count=" + fmt.Sprint(esresult.Total) + " total=" + fmt.Sprint(resultObj.Total))
			resultObj.Result = append(resultObj.Result, esresult.Result...)
			ctx.JSON(200, resultObj)
		} else {
			log.Info("query es error," + err.Error())
			ctx.JSON(200, "")
		}
	}

}

func makePrivateModel(privateModels []*models.AiModelManage, res *SearchRes, Key string, language string) {
	for _, model := range privateModels {
		record := make(map[string]interface{})

		record["id"] = model.ID
		userId := model.OwnerID

		user, errUser := models.GetUserByID(userId)
		if errUser == nil {
			record["owerName"] = user.GetDisplayName()
			record["avatar"] = user.RelAvatarLink()
		}

		repo, errRepo := models.GetRepositoryByID(model.RepoId)
		if errRepo == nil {
			log.Info("repo_url=" + repo.FullName())
			record["repoUrl"] = repo.FullName()
			record["avatar"] = repo.RelAvatarLink()
		} else {
			log.Info("repo err=" + errRepo.Error())
		}
		resultfile := models.QueryModelFileByModelId(model.ID)
		file_name := ""
		if resultfile != nil && len(resultfile) > 0 {
			for _, file := range resultfile {
				file_name += file.Name + ","
			}
			file_name = file_name[0 : len(file_name)-1]
		}
		record["file_name"] = truncLongText(makeHighLight(Key, file_name), true)
		record["name"] = makeHighLight(Key, model.Name)
		record["title"] = makeHighLight(Key, model.Alias)
		if record["title"] == "" {
			record["title"] = record["name"]
		}
		record["real_name"] = model.Name
		record["is_private"] = model.IsPrivate
		record["description"] = truncLongText(makeHighLight(Key, model.Description), true)

		record["label"] = makeHighLight(Key, model.Label)
		record["download_count"] = model.DownloadCount
		record["reference_count"] = model.ReferenceCount
		record["engine"] = model.Engine
		record["created_unix"] = model.CreatedUnix
		record["updated_unix"] = model.UpdatedUnix
		record["updated_html"] = timeutil.TimeSinceUnix(model.UpdatedUnix, language)

		res.Result = append(res.Result, record)
	}
}

func makeModelResult(sRes *elastic.SearchResult, Key string, OnlyReturnNum bool, language string) *SearchRes {
	total := sRes.Hits.TotalHits.Value
	result := make([]map[string]interface{}, 0)
	if !OnlyReturnNum {
		for i, hit := range sRes.Hits.Hits {
			log.Info("this is model query " + fmt.Sprint(i) + " result.")
			recordSource := make(map[string]interface{})
			source, err := hit.Source.MarshalJSON()

			if err == nil {
				err = json.Unmarshal(source, &recordSource)
				if err == nil {
					record := make(map[string]interface{})
					record["id"] = hit.Id
					userIdStr := recordSource["owner_id"].(string)
					userId, cerr := strconv.ParseInt(userIdStr, 10, 64)
					if cerr == nil {
						user, errUser := models.GetUserByID(userId)
						if errUser == nil {
							record["owerName"] = user.GetDisplayName()
							record["avatar"] = user.RelAvatarLink()
						}
					}
					setRepoInfo(recordSource, record)
					record["title"] = getLabelValue("alias", recordSource, hit.Highlight)
					record["name"] = getLabelValue("name", recordSource, hit.Highlight)
					if record["title"] == "" {
						record["title"] = record["name"]
					}
					record["real_name"] = recordSource["name"]

					record["label"] = getLabelValue("label", recordSource, hit.Highlight)
					if recordSource["description"] != nil {
						desc := getLabelValue("description", recordSource, hit.Highlight)
						record["description"] = dealLongText(desc, Key, hit.MatchedQueries)
					} else {
						record["description"] = ""
					}
					record["is_private"] = recordSource["is_private"]
					record["engine"] = recordSource["engine"]
					record["file_name"] = getDatasetFileName(getLabelValue("file_name", recordSource, hit.Highlight))
					record["reference_count"] = recordSource["reference_count"]
					record["download_count"] = recordSource["download_count"]
					record["created_unix"] = recordSource["created_unix"]
					setUpdateHtml(record, recordSource["updated_unix"].(string), language)
					result = append(result, record)
				} else {
					log.Info("deal model source error," + err.Error())
				}
			} else {
				log.Info("deal model source error," + err.Error())
			}
		}
	}
	returnObj := &SearchRes{
		Total:  total,
		Result: result,
	}

	return returnObj
}

// Tokenize 对输入文本进行分词：
//   - 若全为中文（含标点/空格），直接返回原字符串
//   - 若包含英文，对英文部分做驼峰/大写连词拆分，中文部分保持原样
//   - 多个词之间以空格连接
func Tokenize(input string) string {
	if input == "" {
		return ""
	}

	// 检测是否含有英文字母
	hasEnglish := false
	for _, r := range input {
		if unicode.IsLetter(r) && r <= 0x7F { // ASCII 字母 = 英文
			hasEnglish = true
			break
		}
	}

	// 纯中文/无英文 → 直接返回
	if !hasEnglish {
		return input
	}

	// 含英文 → 按字符类型切段，英文段做驼峰拆分
	var tokens []string
	var buf strings.Builder
	inEnglish := false

	flush := func() {
		s := buf.String()
		buf.Reset()
		if s == "" {
			return
		}
		if inEnglish {
			tokens = append(tokens, splitCamelCase(s)...)
		} else {
			tokens = append(tokens, s)
		}
	}

	for _, r := range input {
		isEn := (unicode.IsLetter(r) && r <= 0x7F) || unicode.IsDigit(r)

		if isEn {
			if !inEnglish {
				flush()
				inEnglish = true
			}
			buf.WriteRune(r)
		} else if unicode.IsSpace(r) {
			// 空格：结束当前段，跳过（最终用空格拼接）
			flush()
			inEnglish = false
		} else {
			// 中文字符或其他标点
			if inEnglish {
				flush()
				inEnglish = false
			}
			buf.WriteRune(r)
		}
	}
	flush()

	return strings.Join(tokens, " ")
}

// splitCamelCase 拆分驼峰或大写连写的英文单词
//
//	OpenClaw   → [Open Claw]
//	HTMLParser → [HTML Parser]
//	myHTTPServer → [my HTTP Server]
func splitCamelCase(s string) []string {
	runes := []rune(s)
	var words []string
	start := 0

	for i := 1; i < len(runes); i++ {
		prev, curr := runes[i-1], runes[i]

		// 小写 → 大写：Open|Claw
		if unicode.IsLower(prev) && unicode.IsUpper(curr) {
			words = append(words, string(runes[start:i]))
			start = i
			continue
		}

		// 大写序列 + 下一位是大写紧跟小写：HTM|L → HTML|Parser
		// 即：当前是大写，前一个也是大写，后一个是小写 → 在当前位置前切
		if i+1 < len(runes) {
			next := runes[i+1]
			if unicode.IsUpper(prev) && unicode.IsUpper(curr) && unicode.IsLower(next) {
				words = append(words, string(runes[start:i]))
				start = i
			}
		}
	}
	words = append(words, string(runes[start:]))
	return words
}
