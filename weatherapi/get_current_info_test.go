package weatherapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentInfo(t *testing.T) {
	tests := []struct {
		name           string
		city           string
		mockResponse   string
		mockStatusCode int
		expected       *GetCurrentInfoResponse
		expectError    bool
	}{
		{
			name:           "Valid response",
			city:           "Vancouver",
			mockResponse:   `{"current":{"temp_c":20.0,"temp_f":68.0}}`,
			mockStatusCode: http.StatusOK,
			expected: &GetCurrentInfoResponse{
				TempC: 20.0,
				TempF: 68.0,
				TempK: 293.15,
			},
			expectError: false,
		},
		{
			name:           "Invalid JSON response",
			city:           "InvalidCity",
			mockResponse:   `{invalid json}`,
			mockStatusCode: http.StatusOK,
			expected:       nil,
			expectError:    true,
		},
		{
			name:           "HTTP error",
			city:           "ErrorCity",
			mockResponse:   ``,
			mockStatusCode: http.StatusInternalServerError,
			expected:       nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				_, _ = w.Write([]byte(tt.mockResponse))
			}))

			defer mockServer.Close()
			baseUrl = mockServer.URL
			client := &Client{APIKey: "dummy"}
			ctx := context.Background()
			result, err := client.GetCurrentInfo(ctx, tt.city)

			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}

			if !tt.expectError && result != nil {
				if *result != *tt.expected {
					t.Fatalf("expected: %+v, got: %+v", tt.expected, result)
				}
			}

			if !tt.expectError && result == nil && tt.expected != nil {
				t.Fatalf("expected: %+v, got: nil", tt.expected)
			}
		})
	}
}
