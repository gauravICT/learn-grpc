package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Workspace/learn-grpc/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Recdeived Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("Recdeived PrimeNumberDecomposition RPC: %v", req)

	number := req.GetNumber()
	divisior := int64(2)
	for number > 1 {
		if number%divisior == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisior,
			})
			number = number / divisior
		} else {
			divisior++
			fmt.Printf("Divisor has increased to %v\n", divisior)
		}
	}
	return nil
}

func main() {

	fmt.Println("Calculator Server Start to Serv")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}

}
