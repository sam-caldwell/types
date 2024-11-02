package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Int(0)).Kind() != reflect.Int {
			t.Fatal("type mismatch")
		}
		if mx := Int(math.MaxInt); mx != math.MaxInt {
			t.Fatal("bounds check upper")
		}
		if mn := Int(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Int(math.MaxInt)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Int(math.MaxInt)
		var (
			lhs Int
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
		n := Int(0)
		if typ := n.Type(); typ != "int" {
			t.Fatalf("Int type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
