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

func (plugin *PluginServer) Download(context.Context, *proto.Download) (*proto.BasicResponse, error) {

	return nil, nil
}

func (plugin *PluginServer) Upgrade(context context.Context, download *proto.Download) (response *proto.ProgramInfo, err error) {
	log.Debugf("upgrade plugin %+v ...", download)
	response = &proto.ProgramInfo{Message: "success"}
	pluginProgram, err := model.PluginLoader.DownloadPlugin(download.Name, download.Url, download.Md5, download.Version)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		log.Debugf("upgrade plugin %s fail,response: %v", download.Name, response)
		return
	}
	runner, exist := model.PluginHub.GetPluginRunner(pluginProgram.Name())
	if exist {
		runner.Upgrade(pluginProgram)
		return
	}
	err = model.PluginHub.PushPluginProgram(pluginProgram)
	if err != nil {
		response.Code = 2
		response.Message = err.Error()
		log.Debugf("upgrade plugin %s fail,response: %v", download.Name, response)
	}
	return
}

func (plugin *PluginServer) Plugin(context context.Context, pluginConf *proto.PluginConf) (pluginInfo *proto.ProgramInfo, err error) {
	//pluginInfo = &proto.PluginInfo{}
	//pluginProgram, exist := model.PluginHub.GetPluginProgram(pluginConf.Name)
	//if !exist {
	//	pluginInfo.
	//}
	return pluginInfo, nil
}

func (plugin *PluginServer) Remove(context context.Context, pluginConf *proto.PluginConf) (*proto.ProgramInfo, error) {
	return nil, nil
}

func (plugin *PluginServer) Start(context context.Context, pluginConf *proto.PluginConf) (*proto.ProgramInfo, error) {
	//model.StartPlugin(&model.PluginConf{Name: pluginInfo.Name, Conf: pluginInfo.Conf})
	return nil, nil
}

func (plugin *PluginServer) ReStart(context context.Context, pluginConf *proto.PluginConf) (*proto.ProgramInfo, error) {
	//model.RestartPlugin(&model.PluginConf{Name: pluginInfo.Name, Conf: pluginInfo.Conf})
	return nil, nil
}

func (plugin *PluginServer) Stop(context context.Context, pluginConf *proto.PluginConf) (*proto.ProgramInfo, error) {
	//model.StopPlugin(pluginInfo.Name)
	return nil, nil
}
