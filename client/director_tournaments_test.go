package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/director_tournaments_data_future.json
var sampleDirectorTournamentsFutureResponse string

//go:embed testdata/director_tournaments_data_past.json
var sampleDirectorTournamentsPastResponse string

func TestDirectorTournamentsFuture(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorTournamentsFutureResponse)
	ctx := context.Background()
	resp, err := client.ViewDirectorToursWithResponse(ctx, float32(3478), "FUTURE")
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

	if resp.JSON200.DirectorId == nil || int(*resp.JSON200.DirectorId) != 3478 {
		t.Errorf("expected director_id 3478, got %v", resp.JSON200.DirectorId)
	}
	if *resp.JSON200.TournamentCount != 5 {
		t.Errorf("expected tournament_count 5, got %d", int(*resp.JSON200.TournamentCount))
	}

	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}
	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 5 {
		t.Errorf("expected 5 tournaments, got %d", len(tournaments))
	}

	// verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 105469 {
		t.Errorf("expected first tournament_id 105469, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Flipper Fellowship Thursday Group Matchplay" {
		t.Errorf("expected first tournament_name 'Flipper Fellowship Thursday Group Matchplay', got %v", first.TournamentName)
	}
	if first.RankingSystem == nil || *first.RankingSystem != "MAIN" {
		t.Errorf("expected first ranking_system 'MAIN', got %v", first.RankingSystem)
	}
	if first.City == nil || *first.City != "Arncliffe" {
		t.Errorf("expected first city 'Arncliffe', got %v", first.City)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.EventStartDate == nil || *first.EventStartDate != "2026-03-05" {
		t.Errorf("expected first event_start_date '2026-03-05', got %v", first.EventStartDate)
	}

	// future tournaments should not have player_count
	for i, tour := range tournaments {
		if tour.PlayerCount != nil {
			t.Errorf("tournaments[%d]: expected player_count to be nil for future tournament, got %v", i, tour.PlayerCount)
		}
	}
}

func TestDirectorTournamentsPast(t *testing.T) {
	client, _ := newTestClient(t, sampleDirectorTournamentsPastResponse)
	ctx := context.Background()
	resp, err := client.ViewDirectorToursWithResponse(ctx, float32(3478), "PAST")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.DirectorId == nil || int(*resp.JSON200.DirectorId) != 3478 {
		t.Errorf("expected director_id 3478, got %v", resp.JSON200.DirectorId)
	}
	if *resp.JSON200.TournamentCount != 4 {
		t.Errorf("expected tournament_count 4, got %d", int(*resp.JSON200.TournamentCount))
	}

	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}
	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 4 {
		t.Errorf("expected 4 tournaments, got %d", len(tournaments))
	}

	// verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 105464 {
		t.Errorf("expected first tournament_id 105464, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Flipper Fellowship Flip Frenzy" {
		t.Errorf("expected first tournament_name 'Flipper Fellowship Flip Frenzy', got %v", first.TournamentName)
	}
	if first.EventEndDate == nil || *first.EventEndDate != "2026-02-26" {
		t.Errorf("expected first event_end_date '2026-02-26', got %v", first.EventEndDate)
	}

	// past tournaments must have player_count
	for i, tour := range tournaments {
		if tour.PlayerCount == nil {
			t.Errorf("tournaments[%d]: expected player_count to be non-nil for past tournament", i)
		}
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 24 {
		t.Errorf("expected first player_count 24, got %v", first.PlayerCount)
	}
}

func TestDirectorTournamentsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorTournamentsFutureResponse)
	ctx := context.Background()
	_, err := client.ViewDirectorToursWithResponse(ctx, float32(3478), "FUTURE")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/director/3478/tournaments/FUTURE" {
		t.Errorf("expected path '/director/3478/tournaments/FUTURE', got %s", req.URL.Path)
	}
}
