package subscribe_replica

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/tests/server/dbpeerconnector/peergrpc/peergrpctest"
	"go-skv/util/goutil"
	"go-skv/util/grpcutil"
	"testing"
)

func Test_should_call_usecase_on_server_with_advertised_usecase(t *testing.T) {
	usecase := &peergrpctest.ServerUsecaseMock{}
	controller := peergrpctest.NewController(
		peergrpctest.WithServerUsecase(usecase),
	)
	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpctest.NewConnector(
				peergrpctest.WithAdvertisedAddress("1.2.3.4:1234"),
			)
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()), nil)
			goutil.PanicUnhandledError(err)

			goutil.PanicUnhandledError(gateway.SubscribeReplica(ctx))
		})
	})

	controller.Join()
	assert.Equal(t, "1.2.3.4:1234", usecase.SubscribeReplica_address)
}

func Test_should_send_a_replica_update_back_to_peer_on_client_side_with_key_and_value(t *testing.T) {
	usecase := &peergrpctest.ServerUsecaseMock{SubscribeReplica_ch_Do: func(ch chan<- peerconnectorcontract.ReplicaUpdate) {
		ch <- peerconnectorcontract.ReplicaUpdate{Key: "aaa", Value: "bbb"}
	}}
	peer := &dbpeerconnectortest.PeerMock{}
	controller := peergrpctest.NewController(
		peergrpctest.WithServerUsecase(usecase),
	)
	tests.ContextScope(func(ctx context.Context) {
		goutil.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpctest.NewConnector()
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()), peer)
			goutil.PanicUnhandledError(err)

			goutil.PanicUnhandledError(gateway.SubscribeReplica(ctx))
		})
	})

	controller.Join()
	assert.Equal(t, "aaa", peer.UpdateReplica_key)
	assert.Equal(t, "bbb", peer.UpdateReplica_value)
}
