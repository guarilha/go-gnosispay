package gnosispay

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetCardStatus(cardId string) (*CardStatus, error) {
	var result CardStatus
	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/v1/cards/%s/status", cardId), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetCards() (*[]Card, error) {
	var result []Card
	err := c.doRequest(http.MethodGet, "/api/v1/cards", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PostActivateCard(cardId string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/cards/%s/activate", cardId), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) PostFreezeCard(cardId string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/cards/%s/freeze", cardId), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) PostReportLostCard(cardId string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/cards/%s/lost", cardId), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) PostReportStolenCard(cardId string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/cards/%s/stolen", cardId), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) PostUnfreezeCard(cardId string) (bool, error) {
	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/v1/cards/%s/unfreeze", cardId), nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) GetTransactions(filters *GetTransactionsFilters) (*[]CardEvent, error) {
	var result []CardEvent
	err := c.doRequest(http.MethodGet, getTransactionsURL(filters), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func getTransactionsURL(filters *GetTransactionsFilters) string {
	if filters == nil {
		return "/transactions"
	}
	queryParams := url.Values{}

	if filters.CardTokens != nil && *filters.CardTokens != "" {
		queryParams.Set("cardTokens", *filters.CardTokens)
	}
	if filters.Before != nil && *filters.Before != "" {
		queryParams.Set("before", *filters.Before)
	}
	if filters.After != nil && *filters.After != "" {
		queryParams.Set("after", *filters.After)
	}
	if filters.BillingCurrency != nil && *filters.BillingCurrency != "" {
		queryParams.Set("billingCurrency", *filters.BillingCurrency)
	}
	if filters.TransactionCurrency != nil && *filters.TransactionCurrency != "" {
		queryParams.Set("transactionCurrency", *filters.TransactionCurrency)
	}
	if filters.MCC != nil && *filters.MCC != "" {
		queryParams.Set("mcc", *filters.MCC)
	}

	queryString := queryParams.Encode()
	if queryString == "" {
		return "/transactions"
	}
	return fmt.Sprintf("/transactions?%s", queryString)
}
