package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/setting"
	. "code.gitea.io/gitea/modules/urchin_v2/common"
	. "code.gitea.io/gitea/modules/urchin_v2/module"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

var UClient UrchinClient

type UrchinClient struct {
	addr         string
	header       http.Header
	urchinClient *retryablehttp.Client
}

func (u *UrchinClient) Init(
	ctx context.Context,
	address string,
	reqTimeout int64,
	maxConnection int) {

	u.addr = address

	timeout := time.Duration(reqTimeout) * time.Second

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(
			ctx context.Context,
			network,
			addr string) (net.Conn, error) {
			dialer := &net.Dialer{
				Timeout:   3 * time.Second,  // 连接超时时间
				KeepAlive: 30 * time.Second, // 保持连接时长
			}
			return dialer.DialContext(ctx, network, addr)
		},
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     90 * time.Second,
		MaxIdleConnsPerHost: maxConnection,
	}

	u.urchinClient = retryablehttp.NewClient()
	u.urchinClient.RetryMax = 3
	u.urchinClient.RetryWaitMin = 1 * time.Second
	u.urchinClient.RetryWaitMax = 5 * time.Second
	u.urchinClient.HTTPClient.Transport = transport
	u.urchinClient.HTTPClient.Timeout = timeout

	u.header = make(http.Header)

	u.GetToken()

}

func (u *UrchinClient) SetToken(token string) {
	u.header.Set(UrchinClientHeaderToken, token)
}

type GetTokenResult struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
}

type GetTokenParams struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

const errorIllegalToken = 1005

func (u *UrchinClient) GetToken() error {
	HOST = strings.TrimSuffix(setting.Grampus.Host, "/")

	client := getRestyClient()
	params := GetTokenParams{
		UserName: setting.Grampus.UserName,
		Password: setting.Grampus.Password,
	}

	var result GetTokenResult
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(strings.TrimSuffix(setting.Grampus.Host, "/") + "/openapi/v1/token")
	if err != nil {
		return fmt.Errorf("resty getToken: %v", err)
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("getToken failed:%s", res.String())
	}
	u.header.Set(UrchinClientHeaderToken, result.Token)
	return nil
}

