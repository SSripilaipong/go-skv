package repository

import "context"

func loop(ctx context.Context, handleMessage func(any) bool, ch chan any) {
	defer close(ch)
	isDone := ctx.Done()

	for {
		select {
		case message := <-ch:
			if end := handleMessage(message); end {
				return
			}
		case <-isDone:
			return
		}
	}
}
