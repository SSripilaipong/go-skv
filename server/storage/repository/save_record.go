package repository

import (
	"context"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/repository/message"
)

func saveRecord(ctx context.Context, records map[string]chan<- any) func(msg SaveRecord) {
	return func(msg SaveRecord) {
		defer close(msg.ReplyTo)

		records[msg.Key] = msg.Channel

		goutil.SendWithinCtx[any](ctx, msg.ReplyTo, Ack{Memo: msg.Memo})
	}
}
