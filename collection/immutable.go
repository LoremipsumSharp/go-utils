package collection

import "github.com/huandu/go-clone"

type ImmutableMap[K comparable, V any] struct {
	m map[K]V
}

func NewImmutableMap[K comparable, V any](m map[K]V) ImmutableMap[K, V] {
	return ImmutableMap[K, V]{
		m: clone.Slowly(m).(map[K]V),
	}
}

func (i ImmutableMap[K, V]) Get(key K) (V, bool) {
	v, ok := i.m[key]
	if !ok {
		return v, ok
	}
	return clone.Slowly(v).(V), ok
}


func (i ImmutableMap[K, V]) Keys() []K {
	keys := []K{}
	for k, _ := range i.m {
		keys = append(keys, k)
	}
	return keys
}