package record

import (
	"context"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/record/message"
)

func getValue(ctx context.Context, value *string) func(msg GetValue) {
	return func(msg GetValue) {
		goutil.SendWithinCtx[any](ctx, msg.ReplyTo, Value{Memo: msg.Memo, Value: *value})
	}
}
