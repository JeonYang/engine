package common

import "engine/log"

type PluginProgram interface {
	Name() string
	Version() string
	Start(conf interface{})
	Stop()
}

type PluginProgramEnvironment interface {
	Logger() *log.Logger
	PluginProgram(pluginName string) (*PluginRunner, error)
	Share(key string, moment *Moment)
	OpenShare(key string) (*Moment, bool)
}

type Moment struct {
	ShareUser string
	ShareTime int64
	Mood      interface{}
}

//type PluginProgramBuilder func(env PluginProgramEnvironment) PluginProgram
