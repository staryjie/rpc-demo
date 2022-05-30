# protobuf

## 如何通过proto文件生成Go代码

```bash
# 进入到pb目录
cd rpc-demo/protobuf/pb

# 通过protoc结合protoc-gen-go工具生成go代码
protoc -I=. --go_out=. --go_opt=module="github.com/staryjie/rpc-demo/protobuf/pb" ./hello.proto

# 或者
cd rpc-demo/protobuf
protoc -I=./pb/ --go_out=./pb/ --go_opt=module="github.com/staryjie/rpc-demo/protobuf/pb" hello.proto
```

> 说明：
> - `-I=.`：`-IPATH`，`--proto_path=PATH`，表示`proto`文件查找路径，`.`表示当前路径，不指定则默认值为当前目录
> - `--go_out=.`：`--go`指定使用插件的名称，`protoc-gen`是插件的命名规范,`go`是插件名称，所以这里就表示使用`protoc-gen-go`来生成代码,`--go_out`表示生成的代码存放路径，`.`表示当前路径
> - `--go_opt=module=`，`protoc-gen-go`插件的`opt`参数，这里的`module`指定了`go module`，生成的`go pkg`会去掉`go module`路径，这里表示在生成代码的时候去掉指定的字符串路径，`--go_opt=module="github.com/staryjie/rpc-demo/protobuf/pb"`表示在生成代码的时候去掉`github.com/staryjie/rpc-demo/protobuf/pb`这一层路径
> `./hello.proto`：指定proto文件，protoc通过该文件进行生成代码操作
