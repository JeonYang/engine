package main

import (
	"encoding/json"
	"engine/common"
	"engine/log"
	"time"
)

type Child1 struct {
	environment common.PluginProgramEnvironment
}

func (plugin Child1) Name() string {
	return "child1"
}
func (plugin Child1) Version() string {
	return "1.0.1"
}
func (plugin Child1) Start(conf interface{}) {
	log.Infof("start child1")
	go func() {
		for {
			moment := &common.Moment{
				ShareUser: "child1",
				ShareTime: time.Now().Unix(),
				Mood:      "child1 share",
			}
			plugin.environment.Share("child1", moment)
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			moment, _ := plugin.environment.OpenShare("child2")
			momentJson, _ := json.Marshal(moment)
			log.Infof("child2 share momentJson: %s", string(momentJson))
		}
	}()
}
func (plugin Child1) Stop() {

}

func NewChild1(environment common.PluginProgramEnvironment) common.PluginProgram {
	log.Log = environment.Logger()
	return &Child1{environment: environment}
}
