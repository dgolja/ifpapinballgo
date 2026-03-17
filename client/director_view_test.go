package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/director_id_13.json
var sampleDirectorResponse string

//go:embed testdata/director_id_3478.json
var sampleDirectorResponse3478 string

func TestDirectorViewDirectorWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorResponse)
	ctx := context.Background()
	resp, err := client.ViewDirectorWithResponse(ctx, float32(13))
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

	d := resp.JSON200

	if d.DirectorId == nil || int(*d.DirectorId) != 13 {
		t.Errorf("expected director_id 13, got %v", d.DirectorId)
	}
	if d.Name == nil || *d.Name != "Pete Hendricks" {
		t.Errorf("expected name 'Pete Hendricks', got %v", d.Name)
	}
	if d.DefaultTournamentPhoto == nil || *d.DefaultTournamentPhoto != "https://www.ifpapinball.com/images/ifpa_gray.jpg" {
		t.Errorf("expected default_tournament_photo URL, got %v", d.DefaultTournamentPhoto)
	}

	if d.ProfilePhoto == nil || *d.ProfilePhoto != "" {
		t.Errorf("expected profile_photo to be empty string, got %v", d.ProfilePhoto)
	}

	// null fields
	if d.PlayerId != nil {
		t.Errorf("expected player_id to be nil, got %v", d.PlayerId)
	}
	if d.City != nil {
		t.Errorf("expected city to be nil, got %v", d.City)
	}
	if d.CountryCode != nil {
		t.Errorf("expected country_code to be nil, got %v", d.CountryCode)
	}
	if d.CountryName != nil {
		t.Errorf("expected country_name to be nil, got %v", d.CountryName)
	}
	if d.CountryId != nil {
		t.Errorf("expected country_id to be nil, got %v", d.CountryId)
	}
	if d.Stateprov != nil {
		t.Errorf("expected stateprov to be nil, got %v", d.Stateprov)
	}
	if d.TwitchUsername != nil {
		t.Errorf("expected twitch_username to be nil, got %v", d.TwitchUsername)
	}

	// stats
	if d.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}
	s := d.Stats
	if s.TournamentCount == nil || int(*s.TournamentCount) != 146 {
		t.Errorf("expected tournament_count 146, got %v", s.TournamentCount)
	}
	if s.UniqueLocationCount == nil || int(*s.UniqueLocationCount) != 17 {
		t.Errorf("expected unique_location_count 17, got %v", s.UniqueLocationCount)
	}
	if s.WomenTournamentCount == nil || int(*s.WomenTournamentCount) != 14 {
		t.Errorf("expected women_tournament_count 14, got %v", s.WomenTournamentCount)
	}
	if s.LeagueCount == nil || int(*s.LeagueCount) != 75 {
		t.Errorf("expected league_count 75, got %v", s.LeagueCount)
	}
	if s.TotalPlayerCount == nil || int(*s.TotalPlayerCount) != 5542 {
		t.Errorf("expected total_player_count 5542, got %v", s.TotalPlayerCount)
	}
	if s.UniquePlayerCount == nil || int(*s.UniquePlayerCount) != 943 {
		t.Errorf("expected unique_player_count 943, got %v", s.UniquePlayerCount)
	}
	if s.HighestValue == nil || float32(*s.HighestValue) != 61.29 {
		t.Errorf("expected highest_value 61.29, got %v", s.HighestValue)
	}
	if s.AverageValue == nil || float32(*s.AverageValue) != 14.85 {
		t.Errorf("expected average_value 14.85, got %v", s.AverageValue)
	}
	if s.FirstTimePlayerCount == nil || int(*s.FirstTimePlayerCount) != 466 {
		t.Errorf("expected first_time_player_count 466, got %v", s.FirstTimePlayerCount)
	}
	if s.RepeatPlayerCount == nil || int(*s.RepeatPlayerCount) != 299 {
		t.Errorf("expected repeat_player_count 299, got %v", s.RepeatPlayerCount)
	}
	if s.LargestEventCount == nil || int(*s.LargestEventCount) != 177 {
		t.Errorf("expected largest_event_count 177, got %v", s.LargestEventCount)
	}
	if s.SingleFormatCount == nil || int(*s.SingleFormatCount) != 17 {
		t.Errorf("expected single_format_count 17, got %v", s.SingleFormatCount)
	}
	if s.MultipleFormatCount == nil || int(*s.MultipleFormatCount) != 88 {
		t.Errorf("expected multiple_format_count 88, got %v", s.MultipleFormatCount)
	}
	if s.UnknownFormatCount == nil || int(*s.UnknownFormatCount) != 45 {
		t.Errorf("expected unknown_format_count 45, got %v", s.UnknownFormatCount)
	}

	// formats
	if s.Formats == nil {
		t.Fatal("expected formats to be non-nil")
	}
	if len(*s.Formats) != 4 {
		t.Errorf("expected 4 formats, got %d", len(*s.Formats))
	} else {
		first := (*s.Formats)[0]
		if first.Name == nil || *first.Name != "Match Play" {
			t.Errorf("expected first format name 'Match Play', got %v", first.Name)
		}
		if first.Count == nil || int(*first.Count) != 12 {
			t.Errorf("expected first format count 12, got %v", first.Count)
		}
	}
}

