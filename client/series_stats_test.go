package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/dgolja/ifpapinballgo/types"
)

//go:embed testdata/series_stats_acs.json
var sampleSeriesStatsACSResponse string

//go:embed testdata/series_stats_wnacso.json
var sampleSeriesStatsWNACSOResponse string

func TestSeriesRegionRegionStatsACS(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesStatsACSResponse)
	ctx := context.Background()
	regionCode := "NSW"
	resp, err := client.SeriesRegionRegionStatsWithResponse(ctx, "ACS", &SeriesRegionRegionStatsParams{
		RegionCode: regionCode,
	})
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

	if resp.JSON200.SeriesCode == nil || *resp.JSON200.SeriesCode != "ACS" {
		t.Errorf("expected series_code 'ACS', got %v", resp.JSON200.SeriesCode)
	}

	if resp.JSON200.RegionCode == nil || *resp.JSON200.RegionCode != "NSW" {
		t.Errorf("expected region_code 'NSW', got %v", resp.JSON200.RegionCode)
	}

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2025 {
		t.Errorf("expected year 2025, got %v", resp.JSON200.Year)
	}

	// Verify monthly_stats
	if resp.JSON200.MonthlyStats == nil {
		t.Fatal("expected monthly_stats to be non-nil")
	}
	monthly := *resp.JSON200.MonthlyStats
	if len(monthly) != 12 {
		t.Errorf("expected 12 monthly_stats entries, got %d", len(monthly))
	}

	// Verify first month (January)
	jan := monthly[0]
	if jan.Month == nil || int(*jan.Month) != 1 {
		t.Errorf("expected first month 1, got %v", jan.Month)
	}
	if jan.TournamentCount == nil || int(*jan.TournamentCount) != 5 {
		t.Errorf("expected first tournament_count 5, got %v", jan.TournamentCount)
	}
	if jan.PlayerCount == nil || int(*jan.PlayerCount) != 100 {
		t.Errorf("expected first player_count 100, got %v", jan.PlayerCount)
	}
	if jan.UniquePlayerCount == nil || int(*jan.UniquePlayerCount) != 77 {
		t.Errorf("expected first unique_player_count 77, got %v", jan.UniquePlayerCount)
	}
	if jan.PrizeFund == nil || *jan.PrizeFund != types.StringFloat64(65.0) {
		t.Errorf("expected first prize_fund 65.00, got %v", jan.PrizeFund)
	}

	// Verify all months have sequential month numbers
	for i, m := range monthly {
		if m.Month == nil {
			t.Errorf("monthly[%d]: expected month to be non-nil", i)
		} else if int(*m.Month) != i+1 {
			t.Errorf("monthly[%d]: expected month %d, got %d", i, i+1, int(*m.Month))
		}
	}

	// Verify yearly_stats
	if resp.JSON200.YearlyStats == nil {
		t.Fatal("expected yearly_stats to be non-nil")
	}
	yearly := resp.JSON200.YearlyStats
	if yearly.PlayerCount == nil || int(*yearly.PlayerCount) != 372 {
		t.Errorf("expected yearly player_count 372, got %v", yearly.PlayerCount)
	}
	if yearly.UniquePlayerCount == nil || int(*yearly.UniquePlayerCount) != 372 {
		t.Errorf("expected yearly unique_player_count 372, got %v", yearly.UniquePlayerCount)
	}
	if yearly.TournamentCount == nil || int(*yearly.TournamentCount) != 131 {
		t.Errorf("expected yearly tournament_count 131, got %v", yearly.TournamentCount)
	}
	if yearly.FieldSize == nil || int(*yearly.FieldSize) != 16 {
		t.Errorf("expected yearly field_size 16, got %v", yearly.FieldSize)
	}

	// Verify payouts
	if resp.JSON200.Payouts == nil {
		t.Fatal("expected payouts to be non-nil")
	}
	payouts := *resp.JSON200.Payouts
	if len(payouts) != 16 {
		t.Errorf("expected 16 payouts entries, got %d", len(payouts))
	}

	// Verify first payout
	firstPayout := payouts[0]
	if firstPayout.Position == nil || int(*firstPayout.Position) != 1 {
		t.Errorf("expected first payout position 1, got %v", firstPayout.Position)
	}
	if firstPayout.PrizeFund == nil {
		t.Error("expected first payout prize_fund to be non-nil")
	}

	// Verify all payouts have sequential positions
	for i, p := range payouts {
		if p.Position == nil {
			t.Errorf("payouts[%d]: expected position to be non-nil", i)
		} else if int(*p.Position) != i+1 {
			t.Errorf("payouts[%d]: expected position %d, got %d", i, i+1, int(*p.Position))
		}
	}
}

func TestSeriesRegionRegionStatsWNACSO(t *testing.T) {
	client, _ := newTestClient(t, sampleSeriesStatsWNACSOResponse)
	ctx := context.Background()
	regionCode := "NSW"
	resp, err := client.SeriesRegionRegionStatsWithResponse(ctx, "WNACSO", &SeriesRegionRegionStatsParams{
		RegionCode: regionCode,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.SeriesCode == nil || *resp.JSON200.SeriesCode != "WNACSO" {
		t.Errorf("expected series_code 'WNACSO', got %v", resp.JSON200.SeriesCode)
	}

	if resp.JSON200.MonthlyStats == nil {
		t.Fatal("expected monthly_stats to be non-nil")
	}
	if len(*resp.JSON200.MonthlyStats) != 12 {
		t.Errorf("expected 12 monthly_stats, got %d", len(*resp.JSON200.MonthlyStats))
	}

	// All monthly stats should be zero for this empty series/region
	for i, m := range *resp.JSON200.MonthlyStats {
		if m.TournamentCount == nil || int(*m.TournamentCount) != 0 {
			t.Errorf("monthly[%d]: expected tournament_count 0, got %v", i, m.TournamentCount)
		}
		if m.PlayerCount == nil || int(*m.PlayerCount) != 0 {
			t.Errorf("monthly[%d]: expected player_count 0, got %v", i, m.PlayerCount)
		}
	}

	if resp.JSON200.YearlyStats == nil {
		t.Fatal("expected yearly_stats to be non-nil")
	}
	if resp.JSON200.YearlyStats.TournamentCount == nil || int(*resp.JSON200.YearlyStats.TournamentCount) != 0 {
		t.Errorf("expected yearly tournament_count 0, got %v", resp.JSON200.YearlyStats.TournamentCount)
	}
}

func TestSeriesRegionRegionStatsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesStatsACSResponse)
	ctx := context.Background()
	regionCode := "NSW"
	_, err := client.SeriesRegionRegionStatsWithResponse(ctx, "ACS", &SeriesRegionRegionStatsParams{
		RegionCode: regionCode,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/stats" {
		t.Errorf("expected path '/series/ACS/stats', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("region_code") != "NSW" {
		t.Errorf("expected query param region_code 'NSW', got '%s'", req.URL.Query().Get("region_code"))
	}
}
