package collection

import "sort"

type OrderedMap[K comparable, V any] struct {
	Data        map[K]V
	orderedKeys []K
	keyIndexMap map[K]int // useful for delete operation
	isOrdered   bool
	sorter      Sorter[K]
}


type Sorter[K any] interface {

	Less(a K, b K) bool
}

func (om *OrderedMap[K, V]) ensureOrder() {
	keys := make([]K, 0, len(om.Data))
	for key := range om.Data {
		keys = append(keys, key)
	}


	lessFunc := func(i, j int) bool {
		return om.sorter.Less(keys[i], keys[j])
	}
	sort.Slice(keys, lessFunc)

	om.orderedKeys = keys
	om.keyIndexMap = make(map[K]int)
	for idx, key := range om.orderedKeys {
		om.keyIndexMap[key] = idx
	}
	om.isOrdered = true
}


func (om OrderedMap[K, V]) BuildFrom(
	data map[K]V, sorter Sorter[K],
) OrderedMap[K, V] {
	om.Data = data
	om.sorter = sorter
	om.ensureOrder()
	return om
}

func (om OrderedMap[K, V]) Range() <-chan (K) {
	iterChan := make(chan K)
	go func() {
		defer close(iterChan)
		for _, key := range om.orderedKeys {
			iterChan <- key
		}
	}()
	return iterChan
}

func (om OrderedMap[K, V]) Has(key K) bool {
	_, exists := om.Data[key]
	return exists
}

func (om OrderedMap[K, V]) Len() int {
	return len(om.Data)
}


func (om *OrderedMap[K, V]) Keys() []K {
	if !om.isOrdered {
		om.ensureOrder()
	}
	return om.orderedKeys
}


func (om *OrderedMap[K, V]) Set(key K, val V) {
	om.Data[key] = val
	om.ensureOrder()
}

func (om *OrderedMap[K, V]) Delete(key K) {
	idx, keyExists := om.keyIndexMap[key]
	if keyExists {
		delete(om.Data, key)

		orderedKeys := om.orderedKeys
		orderedKeys = append(orderedKeys[:idx], orderedKeys[idx+1:]...)
		om.orderedKeys = orderedKeys
	}
}

func intIsLess(a, b int) bool {
	return a < b
}

func OrderedMap_Int[V any](data map[int]V) OrderedMap[int, V] {
	omap := OrderedMap[int, V]{}
	return omap.BuildFrom(data, intSorter{})
}

type intSorter struct{}

var _ Sorter[int] = (*intSorter)(nil)

func (sorter intSorter) Less(a, b int) bool {
	return intIsLess(a, b)
}
