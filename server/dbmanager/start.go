package dbmanager

import "fmt"

func (m *manager) Start() error {
	if err := m.dbStorage.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	if err := m.peerServer.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	if err := m.dbServer.Start(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	return nil
}
