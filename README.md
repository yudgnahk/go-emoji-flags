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

Will return a flag, or an empty string if the flag is not found.

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

### Special thanks to those repositories which helps me so much:
 - [go-emoji-flag](https://github.com/jayco/go-emoji-flag)
 - [restcountries](https://github.com/apilayer/restcountries)

... and wikipedia for the understanding about countries code (especially ISO_3166) (https://en.wikipedia.org/wiki/ISO_3166-2:GB)

I was stuck on the countries that belongs to Great Britain, which doesn't have the same format as origin emojis
([go-emoji-flag](https://github.com/jayco/go-emoji-flag) does not support these countries)