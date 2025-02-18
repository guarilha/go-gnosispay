package gnosispay

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/guarilha/go-gnosispay/wallet"
	"github.com/spruceid/siwe-go"
)

// AuthService handles communication with the authentication related
// methods of the Gnosis Pay API.
type AuthService struct {
	client *Client
}

// SignUpRequest represents a request to sign up a new user.
type SignUpRequest struct {
	AuthEmail string `json:"authEmail"`
}

// GetNonce retrieves a new nonce from the server for use in SIWE authentication.
func (s *AuthService) GetNonce(ctx context.Context) (string, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/auth/nonce", nil)
	if err != nil {
		return "", err
	}

	var nonce string
	if err := s.client.Do(ctx, req, &nonce); err != nil {
		return "", err
	}

	return nonce, nil
}

// GetSIWEMessage generates a Sign-In with Ethereum (SIWE) message for the given address.
func (s *AuthService) GetSIWEMessage(ctx context.Context, address common.Address) (string, error) {
	nonce, err := s.GetNonce(ctx)
	if err != nil {
		return "", err
	}

	msg, err := siwe.InitMessage(
		s.client.Domain,
		address.String(),
		s.client.Uri,
		nonce,
		map[string]any{
			"chainId": s.client.ChainID,
		},
	)
	if err != nil {
		return "", err
	}
	return msg.String(), nil
}

// GetAuthToken obtains an authentication token by submitting a signed SIWE message.
func (s *AuthService) GetAuthToken(ctx context.Context, message, signature string) (string, error) {
	type authRequest struct {
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	type authResponse struct {
		Token string `json:"token"`
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/auth/challenge", authRequest{
		Message:   message,
		Signature: signature,
	})
	if err != nil {
		return "", err
	}

	var resp authResponse
	if err := s.client.Do(ctx, req, &resp); err != nil {
		return "", err
	}
	s.client.AuthToken = resp.Token
	return resp.Token, nil
}

// SignUp registers a new user with the provided email address.
func (s *AuthService) SignUp(ctx context.Context, email string) (*SignUpResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/auth/signup", SignUpRequest{
		AuthEmail: email,
	})
	if err != nil {
		return nil, err
	}

	var resp SignUpResponse
	if err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	s.client.AuthToken = resp.Token
	return &resp, nil
}

// AuthenticateWithPrivateKey performs authentication using an Ethereum private key.
func (s *AuthService) AuthenticateWithPrivateKey(ctx context.Context, address common.Address, privateKey *ecdsa.PrivateKey) (string, error) {
	message, err := s.GetSIWEMessage(ctx, address)
	if err != nil {
		return "", err
	}

	signedMessage, err := wallet.SignMessage(message, privateKey)
	if err != nil {
		return "", err
	}

	return s.GetAuthToken(ctx, message, wallet.SignatureToString(signedMessage))
}
