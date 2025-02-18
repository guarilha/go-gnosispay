package gnosispay

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// GetAccountBalances retrieves the account balances from the API.
// Returns an AccountBalances object containing balance information or an error if the request fails.
func (c *Client) GetAccountBalances() (*AccountBalances, error) {
	var result AccountBalances
	err := c.doRequest(http.MethodGet, "/api/v1/account-balances", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSafeConfig retrieves the Safe wallet configuration from the API.
// Returns a SafeConfig object containing wallet settings or an error if the request fails.
func (c *Client) GetSafeConfig() (*SafeConfig, error) {
	var result SafeConfig
	err := c.doRequest(http.MethodGet, "/api/v1/safe-config", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDelayedTransactions retrieves a list of delayed transactions from the API.
// Returns an array of DelayTransaction objects or an error if the request fails.
func (c *Client) GetDelayedTransactions() (*[]DelayTransaction, error) {
	var result []DelayTransaction
	err := c.doRequest(http.MethodGet, "/api/v1/delay-relay", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEoaAccounts retrieves a list of externally owned accounts (EOAs) from the API.
// Returns an array of EoaAccount objects or an error if the request fails.
func (c *Client) GetEoaAccounts() (*[]EoaAccount, error) {
	var result struct {
		Data struct {
			EoaAccounts []EoaAccount `json:"eoaAccounts"`
		} `json:"data"`
	}
	err := c.doRequest(http.MethodGet, "/api/v1/eoa-accounts", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data.EoaAccounts, nil
}

// CreateEoa creates a new externally owned account (EOA) with the specified Ethereum address.
// Parameters:
//   - address: Ethereum address to associate with the new EOA
//
// Returns the created EoaAccount object or an error if the creation fails.
func (c *Client) CreateEoa(address common.Address) (*EoaAccount, error) {
	var result EoaAccount
	err := c.doRequest(http.MethodPost, "/api/v1/eoa-accounts", struct{ address string }{address.String()}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteEoa removes an externally owned account (EOA) with the specified ID.
// Parameters:
//   - id: Unique identifier of the EOA to delete
//
// Returns true if deletion was successful, false otherwise, and any error that occurred.
func (c *Client) DeleteEoa(id string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/eoa-accounts/%s", id), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
