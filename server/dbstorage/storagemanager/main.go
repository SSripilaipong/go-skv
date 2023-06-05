package storagemanager

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerecordfactory"
)

func New(ch chan any, recordFactory storagerecordfactory.Interface) Interface {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	return &manager{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctxWithCancel,
		cancel:        cancel,

		stopped: make(chan struct{}),
		records: make(map[string]storagerecord.DbRecord),
	}
}
