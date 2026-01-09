package emojiflags

import (
	"strings"
)

var SpecialEmojiMap = map[string]string{
	EnglandCode:      "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	ScotlandCode:     "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	WalesCode:        "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
	EnglandShortCode: "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
}

// GetFlag converts a country code (ISO 3166-1 alpha-2, alpha-3, or CIOC) to its corresponding emoji flag.
// It supports 2-letter codes (e.g., "VN"), 3-letter codes (e.g., "VNM" or "GER"),
// and special subdivision codes (e.g., "GB-ENG" for England, "ENG" for England short code).
// Returns an empty string if the country code is not found.
func GetFlag(countryCode string) string {
	countryCode = strings.ToUpper(countryCode)
	switch len(countryCode) {
	case 2:
		if code, ok := Cca2CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
		}
	case 3:
		if code, ok := Cca3CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
		}

		if code, ok := CiocCodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
		}

		if flag, ok := SpecialEmojiMap[countryCode]; ok {
			return flag
		}
	case 6:
		if flag, ok := SpecialEmojiMap[countryCode]; ok {
			return flag
		}
	default:
		return ""
	}

	return ""
}

// GetFlagFuzzy attempts to find a flag using fuzzy matching on country codes.
// It searches for the closest match within all code maps (alpha-2, alpha-3, CIOC, and special codes).
// Returns the flag and the matched code if a close match is found (distance <= 2), otherwise returns empty strings.
// This is useful for handling typos or variations in country code input.
//
// Example:
//
//	flag, code := emojiflags.GetFlagFuzzy("VIETNM")  // Returns Vietnam flag and "VNM"
//	flag, code := emojiflags.GetFlagFuzzy("USA")     // Returns US flag and "US"
func GetFlagFuzzy(input string) (string, string) {
	input = strings.ToUpper(input)

	// Try exact match first
	if flag := GetFlag(input); flag != "" {
		return flag, input
	}

	const maxDistance = 2
	bestMatch := ""
	bestDistance := maxDistance + 1

	// Check alpha-2 codes
	for code := range Cca2CodeMap {
		dist := levenshtein(input, code)
		if dist < bestDistance {
			bestDistance = dist
			bestMatch = code
		}
	}

	// Check alpha-3 codes
	for code := range Cca3CodeMap {
		dist := levenshtein(input, code)
		if dist < bestDistance {
			bestDistance = dist
			bestMatch = code
		}
	}

	// Check CIOC codes
	for code := range CiocCodeMap {
		dist := levenshtein(input, code)
		if dist < bestDistance {
			bestDistance = dist
			bestMatch = code
		}
	}

	// Check special codes
	for code := range SpecialEmojiMap {
		dist := levenshtein(input, code)
		if dist < bestDistance {
			bestDistance = dist
			bestMatch = code
		}
	}

	if bestDistance <= maxDistance && bestMatch != "" {
		flag := GetFlag(bestMatch)
		return flag, bestMatch
	}

	return "", ""
}

// levenshtein calculates the Levenshtein distance between two strings.
// This measures the minimum number of single-character edits (insertions, deletions, or substitutions)
// required to change one string into the other.
func levenshtein(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// Create matrix
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	// Fill matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}

			deletion := matrix[i-1][j] + 1
			insertion := matrix[i][j-1] + 1
			substitution := matrix[i-1][j-1] + cost

			min := deletion
			if insertion < min {
				min = insertion
			}
			if substitution < min {
				min = substitution
			}

			matrix[i][j] = min
		}
	}

	return matrix[len(s1)][len(s2)]
}
