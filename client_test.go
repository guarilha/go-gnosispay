package gnosispay

import (
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		domain  string
		uri     string
		chainID int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid configuration",
			baseURL: "https://api.example.com",
			uri:     "https://api.example.com",
			domain:  "api.example.com",
			wantErr: false,
		},
		{
			name:    "empty base URL",
			baseURL: "",
			uri:     "https://api.example.com",
			domain:  "api.example.com",
			wantErr: true,
			errMsg:  "baseURL and yourAppUri cannot be empty",
		},
		{
			name:    "invalid base URL",
			baseURL: "://invalid",
			uri:     "https://api.example.com",
			domain:  "api.example.com",
			wantErr: true,
			errMsg:  "invalid base URL",
		},
		{
			name:    "empty base URI",
			baseURL: "https://api.example.com",
			uri:     "",
			domain:  "api.example.com",
			wantErr: true,
			errMsg:  "baseURL and yourAppUri cannot be empty",
		},
		{
			name:    "invalid base URI",
			baseURL: "https://api.example.com",
			uri:     "://invalid",
			domain:  "api.example.com",
			wantErr: true,
			errMsg:  "invalid application URI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.baseURL, tt.uri)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewClient() error = nil, want error containing %q", tt.errMsg)
				} else if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("NewClient() error = %v, want error containing %q", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("NewClient() unexpected error = %v", err)
				return
			}

			if client.Domain != tt.domain {
				t.Errorf("NewClient() domain = %v, want %v", client.Domain, tt.domain)
			}
			if client.Uri != tt.uri {
				t.Errorf("NewClient() uri = %v, want %v", client.Uri, tt.uri)
			}
		})
	}
}

func TestIsAuthenticated(t *testing.T) {
	client := &Client{}

	if client.IsAuthenticated() {
		t.Error("IsAuthenticated() = true, want false for empty token")
	}

}
