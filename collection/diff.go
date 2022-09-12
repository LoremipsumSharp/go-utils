package collection

type DiffResult[K any, V any] struct {
	LeftOnly  []K
	RightOnly []V
}
