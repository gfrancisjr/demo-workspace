package main

import (
	"github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
	"google.golang.org/grpc"
)

type grpcClient struct {
	client greeter.GreeterClient
	conn   *grpc.ClientConn
}

// newGrpcServer creates a new gRPC Client for the Simple Greeting service
func newGrpcClient(address string) (grpcClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return grpcClient{}, err
	}
	return grpcClient{
		client: pb.NewGreeterClient(conn),
		conn:   conn,
	}, nil
}

// shutdown closes the connection to the gRPC server
func (g grpcClient) shutdown() error {
	return g.conn.Close()
}
