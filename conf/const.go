package conf

const (
	Version                = ""
	defaultConfName        = "engine.yaml"
	defaultRpcPort         = "9999"
	defaultLogFileName     = "engine.log"
	defaultLogLevel        = "debug"
	defaultLogMaxAge       = 60 * 60 * 24
	defaultLogRotationTime = 60 * 60 * 24 * 7
)

var (
	AppRootDir = ""
	AppAbsDir  = ""
)
