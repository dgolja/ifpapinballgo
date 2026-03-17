package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_country.json
var sampleRankingsCountryResponse string

func TestRankingsCountry(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCountryResponse)
	ctx := context.Background()
	params := &RankingCountryParams{Country: "AU"}
	resp, err := client.RankingCountryWithResponse(ctx, params)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "country" {
		t.Errorf("expected ranking_type 'country', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.RankCountryName == nil || *resp.JSON200.RankCountryName != "Australia" {
		t.Errorf("expected rank_country_name 'Australia', got %v", resp.JSON200.RankCountryName)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 3487 {
		t.Errorf("expected total_count 3487, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.ReturnCount == nil || int(*resp.JSON200.ReturnCount) != 24 {
		t.Errorf("expected return_count 24, got %v", resp.JSON200.ReturnCount)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 24 {
		t.Errorf("expected 24 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - Escher Lefkoff
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 1605 {
		t.Errorf("expected first player_id 1605, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Escher Lefkoff" {
		t.Errorf("expected first name 'Escher Lefkoff', got %v", first.Name)
	}
	if first.Age == nil {
		t.Error("expected first age to be non-nil")
	} else if first.Age.IsEmpty || first.Age.Value != 22 {
		t.Errorf("expected first age 22, got IsEmpty=%v Value=%d", first.Age.IsEmpty, first.Age.Value)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil || int(*first.CurrentRank) != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 2980.3725 {
		t.Errorf("expected first wppr_points 2980.3725, got %v", first.WpprPoints)
	}
	if first.CurrentWpprRank == nil || int(*first.CurrentWpprRank) != 4 {
		t.Errorf("expected first current_wppr_rank 4, got %v", first.CurrentWpprRank)
	}
	if first.LastMonthRank == nil || int(*first.LastMonthRank) != 4 {
		t.Errorf("expected first last_month_rank 4, got %v", first.LastMonthRank)
	}
	if first.Rating == nil || float32(*first.Rating) != 2024.097 {
		t.Errorf("expected first rating 2024.097, got %v", first.Rating)
	}
	if first.RatingDeviation == nil || int(*first.RatingDeviation) != 59 {
		t.Errorf("expected first rating_deviation 59, got %v", first.RatingDeviation)
	}
	if first.EfficiencyPercent == nil || float32(*first.EfficiencyPercent) != 76.57 {
		t.Errorf("expected first efficiency_percent 76.57, got %v", first.EfficiencyPercent)
	}
	if first.EventCount == nil || int(*first.EventCount) != 442 {
		t.Errorf("expected first event_count 442, got %v", first.EventCount)
	}
	if first.BestFinish == nil || *first.BestFinish != "The Open - IFPA World Championship" {
		t.Errorf("expected first best_finish 'The Open - IFPA World Championship', got %v", first.BestFinish)
	}
	if first.BestFinishPosition == nil || int(*first.BestFinishPosition) != 3 {
		t.Errorf("expected first best_finish_position 3, got %v", first.BestFinishPosition)
	}
	if first.BestTournamentId == nil || int(*first.BestTournamentId) != 83318 {
		t.Errorf("expected first best_tournament_id 83318, got %v", first.BestTournamentId)
	}

	// Verify all entries have required fields
	for i, r := range rankings {
		if r.PlayerId == nil {
			t.Errorf("rankings[%d]: expected player_id to be non-nil", i)
		}
		if r.Name == nil || *r.Name == "" {
			t.Errorf("rankings[%d]: expected name to be non-empty", i)
		}
		if r.CountryCode == nil || *r.CountryCode == "" {
			t.Errorf("rankings[%d]: expected country_code to be non-empty", i)
		}
		if r.CurrentRank == nil {
			t.Errorf("rankings[%d]: expected current_rank to be non-nil", i)
		}
		if r.WpprPoints == nil {
			t.Errorf("rankings[%d]: expected wppr_points to be non-nil", i)
		}
	}
}

func TestRankingsCountryRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCountryResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingCountryParams{Country: "AU", StartPos: &startPos, Count: &count}
	_, err := client.RankingCountryWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/country" {
		t.Errorf("expected path '/rankings/country', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("country") != "AU" {
		t.Errorf("expected query param country=AU, got %s", req.URL.Query().Get("country"))
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
