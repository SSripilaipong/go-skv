package dbmanager

import (
	"go-skv/server/dbpeerconnector"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
)

type Manager interface {
	Start() error
	Stop() error
}

func New(peerServer dbpeerconnector.Interface, dbServer dbserver.Interface, dbStorage dbstorage.Repository) Manager {
	return &manager{
		peerConnector: peerServer,
		dbServer:      dbServer,
		dbStorage:     dbStorage,
	}
}

type manager struct {
	peerConnector dbpeerconnector.Interface
	dbServer      dbserver.Interface
	dbStorage     dbstorage.Repository
}