var (
	restyClient *resty.Client
	HOST        string
	TOKEN       string
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func (u *UrchinClient) CreateInitiateMultipartUploadSignedUrl(
	ctx context.Context,
	req *CreateInitiateMultipartUploadSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateInitiateMultipartUploadSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateInitiateMultipartUploadSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateInitiateMultipartUploadSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateInitiateMultipartUploadSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateUploadPartSignedUrl(
	ctx context.Context,
	req *CreateUploadPartSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateUploadPartSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateUploadPartSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateUploadPartSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateUploadPartSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateListPartsSignedUrl(
	ctx context.Context,
	req *CreateListPartsSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateListPartsSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateListPartsSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateListPartsSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateListPartsSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateCompleteMultipartUploadSignedUrl(
	ctx context.Context,
	req *CreateCompleteMultipartUploadSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateCompleteMultipartUploadSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateCompleteMultipartUploadSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateCompleteMultipartUploadSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateCompleteMultipartUploadSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateAbortMultipartUploadSignedUrl(
	ctx context.Context,
	req *CreateAbortMultipartUploadSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateAbortMultipartUploadSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreatePutObjectSignedUrl(
	ctx context.Context,
	req *CreatePutObjectSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreatePutObjectSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreatePutObjectSignedUrlWithoutTask(
	ctx context.Context,
	req *CreatePutObjectSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreatePutObjectSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateGetObjectMetadataSignedUrl(
	ctx context.Context,
	req *CreateGetObjectMetadataSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateGetObjectMetadataSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateGetObjectMetadataSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateGetObjectMetadataSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateGetObjectMetadataSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateGetObjectSignedUrl(
	ctx context.Context,
	req *CreateGetObjectSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateGetObjectSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateGetObjectSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateGetObjectSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateGetObjectSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateListObjectsSignedUrl(
	ctx context.Context,
	req *CreateListObjectsSignedUrlReq) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateListObjectsSignedUrlInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateListObjectsSignedUrlWithoutTask(
	ctx context.Context,
	req *CreateListObjectsSignedUrlReqWithoutTask) (
	err error, resp *CreateSignedUrlResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateSignedUrlResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateListObjectsSignedUrlInterfaceWithoutTask,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) UploadObject(
	ctx context.Context,
	req *UploadObjectReq) (
	err error, resp *UploadObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(UploadObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientUploadObjectInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) UploadFile(
	ctx context.Context,
	req *UploadFileReq) (
	err error, resp *UploadFileResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(UploadFileResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientUploadFileInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) DownloadObject(
	ctx context.Context,
	req *DownloadObjectReq) (
	err error, resp *DownloadObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(DownloadObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientDownloadObjectInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) DownloadFile(
	ctx context.Context,
	req *DownloadFileReq) (
	err error, resp *DownloadFileResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(DownloadFileResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientDownloadFileInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) LoadObject(
	ctx context.Context,
	req *LoadObjectReq) (
	err error, resp *LoadObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(LoadObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientLoadObjectInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) MigrateObject(
	ctx context.Context,
	req *MigrateObjectReq) (
	err error, resp *MigrateObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(MigrateObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientMigrateObjectInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CopyObject(
	ctx context.Context,
	req *CopyObjectReq) (
	err error, resp *CopyObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CopyObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCopyObjectInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) GetObject(
	ctx context.Context,
	req *GetObjectReq) (
	err error, resp *GetObjectResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(GetObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientGetObjectInterface+"?"+values.Encode(),
		http.MethodGet,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) DeleteObject(
	ctx context.Context,
	req *DeleteObjectReq) (
	err error, resp *BaseResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientDeleteObjectInterface+"?"+values.Encode(),
		http.MethodDelete,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) DeleteObjectDeployment(
	ctx context.Context,
	req *DeleteObjectDeploymentReq) (
	err error, resp *BaseResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientDeleteObjectDeploymentInterface+"?"+values.Encode(),
		http.MethodDelete,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) DeleteFile(
	ctx context.Context,
	req *DeleteFileReq) (
	err error, resp *BaseResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientDeleteFileInterface+"?"+values.Encode(),
		http.MethodDelete,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) CreateObject(
	ctx context.Context,
	req *CreateObjectReq) (
	err error, resp *CreateObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientCreateObjectInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) ListObjects(
	ctx context.Context,
	req *ListObjectsReq) (
	err error, resp *ListObjectsResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(ListObjectsResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientListObjectsInterface+"?"+values.Encode(),
		http.MethodGet,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) GetObjectMetadata(
	ctx context.Context,
	req *GetObjectMetadataReq) (
	err error, resp *GetObjectMetadataResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(GetObjectMetadataResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientGetObjectMetadataInterface+"?"+values.Encode(),
		http.MethodGet,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) ListParts(
	ctx context.Context,
	req *ListPartsReq) (
	err error, resp *ListPartsResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(ListPartsResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientListPartsInterface+"?"+values.Encode(),
		http.MethodGet,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) GetTask(
	ctx context.Context,
	req *GetTaskReq) (
	err error, resp *GetTaskResp) {

	values, err := query.Values(req)
	if nil != err {
		return err, resp
	}

	resp = new(GetTaskResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientGetTaskInterface+"?"+values.Encode(),
		http.MethodGet,
		u.header,
		nil,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) FinishTask(
	ctx context.Context,
	req *FinishTaskReq) (
	err error, resp *BaseResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientFinishTaskInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) RetryTask(
	ctx context.Context,
	req *RetryTaskReq) (
	err error, resp *BaseResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientRetryTaskInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) ReportTaskHeartbeat(
	ctx context.Context,
	req *ReportTaskHeartbeatReq) (
	err error, resp *BaseResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientReportTaskHeartbeatInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) ReportSizeChanged(
	ctx context.Context,
	req *ReportSizeChangedReq) (
	err error, resp *BaseResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientReportSizeChangedInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) RelateObject(
	ctx context.Context,
	req *RelateObjectReq) (
	err error, resp *CreateObjectResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(CreateObjectResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientRelateObjectInterface,
		http.MethodPost,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}

func (u *UrchinClient) UpdateObjectLastModifyTime(
	ctx context.Context,
	req *UpdateObjectLastModifyTimeReq) (
	err error, resp *BaseResp) {

	reqBody, err := json.Marshal(req)
	if nil != err {
		return err, resp
	}

	resp = new(BaseResp)
	retry := 0
sendjob:
	err, respBody := Do(
		ctx,
		u.addr+UrchinClientUpdateObjectLastModifyTimeInterface,
		http.MethodPut,
		u.header,
		reqBody,
		u.urchinClient)
	if nil != err {
		return err, resp
	}

	err = json.Unmarshal(respBody, resp)
	if nil != err {
		return err, resp
	}

	if SuccessCode != resp.Code {
		if resp.Code == errorIllegalToken {
			retry++
			_ = u.GetToken()
			goto sendjob
		}
		return nil, resp
	}

	return nil, resp
}
