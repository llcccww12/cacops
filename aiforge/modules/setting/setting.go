// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/generate"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/user"

	shellquote "github.com/kballard/go-shellquote"
	version "github.com/mcuadros/go-version"
	"github.com/unknwon/cae/zip"
	"github.com/unknwon/com"
	ini "gopkg.in/ini.v1"
	"strk.kbt.io/projects/go/libravatar"
)

// Scheme describes protocol types
type Scheme string

// enumerates all the scheme types
const (
	HTTP       Scheme = "http"
	HTTPS      Scheme = "https"
	FCGI       Scheme = "fcgi"
	FCGIUnix   Scheme = "fcgi+unix"
	UnixSocket Scheme = "unix"
)

// LandingPage describes the default page
type LandingPage string

// enumerates all the landing page types
const (
	LandingPageHome          LandingPage = "/"
	LandingPageExplore       LandingPage = "/explore"
	LandingPageOrganizations LandingPage = "/explore/organizations"
	LandingPageLogin         LandingPage = "/user/login"
)

// enumerates all the types of captchas
const (
	ImageCaptcha = "image"
	ReCaptcha    = "recaptcha"
)

type C2NetSequenceInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	ContentEN string `json:"content_en"`
	Loc       string `json:"loc"`
	Type      string `json:"type"`
}

type C2NetSqInfos struct {
	C2NetSqInfo []*C2NetSequenceInfo `json:"sequence"`
}

func applyDemoAiCenterLocations() {
	defaults := []*C2NetSequenceInfo{
		{Name: "cloudbrain_one", Content: "鹏城云脑一号", ContentEN: "Pengcheng Cloudbrain I", Loc: "114.0579,22.5431", Type: "智算中心"},
		{Name: "cloudbrain_two", Content: "鹏城云脑二号", ContentEN: "Pengcheng Cloudbrain II", Loc: "114.0850,22.5550", Type: "智算中心"},
		{Name: "beida", Content: "北大人工智能集群系统", ContentEN: "Peking University AI Center", Loc: "116.3103,39.9927", Type: "智算中心"},
		{Name: "hefei", Content: "合肥类脑智能开放平台", ContentEN: "Hefei AI Center", Loc: "117.2272,31.8206", Type: "智算中心"},
		{Name: "wuhan", Content: "武汉人工智能计算中心", ContentEN: "Wuhan AI Center", Loc: "114.3055,30.5928", Type: "智算中心"},
		{Name: "xian", Content: "西安未来人工智能计算中心", ContentEN: "Xi'an AI Center", Loc: "108.9398,34.2632", Type: "智算中心"},
		{Name: "pclcci", Content: "鹏城云计算所", ContentEN: "Pengcheng Cloud Computing Institute", Loc: "114.0490,22.5600", Type: "智算中心"},
		{Name: "xuchang", Content: "中原人工智能计算中心", ContentEN: "Zhongyuan AI Center", Loc: "113.8526,34.0357", Type: "智算中心"},
		{Name: "chengdu", Content: "成都人工智能计算中心", ContentEN: "Chengdu AI Center", Loc: "104.0665,30.5723", Type: "智算中心"},
		{Name: "hengqin", Content: "横琴先进智能计算中心", ContentEN: "Hengqin AI Center", Loc: "113.5460,22.1200", Type: "智算中心"},
		{Name: "jinan", Content: "国家超级计算济南中心", ContentEN: "Jinan Supercomputing Center", Loc: "117.1205,36.6512", Type: "超算中心"},
		{Name: "tianjin", Content: "国家超级计算天津中心", ContentEN: "Tianjin Supercomputing Center", Loc: "117.2008,39.0842", Type: "超算中心"},
		{Name: "guiyang", Content: "贵安算力枢纽", ContentEN: "Gui'an Computing Hub", Loc: "106.6302,26.6477", Type: "东数西算"},
		{Name: "ningxia", Content: "宁夏中卫算力枢纽", ContentEN: "Zhongwei Computing Hub", Loc: "105.1967,37.5002", Type: "东数西算"},
	}
	if AiCenterCodeAndNameAndLocMapInfo == nil {
		AiCenterCodeAndNameAndLocMapInfo = make(map[string]*C2NetSequenceInfo)
	}
	for _, item := range defaults {
		existing, ok := AiCenterCodeAndNameAndLocMapInfo[item.Name]
		if !ok || existing == nil {
			copied := *item
			AiCenterCodeAndNameAndLocMapInfo[item.Name] = &copied
			continue
		}
		if existing.Loc == "" {
			existing.Loc = item.Loc
		}
		if existing.Type == "" {
			existing.Type = item.Type
		}
		if existing.Content == "" {
			existing.Content = item.Content
		}
		if existing.ContentEN == "" {
			existing.ContentEN = item.ContentEN
		}
	}
}

type AiCenterInfo struct {
	CenterID           string `json:"center_id"`
	Name               string `json:"name"`
	Endpoint           string `json:"endpoint"`
	StorageProxyServer string `json:"storage_proxy_server"`
}

type AiCenterInfos struct {
	Info []*AiCenterInfo `json:"infos"`
}

type StFlavorInfos struct {
	FlavorInfo []*FlavorInfo `json:"flavor_info"`
}

type FlavorInfo struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type StImageInfosModelArts struct {
	ImageInfo []*ImageInfoModelArts `json:"image_info"`
}

