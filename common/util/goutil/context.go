package goutil

import (
	"context"
	"time"
)

func NewContextWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func NewCancelledContext() context.Context {
	newCtx, cancel := context.WithCancel(context.Background())
	cancel()
	return newCtx
}
