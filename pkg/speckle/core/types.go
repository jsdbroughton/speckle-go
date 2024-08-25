package core

import (
	"encoding/json"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/resources"
	"net/http"
)

type UserInfo struct {
	Name    string
	Email   string
	Company string
	ID      string
}

type ServerInfo struct {
	URL string
	// Add other necessary fields
}

type Account struct {
	IsDefault    bool       `json:"isDefault"`
	Token        string     `json:"token"`
	RefreshToken string     `json:"refreshToken"`
	ServerInfo   ServerInfo `json:"serverInfo"`
	UserInfo     UserInfo   `json:"userInfo"`
	ID           string     `json:"id"`
}

func (a *Account) UnmarshalJSON(data []byte) error {
	type Alias Account
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// Client represents the core functionality of the Speckle client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	Server     *resources.ServerResource
	User       *resources.UserResource
	OtherUser  *resources.OtherUserResource
	ActiveUser *resources.ActiveUserResource
	Project    *resources.ProjectResource
	Version    *resources.VersionResource
	Model      *resources.ModelResource
	Object     *resources.ObjectResource
	Subscribe  *resources.SubscriptionResource
}
