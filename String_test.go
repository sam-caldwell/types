package types

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Verify String method", func(t *testing.T) {
		s := String("test")
		if s.String() != "test" {
			t.Fatal("value mismatch")
		}
	})
	t.Run("verify Type method", func(t *testing.T) {
		s := String("test")
		if typ := s.Type(); typ != "string" {
			t.Fatal("type mismatch")
		}
	})
}
