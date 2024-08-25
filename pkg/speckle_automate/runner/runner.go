package runner

import (
	"fmt"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/api"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/transports"
	"github.com/jsdbroughton/speckle-go/pkg/speckle_automate/automation_context"
	"github.com/jsdbroughton/speckle-go/pkg/speckle_automate/schema"
)

// initializeAutomationContext initializes the AutomationContext
func initializeAutomationContext(runData schema.AutomationRunData, speckleToken string) (*automation_context.AutomationContext, error) {
	speckleClient, err := createSpeckleClient(runData.SpeckleServerURL, speckleToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create Speckle client: %w", err)
	}

	serverTransport, err := createServerTransport(runData.ProjectID, speckleClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create server transport: %w", err)
	}

	memoryTransport := createMemoryTransport()

	ctx, err := automation_context.NewAutomationContext(runData, speckleToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create automation context: %w", err)
	}

	// Set the custom-created fields
	ctx.SpeckleClient = speckleClient
	ctx.ServerTransport = serverTransport
	ctx.MemoryTransport = memoryTransport

	return ctx, nil
}

func createSpeckleClient(serverUrl, token string) (*api.Client, error) {
	client := &api.Client{
		ServerUrl: serverUrl,
	}
	if err := client.AuthenticateWithToken(token); err != nil {
		return nil, fmt.Errorf("failed to create Speckle client: %w", err)
	}
	return client, nil
}

func createServerTransport(projectId string, client *api.Client) (*transports.ServerTransport, error) {
	// Implement server transport creation logic
	// This is a placeholder and needs to be implemented based on the Speckle Go SDK
	return &transports.ServerTransport{}, nil
}

func createMemoryTransport() *transports.MemoryTransport {
	// Implement memory transport creation logic
	// This is a placeholder and needs to be implemented based on the Speckle Go SDK
	return &transports.MemoryTransport{}
}
