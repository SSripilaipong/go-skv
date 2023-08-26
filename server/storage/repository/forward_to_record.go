package repository

import (
	"context"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/repository/message"
)

func forwardToRecord(ctx context.Context, records map[string]chan<- any) func(msg ForwardToRecord) {
	return func(msg ForwardToRecord) {
		defer close(msg.ReplyTo)

		record, exists := records[msg.Key]
		if !exists {
			goutil.SendWithinCtx[any](ctx, msg.ReplyTo, RecordNotFound{Key: msg.Key, Memo: msg.Memo})
			return
		}

		goutil.SendWithinCtx[any](ctx, record, msg.Message)
	}
}
