package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt32(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Int32(0)).Kind() != reflect.Int32 {
			t.Fatal("type mismatch")
		}
		if mx := Int32(math.MaxInt32); mx != math.MaxInt32 {
			t.Fatal("bounds check upper")
		}
		if mn := Int32(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Int32(math.MaxInt32)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		rhs := Int32(math.MaxInt32)
		var (
			lhs Int32
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
		n := Int32(0)
		if typ := n.Type(); typ != "int32" {
			t.Fatalf("Int32 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
