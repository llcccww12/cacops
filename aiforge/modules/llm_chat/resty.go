package llm_chat

import (
	"bufio"
	"bytes"
	"code.gitea.io/gitea/modules/httplib"
	"code.gitea.io/gitea/modules/log"
	constants "code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"mime/multipart"
	"net/http"
	"unicode/utf8"
)

var (
	restyClient *resty.Client
)

const (
	urlLLMChat                = "/chat/chat"
	urlKnowledgeBaseChat      = "/chat/knowledge_base_chat"
	urlKnowledgeBaseList      = "/knowledge_base/list_knowledge_bases"
	urlKnowledgeBaseCreate    = "/knowledge_base/create_knowledge_base"
	urlKnowledgeBaseDelete    = "/knowledge_base/delete_knowledge_base"
	urlKnowledgeBaseListFiles = "/knowledge_base/list_files"
	urlKnowledgeBaseSearchDoc = "/knowledge_base/search_docs"
	urlKnowledgeBaseUploadDoc = "/knowledge_base/upload_docs"
	urlKnowledgeBaseDeleteDoc = "/knowledge_base/delete_docs"
	urlKnowledgeBaseUpdateDoc = "/knowledge_base/update_docs"
	urlKnowledgeBaseDownload  = "/knowledge_base/download_doc"
	urlKnowledgeBaseRecreate  = "/knowledge_base/recreate_vector_store"
)

