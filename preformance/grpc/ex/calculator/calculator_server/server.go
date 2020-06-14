package main

import (
	"context"
	"ea/preformance/grpc/ex/calculator/calculatorpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum Rpc: %v", request)
	firstNumber := request.FirstName
	secondNumber := request.SecondName
	sum := firstNumber + secondNumber
	return &calculatorpb.SumResponse{
		SumResult: sum,
	}, nil
}

func main() {
	fmt.Println("Hello world")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
