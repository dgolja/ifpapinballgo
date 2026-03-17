package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_results.json
var sampleTournamentResultsResponse string

func TestViewTournamentResultsWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentResultsResponse)
	ctx := context.Background()
	resp, err := client.ViewTournamentResultsWithResponse(ctx, 85746)
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

	// Verify top-level fields
	if resp.JSON200.TournamentId == nil || int(*resp.JSON200.TournamentId) != 85746 {
		t.Errorf("expected tournament_id 85746, got %v", resp.JSON200.TournamentId)
	}
	if resp.JSON200.RankingSystem == nil || *resp.JSON200.RankingSystem != "MAIN" {
		t.Errorf("expected ranking_system 'MAIN', got %v", resp.JSON200.RankingSystem)
	}

	if resp.JSON200.Results == nil {
		t.Fatal("expected results to be non-nil")
	}

	results := *resp.JSON200.Results
	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}

	// Verify first result (Escher Lefkoff, position 1)
	first := results[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 1605 {
		t.Errorf("expected first player_id 1605, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Escher Lefkoff" {
		t.Errorf("expected first name 'Escher Lefkoff', got %v", first.Name)
	}
	if first.Position == nil || int(*first.Position) != 1 {
		t.Errorf("expected first position 1, got %v", first.Position)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.ExcludedFlag == nil || bool(*first.ExcludedFlag) != false {
		t.Errorf("expected first excluded_flag false, got %v", first.ExcludedFlag)
	}
	if first.WpprProRank == nil || int(*first.WpprProRank) != 2 {
		t.Errorf("expected first wppr_pro_rank 2, got %v", first.WpprProRank)
	}

	// Verify a result with null wppr_pro_rank (Daniel McGorum, position 8)
	found := false
	for _, r := range results {
		if r.PlayerId != nil && int(*r.PlayerId) == 70499 {
			found = true
			if r.WpprProRank != nil {
				t.Errorf("expected wppr_pro_rank to be nil for player 70499, got %v", r.WpprProRank)
			}
			if r.PostWpprProRank != nil {
				t.Errorf("expected post_wppr_pro_rank to be nil for player 70499, got %v", r.PostWpprProRank)
			}
			break
		}
	}
	if !found {
		t.Error("expected to find player 70499 in results")
	}

	// Verify all results have required fields
	for i, r := range results {
		if r.PlayerId == nil {
			t.Errorf("results[%d]: expected player_id to be non-nil", i)
		}
		if r.Name == nil || *r.Name == "" {
			t.Errorf("results[%d]: expected name to be non-empty", i)
		}
		if r.Position == nil {
			t.Errorf("results[%d]: expected position to be non-nil", i)
		}
		if r.Points == nil {
			t.Errorf("results[%d]: expected points to be non-nil", i)
		}
		if r.CountryCode == nil || *r.CountryCode == "" {
			t.Errorf("results[%d]: expected country_code to be non-empty", i)
		}
	}
}

func TestViewTournamentResultsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentResultsResponse)
	ctx := context.Background()
	_, err := client.ViewTournamentResultsWithResponse(ctx, 85746)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/85746/results" {
		t.Errorf("expected path '/tournament/85746/results', got %s", req.URL.Path)
	}
}
