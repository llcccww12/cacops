package modelscope

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

const defaultEndpoint = "https://modelscope.cn"

var httpClient = &http.Client{Timeout: 8 * time.Second}

type msListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Models []msModel `json:"models"`
		Datasets []msDataset `json:"datasets"`
		TotalCount int `json:"total_count"`
		PageNumber int `json:"page_number"`
		PageSize   int `json:"page_size"`
	} `json:"data"`
}

type msModel struct {
	ID           string   `json:"id"`
	DisplayName  string   `json:"display_name"`
	Description  string   `json:"description"`
	Downloads    int64    `json:"downloads"`
	Likes        int      `json:"likes"`
	License      string   `json:"license"`
	Tasks        []string `json:"tasks"`
	Tags         []string `json:"tags"`
	LastModified string   `json:"last_modified"`
	FileSize     int64    `json:"file_size"`
}

type msDataset struct {
	ID           string   `json:"id"`
	DisplayName  string   `json:"display_name"`
	Description  string   `json:"description"`
	Downloads    int64    `json:"downloads"`
	Likes        int      `json:"likes"`
	License      string   `json:"license"`
	Tasks        []string `json:"tasks"`
	Tags         []string `json:"tags"`
	LastModified string   `json:"last_modified"`
	FileSize     int64    `json:"file_size"`
}

func ListModels(search, sort string, page, pageSize int) ([]*entity.AimodelInfo, int64, error) {
	raw, err := fetch("models", search, sort, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var resp msListResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, 0, err
	}
	if !resp.Success {
		return nil, 0, fmt.Errorf("modelscope models api failed")
	}
	out := make([]*entity.AimodelInfo, 0, len(resp.Data.Models))
	for _, item := range resp.Data.Models {
		out = append(out, toAimodelInfo(item))
	}
	return out, int64(resp.Data.TotalCount), nil
}

func ListDatasets(search, sort string, page, pageSize int) ([]*entity.DatasetInfo, int64, error) {
	raw, err := fetch("datasets", search, sort, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var resp msListResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, 0, err
	}
	if !resp.Success {
		return nil, 0, fmt.Errorf("modelscope datasets api failed")
	}
	out := make([]*entity.DatasetInfo, 0, len(resp.Data.Datasets))
	for _, item := range resp.Data.Datasets {
		out = append(out, toDatasetInfo(item))
	}
	return out, int64(resp.Data.TotalCount), nil
}

func fetch(resource, search, sort string, page, pageSize int) ([]byte, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 30
	}
	if pageSize > 100 {
		pageSize = 100
	}
	q := url.Values{}
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("page_size", fmt.Sprintf("%d", pageSize))
	if search != "" {
		q.Set("search", search)
	}
	q.Set("sort", mapSort(sort))
	reqURL := fmt.Sprintf("%s/openapi/v1/%s?%s", defaultEndpoint, resource, q.Encode())
	resp, err := httpClient.Get(reqURL)
	if err != nil {
		log.Error("modelscope fetch failed: %v url=%s", err, reqURL)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("modelscope status %d", resp.StatusCode)
	}
	return body, nil
}

func mapSort(orderBy string) string {
	switch orderBy {
	case "newest", "recentupdate":
		return "last_modified"
	case "downloadcount", "usecount", "derivativecount", "collections":
		return "downloads"
	default:
		return "downloads"
	}
}

func splitID(id string) (owner, name string) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "modelscope", id
}

func parseUnix(ts string) timeutil.TimeStamp {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return timeutil.TimeStamp(time.Now().Unix())
	}
	return timeutil.TimeStamp(t.Unix())
}

func extractLibrary(tags []string) string {
	for _, tag := range tags {
		if strings.HasPrefix(tag, "library:") {
			return strings.TrimPrefix(tag, "library:")
		}
	}
	return ""
}

func toDatasetInfo(item msDataset) *entity.DatasetInfo {
	owner, name := splitID(item.ID)
	labels := make([]string, 0, len(item.Tags))
	for _, tag := range item.Tags {
		if strings.HasPrefix(tag, "custom_tag:") {
			labels = append(labels, strings.TrimPrefix(tag, "custom_tag:"))
		}
	}
	return &entity.DatasetInfo{
		ID:             "ms:" + item.ID,
		Name:           name,
		Alias:          firstNonEmpty(item.DisplayName, name),
		Tags:           labels,
		License:        item.License,
		Tasks:          item.Tasks,
		DownloadCount:  item.Downloads,
		NumStars:       item.Likes,
		UseCount:       int64(item.Likes),
		UpdatedUnix:    parseUnix(item.LastModified),
		Size:           item.FileSize,
		OwnerName:      owner,
		ExternalUrl:    defaultEndpoint + "/datasets/" + item.ID,
		ExternalSource: "modelscope",
	}
}

func toAimodelInfo(item msModel) *entity.AimodelInfo {
	owner, name := splitID(item.ID)
	labelParts := append([]string{}, item.Tasks...)
	if lib := extractLibrary(item.Tags); lib != "" {
		labelParts = append(labelParts, lib)
	}
	return &entity.AimodelInfo{
		ID:            "ms:" + item.ID,
		Name:          name,
		Alias:         firstNonEmpty(item.DisplayName, name),
		Label:         strings.Join(labelParts, " "),
		License:       item.License,
		DownloadCount: int(item.Downloads),
		NumStars:      item.Likes,
		UseCount:      item.Likes,
		UpdatedUnix:   parseUnix(item.LastModified),
		Size:          item.FileSize,
		OwnerName:     owner,
		AimodelType:   2,
		ExternalName:  item.ID,
		ExternalUrl:   defaultEndpoint + "/models/" + item.ID,
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}
