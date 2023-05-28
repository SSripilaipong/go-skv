package dbstoragerecord

import (
	"context"
	"go-skv/server/dbstorage"
)

type recordFactory struct {
	ctx          context.Context
	chBufferSize int
}

func NewFactory(ctx context.Context, channelBufferSize int) dbstorage.RecordFactory {
	return &recordFactory{
		ctx:          ctx,
		chBufferSize: channelBufferSize,
	}
}

func (r *recordFactory) New() dbstorage.DbRecord {
	ctx, ctxCancel := context.WithCancel(r.ctx)
	ch := make(chan any, r.chBufferSize)
	go recordMainLoop(ctx, ch)
	return newRecordInterface(ctx, ctxCancel, ch)
}
