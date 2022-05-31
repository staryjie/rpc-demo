package main

import (
	"context"
	"fmt"

	"github.com/staryjie/rpc-demo/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {
	// 通过grpc建立连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// gRPC我我们生成一个客户端调用服务端的SDK
	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "Jack"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Value)
}
