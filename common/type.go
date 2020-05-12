package common

type PluginProgram interface {
	Name() string
	Version() string
	Start(conf string)
	Stop()
}