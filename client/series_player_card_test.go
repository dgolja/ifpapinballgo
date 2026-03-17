package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/dgolja/ifpapinballgo/types"
)

//go:embed testdata/series_player_card.json
var sampleSeriesPlayerCardResponse string

const TestSeriesPlayerCardPlayerID = float32(90228)

func TestSeriesPlayersCard(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesPlayerCardResponse)
	ctx := context.Background()
	regionCode := "NSW"
	resp, err := client.SeriesPlayersCardWithResponse(ctx, "ACS", TestSeriesPlayerCardPlayerID, &SeriesPlayersCardParams{
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

	if resp.JSON200.PlayerId == nil || int(*resp.JSON200.PlayerId) != 90228 {
		t.Errorf("expected player_id 90228, got %v", resp.JSON200.PlayerId)
	}

	if resp.JSON200.PlayerName == nil || *resp.JSON200.PlayerName != "Andrew Gliatis" {
		t.Errorf("expected player_name 'Andrew Gliatis', got %v", resp.JSON200.PlayerName)
	}

	if resp.JSON200.PlayerCard == nil {
		t.Fatal("expected player_card to be non-nil")
	}

	card := *resp.JSON200.PlayerCard
	if len(card) != 20 {
		t.Errorf("expected 20 player_card entries, got %d", len(card))
	}

	// Verify first entry
	first := card[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 92732 {
		t.Errorf("expected first tournament_id 92732, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Pinawarra 31" {
		t.Errorf("expected first tournament_name 'Pinawarra 31', got %v", first.TournamentName)
	}
	if first.EventEndDate == nil || *first.EventEndDate != "2025-12-13" {
		t.Errorf("expected first event_end_date '2025-12-13', got %v", first.EventEndDate)
	}
	if first.WpprPoints == nil || *first.WpprPoints != types.StringFloat64(18.28) {
		t.Errorf("expected first wppr_points 18.28, got %v", first.WpprPoints)
	}
	if first.RegionEventRank == nil || int(*first.RegionEventRank) != 12 {
		t.Errorf("expected first region_event_rank 12, got %v", first.RegionEventRank)
	}

	// Verify all entries have required fields
	for i, entry := range card {
		if entry.TournamentId == nil {
			t.Errorf("card[%d]: expected tournament_id to be non-nil", i)
		}
		if entry.TournamentName == nil || *entry.TournamentName == "" {
			t.Errorf("card[%d]: expected tournament_name to be non-empty", i)
		}
		if entry.WpprPoints == nil {
			t.Errorf("card[%d]: expected wppr_points to be non-nil", i)
		}
		if entry.RegionEventRank == nil {
			t.Errorf("card[%d]: expected region_event_rank to be non-nil", i)
		}
	}
}

func TestSeriesPlayersCardRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesPlayerCardResponse)
	ctx := context.Background()
	regionCode := "NSW"
	_, err := client.SeriesPlayersCardWithResponse(ctx, "ACS", TestSeriesPlayerCardPlayerID, &SeriesPlayersCardParams{
		RegionCode: regionCode,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/player_card/90228" {
		t.Errorf("expected path '/series/ACS/player_card/90228', got %s", req.URL.Path)
	}
}
