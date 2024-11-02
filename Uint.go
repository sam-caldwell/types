package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Uint - represents an unsigned -bit integer type.
type Uint uint

// String - converts the Uint to a string.
func (u *Uint) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Parse - converts a string to a Uint value.
func (u *Uint) Parse(v string) error {
	sz := int(unsafe.Sizeof(uint(0))) * 8
	n, err := strconv.ParseUint(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Uint(n)
	return nil
}

// Type - returns a type string
func (u *Uint) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
