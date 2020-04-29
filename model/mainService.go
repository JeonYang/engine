package model

var pluginScheduler = &PluginScheduler{}
var pluginHub = make(map[string]PluginFunction)

//var MainService
//
//type MainService struct {
//}

func NotifyPluginFunc(plugin PluginFunction) {
	pluginScheduler.NotifyPluginFunc(plugin)
}

func StartPlugin(pluginConf *PluginConf) {
	pluginScheduler.StartPlugin(pluginConf)
}

func RestartPlugin(pluginConf *PluginConf) {
	pluginScheduler.RestartPlugin(pluginConf)
}

func StopPlugin(pluginName string) {
	pluginScheduler.StopPlugin(pluginName)
}
