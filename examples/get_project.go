package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jsdbroughton/speckle-go/internal/api/client"
)

func main() {
	// Replace these with your actual Speckle server URL and token
	baseURL := "https://speckle.xyz/api"
	token := os.Getenv("SPECKLE_TOKEN")

	if token == "" {
		log.Fatal("SPECKLE_TOKEN environment variable is not set")
	}

	// Replace this with an actual project ID from your Speckle server
	projectID := "2f088c4d2e"

	c := client.NewClient(baseURL, token)

	project, err := c.GetProject(projectID)
	if err != nil {
		log.Fatalf("Error getting project: %v", err)
	}

	fmt.Printf("Project: %+v\n", project)
}
