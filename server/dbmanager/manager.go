package dbmanager

import "fmt"

type Manager interface {
	Start() error
}

func New(peerServer PeerServer) Manager {
	return &manager{peerServer: peerServer}
}

type manager struct {
	peerServer PeerServer
}

func (m *manager) Start() error {
	if err := m.peerServer.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	return nil
}
