package client

import (
	"os"
	"sync"

	"go.temporal.io/sdk/client"
)

var (
	temporalClient client.Client
	once           sync.Once
)

func GetTemporalClient() (client.Client, error) {
	var err error
	once.Do(func() {
		hostPort := os.Getenv("TEMPORAL_HOST_PORT")
		if hostPort == "" {
			hostPort = client.DefaultHostPort
		}
		temporalClient, err = client.Dial(client.Options{HostPort: hostPort})
	})
	return temporalClient, err
}
