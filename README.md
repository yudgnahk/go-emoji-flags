# go-emoji-flags

Converts a string country code to an emoji in Go.

##TODO:
 - [ ] Return emoji with the same actual length (which was printed)
 - [ ] Find the emoji with nearly correct country name/code.


## Install
```
go get -u github.com/yudgnahk/go-emoji-flags
```

## Usage

Will return a flag, or an empty string if the flag does not found.

```go
package main

import (
	"fmt"

	emoji "github.com/yudgnahk/go-emoji-flags"
)

func main() {
	fmt.Println(emoji.GetFlag("VNM"))   // prints 🇻🇳
	fmt.Println(emoji.GetFlag("VN"))    // prints 🇻🇳
	fmt.Println(emoji.GetFlag("BOB"))   // prints
}
```

### Special thanks to those repositories which helps me so much:
 - [go-emoji-flag](https://github.com/jayco/go-emoji-flag)
 - [restcountries](https://github.com/apilayer/restcountries)

... and wikipedia for the understanding about countries code (especially ISO_3166) (https://en.wikipedia.org/wiki/ISO_3166-2:GB)

I was stuck on the countries that belongs to Great Britain, which doesn't have the same format as origin emojis
([go-emoji-flag](https://github.com/jayco/go-emoji-flag) does not support these countries)