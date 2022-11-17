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

func GetFlag(countryCode string) string {
	countryCode = strings.ToUpper(countryCode)
	switch len(countryCode) {
	case 2:
		if code, ok := Cca2CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A') + "\u0020"
		}
	case 3:
		if code, ok := Cca3CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A') + "\u0020"
		}

		if code, ok := CiocCodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A') + "\u0020"
		}
	case 6:
		if code, ok := SpecialCountryMap[countryCode]; ok {
			return SpecialEmojiMap[code]
		}
	default:
		return ""
	}

	return ""
}
