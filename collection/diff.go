package collection

type diffResult[K any, V any] struct {
	LeftOnly  []K
	RightOnly []V
}
