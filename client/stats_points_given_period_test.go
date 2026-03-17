package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/stats_points_given_period_open.json
var sampleStatsPointsPeriodOpenResponse string

//go:embed testdata/stats_points_given_period_women.json
var sampleStatsPointsPeriodWomenResponse string

func TestStatsPointsPeriodOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsPointsPeriodOpenResponse)
	ctx := context.Background()
	rankType := StatsPointsPeriodParamsRankTypeOPEN
	params := &StatsPointsPeriodParams{RankType: &rankType}
	resp, err := client.StatsPointsPeriodWithResponse(ctx, params)
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

	if resp.JSON200.Type == nil || *resp.JSON200.Type != "Points given Period" {
		t.Errorf("expected type 'Points given Period', got %v", resp.JSON200.Type)
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
	if first.PlayerId == nil || int(*first.PlayerId) != 16004 {
		t.Errorf("expected first player_id 16004, got %v", first.PlayerId)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 3049.84 {
		t.Errorf("expected first wppr_points 3049.84, got %v", first.WpprPoints)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
	if first.CountryName == nil || *first.CountryName != "Sweden" {
		t.Errorf("expected first country_name 'Sweden', got %v", first.CountryName)
	}
	if first.CountryCode == nil || *first.CountryCode != "SE" {
		t.Errorf("expected first country_code 'SE', got %v", first.CountryCode)
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
		if s.WpprPoints == nil {
			t.Errorf("stats[%d]: expected wppr_points to be non-nil", i)
		}
		if s.StatsRank == nil {
			t.Errorf("stats[%d]: expected stats_rank to be non-nil", i)
		}
	}
}

func TestStatsPointsPeriodWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleStatsPointsPeriodWomenResponse)
	ctx := context.Background()
	rankType := StatsPointsPeriodParamsRankTypeWOMEN
	params := &StatsPointsPeriodParams{RankType: &rankType}
	resp, err := client.StatsPointsPeriodWithResponse(ctx, params)
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
	if first.PlayerId == nil || int(*first.PlayerId) != 47411 {
		t.Errorf("expected first player_id 47411, got %v", first.PlayerId)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 508.39 {
		t.Errorf("expected first wppr_points 508.39, got %v", first.WpprPoints)
	}
	if first.StatsRank == nil || *first.StatsRank != 1 {
		t.Errorf("expected first stats_rank 1, got %v", first.StatsRank)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
}

func TestStatsPointsPeriodRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleStatsPointsPeriodOpenResponse)
	ctx := context.Background()
	rankType := StatsPointsPeriodParamsRankTypeOPEN
	params := &StatsPointsPeriodParams{RankType: &rankType}
	_, err := client.StatsPointsPeriodWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/stats/points_given_period" {
		t.Errorf("expected path '/stats/points_given_period', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("rank_type") != "OPEN" {
		t.Errorf("expected query param rank_type=OPEN, got %s", req.URL.Query().Get("rank_type"))
	}
}
