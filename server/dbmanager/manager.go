package dbmanager

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
)

type Manager interface {
	Start() error
	Stop() error
}

func New(peerServer peerconnectorcontract.Connector, dbServer dbserver.Interface, dbStorage dbstorage.Repository) Manager {
	ctx, cancelCtx := context.WithCancel(context.Background())
	cancelCtx()
	return manager{
		ctx:           ctx,
		peerConnector: peerServer,
		dbServer:      dbServer,
		dbStorage:     dbStorage,
	}
}

type manager struct {
	peerConnector peerconnectorcontract.Connector
	dbServer      dbserver.Interface
	dbStorage     dbstorage.Repository
	ctx           context.Context
}
