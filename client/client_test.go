package client

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name           string
		server         string
		opts           []ClientOption
		expectedServer string
		expectError    bool
	}{
		{
			name:           "basic client creation",
			server:         "https://api.example.com",
			opts:           nil,
			expectedServer: "https://api.example.com/",
			expectError:    false,
		},
		{
			name:           "server with trailing slash",
			server:         "https://api.example.com/",
			opts:           nil,
			expectedServer: "https://api.example.com/",
			expectError:    false,
		},
		{
			name:           "server with path",
			server:         "https://api.example.com/v1",
			opts:           nil,
			expectedServer: "https://api.example.com/v1/",
			expectError:    false,
		},
		{
			name:   "with custom HTTP client",
			server: "https://api.example.com",
			opts: []ClientOption{
				WithHTTPClient(&MockHTTPClient{}),
			},
			expectedServer: "https://api.example.com/",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.server, tt.opts...)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if client.Server != tt.expectedServer {
				t.Errorf("expected server %q, got %q", tt.expectedServer, client.Server)
			}

			if client.Client == nil {
				t.Error("expected HTTP client to be set")
			}
		})
	}
}

func TestWithHTTPClient(t *testing.T) {
	mockClient := &MockHTTPClient{}

	client, err := NewClient("https://api.example.com", WithHTTPClient(mockClient))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if client.Client != mockClient {
		t.Error("expected custom HTTP client to be set")
	}
}

func TestWithRequestEditorFn(t *testing.T) {
	called := false
	editorFn := func(ctx context.Context, req *http.Request) error {
		called = true
		req.Header.Set("X-Custom-Header", "test-value")
		return nil
	}

	client, err := NewClient("https://api.example.com", WithRequestEditorFn(editorFn))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(client.RequestEditors) != 1 {
		t.Errorf("expected 1 request editor, got %d", len(client.RequestEditors))
	}

	// Test that the request editor is called
	req, _ := http.NewRequest("GET", "https://api.example.com/test", nil)
	ctx := context.Background()

	err = client.RequestEditors[0](ctx, req)
	if err != nil {
		t.Errorf("unexpected error from request editor: %v", err)
	}

	if !called {
		t.Error("expected request editor to be called")
	}

	if req.Header.Get("X-Custom-Header") != "test-value" {
		t.Error("expected custom header to be set")
	}
}

func TestMultipleRequestEditors(t *testing.T) {
	var callOrder []int

	editor1 := func(ctx context.Context, req *http.Request) error {
		callOrder = append(callOrder, 1)
		return nil
	}

	editor2 := func(ctx context.Context, req *http.Request) error {
		callOrder = append(callOrder, 2)
		return nil
	}

	client, err := NewClient("https://api.example.com",
		WithRequestEditorFn(editor1),
		WithRequestEditorFn(editor2),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(client.RequestEditors) != 2 {
		t.Errorf("expected 2 request editors, got %d", len(client.RequestEditors))
	}

	// Test that both editors are called in order
	req, _ := http.NewRequest("GET", "https://api.example.com/test", nil)
	ctx := context.Background()

	for _, editor := range client.RequestEditors {
		err = editor(ctx, req)
		if err != nil {
			t.Errorf("unexpected error from request editor: %v", err)
		}
	}

	if len(callOrder) != 2 || callOrder[0] != 1 || callOrder[1] != 2 {
		t.Errorf("expected call order [1, 2], got %v", callOrder)
	}
}

func TestClientInterface(t *testing.T) {
	// Test that our Client type implements ClientInterface
	var _ ClientInterface = &Client{}
}

// Test parameter types and constants
func TestParameterTypes(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected interface{}
	}{
		{
			name:     "ViewDirectorToursParamsTimePeriod FUTURE",
			value:    FUTURE,
			expected: ViewDirectorToursParamsTimePeriod("FUTURE"),
		},
		{
			name:     "ViewDirectorToursParamsTimePeriod PAST",
			value:    PAST,
			expected: ViewDirectorToursParamsTimePeriod("PAST"),
		},
		{
			name:     "ViewPlayerActiveResultsParamsRankingSystem MAIN",
			value:    ViewPlayerActiveResultsParamsRankingSystemMAIN,
			expected: ViewPlayerActiveResultsParamsRankingSystem("MAIN"),
		},
		{
			name:     "ViewPlayerActiveResultsParamsRankingSystem WOMEN",
			value:    ViewPlayerActiveResultsParamsRankingSystemWOMEN,
			expected: ViewPlayerActiveResultsParamsRankingSystem("WOMEN"),
		},
		{
			name:     "ViewPlayerActiveResultsParamsRankingSystem YOUTH",
			value:    ViewPlayerActiveResultsParamsRankingSystemYOUTH,
			expected: ViewPlayerActiveResultsParamsRankingSystem("YOUTH"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, tt.value)
			}
		})
	}
}

