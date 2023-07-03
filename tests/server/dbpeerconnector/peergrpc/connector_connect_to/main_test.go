package connector_connect_to

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/common/util/grpcutil"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/tests/server/dbpeerconnector/peergrpc/peergrpctest"
	"testing"
)

func Test_should_return_error_for_connecting_unavailable_server(t *testing.T) {
	gatewayConnector := peergrpctest.NewConnector()
	var err error
	tests.ContextScope(func(ctx context.Context) {
		_, err = gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(12345), &dbpeerconnectortest.PeerMock{})
	})
	assert.Equal(t, peerclientcontract.ConnectionError{}, err)
}
