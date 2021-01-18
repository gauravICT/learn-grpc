package main

import (
	"context"
	"fmt"
	"log"

	"github.com/my-projects/learn-grpc/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello my name is client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created connection: %f", c)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Gaurav",
			LastName:  "Haldar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)

}
