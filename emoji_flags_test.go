package emojiflags

import (
	"testing"
	"unicode/utf8"
)

func Test_getFlag(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name        string
		args        args
		expectedLen int
	}{
		{
			"Should handle correct 3 characters input",
			args{"VNM"},
			28,
		},
		{
			"Should handle correct 2 characters input",
			args{"VN"},
			28,
		},
		{
			"Should return empty string if no 3 letters code found",
			args{"BOB"},
			0,
		},
		{
			"Should return empty string if no 2 letters match found",
			args{"AA"},
			0,
		},
		{
			"Should uppercase input",
			args{"vnm"},
			28,
		},
		{
			"Could get England emoji",
			args{"GB-ENG"},
			28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFlag(tt.args.country)
			if !utf8.ValidString(got) {
				t.Errorf("GetFlag() expected valid flag got %v", got)
			}
			if len(got) != tt.expectedLen {
				t.Errorf("expected length emoji of %v got %v", tt.expectedLen, len(got))
			}
		})
	}
}
