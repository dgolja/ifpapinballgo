package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/series_region_reps.json
var sampleSeriesRegionRepsResponse string

func TestSeriesRegionReps(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesRegionRepsResponse)
	ctx := context.Background()
	resp, err := client.SeriesRegionRepsWithResponse(ctx, "ACS")
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

	if resp.JSON200.SeriesCode == nil || *resp.JSON200.SeriesCode != "ACS" {
		t.Errorf("expected series_code 'ACS', got %v", resp.JSON200.SeriesCode)
	}

	if resp.JSON200.Representative == nil {
		t.Fatal("expected representative to be non-nil")
	}

	reps := *resp.JSON200.Representative
	if len(reps) != 8 {
		t.Errorf("expected 8 representatives, got %d", len(reps))
	}

	// Verify first entry
	first := reps[0]
	if first.PlayerKey == nil || int(*first.PlayerKey) != 55833 {
		t.Errorf("expected first player_key 55833, got %v", first.PlayerKey)
	}
	if first.PlayerId == nil || int(*first.PlayerId) != 55833 {
		t.Errorf("expected first player_id 55833, got %v", first.PlayerId)
	}
	if first.Name == nil || *first.Name != "James Todd AUS" {
		t.Errorf("expected first name 'James Todd AUS', got %v", first.Name)
	}
	if first.RegionCode == nil || *first.RegionCode != "ACT" {
		t.Errorf("expected first region_code 'ACT', got %v", first.RegionCode)
	}
	if first.RegionName == nil || *first.RegionName != "Australian Capital Territory" {
		t.Errorf("expected first region_name 'Australian Capital Territory', got %v", first.RegionName)
	}

	// Verify all entries have required fields
	for i, r := range reps {
		if r.PlayerKey == nil {
			t.Errorf("reps[%d]: expected player_key to be non-nil", i)
		}
		if r.PlayerId == nil {
			t.Errorf("reps[%d]: expected player_id to be non-nil", i)
		}
		if r.Name == nil || *r.Name == "" {
			t.Errorf("reps[%d]: expected name to be non-empty", i)
		}
		if r.RegionCode == nil || *r.RegionCode == "" {
			t.Errorf("reps[%d]: expected region_code to be non-empty", i)
		}
		if r.RegionName == nil || *r.RegionName == "" {
			t.Errorf("reps[%d]: expected region_name to be non-empty", i)
		}
	}
}

func TestSeriesRegionRepsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesRegionRepsResponse)
	ctx := context.Background()
	_, err := client.SeriesRegionRepsWithResponse(ctx, "ACS")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/region_reps" {
		t.Errorf("expected path '/series/ACS/region_reps', got %s", req.URL.Path)
	}
}
