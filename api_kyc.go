package gnosispay

import (
	"net/http"
)

func (c *Client) GetKycIntegration() (*KycIntegration, error) {
	var result KycIntegration
	err := c.doRequest(http.MethodGet, "/api/v1/kyc/integration", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetKycSourceOfFunds() (*[]KycQuestion, error) {
	var result []KycQuestion
	err := c.doRequest(http.MethodGet, "/api/v1/source-of-funds", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PostKycSourceOfFunds(answers []KycAnswer) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/source-of-funds", answers, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PostKycPhoneVerification(phone KycPhoneVerification) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/verification", phone, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PostKycPhoneVerificationCheck(code KycPhoneVerificationCheck) (*ApiGenericResponse, error) {
	var result ApiGenericResponse
	err := c.doRequest(http.MethodPost, "/api/v1/verification/check", code, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PostKycImportPartnerApplicant(args KycImportPartnerApplicant) (*KycImportPartnerApplicantResponse, error) {
	var result KycImportPartnerApplicantResponse
	err := c.doRequest(http.MethodPost, "/api/v1/verification/check", args, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
