package auth

import (
	"errors"
	"net/http"
	"testing"
)

type expected struct {
	key string
	err error
}

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected expected
	}{
		{
			name: "valid key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-key"},
			},
			expected: expected{
				key: "valid-key",
				err: nil,
			},
		},
		{
			name:    "empty header",
			headers: http.Header{},
			expected: expected{
				key: "",
				err: ErrNoAuthHeaderIncluded,
			},
		},
		{
			name: "invalid key",
			headers: http.Header{
				"Authorization": []string{"Bearerinvalidtoken"},
			},
			expected: expected{
				key: "fail",
				err: ErrMalformedAuthHeader,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, err := GetAPIKey(test.headers)
			if key != test.expected.key || !errors.Is(err, test.expected.err) {
				t.Errorf("GetAPIKey() = %v, %v; want %v, %v", key, err, test.expected.key, test.expected.err)
			}
		})
	}

}
