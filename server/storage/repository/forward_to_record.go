package repository

import (
	"context"
	"fmt"
	storageMessage "go-skv/server/storage/message"
)

func forwardToRecord(ctx context.Context, records map[string]chan<- any) func(msg storageMessage.ForwardToRecord) {
	return func(msg storageMessage.ForwardToRecord) {
		if record, exists := records[""]; exists {
			fmt.Println("exists!!!")
			select {
			case record <- msg.Message:
			case <-ctx.Done():
			}
		}
	}
}
