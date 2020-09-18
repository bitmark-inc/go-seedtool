package seedtool

// #cgo LDFLAGS: -L/usr/local/lib -lbc-ur
// #cgo CXXFLAGS: -std=c++17
// #include "ur.h"
import "C"
import (
	"encoding/hex"
	"unsafe"
)

type WordStyle int

const (
	Starndard WordStyle = iota
	URI
	Minmal
)

func Encode(b []byte, style WordStyle) string {
	p := C.bytewords_encode(unsafe.Pointer(&b[0]), C.ulong(len(b)))
	s := C.GoString(p)
	return s
}

func Decode(hex string, style WordStyle) []byte {
	var size C.ulong
	p := C.bytewords_decode(&size, C.CString(hex))
	b := C.GoBytes(p, C.int(size))
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
