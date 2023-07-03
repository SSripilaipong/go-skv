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
}

func (m *manager) Start() error {
	go mainLoop(m.ctx, m.ch, m.stopped, m.recordFactory)
	return nil
}

func (m *manager) Stop() error {
	m.cancel()
	<-m.stopped
	return nil
}
