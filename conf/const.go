package conf

const (
	Version                 = "1.0.1"
	defaultConfName         = "engine.yaml"
	defaultRpcPort          = 9999
	defaultEngineBackupDir  = "backup"
	defaultLogDir           = "logs"
	defaultLogFileName      = "engine.log"
	defaultLogPanicFileName = "panic.log"
	defaultLogLevel         = "debug"
	defaultLogMaxAge        = 60 * 60 * 24
	defaultLogRotationTime  = 60 * 60 * 24 * 7
)

type versionType []int32

var (
	AppRootDir = ""
	AppAbsDir  = ""
)
