package dbmanagerTest

import (
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerserver"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage/storagemanager"
)

func NewWithPeerServer(peerServer dbpeerserver.Interface) dbmanager.Manager {
	return dbmanager.New(peerServer, &DbServerMock{}, &DbStorageMock{})
}

func NewWithDbServer(dbServer dbserver.Interface) dbmanager.Manager {
	return dbmanager.New(&PeerServerMock{}, dbServer, &DbStorageMock{})
}

func NewWithDbStorage(dbStorage storagemanager.Interface) dbmanager.Manager {
	return dbmanager.New(&PeerServerMock{}, &DbServerMock{}, dbStorage)
}
