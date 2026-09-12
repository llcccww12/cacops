package avatar

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts_gateway/core"
	"code.gitea.io/gitea/modules/setting"
)

var restyClient *http.Client

type DiffusionPromopt struct {
	Prompt             string  `json:"prompt"`
	NegativePrompt     string  `json:"negative_prompt"`
	Width              int     `json:"width"`
	Height             int     `json:"height"`
	Steps              int     `json:"steps"`
	GuidanceScale      float32 `json:"guidance_scale"`
	NumImagesPerPrompt int     `json:"num_images_per_prompt"`
}
type DiffusionResponse struct {
	ImageBase64 []string `json:"image_base64_list"`
}

func getRestyClient() *http.Client {

	if restyClient == nil {
		restyClient = &http.Client{
			Timeout:   300 * time.Second,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}
	}

	return restyClient
}

func RandomDiffusionBase64Image(size int, promote string) (string, error) {
	var result DiffusionResponse
	client := getRestyClient()
	reqBytes, reqErr := json.Marshal(DiffusionPromopt{Prompt: promote, Width: size, Height: size, Steps: 22, GuidanceScale: 7, NumImagesPerPrompt: 1})
	if reqErr != nil {
		log.Error("Error marshaling request parameters: %v", reqErr)
		return "", fmt.Errorf("error marshaling request parameters: %w", reqErr)
	}

	request, err := http.NewRequest(http.MethodPost, setting.AIAVartarApi, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Error("Error creating new HTTP request: %v", err)
		return "", fmt.Errorf("error creating new HTTP request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")
	if setting.AIAVartarAK != "" {
		s := core.Signer{
			Key:    setting.AIAVartarAK,
			Secret: setting.AIAVartarSK,
		}
		s.Sign(request)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("Error reading response body: %v", err)
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// 将字节数组解析为 JSON
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Error("Error unmarshaling response body: %v", err)
		return "", fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return result.ImageBase64[0], nil
}

func RandomDiffusionImageSize(size int, promote string) (image.Image, []byte, error) {
	var result DiffusionResponse
	client := getRestyClient()
	reqBytes, reqErr := json.Marshal(DiffusionPromopt{Prompt: promote, Width: size, Height: size, Steps: 22, GuidanceScale: 7, NumImagesPerPrompt: 1})
	if reqErr != nil {
		log.Error("Error marshaling request parameters: %v", reqErr)
		return nil, nil, fmt.Errorf("error marshaling request parameters: %w", reqErr)
	}

	request, err := http.NewRequest(http.MethodPost, setting.AIAVartarApi, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Error("Error creating new HTTP request: %v", err)
		return nil, nil, fmt.Errorf("error creating new HTTP request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")
	if setting.AIAVartarAK != "" {
		s := core.Signer{
			Key:    setting.AIAVartarAK,
			Secret: setting.AIAVartarSK,
		}
		s.Sign(request)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		return nil, nil, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("Error reading response body: %v", err)
		return nil, nil, fmt.Errorf("error reading response body: %w", err)
	}

	// 将字节数组解析为 JSON
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Error("Error unmarshaling response body: %v", err)
		return nil, nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return Base64ToImage(result.ImageBase64[0])
}

func Base64ToImage(base64Str string) (image.Image, []byte, error) {
	// 解码Base64字符串
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, nil, err
	}

	// 将字节数组转为image.Image
	imageDecode, _, err := image.Decode(bytes.NewReader(data))

	return imageDecode, data, err
}
