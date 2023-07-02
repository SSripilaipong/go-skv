package peergrpcgateway

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type gateway struct {
	advertisedAddress string
	service           peergrpc.PeerServiceClient
	peer              peerconnectorcontract.Peer
}

func (g gateway) SubscribeReplica(ctx context.Context, onStopped func()) error {
	stream, err := g.service.SubscribeReplica(ctx, &peergrpc.SubscribeReplicaRequest{AdvertisedAddress: g.advertisedAddress})
	goutil.PanicUnhandledError(err)

	go func() {
		defer onStopped()

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
			stt, _ := status.FromError(err)
			if stt.Code() == codes.Canceled {
				return
			}
			goutil.PanicUnhandledError(err)

			goutil.PanicUnhandledError(g.peer.UpdateReplica(update.Key, update.Value))
		}
	}()

	return nil
}
