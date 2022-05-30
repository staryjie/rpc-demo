package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/staryjie/rpc-demo/json_tcp/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

// Server handler
type HelloService struct {
}

// 由于我们是一个rpc服务, 因此参数上面还是有约束：
// 1.request：请求
// 2.response：响应
// request --> name
// response --> hello name
func (p *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("Hello, %s", request)
	return nil
}

// main函数中实现Server端
func main() {
	// 把HelloService注册为一个RPC的receiver
	// 其中rpc.Registry函数会调用对象类型中所有满足RPC规则的对象方法注册为RPC函数
	//所有注册的方法都会放在`HelloService`服务空间之下
	rpc.RegisterName("HelloService", &HelloService{})

	// 然后建立一个唯一的TCP连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}

	// 通过rpc.ServeConn函数在该TCP连接上为对方提供服务
	// 每Accept一个请求，就创建一个goroutie进行处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error", err)
		}
		// 请求连接成功建立之后，开启一个goroutie单独处理该请求
		// go rpc.ServeConn(conn)

		// Server端采用json来编解码
		// 类似于json的josn.Unmarshal和json.Marshal
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
