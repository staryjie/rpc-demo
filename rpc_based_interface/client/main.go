package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/staryjie/rpc-demo/rpc_based_interface/service"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	client rpc.Client
}

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 1.建立连接
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: *client,
	}, nil
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	// 2.通过client.Call发起调用
	//   第一个参数：通过.连接RPC服务名称和方法名
	//   第二个参数：请求参数
	//   第三个参数：请求响应，必须是指针类型，由rpc赋值
	// var response string
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), "Jack", &response)
}

func main() {
	var response string
	// 创建客户端
	client, err := NewHelloServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Hello("Jack", &response); err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
