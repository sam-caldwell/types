package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Int64 - represents an unsigned 64-bit integer type.
type Int64 int64

// String - converts the Int64 to a string.
func (u *Int64) String() string {
	return strconv.FormatInt(int64(*u), 10)
}

// Parse - converts a string to an Int64 value.
func (u *Int64) Parse(v string) error {
	sz := int(unsafe.Sizeof(int64(0))) * 8
	n, err := strconv.ParseInt(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Int64(n)
	return nil
}

// Type - returns a type string
func (u *Int64) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
