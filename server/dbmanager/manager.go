package dbmanager

import (
	"go-skv/server/dbpeerserver"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage/storagemanager"
)

type Manager interface {
	Start() error
	Stop() error
}

func New(peerServer dbpeerserver.Interface, dbServer dbserver.Interface, dbStorage storagemanager.Interface) Manager {
	return &manager{
		peerServer: peerServer,
		dbServer:   dbServer,
		dbStorage:  dbStorage,
	}
}

type manager struct {
	peerServer dbpeerserver.Interface
	dbServer   dbserver.Interface
	dbStorage  storagemanager.Interface
}
