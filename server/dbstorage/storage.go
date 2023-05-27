package dbstorage

import (
	"context"
)

func New(ch chan any) Interface {
	ctx, cancel := context.WithCancel(context.Background())
	return &storage{ch: ch, ctx: ctx, cancel: cancel, stopped: make(chan struct{})}
}

type storage struct {
	ch      chan any
	ctx     context.Context
	cancel  context.CancelFunc
	stopped chan struct{}
}

func (s *storage) Start() error {
	go func() {
		for {
			select {
			case <-s.ch:
			case <-s.ctx.Done():
				goto stop
			}
		}
	stop:
		s.stopped <- struct{}{}
	}()
	return nil
}

func (s *storage) Stop() error {
	s.cancel()
	<-s.stopped
	return nil
}
