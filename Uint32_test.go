package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestUint32(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Uint32(0)).Kind() != reflect.Uint32 {
			t.Fatal("type mismatch")
		}
		if mx := Uint32(math.MaxUint32); mx != math.MaxUint32 {
			t.Fatal("bounds check upper")
		}
		if mn := Uint32(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Uint32(math.MaxUint32)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Uint32(math.MaxUint32)
		var (
			lhs Uint32
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
		n := Uint32(0)
		if typ := n.Type(); typ != "uint32" {
			t.Fatalf("Uint32 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
