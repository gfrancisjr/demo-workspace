package main

import (
	"context"
	"testing"

	pb "github.com/gfrancisjr/demo-workspace/src/greeter/greeter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	name          = "John Doe"
	greetingtests = []struct {
		testName         string
		greetingType     pb.SimpleGreetingRequest_Type
		expectedGreeting string
	}{
		{
			testName:         "hello greeting",
			greetingType:     pb.SimpleGreetingRequest_HELLO,
			expectedGreeting: "Hello " + name,
		},
		{
			testName:         "good morning greeting",
			greetingType:     pb.SimpleGreetingRequest_GOOD_MORNING,
			expectedGreeting: "Good morning " + name,
		},
		{
			testName:         "good night greeting",
			greetingType:     pb.SimpleGreetingRequest_GOOD_NIGHT,
			expectedGreeting: "Good night " + name,
		},
		{
			testName:         "happy birthday greeting",
			greetingType:     pb.SimpleGreetingRequest_HAPPY_BIRTHDAY,
			expectedGreeting: "Happy birthday " + name,
		},
		{
			testName:         "merry christmas greeting",
			greetingType:     pb.SimpleGreetingRequest_MERRY_XMAS,
			expectedGreeting: "Merry Christmas " + name,
		},
	}
)

func Test_createSimpleGreeting(t *testing.T) {
	for _, tt := range greetingtests {
		t.Run(tt.testName, func(t *testing.T) {
			actualGreeting, err := createSimpleGreeting(tt.greetingType, name)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedGreeting, actualGreeting)
		})
	}

	t.Run("invalid greeting type", func(t *testing.T) {
		actualGreeting, err := createSimpleGreeting(pb.SimpleGreetingRequest_UNKNOWN, name)

		assert.Error(t, err)
		assert.Empty(t, actualGreeting)
	})
}

func Test_simpleGreeting(t *testing.T) {
	t.Run("valid greeting request", func(t *testing.T) {
		ctx := context.Background()
		req := pb.SimpleGreetingRequest{
			Type: pb.SimpleGreetingRequest_HELLO,
			Name: name,
		}

		resp, err := simpleGreeting(ctx, &req)

		assert.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, "Hello John Doe", resp.Greeting)

	})

	t.Run("invalid greeting type", func(t *testing.T) {
		ctx := context.Background()
		req := pb.SimpleGreetingRequest{
			Type: pb.SimpleGreetingRequest_UNKNOWN,
			Name: name,
		}

		resp, err := simpleGreeting(ctx, &req)

		assert.Equal(t, errInvalidGreetingType, err)
		require.Nil(t, resp)
	})
	t.Run("invalid name", func(t *testing.T) {
		ctx := context.Background()
		req := pb.SimpleGreetingRequest{
			Type: pb.SimpleGreetingRequest_HELLO,
		}

		resp, err := simpleGreeting(ctx, &req)

		assert.Equal(t, errInvalidName, err)
		assert.Nil(t, resp)
	})
}
