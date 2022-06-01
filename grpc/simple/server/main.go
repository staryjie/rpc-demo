package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/staryjie/rpc-demo/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
// type HelloServiceServer interface {
// 	Hello(context.Context, *Request) (*Response, error)
// 	Channel(HelloService_ChannelServer) error
// 	mustEmbedUnimplementedHelloServiceServer()
// }

// HelloServiceServer must be embedded to have forward compatible implementations.
type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Value: fmt.Sprintf("Hello, %s", req.Value),
	}, nil
}
func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	// 这个for循环只请求单个客户端请求，gRPC框架会为每个客户端分配一个goroutie
	for {
		// 接收请求
		req, err := stream.Recv()
		if err != nil {
			// err如果是io.EOF表示当前客户端关闭
			if err == io.EOF {
				log.Printf("Client Closed!")
			} else {
				log.Printf("Recv error, %s", err)
				return nil
			}
			return err
		}

		// 响应请求
		resp := &pb.Response{Value: fmt.Sprintf("Hello, %s", req.Value)}

		err = stream.Send(resp)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client Closed!")
				return nil
			}
			return err
		}
	}
}

func main() {
	// s grpc.ServiceRegistry  gRPC Server
	// srv HelloService        实现类

	server := grpc.NewServer()

	// 将HelloService这个实现类注册到gRPC server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	// 获取一个监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	log.Printf("gRPC listen addr: 127.0.0.1:1234")
	// 监听socket，HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
