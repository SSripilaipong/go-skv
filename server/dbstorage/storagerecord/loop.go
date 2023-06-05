package storagerecord

import (
	"context"
	"go-skv/util/goutil"
)

func RecordMainLoop(ctx context.Context, ch chan any, stopped chan struct{}) {
	var value string

	for {
		select {
		case raw := <-ch:
			if message, isSetMessage := raw.(SetValueMessage); isSetMessage {
				value = message.Value()
				goutil.PanicUnhandledError(message.Completed(SetValueResponse{}))
			} else if message, isGetMessage := raw.(GetValueMessage); isGetMessage {
				goutil.PanicUnhandledError(message.Completed(GetValueResponse{Value: goutil.Pointer(value)}))
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
