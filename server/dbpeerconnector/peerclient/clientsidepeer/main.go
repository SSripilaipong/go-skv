package clientsidepeer

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"sync"
	"time"
)

func NewFactory(bufferSize int, defaultSendingTimeout time.Duration, replicaUpdaterFactory replicaupdatercontract.ActorFactory) clientsidepeercontract.Factory {
	return factory{
		replicaUpdaterFactory: replicaUpdaterFactory,
		bufferSize:            bufferSize,
		defaultSendingTimeout: defaultSendingTimeout,
	}
}

type factory struct {
	replicaUpdaterFactory replicaupdatercontract.ActorFactory
	bufferSize            int
	defaultSendingTimeout time.Duration
}

func (f factory) New(ctx context.Context) (peerconnectorcontract.Peer, error) {
	wg := sync.WaitGroup{}
	ch := make(chan command, f.bufferSize)
	wg.Add(1)
	go mainLoop(ctx, ch, func() { wg.Done() })
	return interactor{
		defaultTimeout:        f.defaultSendingTimeout,
		replicaUpdaterFactory: f.replicaUpdaterFactory,
		ch:                    ch,
		wg:                    &wg,
	}, nil
}

type interactor struct {
	replicaUpdaterFactory replicaupdatercontract.ActorFactory
	ch                    chan command
	wg                    *sync.WaitGroup
	defaultTimeout        time.Duration
}

func (t interactor) sendCommand(cmd command) error {
	select {
	case t.ch <- cmd:
	case <-time.After(t.defaultTimeout):
		return nil
	}
	return nil
}
