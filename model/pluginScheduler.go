package model

type PluginScheduler struct {
	pluginFuncMap           map[string]*PluginFunc
	pluginFuncNotifyChan    chan PluginFunction
	pluginStartNotifyChan   chan *PluginConf
	pluginRestartNotifyChan chan *PluginConf
	pluginStopNotifyChan    chan string
}

func (scheduler *PluginScheduler) Start(pluginFuncList []PluginFunction) {
	scheduler.pluginFuncMap = make(map[string]*PluginFunc)
	scheduler.pluginFuncNotifyChan = make(chan PluginFunction, 0)
	scheduler.pluginStartNotifyChan = make(chan *PluginConf, 0)
	scheduler.pluginRestartNotifyChan = make(chan *PluginConf, 0)
	scheduler.pluginStopNotifyChan = make(chan string, 0)
	go func() {
		for {
			select {

			case pluginName := <-scheduler.pluginStopNotifyChan:

				pluginFunc, exist := scheduler.pluginFuncMap[pluginName]
				if exist && pluginFunc.Running {
					pluginFunc.PluginFunction.Stop()
				}

			case pluginConf := <-scheduler.pluginStartNotifyChan:

				pluginFunc, exist := scheduler.pluginFuncMap[pluginConf.Name]
				if exist && !pluginFunc.Running {
					pluginFunc.Running = true
					pluginFunc.PluginFunction.Start(pluginConf.Conf)
				}

			case pluginConf := <-scheduler.pluginRestartNotifyChan:

				pluginFunc, exist := scheduler.pluginFuncMap[pluginConf.Name]
				if exist {
					if pluginConf.Conf != "" {
						pluginFunc.CacheConf = pluginConf.Conf
					}
					if pluginFunc.Running {
						pluginFunc.PluginFunction.Stop()
						pluginFunc.PluginFunction.Start(pluginFunc.CacheConf)
					} else {
						pluginFunc.PluginFunction.Start(pluginFunc.CacheConf)
					}
					pluginFunc.Running = true
				}

			case notify := <-scheduler.pluginFuncNotifyChan:

				pluginFunc, exist := scheduler.pluginFuncMap[notify.Name()]
				running := pluginFunc.Running
				if exist && pluginFunc.Running {
					pluginFunc.PluginFunction.Stop()
					notify.Start(pluginFunc.CacheConf)
				} else {
					scheduler.pluginFuncMap[notify.Name()] = &PluginFunc{PluginFunction: notify}
				}
				scheduler.pluginFuncMap[notify.Name()] = &PluginFunc{PluginFunction: notify, Running: running}

			}
		}
	}()
}

func (scheduler *PluginScheduler) NotifyPluginFunc(plugin PluginFunction) {
	scheduler.pluginFuncNotifyChan <- plugin
}

func (scheduler *PluginScheduler) StartPlugin(pluginConf *PluginConf) {
	scheduler.pluginStartNotifyChan <- pluginConf
}

func (scheduler *PluginScheduler) RestartPlugin(pluginConf *PluginConf) {
	scheduler.pluginRestartNotifyChan <- pluginConf
}

func (scheduler *PluginScheduler) StopPlugin(pluginName string) {
	scheduler.pluginStopNotifyChan <- pluginName
}
