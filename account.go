package gnosispay

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// AccountManagementService handles communication with the account management related
// methods of the Gnosis Pay API.
type AccountManagementService struct {
	client *Client
}

// GetBalances retrieves the account balances.
func (s *AccountManagementService) GetBalances(ctx context.Context) (*AccountBalances, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/account-balances", nil)
	if err != nil {
		return nil, err
	}

	var balances AccountBalances
	if err := s.client.Do(ctx, req, &balances); err != nil {
		return nil, err
	}

	return &balances, nil
}

// GetSafeConfig retrieves the Safe wallet configuration.
func (s *AccountManagementService) GetSafeConfig(ctx context.Context) (*SafeConfig, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/safe-config", nil)
	if err != nil {
		return nil, err
	}

	var config SafeConfig
	if err := s.client.Do(ctx, req, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// ListDelayedTransactions retrieves a list of delayed transactions.
func (s *AccountManagementService) ListDelayedTransactions(ctx context.Context) ([]DelayTransaction, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/delay-relay", nil)
	if err != nil {
		return nil, err
	}

	var transactions []DelayTransaction
	if err := s.client.Do(ctx, req, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

// ListEoaAccounts retrieves a list of externally owned accounts (EOAs).
func (s *AccountManagementService) ListEoaAccounts(ctx context.Context) ([]EoaAccount, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/eoa-accounts", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			EoaAccounts []EoaAccount `json:"eoaAccounts"`
		} `json:"data"`
	}
	if err := s.client.Do(ctx, req, &result); err != nil {
		return nil, err
	}

	return result.Data.EoaAccounts, nil
}

// CreateEoa creates a new externally owned account (EOA).
func (s *AccountManagementService) CreateEoa(ctx context.Context, address common.Address) (*EoaAccount, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/eoa-accounts", struct {
		Address string `json:"address"`
	}{
		Address: address.String(),
	})
	if err != nil {
		return nil, err
	}

	var account EoaAccount
	if err := s.client.Do(ctx, req, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// DeleteEoa removes an externally owned account (EOA).
func (s *AccountManagementService) DeleteEoa(ctx context.Context, id string) error {
	path := fmt.Sprintf("/api/v1/eoa-accounts/%s", id)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}
