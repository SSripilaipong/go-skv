package dbstorage

import (
	"go-skv/goutil"
	"time"
)

func New(ch chan any) Interface {
	return &storage{ch: ch}
}

type storage struct {
	ch chan any
}

func (s *storage) Start() error {
	go func() {
		goutil.ReceiveWithTimeout(s.ch, time.Second)
	}()
	return nil
}

func (s *storage) Stop() error {
	return nil
}
