package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	"github.com/adamiklukasz/snippets-go/communication/grpc-gateway/hello"
	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &hello.HelloResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	fmt.Printf("err=%#v\n", err)

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		s.Serve(lis)
	}()

	conn, err := grpc.DialContext(context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Printf("err=%#v\n", err)
	gwmux := runtime.NewServeMux()

	err = hello.RegisterHelloServiceHandler(context.Background(), gwmux, conn)
	fmt.Printf("err=%#v\n", err)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	gwServer.ListenAndServe()

}
