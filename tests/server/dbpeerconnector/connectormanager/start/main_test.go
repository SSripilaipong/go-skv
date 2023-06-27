package start

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_try_to_connect_to_an_existing_peer(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111"}),
		connectormanagertest.WithClient(client),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_connect_to_peer_with_global_context(t *testing.T) {
	ctx := context.WithValue(context.Background(), "test", "this is my context")
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithContext(ctx),
		connectormanagertest.WithClient(client),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, []string{"this is my context"}, goutil.Map(client.ConnectToPeer_ctx_array, func(c context.Context) string {
		return goutil.May(c, func(t context.Context) string { return c.Value("test").(string) })
	}))
}

func Test_should_connect_to_next_peer_if_the_first_peer_cannot_be_connected(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{nil, &connectormanagertest.PeerMock{}},
		ConnectToPeer_Error_array:  []error{peerclientcontract.ConnectionError{}, nil},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		connectormanagertest.WithClient(client),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, []string{"1.1.1.1:1111", "2.2.2.2:2222"}, client.ConnectToPeer_address_array)
}

func Test_should_not_connect_to_next_peer_if_the_first_peer_can_be_connected(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		connectormanagertest.WithClient(client),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_not_panic_when_no_available_peer(t *testing.T) {
	connector := connectormanagertest.New()

	defer goutil.WillPanicUnhandledError(connector.Stop)

	assert.NotPanics(t, goutil.WillPanicUnhandledError(connector.Start))
}

func Test_should_save_connected_peer_to_repository(t *testing.T) {
	connectedPeer := &connectormanagertest.PeerMock{}
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{connectedPeer},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, connectedPeer, peerRepo.Save_peer)
}

func Test_should_save_connected_peer_to_repository_with_its_address_as_its_name(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithAddresses([]string{"1.1.1.1:1111"}),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, "1.1.1.1:1111", peerRepo.Save_name)
}

func Test_should_use_global_context_to_save(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	ctx := context.WithValue(context.Background(), "test", "this is the expected context")
	connector := connectormanagertest.New(
		connectormanagertest.WithContext(ctx),
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.Equal(t, "this is the expected context", peerRepo.Save_ctx.Value("test"))
}

func Test_should_not_save_to_repository_when_cannot_to_connect_to_peer(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	defer goutil.WillPanicUnhandledError(connector.Stop)
	goutil.PanicUnhandledError(connector.Start())

	assert.False(t, peerRepo.Save_IsCalled)
}
