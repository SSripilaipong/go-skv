package stop

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/tests/server/dbmanager/dbmanagertest"
	"go-skv/tests/server/servertest"
	"testing"
)

func Test_should_stop_db_server(t *testing.T) {
	dbServer := &dbmanagertest.DbServerMock{}
	mgr := dbmanagertest.NewWithDbServer(dbServer)

	_ = dbmanagertest.DoStop(mgr)

	assert.True(t, dbServer.Stop_IsCalled)
}

func Test_should_join_db_storage(t *testing.T) {
	dbStorage := &servertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)

	_ = dbmanagertest.DoStop(mgr)

	assert.True(t, dbStorage.Join_IsCalled)
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

func Test_should_not_close_context_for_peer_connector_before_stopping(t *testing.T) {
	peerConnector := &dbmanagertest.PeerConnectorMock{}
	mgr := dbmanagertest.NewWithPeerConnector(peerConnector)
	_ = dbmanagertest.DoStart(mgr)

	_, isClosed := goutil.ReceiveNoBlock(peerConnector.Start_ctx.Done())
	assert.False(t, isClosed)
}

func Test_should_close_context_for_storage(t *testing.T) {
	dbStorage := &servertest.DbStorageMock{}
	mgr := dbmanagertest.NewWithDbStorage(dbStorage)
	_ = dbmanagertest.DoStart(mgr)
	ctx := dbStorage.Start_ctx

	_ = dbmanagertest.DoStop(mgr)

	_, isClosed := goutil.ReceiveNoBlock(ctx.Done())
	assert.True(t, isClosed)
}
