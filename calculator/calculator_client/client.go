package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Workspace/learn-grpc/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello my name is calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created connection: %f", c)
	doUnary(c)
	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting SUM Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:  58,
		SecondNumber: 42,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from SumResult: %v", res.SumResult)

}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting PrimeDecomposition Server Streaming RPC")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 585896542,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happend: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}

}
