package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jsdbroughton/speckle-go/internal/api/models"
)

func TestGetProject(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has the correct path
		if r.URL.Path != "/projects/test-project-id" {
			t.Errorf("Expected to request '/projects/test-project-id', got: %s", r.URL.Path)
		}

		// Check if the Authorization header is set
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Expected Authorization header to be 'Bearer test-token', got: %s", r.Header.Get("Authorization"))
		}

		// Send mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Project{
			ID:   "test-project-id",
			Name: "Test Project",
		})
	}))
	defer server.Close()

	// Create a client using the mock server URL
	client := NewClient(server.URL, "test-token")

	// Test the GetProject method
	project, err := client.GetProject("test-project-id")
	if err != nil {
		t.Fatalf("GetProject returned an error: %v", err)
	}

	// Check the returned project
	if project.ID != "test-project-id" {
		t.Errorf("Expected project ID 'test-project-id', got '%s'", project.ID)
	}
	if project.Name != "Test Project" {
		t.Errorf("Expected project name 'Test Project', got '%s'", project.Name)
	}
}
