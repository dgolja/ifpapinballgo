package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_lucrative_tournaments_open_major.json
var sampleStatsLucrativeTournamentsOpenMajorResponse string

//go:embed testdata/stats_lucrative_tournaments_open_no_major.json
var sampleStatsLucrativeTournamentsOpenNoMajorResponse string

//go:embed testdata/stats_lucrative_tournaments_women_major.json
var sampleStatsLucrativeTournamentsWomenMajorResponse string

func TestStatsLucrativeTournamentsOpenMajor(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsLucrativeTournamentsOpenMajorResponse)
	ctx := context.Background()
	rankType := StatsLucrativeToursParamsRankTypeOPEN
	majorY := "Y"
	params := &StatsLucrativeToursParams{RankType: &rankType, Major: &majorY}
	resp, err := client.StatsLucrativeToursWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Lucrative Tournaments" {
		t.Errorf("expected type 'Lucrative Tournaments', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 25 {
		t.Errorf("expected 25 stats entries, got %d", len(stats))
	}

	// Verify first entry (rank 1)
	first := stats[0]
	if first.CountryName == nil || *first.CountryName != "Sweden" {
		t.Errorf("expected first country_name 'Sweden', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "SE" {
		t.Errorf("expected first country_code 'SE', got %v", first.CountryCode)
	}
	if first.TournamentId == nil || int(*first.TournamentId) != 70912 {
		t.Errorf("expected first tournament_id 70912, got %v", first.TournamentId)
	}
	if first.TournamentValue == nil || *first.TournamentValue != 186.47 {
		t.Errorf("expected first tournament_value 186.47, got %v", first.TournamentValue)
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
		if s.TournamentId == nil {
			t.Errorf("stats[%d]: expected tournament_id to be non-nil", i)
		}
		if s.TournamentName == nil || *s.TournamentName == "" {
			t.Errorf("stats[%d]: expected tournament_name to be non-empty", i)
		}
		if s.TournamentValue == nil {
			t.Errorf("stats[%d]: expected tournament_value to be non-nil", i)
		}
		if s.TournamentDate == nil {
			t.Errorf("stats[%d]: expected tournament_date to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsLucrativeTournamentsOpenNoMajor(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsLucrativeTournamentsOpenNoMajorResponse)
	ctx := context.Background()
	rankType := StatsLucrativeToursParamsRankTypeOPEN
	majorN := "N"
	params := &StatsLucrativeToursParams{RankType: &rankType, Major: &majorN}
	resp, err := client.StatsLucrativeToursWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.RankType == nil || *resp.JSON200.RankType != "OPEN" {
		t.Errorf("expected rank_type 'OPEN', got %v", resp.JSON200.RankType)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := *resp.JSON200.Stats
	if len(stats) != 25 {
		t.Errorf("expected 25 stats entries, got %d", len(stats))
	}

	// First entry should be same as major (major filter removed some entries but rank 1 is unchanged)
	first := stats[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 70912 {
		t.Errorf("expected first tournament_id 70912, got %v", first.TournamentId)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsLucrativeTournamentsWomenMajor(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsLucrativeTournamentsWomenMajorResponse)
	ctx := context.Background()
	rankType := StatsLucrativeToursParamsRankTypeWOMEN
	majorY := "Y"
	params := &StatsLucrativeToursParams{RankType: &rankType, Major: &majorY}
	resp, err := client.StatsLucrativeToursWithResponse(ctx, params)
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
	if len(stats) != 25 {
		t.Errorf("expected 25 stats entries, got %d", len(stats))
	}

	// Verify first entry (rank 1)
	first := stats[0]
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.TournamentId == nil || int(*first.TournamentId) != 82658 {
		t.Errorf("expected first tournament_id 82658, got %v", first.TournamentId)
	}
	if first.TournamentValue == nil || *first.TournamentValue != 189.84 {
		t.Errorf("expected first tournament_value 189.84, got %v", first.TournamentValue)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
}

func TestStatsLucrativeTournamentsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsLucrativeTournamentsOpenMajorResponse)
	ctx := context.Background()
	rankType := StatsLucrativeToursParamsRankTypeOPEN
	majorY := "Y"
	params := &StatsLucrativeToursParams{RankType: &rankType, Major: &majorY}
	_, err := client.StatsLucrativeToursWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/lucrative_tournaments" {
		t.Errorf("expected path '/stats/lucrative_tournaments', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
	if req.URL.Query().Get("major") != "Y" {
		t.Errorf("expected query param major=Y, got %s", req.URL.Query().Get("major"))
	}
}
