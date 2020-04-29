package routers

import (
	"context"
	"engine/model"
	"engine/proto"
)

type PluginServer struct {
	proto.PluginServer
}

func (plugin *PluginServer) Plugin(context.Context, *proto.PluginInfo) (*proto.PluginInfo, error) {

	return nil, nil
}
func (plugin *PluginServer) Remove(context.Context, *proto.PluginInfo) (*proto.BasicResponse, error) {
	return nil, nil
}
func (plugin *PluginServer) Start(context context.Context, pluginInfo *proto.PluginInfo) (*proto.BasicResponse, error) {
	model.StartPlugin(&model.PluginConf{Name: pluginInfo.Name, Conf: pluginInfo.Conf})
	return nil, nil
}
func (plugin *PluginServer) ReStart(context context.Context, pluginInfo *proto.PluginInfo) (*proto.BasicResponse, error) {
	model.RestartPlugin(&model.PluginConf{Name: pluginInfo.Name, Conf: pluginInfo.Conf})
	return nil, nil
}
func (plugin *PluginServer) Stop(context context.Context, pluginInfo *proto.PluginInfo) (*proto.BasicResponse, error) {
	model.StopPlugin(pluginInfo.Name)
	return nil, nil
}
