package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_virtual.json
var sampleRankingsVirtualResponse string

func TestRankingsVirtual(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsVirtualResponse)
	ctx := context.Background()
	resp, err := client.RankingVirtualWithResponse(ctx, nil)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "virtual" {
		t.Errorf("expected ranking_type 'virtual', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 2952 {
		t.Errorf("expected total_count 2952, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.SortOrder == nil || *resp.JSON200.SortOrder != "points" {
		t.Errorf("expected sort_order 'points', got %v", resp.JSON200.SortOrder)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 13 {
		t.Errorf("expected 13 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - has empty age
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 134350 {
		t.Errorf("expected first player_id 134350, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "MadBenHan  " {
		t.Errorf("expected first name 'MadBenHan  ', got %v", first.Name)
	}
	if first.Age == nil {
		t.Error("expected first age to be non-nil")
	} else if !first.Age.IsEmpty {
		t.Errorf("expected first age to be empty, got value %d", first.Age.Value)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil || int(*first.CurrentRank) != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 2735.0725 {
		t.Errorf("expected first wppr_points 2735.0725, got %v", first.WpprPoints)
	}
	if first.TotalWinsLast3Years == nil || int(*first.TotalWinsLast3Years) != 42 {
		t.Errorf("expected first total_wins_last_3_years 42, got %v", first.TotalWinsLast3Years)
	}
	if first.Top3Last3Years == nil || int(*first.Top3Last3Years) != 60 {
		t.Errorf("expected first top_3_last_3_years 60, got %v", first.Top3Last3Years)
	}
	if first.Top10Last3Years == nil || int(*first.Top10Last3Years) != 72 {
		t.Errorf("expected first top_10_last_3_years 72, got %v", first.Top10Last3Years)
	}

	// Verify second entry (rank 2) - has a real age value
	second := rankings[1]
	if second.Age == nil {
		t.Error("expected second age to be non-nil")
	} else if second.Age.IsEmpty || second.Age.Value != 31 {
		t.Errorf("expected second age 31, got IsEmpty=%v Value=%d", second.Age.IsEmpty, second.Age.Value)
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

func TestRankingsVirtualRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsVirtualResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingVirtualParams{StartPos: &startPos, Count: &count}
	_, err := client.RankingVirtualWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/virtual" {
		t.Errorf("expected path '/rankings/virtual', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
