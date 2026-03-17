package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/dgolja/ifpapinballgo/types"
)

//go:embed testdata/series_standings_acs.json
var sampleSeriesStandingsACSResponse string

func TestSeriesRegionRegionStandings(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesStandingsACSResponse)
	ctx := context.Background()
	regionCode := "NSW"
	resp, err := client.SeriesRegionRegionStandingsWithResponse(ctx, "ACS", &SeriesRegionRegionStandingsParams{
		RegionCode: regionCode,
	})
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

	if resp.JSON200.SeriesCode == nil || *resp.JSON200.SeriesCode != "ACS" {
		t.Errorf("expected series_code 'ACS', got %v", resp.JSON200.SeriesCode)
	}

	if resp.JSON200.RegionCode == nil || *resp.JSON200.RegionCode != "NSW" {
		t.Errorf("expected region_code 'NSW', got %v", resp.JSON200.RegionCode)
	}

	if resp.JSON200.RegionName == nil || *resp.JSON200.RegionName != "New South Wales" {
		t.Errorf("expected region_name 'New South Wales', got %v", resp.JSON200.RegionName)
	}

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2024 {
		t.Errorf("expected year 2024, got %v", resp.JSON200.Year)
	}

	if resp.JSON200.PrizeFund == nil || *resp.JSON200.PrizeFund != types.StringFloat64(1238.25) {
		t.Errorf("expected prize_fund 1238.25, got %v", resp.JSON200.PrizeFund)
	}

	if resp.JSON200.Standings == nil {
		t.Fatal("expected standings to be non-nil")
	}

	standings := *resp.JSON200.Standings
	if len(standings) != 29 {
		t.Errorf("expected 29 standings, got %d", len(standings))
	}

	// Verify first entry - rank 1
	first := standings[0]
	if first.SeriesRank == nil || int(*first.SeriesRank) != 1 {
		t.Errorf("expected first series_rank 1, got %v", first.SeriesRank)
	}
	if first.PlayerId == nil || int(*first.PlayerId) != 19853 {
		t.Errorf("expected first player_id 19853, got %v", first.PlayerId)
	}
	if first.PlayerName == nil || *first.PlayerName != "Paul Jones AUS" {
		t.Errorf("expected first player_name 'Paul Jones AUS', got %v", first.PlayerName)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.WpprPoints == nil {
		t.Error("expected first wppr_points to be non-nil")
	}
	if first.EventCount == nil || int(*first.EventCount) != 20 {
		t.Errorf("expected first event_count 20, got %v", first.EventCount)
	}
	if first.WinCount == nil || int(*first.WinCount) != 14 {
		t.Errorf("expected first win_count 14, got %v", first.WinCount)
	}

	// Verify all entries have required fields and ascending ranks
	for i, s := range standings {
		if s.SeriesRank == nil {
			t.Errorf("standings[%d]: expected series_rank to be non-nil", i)
		} else if int(*s.SeriesRank) != i+1 {
			t.Errorf("standings[%d]: expected series_rank %d, got %d", i, i+1, int(*s.SeriesRank))
		}
		if s.PlayerId == nil {
			t.Errorf("standings[%d]: expected player_id to be non-nil", i)
		}
		if s.PlayerName == nil || *s.PlayerName == "" {
			t.Errorf("standings[%d]: expected player_name to be non-empty", i)
		}
		if s.WpprPoints == nil {
			t.Errorf("standings[%d]: expected wppr_points to be non-nil", i)
		}
	}
}

func TestSeriesRegionRegionStandingsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesStandingsACSResponse)
	ctx := context.Background()
	regionCode := "NSW"
	_, err := client.SeriesRegionRegionStandingsWithResponse(ctx, "ACS", &SeriesRegionRegionStandingsParams{
		RegionCode: regionCode,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/standings" {
		t.Errorf("expected path '/series/ACS/standings', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("region_code") != "NSW" {
		t.Errorf("expected query param region_code 'NSW', got '%s'", req.URL.Query().Get("region_code"))
	}
}
