package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_pro_open.json
var sampleRankingsProOpenResponse string

//go:embed testdata/rankings_pro_women.json
var sampleRankingsProWomenResponse string

func TestRankingsProOpen(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsProOpenResponse)
	ctx := context.Background()
	resp, err := client.RankingProWithResponse(ctx, RankingProParamsRankingSystemOPEN)
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

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "pro" {
		t.Errorf("expected ranking_type 'pro', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 248 {
		t.Errorf("expected 248 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - Jason Zahler
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 11938 {
		t.Errorf("expected first player_id 11938, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Jason Zahler" {
		t.Errorf("expected first name 'Jason Zahler', got %v", first.Name)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.CurrentRank == nil || int(*first.CurrentRank) != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.ProPoints == nil || float32(*first.ProPoints) != 3199.76 {
		t.Errorf("expected first pro_points 3199.76, got %v", first.ProPoints)
	}
	if first.OrginalWpprPoints == nil || float32(*first.OrginalWpprPoints) != 3347.08 {
		t.Errorf("expected first orginal_wppr_points 3347.08, got %v", first.OrginalWpprPoints)
	}
	if first.EfficiencyPercent == nil || float32(*first.EfficiencyPercent) != 81.89 {
		t.Errorf("expected first efficiency_percent 81.89, got %v", first.EfficiencyPercent)
	}
	if first.AdjEfficiencyPercent == nil || float32(*first.AdjEfficiencyPercent) != 83.08 {
		t.Errorf("expected first adj_efficiency_percent 83.08, got %v", first.AdjEfficiencyPercent)
	}
	if first.ExcessPercent == nil || float32(*first.ExcessPercent) != 26.0117 {
		t.Errorf("expected first excess_percent 26.0117, got %v", first.ExcessPercent)
	}
	if first.Wpprtunity == nil || float32(*first.Wpprtunity) != 4028.77 {
		t.Errorf("expected first wpprtunity 4028.77, got %v", first.Wpprtunity)
	}
	if first.WpprAdjustment == nil || float32(*first.WpprAdjustment) != 147.32 {
		t.Errorf("expected first wppr_adjustment 147.32, got %v", first.WpprAdjustment)
	}
	if first.SosPercent == nil || *first.SosPercent != 25.36 {
		t.Errorf("expected first sos_percent 25.36, got %v", first.SosPercent)
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
		if r.ProPoints == nil {
			t.Errorf("rankings[%d]: expected pro_points to be non-nil", i)
		}
	}
}

func TestRankingsProWomen(t *testing.T) {
	client, _ := newTestClient(t, sampleRankingsProWomenResponse)
	ctx := context.Background()
	resp, err := client.RankingProWithResponse(ctx, RankingProParamsRankingSystemWOMEN)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.RankingType == nil || *resp.JSON200.RankingType != "pro" {
		t.Errorf("expected ranking_type 'pro', got %v", resp.JSON200.RankingType)
	}

	if resp.JSON200.Rankings == nil {
		t.Fatal("expected rankings to be non-nil")
	}

	rankings := *resp.JSON200.Rankings
	if len(rankings) != 100 {
		t.Errorf("expected 100 rankings entries, got %d", len(rankings))
	}

	// Verify first entry (rank 1) - Jane Verwys
	first := rankings[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 53105 {
		t.Errorf("expected first player_id 53105, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Jane Verwys" {
		t.Errorf("expected first name 'Jane Verwys', got %v", first.Name)
	}
	if first.CurrentRank == nil || int(*first.CurrentRank) != 1 {
		t.Errorf("expected first current_rank 1, got %v", first.CurrentRank)
	}
	if first.ProPoints == nil || float32(*first.ProPoints) != 700.5 {
		t.Errorf("expected first pro_points 700.50, got %v", first.ProPoints)
	}
	if first.SosPercent == nil || *first.SosPercent != 29.28 {
		t.Errorf("expected first sos_percent 29.28, got %v", first.SosPercent)
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
		if r.ProPoints == nil {
			t.Errorf("rankings[%d]: expected pro_points to be non-nil", i)
		}
	}
}

func TestRankingsProRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsProOpenResponse)
	ctx := context.Background()
	_, err := client.RankingProWithResponse(ctx, RankingProParamsRankingSystemOPEN)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/pro/OPEN" {
		t.Errorf("expected path '/rankings/pro/OPEN', got %s", req.URL.Path)
	}
}
