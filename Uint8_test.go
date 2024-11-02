package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestUint8(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Uint8(0)).Kind() != reflect.Uint8 {
			t.Fatal("type mismatch")
		}
		if mx := Uint8(math.MaxUint8); mx != math.MaxUint8 {
			t.Fatal("bounds check upper")
		}
		if mn := Uint8(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Uint8(math.MaxUint8)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		var lhs Uint8
		rhs := Uint8(math.MaxUint8)

		if err := lhs.Parse(fmt.Sprintf("%d", rhs)); err != nil {
			t.Fatalf("Parse failed: '%v' (lhs: %d, rhs: %d)", err, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf("mismatch (lhs: %d, rhs: %d)", lhs, rhs)
		}
	})

	t.Run("verify Type() method", func(t *testing.T) {
		n := Uint8(0)
		if typ := n.Type(); typ != "uint8" {
			t.Fatalf("Uint8 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