type ImageInfoModelArts struct {
	Id    string `json:"id"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

var (
	// AppVer settings
	AppVer         string
	AppBuiltWith   string
	AppName        string
	MLOPS          bool
	MlopsHost      string
	MlopsToken     string
	MlopsSecret    string
	AppURL         string
	AppSubURL      string
	AppSubURLDepth int // Number of slashes
	AppPath        string
	AppDataPath    string
	AppWorkPath    string

	// Server settings
	Protocol             Scheme
	Domain               string
	HTTPAddr             string
	HTTPPort             string
	LocalURL             string
	RedirectOtherPort    bool
	PortToRedirect       string
	OfflineMode          bool
	CertFile             string
	KeyFile              string
	StaticRootPath       string
	StaticCacheTime      time.Duration
	EnableGzip           bool
	LandingPageURL       LandingPage
	UnixSocketPermission uint32
	EnablePprof          bool
	PprofDataPath        string
	EnableLetsEncrypt    bool
	LetsEncryptTOS       bool
	LetsEncryptDirectory string
	LetsEncryptEmail     string
	GracefulRestartable  bool
	GracefulHammerTime   time.Duration
	StartupTimeout       time.Duration
	StaticURLPrefix      string
	EnableWebSocket      bool
	CustomRealUrlHeader  string

	SSH = struct {
		Disabled                 bool           `ini:"DISABLE_SSH"`
		StartBuiltinServer       bool           `ini:"START_SSH_SERVER"`
		BuiltinServerUser        string         `ini:"BUILTIN_SSH_SERVER_USER"`
		Domain                   string         `ini:"SSH_DOMAIN"`
		Port                     int            `ini:"SSH_PORT"`
		ListenHost               string         `ini:"SSH_LISTEN_HOST"`
		ListenPort               int            `ini:"SSH_LISTEN_PORT"`
		RootPath                 string         `ini:"SSH_ROOT_PATH"`
		ServerCiphers            []string       `ini:"SSH_SERVER_CIPHERS"`
		ServerKeyExchanges       []string       `ini:"SSH_SERVER_KEY_EXCHANGES"`
		ServerMACs               []string       `ini:"SSH_SERVER_MACS"`
		KeyTestPath              string         `ini:"SSH_KEY_TEST_PATH"`
		KeygenPath               string         `ini:"SSH_KEYGEN_PATH"`
		AuthorizedKeysBackup     bool           `ini:"SSH_AUTHORIZED_KEYS_BACKUP"`
		MinimumKeySizeCheck      bool           `ini:"-"`
		MinimumKeySizes          map[string]int `ini:"-"`
		CreateAuthorizedKeysFile bool           `ini:"SSH_CREATE_AUTHORIZED_KEYS_FILE"`
		ExposeAnonymous          bool           `ini:"SSH_EXPOSE_ANONYMOUS"`
	}{
		Disabled:           false,
		StartBuiltinServer: false,
		Domain:             "",
		Port:               22,
		ServerCiphers:      []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128"},
		ServerKeyExchanges: []string{"diffie-hellman-group1-sha1", "diffie-hellman-group14-sha1", "ecdh-sha2-nistp256", "ecdh-sha2-nistp384", "ecdh-sha2-nistp521", "curve25519-sha256@libssh.org"},
		ServerMACs:         []string{"hmac-sha2-256-etm@openssh.com", "hmac-sha2-256", "hmac-sha1", "hmac-sha1-96"},
		KeygenPath:         "ssh-keygen",
		MinimumKeySizes:    map[string]int{"ed25519": 256, "ecdsa": 256, "rsa": 2048, "dsa": 1024},
	}

	LFS struct {
		StartServer     bool          `ini:"LFS_START_SERVER"`
		ContentPath     string        `ini:"LFS_CONTENT_PATH"`
		JWTSecretBase64 string        `ini:"LFS_JWT_SECRET"`
		JWTSecretBytes  []byte        `ini:"-"`
		HTTPAuthExpiry  time.Duration `ini:"LFS_HTTP_AUTH_EXPIRY"`
		MaxFileSize     int64         `ini:"LFS_MAX_FILE_SIZE"`
		LocksPagingNum  int           `ini:"LFS_LOCKS_PAGING_NUM"`
	}

	// Security settings
	InstallLock                        bool
	LoginEncrypt                       bool
	SecretKey                          string
	LogInRememberDays                  int
	CookieUserName                     string
	CookieRememberName                 string
	ReverseProxyAuthUser               string
	ReverseProxyAuthEmail              string
	MinPasswordLength                  int
	ImportLocalPaths                   bool
	DisableGitHooks                    bool
	OnlyAllowPushIfGiteaEnvironmentSet bool
	PasswordComplexity                 []string
	PasswordHashAlgo                   string

	// UI settings
	UI = struct {
		ExplorePagingNum      int
		ContributorPagingNum  int
		IssuePagingNum        int
		DatasetPagingNum      int
		RepoSearchPagingNum   int
		MembersPagingNum      int
		FeedMaxCommitNum      int
		GraphMaxCommitNum     int
		CodeCommentLines      int
		ReactionMaxUserNum    int
		ThemeColorMetaTag     string
		MaxDisplayFileSize    int64
		ShowUserEmail         bool
		DefaultShowFullName   bool
		DefaultTheme          string
		Themes                []string
		Reactions             []string
		ReactionsMap          map[string]bool
		SearchRepoDescription bool

		Notification struct {
			MinTimeout             time.Duration
			TimeoutStep            time.Duration
			MaxTimeout             time.Duration
			EventSourceUpdateTime  time.Duration
			RewardNotifyUpdateTime time.Duration
		} `ini:"ui.notification"`

		Admin struct {
			UserPagingNum   int
			RepoPagingNum   int
			NoticePagingNum int
			OrgPagingNum    int
		} `ini:"ui.admin"`
		User struct {
			RepoPagingNum int
		} `ini:"ui.user"`
		Meta struct {
			Author      string
			Description string
			Keywords    string
		} `ini:"ui.meta"`
	}{
		ExplorePagingNum:     20,
		ContributorPagingNum: 50,
		IssuePagingNum:       10,
		DatasetPagingNum:     5,
		RepoSearchPagingNum:  10,
		MembersPagingNum:     20,
		FeedMaxCommitNum:     5,
		GraphMaxCommitNum:    100,
		CodeCommentLines:     4,
		ReactionMaxUserNum:   10,
		ThemeColorMetaTag:    `#6cc644`,
		MaxDisplayFileSize:   8388608,
		DefaultTheme:         `gitea`,
		Themes:               []string{`gitea`, `arc-green`},
		Reactions:            []string{`+1`, `-1`, `laugh`, `hooray`, `confused`, `heart`, `rocket`, `eyes`},
		Notification: struct {
			MinTimeout             time.Duration
			TimeoutStep            time.Duration
			MaxTimeout             time.Duration
			EventSourceUpdateTime  time.Duration
			RewardNotifyUpdateTime time.Duration
		}{
			MinTimeout:             10 * time.Second,
			TimeoutStep:            10 * time.Second,
			MaxTimeout:             60 * time.Second,
			EventSourceUpdateTime:  10 * time.Second,
			RewardNotifyUpdateTime: 2 * time.Second,
		},
		Admin: struct {
			UserPagingNum   int
			RepoPagingNum   int
			NoticePagingNum int
			OrgPagingNum    int
		}{
			UserPagingNum:   50,
			RepoPagingNum:   50,
			NoticePagingNum: 25,
			OrgPagingNum:    50,
		},
		User: struct {
			RepoPagingNum int
		}{
			RepoPagingNum: 15,
		},
		Meta: struct {
			Author      string
			Description string
			Keywords    string
		}{
			Author:      "Gitea - Git with a cup of tea",
			Description: "Gitea (Git with a cup of tea) is a painless self-hosted Git service written in Go",
			Keywords:    "go,git,self-hosted,gitea",
		},
	}

	// Markdown settings
	Markdown = struct {
		EnableHardLineBreak bool
		CustomURLSchemes    []string `ini:"CUSTOM_URL_SCHEMES"`
		FileExtensions      []string
		EnableToc           bool
	}{
		EnableHardLineBreak: true,
		FileExtensions:      strings.Split(".md,.markdown,.mdown,.mkd", ","),
		EnableToc:           true,
	}

	// Admin settings
	Admin struct {
		DisableRegularOrgCreation bool
		DefaultEmailNotification  string
	}

	// Picture settings
	AvatarUploadPath              string
	AvatarMaxWidth                int
	AvatarMaxHeight               int
	AvatarStaticCacheTime         time.Duration
	GravatarSource                string
	GravatarSourceURL             *url.URL
	DisableGravatar               bool
	EnableFederatedAvatar         bool
	LibravatarService             *libravatar.Libravatar
	AvatarMaxFileSize             int64
	RepositoryAvatarUploadPath    string
	RepositoryAvatarFallback      string
	RepositoryAvatarFallbackImage string

	// Log settings
	LogLevel           string
	StacktraceLogLevel string
	LogRootPath        string
	LogDescriptions    = make(map[string]*LogDescription)
	RedirectMacaronLog bool
	DisableRouterLog   bool
	RouterLogLevel     log.Level
	RouterLogMode      string
	EnableAccessLog    bool
	AccessLogTemplate  string
	EnableXORMLog      bool

	// Attachment settings
	Attachment = struct {
		StoreType string
		Path      string
		Minio     struct {
			Endpoint         string
			InternalEndpoint string
			AccessKeyID      string
			SecretAccessKey  string
			UseSSL           bool
			Bucket           string
			Location         string
			BasePath         string
			RealPath         string
		}
		AllowedTypes    string
		MaxSize         int64
		MaxFiles        int
		Enabled         bool
		SIZE_FOR_PYTHON int
	}{
		StoreType: "local",
		Minio: struct {
			Endpoint         string
			InternalEndpoint string
			AccessKeyID      string
			SecretAccessKey  string
			UseSSL           bool
			Bucket           string
			Location         string
			BasePath         string
			RealPath         string
		}{},
		AllowedTypes:    "image/jpeg,image/png,application/zip,application/gzip",
		MaxSize:         4,
		MaxFiles:        5,
		Enabled:         true,
		SIZE_FOR_PYTHON: 500,
	}

	// Time settings
	TimeFormat string
	// UILocation is the location on the UI, so that we can display the time on UI.
	DefaultUILocation = time.Local

	CSRFCookieName     = "_csrf"
	CSRFCookieHTTPOnly = true

	// Mirror settings
	Mirror struct {
		DefaultInterval time.Duration
		MinInterval     time.Duration
	}

	// API settings
	API = struct {
		EnableSwagger          bool
		SwaggerURL             string
		MaxResponseItems       int
		DefaultPagingNum       int
		DefaultGitTreesPerPage int
		DefaultMaxBlobSize     int64
	}{
		EnableSwagger:          false,
		SwaggerURL:             "",
		MaxResponseItems:       50,
		DefaultPagingNum:       30,
		DefaultGitTreesPerPage: 1000,
		DefaultMaxBlobSize:     10485760,
	}

	OAuth2 = struct {
		Enable                     bool
		AccessTokenExpirationTime  int64
		RefreshTokenExpirationTime int64
		InvalidateRefreshTokens    bool
		JWTSecretBytes             []byte `ini:"-"`
		JWTSecretBase64            string `ini:"JWT_SECRET"`
		MaxTokenLength             int
	}{
		Enable:                     true,
		AccessTokenExpirationTime:  3600,
		RefreshTokenExpirationTime: 730,
		InvalidateRefreshTokens:    false,
		MaxTokenLength:             math.MaxInt16,
	}

	U2F = struct {
		AppID         string
		TrustedFacets []string
	}{}

	// Metrics settings
	Metrics = struct {
		Enabled bool
		Token   string
	}{
		Enabled: false,
		Token:   "",
	}

	// I18n settings
	Langs     []string
	Names     []string
	dateLangs map[string]string

	// Highlight settings are loaded in modules/template/highlight.go

	// Other settings
	ShowFooterBranding         bool
	ShowFooterVersion          bool
	ShowFooterTemplateLoadTime bool

	// Global setting objects
	Cfg           *ini.File
	CustomPath    string // Custom directory path
	CustomConf    string
	CustomPID     string
	ProdMode      bool
	RunUser       string
	IsWindows     bool
	HasRobotsTxt  bool
	InternalToken string // internal access token

	// UILocation is the location on the UI, so that we can display the time on UI.
	// Currently only show the default time.Local, it could be added to app.ini after UI is ready
	UILocation = time.Local

	//Machinery config
	Broker        string
	DefaultQueue  string
	ResultBackend string

	//decompress config
	DecompressAddress string
	AuthUser          string
	AuthPassword      string

	//home page
	RecommentRepoAddr string
	ESSearchURL       string
	INDEXPOSTFIX      string
	//notice config
	UserNameOfNoticeRepo string
	RepoNameOfNoticeRepo string
	RefNameOfNoticeRepo  string
	TreePathOfNoticeRepo string
	CacheTimeOutSecond   int
	SDK_NOTICE           string

	//dync config
	DyncConfigRepoOwner  string
	DyncConfigRepoName   string
	DyncConfigBranch     string
	DyncConfigDefaultTTL time.Duration

	// 联通付费算力
	PayConputilityConfigRepoOwner string
	PayConputilityConfigRepoName  string
	PayConputilitycConfigBranch   string

	// 并行科技付费算力
	PeraPayConputilityConfigRepoOwner string
	PeraPayConputilityConfigRepoName  string
	PeraPayConputilitycConfigBranch   string

	// 中国移动云科技付费算力
	ChinaMobilePayConputilityConfigRepoOwner string
	ChinaMobilePayConputilityConfigRepoName  string
	ChinaMobilePayConputilitycConfigBranch   string

	// 算力合作伙伴
	ConputilityPartnerConfigRepoOwner string
	ConputilityPartnerConfigRepoName  string
	ConputilityPartnerConfigBranch    string

	//labelsystem config
	LabelTaskName           string
	LabelDatasetDeleteQueue string
	DecompressOBSTaskName   string

	//cloudbrain config

	CBAuthUser             string
	CBAuthPassword         string
	RestServerHost         string
	JobPath                string
	CBCodePathPrefix       string
	JobType                string
	GpuTypes               string
	SpecialPools           string
	DebugServerHost        string
	ResourceSpecs          string
	MaxDuration            int64
	TrainGpuTypes          string
	TrainResourceSpecs     string
	InferenceGpuTypes      string
	InferenceResourceSpecs string
	MaxModelSize           float64
	MaxModelSizeForWhole   float64
	MaxDatasetNum          int
	CullIdleTimeout        string
	CullInterval           string
	DebugAttachSize        int

	//benchmark config
	IsBenchmarkEnabled     bool
	BenchmarkOwner         string
	BenchmarkName          string
	BenchmarkServerHost    string
	BenchmarkCategory      string
	BenchmarkTypes         string
	BenchmarkGpuTypes      string
	BenchmarkResourceSpecs string
	BenchmarkMaxDuration   int64

	//snn4imagenet config
	IsSnn4imagenetEnabled     bool
	Snn4imagenetOwner         string
	Snn4imagenetName          string
	Snn4imagenetServerHost    string
	ModelBenchmarkMaxDuration int64

	//snn4imagenet config
	IsBrainScoreEnabled  bool
	BrainScoreOwner      string
	BrainScoreName       string
	BrainScoreServerHost string
	BrainScoreRegion     []string

	IsSnn4EcosetEnabled  bool
	Snn4EcosetOwner      string
	Snn4EcosetName       string
	Snn4EcosetServerHost string
	Snn4AttachmentName   string

	IsSim2BrainSnnEnabled    bool
	Sim2BrainSnnOwner        string
	Sim2BrainSnnName         string
	Sim2BrainSnnServerHost   string
	Sim2BrainSnnDatasetTypes []string
	Sim2BrainSnnAttachNames  []string

	//blockchain config
	BlockChainHost  string
	CommitValidDate string

	//obs config
	Endpoint          string
	AccessKeyID       string
	SecretAccessKey   string
	Bucket            string
	Location          string
	BasePath          string
	OutPutPath        string
	TrainJobModelPath string
	CodePathPrefix    string
	UserBasePath      string
	//storage default config
	StorageDefaultType               string
	DatasetSizeChangedQueueBatchSize int
	AimodelSizeChangedQueueBatchSize int

	//urchin config
	UrchinAddress          string
	UrchinReqTimeoutSecond int64
	UrchinMaxConnection    int

	//modelarts config
	ModelArtsHost     string
	IamHost           string
	ProjectID         string
	ProjectName       string
	ModelArtsUsername string
	ModelArtsPassword string
	ModelArtsDomain   string
	AllowedOrg        string
	ProfileID         string
	PoolInfos         string
	FlavorInfos       string
	DebugHost         string
	ImageInfos        string
	Capacity          int
	MaxTempQueryTimes int
	StFlavorInfo      *StFlavorInfos
	StImageInfos      *StImageInfosModelArts
	//train-job
	ResourcePools         string
	Engines               string
	EngineVersions        string
	TrainJobFLAVORINFOS   string
	ModelArtsSpecialPools string
	ModelArtsMultiNode    string
	ModelArtsShareAddr    string
	ModelArtsMountPath    string
	ModelArtsNasType      string
	//kanban
	IsCloudbrainTimingEnabled        bool
	CloudbrainNumTargetTime          string
	CloudbrainDownloadAllConcurrency int

	// modelarts-cd config
	ModelartsCD = struct {
		Enabled     bool
		EndPoint    string
		ProjectID   string
		AccessKey   string
		SecretKey   string
		ImageInfos  string
		FlavorInfos string
	}{}

	//grampus config
	Grampus = struct {
		Env                           string
		Host                          string
		UserName                      string
		Password                      string
		SpecialPools                  string
		C2NetSequence                 string
		SyncScriptProject             string
		LocalCenterID                 string
		GPULocalCenterID              string
		AiCenterInfo                  string
		AiCenterCodeAndNameInfo       string
		AiCenterCodeAndNameAndLocInfo string
		UsageRateBeginTime            string
		GPUImageCommonName            string
		MultiNode                     string
		MMLSparkMaxTime               int64
		NoteBookDomainURL             string
		NoteBookLocalURL              string
	}{}

	ClearStrategy = struct {
		Enabled        bool
		ResultSaveDays int
		BatchSize      int
		DebugJobSize   int
		TrashSaveDays  int
		Cron           string
		RunAtStart     bool
	}{}
	NotebookStrategy = struct {
		ClearEnabled     bool
		ClearBatchSize   int
		MaxNumberPerUser int
		Cron             string
		RunAtStart       bool
	}{}
	CardStatistic = struct {
		Enabled    bool
		Cron       string
		RunAtStart bool
	}{}

	C2NetInfos                       *C2NetSqInfos
	CenterInfos                      *AiCenterInfos
	AiCenterCodeAndNameMapInfo       map[string]*C2NetSequenceInfo
	C2NetLocInfos                    *C2NetSqInfos
	AiCenterCodeAndNameAndLocMapInfo map[string]*C2NetSequenceInfo

	//elk config
	ElkUrl              string
	ElkUser             string
	ElkPassword         string
	Index               string
	TimeField           string
	ElkTimeFormat       string
	PROJECT_LIMIT_PAGES []string

	//wechat config
	WechatApiHost             string
	WechatApiTimeoutSeconds   int
	WechatAppId               string
	WechatAppSecret           string
	WechatQRCodeExpireSeconds int
	WechatAuthSwitch          bool
	WechatSignUpRequired      bool

	//point config
	CloudBrainPaySwitch     bool
	CloudBrainPayDelay      time.Duration
	CloudBrainPayInterval   time.Duration
	DeductTaskRange         time.Duration
	DeductTaskRangeForFirst time.Duration
	DeductTaskMinTimestamp  int64
	UseDynamicData          bool

	//model-migrate config
	UseLocalMinioMigrate bool

	//badge config
	BadgeIconMaxFileSize int64
	BadgeIconMaxWidth    int
	BadgeIconMaxHeight   int
	BadgeIconDefaultSize uint
	IconUploadPath       string

	//wechat auto reply config
	UserNameOfWechatReply  string
	RepoNameOfWechatReply  string
	RefNameOfWechatReply   string
	TreePathOfAutoMsgReply string
	TreePathOfSubscribe    string

	//wechat template msg config
	CloudbrainStartedTemplateId         string
	CloudbrainStartedNotifyList         []string
	CloudbrainStartedTitle              string
	CloudbrainStartedRemark             string
	CloudbrainStoppedTemplateId         string
	CloudbrainStoppedNotifyList         []string
	CloudbrainStoppedTitle              string
	CloudbrainStoppedRemark             string
	CloudbrainComingStopTemplateId      string
	CloudbrainComingStopTitle           string
	CloudbrainComingStopChargeLink      string
	CloudbrainComingStopRemark          string
	CloudbrainComingStopSendFlag        bool
	CloudbrainLongRunningNotifyList     []string
	CloudbrainLongRunningTemplateId     string
	CloudbrainLongRunningNotifyInterval time.Duration
	BindReplySuccess                    string
	BindReplyWechatAccountUsed          string
	BindReplyOpenIAccountUsed           string
	BindReplyFailedDefault              string
	BindReplySignUp                     string
	BindReplySignUpErr4WechatUsed       string
	BindReplyLoginSuccess               string
	BindReplyNoUser                     string

	//repo square config
	IncubationSourceOrgName string
	PaperRepoTopicName      string

	CloudbrainUniquenessLockTime time.Duration
	SyncCloudbrainStatusDuration string

	//ai_task config
	AI_TASK_RANGE                          map[string][]string
	PREPARING_MAX_WAIT_DURATION            time.Duration
	OUTPUT_SHOW_MAX_KEY                    int
	OUTPUT_DOWNLOAD_MAX_KEY                int
	SPECIFICATION_SPECIAL_QUEUE            string
	DEBUG_MODEL_NUM_LIMIT                  int
	DEBUG_MODEL_SIZE_LIMIT_GB              int
	ROLE_MULTI_LIMIT_MAP                   map[string]map[string]int
	AI_TASK_TEMPLATE_FILE_NAME             string
	MAX_MULTI_FINE_TUNE_TASK               int
	MAX_MULTI_SD_FINETUNE_TASK             int
	DEBUG_ADDRESS_CHECK_GLOBAL_TOGGLE      bool
	SD_FINETUNE_MODEL_NUM_LIMIT            int
	SD_FINETUNE_MODEL_SIZE_LIMIT_GB        int
	COMFYUI_EXPERIENCE_MODEL_NUM_LIMIT     int
	COMFYUI_EXPERIENCE_MODEL_SIZE_LIMIT_GB int
	FORCE_OLD_OUTPUT_MIGRATE               bool
	FORCE_OLD_CODE_STORAGE                 bool

	//dataset registry config
	MAX_PREVIEW_FILE_SIZE                   int64
	ENABLE_ORGANIZATION_UPLOAD_DATASET_FILE bool
	HOT_DATASET_LIST_URL                    string
	DATASET_STORAGE_TYPE                    string

	// aimodel
	AIMODEL_STORAGE_TYPE string

	//resource queue
	UseNewSyncQueueAPI bool

	//wenxin url
	BaiduWenXin = struct {
		ModelArtsWenXinURL       string
		BAIDU_API_KEY            string
		BAIDU_SECRET_KEY         string
		BAIDU_TOKEN_URL          string
		PICTURE_BAIDU_REVIEW_URL string
		TEXT_BAIDU_REVIEW_URL    string
		RUN_WORKERS              int
		MODEL_SERVERS            int
	}{}

	//nginx proxy
	PROXYURL string
	RadarMap = struct {
		Impact             float64
		ImpactWatch        float64
		ImpactStar         float64
		ImpactFork         float64
		ImpactCodeDownload float64
		ImpactComments     float64
		ImpactBrowser      float64

		Completeness             float64
		CompletenessIssuesClosed float64
		CompletenessReleases     float64
		CompletenessDevelopAge   float64
		CompletenessDataset      float64
		CompletenessModel        float64
		CompletenessWiki         float64

		Liveness        float64
		LivenessCommit  float64
		LivenessIssue   float64
		LivenessPR      float64
		LivenessRelease float64

		ProjectHealth                   float64
		ProjectHealthIssueCompleteRatio float64
		ProjectHealth0IssueCloseRatio   float64

		TeamHealth                  float64
		TeamHealthContributors      float64
		TeamHealthKeyContributors   float64
		TeamHealthContributorsAdded float64

		Growth             float64
		GrowthCodeLines    float64
		GrowthIssue        float64
		GrowthContributors float64
		GrowthCommit       float64
		GrowthComments     float64
		RecordBeginTime    string
		GrowthBeginTime    string
		IgnoreMirrorRepo   bool
	}{}

	Warn_Notify_Mails []string

	Course = struct {
		OrgName  string
		TeamName string
	}{}

	FileNoteBook = struct {
		ProjectName           string
		ImageGPU              string
		SpecIdGPU             int64
		SpecIdCPU             int64
		ImageIdNPU            string
		SpecIdNPU             int64
		ImageIdNPUCD          string
		SpecIdNPUCD           int64
		ImageCPUDescription   string
		ImageGPUDescription   string
		ImageNPUDescription   string
		ImageNPUCDDescription string
	}{}

	ModelConvert = struct {
		GPU_PYTORCH_IMAGE       string
		GpuQueue                string
		GPU_TENSORFLOW_IMAGE    string
		GPU_PADDLE_IMAGE        string
		GPU_MXNET_IMAGE         string
		NPU_MINDSPORE_16_IMAGE  string
		PytorchOnnxBootFile     string
		PaddleOnnxBootFile      string
		MXnetOnnxBootFile       string
		PytorchTrTBootFile      string
		MindsporeBootFile       string
		TensorFlowNpuBootFile   string
		TensorFlowGpuBootFile   string
		ConvertRepoPath         string
		GPU_Resource_Specs_ID   string
		GPU_AiCenter_Code       string
		GPU_Spec_ID             int64
		GPU_Spec_AccCardType    string
		NPU_FlavorCode          string
		NPU_PoolID              string
		NPU_MINDSPORE_IMAGE_ID  int
		NPU_TENSORFLOW_IMAGE_ID int
	}{}

	ModelSafetyTest = struct {
		NPUBaseDataSetName   string
		NPUBaseDataSetUUID   string
		NPUCombatDataSetName string
		NPUCombatDataSetUUID string

		GPUBaseDataSetName   string
		GPUBaseDataSetUUID   string
		GPUCombatDataSetName string
		GPUCombatDataSetUUID string

		HOST string
		KEY  string
	}{}

	ModelApp = struct {
		DesensitizationUrl string
	}{}

	FLOW_CONTROL = struct {
		ATTACHEMENT_NUM_A_USER_LAST24HOUR int
		ATTACHEMENT_NUM_A_USER_LAST10M    int
		ATTACHEMENT_SIZE_A_USER           int64 //G
		ALL_ATTACHEMENT_NUM_SDK           int
		IGNORE_FLAG                       string
	}{}

	LLM_CHAT_API = struct {
		CHATGLM2_HOST        string
		CHATGLM2_MAX_LENGTH  int
		LLAMA2_HOST          string
		LLAMA2_MAX_LENGTH    int
		COMMON_KB            string
		MAX_FREE_TRIES       int64
		LEGAL_CHECK          bool
		LEGAL_MAX_COUNT      int64
		CHAT_EXPIRED_MINUTES int64
	}{}

	BAIDU_AI = struct {
		API_KEY        string
		SECRET_KEY     string
		URL            string
		LEGAL_TEXT_URL string
	}{}

	ExternalTransfer = struct {
		HfModelOwner   string
		HfDomain       string
		HfToken        string
		MaxModelSizeGb int
		HeartbeatCron  string
		MaxPerUser     int
	}{}

	MODEL_EXPERIENCE = struct {
		MAX_CREATES          int
		PORT_MIN             int
		PORT_MAX             int
		MODELBASE_REPO_OWNER string
	}{}

	//iflytek config
	IFLYTekConfig = struct {
		DataServiceHost       string
		TrainServiceHost      string
		OnlineServiceHost     string
		ModelServiceHost      string
		TOKEN                 string
		DefaultLabelId        string
		OnlineServiceAPIUrl   string
		OnlineServiceAPIToken string
		SystemStructure       string
		ComputeSource         string
		XNum                  string
	}{}

	//deepseek modelarts infer
	DeepSeekModelarts = struct {
		URL_QWEN_1_5B          string
		URL_LLAMA_8B           string
		URL_QWEN_7B            string
		URL_QWEN_14B           string
		URL_QWEN_32B_W8A8      string
		URL_V3_MINDFORMERS     string
		SERVICE_HUB_URL        string
		SERVICE_AUTO_MODEL_URL string
		AccessKey              string
		SecretKey              string
	}{}
)

