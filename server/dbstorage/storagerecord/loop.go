package storagerecord

import (
	"context"
)

func runRecordMainLoop(ctx context.Context, ch chan any, stopped chan struct{}) {
	var value string

	for {
		select {
		case raw := <-ch:
			if message, isSetMessage := raw.(setValueMessage); isSetMessage {
				value = message.value
				message.success(SetValueResponse{Value: message.value})
			} else if message, isGetMessage := raw.(getValueMessage); isGetMessage {
				message.success(GetValueResponse{Value: value})
			}
		case <-ctx.Done():
			goto stop
		}
	}
stop:
	stopped <- struct{}{}
}
