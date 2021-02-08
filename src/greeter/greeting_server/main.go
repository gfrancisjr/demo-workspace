// Package main exposes the Simple Greeting service over gRPC
package main

import (
	"github.com/sirupsen/logrus"
)

// run starts gRPC server and handles any shutdown signal
func run() error {
	app, err := newApp()
	if err != nil {
		return err
	}

	app.start()
	defer app.shutdown()

	select {
	case grpcErr := <-app.grpcServer.error():
		return grpcErr
	case <-app.shutdownCh:
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		logrus.Fatalf("error serving gRPC requests: %v", err)
	}
}
