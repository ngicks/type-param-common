package sync

import (
	sync_ "sync"
)

type Map[K, V any] struct {
	inner sync_.Map
}

func (m *Map[K, V]) Delete(key K) {
	m.inner.Delete(key)
}
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.inner.Load(key)
	if v == nil {
		// It could be untyped nil or whatever inconvertible to type V.
		// Return zero-value in that case.
		return
	}
	return v.(V), ok
}
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.inner.LoadAndDelete(key)
	if v == nil {
		return
	}
	return v.(V), loaded
}
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.inner.LoadOrStore(key, value)
	return v.(V), loaded
}
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.inner.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
func (m *Map[K, V]) Store(key K, value V) {
	m.inner.Store(key, value)
}
