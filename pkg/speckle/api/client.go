package api

import (
	"github.com/jsdbroughton/speckle-go/pkg/speckle/core"
)

// Client extends CoreClient with additional functionality and metrics
type Client struct {
	CoreClient     *core.Client
	metricsEnabled bool
	Account        *core.Account
}

// NewClient initializes a new Speckle Client with extended functionality
func NewClient(baseURL string, useSSL bool, enableMetrics bool) *Client {
	return &Client{
		CoreClient:     core.NewCoreClient(baseURL, useSSL),
		metricsEnabled: enableMetrics,
	}
}

func (c *Client) AuthenticateWithAccount(account *core.Account) error {
	if c.metricsEnabled {
		trackMetric("Client Authenticate With Account", nil)
	}
	c.Account = account
	return c.CoreClient.AuthenticateWithToken(account.Token)
}

func (c *Client) AuthenticateWithToken(token string) error {
	if c.metricsEnabled {
		trackMetric("Client Authenticate With Token", nil)
	}
	return c.CoreClient.AuthenticateWithToken(token)
}

func (c *Client) SetMetricsEnabled(enabled bool) {
	c.metricsEnabled = enabled
}
