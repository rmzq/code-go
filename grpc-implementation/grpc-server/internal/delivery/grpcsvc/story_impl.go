package grpcsvc

import (
	"context"
	pb "grpc-implementation/grpc-server/pb/story"
)

type StoryService struct {
	pb.UnimplementedStoryServiceServer
}

func NewStoryService() *StoryService {
	return &StoryService{}
}

func (s *StoryService) FindAll(ctx context.Context, req *pb.FindAllStoriesRequest) (*pb.Stories, error) {
	stories := []*pb.Story{
		{
			Id:      1,
			Title:   "title 1",
			Content: "content 1",
		},
		{
			Id:      2,
			Title:   "title 2",
			Content: "content 2",
		},
		{
			Id:      3,
			Title:   "title 3",
			Content: "content 3",
		},
	}

	return &pb.Stories{
		Stories: stories,
	}, nil
}

func (s *StoryService) FindByID(ctx context.Context, req *pb.FindStoryByIDRequest) (*pb.Story, error) {
	// assume this payload from usecase story
	story := &pb.Story{
		Id:      1,
		Title:   "title",
		Content: "content",
	}

	return story, nil
}
