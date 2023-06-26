package start

import (
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
	connector := peerconnectortest.NewWithAddressesAndClient([]string{"1.1.1.1:1111"}, client)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}

func Test_should_connect_to_next_peer_if_the_first_peer_cannot_be_connected(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{nil, &peerconnectortest.PeerMock{}},
		ConnectToPeer_Error_array:  []error{peerclientcontract.ConnectionError{}, nil},
	}
	connector := peerconnectortest.NewWithAddressesAndClient([]string{"1.1.1.1:1111", "2.2.2.2:2222"}, client)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111", "2.2.2.2:2222"}, client.ConnectToPeer_address_array)
}

func Test_should_not_connect_to_next_peer_if_the_first_peer_can_be_connected(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&peerconnectortest.PeerMock{}},
	}
	connector := peerconnectortest.NewWithAddressesAndClient([]string{"1.1.1.1:1111", "2.2.2.2:2222"}, client)

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
	connector := peerconnectortest.NewWithAddressesAndClientAndPeerRepo([]string{"1.1.1.1:1111"}, &peerconnectortest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{connectedPeer},
	}, peerRepo)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, connectedPeer, peerRepo.Save_peer)
}
