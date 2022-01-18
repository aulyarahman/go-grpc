package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aulyarahman/go-grpc/calculator/calpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Client Calculator Started")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could connect to server , %v", err)
	}

	defer conn.Close()

	c := calpb.NewCalculatorServiceClient(conn)
	doUnary(c)

}

func doUnary(c calpb.CalculatorServiceClient){
	fmt.Printf("Starting to do unary RPC Calculator")
	req := &calpb.CalculatorRequest{
		Calculators:  &calpb.Calculators{
			PhoneNumber: 9898323,
			Age: 10,
		},
	}

	res,err := c.Calcualtor(context.Background(), req)
	if err != nil {
		log.Fatalf("Error When Call calculator RPC : %v", err)
	}

	log.Printf("Response from Calculator %v", res.Result)
}