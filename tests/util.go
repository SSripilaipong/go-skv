package tests

import (
	"context"
	"time"
)

func ExecuteWithTimeout(duration time.Duration, execute func(ctx context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	return execute(ctx)
}
