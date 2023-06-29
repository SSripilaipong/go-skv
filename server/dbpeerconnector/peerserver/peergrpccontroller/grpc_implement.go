package peergrpccontroller

import (
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
	"go-skv/util/goutil"
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

func (g grpcImplementation) SubscribeReplica(req *peergrpc.SubscribeReplicaRequest, stream peergrpc.PeerService_SubscribeReplicaServer) error {
	ch := make(chan peerconnectorcontract.ReplicaUpdate, 1)
	go goutil.WillPanicUnhandledError(func() error { return g.usecase.SubscribeReplica(req.AdvertisedAddress, ch) })()
	for update := range ch {
		err := stream.Send(&peergrpc.ReplicaUpdate{
			Key:   update.Key,
			Value: update.Value,
		})
		goutil.PanicUnhandledError(err)
	}
	return nil
}
