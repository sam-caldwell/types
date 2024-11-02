package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Uint8 - represents an unsigned 8-bit integer type.
type Uint8 uint8

// String - converts the Uint8 to a string.
func (u *Uint8) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Parse - converts a string to a Uint8 value.
func (u *Uint8) Parse(v string) error {
	sz := int(unsafe.Sizeof(uint8(0))) * 8
	n, err := strconv.ParseUint(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Uint8(n)
	return nil
}

// Type - returns a type string
func (u *Uint8) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
