package subscribe_replica

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway"
	"go-skv/server/dbpeerconnector/peerserver/peergrpccontroller"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/peergrpc/peergrpctest"
	"go-skv/util/goutil"
	"go-skv/util/grpcutil"
	"testing"
)

func Test_should_call_usecase_on_server_with_advertised_usecase(t *testing.T) {
	usecase := &peergrpctest.ServerUsecaseMock{}
	controller := peergrpccontroller.New(0, usecase)
	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpcgateway.NewConnector("1.2.3.4:1234")
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()))
			goutil.PanicUnhandledError(err)

			goutil.PanicUnhandledError(gateway.SubscribeReplica(ctx))
		})
	})

	controller.Join()
	assert.Equal(t, "1.2.3.4:1234", usecase.SubscribeReplica_address)
}
