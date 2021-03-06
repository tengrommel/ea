package main

import (
	"context"
	"ea/hack/algorithm/9_distributed_store/QQ/QQ"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

const Address = "127.0.0.1:8848"

type helloService struct {
}

var HelloService = helloService{}

func (h helloService) GetQQPassWord(ctx context.Context, in *QQ.QQRequest) (*QQ.QQResponse, error) {
	resp := new(QQ.QQResponse)
	resp.QQpassword = fmt.Sprintf("你好%s", in.QQname)
	for i := 0; i < 10; i++ {
		resp.Password = append(resp.Password, strconv.Itoa(i))
	}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	QQ.RegisterQQServer(s, HelloService)
	s.Serve(listen)
}
