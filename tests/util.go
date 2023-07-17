package tests

import (
	"context"
	"go-skv/common/util/goutil"
	"sync"
	"time"
)

func ExecuteWithTimeout(duration time.Duration, execute func(ctx context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	return execute(ctx)
}

func SubContextScope(parent context.Context, f func(ctx context.Context)) {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	f(ctx)
}

func ContextScope(f func(ctx context.Context)) {
	SubContextScope(context.Background(), f)
}

func WaitScope(f func(wd *sync.WaitGroup)) {
	var wg sync.WaitGroup
	f(&wg)
	wg.Wait()
}

func NewClosedContext() context.Context {
	newCtx, cancel := context.WithCancel(context.Background())
	cancel()
	return newCtx
}

func MockWaitUntilCalledNthTimes(wgp **sync.WaitGroup, n int, timeout time.Duration, f func()) bool {
	defer func() {
		*wgp = nil
	}()
	*wgp = &sync.WaitGroup{}
	(*wgp).Add(n)

	f()

	return goutil.WaitWithTimeout(*wgp, timeout)
}

func ContextWithTimeout(timeout time.Duration) context.Context {
	//goland:noinspection ALL
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}
