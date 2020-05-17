package main

import (
	"encoding/json"
	"engine/common"
	"engine/log"
)

type PluginMain struct {
	environment common.PluginProgramEnvironment
}

func (plugin PluginMain) Name() string {
	return "pluginMain"
}
func (plugin PluginMain) Version() string {
	return "1.0.1"
}
func (plugin PluginMain) Start(conf interface{}) {
	confStr, ok := conf.(string)
	if !ok {
		confStr = "{}"
	}
	log.Infof("pluginMainConf: %s", confStr)

	pluginMainConf := make(map[string]string)
	json.Unmarshal([]byte(confStr), &pluginMainConf)

	log.Infof("start child...")
	for pluginName, pluginConf := range pluginMainConf {
		log.Infof("pluginMain start pluginName: %v", pluginName)
		pluginProgram, err := plugin.environment.PluginProgram(pluginName)
		if err != nil {
			log.Errorf("pluginMain load: %s fail,err: %v", pluginName, err)
			continue
		}
		pluginProgram.Start(pluginConf)
		log.Infof("start child %s success", pluginName)
	}
}
func (plugin PluginMain) Stop() {

}

func NewPluginMain(environment common.PluginProgramEnvironment) common.PluginProgram {
	log.Log = environment.Logger()
	return &PluginMain{environment: environment}
}
