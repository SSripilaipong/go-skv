package record

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/storage/record/message"
)

func getValue(ctx context.Context, value *string) func(msg message.GetValue) {
	return func(msg message.GetValue) {
		goutil.SendWithinCtx[any](ctx, msg.ReplyTo, message.Value{Value: *value})
	}
}
