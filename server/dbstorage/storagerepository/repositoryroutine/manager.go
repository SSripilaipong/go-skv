package repositoryroutine

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerecordfactory"
)

type manager struct {
	ch            chan any
	recordFactory storagerecordfactory.Interface
	ctx           context.Context
	cancel        context.CancelFunc

	stopped chan struct{}
	records map[string]storagerecord.DbRecord
}

func (m *manager) Start() error {
	go m.mainLoop()
	return nil
}

func (m *manager) Stop() error {
	m.cancel()
	<-m.stopped
	return nil
}
