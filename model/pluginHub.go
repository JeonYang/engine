package model

import (
	"bytes"
	"encoding/gob"
	"engine/common"
	"fmt"
	"reflect"
	"sync"
)

var PluginHub = pluginHub{runnerHub: make(map[string]*common.PluginRunner)}

type pluginHub struct {
	mutex     sync.Mutex
	runnerHub map[string]*common.PluginRunner
}

func (hub *pluginHub) PushPluginProgram(program common.PluginProgram) error {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	_, exist := hub.runnerHub[program.Name()]
	if exist {
		return fmt.Errorf("the plugin: %s exist in hub.", program.Name())
	}
	hub.runnerHub[program.Name()] = &common.PluginRunner{PluginProgram: program}
	return nil
}

func (hub *pluginHub) deepCopy(src Program) (dst interface{}, err error) {
	dst = reflect.New(reflect.TypeOf(src).Elem()).Interface()
	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(src)
	if err != nil {
		return
	}
	err = gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
	if err != nil {
		return
	}
	return
}

func (hub *pluginHub) GetPluginRunner(name string) (*common.PluginRunner, bool) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	return runner, exist
}

func (hub *pluginHub) DeletePluginRunner(name string) *common.PluginRunner {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	if exist {
		runner.Stop()
		delete(hub.runnerHub, name)
	}
	return runner
}

func (hub *pluginHub) StopAllPlugin() {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	for _, runner := range hub.runnerHub {
		runner.Stop()
	}
}

func (hub *pluginHub) RestartAllPlugin() {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	for _, runner := range hub.runnerHub {
		runner.ReStart(nil)
	}
}
