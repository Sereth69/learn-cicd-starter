package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "Valid API key",
			headers: http.Heder{"Authorization": []string{"ApiKey validAPIKey"}},
			want:    "validAPIKey",
			wantErr: nil,
		},
		{
			name:    "No Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "Malformed Authorization header",
			headers: http.Header{"Authorization": []string{"InvalidHeader"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}
