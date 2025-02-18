package gnosispay

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		httpClient  *http.Client
		opts        []ClientOpt
		wantBaseURL string
		wantErr     bool
	}{
		{
			name:        "default client",
			httpClient:  nil,
			opts:        nil,
			wantBaseURL: defaultBaseURL,
			wantErr:     false,
		},
		{
			name:       "custom base URL",
			httpClient: &http.Client{},
			opts: []ClientOpt{
				SetBaseURL("https://custom.api.com"),
			},
			wantBaseURL: "https://custom.api.com",
			wantErr:     false,
		},
		{
			name:       "invalid base URL",
			httpClient: &http.Client{},
			opts: []ClientOpt{
				SetBaseURL("://invalid.com"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.httpClient, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && client.BaseURL.String() != tt.wantBaseURL {
				t.Errorf("New() BaseURL = %v, want %v", client.BaseURL.String(), tt.wantBaseURL)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	client, _ := New(nil)
	ctx := context.Background()

	tests := []struct {
		name        string
		method      string
		urlStr      string
		body        interface{}
		authToken   string
		wantHeaders http.Header
		wantErr     bool
	}{
		{
			name:   "GET request without body",
			method: "GET",
			urlStr: "/test",
			body:   nil,
			wantHeaders: http.Header{
				"Accept":     []string{"application/json"},
				"User-Agent": []string{userAgent},
			},
			wantErr: false,
		},
		{
			name:      "POST request with body",
			method:    "POST",
			urlStr:    "/test",
			body:      map[string]string{"key": "value"},
			authToken: "test-token",
			wantHeaders: http.Header{
				"Accept":        []string{"application/json"},
				"Content-Type":  []string{"application/json"},
				"User-Agent":    []string{userAgent},
				"Authorization": []string{"Bearer test-token"},
			},
			wantErr: false,
		},
		{
			name:    "invalid URL",
			method:  "GET",
			urlStr:  "://invalid.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.AuthToken = tt.authToken
			req, err := client.NewRequest(ctx, tt.method, tt.urlStr, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			for key, value := range tt.wantHeaders {
				if got := req.Header.Get(key); got != value[0] {
					t.Errorf("NewRequest() header %v = %v, want %v", key, got, value[0])
				}
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type testResponse struct {
		Message string `json:"message"`
	}

	tests := []struct {
		name         string
		statusCode   int
		responseBody interface{}
		wantErr      bool
	}{
		{
			name:         "successful response",
			statusCode:   http.StatusOK,
			responseBody: testResponse{Message: "success"},
			wantErr:      false,
		},
		{
			name:         "error response",
			statusCode:   http.StatusBadRequest,
			responseBody: ApiError{Message: "error message", Error: "error details"},
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				json.NewEncoder(w).Encode(tt.responseBody)
			}))
			defer server.Close()

			client, _ := New(nil, SetBaseURL(server.URL))
			req, _ := client.NewRequest(context.Background(), "GET", "/test", nil)

			var response testResponse
			err := client.Do(context.Background(), req, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_IsAuthenticated(t *testing.T) {
	client, _ := New(nil)

	tests := []struct {
		name      string
		token     string
		want      bool
		setupFunc func() string
	}{
		{
			name:  "empty token",
			token: "",
			want:  false,
		},
		{
			name: "valid token",
			setupFunc: func() string {
				claims := jwt.MapClaims{
					"exp": float64(time.Now().Add(1 * time.Hour).Unix()),
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("test-key"))
				return tokenString
			},
			want: true,
		},
		{
			name: "expired token",
			setupFunc: func() string {
				claims := jwt.MapClaims{
					"exp": float64(time.Now().Add(-1 * time.Hour).Unix()),
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte("test-key"))
				return tokenString
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				tt.token = tt.setupFunc()
			}
			client.AuthToken = tt.token
			if got := client.IsAuthenticated(); got != tt.want {
				t.Errorf("IsAuthenticated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSIWEParams(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		wantHost string
		wantErr  bool
	}{
		{
			name:     "valid URI",
			uri:      "https://example.com/path",
			wantHost: "example.com",
			wantErr:  false,
		},
		{
			name:    "invalid URI",
			uri:     "://invalid.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := New(nil)
			err := SetSIWEParams(tt.uri)(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetSIWEParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && client.Domain != tt.wantHost {
				t.Errorf("SetSIWEParams() domain = %v, want %v", client.Domain, tt.wantHost)
			}
		})
	}
}
