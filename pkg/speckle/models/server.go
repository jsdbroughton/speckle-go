package models

import models2 "github.com/jsdbroughton/speckle-go/internal/api/models"

type ServerInfo struct {
	AdminContact            *string                     `json:"adminContact,omitempty"`
	AuthStrategies          *[]AuthStrategy             `json:"authStrategies,omitempty"`
	Automate                *models2.ServerAutomateInfo `json:"speckle_automate,omitempty"`
	AutomateURL             *string                     `json:"automateUrl,omitempty"`
	BlobSizeLimitBytes      *int                        `json:"blobSizeLimitBytes,omitempty"`
	CanonicalURL            *string                     `json:"canonicalUrl,omitempty"`
	Company                 *string                     `json:"company,omitempty"`
	Description             *string                     `json:"description,omitempty"`
	EnableNewWebUIMessaging *bool                       `json:"enableNewWebUiMessaging,omitempty"`
	Frontend2               *bool                       `json:"frontend2,omitempty"`
	GuestModeEnabled        *bool                       `json:"guestModeEnabled,omitempty"`
	InviteOnly              *bool                       `json:"inviteOnly,omitempty"`
	Migration               *ServerMigration            `json:"migration,omitempty"`
	Name                    *string                     `json:"name,omitempty"`
	ServerRoles             *[]ServerRoleItem           `json:"serverRoles,omitempty"`
	Scopes                  *[]Scope                    `json:"scopes,omitempty"`
	TermsOfService          *string                     `json:"termsOfService,omitempty"`
	Version                 *string                     `json:"version,omitempty"`
	Workspaces              *ServerWorkspacesInfo       `json:"workspaces,omitempty"`
}

type ServerMigration struct {
	MovedTo   *string `json:"movedTo,omitempty"`
	MovedFrom *string `json:"movedFrom,omitempty"`
}

type ServerVersion struct {
	Components []interface{} // Can hold both integers and strings
}

type AuthStrategy struct {
	Color *string `json:"color,omitempty"`
	Icon  *string `json:"icon,omitempty"`
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	URL   *string `json:"url,omitempty"`
}

type Scope struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ServerRoleItem struct {
	ID    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

type ServerWorkspacesInfo struct {
	WorkspacesEnabled *bool `json:"workspacesEnabled,omitempty"`
}
