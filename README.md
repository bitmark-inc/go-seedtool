# go-seedtool

This is a practice to create a go-version seedtool by wrapping
libraries from Blockchain Commons

## Prerequsite

* go 1.14+
* xcode 11.7+
* [libbc-ur](http://github.com/blockchainCommons/bc-ur)

## Build

```
go build
```

## Examples

``` go
import (
	"flag"
	"fmt"

	ur "github.com/bitmark-inc/go-seedtool-cli"
)

func Run() {
	bw := ur.EncodeFromHex("8935a8068526d84da555cdb741a3b8a8", ur.URI)
	fmt.Println("Encode: ", bw)

	s := ur.DecodeToHex(bw, ur.URI)
	fmt.Println("Decode: ", s)
}

```
