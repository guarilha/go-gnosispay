package gnosispay

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetAccountBalances() (*AccountBalances, error) {
	var result AccountBalances
	err := c.doRequest(http.MethodGet, "/api/v1/account-balances", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSafeConfig() (*SafeConfig, error) {
	var result SafeConfig
	err := c.doRequest(http.MethodGet, "/api/v1/safe-config", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetDelayedTransactions() (*[]DelayTransaction, error) {
	var result []DelayTransaction
	err := c.doRequest(http.MethodGet, "/api/v1/delay-relay", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

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

func (c *Client) CreateEoa(address common.Address) (*EoaAccount, error) {
	var result EoaAccount
	err := c.doRequest(http.MethodPost, "/api/v1/eoa-accounts", struct{ address string }{address.String()}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteEoa(id string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/eoa-accounts/%s", id), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
