package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_women_open.json
var sampleRankingsWomenOpenResponse string

//go:embed testdata/rankings_women_women.json
var sampleRankingsWomenWomenResponse string

func TestRankingsWomenOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsWomenOpenResponse)
	ctx := context.Background()
	resp, err := client.RankingWomenOpenWithResponse(ctx, RankingWomenOpenParamsTournamentTypeOPEN, nil)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "women" {
		t.Errorf("expected ranking_type 'women', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.TournamentType == nil || *resp.JSON200.TournamentType != "open" {
		t.Errorf("expected tournament_type 'open', got %v", resp.JSON200.TournamentType)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 5936 {
		t.Errorf("expected total_count 5936, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 17 {
		t.Errorf("expected 17 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - Leslie Ruckman, has real age value
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 47411 {
		t.Errorf("expected first player_id 47411, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Leslie Ruckman" {
		t.Errorf("expected first name 'Leslie Ruckman', got %v", first.Name)
	}
	if first.Age == nil {
		t.Error("expected first age to be non-nil")
	} else if first.Age.IsEmpty || first.Age.Value != 38 {
		t.Errorf("expected first age 38, got IsEmpty=%v Value=%d", first.Age.IsEmpty, first.Age.Value)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil {
		t.Error("expected first current_rank to be non-nil")
	} else if first.CurrentRank.IsEmpty || first.CurrentRank.Value != 1 {
		t.Errorf("expected first current_rank 1, got IsEmpty=%v Value=%d", first.CurrentRank.IsEmpty, first.CurrentRank.Value)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 887.7625 {
		t.Errorf("expected first wppr_points 887.7625, got %v", first.WpprPoints)
	}
	if first.CurrentWpprRank == nil || int(*first.CurrentWpprRank) != 139 {
		t.Errorf("expected first current_wppr_rank 139, got %v", first.CurrentWpprRank)
	}
	if first.LastMonthWpprRank == nil || int(*first.LastMonthWpprRank) != 144 {
		t.Errorf("expected first last_month_wppr_rank 144, got %v", first.LastMonthWpprRank)
	}
	if first.EventCount == nil || int(*first.EventCount) != 582 {
		t.Errorf("expected first event_count 582, got %v", first.EventCount)
	}
	if first.BestFinish == nil || *first.BestFinish != "NW Pinball Championships" {
		t.Errorf("expected first best_finish 'NW Pinball Championships', got %v", first.BestFinish)
	}
	if first.BestFinishPosition == nil || int(*first.BestFinishPosition) != 8 {
		t.Errorf("expected first best_finish_position 8, got %v", first.BestFinishPosition)
	}
	if first.BestTournamentId == nil || int(*first.BestTournamentId) != 55790 {
		t.Errorf("expected first best_tournament_id 55790, got %v", first.BestTournamentId)
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

func TestRankingsWomenWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleRankingsWomenWomenResponse)
	ctx := context.Background()
	resp, err := client.RankingWomenOpenWithResponse(ctx, RankingWomenOpenParamsTournamentTypeWOMEN, nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "women" {
		t.Errorf("expected ranking_type 'women', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.TournamentType == nil || *resp.JSON200.TournamentType != "women" {
		t.Errorf("expected tournament_type 'women', got %v", resp.JSON200.TournamentType)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 9176 {
		t.Errorf("expected total_count 9176, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 17 {
		t.Errorf("expected 17 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - Jane Verwys, has empty age
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 53105 {
		t.Errorf("expected first player_id 53105, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Jane Verwys" {
		t.Errorf("expected first name 'Jane Verwys', got %v", first.Name)
	}
	if first.Age == nil {
		t.Error("expected first age to be non-nil")
	} else if !first.Age.IsEmpty {
		t.Errorf("expected first age to be empty, got value %d", first.Age.Value)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	// current_rank is quoted "1" in women tournament_type
	if first.CurrentRank == nil {
		t.Error("expected first current_rank to be non-nil")
	} else if first.CurrentRank.IsEmpty || first.CurrentRank.Value != 1 {
		t.Errorf("expected first current_rank 1, got IsEmpty=%v Value=%d", first.CurrentRank.IsEmpty, first.CurrentRank.Value)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 875.615 {
		t.Errorf("expected first wppr_points 875.615, got %v", first.WpprPoints)
	}

	// Verify second entry has a real age value (Leslie Ruckman, age 38)
	second := rankings[1]
	if second.Age == nil {
		t.Error("expected second age to be non-nil")
	} else if second.Age.IsEmpty || second.Age.Value != 38 {
		t.Errorf("expected second age 38, got IsEmpty=%v Value=%d", second.Age.IsEmpty, second.Age.Value)
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

func TestRankingsWomenRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsWomenOpenResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingWomenOpenParams{StartPos: &startPos, Count: &count}
	_, err := client.RankingWomenOpenWithResponse(ctx, RankingWomenOpenParamsTournamentTypeOPEN, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/women/OPEN" {
		t.Errorf("expected path '/rankings/women/OPEN', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
