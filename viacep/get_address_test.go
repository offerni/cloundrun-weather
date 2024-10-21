package viacep

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAddress(t *testing.T) {
	tests := []struct {
		name           string
		cep            string
		mockResponse   string
		mockStatusCode int
		expected       *GetAddressResponse
		expectError    bool
	}{
		{
			name:           "Valid response",
			cep:            "01001000",
			mockResponse:   `{"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}`,
			mockStatusCode: http.StatusOK,
			expected: &GetAddressResponse{
				Cep:         "01001-000",
				Logradouro:  "Praça da Sé",
				Complemento: "lado ímpar",
				Bairro:      "Sé",
				Localidade:  "São Paulo",
				Uf:          "SP",
				Ibge:        "3550308",
				Gia:         "1004",
				Ddd:         "11",
				Siafi:       "7107",
			},
			expectError: false,
		},
		{
			name:           "Empty CEP response",
			cep:            "00000000",
			mockResponse:   `{"cep":""}`,
			mockStatusCode: http.StatusOK,
			expected:       nil,
			expectError:    false,
		},
		{
			name:           "Invalid JSON response",
			cep:            "invalid",
			mockResponse:   `{invalid json}`,
			mockStatusCode: http.StatusOK,
			expected:       nil,
			expectError:    true,
		},
		{
			name:           "HTTP error",
			cep:            "00000000",
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
			ctx := context.Background()
			result, err := GetAddress(ctx, tt.cep)
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
