package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_info.json
var sampleTournamentInfoResponse string

//go:embed testdata/tournament_info_past.json
var sampleTournamentInfoPastResponse string

//go:embed testdata/tournament_info_future_rich.json
var sampleTournamentInfoFutureRichResponse string

//go:embed testdata/tournament_info_future_women.json
var sampleTournamentInfoFutureWomenResponse string

//go:embed testdata/tournament_info_future_youth.json
var sampleTournamentInfoFutureYouthResponse string

//go:embed testdata/tournament_info_future_registration.json
var sampleTournamentInfoFutureRegistrationResponse string

func TestTourInfoWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentInfoResponse)
	ctx := context.Background()
	resp, err := client.TourInfoWithResponse(ctx, 107536)
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

	info := resp.JSON200

	// Verify tournament identity
	if info.TournamentId == nil || int(*info.TournamentId) != 107536 {
		t.Errorf("expected tournament_id 107536, got %v", info.TournamentId)
	}
	if info.TournamentName == nil || *info.TournamentName != "Illawarra Frenzy 230426" {
		t.Errorf("expected tournament_name 'Illawarra Frenzy 230426', got %v", info.TournamentName)
	}
	if info.TournamentType == nil || *info.TournamentType != "Tournament" {
		t.Errorf("expected tournament_type 'Tournament', got %v", info.TournamentType)
	}

	// Verify boolean flags (quoted strings in the API response)
	if info.PrestigeFlag == nil || bool(*info.PrestigeFlag) != false {
		t.Errorf("expected prestige_flag false, got %v", info.PrestigeFlag)
	}
	if info.PrivateFlag == nil || bool(*info.PrivateFlag) != false {
		t.Errorf("expected private_flag false, got %v", info.PrivateFlag)
	}

	// Verify location
	if info.City == nil || *info.City != "Wollongong" {
		t.Errorf("expected city 'Wollongong', got %v", info.City)
	}
	if info.Stateprov == nil || *info.Stateprov != "NSW" {
		t.Errorf("expected stateprov 'NSW', got %v", info.Stateprov)
	}
	if info.CountryName == nil || *info.CountryName != "Australia" {
		t.Errorf("expected country_name 'Australia', got %v", info.CountryName)
	}
	if info.CountryCode == nil || *info.CountryCode != "AU" {
		t.Errorf("expected country_code 'AU', got %v", info.CountryCode)
	}

	// Verify director
	if info.DirectorId == nil || int(*info.DirectorId) != 2909 {
		t.Errorf("expected director_id 2909, got %v", info.DirectorId)
	}
	if info.DirectorName == nil || *info.DirectorName != "Paul Soutter" {
		t.Errorf("expected director_name 'Paul Soutter', got %v", info.DirectorName)
	}

	// Verify event details
	if info.EventName == nil || *info.EventName != "Main Tournament" {
		t.Errorf("expected event_name 'Main Tournament', got %v", info.EventName)
	}
	if info.EventStartDate == nil || *info.EventStartDate != "2026-04-23" {
		t.Errorf("expected event_start_date '2026-04-23', got %v", info.EventStartDate)
	}
	if info.EventEndDate == nil || *info.EventEndDate != "2026-04-23" {
		t.Errorf("expected event_end_date '2026-04-23', got %v", info.EventEndDate)
	}

	// Verify formats
	if info.QualifyingFormat == nil || *info.QualifyingFormat != "Flip Frenzy" {
		t.Errorf("expected qualifying_format 'Flip Frenzy', got %v", info.QualifyingFormat)
	}
	if info.FinalsFormat == nil || *info.FinalsFormat != "Match Play" {
		t.Errorf("expected finals_format 'Match Play', got %v", info.FinalsFormat)
	}
	if info.RankingSystem == nil || *info.RankingSystem != "MAIN" {
		t.Errorf("expected ranking_system 'MAIN', got %v", info.RankingSystem)
	}

	// Verify values and player counts now that the tournament has been played
	if info.TournamentValue == nil || float64(*info.TournamentValue) != 8.06 {
		t.Errorf("expected tournament_value 8.06, got %v", info.TournamentValue)
	}
	if info.PlayerCount == nil || int(*info.PlayerCount) != 19 {
		t.Errorf("expected player_count 19, got %v", info.PlayerCount)
	}
	if info.EligiblePlayerCount == nil || int(*info.EligiblePlayerCount) != 17 {
		t.Errorf("expected eligible_player_count 17, got %v", info.EligiblePlayerCount)
	}

	// Verify matchplay_id
	if info.MatchplayId == nil || int(*info.MatchplayId) != 0 {
		t.Errorf("expected matchplay_id 0, got %v", info.MatchplayId)
	}

	// Verify newer fields
	if info.EventClass == nil || *info.EventClass != "NORMAL" {
		t.Errorf("expected event_class 'NORMAL', got %v", info.EventClass)
	}
	if info.ReservedSpots == nil || int(*info.ReservedSpots) != 0 {
		t.Errorf("expected reserved_spots 0, got %v", info.ReservedSpots)
	}
	if info.LocationName != nil || info.RegistrationCost != nil || info.TournamentSystemUrl != nil {
		t.Errorf("expected null location_name, registration_cost and tournament_system_url to be nil")
	}
}

func TestTourInfoRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentInfoResponse)
	ctx := context.Background()
	_, err := client.TourInfoWithResponse(ctx, 107536)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/107536" {
		t.Errorf("expected path '/tournament/107536', got %s", req.URL.Path)
	}
}

