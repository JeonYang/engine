package model

type PluginRunner struct {
	Program
	cacheConf string
	running   bool
}

func (runner *PluginRunner) Upgrade(program Program) {
	if runner.running {
		runner.Program.Stop()
		program.Start(runner.cacheConf)
	}
	runner.Program = program
}

func (runner *PluginRunner) Start(conf string) {
	if runner.running {
		runner.Program.Stop()
	}
	runner.Program.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) ReStart(conf string) {
	if runner.running {
		runner.Program.Stop()
	}
	if conf == "" {
		conf = runner.cacheConf
	}
	runner.Program.Start(conf)
	runner.cacheConf = conf
}

func (runner *PluginRunner) Stop() {
	if runner.running {
		runner.Program.Stop()
	}
}

func (runner *PluginRunner) Conf() string {
	return runner.cacheConf
}
