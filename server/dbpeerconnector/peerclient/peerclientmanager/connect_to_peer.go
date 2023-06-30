package peerclientmanager

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/util/goutil"
)

func (c client) ConnectToPeer(ctx context.Context, address string) (peerconnectorcontract.Peer, error) {
	peer, err := c.peerFactory.New(ctx)
	goutil.PanicUnhandledError(err)

	gateway, err := c.gatewayConnector.ConnectTo(ctx, address, peer)
	goutil.PanicUnhandledError(err)

	gateway.SubscribeReplica(ctx)

	return nil, nil
}