// DateLang transforms standard language locale name to corresponding value in datetime plugin.
func DateLang(lang string) string {
	name, ok := dateLangs[lang]
	if ok {
		return name
	}
	return "en"
}

func getAppPath() (string, error) {
	var appPath string
	var err error
	if IsWindows && filepath.IsAbs(os.Args[0]) {
		appPath = filepath.Clean(os.Args[0])
	} else {
		appPath, err = exec.LookPath(os.Args[0])
	}

	if err != nil {
		return "", err
	}
	appPath, err = filepath.Abs(appPath)
	if err != nil {
		return "", err
	}
	// Note: we don't use path.Dir here because it does not handle case
	//	which path starts with two "/" in Windows: "//psf/Home/..."
	return strings.Replace(appPath, "\\", "/", -1), err
}

func getWorkPath(appPath string) string {
	workPath := AppWorkPath

	if giteaWorkPath, ok := os.LookupEnv("GITEA_WORK_DIR"); ok {
		workPath = giteaWorkPath
	}
	if len(workPath) == 0 {
		i := strings.LastIndex(appPath, "/")
		if i == -1 {
			workPath = appPath
		} else {
			workPath = appPath[:i]
		}
	}
	return strings.Replace(workPath, "\\", "/", -1)
}

func init() {
	IsWindows = runtime.GOOS == "windows"
	// We can rely on log.CanColorStdout being set properly because modules/log/console_windows.go comes before modules/setting/setting.go lexicographically
	log.NewLogger(0, "console", "console", fmt.Sprintf(`{"level": "trace", "colorize": %t, "stacktraceLevel": "none"}`, log.CanColorStdout))

	var err error
	if AppPath, err = getAppPath(); err != nil {
		log.Fatal("Failed to get app path: %v", err)
	}
	AppWorkPath = getWorkPath(AppPath)
}

func forcePathSeparator(path string) {
	if strings.Contains(path, "\\") {
		log.Fatal("Do not use '\\' or '\\\\' in paths, instead, please use '/' in all places")
	}
}

// IsRunUserMatchCurrentUser returns false if configured run user does not match
// actual user that runs the app. The first return value is the actual user name.
// This check is ignored under Windows since SSH remote login is not the main
// method to login on Windows.
func IsRunUserMatchCurrentUser(runUser string) (string, bool) {
	if IsWindows || SSH.StartBuiltinServer {
		return "", true
	}

	currentUser := user.CurrentUsername()
	return currentUser, runUser == currentUser
}

func createPIDFile(pidPath string) {
	currentPid := os.Getpid()
	if err := os.MkdirAll(filepath.Dir(pidPath), os.ModePerm); err != nil {
		log.Fatal("Failed to create PID folder: %v", err)
	}

	file, err := os.Create(pidPath)
	if err != nil {
		log.Fatal("Failed to create PID file: %v", err)
	}
	defer file.Close()
	if _, err := file.WriteString(strconv.FormatInt(int64(currentPid), 10)); err != nil {
		log.Fatal("Failed to write PID information: %v", err)
	}
}

// CheckLFSVersion will check lfs version, if not satisfied, then disable it.
func CheckLFSVersion() {
	if LFS.StartServer {
		//Disable LFS client hooks if installed for the current OS user
		//Needs at least git v2.1.2

		binVersion, err := git.BinVersion()
		if err != nil {
			log.Fatal("Error retrieving git version: %v", err)
		}

		if !version.Compare(binVersion, "2.1.2", ">=") {
			LFS.StartServer = false
			log.Error("LFS server support needs at least Git v2.1.2")
		} else {
			git.GlobalCommandArgs = append(git.GlobalCommandArgs, "-c", "filter.lfs.required=",
				"-c", "filter.lfs.smudge=", "-c", "filter.lfs.clean=")
		}
	}
}

// SetCustomPathAndConf will set CustomPath and CustomConf with reference to the
// GITEA_CUSTOM environment variable and with provided overrides before stepping
// back to the default
func SetCustomPathAndConf(providedCustom, providedConf, providedWorkPath string) {
	if len(providedWorkPath) != 0 {
		AppWorkPath = filepath.ToSlash(providedWorkPath)
	}
	if giteaCustom, ok := os.LookupEnv("GITEA_CUSTOM"); ok {
		CustomPath = giteaCustom
	}
	if len(providedCustom) != 0 {
		CustomPath = providedCustom
	}
	if len(CustomPath) == 0 {
		CustomPath = path.Join(AppWorkPath, "custom")
	} else if !filepath.IsAbs(CustomPath) {
		CustomPath = path.Join(AppWorkPath, CustomPath)
	}

	if len(providedConf) != 0 {
		CustomConf = providedConf
	}
	if len(CustomConf) == 0 {
		CustomConf = path.Join(CustomPath, "conf/app.ini")
	} else if !filepath.IsAbs(CustomConf) {
		CustomConf = path.Join(CustomPath, CustomConf)
		log.Warn("Using 'custom' directory as relative origin for configuration file: '%s'", CustomConf)
	}
}

