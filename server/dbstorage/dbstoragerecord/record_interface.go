package dbstoragerecord

import (
	"context"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
)

type recordInterface struct {
	ctx       context.Context
	ch        chan any
	ctxCancel context.CancelFunc

	stopped chan struct{}
}

func newRecordInterface(ctx context.Context, ctxCancel context.CancelFunc, ch chan any, stopped chan struct{}) dbstorage.DbRecord {
	return &recordInterface{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		ch:        ch,

		stopped: stopped,
	}
}

func (r *recordInterface) SetValue(message dbstorage.SetValueMessage) error {
	if r.isContextEnded() {
		return dbstorage.RecordDestroyedError{}
	}
	r.ch <- message
	return nil
}

func (r *recordInterface) GetValue(message dbstorage.GetValueMessage) error {
	if r.isContextEnded() {
		return dbstorage.RecordDestroyedError{}
	}
	r.ch <- message
	return nil
}

func (r *recordInterface) Destroy() error {
	r.ctxCancel()
	<-r.stopped
	return nil
}

func (r *recordInterface) isContextEnded() bool {
	_, isEnded := goutil.ReceiveNoBlock(r.ctx.Done())
	return isEnded
}
