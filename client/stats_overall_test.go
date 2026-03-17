package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_overall_open.json
var sampleStatsOverallOpenResponse string

//go:embed testdata/stats_overall_women.json
var sampleStatsOverallWomenResponse string

func TestStatsOverallOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsOverallOpenResponse)
	ctx := context.Background()
	systemCode := StatsOverallParamsSystemCodeOPEN
	params := &StatsOverallParams{SystemCode: &systemCode}
	resp, err := client.StatsOverallWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Overall Stats" {
		t.Errorf("expected type 'Overall Stats', got %v", resp.JSON200.Type)
	}

	if resp.JSON200.SystemCode == nil || *resp.JSON200.SystemCode != "OPEN" {
		t.Errorf("expected system_code 'OPEN', got %v", resp.JSON200.SystemCode)
	}

	if resp.JSON200.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}

	stats := resp.JSON200.Stats
	if stats.OverallPlayerCount == nil || *stats.OverallPlayerCount != 149872 {
		t.Errorf("expected overall_player_count 149872, got %v", stats.OverallPlayerCount)
	}
	if stats.ActivePlayerCount == nil || *stats.ActivePlayerCount != 75606 {
		t.Errorf("expected active_player_count 75606, got %v", stats.ActivePlayerCount)
	}
	if stats.TournamentCount == nil || *stats.TournamentCount != 89509 {
		t.Errorf("expected tournament_count 89509, got %v", stats.TournamentCount)
	}
	if stats.TournamentCountLastMonth == nil || *stats.TournamentCountLastMonth != 852 {
		t.Errorf("expected tournament_count_last_month 852, got %v", stats.TournamentCountLastMonth)
	}
	if stats.TournamentCountThisYear == nil || *stats.TournamentCountThisYear != 2008 {
		t.Errorf("expected tournament_count_this_year 2008, got %v", stats.TournamentCountThisYear)
	}
	if stats.TournamentPlayerCount == nil || *stats.TournamentPlayerCount != 2052044 {
		t.Errorf("expected tournament_player_count 2052044, got %v", stats.TournamentPlayerCount)
	}
	if stats.TournamentPlayerCountAverage == nil || *stats.TournamentPlayerCountAverage != 22.9 {
		t.Errorf("expected tournament_player_count_average 22.9, got %v", stats.TournamentPlayerCountAverage)
	}

	if stats.Age == nil {
		t.Fatal("expected age to be non-nil")
	}
	if stats.Age.AgeUnder18 == nil || *stats.Age.AgeUnder18 != 3.57 {
		t.Errorf("expected age_under_18 3.57, got %v", stats.Age.AgeUnder18)
	}
	if stats.Age.Age18To29 == nil || *stats.Age.Age18To29 != 9.74 {
		t.Errorf("expected age_18_to_29 9.74, got %v", stats.Age.Age18To29)
	}
	if stats.Age.Age30To39 == nil || *stats.Age.Age30To39 != 22.86 {
		t.Errorf("expected age_30_to_39 22.86, got %v", stats.Age.Age30To39)
	}
	if stats.Age.Age40To49 == nil || *stats.Age.Age40To49 != 30.81 {
		t.Errorf("expected age_40_to_49 30.81, got %v", stats.Age.Age40To49)
	}
	if stats.Age.Age50To99 == nil || *stats.Age.Age50To99 != 33.02 {
		t.Errorf("expected age_50_to_99 33.02, got %v", stats.Age.Age50To99)
	}
}

func TestStatsOverallWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsOverallWomenResponse)
	ctx := context.Background()
	systemCode := StatsOverallParamsSystemCodeWOMEN
	params := &StatsOverallParams{SystemCode: &systemCode}
	resp, err := client.StatsOverallWithResponse(ctx, params)
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

	stats := resp.JSON200.Stats
	if stats.OverallPlayerCount == nil || *stats.OverallPlayerCount != 149872 {
		t.Errorf("expected overall_player_count 149872, got %v", stats.OverallPlayerCount)
	}

	if stats.Age == nil {
		t.Fatal("expected age to be non-nil")
	}
}

func TestStatsOverallRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsOverallOpenResponse)
	ctx := context.Background()
	systemCode := StatsOverallParamsSystemCodeOPEN
	params := &StatsOverallParams{SystemCode: &systemCode}
	_, err := client.StatsOverallWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/overall" {
		t.Errorf("expected path '/stats/overall', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("system_code") != "OPEN" {
		t.Errorf("expected query param system_code=OPEN, got %s", req.URL.Query().Get("system_code"))
	}
}
