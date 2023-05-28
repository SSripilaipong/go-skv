package dbstoragerecord

import (
	"context"
	"go-skv/goutil"
	"go-skv/server/dbstorage"
)

type recordInterface struct {
	ctx context.Context
	ch  chan any
}

func newRecordInterface(ctx context.Context, ch chan any) dbstorage.DbRecord {
	return &recordInterface{
		ctx: ctx,
		ch:  ch,
	}
}

func (r *recordInterface) SetValue(message dbstorage.SetValueMessage) error {
	if _, isDone := goutil.ReceiveNoBlock(r.ctx.Done()); isDone {
		return dbstorage.RecordDestroyedError{}
	}

	r.ch <- message
	return nil
}

func (r *recordInterface) GetValue(dbstorage.GetValueMessage) error {
	return nil
}

func (r *recordInterface) Destroy() error {
	return nil
}
