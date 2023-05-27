package dbmanager

import "fmt"

func (m *manager) Stop() error {
	if err := m.dbServer.Stop(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	if err := m.dbStorage.Stop(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	return nil
}
