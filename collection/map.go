package collection

// Keys returns a slice containing all keys from the given map.
// The order of keys is not guaranteed and may vary between iterations.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
