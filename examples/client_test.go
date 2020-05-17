package main

import (
	"context"
	"engine/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

var pluginClient proto.PluginClient
var engineClient proto.EngineClient


func init() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	pluginClient = proto.NewPluginClient(conn)
	engineClient = proto.NewEngineClient(conn)
}

func TestLoadPlugin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	childConf:=`{"child1":"child1","child2":"child2"}`
	conf := &proto.PluginConf{Name: "pluginMain",Conf: childConf}
	r, err := pluginClient.LoadPlugin(ctx, conf)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %+v", r)
}

func TestStart(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	childConf:=`{"child1":"child1","child2":"child2"}`
	conf := &proto.PluginConf{Name: "pluginMain",Conf: childConf}
	r, err := pluginClient.Start(ctx, conf)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %+v", r)
}
