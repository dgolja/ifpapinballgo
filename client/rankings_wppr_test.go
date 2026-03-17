package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_wppr.json
var sampleRankingsWpprResponse string

func TestRankingsWppr(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsWpprResponse)
	ctx := context.Background()
	resp, err := client.RankingWpprWithResponse(ctx, nil)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "wppr" {
		t.Errorf("expected ranking_type 'wppr', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 68022 {
		t.Errorf("expected total_count 68022, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.SortOrder == nil || *resp.JSON200.SortOrder != "points" {
		t.Errorf("expected sort_order 'points', got %v", resp.JSON200.SortOrder)
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
	if first.PlayerId == nil || int(*first.PlayerId) != 11938 {
		t.Errorf("expected first player_id 11938, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Jason Zahler" {
		t.Errorf("expected first name 'Jason Zahler', got %v", first.Name)
	}
	if first.Age == nil || first.Age.Value != 21 {
		t.Errorf("expected first age 21, got %v", first.Age)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil || int(*first.CurrentRank) != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 3347.0775 {
		t.Errorf("expected first wppr_points 3347.0775, got %v", first.WpprPoints)
	}
	if first.BestFinish == nil || *first.BestFinish != "IFPA World Pinball Championship" {
		t.Errorf("expected first best_finish 'IFPA World Pinball Championship', got %v", first.BestFinish)
	}
	if first.BestFinishPosition == nil || int(*first.BestFinishPosition) != 1 {
		t.Errorf("expected first best_finish_position 1, got %v", first.BestFinishPosition)
	}
	if first.TotalWinsLast3Years == nil || int(*first.TotalWinsLast3Years) != 26 {
		t.Errorf("expected first total_wins_last_3_years 26, got %v", first.TotalWinsLast3Years)
	}
	if first.Top3Last3Years == nil || int(*first.Top3Last3Years) != 42 {
		t.Errorf("expected first top_3_last_3_years 42, got %v", first.Top3Last3Years)
	}
	if first.Top10Last3Years == nil || int(*first.Top10Last3Years) != 52 {
		t.Errorf("expected first top_10_last_3_years 52, got %v", first.Top10Last3Years)
	}

	// Verify an entry with empty age (rank 9, Gregory Kennedy)
	rank9 := rankings[8]
	if rank9.Age == nil {
		t.Error("expected rank 9 age to be non-nil")
	} else if !rank9.Age.IsEmpty {
		t.Errorf("expected rank 9 age to be empty, got value %d", rank9.Age.Value)
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

func TestRankingsWpprRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsWpprResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingWpprParams{StartPos: &startPos, Count: &count}
	_, err := client.RankingWpprWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/wppr" {
		t.Errorf("expected path '/rankings/wppr', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
