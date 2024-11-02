package types

import (
	"testing"
)

func TestBool(t *testing.T) {
	t.Run("test boolean type", func(t *testing.T) {
		b := Bool(true)
		if b != true {
			t.Fatal("value mismatch: true")
		}
		b = Bool(false)
		if b != false {
			t.Fatal("value mismatch: false")
		}
	})
	t.Run("test boolean String method", func(t *testing.T) {
		b := Bool(true)
		if b.String() != "true" {
			t.Fatal("value mismatch: true")
		}
	})
	t.Run("test boolean Parse", func(t *testing.T) {
		var b Bool
		if err := b.Parse("true"); err != nil {
			t.Fatal(err)
		}
		if !b {
			t.Fatal("value mismatch: true")
		}
		if err := b.Parse("false"); err != nil {
			t.Fatal(err)
		}
		if b {
			t.Fatal("value mismatch: false")
		}
		if err := b.Parse("True"); err != nil {
			t.Fatal(err)
		}
		if !b {
			t.Fatal("value mismatch: true")
		}
		if err := b.Parse("False"); err != nil {
			t.Fatal(err)
		}
		if b {
			t.Fatal("value mismatch: false")
		}
		if err := b.Parse("1"); err == nil {
			t.Fatal("expected error not found(1)")
		}
		if b {
			t.Fatal("value mismatch: false")
		}
		if err := b.Parse("bad"); err == nil {
			t.Fatal("expected error not found(bad)")
		}
		if b {
			t.Fatal("value mismatch: false")
		}
		if err := b.Parse(""); err == nil {
			t.Fatal("expected error not found()")
		}
		if b {
			t.Fatal("value mismatch: false")
		}
	})
}
