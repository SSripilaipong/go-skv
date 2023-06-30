package connect_to_peer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/tests/server/dbpeerconnector/peerclient/peerclientmanager/peerclientmanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_create_new_client_side_peer_with_context(t *testing.T) {
	peerFactory := &peerclientmanagertest.PeerFactoryMock{}
	manager := peerclientmanagertest.New(
		peerclientmanagertest.WithPeerFactory(peerFactory),
	)
	tests.ContextScope(func(ctx context.Context) {
		ctx = context.WithValue(ctx, "test", "1234567890")

		_, err := manager.ConnectToPeer(ctx, "1.1.1.1:1234")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "1234567890", peerFactory.New_ctx.Value("test"))
}

func Test_should_connect_via_gateway_connector_with_context(t *testing.T) {
	gatewayConnector := &peerclientmanagertest.GatewayConnectorMock{}
	manager := peerclientmanagertest.New(
		peerclientmanagertest.WithGatewayConnector(gatewayConnector),
	)
	tests.ContextScope(func(ctx context.Context) {
		ctx = context.WithValue(ctx, "test", "yee")

		_, err := manager.ConnectToPeer(ctx, "1.2.3.4:1234")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "yee", gatewayConnector.ConnectTo_ctx.Value("test"))
}

func Test_should_connect_via_gateway_connector_with_address(t *testing.T) {
	gatewayConnector := &peerclientmanagertest.GatewayConnectorMock{}
	manager := peerclientmanagertest.New(
		peerclientmanagertest.WithGatewayConnector(gatewayConnector),
	)
	tests.ContextScope(func(ctx context.Context) {
		_, err := manager.ConnectToPeer(ctx, "1.2.3.4:1234")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "1.2.3.4:1234", gatewayConnector.ConnectTo_address)
}

func Test_should_connect_via_gateway_connector_with_created_peer(t *testing.T) {
	createdPeer := &dbpeerconnectortest.PeerMock{}
	peerFactory := &peerclientmanagertest.PeerFactoryMock{New_Return: createdPeer}
	gatewayConnector := &peerclientmanagertest.GatewayConnectorMock{}
	manager := peerclientmanagertest.New(
		peerclientmanagertest.WithPeerFactory(peerFactory),
		peerclientmanagertest.WithGatewayConnector(gatewayConnector),
	)
	tests.ContextScope(func(ctx context.Context) {
		_, err := manager.ConnectToPeer(ctx, "1.2.3.4:1234")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, createdPeer, gatewayConnector.ConnectTo_peer)
}
