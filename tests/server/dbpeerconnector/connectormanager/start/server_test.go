package start

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"go-skv/util/goutil"
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
		tests.ContextScope(func(ctx context.Context) {
			goutil.PanicUnhandledError(connector.Start(ctx))
		})
	})
	assert.False(t, server.Start_IsCalled)
}