// NewContext initializes configuration context.
// NOTE: do not print any log except error.
func NewContext() {
	Cfg = ini.Empty()

	if len(CustomPID) > 0 {
		createPIDFile(CustomPID)
	}

	if com.IsFile(CustomConf) {
		if err := Cfg.Append(CustomConf); err != nil {
			log.Fatal("Failed to load custom conf '%s': %v", CustomConf, err)
		}
	} else {
		log.Warn("Custom config '%s' not found, ignore this if you're running first time", CustomConf)
	}
	Cfg.NameMapper = ini.SnackCase

	homeDir, err := com.HomeDir()
	if err != nil {
		log.Fatal("Failed to get home directory: %v", err)
	}
	homeDir = strings.Replace(homeDir, "\\", "/", -1)

	LogLevel = getLogLevel(Cfg.Section("log"), "LEVEL", "Info")
	StacktraceLogLevel = getStacktraceLogLevel(Cfg.Section("log"), "STACKTRACE_LEVEL", "None")
	LogRootPath = Cfg.Section("log").Key("ROOT_PATH").MustString(path.Join(AppWorkPath, "log"))
	forcePathSeparator(LogRootPath)
	RedirectMacaronLog = Cfg.Section("log").Key("REDIRECT_MACARON_LOG").MustBool(false)
	RouterLogLevel = log.FromString(Cfg.Section("log").Key("ROUTER_LOG_LEVEL").MustString("Info"))

	sec := Cfg.Section("server")
	AppName = Cfg.Section("").Key("APP_NAME").MustString("Gitea: Git with a cup of tea")

	Protocol = HTTP
	switch sec.Key("PROTOCOL").String() {
	case "https":
		Protocol = HTTPS
		CertFile = sec.Key("CERT_FILE").String()
		KeyFile = sec.Key("KEY_FILE").String()
		if !filepath.IsAbs(CertFile) && len(CertFile) > 0 {
			CertFile = filepath.Join(CustomPath, CertFile)
		}
		if !filepath.IsAbs(KeyFile) && len(KeyFile) > 0 {
			KeyFile = filepath.Join(CustomPath, KeyFile)
		}
	case "fcgi":
		Protocol = FCGI
	case "fcgi+unix":
		Protocol = FCGIUnix
		UnixSocketPermissionRaw := sec.Key("UNIX_SOCKET_PERMISSION").MustString("666")
		UnixSocketPermissionParsed, err := strconv.ParseUint(UnixSocketPermissionRaw, 8, 32)
		if err != nil || UnixSocketPermissionParsed > 0777 {
			log.Fatal("Failed to parse unixSocketPermission: %s", UnixSocketPermissionRaw)
		}
		UnixSocketPermission = uint32(UnixSocketPermissionParsed)
	case "unix":
		Protocol = UnixSocket
		UnixSocketPermissionRaw := sec.Key("UNIX_SOCKET_PERMISSION").MustString("666")
		UnixSocketPermissionParsed, err := strconv.ParseUint(UnixSocketPermissionRaw, 8, 32)
		if err != nil || UnixSocketPermissionParsed > 0777 {
			log.Fatal("Failed to parse unixSocketPermission: %s", UnixSocketPermissionRaw)
		}
		UnixSocketPermission = uint32(UnixSocketPermissionParsed)
	}
	EnableLetsEncrypt = sec.Key("ENABLE_LETSENCRYPT").MustBool(false)
	LetsEncryptTOS = sec.Key("LETSENCRYPT_ACCEPTTOS").MustBool(false)
	if !LetsEncryptTOS && EnableLetsEncrypt {
		log.Warn("Failed to enable Let's Encrypt due to Let's Encrypt TOS not being accepted")
		EnableLetsEncrypt = false
	}
	LetsEncryptDirectory = sec.Key("LETSENCRYPT_DIRECTORY").MustString("https")
	LetsEncryptEmail = sec.Key("LETSENCRYPT_EMAIL").MustString("")
	Domain = sec.Key("DOMAIN").MustString("localhost")
	HTTPAddr = sec.Key("HTTP_ADDR").MustString("0.0.0.0")
	HTTPPort = sec.Key("HTTP_PORT").MustString("3000")
	GracefulRestartable = sec.Key("ALLOW_GRACEFUL_RESTARTS").MustBool(true)
	GracefulHammerTime = sec.Key("GRACEFUL_HAMMER_TIME").MustDuration(60 * time.Second)
	StartupTimeout = sec.Key("STARTUP_TIMEOUT").MustDuration(0 * time.Second)
	EnableWebSocket = sec.Key("ENABLE_WEBSOCKET").MustBool(true)
	CustomRealUrlHeader = sec.Key("CUSTOM_REAL_URL_HEADER").MustString("")

	defaultAppURL := string(Protocol) + "://" + Domain
	if (Protocol == HTTP && HTTPPort != "80") || (Protocol == HTTPS && HTTPPort != "443") {
		defaultAppURL += ":" + HTTPPort
	}
	AppURL = sec.Key("ROOT_URL").MustString(defaultAppURL)
	AppURL = strings.TrimSuffix(AppURL, "/") + "/"

	// Check if has app suburl.
	appURL, err := url.Parse(AppURL)
	if err != nil {
		log.Fatal("Invalid ROOT_URL '%s': %s", AppURL, err)
	}
	// Suburl should start with '/' and end without '/', such as '/{subpath}'.
	// This value is empty if site does not have sub-url.
	AppSubURL = strings.TrimSuffix(appURL.Path, "/")
	StaticURLPrefix = strings.TrimSuffix(sec.Key("STATIC_URL_PREFIX").MustString(AppSubURL), "/")
	AppSubURLDepth = strings.Count(AppSubURL, "/")

	MLOPS = Cfg.Section("").Key("MLOPS").MustBool(false)
	MlopsHost = Cfg.Section("").Key("MLOPS_HOST").MustString("")
	MlopsToken = Cfg.Section("").Key("MLOPS_TOKEN").MustString("")
	MlopsSecret = Cfg.Section("").Key("MLOPS_SECRET").MustString("")
	// Check if Domain differs from AppURL domain than update it to AppURL's domain
	// TODO: Can be replaced with url.Hostname() when minimal GoLang version is 1.8
	urlHostname := strings.SplitN(appURL.Host, ":", 2)[0]
	if urlHostname != Domain && net.ParseIP(urlHostname) == nil {
		Domain = urlHostname
	}

	var defaultLocalURL string
	switch Protocol {
	case UnixSocket:
		defaultLocalURL = "http://unix/"
	case FCGI:
		defaultLocalURL = AppURL
	case FCGIUnix:
		defaultLocalURL = AppURL
	default:
		defaultLocalURL = string(Protocol) + "://"
		if HTTPAddr == "0.0.0.0" {
			defaultLocalURL += "localhost"
		} else {
			defaultLocalURL += HTTPAddr
		}
		defaultLocalURL += ":" + HTTPPort + "/"
	}
	LocalURL = sec.Key("LOCAL_ROOT_URL").MustString(defaultLocalURL)
	RedirectOtherPort = sec.Key("REDIRECT_OTHER_PORT").MustBool(false)
	PortToRedirect = sec.Key("PORT_TO_REDIRECT").MustString("80")
	OfflineMode = sec.Key("OFFLINE_MODE").MustBool()
	DisableRouterLog = sec.Key("DISABLE_ROUTER_LOG").MustBool()
	StaticRootPath = sec.Key("STATIC_ROOT_PATH").MustString(AppWorkPath)
	StaticCacheTime = sec.Key("STATIC_CACHE_TIME").MustDuration(6 * time.Hour)

	AppDataPath = sec.Key("APP_DATA_PATH").MustString(path.Join(AppWorkPath, "data"))
	EnableGzip = sec.Key("ENABLE_GZIP").MustBool()
	EnablePprof = sec.Key("ENABLE_PPROF").MustBool(false)
	PprofDataPath = sec.Key("PPROF_DATA_PATH").MustString(path.Join(AppWorkPath, "data/tmp/pprof"))
	if !filepath.IsAbs(PprofDataPath) {
		PprofDataPath = filepath.Join(AppWorkPath, PprofDataPath)
	}

	switch sec.Key("LANDING_PAGE").MustString("home") {
	case "explore":
		LandingPageURL = LandingPageExplore
	case "organizations":
		LandingPageURL = LandingPageOrganizations
	case "login":
		LandingPageURL = LandingPageLogin
	default:
		LandingPageURL = LandingPageHome
	}

	if len(SSH.Domain) == 0 {
		SSH.Domain = Domain
	}
	SSH.RootPath = path.Join(homeDir, ".ssh")
	serverCiphers := sec.Key("SSH_SERVER_CIPHERS").Strings(",")
	if len(serverCiphers) > 0 {
		SSH.ServerCiphers = serverCiphers
	}
	serverKeyExchanges := sec.Key("SSH_SERVER_KEY_EXCHANGES").Strings(",")
	if len(serverKeyExchanges) > 0 {
		SSH.ServerKeyExchanges = serverKeyExchanges
	}
	serverMACs := sec.Key("SSH_SERVER_MACS").Strings(",")
	if len(serverMACs) > 0 {
		SSH.ServerMACs = serverMACs
	}
	SSH.KeyTestPath = os.TempDir()
	if err = Cfg.Section("server").MapTo(&SSH); err != nil {
		log.Fatal("Failed to map SSH settings: %v", err)
	}

	SSH.KeygenPath = sec.Key("SSH_KEYGEN_PATH").MustString("ssh-keygen")
	SSH.Port = sec.Key("SSH_PORT").MustInt(22)
	SSH.ListenPort = sec.Key("SSH_LISTEN_PORT").MustInt(SSH.Port)

	// When disable SSH, start builtin server value is ignored.
	if SSH.Disabled {
		SSH.StartBuiltinServer = false
	}

	if !SSH.Disabled && !SSH.StartBuiltinServer {
		if err := os.MkdirAll(SSH.RootPath, 0700); err != nil {
			log.Fatal("Failed to create '%s': %v", SSH.RootPath, err)
		} else if err = os.MkdirAll(SSH.KeyTestPath, 0644); err != nil {
			log.Fatal("Failed to create '%s': %v", SSH.KeyTestPath, err)
		}
	}

	SSH.MinimumKeySizeCheck = sec.Key("MINIMUM_KEY_SIZE_CHECK").MustBool()
	minimumKeySizes := Cfg.Section("ssh.minimum_key_sizes").Keys()
	for _, key := range minimumKeySizes {
		if key.MustInt() != -1 {
			SSH.MinimumKeySizes[strings.ToLower(key.Name())] = key.MustInt()
		}
	}
	SSH.AuthorizedKeysBackup = sec.Key("SSH_AUTHORIZED_KEYS_BACKUP").MustBool(true)
	SSH.CreateAuthorizedKeysFile = sec.Key("SSH_CREATE_AUTHORIZED_KEYS_FILE").MustBool(true)
	SSH.ExposeAnonymous = sec.Key("SSH_EXPOSE_ANONYMOUS").MustBool(false)

	sec = Cfg.Section("server")
	if err = sec.MapTo(&LFS); err != nil {
		log.Fatal("Failed to map LFS settings: %v", err)
	}
	LFS.ContentPath = sec.Key("LFS_CONTENT_PATH").MustString(filepath.Join(AppDataPath, "lfs"))
	if !filepath.IsAbs(LFS.ContentPath) {
		LFS.ContentPath = filepath.Join(AppWorkPath, LFS.ContentPath)
	}
	if LFS.LocksPagingNum == 0 {
		LFS.LocksPagingNum = 50
	}

	LFS.HTTPAuthExpiry = sec.Key("LFS_HTTP_AUTH_EXPIRY").MustDuration(20 * time.Minute)

	if LFS.StartServer {
		LFS.JWTSecretBytes = make([]byte, 32)
		n, err := base64.RawURLEncoding.Decode(LFS.JWTSecretBytes, []byte(LFS.JWTSecretBase64))

		if err != nil || n != 32 {
			LFS.JWTSecretBase64, err = generate.NewJwtSecret()
			if err != nil {
				log.Fatal("Error generating JWT Secret for custom config: %v", err)
				return
			}

			// Save secret
			cfg := ini.Empty()
			if com.IsFile(CustomConf) {
				// Keeps custom settings if there is already something.
				if err := cfg.Append(CustomConf); err != nil {
					log.Error("Failed to load custom conf '%s': %v", CustomConf, err)
				}
			}

			cfg.Section("server").Key("LFS_JWT_SECRET").SetValue(LFS.JWTSecretBase64)

			if err := os.MkdirAll(filepath.Dir(CustomConf), os.ModePerm); err != nil {
				log.Fatal("Failed to create '%s': %v", CustomConf, err)
			}
			if err := cfg.SaveTo(CustomConf); err != nil {
				log.Fatal("Error saving generated JWT Secret to custom config: %v", err)
				return
			}
		}
	}

	if err = Cfg.Section("oauth2").MapTo(&OAuth2); err != nil {
		log.Fatal("Failed to OAuth2 settings: %v", err)
		return
	}

	if OAuth2.Enable {
		OAuth2.JWTSecretBytes = make([]byte, 32)
		n, err := base64.RawURLEncoding.Decode(OAuth2.JWTSecretBytes, []byte(OAuth2.JWTSecretBase64))

		if err != nil || n != 32 {
			OAuth2.JWTSecretBase64, err = generate.NewJwtSecret()
			if err != nil {
				log.Fatal("error generating JWT secret: %v", err)
				return
			}
			cfg := ini.Empty()
			if com.IsFile(CustomConf) {
				if err := cfg.Append(CustomConf); err != nil {
					log.Error("failed to load custom conf %s: %v", CustomConf, err)
					return
				}
			}
			cfg.Section("oauth2").Key("JWT_SECRET").SetValue(OAuth2.JWTSecretBase64)

			if err := os.MkdirAll(filepath.Dir(CustomConf), os.ModePerm); err != nil {
				log.Fatal("failed to create '%s': %v", CustomConf, err)
				return
			}
			if err := cfg.SaveTo(CustomConf); err != nil {
				log.Fatal("error saving generating JWT secret to custom config: %v", err)
				return
			}
		}
	}

	sec = Cfg.Section("admin")
	Admin.DefaultEmailNotification = sec.Key("DEFAULT_EMAIL_NOTIFICATIONS").MustString("enabled")

	sec = Cfg.Section("security")
	InstallLock = sec.Key("INSTALL_LOCK").MustBool(false)
	LoginEncrypt = sec.Key("LOGIN_ENCRYPT").MustBool(false)
	SecretKey = sec.Key("SECRET_KEY").MustString("!#@FDEWREWR&*(")
	LogInRememberDays = sec.Key("LOGIN_REMEMBER_DAYS").MustInt(7)
	CookieUserName = sec.Key("COOKIE_USERNAME").MustString("gitea_awesome")
	CookieRememberName = sec.Key("COOKIE_REMEMBER_NAME").MustString("gitea_incredible")
	ReverseProxyAuthUser = sec.Key("REVERSE_PROXY_AUTHENTICATION_USER").MustString("X-WEBAUTH-USER")
	ReverseProxyAuthEmail = sec.Key("REVERSE_PROXY_AUTHENTICATION_EMAIL").MustString("X-WEBAUTH-EMAIL")
	MinPasswordLength = sec.Key("MIN_PASSWORD_LENGTH").MustInt(6)
	ImportLocalPaths = sec.Key("IMPORT_LOCAL_PATHS").MustBool(false)
	DisableGitHooks = sec.Key("DISABLE_GIT_HOOKS").MustBool(false)
	OnlyAllowPushIfGiteaEnvironmentSet = sec.Key("ONLY_ALLOW_PUSH_IF_GITEA_ENVIRONMENT_SET").MustBool(true)
	PasswordHashAlgo = sec.Key("PASSWORD_HASH_ALGO").MustString("pbkdf2")
	CSRFCookieHTTPOnly = sec.Key("CSRF_COOKIE_HTTP_ONLY").MustBool(true)

	InternalToken = loadInternalToken(sec)

	cfgdata := sec.Key("PASSWORD_COMPLEXITY").Strings(",")
	PasswordComplexity = make([]string, 0, len(cfgdata))
	for _, name := range cfgdata {
		name := strings.ToLower(strings.Trim(name, `"`))
		if name != "" {
			PasswordComplexity = append(PasswordComplexity, name)
		}
	}

	sec = Cfg.Section("attachment")
	Attachment.StoreType = sec.Key("STORE_TYPE").MustString("local")
	switch Attachment.StoreType {
	case "local":
		Attachment.Path = sec.Key("PATH").MustString(path.Join(AppDataPath, "attachments"))
		if !filepath.IsAbs(Attachment.Path) {
			Attachment.Path = path.Join(AppWorkPath, Attachment.Path)
		}
	case "minio":
		Attachment.Minio.Endpoint = sec.Key("MINIO_ENDPOINT").MustString("localhost:9000")
		Attachment.Minio.InternalEndpoint = sec.Key("MINIO_INTERNAL_ENDPOINT").MustString("")
		Attachment.Minio.AccessKeyID = sec.Key("MINIO_ACCESS_KEY_ID").MustString("")
		Attachment.Minio.SecretAccessKey = sec.Key("MINIO_SECRET_ACCESS_KEY").MustString("")
		Attachment.Minio.Bucket = sec.Key("MINIO_BUCKET").MustString("gitea")
		Attachment.Minio.Location = sec.Key("MINIO_LOCATION").MustString("us-east-1")
		Attachment.Minio.BasePath = sec.Key("MINIO_BASE_PATH").MustString("attachments/")
		Attachment.Minio.UseSSL = sec.Key("MINIO_USE_SSL").MustBool(false)
		Attachment.Minio.RealPath = sec.Key("MINIO_REAL_PATH").MustString("/mnt/test/minio/data/")
	}

	Attachment.AllowedTypes = strings.Replace(sec.Key("ALLOWED_TYPES").MustString("image/jpeg,image/png,application/zip,application/gzip"), "|", ",", -1)
	Attachment.MaxSize = sec.Key("MAX_SIZE").MustInt64(4)
	Attachment.MaxFiles = sec.Key("MAX_FILES").MustInt(5)
	Attachment.SIZE_FOR_PYTHON = sec.Key("SIZE_FOR_PYTHON").MustInt(500)
	Attachment.Enabled = sec.Key("ENABLED").MustBool(true)

	timeFormatKey := Cfg.Section("time").Key("FORMAT").MustString("")
	if timeFormatKey != "" {
		TimeFormat = map[string]string{
			"ANSIC":       time.ANSIC,
			"UnixDate":    time.UnixDate,
			"RubyDate":    time.RubyDate,
			"RFC822":      time.RFC822,
			"RFC822Z":     time.RFC822Z,
			"RFC850":      time.RFC850,
			"RFC1123":     time.RFC1123,
			"RFC1123Z":    time.RFC1123Z,
			"RFC3339":     time.RFC3339,
			"RFC3339Nano": time.RFC3339Nano,
			"Kitchen":     time.Kitchen,
			"Stamp":       time.Stamp,
			"StampMilli":  time.StampMilli,
			"StampMicro":  time.StampMicro,
			"StampNano":   time.StampNano,
		}[timeFormatKey]
		// When the TimeFormatKey does not exist in the previous map e.g.'2006-01-02 15:04:05'
		if len(TimeFormat) == 0 {
			TimeFormat = timeFormatKey
			TestTimeFormat, _ := time.Parse(TimeFormat, TimeFormat)
			if TestTimeFormat.Format(time.RFC3339) != "2006-01-02T15:04:05Z" {
				log.Warn("Provided TimeFormat: %s does not create a fully specified date and time.", TimeFormat)
				log.Warn("In order to display dates and times correctly please check your time format has 2006, 01, 02, 15, 04 and 05")
			}
			log.Trace("Custom TimeFormat: %s", TimeFormat)
		}
	}

	zone := Cfg.Section("time").Key("DEFAULT_UI_LOCATION").String()
	if zone != "" {
		DefaultUILocation, err = time.LoadLocation(zone)
		if err != nil {
			log.Fatal("Load time zone failed: %v", err)
		} else {
			log.Info("Default UI Location is %v", zone)
		}
	}
	if DefaultUILocation == nil {
		DefaultUILocation = time.Local
	}

	RunUser = Cfg.Section("").Key("RUN_USER").MustString(user.CurrentUsername())
	// Does not check run user when the install lock is off.
	if InstallLock {
		currentUser, match := IsRunUserMatchCurrentUser(RunUser)
		if !match {
			log.Fatal("Expect user '%s' but current user is: %s", RunUser, currentUser)
		}
	}

	SSH.BuiltinServerUser = Cfg.Section("server").Key("BUILTIN_SSH_SERVER_USER").MustString(RunUser)

	newRepository()

	sec = Cfg.Section("picture")
	AvatarUploadPath = sec.Key("AVATAR_UPLOAD_PATH").MustString(path.Join(AppDataPath, "avatars"))
	forcePathSeparator(AvatarUploadPath)
	if !filepath.IsAbs(AvatarUploadPath) {
		AvatarUploadPath = path.Join(AppWorkPath, AvatarUploadPath)
	}
	RepositoryAvatarUploadPath = sec.Key("REPOSITORY_AVATAR_UPLOAD_PATH").MustString(path.Join(AppDataPath, "repo-avatars"))
	forcePathSeparator(RepositoryAvatarUploadPath)
	if !filepath.IsAbs(RepositoryAvatarUploadPath) {
		RepositoryAvatarUploadPath = path.Join(AppWorkPath, RepositoryAvatarUploadPath)
	}
	RepositoryAvatarFallback = sec.Key("REPOSITORY_AVATAR_FALLBACK").MustString("none")
	RepositoryAvatarFallbackImage = sec.Key("REPOSITORY_AVATAR_FALLBACK_IMAGE").MustString("/img/repo_default.png")
	AvatarMaxWidth = sec.Key("AVATAR_MAX_WIDTH").MustInt(4096)
	AvatarMaxHeight = sec.Key("AVATAR_MAX_HEIGHT").MustInt(3072)
	AvatarStaticCacheTime = sec.Key("AVATAR_STATIC_CACHE_TIME").MustDuration(10 * time.Second)
	AvatarMaxFileSize = sec.Key("AVATAR_MAX_FILE_SIZE").MustInt64(1048576)
	switch source := sec.Key("GRAVATAR_SOURCE").MustString("gravatar"); source {
	case "duoshuo":
		GravatarSource = "http://gravatar.duoshuo.com/avatar/"
	case "gravatar":
		GravatarSource = "https://secure.gravatar.com/avatar/"
	case "libravatar":
		GravatarSource = "https://seccdn.libravatar.org/avatar/"
	default:
		GravatarSource = source
	}
	DisableGravatar = sec.Key("DISABLE_GRAVATAR").MustBool()
	EnableFederatedAvatar = sec.Key("ENABLE_FEDERATED_AVATAR").MustBool(!InstallLock)
	if OfflineMode {
		DisableGravatar = true
		EnableFederatedAvatar = false
	}
	if DisableGravatar {
		EnableFederatedAvatar = false
	}
	if EnableFederatedAvatar || !DisableGravatar {
		GravatarSourceURL, err = url.Parse(GravatarSource)
		if err != nil {
			log.Fatal("Failed to parse Gravatar URL(%s): %v",
				GravatarSource, err)
		}
	}

	if EnableFederatedAvatar {
		LibravatarService = libravatar.New()
		if GravatarSourceURL.Scheme == "https" {
			LibravatarService.SetUseHTTPS(true)
			LibravatarService.SetSecureFallbackHost(GravatarSourceURL.Host)
		} else {
			LibravatarService.SetUseHTTPS(false)
			LibravatarService.SetFallbackHost(GravatarSourceURL.Host)
		}
	}

	if err = Cfg.Section("ui").MapTo(&UI); err != nil {
		log.Fatal("Failed to map UI settings: %v", err)
	} else if err = Cfg.Section("markdown").MapTo(&Markdown); err != nil {
		log.Fatal("Failed to map Markdown settings: %v", err)
	} else if err = Cfg.Section("admin").MapTo(&Admin); err != nil {
		log.Fatal("Fail to map Admin settings: %v", err)
	} else if err = Cfg.Section("api").MapTo(&API); err != nil {
		log.Fatal("Failed to map API settings: %v", err)
	} else if err = Cfg.Section("metrics").MapTo(&Metrics); err != nil {
		log.Fatal("Failed to map Metrics settings: %v", err)
	}

	u := *appURL
	u.Path = path.Join(u.Path, "api", "swagger")
	API.SwaggerURL = u.String()

	newGit()

	sec = Cfg.Section("mirror")
	Mirror.MinInterval = sec.Key("MIN_INTERVAL").MustDuration(10 * time.Minute)
	Mirror.DefaultInterval = sec.Key("DEFAULT_INTERVAL").MustDuration(8 * time.Hour)
	if Mirror.MinInterval.Minutes() < 1 {
		log.Warn("Mirror.MinInterval is too low")
		Mirror.MinInterval = 1 * time.Minute
	}
	if Mirror.DefaultInterval < Mirror.MinInterval {
		log.Warn("Mirror.DefaultInterval is less than Mirror.MinInterval")
		Mirror.DefaultInterval = time.Hour * 8
	}

	Langs = Cfg.Section("i18n").Key("LANGS").Strings(",")
	if len(Langs) == 0 {
		Langs = []string{
			"en-US", "zh-CN", "zh-HK", "zh-TW", "de-DE", "fr-FR", "nl-NL", "lv-LV",
			"ru-RU", "uk-UA", "ja-JP", "es-ES", "pt-BR", "pl-PL", "bg-BG", "it-IT",
			"fi-FI", "tr-TR", "cs-CZ", "sr-SP", "sv-SE", "ko-KR"}
	}
	Names = Cfg.Section("i18n").Key("NAMES").Strings(",")
	if len(Names) == 0 {
		Names = []string{"English", "简体中文", "繁體中文（香港）", "繁體中文（台灣）", "Deutsch",
			"français", "Nederlands", "latviešu", "русский", "Українська", "日本語",
			"español", "português do Brasil", "polski", "български", "italiano",
			"suomi", "Türkçe", "čeština", "српски", "svenska", "한국어"}
	}
	dateLangs = Cfg.Section("i18n.datelang").KeysHash()

	ShowFooterBranding = Cfg.Section("other").Key("SHOW_FOOTER_BRANDING").MustBool(false)
	ShowFooterVersion = Cfg.Section("other").Key("SHOW_FOOTER_VERSION").MustBool(true)
	ShowFooterTemplateLoadTime = Cfg.Section("other").Key("SHOW_FOOTER_TEMPLATE_LOAD_TIME").MustBool(true)

	UI.ShowUserEmail = Cfg.Section("ui").Key("SHOW_USER_EMAIL").MustBool(true)
	UI.DefaultShowFullName = Cfg.Section("ui").Key("DEFAULT_SHOW_FULL_NAME").MustBool(false)
	UI.SearchRepoDescription = Cfg.Section("ui").Key("SEARCH_REPO_DESCRIPTION").MustBool(true)

	HasRobotsTxt = com.IsFile(path.Join(CustomPath, "robots.txt"))

	newMarkup()

	sec = Cfg.Section("U2F")
	U2F.TrustedFacets, _ = shellquote.Split(sec.Key("TRUSTED_FACETS").MustString(strings.TrimRight(AppURL, "/")))
	U2F.AppID = sec.Key("APP_ID").MustString(strings.TrimRight(AppURL, "/"))

	zip.Verbose = false

	UI.ReactionsMap = make(map[string]bool)
	for _, reaction := range UI.Reactions {
		UI.ReactionsMap[reaction] = true
	}

	sec = Cfg.Section("machinery")
	Broker = sec.Key("BROKER").MustString("redis://localhost:6379")
	DefaultQueue = sec.Key("DEFAULT_QUEUE").MustString("DecompressTasksQueue")
	ResultBackend = sec.Key("RESULT_BACKEND").MustString("redis://localhost:6379")

	sec = Cfg.Section("decompress")
	DecompressAddress = sec.Key("HOST").MustString("http://192.168.207.34:39987")
	AuthUser = sec.Key("USER").MustString("")
	AuthPassword = sec.Key("PASSWORD").MustString("")

	sec = Cfg.Section("labelsystem")
	LabelTaskName = sec.Key("LabelTaskName").MustString("LabelRedisQueue")
	LabelDatasetDeleteQueue = sec.Key("LabelDatasetDeleteQueue").MustString("LabelDatasetDeleteQueue")
	DecompressOBSTaskName = sec.Key("DecompressOBSTaskName").MustString("LabelDecompressOBSQueue")

	sec = Cfg.Section("homepage")
	RecommentRepoAddr = sec.Key("Address").MustString("https://openi.pcl.ac.cn/OpenIOSSG/promote/raw/branch/master/")
	ESSearchURL = sec.Key("ESSearchURL").MustString("http://192.168.207.94:9200")
	INDEXPOSTFIX = sec.Key("INDEXPOSTFIX").MustString("")

	sec = Cfg.Section("notice")
	UserNameOfNoticeRepo = sec.Key("USER_NAME").MustString("OpenIOSSG")
	RepoNameOfNoticeRepo = sec.Key("REPO_NAME").MustString("promote")
	RefNameOfNoticeRepo = sec.Key("REF_NAME").MustString("master")
	TreePathOfNoticeRepo = sec.Key("TREE_PATH").MustString("notice.json")
	CacheTimeOutSecond = sec.Key("CACHE_TIME_OUT_SECOND").MustInt(60)
	SDK_NOTICE = sec.Key("SDK_NOTICE").MustString("")

	sec = Cfg.Section("dync_config")
	DyncConfigRepoOwner = sec.Key("OWNER").MustString("OpenIOSSG")
	DyncConfigRepoName = sec.Key("REPO").MustString("promote")
	DyncConfigBranch = sec.Key("BRANCH").MustString("master")
	DyncConfigDefaultTTL = sec.Key("DEFAULT_TTL").MustDuration(2 * time.Minute)
	PayConputilityConfigRepoOwner = sec.Key("ComputilityConfigOWNER").MustString("OpenIOSSG")
	PayConputilityConfigRepoName = sec.Key("ComputilityConfigREPO").MustString("HaoSuanResource")
	PayConputilitycConfigBranch = sec.Key("ComputilityConfigBRANCH").MustString("master")
	PeraPayConputilityConfigRepoOwner = sec.Key("PERA_PAY_CONPUTILITY_CONFIG_OWNER").MustString("OpenIOSSG")
	PeraPayConputilityConfigRepoName = sec.Key("PERA_PAY_CONPUTILITY_CONFIG_REPO").MustString("ParateraResource")
	PeraPayConputilitycConfigBranch = sec.Key("PERA_PAY_CONPUTILITY_CONFIG_BRANCH").MustString("master")

	ChinaMobilePayConputilityConfigRepoOwner = sec.Key("CHINA_MOBILE_PAY_CONPUTILITY_CONFIG_OWNER").MustString("OpenIOSSG")
	ChinaMobilePayConputilityConfigRepoName = sec.Key("CHINA_MOBILE_PAY_CONPUTILITY_CONFIG_REPO").MustString("ecloudResource")
	ChinaMobilePayConputilitycConfigBranch = sec.Key("CHINA_MOBILE_PAY_CONPUTILITY_CONFIG_BRANCH").MustString("master")

	ConputilityPartnerConfigRepoOwner = sec.Key("CONPUTILITY_PARTNER_CONFIG_OWNER").MustString("OpenIOSSG")
	ConputilityPartnerConfigRepoName = sec.Key("CONPUTILITY_PARTNER_CONFIG_REPO").MustString("promote")
	ConputilityPartnerConfigBranch = sec.Key("CONPUTILITY_PARTNER_CONFIG_BRANCH").MustString("master")

	sec = Cfg.Section("cloudbrain")
	CBAuthUser = sec.Key("USER").MustString("")
	CBAuthPassword = sec.Key("PWD").MustString("")
	RestServerHost = sec.Key("REST_SERVER_HOST").MustString("http://192.168.202.73")
	JobPath = sec.Key("JOB_PATH").MustString("/datasets/minio/data/opendata/jobs/")
	CBCodePathPrefix = sec.Key("CODE_PATH_PREFIX").MustString("jobs/")
	DebugServerHost = sec.Key("DEBUG_SERVER_HOST").MustString("http://192.168.202.73")
	JobType = sec.Key("GPU_TYPE_DEFAULT").MustString("openidebug")
	GpuTypes = sec.Key("GPU_TYPES").MustString("")
	ResourceSpecs = sec.Key("RESOURCE_SPECS").MustString("")
	MaxDuration = sec.Key("MAX_DURATION").MustInt64(14400)
	TrainGpuTypes = sec.Key("TRAIN_GPU_TYPES").MustString("")
	TrainResourceSpecs = sec.Key("TRAIN_RESOURCE_SPECS").MustString("")
	MaxModelSize = sec.Key("MAX_MODEL_SIZE").MustFloat64(200)
	MaxModelSizeForWhole = sec.Key("MAX_MODEL_SIZE_FOR_WHOLE").MustFloat64(200)
	CloudbrainDownloadAllConcurrency = sec.Key("DOWNLOAD_ALL_CONCURRENCY").MustInt(2)
	InferenceGpuTypes = sec.Key("INFERENCE_GPU_TYPES").MustString("")
	InferenceResourceSpecs = sec.Key("INFERENCE_RESOURCE_SPECS").MustString("")
	SpecialPools = sec.Key("SPECIAL_POOL").MustString("")

	MaxDatasetNum = sec.Key("MAX_DATASET_NUM").MustInt(5)
	CullIdleTimeout = sec.Key("CULL_IDLE_TIMEOUT").MustString("900")
	CullInterval = sec.Key("CULL_INTERVAL").MustString("60")
	DebugAttachSize = sec.Key("DEBUG_ATTACH_SIZE").MustInt(20)

	CloudbrainUniquenessLockTime = sec.Key("UNIQUENESS_LOCK_TIME").MustDuration(5 * time.Minute)
	SyncCloudbrainStatusDuration = sec.Key("SYNC_STATUS_DURATION").MustString("@every 1m")

	sec = Cfg.Section("ai_task")
	rangeMap := map[string][]string{}
	tmp := sec.Key("NEW_TASK_RANGE").MustString("")
	json.Unmarshal([]byte(tmp), &rangeMap)
	AI_TASK_RANGE = rangeMap
	PREPARING_MAX_WAIT_DURATION = sec.Key("ENABLED").MustDuration(15 * time.Minute)
	OUTPUT_SHOW_MAX_KEY = sec.Key("OUTPUT_SHOW_MAX_KEY").MustInt(100)
	OUTPUT_DOWNLOAD_MAX_KEY = sec.Key("OUTPUT_DOWNLOAD_MAX_KEY").MustInt(1000)
	SPECIFICATION_SPECIAL_QUEUE = sec.Key("SPECIFICATION_SPECIAL_QUEUE").MustString("{}")
	DEBUG_MODEL_NUM_LIMIT = sec.Key("DEBUG_MODEL_NUM_LIMIT").MustInt(5)
	DEBUG_MODEL_SIZE_LIMIT_GB = sec.Key("DEBUG_MODEL_SIZE_LIMIT_GB").MustInt(20)

	roleLimitMap := map[string]map[string]int{}
	tmpRoleLimitStr := sec.Key("ROLE_MULTI_LIMIT_MAP").MustString("")
	json.Unmarshal([]byte(tmpRoleLimitStr), &roleLimitMap)
	ROLE_MULTI_LIMIT_MAP = roleLimitMap

	AI_TASK_TEMPLATE_FILE_NAME = sec.Key("AI_TASK_TEMPLATE_FILE_NAME").MustString(".openi-aitemplate")
	DEBUG_ADDRESS_CHECK_GLOBAL_TOGGLE = sec.Key("DEBUG_ADDRESS_CHECK_GLOBAL_TOGGLE").MustBool(true)
	MAX_MULTI_FINE_TUNE_TASK = sec.Key("MAX_MULTI_FINE_TUNE_TASK").MustInt(3)
	MAX_MULTI_SD_FINETUNE_TASK = sec.Key("MAX_MULTI_SD_FINETUNE_TASK").MustInt(3)
	SD_FINETUNE_MODEL_NUM_LIMIT = sec.Key("SD_FINETUNE_MODEL_NUM_LIMIT").MustInt(10)
	SD_FINETUNE_MODEL_SIZE_LIMIT_GB = sec.Key("SD_FINETUNE_MODEL_SIZE_LIMIT_GB").MustInt(500)
	COMFYUI_EXPERIENCE_MODEL_NUM_LIMIT = sec.Key("COMFYUI_EXPERIENCE_MODEL_NUM_LIMIT").MustInt(10)
	COMFYUI_EXPERIENCE_MODEL_SIZE_LIMIT_GB = sec.Key("COMFYUI_EXPERIENCE_MODEL_SIZE_LIMIT_GB").MustInt(500)

	FORCE_OLD_OUTPUT_MIGRATE = sec.Key("FORCE_OLD_OUTPUT_MIGRATE").MustBool(false)
	FORCE_OLD_CODE_STORAGE = sec.Key("FORCE_OLD_CODE_STORAGE").MustBool(false)

	sec = Cfg.Section("dataset")
	MAX_PREVIEW_FILE_SIZE = sec.Key("MAX_PREVIEW_FILE_SIZE").MustInt64(100 * 1024 * 1024)
	ENABLE_ORGANIZATION_UPLOAD_DATASET_FILE = sec.Key("ENABLE_ORGANIZATION_UPLOAD_DATASET_FILE").MustBool(false)
	HOT_DATASET_LIST_URL = sec.Key("HOT_DATASET_LIST_URL").MustString("")
	DATASET_STORAGE_TYPE = sec.Key("DATASET_STORAGE_TYPE").MustString("")
	DatasetSizeChangedQueueBatchSize = sec.Key("SIZE_CHANGED_QUEUE_BATCH_SIZE").MustInt(10)

	sec = Cfg.Section("model")
	AIMODEL_STORAGE_TYPE = sec.Key("AIMODEL_STORAGE_TYPE").MustString("")
	AimodelSizeChangedQueueBatchSize = sec.Key("SIZE_CHANGED_QUEUE_BATCH_SIZE").MustInt(10)

	sec = Cfg.Section("benchmark")
	IsBenchmarkEnabled = sec.Key("ENABLED").MustBool(false)
	BenchmarkOwner = sec.Key("OWNER").MustString("")
	BenchmarkName = sec.Key("NAME").MustString("")
	BenchmarkServerHost = sec.Key("HOST").MustString("")
	BenchmarkCategory = sec.Key("CATEGORY").MustString("")
	BenchmarkTypes = sec.Key("TYPES").MustString("")
	BenchmarkGpuTypes = sec.Key("GPU_TYPES").MustString("")
	BenchmarkResourceSpecs = sec.Key("RESOURCE_SPECS").MustString("")
	BenchmarkMaxDuration = sec.Key("MAX_DURATION").MustInt64(14400)

	sec = Cfg.Section("snn4imagenet")
	IsSnn4imagenetEnabled = sec.Key("ENABLED").MustBool(false)
	Snn4imagenetOwner = sec.Key("OWNER").MustString("")
	Snn4imagenetName = sec.Key("NAME").MustString("")
	Snn4imagenetServerHost = sec.Key("HOST").MustString("")
	ModelBenchmarkMaxDuration = sec.Key("MAX_DURATION").MustInt64(28800)

	sec = Cfg.Section("brainscore")
	IsBrainScoreEnabled = sec.Key("ENABLED").MustBool(false)
	BrainScoreOwner = sec.Key("OWNER").MustString("")
	BrainScoreName = sec.Key("NAME").MustString("")
	BrainScoreServerHost = sec.Key("HOST").MustString("")
	BrainScoreRegion = strings.Split(sec.Key("REGION").MustString(""), ",")

	sec = Cfg.Section("snn4ecoset")
	IsSnn4EcosetEnabled = sec.Key("ENABLED").MustBool(false)
	Snn4EcosetOwner = sec.Key("OWNER").MustString("")
	Snn4EcosetName = sec.Key("NAME").MustString("")
	Snn4EcosetServerHost = sec.Key("HOST").MustString("")
	Snn4AttachmentName = sec.Key("DATASET").MustString("")

	sec = Cfg.Section("sim2brain_snn")
	IsSim2BrainSnnEnabled = sec.Key("ENABLED").MustBool(false)
	Sim2BrainSnnOwner = sec.Key("OWNER").MustString("")
	Sim2BrainSnnName = sec.Key("NAME").MustString("")
	Sim2BrainSnnServerHost = sec.Key("HOST").MustString("")
	Sim2BrainSnnDatasetTypes = strings.Split(sec.Key("DATASET_TYPE").MustString(""), ",")
	Sim2BrainSnnAttachNames = strings.Split(sec.Key("ATTACH_NAME").MustString(""), ",")

	sec = Cfg.Section("blockchain")
	BlockChainHost = sec.Key("HOST").MustString("http://192.168.136.66:3302/")
	CommitValidDate = sec.Key("COMMIT_VALID_DATE").MustString("2021-01-15")

	sec = Cfg.Section("obs")
	Endpoint = sec.Key("ENDPOINT").MustString("")
	AccessKeyID = sec.Key("ACCESS_KEY_ID").MustString("")
	SecretAccessKey = sec.Key("SECRET_ACCESS_KEY").MustString("")
	Bucket = sec.Key("BUCKET").MustString("")
	Location = sec.Key("LOCATION").MustString("")
	BasePath = sec.Key("BASE_PATH").MustString("")
	TrainJobModelPath = sec.Key("TrainJobModel_Path").MustString("job/")
	OutPutPath = sec.Key("Output_Path").MustString("output/")
	CodePathPrefix = sec.Key("CODE_PATH_PREFIX").MustString("job/")
	UserBasePath = sec.Key("BASE_PATH_USER").MustString("users/")
	PROXYURL = sec.Key("PROXY_URL").MustString("")

	sec = Cfg.Section("urchin")
	UrchinAddress = sec.Key("URCHIN_ADDRESS").MustString("")
	UrchinReqTimeoutSecond = sec.Key("URCHIN_REQ_TIMEOUT_SECOND").MustInt64(0)
	UrchinMaxConnection = sec.Key("URCHIN_MAX_CONNECTION").MustInt(0)

	sec = Cfg.Section("storage")
	StorageDefaultType = sec.Key("DEFAULT").MustString("")
	if StorageDefaultType == "" {
		if Endpoint != "" {
			StorageDefaultType = "obs"
		} else {
			StorageDefaultType = "minio"
		}

	}

	GetModelartsConfig()

	sec = Cfg.Section("elk")
	ElkUrl = sec.Key("ELKURL").MustString("")
	ElkUser = sec.Key("ELKUSER").MustString("")
	ElkPassword = sec.Key("ELKPASSWORD").MustString("")
	Index = sec.Key("INDEX").MustString("")
	TimeField = sec.Key("TIMEFIELD").MustString(" @timestamptest")
	ElkTimeFormat = sec.Key("ELKTIMEFORMAT").MustString("date_time")
	PROJECT_LIMIT_PAGES = strings.Split(sec.Key("project_limit_pages").MustString(""), ",")

	sec = Cfg.Section("wechat")
	WechatApiHost = sec.Key("HOST").MustString("https://api.weixin.qq.com")
	WechatApiTimeoutSeconds = sec.Key("TIMEOUT_SECONDS").MustInt(3)
	WechatAppId = sec.Key("APP_ID").MustString("wxba77b915a305a57d")
	WechatAppSecret = sec.Key("APP_SECRET").MustString("")
	WechatQRCodeExpireSeconds = sec.Key("QR_CODE_EXPIRE_SECONDS").MustInt(120)
	WechatAuthSwitch = sec.Key("AUTH_SWITCH").MustBool(false)
	WechatSignUpRequired = sec.Key("SIGN_UP_REQUIRED").MustBool(false)
	UserNameOfWechatReply = sec.Key("AUTO_REPLY_USER_NAME").MustString("OpenIOSSG")
	RepoNameOfWechatReply = sec.Key("AUTO_REPLY_REPO_NAME").MustString("promote")
	RefNameOfWechatReply = sec.Key("AUTO_REPLY_REF_NAME").MustString("master")
	TreePathOfAutoMsgReply = sec.Key("AUTO_REPLY_TREE_PATH").MustString("wechat/auto_reply.json")
	TreePathOfSubscribe = sec.Key("SUBSCRIBE_TREE_PATH").MustString("wechat/subscribe_reply.json")
	CloudbrainStartedTemplateId = sec.Key("CLOUDBRAIN_STARTED_TEMPLATE_ID").MustString("")
	CloudbrainStartedNotifyList = strings.Split(sec.Key("CLOUDBRAIN_STARTED_NOTIFY_LIST").MustString("DEBUG"), ",")
	CloudbrainStartedTitle = sec.Key("CLOUDBRAIN_STARTED_TITLE").MustString("您好，您提交的算力资源申请已通过，任务已启动，请您关注运行情况。")
	CloudbrainStartedRemark = sec.Key("CLOUDBRAIN_STARTED_REMARK").MustString("感谢您的耐心等待。")
	CloudbrainStoppedTemplateId = sec.Key("CLOUDBRAIN_STOPPED_TEMPLATE_ID").MustString("")
	CloudbrainStoppedNotifyList = strings.Split(sec.Key("CLOUDBRAIN_STOPPED_NOTIFY_LIST").MustString("TRAIN"), ",")
	CloudbrainStoppedTitle = sec.Key("CLOUDBRAIN_STOPPED_TITLE").MustString("您好，您申请的算力资源已结束使用，任务已完成运行，状态为%s，请您关注运行结果")
	CloudbrainStoppedRemark = sec.Key("CLOUDBRAIN_STOPPED_REMARK").MustString("感谢您的耐心等待。")
	CloudbrainComingStopTemplateId = sec.Key("CLOUDBRAIN_COMING_STOP_TEMPLATE_ID").MustString("")
	CloudbrainComingStopTitle = sec.Key("CLOUDBRAIN_COMING_STOP_TITLE").MustString("因算力积分余额不足，您正在运行的任务即将被停止。")
	CloudbrainComingStopChargeLink = sec.Key("CLOUDBRAIN_COMING_STOP_CHARGE_LINK").MustString("请参考“个人中心 > 算力积分”页面的“积分获取说明”获取积分。")
	CloudbrainComingStopRemark = sec.Key("CLOUDBRAIN_COMING_STOP_REMARK").MustString("为了不影响您使用，请您尽快获取算力积分。")
	CloudbrainComingStopSendFlag = sec.Key("CLOUDBRAIN_COMING_STOP_SEND_FLAG").MustBool(true)
	CloudbrainLongRunningNotifyList = strings.Split(sec.Key("CLOUDBRAIN_LONG_RUNNING_NOTIFY_LIST").MustString("GENERAL"), ",")
	CloudbrainLongRunningTemplateId = sec.Key("CLOUDBRAIN_LONG_RUNNING_TEMPLATE_ID").MustString("")
	CloudbrainLongRunningNotifyInterval, _ = time.ParseDuration(sec.Key("CLOUDBRAIN_LONG_RUNNING_NOTIFY_INTERVAL").MustString(""))

	BindReplySuccess = sec.Key("BIND_REPLY_SUCCESS").MustString("扫码成功，您可以使用OpenI启智社区算力环境。")
	BindReplyWechatAccountUsed = sec.Key("BIND_REPLY_WECHAT_ACCOUNT_USED").MustString("认证失败，您的微信号已绑定其他启智账号")
	BindReplyOpenIAccountUsed = sec.Key("BIND_REPLY_OPENI_ACCOUNT_USED").MustString("认证失败，您待认证的启智账号已绑定其他微信号")
	BindReplyFailedDefault = sec.Key("BIND_REPLY_FAILED_DEFAULT").MustString("微信认证失败")
	BindReplySignUp = sec.Key("BIND_REPLY_SIGN_UP").MustString("扫码成功，您正在通过微信扫码注册OpenI启智社区。<a href =\"%s\">继续完成注册流程</a>")
	BindReplySignUpErr4WechatUsed = sec.Key("BIND_REPLY_SIGN_UP_ERR_4_WECHAT_USED").MustString("微信号已被其他用户绑定，请更换微信号")
	BindReplyLoginSuccess = sec.Key("BIND_REPLY_LOGIN_SUCCESS").MustString("扫码成功，您正在通过微信扫码登录OpenI启智社区。")
	BindReplyNoUser = sec.Key("BIND_REPLY_NO_USER").MustString("登录失败，您的微信号未绑定OpenI启智社区账号。")

	sec = Cfg.Section("repo-square")
	IncubationSourceOrgName = sec.Key("INCUBATION_ORG_NAME").MustString("OpenI")
	PaperRepoTopicName = sec.Key("PAPER_REPO_TOPIC_NAME").MustString("openi-paper")

	sec = Cfg.Section("point")
	CloudBrainPaySwitch = sec.Key("CLOUDBRAIN_PAY_SWITCH").MustBool(false)
	CloudBrainPayDelay = sec.Key("CLOUDBRAIN_PAY_DELAY").MustDuration(30 * time.Minute)
	CloudBrainPayInterval = sec.Key("CLOUDBRAIN_PAY_INTERVAL").MustDuration(60 * time.Minute)
	DeductTaskRange = sec.Key("DEDUCT_TASK_RANGE").MustDuration(30 * time.Minute)
	DeductTaskRangeForFirst = sec.Key("DEDUCT_TASK_RANGE_FOR_FIRST").MustDuration(3 * time.Hour)
	DeductTaskMinTimestamp = sec.Key("DEDUCT_TASK_MIN_TIMESTAMP").MustInt64(0)
	UseDynamicData = sec.Key("USE_DYNAMIC_DATA").MustBool(false)

	sec = Cfg.Section("model-migrate")
	UseLocalMinioMigrate = sec.Key("USE_LOCAL_MINIO_MIGRATE").MustBool(false)

	sec = Cfg.Section("icons")
	BadgeIconMaxFileSize = sec.Key("BADGE_ICON_MAX_FILE_SIZE").MustInt64(1048576)
	BadgeIconMaxWidth = sec.Key("BADGE_ICON_MAX_WIDTH").MustInt(4096)
	BadgeIconMaxHeight = sec.Key("BADGE_ICON_MAX_HEIGHT").MustInt(3072)
	BadgeIconDefaultSize = sec.Key("BADGE_ICON_DEFAULT_SIZE").MustUint(200)
	IconUploadPath = sec.Key("ICON_UPLOAD_PATH").MustString(path.Join(AppDataPath, "icons"))

	SetRadarMapConfig()

	sec = Cfg.Section("warn_mail")
	Warn_Notify_Mails = strings.Split(sec.Key("mails").MustString(""), ",")

	sec = Cfg.Section("course")
	Course.OrgName = sec.Key("org_name").MustString("")
	Course.TeamName = sec.Key("team_name").MustString("")

	sec = Cfg.Section("file_notebook")
	FileNoteBook.ProjectName = sec.Key("project_name").MustString("openi-notebook")
	FileNoteBook.ImageIdNPU = sec.Key("imageid_npu").MustString("")
	FileNoteBook.ImageGPU = sec.Key("image_gpu").MustString("")
	FileNoteBook.SpecIdCPU = sec.Key("specid_cpu").MustInt64(-1)
	FileNoteBook.SpecIdGPU = sec.Key("specid_gpu").MustInt64(-1)
	FileNoteBook.SpecIdNPU = sec.Key("specid_npu").MustInt64(-1)
	FileNoteBook.ImageIdNPUCD = sec.Key("imageid_npu_cd").MustString("")
	FileNoteBook.SpecIdNPUCD = sec.Key("specid_npu_cd").MustInt64(-1)
	FileNoteBook.ImageCPUDescription = sec.Key("image_cpu_desc").MustString("")
	FileNoteBook.ImageGPUDescription = sec.Key("image_gpu_desc").MustString("")
	FileNoteBook.ImageNPUDescription = sec.Key("image_npu_desc").MustString("")
	FileNoteBook.ImageNPUCDDescription = sec.Key("image_npu_cd_desc").MustString("")

	sec = Cfg.Section("kanban")
	IsCloudbrainTimingEnabled = sec.Key("ENABLED").MustBool(false)
	CloudbrainNumTargetTime = sec.Key("CLOUDBRAIN_NUM_TARGET_TIME").MustString("2025/06")

	sec = Cfg.Section("wenxin")
	BaiduWenXin.ModelArtsWenXinURL = sec.Key("ModelArts_WenXinURL").MustString("")
	BaiduWenXin.BAIDU_API_KEY = sec.Key("BAIDU_API_KEY").MustString("")
	BaiduWenXin.BAIDU_SECRET_KEY = sec.Key("BAUDY_SECRET_KEY").MustString("")
	BaiduWenXin.BAIDU_TOKEN_URL = sec.Key("BAIDU_TOKEN_URL").MustString("https://aip.baidubce.com/oauth/2.0/token")
	BaiduWenXin.PICTURE_BAIDU_REVIEW_URL = sec.Key("PICTURE_BAIDU_REVIEW_URL").MustString("https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined?access_token=")
	BaiduWenXin.TEXT_BAIDU_REVIEW_URL = sec.Key("TEXT_BAIDU_REVIEW_URL").MustString("https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token=")
	BaiduWenXin.RUN_WORKERS = sec.Key("RUN_WORKERS").MustInt(1)
	BaiduWenXin.MODEL_SERVERS = sec.Key("MODEL_SERVERS").MustInt(1)

	sec = Cfg.Section("llm_chat_api")
	LLM_CHAT_API.CHATGLM2_HOST = sec.Key("CHATGLM2_HOST").MustString("")
	LLM_CHAT_API.CHATGLM2_MAX_LENGTH = sec.Key("CHATGLM2_MAX_LENGTH").MustInt(8192)
	LLM_CHAT_API.LLAMA2_HOST = sec.Key("LLAMA2_HOST").MustString("")
	LLM_CHAT_API.LLAMA2_MAX_LENGTH = sec.Key("LLAMA2_MAX_LENGTH").MustInt(4096)
	LLM_CHAT_API.COMMON_KB = sec.Key("COMMON_KNOWLEDGE_BASE").MustString("")
	LLM_CHAT_API.MAX_FREE_TRIES = sec.Key("MAX_FREE_TRIES").MustInt64(200)
	LLM_CHAT_API.LEGAL_CHECK = sec.Key("LEGAL_CHECK").MustBool(false)
	LLM_CHAT_API.LEGAL_MAX_COUNT = sec.Key("LEGAL_MAX_COUNT").MustInt64(5)
	LLM_CHAT_API.CHAT_EXPIRED_MINUTES = sec.Key("CHAT_EXPIRED_MINUTES").MustInt64(30)

	sec = Cfg.Section("deepseek_modelarts")
	DeepSeekModelarts.URL_QWEN_7B = sec.Key("URL_QWEN_7B").MustString("")
	DeepSeekModelarts.URL_QWEN_1_5B = sec.Key("URL_QWEN_1_5B").MustString("")
	DeepSeekModelarts.URL_LLAMA_8B = sec.Key("URL_LLAMA_8B").MustString("")
	DeepSeekModelarts.URL_QWEN_14B = sec.Key("URL_QWEN_14B").MustString("")
	DeepSeekModelarts.URL_QWEN_32B_W8A8 = sec.Key("URL_QWEN_32B_W8A8").MustString("")
	DeepSeekModelarts.URL_V3_MINDFORMERS = sec.Key("URL_V3_MINDFORMERS").MustString("")
	DeepSeekModelarts.SERVICE_HUB_URL = sec.Key("SERVICE_HUB_URL").MustString("http://192.168.207.34:8111/chat")
	DeepSeekModelarts.SERVICE_AUTO_MODEL_URL = sec.Key("SERVICE_AUTO_MODEL_URL").MustString("http://192.168.206.27:2025")
	DeepSeekModelarts.AccessKey = sec.Key("AccessKey").MustString("")
	DeepSeekModelarts.SecretKey = sec.Key("SecretKey").MustString("")

	sec = Cfg.Section("baidu_ai")
	BAIDU_AI.API_KEY = sec.Key("API_KEY").MustString("")
	BAIDU_AI.SECRET_KEY = sec.Key("SECRET_KEY").MustString("")
	BAIDU_AI.URL = sec.Key("URL").MustString("https://aip.baidubce.com/oauth/2.0/token")
	BAIDU_AI.LEGAL_TEXT_URL = sec.Key("LEGAL_TEXT_URL").MustString("https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token=")

	sec = Cfg.Section("hf_model")
	ExternalTransfer.HfModelOwner = sec.Key("HF_MODEL_OWNER").MustString("")
	ExternalTransfer.HfDomain = sec.Key("HF_DOMAIN").MustString("")
	ExternalTransfer.HfToken = sec.Key("HF_TOKEN").MustString("")
	ExternalTransfer.MaxModelSizeGb = sec.Key("MAX_MODEL_SIZE_GB").MustInt(200)
	ExternalTransfer.HeartbeatCron = sec.Key("HEARTBEAT_CRON").MustString("@every 1m")
	ExternalTransfer.MaxPerUser = sec.Key("MAX_PER_USER").MustInt(1)

	sec = Cfg.Section("model_experience")
	MODEL_EXPERIENCE.MAX_CREATES = sec.Key("MAX_CREATES").MustInt(3)
	MODEL_EXPERIENCE.PORT_MIN = sec.Key("PORT_MIN").MustInt(8000)
	MODEL_EXPERIENCE.PORT_MAX = sec.Key("PORT_MAX").MustInt(8800)
	MODEL_EXPERIENCE.MODELBASE_REPO_OWNER = sec.Key("MODELBASE_REPO_OWNER").MustString("OpenIOSSG")

	sec = Cfg.Section("resource_queue")
	UseNewSyncQueueAPI = sec.Key("USE_NEW_SYNC_QUEUE_API").MustBool(true)

	GetGrampusConfig()
	GetModelartsCDConfig()
	getModelConvertConfig()
	getModelSafetyConfig()
	getModelAppConfig()
	getClearStrategy()
	getNotebookStrategy()
	getCardStatistic()
	NewScreenMapConfig()
	getOauth2Config()
	GetIFLYTekConfig()
	GetDiffusionAvatarConfig()
}

