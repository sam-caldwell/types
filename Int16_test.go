package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt16(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Int16(0)).Kind() != reflect.Int16 {
			t.Fatal("type mismatch")
		}
		if mx := Int16(math.MaxInt16); mx != math.MaxInt16 {
			t.Fatal("bounds check upper")
		}
		if mn := Int16(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Int16(math.MaxInt16)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Int16(math.MaxInt16)
		var (
			lhs Int16
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
		n := Int16(0)
		if typ := n.Type(); typ != "int16" {
			t.Fatalf("Int16 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
