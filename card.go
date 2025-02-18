package gnosispay

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CardService handles communication with the card related
// methods of the Gnosis Pay API.
type CardService struct {
	client *Client
}

// ListTransactionsOptions specifies the optional parameters to the
// CardService.ListTransactions method.
type ListTransactionsOptions struct {
	CardTokens          string `url:"cardTokens,omitempty"`
	Before              string `url:"before,omitempty"`
	After               string `url:"after,omitempty"`
	BillingCurrency     string `url:"billingCurrency,omitempty"`
	TransactionCurrency string `url:"transactionCurrency,omitempty"`
	MCC                 string `url:"mcc,omitempty"`
}

// List returns all cards associated with the authenticated user.
func (s *CardService) List(ctx context.Context) ([]Card, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/cards", nil)
	if err != nil {
		return nil, err
	}

	var cards []Card
	if err := s.client.Do(ctx, req, &cards); err != nil {
		return nil, err
	}

	return cards, nil
}

// GetStatus retrieves the current status of a card.
func (s *CardService) GetStatus(ctx context.Context, cardID string) (*CardStatus, error) {
	path := fmt.Sprintf("/api/v1/cards/%s/status", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var status CardStatus
	if err := s.client.Do(ctx, req, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// Activate activates a card.
func (s *CardService) Activate(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/api/v1/cards/%s/activate", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// Freeze temporarily freezes a card.
func (s *CardService) Freeze(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/api/v1/cards/%s/freeze", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// Unfreeze unfreezes a previously frozen card.
func (s *CardService) Unfreeze(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/api/v1/cards/%s/unfreeze", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// ReportLost reports a card as lost.
func (s *CardService) ReportLost(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/api/v1/cards/%s/lost", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// ReportStolen reports a card as stolen.
func (s *CardService) ReportStolen(ctx context.Context, cardID string) error {
	path := fmt.Sprintf("/api/v1/cards/%s/stolen", cardID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// ListTransactions retrieves card transactions based on the provided options.
func (s *CardService) ListTransactions(ctx context.Context, opts *ListTransactionsOptions) ([]CardEvent, error) {
	path := "/transactions"
	if opts != nil {
		v := url.Values{}
		if opts.CardTokens != "" {
			v.Set("cardTokens", opts.CardTokens)
		}
		if opts.Before != "" {
			v.Set("before", opts.Before)
		}
		if opts.After != "" {
			v.Set("after", opts.After)
		}
		if opts.BillingCurrency != "" {
			v.Set("billingCurrency", opts.BillingCurrency)
		}
		if opts.TransactionCurrency != "" {
			v.Set("transactionCurrency", opts.TransactionCurrency)
		}
		if opts.MCC != "" {
			v.Set("mcc", opts.MCC)
		}

		if len(v) > 0 {
			path = fmt.Sprintf("%s?%s", path, v.Encode())
		}
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var events []CardEvent
	if err := s.client.Do(ctx, req, &events); err != nil {
		return nil, err
	}

	return events, nil
}
