package semaphore

import (
	"fmt"
	"sync"
	"testing"
)

func Test_semaphore_Acquire(t *testing.T) {
	se := New(10)
	wg := sync.WaitGroup{}
	wg.Add(11)
	for i := 0; i < 11; i++ {
		go func() {
			defer wg.Done()
			acquire := se.TryAcquire(1)
			fmt.Println(acquire)
		}()
	}
	wg.Wait()
}
