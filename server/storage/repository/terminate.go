package repository

import (
	"context"
	storageMessage "go-skv/server/storage/message"
)

func terminate(ctx context.Context) func(storageMessage.Terminate) {
	return func(msg storageMessage.Terminate) {
		defer close(msg.Notify)

		select {
		case msg.Notify <- struct{}{}:
		case <-ctx.Done():
		}
	}
}
