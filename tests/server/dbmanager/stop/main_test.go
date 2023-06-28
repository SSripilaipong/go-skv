package stop

import (
	"github.com/stretchr/testify/assert"
	"go-skv/tests/server/dbmanager/dbmanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_stop_db_server(t *testing.T) {
	dbServer := &dbmanagertest.DbServerMock{}
	mgr := dbmanagertest.NewWithDbServer(dbServer)

	_ = dbmanagertest.DoStop(mgr)

	assert.True(t, dbServer.Stop_IsCalled)
}

func Test_should_stop_db_storage(t *testing.T) {
	dbStorage := &dbmanagertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)

	_ = dbmanagertest.DoStop(mgr)

	assert.True(t, dbStorage.Stop_IsCalled)
}

func Test_should_join_peer_connector(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)

	_ = dbmanagertest.DoStop(mgr)

	assert.True(t, peerConnector.Join_IsCalled)
}

func Test_should_close_context_for_peer_connector(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)
	_ = dbmanagertest.DoStart(mgr)
	ctx := peerConnector.Start_ctx

	_ = dbmanagertest.DoStop(mgr)

	_, isClosed := goutil.ReceiveNoBlock(ctx.Done())
	assert.True(t, isClosed)
}
