package main

import (
	"context"
	"fmt"
	pb "github.com/jergoo/go-grpc-example/proto/hello"
	"google.golang.org/grpc"
	"net"
)

const Address = "127.0.0.1:8848"

type helloService struct {
}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("你好%s", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, HelloService)
	s.Serve(listen)
}
