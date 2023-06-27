package tests

import (
	"context"
	"sync"
	"time"
)

func ExecuteWithTimeout(duration time.Duration, execute func(ctx context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	return execute(ctx)
}

func ContextScope(f func(ctx context.Context)) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	f(ctx)
}

func WaitScope(f func(wd *sync.WaitGroup)) {
	var wg sync.WaitGroup
	f(&wg)
	wg.Wait()
}
