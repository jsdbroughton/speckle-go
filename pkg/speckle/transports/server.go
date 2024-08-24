package transports

import (
	"errors"

	"github.com/jsdbroughton/speckle-go/pkg/speckle/api"
)

// ServerTransport represents a transport for server-based operations
type ServerTransport struct {
	projectId string
	client    *api.Client
}

// NewServerTransport creates a new ServerTransport
func NewServerTransport(projectId string, client *api.Client) (*ServerTransport, error) {
	if projectId == "" {
		return nil, errors.New("project ID cannot be empty")
	}
	if client == nil {
		return nil, errors.New("client cannot be nil")
	}
	return &ServerTransport{
		projectId: projectId,
		client:    client,
	}, nil
}
