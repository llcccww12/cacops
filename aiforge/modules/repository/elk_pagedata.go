package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

// 输入elk的json结构begin
type InputInfo struct {
	Batch []Batch `json:"batch"`
}
type Fields struct {
	Field  string `json:"field"`
	Format string `json:"format"`
}
type MatchPhrase struct {
	Message string `json:"message,omitempty"`
	TagName string `json:"tagName.keyword,omitempty"`
}
type Should struct {
	MatchPhrase MatchPhrase `json:"match_phrase"`
}
type Bool struct {
	Should             []Should `json:"should"`
	MinimumShouldMatch int      `json:"minimum_should_match"`
}
type Timestamptest struct {
	// Gte time.Time `json:"gte"`
	Gte    string `json:"gte"`
	Lte    string `json:"lte"`
	Format string `json:"format"`
}
type Range struct {
	Timestamptest Timestamptest `json:"@timestamptest"`
}

type FilterMatchPhrase struct {
	UserName    string `json:"userName.keyword,omitempty"`
	ProjectName string `json:"projectName.keyword,omitempty"`
	TagName     string `json:"tagName.keyword,omitempty"`
}

type Filter struct {
	Bool              *Bool              `json:"bool,omitempty"`
	Range             *Range             `json:"range,omitempty"`
	FilterMatchPhrase *FilterMatchPhrase `json:"match_phrase,omitempty"`
}
type MustNotMatchPhrase struct {
	ProjectName string `json:"projectName"`
}
type MustNot struct {
	MustNotMatchPhrase MustNotMatchPhrase `json:"match_phrase"`
}
type BoolIn struct {
	Filter  []Filter  `json:"filter"`
	MustNot []MustNot `json:"must_not"`
}
type Query struct {
	BoolIn BoolIn `json:"bool"`
}
type Body struct {
	Size   int      `json:"size"`
	Fields []Fields `json:"fields"`
	Query  Query    `json:"query"`
}
type Params struct {
	Index string `json:"index"`
	Body  Body   `json:"body"`
}
type Request struct {
	Params Params `json:"params"`
}
type Batch struct {
	Request Request `json:"request"`
}

//输入elk的json结构end

// elk输出的json结构begin
type Hits struct {
	Total int `json:"total"`
}
type RawResponse struct {
	Hits Hits `json:"hits"`
}
type Result struct {
	RawResponse RawResponse `json:"rawResponse"`
	Loaded      int         `json:"loaded"`
}
type ResultInfo struct {
	Id     int    `json:"id"`
	Result Result `json:"result"`
}

//elk输出的json结构end