func TestTourInfoPastWithResponse(t *testing.T) {
	client, _ := newTestClient(t, sampleTournamentInfoPastResponse)
	ctx := context.Background()
	resp, err := client.TourInfoWithResponse(ctx, 105469)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	info := resp.JSON200

	// Verify tournament identity
	if info.TournamentId == nil || int(*info.TournamentId) != 105469 {
		t.Errorf("expected tournament_id 105469, got %v", info.TournamentId)
	}
	if info.TournamentName == nil || *info.TournamentName != "Flipper Fellowship Thursday Group Matchplay" {
		t.Errorf("expected tournament_name 'Flipper Fellowship Thursday Group Matchplay', got %v", info.TournamentName)
	}

	// Verify location
	if info.City == nil || *info.City != "Arncliffe" {
		t.Errorf("expected city 'Arncliffe', got %v", info.City)
	}
	if info.CountryCode == nil || *info.CountryCode != "AU" {
		t.Errorf("expected country_code 'AU', got %v", info.CountryCode)
	}

	// Verify director
	if info.DirectorId == nil || int(*info.DirectorId) != 3478 {
		t.Errorf("expected director_id 3478, got %v", info.DirectorId)
	}
	if info.DirectorName == nil || *info.DirectorName != "Andrew Gliatis" {
		t.Errorf("expected director_name 'Andrew Gliatis', got %v", info.DirectorName)
	}

	// Verify tournament_value is non-nil for past (completed) tournaments
	if info.TournamentValue == nil {
		t.Errorf("expected tournament_value to be non-nil for a past tournament")
	}

	// Verify player counts
	if info.PlayerCount == nil || int(*info.PlayerCount) != 20 {
		t.Errorf("expected player_count 20, got %v", info.PlayerCount)
	}
	if info.EligiblePlayerCount == nil || int(*info.EligiblePlayerCount) != 19 {
		t.Errorf("expected eligible_player_count 19, got %v", info.EligiblePlayerCount)
	}

	// Verify formats
	if info.QualifyingFormat == nil || *info.QualifyingFormat != "Matchplay Qualifying" {
		t.Errorf("expected qualifying_format 'Matchplay Qualifying', got %v", info.QualifyingFormat)
	}
	if info.FinalsFormat == nil || *info.FinalsFormat != "Match Play" {
		t.Errorf("expected finals_format 'Match Play', got %v", info.FinalsFormat)
	}
}

func TestTourInfoFutureWithResponse(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		id           int
		rankingSys   string
		playerLimit  int
		playerCount  int
		startTime    string // "" means null
		tiebreaker   string // "" means null
		regDate      string
		costPresent  bool
		systemPreset bool
	}{
		{"rich", sampleTournamentInfoFutureRichResponse, 119159, "MAIN", 64, 42, "13:30:00", "1 Full Game", "", true, true},
		{"women", sampleTournamentInfoFutureWomenResponse, 115621, "WOMEN", 92, 0, "", "", "", false, false},
		{"youth", sampleTournamentInfoFutureYouthResponse, 119792, "YOUTH", 0, 0, "", "1 Full Game", "", false, false},
		{"registration", sampleTournamentInfoFutureRegistrationResponse, 120432, "MAIN", 30, 0, "", "1 Full Game", "2026-07-15", false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client, _ := newTestClient(t, tc.body)
			resp, err := client.TourInfoWithResponse(context.Background(), float32(tc.id))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.JSON200 == nil {
				t.Fatal("expected JSON200 to be non-nil")
			}
			info := resp.JSON200

			if info.TournamentId == nil || int(*info.TournamentId) != tc.id {
				t.Errorf("expected tournament_id %d, got %v", tc.id, info.TournamentId)
			}
			if info.RankingSystem == nil || *info.RankingSystem != tc.rankingSys {
				t.Errorf("expected ranking_system %s, got %v", tc.rankingSys, info.RankingSystem)
			}
			if info.PlayerLimit == nil || int(*info.PlayerLimit) != tc.playerLimit {
				t.Errorf("expected player_limit %d, got %v", tc.playerLimit, info.PlayerLimit)
			}
			// future tournaments have no computed value yet; player_count is the registered count
			if info.TournamentValue != nil {
				t.Errorf("expected tournament_value to be nil, got %v", info.TournamentValue)
			}
			if info.PlayerCount == nil || int(*info.PlayerCount) != tc.playerCount {
				t.Errorf("expected player_count %d, got %v", tc.playerCount, info.PlayerCount)
			}
			if info.EventClass == nil || *info.EventClass != "NORMAL" {
				t.Errorf("expected event_class NORMAL, got %v", info.EventClass)
			}

			if (tc.startTime == "") != (info.StartTime == nil) || (tc.startTime != "" && *info.StartTime != tc.startTime) {
				t.Errorf("expected start_time %q, got %v", tc.startTime, info.StartTime)
			}
			if (tc.tiebreaker == "") != (info.TiebreakerFormat == nil) || (tc.tiebreaker != "" && *info.TiebreakerFormat != tc.tiebreaker) {
				t.Errorf("expected tiebreaker_format %q, got %v", tc.tiebreaker, info.TiebreakerFormat)
			}
			if info.RegistrationDate == nil || *info.RegistrationDate != tc.regDate {
				t.Errorf("expected registration_date %q, got %v", tc.regDate, info.RegistrationDate)
			}
			if tc.costPresent {
				if info.RegistrationCost == nil || float64(*info.RegistrationCost) != 10 {
					t.Errorf("expected registration_cost 10, got %v", info.RegistrationCost)
				}
			}
			if tc.systemPreset && (info.TournamentSystem == nil || *info.TournamentSystem != "Match Play Events") {
				t.Errorf("expected tournament_system 'Match Play Events', got %v", info.TournamentSystem)
			}
		})
	}
}
