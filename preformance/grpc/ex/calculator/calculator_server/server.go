package main

import (
	"context"
	"ea/preformance/grpc/ex/calculator/calculatorpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct{}

func (s server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC\n")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err != io.EOF {
			average := float64(sum)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func (s server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Reveived PrimeNumberDecomposition RPC: %v\n", request)
	number := request.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v", divisor)
		}
	}
	return nil
}

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
