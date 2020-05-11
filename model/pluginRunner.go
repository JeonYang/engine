package model

type PluginRunner struct {
	pluginProgram PluginProgram
	cacheConf     string
	running       bool
}

func (runner *PluginRunner) Upgrade(pluginProgram PluginProgram) {
	if runner.running {
		runner.pluginProgram.Stop()
		pluginProgram.Start(runner.cacheConf)
	}
	runner.pluginProgram = pluginProgram
}

func (runner *PluginRunner) Start(conf string) {
	if runner.running {
		runner.pluginProgram.Stop()
	}
	runner.pluginProgram.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) ReStart(conf string) {
	if runner.running {
		runner.pluginProgram.Stop()
	}
	runner.pluginProgram.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) Stop() {
	if runner.running {
		runner.pluginProgram.Stop()
	}
}
