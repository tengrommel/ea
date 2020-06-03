package main

import (
	"context"
	"fmt"
	pb "github.com/jergoo/go-grpc-example/proto/hello"
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
	c := pb.NewHelloClient(conn)
	req := &pb.HelloRequest{
		Name: "你妹",
	}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("收到", res.Message)
}
