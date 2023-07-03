package peergrpccontroller

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newGrpcImplementation(usecase peerserverusecase.Usecase) peergrpc.PeerServiceServer {
	return grpcImplementation{
		usecase: usecase,
	}
}

type grpcImplementation struct {
	peergrpc.UnimplementedPeerServiceServer
	usecase peerserverusecase.Usecase
}

func (g grpcImplementation) HealthCheck(context.Context, *peergrpc.Ping) (*peergrpc.Pong, error) {
	return &peergrpc.Pong{}, nil
}

func (g grpcImplementation) SubscribeReplica(req *peergrpc.SubscribeReplicaRequest, stream peergrpc.PeerService_SubscribeReplicaServer) error {
	ch := make(chan peerconnectorcontract.ReplicaUpdate, 1)
	go goutil.WillPanicUnhandledError(func() error { return g.usecase.SubscribeReplica(req.AdvertisedAddress, ch) })()
	for update := range ch {
		err := stream.Send(&peergrpc.ReplicaUpdate{
			Key:   update.Key,
			Value: update.Value,
		})
		stt, _ := status.FromError(err)
		if stt.Code() == codes.Unavailable {
			return nil
		}
		goutil.PanicUnhandledError(err)
	}
	return nil
}
