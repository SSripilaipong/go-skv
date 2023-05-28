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
}

func newRecordInterface(ctx context.Context, ctxCancel context.CancelFunc, ch chan any) dbstorage.DbRecord {
	return &recordInterface{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		ch:        ch,
	}
}

func (r *recordInterface) SetValue(message dbstorage.SetValueMessage) error {
	if r.isContextEnded() {
		return dbstorage.RecordDestroyedError{}
	}

	r.ch <- message
	return nil
}

func (r *recordInterface) GetValue(dbstorage.GetValueMessage) error {
	if r.isContextEnded() {
		return dbstorage.RecordDestroyedError{}
	}
	return nil
}

func (r *recordInterface) Destroy() error {
	r.ctxCancel()
	return nil
}

func (r *recordInterface) isContextEnded() bool {
	_, isEnded := goutil.ReceiveNoBlock(r.ctx.Done())
	return isEnded
}
