package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Float64 - represents an unsigned 64-bit integer type.
type Float64 float64

// String - converts the Float64 to a string.
func (u *Float64) String() string {
	return fmt.Sprintf("%f", float64(*u))
}

// Parse - converts a string to an Float64 value.
func (u *Float64) Parse(v string) error {
	sz := int(unsafe.Sizeof(float64(0))) * 8
	n, err := strconv.ParseFloat(v, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Float64(n)
	return nil
}

// Type - returns a type string
func (u *Float64) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
