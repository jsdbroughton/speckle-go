package client

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	Account    *Account

	Server     *ServerResource
	User       *UserResource
	OtherUser  *OtherUserResource
	ActiveUser *ActiveUserResource
	Project    *ProjectResource
	Version    *VersionResource
	Model      *ModelResource
	Object     *ObjectResource
	Subscribe  *SubscriptionResource
}

// Account structure
type Account struct {
	Token string
}

// NewClient initializes a new Speckle Client
func NewClient(baseURL string, useSSL, verifyCertificate bool) *Client {
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
	// Initialize the resources with the current client instance
	c.Server = NewServerResource(c)
	c.User = NewUserResource(c)
	c.OtherUser = NewOtherUserResource(c)
	c.ActiveUser = NewActiveUserResource(c)
	c.Project = NewStreamResource(c)
	c.Version = NewCommitResource(c)
	c.Model = NewBranchResource(c)
	c.Object = NewObjectResource(c)
	c.Subscribe = NewSubscriptionResource(c)
}

func (c *Client) AuthenticateWithToken(token string) {
	c.Token = token
	// Perform any additional setup or verification here
}

func (c *Client) AuthenticateWithAccount(account *Account) {
	c.Token = account.Token
	c.Account = account
	// Perform any additional setup or verification here
}
