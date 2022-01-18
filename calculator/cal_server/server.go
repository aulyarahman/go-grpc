package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aulyarahman/go-grpc/calculator/calpb"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) Calcualtor(ctx context.Context, req *calpb.CalculatorRequest) (*calpb.CalculatorResponse, error) {
	fmt.Printf("Calculator invoked %v", req)

	res := &calpb.CalculatorResponse{
		Result: req.GetCalculators(),
	}

	return res, nil
} 

func main() {
	fmt.Println("Server Calculator Running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}