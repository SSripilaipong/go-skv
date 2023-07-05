package storagerepository

import (
	"go-skv/server/dbstorage/dbstoragecontract"
)

func New(storageBufferSize int, recordFactory dbstoragecontract.Factory) dbstoragecontract.Storage {
	m := &manager{
		ch: make(chan command, storageBufferSize),

		stopped: make(chan struct{}),
	}
	go mainLoop(m.ch, m.stopped, recordFactory)
	return m
}
