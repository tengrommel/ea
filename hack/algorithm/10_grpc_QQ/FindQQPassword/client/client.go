package main

import (
	"context"
	FindQQPassword "ea/hack/algorithm/10_grpc_QQ/FindQQPassword/protoc"
	"fmt"
	"google.golang.org/grpc"
)

const Address = "127.0.0.1:8848"

func main() {
	for {
		var QQNum string
		fmt.Scanln(&QQNum)
		conn, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		c := FindQQPassword.NewFindQQPasswordServiceClient(conn)
		req := &FindQQPassword.QQRequest{QQNum: QQNum}
		res, err := c.FindQQPassword(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("收到", res.QQPassword)
	}
}
