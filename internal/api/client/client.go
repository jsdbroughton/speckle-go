package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/specklesystems/speckle-go/internal/api/models"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
		Token: token,
	}
}

func (c *Client) SetToken(token string) {
	c.Token = token
}

func (c *Client) GetProject(projectID string) (*models.Project, error) {
	url := fmt.Sprintf("%s/projects/%s", c.BaseURL, projectID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var project models.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling project data: %w", err)
	}

	return &project, nil
}
