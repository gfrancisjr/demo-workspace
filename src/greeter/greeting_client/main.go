// Package main creates a Simple Greeting client and sends one greeting request
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	address = "localhost:10000"
)

func run() error {
	c, err := newGrpcClient(address)
	if err != nil {
		return err
	}

	defer c.shutdown()

	service := c.client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := getSimpleGreetingReq()
	resp, err := service.SimpleGreeting(ctx, req)
	if err != nil {
		return fmt.Errorf("error processing greeting request: %w", err)
	}

	fmt.Println(resp.Greeting)

	return nil
}

func main() {
	if err := run(); err != nil {
		logrus.Fatalf("error running greeting client: %v", err)
	}
}
