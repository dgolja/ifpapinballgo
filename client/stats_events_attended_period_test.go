package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_events_attended_period_open.json
var sampleStatsEventsAttendedPeriodOpenResponse string

//go:embed testdata/stats_events_attended_period_women.json
var sampleStatsEventsAttendedPeriodWomenResponse string

func TestStatsEventsAttendedPeriodOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsEventsAttendedPeriodOpenResponse)
	ctx := context.Background()
	rankType := StatsEventPeriodParamsRankTypeOPEN
	params := &StatsEventPeriodParams{RankType: &rankType}
	resp, err := client.StatsEventPeriodWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Events attended over a period of time" {
		t.Errorf("expected type 'Events attended over a period of time', got %v", resp.JSON200.Type)
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
	if first.PlayerId == nil || int(*first.PlayerId) != 39336 {
		t.Errorf("expected first player_id 39336, got %v", first.PlayerId)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 25 {
		t.Errorf("expected first tournament_count 25, got %v", first.TournamentCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
	if first.CountryCode == nil || *first.CountryCode != "SI" {
		t.Errorf("expected first country_code 'SI', got %v", first.CountryCode)
	}
	if first.CountryName == nil || *first.CountryName != "Slovenia" {
		t.Errorf("expected first country_name 'Slovenia', got %v", first.CountryName)
	}

	// Verify all entries have required fields
	for i, s := range stats {
		if s.PlayerId == nil {
			t.Errorf("stats[%d]: expected player_id to be non-nil", i)
		}
		if s.FirstName == nil || *s.FirstName == "" {
			t.Errorf("stats[%d]: expected first_name to be non-empty", i)
		}
		if s.LastName == nil || *s.LastName == "" {
			t.Errorf("stats[%d]: expected last_name to be non-empty", i)
		}
		if s.TournamentCount == nil {
			t.Errorf("stats[%d]: expected tournament_count to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsEventsAttendedPeriodWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsEventsAttendedPeriodWomenResponse)
	ctx := context.Background()
	rankType := StatsEventPeriodParamsRankTypeWOMEN
	params := &StatsEventPeriodParams{RankType: &rankType}
	resp, err := client.StatsEventPeriodWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
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
	if first.PlayerId == nil || int(*first.PlayerId) != 89391 {
		t.Errorf("expected first player_id 89391, got %v", first.PlayerId)
	}
	if first.TournamentCount == nil || int(*first.TournamentCount) != 199 {
		t.Errorf("expected first tournament_count 199, got %v", first.TournamentCount)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
}

func TestStatsEventsAttendedPeriodRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsEventsAttendedPeriodOpenResponse)
	ctx := context.Background()
	rankType := StatsEventPeriodParamsRankTypeOPEN
	params := &StatsEventPeriodParams{RankType: &rankType}
	_, err := client.StatsEventPeriodWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/events_attended_period" {
		t.Errorf("expected path '/stats/events_attended_period', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
