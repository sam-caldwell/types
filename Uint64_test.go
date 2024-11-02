package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestUint64(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Uint64(0)).Kind() != reflect.Uint64 {
			t.Fatal("type mismatch")
		}
		if mx := Uint64(math.MaxUint64); mx != math.MaxUint64 {
			t.Fatal("bounds check upper")
		}
		if mn := Uint64(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Uint64(math.MaxUint64)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Uint64(math.MaxUint64)
		var (
			lhs Uint64
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
		n := Uint64(0)
		if typ := n.Type(); typ != "uint64" {
			t.Fatalf("Uint64 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
