package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jsdbroughton/speckle-go/internal/api/models"
)

func TestGetProject(t *testing.T) {
	t.Log("Starting TestGetProject")

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("Received request:", r.Method, r.URL.Path)

		// Check if the request has the correct path
		if r.URL.Path != "/projects/test-project-id" {
			t.Errorf("Expected to request '/projects/test-project-id', got: %s", r.URL.Path)
		} else {
			t.Log("Request path is correct")
		}

		// Check if the Authorization header is set
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-token" {
			t.Errorf("Expected Authorization header to be 'Bearer test-token', got: %s", authHeader)
		} else {
			t.Log("Authorization header is correct")
		}

		// Send mock response
		t.Log("Sending mock response")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Project{
			ID:   "test-project-id",
			Name: "Test Project",
		})
	}))
	defer func() {
		t.Log("Closing mock server")
		server.Close()
	}()

	// Create a client using the mock server URL
	t.Logf("Creating client with server URL: %s", server.URL)
	client := NewClient(server.URL, "test-token")

	// Test the GetProject method
	t.Log("Calling GetProject")
	project, err := client.GetProject("test-project-id")
	if err != nil {
		t.Fatalf("GetProject returned an error: %v", err)
	} else {
		t.Log("GetProject did not return an error")
	}

	// Check the returned project
	if project.ID != "test-project-id" {
		t.Errorf("Expected project ID 'test-project-id', got '%s'", project.ID)
	} else {
		t.Log("Project ID is correct")
	}
	if project.Name != "Test Project" {
		t.Errorf("Expected project name 'Test Project', got '%s'", project.Name)
	} else {
		t.Log("Project name is correct")
	}

	t.Log("TestGetProject completed")
}

func TestGetProject_NotFound(t *testing.T) {
	t.Log("Starting TestGetProject_NotFound")

	// Create a mock server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("Received request: ", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
	}))
	defer server.Close()

	// Create a client using the mock server URL
	client := NewClient(server.URL, "test-token")

	t.Log("Calling GetProject with non-existent project ID")
	_, err := client.GetProject("non-existent-project-id")
	if err == nil {
		t.Fatalf("Expected an error for non-existent project ID, got none")
	}

	t.Log("TestGetProject_NotFound passed")
}
