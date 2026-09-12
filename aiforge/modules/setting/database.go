// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	// SupportedDatabases includes all supported databases type
	SupportedDatabases = []string{"MySQL", "PostgreSQL", "MSSQL"}
	dbTypes            = map[string]string{"MySQL": "mysql", "PostgreSQL": "postgres", "MSSQL": "mssql", "SQLite3": "sqlite3"}

	// EnableSQLite3 use SQLite3, set by build flag
	EnableSQLite3 bool

	// Database holds the database settings
	Database *DBInfo
	DatabaseStatistic *DBInfo
)

type DBInfo struct {
	Type              string
	Host              string
	Name              string
	User              string
	Passwd            string
	Schema            string
	SSLMode           string
	Path              string
	LogSQL            bool
	Charset           string
	Timeout           int // seconds
	UseSQLite3        bool
	UseMySQL          bool
	UseMSSQL          bool
	UsePostgreSQL     bool
	DBConnectRetries  int
	DBConnectBackoff  time.Duration
	MaxIdleConns      int
	MaxOpenConns      int
	ConnMaxLifetime   time.Duration
	IterateBufferSize int
}

// GetDBTypeByName returns the dataase type as it defined on XORM according the given name
func GetDBTypeByName(name string) string {
	return dbTypes[name]
}

// initDBConfig loads the database settings
func initDBConfig(section string, database *DBInfo) {
	sec := Cfg.Section(section)
	database.Type = sec.Key("DB_TYPE").String()
	switch database.Type {
	case "sqlite3":
		database.UseSQLite3 = true
	case "mysql":
		database.UseMySQL = true
	case "postgres":
		database.UsePostgreSQL = true
	case "mssql":
		database.UseMSSQL = true
	}
	database.Host = sec.Key("HOST").String()
	database.Name = sec.Key("NAME").String()
	database.User = sec.Key("USER").String()
	if len(database.Passwd) == 0 {
		database.Passwd = sec.Key("PASSWD").String()
	}
	database.Schema = sec.Key("SCHEMA").String()
	database.SSLMode = sec.Key("SSL_MODE").MustString("disable")
	database.Charset = sec.Key("CHARSET").In("utf8", []string{"utf8", "utf8mb4"})
	database.Path = sec.Key("PATH").MustString(filepath.Join(AppDataPath, "gitea.db"))
	database.Timeout = sec.Key("SQLITE_TIMEOUT").MustInt(500)
	database.MaxIdleConns = sec.Key("MAX_IDLE_CONNS").MustInt(2)
	if database.UseMySQL {
		database.ConnMaxLifetime = sec.Key("CONN_MAX_LIFE_TIME").MustDuration(3 * time.Second)
	} else {
		database.ConnMaxLifetime = sec.Key("CONN_MAX_LIFE_TIME").MustDuration(0)
	}
	database.MaxOpenConns = sec.Key("MAX_OPEN_CONNS").MustInt(0)

	database.IterateBufferSize = sec.Key("ITERATE_BUFFER_SIZE").MustInt(50)
	database.LogSQL = sec.Key("LOG_SQL").MustBool(true)
	database.DBConnectRetries = sec.Key("DB_RETRIES").MustInt(10)
	database.DBConnectBackoff = sec.Key("DB_RETRY_BACKOFF").MustDuration(3 * time.Second)
}

func InitDBConfig() {
	Database = new(DBInfo)
	DatabaseStatistic = new(DBInfo)
	initDBConfig("database", Database)
	initDBConfig("database_statistic", DatabaseStatistic)
}

// DBConnStr returns database connection string
func DBConnStr(database *DBInfo) (string, error) {
	connStr := ""
	var Param = "?"
	if strings.Contains(database.Name, Param) {
		Param = "&"
	}
	switch database.Type {
	case "mysql":
		connType := "tcp"
		if database.Host[0] == '/' { // looks like a unix socket
			connType = "unix"
		}
		tls := database.SSLMode
		if tls == "disable" { // allow (Postgres-inspired) default value to work in MySQL
			tls = "false"
		}
		connStr = fmt.Sprintf("%s:%s@%s(%s)/%s%scharset=%s&parseTime=true&tls=%s",
			database.User, database.Passwd, connType, database.Host, database.Name, Param, database.Charset, tls)
	case "postgres":
		connStr = getPostgreSQLConnectionString(database.Host, database.User, database.Passwd, database.Name, Param, database.SSLMode)
	case "mssql":
		host, port := ParseMSSQLHostPort(database.Host)
		connStr = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;", host, port, database.Name, database.User, database.Passwd)
	case "sqlite3":
		if !EnableSQLite3 {
			return "", errors.New("this binary version does not build support for SQLite3")
		}
		if err := os.MkdirAll(path.Dir(database.Path), os.ModePerm); err != nil {
			return "", fmt.Errorf("Failed to create directories: %v", err)
		}
		connStr = fmt.Sprintf("file:%s?cache=shared&mode=rwc&_busy_timeout=%d&_txlock=immediate", database.Path, database.Timeout)
	default:
		return "", fmt.Errorf("Unknown database type: %s", database.Type)
	}

	return connStr, nil
}

// parsePostgreSQLHostPort parses given input in various forms defined in
// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
// and returns proper host and port number.
func parsePostgreSQLHostPort(info string) (string, string) {
	host, port := "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func getPostgreSQLConnectionString(dbHost, dbUser, dbPasswd, dbName, dbParam, dbsslMode string) (connStr string) {
	host, port := parsePostgreSQLHostPort(dbHost)
	if host[0] == '/' { // looks like a unix socket
		connStr = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
			url.PathEscape(dbUser), url.PathEscape(dbPasswd), port, dbName, dbParam, dbsslMode, host)
	} else {
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
			url.PathEscape(dbUser), url.PathEscape(dbPasswd), host, port, dbName, dbParam, dbsslMode)
	}
	return
}

// ParseMSSQLHostPort splits the host into host and port
func ParseMSSQLHostPort(info string) (string, string) {
	host, port := "127.0.0.1", "1433"
	if strings.Contains(info, ":") {
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	} else if strings.Contains(info, ",") {
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}
