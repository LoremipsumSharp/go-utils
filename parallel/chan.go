package parallel

import "context"

func Pipe[T any](
	ctx context.Context, in <-chan T, out chan<- T,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}
}

func PipeWithSize[T any](
	ctx context.Context, limit uint64, in <-chan T, out chan<- T,
) {
	var size uint64 = 0
	if limit == 0 {
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- v:
				size++
				if size >= limit {
					return
				}

			}
		}
	}
}
