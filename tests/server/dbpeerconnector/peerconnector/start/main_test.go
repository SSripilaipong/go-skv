package start

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests/server/dbpeerconnector/peerconnector/peerconnectortest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_try_to_connect_to_an_existing_peer(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&peerconnectortest.PeerMock{}},
	}
	connector := peerconnectortest.New(
		peerconnectortest.WithAddresses([]string{"1.1.1.1:1111"}),
		peerconnectortest.WithClient(client),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_connect_to_next_peer_if_the_first_peer_cannot_be_connected(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{nil, &peerconnectortest.PeerMock{}},
		ConnectToPeer_Error_array:  []error{peerclientcontract.ConnectionError{}, nil},
	}
	connector := peerconnectortest.New(
		peerconnectortest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		peerconnectortest.WithClient(client),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111", "2.2.2.2:2222"}, client.ConnectToPeer_address_array)
}

func Test_should_not_connect_to_next_peer_if_the_first_peer_can_be_connected(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&peerconnectortest.PeerMock{}},
	}
	connector := peerconnectortest.New(
		peerconnectortest.WithAddresses([]string{"1.1.1.1:1111", "2.2.2.2:2222"}),
		peerconnectortest.WithClient(client),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_not_panic_when_no_available_peer(t *testing.T) {
	connector := peerconnectortest.New()

	defer goutil.WillPanicUnhandledError(connector.Stop)

	assert.NotPanics(t, goutil.WillPanicUnhandledError(connector.Start))
}

func Test_should_save_connected_peer_to_repository(t *testing.T) {
	connectedPeer := &peerconnectortest.PeerMock{}
	peerRepo := &peerconnectortest.PeerRepositoryMock{}
	connector := peerconnectortest.New(
		peerconnectortest.WithNonEmptyAddresses(),
		peerconnectortest.WithClient(&peerconnectortest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{connectedPeer},
		}),
		peerconnectortest.WithPeerRepo(peerRepo),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, connectedPeer, peerRepo.Save_peer)
}

func Test_should_save_connected_peer_to_repository_with_its_address_as_its_name(t *testing.T) {
	peerRepo := &peerconnectortest.PeerRepositoryMock{}
	connector := peerconnectortest.New(
		peerconnectortest.WithAddresses([]string{"1.1.1.1:1111"}),
		peerconnectortest.WithClient(&peerconnectortest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&peerconnectortest.PeerMock{}},
		}),
		peerconnectortest.WithPeerRepo(peerRepo),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, "1.1.1.1:1111", peerRepo.Save_name)
}

func Test_should_use_global_context_to_save(t *testing.T) {
	peerRepo := &peerconnectortest.PeerRepositoryMock{}
	ctx := context.WithValue(context.Background(), "test", "this is the expected context")
	connector := peerconnectortest.New(
		peerconnectortest.WithContext(ctx),
		peerconnectortest.WithNonEmptyAddresses(),
		peerconnectortest.WithClient(&peerconnectortest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&peerconnectortest.PeerMock{}},
		}),
		peerconnectortest.WithPeerRepo(peerRepo),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, "this is the expected context", peerRepo.Save_ctx.Value("test"))
}

func Test_should_not_save_to_repository_when_cannot_to_connect_to_peer(t *testing.T) {
	peerRepo := &peerconnectortest.PeerRepositoryMock{}
	connector := peerconnectortest.New(
		peerconnectortest.WithPeerRepo(peerRepo),
	)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.False(t, peerRepo.Save_IsCalled)
}
