package peergrpccontroller

import (
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

func (g grpcImplementation) SubscribeReplica(req *peergrpc.SubscribeReplicaRequest, steam peergrpc.PeerService_SubscribeReplicaServer) error {
	goutil.PanicUnhandledError(g.usecase.SubscribeReplica(req.AdvertisedAddress, nil))
	return nil
}
