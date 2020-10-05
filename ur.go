package seedtool

// #cgo LDFLAGS: -L/usr/local/lib -lbc-ur
// #cgo CXXFLAGS: -std=c++17
// #include <stdlib.h> // without this the `C.free` will not be found
// #include "ur.h"
import "C"
import (
	"encoding/hex"
	"strings"
	"unsafe"
)

type WordStyle int

const (
	Starndard WordStyle = iota
	URI
	Minmal
)

func Encode(b []byte, style WordStyle) string {
	p := (*C.uchar)(&b[0])
	ret := C.bytewords_encode(p, C.ulong(len(b)))
	return C.GoString(ret)
}

func Decode(hex string, style WordStyle) []byte {
	var size C.ulong

	cs := C.CString(hex)
	defer C.free(unsafe.Pointer(cs))

	p := C.bytewords_decode(&size, cs)
	b := C.GoBytes(unsafe.Pointer(p), C.int(size))
	return b
}

func EncodeFromHex(s string, style WordStyle) string {
	b, err := hex.DecodeString(s)
	if err != nil {
		return ""
	}

	return Encode(b, style)
}

func DecodeToHex(s string, style WordStyle) string {
	src := Decode(s, style)
	return hex.EncodeToString(src)
}

// UREncodeSeed accepts a CBOR buffer of a seed and returns UR strings
func UREncodeSeed(seed []byte, fragementSize uint64) []string {
	fSize := C.ulong(fragementSize)

	p := (*C.uchar)(&seed[0])
	ret := C.ur_encode_seed(p, C.ulong(len(seed)), fSize)

	return strings.Split(C.GoString(ret), "\n")
}

// URDecodeSeed accepts a CBOR buffer of a seed and returns UR strings
func URDecodeSeed(urParts []string) []byte {

	cParts := make([]*C.char, len(urParts))
	for i := 0; i < len(cParts); i++ {
		cs := C.CString(urParts[i])
		defer C.free(unsafe.Pointer(cs)) // can not free here
		cParts[i] = cs
	}

	var retSize C.ulong
	// ret := C.ur_decode_seed(&retSize, (**C.char)(cParts))
	ret := C.ur_decode_seed(&retSize, C.ulong(len(urParts)), &cParts[0])
	return C.GoBytes(unsafe.Pointer(ret), C.int(retSize))
}
