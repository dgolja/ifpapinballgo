package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/dgolja/ifpapinballgo/types"
)

//go:embed testdata/series_tournaments_acs_qld.json
var sampleSeriesTournamentsACSQLDResponse string

//go:embed testdata/series_tournaments_wnacso_il.json
var sampleSeriesTournamentsWNACSOILResponse string

func TestSeriesRegionTourACSQLD(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesTournamentsACSQLDResponse)
	ctx := context.Background()
	regionCode := "QLD"
	resp, err := client.SeriesRegionTourWithResponse(ctx, "ACS", &SeriesRegionTourParams{
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

	if resp.JSON200.RegionCode == nil || *resp.JSON200.RegionCode != "QLD" {
		t.Errorf("expected region_code 'QLD', got %v", resp.JSON200.RegionCode)
	}

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2025 {
		t.Errorf("expected year 2025, got %v", resp.JSON200.Year)
	}

	if resp.JSON200.SubmittedTournaments == nil {
		t.Fatal("expected submitted_tournaments to be non-nil")
	}

	tournaments := *resp.JSON200.SubmittedTournaments
	if len(tournaments) != 222 {
		t.Errorf("expected 222 submitted_tournaments, got %d", len(tournaments))
	}

	// Verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 88669 {
		t.Errorf("expected first tournament_id 88669, got %v", first.TournamentId)
	}
	if first.EventEndDate == nil || *first.EventEndDate != "2025-12-29" {
		t.Errorf("expected first event_end_date '2025-12-29', got %v", first.EventEndDate)
	}
	if first.TournamentName == nil || *first.TournamentName != "Empire Cairns Pizza + Pinball Strikes" {
		t.Errorf("expected first tournament_name 'Empire Cairns Pizza + Pinball Strikes', got %v", first.TournamentName)
	}
	if first.EventName == nil || *first.EventName != "Main Tournament" {
		t.Errorf("expected first event_name 'Main Tournament', got %v", first.EventName)
	}
	if first.WpprPoints == nil || *first.WpprPoints != types.StringFloat64(12.01) {
		t.Errorf("expected first wppr_points 12.0100, got %v", first.WpprPoints)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 47 {
		t.Errorf("expected first player_count 47, got %v", first.PlayerCount)
	}
	if first.Winner == nil {
		t.Fatal("expected first winner to be non-nil")
	}
	if first.Winner.PlayerId == nil || int(*first.Winner.PlayerId) != 114854 {
		t.Errorf("expected first winner player_id 114854, got %v", first.Winner.PlayerId)
	}
	if first.Winner.Name == nil || *first.Winner.Name != "Ronan Wolfe" {
		t.Errorf("expected first winner name 'Ronan Wolfe', got %v", first.Winner.Name)
	}

	// Verify all tournaments have required fields
	for i, tour := range tournaments {
		if tour.TournamentId == nil {
			t.Errorf("tournaments[%d]: expected tournament_id to be non-nil", i)
		}
		if tour.TournamentName == nil || *tour.TournamentName == "" {
			t.Errorf("tournaments[%d]: expected tournament_name to be non-empty", i)
		}
		if tour.Winner == nil {
			t.Errorf("tournaments[%d]: expected winner to be non-nil", i)
		}
	}
}

func TestSeriesRegionTourWNACSOIL(t *testing.T) {
	client, _ := newTestClient(t, sampleSeriesTournamentsWNACSOILResponse)
	ctx := context.Background()
	resp, err := client.SeriesRegionTourWithResponse(ctx, "WNACSO", &SeriesRegionTourParams{
		RegionCode: "IL",
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

	if resp.JSON200.RegionCode == nil || *resp.JSON200.RegionCode != "IL" {
		t.Errorf("expected region_code 'IL', got %v", resp.JSON200.RegionCode)
	}

	if resp.JSON200.SubmittedTournaments == nil {
		t.Fatal("expected submitted_tournaments to be non-nil")
	}

	tournaments := *resp.JSON200.SubmittedTournaments
	if len(tournaments) != 331 {
		t.Errorf("expected 331 submitted_tournaments, got %d", len(tournaments))
	}

	// Verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 106024 {
		t.Errorf("expected first tournament_id 106024, got %v", first.TournamentId)
	}
	if first.EventEndDate == nil || *first.EventEndDate != "2025-12-31" {
		t.Errorf("expected first event_end_date '2025-12-31', got %v", first.EventEndDate)
	}
	if first.WpprPoints == nil || *first.WpprPoints != types.StringFloat64(0.0) {
		t.Errorf("expected first wppr_points 0.0000, got %v", first.WpprPoints)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 10 {
		t.Errorf("expected first player_count 10, got %v", first.PlayerCount)
	}
}

func TestSeriesRegionTourRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesTournamentsACSQLDResponse)
	ctx := context.Background()
	_, err := client.SeriesRegionTourWithResponse(ctx, "ACS", &SeriesRegionTourParams{
		RegionCode: "QLD",
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/tournaments" {
		t.Errorf("expected path '/series/ACS/tournaments', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("region_code") != "QLD" {
		t.Errorf("expected query param region_code 'QLD', got '%s'", req.URL.Query().Get("region_code"))
	}
}