func TestDirectorViewDirectorWithResponse3478(t *testing.T) {
	client, _ := newTestClient(t, sampleDirectorResponse3478)
	ctx := context.Background()
	resp, err := client.ViewDirectorWithResponse(ctx, float32(3478))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode())
	}

	if resp.JSON200 == nil {
		t.Fatal("expected JSON200 to be non-nil")
	}

	d := resp.JSON200

	if d.DirectorId == nil || int(*d.DirectorId) != 3478 {
		t.Errorf("expected director_id 3478, got %v", d.DirectorId)
	}
	if d.Name == nil || *d.Name != "Andrew Gliatis" {
		t.Errorf("expected name 'Andrew Gliatis', got %v", d.Name)
	}
	if d.ProfilePhoto == nil || *d.ProfilePhoto != "https://www.ifpapinball.com/images/profiles/director/3478.jpg" {
		t.Errorf("expected profile_photo URL, got %v", d.ProfilePhoto)
	}
	if d.DefaultTournamentPhoto == nil || *d.DefaultTournamentPhoto != "https://www.ifpapinball.com/images/ifpa_gray.jpg" {
		t.Errorf("expected default_tournament_photo URL, got %v", d.DefaultTournamentPhoto)
	}

	// player_id is present in this fixture
	if d.PlayerId == nil || int(*d.PlayerId) != 90228 {
		t.Errorf("expected player_id 90228, got %v", d.PlayerId)
	}

	// empty-string fields (non-null but empty)
	if d.City == nil || *d.City != "" {
		t.Errorf("expected city to be empty string, got %v", d.City)
	}
	if d.Stateprov == nil || *d.Stateprov != "" {
		t.Errorf("expected stateprov to be empty string, got %v", d.Stateprov)
	}
	if d.TwitchUsername == nil || *d.TwitchUsername != "" {
		t.Errorf("expected twitch_username to be empty string, got %v", d.TwitchUsername)
	}

	// country fields are populated
	if d.CountryName == nil || *d.CountryName != "Australia" {
		t.Errorf("expected country_name 'Australia', got %v", d.CountryName)
	}
	if d.CountryCode == nil || *d.CountryCode != "AU" {
		t.Errorf("expected country_code 'AU', got %v", d.CountryCode)
	}
	if d.CountryId == nil || int(*d.CountryId) != 7 {
		t.Errorf("expected country_id 7, got %v", d.CountryId)
	}

	// stats
	if d.Stats == nil {
		t.Fatal("expected stats to be non-nil")
	}
	s := d.Stats
	if s.TournamentCount == nil || int(*s.TournamentCount) != 52 {
		t.Errorf("expected tournament_count 52, got %v", s.TournamentCount)
	}
	if s.UniqueLocationCount == nil || int(*s.UniqueLocationCount) != 2 {
		t.Errorf("expected unique_location_count 2, got %v", s.UniqueLocationCount)
	}
	if s.WomenTournamentCount == nil || int(*s.WomenTournamentCount) != 0 {
		t.Errorf("expected women_tournament_count 0, got %v", s.WomenTournamentCount)
	}
	if s.LeagueCount == nil || int(*s.LeagueCount) != 0 {
		t.Errorf("expected league_count 0, got %v", s.LeagueCount)
	}
	if s.HighestValue == nil || float32(*s.HighestValue) != 36.92 {
		t.Errorf("expected highest_value 36.92, got %v", s.HighestValue)
	}
	if s.AverageValue == nil || float32(*s.AverageValue) != 13.30 {
		t.Errorf("expected average_value 13.30, got %v", s.AverageValue)
	}
	if s.TotalPlayerCount == nil || int(*s.TotalPlayerCount) != 1200 {
		t.Errorf("expected total_player_count 1200, got %v", s.TotalPlayerCount)
	}
	if s.UniquePlayerCount == nil || int(*s.UniquePlayerCount) != 152 {
		t.Errorf("expected unique_player_count 152, got %v", s.UniquePlayerCount)
	}
	if s.FirstTimePlayerCount == nil || int(*s.FirstTimePlayerCount) != 56 {
		t.Errorf("expected first_time_player_count 56, got %v", s.FirstTimePlayerCount)
	}
	if s.RepeatPlayerCount == nil || int(*s.RepeatPlayerCount) != 35 {
		t.Errorf("expected repeat_player_count 35, got %v", s.RepeatPlayerCount)
	}
	if s.LargestEventCount == nil || int(*s.LargestEventCount) != 44 {
		t.Errorf("expected largest_event_count 44, got %v", s.LargestEventCount)
	}
	if s.SingleFormatCount == nil || int(*s.SingleFormatCount) != 1 {
		t.Errorf("expected single_format_count 1, got %v", s.SingleFormatCount)
	}
	if s.MultipleFormatCount == nil || int(*s.MultipleFormatCount) != 100 {
		t.Errorf("expected multiple_format_count 100, got %v", s.MultipleFormatCount)
	}
	if s.UnknownFormatCount == nil || int(*s.UnknownFormatCount) != 0 {
		t.Errorf("expected unknown_format_count 0, got %v", s.UnknownFormatCount)
	}

	// formats
	if s.Formats == nil {
		t.Fatal("expected formats to be non-nil")
	}
	if len(*s.Formats) != 1 {
		t.Errorf("expected 1 format, got %d", len(*s.Formats))
	} else {
		first := (*s.Formats)[0]
		if first.Name == nil || *first.Name != "Single Elimination" {
			t.Errorf("expected first format name 'Single Elimination', got %v", first.Name)
		}
		if first.Count == nil || int(*first.Count) != 1 {
			t.Errorf("expected first format count 1, got %v", first.Count)
		}
	}
}

func TestDirectorViewDirectorRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorResponse)
	ctx := context.Background()
	_, err := client.ViewDirectorWithResponse(ctx, float32(13))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/director/13" {
		t.Errorf("expected path '/director/13', got %s", req.URL.Path)
	}
}
