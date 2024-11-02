package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestUint16(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Uint16(0)).Kind() != reflect.Uint16 {
			t.Fatal("type mismatch")
		}
		if mx := Uint16(math.MaxUint16); mx != math.MaxUint16 {
			t.Fatal("bounds check upper")
		}
		if mn := Uint16(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Uint16(math.MaxUint16)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Uint16(math.MaxUint16)
		var (
			lhs Uint16
		)

		if err := lhs.Parse(fmt.Sprintf("%d", rhs)); err != nil {
			t.Fatal(err)
		} else {
			if lhs != rhs {
				t.Fatalf("mismatch (lhs: %d, rhs: %d)", lhs, rhs)
			}
		}
	})

	t.Run("verify Type() method", func(t *testing.T) {
		n := Uint16(0)
		if typ := n.Type(); typ != "uint16" {
			t.Fatalf("Uint16 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
