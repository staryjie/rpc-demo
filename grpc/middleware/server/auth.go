package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderAccessKey = "client-id"
	ClientHeaderSecretKey = "client-secret"
)

func NewClientCredential(cak, csk string) metadata.MD {
	return metadata.MD{
		ClientHeaderAccessKey: []string{cak},
		ClientHeaderSecretKey: []string{csk},
	}
}

type GrpcAuther struct {
}

func NewAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&GrpcAuther{}).UnaryServerInterceptor
}

func NewAuthStreamServerInterception() grpc.StreamServerInterceptor {
	return (&GrpcAuther{}).StreamServerInterception
}

func (a *GrpcAuther) UnaryServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 1.读取凭证，凭证放在meta信息中[类似http2 header的结构]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context!")
	}
	fmt.Printf("gRPC header info: %s", md)

	// 从metadata中获取客户端传递过来的凭证
	clientId, clientSecret := a.getClientCredentialsFromMeta(md)

	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	// 认证通过之后，请求继续往后传递
	return handler(ctx, req)
}

// Stream rpc interceptor
func (a *GrpcAuther) StreamServerInterception(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// 读取凭证，凭证放在meta信息中
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return fmt.Errorf("ctx is not an grpc incoming context!")
	}
	fmt.Printf("gRPC header info: %s", md)

	// 从metadata中获取客户端传递过来的凭证
	clientId, clientSecret := a.getClientCredentialsFromMeta(md)

	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return err
	}

	// 认证通过之后，请求继续往后传递
	return handler(srv, ss)
}

// 从metada中获取信息
func (a *GrpcAuther) getClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cakList := md[ClientHeaderAccessKey]
	if len(cakList) > 0 {
		clientId = cakList[0]
	}
	cskList := md[ClientHeaderSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}

	return
}

func (a *GrpcAuther) validateServiceCredential(clientId, clientSecret string) error {

	if !(clientId == "admin" && clientSecret == "123456") {
		// 认证错误，返回错误
		return status.Errorf(codes.Unauthenticated, "clientId or clientSecret not correct!")
	}
	return nil
}
