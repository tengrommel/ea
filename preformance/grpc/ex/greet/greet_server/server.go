package main

import (
	"context"
	"ea/preformance/grpc/ex/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked")
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
