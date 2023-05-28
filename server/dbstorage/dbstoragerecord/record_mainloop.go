package dbstoragerecord

import (
	"context"
	"go-skv/server/dbstorage"
	goutil2 "go-skv/util/goutil"
)

func recordMainLoop(ctx context.Context, ch chan any, stopped chan struct{}) {
	var value string

	for {
		select {
		case raw := <-ch:
			if message, isSetMessage := raw.(dbstorage.SetValueMessage); isSetMessage {
				value = message.Value()
				goutil2.PanicUnhandledError(message.Completed(dbstorage.SetValueResponse{}))
			} else if message, isGetMessage := raw.(dbstorage.GetValueMessage); isGetMessage {
				goutil2.PanicUnhandledError(message.Completed(dbstorage.GetValueResponse{Value: goutil2.Pointer(value)}))
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
