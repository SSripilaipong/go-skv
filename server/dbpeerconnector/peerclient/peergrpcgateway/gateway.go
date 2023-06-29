package peergrpcgateway

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
	"io"
)

type gateway struct {
	advertisedAddress string
	service           peergrpc.PeerServiceClient
	peer              peerconnectorcontract.Peer
}

func (g gateway) SubscribeReplica(ctx context.Context) error {
	stream, err := g.service.SubscribeReplica(ctx, &peergrpc.SubscribeReplicaRequest{AdvertisedAddress: g.advertisedAddress})
	goutil.PanicUnhandledError(err)

	for {
		select {
		case <-ctx.Done():
			break
		default:
		}
		update, err := stream.Recv()
		if err == io.EOF {
			break
		}
		goutil.PanicUnhandledError(err)

		goutil.PanicUnhandledError(g.peer.UpdateReplica(update.Key, update.Value))
	}

	return nil
}
