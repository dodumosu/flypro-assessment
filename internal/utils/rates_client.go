package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"flypro-assessment/internal/config"

	"github.com/go-resty/resty/v2"
)

const RATES_URL = "https://api.apilayer.com/exchangerates_data/latest"
const SYMBOLS_URL = "https://api.apilayer.com/exchangerates_data/symbols"

type CurrencyRatesClient struct {
	apiKey     string
	client     *resty.Client
	ratesURL   string
	symbolsURL string
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0) // Convert Unix timestamp to time.Time
	return nil
}

type CurrencyRatesResponse struct {
	Success   bool               `json:"success"`
	Timestamp UnixTime           `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

type SymbolsResponse struct {
	Success bool              `json:"success"`
	Symbols map[string]string `json:"symbols"`
}

func NewCurrencyRatesClient(cfg config.RateAPIConfig) *CurrencyRatesClient {
	return &CurrencyRatesClient{
		apiKey:     cfg.APIKey,
		client:     resty.New(),
		ratesURL:   RATES_URL,
		symbolsURL: SYMBOLS_URL,
	}
}

func (c *CurrencyRatesClient) GetCurrentRates() (map[string]float64, error) {
	resp, err := c.client.R().SetHeader(
		"apikey", c.apiKey,
	).SetQueryParam(
		"base", "USD",
	).Get(c.ratesURL)

	if err != nil {
		return nil, fmt.Errorf("currency rates retrieval failed: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("currency rates retrieval failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	var apiResponse CurrencyRatesResponse
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		return nil, fmt.Errorf("parsing currency rates API response failed: %w", err)
	}

	if !apiResponse.Success {
		return nil, fmt.Errorf("currency rates API call did not succeed")
	}

	return apiResponse.Rates, nil
}

func (c *CurrencyRatesClient) GetSymbols() (map[string]string, error) {
	resp, err := c.client.R().SetHeader(
		"apikey", c.apiKey,
	).Get(c.symbolsURL)

	if err != nil {
		return nil, fmt.Errorf("symbols retrieval failed: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("symbols retrieval failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	var apiResponse SymbolsResponse
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		return nil, fmt.Errorf("parsing symbols API response failed: %w", err)
	}

	if !apiResponse.Success {
		return nil, fmt.Errorf("symbols API call did not succeed")
	}

	return apiResponse.Symbols, nil
}
