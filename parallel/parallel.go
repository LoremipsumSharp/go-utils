package parallel

import  (
	"github.com/hashicorp/go-multierror"
	"sync"
)



type empty struct{}
type Processor func(idx int) error



func ForEach(total int, n int, process Processor) error {
	semaphore := make(chan empty, n)
	errors := make([]error, total)

	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(i int) {
			semaphore <- empty{}
			errors[i] = process(i)
			<-semaphore
			wg.Done()
		}(i)
	}
	wg.Wait()
	return multierror.Append(nil, errors...).ErrorOrNil()
}