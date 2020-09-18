package main

import (
	"flag"
	"fmt"

	"github.com/bitmark-inc/go-seedtool"
)

func main() {
	var seed string
	flag.StringVar(&seed, "seed", "", "seed for encode/decode")
	flag.Parse()

	bw := seedtool.EncodeFromHex(seed, seedtool.URI)
	fmt.Println("Encode: ", bw)

	s := seedtool.DecodeToHex(bw, seedtool.URI)
	fmt.Println("Decode: ", s)
}
