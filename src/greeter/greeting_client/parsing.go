package main

import (
	"os"
	"strings"

	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
)

const (
	defaultName = "George"
)

var (
	defaultReq = pb.SimpleGreetingRequest{Type: pb.SimpleGreetingRequest_HELLO, Name: defaultName}
)

func getSimpleGreetingReq() *pb.SimpleGreetingRequest {
	gType := os.Getenv("GREETING_TYPE")
	name := os.Getenv("RECEPIENT")

	if gType != "" && name != "" {
		greetingType := pb.SimpleGreetingRequest_Type(
			pb.SimpleGreetingRequest_Type_value[strings.ToUpper(gType)])
		return &pb.SimpleGreetingRequest{Type: greetingType, Name: name}
	}

	return &defaultReq
}
