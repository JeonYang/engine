package routers

import (
	"context"
	"engine/proto"
)

type EngineServer struct {
	proto.EngineServer
}

func (engin *EngineServer) Command(context context.Context, command *proto.CommandRequest) (response *proto.BasicResponse, err error) {
	switch command.Command {
	case "restart":
	case "stop":
	default:
		response = &proto.BasicResponse{Code: 0, Stderr: "fail command,command mast in [restart,stop]."}
	}
	return response, nil
}
func (engin *EngineServer) Upgrade(context.Context, *proto.UpgradeRequest) (*proto.BasicResponse, error) {
	return nil, nil
}