func GetModelartsConfig() {
	sec := Cfg.Section("modelarts")
	ModelArtsHost = sec.Key("ENDPOINT").MustString("")
	IamHost = sec.Key("IAMHOST").MustString("")
	ProjectID = sec.Key("PROJECT_ID").MustString("")
	ProjectName = sec.Key("PROJECT_NAME").MustString("")
	ModelArtsUsername = sec.Key("USERNAME").MustString("")
	ModelArtsPassword = sec.Key("PASSWORD").MustString("")
	ModelArtsDomain = sec.Key("DOMAIN").MustString("")
	AllowedOrg = sec.Key("ORGANIZATION").MustString("")
	ProfileID = sec.Key("PROFILE_ID").MustString("")
	PoolInfos = sec.Key("POOL_INFOS").MustString("")
	ImageInfos = sec.Key("IMAGE_INFOS").MustString("")
	Capacity = sec.Key("CAPACITY").MustInt(100)
	MaxTempQueryTimes = sec.Key("MAX_TEMP_QUERY_TIMES").MustInt(30)
	ResourcePools = sec.Key("Resource_Pools").MustString("")
	Engines = sec.Key("Engines").MustString("")
	EngineVersions = sec.Key("Engine_Versions").MustString("")
	FlavorInfos = sec.Key("FLAVOR_INFOS").MustString("")
	TrainJobFLAVORINFOS = sec.Key("TrainJob_FLAVOR_INFOS").MustString("")
	ModelArtsSpecialPools = sec.Key("SPECIAL_POOL").MustString("")
	ModelArtsMultiNode = sec.Key("MULTI_NODE").MustString("")
	ModelArtsShareAddr = sec.Key("ModelArts_Share_Addr").MustString("192.168.0.30:/")
	ModelArtsMountPath = sec.Key("ModelArts_Mount_Path").MustString("/cache/sfs")
	ModelArtsNasType = sec.Key("ModelArts_Nas_Type").MustString("nfs")

	getFineTuneConfig()
	getFlowControlConfig()
}

