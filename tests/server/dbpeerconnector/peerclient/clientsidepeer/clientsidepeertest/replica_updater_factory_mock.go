package clientsidepeertest

import (
	"context"
	"go-skv/common/util/goutil"
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

	result := goutil.Coalesce[replicaupdatercontract.InboundUpdater](r.NewInboundUpdater_Return, &ReplicaInboundUpdaterMock{})
	return result, nil
}

var _ replicaupdatercontract.Factory = &ReplicaUpdaterFactoryMock{}
