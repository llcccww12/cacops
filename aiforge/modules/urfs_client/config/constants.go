package config

import (
	"time"
)

// Reason of backing to source.
const (
	BackSourceReasonNone          = 0
	BackSourceReasonRegisterFail  = 1
	BackSourceReasonMd5NotMatch   = 2
	BackSourceReasonDownloadError = 3
	BackSourceReasonNoSpace       = 4
	BackSourceReasonInitError     = 5
	BackSourceReasonWriteError    = 6
	BackSourceReasonHostSysError  = 7
	BackSourceReasonNodeEmpty     = 8
	BackSourceReasonSourceError   = 10
	BackSourceReasonUserSpecified = 100
	ForceNotBackSourceAddition    = 1000
)

// Download pattern.
const (
	PatternP2P      = "p2p"
	PatternSeedPeer = "seed-peer"
	PatternSource   = "source"
)

//// Download limit.
//const (
//	DefaultPerPeerDownloadLimit = 20 * unit.MB
//	DefaultTotalDownloadLimit   = 100 * unit.MB
//	DefaultUploadLimit          = 100 * unit.MB
//	DefaultMinRate              = 20 * unit.MB
//)

// Others.
const (
	DefaultTimestampFormat = "2006-01-02 15:04:05"
	SchemaHTTP             = "http"

	DefaultTaskExpireTime  = 6 * time.Hour
	DefaultGCInterval      = 1 * time.Minute
	DefaultDaemonAliveTime = 5 * time.Minute
	DefaultScheduleTimeout = 5 * time.Minute
	DefaultDownloadTimeout = 5 * time.Minute

	DefaultSchedulerSchema = "http"
	DefaultSchedulerIP     = "127.0.0.1"
	DefaultSchedulerPort   = 8002

	DefaultPieceChanSize     = 16
	DefaultObjectMaxReplicas = 3
)

// Dfcache subcommand names.
const (
	CmdStat   = "stat"
	CmdImport = "import"
	CmdExport = "export"
	CmdDelete = "delete"
)

// Service defalut port of listening.
const (
	DefaultEndPort                = 65535
	DefaultPeerStartPort          = 65000
	DefaultUploadStartPort        = 65002
	DefaultObjectStorageStartPort = 65004
	DefaultHealthyStartPort       = 40901
)

var (
	// DefaultCertValidityPeriod is default validity period of certificate.
	DefaultCertValidityPeriod = 180 * 24 * time.Hour
)
