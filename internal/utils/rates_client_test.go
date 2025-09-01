package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"flypro-assessment/internal/config"
)

func TestNewRatesClient(t *testing.T) {
	cfg := config.RateAPIConfig{
		APIKey: "new-api-key",
	}

	client := NewCurrencyRatesClient(cfg)
	if client == nil {
		t.Error("NewCurrencyRatesClient should return a non-nil client")
	}
}

func TestUnixTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid timestamp",
			input:   []byte("1609459200"), // 2021-01-01 00:00:00 UTC
			want:    time.Unix(1609459200, 0).UTC(),
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   []byte("\"invalid\""),
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u UnixTime
			err := u.UnmarshalJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !u.Time.Equal(tt.want) {
				t.Errorf("UnixTime.UnmarshalJSON() = %v, want %v", u.Time, tt.want)
			}
		})
	}
}

func TestCurrencyRatesClient_GetCurrentRates(t *testing.T) {
	// Mock server for successful response
	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers
		if r.Header.Get("apikey") != "test-api-key" {
			t.Errorf("Expected apikey header 'test-api-key', got '%s'", r.Header.Get("apikey"))
		}

		// Verify query parameters
		if r.URL.Query().Get("base") != "USD" {
			t.Errorf("Expected base query parameter 'USD', got '%s'", r.URL.Query().Get("base"))
		}

		// Return mock response
		response := map[string]interface{}{
			"success":   true,
			"timestamp": 1609459200,
			"base":      "USD",
			"rates": map[string]float64{
				"EUR": 0.82,
				"GBP": 0.73,
				"JPY": 103.50,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer successServer.Close()

	// Mock server for HTTP error
	httpErrorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer httpErrorServer.Close()

	// Mock server for API error (success=false)
	apiErrorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"success": false,
			"rates":   nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer apiErrorServer.Close()

	// Mock server for invalid JSON
	invalidJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{ invalid json }"))
	}))
	defer invalidJSONServer.Close()

	tests := []struct {
		name          string
		serverURL     string
		want          map[string]float64
		wantErr       bool
		errorContains string
	}{
		{
			name:      "successful response",
			serverURL: successServer.URL,
			want: map[string]float64{
				"EUR": 0.82,
				"GBP": 0.73,
				"JPY": 103.50,
			},
			wantErr: false,
		},
		{
			name:          "HTTP error",
			serverURL:     httpErrorServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "currency rates retrieval failed with status 500",
		},
		{
			name:          "API error",
			serverURL:     apiErrorServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "currency rates API call did not succeed",
		},
		{
			name:          "invalid JSON response",
			serverURL:     invalidJSONServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "parsing currency rates API response failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.RateAPIConfig{
				APIKey: "test-api-key",
			}
			client := NewCurrencyRatesClient(cfg)
			// Override the URL for testing
			client.ratesURL = tt.serverURL

			got, err := client.GetCurrentRates()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyRatesClient.GetCurrentRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("CurrencyRatesClient.GetCurrentRates() error = %v, want error containing %v", err, tt.errorContains)
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyRatesClient.GetCurrentRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyRatesClient_GetSymbols(t *testing.T) {
	// Mock server for successful response
	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers
		if r.Header.Get("apiKey") != "test-api-key" {
			t.Errorf("Expected apiKey header 'test-api-key', got '%s'", r.Header.Get("apiKey"))
		}

		// Return mock response
		response := map[string]interface{}{
			"success": true,
			"symbols": map[string]string{
				"USD": "United States Dollar",
				"EUR": "Euro",
				"GBP": "British Pound",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer successServer.Close()

	// Mock server for HTTP error
	httpErrorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer httpErrorServer.Close()

	// Mock server for API error (success=false)
	apiErrorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"success": false,
			"symbols": nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer apiErrorServer.Close()

	// Mock server for invalid JSON
	invalidJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{ invalid json }"))
	}))
	defer invalidJSONServer.Close()

	tests := []struct {
		name          string
		serverURL     string
		want          map[string]string
		wantErr       bool
		errorContains string
	}{
		{
			name:      "successful response",
			serverURL: successServer.URL,
			want: map[string]string{
				"USD": "United States Dollar",
				"EUR": "Euro",
				"GBP": "British Pound",
			},
			wantErr: false,
		},
		{
			name:          "HTTP error",
			serverURL:     httpErrorServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "symbols retrieval failed with status 500",
		},
		{
			name:          "API error",
			serverURL:     apiErrorServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "symbols API call did not succeed",
		},
		{
			name:          "invalid JSON response",
			serverURL:     invalidJSONServer.URL,
			want:          nil,
			wantErr:       true,
			errorContains: "parsing symbols API response failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.RateAPIConfig{
				APIKey: "test-api-key",
			}
			client := NewCurrencyRatesClient(cfg)
			// Override the URL for testing
			client.symbolsURL = tt.serverURL

			got, err := client.GetSymbols()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrencyRatesClient.GetSymbols() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("CurrencyRatesClient.GetSymbols() error = %v, want error containing %v", err, tt.errorContains)
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrencyRatesClient.GetSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
