package types

import (
	"fmt"
	"math"
	"testing"
)

func TestFloat64(t *testing.T) {

	n := Float64(math.MaxFloat64)
	s := fmt.Sprintf("%f", math.MaxFloat64)

	t.Run("test float64 type", func(t *testing.T) {
		if float64(n) != math.MaxFloat64 {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test String() method", func(t *testing.T) {
		if s != n.String() {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test Parse() method", func(t *testing.T) {
		var m Float64
		if err := m.Parse(s); err != nil {
			t.Fatal(err)
		}
		if m != n {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test Type method", func(t *testing.T) {
		if typ := n.Type(); typ != "float64" {
			t.Fatal("type mismatch")
		}
	})

}
