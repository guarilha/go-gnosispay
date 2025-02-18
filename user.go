package gnosispay

import (
	"context"
	"net/http"
)

// UserService handles communication with the user related
// methods of the Gnosis Pay API.
type UserService struct {
	client *Client
}

// Get retrieves the authenticated user's profile information.
func (s *UserService) Get(ctx context.Context) (*User, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/user", nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := s.client.Do(ctx, req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetReferrals retrieves the user's referral information.
func (s *UserService) GetReferrals(ctx context.Context) (*UserReferrals, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/user/referrals", nil)
	if err != nil {
		return nil, err
	}

	var referrals UserReferrals
	if err := s.client.Do(ctx, req, &referrals); err != nil {
		return nil, err
	}

	return &referrals, nil
}

// CreateReferralCode generates a new referral code for the authenticated user.
func (s *UserService) CreateReferralCode(ctx context.Context) (*UserReferralCode, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/user/referrer-code", nil)
	if err != nil {
		return nil, err
	}

	var code UserReferralCode
	if err := s.client.Do(ctx, req, &code); err != nil {
		return nil, err
	}

	return &code, nil
}
