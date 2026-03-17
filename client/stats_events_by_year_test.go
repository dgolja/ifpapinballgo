package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_events_by_year_open.json
var sampleStatsEventsByYearOpenResponse string

//go:embed testdata/stats_events_by_year_women.json
var sampleStatsEventsByYearWomenResponse string

func TestStatsEventsByYearOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsEventsByYearOpenResponse)
	ctx := context.Background()
	rankType := StatsEventsByYearParamsRankTypeOPEN
	params := &StatsEventsByYearParams{RankType: &rankType}
	resp, err := client.StatsEventsByYearWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Events Per Year" {
		t.Errorf("expected type 'Events Per Year', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 10 {
		t.Errorf("expected 10 stats entries, got %d", len(stats))
	}

	// Verify first entry (2026, rank 1)
	first := stats[0]
	if first.Year == nil || int(*first.Year) != 2026 {
		t.Errorf("expected first year 2026, got %v", first.Year)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 113 {
		t.Errorf("expected first tournament_count 113, got %v", first.TournamentCount)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 2950 {
		t.Errorf("expected first player_count 2950, got %v", first.PlayerCount)
	}
	if first.CountryCount == nil || int(*first.CountryCount) != 1 {
		t.Errorf("expected first country_count 1, got %v", first.CountryCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}

	// Verify all entries have required fields
	for i, s := range stats {
		if s.Year == nil {
			t.Errorf("stats[%d]: expected year to be non-nil", i)
		}
		if s.TournamentCount == nil {
			t.Errorf("stats[%d]: expected tournament_count to be non-nil", i)
		}
		if s.PlayerCount == nil {
			t.Errorf("stats[%d]: expected player_count to be non-nil", i)
		}
		if s.CountryCount == nil {
			t.Errorf("stats[%d]: expected country_count to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsEventsByYearWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsEventsByYearWomenResponse)
	ctx := context.Background()
	rankType := StatsEventsByYearParamsRankTypeWOMEN
	params := &StatsEventsByYearParams{RankType: &rankType}
	resp, err := client.StatsEventsByYearWithResponse(ctx, params)
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
	if len(stats) != 10 {
		t.Errorf("expected 10 stats entries, got %d", len(stats))
	}

	// Verify first entry (2026, rank 1)
	first := stats[0]
	if first.Year == nil || int(*first.Year) != 2026 {
		t.Errorf("expected first year 2026, got %v", first.Year)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 227 {
		t.Errorf("expected first tournament_count 227, got %v", first.TournamentCount)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 3330 {
		t.Errorf("expected first player_count 3330, got %v", first.PlayerCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsEventsByYearRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsEventsByYearOpenResponse)
	ctx := context.Background()
	rankType := StatsEventsByYearParamsRankTypeOPEN
	params := &StatsEventsByYearParams{RankType: &rankType}
	_, err := client.StatsEventsByYearWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/events_by_year" {
		t.Errorf("expected path '/stats/events_by_year', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
