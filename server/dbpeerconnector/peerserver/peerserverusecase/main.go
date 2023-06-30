package peerserverusecase

import (
	"fmt"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerserver/peerserverusecase/peerserverusecase"
	"time"
)

func New() peerserverusecase.Usecase {
	return tempUsecase{}
}

type tempUsecase struct {
}

func (t tempUsecase) SubscribeReplica(address string, ch chan<- peerconnectorcontract.ReplicaUpdate) error {
	var i int
	for {
		ch <- peerconnectorcontract.ReplicaUpdate{
			Key:   fmt.Sprintf("key%d", i),
			Value: fmt.Sprintf("value%d", i),
		}
		i += 1

		time.Sleep(1 * time.Second)
	}
}
