package dbmanager

import (
	"fmt"
	"go-skv/server/dbpeerserver"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
)

type Manager interface {
	Start() error
	Stop() error
}

func New(peerServer dbpeerserver.Interface, dbServer dbserver.Interface, dbStorage dbstorage.Interface) Manager {
	return &manager{
		peerServer: peerServer,
		dbServer:   dbServer,
		dbStorage:  dbStorage,
	}
}

type manager struct {
	peerServer dbpeerserver.Interface
	dbServer   dbserver.Interface
	dbStorage  dbstorage.Interface
}

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

func (m *manager) Stop() error {
	if err := m.dbServer.Stop(); err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	return nil
}
