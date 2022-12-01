package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/WeslleyRibeiro-1999/RPC-golang/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("alou")

	client := pb.NewHelloServiceClient(conn)

	request := &pb.HelloRequest{
		Msg: "Ola Mundo",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.HelloMessage(ctx, request)
	if err != nil {
		log.Fatalf("could not execute: %v", err)
	}

	log.Println(res.Text)
}
