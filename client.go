package gnosispay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	AuthToken  string
	Domain     string
	Uri        string
	ChainID    int
}

// NewClient creates a new Gnosis Pay API client.
// baseURL is the API endpoint (e.g., "https://api.gnosispay.com")
// domain is your application's domain for SIWE authentication
// uri is your application's URI for SIWE authentication
// chainID is the blockchain network ID (e.g., 100 for Gnosis Chain)
func NewClient(baseURL, domain, uri string, chainID int) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	return &Client{
		BaseURL: parsedURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		AuthToken: "",
		Domain:    domain,
		Uri:       uri,
		ChainID:   chainID,
	}, nil
}

// buildURL ensures safe URL concatenation
func (c *Client) buildURL(endpoint string) (string, error) {
	if endpoint == "" {
		return "", fmt.Errorf("endpoint cannot be empty")
	}

	relPath, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("invalid endpoint path: %w", err)
	}

	fullURL := c.BaseURL.ResolveReference(relPath)
	return fullURL.String(), nil
}

// doRequest handles API requests with authentication
func (c *Client) doRequest(method, endpoint string, body interface{}, result interface{}) error {
	if !c.IsAuthenticated() && !strings.Contains(endpoint, "/auth/") {
		return fmt.Errorf("client is not authenticated")
	}

	url, err := c.buildURL(endpoint)
	if err != nil {
		return err
	}

	// Marshal body if provided
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode >= 400 {
		var apiError ApiError
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}

		return fmt.Errorf("API error (status %d): %v %v", resp.StatusCode, apiError.Message, apiError.Error)
	}

	// Decode response
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
