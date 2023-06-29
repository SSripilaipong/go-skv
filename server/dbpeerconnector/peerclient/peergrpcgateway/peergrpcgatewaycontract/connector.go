package peergrpcgatewaycontract

import "context"

type GatewayConnector interface {
	ConnectTo(ctx context.Context, address string) (Gateway, error)
}
