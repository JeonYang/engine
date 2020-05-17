package common

type PluginRunner struct {
	PluginProgram
	cacheConf interface{}
	running   bool
}

func (runner *PluginRunner) Upgrade(program PluginProgram) {
	if runner.running {
		runner.PluginProgram.Stop()
		program.Start(runner.cacheConf)
	}
	runner.PluginProgram = program
}

func (runner *PluginRunner) Start(conf interface{}) {
	if runner.running {
		runner.PluginProgram.Stop()
	}
	runner.PluginProgram.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) ReStart(conf interface{}) {
	if runner.running {
		runner.PluginProgram.Stop()
	}
	if conf == nil {
		conf = runner.cacheConf
	}
	runner.PluginProgram.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) Stop() {
	if runner.running {
		runner.PluginProgram.Stop()
	}
}

func (runner *PluginRunner) Conf() interface{} {
	return runner.cacheConf
}
