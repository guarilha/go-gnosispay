package gnosispay

import (
	"net/http"
)

func (c *Client) GetIbanAvailability() (bool, error) {
	err := c.doRequest(http.MethodGet, "/api/v1/ibans/available", nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) PostActivateIban() (bool, error) {
	err := c.doRequest(http.MethodPost, "/api/v1/ibans/monerium-profile", nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

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
