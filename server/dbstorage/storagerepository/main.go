package storagerepository

import (
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordFactory storagerecord.Factory) Interface {
	ch := make(chan any, storageBufferSize)
	m := &manager{
		ch: ch,

		stopped: make(chan struct{}),
	}
	go mainLoop(m.ch, m.stopped, recordFactory)
	return m
}
