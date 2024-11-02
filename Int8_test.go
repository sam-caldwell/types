package types

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt8(t *testing.T) {

	t.Run("confirm type", func(t *testing.T) {
		if reflect.TypeOf(Int8(0)).Kind() != reflect.Int8 {
			t.Fatal("type mismatch")
		}
		if mx := Int8(math.MaxInt8); mx != math.MaxInt8 {
			t.Fatal("bounds check upper")
		}
		if mn := Int8(0); mn != 0 {
			t.Fatal("bounds check lower")
		}
	})

	t.Run("confirm String() method", func(t *testing.T) {
		value := Int8(math.MaxInt8)
		expected := fmt.Sprintf("%d", value)
		actual := value.String()
		if actual != expected {
			t.Fatalf("mismatch (actual: %s, expected: %s)", actual, expected)
		}
	})

	t.Run("confirm Parse() method", func(t *testing.T) {
		var lhs Int8
		rhs := Int8(math.MaxInt8)

		if err := lhs.Parse(fmt.Sprintf("%d", rhs)); err != nil {
			t.Fatalf("Parse failed: '%v' (lhs: %d, rhs: %d)", err, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf("mismatch (lhs: %d, rhs: %d)", lhs, rhs)
		}
	})

	t.Run("verify Type() method", func(t *testing.T) {
		n := Int8(0)
		if typ := n.Type(); typ != "int8" {
			t.Fatalf("Int8 type mismatch. Got: '%d', Want: '%s'", n, typ)
		}
	})
}
