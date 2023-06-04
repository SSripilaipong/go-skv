package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
)

type recordFactory struct {
	ctx          context.Context
	chBufferSize int
}

func (r *recordFactory) New() storagemanager.DbRecord {
	ctx, ctxCancel := context.WithCancel(r.ctx)
	ch := make(chan any, r.chBufferSize)
	stopped := make(chan struct{})

	go recordMainLoop(ctx, ch, stopped)
	return newRecordInteractor(ctx, ctxCancel, ch, stopped)
}
