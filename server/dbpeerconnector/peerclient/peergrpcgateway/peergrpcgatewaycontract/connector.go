package peergrpcgatewaycontract

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type GatewayConnector interface {
	ConnectTo(ctx context.Context, address string, peer peerconnectorcontract.Peer) (Gateway, error)
}
