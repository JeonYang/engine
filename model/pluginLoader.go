package model

import (
	"engine/common"
	"engine/conf"
	"engine/util"
	"fmt"
	"os"
	"plugin"
	"reflect"
)

var PluginLoader = &pluginLoader{}

type pluginLoader struct {
}

func (loader *pluginLoader) Delete(pluginName string) error {
	return os.Remove(conf.EngineConf.PluginPath(pluginName))
}

func (loader *pluginLoader) DownloadPlugin(pluginName, url, md5 string) (program Program, err error) {
	md5Str, err := util.WGet(url, conf.EngineConf.PluginPath(pluginName))
	if err != nil {
		return nil, err
	}
	if md5Str != md5 {
		err = fmt.Errorf("request md5: %s,plugin md5: %s.", md5, md5Str)
		return
	}
	return loader.LoadPluginProgram(pluginName)
}

func (loader *pluginLoader) DownloadEngine(url string) (string, error) {
	return util.WGet(url, conf.EngineConf.EngineDownLoadPath())
}

// 使用插件中的 New+PluginName function初始化
func (loader *pluginLoader) LoadPluginProgram(pluginName string) (p Program, err error) {
	plugin, err := plugin.Open(fmt.Sprintf("%s.so", conf.EngineConf.PluginPath(pluginName)))
	if err != nil {
		err = fmt.Errorf("open plugin: %s fail,err: %v", pluginName, err)
		return
	}
	pluginName = strFirstToUpper(pluginName)
	pluginBuildFun, err := plugin.Lookup(fmt.Sprintf("New%s", pluginName))
	if err != nil {
		err = fmt.Errorf("lookup plugin: %s fail,err: %v", pluginName, err)
		return
	}
	pluginProgramBuild, ok := pluginBuildFun.(func(common.PluginProgramEnvironment) common.PluginProgram)
	if !ok {
		err = fmt.Errorf("the pluginBuildFun type: %s fail.", reflect.TypeOf(pluginBuildFun).String())
		return
	}
	pluginProgram := pluginProgramBuild(PluginProgramEnvironment)
	return &program{PluginProgram: pluginProgram}, nil
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	list := []rune(str)
	if list[0] >= 97 && list[0] <= 122 {
		list[0] -= 32
	}
	return string(list)
}
