# go-emoji-flags

Converts a string country code to an emoji in Go.

## Features

- ✅ Support for ISO 3166-1 alpha-2 codes (e.g., "VN")
- ✅ Support for ISO 3166-1 alpha-3 codes (e.g., "VNM") 
- ✅ Support for CIOC codes (e.g., "GER")
- ✅ Support for Great Britain subdivisions (England, Scotland, Wales)
- ✅ Fuzzy matching for typos and variations (e.g., "USA" → "US", "GERM" → "GER")
- ✅ Consistent emoji output length (no trailing spaces)

## Install
```
go get -u github.com/yudgnahk/go-emoji-flags
```

## Usage

### Basic Usage

Convert country codes to flag emojis:

```go
package main

import (
	"fmt"

	emoji "github.com/yudgnahk/go-emoji-flags"
)

func main() {
	fmt.Println(emoji.GetFlag("VNM"))   // prints 🇻🇳
	fmt.Println(emoji.GetFlag("VN"))    // prints 🇻🇳
	fmt.Println(emoji.GetFlag("BOB"))   // prints (empty string)
}
```

### Reverse Lookup (Flag to Code)

Convert flag emojis back to country codes:

```go
code := emoji.GetCode("🇻🇳")    // Returns "VN"
code := emoji.GetCode("🏴󠁧󠁢󠁥󠁮󠁧󠁿")  // Returns "GB-ENG"
```

### Get Country Names

Get country names from codes or flags:

```go
name := emoji.GetName("VN")     // Returns "Vietnam"
name := emoji.GetName("VNM")    // Returns "Vietnam"
name := emoji.GetName("GER")    // Returns "Germany"
name := emoji.GetName("🇻🇳")     // Returns "Vietnam"
```

### Get Flag by Country Name

Find flags using country names (with fuzzy matching):

```go
flag, code := emoji.GetFlagByName("Vietnam")        // Returns "🇻🇳", "VN"
flag, code := emoji.GetFlagByName("United States")  // Returns "🇺🇸", "US"
flag, code := emoji.GetFlagByName("USA")            // Returns "🇺🇸", "US" (alias support)
flag, code := emoji.GetFlagByName("UK")             // Returns "🇬🇧", "GB" (alias support)
```

### Fuzzy Matching

Use `GetFlagFuzzy()` to handle typos or variations in country codes:

```go
package main

import (
	"fmt"

	emoji "github.com/yudgnahk/go-emoji-flags"
)

func main() {
	// Exact match still works
	flag, code := emoji.GetFlagFuzzy("VNM")
	fmt.Printf("%s (matched: %s)\n", flag, code) // 🇻🇳 (matched: VNM)
	
	// Fuzzy matching handles typos (within distance of 2)
	flag, code = emoji.GetFlagFuzzy("USA")
	fmt.Printf("%s (matched: %s)\n", flag, code) // 🇺🇸 (matched: USA)
	
	flag, code = emoji.GetFlagFuzzy("GERM")
	fmt.Printf("%s (matched: %s)\n", flag, code) // 🇩🇪 (matched: GER)
	
	flag, code = emoji.GetFlagFuzzy("GB-EN")
	fmt.Printf("%s (matched: %s)\n", flag, code) // 🏴󠁧󠁢󠁥󠁮󠁧󠁿 (matched: GB-ENG)
	
	// Returns empty if no close match found
	flag, code = emoji.GetFlagFuzzy("ZZZZZ")
	fmt.Printf("'%s' (matched: '%s')\n", flag, code) // '' (matched: '')
}
```

## Advanced Usage

### Error Handling

The `GetFlag()` function returns an empty string when no match is found:

```go
flag := emoji.GetFlag("INVALID")
if flag == "" {
    log.Println("Country code not found")
}
```

### Fuzzy Matching with Validation

Use fuzzy matching to suggest corrections to users:

```go
userInput := "GERM"
flag, matchedCode := emoji.GetFlagFuzzy(userInput)

if flag == "" {
    fmt.Printf("No match found for: %s\n", userInput)
} else if matchedCode != userInput {
    fmt.Printf("Did you mean %s? %s\n", matchedCode, flag)
} else {
    fmt.Printf("Found: %s\n", flag)
}
```