func getModelSafetyConfig() {
	sec := Cfg.Section("model_safety_test")
	ModelSafetyTest.GPUBaseDataSetName = sec.Key("GPUBaseDataSetName").MustString("")
	ModelSafetyTest.GPUBaseDataSetUUID = sec.Key("GPUBaseDataSetUUID").MustString("")
	ModelSafetyTest.GPUCombatDataSetName = sec.Key("GPUCombatDataSetName").MustString("")
	ModelSafetyTest.GPUCombatDataSetUUID = sec.Key("GPUCombatDataSetUUID").MustString("")
	ModelSafetyTest.NPUBaseDataSetName = sec.Key("NPUBaseDataSetName").MustString("")
	ModelSafetyTest.NPUBaseDataSetUUID = sec.Key("NPUBaseDataSetUUID").MustString("")
	ModelSafetyTest.NPUCombatDataSetName = sec.Key("NPUCombatDataSetName").MustString("")
	ModelSafetyTest.NPUCombatDataSetUUID = sec.Key("NPUCombatDataSetUUID").MustString("")
	ModelSafetyTest.HOST = sec.Key("HOST").MustString("")
	ModelSafetyTest.KEY = sec.Key("KEY").MustString("")
}

func getModelConvertConfig() {
	sec := Cfg.Section("model_convert")
	ModelConvert.GPU_PYTORCH_IMAGE = sec.Key("GPU_PYTORCH_IMAGE").MustString("192.168.204.28:5000/default-workspace/99280a9940ae44ca8f5892134386fddb/image:tensorrt_8.0_for_c2net_1")
	ModelConvert.GpuQueue = sec.Key("GpuQueue").MustString("openidgx")
	ModelConvert.GPU_TENSORFLOW_IMAGE = sec.Key("GPU_TENSORFLOW_IMAGE").MustString("dockerhub.pcl.ac.cn:5000/user-images/openi:tf2onnx")
	ModelConvert.NPU_MINDSPORE_16_IMAGE = sec.Key("NPU_MINDSPORE_16_IMAGE").MustString("swr.cn-south-222.ai.pcl.cn/openi/mindspore1.8.1_train_openi_new:v1")
	ModelConvert.PytorchOnnxBootFile = sec.Key("PytorchOnnxBootFile").MustString("convert_pytorch.py")
	ModelConvert.PytorchTrTBootFile = sec.Key("PytorchTrTBootFile").MustString("convert_pytorch_tensorrt.py")
	ModelConvert.MindsporeBootFile = sec.Key("MindsporeBootFile").MustString("convert_mindspore.py")
	ModelConvert.TensorFlowNpuBootFile = sec.Key("TensorFlowNpuBootFile").MustString("convert_tensorflow.py")
	ModelConvert.TensorFlowGpuBootFile = sec.Key("TensorFlowGpuBootFile").MustString("convert_tensorflow_gpu.py")
	ModelConvert.ConvertRepoPath = sec.Key("ConvertRepoPath").MustString("https://openi.pcl.ac.cn/zouap/npu_test")
	ModelConvert.GPU_Resource_Specs_ID = sec.Key("GPU_Resource_Specs_ID").MustString("e677036134cd11ed9c2a06e51cc0c06b")
	ModelConvert.GPU_AiCenter_Code = sec.Key("GPU_AiCenter_Code").MustString("cloudbrain1")
	ModelConvert.GPU_Spec_ID = sec.Key("GPU_Spec_ID").MustInt64(183)
	ModelConvert.GPU_Spec_AccCardType = sec.Key("GPU_Spec_AccCardType").MustString("A100")
	ModelConvert.NPU_FlavorCode = sec.Key("NPU_FlavorCode").MustString("modelarts.bm.910.arm.public.1")
	ModelConvert.NPU_PoolID = sec.Key("NPU_PoolID").MustString("pool7908321a")
	ModelConvert.NPU_MINDSPORE_IMAGE_ID = sec.Key("NPU_MINDSPORE_IMAGE_ID").MustInt(37)
	ModelConvert.NPU_TENSORFLOW_IMAGE_ID = sec.Key("NPU_TENSORFLOW_IMAGE_ID").MustInt(38)
	ModelConvert.GPU_PADDLE_IMAGE = sec.Key("GPU_PADDLE_IMAGE").MustString("dockerhub.pcl.ac.cn:5000/user-images/openi:paddle2.3.0_gpu_cuda11.2_cudnn8")
	ModelConvert.GPU_MXNET_IMAGE = sec.Key("GPU_MXNET_IMAGE").MustString("192.168.204.22:5000/default-workspace/99280a9940ae44ca8f5892134386fddb/image:mxnet191cu_cuda102_py37_c2net")
	ModelConvert.PaddleOnnxBootFile = sec.Key("PaddleOnnxBootFile").MustString("convert_paddle.py")
	ModelConvert.MXnetOnnxBootFile = sec.Key("MXnetOnnxBootFile").MustString("convert_mxnet.py")
}
func getFlowControlConfig() {
	sec := Cfg.Section("flow_control")
	FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK = sec.Key("ALL_ATTACHEMENT_NUM_SDK").MustInt(100)
	FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR = sec.Key("ATTACHEMENT_NUM_A_USER_LAST24HOUR").MustInt(1000)
	FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M = sec.Key("ATTACHEMENT_NUM_A_USER_LAST10M").MustInt(10)
	FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER = sec.Key("ATTACHEMENT_SIZE_A_USER").MustInt64(500)
	FLOW_CONTROL.IGNORE_FLAG = sec.Key("IGNORE_FLAG").MustString("")
}

