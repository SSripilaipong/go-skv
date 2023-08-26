package repository

import (
	"context"
	"go-skv/common/util/goutil"
	storageMessage "go-skv/server/storage/message"
)

func saveRecord(ctx context.Context, records map[string]chan<- any) func(msg storageMessage.SaveRecord) {
	return func(msg storageMessage.SaveRecord) {
		defer close(msg.ReplyTo)

		records[msg.Key] = msg.Channel

		goutil.SendWithinCtx[any](ctx, msg.ReplyTo, storageMessage.Ack{Memo: msg.Memo})
	}
}
