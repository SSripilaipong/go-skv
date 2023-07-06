package clientsidepeertest

import "go-skv/server/replicaupdater/replicaupdatercontract"

type ReplicaUpdaterFactoryMock struct {
	NewInboundUpdater_IsCalled  bool
	NewInboundUpdater_CallCount int
	NewInboundUpdater_Return    replicaupdatercontract.InboundUpdater
}

func (r *ReplicaUpdaterFactoryMock) NewInboundUpdater() (replicaupdatercontract.InboundUpdater, error) {
	r.NewInboundUpdater_IsCalled = true
	r.NewInboundUpdater_CallCount += 1
	return r.NewInboundUpdater_Return, nil
}

var _ replicaupdatercontract.Factory = &ReplicaUpdaterFactoryMock{}