func GetEndpoint(modelName string) string {
	//endpoint := constants.LLM_CHAT_API.HOST + ":"
	var endpoint string
	switch modelName {
	case "chatglm2-6b":
		endpoint = constants.LLM_CHAT_API.CHATGLM2_HOST
	case "llama2-7b-chat-hf":
		endpoint = constants.LLM_CHAT_API.LLAMA2_HOST
	default:
		endpoint = constants.LLM_CHAT_API.CHATGLM2_HOST
	}

	return endpoint
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

// Custom function to split by character
func scanByCharacter(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	return 1, data[:1], nil
}

func StreamLLMChat(params api.LLMChatMessage, resultChan chan string, errChan chan error, done chan struct{}) {
	client := httplib.NewClientTimeOut(&httplib.Config{})
	endpoint := GetEndpoint(params.ModelName)
	requestBody, _ := json.Marshal(params)
	log.Info("Request body: %s\n", requestBody)
	request, err := http.NewRequest("POST", endpoint+urlLLMChat, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error creating request: %v", err)
		errChan <- err
		return
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		errChan <- err
		return
	}
	defer resp.Body.Close()
	log.Info("Response status: %s\n", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(scanByCharacter)
	var invalidCharBuffer string
	for scanner.Scan() {
		char := scanner.Text()
		if len(invalidCharBuffer) > 0 {
			char = invalidCharBuffer + char
			invalidCharBuffer = ""
		}
		if utf8.ValidString(char) {
			//runes := []rune(char)
			//log.Info("%s -> %U \n", char, runes[0])
			resultChan <- char
		} else {
			invalidCharBuffer += char
		}
	}
	if len(invalidCharBuffer) > 0 {
		log.Info("Unprocessed invalid UTF-8 characters: %s\n", invalidCharBuffer)
	}
	close(done)

	if scanner.Err() != nil {
		errChan <- scanner.Err()
	}
}

func StreamKBChat(params api.KBChatMessage, resultChan chan string, errChan chan error, done chan struct{}) {
	client := httplib.NewClientTimeOut(&httplib.Config{})
	endpoint := GetEndpoint(params.ModelName)
	requestBody, _ := json.Marshal(params)
	log.Info("Request body: %s\n", requestBody)
	request, err := http.NewRequest("POST", endpoint+urlKnowledgeBaseChat, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error creating request: %v", err)
		errChan <- err
		return
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		errChan <- err
		return
	}
	defer resp.Body.Close()
	log.Info("Response status: %s\n", resp.Status)

	//Create a buffer to read 2048-byte blocks
	buffer := make([]byte, 4096)
	for {
		// Read a 4096-byte block from the response body
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err != io.EOF {
				errChan <- err
			}
			break
		}
		resultChan <- string(buffer[:n])
	}
	close(done)
}

func SendLLMChat(params api.LLMChatMessage) (*LLMChatResponse, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint(params.ModelName)

	request, _ := json.Marshal(params)
	log.Info("resty request body: %s", request)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		Post(endpoint + urlLLMChat)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	result := LLMChatResponse{
		Answer: res.String(),
	}
	log.Info("resty response: %+v", result)

	if err != nil {
		return &result, fmt.Errorf("resty SendLLMChat(): %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func SendKBChat(params api.KBChatMessage) (*KBChatResponse, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint(params.ModelName)
	var result KBChatResponse

	request, _ := json.Marshal(params)
	log.Info("resty request body: %s", request)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		Post(endpoint + urlKnowledgeBaseChat)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response := res.String()
	json.Unmarshal([]byte(response), &result)
	log.Info("resty response: %+v", result)

	if err != nil {
		return &result, fmt.Errorf("resty SendLLMChat(): %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func ListKnowledgeBase() (*LLMBasicMsgWithData, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMBasicMsgWithData

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&result).
		Get(endpoint + urlKnowledgeBaseList)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if err != nil {
		return &result, fmt.Errorf("resty ListKnowledgeBase(): %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func CreateKnowledgeBase(params api.CreateKnowledgeBaseParams) (*LLMBasicMsg, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMBasicMsg

	request, _ := json.Marshal(params)
	log.Info("resty request body: %s", request)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(endpoint + urlKnowledgeBaseCreate)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if err != nil {
		return &result, fmt.Errorf("resty CreateKnowledgeBase(): %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func DeleteKnowledgeBase(knowledgeBaseName string) (*LLMBasicMsgWithData, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMBasicMsgWithData

	log.Info("resty request body: %s", knowledgeBaseName)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/text").
		SetBody(knowledgeBaseName).
		SetResult(&result).
		Post(endpoint + urlKnowledgeBaseDelete)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if err != nil {
		log.Error("resty DeleteKnowledgeBase(): %s", err)
		return &result, fmt.Errorf("resty DeleteKnowledgeBase(): %s", err)
	}
	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func KBListFiles(knowledgeBaseName string) (*LLMBasicMsgWithData, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMBasicMsgWithData

	log.Info("resty request body: %s", knowledgeBaseName)

sendjob:
	res, err := client.R().
		SetQueryParams(map[string]string{
			"knowledge_base_name": knowledgeBaseName,
		}).
		SetHeader("Content-Type", "application/text").
		SetBody(knowledgeBaseName).
		SetResult(&result).
		Get(endpoint + urlKnowledgeBaseListFiles)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if err != nil {
		log.Error("resty KBListFiles(): %s", err)
		return &result, fmt.Errorf("resty KBListFiles(): %s", err)
	}
	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func KBSearchDoc(params api.SearchDocParams) (*SearchDocResponse, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result []SearchDocResult

	log.Info("resty request body: %+v", params)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(endpoint + urlKnowledgeBaseSearchDoc)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	resultAPI := SearchDocResponse{
		Results: result,
	}
	response, _ := json.Marshal(resultAPI)
	log.Info("resty response: %s", response)

	if err != nil {
		log.Error("resty KBListFiles(): %s", err)
		return &resultAPI, fmt.Errorf("resty KBListFiles(): %s", err)
	}
	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &resultAPI, nil
}

func KBDeleteDoc(params api.DeleteDocParams) (interface{}, error) { // *LLMDeleteDocMsg, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMDeleteDocMsg

	request, _ := json.Marshal(params)
	log.Info("resty request body: %s", request)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(endpoint + urlKnowledgeBaseDeleteDoc)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)
	if err != nil {
		var errResult LLMErrorMsg
		json.Unmarshal([]byte(res.String()), &errResult)
		return &errResult, fmt.Errorf("resty KBDeleteDoc(): %s", err)
	}

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func KBUpdateDoc(params api.UpdateDocParams) (*LLMBasicMsg, error) {
	client := getRestyClient()
	retry := 0
	endpoint := GetEndpoint("")
	var result LLMBasicMsg

	request, _ := json.Marshal(params)
	log.Info("resty request body: %s", request)

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(endpoint + urlKnowledgeBaseUpdateDoc)

	log.Info("resty status: %+v, route: %+v", res.StatusCode(), res.Request.URL)

	response, _ := json.Marshal(result)
	log.Info("resty response: %s", response)

	if err != nil {
		return &result, fmt.Errorf("resty KBDeleteDoc(): %s", err)
	}

	if res.StatusCode() == http.StatusUnauthorized && retry < 1 {
		retry++
		goto sendjob
	}
	return &result, nil
}

func KBRecreateVectorStore(params api.RecreateVectorStoreParams, resultChan chan string, errChan chan error, done chan struct{}) {
	client := httplib.NewClientTimeOut(&httplib.Config{})
	endpoint := GetEndpoint("")
	requestBody, _ := json.Marshal(params)
	log.Info("Request body: %s\n", requestBody)
	request, err := http.NewRequest("POST", endpoint+urlKnowledgeBaseRecreate, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error creating request: %v", err)
		errChan <- err
		return
	}

	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Error("Error sending request: %v", err)
		errChan <- err
		return
	}
	defer resp.Body.Close()
	log.Info("Response status: %s\n", resp.Status)

	// Create a buffer to read 2048-byte blocks
	buffer := make([]byte, 4096)
	for {
		// Read a 2048-byte block from the response body
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err != io.EOF {
				errChan <- err
			}
			break
		}
		resultChan <- string(buffer[:n])
	}
	close(done)
}

func GetUploadDocUrl() (string, error) {
	endpoint := GetEndpoint("") + urlKnowledgeBaseUploadDoc
	log.Info("resty GetUploadDocUrl: %s", endpoint)
	return endpoint, nil
}

func writeDocs(fileHeader *multipart.FileHeader, writer *multipart.Writer) error {
	filename := fileHeader.Filename
	file, err := fileHeader.Open()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer file.Close()
	part, err := writer.CreateFormFile("files", filename)
	if err != nil {
		log.Error("Error creating form file:", err)
		return err
	}
	_, err = io.Copy(part, file)
	return nil
}

func writeBody(requestBody *bytes.Buffer, form api.LLMChatUploadForm) (string, error) {
	writer := multipart.NewWriter(requestBody)
	defer writer.Close()
	err := writer.WriteField("knowledge_base_name", form.KnowledgeBaseName)
	if err != nil {
		log.Error("failed to create upload_doc() writer")
		return "", err
	}
	for _, fileHeader := range form.Files {
		err = writeDocs(fileHeader, writer)
		if err != nil {
			log.Error("Error getting doc content: %s", err)
			return "", err
		}
	}
	return writer.FormDataContentType(), nil
}

func UploadDocs(modelName string, form api.LLMChatUploadForm) (*map[string]interface{}, error) {
	log.Info("######### received by resty\n")

	var requestBody bytes.Buffer
	headerValue, err := writeBody(&requestBody, form)
	if err != nil {
		log.Error("upload docs write body failed.")
		return nil, err
	}

	endpoint := GetEndpoint(modelName)
	req, err := http.NewRequest("POST", endpoint+urlKnowledgeBaseUploadDoc, &requestBody)
	if err != nil {
		log.Info("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", headerValue)

	client := httplib.NewClientTimeOut(&httplib.Config{})
	resp, err := client.Do(req)
	if err != nil {
		log.Info("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Info("############## Response Status:", resp.Status)

	var errResult map[string]interface{}
	//if resp.StatusCode == http.StatusUnprocessableEntity {
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Info("Error reading response body:", err)
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &errResult)
		log.Error("##############upload_docs() errResult: %+v\n", errResult)
		return &errResult, nil
	}
	log.Info("############## Response Body: %+v\n", resp.Body)

	// Parse the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Info("Error decoding response:", err)
		return nil, err
	}
	return &result, nil
}

func GetDownloadDocUrl(knowledgeBaseName string, fileName string) (string, error) {
	endpoint := GetEndpoint("") + urlKnowledgeBaseDownload
	params := "?knowledge_base_name=" + knowledgeBaseName + "&file_name=" + fileName
	log.Info("resty GetDownloadDocUrl: %s", endpoint+params)
	return endpoint + params, nil
}
