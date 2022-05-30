package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/staryjie/rpc-demo/pbrpc/codec/client"
	"github.com/staryjie/rpc-demo/pbrpc/service"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 1. 建立socket连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端采用Json 格式来编解码
	client := rpc.NewClientWithCodec(client.NewClientCodec(conn))
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request *service.Request, response *service.Response) error {
	// 执行RPC方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var resp service.Response
	if err := c.Hello(&service.Request{Value: "Jack"}, &resp); err != nil {
		panic(err)
	}

	fmt.Println(resp.Value)
}
