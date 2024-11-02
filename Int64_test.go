package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt64(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Int64(0)).Kind() != reflect.Int64 {
			t.Fatal("type mismatch")
		}
		if mx := Int64(math.MaxInt64); mx != math.MaxInt64 {
			t.Fatal("bounds check upper")
		}
		if mn := Int64(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Int64(math.MaxInt64)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Int64(math.MaxInt64)
		var (
			lhs Int64
		)
		stringValue := fmt.Sprintf("%d", rhs)
		if err := lhs.Parse(stringValue); err == nil {
			if lhs != rhs {
				t.Fatalf("mismatch (lhs: %d, rhs: %d, stringValue: '%s')", lhs, rhs, stringValue)
			}
		} else {
			t.Fatalf("Parse error: '%v' (lhs: %d, rhs: %d, stringValue: '%s')", err, lhs, rhs, stringValue)
		}
	})

	t.Run("verify Type() method", func(t *testing.T) {
		n := Int64(0)
		if typ := n.Type(); typ != "int64" {
			t.Fatalf("Int64 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
