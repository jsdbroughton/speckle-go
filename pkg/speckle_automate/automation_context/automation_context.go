package automation_context

import (
	"fmt"
	"github.com/jsdbroughton/speckle-go/pkg/speckle_automate/schema"
	"time"

	"github.com/jsdbroughton/speckle-go/pkg/speckle/api"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/transports"
)

type AutomationContext struct {
	AutomationRunData schema.AutomationRunData
	SpeckleClient     *api.Client
	ServerTransport   *transports.ServerTransport
	MemoryTransport   *transports.MemoryTransport
	speckleToken      string
	initTime          time.Time
	AutomationResult  schema.AutomationResult
}

func NewAutomationContext(runData schema.AutomationRunData, speckleToken string) (*AutomationContext, error) {
	client, err := api.NewClient(runData.SpeckleServerURL, true)
	if err != nil {
		return nil, fmt.Errorf("failed to create Speckle client: %w", err)
	}

	err = client.Authenticate(speckleToken)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	serverTransport, err := transports.NewServerTransport(runData.ProjectID, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create server transport: %w", err)
	}

	return &AutomationContext{
		AutomationRunData: runData,
		SpeckleClient:     client,
		ServerTransport:   serverTransport,
		MemoryTransport:   transports.NewMemoryTransport(),
		speckleToken:      speckleToken,
		initTime:          time.Now(),
		AutomationResult: schema.AutomationResult{
			RunStatus: schema.StatusInitializing,
		},
	}, nil
}

func (ac *AutomationContext) ContextView() *string {
	return ac.AutomationResult.ResultView
}

func (ac *AutomationContext) RunStatus() string {
	return string(ac.AutomationResult.RunStatus)
}

func (ac *AutomationContext) StatusMessage() *string {
	return ac.AutomationResult.StatusMessage
}

func (ac *AutomationContext) Elapsed() time.Duration {
	return time.Since(ac.initTime)
}
