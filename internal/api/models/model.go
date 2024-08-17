package models

import "time"

// Model represents a Speckle model (formerly known as branch)
type Model struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Add other fields as necessary
}
