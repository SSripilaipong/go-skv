package peerclientmanagertest

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type GatewayConnectorMock struct {
	ConnectTo_address string
	ConnectTo_ctx     context.Context
}

func (g *GatewayConnectorMock) ConnectTo(ctx context.Context, address string, peer peerconnectorcontract.Peer) (peergrpcgatewaycontract.Gateway, error) {
	g.ConnectTo_ctx = ctx
	g.ConnectTo_address = address
	return nil, nil
}
