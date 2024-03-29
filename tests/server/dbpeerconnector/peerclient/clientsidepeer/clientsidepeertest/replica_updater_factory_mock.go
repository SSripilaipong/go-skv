package clientsidepeertest

import (
	"context"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

type ReplicaUpdaterFactoryMock struct {
	NewInboundUpdater_IsCalled  bool
	NewInboundUpdater_CallCount int
	NewInboundUpdater_Return    replicaupdatercontract.InboundUpdater
	NewInboundUpdater_ctx       context.Context
}

func (r *ReplicaUpdaterFactoryMock) NewInboundUpdater(ctx context.Context) (replicaupdatercontract.InboundUpdater, error) {
	r.NewInboundUpdater_IsCalled = true
	r.NewInboundUpdater_ctx = ctx
	r.NewInboundUpdater_CallCount += 1

	return r.NewInboundUpdater_Return, nil
}

var _ replicaupdatercontract.Factory = &ReplicaUpdaterFactoryMock{}

type ReplicaUpdaterFactory2Mock struct {
	NewInboundUpdater_IsCalled  bool
	NewInboundUpdater_CallCount int
	NewInboundUpdater_Return    chan<- any
	NewInboundUpdater_ctx       context.Context
}

func (r *ReplicaUpdaterFactory2Mock) NewInboundUpdater(ctx context.Context) (chan<- any, error) {
	r.NewInboundUpdater_IsCalled = true
	r.NewInboundUpdater_ctx = ctx
	r.NewInboundUpdater_CallCount += 1

	return r.NewInboundUpdater_Return, nil
}

var _ replicaupdatercontract.ActorFactory = &ReplicaUpdaterFactory2Mock{}
