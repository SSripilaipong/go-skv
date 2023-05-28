package dbstoragerecord

import (
	"context"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
)

func recordMainLoop(ctx context.Context, ch chan any, stopped chan struct{}) {
	var value string

	for {
		select {
		case raw := <-ch:
			if message, isSetMessage := raw.(dbstorage.SetValueMessage); isSetMessage {
				value = message.Value()
				goutil.PanicUnhandledError(message.Completed(dbstorage.SetValueResponse{}))
			} else if message, isGetMessage := raw.(dbstorage.GetValueMessage); isGetMessage {
				goutil.PanicUnhandledError(message.Completed(dbstorage.GetValueResponse{Value: goutil.Pointer(value)}))
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
