package storagemanager

import (
	"context"
)

type manager struct {
	ch            chan any
	recordFactory RecordFactory
	ctx           context.Context
	cancel        context.CancelFunc

	stopped chan struct{}
	records map[string]DbRecord
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
