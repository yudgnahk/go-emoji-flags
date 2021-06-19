package emojiflags

import (
	"testing"
	"unicode"
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
			7,
		},
		{
			"Should handle correct 2 characters input",
			args{"VN"},
			7,
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
			7,
		},
		{
			"Could get England emoji",
			args{"GB-ENG"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFlag(tt.args.country)
			if !utf8.ValidString(got) {
				t.Errorf("GetFlag() expected valid flag got %v", got)
			}
			if GetPrintableLength(got) != tt.expectedLen {
				t.Errorf("expected length emoji of %v got %v", tt.expectedLen, GetPrintableLength(got))
			}
		})
	}
}

func GetPrintableLength(s string) int {
	res := 0
	for i := range s {
		if unicode.IsPrint(rune(s[i])) {
			res++
		}
	}

	return res
}
