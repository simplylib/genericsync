package genericsync

import (
	"fmt"
	"testing"
)

func TestMap_Delete(t *testing.T) {
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
}

func TestMap_Load(t *testing.T) {
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
}

func TestMap_LoadAndDelete(t *testing.T) {
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
}

func TestMap_LoadOrStore(t *testing.T) {
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
}

func TestMap_Store(t *testing.T) {
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
}

func TestMap_Range(t *testing.T) {
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
			t.Fatalf("Panic occurred during testing Map.Range (%v)", r)
		}
	}()
	mm := Map[error, error]{}
	mm.Store(nil, nil)
	mm.Range(func(k error, v error) bool {
		return true
	})
	mm.Store(nil, fmt.Errorf("test"))
	mm.Range(func(k error, v error) bool {
		return true
	})
	mm.Store(fmt.Errorf("test"), nil)
	mm.Range(func(k error, v error) bool {
		return true
	})
}
