package point

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// 配置参数
const (
	dbHost         = "127.0.0.1"
	dbPort         = 55435
	dbUser         = "gitea"
	dbPassword     = "gitea"
	dbName         = "gitea"
	db2Host        = "127.0.0.1"
	db2Port        = 55435
	db2User        = "gitea"
	db2Password    = "gitea"
	db2Name        = "gitea"
	apiBaseURL     = "http://localhost:3000/api/v1/admin/reward/action/historic?access_token=de4ff1a5f918b34ff2db7548e43d9eb9a79d94bd"
	pageSize       = 10
	throttleDelay  = 1500 * time.Millisecond
	maxRetries     = 3
	requestTimeout = 10 * time.Second
)

// 数据库连接
var db *sql.DB
var db2 *sql.DB

// 主函数
func main() {
	// 初始化数据库连接
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	if err := initDB2(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db2.Close()
	var lastId int64
	// 分页处理
	page := 0
	for {
		// 查询主数据
		records, err := fetchRecords(lastId)
		if err != nil {
			log.Printf("Failed to fetch records: %v", err)
			break
		}

		if len(records) == 0 {
			log.Println("Processing completed")
			break
		}

		log.Printf("Processing page %d with %d records", page+1, len(records))

		// 处理记录
		id, payload := processRecords(records)
		if lastId < id {
			log.Printf("old last Id = %d   new last id = %d", lastId, id)
			lastId = id

		}
		if len(payload) == 0 {
			log.Println("No valid data in this page")
			page++
			continue
		}

		// 发送API请求
		if err := sendAPIRequest(payload); err != nil {
			log.Printf("Failed to send API request for page %d: %v", page+1, err)
		} else {
			log.Printf("Successfully processed page %d", page+1)
		}

		// 压力控制
		time.Sleep(throttleDelay)
		page++
	}
}

// 初始化数据库连接
func initDB() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.PingContext(ctx)
}

// 初始化数据库连接
func initDB2() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db2Host, db2Port, db2User, db2Password, db2Name)

	var err error
	db2, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db2.PingContext(ctx)
}

// 分页查询主数据
func fetchRecords(lastId int64) ([]map[string]interface{}, error) {
	query := `
		SELECT * FROM reward_operate_record
		WHERE (source_content = '' OR source_content IS NULL) and id > $1 and source_type = 'ACCOMPLISH_TASK'
		ORDER BY id
		LIMIT $2
	`
	log.Println("id = %d  pageSize= %d", lastId, pageSize)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := db2.QueryContext(ctx, query, lastId, pageSize)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	var records []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		record := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				record[col] = string(b)
			} else {
				record[col] = val
			}
		}
		records = append(records, record)
	}

	return records, nil
}

// 处理记录
func processRecords(records []map[string]interface{}) (int64, []map[string]interface{}) {
	var payload []map[string]interface{}
	var lastId int64
	for _, record := range records {
		log.Println("record=%+v", record)
		sourceType, ok1 := record["source_type"].(string)
		sourceID, ok2 := record["source_id"].(string)
		id, ok3 := record["id"].(int64)
		if !ok1 || !ok2 || !ok3 {
			log.Println("Invalid record: missing source_type or source_id")
			continue
		}
		log.Println("id = %d ,lastId=%d", id, lastId)
		if lastId < id {
			lastId = id
		}

		relatedData, err := fetchRelatedData(sourceType, sourceID)
		if err != nil {
			log.Printf("Failed to fetch related data: %v", err)
			continue
		}

		if relatedData == nil {
			continue
		}

		content, err := json.Marshal(relatedData)
		if err != nil {
			log.Printf("Failed to marshal JSON: %v", err)
			continue
		}

		payload = append(payload, map[string]interface{}{
			"type":    sourceType,
			"content": string(content),
		})
	}

	return lastId, payload
}

func fetchRelatedData(sourceType, sourceID string) (map[string]interface{}, error) {
	var query string
	var args []interface{}

	switch sourceType {
	case "ACCOMPLISH_TASK":
		id, err := strconv.ParseInt(sourceID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid source_id: %w", err)
		}
		log.Printf("action id =  %d", id)

		query = "SELECT * FROM action WHERE id = $1"
		args = []interface{}{id}
	case "ADMIN_OPERATE":
		query = "SELECT * FROM reward_admin_log WHERE log_id = $1"
		args = []interface{}{sourceID}
	case "RUN_CLOUDBRAIN_TASK":
		id, err := strconv.ParseInt(sourceID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid source_id: %w", err)
		}
		query = "SELECT *FROM cloudbrain WHERE id = $1"
		args = []interface{}{id}
	default:
		log.Printf("Unknown source_type: %s", sourceType)
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("row error: %w", err)
		}
		log.Printf("No related data found for type=%s, id=%s", sourceType, sourceID)
		return nil, nil
	}

	values := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range values {
		pointers[i] = &values[i]
	}

	if err := rows.Scan(pointers...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	result := make(map[string]interface{})
	for i, col := range columns {
		val := values[i]
		// 检查是否是[]byte类型，并尝试转换为字符串或浮点数
		if b, ok := val.([]byte); ok {
			// 尝试将字节切片转换为浮点数（针对numeric字段）
			if f, err := strconv.ParseFloat(string(b), 64); err == nil {
				result[toCamelCase(col)] = f
			} else {
				// 如果无法转换为浮点数，则直接存储为字符串
				result[toCamelCase(col)] = string(b)
			}
		} else {
			// 其他类型直接存储
			result[toCamelCase(col)] = val
		}
	}

	return result, nil
}

// 发送API请求
func sendAPIRequest(payload []map[string]interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	params := url.Values{}
	params.Add("list", string(jsonData))

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		req, err := http.NewRequestWithContext(ctx, "POST", apiBaseURL+"&"+params.Encode(), nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}

		lastErr = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		if resp.StatusCode == http.StatusTooManyRequests {
			backoff := time.Duration(i+1) * time.Second
			time.Sleep(backoff)
		}
	}

	return lastErr
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i == 0 {
			parts[i] = strings.Title(parts[i]) // 首字母大写
		} else {
			parts[i] = strings.Title(parts[i])
		}
	}
	return strings.Join(parts, "")
}
