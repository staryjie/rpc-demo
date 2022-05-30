package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/staryjie/rpc-demo/pbrpc/codec/server"
	"github.com/staryjie/rpc-demo/pbrpc/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

// service handler
type HelloService struct {
}

// request: 请求
// response 响应
// request  -->   name
// response <--  hello, name
func (s *HelloService) Hello(request *service.Request, response *service.Response) error {
	response.Value = fmt.Sprintf("hello, %s", request)
	return nil
}

// main 里面编写Server
func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})

	// 准备socket
	// 然后我们建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 获取连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// server端采用json来进行编解码, 类似于json.Unmarshal和Marshal
		go rpc.ServeCodec(server.NewServerCodec(conn))
	}

}
