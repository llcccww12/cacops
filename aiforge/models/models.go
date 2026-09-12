// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"code.gitea.io/gitea/modules/setting"

	// Needed for the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
	"xorm.io/xorm/schemas"

	// Needed for the Postgresql driver
	_ "github.com/lib/pq"

	// Needed for the MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

// Engine represents a xorm engine or session.
type Engine interface {
	Table(tableNameOrBean interface{}) *xorm.Session
	Count(...interface{}) (int64, error)
	Decr(column string, arg ...interface{}) *xorm.Session
	Delete(interface{}) (int64, error)
	Exec(...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Incr(column string, arg ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *xorm.Session
	SQL(interface{}, ...interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
	Asc(colNames ...string) *xorm.Session
	Limit(limit int, start ...int) *xorm.Session
	SumInt(bean interface{}, columnName string) (res int64, err error)
}

const (
	// When queries are broken down in parts because of the number
	// of parameters, attempt to break by this amount
	maxQueryParameters = 300
)

var (
	x      *xorm.Engine
	tables []interface{}

	xStatistic      *xorm.Engine
	tablesStatistic []interface{}

	// HasEngine specifies if we have a xorm.Engine
	HasEngine bool
)

func init() {
	tables = append(tables,
		new(User),
		new(PublicKey),
		new(AccessToken),
		new(Repository),
		new(DeployKey),
		new(Collaboration),
		new(Access),
		new(Upload),
		new(Watch),
		new(Star),
		new(Follow),
		new(Action),
		new(Issue),
		new(PullRequest),
		new(Comment),
		new(Attachment),
		new(Label),
		new(IssueLabel),
		new(Milestone),
		new(Mirror),
		new(Release),
		new(LoginSource),
		new(Webhook),
		new(HookTask),
		new(Team),
		new(OrgUser),
		new(TeamUser),
		new(TeamRepo),
		new(Notice),
		new(EmailAddress),
		new(Notification),
		new(IssueUser),
		new(LFSMetaObject),
		new(TwoFactor),
		new(GPGKey),
		new(GPGKeyImport),
		new(RepoUnit),
		new(RepoRedirect),
		new(ExternalLoginUser),
		new(ProtectedBranch),
		new(UserOpenID),
		new(IssueWatch),
		new(CommitStatus),
		new(Stopwatch),
		new(TrackedTime),
		new(DeletedBranch),
		new(RepoIndexerStatus),
		new(IssueDependency),
		new(LFSLock),
		new(Reaction),
		new(IssueAssignees),
		new(U2FRegistration),
		new(TeamUnit),
		new(Review),
		new(OAuth2Application),
		new(OAuth2AuthorizationCode),
		new(OAuth2Grant),
		new(Task),
		new(LanguageStat),
		new(EmailHash),
		new(Dataset),
		new(DatasetStar),
		new(Cloudbrain),
		new(Image),
		new(ImageStar),
		new(ImageTopic),
		new(ImageTopicRelation),
		new(FileChunk),
		new(ModelFileChunk),
		new(BlockChain),
		new(RecommendOrg),
		new(AiModelManage),
		new(OfficialTag),
		new(OfficialTagRepos),
		new(WechatBindLog),
		new(OrgStatistic),
		new(SearchRecord),
		new(TaskConfig),
		new(TaskAccomplishLog),
		new(RewardOperateRecord),
		new(LimitConfig),
		new(RewardPeriodicTask),
		new(PointAccountLog),
		new(PointAccount),
		new(RewardAdminLog),
		new(AiModelConvert),
		new(ResourceQueue),
		new(ResourceSpecification),
		new(ResourceScene),
		new(ResourceSceneSpec),
		new(AdminOperateLog),
		new(CloudbrainSpec),
		new(CloudbrainTemp),
		new(DatasetReference),
		new(ScheduleRecord),
		new(BadgeCategory),
		new(Badge),
		new(BadgeUser),
		new(BadgeUserLog),
		new(TechConvergeBaseInfo),
		new(RepoConvergeInfo),
		new(UserRole),
		new(AiModelCollect),
		new(AiModelFile),
		new(ModelMigrateRecord),
		new(IPLocation),
		new(ModelartsDeploy),
		new(ModelartsDeployQueue),
		new(CloudbrainConfig),
		new(ResourceExclusivePool),
		new(CardRequest),
		new(CardRequestSpec),
		new(HfModelFile),
		new(HfModelOperation),
		new(ResourceAICenter),

		new(Agent),
		new(UserAgent),
		new(UserAgentRecord),
		new(AgentExecutionRecord),

		new(AiforgeRole),
		new(AiforgeUserRole),
		new(AiforgeOperation),

		new(OfficialModel),
		new(OfficialDataset),
		new(DatasetCollection),
		new(DatasetRegistry),
		new(SubjectAccess),
		new(SubjectCollaboration),
		new(UploadChunk),
		new(TeamSubject),
		new(OfficialDatasetRegistry),
		new(OldDatasetProcessRecord),
		new(StorageDeleteFailedDataset),
		new(AITaskTemplate),
		new(TemplateDatasets),
		new(TemplateModels),
		new(AITaskTemplateCollection),
		new(OAuth2UserScope),
		new(OAuth2GrantEvent),
		new(OAuth2ApplicationScope),
		new(RepositoryPinnedRecord),
		new(ResourceUpdateBatch),
		new(ResourceUpdateBatchItem),
	)

	tablesStatistic = append(tablesStatistic,
		new(RepoStatistic),
		new(SummaryStatistic),
		new(UserBusinessAnalysis),
		new(UserBusinessAnalysisAll),
		new(UserBusinessAnalysisCurrentYear),
		new(UserBusinessAnalysisLast30Day),
		new(UserBusinessAnalysisLastMonth),
		new(UserBusinessAnalysisCurrentMonth),
		new(UserBusinessAnalysisCurrentWeek),
		new(UserBusinessAnalysisYesterday),
		new(UserBusinessAnalysisLastWeek),
		new(UserLoginLog),
		new(UserLoginActionLog),
		new(UserOtherInfo),
		new(UserMetrics),
		new(UserAnalysisPara),
		new(Invitation),
		new(CloudbrainDurationStatistic),
		new(CloudbrainTaskNumStatistic),
		new(DateCloudbrainNumJson),
		new(QueueTaskDurationStatistic),
		new(QueueTaskNumStatistic),
		new(UserSummaryCurrentYear),
		new(ModelApp),
		new(LlmChat),
		new(LlmChatVisit),
		new(XPUInfoBase),
		new(XPUInfoStatistic),
	)

	gonicNames := []string{"SSL", "UID"}
	for _, name := range gonicNames {
		names.LintGonicMapper[name] = true
	}
}

func getEngine(database *setting.DBInfo) (*xorm.Engine, error) {
	connStr, err := setting.DBConnStr(database)
	if err != nil {
		return nil, err
	}

	engine, err := xorm.NewEngine(setting.Database.Type, connStr)
	if err != nil {
		return nil, err
	}
	if setting.Database.Type == "mysql" {
		engine.Dialect().SetParams(map[string]string{"rowFormat": "DYNAMIC"})
	}
	engine.SetSchema(setting.Database.Schema)

	return engine, nil
}

// NewTestEngine sets a new test xorm.Engine
func NewTestEngine(x *xorm.Engine) (err error) {
	x, err = getEngine(setting.Database)
	if err != nil {
		return fmt.Errorf("Connect to database: %v", err)
	}

	x.SetMapper(names.GonicMapper{})
	x.SetLogger(NewXORMLogger(!setting.ProdMode))
	x.ShowSQL(!setting.ProdMode)
	return x.StoreEngine("InnoDB").Sync2(tables...)
}

// setEngine sets the xorm.Engine
func setEngine(engine *xorm.Engine, table []interface{}, database *setting.DBInfo) (err error) {
	engine.SetMapper(names.GonicMapper{})
	// WARNING: for serv command, MUST remove the output to os.stdout,
	// so use log file to instead print to stdout.
	engine.SetLogger(NewXORMLogger(database.LogSQL))
	engine.ShowSQL(database.LogSQL)
	engine.SetMaxOpenConns(database.MaxOpenConns)
	engine.SetMaxIdleConns(database.MaxIdleConns)
	engine.SetConnMaxLifetime(database.ConnMaxLifetime)
	engine.Sync2(table...)

	return nil
}

func SetEngine() (err error) {
	x, err = getEngine(setting.Database)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	if err = setEngine(x, tables, setting.Database); err != nil {
		return err
	}

	xStatistic, err = getEngine(setting.DatabaseStatistic)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	if err = setEngine(xStatistic, tablesStatistic, setting.DatabaseStatistic); err != nil {
		return err
	}

	return nil
}

func NewEngine(ctx context.Context, migrateFunc func(*xorm.Engine) error) (err error) {
	x, err = getEngine(setting.Database)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	if err = newEngine(ctx, migrateFunc, x, tables, setting.Database); err != nil {
		return fmt.Errorf("newEngine failed: %v", err)
	}
	MigrateCustom(x)
	xStatistic, err = getEngine(setting.DatabaseStatistic)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	if err = newEngine(ctx, migrateFunc, xStatistic, tablesStatistic, setting.DatabaseStatistic); err != nil {
		return fmt.Errorf("newEngine statistic failed: %v", err)
	}
	MigrateCustomStatic(x, xStatistic)

	HasEngine = true

	return nil
}

// newEngine initializes a new xorm.Engine
func newEngine(ctx context.Context, migrateFunc func(*xorm.Engine) error, engine *xorm.Engine, table []interface{}, database *setting.DBInfo) (err error) {
	if err = setEngine(engine, table, database); err != nil {
		return err
	}

	engine.SetDefaultContext(ctx)

	if err = engine.Ping(); err != nil {
		return err
	}

	if err = migrateFunc(engine); err != nil {
		return fmt.Errorf("migrate: %v", err)
	}

	if err = engine.StoreEngine("InnoDB").Sync2(table...); err != nil {
		return fmt.Errorf("sync database struct error: %v", err)
	}

	return nil
}

// Statistic contains the database statistics
type Statistic struct {
	Counter struct {
		User, Org, PublicKey,
		Repo, Watch, Star, Action, Access,
		Issue, Comment, Oauth, Follow,
		Mirror, Release, LoginSource, Webhook,
		Milestone, Label, HookTask,
		Team, UpdateTask, Attachment int64
	}
}

// GetStatistic returns the database statistics
func GetStatistic() (stats Statistic) {
	stats.Counter.User = CountUsers()
	stats.Counter.Org = CountOrganizations()
	stats.Counter.PublicKey, _ = x.Count(new(PublicKey))
	stats.Counter.Repo = CountRepositories(true)
	stats.Counter.Watch, _ = x.Count(new(Watch))
	stats.Counter.Star, _ = x.Count(new(Star))
	//stats.Counter.Action, _ = x.Count(new(Action))
	stats.Counter.Access, _ = x.Count(new(Access))
	stats.Counter.Issue, _ = x.Count(new(Issue))
	stats.Counter.Comment, _ = x.Count(new(Comment))
	stats.Counter.Oauth = 0
	stats.Counter.Follow, _ = x.Count(new(Follow))
	stats.Counter.Mirror, _ = x.Count(new(Mirror))
	stats.Counter.Release, _ = x.Count(new(Release))
	stats.Counter.LoginSource = CountLoginSources()
	stats.Counter.Webhook, _ = x.Count(new(Webhook))
	stats.Counter.Milestone, _ = x.Count(new(Milestone))
	stats.Counter.Label, _ = x.Count(new(Label))
	stats.Counter.HookTask, _ = x.Count(new(HookTask))
	stats.Counter.Team, _ = x.Count(new(Team))
	stats.Counter.Attachment, _ = x.Count(new(Attachment))
	return
}

// Ping tests if database is alive
func Ping() error {
	if x != nil {
		return x.Ping()
	}

	if xStatistic != nil {
		return xStatistic.Ping()
	}

	return errors.New("database not configured")
}

// DumpDatabase dumps all data from database according the special database SQL syntax to file system.
func DumpDatabase(filePath string, dbType string) error {
	var tbs []*schemas.Table
	for _, t := range tables {
		t, err := x.TableInfo(t)
		if err != nil {
			return err
		}
		tbs = append(tbs, t)
	}
	if len(dbType) > 0 {
		return x.DumpTablesToFile(tbs, filePath, schemas.DBType(dbType))
	}
	return x.DumpTablesToFile(tbs, filePath)
}

// MaxBatchInsertSize returns the table's max batch insert size
func MaxBatchInsertSize(bean interface{}) int {
	t, err := x.TableInfo(bean)
	if err != nil {
		return 50
	}
	return 999 / len(t.ColumnsSeq())
}

// Count returns records number according struct's fields as database query conditions
func Count(bean interface{}) (int64, error) {
	return x.Count(bean)
}

// IsTableNotEmpty returns true if table has at least one record
func IsTableNotEmpty(tableName string) (bool, error) {
	return x.Table(tableName).Exist()
}

// DeleteAllRecords will delete all the records of this table
func DeleteAllRecords(tableName string) error {
	_, err := x.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	return err
}

// GetMaxID will return max id of the table
func GetMaxID(beanOrTableName interface{}) (maxID int64, err error) {
	_, err = x.Select("MAX(id)").Table(beanOrTableName).Get(&maxID)
	return
}

// FindByMaxID filled results as the condition from database
func FindByMaxID(maxID int64, limit int, results interface{}) error {
	return x.Where("id <= ?", maxID).
		OrderBy("id DESC").
		Limit(limit).
		Find(results)
}
