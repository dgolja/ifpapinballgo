package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_id_related.json
var sampleTournamentRelatedResponse string

func TestTournamentRelated(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentRelatedResponse)
	ctx := context.Background()

	resp, err := client.TourRelatedWithResponse(ctx, 73485)
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

	if resp.JSON200.Tournament == nil {
		t.Fatal("expected tournament to be non-nil")
	}
	tournaments := *resp.JSON200.Tournament
	if len(tournaments) != 10 {
		t.Errorf("expected 10 tournament entries, got %d", len(tournaments))
	}

	// verify first entry (future tournament — no winner yet)
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 105255 {
		t.Errorf("expected first tournament_id 105255, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Brisbane Masters" {
		t.Errorf("expected first tournament_name 'Brisbane Masters', got %v", first.TournamentName)
	}
	if first.EventName == nil || *first.EventName != "Main Tournament" {
		t.Errorf("expected first event_name 'Main Tournament', got %v", first.EventName)
	}
	if first.RankingSystem == nil || *first.RankingSystem != "MAIN" {
		t.Errorf("expected first ranking_system 'MAIN', got %v", first.RankingSystem)
	}
	if first.EventStartDate == nil || *first.EventStartDate != "2026-06-20" {
		t.Errorf("expected first event_start_date '2026-06-20', got %v", first.EventStartDate)
	}

	// future tournament: winner fields should be nil (empty string in JSON)
	if first.Winner != nil && first.Winner.PlayerId != nil {
		t.Errorf("expected winner.player_id to be nil for future tournament, got %v", first.Winner.PlayerId)
	}

	// verify second entry (future women's division)
	second := tournaments[1]
	if second.TournamentId == nil || int(*second.TournamentId) != 105256 {
		t.Errorf("expected second tournament_id 105256, got %v", second.TournamentId)
	}
	if second.RankingSystem == nil || *second.RankingSystem != "WOMEN" {
		t.Errorf("expected second ranking_system 'WOMEN', got %v", second.RankingSystem)
	}
	if second.Winner != nil && second.Winner.PlayerId != nil {
		t.Errorf("expected winner.player_id to be nil for future tournament, got %v", second.Winner.PlayerId)
	}

	// verify third entry (past tournament with winner)
	third := tournaments[2]
	if third.TournamentId == nil || int(*third.TournamentId) != 73485 {
		t.Errorf("expected third tournament_id 73485, got %v", third.TournamentId)
	}
	if third.Winner == nil {
		t.Fatal("expected winner to be non-nil for past tournament")
	}
	if third.Winner.PlayerId == nil || int(*third.Winner.PlayerId) != 1605 {
		t.Errorf("expected winner player_id 1605, got %v", third.Winner.PlayerId)
	}
	if third.Winner.Name == nil || *third.Winner.Name != "Escher Lefkoff" {
		t.Errorf("expected winner name 'Escher Lefkoff', got %v", third.Winner.Name)
	}
	if third.Winner.CountryName == nil || *third.Winner.CountryName != "Australia" {
		t.Errorf("expected winner country_name 'Australia', got %v", third.Winner.CountryName)
	}
	if third.Winner.CountryCode == nil || *third.Winner.CountryCode != "AU" {
		t.Errorf("expected winner country_code 'AU', got %v", third.Winner.CountryCode)
	}

	// verify last entry
	last := tournaments[9]
	if last.TournamentId == nil || int(*last.TournamentId) != 10383 {
		t.Errorf("expected last tournament_id 10383, got %v", last.TournamentId)
	}
	if last.Winner == nil || last.Winner.Name == nil || *last.Winner.Name != "Peter Watt" {
		t.Errorf("expected last winner name 'Peter Watt', got %v", last.Winner)
	}
}

func TestTournamentRelatedRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentRelatedResponse)
	ctx := context.Background()

	_, err := client.TourRelatedWithResponse(ctx, 73485)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/73485/related" {
		t.Errorf("expected path '/tournament/73485/related', got %s", req.URL.Path)
	}
}
