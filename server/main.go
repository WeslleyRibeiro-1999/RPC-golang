package main

import (
	"context"
	"log"
	"net"

	pb "github.com/WeslleyRibeiro-1999/RPC-golang/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (*Server) HelloMessage(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello, " + req.GetName()

	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServe := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServe, &Server{})

	log.Println("Rodando")

	if err := grpcServe.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
