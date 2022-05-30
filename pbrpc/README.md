# 代码生成



```sh
# 生成 serivce pb编译文件
$ protoc -I=./service/pb/ --go_out=./service/ --go_opt=module="gitee.com/go-course/rpc-g7/pbrpc/service"  hello.proto

# 生成 codec的protobuf编译文件
protoc -I=./codec/pb/ --go_out=./codec/pb/ --go_opt=module="gitee.com/go-course/rpc-g7/pbrpc/codec/pb"  header.proto
```