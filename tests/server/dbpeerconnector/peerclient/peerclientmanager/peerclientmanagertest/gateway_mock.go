package peerclientmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
)

type GatewayMock struct {
	SubscribeReplica_ctx context.Context
}

func (g *GatewayMock) SubscribeReplica(ctx context.Context, onStopped func()) error {
	g.SubscribeReplica_ctx = ctx
	return nil
}

var _ peergrpcgatewaycontract.Gateway = &GatewayMock{}
