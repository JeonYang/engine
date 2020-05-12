package model

import (
	"fmt"
	"sync"
)

var PluginHub = pluginHub{runnerHub: make(map[string]*PluginRunner)}

type pluginHub struct {
	mutex     sync.Mutex
	runnerHub map[string]*PluginRunner
}

func (hub *pluginHub) GetPluginProgram(name string) (PluginProgram, bool) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	return runner.pluginProgram, exist
}

func (hub *pluginHub) PushPluginProgram(pluginProgram PluginProgram) error {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	_, exist := hub.runnerHub[pluginProgram.Name()]
	if exist {
		return fmt.Errorf("the plugin: %s exist in hub.", pluginProgram.Name())
	}
	hub.runnerHub[pluginProgram.Name()] = &PluginRunner{pluginProgram: pluginProgram}
	return nil
}

func (hub *pluginHub) GetPluginRunner(name string) (*PluginRunner, bool) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	return runner, exist
}

func (hub *pluginHub) DeletePluginRunner(name string) (*PluginRunner) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	runner, exist := hub.runnerHub[name]
	if exist {
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
