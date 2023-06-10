package repositoryroutine

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func New(ch chan any, recordFactory storagerecord.Factory) Interface {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	return &manager{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctxWithCancel,
		cancel:        cancel,

		stopped: make(chan struct{}),
		records: make(map[string]storagerecord.Interface),
	}
}
