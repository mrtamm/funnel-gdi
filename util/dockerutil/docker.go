package dockerutil

import (
	"context"
	"os"
	"time"

	"github.com/docker/docker/client"
)

// NewDockerClient returns a new docker client. This util handles
// working around some client/server API version mismatch issues.
func NewDockerClient() (*client.Client, error) {
	dclient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	// If the api version is not set test if the client can communicate with the
	// server; if not infer API version from error message and inform the client
	// to use that version for future communication
	if os.Getenv("DOCKER_API_VERSION") == "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		dclient.NegotiateAPIVersion(ctx)
	}

	return dclient, nil
}
