package collection

import (
	"cmp"
	"slices"
)

// SortSubset sorts the sub slice based on the order of elements in the already-sorted super slice.
// T is the type of elements in both slices.
// K is the comparable key type returned by the selector functions.
// superSelector extracts the key from super elements.
// subSelector extracts the key from sub elements.
func SortSubset[T any, K comparable](super []T, sub []T, superSelector func(T) K, subSelector func(T) K) {
	// 1. Map each value in super to its index (rank)
	rank := make(map[K]int, len(super))
	for i, val := range super {
		rank[superSelector(val)] = i
	}

	// 2. Sort sub using the rank map
	slices.SortFunc(sub, func(a, b T) int {
		return cmp.Compare(rank[subSelector(a)], rank[subSelector(b)])
	})
}
