package gnosispay

import (
	"context"
	"net/http"
)

// KYCService handles communication with the KYC related
// methods of the Gnosis Pay API.
type KYCService struct {
	client *Client
}

// GetIntegration retrieves the KYC integration configuration.
func (s *KYCService) GetIntegration(ctx context.Context) (*KycIntegration, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/kyc/integration", nil)
	if err != nil {
		return nil, err
	}

	var integration KycIntegration
	if err := s.client.Do(ctx, req, &integration); err != nil {
		return nil, err
	}

	return &integration, nil
}

// ListSourceOfFunds retrieves the list of KYC source of funds questions.
func (s *KYCService) ListSourceOfFunds(ctx context.Context) ([]KycQuestion, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "/api/v1/source-of-funds", nil)
	if err != nil {
		return nil, err
	}

	var questions []KycQuestion
	if err := s.client.Do(ctx, req, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}

// SubmitSourceOfFunds submits answers to the source of funds questionnaire.
func (s *KYCService) SubmitSourceOfFunds(ctx context.Context, answers []KycAnswer) (*ApiGenericResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/source-of-funds", answers)
	if err != nil {
		return nil, err
	}

	var response ApiGenericResponse
	if err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// InitiatePhoneVerification initiates a phone verification process.
func (s *KYCService) InitiatePhoneVerification(ctx context.Context, phone KycPhoneVerification) (*ApiGenericResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/verification", phone)
	if err != nil {
		return nil, err
	}

	var response ApiGenericResponse
	if err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// VerifyPhone verifies a phone verification code.
func (s *KYCService) VerifyPhone(ctx context.Context, code KycPhoneVerificationCheck) (*ApiGenericResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/verification/check", code)
	if err != nil {
		return nil, err
	}

	var response ApiGenericResponse
	if err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ImportPartnerApplicant imports a KYC applicant from a partner system.
func (s *KYCService) ImportPartnerApplicant(ctx context.Context, args KycImportPartnerApplicant) (*KycImportPartnerApplicantResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/api/v1/kyc/import-partner-applicant", args)
	if err != nil {
		return nil, err
	}

	var response KycImportPartnerApplicantResponse
	if err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
