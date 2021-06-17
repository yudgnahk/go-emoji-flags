package emojiflags

import (
	"strings"
)

var SpecialEmojiMap = map[string]string{
	EnglandCode:  "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	ScotlandCode: "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	WalesCode:    "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
}

func GetFlag(countryCode string) string {
	countryCode = strings.ToUpper(countryCode)
	switch len(countryCode) {
	case 2:
		if code, ok := Alpha2CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
		}
	case 3:
		if code, ok := Alpha3CodeMap[countryCode]; ok {
			return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
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
