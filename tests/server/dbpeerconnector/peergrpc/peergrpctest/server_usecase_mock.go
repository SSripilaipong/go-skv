package peergrpctest

import (
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
)

type ServerUsecaseMock struct {
	SubscribeReplica_ch_Do   func(ch chan<- peerconnectorcontract.ReplicaUpdate)
	SubscribeReplica_address string
}

func (u *ServerUsecaseMock) SubscribeReplica(address string, ch chan<- peerconnectorcontract.ReplicaUpdate) error {
	u.SubscribeReplica_address = address
	if u.SubscribeReplica_ch_Do != nil {
		u.SubscribeReplica_ch_Do(ch)
	}
	close(ch)
	return nil
}

var _ peerserverusecase.Usecase = &ServerUsecaseMock{}
