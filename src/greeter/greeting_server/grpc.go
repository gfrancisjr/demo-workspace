package main

import (
	"net"
	"time"

	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
	port     string
}

var (
	logrusLogger *logrus.Logger = logrus.New()
)

type listenFunc func(string, string) (net.Listener, error)

// newGrpcServer creates a new gRPC Server for the Simple Greeting service
func newGrpcServer(service pb.GreeterServer, port string, listen listenFunc) (grpcServer, error) {
	lis, err := listen("tcp", ":"+port)
	if err != nil {
		return grpcServer{}, err
	}

	server := grpc.NewServer(middleware())
	pb.RegisterGreeterServer(server, service)

	return grpcServer{
		server:   server,
		listener: lis,
		port:     port,
		errCh:    make(chan error),
	}, nil
}

// start starts the server in the background
func (g grpcServer) start() {
	go func() {
		logrus.Infof("starting gRPC server on port: %s", g.port)
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// stop stops the gRPC server
func (g grpcServer) stop() {
	logrus.Info("shutting down gRPC server...")
	g.server.GracefulStop()
}

// Error returns the server's error channel
func (g grpcServer) error() chan error {
	return g.errCh
}

// middleware returns the middleware for the gRPC server
func middleware() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		loggingUnaryInterceptor(),
	))
}

// loggingUnaryInterceptor returns a UnaryServerInterceptor for logging gRPC requests
func loggingUnaryInterceptor() grpc.UnaryServerInterceptor {
	logrusEntry := logrus.NewEntry(logrusLogger)
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ms", duration.Milliseconds()
		}),
	}
	return grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...)
}
