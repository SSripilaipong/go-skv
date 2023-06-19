package start

import (
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerconnector"
	"go-skv/tests/server/dbpeerconnector/peerconnector/peerconnectortest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_try_to_connect_to_an_existing_peer(t *testing.T) {
	client := &peerconnectortest.PeerClientMock{}
	connector := peerconnector.New([]string{"1.1.1.1:1111"}, client)

	goutil.PanicUnhandledError(connector.Start())

	goutil.PanicUnhandledError(connector.Stop())
	assert.Equal(t, []string{"1.1.1.1:1111"}, client.ConnectToPeer_address_array)
}
