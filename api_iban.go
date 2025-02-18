package gnosispay

import (
	"net/http"
)

// GetIbanAvailability checks if IBAN services are available for the authenticated user.
func (c *Client) GetIbanAvailability() (bool, error) {
	err := c.doRequest(http.MethodGet, "/api/v1/ibans/available", nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// PostActivateIban activates IBAN services for the authenticated user.
// Returns true if activation was successful.
func (c *Client) PostActivateIban() (bool, error) {
	err := c.doRequest(http.MethodPost, "/api/v1/ibans/monerium-profile", nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetIbanDetails retrieves the IBAN details for the authenticated user.
func (c *Client) GetIbanDetails() (*IbanDetails, error) {
	var result struct {
		Data IbanDetails `json:"data"`
	}
	err := c.doRequest(http.MethodGet, "/api/v1/ibans/details", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// GetIbanOrderHistory retrieves the history of IBAN-related orders.
func (c *Client) GetIbanOrderHistory() (*[]IbanOrder, error) {
	var result struct {
		Data []IbanOrder `json:"data"`
	}
	err := c.doRequest(http.MethodGet, "/api/v1/ibans/orders", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}
