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
