package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_state_tournaments_open.json
var sampleStatsStateTournamentsOpenResponse string

//go:embed testdata/stats_state_tournaments_women.json
var sampleStatsStateTournamentsWomenResponse string

func TestStatsStateTourCountOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsStateTournamentsOpenResponse)
	ctx := context.Background()
	rankType := StatsStateTourCountParamsRankTypeOPEN
	params := &StatsStateTourCountParams{RankType: &rankType}
	resp, err := client.StatsStateTourCountWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Tournaments by State (North America)" {
		t.Errorf("expected type 'Tournaments by State (North America)', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 13 {
		t.Errorf("expected 13 stats entries, got %d", len(stats))
	}

	// Verify first entry (WA, rank 1)
	first := stats[0]
	if first.Stateprov == nil || *first.Stateprov != "WA" {
		t.Errorf("expected first stateprov 'WA', got %v", first.Stateprov)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 5869 {
		t.Errorf("expected first tournament_count 5869, got %v", first.TournamentCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}

	// Verify all entries have required fields
	for i, s := range stats {
		if s.Stateprov == nil || *s.Stateprov == "" {
			t.Errorf("stats[%d]: expected stateprov to be non-empty", i)
		}
		if s.TournamentCount == nil {
			t.Errorf("stats[%d]: expected tournament_count to be non-nil", i)
		}
		if s.TotalPointsAll == nil {
			t.Errorf("stats[%d]: expected total_points_all to be non-nil", i)
		}
		if s.TotalPointsTournamentValue == nil {
			t.Errorf("stats[%d]: expected total_points_tournament_value to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsStateTourCountWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsStateTournamentsWomenResponse)
	ctx := context.Background()
	rankType := StatsStateTourCountParamsRankTypeWOMEN
	params := &StatsStateTourCountParams{RankType: &rankType}
	resp, err := client.StatsStateTourCountWithResponse(ctx, params)
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
	if len(stats) != 29 {
		t.Errorf("expected 29 stats entries, got %d", len(stats))
	}

	// Verify first entry (TX, rank 1)
	first := stats[0]
	if first.Stateprov == nil || *first.Stateprov != "TX" {
		t.Errorf("expected first stateprov 'TX', got %v", first.Stateprov)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 492 {
		t.Errorf("expected first tournament_count 492, got %v", first.TournamentCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsStateTourCountRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsStateTournamentsOpenResponse)
	ctx := context.Background()
	rankType := StatsStateTourCountParamsRankTypeOPEN
	params := &StatsStateTourCountParams{RankType: &rankType}
	_, err := client.StatsStateTourCountWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/state_tournaments" {
		t.Errorf("expected path '/stats/state_tournaments', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
