package clientsidepeer

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"sync"
)

func NewFactory(replicaUpdaterFactory replicaupdatercontract.Factory) clientsidepeercontract.Factory {
	return factory{
		replicaUpdaterFactory: replicaUpdaterFactory,
		bufferSize:            1,
	}
}

type factory struct {
	replicaUpdaterFactory replicaupdatercontract.Factory
	bufferSize            int
}

func (f factory) New(ctx context.Context) (peerconnectorcontract.Peer, error) {
	wg := sync.WaitGroup{}
	ch := make(chan command, f.bufferSize)
	wg.Add(1)
	go mainLoop(ctx, ch, func() { wg.Done() })
	return interactor{
		replicaUpdaterFactory: f.replicaUpdaterFactory,
		ch:                    ch,
		wg:                    &wg,
	}, nil
}

type interactor struct {
	replicaUpdaterFactory replicaupdatercontract.Factory
	ch                    chan command
	wg                    *sync.WaitGroup
}
