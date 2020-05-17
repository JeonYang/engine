package routers

import (
	"context"
	"engine/log"
	"engine/model"
	"engine/proto"
)

type PluginServer struct {
	proto.PluginServer
}

func (plugin *PluginServer) Download(context context.Context, download *proto.Download) (basicResponse *proto.BasicResponse, _ error) {
	log.Infof("download plugin %s ...", download.Name)
	basicResponse = &proto.BasicResponse{Message: "success"}
	_, err := model.PluginLoader.DownloadPlugin(download.Name, download.Url, download.Md5)
	if err != nil {
		basicResponse.Code = 1
		basicResponse.Message = "download fail"
		log.Errorf("download plugin %s fail,response: %s", download.Name, err.Error())
	}
	log.Infof("download plugin %s success", download.Name)
	return
}

func (plugin *PluginServer) LoadPlugin(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("load plugin %s ...", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success"}
	program, err := model.PluginLoader.LoadPluginProgram(pluginConf.Name)
	if err != nil {
		programInfo.Code = 1
		programInfo.Message = "load fail"
		log.Errorf("load plugin %s fail,err: %s", pluginConf.Name, err.Error())
		return
	}

	runner, exist := model.PluginHub.GetPluginRunner(program.Name())
	if exist {
		runner.Upgrade(program)
		programInfo.Name = runner.Name()
		programInfo.Version = runner.Version()
		//programInfo.Md5 = runner.Md5()
		//programInfo.Conf = runner.Conf()
		log.Infof("load plugin %s ,info: %+v", pluginConf.Name, programInfo)
		return
	}
	err = model.PluginHub.PushPluginProgram(program)
	if err != nil {
		programInfo.Code = 2
		programInfo.Message = "push fast..."
		log.Errorf("load plugin %s fail,err: %v", program.Name(), err.Error())
	}
	return
}

func (plugin *PluginServer) Upgrade(context context.Context, download *proto.Download) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("upgrade plugin %s ...", download.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: download.Name, Version: download.Version}
	program, err := model.PluginLoader.DownloadPlugin(download.Name, download.Url, download.Md5)
	if err != nil {
		programInfo.Code = 1
		programInfo.Message = "download fail"
		log.Errorf("upgrade plugin %s fail,err: %s", download.Name, err.Error())
		return
	}
	runner, exist := model.PluginHub.GetPluginRunner(program.Name())
	if exist {
		runner.Upgrade(program)
		programInfo.Name = runner.Name()
		programInfo.Version = runner.Version()
		//programInfo.Md5 = runner.Md5()
		//programInfo.Conf = runner.Conf()
		log.Infof("upgrade plugin %s success", download.Name)
		return
	}
	err = model.PluginHub.PushPluginProgram(program)
	if err != nil {
		programInfo.Code = 2
		programInfo.Message = "upgrade fast..."
		log.Errorf("upgrade plugin %s fail,err: %v", download.Name, err.Error())
	}
	return
}

func (plugin *PluginServer) Plugin(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("get plugin %s ...", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: pluginConf.Name}
	runner, exist := model.PluginHub.GetPluginRunner(pluginConf.Name)
	if !exist {
		programInfo.Code = 1
		programInfo.Message = "plugin not fond"
		log.Errorf("get plugin %s fail,plugin not fond", pluginConf.Name)
		return
	}
	programInfo.Name = runner.Name()
	programInfo.Version = runner.Version()
	//programInfo.Md5 = runner.Md5()
	//programInfo.Conf = runner.Conf()
	log.Infof("get plugin %s success,plugin: %+v", runner.Name(), programInfo)
	return
}

func (plugin *PluginServer) Remove(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("remove plugin %s ...", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: pluginConf.Name}
	runner, exist := model.PluginHub.GetPluginRunner(pluginConf.Name)
	if !exist {
		return
	}
	if err := model.PluginLoader.Delete(runner.Name()); err != nil {
		programInfo.Code = 1
		programInfo.Message = "remove plugin file err"
		log.Errorf("remove plugin err: %v", err)
		return
	}
	model.PluginHub.DeletePluginRunner(runner.Name())
	programInfo.Name = runner.Name()
	programInfo.Version = runner.Version()
	//programInfo.Md5 = runner.Md5()
	//programInfo.Conf = runner.Conf()
	log.Infof("remove plugin %s success.plugin: %+v", pluginConf.Name, programInfo)
	return
}

func (plugin *PluginServer) Start(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("start plugin: %s", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: pluginConf.Name}
	runner, exist := model.PluginHub.GetPluginRunner(pluginConf.Name)
	if !exist {
		programInfo.Code = 1
		programInfo.Message = "plugin not fond"
		log.Errorf("start plugin: %s fail,err: %s", pluginConf.Name, programInfo.Message)
		return
	}
	runner.Start(pluginConf.Conf)
	programInfo.Name = runner.Name()
	programInfo.Version = runner.Version()
	//programInfo.Md5 = runner.Md5()
	//programInfo.Conf = runner.Conf()
	log.Infof("start plugin: %s success", runner.Name())
	return
}

func (plugin *PluginServer) ReStart(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("restart plugin: %s", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: pluginConf.Name}
	runner, exist := model.PluginHub.GetPluginRunner(pluginConf.Name)
	if !exist {
		programInfo.Code = 1
		programInfo.Message = "plugin not fond"
		log.Infof("restart plugin: %s fail: %s", pluginConf.Name, programInfo.Message)
		return
	}
	runner.ReStart(pluginConf.Conf)
	programInfo.Name = runner.Name()
	programInfo.Version = runner.Version()
	//programInfo.Md5 = runner.Md5()
	//programInfo.Conf = runner.Conf()
	log.Infof("restart plugin: %s success", pluginConf.Name)
	return
}

func (plugin *PluginServer) Stop(context context.Context, pluginConf *proto.PluginConf) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("stop plugin: %s ", pluginConf.Name)
	programInfo = &proto.ProgramInfo{Message: "success", Name: pluginConf.Name}
	runner, exist := model.PluginHub.GetPluginRunner(pluginConf.Name)
	if !exist {
		programInfo.Code = 1
		programInfo.Message = "plugin not fond"
		log.Infof("stop plugin : %s fail: %s", pluginConf.Name, programInfo.Message)
		return
	}
	runner.Stop()
	programInfo.Name = runner.Name()
	programInfo.Version = runner.Version()
	//programInfo.Md5 = runner.Md5()
	//programInfo.Conf = runner.Conf()
	log.Infof("stop plugin: %s success", pluginConf.Name)
	return
}
