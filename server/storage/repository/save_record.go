package repository

import (
	"context"
	"fmt"
	storageMessage "go-skv/server/storage/message"
)

func saveRecord(ctx context.Context, records map[string]chan<- any) func(msg storageMessage.SaveRecord) {
	return func(msg storageMessage.SaveRecord) {
		fmt.Println("Saving")
		records[""] = msg.Channel

		select {
		case msg.ReplyTo <- storageMessage.Ack{Memo: msg.Memo}:
		case <-ctx.Done():
		}
	}
}