func getModelAppConfig() {
	sec := Cfg.Section("model_app")
	ModelApp.DesensitizationUrl = sec.Key("desensitization_url").MustString("")
}

func GetModelartsCDConfig() {
	sec := Cfg.Section("modelarts-cd")

	ModelartsCD.Enabled = sec.Key("ENABLED").MustBool(false)
	ModelartsCD.EndPoint = sec.Key("ENDPOINT").MustString("https://modelarts.cn-southwest-228.cdzs.cn")
	ModelartsCD.ProjectID = sec.Key("PROJECT_ID").MustString("")
	ModelartsCD.AccessKey = sec.Key("ACCESS_KEY").MustString("")
	ModelartsCD.SecretKey = sec.Key("SECRET_KEY").MustString("")
	ModelartsCD.ImageInfos = sec.Key("IMAGE_INFOS").MustString("")
	ModelartsCD.FlavorInfos = sec.Key("FLAVOR_INFOS").MustString("")

	getNotebookImageInfos()
	getNotebookFlavorInfos()
}

func getClearStrategy() {

	sec := Cfg.Section("clear_strategy")
	ClearStrategy.Enabled = sec.Key("ENABLED").MustBool(false)
	ClearStrategy.ResultSaveDays = sec.Key("RESULT_SAVE_DAYS").MustInt(30)
	ClearStrategy.BatchSize = sec.Key("BATCH_SIZE").MustInt(500)
	ClearStrategy.DebugJobSize = sec.Key("DEBUG_BATCH_SIZE").MustInt(100)
	ClearStrategy.TrashSaveDays = sec.Key("TRASH_SAVE_DAYS").MustInt(90)
	ClearStrategy.Cron = sec.Key("CRON").MustString("0 0/30 2-8 * * ?")
	ClearStrategy.RunAtStart = sec.Key("RUN_AT_START").MustBool(false)
}

func getNotebookStrategy() {

	sec := Cfg.Section("notebook_strategy")
	NotebookStrategy.ClearEnabled = sec.Key("CLEAR_ENABLED").MustBool(false)
	NotebookStrategy.ClearBatchSize = sec.Key("CLEAR_BATCH_SIZE").MustInt(300)
	NotebookStrategy.MaxNumberPerUser = sec.Key("MAX_NUMBER").MustInt(5)
	NotebookStrategy.Cron = sec.Key("CRON").MustString("* 0,0 2-8 * * ?")
	NotebookStrategy.RunAtStart = sec.Key("RUN_AT_START").MustBool(false)
}

func getCardStatistic() {

	sec := Cfg.Section("card_statistic")
	CardStatistic.Enabled = sec.Key("ENABLED").MustBool(false)
	CardStatistic.Cron = sec.Key("CRON").MustString("@daily")
	CardStatistic.RunAtStart = sec.Key("RUN_AT_START").MustBool(false)
}

