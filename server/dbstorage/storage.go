package dbstorage

import (
	"context"
)

func New(ch chan any, recordFactory RecordFactory) Interface {
	ctx, cancel := context.WithCancel(context.Background())
	return &storage{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctx,
		cancel:        cancel,
		stopped:       make(chan struct{}),
	}
}

type storage struct {
	ch            chan any
	recordFactory RecordFactory
	ctx           context.Context
	cancel        context.CancelFunc
	stopped       chan struct{}
}

func (s *storage) Start() error {
	go s.mainLoop()
	return nil
}

func (s *storage) Stop() error {
	s.cancel()
	<-s.stopped
	return nil
}
