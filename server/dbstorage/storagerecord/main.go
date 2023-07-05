package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func NewFactory(channelBufferSize int) dbstoragecontract.Factory {
	return recordFactory{
		chBufferSize: channelBufferSize,
	}
}

type recordFactory struct {
	chBufferSize int
}

func (r recordFactory) New(ctx context.Context) dbstoragecontract.Record {
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
