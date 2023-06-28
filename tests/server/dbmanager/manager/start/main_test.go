package start

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbmanager/dbmanagertest"
	"testing"
)

func Test_should_start_peer_connector_with_context(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)

	_ = doStart(mgr)

	assert.NotNil(t, peerConnector.Start_ctx)
}

func Test_should_start_db_server(t *testing.T) {
	dbServer := &dbmanagertest.DbServerMock{}
	mgr := dbmanagertest.NewWithDbServer(dbServer)

	_ = doStart(mgr)

	assert.True(t, dbServer.Start_IsCalled)
}

func Test_should_start_db_storage(t *testing.T) {
	dbStorage := &dbmanagertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)

	_ = doStart(mgr)

	assert.True(t, dbStorage.Start_IsCalled)
}
