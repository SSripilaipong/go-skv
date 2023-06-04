package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/util/goutil"
)

func recordMainLoop(ctx context.Context, ch chan any, stopped chan struct{}) {
	var value string

	for {
		select {
		case raw := <-ch:
			if message, isSetMessage := raw.(storagemanager.SetValueMessage); isSetMessage {
				value = message.Value()
				goutil.PanicUnhandledError(message.Completed(storagemanager.SetValueResponse{}))
			} else if message, isGetMessage := raw.(storagemanager.GetValueMessage); isGetMessage {
				goutil.PanicUnhandledError(message.Completed(storagemanager.GetValueResponse{Value: goutil.Pointer(value)}))
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
