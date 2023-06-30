package peergrpcgateway

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type connector struct {
	advertisedAddress string
}

func (f connector) ConnectTo(ctx context.Context, address string, peer peerconnectorcontract.Peer) (peergrpcgatewaycontract.Gateway, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

	service := peergrpc.NewPeerServiceClient(conn)
	go func() {
		select {
		case <-ctx.Done():
			goutil.PanicUnhandledError(conn.Close())
		}
	}()
	if pong, _ := service.HealthCheck(ctx, &peergrpc.Ping{}); pong == nil {
		return nil, peerconnectorcontract.CannotConnectToPeerError{}
	}

	return gateway{
		advertisedAddress: f.advertisedAddress,

		service: service,
		peer:    peer,
	}, nil
}

var _ peergrpcgatewaycontract.GatewayConnector = connector{}
