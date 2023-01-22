package collection

func ToSet[TSlice ~[]T, T any, K comparable](slice TSlice, keySelector func(T) K) Set[K] {
	set := NewSet(make([]K, 0)...)
	for _, t := range slice {
		set.Add(keySelector(t))
	}
	return set
}

func ToMap[TSlice ~[]T, T any, K comparable, V any](slice TSlice, keySelector func(T) K, valSelector func(T) V) map[K]V {
	m := make(map[K]V)
	for _, t := range slice {
		m[keySelector(t)] = valSelector(t)
	}

	return m
}

func DistinctBy[TSlice ~[]T, T any, K comparable](slice TSlice, keySelector func(T) K) TSlice {
	newSlice := make([]T, 0, len(slice))
	seen := NewSet[K]()
	for _, s := range slice {
		key := keySelector(s)
		if !seen.Contains(key) {
			seen.Add(key)
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func FirstMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) T {
	for _, t := range slice {
		if predicate(t) {
			return t
		}
	}

	var zero T
	return zero
}

func Select[TSlice ~[]T, T any, V any](source TSlice, selector func(t T) V) []V {
	var result []V
	for _, t := range source {
		selected := selector(t)
		result = append(result, selected)
	}
	return result
}

func Filter[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) TSlice {
	selected := make(TSlice, 0)
	for _, t := range slice {
		if predicate(t) {
			selected = append(selected, t)
		}
	}

	return selected
}

// AnyMatch returns true if any element passes the predicate function
func AnyMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) bool {
	for _, t := range slice {
		if predicate(t) {
			return true
		}
	}

	return false
}

// AllMatch returns true if all elements pass the predicate function
func AllMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) bool {
	for _, t := range slice {
		if !predicate(t) {
			return false
		}
	}

	return true
}

// Contains returns true if find appears in slice
func Contains[TSlice ~[]T, T comparable](slice TSlice, find T) bool {
	for _, t := range slice {
		if t == find {
			return true
		}
	}

	return false
}

func IndexOf[TSlice ~[]T, T comparable](slice TSlice, find T) int {
	for i, t := range slice {
		if t == find {
			return i
		}
	}

	return -1
}

// GroupBy returns a map that is keyed by keySelector and contains a slice of elements returned by valSelector
func GroupBy[TSlice ~[]T, T any, K comparable, V any](slice TSlice, keySelector func(T) K, valSelector func(T) V) map[K][]V {
	grouping := make(map[K][]V)
	for _, t := range slice {
		key := keySelector(t)
		grouping[key] = append(grouping[key], valSelector(t))
	}

	return grouping
}

// Resolve the difference between two slice
func Diff[TSlice ~[]T, VSlice ~[]V, T any, V any, K comparable](left TSlice, right VSlice, leftKeySelector func(T) K, rightKeySelector func(V) K) DiffResult[T, V] {
	leftMap := ToMap(left, leftKeySelector, func(t T) T {
		return t
	})
	leftMapKeys := make([]K, 0, len(leftMap))
	for k := range leftMap {
		leftMapKeys = append(leftMapKeys, k)
	}
	leftMapKeySet := NewSet(leftMapKeys...)
	rightMap := ToMap(right, rightKeySelector, func(v V) V {
		return v
	})
	rightMapKeys := make([]K, 0, len(rightMap))
	for k := range rightMap {
		rightMapKeys = append(rightMapKeys, k)
	}
	rightMapKeySet := NewSet(rightMapKeys...)

	leftOnlyKeys := leftMapKeySet.Diff(rightMapKeySet).List()
	rightOnlyKeys := rightMapKeySet.Diff(leftMapKeySet).List()

	var leftOnly []T
	var rightOnly []V

	for _, k := range leftOnlyKeys {
		leftOnly = append(leftOnly, leftMap[k])
	}
	for _, k := range rightOnlyKeys {
		rightOnly = append(rightOnly, rightMap[k])
	}

	return DiffResult[T, V]{
		LeftOnly:  leftOnly,
		RightOnly: rightOnly,
	}

}
