package start

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-skv/common/test"
	"go-skv/common/util/goutil"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"testing"
)

func Test_should_not_start_server_if_connecting_to_existing_peers_panics(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{ConnectToPeer_Panics_array: []error{errors.New("boom")}}
	server := &connectormanagertest.PeerServerMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(client),
		connectormanagertest.WithServer(server),
	)

	assert.Panics(t, func() {
		test.ContextScope(func(ctx context.Context) {
			goutil.PanicUnhandledError(connector.Start(ctx))
		})
	})
	assert.False(t, server.Start_IsCalled)
}
