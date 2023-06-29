package peerserverusecase

import (
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type Usecase interface {
	SubscribeReplica(address string, ch chan<- peerconnectorcontract.ReplicaUpdate) error
}
