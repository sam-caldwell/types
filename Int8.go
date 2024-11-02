package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Int8 - represents an unsigned 8-bit integer type.
type Int8 int8

// String - converts the Int8 to a string.
func (u *Int8) String() string {
	return strconv.FormatInt(int64(*u), 10)
}

// Parse - converts a string to an Int8 value.
func (u *Int8) Parse(v string) error {
	sz := int(unsafe.Sizeof(int8(0))) * 8
	n, err := strconv.ParseInt(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Int8(n)
	return nil
}

// Type - returns a type string
func (u *Int8) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
