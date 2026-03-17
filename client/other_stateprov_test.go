package client

import (
	"context"
	_ "embed"
	"net/http"
	"testing"
)

//go:embed testdata/other_stateprov.json
var sampleOtherStateProvResponse string

func TestOtherStateProvWithResponse(t *testing.T) {
	client, mockClient := newTestClient(t, sampleOtherStateProvResponse)
	ctx := context.Background()
	resp, err := client.OtherStateProvWithResponse(ctx)
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
	if resp.JSON200.Stateprov == nil {
		t.Fatal("expected Stateprov to be non-nil")
	}

	stateprovs := *resp.JSON200.Stateprov
	if len(stateprovs) != 3 {
		t.Errorf("expected 3 countries, got %d", len(stateprovs))
	}

	// Verify Australia (first entry)
	au := stateprovs[0]
	if au.CountryId == nil || int(*au.CountryId) != 7 {
		t.Errorf("expected Australia country_id 7, got %v", au.CountryId)
	}
	if au.CountryName == nil || *au.CountryName != "Australia" {
		t.Errorf("expected country_name 'Australia', got %v", au.CountryName)
	}
	if au.CountryCode == nil || *au.CountryCode != "AU" {
		t.Errorf("expected country_code 'AU', got %v", au.CountryCode)
	}
	if au.Regions == nil {
		t.Fatal("expected Australia regions to be non-nil")
	}
	if len(*au.Regions) != 8 {
		t.Errorf("expected 8 regions for Australia, got %d", len(*au.Regions))
	}

	// Verify Canada (second entry)
	ca := stateprovs[1]
	if ca.CountryId == nil || int(*ca.CountryId) != 29 {
		t.Errorf("expected Canada country_id 29, got %v", ca.CountryId)
	}
	if ca.CountryCode == nil || *ca.CountryCode != "CA" {
		t.Errorf("expected country_code 'CA', got %v", ca.CountryCode)
	}
	if ca.Regions == nil || len(*ca.Regions) != 8 {
		t.Errorf("expected 8 regions for Canada, got %v", ca.Regions)
	}

	// Verify United States (third entry)
	us := stateprovs[2]
	if us.CountryId == nil || int(*us.CountryId) != 168 {
		t.Errorf("expected US country_id 168, got %v", us.CountryId)
	}
	if us.CountryCode == nil || *us.CountryCode != "US" {
		t.Errorf("expected country_code 'US', got %v", us.CountryCode)
	}
	if us.Regions == nil || len(*us.Regions) != 52 {
		t.Errorf("expected 52 regions for US, got %v", func() int {
			if us.Regions == nil {
				return 0
			}
			return len(*us.Regions)
		}())
	}

	// Spot-check a US region
	usRegions := *us.Regions
	found := false
	for _, r := range usRegions {
		if r.RegionCode != nil && *r.RegionCode == "OH" && r.RegionName != nil && *r.RegionName == "Ohio" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find Ohio (OH) in US regions")
	}

	// Verify all countries have required fields and at least one region
	for i, sp := range stateprovs {
		if sp.CountryId == nil {
			t.Errorf("stateprov[%d]: expected country_id to be non-nil", i)
		}
		if sp.CountryName == nil || *sp.CountryName == "" {
			t.Errorf("stateprov[%d]: expected country_name to be non-empty", i)
		}
		if sp.CountryCode == nil || *sp.CountryCode == "" {
			t.Errorf("stateprov[%d]: expected country_code to be non-empty", i)
		}
		if sp.Regions == nil || len(*sp.Regions) == 0 {
			t.Errorf("stateprov[%d]: expected at least one region", i)
		}
		for j, r := range *sp.Regions {
			if r.RegionName == nil || *r.RegionName == "" {
				t.Errorf("stateprov[%d] region[%d]: expected region_name to be non-empty", i, j)
			}
			if r.RegionCode == nil || *r.RegionCode == "" {
				t.Errorf("stateprov[%d] region[%d]: expected region_code to be non-empty", i, j)
			}
		}
	}
}

func TestOtherStateProvRequestPath(t *testing.T) {
	client, mockClient := newTestClient(t, sampleOtherStateProvResponse)
	ctx := context.Background()
	_, err := client.OtherStateProvWithResponse(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Method != "GET" {
		t.Errorf("expected GET method, got %s", req.Method)
	}
	if req.URL.Path != "/other/stateprovs" {
		t.Errorf("expected path '/other/stateprovs', got %s", req.URL.Path)
	}
}
