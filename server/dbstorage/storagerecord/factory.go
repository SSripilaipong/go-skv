package storagerecord

import (
	"context"
)

type recordFactory struct {
	chBufferSize int
}

func (r recordFactory) New(ctx context.Context) Interface {
	recordCtx, ctxCancel := context.WithCancel(ctx)
	ch := make(chan any, r.chBufferSize)
	stopped := make(chan struct{})

	go runRecordMainLoop(recordCtx, ch, stopped)
	return newRecordInteractor(recordCtx, ctxCancel, ch, stopped)
}
