package hf_model_service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

const HF_MODEL_INFO_URL = "/api/models"
const HF_TOKEN_ACCESS_URL = "/resolve/main/.gitattributes"

var restyClient *resty.Client

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

type Lfs struct {
	Sha256      string `json:"sha256"`
	Size        int64  `json:"size"`
	PointerSize int    `json:"pointerSize"`
}

type Sibling struct {
	Rfilename string `json:"rfilename"`
	BlobId    string `json:"blobId"`
	Size      int64  `json:"size"`
	Lfs       *Lfs   `json:"lfs,omitempty"`
}

//type CardData struct {
//	License     string `json:"license"`
//	LicenseName string `json:"license_name"`
//	LicenseLink string `json:"license_link"`
//}

type HfModelInfo struct {
	ModelId string   `json:"modelId"`
	Sha     string   `json:"sha"`
	Tags    []string `json:"tags"`
	//CardData     CardData  `json:"cardData"`
	Siblings     []Sibling `json:"siblings"`
	Gated        string    `json:"gated"`
	LastModified time.Time `json:"lastModified"`
	CreatedAt    time.Time `json:"createdAt"`
}

// UnmarshalJSON
// "gated" returned by HF API can be either a boolean or a string
func (h *HfModelInfo) UnmarshalJSON(data []byte) error {
	type Alias HfModelInfo
	tmp := struct {
		*Alias
		Gated json.RawMessage `json:"gated"`
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var b bool
	if err := json.Unmarshal(tmp.Gated, &b); err == nil {
		// If it's a boolean, convert it to a string
		h.Gated = strconv.FormatBool(b)
		return nil
	}

	return json.Unmarshal(tmp.Gated, &h.Gated)
}

func GetHfModelInfoResty(hfRepoId string, hfToken string) (*HfModelInfo, int, error) {
	var result HfModelInfo

	if hfToken == "" {
		hfToken = setting.ExternalTransfer.HfToken
	}
	url := setting.ExternalTransfer.HfDomain + HF_MODEL_INFO_URL + "/" + hfRepoId
	client := getRestyClient()

	retry := 0

sendjob:
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+hfToken).
		SetQueryParam("blobs", "True").
		SetResult(&result).
		Get(url)

	if err != nil {
		log.Error("[hf_model] %s; url: %s", err.Error(), url)
		return nil, 500, err
	}

	if resp.StatusCode() != 200 {
		if retry < 1 {
			retry++
			goto sendjob
		} else {
			log.Error("[hf_model] Get %d; resp: %s; url: %s", resp.StatusCode(), resp.RawBody(), url)
			return nil, resp.StatusCode(), fmt.Errorf(resp.String())
		}
	}

	return &result, resp.StatusCode(), nil
}

func ValidTokenAccess(hfRepoId string, hfToken string) (int, error) {
	var result HfModelInfo

	if hfToken == "" {
		hfToken = setting.ExternalTransfer.HfToken
	}
	url := setting.ExternalTransfer.HfDomain + "/" + hfRepoId + HF_TOKEN_ACCESS_URL
	client := getRestyClient()
	fmt.Println("url: ", url)

	retry := 0

sendjob:
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+hfToken).
		SetResult(&result).
		Head(url)

	if err != nil {
		log.Error("[hf_model] %s; url: %s", err.Error(), url)
		return 500, err
	}

	if resp.StatusCode() != 200 && retry < 1 {
		retry++
		goto sendjob
	}

	return resp.StatusCode(), nil
}
