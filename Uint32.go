package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Uint32 - represents an unsigned 32-bit integer type.
type Uint32 uint32

// String - converts the Uint32 to a string.
func (u *Uint32) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Parse - converts a string to a Uint32 value.
func (u *Uint32) Parse(v string) error {
	sz := int(unsafe.Sizeof(uint32(0))) * 8
	n, err := strconv.ParseUint(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Uint32(n)
	return nil
}

// Type - returns a type string
func (u *Uint32) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
