package genericsync

import (
	"fmt"
	"testing"
)

func TestMap_CompareAndDelete(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		m := Map[int, int]{}

		m.Store(0, 1)
		if !m.CompareAndDelete(0, 1) {
			t.Fatalf("expected CompareAndDelete to return true")
		}

		if _, ok := m.Load(0); ok {
			t.Fatalf("Map contained value after it should have been deleted")
		}

		m.Store(1, 0)
		if m.CompareAndDelete(1, 1) {
			t.Fatalf("CompareAndDelete with wrong value should not have worked")
		}
	})

	t.Run("pointer types", func(t *testing.T) {
		t.Parallel()

		m := Map[*int, *int]{}

		zero, one := 0, 1

		m.Store(&zero, &one)
		if !m.CompareAndDelete(&zero, &one) {
			t.Fatalf("expected CompareAndDelete to return true")
		}

		if _, ok := m.Load(&zero); ok {
			t.Fatalf("Map contained value after it should have been deleted")
		}

		m.Store(&one, &zero)
		if m.CompareAndDelete(&one, &one) {
			t.Fatalf("CompareAndDelete with wrong value should not have worked")
		}
	})
}

func TestMap_CompareAndSwap(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		m := Map[int, int]{}

		m.Store(0, 1)

		v, ok := m.Load(0)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}
		if v != 1 {
			t.Fatalf("value not 1, instead %v", v)
		}

		if !m.CompareAndSwap(0, 1, 2) {
			t.Fatalf("value isn't 1")
		}

		v, ok = m.Load(0)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}

		if v != 2 {
			t.Fatalf("key 0 should have been 2, instead was %v", v)
		}
	})
	t.Run("pointer types", func(t *testing.T) {
		t.Parallel()

		m := Map[*int, *int]{}

		zero, one, two := 0, 1, 2

		m.Store(&zero, &one)

		v, ok := m.Load(&zero)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}
		if v != &one {
			t.Fatalf("value not 1, instead %v", v)
		}

		if !m.CompareAndSwap(&zero, &one, &two) {
			t.Fatalf("value isn't 1")
		}

		v, ok = m.Load(&zero)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}

		if v != &two {
			t.Fatalf("key 0 should have been 2, instead was %v", v)
		}
	})
}

func TestMap_Delete(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var (
			m  Map[int, int]
			ok bool
		)
		for i := 0; i < 10; i++ {
			m.Store(i, i+1)
			m.Delete(i)
			_, ok = m.Load(i)
			if ok {
				t.Fatalf("could not delete (%v)", i)
			}
		}
	})
}

func TestMap_Load(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var (
			m     Map[int, int]
			value int
			ok    bool
		)
		for i := 0; i < 10; i++ {
			m.m.Store(i, i+1)
			value, ok = m.Load(i)
			if !ok {
				t.Fatal("value does not exist")
			}
			if value != i+1 {
				t.Fatalf("value expected (%v) got (%v)", i+1, value)
			}
		}

		v, ok := m.Load(20)
		if ok {
			t.Fatalf("map contains key (20), when not set with value (%v)", v)
		}
	})
}

func TestMap_LoadAndDelete(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var (
			m      Map[int, int]
			v      int
			loaded bool
		)
		for i := 0; i < 10; i++ {
			m.Store(i, i+1)
			v, loaded = m.LoadAndDelete(i)
			if !loaded {
				t.Fatalf("key (%v) doesnt exist after setting value (%v)", i, i+1)
			}
			if v != i+1 {
				t.Fatalf("key (%v) expected value (%v) got (%v)", i, i+1, v)
			}
		}

		v, loaded = m.LoadAndDelete(20)
		if loaded {
			t.Fatalf("loaded key (20) when key never set with value (%v)", v)
		}
	})
}

func TestMap_LoadOrStore(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var (
			m      Map[int, int]
			actual int
			loaded bool
		)
		for i := 0; i < 10; i++ {
			actual, loaded = m.LoadOrStore(i, i+1)
			if loaded {
				t.Fatalf("key (%v) loaded after not being set, value (%v)", i, actual)
			}

			if actual != i+1 {
				t.Fatalf("key (%v) value (%v) not (%v))", i, actual, i+1)
			}
		}

		actual, loaded = m.LoadOrStore(0, 0)
		if !loaded {
			t.Fatal("key (0) not loaded after being stored")
		}

		if actual != 1 {
			t.Fatal("key (0) not value (1)")
		}
	})
}

func TestMap_Store(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var (
			m     Map[int, int]
			value int
			ok    bool
		)
		for i := 0; i < 10; i++ {
			m.Store(i, i+1)
			value, ok = m.Load(i)
			if !ok {
				t.Fatalf("key (%v) did not load after being set", i)
			}
			if value != i+1 {
				t.Fatalf("key (%v) expected value (%v) got (%v)", i, i+1, value)
			}
		}
	})
}

func TestMap_Range(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		var m Map[int, int]
		const count = 10
		for i := 0; i < count; i++ {
			m.Store(i, i+1)
		}

		vals := map[int]int{}
		m.Range(func(k int, v int) bool {
			vals[k] = v
			return true
		})
		if len(vals) != count {
			t.Fatalf("expected len(vals) to be (%v) got (%v)", count, len(vals))
		}
		for i, val := range vals {
			if val != i+1 {
				t.Fatalf("key (%v) expected value (%v) got (%v)", i, i+1, val)
			}
		}

		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Panic occurred during testing Map.Range (type: [%T], value: [%#+v])", r, r)
			}
		}()
		mm := Map[error, error]{}
		mm.Store(nil, nil)
		mm.Range(func(_ error, _ error) bool {
			return true
		})
		mm.Store(nil, fmt.Errorf("test"))
		mm.Range(func(_ error, _ error) bool {
			return true
		})
		mm.Store(fmt.Errorf("test"), nil)
		mm.Range(func(_ error, _ error) bool {
			return true
		})
	})
}

func TestMap_Swap(t *testing.T) {
	t.Parallel()

	t.Run("golden path", func(t *testing.T) {
		t.Parallel()

		m := Map[int, int]{}

		m.Store(0, 1)

		v, ok := m.Load(0)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}
		if v != 1 {
			t.Fatalf("value not 1, instead %v", v)
		}

		previous, loaded := m.Swap(0, 2)
		if !loaded {
			t.Fatalf("map didn't contain key 0")
		}
		if previous != 1 {
			t.Fatalf("previous value should have been 1")
		}

		v, ok = m.Load(0)
		if !ok {
			t.Fatalf("key 0 should have been in the map, but wasn't")
		}
		if v != 2 {
			t.Fatalf("value not 2, instead %v", v)
		}
	})
}
