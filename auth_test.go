package gnosispay

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createTestToken(exp int64) string {
	// Create a test token with given expiration
	claims := jwt.MapClaims{
		"exp": float64(exp),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a test key (we don't verify signature in isTokenExpired)
	tokenString, _ := token.SignedString([]byte("test-key"))
	return tokenString
}

func TestIsTokenExpired(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		wantExpired bool
		wantErr     bool
	}{
		{
			name:        "expired token",
			token:       createTestToken(time.Now().Add(-2 * time.Hour).Unix()),
			wantExpired: true,
			wantErr:     false,
		},
		{
			name:        "valid token",
			token:       createTestToken(time.Now().Add(1 * time.Hour).Unix()),
			wantExpired: false,
			wantErr:     false,
		},
		{
			name:        "malformed token",
			token:       "invalid-token",
			wantExpired: true,
			wantErr:     true,
		},
		{
			name:        "empty token",
			token:       "",
			wantExpired: true,
			wantErr:     true,
		},
		{
			name: "token without exp claim",
			token: func() string {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
				tokenString, _ := token.SignedString([]byte("test-key"))
				return tokenString
			}(),
			wantExpired: true,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExpired, err := isTokenExpired(tt.token)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("isTokenExpired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check expiration result
			if gotExpired != tt.wantExpired {
				t.Errorf("isTokenExpired() = %v, want %v", gotExpired, tt.wantExpired)
			}
		})
	}
}
