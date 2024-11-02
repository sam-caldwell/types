package types

import (
	"fmt"
	"math"
	"testing"
)

func TestFloat32(t *testing.T) {

	n := Float32(math.MaxFloat32)
	s := fmt.Sprintf("%f", math.MaxFloat32)

	t.Run("test float32 type", func(t *testing.T) {
		if float32(n) != math.MaxFloat32 {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test String() method", func(t *testing.T) {
		if s != n.String() {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test Parse() method", func(t *testing.T) {
		var m Float32
		if err := m.Parse(s); err != nil {
			t.Fatal(err)
		}
		if m != n {
			t.Fatal("value mismatch")
		}
	})

	t.Run("test Type method", func(t *testing.T) {
		if typ := n.Type(); typ != "float32" {
			t.Fatal("type mismatch")
		}
	})

}
