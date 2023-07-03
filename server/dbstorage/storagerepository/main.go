package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordFactory storagerecord.Factory) (Interface, Interactor) {
	ch := make(chan any, storageBufferSize)
	return NewRepository(ch, recordFactory), NewInteractor(ch)
}

func NewRepository(ch chan any, recordFactory storagerecord.Factory) Interface {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	return &manager{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctxWithCancel,
		cancel:        cancel,

		stopped: make(chan struct{}),
	}
}

func NewInteractor(ch chan<- any) Interactor {
	return interactor{ch: ch}
}
