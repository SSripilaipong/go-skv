package subscribe_replica

import (
	"context"
	"github.com/stretchr/testify/assert"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/common/util/grpcutil"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/tests"
	"go-skv/tests/server/dbpeerconnector/dbpeerconnectortest"
	"go-skv/tests/server/dbpeerconnector/peergrpc/peergrpctest"
	"testing"
)

func Test_should_call_usecase_on_server_with_advertised_usecase(t *testing.T) {
	usecase := &peergrpctest.ServerUsecaseMock{}
	controller := peergrpctest.NewController(
		peergrpctest.WithServerUsecase(usecase),
	)
	tests.ContextScope(func(ctx context.Context) {
		goutil2.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpctest.NewConnector(
				peergrpctest.WithAdvertisedAddress("1.2.3.4:1234"),
			)
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()), nil)
			goutil2.PanicUnhandledError(err)

			gatewayStopped := make(chan struct{})
			goutil2.PanicUnhandledError(gateway.SubscribeReplica(ctx, func() { gatewayStopped <- struct{}{} }))
			goutil2.ReceiveWithTimeoutOrPanic(gatewayStopped, defaultTimeout)
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
		goutil2.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpctest.NewConnector()
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()), peer)
			goutil2.PanicUnhandledError(err)

			gatewayStopped := make(chan struct{})
			goutil2.PanicUnhandledError(gateway.SubscribeReplica(ctx, func() { gatewayStopped <- struct{}{} }))
			goutil2.ReceiveWithTimeout(gatewayStopped, defaultTimeout)
		})
	})

	controller.Join()
	assert.Equal(t, "aaa", peer.UpdateReplica_key)
	assert.Equal(t, "bbb", peer.UpdateReplica_value)
}

func Test_should_wait_for_update_in_background(t *testing.T) {
	triggerSendUpdate := make(chan struct{})
	var waitInBackground bool
	usecase := &peergrpctest.ServerUsecaseMock{SubscribeReplica_ch_Do: func(ch chan<- peerconnectorcontract.ReplicaUpdate) {
		_, waitInBackground = goutil2.ReceiveWithTimeout(triggerSendUpdate, defaultTimeout)
		ch <- peerconnectorcontract.ReplicaUpdate{}
	}}
	peer := &dbpeerconnectortest.PeerMock{}
	controller := peergrpctest.NewController(
		peergrpctest.WithServerUsecase(usecase),
	)
	tests.ContextScope(func(ctx context.Context) {
		goutil2.PanicUnhandledError(controller.Start(ctx))
		tests.SubContextScope(ctx, func(ctx context.Context) {
			gatewayConnector := peergrpctest.NewConnector()
			gateway, err := gatewayConnector.ConnectTo(ctx, grpcutil.LocalAddress(controller.Port()), peer)
			goutil2.PanicUnhandledError(err)

			goutil2.PanicUnhandledError(gateway.SubscribeReplica(ctx, func() {}))

			triggerSendUpdate <- struct{}{}
		})
	})

	assert.True(t, waitInBackground)
}
