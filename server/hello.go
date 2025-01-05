package server

import (
	"context"
	"gomod.pri/helloGrpc/pb/xcommon"
	"gomod.pri/helloGrpc/pb/xinternal/hello"
	"log"
)

func newHelloServer() hello.HelloServiceServer {
	return new(helloServiceServer)
}

type helloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

// SayHello 实现了 HelloServiceServer 接口中的 SayHello 方法
func (s *helloServiceServer) SayHello(ctx context.Context, in *hello.SayHelloRequest) (*hello.SayHelloResponse, error) {
	log.Printf("hello %s, your age is %d", in.GetMsg().GetName(), in.GetMsg().GetAge())
	return &hello.SayHelloResponse{Status: xcommon.Status_STATUS_NORMAL}, nil
}
