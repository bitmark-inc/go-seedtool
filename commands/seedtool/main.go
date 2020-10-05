package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"time"

	"github.com/bitmark-inc/go-seedtool"
	"github.com/fxamacker/cbor/v2"
)

// A CBOR structure for seed data
type SeedData struct {
	Seed      []byte   `cbor:"1,keyasint,omitempty"`
	Timestamp cbor.Tag `cbor:"2,keyasint,omitempty"`
}

func NewSeedData(seed string) SeedData {
	bb, err := hex.DecodeString(seed)
	if err != nil {
		panic(err)
	}

	v := SeedData{
		Seed: bb,
		Timestamp: cbor.Tag{
			Number:  100,
			Content: uint(time.Now().Unix() / 86400),
		},
	}

	return v
}

func UREncodeSeed(seed string) []string {
	data := NewSeedData(seed)
	b, err := cbor.Marshal(data) // encode v to []byte b

	if err != nil {
		panic(err)
	}

	return seedtool.UREncodeSeed(b, 64)
}

func URDecodeSeed(urParts []string) string {
	b := seedtool.URDecodeSeed(urParts)

	var data SeedData
	if err := cbor.Unmarshal(b, &data); err != nil {
		panic(err)
	}

	return hex.EncodeToString(data.Seed)
}

func main() {
	var seed string
	flag.StringVar(&seed, "seed", "", "seed for encode/decode")
	flag.Parse()

	bw := seedtool.EncodeFromHex(seed, seedtool.URI)
	fmt.Println("=====> Bytewords Encode: ", bw)

	s := seedtool.DecodeToHex(bw, seedtool.URI)
	fmt.Println("=====> Bytewords Decode: ", s)

	fmt.Println("=====> UR Encode Seed <=====")
	urParts := UREncodeSeed(seed)
	for _, s := range urParts {
		fmt.Println(s)
	}
	fmt.Println("=====> UR Decode Seed <=====")
	fmt.Println(URDecodeSeed(urParts))
}
