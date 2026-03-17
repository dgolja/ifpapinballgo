package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"

	"github.com/dgolja/ifpapinballgo/types"
)

//go:embed testdata/player_multi.json
var samplePlayerResponse string

func TestPlayerViewPlayerMultiWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, samplePlayerResponse)
	ctx := context.Background()
	resp, err := client.ViewPlayerMultiWithResponse(ctx, &ViewPlayerMultiParams{Players: "13,44334,83231"})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}

	if len(*resp.JSON200.Player) != 3 {
		t.Errorf("expected 3 players, got %d", len(*resp.JSON200.Player))
	}

	if int(*(*resp.JSON200.Player)[0].PlayerId) != 13 {
		t.Errorf("expected first player ID '13', got '%s'", (*resp.JSON200.Player)[0].PlayerId)
	}

	firstPlayer := (*resp.JSON200.Player)[0]

	if firstPlayer.PlayerStats.System.Open == nil {
		t.Errorf("expected player ID '%d', to have PlayerStats.System.Open data", firstPlayer.PlayerId)
	} else {
		if firstPlayer.PlayerStats.System.Open.HighestRankDate == nil {
			t.Errorf("expected highest_rank_date to be non-nil for first player")
		}
		if firstPlayer.PlayerStats.System.Open.ProRank == nil {
			t.Errorf("expected pro_rank to be non-nil for first player")
		}
	}

	if firstPlayer.VirtualPlayerFlag == nil {
		t.Errorf("expected virtual_player_flag to be non-nil for first player")
	} else if bool(*firstPlayer.VirtualPlayerFlag) {
		t.Errorf("expected virtual_player_flag to be false for first player")
	}

	if firstPlayer.PlayerStats.YearsActive == nil {
		t.Errorf("expected years_active to be non-nil for first player")
	}

	if int(*(*resp.JSON200.Player)[1].PlayerId) != 44334 {
		t.Errorf("expected second player ID '44334', got '%s'", (*resp.JSON200.Player)[1].PlayerId)
	}

	if int(*(*resp.JSON200.Player)[2].PlayerId) != 83231 {
		t.Errorf("expected third player ID '83231', got '%s'", (*resp.JSON200.Player)[2].PlayerId)
	}
}

//go:embed testdata/player_data.json
var samplePlayerIDResponse string

//go:embed testdata/player_data_woman.json
var samplePlayerIDWomanResponse string

func TestPlayerViewPlayerWithResponse(t *testing.T) {
	tests := []struct {
		name               string
		viewPlayerResponse string
		playerId           int
		woman              bool
	}{
		{
			name:               "Basic ViewPlayerResponse",
			viewPlayerResponse: samplePlayerIDResponse,
			playerId:           44334,
		},
		{
			name:               "Woman ViewPlayerResponse",
			viewPlayerResponse: samplePlayerIDWomanResponse,
			playerId:           83231,
			woman:              true,
		},
	}

	for _, tt := range tests {
		mockClient := &MockHTTPClient{Response: createMockResponse(200, tt.viewPlayerResponse)}
		client, err := NewClientWithResponses("https://api.ifpapinball.com", WithHTTPClient(mockClient))
		if err != nil {
			t.Fatalf("failed to create client: %v", err)
		}
		ctx := context.Background()
		resp, err := client.ViewPlayerWithResponse(ctx, float32(tt.playerId))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if resp.StatusCode() != http.StatusOK {
			t.Errorf("expected status code 200, got %d", resp.StatusCode())
		}

		// Verify the request was made
		if len(mockClient.Requests) != 1 {
			t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
		}

		if resp.JSON200 == nil {
			t.Error("expected JSON200 to be non-nil")
		}

		if tt.playerId != int(*(*resp.JSON200.Player)[0].PlayerId) {
			t.Errorf("expected third player ID '%d', got '%s'", tt.playerId, (*resp.JSON200.Player)[0].PlayerId)
		}

		player := (*resp.JSON200.Player)[0]

		if player.PlayerStats.System.Open == nil {
			t.Errorf("expected player ID '%d', to have PlayerStats.System.Open data", tt.playerId)
		} else {
			if player.PlayerStats.System.Open.HighestRankDate == nil {
				t.Errorf("expected highest_rank_date to be non-nil for player %d", tt.playerId)
			}
			if player.PlayerStats.System.Open.ProRank == nil {
				t.Errorf("expected pro_rank to be non-nil for player %d", tt.playerId)
			}
		}

		if player.VirtualPlayerFlag == nil {
			t.Errorf("expected virtual_player_flag to be non-nil for player %d", tt.playerId)
		} else if bool(*player.VirtualPlayerFlag) {
			t.Errorf("expected virtual_player_flag to be false for player %d", tt.playerId)
		}

		if player.PlayerStats.YearsActive == nil {
			t.Errorf("expected years_active to be non-nil for player %d", tt.playerId)
		}

		if tt.woman {
			if player.PlayerStats.System.Womens == nil {
				t.Errorf("expected player ID '%d', to have PlayerStats.System.Womens data", tt.playerId)
			} else {
				if player.PlayerStats.System.Womens.HighestRankDate == nil {
					t.Errorf("expected womens highest_rank_date to be non-nil for player %d", tt.playerId)
				}
				if player.PlayerStats.System.Womens.ProRank == nil {
					t.Errorf("expected womens pro_rank to be non-nil for player %d", tt.playerId)
				}
			}
		}
	}
}

