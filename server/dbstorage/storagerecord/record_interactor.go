package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/util/goutil"
)

type recordInteractor struct {
	ctx       context.Context
	ch        chan any
	ctxCancel context.CancelFunc

	stopped chan struct{}
}

func newRecordInteractor(ctx context.Context, ctxCancel context.CancelFunc, ch chan any, stopped chan struct{}) storagemanager.DbRecord {
	return &recordInteractor{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		ch:        ch,

		stopped: stopped,
	}
}

func (r *recordInteractor) SetValue(message storagemanager.SetValueMessage) error {
	if r.isContextEnded() {
		return storagemanager.RecordDestroyedError{}
	}
	r.ch <- message
	return nil
}

func (r *recordInteractor) GetValue(message storagemanager.GetValueMessage) error {
	if r.isContextEnded() {
		return storagemanager.RecordDestroyedError{}
	}
	r.ch <- message
	return nil
}

func (r *recordInteractor) Destroy() error {
	r.ctxCancel()
	<-r.stopped
	return nil
}

func (r *recordInteractor) isContextEnded() bool {
	_, isEnded := goutil.ReceiveNoBlock(r.ctx.Done())
	return isEnded
}
