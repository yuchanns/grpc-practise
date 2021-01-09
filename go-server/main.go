package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/yuchanns/grpc-practise/proto/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}
	srv := grpc.NewServer()

	greeterServer := &GreeterServer{}
	greeter.RegisterGreeterServer(srv, greeterServer)

	reflection.Register(srv)

	log.Println("start at :9090")

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
