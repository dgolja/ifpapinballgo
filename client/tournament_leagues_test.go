package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_leagues_active.json
var sampleTournamentLeaguesActiveResponse string

//go:embed testdata/tournament_leagues_upcoming.json
var sampleTournamentLeaguesUpcomingResponse string

func TestTournamentLeaguesActiveWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentLeaguesActiveResponse)
	ctx := context.Background()
	resp, err := client.ViewLeagueInfoWithResponse(ctx, ViewLeagueInfoParamsTimePeriodACTIVE)
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

	if resp.JSON200.TotalEntries == nil {
		t.Fatal("expected TotalEntries to be non-nil")
	}
	if int(*resp.JSON200.TotalEntries) != 228 {
		t.Errorf("expected TotalEntries 228, got %d", int(*resp.JSON200.TotalEntries))
	}

	if resp.JSON200.LeagueStatus == nil {
		t.Fatal("expected LeagueStatus to be non-nil")
	}
	if *resp.JSON200.LeagueStatus != "active" {
		t.Errorf("expected LeagueStatus 'active', got '%s'", *resp.JSON200.LeagueStatus)
	}

	if resp.JSON200.Leagues == nil {
		t.Fatal("expected Leagues to be non-nil")
	}
	leagues := *resp.JSON200.Leagues
	if len(leagues) != 227 {
		t.Errorf("expected 227 leagues, got %d", len(leagues))
	}

	// Verify first league entry
	first := leagues[0]
	if first.TournamentId == nil || *first.TournamentId != "100860" {
		t.Errorf("expected first TournamentId '100860', got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "Flipperliga Regensburg" {
		t.Errorf("expected first TournamentName 'Flipperliga Regensburg', got %v", first.TournamentName)
	}
	if first.City == nil || *first.City != "Regensburg" {
		t.Errorf("expected first City 'Regensburg', got %v", first.City)
	}
	if first.CountryCode == nil || *first.CountryCode != "DE" {
		t.Errorf("expected first CountryCode 'DE', got %v", first.CountryCode)
	}
	if first.DirectorId == nil || int(*first.DirectorId) != 2431 {
		t.Errorf("expected first DirectorId 2431, got %v", first.DirectorId)
	}
	if first.PrivateFlag == nil || bool(*first.PrivateFlag) != true {
		t.Errorf("expected first PrivateFlag true, got %v", first.PrivateFlag)
	}
	if first.RankingSystem == nil || *first.RankingSystem != "MAIN" {
		t.Errorf("expected first RankingSystem 'MAIN', got %v", first.RankingSystem)
	}
}

func TestTournamentLeaguesUpcomingWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentLeaguesUpcomingResponse)
	ctx := context.Background()
	resp, err := client.ViewLeagueInfoWithResponse(ctx, ViewLeagueInfoParamsTimePeriodUPCOMING)
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

	if resp.JSON200.TotalEntries == nil {
		t.Fatal("expected TotalEntries to be non-nil")
	}
	if int(*resp.JSON200.TotalEntries) != 249 {
		t.Errorf("expected TotalEntries 249, got %d", int(*resp.JSON200.TotalEntries))
	}

	if resp.JSON200.LeagueStatus == nil {
		t.Fatal("expected LeagueStatus to be non-nil")
	}
	if *resp.JSON200.LeagueStatus != "upcoming" {
		t.Errorf("expected LeagueStatus 'upcoming', got '%s'", *resp.JSON200.LeagueStatus)
	}

	if resp.JSON200.Leagues == nil {
		t.Fatal("expected Leagues to be non-nil")
	}
	leagues := *resp.JSON200.Leagues
	if len(leagues) != 248 {
		t.Errorf("expected 248 leagues, got %d", len(leagues))
	}

	// Verify first league entry
	first := leagues[0]
	if first.TournamentId == nil || *first.TournamentId != "110460" {
		t.Errorf("expected first TournamentId '110460', got %v", first.TournamentId)
	}
	if first.TournamentName == nil || *first.TournamentName != "The Pinball Room - League" {
		t.Errorf("expected first TournamentName 'The Pinball Room - League', got %v", first.TournamentName)
	}
	if first.DirectorId == nil || int(*first.DirectorId) != 3707 {
		t.Errorf("expected first DirectorId 3707, got %v", first.DirectorId)
	}
	if first.PrivateFlag == nil || bool(*first.PrivateFlag) != true {
		t.Errorf("expected first PrivateFlag true, got %v", first.PrivateFlag)
	}
}

func TestTournamentLeaguesRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentLeaguesActiveResponse)
	ctx := context.Background()
	_, err := client.ViewLeagueInfoWithResponse(ctx, ViewLeagueInfoParamsTimePeriodACTIVE)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/leagues/ACTIVE" {
		t.Errorf("expected path '/tournament/leagues/ACTIVE', got %s", req.URL.Path)
	}
}

func TestTournamentLeaguesPrivateFlagFalse(t *testing.T) {
	client, _ := newTestClient(t, sampleTournamentLeaguesActiveResponse)
	ctx := context.Background()
	resp, err := client.ViewLeagueInfoWithResponse(ctx, ViewLeagueInfoParamsTimePeriodACTIVE)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.Leagues == nil {
		t.Fatal("expected JSON200 and Leagues to be non-nil")
	}

	// The 4th entry (index 3) has private_flag: "false"
	leagues := *resp.JSON200.Leagues
	fourth := leagues[3]
	if fourth.PrivateFlag == nil || bool(*fourth.PrivateFlag) != false {
		t.Errorf("expected fourth league PrivateFlag false, got %v", fourth.PrivateFlag)
	}
	if fourth.TournamentId == nil || *fourth.TournamentId != "99118" {
		t.Errorf("expected fourth TournamentId '99118', got %v", fourth.TournamentId)
	}
}

func TestTournamentLeaguesWomenRankingSystem(t *testing.T) {
	client, _ := newTestClient(t, sampleTournamentLeaguesActiveResponse)
	ctx := context.Background()
	resp, err := client.ViewLeagueInfoWithResponse(ctx, ViewLeagueInfoParamsTimePeriodACTIVE)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.Leagues == nil {
		t.Fatal("expected JSON200 and Leagues to be non-nil")
	}

	// Find a league with WOMEN ranking system
	leagues := *resp.JSON200.Leagues
	found := false
	for _, league := range leagues {
		if league.RankingSystem != nil && *league.RankingSystem == "WOMEN" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find at least one league with WOMEN ranking system")
	}
}
