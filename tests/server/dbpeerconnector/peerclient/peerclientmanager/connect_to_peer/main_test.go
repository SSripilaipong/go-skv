package connect_to_peer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerclient/peerclientmanager"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/peerclient/peerclientmanager/peerclientmanagertest"
	"go-skv/util/goutil"
	"testing"
)

func Test_should_create_new_client_side_peer_with_context(t *testing.T) {
	peerFactory := &peerclientmanagertest.PeerFactoryMock{}
	manager := peerclientmanager.New(peerFactory)
	tests.ContextScope(func(ctx context.Context) {
		ctx = context.WithValue(ctx, "test", "1234567890")

		_, err := manager.ConnectToPeer(ctx, "1.1.1.1:1234")
		goutil.PanicUnhandledError(err)
	})

	assert.Equal(t, "1234567890", peerFactory.New_ctx.Value("test"))
}
