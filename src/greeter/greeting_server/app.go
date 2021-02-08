package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	grpcPort = "10000"
)

type app struct {
	grpcServer grpcServer
	// listens for anapplication termination signal
	shutdownCh chan os.Signal
}

// start starts the Simple Greeting service over GRPC
func (a app) start() {
	a.grpcServer.start()
}

// shutdown stops the Simple Greeting gRPC service
func (a app) shutdown() {
	a.grpcServer.stop()
}

// newApp creates a new app with a gRPC server
// it performs all app related initialization
func newApp() (app, error) {
	gs, err := newGrpcServer(&service{}, grpcPort, net.Listen)
	if err != nil {
		return app{}, err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	return app{
		grpcServer: gs,
		shutdownCh: quit,
	}, nil
}
