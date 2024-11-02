package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Int - represents an unsigned -bit integer type.
type Int int

// String - converts the Uint to a string.
func (u *Int) String() string {
	return strconv.FormatInt(int64(*u), 10)
}

// Parse - converts a string to a Uint value.
func (u *Int) Parse(v string) error {
	sz := int(unsafe.Sizeof(0)) * 8
	n, err := strconv.ParseInt(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Int(n)
	return nil
}

// Type - returns a type string
func (u *Int) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
