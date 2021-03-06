package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aulyarahman/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet func was invoked with %v", req)
 firstName := req.GetGreeting().GetFirstName()
 result := "Hello" + firstName
 res := &greetpb.GreetResponse{
	 Result: result,
 }

 return res, nil
}

func main() {
	fmt.Println("Server Service Greet Running")

	lis, err:=net.Listen("tcp", "0.0.0.0:50051")  //network string, address
	if err!= nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	

	if err:= s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}