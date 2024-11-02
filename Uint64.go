package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Uint64 - represents an unsigned 64-bit integer type.
type Uint64 uint64

// String - converts the Uint64 to a string.
func (u *Uint64) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Parse - converts a string to a Uint64 value.
func (u *Uint64) Parse(v string) error {
	sz := int(unsafe.Sizeof(uint64(0))) * 8
	n, err := strconv.ParseUint(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Uint64(n)
	return nil
}

// Type - returns a type string
func (u *Uint64) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
