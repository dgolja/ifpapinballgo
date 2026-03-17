package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/series_regions.json
var sampleSeriesRegionsResponse string

func TestSeriesRegions(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesRegionsResponse)
	ctx := context.Background()
	resp, err := client.SeriesRegionsWithResponse(ctx, ACS, &SeriesRegionsParams{})
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

	if resp.JSON200.Year == nil || int(*resp.JSON200.Year) != 2025 {
		t.Errorf("expected year 2025, got %v", resp.JSON200.Year)
	}

	if resp.JSON200.ActiveRegions == nil {
		t.Fatal("expected active_regions to be non-nil")
	}

	regions := *resp.JSON200.ActiveRegions
	if len(regions) != 7 {
		t.Errorf("expected 7 regions, got %d", len(regions))
	}

	// Verify first entry
	first := regions[0]
	if first.RegionName == nil || *first.RegionName != "Australian Capital Territory" {
		t.Errorf("expected first region_name 'Australian Capital Territory', got %v", first.RegionName)
	}
	if first.RegionCode == nil || *first.RegionCode != "ACT" {
		t.Errorf("expected first region_code 'ACT', got %v", first.RegionCode)
	}

	// Verify all entries have required fields
	for i, r := range regions {
		if r.RegionName == nil || *r.RegionName == "" {
			t.Errorf("regions[%d]: expected region_name to be non-empty", i)
		}
		if r.RegionCode == nil || *r.RegionCode == "" {
			t.Errorf("regions[%d]: expected region_code to be non-empty", i)
		}
	}
}

func TestSeriesRegionsRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesRegionsResponse)
	ctx := context.Background()
	_, err := client.SeriesRegionsWithResponse(ctx, ACS, &SeriesRegionsParams{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/ACS/regions" {
		t.Errorf("expected path '/series/ACS/regions', got %s", req.URL.Path)
	}
}
