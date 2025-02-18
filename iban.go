package gnosispay

import (
	"context"
	"net/http"
)

// IBANService handles communication with the IBAN related
// methods of the Gnosis Pay API.
type IBANService struct {
	client *Client
}

// CheckAvailability checks if IBAN services are available for the authenticated user.
func (s *IBANService) CheckAvailability(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/ibans/available", nil)
	if err != nil {
		return false, err
	}

	err = s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Activate activates IBAN services for the authenticated user.
func (s *IBANService) Activate(ctx context.Context) error {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/ibans/monerium-profile", nil)
	if err != nil {
		return err
	}

	return s.client.Do(ctx, req, nil)
}

// GetDetails retrieves the IBAN details for the authenticated user.
func (s *IBANService) GetDetails(ctx context.Context) (*IbanDetails, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/ibans/details", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data IbanDetails `json:"data"`
	}
	if err := s.client.Do(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

// ListOrders retrieves the history of IBAN-related orders.
func (s *IBANService) ListOrders(ctx context.Context) ([]IbanOrder, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/ibans/orders", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []IbanOrder `json:"data"`
	}
	if err := s.client.Do(ctx, req, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
