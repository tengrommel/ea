package main

import (
	"ea/preformance/grpc/ex/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Create client: %f", c)
}
