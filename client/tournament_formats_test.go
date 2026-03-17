package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/tournament_formats.json
var sampleTournamentFormatsResponse string

func TestTourFormatsWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentFormatsResponse)
	ctx := context.Background()
	resp, err := client.TourFormatsWithResponse(ctx)
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

	// Verify qualifying formats
	if resp.JSON200.QualifyingFormats == nil {
		t.Fatal("expected QualifyingFormats to be non-nil")
	}
	qualifyingFormats := *resp.JSON200.QualifyingFormats
	if len(qualifyingFormats) != 11 {
		t.Errorf("expected 11 qualifying formats, got %d", len(qualifyingFormats))
	}

	// Verify first qualifying format (Best Game, id=4)
	first := qualifyingFormats[0]
	if first.FormatId == nil || int(*first.FormatId) != 4 {
		t.Errorf("expected first qualifying format_id 4, got %v", first.FormatId)
	}
	if first.Name == nil || *first.Name != "Best Game" {
		t.Errorf("expected first qualifying name 'Best Game', got %v", first.Name)
	}

	// Spot-check Strike Knockout in qualifying formats
	found := false
	for _, f := range qualifyingFormats {
		if f.FormatId != nil && int(*f.FormatId) == 10 && f.Name != nil && *f.Name == "Strike Knockout" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find 'Strike Knockout' (id=10) in qualifying formats")
	}

	// Verify finals formats
	if resp.JSON200.FinalsFormats == nil {
		t.Fatal("expected FinalsFormats to be non-nil")
	}
	finalsFormats := *resp.JSON200.FinalsFormats
	if len(finalsFormats) != 14 {
		t.Errorf("expected 14 finals formats, got %d", len(finalsFormats))
	}

	// Verify first finals format (Amazing Race, id=28)
	firstFinal := finalsFormats[0]
	if firstFinal.FormatId == nil || int(*firstFinal.FormatId) != 28 {
		t.Errorf("expected first finals format_id 28, got %v", firstFinal.FormatId)
	}
	if firstFinal.Name == nil || *firstFinal.Name != "Amazing Race" {
		t.Errorf("expected first finals name 'Amazing Race', got %v", firstFinal.Name)
	}

	// Spot-check Single Elimination in finals formats
	found = false
	for _, f := range finalsFormats {
		if f.FormatId != nil && int(*f.FormatId) == 2 && f.Name != nil && *f.Name == "Single Elimination" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find 'Single Elimination' (id=2) in finals formats")
	}

	// Verify all formats have required fields
	for i, f := range qualifyingFormats {
		if f.FormatId == nil {
			t.Errorf("qualifying_formats[%d]: expected format_id to be non-nil", i)
		}
		if f.Name == nil || *f.Name == "" {
			t.Errorf("qualifying_formats[%d]: expected name to be non-empty", i)
		}
	}
	for i, f := range finalsFormats {
		if f.FormatId == nil {
			t.Errorf("finals_formats[%d]: expected format_id to be non-nil", i)
		}
		if f.Name == nil || *f.Name == "" {
			t.Errorf("finals_formats[%d]: expected name to be non-empty", i)
		}
	}
}

func TestTourFormatsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleTournamentFormatsResponse)
	ctx := context.Background()
	_, err := client.TourFormatsWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/tournament/formats" {
		t.Errorf("expected path '/tournament/formats', got %s", req.URL.Path)
	}
}
