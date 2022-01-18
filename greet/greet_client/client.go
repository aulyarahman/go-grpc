package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aulyarahman/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could connect to server %v" , err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created Client %f" ,c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Printf("Starting to do unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Aulia",
			LastName: "Rahman",
		},
	}

	res, err := c.Greet(context.Background(), req) 
	if err != nil {
		log.Fatalf("error when call Greet RPC : &v", err)
	}

	log.Printf("Response from Greet %v", res.Result)
}