package gnosispay

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/guarilha/go-gnosispay/wallet"
	"github.com/spruceid/siwe-go"
)

func fetchPlainText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func (c *Client) GetNonce() (*string, error) {
	url, err := c.buildURL("/api/v1/auth/nonce")
	if err != nil {
		return nil, err
	}

	nonce, err := fetchPlainText(url)
	if err != nil {
		return nil, err
	}

	return &nonce, nil
}

func (c *Client) AuthenticateWithPrivateKey(address common.Address, privateKey *ecdsa.PrivateKey) (*string, error) {
	message, err := c.GetSIWEMessage(address)
	if err != nil {
		return nil, err
	}

	signedMessage, err := wallet.SignMessage(message, privateKey)
	if err != nil {
		return nil, err
	}

	return c.GetAuthToken(message, wallet.SignatureToString(signedMessage))
}

func (c *Client) GetSIWEMessage(address common.Address) (string, error) {
	nonce, err := c.GetNonce()
	if err != nil {
		return "", err
	}

	msg, err := siwe.InitMessage(
		c.Domain,
		address.String(),
		c.Uri,
		*nonce,
		map[string]any{
			"chainId": c.ChainID,
		},
	)
	if err != nil {
		return "", err
	}

	return msg.String(), nil
}

func (c *Client) GetAuthToken(message string, signature string) (*string, error) {
	var jwt struct {
		Token string `json:"token"`
	}
	err := c.doRequest(http.MethodPost, "/api/v1/auth/challenge", map[string]string{
		"message":   message,
		"signature": signature,
	}, &jwt)
	if err != nil {
		return nil, err
	}

	return &jwt.Token, nil
}

func (c *Client) SignUp(ctx context.Context, email string) (*SignUpResponse, error) {
	var resp SignUpResponse
	err := c.doRequest(http.MethodPost, "/api/v1/auth/signup", map[string]string{
		"authEmail": email,
	}, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
