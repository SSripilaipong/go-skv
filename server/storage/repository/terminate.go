package repository

import (
	"context"
	"go-skv/common/util/goutil"
	. "go-skv/server/storage/repository/message"
)

func terminate(ctx context.Context) func(Terminate) {
	return func(msg Terminate) {
		defer close(msg.Notify)

		goutil.SendWithinCtx(ctx, msg.Notify, struct{}{})
	}
}
