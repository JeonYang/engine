package main

import (
	"context"
	"engine/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:9999"
	defaultName = "pb.Add"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewPluginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	download := &proto.Download{Name: "pluginName"}
	r, err := c.Upgrade(ctx, download)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %+v", r)
}
