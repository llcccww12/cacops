package paratera

import (
	"code.gitea.io/gitea/modules/log"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
	"time"
)

/*
	并行科技对接api
	参考文档:https://ai.paratera.com/document/openapi/computer/production/DescribeInstanceTypes
*/

type ParateraAPI struct {
	ak string
	sk string
	//log *log.Helper
	client  *resty.Client
	baseUrl string
}

type BaseRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StopInstancesReq struct {
	ZoneCode    string   `json:"zoneCode"`
	EcsUuids    []string `json:"ecsUuids"`
	StopType    string   `json:"stopType"`
	StoppedMode string   `json:"stoppedMode"`
	ReleaseEip  bool     `json:"releaseEip"`
}

type DescribeInstanceTypesRsp struct {
	BaseRsp
	Data []*struct {
		VhostCpus     int64    `json:"vhostCpus"`
		VhostType     string   `json:"vhostType"`
		VhostModel    string   `json:"vhostModel"`
		VhostMemory   int64    `json:"vhostMemory"`
		VhostGpus     float64  `json:"vhostGpus"`
		VhostGpumem   int64    `json:"vhostGpumem"`
		RootDiskTypes []string `json:"rootDiskTypes"`
		Zone          *struct {
			ZoneCode string `json:"zoneCode"`
		} `json:"zone"`
	} `json:"data"`
}

func NewParateraAPI() (*ParateraAPI, error) {
	// 初始化基础的配置
	return &ParateraAPI{
		client:  resty.New(),
		ak:      "",
		sk:      "",
		baseUrl: "https://ai.blsc.cn",
	}, nil
}

type Request struct {
	HttpMethod  string
	QueryString string
	Body        interface{}
	Service     string
	Action      string
	Url         string
}

func (p *ParateraAPI) DoRequest(ctx context.Context, r *Request, response interface{}) error {
	var err error
	canonicalHeaders := "content-type:application/json; charset=utf-8\nhost:ai.blsc.cn"
	signedHeaders := "content-type;host"

	var body []byte
	b, ok := r.Body.(string)
	if ok {
		body = []byte(b)
	} else {
		body, err = json.Marshal(r.Body)
		if err != nil {
			return err
		}

	}
	body = []byte(strings.ReplaceAll(string(body), "\n", ""))
	sum256rp := sha256.Sum256(body)
	hashedRequestPayload := hex.EncodeToString(sum256rp[:])
	canonicalRequest := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v", r.HttpMethod, r.Url, r.QueryString, canonicalHeaders, signedHeaders, hashedRequestPayload)
	sum256cr := sha256.Sum256([]byte(canonicalRequest))
	hashedCanonicalRequest := hex.EncodeToString(sum256cr[:])

	stringToSign := fmt.Sprintf("HmacSHA256\nV3\n%v\n%v\nparatera/aicloud/%v\n%v", p.ak, r.Service, r.Service, hashedCanonicalRequest)
	mac := hmac.New(sha256.New, []byte("BC_SIGNATURE&"+p.sk))
	mac.Write([]byte(stringToSign))
	signature := hex.EncodeToString(mac.Sum(nil))

	headers := map[string]string{
		"X-AIC-Version":       "V3",
		"X-AIC-Action":        r.Action,
		"X-AIC-Timestamp":     strconv.Itoa(int(time.Now().Unix())),
		"X-AIC-AccessKey":     p.ak,
		"X-AIC-SignedHeaders": "content-type;host",
		"X-AIC-Signature":     signature,
		"X-AIC-Service":       r.Service,
		"Content-Type":        "application/json; charset=utf-8",
		"HOST":                "ai.blsc.cn",
	}
	url := fmt.Sprintf("%v%v", p.baseUrl, r.Url)
	req := p.client.R().SetHeaders(headers).SetBody(r.Body)
	var rsp *resty.Response
	switch strings.ToUpper(r.HttpMethod) {
	case "POST":
		rsp, err = req.Post(url)
	case "GET":
		rsp, err = req.Get(url)
	}
	if err != nil {
		log.Error("response error:%v", err)
		return err
	}

	if rsp.StatusCode() != 200 {
		br := &BaseRsp{}
		if err := json.Unmarshal([]byte(rsp.String()), br); err == nil {
			if br.Code != 0 {
				return err
			}
		}

		return err
	}

	if err := json.Unmarshal([]byte(rsp.String()), response); err != nil {
		return err
	}

	return nil
}

func (p *ParateraAPI) DescribeInstanceTypes(ctx context.Context) (*DescribeInstanceTypesRsp, error) {
	rsp := &DescribeInstanceTypesRsp{}

	err := p.DoRequest(ctx, &Request{
		HttpMethod:  "POST",
		QueryString: "",
		Body:        "{}",
		Service:     "product",
		Action:      "DescribeInstanceTypes",
		Url:         "/platform/v3/product/DescribeInstanceTypes",
	}, rsp)
	if err != nil {
		return nil, err
	}

	if rsp.Data == nil || len(rsp.Data) == 0 {
		return nil, err
	}

	return rsp, nil
}
