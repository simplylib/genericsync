//go:build go1.23

package genericsync

import (
	"testing"
)

func TestMap_Clear(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		m := Map[int, int]{}

		zero, one := 0, 1

		m.Store(zero, one)
		v, ok := m.Load(zero)
		if !ok {
			t.Fatalf("value %v should have existed at %v, but didn't", one, zero)
		}

		if v != one {
			t.Fatalf("value should have been %v but was loaded as %v", one, v)
		}

		m.Clear()

		if v, ok = m.Load(zero); ok {
			t.Fatalf("should be no values in map, but previous value existed: %v:%v", zero, v)
		}
	})

	t.Run("pointer types", func(t *testing.T) {
		t.Parallel()

		m := Map[*int, *int]{}

		zero, one := 0, 1

		m.Store(&zero, &one)
		v, ok := m.Load(&zero)
		if !ok {
			t.Fatalf("value %v should have existed at %v, but didn't", one, zero)
		}

		if v != &one {
			t.Fatalf("value should have been %v but was loaded as %v", one, v)
		}

		m.Clear()

		if v, ok = m.Load(&zero); ok {
			t.Fatalf("should be no values in map, but previous value existed: %v:%v", zero, v)
		}
	})
}
