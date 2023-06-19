package dbmanagertest

import (
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
)

func NewWithPeerConnector(peerServer peerconnectorcontract.Connector) dbmanager.Manager {
	return dbmanager.New(peerServer, &DbServerMock{}, &DbStorageMock{})
}

func NewWithDbServer(dbServer dbserver.Interface) dbmanager.Manager {
	return dbmanager.New(&PeerConnectorMock{}, dbServer, &DbStorageMock{})
}

func NewWithDbStorage(dbStorage dbstorage.Repository) dbmanager.Manager {
	return dbmanager.New(&PeerConnectorMock{}, &DbServerMock{}, dbStorage)
}