### Batch Processing

Process multiple country codes efficiently:

```go
codes := []string{"VN", "US", "GB", "INVALID"}
for _, code := range codes {
    if flag := emoji.GetFlag(code); flag != "" {
        fmt.Printf("%s: %s\n", code, flag)
    } else {
        fmt.Printf("%s: Not found\n", code)
    }
}
```

## Performance Considerations

- **GetFlag()**: Fast O(1) map lookup - use for known-good codes
- **GetFlagFuzzy()**: Slower O(n) search - use only for user input that may contain typos
- **Fuzzy matching distance**: Limited to 2 character edits for performance
- **Caching**: Consider caching fuzzy results if you process the same queries repeatedly

### Benchmarks

```
BenchmarkGetFlag-12              35772781    33.55 ns/op      8 B/op   1 allocs/op
BenchmarkGetFlagFuzzy-12         28222648    41.78 ns/op      8 B/op   1 allocs/op
BenchmarkGetFlagFuzzyClose-12       13245    90293 ns/op  194969 B/op   4261 allocs/op
```

*Run `go test -bench=. -benchmem` to see all benchmarks*

## Supported Codes

### ISO 3166-1 Codes
- **Alpha-2**: 2-letter codes (e.g., VN, US, GB)
- **Alpha-3**: 3-letter codes (e.g., VNM, USA, GBR)
- **CIOC**: Olympic codes (e.g., GER for Germany, SUI for Switzerland)

### Special Subdivisions
- **England**: `GB-ENG` or `ENG`
- **Scotland**: `GB-SCT` or `SCT`
- **Wales**: `GB-WLS` or `WLS`

## API Reference

### `GetFlag(countryCode string) string`
Converts a country code to its emoji flag. Supports ISO 3166-1 alpha-2, alpha-3, CIOC codes, and special subdivisions.

**Parameters:**
- `countryCode` - 2-letter (VN), 3-letter (VNM, GER), or special codes (GB-ENG)

**Returns:** Flag emoji string, or empty string if not found

### `GetFlagFuzzy(input string) (string, string)`
Finds a flag using fuzzy matching (Levenshtein distance ≤ 2). Prefers shorter codes when multiple matches exist.

**Parameters:**
- `input` - Country code (possibly with typos)

**Returns:** Flag emoji and matched code, or empty strings if no match

### `GetCode(flag string) string`
Converts a flag emoji to its ISO 3166-1 alpha-2 country code.

**Parameters:**
- `flag` - Flag emoji (e.g., 🇻🇳)

**Returns:** Country code (e.g., "VN"), or empty string if not recognized

### `GetName(input string) string`
Converts a country code or flag emoji to the country name.

**Parameters:**
- `input` - Country code (alpha-2, alpha-3, CIOC) or flag emoji

**Returns:** Country name, or empty string if not found

### `GetFlagByName(name string) (string, string)`
Finds a flag by country name using exact or fuzzy matching (Levenshtein distance ≤ 2). Supports common aliases.

**Parameters:**
- `name` - Country name or alias (e.g., "Vietnam", "USA", "UK")

**Returns:** Flag emoji and matched code, or empty strings if no match

## Data Maps

The library provides access to several data maps:

- `CountryNames` - map[string]string: ISO alpha-2 codes to country names
- `CountryAliases` - map[string]string: Common aliases to ISO alpha-2 codes
- `Cca2CodeMap` - map[string]string: ISO alpha-2 code mappings
- `Cca3CodeMap` - map[string]string: ISO alpha-3 to alpha-2 code mappings
- `CiocCodeMap` - map[string]string: CIOC to alpha-2 code mappings
- `SpecialEmojiMap` - map[string]string: Special subdivision codes to emoji flags

### Special thanks to those repositories which helps me so much:
 - [go-emoji-flag](https://github.com/jayco/go-emoji-flag)
 - [restcountries](https://github.com/apilayer/restcountries)

... and wikipedia for the understanding about countries code (especially ISO_3166) (https://en.wikipedia.org/wiki/ISO_3166-2:GB)

I was stuck on the countries that belongs to Great Britain, which doesn't have the same format as origin emojis
([go-emoji-flag](https://github.com/jayco/go-emoji-flag) does not support these countries)