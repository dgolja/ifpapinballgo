package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/series_list.json
var sampleSeriesListResponse string

func TestSeriesList(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesListResponse)
	ctx := context.Background()
	resp, err := client.SeriesListWithResponse(ctx)
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

	if resp.JSON200.Series == nil {
		t.Fatal("expected series to be non-nil")
	}

	series := *resp.JSON200.Series
	if len(series) != 4 {
		t.Errorf("expected 4 series, got %d", len(series))
	}

	// Verify first entry - NACS
	first := series[0]
	if first.Code == nil || *first.Code != "NACS" {
		t.Errorf("expected first code 'NACS', got %v", first.Code)
	}
	if first.Title == nil || *first.Title != "North American Championship Series" {
		t.Errorf("expected first title 'North American Championship Series', got %v", first.Title)
	}
	if first.Years == nil || len(*first.Years) != 5 {
		t.Errorf("expected 5 years for NACS, got %v", first.Years)
	}

	// Verify all entries have required fields and years
	for i, s := range series {
		if s.Code == nil || *s.Code == "" {
			t.Errorf("series[%d]: expected code to be non-empty", i)
		}
		if s.Title == nil || *s.Title == "" {
			t.Errorf("series[%d]: expected title to be non-empty", i)
		}
		if s.Years == nil || len(*s.Years) == 0 {
			t.Errorf("series[%d]: expected years to be non-empty", i)
		}
	}
}

func TestSeriesListRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleSeriesListResponse)
	ctx := context.Background()
	_, err := client.SeriesListWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/series/list" {
		t.Errorf("expected path '/series/list', got %s", req.URL.Path)
	}
}
