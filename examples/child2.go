package main

import (
	"encoding/json"
	"engine/common"
	"engine/log"
	"time"
)

type child2 struct {
	environment common.PluginProgramEnvironment
}

func (plugin child2) Name() string {
	return "child2"
}
func (plugin child2) Version() string {
	return "1.0.1"
}
func (plugin child2) Start(conf interface{}) {
	log.Infof("start child2")
	go func() {
		for {
			moment := &common.Moment{
				ShareUser: "child2",
				ShareTime: time.Now().Unix(),
				Mood:      "child2 share",
			}
			plugin.environment.Share("child2", moment)
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			moment, _ := plugin.environment.OpenShare("child1")
			momentJson, _ := json.Marshal(moment)
			log.Infof("child1 share momentJson: %s", string(momentJson))
		}
	}()
}
func (plugin child2) Stop() {

}

func NewChild2(environment common.PluginProgramEnvironment) common.PluginProgram {
	log.Log = environment.Logger()
	return &child2{environment: environment}
}
