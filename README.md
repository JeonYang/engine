
## 思想
1。主程序维护插件为更新与执行。
2。所有业务都动态加载在插件中。

这里的主程序采用的是grpc。当然也可以使用其他方式，比如：tcp与manager连接。

### grpc
protoc --go_out=plugins=grpc:. proto/*.proto

### plugin

go build -buildmode=plugin -o data/plugin/plugin1.so examples/plugin1.go

