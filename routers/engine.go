package routers

import (
	"context"
	"engine/model"
	"engine/proto"
	"os"
)

type EngineServer struct {
	proto.EngineServer
}

func (engine *EngineServer) Command(context context.Context, command *proto.CommandRequest) (response *proto.BasicResponse, err error) {
	response = &proto.BasicResponse{Message: "success"}
	switch command.Command {
	case "restart":
		model.PluginHub.RestartAllPlugin()
	case "stop":
		model.PluginHub.StopAllPlugin()
		os.Exit(0)
	default:
		response.Code = 1
		response.Message = "fail command,command mast in [restart,stop]."
	}
	return response, nil
}

func (engine *EngineServer) Upgrade(context context.Context, download *proto.Download) (*proto.ProgramInfo, error) {
	response := &proto.ProgramInfo{Message: "success"}
	err := model.PluginLoader.DownloadEngine(download.Url, download.Md5, download.Version)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
	}
	return response, nil
}
