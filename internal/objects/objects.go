package objects

import (
	"time"
)

// Base represents the base Speckle object
type Base struct {
	ID                string                 `json:"id"`
	SpeckleType       string                 `json:"speckleType"`
	ApplicationID     string                 `json:"applicationId,omitempty"`
	CreatedAt         time.Time              `json:"createdAt,omitempty"`
	UpdatedAt         time.Time              `json:"updatedAt,omitempty"`
	TotalChildrenCount int                   `json:"totalChildrenCount"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
}

// NewBase creates a new Base object
func NewBase() *Base {
	return &Base{
		SpeckleType: "Base",
		Properties:  make(map[string]interface{}),
	}
}

