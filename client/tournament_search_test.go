package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_search_past_main_au.json
var sampleTournamentSearchPastMainAU string

//go:embed testdata/tournament_search_future_main_au.json
var sampleTournamentSearchFutureMainAU string

//go:embed testdata/tournament_search_past_women_us.json
var sampleTournamentSearchPastWomenUS string

func TestTournamentSearchPastMainAU(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentSearchPastMainAU)
	ctx := context.Background()

	country := "AU"
	stateprov := "NSW"
	startDate := "2026-01-01"
	endDate := "2026-02-01"
	rankType := MAIN
	eventType := Tournament
	params := &TourSearchParams{
		Country:   &country,
		Stateprov: &stateprov,
		StartDate: &startDate,
		EndDate:   &endDate,
		RankType:  &rankType,
		EventType: &eventType,
	}

	resp, err := client.TourSearchWithResponse(ctx, params)
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

	// verify total_results (quoted int in JSON)
	if resp.JSON200.TotalResults == nil || int(*resp.JSON200.TotalResults) != 13 {
		t.Errorf("expected total_results 13, got %v", resp.JSON200.TotalResults)
	}

	// verify search_filter
	if resp.JSON200.SearchFilter == nil {
		t.Fatal("expected search_filter to be non-nil")
	}
	sf := resp.JSON200.SearchFilter
	if sf.Country == nil || *sf.Country != "AU" {
		t.Errorf("expected search_filter.country 'AU', got %v", sf.Country)
	}
	if sf.Stateprov == nil || *sf.Stateprov != "NSW" {
		t.Errorf("expected search_filter.stateprov 'NSW', got %v", sf.Stateprov)
	}
	if sf.RankType == nil || *sf.RankType != "MAIN" {
		t.Errorf("expected search_filter.rank_type 'MAIN', got %v", sf.RankType)
	}

	// verify tournaments array
	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}
	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 12 {
		t.Errorf("expected 12 tournaments, got %d", len(tournaments))
	}

	// verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 92305 {
		t.Errorf("expected first tournament_id 92305, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Pinhaven 11" {
		t.Errorf("expected first tournament_name 'Pinhaven 11', got %v", first.TournamentName)
	}
	if first.EventName == nil || *first.EventName != "Main Tournament" {
		t.Errorf("expected first event_name 'Main Tournament', got %v", first.EventName)
	}
	if first.City == nil || *first.City != "Nowra" {
		t.Errorf("expected first city 'Nowra', got %v", first.City)
	}
	if first.Stateprov == nil || *first.Stateprov != "NSW" {
		t.Errorf("expected first stateprov 'NSW', got %v", first.Stateprov)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.RankingSystem == nil || *first.RankingSystem != "MAIN" {
		t.Errorf("expected first ranking_system 'MAIN', got %v", first.RankingSystem)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 14 {
		t.Errorf("expected first player_count 14, got %v", first.PlayerCount)
	}
	if first.DirectorId == nil || int(*first.DirectorId) != 2909 {
		t.Errorf("expected first director_id 2909, got %v", first.DirectorId)
	}
	if first.PrivateFlag == nil || bool(*first.PrivateFlag) != false {
		t.Errorf("expected first private_flag false, got %v", first.PrivateFlag)
	}
	if first.EventStartDate == nil || *first.EventStartDate != "2026-01-03" {
		t.Errorf("expected first event_start_date '2026-01-03', got %v", first.EventStartDate)
	}

	// verify winner for past tournament
	if first.Winner == nil {
		t.Fatal("expected winner to be non-nil for past tournament")
	}
	if first.Winner.PlayerId == nil || int(*first.Winner.PlayerId) != 119126 {
		t.Errorf("expected winner player_id 119126, got %v", first.Winner.PlayerId)
	}
	if first.Winner.PlayerName == nil || *first.Winner.PlayerName != "Ron Lengling" {
		t.Errorf("expected winner player_name 'Ron Lengling', got %v", first.Winner.PlayerName)
	}
	if first.Winner.WpprPoints == nil || float32(*first.Winner.WpprPoints) != float32(8.66) {
		t.Errorf("expected winner wppr_points ~8.66, got %v", first.Winner.WpprPoints)
	}
	if first.Winner.CountryCd == nil || *first.Winner.CountryCd != "AU" {
		t.Errorf("expected winner country_cd 'AU', got %v", first.Winner.CountryCd)
	}
	if first.Winner.ExcludedFlag == nil || bool(*first.Winner.ExcludedFlag) != false {
		t.Errorf("expected winner excluded_flag false, got %v", first.Winner.ExcludedFlag)
	}
}

func TestTournamentSearchFutureMainAU(t *testing.T) {
	client, _ := newTestClient(t, sampleTournamentSearchFutureMainAU)
	ctx := context.Background()

	country := "AU"
	stateprov := "NSW"
	rankType := MAIN
	params := &TourSearchParams{
		Country:   &country,
		Stateprov: &stateprov,
		RankType:  &rankType,
	}

	resp, err := client.TourSearchWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	if resp.JSON200.TotalResults == nil || int(*resp.JSON200.TotalResults) != 22 {
		t.Errorf("expected total_results 22, got %v", resp.JSON200.TotalResults)
	}

	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}
	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 22 {
		t.Errorf("expected 22 tournaments, got %d", len(tournaments))
	}

	// verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 106130 {
		t.Errorf("expected first tournament_id 106130, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Central Coast Pinball @Lost Souls" {
		t.Errorf("expected first tournament_name 'Central Coast Pinball @Lost Souls', got %v", first.TournamentName)
	}
	if first.EventStartDate == nil || *first.EventStartDate != "2026-04-01" {
		t.Errorf("expected first event_start_date '2026-04-01', got %v", first.EventStartDate)
	}

	// future tournaments: player_count should be 0 and winner.player_id should be nil
	for i, tour := range tournaments {
		if tour.Winner != nil && tour.Winner.PlayerId != nil {
			t.Errorf("tournaments[%d]: expected winner.player_id to be nil for future tournament, got %v", i, tour.Winner.PlayerId)
		}
		if tour.Winner != nil && tour.Winner.WpprPoints != nil {
			t.Errorf("tournaments[%d]: expected winner.wppr_points to be nil for future tournament, got %v", i, tour.Winner.WpprPoints)
		}
	}
}

func TestTournamentSearchPastWomenUS(t *testing.T) {
	client, _ := newTestClient(t, sampleTournamentSearchPastWomenUS)
	ctx := context.Background()

	country := "US"
	rankType := WOMEN
	params := &TourSearchParams{
		Country:  &country,
		RankType: &rankType,
	}

	resp, err := client.TourSearchWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	// total_results reflects server total (133), not number in response array (50)
	if resp.JSON200.TotalResults == nil || int(*resp.JSON200.TotalResults) != 133 {
		t.Errorf("expected total_results 133, got %v", resp.JSON200.TotalResults)
	}

	if resp.JSON200.SearchFilter == nil {
		t.Fatal("expected search_filter to be non-nil")
	}
	sf := resp.JSON200.SearchFilter
	if sf.Country == nil || *sf.Country != "US" {
		t.Errorf("expected search_filter.country 'US', got %v", sf.Country)
	}
	if sf.RankType == nil || *sf.RankType != "WOMEN" {
		t.Errorf("expected search_filter.rank_type 'WOMEN', got %v", sf.RankType)
	}

	if resp.JSON200.Tournaments == nil {
		t.Fatal("expected tournaments to be non-nil")
	}
	tournaments := *resp.JSON200.Tournaments
	if len(tournaments) != 50 {
		t.Errorf("expected 50 tournaments in response, got %d", len(tournaments))
	}

	// verify first tournament
	first := tournaments[0]
	if first.TournamentId == nil || int(*first.TournamentId) != 97228 {
		t.Errorf("expected first tournament_id 97228, got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Belles & Chimes @ Player 1UP" {
		t.Errorf("expected first tournament_name 'Belles & Chimes @ Player 1UP', got %v", first.TournamentName)
	}
	if first.RankingSystem == nil || *first.RankingSystem != "WOMEN" {
		t.Errorf("expected first ranking_system 'WOMEN', got %v", first.RankingSystem)
	}
	if first.CountryCode == nil || *first.CountryCode != "US" {
		t.Errorf("expected first country_code 'US', got %v", first.CountryCode)
	}
	if first.PlayerCount == nil || int(*first.PlayerCount) != 9 {
		t.Errorf("expected first player_count 9, got %v", first.PlayerCount)
	}

	// verify winner
	if first.Winner == nil {
		t.Fatal("expected winner to be non-nil for past tournament")
	}
	if first.Winner.PlayerId == nil || int(*first.Winner.PlayerId) != 96033 {
		t.Errorf("expected winner player_id 96033, got %v", first.Winner.PlayerId)
	}
	if first.Winner.PlayerName == nil || *first.Winner.PlayerName != "Kelly Walker" {
		t.Errorf("expected winner player_name 'Kelly Walker', got %v", first.Winner.PlayerName)
	}
	if first.Winner.CountryCd == nil || *first.Winner.CountryCd != "US" {
		t.Errorf("expected winner country_cd 'US', got %v", first.Winner.CountryCd)
	}

	// all tournaments should have ranking_system WOMEN
	for i, tour := range tournaments {
		if tour.RankingSystem == nil || *tour.RankingSystem != "WOMEN" {
			t.Errorf("tournaments[%d]: expected ranking_system 'WOMEN', got %v", i, tour.RankingSystem)
		}
	}
}

func TestTournamentSearchRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentSearchPastMainAU)
	ctx := context.Background()

	country := "AU"
	stateprov := "NSW"
	startDate := "2026-01-01"
	endDate := "2026-02-01"
	rankType := MAIN
	params := &TourSearchParams{
		Country:   &country,
		Stateprov: &stateprov,
		StartDate: &startDate,
		EndDate:   &endDate,
		RankType:  &rankType,
	}

	_, err := client.TourSearchWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/search" {
		t.Errorf("expected path '/tournament/search', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("country") != "AU" {
		t.Errorf("expected query param country=AU, got %s", req.URL.Query().Get("country"))
	}
	if req.URL.Query().Get("stateprov") != "NSW" {
		t.Errorf("expected query param stateprov=NSW, got %s", req.URL.Query().Get("stateprov"))
	}
	if req.URL.Query().Get("rank_type") != "MAIN" {
		t.Errorf("expected query param rank_type=MAIN, got %s", req.URL.Query().Get("rank_type"))
	}
	if req.URL.Query().Get("start_date") != "2026-01-01" {
		t.Errorf("expected query param start_date=2026-01-01, got %s", req.URL.Query().Get("start_date"))
	}
	if req.URL.Query().Get("end_date") != "2026-02-01" {
		t.Errorf("expected query param end_date=2026-02-01, got %s", req.URL.Query().Get("end_date"))
	}
}
