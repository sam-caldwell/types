package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Float32 - represents an unsigned 64-bit integer type.
type Float32 float32

// String - converts the Float32 to a string.
func (u *Float32) String() string {
	return fmt.Sprintf("%f", float32(*u))
}

// Parse - converts a string to an Float32 value.
func (u *Float32) Parse(v string) error {
	sz := int(unsafe.Sizeof(float32(0))) * 8
	n, err := strconv.ParseFloat(v, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Float32(n)
	return nil
}

// Type - returns a type string
func (u *Float32) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
