package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_custom_id.json
var sampleRankingsCustomIDResponse string

func TestRankingsCustomID(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCustomIDResponse)
	ctx := context.Background()
	resp, err := client.RankingCustomIDWithResponse(ctx, 365, nil)
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

	if resp.JSON200.ViewId == nil || int(*resp.JSON200.ViewId) != 365 {
		t.Errorf("expected view_id 365, got %v", resp.JSON200.ViewId)
	}
	if resp.JSON200.Title == nil || *resp.JSON200.Title != "2024 - IFPA Slovenian Championship Series" {
		t.Errorf("expected title '2024 - IFPA Slovenian Championship Series', got %v", resp.JSON200.Title)
	}
	if resp.JSON200.StartPosition == nil || int(*resp.JSON200.StartPosition) != 1 {
		t.Errorf("expected start_position 1, got %v", resp.JSON200.StartPosition)
	}
	if resp.JSON200.ReturnCount == nil || int(*resp.JSON200.ReturnCount) != 13 {
		t.Errorf("expected return_count 13, got %v", resp.JSON200.ReturnCount)
	}
	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 109 {
		t.Errorf("expected total_count 109, got %v", resp.JSON200.TotalCount)
	}

	// Verify view_results
	if resp.JSON200.ViewResults == nil {
		t.Fatal("expected view_results to be non-nil")
	}

	results := *resp.JSON200.ViewResults
	if len(results) != 13 {
		t.Errorf("expected 13 view_results entries, got %d", len(results))
	}

	// Verify first entry - Arno Nöbl
	first := results[0]
	if first.PlayerId == nil || int(*first.PlayerId) != 21106 {
		t.Errorf("expected first player_id 21106, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "Arno Nöbl" {
		t.Errorf("expected first name 'Arno Nöbl', got %v", first.Name)
	}
	if first.CountryCode == nil || *first.CountryCode != "AT" {
		t.Errorf("expected first country_code 'AT', got %v", first.CountryCode)
	}
	if first.WpprRank == nil || int(*first.WpprRank) != 454 {
		t.Errorf("expected first wppr_rank 454, got %v", first.WpprRank)
	}
	if first.WpprPoints == nil || float32(*first.WpprPoints) != 90.82 {
		t.Errorf("expected first wppr_points 90.82, got %v", first.WpprPoints)
	}
	if first.EventCount == nil || int(*first.EventCount) != 3 {
		t.Errorf("expected first event_count 3, got %v", first.EventCount)
	}
	if first.Position == nil || int(*first.Position) != 1 {
		t.Errorf("expected first position 1, got %v", first.Position)
	}

	// Verify all view_results entries have required fields
	for i, r := range results {
		if r.PlayerId == nil {
			t.Errorf("view_results[%d]: expected player_id to be non-nil", i)
		}
		if r.Name == nil || *r.Name == "" {
			t.Errorf("view_results[%d]: expected name to be non-empty", i)
		}
		if r.Position == nil {
			t.Errorf("view_results[%d]: expected position to be non-nil", i)
		}
	}

	// Verify tournaments
	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}

	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 14 {
		t.Errorf("expected 14 tournament entries, got %d", len(tournaments))
	}

	firstTournament := tournaments[0]
	if firstTournament.TournamentId == nil || int(*firstTournament.TournamentId) != 71194 {
		t.Errorf("expected first tournament_id 71194, got %v", firstTournament.TournamentId)
	}
	if firstTournament.TournamentName == nil || *firstTournament.TournamentName != "LOO-BLAH-NAH 2024 kickoff" {
		t.Errorf("expected first tournament_name 'LOO-BLAH-NAH 2024 kickoff', got %v", firstTournament.TournamentName)
	}
	if firstTournament.EventDate == nil || *firstTournament.EventDate != "2024-01-13" {
		t.Errorf("expected first event_date '2024-01-13', got %v", firstTournament.EventDate)
	}

	// Verify view_filters
	if resp.JSON200.ViewFilters == nil {
		t.Fatal("expected view_filters to be non-nil")
	}

	filters := *resp.JSON200.ViewFilters
	if len(filters) != 2 {
		t.Errorf("expected 2 view_filters entries, got %d", len(filters))
	}
	if filters[0].Name == nil || *filters[0].Name != "Tournament Year" {
		t.Errorf("expected first filter name 'Tournament Year', got %v", filters[0].Name)
	}
	if filters[0].Setting == nil || *filters[0].Setting != "2024" {
		t.Errorf("expected first filter setting '2024', got %v", filters[0].Setting)
	}
}

func TestRankingsCustomIDRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCustomIDResponse)
	ctx := context.Background()
	startPos := float32(1)
	count := float32(25)
	params := &RankingCustomIDParams{StartPos: &startPos, Count: &count}
	_, err := client.RankingCustomIDWithResponse(ctx, 365, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/custom/365" {
		t.Errorf("expected path '/rankings/custom/365', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("start_pos") != "1" {
		t.Errorf("expected query param start_pos=1, got %s", req.URL.Query().Get("start_pos"))
	}
	if req.URL.Query().Get("count") != "25" {
		t.Errorf("expected query param count=25, got %s", req.URL.Query().Get("count"))
	}
}
