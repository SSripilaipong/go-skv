package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordFactory storagerecord.Factory) (Interface, Interactor) {
	ch := make(chan any, storageBufferSize)
	r := NewRepository(ch, recordFactory)
	return r, r
}

func NewRepository(ch chan any, recordFactory storagerecord.Factory) *Manager {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	m := &Manager{
		ch:     ch,
		ctx:    ctxWithCancel,
		cancel: cancel,

		stopped: make(chan struct{}),
	}
	go mainLoop(m.ctx, m.ch, m.stopped, recordFactory)

	return m
}

func NewInteractor(ch chan any) Interactor {
	r := NewRepository(ch, nil)
	return r
}
