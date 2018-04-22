package main

import (
	"context"
	"fmt"
	"github.com/nmiculinic/go-generate-talk/gRPC/summer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Math int

func (Math) Sum(ctx context.Context, req *summer.SumRequest) (*summer.SumResponse, error) {
	if len(req.A) == 0 {
		return nil, fmt.Errorf("no elements to sum")
	}
	var sol int32 = 0
	for _, x := range req.A {
		sol += x
	}
	return &summer.SumResponse{
		Sum: sol,
	}, nil
}

func server() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	summer.RegisterMathServer(s, Math(0))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	go server()
	log.Println("go to http://localhost:8080/Math/Sum")
	summer.Serve("[::]:8080", "localhost:50051", summer.DefaultHtmlStringer, grpc.WithInsecure())
}
