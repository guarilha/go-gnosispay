package gnosispay

import (
	"context"
	"net/http"
)

// GetUser retrieves the authenticated user's profile information.
func (c *Client) GetUser() (*User, error) {
	var result User
	err := c.doRequest(http.MethodGet, "/api/v1/user", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserReferrals retrieves the user's referral information.
func (c *Client) GetUserReferrals(ctx context.Context) (*UserReferrals, error) {
	var result UserReferrals
	err := c.doRequest(http.MethodGet, "/api/v1/user/referrals", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateReferralCode generates a new referral code for the authenticated user.
func (c *Client) CreateReferralCode(ctx context.Context) (*UserReferralCode, error) {
	var result UserReferralCode
	err := c.doRequest(http.MethodPost, "/api/v1/user/referrer-code", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
