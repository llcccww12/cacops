package blockchain

import (
	"fmt"
	"strconv"

	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
)

const (
	UrlCreateAccount = "createAccount"
	UrlGetBalance    = "getBalance"
	UrlNewRepo       = "newRepo"
	UrlContribute    = "contribute"
	UrlSetIssue      = "setIssue"

	Success = 0
)

type CreateAccountResult struct {
	Code    int                    `json:"code"`
	Msg     string                 `json:"message"`
	Payload map[string]interface{} `json:"data"`
}

type GetBalanceResult struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data string `json:"data"`
}

type NewRepoResult struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	//Data 	string					`json:"data"`
}

type ContributeResult struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	//Payload map[string]interface{}	`json:"data"`
}

type SetIssueResult struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	//Data	string					`json:"data"`
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}

func CreateBlockchainAccount() (*CreateAccountResult, error) {
	client := getRestyClient()
	var result CreateAccountResult

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&result).
		Get(setting.BlockChainHost + UrlCreateAccount)

	if err != nil {
		return nil, fmt.Errorf("resty create account: %s", err)
	}

	if result.Code != Success {
		return &result, fmt.Errorf("CreateAccount err: %s", res.String())
	}

	return &result, nil
}

func NewRepo(repoID, publicKey, repoName string) (*NewRepoResult, error) {
	client := getRestyClient()
	var result NewRepoResult

	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"repoId":   repoID,
			"creator":  publicKey,
			"repoName": repoName,
		}).
		SetResult(&result).
		Get(setting.BlockChainHost + UrlNewRepo)

	if err != nil {
		return nil, fmt.Errorf("resty newRepo: %v", err)
	}

	if result.Code != Success {
		return &result, fmt.Errorf("newRepo err: %s", res.String())
	}

	return &result, nil
}

func GetBalance(contractAddress, contributor string) (*GetBalanceResult, error) {
	client := getRestyClient()
	var result GetBalanceResult

	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"contractAddress": contractAddress,
			"contributor":     contributor,
		}).
		SetResult(&result).
		Get(setting.BlockChainHost + UrlGetBalance)

	if err != nil {
		return nil, fmt.Errorf("resty getBalance: %v", err)
	}

	if result.Code != Success {
		return &result, fmt.Errorf("getBalance err: %s", res.String())
	}

	return &result, nil
}

func Contribute(contractAddress, contributor, commitId string, amount int64) (*ContributeResult, error) {
	client := getRestyClient()
	var result ContributeResult

	strAmount := strconv.FormatInt(amount, 10)
	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"contractAddress": contractAddress,
			"contributor":     contributor,
			"commitId":        commitId,
			"amount":          strAmount,
		}).
		SetResult(&result).
		Get(setting.BlockChainHost + UrlContribute)

	if err != nil {
		return nil, fmt.Errorf("resty contribute: %v", err)
	}

	if result.Code != Success {
		return &result, fmt.Errorf("contribute err: %s", res.String())
	}

	return &result, nil
}

func SetIssue(contractAddress, contributor string, issueId int64, amount int64) (*SetIssueResult, error) {
	client := getRestyClient()
	var result SetIssueResult

	strAmount := strconv.FormatInt(amount, 10)
	strIssue := strconv.FormatInt(issueId, 10)
	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"contractAddress": contractAddress,
			"contributor":     contributor,
			"issueId":         strIssue,
			"amount":          strAmount,
		}).
		SetResult(&result).
		Get(setting.BlockChainHost + UrlSetIssue)

	if err != nil {
		return nil, fmt.Errorf("resty SetIssue: %v", err)
	}

	if result.Code != Success {
		return &result, fmt.Errorf("SetIssue err: %s", res.String())
	}

	return &result, nil
}
