package gnosispay

import (
	"net/http"
)

// GetKycIntegration retrieves the KYC integration configuration from the API.
// Returns a KycIntegration object containing integration settings or an error if the request fails.
func (c *Client) GetKycIntegration() (*KycIntegration, error) {
	var result KycIntegration
	err := c.doRequest(http.MethodGet, "/api/v1/kyc/integration", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetKycSourceOfFunds retrieves the list of KYC source of funds questions from the API.
// Returns an array of KycQuestion objects or an error if the request fails.
func (c *Client) GetKycSourceOfFunds() (*[]KycQuestion, error) {
	var result []KycQuestion
	err := c.doRequest(http.MethodGet, "/api/v1/source-of-funds", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PostKycSourceOfFunds submits answers to the source of funds questionnaire.
// Parameters:
//   - answers: Array of KycAnswer objects containing user responses
//
// Returns an ApiGenericResponse or an error if the submission fails.
func (c *Client) PostKycSourceOfFunds(answers []KycAnswer) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/source-of-funds", answers, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PostKycPhoneVerification initiates a phone verification process.
// Parameters:
//   - phone: KycPhoneVerification object containing the phone number to verify
//
// Returns an ApiGenericResponse or an error if the verification initiation fails.
func (c *Client) PostKycPhoneVerification(phone KycPhoneVerification) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/verification", phone, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PostKycPhoneVerificationCheck verifies a phone verification code.
// Parameters:
//   - code: KycPhoneVerificationCheck object containing the verification code
//
// Returns an ApiGenericResponse or an error if the code verification fails.
func (c *Client) PostKycPhoneVerificationCheck(code KycPhoneVerificationCheck) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/verification/check", code, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PostKycImportPartnerApplicant imports a KYC applicant from a partner system.
// Parameters:
//   - args: KycImportPartnerApplicant object containing the applicant details
//
// Returns a KycImportPartnerApplicantResponse or an error if the import fails.
func (c *Client) PostKycImportPartnerApplicant(args KycImportPartnerApplicant) (*KycImportPartnerApplicantResponse, error) {
	var result KycImportPartnerApplicantResponse
	err := c.doRequest(http.MethodPost, "/api/v1/kyc/import-partner-applicant", args, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
