package routers

import (
	"context"
	"engine/conf"
	"engine/log"
	"engine/model"
	"engine/proto"
	"engine/util"
	"fmt"
	"os"
)

type EngineServer struct {
	proto.EngineServer
}

func (engine *EngineServer) Command(context context.Context, command *proto.CommandRequest) (response *proto.BasicResponse, err error) {
	log.Infof("engine command: %+v", command)
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
	log.Infof("engine command response: %+v", response)
	return response, nil
}

func (engine *EngineServer) Upgrade(context context.Context, download *proto.Download) (programInfo *proto.ProgramInfo, _ error) {
	log.Infof("engine upgrade: %+v", download)
	programInfo = &proto.ProgramInfo{Message: "success", Name: "engine"}
	md5Str, err := model.PluginLoader.DownloadEngine(download.Url)
	if err != nil {
		log.Errorf("download engine fail,err: %v", err)
		programInfo.Code = 1
		programInfo.Message = "download engine fail"
		return
	}
	if md5Str != download.Md5 {
		programInfo.Code = 2
		programInfo.Message = fmt.Sprintf("request md5: %s,engine md5: %s.", download.Md5, md5Str)
		log.Errorf("compare md5 fail,err: %s", programInfo.Message)
		return
	}
	if util.CompareVersion(conf.Version, download.Version) <= 0 {
		programInfo.Code = 3
		programInfo.Message = fmt.Sprintf("request version: %s,engine version: %s.", download.Version, conf.Version)
		log.Errorf("compare engine fail,err: %s", programInfo.Message)
		return
	}
	log.Infof("start upgrade, new version", download.Version)
	err = model.Engine.Update(conf.EngineConf.EngineAppPath(), conf.EngineConf.EngineBackupPath(), conf.EngineConf.EngineDownLoadPath())
	if err != nil {
		programInfo.Code = 4
		programInfo.Message = fmt.Sprintf("upgrade engine fail")
		log.Errorf("upgrade engine fail,err: %s", programInfo.Message)
		return
	}
	programInfo.Version = download.Version
	programInfo.Md5 = download.Md5
	return programInfo, nil
}
