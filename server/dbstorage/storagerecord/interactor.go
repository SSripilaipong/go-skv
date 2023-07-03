package storagerecord

import (
	"context"
	"go-skv/common/commoncontract"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type recordInteractor struct {
	ctx       context.Context
	ch        chan any
	ctxCancel context.CancelFunc

	stopped chan struct{}
}

func newRecordInteractor(ctx context.Context, ctxCancel context.CancelFunc, ch chan any, stopped chan struct{}) Interface {
	return recordInteractor{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		ch:        ch,

		stopped: stopped,
	}
}

func (r recordInteractor) SetValue(ctx context.Context, value string, success func(response SetValueResponse)) error {
	if r.isContextEnded() {
		return dbstoragecontract.RecordDestroyedError{}
	}
	select {
	case r.ch <- setValueMessage{value: value, success: success}:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func (r recordInteractor) GetValue(ctx context.Context, success func(response GetValueResponse)) error {
	if r.isContextEnded() {
		return dbstoragecontract.RecordDestroyedError{}
	}
	select {
	case r.ch <- getValueMessage{success: success}:
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
	return nil
}

func (r recordInteractor) Destroy() error {
	r.ctxCancel()
	<-r.stopped
	return nil
}

func (r recordInteractor) isContextEnded() bool {
	_, isEnded := goutil.ReceiveNoBlock(r.ctx.Done())
	return isEnded
}
