package start

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/connectormanager/connectormanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_connect_to_peer_with_global_context(t *testing.T) {
	client := &connectormanagertest.PeerClientMock{
		ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
	}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(client),
	)

	tests.ContextScope(func(ctx context.Context) {
		ctx = context.WithValue(ctx, "test", "this is my context")
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, []string{"this is my context"}, goutil.Map(client.ConnectToPeer_ctx_array, func(c context.Context) string {
		return goutil.May(c, func(t context.Context) string { return c.Value("test").(string) })
	}))
}

func Test_should_use_global_context_to_save(t *testing.T) {
	peerRepo := &connectormanagertest.PeerRepositoryMock{}
	connector := connectormanagertest.New(
		connectormanagertest.WithNonEmptyAddresses(),
		connectormanagertest.WithClient(&connectormanagertest.PeerClientMock{
			ConnectToPeer_Return_array: []peerconnectorcontract.Peer{&connectormanagertest.PeerMock{}},
		}),
		connectormanagertest.WithPeerRepo(peerRepo),
	)

	tests.ContextScope(func(ctx context.Context) {
		ctx = context.WithValue(ctx, "test", "this is the expected context")
		goutil.PanicUnhandledError(connector.Start(ctx))
	})

	assert.Equal(t, "this is the expected context", peerRepo.Save_ctx.Value("test"))
}
