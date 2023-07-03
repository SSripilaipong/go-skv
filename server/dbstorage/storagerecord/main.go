package storagerecord

import "context"

func NewFactory(channelBufferSize int) Factory {
	return recordFactory{
		chBufferSize: channelBufferSize,
	}
}

type recordFactory struct {
	chBufferSize int
}

func (r recordFactory) New(ctx context.Context) Interface {
	recordCtx, ctxCancel := context.WithCancel(ctx)
	ch := make(chan command, r.chBufferSize)
	stopped := make(chan struct{})

	go runRecordMainLoop(recordCtx, ch, stopped)
	return recordInteractor{
		ctx:       recordCtx,
		ctxCancel: ctxCancel,
		ch:        ch,

		stopped: stopped,
	}
}
