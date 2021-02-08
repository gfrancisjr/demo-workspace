package main

import (
	"context"
	"fmt"

	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
}

var (
	errInvalidGreetingType = status.Error(codes.InvalidArgument, "greeting type is invalid")
	errInvalidName         = status.Error(codes.InvalidArgument, "name is required")
)

// SimpleGreeting creates a simple greeting
func (s *service) SimpleGreeting(ctx context.Context, req *pb.SimpleGreetingRequest) (*pb.SimpleGreetingResponse, error) {
	return simpleGreeting(ctx, req)
}

func simpleGreeting(ctx context.Context, req *pb.SimpleGreetingRequest) (*pb.SimpleGreetingResponse, error) {
	if req.Type == pb.SimpleGreetingRequest_UNKNOWN {
		return nil, errInvalidGreetingType
	}

	if req.Name == "" {
		return nil, errInvalidName
	}

	greeting, err := createSimpleGreeting(req.Type, req.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.SimpleGreetingResponse{Greeting: greeting}, nil
}

// createSimpleGreeting creates a simple greeting from the greeting type and name
func createSimpleGreeting(greetingType pb.SimpleGreetingRequest_Type, name string) (string, error) {
	switch greetingType {
	case pb.SimpleGreetingRequest_HELLO:
		return fmt.Sprintf("Hello %s", name), nil
	case pb.SimpleGreetingRequest_GOOD_MORNING:
		return fmt.Sprintf("Good morning %s", name), nil
	case pb.SimpleGreetingRequest_GOOD_NIGHT:
		return fmt.Sprintf("Good night %s", name), nil
	case pb.SimpleGreetingRequest_HAPPY_BIRTHDAY:
		return fmt.Sprintf("Happy birthday %s", name), nil
	case pb.SimpleGreetingRequest_MERRY_XMAS:
		return fmt.Sprintf("Merry Christmas %s", name), nil
	default:
		return "", fmt.Errorf("unsupported greeting type: %s", greetingType.String())
	}
}