func TestSearchDirectorsParams(t *testing.T) {
	name := "John Doe"
	count := 25

	params := SearchDirectorsParams{
		Name:  &name,
		Count: &count,
	}

	if *params.Name != "John Doe" {
		t.Errorf("expected name 'John Doe', got %q", *params.Name)
	}

	if *params.Count != 25 {
		t.Errorf("expected count 25, got %d", *params.Count)
	}
}

func TestViewPlayerMultiParams(t *testing.T) {
	players := "123,456,789"

	params := ViewPlayerMultiParams{
		Players: players,
	}

	if params.Players != "123,456,789" {
		t.Errorf("expected players '123,456,789', got %q", params.Players)
	}
}

func TestSearchPlayersParams(t *testing.T) {
	name := "Jane Smith"
	country := "US"
	stateprov := "CA"
	tournament := "World Championship"
	tourpos := float32(1.0)

	params := SearchPlayersParams{
		Name:       &name,
		Country:    &country,
		Stateprov:  &stateprov,
		Tournament: &tournament,
		Tourpos:    &tourpos,
	}

	if *params.Name != "Jane Smith" {
		t.Errorf("expected name 'Jane Smith', got %q", *params.Name)
	}

	if *params.Country != "US" {
		t.Errorf("expected country 'US', got %q", *params.Country)
	}

	if *params.Stateprov != "CA" {
		t.Errorf("expected stateprov 'CA', got %q", *params.Stateprov)
	}

	if *params.Tournament != "World Championship" {
		t.Errorf("expected tournament 'World Championship', got %q", *params.Tournament)
	}

	if *params.Tourpos != 1.0 {
		t.Errorf("expected tourpos 1.0, got %f", *params.Tourpos)
	}
}

func TestRankingCountryParams(t *testing.T) {
	startPos := float32(10)
	count := float32(50)
	country := "US"

	params := RankingCountryParams{
		StartPos: &startPos,
		Count:    &count,
		Country:  country,
	}

	if *params.StartPos != 10 {
		t.Errorf("expected startPos 10, got %f", *params.StartPos)
	}

	if *params.Count != 50 {
		t.Errorf("expected count 50, got %f", *params.Count)
	}

	if params.Country != "US" {
		t.Errorf("expected country 'US', got %q", params.Country)
	}
}

func TestSeriesRegionsParamsSeriesCode(t *testing.T) {
	testCases := []SeriesRegionsParamsSeriesCode{
		ACS,
		NACS,
		WNACSO,
		WNACSW,
	}

	expectedValues := []string{"ACS", "NACS", "WNACSO", "WNACSW"}

	for i, code := range testCases {
		if string(code) != expectedValues[i] {
			t.Errorf("expected %q, got %q", expectedValues[i], string(code))
		}
	}
}

