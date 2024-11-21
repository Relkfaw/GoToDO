package entity

import (
	"testing"
)

func Test_newClient(t *testing.T) {
	tests := []struct {
		name        string
		givenName   string
		wantedError error
	}{
		{
			name:        "Valid Name",
			givenName:   "Michel",
			wantedError: nil,
		},
		{
			name:        "Empty Name",
			givenName:   "",
			wantedError: ErrEmptyName,
		},
		{
			name:        "Bad format Name",
			givenName:   "Michel@",
			wantedError: ErrFormatName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newClient(tt.givenName)
			if err != tt.wantedError {
				t.Errorf("Expected error to be %v, but got %v", tt.wantedError, err)
			}
		})
	}
}
