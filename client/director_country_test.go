package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/director_country.json
var sampleDirectorCountryResponse string

func TestDirectorViewCountryDirectors(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorCountryResponse)
	ctx := context.Background()
	resp, err := client.ViewCountryDirectorsWithResponse(ctx)
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

	if *resp.JSON200.Count != 35 {
		t.Errorf("expected count 35, got %v", *resp.JSON200.Count)
	}

	if resp.JSON200.CountryDirectors == nil {
		t.Fatal("expected country_directors to be non-nil")
	}
	directors := *resp.JSON200.CountryDirectors
	if len(directors) != 35 {
		t.Errorf("expected 35 country directors, got %d", len(directors))
	}

	// verify first entry
	first := directors[0]
	if first.PlayerProfile == nil {
		t.Fatal("expected player_profile to be non-nil for first director")
	}
	p := first.PlayerProfile
	if p.PlayerId == nil || int(*p.PlayerId) != 86571 {
		t.Errorf("expected first player_id 86571, got %v", p.PlayerId)
	}
	if p.Name == nil || *p.Name != "Esteban Mazzoli" {
		t.Errorf("expected first name 'Esteban Mazzoli', got %v", p.Name)
	}
	if p.CountryCode == nil || *p.CountryCode != "AR" {
		t.Errorf("expected first country_code 'AR', got %v", p.CountryCode)
	}
	if p.CountryName == nil || *p.CountryName != "Argentina" {
		t.Errorf("expected first country_name 'Argentina', got %v", p.CountryName)
	}

	// verify last entry
	last := directors[len(directors)-1]
	if last.PlayerProfile == nil {
		t.Fatal("expected player_profile to be non-nil for last director")
	}
	lp := last.PlayerProfile
	if lp.PlayerId == nil || int(*lp.PlayerId) != 4 {
		t.Errorf("expected last player_id 4, got %v", lp.PlayerId)
	}
	if lp.Name == nil || *lp.Name != "Josh Sharpe" {
		t.Errorf("expected last name 'Josh Sharpe', got %v", lp.Name)
	}
	if lp.CountryCode == nil || *lp.CountryCode != "US" {
		t.Errorf("expected last country_code 'US', got %v", lp.CountryCode)
	}

	// all entries must have a player_profile with player_id and name
	for i, d := range directors {
		if d.PlayerProfile == nil {
			t.Errorf("directors[%d]: expected player_profile to be non-nil", i)
			continue
		}
		if d.PlayerProfile.PlayerId == nil {
			t.Errorf("directors[%d]: expected player_id to be non-nil", i)
		}
		if d.PlayerProfile.Name == nil || *d.PlayerProfile.Name == "" {
			t.Errorf("directors[%d]: expected name to be non-empty", i)
		}
	}
}

func TestDirectorViewCountryDirectorsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorCountryResponse)
	ctx := context.Background()
	_, err := client.ViewCountryDirectorsWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/director/country" {
		t.Errorf("expected path '/director/country', got %s", req.URL.Path)
	}
}
