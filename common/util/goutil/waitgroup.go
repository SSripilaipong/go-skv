package goutil

import (
	"sync"
	"time"
)

func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	completed := make(chan struct{})
	go func() {
		wg.Wait()
		completed <- struct{}{}
	}()
	select {
	case <-completed:
		return true
	case <-time.After(timeout):
		return false
	}
}
