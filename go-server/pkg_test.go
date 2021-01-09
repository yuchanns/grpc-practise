package main

import (
	"context"
	"testing"

	"github.com/yuchanns/grpc-practise/proto/greeter"
	"google.golang.org/grpc"
)

func TestSayHello(t *testing.T) {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial grpc: %s\n", err)
	}

	c := greeter.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &greeter.HelloRequest{
		Name: "yuchanns",
	})

	if err != nil {
		t.Fatalf("failed to SayHello: %s\n", err)
	}

	t.Logf("get response from SayHello: %s\n", r.Msg)
}
