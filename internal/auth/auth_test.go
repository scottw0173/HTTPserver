package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expected    string
		expectError bool
	}{
		{
			name:        "valid bearer token",
			headerValue: "Bearer abc123",
			expected:    "abc123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headerValue: "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "missing bearer prefix",
			headerValue: "abc123",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty bearer token",
			headerValue: "Bearer ",
			expected:    "",
			expectError: true,
		},
		{
			name:        "extra whitespace",
			headerValue: "Bearer    abc123",
			expected:    "abc123",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}

			if tc.headerValue != "" {
				headers.Set("Authorization", tc.headerValue)
			}

			token, err := GetBearerToken(headers)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if token != tc.expected {
				t.Errorf("expected token %q, got %q", tc.expected, token)
			}
		})
	}
}
