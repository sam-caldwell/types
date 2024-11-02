package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestUint(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Uint(0)).Kind() != reflect.Uint {
			t.Fatal("type mismatch")
		}
		if mx := Uint(math.MaxUint); mx != math.MaxUint {
			t.Fatal("bounds check upper")
		}
		if mn := Uint(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Uint(math.MaxUint)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Uint(math.MaxUint)
		var (
			lhs Uint
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
		n := Uint(0)
		if typ := n.Type(); typ != "uint" {
			t.Fatalf("Uint type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
