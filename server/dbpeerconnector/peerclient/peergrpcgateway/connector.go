package peergrpcgateway

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type connector struct {
	advertisedAddress string
}

func (f connector) ConnectTo(ctx context.Context, address string) (peergrpcgatewaycontract.Gateway, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

	service := peergrpc.NewPeerServiceClient(conn)
	go func() {
		select {
		case <-ctx.Done():
			goutil.PanicUnhandledError(conn.Close())
		}
	}()
	return gateway{
		advertisedAddress: f.advertisedAddress,
		service:           service,
	}, nil
}

var _ peergrpcgatewaycontract.GatewayConnector = connector{}
