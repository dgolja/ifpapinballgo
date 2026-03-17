package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_players_by_year.json
var sampleStatsPlayersByYearResponse string

func TestStatsPlayersByYear(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsPlayersByYearResponse)
	ctx := context.Background()

	resp, err := client.StatsPlayerByYearWithResponse(ctx)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Players by Year" {
		t.Errorf("expected type 'Players by Year', got %v", resp.JSON200.Type)
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

	// verify first entry (most recent year)
	first := stats[0]
	if first.Year == nil || int(*first.Year) != 2025 {
		t.Errorf("expected first year 2025, got %v", first.Year)
	}
	if first.CurrentYearCount == nil || int(*first.CurrentYearCount) != 43745 {
		t.Errorf("expected first current_year_count 43745, got %v", first.CurrentYearCount)
	}
	if first.PreviousYearCount == nil || int(*first.PreviousYearCount) != 18452 {
		t.Errorf("expected first previous_year_count 18452, got %v", first.PreviousYearCount)
	}
	if first.Previous2YearCount == nil || int(*first.Previous2YearCount) != 8279 {
		t.Errorf("expected first previous_2_year_count 8279, got %v", first.Previous2YearCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}

	// verify stats_rank is sequential
	for i, s := range stats {
		if s.StatsRank == nil || *s.StatsRank != i+1 {
			t.Errorf("stats[%d]: expected stats_rank %d, got %v", i, i+1, s.StatsRank)
		}
	}

	// verify last entry
	last := stats[9]
	if last.Year == nil || int(*last.Year) != 2016 {
		t.Errorf("expected last year 2016, got %v", last.Year)
	}
	if last.CurrentYearCount == nil || int(*last.CurrentYearCount) != 17230 {
		t.Errorf("expected last current_year_count 17230, got %v", last.CurrentYearCount)
	}
	if last.StatsRank == nil || *last.StatsRank != 10 {
		t.Errorf("expected last stats_rank 10, got %v", last.StatsRank)
	}
}

func TestStatsPlayersByYearRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsPlayersByYearResponse)
	ctx := context.Background()

	_, err := client.StatsPlayerByYearWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/players_by_year" {
		t.Errorf("expected path '/stats/players_by_year', got %s", req.URL.Path)
	}
}
