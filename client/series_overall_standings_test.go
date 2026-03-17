package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/series_overall_standings_acs.json
var sampleSeriesOverallStandingsACSResponse string

//go:embed testdata/series_overall_standings_wnacso.json
var sampleSeriesOverallStandingsWNACSOResponse string

func TestSeriesRegionOverallStandingsACS(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesOverallStandingsACSResponse)
	ctx := context.Background()
	resp, err := client.SeriesRegionOverallStandingsWithResponse(ctx, "ACS", &SeriesRegionOverallStandingsParams{})
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

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2024 {
		t.Errorf("expected year 2024, got %v", resp.JSON200.Year)
	}

	if resp.JSON200.ChampionshipPrizeFund == nil {
		t.Error("expected championship_prize_fund to be non-nil")
	}

	if resp.JSON200.OverallResults == nil {
		t.Fatal("expected overall_results to be non-nil")
	}

	results := *resp.JSON200.OverallResults
	if len(results) != 7 {
		t.Errorf("expected 7 overall_results, got %d", len(results))
	}

	// Verify first entry - ACT
	first := results[0]
	if first.RegionCode == nil || *first.RegionCode != "ACT" {
		t.Errorf("expected first region_code 'ACT', got %v", first.RegionCode)
	}
	if first.RegionName == nil || *first.RegionName != "Australian Capital Territory" {
		t.Errorf("expected first region_name 'Australian Capital Territory', got %v", first.RegionName)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 268 {
		t.Errorf("expected first player_count 268, got %v", first.PlayerCount)
	}
	if first.UniquePlayerCount == nil || int(*first.UniquePlayerCount) != 64 {
		t.Errorf("expected first unique_player_count 64, got %v", first.UniquePlayerCount)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 14 {
		t.Errorf("expected first tournament_count 14, got %v", first.TournamentCount)
	}
	if first.PrizeFund == nil {
		t.Error("expected first prize_fund to be non-nil")
	}
	if first.CurrentLeader == nil {
		t.Fatal("expected first current_leader to be non-nil")
	}
	if first.CurrentLeader.PlayerId == nil || int(*first.CurrentLeader.PlayerId) != 14456 {
		t.Errorf("expected first current_leader player_id 14456, got %v", first.CurrentLeader.PlayerId)
	}
	if first.CurrentLeader.PlayerName == nil || *first.CurrentLeader.PlayerName != "Curtis Sahariv" {
		t.Errorf("expected first current_leader player_name 'Curtis Sahariv', got %v", first.CurrentLeader.PlayerName)
	}

	// Verify all entries have required fields
	for i, r := range results {
		if r.RegionCode == nil || *r.RegionCode == "" {
			t.Errorf("results[%d]: expected region_code to be non-empty", i)
		}
		if r.RegionName == nil || *r.RegionName == "" {
			t.Errorf("results[%d]: expected region_name to be non-empty", i)
		}
		if r.PlayerCount == nil {
			t.Errorf("results[%d]: expected player_count to be non-nil", i)
		}
		if r.CurrentLeader == nil {
			t.Errorf("results[%d]: expected current_leader to be non-nil", i)
		}
	}
}

func TestSeriesRegionOverallStandingsWNACSO(t *testing.T) {
	client, _ := newTestClient(t, sampleSeriesOverallStandingsWNACSOResponse)
	ctx := context.Background()
	resp, err := client.SeriesRegionOverallStandingsWithResponse(ctx, "WNACSO", &SeriesRegionOverallStandingsParams{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.SeriesCode == nil || *resp.JSON200.SeriesCode != "WNACSO" {
		t.Errorf("expected series_code 'WNACSO', got %v", resp.JSON200.SeriesCode)
	}

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2024 {
		t.Errorf("expected year 2024, got %v", resp.JSON200.Year)
	}

	if resp.JSON200.OverallResults == nil {
		t.Fatal("expected overall_results to be non-nil")
	}

	results := *resp.JSON200.OverallResults
	if len(results) != 58 {
		t.Errorf("expected 58 overall_results, got %d", len(results))
	}

	// Verify first entry - Alabama
	first := results[0]
	if first.RegionCode == nil || *first.RegionCode != "AL" {
		t.Errorf("expected first region_code 'AL', got %v", first.RegionCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 279 {
		t.Errorf("expected first player_count 279, got %v", first.PlayerCount)
	}
	if first.CurrentLeader == nil || first.CurrentLeader.PlayerName == nil || *first.CurrentLeader.PlayerName != "Rose Gwathney" {
		t.Errorf("expected first current_leader 'Rose Gwathney', got %v", first.CurrentLeader)
	}
}

func TestSeriesRegionOverallStandingsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesOverallStandingsACSResponse)
	ctx := context.Background()
	_, err := client.SeriesRegionOverallStandingsWithResponse(ctx, "ACS", &SeriesRegionOverallStandingsParams{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/overall_standings" {
		t.Errorf("expected path '/series/ACS/overall_standings', got %s", req.URL.Path)
	}
}
