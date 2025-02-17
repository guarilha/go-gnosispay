package gnosispay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
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

func NewClient(baseURL, authToken, domain, uri string, chainID int) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	return &Client{
		BaseURL: parsedURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		AuthToken: authToken,
		Domain:    domain,
		Uri:       uri,
		ChainID:   chainID,
	}, nil
}

// buildURL ensures safe URL concatenation
func (c *Client) buildURL(endpoint string) (string, error) {
	relPath, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("invalid endpoint path: %w", err)
	}

	fullURL := c.BaseURL.ResolveReference(relPath)
	return fullURL.String(), nil
}

// doRequest handles API requests with authentication
func (c *Client) doRequest(method, endpoint string, body interface{}, result interface{}) error {
	url, err := c.buildURL(endpoint)
	if err != nil {
		return err
	}

	slog.Info("gnosispay.doRequest", "method", method, "endpoint", endpoint, "url", url)

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
		fmt.Println()
		return fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	// Decode response
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
