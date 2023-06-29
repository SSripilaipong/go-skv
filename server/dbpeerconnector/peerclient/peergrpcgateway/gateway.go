package peergrpcgateway

import (
	"context"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
)

type gateway struct {
	advertisedAddress string
	service           peergrpc.PeerServiceClient
}

func (g gateway) SubscribeReplica(ctx context.Context) error {
	_, err := g.service.SubscribeReplica(ctx, &peergrpc.SubscribeReplicaRequest{AdvertisedAddress: g.advertisedAddress})
	goutil.PanicUnhandledError(err)

	return nil
}
