package repository

import (
	"context"
	storageMessage "go-skv/server/storage/message"
)

func saveRecord(ctx context.Context) func(msg storageMessage.SaveRecord) {
	return func(msg storageMessage.SaveRecord) {
		select {
		case msg.ReplyTo <- storageMessage.Ack{Memo: msg.Memo}:
		case <-ctx.Done():
		}
	}
}
