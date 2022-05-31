# Hello gRPC


###  通过hello.proto生成Go代码

```bash
cd rpc-demo/grpc/simple/server/
protoc -I=./pb/ --go_out=. --go_opt=module="github.com/staryjie/rpc-demo/grpc/simple/server" ./pb/hello.proto
```


### rpc接口定义的protobuf代码生成
```bash
# 安装Go官方推荐的grpc代码生成插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 代码生成
cd rpc-demo/grpc/simple/server/
protoc -I=. --go_out=. --go_opt=module="github.com/staryjie/rpc-demo/grpc/simple/server" --go-grpc_out=. --go-grpc_opt=module="github.com/staryjie/rpc-demo/grpc/simple/
server" pb/hello.proto
```