func TestTourSearchParamsTypes(t *testing.T) {
	// Test RankType constants
	if MAIN != TourSearchParamsRankType("MAIN") {
		t.Errorf("expected MAIN to equal 'MAIN'")
	}

	if WOMEN != TourSearchParamsRankType("WOMEN") {
		t.Errorf("expected WOMEN to equal 'WOMEN'")
	}

	// Test EventType constants
	if Tournament != TourSearchParamsEventType("Tournament") {
		t.Errorf("expected Tournament to equal 'Tournament'")
	}

	if Keague != TourSearchParamsEventType("Keague") {
		t.Errorf("expected Keague to equal 'Keague'")
	}

	// Test PreRegistration constants
	if TourSearchParamsPreRegistrationY != TourSearchParamsPreRegistration("Y") {
		t.Errorf("expected Y to equal 'Y'")
	}

	if TourSearchParamsPreRegistrationN != TourSearchParamsPreRegistration("N") {
		t.Errorf("expected N to equal 'N'")
	}
}

func TestApi_keyScopes(t *testing.T) {
	if API_KEYScopes != "API_KEY.Scopes" {
		t.Errorf("expected 'API_KEY.Scopes', got %q", API_KEYScopes)
	}
}

// Test that we can create parameters with nil pointers (optional fields)
func TestOptionalParameters(t *testing.T) {
	// Test SearchDirectorsParams with nil values
	params := SearchDirectorsParams{}
	if params.Name != nil {
		t.Error("expected Name to be nil")
	}
	if params.Count != nil {
		t.Error("expected Count to be nil")
	}

	// Test RankingCountryParams with required and optional fields
	countryParams := RankingCountryParams{
		Country: "US", // required
		// StartPos and Count are optional (nil)
	}
	if countryParams.StartPos != nil {
		t.Error("expected StartPos to be nil")
	}
	if countryParams.Count != nil {
		t.Error("expected Count to be nil")
	}
	if countryParams.Country != "US" {
		t.Errorf("expected Country 'US', got %q", countryParams.Country)
	}
}

