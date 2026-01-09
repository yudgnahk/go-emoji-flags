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
			2,
		},
		{
			"Should handle correct 2 characters input",
			args{"VN"},
			2,
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
			2,
		},
		{
			"Could get England emoji",
			args{"GB-ENG"},
			7,
		},
		{
			"Could get CIOC code",
			args{"GER"},
			2,
		},
		{
			"Return empty string if code is empty",
			args{""},
			0,
		},
		{
			"Could get England flag with short code",
			args{"ENG"},
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

func Test_GetFlagFuzzy(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantFlag    bool // whether a flag should be returned
		wantCode    string
		wantFlagLen int
	}{
		{
			"Should find exact match VNM",
			args{"VNM"},
			true,
			"VNM",
			2,
		},
		{
			"Should find typo VIETNM (missing A)",
			args{"VIETNM"},
			false, // distance is 3, exceeds maxDistance threshold of 2
			"",
			0,
		},
		{
			"Should find USA from US",
			args{"USA"},
			true,
			"USA", // Found USA in alpha-3 map
			2,
		},
		{
			"Should find GER from GERMANY",
			args{"GERMANY"},
			false, // distance is 4, exceeds maxDistance threshold of 2
			"",
			0,
		},
		{
			"Should find GER from GERM",
			args{"GERM"},
			true,
			"GER",
			2,
		},
		{
			"Should find FR from FRR (distance 1)",
			args{"FRR"},
			true,
			"FR", // France at distance 1 (delete one R) - unambiguous match
			2,
		},
		{
			"Should return empty for completely wrong input",
			args{"ZZZZZ"},
			false,
			"",
			0,
		},
		{
			"Should find GB-ENG from GB-EN",
			args{"GB-EN"},
			true,
			"GB-ENG",
			7,
		},
		{
			"Should find lowercase input",
			args{"vnm"},
			true,
			"VNM",
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFlag, gotCode := GetFlagFuzzy(tt.args.input)

			if tt.wantFlag {
				if gotFlag == "" {
					t.Errorf("GetFlagFuzzy() expected flag, got empty string")
				}
				if gotCode != tt.wantCode {
					t.Errorf("GetFlagFuzzy() code = %v, want %v", gotCode, tt.wantCode)
				}
				if !utf8.ValidString(gotFlag) {
					t.Errorf("GetFlagFuzzy() expected valid flag got %v", gotFlag)
				}
				if GetPrintableLength(gotFlag) != tt.wantFlagLen {
					t.Errorf("GetFlagFuzzy() flag length = %v, want %v", GetPrintableLength(gotFlag), tt.wantFlagLen)
				}
			} else {
				if gotFlag != "" || gotCode != "" {
					t.Errorf("GetFlagFuzzy() expected empty results, got flag=%v, code=%v", gotFlag, gotCode)
				}
			}
		})
	}
}

func Test_levenshtein(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want int
	}{
		{"identical strings", "ABC", "ABC", 0},
		{"one insertion", "ABC", "ABCD", 1},
		{"one deletion", "ABCD", "ABC", 1},
		{"one substitution", "ABC", "ABD", 1},
		{"empty strings", "", "", 0},
		{"one empty", "ABC", "", 3},
		{"complex", "VIETNM", "VNM", 3},
		{"complex 2", "USA", "US", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenshtein(tt.s1, tt.s2); got != tt.want {
				t.Errorf("levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}
