package model

import (
	"bytes"
	"encoding/gob"
	"engine/log"
	"fmt"
	"reflect"
	"sync"
)

var PluginHub = pluginHub{runnerHub: make(map[string]*PluginRunner)}

type pluginHub struct {
	mutex     sync.Mutex
	runnerHub map[string]*PluginRunner
}

func (hub *pluginHub) PushPluginProgram(program Program) error {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	_, exist := hub.runnerHub[program.Name()]
	if exist {
		return fmt.Errorf("the plugin: %s exist in hub.", program.Name())
	}
	hub.runnerHub[program.Name()] = &PluginRunner{Program: program}
	return nil
}

func (hub *pluginHub) CopyProgram(name string) (Program, bool) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	if !exist {
		return nil, exist
	}
	//var program Program
	//var err error
	program, err := hub.deepCopy(runner.Program)
	if err != nil {
		log.Errorf("deep copy program fail,[err: %v]", err)
		return nil, false
	}
	pro, ok := program.(Program)
	if !ok {
		log.Errorf("the program type fail,[type: %v]", reflect.TypeOf(program))
		return nil, false
	}
	return pro, exist
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

func (hub *pluginHub) GetPluginRunner(name string) (*PluginRunner, bool) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	return runner, exist
}

func (hub *pluginHub) DeletePluginRunner(name string) *PluginRunner {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	if exist {
		if runner.running {
			runner.Stop()
		}
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
		runner.ReStart(runner.cacheConf)
	}
}
