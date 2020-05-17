package model

import (
	"engine/common"
	"engine/log"
)

var PluginProgramEnvironment = &pluginProgramEnvironment{share: map[string]*common.Moment{}}

type pluginProgramEnvironment struct {
	share map[string]*common.Moment
}

func (env *pluginProgramEnvironment) Logger() *log.Logger {
	return log.Log
}

func (env *pluginProgramEnvironment) PluginProgram(pluginName string) (*common.PluginRunner, error) {
	program, err := PluginLoader.LoadPluginProgram(pluginName)
	if err != nil {
		return nil, err
	}
	return &common.PluginRunner{PluginProgram: program}, nil
}

func (env *pluginProgramEnvironment) Share(key string, moment *common.Moment) {
	env.share[key] = moment
	return
}

func (env *pluginProgramEnvironment) OpenShare(key string) (*common.Moment, bool) {
	moment, exist := env.share[key]
	return moment, exist
}
