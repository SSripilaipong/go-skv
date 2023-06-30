package clientsidepeer

import (
	"fmt"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New() peerconnectorcontract.Peer {
	return tempInteractor{}
}

type tempInteractor struct{}

func (tempInteractor) UpdateReplica(key string, value string) error {
	fmt.Printf("log: client receive replica (%s, %s)\n", key, value)
	return nil
}
