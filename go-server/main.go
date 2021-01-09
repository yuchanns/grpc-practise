package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/yuchanns/grpc-practise/proto/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	endpoint := ":9090"
	addr := ":8080"
	// grpc
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}
	srv := grpc.NewServer()

	greeterServer := &GreeterServer{}
	greeter.RegisterGreeterServer(srv, greeterServer)

	reflection.Register(srv)

	// grpc-gateway
	mux := runtime.NewServeMux()
	greeter.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{
		grpc.WithInsecure(),
	})

	log.Printf("grpc-server start at %s and grpc-gateway start at %s\n", endpoint, addr)

	go func() {
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatalf("failed to start grpc gateway: %+v", err)
		}
	}()

	if err := srv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// GreeterServer implements greeter.GreeterServer
type GreeterServer struct{}

// SayHello returns a grpc response
func (s *GreeterServer) SayHello(c context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Msg: fmt.Sprintf("hello, %s", req.Name),
	}, nil
}
