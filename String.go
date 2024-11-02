package types

import (
	"reflect"
	"strings"
)

// String - represents a string
type String string

// String - converts the Int8 to a string.
func (u *String) String() string {
	return string(*u)
}

// Type - returns a type string
func (u *String) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
