package join

import (
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/goutil"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"testing"
)

func Test_should_wait_for_all_peers_in_client_to_be_disconnected(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(client),
	)

	goutil.PanicUnhandledError(connector.Join())

	assert.True(t, client.WaitForAllToBeDisconnected_IsCalled)
}
