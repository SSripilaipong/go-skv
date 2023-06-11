package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type manager struct {
	ch            chan any
	recordFactory storagerecord.Factory
	ctx           context.Context
	cancel        context.CancelFunc

	stopped chan struct{}
	records map[string]storagerecord.Interface
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
