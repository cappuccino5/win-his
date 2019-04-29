微服务测试用例
# 目录说明
```
win-his
│
└───api          // micro**.proto
│   │  gateway       // http API .proto和api的*.proto方法一样
│   │──rpc    // 公用部分
│   │   │client       // client
│   │   │gw       // http
│   │   │server    // server
│   │  server      // 服务具体实现
│   │   vendor         // 第三方包管理
│   │	README

```
# 这是一个micro 服务测试用例
```

一、编译*.proto,生成micro
win-his/api/*.proto:
	protoc --go_out=plugins=grpc:. *.proto

生成micro,**.micro.go
	protoc  --proto_path=. --micro_out=. --go_out=. *.proto
```
# 二、编译*.proto,生成gateway
```
win-his/gateway/*.proto:
 生成 *.pb.go
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. *.proto

 生成*.pb.gw.go	
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. *.proto

 生成 .swagger.json
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. *.proto
```
# 三、运行测试用例:
```
start:
 启动consul:
    consul agent -dev
 micro默认支持consul：
    go run main.go
 查看启动的微服务：
    micro list services
 也可以注册服务到consul并指定端口:
    go run main.go  --registry=consul --server_address=localhost:9000
 服务注册到consul的前提下，使用游览器可以查看到consul中服务:
     http://127.0.0.1:8500
```

# END
```
测试done时出现以下情况,附解决方法
1、// go run main.go ： const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package
 谢谢@lenvs给了我灵感！我也解决了这个问题！
你可以这样做。
rm -rf vendor/github.com/golang/protobuf
cd vendor/github.com/golang/
git clone https://github.com/golang/protobuf.git
go install
2、生成gw包的时候需要加路径依赖:
	a. -I 需要protoc的 include/
	b. -I 需要github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis 依赖

```