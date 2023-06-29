package peergrpctest

import (
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
)

type ServerUsecaseMock struct {
	SubscribeReplica_address string
}

func (u *ServerUsecaseMock) SubscribeReplica(address string, ch chan<- peerconnectorcontract.ReplicaUpdate) error {
	u.SubscribeReplica_address = address
	return nil
}

var _ peerserverusecase.Usecase = &ServerUsecaseMock{}
