package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-implementation/grpc-server/pb/story"
)

const storygRPCHost = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(storygRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to create grpc client, error: %v", err)
	}

	defer conn.Close()

	storyClient := pb.NewStoryServiceClient(conn)

	result, err := storyClient.FindAll(context.Background(), &pb.FindAllStoriesRequest{})
	if err != nil {
		log.Panicf("failed to call FindAll, error: %v", err)
	}

	bStories := result.GetStories()
	fmt.Println(bStories)

	for _, bStory := range bStories {
		log.Printf("id: %d, title: %s, content: %s", bStory.GetId(), bStory.GetTitle(), bStory.GetContent())
	}

}
