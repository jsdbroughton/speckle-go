package models

import "time"

// Version represents a Speckle version (formerly known as commit)
type Version struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	AuthorId  string    `json:"authorId"`
	// Add other fields as necessary
}
