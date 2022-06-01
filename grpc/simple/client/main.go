package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/staryjie/rpc-demo/grpc/middleware/server"
	"github.com/staryjie/rpc-demo/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 通过grpc建立连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// gRPC我我们生成一个客户端调用服务端的SDK
	client := pb.NewHelloServiceClient(conn)

	// req <--> resp
	crendential := server.NewClientCredential("admin", "123456")

	ctx := metadata.NewOutgoingContext(context.Background(), crendential)

	resp, err := client.Hello(ctx, &pb.Request{Value: "Jack From method Hello"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Value)

	// stream
	stream, err := client.Channel(ctx)
	if err != nil {
		panic(err)
	}

	reqCount := 0
	// 启用一个Goroutine来发送请求
	go func() {
		for {
			recover()
			reqStr := "For Jack" + strconv.Itoa(reqCount+1)
			err := stream.Send(&pb.Request{Value: reqStr})
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			reqCount += 1
		}
	}()

	// 主循环 负责接收服务端响应
	for {
		resp, err = stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Value)
	}
}
