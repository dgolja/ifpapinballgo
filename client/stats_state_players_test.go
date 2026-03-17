package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_state_players_open.json
var sampleStatsStatePlayersOpenResponse string

//go:embed testdata/stats_state_players_women.json
var sampleStatsStatePlayersWomenResponse string

func TestStatsStatePlayerCountOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsStatePlayersOpenResponse)
	ctx := context.Background()
	rankType := StatsStatePlayerCountParamsRankTypeOPEN
	params := &StatsStatePlayerCountParams{RankType: &rankType}
	resp, err := client.StatsStatePlayerCountWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Players by State (North America)" {
		t.Errorf("expected type 'Players by State (North America)', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 38 {
		t.Errorf("expected 38 stats entries, got %d", len(stats))
	}

	// Verify first entry (Unknown)
	first := stats[0]
	if first.Stateprov == nil || *first.Stateprov != "Unknown" {
		t.Errorf("expected first stateprov 'Unknown', got %v", first.Stateprov)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 38712 {
		t.Errorf("expected first player_count 38712, got %v", first.PlayerCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}

	// Verify a known mid-list entry (CA, rank 2)
	second := stats[1]
	if second.Stateprov == nil || *second.Stateprov != "CA" {
		t.Errorf("expected second stateprov 'CA', got %v", second.Stateprov)
	}
	if second.PlayerCount == nil || int(*second.PlayerCount) != 678 {
		t.Errorf("expected second player_count 678, got %v", second.PlayerCount)
	}

	// Verify all entries have required fields
	for i, s := range stats {
		if s.Stateprov == nil || *s.Stateprov == "" {
			t.Errorf("stats[%d]: expected stateprov to be non-empty", i)
		}
		if s.PlayerCount == nil {
			t.Errorf("stats[%d]: expected player_count to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsStatePlayerCountWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsStatePlayersWomenResponse)
	ctx := context.Background()
	rankType := StatsStatePlayerCountParamsRankTypeWOMEN
	params := &StatsStatePlayerCountParams{RankType: &rankType}
	resp, err := client.StatsStatePlayerCountWithResponse(ctx, params)
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
	if len(stats) != 22 {
		t.Errorf("expected 22 stats entries, got %d", len(stats))
	}

	// Verify first entry (Unknown)
	first := stats[0]
	if first.Stateprov == nil || *first.Stateprov != "Unknown" {
		t.Errorf("expected first stateprov 'Unknown', got %v", first.Stateprov)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 5257 {
		t.Errorf("expected first player_count 5257, got %v", first.PlayerCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsStatePlayerCountRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsStatePlayersOpenResponse)
	ctx := context.Background()
	rankType := StatsStatePlayerCountParamsRankTypeOPEN
	params := &StatsStatePlayerCountParams{RankType: &rankType}
	_, err := client.StatsStatePlayerCountWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/state_players" {
		t.Errorf("expected path '/stats/state_players', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