// 处理返回的elk数据，只保留totalView，即访问量；loaded是分片载入次数，用来判断返回的数据是否准确
func GetResultFromElk(resultInfo ResultInfo, jsonStr []byte) (loaded int, totalView int, err error) {
	ElkBase64Init := setting.ElkUser + ":" + setting.ElkPassword
	ElkBase64 := base64.StdEncoding.EncodeToString([]byte(ElkBase64Init))
	BasicElkBase64 := "Basic" + " " + ElkBase64
	url := setting.ElkUrl
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("kbn-version", "7.13.2")
	req.Header.Set("Authorization", BasicElkBase64)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return 0, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error("ConnectToElk failed:%s", resp.Status)
		return 0, 0, fmt.Errorf("ConnectToElk failed:%s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	err = json.Unmarshal([]byte(string(body)), &resultInfo)
	if err != nil {
		log.Error("Get resultJson failed", err)
		return 0, 0, err
	}

	return resultInfo.Result.Loaded, resultInfo.Result.RawResponse.Hits.Total, err
}

// 初始化传给elk的数据结构，给定用户名和项目名，查询的起止时间，返回初始化后的结构
func ProjectViewInit(User string, Project string, Gte string, Lte string) (projectViewInit InputInfo) {
	var inputStruct InputInfo
	inputStruct.Batch = make([]Batch, 1)
	inputStruct.Batch[0].Request.Params.Index = setting.Index
	inputStruct.Batch[0].Request.Params.Body.Size = 0
	inputStruct.Batch[0].Request.Params.Body.Fields = make([]Fields, 1)
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Field = setting.TimeField
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Format = setting.ElkTimeFormat
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter = make([]Filter, 4)
	//限定查询时间
	var timeRange Range
	timeRange.Timestamptest.Gte = Gte
	timeRange.Timestamptest.Lte = Lte
	timeRange.Timestamptest.Format = "strict_date_optional_time"
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[0].Range = &timeRange
	//限定用户
	var userName FilterMatchPhrase
	userName.UserName = User
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[1].FilterMatchPhrase = &userName
	//限定项目
	var projectName FilterMatchPhrase
	projectName.ProjectName = Project
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[2].FilterMatchPhrase = &projectName
	//限定页面
	var bool Bool
	bool.Should = make([]Should, len(setting.PROJECT_LIMIT_PAGES))
	for i, pageName := range setting.PROJECT_LIMIT_PAGES {
		bool.Should[i].MatchPhrase.TagName = pageName
	}
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[3].Bool = &bool
	return inputStruct
}

// 初始化传给elk的数据结构，给定查询信息和非项目名，查询的起止时间，返回初始化后的结构
func AllProjectViewInit(MessageInfo string, NotProject string, Gte string, Lte string) (allProjectViewInit InputInfo) {
	var inputStruct InputInfo
	inputStruct.Batch = make([]Batch, 1)
	inputStruct.Batch[0].Request.Params.Index = setting.Index
	inputStruct.Batch[0].Request.Params.Body.Size = 0
	inputStruct.Batch[0].Request.Params.Body.Fields = make([]Fields, 1)
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Field = setting.TimeField
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Format = setting.ElkTimeFormat
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter = make([]Filter, 2)
	//限定message
	var bool Bool
	bool.Should = make([]Should, 1)
	bool.Should[0].MatchPhrase.Message = MessageInfo
	bool.MinimumShouldMatch = 1
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[0].Bool = &bool
	//限定查询时间
	var timeRange Range
	timeRange.Timestamptest.Gte = Gte
	timeRange.Timestamptest.Lte = Lte
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[1].Range = &timeRange
	//限定非项目
	// var boolIn BoolIn
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.MustNot = make([]MustNot, 1)
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.MustNot[0].MustNotMatchPhrase.ProjectName = NotProject
	return inputStruct
}

// 初始化传给elk的数据结构，给定查询信息和tagName，查询的起止时间，返回初始化后的结构
func TagNameInit(MessageInfo string, Tagname string, Gte string, Lte string) (projectViewInit InputInfo) {
	var inputStruct InputInfo
	inputStruct.Batch = make([]Batch, 1)
	inputStruct.Batch[0].Request.Params.Index = setting.Index
	inputStruct.Batch[0].Request.Params.Body.Size = 0
	inputStruct.Batch[0].Request.Params.Body.Fields = make([]Fields, 1)
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Field = setting.TimeField
	inputStruct.Batch[0].Request.Params.Body.Fields[0].Format = setting.ElkTimeFormat
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter = make([]Filter, 3)
	//限定message
	var bool Bool
	bool.Should = make([]Should, 1)
	bool.Should[0].MatchPhrase.Message = MessageInfo
	bool.MinimumShouldMatch = 1
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[0].Bool = &bool
	//限定tagName
	var tagName FilterMatchPhrase
	tagName.TagName = Tagname
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[1].FilterMatchPhrase = &tagName
	//限定查询时间
	var timeRange Range
	timeRange.Timestamptest.Gte = Gte
	timeRange.Timestamptest.Lte = Lte
	inputStruct.Batch[0].Request.Params.Body.Query.BoolIn.Filter[2].Range = &timeRange
	return inputStruct
}

// 向elk发送请求，将获取的结果只保留访问量，输入是初始化后的数据结构，返回访问量
func ViewInfo(viewInfo InputInfo) (totalView int, err error) {
	// jsons, err := json.Marshal(viewInfo)
	// if err != nil {
	// 	log.Error("jsons failed", err)
	// 	return 0, err
	// }
	// var jsonStr = []byte(jsons)
	// var resultInfo ResultInfo
	// loaded, totalView, err := GetResultFromElk(resultInfo, jsonStr)

	// time := 0
	// for {
	// 	if loaded == 0 {
	// 		loaded_next, totalView, err := GetResultFromElk(resultInfo, jsonStr)
	// 		time++
	// 		if loaded_next != 0 && time < 100 {
	// 			return totalView, err
	// 		}
	// 		if time > 100 {
	// 			break
	// 		}
	// 	} else {
	// 		break
	// 	}
	// }
	return totalView, err
}

// @title    AppointProjectView
// @description   获取指定用户和项目的访问量
// @param     User       string      "用户名"
// @param     Project    string      "项目名"
// @param     Gte        string      "起始时间"  如time.Now().AddDate(0, 0, -1).Format(time.RFC3339)
// @param     Lte        string      "结束时间"  如time.Now().Format(time.RFC3339)
// @return    totalView      int         "访问量"
func AppointProjectView(User string, Project string, Gte string, Lte string) (totalView int, err error) {
	ProjectViewInitInfo := ProjectViewInit(User, Project, Gte, Lte)
	ProjectTotalView, err := ViewInfo(ProjectViewInitInfo)
	return ProjectTotalView, err
}

// 统计项目相关页面的访问量
type ProjectInfo struct {
	/* 统计所有项目中该页面的浏览情况，不需要区分项目。以aiforge项目为例 */
	//地址：https://git.openi.org.cn/OpenI/aiforge/datasets?type=0
	Project_dataset_type_0 int
	//地址：https://git.openi.org.cn/OpenI/aiforge/datasets?type=1
	Project_dataset_type_1 int
	//地址：https://git.openi.org.cn/OpenI/aiforge/issues
	Project_issues int
	//地址：https://git.openi.org.cn/OpenI/aiforge/labels
	Project_labels int
	//地址：https://git.openi.org.cn/OpenI/aiforge/milestones
	Project_milestones int
	//地址：https://git.openi.org.cn/OpenI/aiforge/pulls
	Project_pulls int
	//地址：https://git.openi.org.cn/OpenI/aiforge/release
	Project_release int
	//地址：https://git.openi.org.cn/OpenI/aiforge/wiki
	Project_wiki int
	//地址：https://git.openi.org.cn/OpenI/aiforge/activity
	Project_activity int
	//地址：https://git.openi.org.cn/OpenI/aiforge/cloudbrain
	Project_cloudbrain int
	//地址：https://git.openi.org.cn/OpenI/aiforge/modelarts
	Project_modelarts int
	//地址：https://git.openi.org.cn/OpenI/aiforge/blockchain
	Project_blockchain int
	//地址：https://git.openi.org.cn/OpenI/aiforge/watchers
	Project_watchers int
	//地址：https://git.openi.org.cn/OpenI/aiforge/stars
	Project_stars int
	//地址：https://git.openi.org.cn/OpenI/aiforge/forks
	Project_forks int
}

type ErrorInfo struct {
	Project_dataset_type_0 error
	Project_dataset_type_1 error
	Project_issues         error
	Project_labels         error
	Project_milestones     error
	Project_pulls          error
	Project_release        error
	Project_wiki           error
	Project_activity       error
	Project_cloudbrain     error
	Project_modelarts      error
	Project_blockchain     error
	Project_watchers       error
	Project_stars          error
	Project_forks          error
}

// @title    AllProjectView
// @description   获取指定用户和项目的访问量
// @param     Gte        string      "起始时间"  如time.Now().AddDate(0, 0, -1).Format(time.RFC3339)
// @param     Lte        string      "结束时间"
// @return    projectInfo      ProjectInfo         "统计所有项目中页面的浏览情况，不需要区分项目"
func AllProjectView(Gte string, Lte string) (projectViewInfo ProjectInfo, errorInfo ErrorInfo) {
	projectViewInfo.Project_dataset_type_0, errorInfo.Project_dataset_type_0 = ViewInfo(AllProjectViewInit("/datasets?type=0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_dataset_type_1, errorInfo.Project_dataset_type_1 = ViewInfo(AllProjectViewInit("/datasets?type=1", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_issues, errorInfo.Project_issues = ViewInfo(AllProjectViewInit("/issues HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_labels, errorInfo.Project_labels = ViewInfo(TagNameInit("/labels HTTP/2.0", "labels", Gte, Lte))
	projectViewInfo.Project_milestones, errorInfo.Project_milestones = ViewInfo(AllProjectViewInit("/milestones HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_pulls, errorInfo.Project_pulls = ViewInfo(AllProjectViewInit("/pulls HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_release, errorInfo.Project_release = ViewInfo(AllProjectViewInit("/release HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_wiki, errorInfo.Project_wiki = ViewInfo(AllProjectViewInit("/wiki HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_activity, errorInfo.Project_activity = ViewInfo(AllProjectViewInit("/activity HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_cloudbrain, errorInfo.Project_cloudbrain = ViewInfo(AllProjectViewInit("/cloudbrain HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_modelarts, errorInfo.Project_modelarts = ViewInfo(AllProjectViewInit("/modelarts HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_blockchain, errorInfo.Project_blockchain = ViewInfo(AllProjectViewInit("/blockchain HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_watchers, errorInfo.Project_watchers = ViewInfo(AllProjectViewInit("/watchers HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_stars, errorInfo.Project_stars = ViewInfo(AllProjectViewInit("/stars HTTP/2.0", "%{[request][2]}", Gte, Lte))
	projectViewInfo.Project_forks, errorInfo.Project_forks = ViewInfo(AllProjectViewInit("/forks HTTP/2.0", "%{[request][2]}", Gte, Lte))
	return projectViewInfo, errorInfo
}
