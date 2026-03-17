package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_country_players_open.json
var sampleStatsCountryPlayersOpenResponse string

//go:embed testdata/stats_country_players_women.json
var sampleStatsCountryPlayersWomenResponse string

func TestStatsCountryPlayerCountOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsCountryPlayersOpenResponse)
	ctx := context.Background()
	rankType := StatsCountryPlayerCountParamsRankTypeOPEN
	params := &StatsCountryPlayerCountParams{RankType: &rankType}
	resp, err := client.StatsCountryPlayerCountWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Players by Country" {
		t.Errorf("expected type 'Players by Country', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 20 {
		t.Errorf("expected 20 stats entries, got %d", len(stats))
	}

	// Verify first entry (United States)
	first := stats[0]
	if first.CountryName == nil || *first.CountryName != "United States" {
		t.Errorf("expected first country_name 'United States', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 47944 {
		t.Errorf("expected first player_count 47944, got %v", first.PlayerCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}

	// Verify all entries have required fields
	for i, s := range stats {
		if s.CountryName == nil || *s.CountryName == "" {
			t.Errorf("stats[%d]: expected country_name to be non-empty", i)
		}
		if s.CountryCode == nil || *s.CountryCode == "" {
			t.Errorf("stats[%d]: expected country_code to be non-empty", i)
		}
		if s.PlayerCount == nil {
			t.Errorf("stats[%d]: expected player_count to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsCountryPlayerCountWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsCountryPlayersWomenResponse)
	ctx := context.Background()
	rankType := StatsCountryPlayerCountParamsRankTypeWOMEN
	params := &StatsCountryPlayerCountParams{RankType: &rankType}
	resp, err := client.StatsCountryPlayerCountWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "WOMEN" {
		t.Errorf("expected rank_type 'WOMEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 3 {
		t.Errorf("expected 3 stats entries, got %d", len(stats))
	}

	// Verify first entry (United States)
	first := stats[0]
	if first.CountryName == nil || *first.CountryName != "United States" {
		t.Errorf("expected first country_name 'United States', got %v", first.CountryName)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 7354 {
		t.Errorf("expected first player_count 7354, got %v", first.PlayerCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsCountryPlayerCountRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsCountryPlayersOpenResponse)
	ctx := context.Background()
	rankType := StatsCountryPlayerCountParamsRankTypeOPEN
	params := &StatsCountryPlayerCountParams{RankType: &rankType}
	_, err := client.StatsCountryPlayerCountWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/country_players" {
		t.Errorf("expected path '/stats/country_players', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
