package main

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
)

const bufSize = 1024 * 1024

var mockListener *bufconn.Listener
var mockListen func(string, string) (net.Listener, error)
var mockDialer func(context.Context, string) (net.Conn, error)

func TestMain(m *testing.M) {
	mockListener = bufconn.Listen(bufSize)
	mockListen = func(network string, address string) (net.Listener, error) {
		return mockListener, nil
	}
	mockDialer = func(context.Context, string) (net.Conn, error) {
		return mockListener.Dial()
	}
	os.Exit(m.Run())
}

func TestSimpleGreeting(t *testing.T) {
	gs, err := newGrpcServer(&service{}, grpcPort, mockListen)
	require.NoError(t, err)

	gs.start()
	defer gs.stop()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "foobar", grpc.WithContextDialer(mockDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	req := pb.SimpleGreetingRequest{
		Type: pb.SimpleGreetingRequest_HELLO,
		Name: "John Doe",
	}
	resp, err := client.SimpleGreeting(ctx, &req)
	require.NoError(t, err)
	assert.Equal(t, "Hello John Doe", resp.Greeting)
}
