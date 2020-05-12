package conf

const (
	Version                 = ""
	defaultConfName         = "engine.yaml"
	defaultRpcPort          = 9999
	defaultLogDir           = "logs"
	defaultLogFileName      = "engine.log"
	defaultLogPanicFileName = "panic.log"
	defaultLogLevel         = "debug"
	defaultLogMaxAge        = 60 * 60 * 24
	defaultLogRotationTime  = 60 * 60 * 24 * 7
)

var (
	AppRootDir = ""
	AppAbsDir  = ""
)
