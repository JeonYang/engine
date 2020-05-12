package main

import (
	"engine/conf"
	"engine/log"
	"engine/routers"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	confFile *string
)

func main() {
	confFile = flag.String("f", "", "input conf file")
	err := conf.InitConf(*confFile)
	if err != nil {
		panic(err)
	}
	err = log.InitLog()
	if err != nil {
		panic(err)
	}

	panicFile, err := os.OpenFile(conf.EngineConf.PanicFile(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		log.Warnf("open %s fail: %v", conf.EngineConf.PanicFile(), err)
		return
	}
	panicFile.WriteString(fmt.Sprintf("\n%v opened panic.log at %v\n", os.Getpid(), time.Now()))
	os.Stderr = panicFile

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.EngineConf.RpcPort))
	if err != nil {
		log.Errorf("failed to listen,port: %d,err: %v", conf.EngineConf.RpcPort, err)
		return
	}
	routers.Server().Serve(lis)
}
