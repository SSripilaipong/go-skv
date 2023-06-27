package stop

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_close_context_used_by_peers(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(client),
	)
	goutil.PanicUnhandledError(connector.Start())
	ctxUsedByPeers, _ := goutil.ElementAt(client.ConnectToPeer_ctx_array, 0)

	goutil.PanicUnhandledError(connector.Stop())

	_, isClosed := goutil.ReceiveNoBlock(ctxUsedByPeers.Done())
	assert.True(t, isClosed)
}

func Test_should_not_close_context_used_by_peers_before_stopping(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(client),
	)
	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())
	ctxUsedByPeers, _ := goutil.ElementAt(client.ConnectToPeer_ctx_array, 0)

	_, isClosed := goutil.ReceiveNoBlock(ctxUsedByPeers.Done())
	assert.False(t, isClosed)
}
