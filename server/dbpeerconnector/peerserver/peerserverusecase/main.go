package peerserverusecase

import (
	"fmt"
	"go-skv/common/util/goutil"
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
	token := goutil.RandomString(8)
	fmt.Printf("PeerConnector: connected by Peer(%#v)\n", token) // TODO: remove demo log

	var i int
	for {
		key := goutil.RandomString(4)
		value := goutil.RandomString(4)
		ch <- peerconnectorcontract.ReplicaUpdate{Key: key, Value: value}
		i += 1

		fmt.Printf("PeerConnector: send ReplicaUpdate(%#v, %#v) to Peer(%#v)\n", key, value, token) // TODO: remove demo log
		time.Sleep(10 * time.Second)
	}
}
