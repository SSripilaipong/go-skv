package storagerepository

import (
	"context"
)

type Manager struct {
	ch     chan any
	ctx    context.Context
	cancel context.CancelFunc

	stopped chan struct{}
}

func (m *Manager) Start(context.Context) error {
	return nil
}

func (m *Manager) Join() error {
	m.cancel()
	<-m.stopped
	return nil
}
