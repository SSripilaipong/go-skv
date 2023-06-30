package clientsidepeer

import (
	"context"
	"fmt"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func NewFactory() clientsidepeercontract.Factory {
	return tempFactory{}
}

type tempFactory struct{}

func (t tempFactory) New(ctx context.Context) (peerconnectorcontract.Peer, error) {
	return tempInteractor{}, nil
}

type tempInteractor struct{}

func (tempInteractor) UpdateReplica(key string, value string) error {
	fmt.Printf("log: client receive replica (%s, %s)\n", key, value)
	return nil
}
