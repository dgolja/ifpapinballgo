package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/other_countries.json
var sampleOtherCountriesResponse string

func TestOtherCountriesWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleOtherCountriesResponse)
	ctx := context.Background()
	resp, err := client.OtherCountriesWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}
	if resp.JSON200.Country == nil {
		t.Fatal("expected Country to be non-nil")
	}

	countries := *resp.JSON200.Country
	if len(countries) != 62 {
		t.Errorf("expected 62 countries, got %d", len(countries))
	}

	// Verify first country (Argentina)
	first := countries[0]
	if first.CountryId == nil || int(*first.CountryId) != 5 {
		t.Errorf("expected first country_id 5, got %v", first.CountryId)
	}
	if first.CountryName == nil || *first.CountryName != "Argentina" {
		t.Errorf("expected first country_name 'Argentina', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "AR" {
		t.Errorf("expected first country_code 'AR', got %v", first.CountryCode)
	}
	if first.ActiveFlag == nil || *first.ActiveFlag != "Y" {
		t.Errorf("expected first active_flag 'Y', got %v", first.ActiveFlag)
	}

	// Verify all countries have required fields populated
	for i, c := range countries {
		if c.CountryId == nil {
			t.Errorf("country[%d]: expected country_id to be non-nil", i)
		}
		if c.CountryName == nil || *c.CountryName == "" {
			t.Errorf("country[%d]: expected country_name to be non-empty", i)
		}
		if c.CountryCode == nil || *c.CountryCode == "" {
			t.Errorf("country[%d]: expected country_code to be non-empty", i)
		}
	}
}

func TestOtherCountriesRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleOtherCountriesResponse)
	ctx := context.Background()
	_, err := client.OtherCountriesWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/other/countries" {
		t.Errorf("expected path '/other/countries', got %s", req.URL.Path)
	}
}
