package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jsdbroughton/speckle-go/internal/api/client"
	"github.com/jsdbroughton/speckle-go/internal/api/models"
	"net/http"
)

type ProjectResource struct {
	Client *client.Client
}

func NewProjectResource(client *client.Client) *ProjectResource {
	return &ProjectResource{Client: client}
}

func (sr *ProjectResource) Get(id string, modelLimit, versionLimit int) (*models.Project, error) {
	url := fmt.Sprintf("%s/streams/%s?branchLimit=%d&commitLimit=%d", sr.Client.BaseURL, id, modelLimit, versionLimit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var project models.Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &project, nil
}

func (sr *ProjectResource) List(projectLimit int) ([]models.Project, error) {
	url := fmt.Sprintf("%s/streams?limit=%d", sr.Client.BaseURL, projectLimit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var projects []models.Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return projects, nil
}

func (sr *ProjectResource) Create(name, description string, isPublic bool) (string, error) {
	url := fmt.Sprintf("%s/streams", sr.Client.BaseURL)

	payload := map[string]interface{}{
		"name":        name,
		"description": description,
		"isPublic":    isPublic,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response body: %w", err)
	}

	return result["id"], nil
}
