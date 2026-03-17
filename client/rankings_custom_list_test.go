package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/rankings_custom_list.json
var sampleRankingsCustomListResponse string

func TestRankingsCustomList(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCustomListResponse)
	ctx := context.Background()
	resp, err := client.RankingCustomListWithResponse(ctx)
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

	if resp.JSON200.TotalCount == nil || int(*resp.JSON200.TotalCount) != 78 {
		t.Errorf("expected total_count 78, got %v", resp.JSON200.TotalCount)
	}

	if resp.JSON200.CustomView == nil {
		t.Fatal("expected custom_view to be non-nil")
	}

	views := *resp.JSON200.CustomView
	if len(views) != 78 {
		t.Errorf("expected 78 custom_view entries, got %d", len(views))
	}

	// Verify first entry - Lund Pinball Academy Rankings
	first := views[0]
	if first.ViewId == nil || int(*first.ViewId) != 14 {
		t.Errorf("expected first view_id 14, got %v", first.ViewId)
	}
	if first.Title == nil || *first.Title != "Lund Pinball Academy Rankings" {
		t.Errorf("expected first title 'Lund Pinball Academy Rankings', got %v", first.Title)
	}
	if first.Description == nil || *first.Description != "WPPR points earned at all Lund Pinball Academy events" {
		t.Errorf("expected first description 'WPPR points earned at all Lund Pinball Academy events', got %v", first.Description)
	}

	// Verify all entries have required fields
	for i, v := range views {
		if v.ViewId == nil {
			t.Errorf("views[%d]: expected view_id to be non-nil", i)
		}
		if v.Title == nil || *v.Title == "" {
			t.Errorf("views[%d]: expected title to be non-empty", i)
		}
	}
}

func TestRankingsCustomListRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleRankingsCustomListResponse)
	ctx := context.Background()
	_, err := client.RankingCustomListWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/rankings/custom/list" {
		t.Errorf("expected path '/rankings/custom/list', got %s", req.URL.Path)
	}
}
