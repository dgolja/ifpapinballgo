package client

import (
	"context"
	"testing"
)

func TestExampleParameterCombinations(t *testing.T) {
	name := "Test Director"
	count10 := 10
	count25 := 25
	tests := []struct {
		name   string
		params *SearchDirectorsParams
	}{
		{name: "with both name and count", params: &SearchDirectorsParams{Name: &name, Count: &count10}},
		{name: "with name only", params: &SearchDirectorsParams{Name: &name}},
		{name: "with count only", params: &SearchDirectorsParams{Count: &count25}},
		{name: "with no parameters", params: &SearchDirectorsParams{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{Response: createMockResponse(200, `{}`)}
			client, err := NewClient(defaultBaseURL, WithHTTPClient(mockClient))
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			ctx := context.Background()
			resp, err := client.SearchDirectors(ctx, tt.params)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if resp.StatusCode != 200 {
				t.Errorf("expected status code 200, got %d", resp.StatusCode)
			}
		})
	}
}
