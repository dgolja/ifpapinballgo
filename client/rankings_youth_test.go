package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_youth.json
var sampleRankingsYouthResponse string

func TestRankingsYouth(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsYouthResponse)
	ctx := context.Background()
	resp, err := client.RankingYouthWithResponse(ctx, nil)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "youth" {
		t.Errorf("expected ranking_type 'youth', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 332 {
		t.Errorf("expected total_count 332, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 23 {
		t.Errorf("expected 23 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1)
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 49549 {
		t.Errorf("expected first player_id 49549, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Arvid Flygare" {
		t.Errorf("expected first name 'Arvid Flygare', got %v", first.Name)
	}
	if first.Age == nil || *first.Age != 17 {
		t.Errorf("expected first age 17, got %v", first.Age)
	}
	if first.CountryCode == nil || *first.CountryCode != "SE" {
		t.Errorf("expected first country_code 'SE', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil || *first.CurrentRank != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.CurrentWpprRank == nil || int(*first.CurrentWpprRank) != 3 {
		t.Errorf("expected first current_wppr_rank 3, got %v", first.CurrentWpprRank)
	}
	if first.BestFinish == nil || *first.BestFinish != "IFPA World Pinball Championship" {
		t.Errorf("expected first best_finish 'IFPA World Pinball Championship', got %v", first.BestFinish)
	}
	if first.BestFinishPosition == nil || int(*first.BestFinishPosition) != 2 {
		t.Errorf("expected first best_finish_position 2, got %v", first.BestFinishPosition)
	}
	if first.BestTournamentId == nil || int(*first.BestTournamentId) != 78171 {
		t.Errorf("expected first best_tournament_id 78171, got %v", first.BestTournamentId)
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

func TestRankingsYouthRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsYouthResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingYouthParams{StartPos: &startPos, Count: &count}
	_, err := client.RankingYouthWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/youth" {
		t.Errorf("expected path '/rankings/youth', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
