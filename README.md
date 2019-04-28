微服务测试用例

micro 服务:
生成pb：pb.go
	protoc --go_out=plugins=grpc:. *.proto

生成micro,**.micro.go
	protoc  --proto_path=. --micro_out=. --go_out=. *.proto
 

api gateway: 
 生成 *.pb.go
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. *.proto

 生成*.pb.gw.go	
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. *.proto

 生成 .swagger.json
	protoc -IC:/Go/pkg/protoc-3.7.1-win64/include/ -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. *.proto

start:	
 注册微服务到consul并指定端口:
	go run main.go  --registry=consul --server_address=localhost:9000
 
 查看启动的微服务：
	micro list services
 
END:
测试done时出现以下情况,附解决方法
// go run main.go ： const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package
 
 谢谢@lenvs给了我灵感！我也解决了这个问题！
你可以这样做。

rm -rf vendor/github.com/golang/protobuf
cd vendor/github.com/golang/
git clone https://github.com/golang/protobuf.git
go install
