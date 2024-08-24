package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/api"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/models"
	"github.com/machinebox/graphql"
	"net/http"
	"regexp"
	"strconv"
)

type ServerResource struct {
	Client *api.Client
}

func NewServerResource(client *api.Client) *ServerResource {
	return &ServerResource{Client: client}
}

func (sr *ServerResource) Get() (*models.ServerInfo, error) {
	gqlClient := graphql.NewClient(fmt.Sprintf("%s/graphql", sr.Client.BaseURL))

	// Define the GraphQL query
	query := `
	query Server {
		serverInfo {
			name
			company
			description
			adminContact
			canonicalUrl
			version
			roles {
				name
				description
				resourceTarget
			}
			scopes {
				name
				description
			}
			authStrategies {
				id
				name
				icon
			}
		}
	}`

	// Create a request
	req := graphql.NewRequest(query)

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	// Define a response struct
	var response struct {
		ServerInfo models.ServerInfo `json:"serverInfo"`
	}

	// Perform the query
	if err := gqlClient.Run(context.Background(), req, &response); err != nil {
		return nil, fmt.Errorf("error executing GraphQL query: %w", err)
	}

	serverInfo := response.ServerInfo

	// Check the `canonicalUrl` as in the Python code
	if serverInfo.CanonicalURL != nil {
		resp, err := sr.Client.HTTPClient.Get(*serverInfo.CanonicalURL)
		if err != nil {
			return nil, fmt.Errorf("error checking canonical URL: %w", err)
		}
		defer func() {
			if closeError := resp.Body.Close(); closeError != nil {
				// Log the error or return it if necessary
				fmt.Printf("error closing response body: %v\n", closeError)
			}
		}()

		if _, ok := resp.Header["X-Speckle-Frontend-2"]; ok {
			frontend2 := true
			serverInfo.Frontend2 = &frontend2
		} else {
			frontend2 := false
			serverInfo.Frontend2 = &frontend2
		}
	}

	return &serverInfo, nil
}

func (sr *ServerResource) Version() (*models.ServerVersion, error) {
	url := fmt.Sprintf("%s/server/version", sr.Client.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer func() {
		if closeError := resp.Body.Close(); closeError != nil {
			// Log the error or return it if necessary
			fmt.Printf("error closing response body: %v\n", closeError)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var versionString string
	if err := json.NewDecoder(resp.Body).Decode(&versionString); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	// Split the version string into components
	var components []interface{}
	for _, segment := range regexp.MustCompile(`[.-]`).Split(versionString, -1) {
		if num, err := strconv.Atoi(segment); err == nil {
			components = append(components, num)
		} else {
			components = append(components, segment)
		}
	}

	version := &models.ServerVersion{
		Components: components,
	}

	return version, nil
}

func (sr *ServerResource) Apps() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/server/apps", sr.Client.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer func() {
		if closeError := resp.Body.Close(); closeError != nil {
			// Log the error or return it if necessary
			fmt.Printf("error closing response body: %v\n", closeError)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apps map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apps); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return apps, nil
}

func (sr *ServerResource) CreateToken(name string, scopes []string, lifespan int) (string, error) {
	url := fmt.Sprintf("%s/server/tokens", sr.Client.BaseURL)

	payload := map[string]interface{}{
		"name":     name,
		"scopes":   scopes,
		"lifespan": lifespan,
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
	defer func() {
		if closeError := resp.Body.Close(); closeError != nil {
			// Log the error or return it if necessary
			fmt.Printf("error closing response body: %v\n", closeError)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response body: %w", err)
	}

	return result["token"], nil
}

func (sr *ServerResource) RevokeToken(token string) (bool, error) {
	url := fmt.Sprintf("%s/server/tokens/%s", sr.Client.BaseURL, token)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+sr.Client.Token)

	resp, err := sr.Client.HTTPClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error sending request: %w", err)
	}
	defer func() {
		if closeError := resp.Body.Close(); closeError != nil {
			// Log the error or return it if necessary
			fmt.Printf("error closing response body: %v\n", closeError)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return true, nil
}
