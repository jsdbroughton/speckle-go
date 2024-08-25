package core

import (
	"errors"
	"fmt"
	resources2 "github.com/jsdbroughton/speckle-go/pkg/speckle/resources"
	"net/http"
	"time"

	"github.com/jsdbroughton/speckle-go/internal/api/resources"
)

// NewCoreClient initializes a new core Speckle Client
func NewCoreClient(baseURL string, useSSL bool) *Client {
	scheme := "http"
	if useSSL {
		scheme = "https"
	}
	client := &Client{
		BaseURL: fmt.Sprintf("%s://%s", scheme, baseURL),
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
	client.initializeResources()
	return client
}

func (c *Client) initializeResources() {
	c.Server = resources2.NewServerResource(c)
	c.User = resources.NewUserResource(c)
	c.OtherUser = resources.NewOtherUserResource(c)
	c.ActiveUser = resources.NewActiveUserResource(c)
	c.Project = resources2.NewProjectResource(c)
	c.Version = resources.NewVersionResource(c)
	c.Model = resources.NewModelResource(c)
	c.Object = resources.NewObjectResource(c)
	c.Subscribe = resources.NewSubscriptionResource(c)
}

func (c *Client) AuthenticateWithToken(token string) error {
	c.Token = token
	return c.verifyToken()
}

func (c *Client) verifyToken() error {
	if c.Token == "" {
		return errors.New("token is empty")
	}
	// TODO: Implement actual token verification logic
	return nil
}
