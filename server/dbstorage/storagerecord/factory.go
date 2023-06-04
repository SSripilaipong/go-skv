package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
)

type recordFactory struct {
	chBufferSize int
}

func (r *recordFactory) New(ctx context.Context) storagemanager.DbRecord {
	recordCtx, ctxCancel := context.WithCancel(ctx)
	ch := make(chan any, r.chBufferSize)
	stopped := make(chan struct{})

	go recordMainLoop(recordCtx, ch, stopped)
	return newRecordInteractor(recordCtx, ctxCancel, ch, stopped)
}
