package client

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func createMockResponse(statusCode int, body string) *http.Response {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     header,
	}
}

func newTestClient(t *testing.T, body string) (*ClientWithResponses, *MockHTTPClient) {
	t.Helper()
	mockClient := &MockHTTPClient{Response: createMockResponse(200, body)}
	client, err := NewClientWithResponses(defaultBaseURL, WithHTTPClient(mockClient))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client, mockClient
}
