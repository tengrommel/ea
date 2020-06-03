package main

import (
	"context"
	"ea/hack/algorithm/9_distributed_store/QQ/QQ"
	"fmt"
	"google.golang.org/grpc"
)

const Address = "127.0.0.1:8848"

type helloService struct {
}

var HelloService = helloService{}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := QQ.NewQQClient(conn)
	req := &QQ.QQRequest{
		QQnum:  "1343",
		QQname: "你大爷",
	}
	res, err := c.GetQQPassWord(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("收到", res.Password)
	fmt.Println("收到", res.QQpassword)
}
