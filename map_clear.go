//go:build go1.23

package genericsync

func (m *Map[K, V]) Clear() {
	m.m.Clear()
}
