package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Int16 - represents an unsigned 16-bit integer type.
type Int16 int16

// String - converts the Int16 to a string.
func (u *Int16) String() string {
	return strconv.FormatInt(int64(*u), 10)
}

// Parse - converts a string to an Int16 value.
func (u *Int16) Parse(v string) error {
	sz := int(unsafe.Sizeof(int16(0))) * 8
	n, err := strconv.ParseInt(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Int16(n)
	return nil
}

// Type - returns a type string
func (u *Int16) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
