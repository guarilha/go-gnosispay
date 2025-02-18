package gnosispay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	defaultBaseURL = "https://api.gnosispay.com"
	userAgent      = "go-gnosispay"
)

// Client manages communication with Gnosis Pay API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the Gnosis Pay API.
	UserAgent string

	// Auth token for API requests
	AuthToken string

	// Domain and URI for SIWE authentication
	Domain string
	Uri    string

	// Chain ID for the network
	ChainID int

	// Services used for communicating with different parts of the Gnosis Pay API.
	Auth    *AuthService
	User    *UserService
	Cards   *CardService
	KYC     *KYCService
	IBAN    *IBANService
	Account *AccountManagementService
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New returns a new Gnosis Pay API client.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		ChainID:   100, // Gnosis Chain ID (mainnet)
	}

	// Apply any given client options.
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// Create all the services.
	c.Auth = &AuthService{client: c}
	c.User = &UserService{client: c}
	c.Cards = &CardService{client: c}
	c.KYC = &KYCService{client: c}
	c.IBAN = &IBANService{client: c}
	c.Account = &AccountManagementService{client: c}

	return c, nil
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = ua
		return nil
	}
}

// SetAuthToken is a client option for setting the authentication token.
func SetAuthToken(token string) ClientOpt {
	return func(c *Client) error {
		c.AuthToken = token
		return nil
	}
}

// SetSIWEParams is a client option for setting the SIWE authentication parameters.
func SetSIWEParams(uri string) ClientOpt {
	return func(c *Client) error {
		parsedURI, err := url.Parse(uri)
		if err != nil {
			return fmt.Errorf("invalid application URI: %w", err)
		}

		c.Domain = parsedURI.Host
		c.Uri = uri
		return nil
	}
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Debug: Print response body
	slog.Debug("Response Status", "status", resp.StatusCode)
	slog.Debug("Response Body", "body", string(bodyBytes))

	if resp.StatusCode >= 400 {
		var apiError ApiError
		if err := json.Unmarshal(bodyBytes, &apiError); err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}
		return fmt.Errorf("API error (status %d): %v %v", resp.StatusCode, apiError.Message, apiError.Error)
	}

	if v != nil {
		if str, ok := v.(*string); ok {
			*str = string(bodyBytes)
		} else if err := json.Unmarshal(bodyBytes, v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// IsAuthenticated checks if the client has a valid, non-expired authentication token.
func (c *Client) IsAuthenticated() bool {
	if c.AuthToken == "" {
		return false
	}

	expired, err := isTokenExpired(c.AuthToken)
	if err != nil || expired {
		return false
	}

	return true
}

func isTokenExpired(tokenString string) (bool, error) {
	// Parse the token
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return true, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true, fmt.Errorf("invalid token claims")
	}

	// Check expiration time
	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		return expirationTime.Before(time.Now()), nil
	}

	return true, fmt.Errorf("no expiration claim found")
}
