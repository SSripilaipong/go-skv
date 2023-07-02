package peergrpcgatewaycontract

import "context"

type Gateway interface {
	SubscribeReplica(ctx context.Context, onStopped func()) error
}
