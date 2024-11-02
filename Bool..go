package types

import (
	"fmt"
	"reflect"
	"strings"
)

// Bool - represents a boolean value
type Bool bool

// String - converts the Bool to a string.
func (u *Bool) String() string {
	if *u {
		return "true"
	}
	return "false"
}

// Parse - converts a string to an Bool value.
func (u *Bool) Parse(v string) error {
	switch v := strings.ToLower(strings.TrimSpace(v)); v {
	case "true":
		*u = true
	case "false":
		*u = false
	default:
		*u = false
		return fmt.Errorf("unknown value (expect true/false): %s", v)
	}
	return nil
}

// Type - returns a type string
func (u *Bool) Type() string {
	return strings.ToLower(strings.TrimSpace(strings.Split(reflect.TypeOf(*u).String(), ".")[1]))
}
