package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_country_list.json
var sampleRankingsCountryListResponse string

func TestRankingsCountryList(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCountryListResponse)
	ctx := context.Background()
	resp, err := client.RankingCountryListWithResponse(ctx)
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

	if resp.JSON200.Count == nil || int(*resp.JSON200.Count) != 51 {
		t.Errorf("expected count 51, got %v", resp.JSON200.Count)
	}

	if resp.JSON200.Country == nil {
		t.Fatal("expected country to be non-nil")
	}

	countries := *resp.JSON200.Country
	if len(countries) != 51 {
		t.Errorf("expected 51 country entries, got %d", len(countries))
	}

	// Verify first entry - Argentina
	first := countries[0]
	if first.CountryName == nil || *first.CountryName != "Argentina" {
		t.Errorf("expected first country_name 'Argentina', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "AR" {
		t.Errorf("expected first country_code 'AR', got %v", first.CountryCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 142 {
		t.Errorf("expected first player_count 142, got %v", first.PlayerCount)
	}

	// Verify United States entry (last)
	last := countries[len(countries)-2]
	if last.CountryName == nil || *last.CountryName != "United States" {
		t.Errorf("expected second-to-last country_name 'United States', got %v", last.CountryName)
	}
	if last.CountryCode == nil || *last.CountryCode != "US" {
		t.Errorf("expected second-to-last country_code 'US', got %v", last.CountryCode)
	}
	if last.PlayerCount == nil || int(*last.PlayerCount) != 47944 {
		t.Errorf("expected second-to-last player_count 47944, got %v", last.PlayerCount)
	}

	// Verify all entries have required fields
	for i, c := range countries {
		if c.CountryName == nil || *c.CountryName == "" {
			t.Errorf("countries[%d]: expected country_name to be non-empty", i)
		}
		if c.CountryCode == nil || *c.CountryCode == "" {
			t.Errorf("countries[%d]: expected country_code to be non-empty", i)
		}
		if c.PlayerCount == nil {
			t.Errorf("countries[%d]: expected player_count to be non-nil", i)
		}
	}
}

func TestRankingsCountryListRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCountryListResponse)
	ctx := context.Background()
	_, err := client.RankingCountryListWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/country_list" {
		t.Errorf("expected path '/rankings/country_list', got %s", req.URL.Path)
	}
}
