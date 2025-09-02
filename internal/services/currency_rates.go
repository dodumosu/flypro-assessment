package services

import (
	"context"
	"strings"
	"time"

	"flypro-assessment/internal/utils"
)

type CurrencyRateService interface {
	GetUSDRate(ctx context.Context, symbol string) (float64, error)
	LookupSymbol(ctx context.Context, symbol string) (bool, error)
	LoadRates(ctx context.Context) error
}

type currencyRateService struct {
	cache  Cache
	client *utils.CurrencyRatesClient
}

func (c *currencyRateService) LookupSymbol(ctx context.Context, symbol string) (bool, error) {
	symbol = strings.ToUpper(symbol)

	result, err := c.cache.GetSymbol(ctx, symbol)
	if err != nil {
		return false, err
	}

	if result == "" {
		// symbol does not exist, perhaps look it up?
		currenciesMap, err := c.client.GetSymbols()
		if err != nil {
			return false, err
		}

		// store the symbols in the cache
		var symbolFound bool = false
		for sym := range currenciesMap {
			c.cache.SetSymbol(ctx, sym, sym, 6*time.Hour)

			if sym == symbol {
				symbolFound = true
			}
		}

		return symbolFound, nil
	}

	return true, nil
}

func (c *currencyRateService) GetUSDRate(ctx context.Context, symbol string) (float64, error) {
	symbol = strings.ToUpper(symbol)

	symbolPresent, err := c.LookupSymbol(ctx, symbol)
	if err != nil {
		return 0, err
	}

	if !symbolPresent {
		return 0, nil
	}

	rate, err := c.cache.GetUSDRate(ctx, symbol)
	if err != nil {
		return 0, err
	}

	if rate == 0 {
		err = c.LoadRates(ctx)
		if err != nil {
			return 0, err
		}

		// retry
		rate, err = c.cache.GetUSDRate(ctx, symbol)
		if err != nil {
			return 0, err
		}
	}

	return rate, nil
}

func (c *currencyRateService) LoadRates(ctx context.Context) error {
	rates, err := c.client.GetCurrentRates()
	if err != nil {
		return err
	}

	for symbol, rate := range rates {
		err = c.cache.SetSymbol(ctx, symbol, symbol, 24*time.Hour)
		if err != nil {
			return err
		}
		err = c.cache.SetUSDRate(ctx, symbol, rate, 6*time.Hour)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewCurrencyRatesService(cache Cache, client *utils.CurrencyRatesClient) CurrencyRateService {
	return &currencyRateService{
		cache:  cache,
		client: client,
	}
}
