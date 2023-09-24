package models

import (
	"os"
	"time"
)

type SSEInterface interface {
	// Stream return a channel of event that will send data from param "adress" during param "duration".
	// Return an error if there is one and a nil channel
	Stream(address string, duration time.Duration) (chan Event, error)
}

func GetNewClient() SSEInterface {
	if os.Getenv("ENVIRONMENT") == "test" {
		return StubSSEClient{}
	}
	return SSEClient{}
}
