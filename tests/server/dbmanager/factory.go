package dbmanagerTest

import (
	"go-skv/server/dbmanager"
)

func NewWithPeerServer(peerServer dbmanager.PeerServer) dbmanager.Manager {
	return dbmanager.New(peerServer, &DbServerMock{})
}

func NewWithDbServer(dbServer dbmanager.DbServer) dbmanager.Manager {
	return dbmanager.New(&PeerServerMock{}, dbServer)
}
