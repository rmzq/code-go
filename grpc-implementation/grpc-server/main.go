package main

import (
	"fmt"
	"grpc-implementation/grpc-server/internal/delivery/grpcsvc"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-implementation/grpc-server/pb/story"
)

const port = 50051

func main() {
	grpcServer := grpc.NewServer()
	storyService := grpcsvc.NewStoryService()

	pb.RegisterStoryServiceServer(grpcServer, storyService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Panicf("failed to create http listener, error: %v", err)
	}

	log.Printf("grpc server is running with port: %d", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Panicf("failed to run grpc server, error: %v", err)
	}

}
