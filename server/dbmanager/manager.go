package dbmanager

import "fmt"

type Manager interface {
	Start() error
}

func New(peerServer PeerServer, dbServer DbServer) Manager {
	return &manager{peerServer: peerServer, dbServer: dbServer}
}

type manager struct {
	peerServer PeerServer
	dbServer   DbServer
}

func (m *manager) Start() error {
	if err := m.peerServer.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	if err := m.dbServer.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	return nil
}
