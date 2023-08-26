package repository

import (
	"context"
	"go-skv/common/util/goutil"
	storageMessage "go-skv/server/storage/message"
)

func forwardToRecord(ctx context.Context, records map[string]chan<- any) func(msg storageMessage.ForwardToRecord) {
	return func(msg storageMessage.ForwardToRecord) {
		defer close(msg.ReplyTo)

		record, exists := records[msg.Key]
		if !exists {
			goutil.SendWithinCtx[any](ctx, msg.ReplyTo, storageMessage.RecordNotFound{Key: msg.Key, Memo: msg.Memo})
			return
		}

		goutil.SendWithinCtx[any](ctx, record, msg.Message)
	}
}
