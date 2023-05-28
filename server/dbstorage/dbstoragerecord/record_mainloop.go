package dbstoragerecord

import (
	"context"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
)

func recordMainLoop(ctx context.Context, ch chan any) {
	select {
	case raw := <-ch:
		message := goutil.CastOrPanic[dbstorage.SetValueMessage](raw)
		goutil.PanicUnhandledError(message.Completed(dbstorage.SetValueResponse{}))
	case <-ctx.Done():
	}
}
