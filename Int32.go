package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Int32 - represents an unsigned 32-bit integer type.
type Int32 int32

// String - converts the Int32 to a string.
func (u *Int32) String() string {
	return strconv.FormatInt(int64(*u), 10)
}

// Parse - converts a string to an Int32 value.
func (u *Int32) Parse(v string) error {
	sz := int(unsafe.Sizeof(int32(0))) * 8
	n, err := strconv.ParseInt(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Int32(n)
	return nil
}

// Type - returns a type string
func (u *Int32) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
