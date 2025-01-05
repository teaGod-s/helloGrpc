package server

import (
	"context"
	"errors"
	"net"
	"net/http"

	authipb "gomod.pri/helloGrpc/pb/xinternal/auth"
	"gomod.pri/helloGrpc/pb/xinternal/hello"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			grpclog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, newHelloServer())
	authipb.RegisterAuthServiceServer(s, newAuthServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)

	hello.RegisterHelloServiceHandlerServer(ctx, mux, newHelloServer())
	authipb.RegisterAuthServiceHandlerServer(ctx, mux, newAuthServer())

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		grpclog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			grpclog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		grpclog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil

}
