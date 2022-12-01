package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/WeslleyRibeiro-1999/RPC-golang/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (*Server) HelloMessage(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello World" + req.GetMsg()

	res := &pb.HelloResponse{
		Text: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServe := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServe, &Server{})

	if err := grpcServe.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	fmt.Println("Rodando")
}
