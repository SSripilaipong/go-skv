package start

import (
	"github.com/stretchr/testify/assert"
	dbmanagerTest "go-skv/tests/server/dbmanager"
	"testing"
)

func Test_should_start_peer_server(t *testing.T) {
	peerServer := &dbmanagerTest.PeerServerMock{}
	mgr := dbmanagerTest.NewWithPeerServer(peerServer)

	_ = doStart(mgr)

	assert.True(t, peerServer.Start_IsCalled)
}

func Test_should_start_db_server(t *testing.T) {
	dbServer := &dbmanagerTest.DbServerMock{}
	mgr := dbmanagerTest.NewWithDbServer(dbServer)

	_ = doStart(mgr)

	assert.True(t, dbServer.Start_IsCalled)
}
