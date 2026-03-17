package client

import (
	"context"
	_ "embed"
	"net/http"
	"strings"
	"testing"
)

//go:embed testdata/director_search.json
var sampleDirectorSearchResponse string

func TestDirectorSearch(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorSearchResponse)
	ctx := context.Background()
	searchQuery := "Andrew"
	params := &SearchDirectorsParams{Name: &searchQuery}
	resp, err := client.SearchDirectorsWithResponse(ctx, params)
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

	if resp.JSON200.SearchTerm == nil || *resp.JSON200.SearchTerm != "Andrew" {
		t.Errorf("expected search_term 'Andrew', got %v", resp.JSON200.SearchTerm)
	}
	if *resp.JSON200.Count != 16 {
		t.Errorf("expected count 16, got %d", int(*resp.JSON200.Count))
	}

	if resp.JSON200.Directors == nil {
		t.Fatal("expected directors to be non-nil")
	}
	directors := *resp.JSON200.Directors
	if len(directors) != 16 {
		t.Errorf("expected 16 directors, got %d", len(directors))
	}

	// all names must start with "Andrew"
	for _, director := range directors {
		if director.Name == nil || !strings.HasPrefix(*director.Name, "Andrew") {
			t.Errorf("expected name to start with 'Andrew', got %v (director_id: %v)", director.Name, director.DirectorId)
		}
	}

	// verify first result
	first := directors[0]
	if first.DirectorId == nil || int(*first.DirectorId) != 3478 {
		t.Errorf("expected first director_id 3478, got %v", first.DirectorId)
	}
	if first.Name == nil || *first.Name != "Andrew Gliatis" {
		t.Errorf("expected first name 'Andrew Gliatis', got %v", first.Name)
	}
	if first.CountryCode == nil || *first.CountryCode != "AU" {
		t.Errorf("expected first country_code 'AU', got %v", first.CountryCode)
	}
	if first.CountryName == nil || *first.CountryName != "Australia" {
		t.Errorf("expected first country_name 'Australia', got %v", first.CountryName)
	}

	// verify first result stats
	if first.Stats == nil {
		t.Fatal("expected stats to be non-nil for first director")
	}
	if first.Stats.EventCount == nil || int(*first.Stats.EventCount) != 50 {
		t.Errorf("expected first stats.event_count 50, got %v", first.Stats.EventCount)
	}
	if first.Stats.UniquePlayerCount == nil || int(*first.Stats.UniquePlayerCount) != 149 {
		t.Errorf("expected first stats.unique_player_count 149, got %v", first.Stats.UniquePlayerCount)
	}
	if first.Stats.LastEventDate == nil || *first.Stats.LastEventDate != "2026-02-26" {
		t.Errorf("expected first stats.last_event_date '2026-02-26', got %v", first.Stats.LastEventDate)
	}
}

func TestDirectorSearchRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleDirectorSearchResponse)
	ctx := context.Background()
	searchQuery := "Andrew"
	params := &SearchDirectorsParams{Name: &searchQuery}
	_, err := client.SearchDirectorsWithResponse(ctx, params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/director/search" {
		t.Errorf("expected path '/director/search', got %s", req.URL.Path)
	}
	if req.URL.Query().Get("name") != "Andrew" {
		t.Errorf("expected query param name=Andrew, got %s", req.URL.Query().Get("name"))
	}
}
