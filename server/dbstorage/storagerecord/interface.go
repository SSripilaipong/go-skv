package storagerecord

import (
	"context"
	"go-skv/util/goutil"
)

type DbRecord interface {
	SetValue(SetValueMessage) error
	GetValue(GetValueMessage) error
	Destroy() error
}

type recordInteractor struct {
	ctx       context.Context
	ch        chan any
	ctxCancel context.CancelFunc

	stopped chan struct{}
}

func NewRecordInteractor(ctx context.Context, ctxCancel context.CancelFunc, ch chan any, stopped chan struct{}) DbRecord {
	return recordInteractor{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		ch:        ch,

		stopped: stopped,
	}
}

func (r recordInteractor) SetValue(message SetValueMessage) error {
	if r.isContextEnded() {
		return RecordDestroyedError{}
	}
	r.ch <- message
	return nil
}

func (r recordInteractor) GetValue(message GetValueMessage) error {
	if r.isContextEnded() {
		return RecordDestroyedError{}
	}
	r.ch <- message
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