// Benchmark tests for client creation
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewClient("https://api.example.com")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkNewClientWithOptions(b *testing.B) {
	mockClient := &MockHTTPClient{}
	editor := func(ctx context.Context, req *http.Request) error {
		return nil
	}

	for i := 0; i < b.N; i++ {
		_, err := NewClient("https://api.example.com",
			WithHTTPClient(mockClient),
			WithRequestEditorFn(editor),
		)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func TestClient_WithRequestEditor(t *testing.T) {
	mockClient := &MockHTTPClient{Response: createMockResponse(200, `{"test": "response"}`)}

	requestEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-Custom-Header", "test-value")
		req.Header.Set("Authorization", "Bearer token123")
		return nil
	}

	client, err := NewClient(defaultBaseURL,
		WithHTTPClient(mockClient),
		WithRequestEditorFn(requestEditor),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	_, err = client.ViewCountryDirectors(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	if req.Header.Get("X-Custom-Header") != "test-value" {
		t.Errorf("expected X-Custom-Header 'test-value', got %q", req.Header.Get("X-Custom-Header"))
	}
	if req.Header.Get("Authorization") != "Bearer token123" {
		t.Errorf("expected Authorization 'Bearer token123', got %q", req.Header.Get("Authorization"))
	}
}

func TestClient_RequestEditorError(t *testing.T) {
	mockClient := &MockHTTPClient{Response: createMockResponse(200, `{"test": "response"}`)}

	requestEditor := func(ctx context.Context, req *http.Request) error {
		return errors.New("request editor error")
	}

	client, err := NewClient(defaultBaseURL,
		WithHTTPClient(mockClient),
		WithRequestEditorFn(requestEditor),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	_, err = client.ViewCountryDirectors(ctx, requestEditor)
	if err == nil {
		t.Error("expected error from request editor, got nil")
	}
	if !strings.Contains(err.Error(), "request editor error") {
		t.Errorf("expected error message to contain 'request editor error', got %q", err.Error())
	}
}

func TestClient_HTTPError(t *testing.T) {
	mockClient := &MockHTTPClient{Error: errors.New("network error")}

	client, err := NewClient(defaultBaseURL, WithHTTPClient(mockClient))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	_, err = client.ViewCountryDirectors(ctx)
	if err == nil {
		t.Error("expected error from HTTP client, got nil")
	}
	if !strings.Contains(err.Error(), "network error") {
		t.Errorf("expected error message to contain 'network error', got %q", err.Error())
	}
}

func TestClient_MethodOverride(t *testing.T) {
	tests := []struct {
		name     string
		method   func(client *Client, ctx context.Context) (*http.Response, error)
		wantPath string
	}{
		{
			name: "ViewCountryDirectors",
			method: func(client *Client, ctx context.Context) (*http.Response, error) {
				return client.ViewCountryDirectors(ctx)
			},
			wantPath: "director/country",
		},
		{
			name: "OtherCountries",
			method: func(client *Client, ctx context.Context) (*http.Response, error) {
				return client.OtherCountries(ctx)
			},
			wantPath: "other/countries",
		},
		{
			name: "OtherStateProv",
			method: func(client *Client, ctx context.Context) (*http.Response, error) {
				return client.OtherStateProv(ctx)
			},
			wantPath: "other/stateprovs",
		},
		{
			name: "RankingCountryList",
			method: func(client *Client, ctx context.Context) (*http.Response, error) {
				return client.RankingCountryList(ctx)
			},
			wantPath: "rankings/country_list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{Response: createMockResponse(200, `{}`)}

			client, err := NewClient(defaultBaseURL, WithHTTPClient(mockClient))
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}

			ctx := context.Background()
			_, err = tt.method(client, ctx)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			req := mockClient.Requests[0]
			if !strings.Contains(req.URL.Path, tt.wantPath) {
				t.Errorf("expected URL path to contain %q, got %s", tt.wantPath, req.URL.Path)
			}
		})
	}
}

func TestClient_ParameterSerialization(t *testing.T) {
	mockClient := &MockHTTPClient{Response: createMockResponse(200, `{}`)}

	client, err := NewClient(defaultBaseURL, WithHTTPClient(mockClient))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	_, err = client.SearchDirectors(ctx, &SearchDirectorsParams{Name: nil, Count: nil})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	req := mockClient.Requests[0]
	query := req.URL.Query()
	if query.Has("name") {
		t.Error("expected 'name' parameter to be absent when nil")
	}
	if query.Has("count") {
		t.Error("expected 'count' parameter to be absent when nil")
	}
}

func TestClient_MultipleRequestEditors(t *testing.T) {
	mockClient := &MockHTTPClient{Response: createMockResponse(200, `{}`)}

	var callOrder []string

	editor1 := func(ctx context.Context, req *http.Request) error {
		callOrder = append(callOrder, "editor1")
		req.Header.Set("X-Editor-1", "value1")
		return nil
	}

	editor2 := func(ctx context.Context, req *http.Request) error {
		callOrder = append(callOrder, "editor2")
		req.Header.Set("X-Editor-2", "value2")
		return nil
	}

	client, err := NewClient(defaultBaseURL,
		WithHTTPClient(mockClient),
		WithRequestEditorFn(editor1),
		WithRequestEditorFn(editor2),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	_, err = client.ViewCountryDirectors(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(callOrder) != 2 || callOrder[0] != "editor1" || callOrder[1] != "editor2" {
		t.Errorf("expected call order [editor1, editor2], got %v", callOrder)
	}

	req := mockClient.Requests[0]
	if req.Header.Get("X-Editor-1") != "value1" {
		t.Error("expected X-Editor-1 header to be set")
	}
	if req.Header.Get("X-Editor-2") != "value2" {
		t.Error("expected X-Editor-2 header to be set")
	}
}
