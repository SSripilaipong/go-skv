package repository

import (
	"context"
	"go-skv/common/util/goutil"
	storageMessage "go-skv/server/storage/message"
)

func terminate(ctx context.Context) func(storageMessage.Terminate) {
	return func(msg storageMessage.Terminate) {
		defer close(msg.Notify)

		goutil.SendWithinCtx(ctx, msg.Notify, struct{}{})
	}
}
