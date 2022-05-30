package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	// 1.建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dailing error ", err)
	}

	// 2.通过client.Call发起调用
	//   第一个参数：通过.连接RPC服务名称和方法名
	//   第二个参数：请求参数
	//   第三个参数：请求响应，必须是指针类型，由rpc赋值
	var response string
	err = client.Call("HelloService.Hello", "Jack", &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
