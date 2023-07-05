package start

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbmanager/dbmanagertest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_start_peer_connector_with_context(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)

	_ = dbmanagertest.DoStart(mgr)

	assert.NotNil(t, peerConnector.Start_ctx)
}

func Test_should_start_db_server(t *testing.T) {
	dbServer := &dbmanagertest.DbServerMock{}
	mgr := dbmanagertest.NewWithDbServer(dbServer)

	_ = dbmanagertest.DoStart(mgr)

	assert.True(t, dbServer.Start_IsCalled)
}

func Test_should_start_db_storage_with_context(t *testing.T) {
	dbStorage := &servertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)

	_ = dbmanagertest.DoStart(mgr)

	assert.NotNil(t, dbStorage.Start_ctx)
}
