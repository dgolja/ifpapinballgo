package client

import (
	"net/http"
)

// MockHTTPClient is a mock implementation of HttpRequestDoer for testing
type MockHTTPClient struct {
	Response *http.Response
	Error    error
	Requests []*http.Request
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.Requests = append(m.Requests, req)
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Response, nil
}
