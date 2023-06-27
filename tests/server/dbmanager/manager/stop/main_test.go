package stop

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbmanager/dbmanagertest"
	"testing"
)

func Test_should_stop_db_server(t *testing.T) {
	dbServer := &dbmanagertest.DbServerMock{}
	mgr := dbmanagertest.NewWithDbServer(dbServer)

	_ = doStop(mgr)

	assert.True(t, dbServer.Stop_IsCalled)
}

func Test_should_stop_db_storage(t *testing.T) {
	dbStorage := &dbmanagertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)

	_ = doStop(mgr)

	assert.True(t, dbStorage.Stop_IsCalled)
}

func Test_should_join_peer_connector(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)

	_ = doStop(mgr)

	assert.True(t, peerConnector.Join_IsCalled)
}
