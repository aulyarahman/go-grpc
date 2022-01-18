package main

import (
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
	fmt.Printf("Created Client %f" ,c)
}