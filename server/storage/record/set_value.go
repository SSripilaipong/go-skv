package record

import (
	"context"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/record/message"
)

func setValue(ctx context.Context, value *string) func(msg SetValue) {
	return func(msg SetValue) {
		defer close(msg.ReplyTo)

		*value = msg.Value
		goutil.SendWithinCtx[any](ctx, msg.ReplyTo, Ack{Memo: msg.Memo})
	}
}