//go:embed testdata/player_pvp.json
var sampleViewPlayerPVPWithResponse string

func TestPlayerViewPlayerPVP(t *testing.T) {
	client, mockClient := newTestClient(t, sampleViewPlayerPVPWithResponse)
	ctx := context.Background()
	resp, err := client.ViewPlayerPVPWithResponse(ctx, 83231)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}

	if *resp.JSON200.PlayerId != 83231 {
		t.Error("expected PlayerId to be 8321")
	}
	if *resp.JSON200.TotalCompetitors != 5 {
		t.Error("expected TotalCompetitors to be 5")
	}
}

//go:embed testdata/player_pvp_to_player.json
var sampleViewPlayerPVPToPlayer string

func TestPlayerViewPlayerPVPToPlayer(t *testing.T) {
	client, mockClient := newTestClient(t, sampleViewPlayerPVPToPlayer)
	ctx := context.Background()
	resp, err := client.ViewPlayerPVPToPlayerWithResponse(ctx, 19853, 16976)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}

	if int(*(resp.JSON200.Player1).PlayerId) != 19853 {
		t.Error("expected PlayerId to be 19853")
	}

	if int(*(resp.JSON200.Player2).PlayerId) != 16976 {
		t.Error("expected PlayerId to be 16976")
	}

	if resp.JSON200.Pvp == nil || len(*resp.JSON200.Pvp) == 0 {
		t.Fatal("expected pvp to be non-nil and non-empty")
	}
	firstPvp := (*resp.JSON200.Pvp)[0]
	if firstPvp.EventEndDt == nil || *firstPvp.EventEndDt != "2015-11-15" {
		t.Errorf("expected event_end_dt '2015-11-15', got %v", firstPvp.EventEndDt)
	}
}

//go:embed testdata/player_rank_history.json
var sampleViewPlayerRankHistory string

func TestPlayerViewPlayerRankHistory(t *testing.T) {
	client, mockClient := newTestClient(t, sampleViewPlayerRankHistory)
	ctx := context.Background()
	resp, err := client.ViewPlayerRankHistoryWithResponse(ctx, 44434)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}

	if bool(*resp.JSON200.ActiveFlag) {
		t.Error("expected ActiveFlag to be false")
	}
}

// viewPlayerActiveResults
//
//go:embed testdata/player_active_results.json
var sampleViewPlayerActive string

func TestPlayerViewPlayerActiveResults(t *testing.T) {
	client, mockClient := newTestClient(t, sampleViewPlayerActive)
	ctx := context.Background()
	resp, err := client.ViewPlayerActiveResultsWithResponse(ctx, 44334, ViewPlayerActiveResultsParamsRankingSystem("MAIN"), ViewPlayerActiveResultsParamsType("ACTIVE"))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}

	if int(*(resp.JSON200.PlayerId)) != 44334 {
		t.Error("expected PlayerId to be 44334")
	}
	if *resp.JSON200.ResultsType != "active" {
		t.Error("expected ResultsType to be active")
	}

	if len(*resp.JSON200.Results) != int(*resp.JSON200.ResultsCount) {
		t.Errorf("expected Results array of lenght %d to match result count %d", len(*resp.JSON200.Results), int(*resp.JSON200.ResultsCount))
	}

	if *resp.JSON200.RankType != "MAIN" {
		t.Errorf("expected RankType to be 'MAIN', got '%s'", *resp.JSON200.RankType)
	}
}

//go:embed testdata/player_search.json
var sampleSearchPlayers string

func TestPlayerSearchPlayers(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSearchPlayers)
	ctx := context.Background()
	resp, err := client.SearchPlayersWithResponse(ctx, &SearchPlayersParams{Name: types.Ptr("Golja")})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	// Verify the request was made
	if len(mockClient.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(mockClient.Requests))
	}

	if resp.JSON200 == nil {
		t.Error("expected JSON200 to be non-nil")
	}
}
