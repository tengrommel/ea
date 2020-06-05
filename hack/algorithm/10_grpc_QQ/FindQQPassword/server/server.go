package main

import (
	"context"
	FindQQPassword "ea/hack/algorithm/10_grpc_QQ/FindQQPassword/protoc"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const Address = "127.0.0.1:8848"

type helloService struct{}

func (h helloService) FindQQPassword(ctx context.Context, request *FindQQPassword.QQRequest) (*FindQQPassword.QQResponse, error) {
	resp := new(FindQQPassword.QQResponse)
	resp.QQPassword = fmt.Sprintf("你好%s", request.QQNum)
	return resp, nil
}

var HelloService = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	FindQQPassword.RegisterFindQQPasswordServiceServer(s, HelloService)
	s.Serve(listen)
}
