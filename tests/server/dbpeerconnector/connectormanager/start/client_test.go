package start

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_try_to_connect_to_an_existing_peer(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&dbpeerconnectortest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111"}),
		connectormanagertest.WithClient(client),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_connect_to_next_peer_if_the_first_peer_cannot_be_connected(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{nil, &dbpeerconnectortest.PeerMock{}},
		ConnectToPeer_Error_array:  []error{peerclientcontract.ConnectionError{}, nil},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		connectormanagertest.WithClient(client),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, []string{"1.1.1.1:1111", "2.2.2.2:2222"}, client.ConnectToPeer_address_array)
}

func Test_should_not_connect_to_next_peer_if_the_first_peer_can_be_connected(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&dbpeerconnectortest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		connectormanagertest.WithClient(client),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_not_panic_when_no_available_peer(t *testing.T) {
	connector := connectormanagertest.New()

	tests.ContextScope(func(ctx context.Context) {
		assert.NotPanics(t, goutil.WillPanicUnhandledError(func() error { return connector.Start(ctx) }))
	})
}

func Test_should_save_connected_peer_to_repository(t *testing.T) {
	connectedPeer := &dbpeerconnectortest.PeerMock{}
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{connectedPeer},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, connectedPeer, peerRepo.Save_peer)
}

func Test_should_save_connected_peer_to_repository_with_its_address_as_its_name(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111"}),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&dbpeerconnectortest.PeerMock{}},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, "1.1.1.1:1111", peerRepo.Save_name)
}

func Test_should_not_save_to_repository_when_cannot_to_connect_to_peer(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.False(t, peerRepo.Save_IsCalled)
}
