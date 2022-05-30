package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/staryjie/rpc-demo/json_http/service"
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

// 求两数之和
func (p *HelloService) Calc(request *service.CalcRequest, response *int) error {
	*response = request.A + request.B
	return nil
}

type RPCReadWriterCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriterCloser(w http.ResponseWriter, r *http.Request) *RPCReadWriterCloser {
	return &RPCReadWriterCloser{
		w,
		r.Body,
	}
}

// main函数中实现Server端
func main() {
	// 把HelloService注册为一个RPC的receiver
	// 其中rpc.Registry函数会调用对象类型中所有满足RPC规则的对象方法注册为RPC函数
	//所有注册的方法都会放在`HelloService`服务空间之下
	rpc.RegisterName("HelloService", &HelloService{})

	// 通过 /jsonrpc 处理rpc请求
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriterCloser(w, r)))
	})

	// 通过HTTP协议接受rpc请求
	http.ListenAndServe(":1234", nil)
}
