package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_largest_tournaments_open.json
var sampleStatsLargestTournamentsOpenResponse string

//go:embed testdata/stats_largest_tournaments_women.json
var sampleStatsLargestTournamentsWomenResponse string

func TestStatsLargestTournamentsOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsLargestTournamentsOpenResponse)
	ctx := context.Background()
	rankType := StatsLargestTournamentsParamsRankTypeOPEN
	params := &StatsLargestTournamentsParams{RankType: &rankType}
	resp, err := client.StatsLargestTournamentsWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Largest Tournaments" {
		t.Errorf("expected type 'Largest Tournaments', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 14 {
		t.Errorf("expected 14 stats entries, got %d", len(stats))
	}

	// Verify first entry (rank 1)
	first := stats[0]
	if first.CountryName == nil || *first.CountryName != "Sweden" {
		t.Errorf("expected first country_name 'Sweden', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "SE" {
		t.Errorf("expected first country_code 'SE', got %v", first.CountryCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 317 {
		t.Errorf("expected first player_count 317, got %v", first.PlayerCount)
	}
	if first.TournamentId == nil || int(*first.TournamentId) != 100743 {
		t.Errorf("expected first tournament_id 100743, got %v", first.TournamentId)
	}
	if first.TournamentDate == nil {
		t.Error("expected first tournament_date to be non-nil")
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
		if s.TournamentId == nil {
			t.Errorf("stats[%d]: expected tournament_id to be non-nil", i)
		}
		if s.TournamentName == nil || *s.TournamentName == "" {
			t.Errorf("stats[%d]: expected tournament_name to be non-empty", i)
		}
		if s.TournamentDate == nil {
			t.Errorf("stats[%d]: expected tournament_date to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsLargestTournamentsWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsLargestTournamentsWomenResponse)
	ctx := context.Background()
	rankType := StatsLargestTournamentsParamsRankTypeWOMEN
	params := &StatsLargestTournamentsParams{RankType: &rankType}
	resp, err := client.StatsLargestTournamentsWithResponse(ctx, params)
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
	if len(stats) != 15 {
		t.Errorf("expected 15 stats entries, got %d", len(stats))
	}

	// Verify first entry (rank 1)
	first := stats[0]
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 127 {
		t.Errorf("expected first player_count 127, got %v", first.PlayerCount)
	}
	if first.TournamentId == nil || int(*first.TournamentId) != 34627 {
		t.Errorf("expected first tournament_id 34627, got %v", first.TournamentId)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsLargestTournamentsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsLargestTournamentsOpenResponse)
	ctx := context.Background()
	rankType := StatsLargestTournamentsParamsRankTypeOPEN
	params := &StatsLargestTournamentsParams{RankType: &rankType}
	_, err := client.StatsLargestTournamentsWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/largest_tournaments" {
		t.Errorf("expected path '/stats/largest_tournaments', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
