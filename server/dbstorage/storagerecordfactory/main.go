package storagerecordfactory

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory(channelBufferSize int) Interface {
	return recordFactory{
		chBufferSize: channelBufferSize,
	}
}

type recordFactory struct {
	chBufferSize int
}

func (r recordFactory) New(ctx context.Context) storagerecord.DbRecord {
	recordCtx, ctxCancel := context.WithCancel(ctx)
	ch := make(chan any, r.chBufferSize)
	stopped := make(chan struct{})

	go storagerecord.RecordMainLoop(recordCtx, ch, stopped)
	return storagerecord.NewRecordInteractor(recordCtx, ctxCancel, ch, stopped)
}
