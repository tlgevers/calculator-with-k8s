// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "addition/pkg/addition"
	"google.golang.org/grpc"
	"strconv"
	"strings"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAdderServer
}

// Add accepts a string eg:"1+1+2" splits, converts to float, sums and returns those values
func Add(a string) (sum string) {
	digitsStr := strings.Split(a, "+")
	fmt.Println("digits", digitsStr)
	var computeSum float64
	for i, d := range digitsStr {
		fmt.Println(i, d)
		f, err := strconv.ParseFloat(d, 64)
		if err != nil {
			e := fmt.Errorf("Invalid argument %q", err)
			fmt.Println(e.Error())
		}
		computeSum += f
		fmt.Println("computeSum", computeSum)
	}
	sum = fmt.Sprintf("%f", computeSum)
	return
}

// sayHello implements helloworld.GreeterServer.SayHello
func (s *server) CalcAdd(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v", in.GetArgs())
	sum := Add(in.GetArgs())
	return &pb.AddResponse{Sum: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
