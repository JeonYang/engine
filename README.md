# engine

protoc --go_out=plugins=grpc:. proto/*.proto

pluginInterface
Version()[]int
Name()string
Stop()error
Start()error
NotifyConf(conf *Conf)
Notify( map[string]interface)

conf: {
"command":"stop/restart",
"pluginVersion":[{"name":"name","version":[1,2,0],"url":""}],
"pluginConf":"{"pluginName":{"status":"start","conf":{}}},
"":"",
}

功能
1. 更新配置
2. 根据配置下载插件
2. 根据配置对主程序进行重启/停止,