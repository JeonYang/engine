package routers

import (
	"engine/proto"
	"google.golang.org/grpc"
)

func Server() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterEngineServer(s, &EngineServer{})
	proto.RegisterPluginServer(s, &PluginServer{})
	return s
}