func GetGrampusConfig() {
	sec := Cfg.Section("grampus")

	Grampus.Env = sec.Key("ENV").MustString("TEST")
	Grampus.Host = sec.Key("SERVER_HOST").MustString("")
	Grampus.UserName = sec.Key("USERNAME").MustString("")
	Grampus.Password = sec.Key("PASSWORD").MustString("")
	Grampus.SpecialPools = sec.Key("SPECIAL_POOL").MustString("")
	//Grampus.C2NetSequence = sec.Key("C2NET_SEQUENCE").MustString("{\"sequence\":[{\"id\":1,\"name\":\"cloudbrain_one\",\"content\":\"鹏城云脑一号\",\"content_en\":\"Pencheng Cloudbrain Ⅰ\"},{\"id\":2,\"name\":\"cloudbrain_two\",\"content\":\"鹏城云脑二号\",\"content_en\":\"Pencheng Cloudbrain Ⅱ\"},{\"id\":3,\"name\":\"beida\",\"content\":\"北大人工智能集群系统\",\"content_en\":\"Peking University AI Center\"},{\"id\":4,\"name\":\"hefei\",\"content\":\"合肥类脑智能开放平台\",\"content_en\":\"Hefei AI Center\"},{\"id\":5,\"name\":\"wuhan\",\"content\":\"武汉人工智能计算中心\",\"content_en\":\"Wuhan AI Center\"},{\"id\":6,\"name\":\"xian\",\"content\":\"西安未来人工智能计算中心\",\"content_en\":\"Xi'an AI Center\"},{\"id\":7,\"pclcci\":\"more\",\"content\":\"鹏城云计算所\",\"content_en\":\"Pengcheng Cloud Computing Institute\"},{\"id\":8,\"name\":\"xuchang\",\"content\":\"中原人工智能计算中心\",\"content_en\":\"Zhongyuan AI Center\"},{\"id\":9,\"name\":\"chengdu\",\"content\":\"成都人工智能计算中心\",\"content_en\":\"Chengdu AI Center\"},{\"id\":10,\"name\":\"more\",\"content\":\"横琴先进智能计算中心\",\"content_en\":\"Hengqin AI Center\"},{\"id\":11,\"name\":\"more\",\"content\":\"国家超级计算济南中心\",\"content_en\":\"HPC & AI Center\"}]}")
	Grampus.AiCenterCodeAndNameInfo = sec.Key("AI_CENTER_CODE_AND_NAME").MustString("{\"sequence\":[{\"id\":1,\"name\":\"cloudbrain_one\",\"content\":\"鹏城云脑一号\",\"content_en\":\"Pencheng Cloudbrain Ⅰ\"},{\"id\":2,\"name\":\"cloudbrain_two\",\"content\":\"鹏城云脑二号\",\"content_en\":\"Pencheng Cloudbrain Ⅱ\"},{\"id\":3,\"name\":\"beida\",\"content\":\"北大人工智能集群系统\",\"content_en\":\"Peking University AI Center\"},{\"id\":4,\"name\":\"hefei\",\"content\":\"合肥类脑智能开放平台\",\"content_en\":\"Hefei AI Center\"},{\"id\":5,\"name\":\"wuhan\",\"content\":\"武汉人工智能计算中心\",\"content_en\":\"Wuhan AI Center\"},{\"id\":6,\"name\":\"xian\",\"content\":\"西安未来人工智能计算中心\",\"content_en\":\"Xi'an AI Center\"},{\"id\":7,\"pclcci\":\"more\",\"content\":\"鹏城云计算所\",\"content_en\":\"Pengcheng Cloud Computing Institute\"},{\"id\":8,\"name\":\"xuchang\",\"content\":\"中原人工智能计算中心\",\"content_en\":\"Zhongyuan AI Center\"},{\"id\":9,\"name\":\"chengdu\",\"content\":\"成都人工智能计算中心\",\"content_en\":\"Chengdu AI Center\"},{\"id\":10,\"name\":\"more\",\"content\":\"横琴先进智能计算中心\",\"content_en\":\"Hengqin AI Center\"},{\"id\":11,\"name\":\"more\",\"content\":\"国家超级计算济南中心\",\"content_en\":\"HPC & AI Center\"}]}")
	Grampus.AiCenterCodeAndNameAndLocInfo = sec.Key("AI_CENTER_CODE_AND_NAME_AND_LOC").MustString("{\"sequence\":[{\"id\":1,\"name\":\"cloudbrain_one\",\"content\":\"鹏城云脑一号\",\"content_en\":\"Pencheng Cloudbrain Ⅰ\"},{\"id\":2,\"name\":\"cloudbrain_two\",\"content\":\"鹏城云脑二号\",\"content_en\":\"Pencheng Cloudbrain Ⅱ\"},{\"id\":3,\"name\":\"beida\",\"content\":\"北大人工智能集群系统\",\"content_en\":\"Peking University AI Center\"},{\"id\":4,\"name\":\"hefei\",\"content\":\"合肥类脑智能开放平台\",\"content_en\":\"Hefei AI Center\"},{\"id\":5,\"name\":\"wuhan\",\"content\":\"武汉人工智能计算中心\",\"content_en\":\"Wuhan AI Center\"},{\"id\":6,\"name\":\"xian\",\"content\":\"西安未来人工智能计算中心\",\"content_en\":\"Xi'an AI Center\"},{\"id\":7,\"pclcci\":\"more\",\"content\":\"鹏城云计算所\",\"content_en\":\"Pengcheng Cloud Computing Institute\"},{\"id\":8,\"name\":\"xuchang\",\"content\":\"中原人工智能计算中心\",\"content_en\":\"Zhongyuan AI Center\"},{\"id\":9,\"name\":\"chengdu\",\"content\":\"成都人工智能计算中心\",\"content_en\":\"Chengdu AI Center\"},{\"id\":10,\"name\":\"more\",\"content\":\"横琴先进智能计算中心\",\"content_en\":\"Hengqin AI Center\"},{\"id\":11,\"name\":\"more\",\"content\":\"国家超级计算济南中心\",\"content_en\":\"HPC & AI Center\"}]}")

	Grampus.UsageRateBeginTime = sec.Key("USAGE_RATE_BEGIN_TIME").MustString("2021-01-01 00:00:00")
	Grampus.GPUImageCommonName = sec.Key("GPU_IMAGE_COMMON_NAME").MustString("image")

	if Grampus.AiCenterCodeAndNameAndLocInfo != "" {
		if err := json.Unmarshal([]byte(Grampus.AiCenterCodeAndNameAndLocInfo), &C2NetLocInfos); err != nil {
			log.Error("Unmarshal(AiCenterCodeAndNameLocInfo) failed:%v", err)
		}
		AiCenterCodeAndNameAndLocMapInfo = make(map[string]*C2NetSequenceInfo)
		for _, value := range C2NetLocInfos.C2NetSqInfo {
			AiCenterCodeAndNameAndLocMapInfo[value.Name] = value
		}
		applyDemoAiCenterLocations()
	}
	if Grampus.AiCenterCodeAndNameInfo != "" {
		if err := json.Unmarshal([]byte(Grampus.AiCenterCodeAndNameInfo), &C2NetInfos); err != nil {
			log.Error("Unmarshal(AiCenterCodeAndNameInfo) failed:%v", err)
		}
		AiCenterCodeAndNameMapInfo = make(map[string]*C2NetSequenceInfo)
		for _, value := range C2NetInfos.C2NetSqInfo {
			AiCenterCodeAndNameMapInfo[value.Name] = value
		}
	}

	Grampus.SyncScriptProject = sec.Key("SYNC_SCRIPT_PROJECT").MustString("script_for_grampus")
	Grampus.LocalCenterID = sec.Key("LOCAL_CENTER_ID").MustString("cloudbrain2")
	Grampus.GPULocalCenterID = sec.Key("GPU_LOCAL_CENTER_ID").MustString("openi")
	Grampus.AiCenterInfo = sec.Key("AI_CENTER_INFO").MustString("")
	if Grampus.AiCenterInfo != "" {
		if err := json.Unmarshal([]byte(Grampus.AiCenterInfo), &CenterInfos); err != nil {
			log.Error("Unmarshal(AiCenterInfo) failed:%v", err)
		}
	}
	Grampus.MultiNode = sec.Key("MULTI_NODE").MustString("")

	Grampus.MMLSparkMaxTime = sec.Key("MMLSparkMaxTime").MustInt64(8 * 3600)

	Grampus.NoteBookDomainURL = sec.Key("NoteBookDomainURL").MustString("https://mlunotebook.openi.org.cn")
	Grampus.NoteBookLocalURL = sec.Key("NoteBookLocalURL").MustString("http://192.168.242.50")
}

func GetIFLYTekConfig() {
	sec := Cfg.Section("iflytek")

	IFLYTekConfig.TOKEN = sec.Key("TOKEN").MustString("")
	IFLYTekConfig.DataServiceHost = sec.Key("DATA_SERVICE_HOST").MustString("")
	IFLYTekConfig.TrainServiceHost = sec.Key("TRAIN_SERVICE_HOST").MustString("")
	IFLYTekConfig.OnlineServiceHost = sec.Key("ONLINE_SERVICE_HOST").MustString("")
	IFLYTekConfig.ModelServiceHost = sec.Key("MODEL_SERVICE_HOST").MustString("")
	IFLYTekConfig.OnlineServiceAPIUrl = sec.Key("ONLINE_SERVICE_API_URL").MustString("")
	IFLYTekConfig.OnlineServiceAPIToken = sec.Key("ONLINE_SERVICE_API_TOKEN").MustString("")
	IFLYTekConfig.DefaultLabelId = sec.Key("LABEL_ID").MustString("")
	IFLYTekConfig.SystemStructure = sec.Key("SYSTEM_STRUCTURE").MustString("")
	IFLYTekConfig.ComputeSource = sec.Key("COMPUTE_SOURCE").MustString("")
	IFLYTekConfig.XNum = sec.Key("X_NUM").MustString("")
}

func SetRadarMapConfig() {
	sec := Cfg.Section("radar_map")

	RadarMap.Impact = sec.Key("impact").MustFloat64(0.3)
	RadarMap.ImpactWatch = sec.Key("impact_watch").MustFloat64(0.1)
	RadarMap.ImpactStar = sec.Key("impact_star").MustFloat64(0.2)
	RadarMap.ImpactFork = sec.Key("impact_fork").MustFloat64(0.3)
	RadarMap.ImpactCodeDownload = sec.Key("impact_code_download").MustFloat64(0.2)
	RadarMap.ImpactComments = sec.Key("impact_comments").MustFloat64(0.1)
	RadarMap.ImpactBrowser = sec.Key("impact_browser").MustFloat64(0.1)
	RadarMap.Completeness = sec.Key("completeness").MustFloat64(0.1)
	RadarMap.CompletenessIssuesClosed = sec.Key("completeness_issues_closed").MustFloat64(0.2)
	RadarMap.CompletenessReleases = sec.Key("completeness_releases").MustFloat64(0.3)
	RadarMap.CompletenessDevelopAge = sec.Key("completeness_develop_age").MustFloat64(0.1)
	RadarMap.CompletenessDataset = sec.Key("completeness_dataset").MustFloat64(0)
	RadarMap.CompletenessModel = sec.Key("completeness_model").MustFloat64(0.1)
	RadarMap.CompletenessWiki = sec.Key("completeness_wiki").MustFloat64(0.1)
	RadarMap.Liveness = sec.Key("liveness").MustFloat64(0.3)
	RadarMap.LivenessCommit = sec.Key("liveness_commit").MustFloat64(0.2)
	RadarMap.LivenessIssue = sec.Key("liveness_issue").MustFloat64(0.2)
	RadarMap.LivenessPR = sec.Key("liveness_pr").MustFloat64(0.2)
	RadarMap.LivenessRelease = sec.Key("liveness_release").MustFloat64(0.4)
	RadarMap.ProjectHealth = sec.Key("project_health").MustFloat64(0.1)
	RadarMap.ProjectHealthIssueCompleteRatio = sec.Key("project_health_issue_complete_ratio").MustFloat64(100)
	RadarMap.ProjectHealth0IssueCloseRatio = sec.Key("project_health_0_issue_close_ratio").MustFloat64(0.0)
	RadarMap.TeamHealth = sec.Key("team_health").MustFloat64(0.1)
	RadarMap.TeamHealthContributors = sec.Key("team_health_contributors").MustFloat64(0.2)
	RadarMap.TeamHealthKeyContributors = sec.Key("team_health_key_contributors").MustFloat64(0.6)
	RadarMap.TeamHealthContributorsAdded = sec.Key("team_health_contributors_added").MustFloat64(0.2)
	RadarMap.Growth = sec.Key("growth").MustFloat64(0.1)
	RadarMap.GrowthCodeLines = sec.Key("growth_code_lines").MustFloat64(0.2)
	RadarMap.GrowthIssue = sec.Key("growth_issue").MustFloat64(0.2)
	RadarMap.GrowthContributors = sec.Key("growth_contributors").MustFloat64(0.2)
	RadarMap.GrowthCommit = sec.Key("growth_commit").MustFloat64(0.2)
	RadarMap.GrowthComments = sec.Key("growth_comments").MustFloat64(0.2)
	RadarMap.RecordBeginTime = sec.Key("record_begin_time").MustString("2021-11-05")
	RadarMap.GrowthBeginTime = sec.Key("growth_begin_time").MustString("2022-05-20")
	RadarMap.IgnoreMirrorRepo = sec.Key("ignore_mirror_repo").MustBool(true)

}

func loadInternalToken(sec *ini.Section) string {
	uri := sec.Key("INTERNAL_TOKEN_URI").String()
	if len(uri) == 0 {
		return loadOrGenerateInternalToken(sec)
	}
	tempURI, err := url.Parse(uri)
	if err != nil {
		log.Fatal("Failed to parse INTERNAL_TOKEN_URI (%s): %v", uri, err)
	}
	switch tempURI.Scheme {
	case "file":
		fp, err := os.OpenFile(tempURI.RequestURI(), os.O_RDWR, 0600)
		if err != nil {
			log.Fatal("Failed to open InternalTokenURI (%s): %v", uri, err)
		}
		defer fp.Close()

		buf, err := ioutil.ReadAll(fp)
		if err != nil {
			log.Fatal("Failed to read InternalTokenURI (%s): %v", uri, err)
		}
		// No token in the file, generate one and store it.
		if len(buf) == 0 {
			token, err := generate.NewInternalToken()
			if err != nil {
				log.Fatal("Error generate internal token: %v", err)
			}
			if _, err := io.WriteString(fp, token); err != nil {
				log.Fatal("Error writing to InternalTokenURI (%s): %v", uri, err)
			}
			return token
		}

		return string(buf)
	default:
		log.Fatal("Unsupported URI-Scheme %q (INTERNAL_TOKEN_URI = %q)", tempURI.Scheme, uri)
	}
	return ""
}

func loadOrGenerateInternalToken(sec *ini.Section) string {
	var err error
	token := sec.Key("INTERNAL_TOKEN").String()
	if len(token) == 0 {
		token, err = generate.NewInternalToken()
		if err != nil {
			log.Fatal("Error generate internal token: %v", err)
		}

		// Save secret
		cfgSave := ini.Empty()
		if com.IsFile(CustomConf) {
			// Keeps custom settings if there is already something.
			if err := cfgSave.Append(CustomConf); err != nil {
				log.Error("Failed to load custom conf '%s': %v", CustomConf, err)
			}
		}

		cfgSave.Section("security").Key("INTERNAL_TOKEN").SetValue(token)

		if err := os.MkdirAll(filepath.Dir(CustomConf), os.ModePerm); err != nil {
			log.Fatal("Failed to create '%s': %v", CustomConf, err)
		}
		if err := cfgSave.SaveTo(CustomConf); err != nil {
			log.Fatal("Error saving generated INTERNAL_TOKEN to custom config: %v", err)
		}
	}
	return token
}

func ensureLFSDirectory() {
	if LFS.StartServer {
		if err := os.MkdirAll(LFS.ContentPath, 0700); err != nil {
			log.Fatal("Failed to create '%s': %v", LFS.ContentPath, err)
		}
	}
}

func getNotebookImageInfos() {

	if ModelartsCD.Enabled {
		json.Unmarshal([]byte(ModelartsCD.ImageInfos), &StImageInfos)
	} else {
		json.Unmarshal([]byte(ImageInfos), &StImageInfos)
	}
}

func getNotebookFlavorInfos() {

	if ModelartsCD.Enabled {
		json.Unmarshal([]byte(ModelartsCD.FlavorInfos), &StFlavorInfo)
	} else {
		json.Unmarshal([]byte(FlavorInfos), &StFlavorInfo)
	}
}

// NewServices initializes the services
func NewServices() {
	InitDBConfig()
	newService()
	NewLogServices(false)
	ensureLFSDirectory()
	newCacheService()
	newSessionService()
	newCORSService()
	newMailService()
	newRegisterMailService()
	newNotifyMailService()
	newWebhookService()
	newMigrationsService()
	newIndexerService()
	newTaskService()
	NewQueueService()
	newPhoneService()
	newGuestInfo()
	initRateLimitConfig()
	iniAimConfig()
	iniMonitorTaskConfig()
	newEvalService()
	InitOTEL()
}
